package main

import "sort"

func main() {

}

type DSU struct {
	parent []int
	size   []int
}

func NewDSU(n int) *DSU {
	dsu := &DSU{
		parent: make([]int, n),
		size:   make([]int, n),
	}

	for i := 0; i < n; i++ {
		dsu.parent[i] = i
		dsu.size[i] = 1
	}

	return dsu
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) Union(x, y int) {
	rootX := d.Find(x)
	rootY := d.Find(y)

	if rootX == rootY {
		return
	}

	if d.size[rootX] < d.size[rootY] {
		rootX, rootY = rootY, rootX
	}

	d.parent[rootY] = rootX
	d.size[rootX] += d.size[rootY]
}

func accountsMerge(accounts [][]string) [][]string {
	dsu := NewDSU(len(accounts))
	emailToIdx := map[string]int{}

	for i, a := range accounts {
		for j := 1; j < len(a); j++ {
			if aidx, ok := emailToIdx[a[j]]; ok {
				dsu.Union(aidx, i)
			} else {
				emailToIdx[a[j]] = i
			}
		}
	}

	rootToEmail := map[int]map[string]bool{}
	roots := make([]int, 0, len(accounts))
	for i := range accounts {
		roots = append(roots, dsu.Find(i))
	}
	for i, a := range accounts {
		for j := 1; j < len(a); j++ {
			m := rootToEmail[roots[i]]
			if m == nil {
				m = map[string]bool{}
				rootToEmail[roots[i]] = m
			}
			m[a[j]] = true
		}
	}

	merged := [][]string{}
	for i, emailsMap := range rootToEmail {
		account := []string{accounts[i][0]}
		emails := make([]string, 0, len(emailsMap))
		for k := range emailsMap {
			emails = append(emails, k)
		}
		sort.Strings(emails)
		account = append(account, emails...)
		merged = append(merged, account)
	}

	return merged
}

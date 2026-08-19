package main

func main() {

}

type TimeValue struct {
	value string
	timestamp int
}

type TimeMap struct {
	valMap map[string][]TimeValue
}

func Constructor() TimeMap {
    return TimeMap{
		valMap: map[string][]TimeValue{},
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int)  {
    if _, ok := this.valMap[key]; !ok {
		this.valMap[key] = []TimeValue{}
	}
	this.valMap[key] = append(this.valMap[key], TimeValue{value, timestamp})
}


func (this *TimeMap) Get(key string, timestamp int) string {
	if len(this.valMap[key]) == 0 {
		return ""
	}

    left := 0
	right := len(this.valMap[key])-1
	ans := -1
	for left <= right {
		mid := left + (right - left) / 2
		if this.valMap[key][mid].timestamp <= timestamp {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if ans >= 0 {
		return this.valMap[key][ans].value 
	}
	return ""
}
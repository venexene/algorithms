package main

import "fmt"

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	result := exist(board, word)
	fmt.Println(result)
}

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				if check(board, word, i, j, 0) {
					return true
				}
			}
		}
	}
	return false
}

func check(board [][]byte, word string, r, c, l int) bool {
	if l == len(word) {
		return true
	}

	if r >= len(board) || r < 0 || c >= len(board[0]) || c < 0 {
		return false
	}

	if board[r][c] != word[l] {
		return false
	}

	temp := board[r][c]
	board[r][c] = '#'

	found := check(board, word, r+1, c, l+1) ||
		check(board, word, r-1, c, l+1) ||
		check(board, word, r, c+1, l+1) ||
		check(board, word, r, c-1, l+1)

	board[r][c] = temp
	return found
}

package main

import "fmt"

func main() {
	image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	sr := 1
	sc := 1
	color := 2
	result := floodFill(image, sr, sc, color)
	fmt.Println(result)
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	if image == nil {
		return nil
	}
	if len(image) == 0 {
		return image
	}
	if image[sr][sc] == color {
		return image
	}

	w := len(image)
	h := len(image[0])
	orig := image[sr][sc]

	image[sr][sc] = color
	if sr + 1 < w && image[sr + 1][sc] == orig {
		image = floodFill(image, sr + 1, sc, color)
	}
	if sr - 1 >= 0 && image[sr - 1][sc] == orig {
		image = floodFill(image, sr - 1, sc, color)
	}
	if sc + 1 < h && image[sr][sc + 1] == orig {
		image = floodFill(image, sr, sc + 1, color)
	}
	if sc - 1 >= 0 && image[sr][sc - 1] == orig {
		image = floodFill(image, sr, sc - 1, color)
	}

	return image
}
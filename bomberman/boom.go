package main

import (
	"fmt"
)

type tBomb struct {
	tick int
	mine byte
}

func Tick(fld [][]tBomb) {
	for row := range fld {
		for cell := range fld[row] {
			if fld[row][cell].mine == 'O' {
				fld[row][cell].tick--
			}
		}
	}
}

func Fill(fld [][]tBomb) {
	for row := range fld {
		for cell := range fld[row] {
			if fld[row][cell].mine == '.' {
				fld[row][cell].mine = 'O'
				fld[row][cell].tick = 3
			}
		}
	}
}

func Boom(fld [][]tBomb) {
	H := len(fld)
	W := len(fld[0])
	for row := range fld {
		for cell := range fld[row] {
			if fld[row][cell].mine == 'O' {
				if fld[row][cell].tick == 0 {
					//	boom destroed self...
					fld[row][cell].mine = '.'
					// and nearst cells
					if cell < W-1 {
						if (fld[row][cell+1].mine == 'O') && (fld[row][cell+1].tick > 0) {
							fld[row][cell+1].mine = '.'
						}
					}
					if row < H-1 {
						if (fld[row+1][cell].mine == 'O') && (fld[row+1][cell].tick > 0) {
							fld[row+1][cell].mine = '.'
						}
					}
				} else {
					if cell < W-1 {
						if (fld[row][cell+1].mine == 'O') && (fld[row][cell+1].tick == 0) {
							fld[row][cell].mine = '.'
						}
					}
					if row < H-1 {
						if (fld[row+1][cell].mine == 'O') && (fld[row+1][cell].tick == 0) {
							fld[row][cell].mine = '.'
						}
					}
				}
			}
		}
	}
}

// Complete the bomberMan function below.
func bomberMan(n int, grid []string) (res []string) {
	var (
		W, H int
	)

	W = len(grid[0])
	H = len(grid)
	m_fld := make([][]tBomb, H)
	for y := 0; y < H; y++ {
		m_fld[y] = make([]tBomb, W)
		for x := 0; x < W; x++ {
			m_fld[y][x].mine = grid[y][x]
			if grid[y][x] == 'O' {
				m_fld[y][x].tick = 3
			}
		}
	}

	if n > 6 {
		n = (n-2)%4 + 2
	}
	Tick(m_fld)
	n--
	for i := 0; i < n; i++ {
		Tick(m_fld)
		Fill(m_fld)
		Boom(m_fld)
	}

	res = make([]string, len(grid))
	for row := range m_fld {
		res[row] = ""
		for cell := range m_fld[row] {
			res[row] += string(m_fld[row][cell].mine)
		}
	}
	return res
}

func main() {
	grd := []string{
		"OOOOOOO",
		"O..O..O",
		"O.....O",
		"O..O..O",
		"OOOOOOO",
	}

	for i := 0; i < 15; i++ {
		fmt.Println(i-2, (i-2)%4)
		res := bomberMan(i, grd)
		for _, s := range res {
			fmt.Println(s)
		}
		fmt.Println()
	}
}

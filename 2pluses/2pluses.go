package main

import (
	"fmt"
)

type tPoint struct {
	x, y int
}

type tPlus struct {
	center tPoint
	size   int
}

//Area площадь плюса
func (p tPlus) Area() int {
	return p.size*2 - 1
}

//Across пересечение плюса с другим
func (p tPlus) Across(l tPlus) bool {
	var (
		dx, dy     int
		maxS, minS int
	)
	dx = p.center.x - l.center.x
	dy = p.center.y - l.center.y
	maxS = p.size
	if l.size > maxS {
		maxS = l.size
		minS = p.size
	} else {
		minS = l.size
	}

	if dx < 0 {
		dx = -dx
	}
	if dy < 0 {
		dy = -dy
	}
	maxS /= 2
	minS /= 2

	if dx == 0 {
		return dy <= (minS + maxS)

	} else if dx <= (minS) {
		return dy <= (maxS)

	} else if dx <= (maxS) {
		return dy <= (minS)
	}

	if dy == 0 {
		return dx <= (minS + maxS)
	}
	return false
}

func (p tPlus) String() string {
	return fmt.Sprintf("Plus%d(%d,%d)", p.size, p.center.x, p.center.y)
}

func allPluses(grid []string) (pls []tPlus) {
	var (
		l_pl tPlus
		s    int
		H, W int
	)

	H = len(grid)
	W = len(grid[0])
	// all plus maxumum sizes
	pls = make([]tPlus, 0, 2)
	fmt.Println("first plus iteration:")
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			if grid[y][x] != 'G' {
				continue
			}
			l_pl = tPlus{
				center: tPoint{x, y},
				size:   1,
			}
			s = 0
			for {
				s++
				if (x-s) < 0 || (x+s) >= W ||
					(y-s) < 0 || (y+s) >= H {
					break
				}
				if (grid[y][x-s] != 'G') || (grid[y][x+s] != 'G') ||
					(grid[y-s][x] != 'G') || (grid[y+s][x] != 'G') {
					break
				}
				l_pl.size += 2
			}

			pls = append(pls, l_pl)
			fmt.Println(l_pl)
		}
	}
	return
}

// Complete the twoPluses function below.
func twoPluses(grid []string) int {
	var (
		pl1, pl2 tPlus
		pls      []tPlus
		s, maxS  int
	)

	pls = allPluses(grid)
	lenPls := len(pls)
	// search two plus maximum multiple os area
	fmt.Println("search two pluses iteration")
	for x := 0; x < lenPls-1; x++ {
		pl1 = pls[x]
		for pl1.size > 0 {
			for y := x + 1; y < len(pls); y++ {
				pl2 = pls[y]
				for pl2.size > 0 {
					if pl1.Across(pl2) {
						pl2.size -= 2
						continue
					}
					s = pl1.Area() * pl2.Area()
					if s > maxS {
						maxS = s
						fmt.Println(pl1, pl2, maxS)
					}
					pl2.size -= 2
				}
			}
			pl1.size -= 2
		}
	}

	return maxS
}

func main() {
	grd := []string{
		"BBBGBGBBB",
		"BBBGBGBBB",
		"BBBGBGBBB",
		"GGGGGGGGG",
		"BBBGBGBBB",
		"BBBGBGBBB",
		"GGGGGGGGG",
		"BBBGBGBBB",
		"BBBGBGBBB",
		"BBBGBGBBB",
	}

	fmt.Println(twoPluses(grd))
}

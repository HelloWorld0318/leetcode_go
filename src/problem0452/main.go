package problem0452

import "sort"

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	pairs := make([]*Pair, 0)
	for _, point := range points {
		pairs = append(pairs, &Pair{
			begin: point[0],
			end:   point[1],
		})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].begin < pairs[j].begin
	})
	var shootNum, shootEnd = 1, pairs[0].end
	i := 1
	for i < len(pairs) {
		if pairs[i].begin <= shootEnd {
			if pairs[i].end < shootEnd {
				shootEnd = pairs[i].end
			}
		} else {
			shootNum++
			shootEnd = pairs[i].end
		}
		i++
	}
	return 0
}

type Pair struct {
	begin int
	end   int
}

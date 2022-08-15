package problem0200

var (
	dx = []int{-1, 1, 0, 0}
	dy = []int{0, 0, 1, -1}
)

//求地图中岛屿的数量
//1.设置岛屿数量numOfLands=0
//2.设置mark数组，并初始化
//3.遍历地图grid的上所有的点，如果当前点是陆地，且未被访问过（mark[i][j]=0）,调用搜索接口search(grid,mark,i,j)
//search可以是DFS或者BFS，完成搜索后numOfLands++
func numIslands(grid [][]byte) int {
	numOfLands := 0
	if len(grid) == 0 {
		return 0
	}
	row, column := len(grid), len(grid[0])
	mark := make([][]int, row)
	i := 0
	for i < row {
		mark[i] = make([]int, column)
		i++
	}
	i = 0
	for i < row {
		j := 0
		for j < column {
			if grid[i][j] == '1' && mark[i][j] == 0 {
				bfs(grid, mark, i, j)
				numOfLands++
			}
			j++
		}
		i++
	}
	return numOfLands
}

func dfs(grid [][]byte, mark [][]int, x int, y int) {
	mark[x][y] = 1
	//朝方向数组4个方向遍历深度优先搜索
	row, column := len(grid), len(grid[0])
	i := 0
	for i < 4 {
		newX := x + dx[i]
		newY := y + dy[i]
		if newX >= 0 && newX < row && newY >= 0 && newY < column && grid[newX][newY] == '1' &&
			mark[newX][newY] == 0 {
			dfs(grid, mark, newX, newY)
		}
		i++
	}
}

func bfs(grid [][]byte, mark [][]int, x int, y int) {
	mark[x][y] = 1
	queue := make([]*Pair, 0)
	pair := &Pair{X: x, Y: y}
	queue = append(queue, pair)
	for len(queue) != 0 {
		pair, queue = queue[0], queue[1:]
		i := 0
		for i < 4 {
			newX := pair.X + dx[i]
			newY := pair.Y + dy[i]
			for newX >= 0 && newX < len(grid) && newY >= 0 && newY < len(grid[0]) && grid[newX][newY] == '1' &&
				mark[newX][newY] == 0 {
				mark[newX][newY] = 1
				queue = append(queue, &Pair{X: newX, Y: newY})
			}
			i++
		}
	}
}

type Pair struct {
	X int
	Y int
}

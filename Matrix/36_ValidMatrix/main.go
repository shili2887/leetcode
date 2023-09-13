package main

func isValidSudoku(board [][]byte) bool {
	rows := [9][9]int{}
	columns := [9][9]int{}
	subMax := [3][3][9]int{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			numBy := board[i][j]
			if numBy == '.' {
				continue
			}
			num := numBy - '0'
			rows[i][num-1]++
			if rows[i][num-1] > 1 {
				return false
			}
			columns[j][num-1]++
			if columns[j][num-1] > 1 {
				return false
			}
			subMax[i/3][j/3][num-1]++
			if subMax[i/3][j/3][num-1] > 1 {
				return false
			}
		}
	}

	return true
}

func main() {

}

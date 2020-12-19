package main

var board = [3][3]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}

const (
	empty  = 0
	cross  = 1
	circle = 2
)

func markSquare(row, col, shape int) {
	board[row][col] = shape
}

func checkWin() (shape int) {
	if a := checkHorizontal(); a > 0 {
		return a
	} else if a := checkVertical(); a > 0 {
		return a
	} else if a := checkDiagonal(); a > 0 {
		return a
	}
	return 0
}

func checkHorizontal() (shape int) {
	for i := 0; i < 3; i++ {
		threeInARow := true
		temp := board[i][0]
		for j := 1; j < 3; j++ {
			if board[i][j] != temp || board[i][j] == 0 {
				threeInARow = false
				j = 3
			}
		}
		if threeInARow {
			return temp
		}
	}
	return 0
}

func checkVertical() (shape int) {
	for j := 0; j < 3; j++ {
		threeInARow := true
		temp := board[0][j]
		for i := 1; i < 3; i++ {
			if board[i][j] != temp || board[i][j] == 0 {
				threeInARow = false
				i = 3
			}
		}
		if threeInARow {
			return temp
		}
	}
	return 0
}

func checkDiagonal() (shape int) {
	//checks if all spaces match (nonzero) left to right
	threeInARow := true
	temp := board[0][0]
	j := 1
	for i := 1; i < 3; i++ {
		if board[i][j] != temp || board[i][j] == 0 {
			threeInARow = false
			i = 3
		}
		j++
	}
	if threeInARow {
		return temp
	}

	//checks if all spaces match (nonzero) left to right
	threeInARow = true
	temp = board[2][0]
	j = 1
	for i := 1; i >= 0; i-- {
		if board[i][j] != temp || board[i][j] == 0 {
			threeInARow = false
			i = -1
		}
		j++
	}
	if threeInARow {
		return temp
	}
	return 0
}

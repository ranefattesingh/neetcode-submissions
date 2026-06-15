func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		col := make(map[byte]struct{})
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			_, ok := col[board[i][j]]
			if ok {
				return false
			}

			col[board[i][j]] = struct{}{}
		}

		

		row := make(map[byte]struct{})
		for j := 0; j < 9; j++ {
			if board[j][i] == '.' {
				continue
			}

			_, ok := row[board[j][i]]
			if ok {
				return false
			}

			row[board[j][i]] = struct{}{}
		}
	}




	for i := 0; i < 9; i++ {
		row := ((i * 3) / 9) * 3
		col := (i * 3) % 9
		box := make(map[byte]struct{})
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				if board[row + j][col + k] == '.' {
					continue
				}

				_, ok := box[board[row + j][col + k]]
				if ok {
					return false
				}

				box[board[row + j][col + k]] = struct{}{}
			}
		}
	}

	return true
}


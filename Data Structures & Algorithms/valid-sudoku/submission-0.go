func isValidSudoku(board [][]byte) bool {
	

	// ROW VALIDATION
	for i := 0; i < 9; i++ {
		visited := make(map[byte]struct{})
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			if _, ok := visited[board[i][j]]; ok {
				return false
			}

			visited[board[i][j]] = struct{}{}
		}
	}

	// COLUMN VALIDATION
	for i := 0; i < 9; i++ {
		visited := make(map[byte]struct{})
		for j := 0; j < 9; j++ {
			if board[j][i] == '.' {
				continue
			}

			if _, ok := visited[board[j][i]]; ok {
				return false
			}

			visited[board[j][i]] = struct{}{}
		}
	}

	// BOX VALIDATION
	for b := 0; b < 9; b++ {
		visited := make(map[byte]struct{})
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				row := (b / 3) * 3 + i
                col := (b % 3) * 3 + j
				if board[row][col] == '.' {
					continue
				}

				if _, ok := visited[board[row][col]]; ok {
					return false
				}

				visited[board[row][col]] = struct{}{}
			}
		}
	}

	return true
}

func searchMatrix(matrix [][]int, target int) bool {
    rows, cols := len(matrix), len(matrix[0])
    l, r := 0, rows*cols-1

    for l <= r {
        m := l+(r-l)/2
        row, col := m/cols, m%cols
        middleVal := matrix[row][col]

        if middleVal == target {
            return true
        }

        if middleVal < target {
            l = m + 1
        } else {
            r = m - 1
        }
    }

    return false
}

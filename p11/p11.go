package p11


// initial approach: scan the nxn array in 4x4 chunks.

func Solve(mat [][]int) int {

	result := 0
	for i := 0; i < len(mat) - 4; i++ {
		for j := 0; j < len(mat[i]) - 4; j++ {
			result = max(result, processChunk(mat, i, j))
		}
	}

	return result
}


// takes matrix, upper left index of 4x4
// return max product in 4x4 chunk

func processChunk(mat [][]int, r int, c int) int {

	maximum := 0

	for i := 0; i < 4; i++ {
		maximum = max(maximum, mat[r + i][c] * mat[r + i][c + 1] * mat[r + i][c + 2] * mat[r + i][c + 3])
		maximum = max(maximum, mat[r][c + i] * mat[r + 1][c + i] * mat[r + 2][c + i] * mat[r + 3][c + i])
	}

	maximum = max(maximum, mat[r][c] * mat[r + 1][c + 1] * mat[r + 2][c + 2] * mat[r + 3][c + 3])
	maximum = max(maximum, mat[r][c + 3] * mat[r + 1][c + 2] * mat[r + 2][c + 1] * mat[r + 3][c])

	return maximum
}

package Hot100

//48. 旋转图像
//给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
//
//你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。
//
//
//
//示例 1：
//
//
//输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[[7,4,1],[8,5,2],[9,6,3]]
//示例 2：
//
//
//输入：matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
//输出：[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
//示例 3：
//
//输入：matrix = [[1]]
//输出：[[1]]
//示例 4：
//
//输入：matrix = [[1,2],[3,4]]
//输出：[[3,1],[4,2]]
//
//
//提示：
//
//matrix.length == n
//matrix[i].length == n
//1 <= n <= 20
//-1000 <= matrix[i][j] <= 1000

// 基于圈进行旋转 + 递归
func rotate(matrix [][]int) [][]int {
	size := len(matrix)

	var rotateCircle func(left int, right int)
	rotateCircle = func(left int, right int) {
		if left > right {
			return
		}

		// 按圈进行旋转
		for i := left; i < right; i++ {
			matrix[left][i], matrix[i][right], matrix[right][right-i+left], matrix[right-i+left][left] =
				matrix[right-i+left][left], matrix[left][i], matrix[i][right], matrix[right][right-i+left]
		}

		rotateCircle(left+1, right-1)
	}

	rotateCircle(0, size-1)
	return matrix
}

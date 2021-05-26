package Hot100

//22. 括号生成
//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
//
//
//示例 1：
//
//输入：n = 3
//输出：["((()))","(()())","(())()","()(())","()()()"]
//示例 2：
//
//输入：n = 1
//输出：["()"]
//
//
//提示：
//
//1 <= n <= 8

// 动态规划求解 P(n) = '('+ P(i) + ')' + P(n-1-i)
func generateParenthesis(n int) []string {
	parenthesises := make([][]string, 0)
	parenthesises = append(parenthesises, []string{""})

	for i := 1; i <= n; i++ {
		parenthesis_i := []string{}
		for j := 0; j <= i-1; j++ {
			p2 := parenthesises[j]
			p1 := parenthesises[i-1-j]
			for _, tmp1 := range p1 {
				for _, tmp2 := range p2 {
					parenthesis_i = append(parenthesis_i, "("+tmp1+")"+tmp2)
				}
			}
		}
		parenthesises = append(parenthesises, parenthesis_i)
	}
	return parenthesises[n]

}

var parenthesis []string

// 根据“剩余左括号总数要小于等于右括号”的规则生成所有可能结果 + 递归实现
func generateParenthesis2(n int) []string {
	parenthesis = make([]string, 0)
	getParenthesis(n, n, "")
	return parenthesis
}

func getParenthesis(left int, right int, cur string) {
	if left == 0 && right == 0 {
		parenthesis = append(parenthesis, cur)
		return
	}

	if left > right {
		return
	} else {
		if left > 0 {
			getParenthesis(left-1, right, cur+"(")
		}
		if right > 0 {
			getParenthesis(left, right-1, cur+")")
		}
	}

}

// 暴力穷举判断是否合法（n=5超时）
var parenthesisSet map[string]bool

func generateParenthesis1(n int) []string {
	parenthesis = make([]string, 0)
	parenthesisSet = make(map[string]bool, 0)
	helper(0, n, n, "")
	return parenthesis
}

func helper(id int, index int, size int, cur string) {
	if index == 0 {
		for i := len(cur); i < 2*size; i++ {
			cur += ")"
		}
		if isParenthesis(cur) && !parenthesisSet[cur] {
			parenthesisSet[cur] = true
			parenthesis = append(parenthesis, cur)
		}
		return
	}
	for i := id; i < 2*size; i++ {
		helper(id+1, index-1, size, cur+"(")
		if index <= (2*size - id) {
			helper(id+1, index, size, cur+")")
		}
	}

}

func isParenthesis(str string) bool {
	left, right := 0, 0
	size := len(str)
	for i := 0; i < size; i++ {
		if str[i] == '(' {
			left++
		} else {
			right++
		}
		if left < right {
			return false
		}
	}
	for i := size - 1; i >= 0; i-- {
		if str[i] == '(' {
			left++
		} else {
			right++
		}
		if left > right {
			return false
		}
	}
	return left == right
}

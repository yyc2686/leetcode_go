package Hot100

//121. 买卖股票的最佳时机
//给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
//
//你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
//
//返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
//
//
//
//示例 1：
//
//输入：[7,1,5,3,6,4]
//输出：5
//解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
//注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
//示例 2：
//
//输入：prices = [7,6,4,3,1]
//输出：0
//解释：在这种情况下, 没有交易完成, 所以最大利润为 0。
//
//
//提示：
//
//1 <= prices.length <= 105
//0 <= prices[i] <= 104

// 维护当前最小值（优化空间，原数组保存当前最小价格）（动态规划）
func maxProfit(prices []int) int {
	ret := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			if ret < prices[i]-prices[i-1] {
				ret = prices[i] - prices[i-1]
			}
			prices[i] = prices[i-1]
		}
	}
	return ret
}

// 维护当前最小值（单调栈改进）
func maxProfit2(prices []int) int {
	ret := 0
	minPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if ret < prices[i]-minPrice {
			ret = prices[i] - minPrice
		}
	}
	return ret
}

// 暴力解法：二重循环（200 / 211 个通过测试用例）
func maxProfit1(prices []int) int {
	ret := 0
	size := len(prices)
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			if prices[j] > prices[i] && ret < prices[j]-prices[i] {
				ret = prices[j] - prices[i]
			}
		}
	}
	return ret
}

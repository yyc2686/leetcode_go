package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type ThreeSumInput struct {
	nums []int
}

type ThreeSumOutput struct {
	ret [][]int
}

func TestThreeSum(t *testing.T) {
	inputs := []ThreeSumInput{
		{[]int{-2, 0, 0, 2, 2}},
		{[]int{-3, 1, -5, -1, 0, -1, 3, -4, 1, 2, -1, -1, -4, -4}},
		{[]int{0, 0, 0, 0}},
		{[]int{-1, 0, 1}},
		{[]int{-1, 0, 1, 2, -1, -4}},
		{[]int{}},
		{[]int{0}},
	}

	expected := []ThreeSumOutput{
		{[][]int{{-2, 0, 2}}},
		{[][]int{{-5, 2, 3}, {-4, 1, 3}, {-3, 0, 3}, {-3, 1, 2}, {-1, -1, 2}, {-1, 0, 1}}},
		{[][]int{{0, 0, 0}}},
		{[][]int{{-1, 0, 1}}},
		{[][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[][]int{}},
		{[][]int{}},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := threeSum(input.nums)
		assert.Equal(t, expected[i].ret, ret)
	}

}

func _BenchmarkThreeSumOperator(b *testing.B) {

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		threeSumOperator([]int{-3, 1, -5, -1, 0, -1, 3, -4, 1, 2, -1, -1, -4, -4})
		//fmt.Println("Actual", threeSum([]int{-3, 1, -5, -1, 0, -1, 3, -4, 1, 2, -1, -1, -4, -4}))
		//fmt.Println("Expected", [][]int{{-5, 2, 3}, {-4, 1, 3}, {-3, 0, 3}, {-3, 1, 2}, {-1, -1, 2}, {-1, 0, 1}})
	}

	b.StopTimer()

}

func _BenchmarkThreeSumSprintf(b *testing.B) {

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		threeSumSprintf([]int{-3, 1, -5, -1, 0, -1, 3, -4, 1, 2, -1, -1, -4, -4})
		//fmt.Println("Actual", threeSum([]int{-3, 1, -5, -1, 0, -1, 3, -4, 1, 2, -1, -1, -4, -4}))
		//fmt.Println("Expected", [][]int{{-5, 2, 3}, {-4, 1, 3}, {-3, 0, 3}, {-3, 1, 2}, {-1, -1, 2}, {-1, 0, 1}})
	}

	b.StopTimer()

}

func _BenchmarkThreeSumJoin(b *testing.B) {

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		threeSumJoin([]int{-3, 1, -5, -1, 0, -1, 3, -4, 1, 2, -1, -1, -4, -4})
		//fmt.Println("Actual", threeSum([]int{-3, 1, -5, -1, 0, -1, 3, -4, 1, 2, -1, -1, -4, -4}))
		//fmt.Println("Expected", [][]int{{-5, 2, 3}, {-4, 1, 3}, {-3, 0, 3}, {-3, 1, 2}, {-1, -1, 2}, {-1, 0, 1}})
	}

	b.StopTimer()

}

func _BenchmarkThreeSumBuffer(b *testing.B) {

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		threeSumBuffer([]int{-3, 1, -5, -1, 0, -1, 3, -4, 1, 2, -1, -1, -4, -4})
		//fmt.Println("Actual", threeSum([]int{-3, 1, -5, -1, 0, -1, 3, -4, 1, 2, -1, -1, -4, -4}))
		//fmt.Println("Expected", [][]int{{-5, 2, 3}, {-4, 1, 3}, {-3, 0, 3}, {-3, 1, 2}, {-1, -1, 2}, {-1, 0, 1}})
	}

	b.StopTimer()

}

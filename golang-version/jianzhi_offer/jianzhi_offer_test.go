package jianzhi_offer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRepeatedNumInArray(t *testing.T) {
	var testCase = []struct {
		input  []int
		result int
	}{
		{[]int{2, 3, 1, 0, 2, 5, 3}, 2},
		{[]int{0, 1, 2, 3, 4}, -1}, //no repeat data case.
		{[]int{1, 0, 3, 2, 5, 4}, -1},
	}
	for i := 0; i < len(testCase); i++ {
		ret := FindRepeatedNumInArray(testCase[i].input)
		assert.Equal(t, ret, testCase[i].result)
	}

}

func TestFindValueInMatrix(t *testing.T) {
	var data [][]int = [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("row nums: ", len(data), ", col nums: ", len(data[0]))

	var testCase = []struct {
		input  [][]int
		retRow int
		retCol int
		dst    int
	}{
		{[][]int{
			{1, 2, 8, 9},
			{2, 4, 9, 12},
			{4, 7, 10, 13},
			{6, 8, 11, 15},
		},
			2,
			1,
			7,
		},
		{[][]int{
			{1, 2, 8, 9},
			{2, 4, 9, 12},
			{4, 7, 10, 13},
			{6, 8, 11, 15},
		},
			-1,
			-1,
			5,
		},
	}

	for i := 0; i < len(testCase); i++ {
		rRow, rCol := FindValueInMatrix(testCase[i].input, testCase[i].dst)
		assert.Equal(t, rRow, testCase[i].retRow)
		assert.Equal(t, rCol, testCase[i].retCol)
	}
}

func TestReplaceSpaceInStr(t *testing.T) {
	var testCase = []struct {
		input string
		ret   string
	}{
		{"hello world!", "hello%20world!"},
		{"  hello  world!", "%20%20hello%20%20world!"},
	}
	for i := 0; i < len(testCase); i++ {
		ret := ReplaceSpaceInStr(testCase[i].input)
		assert.Equal(t, ret, testCase[i].ret)
	}
}

func TestReverseSingleLink(t *testing.T) {
	var testCase = []struct {
		input  []int
		result []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{nil, nil},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{2, 1}},
	}

	for i := 0; i < len(testCase); i++ {
		s1 := NewSingleLinkBySlice(testCase[i].input)
		assert.NotEqual(t, s1, nil)

		s2 := ReverseSingleLink(s1)
		assert.NotEqual(t, s2, nil)

		r2 := IterSingleLink(s2)
		assert.Equal(t, len(testCase[i].input), len(r2))
		//fmt.Println(testCase[i].result)
		//fmt.Println(r2)
		for j := 0; j < len(r2); j++ {
			assert.Equal(t, testCase[i].result[j], r2[j])
		}
	}
}

func TestQueueImplementInStack(t *testing.T) {
	var testCase = []struct {
		input  []int
		result []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}},
	}

	for i := 0; i < len(testCase); i++ {
		stackItem := NewQueueImplementInStack()
		for jj := 0; jj < len(testCase[i].input); jj++ {
			stackItem.AppendTail(testCase[i].input[jj])
		}
		//
		for jj := 0; jj < len(testCase[i].result); jj++ {
			v := stackItem.DeleteHead()
			assert.Equal(t, testCase[i].result[jj], v)
		}
	}
}

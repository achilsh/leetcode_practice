package jianzhi_offer

import (
	"fmt"
	"reflect"
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

func TestFibonacciInCycle(t *testing.T) {
	var testCase = []struct {
		input int
		ret   int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
	}
	for i := 0; i < len(testCase); i++ {
		ret := FibonacciInCycle(testCase[i].input)
		assert.Equal(t, ret, testCase[i].ret)
	}
}

func TestQinWaJump(t *testing.T) {
	var testCase = []struct {
		input int
		ret   int
	}{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		{6, 13},
		{7, 21},
		{8, 34},
	}
	for i := 0; i < len(testCase); i++ {
		ret := QinWaJump(testCase[i].input)
		assert.Equal(t, testCase[i].ret, ret)
	}
}

func TestBinarySearch(t *testing.T) {
	var testCase = []struct {
		input []int
		dst   int
		ret   int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 4, 3},
		{[]int{1, 2, 3, 4, 5, 6}, 100, -1},
		{[]int{1, 2, 3, 4, 5, 6}, 6, 5},
		{[]int{1, 2, 3, 4, 5, 6}, 1, 0},
	}

	for i := 0; i < len(testCase); i++ {
		ret := BinarySearch(testCase[i].input, testCase[i].dst)
		assert.Equal(t, testCase[i].ret, ret)

	}
}

func TestBubbleSort(t *testing.T) {
	var testCase = []struct {
		input []int
		ret   []int
	}{
		{[]int{1, 6, 2, 8, 3, 10, 9, 7}, []int{1, 2, 3, 6, 7, 8, 9, 10}},
		{[]int{1, 2, 3, 4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{1}, []int{1}},
	}
	for i := 0; i < len(testCase); i++ {
		ret := BubbleSort(testCase[i].input)
		t.Logf("ret: %v", ret)
		assert.Equal(t, reflect.DeepEqual(ret, testCase[i].ret), true)
	}
}

func TestInsertSort(t *testing.T) {
	var testCase = []struct {
		input []int
		ret   []int
	}{
		{[]int{1, 6, 2, 8, 3, 10, 9, 7}, []int{1, 2, 3, 6, 7, 8, 9, 10}},
		{[]int{1, 2, 3, 4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{1}, []int{1}},
	}
	for i := 0; i < len(testCase); i++ {
		ret := InsertSort(testCase[i].input)
		t.Logf("ret: %v", ret)
		assert.Equal(t, reflect.DeepEqual(ret, testCase[i].ret), true)
	}
}

func TestChoseSort(t *testing.T) {
	var testCase = []struct {
		input []int
		ret   []int
	}{
		{[]int{1, 6, 2, 8, 3, 10, 9, 7}, []int{1, 2, 3, 6, 7, 8, 9, 10}},
		{[]int{1, 2, 3, 4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{1}, []int{1}},
	}
	for i := 0; i < len(testCase); i++ {
		ret := ChoseSort(testCase[i].input)
		t.Logf("ret: %v", ret)
		assert.Equal(t, reflect.DeepEqual(ret, testCase[i].ret), true)
	}
}

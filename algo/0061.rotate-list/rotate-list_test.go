package problem0061

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

type para struct {
	one []int
	two int
}

type ans struct {
	one []int
}

func Test_Problem0061(t *testing.T) {
	ast := assert.New(t)
	qs := []question{
		question{
			para{
				[]int{1, 2, 3, 4, 5},
				2,
			},
			ans{
				[]int{4, 5, 1, 2, 3},
			},
		},
		question{
			para{
				[]int{0, 1, 2},
				4,
			},
			ans{
				[]int{2, 0, 1},
			},
		},
		question{
			para{
				[]int{1},
				0,
			},
			ans{
				[]int{1},
			},
		},
		question{
			para{
				[]int{1, 2},
				1,
			},
			ans{
				[]int{2, 1},
			},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)
		ast.Equal(a.one, l2s(rotateRight(s2l(p.one), p.two)), "输入:%v", p)
	}

}

// convert *ListNode to []int
func l2s(head *ListNode) []int {
	res := []int{}

	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

// convert []int to *ListNode
func s2l(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	res := &ListNode{
		Val: nums[0],
	}
	temp := res
	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{
			Val: nums[i],
		}
		temp = temp.Next
	}

	return res
}

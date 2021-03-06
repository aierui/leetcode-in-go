package problem0450

import (
	"fmt"
	"testing"

	"github.com/aierui/leetcode-in-go/kit"
	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para1 []int
	para2 []int
	key   int
	ans1  []int
	ans2  []int
}{
	{
		[]int{2, 1, 3, 4},
		[]int{1, 2, 3, 4},
		4,
		[]int{2, 1, 3},
		[]int{1, 2, 3},
	},
}

func Test_DeleteNode(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		root := kit.PreIn2Tree(q.para1, q.para2)
		tree := kit.PreIn2Tree(q.ans1, q.ans2)

		ast.Equal(tree, deleteNode(root, q.key), "输入：%v\n", q)
	}
}

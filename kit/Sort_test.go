package kit

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para []int
	ans  []int
}{
	{
		[]int{2, 3, 32, 3, 4, 1, 7, 10},
		[]int{1, 2, 3, 3, 4, 7, 10, 32},
	},
	{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	},
	{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	},
}

func Test_Bubble(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		Bubble(q.para)
		ast.Equal(q.ans, q.para, "输入：%v\n", q)
	}
}

func Benchmark_Bubble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			Bubble(q.para)
		}
	}
}

func Test_Insert(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		Insert(q.para)
		ast.Equal(q.ans, q.para, "输入：%v\n", q)
	}
}

func Benchmark_Insert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			Insert(q.para)
		}
	}
}

func Test_Selection(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		Selection(q.para)
		ast.Equal(q.ans, q.para, "输入：%v\n", q)
	}
}

func Benchmark_Selection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			Selection(q.para)
		}
	}
}

func Test_Quick(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		Quick(q.para)
		ast.Equal(q.ans, q.para, "输入：%v\n", q)
	}
}

func Benchmark_Quick(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			Quick(q.para)
		}
	}
}

func Test_Merge(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		Merge(q.para)
		ast.Equal(q.ans, q.para, "输入：%v\n", q)
	}
}

func Benchmark_Merge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			Merge(q.para)
		}
	}
}

func Test_Heap(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		Heap(q.para)
		ast.Equal(q.ans, q.para, "输入：%v\n", q)
	}
}

func Benchmark_Heap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			Heap(q.para)
		}
	}
}

func Test_Shell(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		Shell(q.para)
		ast.Equal(q.ans, q.para, "输入：%v\n", q)
	}
}

func Benchmark_Shell(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			Shell(q.para)
		}
	}
}

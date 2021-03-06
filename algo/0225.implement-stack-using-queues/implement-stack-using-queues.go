package problem0225

type MyStack struct {
	Element []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	var myStack MyStack

	return myStack
}

// Push O(n)
/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.Element = append(this.Element, x)

	n := len(this.Element)

	for i := 0; i < n-1; i++ { // reverse element
		first := this.Element[0]
		this.Element = this.Element[1:]
		this.Element = append(this.Element, first)
	}

}

// Pop O(1)
/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	tmp := this.Element[0]
	this.Element = this.Element[1:]

	return tmp
}

// Top O(1)
/** Get the top element. */
func (this *MyStack) Top() int {
	return this.Element[0]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.Element) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

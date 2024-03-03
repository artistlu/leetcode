package min_stack01

type MinStack struct {
	stack  []int
	helper []int
}

func Constructor() MinStack {
	return MinStack{
		stack:  make([]int, 0),
		helper: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.helper) == 0 || this.GetMin() >= val {
		this.helper = append(this.helper, val)
	}
}

func (this *MinStack) Pop() {
	if this.Top() == this.GetMin() {
		this.helper = this.helper[:len(this.helper)-1]
	}

	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.helper[len(this.helper)-1]
}

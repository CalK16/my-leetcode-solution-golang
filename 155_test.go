package main

type MinStack struct {
	min  []int
	full []int
}

func Constructor() MinStack {
	return MinStack{
		min:  make([]int, 0),
		full: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.full = append(this.full, val)
	if len(this.min) == 0 || this.min[len(this.min)-1] >= val {
		this.min = append(this.min, val)
	}
}

func (this *MinStack) Pop() {
	last := this.full[len(this.full)-1]
	if len(this.min) > 0 && this.min[len(this.min)-1] == last {
		this.min = this.min[:len(this.min)-1]
	}
	this.full = this.full[:len(this.full)-1]
}

func (this *MinStack) Top() int {
	return this.full[len(this.full)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

package medium

type MinStack struct {
	stack    []int
	minStack []int
	minStart int
	start    int
}

func ConstructorMinStack() MinStack {
	return MinStack{
		stack:    make([]int, 0),
		minStack: make([]int, 0),
		minStart: -1,
		start:    -1,
	}
}

func (this *MinStack) Push(val int) {

	if len(this.stack) >= this.start+1 {
		this.stack = append(this.stack, val)
	} else {
		this.stack[this.start+1] = val
	}
	this.start++

	if (this.minStart >= 0 && this.minStack[this.minStart] >= val) || this.minStart == -1 {
		if len(this.minStack) >= this.minStart+1 {
			this.minStack = append(this.minStack, val)
		} else {
			this.minStack[this.minStart+1] = val
		}
		this.minStart++
	}

}

func (this *MinStack) Pop() {

	val := this.stack[this.start]
	if this.minStart >= 0 && this.minStack[this.minStart] == val {
		this.minStart--
	}
	this.start--
}

func (this *MinStack) Top() int {
	return this.stack[this.start]
}

func (this *MinStack) GetMin() int {
	return this.minStack[this.minStart]
}

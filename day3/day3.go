package day3

type CustomStack struct {
	maxSize  int
	stackLen int
	stack    []int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{maxSize, 0, []int{}}
}

func (this *CustomStack) Push(x int) {
	//满了就不加了
	if this.maxSize == this.stackLen {
		return
	}
	//添加一个元素并且栈长度加1
	this.stack = append(this.stack, x)
	this.stackLen++
}

func (this *CustomStack) Pop() int {
	if this.stackLen == 0 {
		return -1
	}
	popValue := this.stack[this.stackLen-1]
	this.stackLen--
	this.stack = this.stack[:this.stackLen]
	return popValue
}

func (this *CustomStack) Increment(k int, val int) {
	if k > this.stackLen {
		k = this.stackLen
	}
	for i := 0; i < k; i++ {
		this.stack[i] = this.stack[i] + val
	}
}

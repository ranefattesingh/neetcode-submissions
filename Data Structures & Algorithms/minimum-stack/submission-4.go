type MinStack struct {
    minStack []int
    stack []int
}

func Constructor() MinStack {
    return MinStack{
        minStack: []int{},
        stack: []int{},
    }
}

func (this *MinStack) Push(val int) {
    this.stack = append(this.stack, val)

    if len(this.minStack) == 0 {
        this.minStack = append(this.minStack, val)
        return
    }

    currentMin := this.GetMin()
    if val <= currentMin {
        this.minStack = append(this.minStack, val)
    }
}

func (this *MinStack) Pop() {
    top := this.stack[len(this.stack)-1]
    this.stack = this.stack[:len(this.stack)-1]

    if top == this.GetMin() {
        this.minStack = this.minStack[:len(this.minStack)-1]
    }
}

func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
    if len(this.minStack) == 0 {
        return math.MaxInt64
    }

    return this.minStack[len(this.minStack)-1]
}

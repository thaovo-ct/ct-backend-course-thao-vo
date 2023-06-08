package main

import "fmt"

type Stack struct {
	data []interface{}
}

// Create a new stack
func New() *Stack {
	stack := Stack{make([]interface{}, 0)}
	return &stack
}

// View the top item on the stack
func (this *Stack) Peek() interface{} {
	return this.data[this.Len()-1]
}

// Return the number of items in the stack
func (this *Stack) Len() int {
	return len(this.data)
}

// Pop the top item of the stack and return it
func (this *Stack) Pop() interface{} {
	popItem := this.Peek()
	this.data = this.data[:this.Len()-1]
	return popItem
}

// Push a value onto the top of the stack
func (this *Stack) Push(value interface{}) {
	this.data = append(this.data, value)
}

func isValid(s string) bool {
    var stack = New()
    var m = map[string]int{
        "(": 1,
        ")": -1,
        "{": 2,
        "}": -2,
        "[": 3, 
        "]": -3}
    for i := 0; i < len(s); i++{
        c := string(s[i])
        if m[c] > 0{
            stack.Push(m[c])
        } else {
            if stack.Len() == 0 {
                return false
            }
            sign := stack.Pop()
            if sign.(int) + m[c] != 0{
                return false
            }
        }
    }
    if stack.Len() > 0{
        return false
    }
    return true
}


func main(){
	s := ")"
	fmt.Println(isValid(s))
}
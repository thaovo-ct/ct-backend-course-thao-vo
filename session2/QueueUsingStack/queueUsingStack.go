package main

import (
	"fmt"
)
type MyQueue struct {
	s1 []int
	s2 []int
	front int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
    // only the first time
	if len(this.s1) == 0 {
		this.front = x
	}
	this.s1 = append(this.s1, x)
}

func (this *MyQueue) Pop() int {
    // Only if s2 is empty
	if len(this.s2) == 0 {
		for len(this.s1) > 0 {
			topElement := this.s1[len(this.s1)-1]
			this.s1 = this.s1[:len(this.s1)-1]
			this.s2 = append(this.s2, topElement)
		}
	}
	topElement := this.s2[len(this.s2)-1]
	this.s2 = this.s2[:len(this.s2)-1]
	return topElement
}

func (this *MyQueue) Peek() int {
	if len(this.s2) != 0 {
		return this.s2[len(this.s2)-1]
	}
	return this.front
}

func (this *MyQueue) Empty() bool {
	return len(this.s1) == 0 && len(this.s2) == 0
}

func main (){
	obj := Constructor();
	obj.Push(1);
	obj.Push(2);
	param_2 := obj.Pop();
	param_3 := obj.Peek();
	param_4 := obj.Empty();
	fmt.Println(param_2)
	fmt.Println(param_3)
	fmt.Println(param_4)

}
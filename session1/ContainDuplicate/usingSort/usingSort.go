package main

import (
    "sort" 
    "fmt"
)

func containsDuplicate(nums []int) bool {
    sort.Ints(nums)
    for i := 1; i < len(nums); i++{
        if nums[i] == nums[i-1] {
            return true
        }
    }
    return false
}

type Animal interface {
    Name() string
    Type() string
}

type Dog struct {
    name string
}

func (d *Dog) Name() string {
    return d.name
}

func (d *Dog) Type() string {
    return d.name
}

func animalName(a Animal) {
    fmt.Println("Say hello" + a.Name())
}


type Student struct {
    Name string
    Age int
}

func SortByAgeFn(a, b any) bool {
    return a.Age < b.Age
}


func main(){
    nums := []int{1, 2, 3, 1}
    fmt.Println(containsDuplicate(nums));
    d := Dog{}
    d.name = "123"
    a := make([]Student, 10)
    sort.SliceIsSorted(a, SortByAgeFn(a))
    animalName(&d)
}
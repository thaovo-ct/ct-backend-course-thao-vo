package main

import "fmt"


func containsDuplicate(nums []int) bool {
    m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
        _, existed := m[nums[i]]
        if existed == true {
            return true
        } else {
            m[nums[i]] = 1
        }
    }
    return false
}

func main() {
	nums := []int{1, 2, 3, 4}
    fmt.Println(containsDuplicate(nums));
}
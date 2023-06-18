package main

import (
	"fmt"
	"sort"
	"strings"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t){
        return false
    }
    sliceT := strings.Split(t, "")
	sort.Strings(sliceT)

    sliceS := strings.Split(s, "")
    sort.Strings(sliceS)

    s = strings.Join(sliceS, "")
    t = strings.Join(sliceT, "")
    if s == t {
        return true
    } else {
        return false
    }
}

func main (){
	fmt.Println(isAnagram("anaggram", "gramana"))
}
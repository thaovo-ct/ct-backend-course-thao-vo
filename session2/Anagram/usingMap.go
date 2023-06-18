package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t){
        return false
    }
	mapS := make(map[byte]int)
	mapT := make(map[byte]int)
	for i := 0; i < len(s); i++{
		_, existedS := mapS[s[i]]
        if existedS == true {
            mapS[s[i]] += 1
        } else {
            mapS[s[i]] = 0
        }

		_, existedT := mapT[t[i]]
        if existedT == true {
            mapT[t[i]] += 1
        } else {
            mapT[t[i]] = 0
        }
	}
	for i:= 0; i < len(s); i++{
		if mapS[s[i]] != mapT[s[i]]{
			return false
		}
	}
	return true
}

func main (){
	fmt.Println(isAnagram("anagram", "graamana"))
}
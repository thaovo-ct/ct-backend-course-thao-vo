package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	_ "strconv"
)

func main(){
	out, _ := ReadFromFile("input.txt")

	for _, v := range out {
		var user User
		if err := json.Unmarshal([]byte(v), &user); err != nil {
			panic(err)
		}
		fmt.Println(*user.Address)
	}
	
}

// INPUT(byte)     -> DEcode        -> Process   -> Encode         -> OUTPUT(byte)
// io.Reader       json unmarshal                 json marshal         io.Writer



//check duplicate
// requirement
// todo#1: input from a file
//     read into from file
//info json accountid age

func ReadFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var output [] string
	for scanner.Scan() {
		line := scanner.Text()

		// fmt.Println(line)
		if err != nil{
			return nil, err
		}
		output = append(output, line)		
	}
	return output,nil
}

func WriteFile(filename string) {
	
}


func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
    for i := 1; i < len(nums); i++{
        if nums[i] == nums[i-1] {
            return true
        }
    }
    return false
}

type User struct {
	Name string `json:"name"`
	AccountId int `json:"accountId"`
	Age int `json:"age"`
	Address *string `json:"address"`
}
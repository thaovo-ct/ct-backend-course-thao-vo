// You can edit this code!
// Click here and start typing.
package main

import (
	"bufio"
	"encoding/json"
	"reflect"

	"fmt"
	"log"
	"os"
)

func main() {
	out, err := ReadFromFile("input.txt")
	if err != nil {
		panic(err)
	}

	var listUser []User

	for _, v := range out {
		//var user map[string]interface{}
		var user User
		if err := json.Unmarshal([]byte(v), &user); err != nil {
			panic(err)
		}

		listUser = append(listUser, user)
		//printStructFields(user)
	}

	writeStructToJSONFile(listUser[0], "out_1.txt")

	fmt.Println(len(out))

}

func writeStructToJSONFile(data interface{}, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	n, err := file.Write(b)
	if err != nil {
		return err
	}

	if len(b) != n {
		panic("not correct")
	}

	return nil
}

func ReadFromFile(filename string) ([]string, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var output []string
	// Read line by line
	for scanner.Scan() {
		line := scanner.Text()
		// Process the line here

		output = append(output, line)
	}

	// Check for any errors encountered during scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return output, nil
}

// Check Duplicate
// TODO #1: input from a file
// Read info from file
// Info Json: AccountID, Name, Age: check duplicate accountID

func containsDuplicate(nums []int) bool {
	numCount := make(map[int]int)
	for _, num := range nums {
		numCount[num]++
		if numCount[num] > 1 {
			return true
		}
	}
	return false
}

type User struct {
	AccountId int     `json:"account_id"`
	Name      string  `json:"name"`
	Age       int     `json:"age"`
	Address   Address `json:"address"`
}

type Address struct {
	Street string `json:"street"`
	Ward   string `json:"ward"`
}

func printStructFields(s interface{}) {
	v := reflect.ValueOf(s)
	t := v.Type()

	// Iterate over the fields of the struct
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)

		// Get the field name and value
		fieldName := field.Name
		fieldValueInterface := fieldValue.Interface()

		// Print the field name and value
		fmt.Printf("%s: %v\n", fieldName, fieldValueInterface)
	}
}
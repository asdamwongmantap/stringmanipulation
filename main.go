package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main()  {
	var kata1 string = "My Name is Asdam"
	var kata2 string = "my name is asdam"
	var kata3 = []string{"asdam","khan","student"}
	//Convert lowercase to uppercase
	fmt.Println(strings.ToUpper(kata1))
	//Convert uppercase to lowercase
	fmt.Println(strings.ToLower(kata1))
	//Check if string contains specific character
	fmt.Println(strings.Contains(kata1,"is"))
	//Convert first character in word to uppercase
	fmt.Println(strings.Title(kata1))
	//Compare word / sentence with non case sensitive
	fmt.Println(strings.EqualFold(kata1,kata2))
	//Replace character or word with another
	fmt.Println(strings.Replace(kata2,"asdam","khan",-1))
	//Join word
	fmt.Println(strings.Join(kata3," and "))
	//Split word
	fmt.Println(strings.Split(kata1," "))
	//Count specific character
	fmt.Println(strings.Count(kata1,"m"))
	//Convert to integer
	fmt.Println(strconv.Atoi("1"))
	fmt.Println(strconv.ParseInt("1",10,64))
	//Convert to float64
	fmt.Println(strconv.ParseFloat("1",64))
}

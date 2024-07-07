package main

import "fmt"

func main() {

	// Q6 answer

	word := " cafe,face "

	switch word {

	case " donut,peanut":
		{
			fmt.Println("not anagram")
		}

	case "cafe,face":
		{
			fmt.Println("anagram")
		}

	}

}

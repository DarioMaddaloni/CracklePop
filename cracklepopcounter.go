package main

import "fmt"

func main() {
	crackle := "Crackle"
	pop := "Pop"
	for i:=1; i<=100; i++ {
		if i%3==0 { // Check i%3==0 before i%5==0, because it is more frequent
			if i%5==0 {
				fmt.Println(crackle+pop)
			} else {
				fmt.Println(crackle)
			}
		} else if i%5==0 { // if this line is executed, it can be assumed that i%3!=0
			fmt.Println(pop)
		} else {
			fmt.Println(i)
		}
	}
}

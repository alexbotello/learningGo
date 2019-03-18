package main

import "fmt"

func main() {
	billy := Engineer{
		name: "Billy",
		age:  24,
	}

	susan := Teacher{
		name: "Susan",
		age:  21,
	}

	fmt.Println(susan.Teach())
	fmt.Println(billy.Build())
	fmt.Println(TellJoke(billy))
}

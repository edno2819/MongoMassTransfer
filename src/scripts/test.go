package main

import "fmt"

func main() {
	b := 10
	p := &b
	d := &p

	fmt.Println(p)
	fmt.Println(&p)
	fmt.Println(*p)
	fmt.Println(**d)

}

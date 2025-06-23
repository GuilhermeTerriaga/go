package main

import "fmt"

func main()  {
	s:= "Hello, golang"
	s2:=`Hello 
			golang
		gopher!` 
	fmt.Printf("%v\n%T\n", s,s)
	fmt.Printf("%v\n", s2)
}

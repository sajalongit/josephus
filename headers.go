package main

type PersonNode struct {
	P     *Person
	NextP *PersonNode
}
type Person struct {
	Name int
}

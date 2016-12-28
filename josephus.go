package main

import (
	"fmt"
)

func main() {
	personWithSword, err := MakePeople(100) //first person with sword

	if err != nil {
		panic(err)
	}

	personSurvived := ExecutePeople(personWithSword)

	fmt.Printf("%d survived", personSurvived.P.Name)
}

func NewPersonNode(n int) *PersonNode {
	newPerson := NewPerson(n)
	return &PersonNode{P: newPerson}
}

func NewPerson(n int) *Person {
	return &Person{Name: n}
}

func MakePeople(n int) (*PersonNode, error) {
	if n < 1 {
		return nil, fmt.Errorf("MakePeople: n:%d cannot be less than 1", n)
	}

	var firstPersonNode *PersonNode //first person with sword

	//Setup 99 more people
	var previousPersonNode *PersonNode

	for i := 1; i <= n; i++ {
		newPersonNode := NewPersonNode(i)

		if firstPersonNode == nil {
			firstPersonNode = newPersonNode
		}

		if previousPersonNode == nil {
			previousPersonNode = newPersonNode
		}

		previousPersonNode.NextP = newPersonNode

		if i == n {
			newPersonNode.NextP = firstPersonNode
		}

		previousPersonNode = newPersonNode
	}

	return firstPersonNode, nil
}

func ExecutePerson(personWithSword *PersonNode) *PersonNode {
	personWithSword.NextP = personWithSword.NextP.NextP

	return personWithSword.NextP
}

func ExecutePeople(personWithSword *PersonNode) *PersonNode {
	var personSurvived *PersonNode = personWithSword
	for {
		personWithSword = ExecutePerson(personWithSword)

		if personWithSword.NextP == personWithSword {
			personSurvived = personWithSword
			break
		}
	}
	return personSurvived
}

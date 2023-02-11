package main

import (
	"design_patterns/behavioral/iterator/collection"
	"design_patterns/behavioral/iterator/user"
	"fmt"
)

func main() {
	userCollection := &collection.UserCollection{
		Users: []*user.User{
			{
				Name: "A",
				Age:  30,
			},
			{
				Name: "B",
				Age:  20,
			},
		},
	}

	iterator := userCollection.CreateIterator()
	for iterator.HasNext() {
		user := iterator.GetNext()
		fmt.Printf("User is %+v\n", user)
	}

}

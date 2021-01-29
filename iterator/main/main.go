package main

import (
	"design_patterns/iterator"
	"fmt"
)

func main() {
	//创建用户并且放入到集合中
	user1 := &iterator.User{
		Name: "a",
		Age:  30,
	}

	user2 := &iterator.User{
		Name: "b",
		Age:  20,
	}

	//遍历集合
	userCollection := iterator.UserCollection{Users: []*iterator.User{user1, user2}}
	userIterator := userCollection.CreateIterator()
	for userIterator.HasNext() {
		user := userIterator.GetNext()
		fmt.Println(user)
	}

}

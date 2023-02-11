package collection

import (
	"design_patterns/behavioral/iterator/iterator"
	"design_patterns/behavioral/iterator/user"
)

type UserCollection struct {
	Users []*user.User
}

func (u *UserCollection) CreateIterator() iterator.Iterator {
	// 返回用户迭代器
	return &iterator.UserIterator{
		Users: u.Users,
	}

}

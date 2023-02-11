package iterator

import "design_patterns/behavioral/iterator/user"

type Iterator interface {
	HasNext() bool // 是否还有下一个
	GetNext() *user.User
}

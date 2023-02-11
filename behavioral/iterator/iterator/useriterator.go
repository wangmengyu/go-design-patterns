package iterator

import "design_patterns/behavioral/iterator/user"

type UserIterator struct {
	//  具体实现迭代器的类
	index int          //  当前索引位置
	Users []*user.User // 用户列表
}

func (u *UserIterator) HasNext() bool {
	return u.index < len(u.Users)
}

func (u *UserIterator) GetNext() *user.User {
	if u.HasNext() {
		nextUser := u.Users[u.index]
		u.index++ // 顺序执行就行
		return nextUser
	}
	return nil
}

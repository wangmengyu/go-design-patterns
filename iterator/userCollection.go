package iterator

type UserCollection struct {
	Users []*User
}

func (uc *UserCollection) CreateIterator() Iterator {
	return &UserIterator{
		Index: 0,
		Users: uc.Users,
	}
}

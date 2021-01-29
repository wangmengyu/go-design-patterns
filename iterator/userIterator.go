package iterator

type UserIterator struct {
	Index int
	Users []*User
}

func (ui *UserIterator) HasNext() bool {
	if ui.Index > len(ui.Users)-1 {
		return false
	}
	return true
}
func (ui *UserIterator) GetNext() interface{} {
	user := ui.Users[ui.Index]
	if ui.HasNext() {
		ui.Index++
	}
	return user
}

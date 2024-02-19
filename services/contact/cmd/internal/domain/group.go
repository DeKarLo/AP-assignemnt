package domain

type Group struct {
	ID   int
	Name string
}

func NewGroup(id int, name string) *Group {
	return &Group{
		ID:   id,
		Name: name,
	}
}

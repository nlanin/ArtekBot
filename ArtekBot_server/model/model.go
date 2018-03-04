package model

//import (
//	"fmt"
//)

type Model struct {
	db
}

func New(db db) *Model {
	return &Model{
		db: db,
	}
}

func (m *Model) People() ([]*Person, error) {
	//v,err := m.SelectPeople()
	//fmt.Printf("%v %s",m.SelectPeople())
	//fmt.Println(&v,err)
	return m.SelectPeople()
}

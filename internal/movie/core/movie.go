package core

import "fmt"

type (
	Movie interface {
		//temp variable name and type, change depend on value to be saved
		SaveLog(title string) (err error)
	}
	movie struct{}
)

func (m *movie) SaveLog(title string) (err error) {
	//dummy save to db
	fmt.Println("save to DB : ", title)
	return nil
}

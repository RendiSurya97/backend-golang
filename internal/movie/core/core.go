package core

import "log"

//core is for something related to DB (and redis)
func New() Movie {
	err := prepare()
	if err != nil {
		log.Println("[Core] failed to prepare: ", err.Error())
	}

	return &movie{}
}

// prepare is an internal function that init package (ex, init db) for this package
func prepare() (err error) {
	//prepare database (and redis if needed)
	return nil
}

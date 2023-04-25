package greeting

import (
	"database/sql"
	"fmt"
)

type Hoge struct {
	name string
}

func NewHoge(name string) *Hoge {
	return &Hoge{
		name: name,
	}
}

func (h *Hoge) Greeting() {
	fmt.Println("hello " + h.name + "!")
}

type Fuga struct {
	name string
}

func NewFuga(name string) *Fuga {
	return &Fuga{
		name: name,
	}
}

func (f *Fuga) Greeting() {
	fmt.Println("hello " + f.name)
}

func SQLInjection(h string) {
	f := NewFuga("SELECT * FROM hoge WHERE id = " + h + ";")
	f.Greeting()
}

func SQLInjectionForDB(h string) {
	var db *sql.DB
	rows, _ := db.Query("SELECT * FROM hoge WHERE id = " + h + ";")
	for rows.Next() {
	}
}

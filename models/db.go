package models

import (
	"github.com/gobuffalo/pop"
)

func NewDB() (*pop.Connection, error) {
	return pop.Connect("development")
}

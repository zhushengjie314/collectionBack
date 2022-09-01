package db

import (
	"fmt"
)

type Name = string
type Ip = string
type Port = int32
type Pass = string
type User = string
type Base = string

type Db interface {
	Type() string
	Connect()
}

type DbMap map[Name]Db

var Conn DbMap

func init() {
	Conn = make(map[Name]Db)
	fmt.Println(Conn)
}

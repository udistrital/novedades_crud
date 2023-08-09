package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AgregarCampoTabla_20230124_112600 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarCampoTabla_20230124_112600{}
	m.Created = "20230124_112600"

	migration.Register("AgregarCampoTabla_20230124_112600", m)
}

// Run the migrations
func (m *AgregarCampoTabla_20230124_112600) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20230124_112600_agregar_campo_up.sql")
	if err != nil {
		fmt.Println(err)
	}
	requests := strings.Split(string(file), ";")
	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}

// Reverse the migrations
func (m *AgregarCampoTabla_20230124_112600) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20230124_112600_agregar_campo_down.sql")

	if err != nil {
		fmt.Println(err)
	}
	requests := strings.Split(string(file), ";")
	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}

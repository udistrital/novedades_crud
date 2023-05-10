package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AgregarCampoDocumento_20230504_054753 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarCampoDocumento_20230504_054753{}
	m.Created = "20230504_054753"

	migration.Register("AgregarCampoDocumento_20230504_054753", m)
}

// Run the migrations
func (m *AgregarCampoDocumento_20230504_054753) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20230504_054753_agregar_campo_documento_up.sql")
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
func (m *AgregarCampoDocumento_20230504_054753) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20230504_054753_agregar_campo_documento_down.sql")

	if err != nil {
		fmt.Println(err)
	}
	requests := strings.Split(string(file), ";")
	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}

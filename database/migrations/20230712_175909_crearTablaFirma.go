package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CrearTablaFirma_20230712_175909 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaFirma_20230712_175909{}
	m.Created = "20230712_175909"

	migration.Register("CrearTablaFirma_20230712_175909", m)
}

// Run the migrations
func (m *CrearTablaFirma_20230712_175909) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20230707_104407_crear_tabla_firma_up.sql")
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
func (m *CrearTablaFirma_20230712_175909) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20230707_104407_crear_tabla_firma_down.sql")

	if err != nil {
		fmt.Println(err)
	}
	requests := strings.Split(string(file), ";")
	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}

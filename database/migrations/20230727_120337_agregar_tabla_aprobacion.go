package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AgregarTablaAprobacion_20230727_120337 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarTablaAprobacion_20230727_120337{}
	m.Created = "20230727_120337"

	migration.Register("AgregarTablaAprobacion_20230727_120337", m)
}

// Run the migrations
func (m *AgregarTablaAprobacion_20230727_120337) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20230727_120337_agregar_tabla_aprobacion_up.sql")
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
func (m *AgregarTablaAprobacion_20230727_120337) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20230727_120337_agregar_tabla_aprobacion_down.sql")

	if err != nil {
		fmt.Println(err)
	}
	requests := strings.Split(string(file), ";")
	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}

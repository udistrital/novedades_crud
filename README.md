# novedades_crud
novedades_crud, API crud con modelo relacional en postgresSQL, el proyecto está programado en el lenguaje Go y creado con el [framework beego](https://beego.me/).

el api provee la gestion de las diferentes novedades que puedan ser creadas en el sistemas

***Instlaciones Previas:***
* [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
* [BeeGo](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)


## Especificaciones Técnicas

### Instalación

#### Opción 1
```shell
go get github.com/udistrital/novedades_crud
```
#### Opción 2
```shell
# Para clonar el proyecto en la carpeta local go/src/github.com/udistrital
cd go/src/github.com/udistrital

# clonar repo
git clone https://github.com/udistrital/novedades_crud.git

#  Ir a la carpeta del proyecto
cd novedades_crud

# Instalar dependencias del proyecto
go get
```

### Variables de Entorno

- NOVEDADES_CRUD__PGDB=[nombre de la base de datos]
- NOVEDADES_CRUD__PGPASS=[password del usuario]
- NOVEDADES_CRUD__PGURLS=[direccion de la base de datos]
- NOVEDADES_CRUD__PGUSER=[usuario con acceso a la base de datos]
- NOVEDADES_CRUD__PGSCHEMA=[esquema donde se ubican las tablas]
- NOVEDADES_HTTP_PORT=[puerto de ejecucion] bee run


### Ejecución del proyecto

* Ubicado en la raíz del proyecto, ejecutar:
```bash
NOVEDADES_CRUD__PGDB=XXX NOVEDADES_CRUD__PGPASS=XXX NOVEDADES_CRUD__PGURLS=XXX NOVEDADES_CRUD__PGUSER=XXX NOVEDADES_CRUD__PGSCHEMA=XXX NOVEDADES_HTTP_PORT=XXX bee run
```
* O si se quiere ejecutar el swager:
```shell
NOVEDADES_CRUD__PGDB=XXX NOVEDADES_CRUD__PGPASS=XXX NOVEDADES_CRUD__PGURLS=XXX NOVEDADES_CRUD__PGUSER=XXX NOVEDADES_CRUD__PGSCHEMA=XXX NOVEDADES_HTTP_PORT=XXX bee run -downdoc=true -gendoc=true
```

### Puertos

* El servidor se expone por defecto en el puerto: localhost:8080

* Para ver la documentación de swagger: [localhost:8080/swagger/](http://localhost:8080/swagger/)
    *Nota*: En el swagger sale un error, hacer caso omiso.

### EndPoints

Al ejecutar el swagger se puede tener mayor apreciacion de los diferentes metodos de peticion por cada endpoint cuales son los distinpos endpoint disponibles y como usarlos.

### Modelo de datos de novedades_crud
El modelo de datos del API crud de novedades se muestra en la siguiente imágen:
![novedades](https://user-images.githubusercontent.com/28914781/65917368-d0438500-e39c-11e9-8831-c13f4048309f.png)


## Licencia

This file is part of novedades_crud.

novedades_crud is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

Foobar is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with Foobar. If not, see https://www.gnu.org/licenses/.

# novedades_crud
El API provee la gestion de las diferentes novedades pos-contractuales sobre diferentes contratos.

## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones
* [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
* [BeeGo](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)
* [Docker](https://docs.docker.com/engine/install/ubuntu/)
* [Docker Compose](https://docs.docker.com/compose/)

### Variables de Entorno
```shell
NOVEDADES_CRUD__PGDB=[nombre de la base de datos]
NOVEDADES_CRUD__PGPASS=[password del usuario]
NOVEDADES_CRUD__PGURLS=[direccion de la base de datos]
NOVEDADES_CRUD__PGUSER=[usuario con acceso a la base de datos]
NOVEDADES_CRUD__PGSCHEMA=[esquema donde se ubican las tablas]
NOVEDADES_HTTP_PORT=[puerto de ejecucion] bee run
```

**NOTA:** Las variables se pueden ver en el fichero conf/app.conf y están identificadas con NOVEDADES_CRUD_...

### Ejecución del Proyecto
```shell
#1. Obtener el repositorio con Go
go get github.com/udistrital/novedades_crud

#2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/udistrital/novedades_crud

# 3. Moverse a la rama **develop**
git pull origin develop && git checkout develop

# 4. alimentar todas las variables de entorno que utiliza el proyecto.
NOVEDADES_CRUD_PORT=8080 NOVEDADES_CRUD_DB_HOST=127.0.0.1:27017 NOVEDADES_CRUD_SOME_VARIABLE=some_value bee run
```

### Ejecución Dockerfile
```shell
# docker build --tag=novedades_crud . --no-cache
# docker run -p 80:80 novedades_crud
```

### Ejecución docker-compose
```shell
#1. Clonar el repositorio
git clone -b develop https://github.com/udistrital/novedades_crud

#2. Moverse a la carpeta del repositorio
cd novedades_crud

#3. Crear un fichero con el nombre **custom.env**
# En windows ejecutar:* ` ni custom.env`
touch custom.env

#4. Crear la network **back_end** para los contenedores
docker network create back_end

#5. Ejecutar el compose del contenedor
docker-compose up --build

#6. Comprobar que los contenedores estén en ejecución
docker ps
```

### Ejecución Pruebas

Pruebas unitarias
```shell
# En Proceso
```
## Estado CI

| Develop | Relese 0.0.1 | Master |
| -- | -- | -- |
| [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/novedades_crud/status.svg?ref=refs/heads/develop)](https://hubci.portaloas.udistrital.edu.co/udistrital/novedades_crud) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/novedades_crud/status.svg?ref=refs/heads/release/0.0.1)](https://hubci.portaloas.udistrital.edu.co/udistrital/novedades_crud) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/novedades_crud/status.svg)](https://hubci.portaloas.udistrital.edu.co/udistrital/novedades_crud) |


## Modelo de Datos
[Modelo de Datos API CRUD Novedades](https://user-images.githubusercontent.com/28914781/65917368-d0438500-e39c-11e9-8831-c13f4048309f.png)


## Licencia

This file is part of novedades_crud.

novedades_crud is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

novedades_crud is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with novedades_crud. If not, see https://www.gnu.org/licenses/.

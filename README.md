# novedades_crud
novedades_crud, API crud con modelo relacional en postgresSQL, el proyecto está programado en el lenguaje Go y creado con el [framework beego](https://beego.me/).

***Instlaciones Previas:***
* [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
* [BeeGo](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)

## Modelo de datos de novedades_crud
El modelo de datos del API crud de novedades se muestra en la siguiente imágen: 
![novedades](https://user-images.githubusercontent.com/28914781/65917368-d0438500-e39c-11e9-8831-c13f4048309f.png)


## Configuración del Proyecto.

### Opción 1
Ejecutar desde la terminal 'go get repositorio':
```shell 
go get github.com/udistrital/novedades_crud
```
### Opción 2
Para instalar el proyecto realizar los siguientes pasos:
- Para clonar el proyecto en la carpeta local go/src/github.com/udistrital ir a la consola y ejecutar:
```shell 
    cd go/src/github.com/udistrital
```
- Ejecutar:
```shell 
    git clone https://github.com/udistrital/novedades_crud.git
```

- Ir a la carpeta del proyecto:
```shell 
    cd novedades_crud
```

- Instalar dependencias del proyecto:
```shell 
    go get
```

## Variables de entorno 

* El puerto por el que se expone la api **httpport = 8080**; si se cambia de puerto se debe editar la configuración en el [cliente](https://github.com/udistrital/novedades_cliente), especificamente la varible de entorno NOVEDADES_CRUD_SERVICE.
* La variable de entorno corresponde al puerto en donde se desplegará el API y corresponde a la siguiente :
```shell 
    NOVEDADES_CRUD_HTTP_PORT=8080
```

## Ejecución del proyecto

* Ubicado en la raíz del proyecto, ejecutar:
```shell 
    NOVEDADES_API_HTTP_PORT=8080 bee run
```
* O si se quiere ejecutar el swager:
```shell 
    NOVEDADES_API_HTTP_PORT=8080 bee run -downdoc=true -gendoc=true
```

### Puertos

* El servidor se expone en el puerto: localhost:8080

* Para ver la documentación de swagger: [localhost:8080/swagger/](http://localhost:8080/swagger/)
    *Nota*: En el swagger sale un error, hacer caso omiso.

### EndPoints

Cada controlador tiene los metodos :
* Post
 los endpoint a los cuales apuntar son los siguientes:


||End Point|
|----------------|------------------------|
| **registroNovedad** | `[host de la maquina]:[puerto]/v1/trNovedad` |


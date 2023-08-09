package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:FechasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:FechasController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:FechasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:FechasController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:FechasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:FechasController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:FechasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:FechasController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:FechasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:FechasController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:NovedadesPoscontractualesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:NovedadesPoscontractualesController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:NovedadesPoscontractualesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:NovedadesPoscontractualesController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:NovedadesPoscontractualesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:NovedadesPoscontractualesController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:NovedadesPoscontractualesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:NovedadesPoscontractualesController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:NovedadesPoscontractualesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:NovedadesPoscontractualesController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:PolizaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:PolizaController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:PolizaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:PolizaController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:PolizaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:PolizaController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:PolizaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:PolizaController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:PolizaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:PolizaController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:PropiedadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:PropiedadController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:PropiedadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:PropiedadController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:PropiedadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:PropiedadController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:PropiedadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:PropiedadController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:PropiedadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:PropiedadController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:TipoFechaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:TipoFechaController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:TipoFechaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:TipoFechaController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:TipoFechaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:TipoFechaController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:TipoFechaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:TipoFechaController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:TipoFechaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:TipoFechaController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:TipoNovedadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:TipoNovedadController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:TipoNovedadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:TipoNovedadController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:TipoNovedadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:TipoNovedadController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:TipoNovedadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:TipoNovedadController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:TipoNovedadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:TipoNovedadController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:TipoPropiedadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:TipoPropiedadController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:TipoPropiedadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:TipoPropiedadController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:TipoPropiedadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:TipoPropiedadController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:TipoPropiedadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:TipoPropiedadController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:TipoPropiedadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:TipoPropiedadController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:Tr_novedad_poscontractualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:Tr_novedad_poscontractualController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:Tr_novedad_poscontractualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:Tr_novedad_poscontractualController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:Tr_novedad_poscontractualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/novedades_crud/controllers:Tr_novedad_poscontractualController"],
		beego.ControllerComments{
			Method:           "PostPoliza",
			Router:           "/trnovedadpoliza/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}

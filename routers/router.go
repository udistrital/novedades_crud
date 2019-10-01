// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/novedades_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/novedades_poscontractuales",
			beego.NSInclude(
				&controllers.NovedadesPoscontractualesController{},
			),
		),

		beego.NSNamespace("/tipo_fecha",
			beego.NSInclude(
				&controllers.TipoFechaController{},
			),
		),

		beego.NSNamespace("/tipo_propiedad",
			beego.NSInclude(
				&controllers.TipoPropiedadController{},
			),
		),

		beego.NSNamespace("/fechas",
			beego.NSInclude(
				&controllers.FechasController{},
			),
		),

		beego.NSNamespace("/poliza",
			beego.NSInclude(
				&controllers.PolizaController{},
			),
		),

		beego.NSNamespace("/propiedad",
			beego.NSInclude(
				&controllers.PropiedadController{},
			),
		),

		beego.NSNamespace("/tipo_novedad",
			beego.NSInclude(
				&controllers.TipoNovedadController{},
			),
		),
	)
	beego.AddNamespace(ns)
}

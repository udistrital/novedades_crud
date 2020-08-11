package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	"github.com/udistrital/utils_oas/time_bogota"
)

type TrNovedadesPoscontractuales struct {
	NovedadPoscontractual *NovedadesPoscontractuales
	Fechas                *[]Fechas
	Propiedad             *[]Propiedad
	//Poliza                *Poliza
}
type TrNovedadesPoscontractualesPoliza struct {
	NovedadPoscontractual *NovedadesPoscontractuales
	Fechas                *[]Fechas
	Propiedad             *[]Propiedad
	Poliza                *[]Poliza
}

// AddTransaccionProyectoAcademica Transacción para registrar toda la información de un proyecto academico
func AddTransaccionNovedadPoscontractualNoPoliza(m *TrNovedadesPoscontractuales) (err error, id int) {
	o := orm.NewOrm()
	err = o.Begin()

	horaRegistro := time_bogota.TiempoBogotaFormato()
	m.NovedadPoscontractual.FechaCreacion = horaRegistro
	m.NovedadPoscontractual.FechaModificacion = horaRegistro
	var idnovedadregistro int = 0

	if idNovedad, errTr := o.Insert(m.NovedadPoscontractual); errTr == nil {
		fmt.Println(idNovedad)
		idnovedadregistro = int(idNovedad)

		for _, v := range *m.Fechas {
			v.IdNovedadesPoscontractuales.Id = int(idNovedad)
			v.FechaCreacion = horaRegistro
			v.FechaModificacion = horaRegistro
			if _, errTr = o.Insert(&v); errTr != nil {
				err = errTr
				fmt.Println(err)
				_ = o.Rollback()
				return errTr, 0
			}
		}

		for _, v := range *m.Propiedad {
			v.IdNovedadesPoscontractuales.Id = int(idNovedad)
			v.FechaCreacion = horaRegistro
			v.FechaModificacion = horaRegistro
			if _, errTr = o.Insert(&v); errTr != nil {
				err = errTr
				fmt.Println(err)
				_ = o.Rollback()
				return errTr, 0
			}
		}

		_ = o.Commit()
	} else {
		err = errTr
		fmt.Println(err)
		_ = o.Rollback()
	}

	return nil, idnovedadregistro
}

// AddTransaccionProyectoAcademica Transacción para registrar toda la información de un proyecto academico
func AddTransaccionNovedadPoscontractualPoliza(m *TrNovedadesPoscontractualesPoliza) (err error, id int) {
	o := orm.NewOrm()
	err = o.Begin()

	horaRegistro := time_bogota.TiempoBogotaFormato()
	m.NovedadPoscontractual.FechaCreacion = horaRegistro
	m.NovedadPoscontractual.FechaModificacion = horaRegistro
	var idnovedadregistro int = 0

	if idNovedad, errTr := o.Insert(m.NovedadPoscontractual); errTr == nil {
		fmt.Println(idNovedad)
		idnovedadregistro = int(idNovedad)

		for _, v := range *m.Fechas {
			v.IdNovedadesPoscontractuales.Id = int(idNovedad)
			v.FechaCreacion = horaRegistro
			v.FechaModificacion = horaRegistro
			if _, errTr = o.Insert(&v); errTr != nil {
				err = errTr
				fmt.Println(err)
				_ = o.Rollback()
				return errTr, 0
			}
		}

		for _, v := range *m.Propiedad {
			v.IdNovedadesPoscontractuales.Id = int(idNovedad)
			v.FechaCreacion = horaRegistro
			v.FechaModificacion = horaRegistro
			if _, errTr = o.Insert(&v); errTr != nil {
				err = errTr
				fmt.Println(err)
				_ = o.Rollback()
				return errTr, 0
			}
		}

		for _, v := range *m.Poliza {
			v.IdNovedadesPoscontractuales.Id = int(idNovedad)
			v.FechaCreacion = horaRegistro
			v.FechaModificacion = horaRegistro
			if _, errTr = o.Insert(&v); errTr != nil {
				err = errTr
				fmt.Println(err)
				_ = o.Rollback()
				return errTr, 0
			}
		}

		_ = o.Commit()
	} else {
		err = errTr
		fmt.Println(err)
		_ = o.Rollback()
	}

	return nil, idnovedadregistro
}

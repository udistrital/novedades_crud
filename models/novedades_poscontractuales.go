package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type NovedadesPoscontractuales struct {
	Id                int    `orm:"column(id);pk;auto"`
	NumeroSolicitud   string `orm:"column(numero_solicitud)"`
	ContratoId        int    `orm:"column(contrato_id)"`
	NumeroCdpId       int    `orm:"column(numero_cdp_id)"`
	Motivo            string `orm:"column(motivo);null"`
	Aclaracion        string `orm:"column(aclaracion);null"`
	Observacion       string `orm:"column(observacion);null"`
	Vigencia          int    `orm:"column(vigencia)"`
	VigenciaCdp       int    `orm:"column(vigencia_cdp)"`
	FechaCreacion     string `orm:"auto_now_add;column(fecha_creacion);type(timestamp without time zone)"`
	FechaModificacion string `orm:"auto_now;column(fecha_modificacion);type(timestamp without time zone)"`
	Activo            bool   `orm:"column(activo)"`
	TipoNovedad       int    `orm:"column(tipo_novedad)"`
	OficioSupervisor  string `orm:"column(oficio_supervisor)"`
	OficioOrdenador   string `orm:"column(oficio_ordenador)"`
	Estado            string `orm:"column(estado)"`
	EnlaceDocumento   string `orm:"column(enlace_documento)"`
}

func (t *NovedadesPoscontractuales) TableName() string {
	return "novedades_poscontractuales"
}

func init() {
	orm.RegisterModel(new(NovedadesPoscontractuales))
}

// AddNovedadesPoscontractuales insert a new NovedadesPoscontractuales into database and returns
// last inserted Id on success.
func AddNovedadesPoscontractuales(m *NovedadesPoscontractuales) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetNovedadesPoscontractualesById retrieves NovedadesPoscontractuales by Id. Returns error if
// Id doesn't exist
func GetNovedadesPoscontractualesById(id int) (v *NovedadesPoscontractuales, err error) {
	o := orm.NewOrm()
	v = &NovedadesPoscontractuales{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllNovedadesPoscontractuales retrieves all NovedadesPoscontractuales matches certain condition. Returns empty list if
// no records exist
func GetAllNovedadesPoscontractuales(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(NovedadesPoscontractuales)).RelatedSel()
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []NovedadesPoscontractuales
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateNovedadesPoscontractuales updates NovedadesPoscontractuales by Id and returns error if
// the record to be updated doesn't exist
func UpdateNovedadesPoscontractualesById(m *NovedadesPoscontractuales) (err error) {
	o := orm.NewOrm()
	v := NovedadesPoscontractuales{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteNovedadesPoscontractuales deletes NovedadesPoscontractuales by Id and returns error if
// the record to be deleted doesn't exist
func DeleteNovedadesPoscontractuales(id int) (err error) {
	o := orm.NewOrm()
	v := NovedadesPoscontractuales{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&NovedadesPoscontractuales{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

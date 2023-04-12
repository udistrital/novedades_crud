package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/udistrital/novedades_crud/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

//  Tr_novedad_poscontractualController operations for Tr_novedad_poscontractual
type Tr_novedad_poscontractualController struct {
	beego.Controller
}

// URLMapping ...
func (c *Tr_novedad_poscontractualController) URLMapping() {
	c.Mapping("Post", c.Post)
	//c.Mapping("GetOne", c.GetOne)
	//c.Mapping("GetAll", c.GetAll)
	// c.Mapping("Put", c.Put)
	//c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Tr_novedad_poscontractual
// @Param	body		body 	models.TrNovedadesPoscontractuales	true		"body for Tr_novedad_poscontractual content"
// @Success 201 {int} models.TrNovedadesPoscontractuales
// @Failure 403 body is empty
// @router / [post]
func (c *Tr_novedad_poscontractualController) Post() {
	var v models.TrNovedadesPoscontractuales
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		fmt.Println("1")
		if err, idnovedad := models.AddTransaccionNovedadPoscontractualNoPoliza(&v); err == nil {
			fmt.Println("2")
			c.Ctx.Output.SetStatus(201)
			v.NovedadPoscontractual.Id = idnovedad
			fmt.Println("Aqui se muestra Id de la inserción \n", idnovedad)
			c.Data["json"] = v
		} else {
			logs.Error(err)
			//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
			c.Data["system"] = err
			c.Abort("400")
		}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("400")
	}
	c.ServeJSON()
}

// Post ...
// @Title Post
// @Description Transacción para crear novedad con poliza
// @Param	body		body 	models.TrNovedadesPoscontractualesPoliza	true		"body for Tr_novedad_poscontractual content"
// @Success 201 {int} models.TrNovedadesPoscontractualesPoliza
// @Failure 400 the request contains incorrect syntax
// @router /trnovedadpoliza/ [post]
func (c *Tr_novedad_poscontractualController) PostPoliza() {
	var v models.TrNovedadesPoscontractualesPoliza
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		fmt.Println("1")
		if err, idnovedad := models.AddTransaccionNovedadPoscontractualPoliza(&v); err == nil {
			fmt.Println("2")
			c.Ctx.Output.SetStatus(201)
			v.NovedadPoscontractual.Id = idnovedad
			fmt.Println("Aqui se muestra Id de la inserción \n", idnovedad)
			fmt.Println(v)
			c.Data["json"] = v
		} else {
			logs.Error(err)
			//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
			c.Data["system"] = err
			c.Data["Code"] = 400
			c.Ctx.Output.SetStatus(400)
			fmt.Println("entro al error en el crud")
			c.Abort("400")
		}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("400")
		c.Ctx.Output.SetStatus(400)
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Tr_novedad_poscontractual by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Tr_novedad_poscontractual
// @Failure 403 :id is empty
// @router /:id [get]
// func (c *Tr_novedad_poscontractualController) GetOne() {
// 	idStr := c.Ctx.Input.Param(":id")
// 	id, _ := strconv.ParseInt(idStr, 0, 64)
// 	v, err := models.GetTr_novedad_poscontractualById(id)
// 	if err != nil {
// 		c.Data["json"] = err.Error()
// 	} else {
// 		c.Data["json"] = v
// 	}
// 	c.ServeJSON()
// }

// // GetAll ...
// // @Title Get All
// // @Description get Tr_novedad_poscontractual
// // @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// // @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// // @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// // @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// // @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// // @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// // @Success 200 {object} models.Tr_novedad_poscontractual
// // @Failure 403
// // @router / [get]
// func (c *Tr_novedad_poscontractualController) GetAll() {
// 	var fields []string
// 	var sortby []string
// 	var order []string
// 	var query = make(map[string]string)
// 	var limit int64 = 10
// 	var offset int64

// 	// fields: col1,col2,entity.col3
// 	if v := c.GetString("fields"); v != "" {
// 		fields = strings.Split(v, ",")
// 	}
// 	// limit: 10 (default is 10)
// 	if v, err := c.GetInt64("limit"); err == nil {
// 		limit = v
// 	}
// 	// offset: 0 (default is 0)
// 	if v, err := c.GetInt64("offset"); err == nil {
// 		offset = v
// 	}
// 	// sortby: col1,col2
// 	if v := c.GetString("sortby"); v != "" {
// 		sortby = strings.Split(v, ",")
// 	}
// 	// order: desc,asc
// 	if v := c.GetString("order"); v != "" {
// 		order = strings.Split(v, ",")
// 	}
// 	// query: k:v,k:v
// 	if v := c.GetString("query"); v != "" {
// 		for _, cond := range strings.Split(v, ",") {
// 			kv := strings.SplitN(cond, ":", 2)
// 			if len(kv) != 2 {
// 				c.Data["json"] = errors.New("Error: invalid query key/value pair")
// 				c.ServeJSON()
// 				return
// 			}
// 			k, v := kv[0], kv[1]
// 			query[k] = v
// 		}
// 	}

// 	l, err := models.GetAllTr_novedad_poscontractual(query, fields, sortby, order, offset, limit)
// 	if err != nil {
// 		c.Data["json"] = err.Error()
// 	} else {
// 		c.Data["json"] = l
// 	}
// 	c.ServeJSON()
// }

// Put ...
// @Title Put
// @Description update the Tr_novedad_poscontractual
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Tr_novedad_poscontractual	true		"body for Tr_novedad_poscontractual content"
// @Success 200 {object} models.Tr_novedad_poscontractual
// @Failure 403 :id is not int
// @router /:id [put]
func (c *Tr_novedad_poscontractualController) Put() {
	// idStr := c.Ctx.Input.Param(":id")
	// id, _ := strconv.ParseInt(idStr, 0, 64)
	// v := models.Tr_novedad_poscontractual{Id: id}
	// json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	// if err := models.UpdateTr_novedad_poscontractualById(&v); err == nil {
	// 	c.Data["json"] = "OK"
	// } else {
	// 	c.Data["json"] = err.Error()
	// }
	// c.ServeJSON()
}

// // Delete ...
// // @Title Delete
// // @Description delete the Tr_novedad_poscontractual
// // @Param	id		path 	string	true		"The id you want to delete"
// // @Success 200 {string} delete success!
// // @Failure 403 id is empty
// // @router /:id [delete]
// func (c *Tr_novedad_poscontractualController) Delete() {
// 	idStr := c.Ctx.Input.Param(":id")
// 	id, _ := strconv.ParseInt(idStr, 0, 64)
// 	if err := models.DeleteTr_novedad_poscontractual(id); err == nil {
// 		c.Data["json"] = "OK"
// 	} else {
// 		c.Data["json"] = err.Error()
// 	}
// 	c.ServeJSON()
// }

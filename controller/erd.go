package controller

import (
	"fmt"
	"net/http"
	"plm/models"
	"plm/responsegenr"

	"github.com/labstack/echo"
)

type Erdget struct {
	Email      string `json:"email"`
	Id_usecase int    `json:"id_usecase"`
	Link_erd   string `json:"link_erd"`
}

type ErdEditget struct {
	Email    string `json:"email"`
	Id       int    `json:"id"`
	Link_erd string `json:"link_erd"`
}

type Tableget struct {
	Email       string `json:"email"`
	Id_erd      int    `json:"id_erd"`
	Table_name  string `json:"table_name"`
	Description string `json:"description"`
}

type TableEditget struct {
	Email       string `json:"email"`
	Id          int    `json:"id"`
	Table_name  string `json:"table_name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type Delget struct {
	Email string `json:"email"`
	Id    int    `json:"id"`
}

type Fieldget struct {
	Email        string `json:"email"`
	Id_table     int    `json:"id_table"`
	Field_name   string `json:"field_name"`
	Field_type   string `json:"field_type"`
	Field_length int    `json:"field_length"`
}

type FielEditdget struct {
	Email        string `json:"email"`
	Id           int    `json:"id`
	Field_name   string `json:"field_name"`
	Field_type   string `json:"field_type"`
	Field_length int    `json:"field_length"`
}

type ViewTableget struct {
	Email  string `json:"email"`
	Id_erd int    `json:"id_erd"`
}

type ViewFieldget struct {
	Email    string `json:"email"`
	Id_table int    `json:"id_table"`
}

type ErdUsecaseViewget struct {
	Email      string `json:"email"`
	Id_usecase int    `json:"id_usecase"`
}

type ViewErdProjget struct {
	Email      string `json:"email"`
	Id_project int    `json:"id_project"`
}

type DelErd struct {
	Id int `json:"id"`
}

//add erd
func (Controller Controller) AddErd(c echo.Context) error {
	a := new(Erdget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.Erdtask{
		Email:      a.Email,
		Id_usecase: a.Id_usecase,
		Link_erd:   a.Link_erd,
	}
	fmt.Println(a.Email)

	addErd := Controller.ma.AddErd(ab)

	if addErd {

		getadd := Erdget{a.Email, a.Id_usecase, a.Link_erd}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil create erd",
			Data:    getadd,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal create erd",
		}
		return c.JSON(http.StatusOK, res)

	}
}

//edit erd
func (Controller Controller) EditErd(c echo.Context) error {
	a := new(ErdEditget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.Erdtask{
		Email:    a.Email,
		Id:       a.Id,
		Link_erd: a.Link_erd,
	}
	fmt.Println(a.Email)

	editErd := Controller.ma.EditErd(ab)

	if editErd {

		getedit := ErdEditget{a.Email, a.Id, a.Link_erd}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil update erd",
			Data:    getedit,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal update erd",
		}
		return c.JSON(http.StatusOK, res)

	}
}

//add table
func (Controller Controller) AddErdTable(c echo.Context) error {
	a := new(Tableget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.Erdtask{
		Email:       a.Email,
		Id_erd:      a.Id_erd,
		Table_name:  a.Table_name,
		Description: a.Description,
	}
	fmt.Println(a.Email)

	addErdTable := Controller.ma.AddErdTable(ab)

	if addErdTable {

		getadd := Tableget{a.Email, a.Id_erd, a.Table_name, a.Description}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil create table erd",
			Data:    getadd,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal create table erd",
		}
		return c.JSON(http.StatusOK, res)

	}
}

//edit table
func (Controller Controller) EditErdTable(c echo.Context) error {
	a := new(TableEditget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.Erdtask{
		Email:       a.Email,
		Id:          a.Id,
		Table_name:  a.Table_name,
		Description: a.Description,
		Status:      a.Status,
	}
	fmt.Println(a.Email)

	editTable := Controller.ma.EditErdTable(ab)

	if editTable {

		getedit := TableEditget{a.Email, a.Id, a.Table_name, a.Description, a.Status}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil update erd",
			Data:    getedit,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal update erd",
		}
		return c.JSON(http.StatusOK, res)

	}
}

//delete table
func (Controller Controller) DelErdTable(c echo.Context) error {
	a := new(Delget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.Erdtask{
		Email: a.Email,
		Id:    a.Id,
	}
	fmt.Println(a.Email)

	editUsecDesc := Controller.ma.DelErdTable(ab.Id)

	if editUsecDesc {

		getedit := Delget{a.Email, a.Id}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil delete table erd",
			Data:    getedit,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal delete table erd",
		}
		return c.JSON(http.StatusOK, res)

	}
}

//add field
func (Controller Controller) AddErdField(c echo.Context) error {
	a := new(Fieldget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.Erdtask{
		Email:        a.Email,
		Id_table:     a.Id_table,
		Field_name:   a.Field_name,
		Field_type:   a.Field_type,
		Field_length: a.Field_length,
	}
	fmt.Println(a.Email)

	addErdField := Controller.ma.AddErdField(ab)

	if addErdField {

		getadd := Fieldget{a.Email, a.Id_table, a.Field_name, a.Field_type, a.Field_length}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil create field erd",
			Data:    getadd,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal create field erd",
		}
		return c.JSON(http.StatusOK, res)

	}
}

//edit field
func (Controller Controller) EditErdField(c echo.Context) error {
	a := new(FielEditdget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.Erdtask{
		Email:        a.Email,
		Id:           a.Id,
		Field_name:   a.Field_name,
		Field_type:   a.Field_type,
		Field_length: a.Field_length,
	}
	fmt.Println(a.Email)

	addErdField := Controller.ma.EditErdField(ab)

	if addErdField {

		getadd := Fieldget{a.Email, a.Id, a.Field_name, a.Field_type, a.Field_length}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil update field erd ",
			Data:    getadd,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal update field erd",
		}
		return c.JSON(http.StatusOK, res)

	}
}

//delete field
func (Controller Controller) DelErdField(c echo.Context) error {
	a := new(Delget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.Erdtask{
		Email: a.Email,
		Id:    a.Id,
	}
	fmt.Println(a.Email)

	editUsecDesc := Controller.ma.DelErdField(ab.Id)

	if editUsecDesc {

		getedit := Delget{a.Email, a.Id}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil delete field erd",
			Data:    getedit,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal delete field erd",
		}
		return c.JSON(http.StatusOK, res)

	}
}

//view erd by usecase
func (Controller Controller) ViewErdUsec(c echo.Context) error {
	a := new(ErdUsecaseViewget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.ErdUsecaseViewtask{
		Email:      a.Email,
		Id_usecase: a.Id_usecase,
	}
	view := Controller.ma.ViewErdUsec(ab)

	res := responsegenr.ResponseGenericGet{
		Status:  "Success",
		Message: "Berhasil dapatkan data  erd by usecase",
		Data:    view,
	}
	return c.JSON(http.StatusOK, res)

}

//view table
func (Controller Controller) ViewErdTable(c echo.Context) error {
	a := new(ViewTableget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.ErdtableViewtask{
		Email:  a.Email,
		Id_erd: a.Id_erd,
	}
	view := Controller.ma.ViewErdTable(ab)

	res := responsegenr.ResponseGenericGet{
		Status:  "Success",
		Message: "Berhasil dapatkan data  table",
		Data:    view,
	}
	return c.JSON(http.StatusOK, res)

}

//view field
func (Controller Controller) ViewErdFeild(c echo.Context) error {
	a := new(ViewFieldget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.ErdfieldViewtask{
		Email:    a.Email,
		Id_table: a.Id_table,
	}
	view := Controller.ma.ViewErdField(ab)

	res := responsegenr.ResponseGenericGet{
		Status:  "Success",
		Message: "Berhasil dapatkan data  field",
		Data:    view,
	}
	return c.JSON(http.StatusOK, res)

}

//view table all
func (Controller Controller) ViewErdTableAll(c echo.Context) error {
	a := new(ViewTableget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.ErdtableViewtask{
		Email:  a.Email,
		Id_erd: a.Id_erd,
	}
	view := Controller.ma.ViewErdTable(ab)

	res := responsegenr.ResponseGenericGet{
		Status:  "Success",
		Message: "Berhasil dapatkan data  table all",
		Data:    view,
	}
	return c.JSON(http.StatusOK, res)

}

//view proj
func (Controller Controller) ViewErdProj(c echo.Context) error {
	a := new(ViewErdProjget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.ErdprojViewtask{
		Email:      a.Email,
		Id_project: a.Id_project,
	}
	view := Controller.ma.ViewErdProj(ab)

	res := responsegenr.ResponseGenericGet{
		Status:  "Success",
		Message: "Berhasil dapatkan data erd by project",
		Data:    view,
	}
	return c.JSON(http.StatusOK, res)

}

// DeleteErd delete timeline
func (Controller Controller) DeleteErd(c echo.Context) error {
	a := new(DelErd)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.Erdtask{
		Id: a.Id,
	}
	fmt.Println(a.Id)

	editUsecDesc := Controller.ma.DeleteErd(ab.Id)

	if editUsecDesc {

		getdelete := DelErd{a.Id}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil delete ERD",
			Data:    getdelete,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal delete ERD",
		}
		return c.JSON(http.StatusOK, res)

	}
}

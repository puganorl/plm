package controller

import (
	"fmt"
	"net/http"
	"plm/models"
	"plm/responsegenr"

	"github.com/labstack/echo"
)

type Dataflowget struct {
	Email           string `json:"email"`
	Id_usecase_desc int    `json:"id_usecase_desc"`
	Link_dataflow   string `json:"link_dataflow"`
	Description     string `json:"description"`
}

type DataflowEditget struct {
	Email         string `json:"email"`
	Id            int    `json:"id"`
	Link_dataflow string `json:"link_dataflow"`
	Description   string `json:"description"`
	Status        string `json:"status"`
}

type DataflowDeleteget struct {
	Id int `json:"id"`
}

type ViewUDescget struct {
	Email           string `json:"email"`
	Id_usecase_desc int    `json:"id_usecase_desc"`
}

type ViewUjourget struct {
	Email          string `json:"email"`
	Id_userjourney int    `json:"id_userjourney"`
}

type ViewDflowStructget struct {
	Email                 string `json:"email"`
	Id_dataflow_structure int    `json:"id_dataflow_structure"`
}

type ViewDflowProjget struct {
	Email       string `json:"email"`
	Id_project  int    `json:"id_project"`
	Dataflow    string `json:"dataflow"`
	Description string `json:"description"`
}

// AddDataflow add data
func (Controller Controller) AddDataflow(c echo.Context) error {
	a := new(Dataflowget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.Dataflowtask{
		Email:           a.Email,
		Id_usecase_desc: a.Id_usecase_desc,
		Link_dataflow:   a.Link_dataflow,
		Description:     a.Description,
	}
	fmt.Println(a.Email)

	addDataflow := Controller.ma.AddDataflow(ab)

	if addDataflow {
		getadd := Dataflowget{a.Email, a.Id_usecase_desc, a.Link_dataflow, a.Description}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil create dataflow",
			Data:    getadd,
		}
		return c.JSON(http.StatusOK, res)

	}

	res := responsegenr.ResponseGeneric{
		Status:  "Error",
		Message: "Gagal create dataflow",
	}

	return c.JSON(http.StatusOK, res)
}

// EditDataflow edit dataflow
func (Controller Controller) EditDataflow(c echo.Context) error {
	a := new(DataflowEditget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.Dataflowtask{
		Email:         a.Email,
		Id:            a.Id,
		Link_dataflow: a.Link_dataflow,
		Description:   a.Description,
		Status:        a.Status,
	}

	fmt.Println(a.Email)

	editDataflow := Controller.ma.EditDataflow(ab)

	if editDataflow {
		getedit := DataflowEditget{a.Email, a.Id, a.Link_dataflow, a.Description, a.Status}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil update dataflow",
			Data:    getedit,
		}
		return c.JSON(http.StatusOK, res)

	}

	res := responsegenr.ResponseGeneric{
		Status:  "Error",
		Message: "Gagal update dataflow",
	}

	return c.JSON(http.StatusOK, res)
}

// ViewDataflowDesc view dataflow by Descripton
func (Controller Controller) ViewDataflowDesc(c echo.Context) error {
	a := new(ViewUDescget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.DataflowDescViewtask{
		Email:           a.Email,
		Id_usecase_desc: a.Id_usecase_desc,
	}
	view := Controller.ma.ViewDataflowDesc(ab)

	res := responsegenr.ResponseGenericGet{
		Status:  "Success",
		Message: "Berhasil dapatkan data dataflow by description",
		Data:    view,
	}
	return c.JSON(http.StatusOK, res)

}

// ViewUJourDesc view dataflow by Userjourney
func (Controller Controller) ViewUJourDesc(c echo.Context) error {
	a := new(ViewUjourget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.DataflowUjourViewtask{
		Email:          a.Email,
		Id_userjourney: a.Id_userjourney,
	}
	view := Controller.ma.ViewUjour(ab)
	res := responsegenr.ResponseGenericGet{
		Status:  "Success",
		Message: "Berhasil dapatkan data dataflow by userjourney",
		Data:    view,
	}
	return c.JSON(http.StatusOK, res)

}

// ViewDflowStruct view dataflow by dataflow structure
func (Controller Controller) ViewDflowStruct(c echo.Context) error {
	a := new(ViewDflowStructget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.DataflowStructViewtask{
		Email:                 a.Email,
		Id_dataflow_structure: a.Id_dataflow_structure,
	}
	view := Controller.ma.ViewDflowStruct(ab)

	res := responsegenr.ResponseGenericGet{
		Status:  "Success",
		Message: "Berhasil dapatkan data dataflow by dataflow structure",
		Data:    view,
	}
	return c.JSON(http.StatusOK, res)

}

// ViewDflowProj view dataflow by proj
func (Controller Controller) ViewDflowProj(c echo.Context) error {
	a := new(ViewDflowProjget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.DflowProjectView{
		Email:      a.Email,
		Id_project: a.Id_project,
	}
	view := Controller.ma.ViewDflowProj(ab)
	res := responsegenr.ResponseGenericGet{
		Status:  "Success",
		Message: "Berhasil dapatkan data dataflow by project",
		Data:    view,
	}

	return c.JSON(http.StatusOK, res)
}

// DeleteDataflow delete dataflow
func (Controller Controller) DeleteDataflow(c echo.Context) error {
	a := new(DataflowDeleteget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.Dataflowtask{
		Id: a.Id,
	}

	fmt.Println(a.Id)

	deleteDataflow := Controller.ma.DeleteDataflow(ab.Id)

	if deleteDataflow {
		getdelete := DataflowDeleteget{a.Id}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil delete dataflow",
			Data:    getdelete,
		}

		return c.JSON(http.StatusOK, res)
	}

	res := responsegenr.ResponseGeneric{
		Status:  "Error",
		Message: "Gagal delete dataflow",
	}
	return c.JSON(http.StatusOK, res)
}

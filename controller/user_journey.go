package controller

import (
	"fmt"
	"net/http"
	"plm/models"
	"plm/responsegenr"

	"github.com/labstack/echo"
)

type Userjrnyget struct {
	Email          string `json:"email"`
	Id_dataflow    int    `json:"id_dataflow"`
	Journey_name   string `json:"journey_name"`
	Initiate_state int    `json:"initiate_state"`
	Journey_link   string `json:"jouney_link"`
	Journey_index  int    `json:"journey_index"`
	Description    string `json:"description"`
}

type UserjrnyEditget struct {
	Email          string `json:"email"`
	Id_journey     int    `json:"id"`
	Journey_name   string `json:"journey_name"`
	Initiate_state int    `json:"initiate_state"`
	Journey_link   string `json:"jouney_link"`
	Journey_index  int    `json:"journey_index"`
	Status         string `json:"status"`
}

type ViewUserjrnyDflowget struct {
	Email       string `json:"email"`
	Id_dataflow int    `json:"id_dataflow"`
}

type ViewUserjrnyDflowStructget struct {
	Email                 string `json:"email"`
	Id_dataflow_structure int    `json:"id_dataflow_structure"`
}

type ViewUserjrnyProjget struct {
	Email      string `json:"email"`
	Id_project int    `json:"id_project"`
}

type DelUserjrny struct {
	Id int `json:"id"`
}

// AddUserjrnys add userjourney
func (Controller Controller) AddUserjrny(c echo.Context) error {
	a := new(Userjrnyget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.Userjrnytask{
		Email:          a.Email,
		Id_dataflow:    a.Id_dataflow,
		Journey_name:   a.Journey_name,
		Initiate_state: a.Initiate_state,
		Journey_link:   a.Journey_link,
		Journey_index:  a.Journey_index,
		Description:	a.Description,
	}
	fmt.Println(a.Email)

	addUserjrny := Controller.ma.AddUserjrnys(ab)

	if addUserjrny {

		getadd := Userjrnyget{a.Email, a.Id_dataflow, a.Journey_name, a.Initiate_state, a.Journey_link, a.Journey_index, a.Description}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil create dataflow structure",
			Data:    getadd,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal create dataflow structure",
		}
		return c.JSON(http.StatusOK, res)

	}
}

// EditUserjrnys edit userjourney
func (Controller Controller) EditUserjrny(c echo.Context) error {
	a := new(UserjrnyEditget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.Userjrnytask{
		Email:          a.Email,
		Id_journey:     a.Id_journey,
		Journey_name:   a.Journey_name,
		Initiate_state: a.Initiate_state,
		Journey_link:   a.Journey_link,
		Journey_index:  a.Journey_index,
		Status:         a.Status,
	}
	fmt.Println(a.Email)

	addUserjrny := Controller.ma.EditUserjrnys(ab)

	if addUserjrny {

		getadd := UserjrnyEditget{a.Email, a.Id_journey, a.Journey_name, a.Initiate_state, a.Journey_link, a.Journey_index, a.Status}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil update userjourney",
			Data:    getadd,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal update userjorney",
		}
		return c.JSON(http.StatusOK, res)

	}
}

// ViewUserjrnyDflow view userjourney by dataflow
func (Controller Controller) ViewUserjrnyDflow(c echo.Context) error {
	a := new(ViewUserjrnyDflowget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.UserjrnyDflowViewtask{
		Id_dataflow: a.Id_dataflow,
	}
	view := Controller.ma.ViewUserjrnyDflow(ab)

	res := responsegenr.ResponseGenericGet{
		Status:  "Success",
		Message: "Berhasil dapatkan data userjourney by dataflow ",
		Data:    view,
	}
	return c.JSON(http.StatusOK, res)

}

// ViewUserjrnyDflowStruct view userjourney by dataflow structure
func (Controller Controller) ViewUserjrnyDflowStruct(c echo.Context) error {
	a := new(ViewUserjrnyDflowStructget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.UserjrnyDflowStructViewtask{
		Email:                 a.Email,
		Id_dataflow_structure: a.Id_dataflow_structure,
	}
	view := Controller.ma.ViewUserjrnyDflowStruct(ab)

	res := responsegenr.ResponseGenericGet{
		Status:  "Success",
		Message: "Berhasil dapatkan data userjouney by dataflow structure ",
		Data:    view,
	}
	return c.JSON(http.StatusOK, res)

}

// ViewUserjrnyProj view userjourney by project
func (Controller Controller) ViewUserjrnyProj(c echo.Context) error {
	a := new(ViewUserjrnyProjget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.UserjrnyProjViewtask{
		Email:      a.Email,
		Id_project: a.Id_project,
	}
	view := Controller.ma.ViewUserjrnyProj(ab)

	res := responsegenr.ResponseGenericGet{
		Status:  "Success",
		Message: "Berhasil dapatkan data userjourney by project ",
		Data:    view,
	}
	return c.JSON(http.StatusOK, res)

}

// DeleteUserjrny delete timeline
func (Controller Controller) DeleteUserjrny(c echo.Context) error {
	a := new(DelUserjrny)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.UserjrnyDflowViewtask{
		Id: a.Id,
	}
	fmt.Println(a.Id)

	editUsecDesc := Controller.ma.DeleteUserjrny(ab.Id)

	if editUsecDesc {

		getdelete := DelUserjrny{a.Id}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil delete userjourney ",
			Data:    getdelete,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal delete userjourney ",
		}
		return c.JSON(http.StatusOK, res)

	}
}

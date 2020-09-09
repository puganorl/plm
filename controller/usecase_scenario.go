package controller

import (
	"fmt"
	"net/http"
	"plm/models"
	"plm/responsegenr"

	"github.com/labstack/echo"
)

type UsecScenget struct {
	Email           string `json:"email"`
	Id_usecase_desc int    `json:"id_usecase_desc"`
	Case_type       string `json:"case_type"`
	Initiate_state  string `json:"initiate_state"`
	Request         string `json:"request"`
	Response        string `json:"response"`
	Expectation     string `json:"expectation"`
	Valid           bool   `json:"valid"`
}

type UsecScenEditget struct {
	Email          string `json:"email"`
	Id             int    `json:"id"`
	Case_type      string `json:"case_type"`
	Initiate_state string `json:"initiate_state"`
	Request        string `json:"request"`
	Response       string `json:"response"`
	Expectation    string `json:"expectation"`
	Valid          bool   `json:"valid"`
	Status         string `json:"status"`
}

type ViewScenget struct {
	Email           string `json:"email"`
	Id_usecase_desc int    `json:"id_usecase_desc"`
}

type ViewScenprojget struct {
	Email      string `json:"email"`
	Id_project int    `json:"id_project"`
}

type DelScen struct {
	Id int `json:"id"`
}

//add scenario
func (Controller Controller) AddScenario(c echo.Context) error {
	a := new(UsecScenget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.UsecScentask{
		Email:           a.Email,
		Id_usecase_desc: a.Id_usecase_desc,
		Case_type:       a.Case_type,
		Initiate_state:  a.Initiate_state,
		Request:         a.Request,
		Response:        a.Response,
		Expectation:     a.Expectation,
		Valid:           a.Valid,
	}
	fmt.Println(a.Email)

	addScenario := Controller.ma.AddScenario(ab)

	if addScenario {

		getadd := UsecScenget{a.Email, a.Id_usecase_desc, a.Case_type, a.Initiate_state,
			a.Request, a.Response, a.Expectation, a.Valid}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil create scenario",
			Data:    getadd,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal create scenario",
		}
		return c.JSON(http.StatusOK, res)

	}
}

//edit scenario
func (Controller Controller) EditScenario(c echo.Context) error {
	a := new(UsecScenEditget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.UsecScentask{
		Email:          a.Email,
		Id:             a.Id,
		Case_type:      a.Case_type,
		Initiate_state: a.Initiate_state,
		Request:        a.Request,
		Response:       a.Response,
		Expectation:    a.Expectation,
		Valid:          a.Valid,
		Status:         a.Status,
	}
	fmt.Println(a.Email)

	editScenario := Controller.ma.EditScenario(ab)

	if editScenario {

		getedit := UsecScenEditget{a.Email, a.Id, a.Case_type, a.Initiate_state,
			a.Request, a.Response, a.Expectation, a.Valid, a.Status}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil edit scenario",
			Data:    getedit,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal edit scenario",
		}
		return c.JSON(http.StatusOK, res)

	}
}

//view scenario
func (Controller Controller) ViewScen(c echo.Context) error {
	a := new(ViewScenget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.ViewUsecScentask{
		Email:           a.Email,
		Id_usecase_desc: a.Id_usecase_desc,
	}
	view := Controller.ma.ViewScen(ab)

	res := responsegenr.ResponseGenericGet{
		Status:  "Success",
		Message: "Berhasil dapatkan data scenario",
		Data:    view,
	}
	return c.JSON(http.StatusOK, res)

}

//view scenario by project
func (Controller Controller) ViewScenProj(c echo.Context) error {
	a := new(ViewScenprojget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.ViewUsecScenProjtask{
		Email:      a.Email,
		Id_project: a.Id_project,
	}
	view := Controller.ma.ViewScenProj(ab)

	res := responsegenr.ResponseGenericGet{
		Status:  "Success",
		Message: "Berhasil dapatkan data scenario bye project",
		Data:    view,
	}
	return c.JSON(http.StatusOK, res)

}

// DeleteScenario delete timeline
func (Controller Controller) DeleteScenario(c echo.Context) error {
	a := new(DelScen)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.UsecScentask{
		Id: a.Id,
	}
	fmt.Println(a.Id)

	editUsecDesc := Controller.ma.DeleteScenario(ab.Id)

	if editUsecDesc {

		getdelete := DelScen{a.Id}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil delete scenario ",
			Data:    getdelete,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal delete scenario ",
		}
		return c.JSON(http.StatusOK, res)

	}
}

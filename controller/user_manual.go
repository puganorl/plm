package controller

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"plm/models"
	"plm/responsegenr"
)

type Umanualget struct {
	Email           string `json:"email"`
	Id_project      int    `json:"id_project"`
	Link_usermanual string `json:"link_usermanual"`
}

type UmanualEditget struct {
	Email           string `json:"email"`
	Id_project      int    `json:"id_project"`
	Id              int    `json:"id"`
	Link_usermanual string `json:"link_usermanual"`
	Status          string `json:"status"`
}

type ViewUmanualget struct {
	Email      string `json:"email"`
	Id_project int    `json:"id_project"`
}

//add UI UX
func (Controller Controller) AddUmanual(c echo.Context) error {
	a := new(Umanualget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.Umanualtask{
		Email:           a.Email,
		Id_project:      a.Id_project,
		Link_usermanual: a.Link_usermanual,
	}
	fmt.Println(a.Email)

	addSeqDiag := Controller.ma.AddUmanual(ab)

	if addSeqDiag {

		getadd := Umanualget{a.Email, a.Id_project, a.Link_usermanual}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil create usermanual",
			Data:    getadd,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal create usermanual",
		}
		return c.JSON(http.StatusOK, res)

	}
}

//edit UI UX
func (Controller Controller) EditUmanual(c echo.Context) error {
	a := new(UmanualEditget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.Umanualtask{
		Email:           a.Email,
		Id_project:      a.Id_project,
		Id:              a.Id,
		Link_usermanual: a.Link_usermanual,
		Status:          a.Status,
	}
	fmt.Println(a.Email)

	editSeqDiag := Controller.ma.EditUmanual(ab)

	if editSeqDiag {

		getadd := UmanualEditget{a.Email, a.Id_project, a.Id, a.Link_usermanual, a.Status}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil update usermanual",
			Data:    getadd,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal update UI UX",
		}
		return c.JSON(http.StatusOK, res)

	}
}

//view UI UX
func (Controller Controller) ViewUmanual(c echo.Context) error {
	a := new(ViewUmanualget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.ViewUmanualtask{
		Email:      a.Email,
		Id_project: a.Id_project,
	}
	view := Controller.ma.ViewUmanual(ab)

	res := responsegenr.ResponseGenericGet{
		Status:  "Success",
		Message: "Berhasil dapatkan data UI UX",
		Data:    view,
	}
	return c.JSON(http.StatusOK, res)

}

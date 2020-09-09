package controller

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"plm/models"
	"plm/responsegenr"
)

type Uiget struct {
	Email                 string `json:"email"`
	Id_journey            int    `json:"id_journey"`
	Id_dataflow_structure int    `json:"id_dataflow_structure"`
	Link_design           string `json:"link_design"`
	Description           string `json:"description"`
}

type UiEditget struct {
	Email                 string `json:"email"`
	Id_journey            int    `json:"id_journey"`
	Id_dataflow_structure int    `json:"id_dataflow_structure"`
	Link_design           string `json:"link_design"`
	Description           string `json:"description"`
	Status                string `json:"status"`
}

type Viewuiget struct {
	Email      string `json:"email"`
	Id_journey int    `json:"id_journey"`
}

//add UI UX
func (Controller Controller) AddUi(c echo.Context) error {
	a := new(Uiget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.Uitask{
		Email:                 a.Email,
		Id_journey:            a.Id_journey,
		Id_dataflow_structure: a.Id_dataflow_structure,
		Link_design:           a.Link_design,
		Description:           a.Description,
	}
	fmt.Println(a.Email)

	addSeqDiag := Controller.ma.AddUi(ab)

	if addSeqDiag {

		getadd := Uiget{a.Email, a.Id_journey, a.Id_dataflow_structure, a.Link_design, a.Description}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil create UI UX",
			Data:    getadd,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal create UI UX",
		}
		return c.JSON(http.StatusOK, res)

	}
}

//edit UI UX
func (Controller Controller) EditUi(c echo.Context) error {
	a := new(UiEditget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.Uitask{
		Email:                 a.Email,
		Id_journey:            a.Id_journey,
		Id_dataflow_structure: a.Id_dataflow_structure,
		Link_design:           a.Link_design,
		Description:           a.Description,
		Status:                a.Status,
	}
	fmt.Println(a.Email)

	editSeqDiag := Controller.ma.EditUi(ab)

	if editSeqDiag {

		getadd := UiEditget{a.Email, a.Id_journey, a.Id_dataflow_structure, a.Link_design, a.Description, a.Status}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil update UI UX",
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
func (Controller Controller) ViewUi(c echo.Context) error {
	a := new(Viewuiget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.ViewUitask{
		Email:      a.Email,
		Id_journey: a.Id_journey,
	}
	view := Controller.ma.ViewUi(ab)

	res := responsegenr.ResponseGenericGet{
		Status:  "Success",
		Message: "Berhasil dapatkan data UI UX",
		Data:    view,
	}
	return c.JSON(http.StatusOK, res)

}

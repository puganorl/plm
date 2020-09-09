package controller

import (
	"fmt"
	"net/http"
	"plm/models"
	"plm/responsegenr"

	"github.com/labstack/echo"
)

type DataFstrucget struct {
	Email         string `json:"email"`
	Id_dataflow   int    `json:"id_dataflow"`
	Index         int    `json:"index"`
	Dataflow_name string `json:"dataflow_name"`
	Protocol      int    `json:"protocol"`
	Type          int    `json:"type"`
	Address       string `json:"address"`
	Request_type  string `json:"request_type"`
	Response_type string `json:"response_type"`
}

type EditDataFstrucget struct {
	Id            int    `json:"id"`
	Email         string `json:"email"`
	Id_dataflow   int    `json:"id_dataflow"`
	Index         int    `json:"index"`
	Dataflow_name string `json:"dataflow_name"`
	Protocol      int    `json:"protocol"`
	Type          int    `json:"type"`
	Address       string `json:"address"`
	Request_type  string `json:"request_type"`
	Response_type string `json:"response_type"`
	Status        string `json:"status"`
}

type ViewDflowget struct {
	Email       string `json:"email"`
	Id_dataflow int    `json:"id_dataflow"`
}

type ViewDflowSequenceget struct {
	Email       string `json:"email"`
	Id_sequence int    `json:"id_sequence"`
}

type ViewDflowDesignget struct {
	Email     string `json:"email"`
	Id_design int    `json:"id_design"`
}

type ViewDflowStructProjget struct {
	Email      string `json:"email"`
	Id_project int    `json:"id_project"`
}

type DelDataFStruc struct {
	Id int `json:"id"`
}

//add dataflow_sctructure
func (Controller Controller) AddDataFStruc(c echo.Context) error {
	a := new(DataFstrucget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.DataFstructask{
		Email:         a.Email,
		Id_dataflow:   a.Id_dataflow,
		Index:         a.Index,
		Dataflow_name: a.Dataflow_name,
		Protocol:      a.Protocol,
		Type:          a.Type,
		Address:       a.Address,
		Request_type:  a.Request_type,
		Response_type: a.Response_type,
	}
	fmt.Println(a.Email)

	addDataFStruc := Controller.ma.AddDataFStrucs(ab)

	if addDataFStruc {

		getadd := DataFstrucget{a.Email, a.Id_dataflow, a.Index, a.Dataflow_name, a.Protocol, a.Type,
			a.Address, a.Request_type, a.Response_type}
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

//edit dataflow_sctructure
func (Controller Controller) EditDataFStruc(c echo.Context) error {
	a := new(EditDataFstrucget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.DataFstructask{
		Id:            a.Id,
		Email:         a.Email,
		Id_dataflow:   a.Id_dataflow,
		Index:         a.Index,
		Dataflow_name: a.Dataflow_name,
		Protocol:      a.Protocol,
		Type:          a.Type,
		Address:       a.Address,
		Request_type:  a.Request_type,
		Response_type: a.Response_type,
		Status:        a.Status,
	}
	fmt.Println(a.Email)

	addDataFStruc := Controller.ma.EditDataFStrucs(ab)

	if addDataFStruc {

		getedit := EditDataFstrucget{a.Id, a.Email, a.Id_dataflow, a.Index, a.Dataflow_name, a.Protocol,
			a.Type, a.Address, a.Request_type, a.Response_type, a.Status}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil update dataflow structure",
			Data:    getedit,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal update dataflow structure",
		}
		return c.JSON(http.StatusOK, res)

	}
}

//view  dataflow structure by dataflow
func (Controller Controller) ViewDflow(c echo.Context) error {
	a := new(ViewDflowget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.DataflowView{
		Id_dataflow: a.Id_dataflow,
	}
	view := Controller.ma.ViewDflow(ab)

	res := responsegenr.ResponseGenericGet{
		Status:  "Success",
		Message: "Berhasil dapatkan data dataflow structure by dataflow ",
		Data:    view,
	}
	return c.JSON(http.StatusOK, res)

}

//view  dataflow structure by sequence
func (Controller Controller) ViewDflowSequence(c echo.Context) error {
	a := new(ViewDflowSequenceget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.DataflowSequnceViewtask{
		Email:       a.Email,
		Id_sequence: a.Id_sequence,
	}
	view := Controller.ma.ViewDflowSequence(ab)

	res := responsegenr.ResponseGenericGet{
		Status:  "Success",
		Message: "Berhasil dapatkan data dataflow structure by sequence ",
		Data:    view,
	}
	return c.JSON(http.StatusOK, res)

}

//view  dataflow structure by design
func (Controller Controller) ViewDflowDesign(c echo.Context) error {
	a := new(ViewDflowDesignget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.DataflowDesignViewtask{
		Email:     a.Email,
		Id_design: a.Id_design,
	}
	view := Controller.ma.ViewDflowDesign(ab)

	res := responsegenr.ResponseGenericGet{
		Status:  "Success",
		Message: "Berhasil dapatkan data dataflow structure by design ",
		Data:    view,
	}
	return c.JSON(http.StatusOK, res)

}

//view  dataflow structure by project
func (Controller Controller) ViewDflowStructProj(c echo.Context) error {
	a := new(ViewDflowStructProjget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.DataflowProjViewtask{
		Email:      a.Email,
		Id_project: a.Id_project,
	}
	view := Controller.ma.ViewDflowStructProj(ab)

	res := responsegenr.ResponseGenericGet{
		Status:  "Success",
		Message: "Berhasil dapatkan data dataflow structure by project ",
		Data:    view,
	}
	return c.JSON(http.StatusOK, res)

}

// DeleteDataFStruc delete timeline
func (Controller Controller) DeleteDataFStruc(c echo.Context) error {
	a := new(DelDataFStruc)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.DataFstructask{
		Id: a.Id,
	}
	fmt.Println(a.Id)

	editUsecDesc := Controller.ma.DeleteDataFStruc(ab.Id)

	if editUsecDesc {

		getdelete := DelDataFStruc{a.Id}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil delete dataflow structure",
			Data:    getdelete,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal delete dataflow structure",
		}
		return c.JSON(http.StatusOK, res)

	}
}

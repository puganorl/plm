package controller

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"plm/models"
	"plm/responsegenr"
)

type Memberget struct {
	Email      string `json:"email"`
	Id_member  int 		`json:"id_member"`
	Role       string `json:"role"`
	Id_project int    `json:"id_project"`
}

type Memberadd struct {
	Email      string `json:"email"`
	Id_user    string `json:"id_user"`
	Role       int	  `json:"role"`
	Id_project int    `json:"id_project"`
}

type DelMemberget struct {
	Email      string `json:"email"`
	Id_member  int    `json:"id_member"`
	Id_project int    `json:"id_project"`
}
type ViewMemberget struct {
	Email      string `json:"email"`
	Id_project int    `json:"id_project"`
}

//add member
func (Controller Controller) AddMember(c echo.Context) error {
	a := new(Memberadd)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.Membertask{
		Email:      a.Email,
		Id_user:  	a.Id_user,
		Role:       a.Role,
		Id_project: a.Id_project,
	}
	fmt.Println(a.Email)

	addMember := Controller.ma.AddMember(ab)

	if addMember {

		getadd := Memberadd{a.Email, a.Id_user, a.Role, a.Id_project}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil create member",
			Data:    getadd,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal create member",
		}
		return c.JSON(http.StatusOK, res)

	}
}

// view member
func (Controller Controller) ViewMember(c echo.Context) error {
	a := new(ViewMemberget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.MemberViewtask{
		Id_project: a.Id_project,
	}
	view := Controller.ma.ViewMember(ab)

	res := responsegenr.ResponseGenericGet{
		Status:  "Success",
		Message: "Berhasil dapatkan data member",
		Data:    view,
	}
	return c.JSON(http.StatusOK, res)

}

// view member not in project
func (Controller Controller) ViewMemberNotInProject(c echo.Context) error {
	a := new(ViewMemberget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.MemberNotInProjectViewtask{
		Id_project: a.Id_project,
	}
	view := Controller.ma.ViewMemberNotInProject(ab)

	res := responsegenr.ResponseGenericGet{
		Status:  "Success",
		Message: "Berhasil dapatkan data member",
		Data:    view,
	}
	return c.JSON(http.StatusOK, res)

}

//delete member
func (Controller Controller) DelMember(c echo.Context) error {
	a := new(DelMemberget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.Memberdel{
		Email:      a.Email,
		Id_member:  a.Id_member,
		Id_project: a.Id_project,
	}
	fmt.Println(a.Email)

	delMember := Controller.ma.DelMember(ab.Id_member)

	if delMember {

		getadd := DelMemberget{a.Email, a.Id_member, a.Id_project}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil hapus member",
			Data:    getadd,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal hapus member",
		}
		return c.JSON(http.StatusOK, res)

	}
}

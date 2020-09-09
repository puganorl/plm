package controller

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"plm/models"
	"plm/responsegenr"
)

type TaskProject struct {
	Id           int    `json:"id"`
	Id_project   int    `json:"id_project"`
	Email        string `json:"email"`
	Project_name string `json:"project_name"`
	Deskripsi    string `json:"deskripsi"`
	//Id_role      int    `json:"id_role"`
}

type EditGet struct {
	Email        string `json:"email"`
	Project_name string `json:"project_name"`
	Deskripsi    string `json:"deskripsi"`
	Id_project   int    `json:"id_project"`
}

type CreateGets struct {
	Id_project   int    `json:"id_project"`
	Project_name string `json:"project_name"`
	Deskripsi    string `json:"deskripsi"`
	//Id_role      int    `json:"id_role"`
}

type ProjectViewget struct {
	Email      string `json:"email"`
	Id_project int    `json:"id_project"`
}

type ProjectAllwget struct {
	Email      string `json:"email"`
}

type Block struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

type Exception interface{}

func Throw(up Exception) {
	panic(up)
}

func (tcf Block) Do() {
	if tcf.Finally != nil {

		defer tcf.Finally()
	}
	if tcf.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.Catch(r)
			}
		}()
	}
	tcf.Try()
}

//create
func (Controller Controller) PostsCreate(c echo.Context) error {
	a := new(TaskProject)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.TaskProject{
		Email:        a.Email,
		Project_name: a.Project_name,
		Deskripsi:    a.Deskripsi,
		//Id_role:      a.Id_role,
	}

	fmt.Println(a.Email)

	insertTbl_project := Controller.ma.InsertTbl_Project(ab)
	GetMaxId := Controller.ma.MaxIdProject(ab)
	//cekIdProject := Controller.ma.MaxIdProject(ab)

	Block{
		Try: func() {
			Controller.ma.InsertTbl_member_belongto_project(ab, GetMaxId)
			Controller.ma.InsertTbl_role_belongto_member_project(ab, GetMaxId)
		},
		Catch: func(e Exception) {
			fmt.Printf("Caught %v\n", e)
		},
	}.Do()

	fmt.Println(GetMaxId)

	if insertTbl_project {

		getCreate := CreateGets{Id_project: GetMaxId, Project_name: a.Project_name, Deskripsi: a.Deskripsi}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil create project",
			Data:    getCreate,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal create project",
		}
		return c.JSON(http.StatusOK, res)

	}
	//} else {
	//
	//	res := responsegenr.ResponseGeneric{
	//		Status:  "Error",
	//		Message: "Id Project tidak ditemukan",
	//	}
	//	return c.JSON(http.StatusOK, res)
	//
	//}
}

//edit
func (Controller Controller) EditProject(c echo.Context) error {
	a := new(TaskProject)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.TaskProject{
		Email:        a.Email,
		Project_name: a.Project_name,
		Deskripsi:    a.Deskripsi,
		Id_project:   a.Id_project,
	}

	cekId_Project := Controller.ma.CekId_Project(ab)
	editTbl_Project := Controller.ma.EditTbl_Project(ab)

	if cekId_Project {

		if editTbl_Project {

			getCreate := EditGet{a.Email, a.Project_name, a.Deskripsi, a.Id_project}
			res := responsegenr.ResponseGenericGet{
				Status:  "Success",
				Message: "Berhasil update project",
				Data:    getCreate,
			}
			return c.JSON(http.StatusOK, res)

		} else {
			res := responsegenr.ResponseGeneric{
				Status:  "Error",
				Message: "Gagal update project",
			}
			return c.JSON(http.StatusOK, res)
		}
	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Id_project tidak tersedia",
		}
		return c.JSON(http.StatusOK, res)

	}
}

//view
func (Controller Controller) ViewProject(c echo.Context) error {
	a := new(ProjectViewget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.TaskProjectView{
		Email:      a.Email,
		Id_project: a.Id_project,
	}
	view := Controller.ma.ViewProject(ab)

	res := responsegenr.ResponseGenericGet{
		Status:  "Success",
		Message: "Berhasil dapatkan data project",
		Data:    view,
	}
	return c.JSON(http.StatusOK, res)

}


// all projects
func (Controller Controller) AllProjects(c echo.Context) error {
	a := new(ProjectAllwget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.TaskProjectAll{
		Email:      a.Email,
	}
	view := Controller.ma.AllProjects(ab)

	res := responsegenr.ResponseGenericGet{
		Status:  "Success",
		Message: "Berhasil dapatkan data project",
		Data:    view,
	}
	return c.JSON(http.StatusOK, res)

}

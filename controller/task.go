package controller

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"plm/models"
	"plm/responsegenr"
)

type Controller struct {
	ma models.Models
}

// Interface
type ExampleControllerInterface interface {
	GetPostsController(c echo.Context) error
}

// Interface
type ControllerInsertDB interface {
	GetPostsDB(c echo.Context) error
}

type Task struct {
	Id           int    `json:"id"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	Password     string `json:"password"`
	Old_password string `json:"old_password"`
}
type LoginTask struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type RegisTask struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
type LoginGet struct {
	Email string `json:"email"`
	Nama  string `json:"nama"`
}

//registration
func (Controller Controller) PostsRegis(c echo.Context) error {
	a := new(RegisTask)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.Task{
		Email:    a.Email,
		Name:     a.Name,
		Password: a.Password,
	}

	fmt.Println(a.Email)

	cekEmail := Controller.ma.CekEmailUser(ab.Email)
	insertUser := Controller.ma.InsertTbl_User(ab)
	insertUserAuth := Controller.ma.InsertTbl_UserAuth(ab)

	if cekEmail {

		if insertUser && insertUserAuth {
			getRegis := LoginGet{a.Email, a.Name}
			res := responsegenr.ResponseGenericGet{
				Status:  "Success",
				Message: "Berhasil Input Data",
				Data:    getRegis,
			}
			return c.JSON(http.StatusOK, res)
		} else {
			res := responsegenr.ResponseGeneric{
				Status:  "Error",
				Message: "Gagal Input User",
			}
			return c.JSON(http.StatusOK, res)
		}

	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Email Telah Digunakan",
		}
		return c.JSON(http.StatusOK, res)

	}
}

//edit
func (Controller Controller) UpEdit(c echo.Context) error {
	a := new(Task)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.Task{
		Email:        a.Email,
		Name:         a.Name,
		Password:     a.Password,
		Old_password: a.Old_password,
	}

	EditTbl_user := Controller.ma.EditTbl_user(ab)
	EditTbl_user_auth := Controller.ma.EditTbl_user_auth(ab)

	if EditTbl_user && EditTbl_user_auth {

		getUpdate := LoginGet{a.Email, a.Name}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil update data",
			Data:    getUpdate,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal Update Data",
		}
		return c.JSON(http.StatusOK, res)

	}
}

//login
func (Controller Controller) GetLogin(c echo.Context) error {
	id := new(LoginTask)
	if err := c.Bind(id); err != nil {
		return err
	}

	ab := models.Task{
		Email:    id.Email,
		Password: id.Password,
	}
	cekLogin := Controller.ma.Ceklogin(ab.Email, ab.Password)
	positions := Controller.ma.GetPositionUserLogin(ab)
	posts := Controller.ma.LoginTask(ab, positions.Position)

	if cekLogin {
		if posts.Status {
			res := responsegenr.ResponseGenericGet{
				Status:  "Success",
				Message: "Login berhasil",
				Data:    posts.ResLogin,
			}
			return c.JSON(http.StatusOK, res)

		} else {

			res := responsegenr.ResponseGenericGet{
				Status:  "Error",
				Message: "Login Gagal ",
			}
			return c.JSON(http.StatusOK, res)
		}
	}else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Salah email/password ",
		}
		return c.JSON(http.StatusOK, res)

	}
}

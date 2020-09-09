package controller

import (
	"fmt"
	"net/http"
	"plm/models"
	"plm/responsegenr"
	"github.com/labstack/echo"
)

type Archget struct {
	Id_project int    `json:"id_project"`
	Link_arch  string `json:"link_arch"`
	Email      string `json:"email"`
}

type ArchEditget struct {
	Email        string `json:"email"`
	Id_arch_diag int    `json:"id_arch_diag"`
	Link_arch    string `json:"link_arch"`
	Status       string `json:"status"`
}

type ArchDesget struct {
	Email        string `json:"email"`
	Id_arch_diag int    `json:"id_arch_diag"`
	Description  string `json:"description"`
	Desc_index   string `json:"desc_index"`
}

type ArchEditDesget struct {
	Email        string `json:"email"`
	Id           int    `json:"id"`
	Id_arch_diag int    `json:"id_arch_diag"`
	Description  string `json:"description"`
	Desc_index   string `json:"desc_index"`
}

type ArchDelDesget struct {
	Email string `json:"email"`
	Id    int    `json:"id"`
	Index    int    `json:"index"`
}

type ViewArchget struct {
	Email      string `json:"email"`
	Id_project int    `json:"id_project"`
}

//add arch
func (Controller Controller) AddArch(c echo.Context) error {
	a := new(Archget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.Archtask{
		Email:      a.Email,
		Id_project: a.Id_project,
		Link_arch:  a.Link_arch,
	}
	fmt.Println(a.Email)

	addArch := Controller.ma.AddArch(ab)

	if addArch {

		getadd := Archget{a.Id_project, a.Link_arch, a.Email}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil create arch",
			Data:    getadd,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal create arch",
		}
		return c.JSON(http.StatusOK, res)

	}
}

//edit arch
func (Controller Controller) EditArch(c echo.Context) error {
	a := new(ArchEditget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.Archtask{
		Email:        a.Email,
		Id_arch_diag: a.Id_arch_diag,
		Link_arch:    a.Link_arch,
		Status:       a.Status,
	}
	fmt.Println(a.Email)

	editArch := Controller.ma.EditArch(ab)

	if editArch {

		getEdit := ArchEditget{a.Email, a.Id_arch_diag, a.Link_arch, a.Status}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil update arch",
			Data:    getEdit,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal update arch",
		}
		return c.JSON(http.StatusOK, res)

	}
}

//add desc arch
func (Controller Controller) AddArchDes(c echo.Context) error {
	a := new(ArchDesget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.ArchDestask{
		Email:        a.Email,
		Id_arch_diag: a.Id_arch_diag,
		Description:  a.Description,
		Desc_index:   a.Desc_index,
	}
	fmt.Println(a.Email)

	addArchdes := Controller.ma.AddArchDesc(ab)

	if addArchdes {

		getadd := ArchDesget{a.Email, a.Id_arch_diag, a.Description, a.Desc_index}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil create arch description",
			Data:    getadd,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal create arch description",
		}
		return c.JSON(http.StatusOK, res)

	}
}

//edit desc arch
func (Controller Controller) EditArchDes(c echo.Context) error {
	a := new(ArchEditDesget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.ArchDestask{
		Email:        a.Email,
		Id:           a.Id,
		Id_arch_diag: a.Id_arch_diag,
		Description:  a.Description,
		Desc_index:   a.Desc_index,
	}
	fmt.Println(a.Email)

	editArch := Controller.ma.EditArchDesc(ab)

	if editArch {

		getEditDes := ArchEditDesget{a.Email, a.Id, a.Id_arch_diag, a.Description, a.Desc_index}

		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil update arch description",
			Data:    getEditDes,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal update arch description",
		}
		return c.JSON(http.StatusOK, res)

	}
}

//view arch
func (Controller Controller) ViewArch(c echo.Context) error {
	a := new(ViewArchget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.ArchitectureView{
		Id_project: a.Id_project,
	}
	view := Controller.ma.ViewArch(ab)

	res := responsegenr.ResponseGenericGet{
		Status:  "Success",
		Message: "Berhasil dapatkan data architecture",
		Data:    view,
	}
	return c.JSON(http.StatusOK, res)

}

//delete desc arch
func (Controller Controller) DelArchDes(c echo.Context) error {
	a := new(ArchDelDesget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.ArchDestask{
		Email: a.Email,
		Id:    a.Id,
		Index: a.Index,
	}
	fmt.Println(a.Email)

	delArchdes := Controller.ma.DelArchDesc(ab.Id, ab.Index)
	fmt.Println(ab.Id)
	fmt.Println(ab.Index)

	if delArchdes {

		getdel := ArchDelDesget{a.Email, a.Id, a.Index}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil hapus arch description",
			Data:    getdel,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal hapus arch description",
		}
		return c.JSON(http.StatusOK, res)

	}
}

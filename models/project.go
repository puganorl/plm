package models

import (
	"fmt"
)

type TaskProject struct {
	Name         string `json:"name"`
	Id_project   int    `json:"id_project"`
	Email        string `json:"email"`
	Project_name string `json:"project_name"`
	Deskripsi    string `json:"deskripsi"`
	Id_role      int    `json:"id_role"`
}

type TaskProjectView struct {
	Email        string `json:"email"`
	Id_project   int    `json:"id_project"`
	Project_name string `json:"project_name"`
	Deskripsi    string `json:"deskripsi"`
	Create_date  string `json:"create_date"`
	Id_member    int 	`json:"id_member"`
}

type TaskProjectAll struct {
	Email				 string `json:"email"`
	Id_project   int    `json:"id_project"`
	Name 				 string `json:"project_name"`
	Deskripsi    string `json:"deskripsi"`
	Create_date  string `json:"create_date"`
}

type Taskcekemail struct {
	Email        string `json:"email"`
	Id_project   int    `json:"id_project"`
	Project_name string `json:"project_name"`
	Deskripsi    string `json:"deskripsi"`
	Create_date  string `json:"create_date"`
}

type ProjectRes struct {
	Id_project   string `json:"id_project"`
	Project_name string `json:"project_name"`
	Deskripsi    string `json:"deskripsi"`
}

type ProjectsView struct {
	ProjectView []TaskProjectView `json:"project_view"`
}

type ProjectsAll struct {
	ProjectView []TaskProjectAll `json:"project_all"`
}

type ProjectTask struct {
	Id_project   string `json:"id_project"`
	Project_name string `json:"project_name"`
	Creator      string `json:"creator"`
	Deskripsi    string `json:"deskripsi"`
}
type CekId struct {
	Max int `json:"max"`
}

//create
func (ExampleModel Models) InsertTbl_Project(Create TaskProject) bool {

	sqlStatement2 := "INSERT INTO tbl_project (name, creator, description, create_date) " +
		"VALUES ($1,$2 ,$3, now()::timestamp)"
	res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
		Create.Project_name,
		Create.Email,
		Create.Deskripsi,
	)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	if err2 != nil {
		fmt.Println(err2)
		return false
	} else {
		fmt.Println(res2)
		return true
	}
}
func (ExampleModel Models) MaxIdProject(Create TaskProject) int {

	//var MaxId int
	//MaxId = Create.Id_project

	Maxx := CekId{}
	sqlStatement3 := " SELECT  max(id) FROM tbl_project "
	res3, err3 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement3)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Println(res3)
	}

	//task := TaskProject{}
	for res3.Next() {
		err := res3.Scan(&Maxx.Max)
		// Exit if we get an error
		if err != nil {
			fmt.Println(err)
		}

	}
	fmt.Println(Maxx)

	return Maxx.Max
}
func (ExampleModel Models) InsertTbl_member_belongto_project(Create TaskProject, MaxId int) bool {
	sqlStatement := "INSERT INTO tbl_member_belongto_project (id_user, id_project, creator, status, create_date) " +
		"VALUES ($1,$2 ,$3, $4, now()::timestamp)"
	res, err := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement,
		Create.Email,
		MaxId,
		Create.Email,
		true,
	)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	fmt.Println(MaxId)
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		fmt.Println(res)
		return true
	}

}
func (ExampleModel Models) InsertTbl_role_belongto_member_project(Create TaskProject, MaxId int) bool {
	sqlStatement := "INSERT INTO tbl_role_belongto_member_project (id_role, id_member_project, creator, create_date) " +
		"VALUES ($1,$2 ,$3, now()::timestamp)"
	res, err := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement,
		1,
		MaxId,
		Create.Email,
	)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	fmt.Println(Create.Id_role)
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		fmt.Println(res)
		return true
	}

}

//edit
func (ExampleModel Models) CekId_Project(Edit TaskProject) bool {
	//cek id_project
	sqlStatement3 := "SELECT tbl_project.creator FROM tbl_project " +
		"WHERE tbl_project.id =$1"
	res3, err3 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement3,
		Edit.Id_project,
	)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	if err3 != nil {
		fmt.Println(err3)

	} else {
		fmt.Println(res3)
	}
	task := ProjectTask{}
	for res3.Next() {
		err3 := res3.Scan(&task.Id_project)
		// Exit if we get an error
		if err3 != nil {
			fmt.Println(err3)
		}
	}

	if task.Id_project != task.Id_project {
		return false
	}
	return true
}

func (ExampleModel Models) EditTbl_Project(Edit TaskProject) bool {

	sqlStatement := " UPDATE tbl_project " +
		"SET  name = $3, description = $4 " +
		"WHERE tbl_project.id =$1 AND tbl_project.creator =$2"
	res, err := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement,
		Edit.Id_project,
		Edit.Email,
		Edit.Project_name,
		Edit.Deskripsi,
	)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		fmt.Println(res)
		return true
	}
}

//view
func (ExampleModel Models) ViewProject(View TaskProjectView) ProjectsView {

	sqlStatement3 := "SELECT *  FROM tbl_project " +
		"WHERE tbl_project.id=$1 "
	res3, err3 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement3,
		View.Id_project,
	)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	if err3 != nil {
		fmt.Println(err3)

	} else {
		fmt.Println(res3)
	}
	result := ProjectsView{}

	for res3.Next() {
		task := TaskProjectView{}
		err3 := res3.Scan(&task.Id_project, &task.Project_name, &task.Email, &task.Deskripsi, &task.Id_member)
		// Exit if we get an error
		if err3 != nil {
			fmt.Println(err3)

		}
		result.ProjectView = append(result.ProjectView, task)
	}
	return result
}
//all
func (ExampleModel Models) AllProjects(View TaskProjectAll) ProjectsAll {
	sqlStatement3 := "SELECT tbl_project.id as id_project, tbl_project.name, tbl_project.description as deskripsi, tbl_project.create_date FROM tbl_project " +
		"INNER JOIN tbl_member_belongto_project ON tbl_project.id =  tbl_member_belongto_project.id_project " +
		"INNER JOIN tbl_role_belongto_member_project ON tbl_member_belongto_project.id_project = tbl_role_belongto_member_project.id_member_project " +
		"WHERE tbl_member_belongto_project.id_user = $1 OR tbl_project.creator = $1 " +
		"GROUP BY tbl_project.id"
	res3, err3 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement3,
		View.Email,
	)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	if err3 != nil {
		fmt.Println(err3)

	} else {
		fmt.Println(res3)
	}
	result := ProjectsAll{}

	for res3.Next() {
		task := TaskProjectAll{}
		err3 := res3.Scan(&task.Id_project, &task.Name, &task.Deskripsi, &task.Create_date)
		// Exit if we get an error
		if err3 != nil {
			fmt.Println(err3)

		}
		result.ProjectView = append(result.ProjectView, task)
	}

	return result

}

package models

import (
	"fmt"
	"plm/settingdb"
)

type Task struct {
	Id_user      string `json:"id_user"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	Password     string `json:"password"`
	Old_password string `json:"old_password"`
}

type LoginTask struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type UserTask struct {
	Email    string `json:"email"`
	Position string `json:"position"`
	Name     string `json:"name"`
}

type CekTask struct {
	Id_user  string `json:"id_user"`
	Password string `json:"password"`
}

type LoginTask1 struct {
	Status   bool           `json:"status"`
	ResLogin ProjectsLogin1 `json:"resLogin"`
}

type ProjectsLogin1 struct {
	Email    string          `json:"email"`
	Projects []ProjectsLogin `json:"projects"`
}

type ProjectsLogin struct {
	Position string       `json:"position"`
	Project  ProjectLogin `json:"project"`
}

type ProjectLogin struct {
	Id_project   int    `json:"id_project"`
	Project_name string `json:"project_name"`
	Description  string `json:"description"`
	Id_role      int    `json:"id_role"`
}

type Models struct {
	db settingdb.DatabaseConfig
}

//registration
func (ExampleModel Models) InsertTbl_UserAuth(Regis Task) bool {

	sqlStatement3 := "INSERT INTO tbl_user_auth (id_user, id_authentication, password, id_method, status, \"limit\", create_date) " +
		"VALUES ($1,$2 ,$3, $4, $5,now()::time, now()::timestamp)"
	res3, err3 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement3,
		Regis.Email,
		'-',
		Regis.Password,
		1,
		true,
	)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	if err3 != nil {
		fmt.Println(err3)
		return false
	} else {
		fmt.Println(res3)
		return true
	}
}

func (ExampleModel Models) InsertTbl_User(Regis Task) bool {
	sqlStatement := "INSERT INTO tbl_user (email, name, position) " +
		"VALUES ($1, $2, $3)"
	res, err := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement,
		Regis.Email,
		Regis.Name,
		"",
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

func (ExampleModel Models) CekEmailUser(Email string) bool {
	// CEK EMAIL TERDAFTAR
	sqlStatement2 := "SELECT * FROM tbl_user " +
		"WHERE email=$1"
	res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
		Email,
	)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	if err2 != nil {
		fmt.Println(res2)
	} else {
		fmt.Println(err2)
	}

	task := UserTask{}
	for res2.Next() {
		err2 := res2.Scan(&task.Email, &task.Name, &task.Position)
		// Exit if we get an error
		if err2 != nil {
			fmt.Println(err2)
		}
	}

	if task.Email != "" {
		return false
	}
	return true
}

//edit
func (ExampleModel Models) EditTbl_user(Edit Task) bool {

	sqlStatement := " UPDATE tbl_user " +
		" SET name = $1 " +
		" WHERE email = $2"
	res, err := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement,
		Edit.Name,
		Edit.Email,
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

func (ExampleModel Models) EditTbl_user_auth(Edit Task) bool {
	//defer ExampleModel.db.GetDatabaseConfig().Close()
	sqlStatement3 := "UPDATE tbl_user_auth " +
		"SET id_authentication = $3, password = $1, id_method = 1, status = $4, \"limit\" = now()::time, create_date = now()::timestamp " +
		"WHERE id_user = $2"

	res3, err3 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement3,
		Edit.Password,
		Edit.Email,
		'-',
		true,
	)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	if err3 != nil {
		fmt.Println(err3)
		return false
	} else {
		fmt.Println(res3)
		return true
	}
}

//login
func (ExampleModel Models) Ceklogin(Email string, Password string) bool {
	// CEK EMAIL TERDAFTAR
	sqlStatement2 := "SELECT tbl_user_auth.id_user, tbl_user_auth.password FROM tbl_user_auth " +
		"WHERE id_user=$1 AND password=$2 "
	res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
		Email,
		Password,
	)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	if err2 != nil {
		fmt.Println(res2)
	} else {
		fmt.Println(err2)
	}

	task := CekTask{}
	for res2.Next() {
		err2 := res2.Scan(&task.Id_user, &task.Password)
		// Exit if we get an error
		if err2 != nil {
			fmt.Println(err2)
		}
	}

	if task.Id_user != Email {
		return false
	} else if task.Password != Password {
		return false
	}
	return true
}
func (ExampleModel Models) GetPositionUserLogin(Login Task) UserTask {

	task := UserTask{}
	sqlStatement := "SELECT  tbl_user.position FROM tbl_user " +
		"INNER JOIN tbl_user_auth ON tbl_user.email = tbl_user_auth.id_user " +
		"WHERE tbl_user.email=$1 AND tbl_user_auth.password=$2 "
	res, err := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement,
		Login.Email,
		Login.Password,
	)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	if err != nil {
		fmt.Println(err)
		return task
	} else {
		fmt.Println(res)
	}

	for res.Next() {
		err := res.Scan(&task.Position)
		// Exit if we get an error
		if err != nil {
			fmt.Println(err)
		}
	}
	return task

}

func (ExampleModel Models) LoginTask(Login Task, Positions string) LoginTask1 {
	getlogin := LoginTask1{}

	sqlStatement2 := "SELECT  tbl_project.id,tbl_project.name,tbl_project.description,tbl_role_belongto_member_project.id_role FROM tbl_project " +
		"INNER JOIN tbl_member_belongto_project ON tbl_project.id =  tbl_member_belongto_project.id_project " +
		"INNER JOIN tbl_role_belongto_member_project ON tbl_member_belongto_project.id_project = tbl_role_belongto_member_project.id_member_project " +
		"WHERE tbl_member_belongto_project.id_user = $1 OR tbl_project.creator = $1 " //+
	//"GROUP BY tbl_project.id"

	res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
		Login.Email,
	)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	if err2 != nil {
		fmt.Println(err2)

	} else {
		fmt.Println(res2)
	}

	loginResponse := ProjectsLogin1{}

	for res2.Next() {
		projects := ProjectsLogin{}
		err2 := res2.Scan(&projects.Project.Id_project, &projects.Project.Project_name, &projects.Project.Description, &projects.Project.Id_role)
		// Exit if we get an error
		if err2 != nil {
			fmt.Println(err2)
		}
		projects.Position = Positions
		loginResponse.Projects = append(loginResponse.Projects, projects)
	}

	loginResponse.Email = Login.Email

	getlogin.Status = true
	getlogin.ResLogin = loginResponse
	return getlogin
}

package models

import "fmt"

type Membertask struct {
	Email      string `json:"email"`
	Id_user    string `json:"id_member"`
	Role       int    `json:"role"`
	Id_project int    `json:"id_project"`
}

type Memberdel struct {
	Email      string `json:"email"`
	Id_member  int    `json:"id_member"`
	Id_project int    `json:"id_project"`
}

type MemberViewtask struct {
	Id_project  int    `json:"id_project"`
	Name        string `json:"name"`
	Position    string `json:"position"`
}

type MemberNotInProjectViewtask struct {
	Id_project int    `json:"id_project"`
	Email      string `json:"email"`
	Name       string `json:"name"`
}

type MemberView struct {
	MemberView []MemberViewtask `json:"MemberView"`
}

type MemberNotInProjectView struct {
	MemberNotInProjectView []MemberNotInProjectViewtask `json:"MemberNotInProjectView"`
}

//add member
func (ExampleModel Models) AddMember(Add Membertask) bool {

	sqlStatement2 := "INSERT INTO  tbl_member_belongto_project(id_user, id_project, creator, status, create_date) " +
		"VALUES ($1,$2 ,$3, $4, now()::timestamp)"
	res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
		Add.Id_user,
		Add.Id_project,
		Add.Email,
		true,
	)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	if err2 != nil {
		fmt.Println(err2)
		return false
	} else {
		fmt.Println(res2)
		rows, err := ExampleModel.db.GetDatabaseConfig().Query("SELECT id FROM tbl_member_belongto_project ORDER BY create_date DESC LIMIT 1")
		if err != nil {
			fmt.Println(err)
		}
		defer rows.Close()

		var (
			id          int
			inserted_id int
		)

		for rows.Next() {
			err := rows.Scan(&id)
			if err != nil {
				fmt.Println(err)
			}
			inserted_id = id
		}

		sqlStatement2 := "INSERT INTO  tbl_role_belongto_member_project(id_role, id_member_project, creator, create_date) " +
			"VALUES ($1, $2, $3, now()::timestamp)"
		res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
			Add.Role,
			inserted_id,
			Add.Email,
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
}

//view member
func (ExampleModel Models) ViewMember(View MemberViewtask) MemberView {

	sqlStatement3 := "SELECT tbl_member_belongto_project.id AS id_member, tbl_role.rolename as rolename, tbl_role_belongto_member_project.id, id_user, id_project, tbl_user.name as name, tbl_user.position as position, tbl_role_belongto_member_project.creator, status, tbl_role_belongto_member_project.create_date " +
		"FROM tbl_member_belongto_project " +
		"INNER JOIN tbl_user ON (tbl_user.email = tbl_member_belongto_project.id_user) " +
		"INNER JOIN tbl_role_belongto_member_project ON (tbl_role_belongto_member_project.id_member_project = tbl_member_belongto_project.id) " +
		"INNER JOIN tbl_role ON (tbl_role_belongto_member_project.id_role = tbl_role.id) " +
		"WHERE tbl_member_belongto_project.id_project = $1"
	res3, err3 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement3,
		View.Id_project,
	)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	if err3 != nil {
		fmt.Println(err3)

	} else {
		fmt.Println(res3)
	}
	result := MemberView{}

	for res3.Next() {
		task := MemberViewtask{}
		err3 := res3.Scan(&task.Id_project, &task.Name, &task.Position)
		// Exit if we get an error
		if err3 != nil {
			fmt.Println(err3)

		}
		result.MemberView = append(result.MemberView, task)
	}

	return result
}

//view member
func (ExampleModel Models) ViewMemberNotInProject(View MemberNotInProjectViewtask) MemberNotInProjectView {
	sqlStatement3 := "SELECT email, name FROM public.tbl_user " +
		"LEFT JOIN tbl_member_belongto_project ON tbl_member_belongto_project.id_user = tbl_user.email " +
		"WHERE tbl_member_belongto_project.id_user IS NULL " +
		"ORDER BY name DESC "
	res3, err3 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement3)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	if err3 != nil {
		fmt.Println(err3)

	} else {
		fmt.Println(res3)
	}
	result := MemberNotInProjectView{}

	for res3.Next() {
		task := MemberNotInProjectViewtask{}
		err3 := res3.Scan(&task.Email, &task.Name)
		// Exit if we get an error
		if err3 != nil {
			fmt.Println(err3)

		}
		result.MemberNotInProjectView = append(result.MemberNotInProjectView, task)
	}

	return result
}

//delete member
func (ExampleModel Models) DelMember(Id_member int) bool {

	sqlStatement2 := "DELETE FROM tbl_member_belongto_project " +
		"WHERE id = $1"
	res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
		Id_member,
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

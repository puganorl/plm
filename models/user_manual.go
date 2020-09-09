package models

import "fmt"

type Umanualtask struct {
	Email           string `json:"email"`
	Id_project      int    `json:"id_project"`
	Link_usermanual string `json:"link_usermanual"`
	Id              int    `json:"id"`
	Status          string `json:"status"`
}

type ViewUmanualtask struct {
	Email      string `json:"email"`
	Id_project int    `json:"id_project"`
}

type UmanualView struct {
	UmanualView []ViewUmanualtask `json:"ui_view"`
}

//add UI UX
func (ExampleModel Models) AddUmanual(Add Umanualtask) bool {

	sqlStatement2 := "INSERT INTO  tbl_design_interface (id, design_interface, description, creator, create_date, approval) " +
		"VALUES ($1, $2 ,$3, $4, now()::timestamp, $5)"
	res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
		Add.Id_project,
		Add.Link_usermanual,
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

//edit UI UX
func (ExampleModel Models) EditUmanual(Edit Umanualtask) bool {

	sqlStatement2 := "UPDATE tbl_design_interface " +
		"SET design_interface = $1, description = $2, approval = $3 " +
		"WHERE id = $4  "
	res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
		Edit.Id_project,
		Edit.Id,
		Edit.Status,
		Edit.Link_usermanual,
		Edit.Status,
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

//view UI UX
func (ExampleModel Models) ViewUmanual(View ViewUmanualtask) UmanualView {

	sqlStatement3 := "SELECT  tbl_design_interface.id, tbl_member_belongto_project.creator FROM tbl_design_interface " +
		"INNER JOIN tbl_member_belongto_project ON tbl_design_interface.creator = tbl_member_belongto_project.id " +
		"WHERE tbl_design_interface.id=$1 "
	res3, err3 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement3,
		View.Id_project,
	)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	if err3 != nil {
		fmt.Println(err3)

	} else {
		fmt.Println(res3)
	}

	result := UmanualView{}

	for res3.Next() {
		task := ViewUmanualtask{}
		err3 := res3.Scan(&task.Id_project, &task.Email)
		// Exit if we get an error
		if err3 != nil {
			fmt.Println(err3)
		}
		result.UmanualView = append(result.UmanualView, task)
	}
	return result
}

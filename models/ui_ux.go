package models

import "fmt"

type Uitask struct {
	Email                 string `json:"email"`
	Id_journey            int    `json:"id_journey"`
	Id_dataflow_structure int    `json:"id_dataflow_structure"`
	Description           string `json:"description"`
	Link_design           string `json:"link_design"`
	Status                string `json:"status"`
}

type ViewUitask struct {
	Email      string `json:"email"`
	Id_journey int    `json:"id_journey"`
}

type UiView struct {
	UiView []ViewUitask `json:"ui_view"`
}

//add UI UX
func (ExampleModel Models) AddUi(Add Uitask) bool {

	sqlStatement2 := "INSERT INTO  tbl_design_interface (id, design_interface, description, creator, create_date, approval) " +
		"VALUES ($1, $2 ,$3, $4, now()::timestamp, $5)"
	res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
		Add.Id_journey,
		Add.Link_design,
		Add.Description,
		2,
		"pending",
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
func (ExampleModel Models) EditUi(Edit Uitask) bool {

	sqlStatement2 := "UPDATE tbl_design_interface " +
		"SET design_interface = $1, description = $2, approval = $3 " +
		"WHERE id = $4  "
	res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
		Edit.Link_design,
		Edit.Description,
		Edit.Status,
		Edit.Id_journey,
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
func (ExampleModel Models) ViewUi(View ViewUitask) UiView {
	sqlStatement3 := "SELECT  tbl_design_interface.id, tbl_member_belongto_project.creator FROM tbl_design_interface " +
		"INNER JOIN tbl_member_belongto_project ON tbl_design_interface.creator = tbl_member_belongto_project.id " +
		"WHERE tbl_design_interface.id=$1 "
	res3, err3 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement3,
		View.Id_journey,
	)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	if err3 != nil {
		fmt.Println(err3)

	} else {
		fmt.Println(res3)
	}

	result := UiView{}

	for res3.Next() {
		task := ViewUitask{}
		err3 := res3.Scan(&task.Id_journey, &task.Email)
		// Exit if we get an error
		if err3 != nil {
			fmt.Println(err3)

		}
		result.UiView = append(result.UiView, task)
	}
	return result
}

package models

import "fmt"

type Userjrnytask struct {
	Id_journey     int    `json:"id_journey"`
	Email          string `json:"email"`
	Id_dataflow    int    `json:"id_dataflow"`
	Journey_name   string `json:"journey_name"`
	Initiate_state int    `json:"initiate_state"`
	Journey_link   string `json:"journey_link"`
	Journey_index  int    `json:"journey_index"`
	Status         string `json:"status"`
	Description    string `json:"description"`
}

type UserjrnyDflowViewtask struct {
	Id             int    `json:"id"`
	Id_dataflow    int    `json:"id_dataflow"`
	Journey_name   string `json:"journey_name"`
	Initiate_state int    `json:"initiate_state"`
	Userjourney    string `json:"userjourney"`
	Index          int    `json:"index"`
	Description    string `json:"description"`
	Status         string `json:"status"`
	Approval       string `json:"approval"`
}

type UserjrnyDflowView struct {
	UserjrnyDflowView []UserjrnyDflowViewtask `json:"userjrny_dflow_view"`
}

type UserjrnyDflowStructViewtask struct {
	Email                 string `json:"email"`
	Id_dataflow_structure int    `json:"id_dataflow_structure"`
}

type UserjrnyDflowStructView struct {
	UserjrnyDflowStructView []UserjrnyDflowStructViewtask `json:"userjrny_dflow_struct_view"`
}

type UserjrnyProjViewtask struct {
	Email      string `json:"email"`
	Id_project int    `json:"id_project"`
}

type UserjrnyProjView struct {
	UserjrnyProjView []UserjrnyProjViewtask `json:"userjrny_proj_view"`
}

//add userjourney
func (ExampleModel Models) AddUserjrnys(Add Userjrnytask) bool {

	sqlStatement2 := "INSERT INTO  tbl_user_journey (id_dataflow, name, initiate_state, userjourney, index, creator, description, status, create_date, approval) " +
		"VALUES ($1, $2 ,$3, $4, $5, $6, $7, $8, now()::timestamp, $9)"
	res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
		Add.Id_dataflow,
		Add.Journey_name,
		Add.Initiate_state,
		Add.Journey_link,
		Add.Journey_index,
		2,
		Add.Description,
		'-',
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

//edit scenario
func (ExampleModel Models) EditUserjrnys(Edit Userjrnytask) bool {

	sqlStatement2 := "UPDATE  tbl_user_journey " +
		"SET  name=$1, initiate_state=$2, userjourney=$3, index=$4, creator=$5, description=$6, " +
		"status=$7, create_date=now()::timestamp, approval=$8 " +
		"WHERE id = $9  "
	res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
		Edit.Journey_name,
		Edit.Initiate_state,
		Edit.Journey_link,
		Edit.Journey_index,
		2,
		'-',
		Edit.Status,
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

//view scenario by dataflow
func (ExampleModel Models) ViewUserjrnyDflow(View UserjrnyDflowViewtask) UserjrnyDflowView {

	sqlStatement3 := "SELECT usj.id, usj.id_dataflow, usj.name as journey_name, usj.initiate_state, usj.userjourney, usj.index, usj.description, usj.status, usj.approval " +
		"FROM tbl_user_journey usj WHERE usj.id_dataflow = $1"
	res3, err3 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement3,
		View.Id_dataflow,
	)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	if err3 != nil {
		fmt.Println(err3)

	} else {
		fmt.Println(res3)
	}
	result := UserjrnyDflowView{}

	for res3.Next() {
		task := UserjrnyDflowViewtask{}
		err3 := res3.Scan(
			&task.Id,
			&task.Id_dataflow,
			&task.Journey_name,
			&task.Initiate_state,
			&task.Userjourney,
			&task.Index,
			&task.Description,
			&task.Status,
			&task.Approval,
		)
		// Exit if we get an error
		if err3 != nil {
			fmt.Println(err3)

		}
		result.UserjrnyDflowView = append(result.UserjrnyDflowView, task)
	}

	return result

}

//view scenario by dataflow structure
func (ExampleModel Models) ViewUserjrnyDflowStruct(View UserjrnyDflowStructViewtask) UserjrnyDflowStructView {

	sqlStatement3 := "SELECT tbl_dataflow_structure.id, tbl_member_belongto_project.id_user  FROM tbl_user_journey " +
		"INNER JOIN tbl_member_belongto_project ON tbl_user_journey.creator = tbl_member_belongto_project.id " +
		"INNER JOIN tbl_dataflow_structure ON tbl_user_journey.id_dataflow = tbl_dataflow_structure.id_dataflow " +
		"WHERE tbl_dataflow_structure.id=$1 "
	res3, err3 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement3,
		View.Id_dataflow_structure,
	)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	if err3 != nil {
		fmt.Println(err3)

	} else {
		fmt.Println(res3)
	}
	result := UserjrnyDflowStructView{}

	for res3.Next() {
		task := UserjrnyDflowStructViewtask{}
		err3 := res3.Scan(&task.Id_dataflow_structure, &task.Email)
		// Exit if we get an error
		if err3 != nil {
			fmt.Println(err3)

		}
		result.UserjrnyDflowStructView = append(result.UserjrnyDflowStructView, task)
	}

	return result

}

//view scenario by project
func (ExampleModel Models) ViewUserjrnyProj(View UserjrnyProjViewtask) UserjrnyProjView {

	sqlStatement3 := "SELECT  tbl_member_belongto_project.id_project, tbl_member_belongto_project.creator  FROM tbl_user_journey " +
		"INNER JOIN tbl_member_belongto_project ON tbl_user_journey.creator = tbl_member_belongto_project.id  " +
		"WHERE tbl_member_belongto_project.id_project=$1 "
	res3, err3 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement3,
		View.Id_project,
	)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	if err3 != nil {
		fmt.Println(err3)

	} else {
		fmt.Println(res3)
	}
	result := UserjrnyProjView{}

	for res3.Next() {
		task := UserjrnyProjViewtask{}
		err3 := res3.Scan(&task.Id_project, &task.Email)
		// Exit if we get an error
		if err3 != nil {
			fmt.Println(err3)

		}
		result.UserjrnyProjView = append(result.UserjrnyProjView, task)
	}

	return result

}

//delete dataflow
func (ExampleModel Models) DeleteUserjrny(Id int) bool {
	sqlStatement2 := "DELETE FROM tbl_user_journey " +
		"WHERE id = $1"
	res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
		Id,
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

package models

import "fmt"

type Usecasetask struct {
	Id         int    `json:"id"`
	Email      string `json:"email"`
	Id_project int    `json:"id_project"`
	Link_usec  string `json:"link_usec"`
}

type UsecDescTask struct {
	Id          int    `json:"id"`
	Email       string `json:"email"`
	Id_usecase  int    `json:"id_usecase"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Index       int    `json:"index"`
}

type UsecDesc struct {
	Id          int    `json:"id"`
	Id_usecase  int    `json:"id_usecase"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Index       int    `json:"index"`
}

type UsecViewtask struct {
	Id          int    `json:"id"`
	Id_project  int    `json:"id_project"`
	Email       string `json:"email"`
	Usecase     string `json:"usecase"`
	Creator     string `json:"creator"`
	Create_date string `json:"create_date"`
	Approval    string `json:"approval"`
}

type UsecViewErdtask struct {
	Id_erd int    `json:"id_erd"`
	Email  string `json:"email"`
}

type UsecView struct {
	UsecView []UsecViewtask `json:"usec_view"`
}

type UsecViewDesc struct {
	UsecViewDesc []UsecDesc `json:"usec_desc_view"`
}

type UsecViewErd struct {
	UsecErdView []UsecViewErdtask `json:"usec_erd_view"`
}

type UsecViewScenartask struct {
	Id_usecase_scenario int    `json:"id_usecase_scenario"`
	Email               string `json:"email"`
}

type UsecViewScenar struct {
	UsecScenarView []UsecViewScenartask `json:"usec_scenar_view"`
}

type UsecViewDflowtask struct {
	Email       string `json:"email"`
	Id_dataflow int    `json:"id_dataflow"`
}

type UsecViewDflow struct {
	UsecDflowView []UsecViewDflowtask `json:"usec_dflow_view"`
}

//add usecase
func (ExampleModel Models) AddUsec(Add Usecasetask) bool {

	sqlStatement2 := "INSERT INTO  tbl_usecase (id,usecase, creator, approval) " +
		"VALUES ($1,$2 ,$3, $4)"
	res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
		Add.Id_project,
		Add.Link_usec,
		135,
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

//edit usecase
func (ExampleModel Models) EditUsec(Edit Usecasetask) bool {

	sqlStatement2 := "UPDATE tbl_usecase " +
		"SET usecase = $1, creator = $2, approval =  $3 " +
		"WHERE id = $4  "
	res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
		Edit.Link_usec,
		2,
		"pending",
		Edit.Id,
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

//add usecase description
func (ExampleModel Models) AddUsecDesc(Add UsecDescTask) bool {

	sqlStatement2 := "INSERT INTO  tbl_usecase_desc (id_usecase, index, name, description, creator, create_date) " +
		"VALUES ($1,$2 ,$3, $4, $5, now()::timestamp)"
	res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
		Add.Id_usecase,
		Add.Index,
		Add.Name,
		Add.Description,
		135,
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

//edit usecase descriprion
func (ExampleModel Models) EditUsecDesc(Edit UsecDescTask) bool {

	sqlStatement2 := "UPDATE tbl_usecase_desc " +
		"SET name = $1, description = $2, creator = $3 " +
		"WHERE tbl_usecase_desc.index = $4 AND id = $5  "
	res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
		Edit.Name,
		Edit.Description,
		135,
		Edit.Index,
		Edit.Id,
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

//view usecase
func (ExampleModel Models) ViewUsec(View UsecViewtask) UsecView {
	sqlStatement3 := "SELECT  tbl_usecase.id, tbl_usecase.id, tbl_usecase.usecase, tbl_member_belongto_project.id_user as creator, tbl_member_belongto_project.id_user as email, tbl_usecase.create_date, tbl_usecase.approval FROM tbl_usecase " +
		"INNER JOIN tbl_member_belongto_project ON tbl_usecase.creator = tbl_member_belongto_project.id " +
		"WHERE tbl_usecase.id=$1 "
	res3, err3 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement3,
		View.Id_project,
	)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	if err3 != nil {
		fmt.Println(err3)

	} else {
		fmt.Println(res3)
	}
	result := UsecView{}

	for res3.Next() {
		task := UsecViewtask{}
		// err3 := res3.Scan(&task.Id_project, &task.Email)
		err3 := res3.Scan(&task.Id, &task.Id_project, &task.Usecase, &task.Email, &task.Creator, &task.Create_date, &task.Approval)
		// Exit if we get an error
		if err3 != nil {
			fmt.Println(err3)

		}
		result.UsecView = append(result.UsecView, task)
	}

	return result

}

//view usecase by erd
func (ExampleModel Models) ViewUsecErd(View UsecViewErdtask) UsecViewErd {

	sqlStatement3 := "SELECT tbl_erd.id, tbl_member_belongto_project.id_user  FROM tbl_usecase " +
		"INNER JOIN tbl_member_belongto_project ON tbl_usecase.creator = tbl_member_belongto_project.id  " +
		"INNER JOIN tbl_erd ON tbl_usecase.id = tbl_erd.id_usecase " +
		"WHERE tbl_erd.id=$1 "
	res3, err3 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement3,
		View.Id_erd,
	)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	if err3 != nil {
		fmt.Println(err3)

	} else {
		fmt.Println(res3)
	}
	result := UsecViewErd{}

	for res3.Next() {
		task := UsecViewErdtask{}
		err3 := res3.Scan(&task.Id_erd, &task.Email)
		//err3 := res3.Scan(&task.Id, &task.Id_Usecase, &task.Index, &task.Name, &task.Description, &task.Creator, &task.Create_date)
		// Exit if we get an error
		if err3 != nil {
			fmt.Println(err3)

		}
		result.UsecErdView = append(result.UsecErdView, task)
	}
	return result
}

//view usecase by Scenario
func (ExampleModel Models) ViewUsecScenar(View UsecViewScenartask) UsecViewScenar {
	sqlStatement3 := "SELECT tbl_usecase_scenario.id, tbl_member_belongto_project.id_user  FROM tbl_usecase_desc " +
		"INNER JOIN tbl_member_belongto_project ON tbl_usecase_desc.creator = tbl_member_belongto_project.id " +
		"INNER JOIN tbl_usecase_scenario ON tbl_usecase_desc.id = tbl_usecase_scenario.id_usecase " +
		"WHERE tbl_usecase_scenario.id=$1 "
	res3, err3 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement3,
		View.Id_usecase_scenario,
	)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	if err3 != nil {
		fmt.Println(err3)

	} else {
		fmt.Println(res3)
	}
	result := UsecViewScenar{}

	for res3.Next() {
		task := UsecViewScenartask{}
		err3 := res3.Scan(&task.Id_usecase_scenario, &task.Email)
		//err3 := res3.Scan(&task.Id, &task.Id_Usecase, &task.Case_type, &task.Initiate_state, &task.Request,
		//	&task.Response, &task.Expectation, &task.Valid, &task.Create_date, &task.Mod_date, &task.Approval)
		// Exit if we get an error
		if err3 != nil {
			fmt.Println(err3)

		}
		result.UsecScenarView = append(result.UsecScenarView, task)
	}
	return result
}

//view usecase by dataflow
func (ExampleModel Models) ViewUsecDataflow(View UsecViewDflowtask) UsecViewDflow {
	sqlStatement3 := "SELECT tbl_dataflow.id, tbl_member_belongto_project.id_user  FROM tbl_usecase_desc " +
		"INNER JOIN tbl_member_belongto_project ON tbl_usecase_desc.creator = tbl_member_belongto_project.id " +
		"INNER JOIN tbl_dataflow ON tbl_usecase_desc.id = tbl_dataflow.id_usecase " +
		"WHERE tbl_dataflow.id=$1 "
	res3, err3 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement3,
		View.Id_dataflow,
	)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	if err3 != nil {
		fmt.Println(err3)

	} else {
		fmt.Println(res3)
	}
	result := UsecViewDflow{}

	for res3.Next() {
		task := UsecViewDflowtask{}
		err3 := res3.Scan(&task.Id_dataflow, &task.Email)
		//err3 := res3.Scan(&task.Id, &task.Id_usecase, &task.Description, &task.Creator, &task.Create_date, &task.Approval)
		// Exit if we get an error
		if err3 != nil {
			fmt.Println(err3)

		}
		result.UsecDflowView = append(result.UsecDflowView, task)
	}
	return result
}

// view usecase desc by usecase id
func (ExampleModel Models) ViewUsecDesc(View UsecDesc) UsecViewDesc {
	sqlStatement3 := "SELECT id, id_usecase, name, description, index " +
		"FROM tbl_usecase_desc WHERE tbl_usecase_desc.id_usecase = $1"
	res3, err3 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement3,
		View.Id_usecase,
	)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Println(res3)
	}

	result := UsecViewDesc{}

	for res3.Next() {
		task := UsecDesc{}
		err3 := res3.Scan(&task.Id, &task.Id_usecase, &task.Name, &task.Description, &task.Index)
		// Exit if we get an error
		if err3 != nil {
			fmt.Println(err3)

		}
		result.UsecViewDesc = append(result.UsecViewDesc, task)
	}
	return result
}

//delete usecase descriprion
func (ExampleModel Models) DelUsecDesc(Id int) bool {

	sqlStatement2 := "DELETE FROM tbl_usecase_desc " +
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

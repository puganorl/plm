package models

import "fmt"

type UsecScentask struct {
	Email           string `json:"email"`
	Id              int    `json:"id"`
	Id_usecase_desc int    `json:"id_usecase_desc"`
	Case_type       string `json:"case_type"`
	Initiate_state  string `json:"initiate_state"`
	Request         string `json:"request"`
	Response        string `json:"response"`
	Expectation     string `json:"expectation"`
	Valid           bool   `json:"valid"`
	Status          string `json:"status"`
}

type ViewUsecScentask struct {
	Email           string `json:"email"`
	Id              int    `json:"id"`
	Id_usecase_desc int    `json:"id_usecase_desc"`
	Case_type       string `json:"case_type"`
	Initiate_state  string `json:"initiate_state"`
	Request         string `json:"request"`
	Response        string `json:"response"`
	Expectation     string `json:"expectation"`
	Valid           bool   `json:"valid"`
	Status          string `json:"status"`
}

type UsecScenView struct {
	UsecScenView []ViewUsecScentask `json:"usec_scen_view"`
}

type ViewUsecScenProjtask struct {
	Email      string `json:"email"`
	Id_project int    `json:"id_project"`
}

type UsecScenProjView struct {
	UsecScenViewProj []ViewUsecScenProjtask `json:"usec_scen_view_proj"`
}

//add scenario
func (ExampleModel Models) AddScenario(Add UsecScentask) bool {

	sqlStatement2 := "INSERT INTO  tbl_usecase_scenario (id_usecase, case_type, initiate_state, request, response, expectation, valid, creator, create_date, mod_date, approval) " +
		"VALUES ($1, $2 ,$3, $4, $5, $6, $7, $8, now()::timestamp, now()::timestamp, $9)"
	res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
		Add.Id_usecase_desc,
		Add.Case_type,
		Add.Initiate_state,
		Add.Request,
		Add.Response,
		Add.Expectation,
		Add.Valid,
		2,
		"pending",
	)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	fmt.Println(Add.Id_usecase_desc)
	if err2 != nil {
		fmt.Println(err2)
		return false
	} else {
		fmt.Println(res2)
		return true
	}
}

//edit scenario
func (ExampleModel Models) EditScenario(Edit UsecScentask) bool {

	sqlStatement2 := "UPDATE  tbl_usecase_scenario " +
		"SET case_type=$1, initiate_state=$2, request=$3, response=$4, " +
		"expectation=$5, valid=$6, mod_date=now()::timestamp, approval=$7 " +
		"WHERE id = $8  "
	res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
		Edit.Case_type,
		Edit.Initiate_state,
		Edit.Request,
		Edit.Response,
		Edit.Expectation,
		Edit.Valid,
		Edit.Status,
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

//view scenario
func (ExampleModel Models) ViewScen(View ViewUsecScentask) UsecScenView {
	sqlStatement3 := "SELECT scen.id, scen.id_usecase AS id_usecase_desc, case_type, " +
		"initiate_state, request, response, expectation, valid, approval as status " +
		"FROM tbl_usecase_scenario AS scen " +
		"JOIN tbl_usecase_desc AS udesc ON scen.id_usecase = udesc.id " +
		"WHERE scen.id_usecase = $1"
	res3, err3 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement3,
		View.Id_usecase_desc,
	)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	if err3 != nil {
		fmt.Println(err3)

	} else {
		fmt.Println(res3)
	}

	result := UsecScenView{}

	for res3.Next() {
		task := ViewUsecScentask{}
		err3 := res3.Scan(&task.Id, &task.Id_usecase_desc, &task.Case_type, &task.Initiate_state, &task.Request, &task.Response, &task.Expectation, &task.Valid, &task.Status)
		// Exit if we get an error
		if err3 != nil {
			fmt.Println(err3)
		}
		result.UsecScenView = append(result.UsecScenView, task)
	}

	return result

}

//view scenario by proj
func (ExampleModel Models) ViewScenProj(View ViewUsecScenProjtask) UsecScenProjView {

	sqlStatement3 := "SELECT  tbl_member_belongto_project.id_project, tbl_member_belongto_project.creator  FROM tbl_usecase_scenario " +
		"INNER JOIN tbl_member_belongto_project ON tbl_usecase_scenario.creator = tbl_member_belongto_project.id " +
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

	result := UsecScenProjView{}

	for res3.Next() {
		task := ViewUsecScenProjtask{}
		err3 := res3.Scan(&task.Id_project, &task.Email)
		// Exit if we get an error
		if err3 != nil {
			fmt.Println(err3)

		}
		result.UsecScenViewProj = append(result.UsecScenViewProj, task)
	}
	return result
}

//delete dataflow
func (ExampleModel Models) DeleteScenario(Id int) bool {
	sqlStatement2 := "DELETE FROM tbl_usecase_scenario " +
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

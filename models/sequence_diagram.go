package models

import "fmt"

type SeqDiagtask struct {
	Id                    int    `json:"id"`
	Email                 string `json:"email"`
	Id_dataflow_structure int    `json:"id_dataflow_structure"`
	Sequence_link         string `json:"sequence_link"`
	Description           string `json:"description"`
	Status                string `json:"status"`
}

type ViewSeqDiagtask struct {
	Id                    int    `json:"id"`
	Id_dataflow_structure int    `json:"id_dataflow_structure"`
	Sequence_link         string `json:"sequence_link"`
	Description           string `json:"description"`
	Status                string `json:"status"`
}

type SeqDiagView struct {
	SeqDiagView []ViewSeqDiagtask `json:"seq_diag_view"`
}

type ViewSeqDiagProjtask struct {
	Email      string `json:"email"`
	Id_project int    `json:"id_project"`
}

type SeqDiagProjView struct {
	SeqDiagProjView []ViewSeqDiagProjtask `json:"seq_diag_proj_view"`
}

//add sequence diagram
func (ExampleModel Models) AddSeqDiags(Add SeqDiagtask) bool {

	sqlStatement2 := "INSERT INTO  tbl_sequence_diagram (id_dataflow_structure, sequence_diagram, description, creator, create_date, approval) " +
		"VALUES ($1, $2 ,$3, $4, now()::timestamp, $5)"
	res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
		Add.Id_dataflow_structure,
		Add.Sequence_link,
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

//edit sequence diagram
func (ExampleModel Models) EditSeqDiags(Edit SeqDiagtask) bool {

	sqlStatement2 := "UPDATE tbl_sequence_diagram " +
		"SET id_dataflow_structure = $1, sequence_diagram = $2, description = $3, creator = $4, approval = $5 " +
		"WHERE id = $6  "
	res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
		Edit.Id_dataflow_structure,
		Edit.Sequence_link,
		Edit.Description,
		2,
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

//view sequence diagram
func (ExampleModel Models) ViewSeqDiag(View ViewSeqDiagtask) SeqDiagView {

	sqlStatement3 := "SELECT seq.id, seq.id_dataflow_structure, seq.sequence_diagram AS sequence_link, seq.description, seq.approval AS status " +
		"FROM tbl_sequence_diagram seq WHERE id_dataflow_structure = $1"
	res3, err3 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement3,
		View.Id_dataflow_structure,
	)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	if err3 != nil {
		fmt.Println(err3)

	} else {
		fmt.Println(res3)
	}

	result := SeqDiagView{}

	for res3.Next() {
		task := ViewSeqDiagtask{}
		err3 := res3.Scan(
			&task.Id,
			&task.Id_dataflow_structure,
			&task.Sequence_link,
			&task.Description,
			&task.Status,
		)

		// Exit if we get an error
		if err3 != nil {
			fmt.Println(err3)

		}

		//result.Email = Email
		result.SeqDiagView = append(result.SeqDiagView, task)
	}

	return result

}

//view sequence diagram by project
func (ExampleModel Models) ViewSeqDiagProj(View ViewSeqDiagProjtask) SeqDiagProjView {

	sqlStatement3 := "SELECT  tbl_member_belongto_project.id_project, tbl_member_belongto_project.creator  FROM tbl_sequence_diagram " +
		"INNER JOIN tbl_member_belongto_project ON tbl_sequence_diagram.creator = tbl_member_belongto_project.id " +
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

	result := SeqDiagProjView{}

	for res3.Next() {
		task := ViewSeqDiagProjtask{}
		err3 := res3.Scan(&task.Id_project, &task.Email)
		// Exit if we get an error
		if err3 != nil {
			fmt.Println(err3)

		}

		//result.Email = Email
		result.SeqDiagProjView = append(result.SeqDiagProjView, task)
	}
	return result
}

//delete dataflow
func (ExampleModel Models) DeleteSeqDiag(Id int) bool {
	sqlStatement2 := "DELETE FROM tbl_sequence_diagram " +
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

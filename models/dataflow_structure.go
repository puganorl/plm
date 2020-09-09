package models

import "fmt"

type DataFstructask struct {
	Id            int    `json:"id"`
	Email         string `json:"email"`
	Id_dataflow   int    `json:"id_dataflow"`
	Index         int    `json:"index"`
	Dataflow_name string `json:"dataflow_name"`
	Protocol      int    `json:"protocol"`
	Type          int    `json:"type"`
	Address       string `json:"address"`
	Request_type  string `json:"request_type"`
	Response_type string `json:"response_type"`
	Status        string `json:"status"`
}

type DataflowView struct {
	Id            int    `json:"id"`
	Id_dataflow   int    `json:"id_dataflow"`
	Index         int    `json:"index"`
	Dataflow_name string `json:"dataflow_name"`
	Protocol      int    `json:"protocol"`
	Type          int    `json:"type"`
	Address       string `json:"address"`
	Request_type  string `json:"request_type"`
	Response_type string `json:"response_type"`
	Approval      string `json:"approval"`
}

type DflowView struct {
	DataflowView []DataflowView `json:"dataflow_userjour_view"`
}

type DataflowSequnceViewtask struct {
	Email       string `json:"email"`
	Id_sequence int    `json:"id_sequence"`
}

type DataflowSequenceView struct {
	DataflowSequenceView []DataflowSequnceViewtask `json:"dataflow_sequence_view"`
}

type DataflowProjViewtask struct {
	Email      string `json:"email"`
	Id_project int    `json:"id_project"`
}

type DataflowProjView struct {
	DataflowProjView []DataflowProjViewtask `json:"dataflow_proj_view"`
}

type DataflowDesignViewtask struct {
	Email     string `json:"email"`
	Id_design int    `json:"id_design"`
}

type DataflowDesignView struct {
	DataflowDesignView []DataflowDesignViewtask `json:"dataflow_design_view"`
}

//add dataflow_sctructure
func (ExampleModel Models) AddDataFStrucs(Add DataFstructask) bool {

	sqlStatement2 := "INSERT INTO  tbl_dataflow_structure (id_dataflow, index, name, protocol, type, address, request_type, response_type, creator, create_date, approval) " +
		"VALUES ($1, $2 ,$3, $4, $5, $6, $7, $8, $9, now()::timestamp, $10)"
	res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
		Add.Id_dataflow,
		Add.Index,
		Add.Dataflow_name,
		Add.Protocol,
		Add.Type,
		Add.Address,
		Add.Request_type,
		Add.Response_type,
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

//edit dataflow_sctructure
func (ExampleModel Models) EditDataFStrucs(Add DataFstructask) bool {

	sqlStatement2 := "UPDATE tbl_dataflow_structure " +
		"SET id_dataflow=$1, index=$2, name=$3, protocol=$4, type=$5, address=$6, " +
		"request_type=$7, response_type=$8, creator=$9, approval=$10 " +
		"WHERE id = $11 "
	res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
		Add.Id_dataflow,
		Add.Index,
		Add.Dataflow_name,
		Add.Protocol,
		Add.Type,
		Add.Address,
		Add.Request_type,
		Add.Response_type,
		2,
		Add.Status,
		Add.Id,
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

//view dataflow structure by dataflow
func (ExampleModel Models) ViewDflow(View DataflowView) DflowView {

	sqlStatement3 := "SELECT dfs.id, dfs.id_dataflow, dfs.index, dfs.name AS dataflow_name, dfs.protocol, dfs.type, dfs.address, dfs.request_type, dfs.response_type, dfs.approval " +
		"FROM tbl_dataflow_structure dfs " +
		"JOIN tbl_dataflow ON tbl_dataflow.id = dfs.id_dataflow " +
		"WHERE dfs.id_dataflow = $1"
	res3, err3 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement3,
		View.Id_dataflow,
	)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	if err3 != nil {
		fmt.Println(err3)

	} else {
		fmt.Println(res3)
	}
	result := DflowView{}

	for res3.Next() {
		task := DataflowView{}
		err3 := res3.Scan(
			&task.Id,
			&task.Id_dataflow,
			&task.Index,
			&task.Dataflow_name,
			&task.Protocol,
			&task.Type,
			&task.Address,
			&task.Request_type,
			&task.Response_type,
			&task.Approval,
		)
		// Exit if we get an error
		if err3 != nil {
			fmt.Println(err3)

		}
		result.DataflowView = append(result.DataflowView, task)
	}

	return result

}

//view dataflow structure by sequence
func (ExampleModel Models) ViewDflowSequence(View DataflowSequnceViewtask) DataflowSequenceView {

	sqlStatement3 := "SELECT tbl_sequence_diagram.id, tbl_member_belongto_project.id_user  FROM tbl_dataflow_structure " +
		"INNER JOIN tbl_member_belongto_project ON tbl_dataflow_structure.creator = tbl_member_belongto_project.id " +
		"INNER JOIN tbl_sequence_diagram ON tbl_dataflow_structure.id = tbl_sequence_diagram.id_dataflow_structure " +
		"WHERE tbl_sequence_diagram.id=$1 "
	res3, err3 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement3,
		View.Id_sequence,
	)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	if err3 != nil {
		fmt.Println(err3)

	} else {
		fmt.Println(res3)
	}
	result := DataflowSequenceView{}

	for res3.Next() {
		task := DataflowSequnceViewtask{}
		err3 := res3.Scan(&task.Id_sequence, &task.Email)
		//err3 := res3.Scan(&task.Id, &task.Id_dataflow, &task.Index, &task.Name, &task.Protocol, &task.Type,
		//	&task.Index, &task.Request_type, &task.Response_type, &task.Creator, &task.Create_date, &task.Approval)
		// Exit if we get an error
		if err3 != nil {
			fmt.Println(err3)

		}
		result.DataflowSequenceView = append(result.DataflowSequenceView, task)
	}

	return result

}

//view dataflow structure by design
func (ExampleModel Models) ViewDflowDesign(View DataflowDesignViewtask) DataflowDesignView {

	sqlStatement3 := "SELECT tbl_design_belongto_dataflow_structure.id_design, tbl_member_belongto_project.id_user  FROM tbl_dataflow_structure " +
		"INNER JOIN tbl_member_belongto_project ON tbl_dataflow_structure.creator = tbl_member_belongto_project.id " +
		"INNER JOIN tbl_design_belongto_dataflow_structure ON tbl_dataflow_structure.id = tbl_design_belongto_dataflow_structure.id_dataflow_structure " +
		"WHERE tbl_design_belongto_dataflow_structure.id_design=$1 "
	res3, err3 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement3,
		View.Id_design,
	)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	if err3 != nil {
		fmt.Println(err3)

	} else {
		fmt.Println(res3)
	}
	result := DataflowDesignView{}

	for res3.Next() {
		task := DataflowDesignViewtask{}
		err3 := res3.Scan(&task.Id_design, &task.Email)
		// Exit if we get an error
		if err3 != nil {
			fmt.Println(err3)

		}
		result.DataflowDesignView = append(result.DataflowDesignView, task)
	}

	return result

}

//view dataflow structure by project
func (ExampleModel Models) ViewDflowStructProj(View DataflowProjViewtask) DataflowProjView {

	sqlStatement3 := "SELECT  tbl_member_belongto_project.id_project, tbl_member_belongto_project.creator  FROM tbl_dataflow " +
		"INNER JOIN tbl_member_belongto_project ON tbl_dataflow.creator = tbl_member_belongto_project.id  " +
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
	result := DataflowProjView{}

	for res3.Next() {
		task := DataflowProjViewtask{}
		err3 := res3.Scan(&task.Id_project, &task.Email)
		// Exit if we get an error
		if err3 != nil {
			fmt.Println(err3)

		}
		result.DataflowProjView = append(result.DataflowProjView, task)
	}

	return result

}

//delete dataflow
func (ExampleModel Models) DeleteDataFStruc(Id int) bool {

	sqlStatement2 := "DELETE FROM tbl_dataflow_structure " +
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

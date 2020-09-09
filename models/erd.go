package models

import (
	"fmt"
)

type Erdtask struct {
	Id           int    `json:"id"`
	Email        string `json:"email"`
	Id_erd       int    `json:"id_erd"`
	Table_name   string `json:"table_name"`
	Id_usecase   int    `json:"id_usecase"`
	Link_erd     string `json:"link_erd"`
	Description  string `json:"description"`
	Status       string `json:"status"`
	Id_table     int    `json:"id_table"`
	Field_name   string `json:"field_name"`
	Field_type   string `json:"field_type"`
	Field_length int    `json:"field_length"`
}

type ErdUsecaseViewtask struct {
	Email      string `json:"email"`
	Id_usecase int    `json:"id_usecase"`
}

type ErdUsecaseView struct {
	ErdUsecaseView []ErdUsecaseViewtask `json:"erd_usecase_view"`
}

type ErdtableViewtask struct {
	Id_erd int    `json:"id_erd"`
	Email  string `json:"email"`
}

type ErdtableView struct {
	ErdtableView []ErdtableViewtask `json:"erdtable_view"`
}

type ErdfieldViewtask struct {
	Id_table int    `json:"id_table"`
	Email    string `json:"email"`
}

type ErdfieldView struct {
	ErdfieldView []ErdfieldViewtask `json:"erdfield_view"`
}

type ErdprojViewtask struct {
	Id          int    `json:"id"`
	Id_project  int    `json:"id_project"`
	Erd         string `json:"erd"`
	Creator     int    `json:"creator"`
	Create_date string `json:"create_date"`
	Email       string `json:"email"`
}

type ErdprojView struct {
	ErdprojView []ErdprojViewtask `json:"erdproj_view"`
}

//add erd
func (ExampleModel Models) AddErd(Add Erdtask) bool {

	sqlStatement2 := "INSERT INTO  tbl_erd (id_usecase, erd, creator, create_date) " +
		"VALUES ($1,$2 ,$3, now()::timestamp)"
	res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
		Add.Id_usecase,
		Add.Link_erd,
		2,
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

//edit erd
func (ExampleModel Models) EditErd(Edit Erdtask) bool {

	sqlStatement2 := "UPDATE  tbl_erd " +
		"SET erd = $1, creator = $2 " +
		"WHERE id = $3  "
	res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
		Edit.Link_erd,
		2,
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

//add table
func (ExampleModel Models) AddErdTable(Add Erdtask) bool {

	sqlStatement2 := "INSERT INTO  tbl_table_belongto_erd (id_erd, name, description, create_date, creator, approval) " +
		"VALUES ($1,$2 ,$3, now()::timestamp, $4, $5)"
	res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
		Add.Id_erd,
		Add.Table_name,
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

//edit table
func (ExampleModel Models) EditErdTable(Edit Erdtask) bool {

	sqlStatement2 := "UPDATE  tbl_table_belongto_erd " +
		"SET name = $1, description = $2, approval = $3 " +
		"WHERE id = $4  "
	res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
		Edit.Table_name,
		Edit.Description,
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

//delete table
func (ExampleModel Models) DelErdTable(Id int) bool {

	sqlStatement2 := "DELETE FROM tbl_table_belongto_erd " +
		"WHERE id =$1"
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

//add field
func (ExampleModel Models) AddErdField(Add Erdtask) bool {

	sqlStatement2 := "INSERT INTO  tbl_field_belongto_table (id_table, field_name, field_type, length, creator, create_date) " +
		"VALUES ($1,$2 ,$3, $4, $5, now()::timestamp)"
	res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
		Add.Id_table,
		Add.Field_name,
		Add.Field_type,
		Add.Field_length,
		2,
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

//edit field
func (ExampleModel Models) EditErdField(Edit Erdtask) bool {

	sqlStatement2 := "UPDATE  tbl_field_belongto_table " +
		"SET field_name = $1, field_type = $2, length = $3 " +
		"WHERE id = $4  "
	res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
		Edit.Field_name,
		Edit.Field_type,
		Edit.Field_length,
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

//delete field
func (ExampleModel Models) DelErdField(Id int) bool {

	sqlStatement2 := "DELETE FROM tbl_field_belongto_table " +
		"WHERE id =$1"
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

//view  erd by usecase
func (ExampleModel Models) ViewErdUsec(View ErdUsecaseViewtask) ErdUsecaseView {

	sqlStatement3 := "SELECT tbl_usecase.id, tbl_member_belongto_project.id_user  FROM tbl_erd " +
		"INNER JOIN tbl_member_belongto_project ON tbl_erd.creator = tbl_member_belongto_project.id " +
		"INNER JOIN tbl_usecase ON tbl_erd.id_usecase = tbl_usecase.id " +
		"WHERE tbl_usecase.id=$1 "
	res3, err3 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement3,
		View.Id_usecase,
	)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	if err3 != nil {
		fmt.Println(err3)

	} else {
		fmt.Println(res3)
	}

	result := ErdUsecaseView{}

	for res3.Next() {
		task := ErdUsecaseViewtask{}
		err3 := res3.Scan(&task.Id_usecase, &task.Email)
		// Exit if we get an error
		if err3 != nil {
			fmt.Println(err3)

		}

		//result.Email = Email
		result.ErdUsecaseView = append(result.ErdUsecaseView, task)
	}

	return result

}

//view table
func (ExampleModel Models) ViewErdTable(View ErdtableViewtask) ErdtableView {

	sqlStatement3 := "SELECT  tbl_table_belongto_erd.id_erd, tbl_member_belongto_project.creator FROM tbl_table_belongto_erd " +
		"INNER JOIN tbl_member_belongto_project ON tbl_table_belongto_erd.creator = tbl_member_belongto_project.id " +
		"WHERE tbl_table_belongto_erd.id_erd=$1 "
	res3, err3 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement3,
		View.Id_erd,
	)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	if err3 != nil {
		fmt.Println(err3)

	} else {
		fmt.Println(res3)
	}

	result := ErdtableView{}

	for res3.Next() {
		task := ErdtableViewtask{}
		err3 := res3.Scan(&task.Id_erd, &task.Email)
		// Exit if we get an error
		if err3 != nil {
			fmt.Println(err3)

		}

		//result.Email = Email
		result.ErdtableView = append(result.ErdtableView, task)
	}

	return result

}

//view field
func (ExampleModel Models) ViewErdField(View ErdfieldViewtask) ErdfieldView {

	sqlStatement3 := "SELECT  tbl_field_belongto_table.id_table, tbl_member_belongto_project.creator FROM tbl_field_belongto_table " +
		"INNER JOIN tbl_member_belongto_project ON tbl_field_belongto_table.creator = tbl_member_belongto_project.id " +
		"WHERE tbl_field_belongto_table.id_table=$1 "
	res3, err3 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement3,
		View.Id_table,
	)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	if err3 != nil {
		fmt.Println(err3)

	} else {
		fmt.Println(res3)
	}

	result := ErdfieldView{}

	for res3.Next() {
		task := ErdfieldViewtask{}
		err3 := res3.Scan(&task.Id_table, &task.Email)
		// Exit if we get an error
		if err3 != nil {
			fmt.Println(err3)

		}
		result.ErdfieldView = append(result.ErdfieldView, task)
	}

	return result

}

//view table all
func (ExampleModel Models) ViewErdTableAll(View ErdtableViewtask) ErdtableView {

	sqlStatement3 := "SELECT  tbl_table_belongto_erd.id_erd, tbl_member_belongto_project.creator FROM tbl_table_belongto_erd " +
		"INNER JOIN tbl_member_belongto_project ON tbl_table_belongto_erd.creator = tbl_member_belongto_project.id " +
		"WHERE tbl_table_belongto_erd.id_erd=$1 "
	res3, err3 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement3,
		View.Id_erd,
	)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	if err3 != nil {
		fmt.Println(err3)

	} else {
		fmt.Println(res3)
	}

	result := ErdtableView{}

	for res3.Next() {
		task := ErdtableViewtask{}
		err3 := res3.Scan(&task.Id_erd, &task.Email)
		// Exit if we get an error
		if err3 != nil {
			fmt.Println(err3)

		}

		//result.Email = Email
		result.ErdtableView = append(result.ErdtableView, task)
	}

	return result

}

//view erd proj
func (ExampleModel Models) ViewErdProj(View ErdprojViewtask) ErdprojView {

	sqlStatement3 := "SELECT * FROM tbl_erd WHERE tbl_erd.id_usecase = $1 LIMIT 1"
	res3, err3 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement3,
		View.Id_project,
	)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	if err3 != nil {
		fmt.Println(err3)

	} else {
		fmt.Println(res3)
	}

	result := ErdprojView{}

	for res3.Next() {
		task := ErdprojViewtask{}
		err3 := res3.Scan(&task.Id, &task.Id_project, &task.Erd, &task.Creator, &task.Create_date)
		// Exit if we get an error
		if err3 != nil {
			fmt.Println(err3)

		}
		result.ErdprojView = append(result.ErdprojView, task)
	}

	return result

}

// DeleteErd delete dataflow
func (ExampleModel Models) DeleteErd(Id int) bool {
	sqlStatement2 := "DELETE FROM tbl_erd " +
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

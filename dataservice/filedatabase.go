package dataservice

import (
	"fmt"
	"go_API_CBCardDev/models"
	"log"

	"golang.org/x/net/context"
)

func GetProcessedfilesFromDb(processedfiles []models.ProcessedFile) []models.ProcessedFile {

	db := ConnectToDb()
	fmt.Println("\nin GetProcessedfilesFromDb ", db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rows, err := db.QueryContext(ctx, "SELECT * FROM MANISH_ORACLE_DB.processedfile")
	fmt.Println(rows, err)
	if err != nil {
		fmt.Println("-------------------", err)
		log.Fatal(err)
	}
	//defer rows.Close()
	for rows.Next() {
		var a models.ProcessedFile
		err := rows.Scan(&a.Fileid, &a.Source_system, &a.Platform, &a.Isproduction, &a.Header_date,
			&a.Process_date, &a.Start_seq, &a.End_seq, &a.File_name, &a.Mabx_csm_fileid)

		if err != nil {
			log.Fatal(err)
		}
		processedfiles = append(processedfiles, a)
		fmt.Println(processedfiles)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return processedfiles
}

// Get the data requested by client with Fileid
func GetprocessedfilesByfileidFromDb(fileid int) (models.ProcessedFile, error) {
	fmt.Println("----In DB--------")
	db := ConnectToDb()
	fmt.Println("\nin GetprocessedfilesByfileidFromDb ", db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var processedfiles models.ProcessedFile
	query := `SELECT * FROM MANISH_ORACLE_DB.processedfile WHERE Fileid = :1`
	fmt.Println("Query:", query)

	// Prepare and execute the query
	rows, err := db.QueryContext(ctx, query, fileid)
	if err != nil {
		return processedfiles, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(
			&processedfiles.Fileid, &processedfiles.Source_system, &processedfiles.Platform,
			&processedfiles.Isproduction, &processedfiles.Header_date, &processedfiles.Process_date,
			&processedfiles.Start_seq, &processedfiles.End_seq, &processedfiles.File_name,
			&processedfiles.Mabx_csm_fileid)
		if err != nil {
			return processedfiles, err
		}
		fmt.Println("Got a particular Processed File from the table successfully")
	}

	return processedfiles, nil
}

func PostAddProcessedfileToDb(processedfiles models.ProcessedFile) {
	db := ConnectToDb()
	fmt.Println("\nin PostaddFILEToDb ", db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query := `INSERT INTO MANISH_ORACLE_DB.processedfile (Fileid, Source_system, Platform,
		Isproduction,Header_date,Process_date, Start_seq, End_seq, File_name, Mabx_csm_fileid) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10)`

	fmt.Println("QUERY:", query)

	_, err := db.ExecContext(ctx, query,
		processedfiles.Fileid, processedfiles.Platform, processedfiles.Source_system,
		processedfiles.Isproduction, processedfiles.Header_date, processedfiles.Process_date,
		processedfiles.Start_seq, processedfiles.End_seq, processedfiles.File_name,
		processedfiles.Mabx_csm_fileid)
	if err != nil {
		fmt.Println("Error in executing query")
		log.Fatal(err)
	}
	fmt.Println("Added New Processed file successfully")
}

func DeleteProcessedfileFromDb(fileid int) {
	db := ConnectToDb()
	fmt.Println("\nin GetPrpcessedfileFromDb ", db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query := fmt.Sprintf("DELETE FROM MANISH_ORACLE_DB.processedfile WHERE fileid = %d", fileid)
	fmt.Println("Query:", query)

	// Execute the DELETE query
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("productioncard with contractnbr:", fileid, "deleted successfully")
}

func UpdateProcessedfileInDb(processedfiles models.ProcessedFile) {
	db := ConnectToDb()
	fmt.Println("\nin GetProcessedfileFromDb ", db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query := `UPDATE MANISH_ORACLE_DB.processedfile SET Fileid =:1, Source_system = :2, Platform= :3,
	Isproduction = :4,Header_date = :5,Process_date = :6, Start_seq = :7, End_seq = :8, File_name = :9, Mabx_csm_fileid = :10`

	fmt.Println("Query:", query)

	_, err := db.ExecContext(ctx, query,
		processedfiles.Fileid, processedfiles.Source_system, processedfiles.Platform,
		processedfiles.Isproduction, processedfiles.Header_date, processedfiles.Process_date,
		processedfiles.Start_seq, processedfiles.End_seq, processedfiles.File_name,
		processedfiles.Mabx_csm_fileid)
	if err != nil {
		fmt.Println("Error in executing query")
		log.Fatal(err)
	}
	fmt.Println("File updated successfully.")
}

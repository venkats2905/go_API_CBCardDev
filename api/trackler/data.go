package trackler

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"tracking/model"

	_ "github.com/sijms/go-ora/v2"
)

var localDB = map[string]string{
	"service":  "xe",
	"username": "system",
	"server":   "localhost",
	"port":     "1521",
	"password": "password",
}

func ConnectToDb() *sql.DB {
	connectionString := "oracle://" + localDB["username"] + ":" + localDB["password"] + "@" + localDB["server"] + ":" + localDB["port"] + "/" + localDB["service"]
	db, err := sql.Open("oracle", connectionString)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging the database:", err)
	}

	fmt.Println("Connected to the Oracle database!")
	fmt.Println("\n------", db)
	return db
}

func GetTracking(Trackrecord []model.Trackingsequence_archive, db *sql.DB) []model.Trackingsequence_archive {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rows, err := db.QueryContext(ctx, "SELECT * FROM MANI_ORACLE_DB.Trackingsequence_archive")
	fmt.Println(rows, err)
	if err != nil {
		fmt.Println("-------------------", err)
		log.Fatal(err)
	}
	//defer rows.Close()
	for rows.Next() {
		var a model.Trackingsequence_archive
		err := rows.Scan(&a.SID, &a.CONTRACTNBR, &a.SUBSEQ, &a.GROUPSEQ, &a.GROUPNBR,
			&a.SUFFIXNBR, &a.HOMEPLAN, &a.IMB, &a.MATRLDIST, &a.DISPCODE,
			&a.MAILDATE, &a.SENTFLAG, &a.CARDTEMPLATECODE, &a.SETUPNAME, &a.REQ_DATE, &a.SEARCHCODE,
			&a.RELEASE_DT, &a.REPTTYPE, &a.FEPFLAG)

		if err != nil {
			log.Fatal(err)
		}
		Trackrecord = append(Trackrecord, a)
		fmt.Println(Trackrecord)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return Trackrecord
}

func trackbySID(SID int, db *sql.DB) (model.Trackingsequence_archive, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var Trackrecord model.Trackingsequence_archive
	query := fmt.Sprintf("SELECT * FROM MANI_ORACLE_DB.Trackingsequence_archive WHERE SID = %d", SID)
	fmt.Println("Query:", query)

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return Trackrecord, err
	}
	defer rows.Close()

	var a model.Trackingsequence_archive
	if rows.Next() {
		err := rows.Scan(
			&a.SID, &a.CONTRACTNBR, &a.SUBSEQ, &a.GROUPSEQ, &a.GROUPNBR,
			&a.SUFFIXNBR, &a.HOMEPLAN, &a.IMB, &a.MATRLDIST, &a.DISPCODE,
			&a.MAILDATE, &a.SENTFLAG, &a.CARDTEMPLATECODE, &a.SETUPNAME, &a.REQ_DATE, &a.SEARCHCODE,
			&a.RELEASE_DT, &a.REPTTYPE, &a.FEPFLAG,
		)
		if err != nil {
			return Trackrecord, err
		}
		fmt.Println("Got a particular production card from the table successfully")
	}

	return a, nil
}

func Postrecord(Trackrecord model.Trackingsequence_archive, db *sql.DB) {
	fmt.Println("\nin Postrecord ", db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query := `INSERT INTO MANI_ORACLE_DB.Trackingsequence_archive (SID, CONTRACTNBR, SUBSEQ, GROUPSEQ, GROUPNBR,
		SUFFIXNBR, HOMEPLAN, IMB, MATRLDIST,DISPCODE,MAILDATE, SENTFLAG, CARDTEMPLATECODE, SETUPNAME, REQ_DATE,SEARCHCODE,
		RELEASE_DT, REPTTYPE, FEPFLAG) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)`

	fmt.Println("QUERY:", query)

	_, err := db.ExecContext(ctx, query, Trackrecord.SID, Trackrecord.CONTRACTNBR, Trackrecord.SUBSEQ, Trackrecord.GROUPSEQ, Trackrecord.GROUPNBR,
		Trackrecord.SUFFIXNBR, Trackrecord.HOMEPLAN, Trackrecord.IMB, Trackrecord.MATRLDIST, Trackrecord.DISPCODE, Trackrecord.MAILDATE, Trackrecord.SENTFLAG, Trackrecord.CARDTEMPLATECODE, Trackrecord.SETUPNAME, Trackrecord.REQ_DATE, Trackrecord.SEARCHCODE,
		Trackrecord.RELEASE_DT, Trackrecord.REPTTYPE, Trackrecord.FEPFLAG)
	if err != nil {
		fmt.Println("Error in executing query")
		log.Fatal(err)
	}
	fmt.Println("Added New Production card successfully")
}

func Updaterecord(Trackrecord model.Trackingsequence_archive, db *sql.DB, SID string) {
	fmt.Println("\nin GetProductionCardFromDb ", db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query := `UPDATE MANI_ORACLE_DB.Trackingsequence_archive SET SID = :1, CONTRACTNBR = :2, SUBSEQ = :3, GROUPSEQ = :4, GROUPNBR = :5,
	SUFFIXNBR = :6, HOMEPLAN = :7, IMB = :8, MATRLDIST = :9,DISPCODE = :10,MAILDATE = :11, SENTFLAG = :12, CARDTEMPLATECODE = :13, SETUPNAME = :14, REQ_DATE = :15,SEARCHCODE = :16,
	RELEASE_DT = :17, REPTTYPE = :18, FEPFLAG = :19  WHERE SID = :7`

	fmt.Println("Query:", query)

	_, err := db.ExecContext(ctx, query,
		Trackrecord.SID, Trackrecord.CONTRACTNBR, Trackrecord.SUBSEQ, Trackrecord.GROUPSEQ, Trackrecord.GROUPNBR,
		Trackrecord.SUFFIXNBR, Trackrecord.HOMEPLAN, Trackrecord.IMB, Trackrecord.MATRLDIST, Trackrecord.DISPCODE, Trackrecord.MAILDATE, Trackrecord.SENTFLAG, Trackrecord.CARDTEMPLATECODE, Trackrecord.SETUPNAME, Trackrecord.REQ_DATE, Trackrecord.SEARCHCODE,
		Trackrecord.RELEASE_DT, Trackrecord.REPTTYPE, Trackrecord.FEPFLAG, SID)
	if err != nil {
		fmt.Println("Error in executing query")
		log.Fatal(err)
	}
	fmt.Println("Record updated successfully.")
}

func DeleteSID(SID int, db *sql.DB) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query := fmt.Sprintf("DELETE FROM MANI_ORACLE_DB.Trackingsequence_archive WHERE SID = %d", SID)
	fmt.Println("Query:", query)
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("productioncard with contractnbr:", SID, "deleted successfully")
}

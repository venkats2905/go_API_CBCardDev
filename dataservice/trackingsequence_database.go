package dataservice

import (
	"database/sql"
	"db/models"
	"fmt"
	"log"

	_ "github.com/sijms/go-ora/v2"
	"golang.org/x/net/context"
)

var localDB = map[string]string{
	"service":  "xe",
	"username": "system",
	"server":   "localhost",
	"port":     "1521",
	"password": "mysecretpassword",
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

func GetTracking_SeqFromDb(trackingtemp []models.Tracking_seq) []models.Tracking_seq {

	db := ConnectToDb()
	fmt.Println("\nin GetTracking_SeqFromDb ", db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rows, err := db.QueryContext(ctx, "SELECT * FROM VISHNU_ARIKATLA_DB.TRACKINGSEQUENCE_TEMP")
	fmt.Println(rows, err)
	if err != nil {
		fmt.Println("-------------------", err)
		log.Fatal(err)
	}
	//defer rows.Close()
	for rows.Next() {
		var a models.Tracking_seq
		err := rows.Scan(&a.CONTRACTNBR, &a.SUBSEQ, &a.GROUPSEQ, &a.GROUPNBR, &a.SUFFIXNBR, &a.HOMEPLAN, &a.MATRLDIST, &a.DISPCODE, &a.SENTFLAG, &a.CARDTEMPLATECODE, &a.SETUPNAME, &a.REQ_DATE, &a.SEARCHCODE, &a.REPTTYPE)

		if err != nil {
			log.Fatal(err)
		}
		trackingtemp = append(trackingtemp, a)
		fmt.Println(trackingtemp)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return trackingtemp
}

func GetTracking_SeqByContractnbrFromDb(CONTRACTNBR string) (models.Tracking_seq, error) {
	fmt.Println("----In DB--------")
	db := ConnectToDb()
	fmt.Println("\nin VISHNU_ARIKATLA_DB.TRACKINGSEQUENCE_TEMP ", db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var tracking_seq models.Tracking_seq
	query := `SELECT * FROM VISHNU_ARIKATLA_DB.TRACKINGSEQUENCE_TEMP WHERE CONTRACTNBR = :1`
	fmt.Println("Query:", query)

	// Prepare and execute the query
	rows, err := db.QueryContext(ctx, query, CONTRACTNBR)
	if err != nil {
		return tracking_seq, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(
			&tracking_seq.CONTRACTNBR, &tracking_seq.SUBSEQ, &tracking_seq.GROUPSEQ,
			&tracking_seq.GROUPNBR, &tracking_seq.SUFFIXNBR, &tracking_seq.HOMEPLAN,
			&tracking_seq.MATRLDIST, &tracking_seq.DISPCODE, &tracking_seq.SENTFLAG,
			&tracking_seq.CARDTEMPLATECODE, &tracking_seq.SETUPNAME, &tracking_seq.REQ_DATE,
			&tracking_seq.SEARCHCODE, &tracking_seq.REPTTYPE,
		)
		if err != nil {
			return tracking_seq, err
		}
		fmt.Println("Got a particular tracking sequence from the table successfully")
	}

	return tracking_seq, nil
}

func PostAddTracking_SeqAlbumToDb(trackingtemp models.Tracking_seq) {
	db := ConnectToDb()
	fmt.Println("\nin PostAddAlbumToDb ", db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query := `INSERT INTO VISHNU_ARIKATLA_DB.TRACKINGSEQUENCE_TEMP (CONTRACTNBR, SUBSEQ, GROUPSEQ, GROUPNBR, SUFFIXNBR, HOMEPLAN, MATRLDIST, DISPCODE, SENTFLAG, CARDTEMPLATECODE, SETUPNAME, REQ_DATE, SEARCHCODE, REPTTYPE  ) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14)`

	fmt.Println("QUERY:", query)

	_, err := db.ExecContext(ctx, query,
		trackingtemp.CONTRACTNBR, trackingtemp.SUBSEQ, trackingtemp.GROUPSEQ,
		trackingtemp.GROUPNBR, trackingtemp.SUFFIXNBR, trackingtemp.HOMEPLAN,
		trackingtemp.MATRLDIST, trackingtemp.DISPCODE, trackingtemp.SENTFLAG,
		trackingtemp.CARDTEMPLATECODE, trackingtemp.SETUPNAME, trackingtemp.REQ_DATE,
		trackingtemp.SEARCHCODE, trackingtemp.REPTTYPE)
	if err != nil {
		fmt.Println("Error in executing query")
		log.Fatal(err)
	}
	fmt.Println("Added New tracking sequence successfully")
}

func DeleteTracking_SeqAlbumFromDb(contractnbr string) {
	db := ConnectToDb()
	fmt.Println("\nin DeleteTrackingSequenceFromDb ", db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query := fmt.Sprintf("DELETE FROM VISHNU_ARIKATLA_DB.TRACKINGSEQUENCE_TEMP WHERE CONTRACTNBR = %s", contractnbr)
	fmt.Println("Query:", query)

	// Execute the DELETE query
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tracking sequence with contractnbr:", contractnbr, "deleted successfully")
}

func UpdateTracking_SeqInDb(trackingtemp models.Tracking_seq) {
	db := ConnectToDb()
	fmt.Println("\nin UpdateTrackingSequenceInDb ", db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query := `UPDATE VISHNU_ARIKATLA_DB.TRACKINGSEQUENCE_TEMP SET CONTRACTNBR = :1, SUBSEQ = :2, GROUPSEQ = :3, GROUPNBR = :4, SUFFIXNBR = :5, HOMEPLAN = :6, MATRLDIST = :7, DISPCODE = :8, SENTFLAG = :9, CARDTEMPLATECODE = :10, SETUPNAME = :11, REQ_DATE = :12, SEARCHCODE = :13, REPTTYPE = :14   WHERE CONTRACTNBR = :1`

	fmt.Println("Query:", query)

	_, err := db.ExecContext(ctx, query,
		trackingtemp.CONTRACTNBR, trackingtemp.SUBSEQ, trackingtemp.GROUPSEQ,
		trackingtemp.GROUPNBR, trackingtemp.SUFFIXNBR, trackingtemp.HOMEPLAN,
		trackingtemp.MATRLDIST, trackingtemp.DISPCODE, trackingtemp.SENTFLAG,
		trackingtemp.CARDTEMPLATECODE, trackingtemp.SETUPNAME, trackingtemp.REQ_DATE,
		trackingtemp.SEARCHCODE, trackingtemp.REPTTYPE)
	if err != nil {
		fmt.Println("Error in executing query")
		log.Fatal(err)
	}
	fmt.Println("Record updated successfully.")
}

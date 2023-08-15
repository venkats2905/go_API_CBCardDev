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
	"password": "oracle",
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

func GetProductionCardFromDb(productioncards []models.Productioncard) []models.Productioncard {

	db := ConnectToDb()
	fmt.Println("\nin GetProductionCardFromDb ", db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rows, err := db.QueryContext(ctx, "SELECT * FROM ORACLEDATABASE.productioncard")
	fmt.Println(rows, err)
	if err != nil {
		fmt.Println("-------------------", err)
		log.Fatal(err)
	}
	//defer rows.Close()
	for rows.Next() {
		var a models.Productioncard
		err := rows.Scan(&a.Contractnbr, &a.Requestdate, &a.Status, &a.Statusdate, &a.Statusby,
			&a.Searchcode, &a.Cardcount, &a.Jobname, &a.Producedby, &a.Produceddate,
			&a.Scheduleddate, &a.Cardtemplatecode, &a.Groupnbr, &a.Suffixnbr, &a.Matrldist, &a.Trancd,
			&a.Reasoncd, &a.Reptype, &a.Litcode)

		if err != nil {
			log.Fatal(err)
		}
		productioncards = append(productioncards, a)
		fmt.Println(productioncards)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return productioncards
}

func GetProductioncardByContractnbrFromDb(contractnbr string) (models.Productioncard, error) {
	fmt.Println("----In DB--------")
	db := ConnectToDb()
	fmt.Println("\nin GetProductioncardByContractnbrFromDb ", db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var productioncard models.Productioncard
	query := `SELECT * FROM ORACLEDATABASE.productioncard WHERE contractnbr = :1`
	fmt.Println("Query:", query)

	// Prepare and execute the query
	rows, err := db.QueryContext(ctx, query, contractnbr)
	if err != nil {
		return productioncard, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(
			&productioncard.Contractnbr, &productioncard.Requestdate, &productioncard.Status,
			&productioncard.Statusdate, &productioncard.Statusby, &productioncard.Searchcode,
			&productioncard.Cardcount, &productioncard.Jobname, &productioncard.Producedby,
			&productioncard.Produceddate, &productioncard.Scheduleddate, &productioncard.Cardtemplatecode,
			&productioncard.Groupnbr, &productioncard.Suffixnbr, &productioncard.Matrldist,
			&productioncard.Trancd, &productioncard.Reasoncd, &productioncard.Reptype,
			&productioncard.Litcode,
		)
		if err != nil {
			return productioncard, err
		}
		fmt.Println("Got a particular production card from the table successfully")
	}

	return productioncard, nil
}

func PostAddProductionToDb(productioncards models.Productioncard) {
	db := ConnectToDb()
	fmt.Println("\nin PostAddAlbumToDb ", db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query := `INSERT INTO ORACLEDATABASE.productioncard (Contractnbr, Requestdate, Status, Statusdate, Statusby,
		Searchcode, Cardcount, Jobname, Producedby, Produceddate,
		Scheduleddate, Cardtemplatecode, Groupnbr, Suffixnbr, Matrldist, Trancd,
		Reasoncd, Reptype, Litcode) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)`

	fmt.Println("QUERY:", query)

	_, err := db.ExecContext(ctx, query,
		productioncards.Contractnbr, productioncards.Requestdate, productioncards.Status,
		productioncards.Statusdate, productioncards.Statusby, productioncards.Searchcode,
		productioncards.Cardcount, productioncards.Jobname, productioncards.Producedby,
		productioncards.Produceddate, productioncards.Scheduleddate, productioncards.Cardtemplatecode,
		productioncards.Groupnbr, productioncards.Suffixnbr, productioncards.Matrldist,
		productioncards.Trancd, productioncards.Reasoncd, productioncards.Reptype,
		productioncards.Litcode)
	if err != nil {
		fmt.Println("Error in executing query")
		log.Fatal(err)
	}
	fmt.Println("Added New Production card successfully")
}

func DeleteAlbumFromDb(contractnbr string) {
	db := ConnectToDb()
	fmt.Println("\nin GetProductionCardFromDb ", db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query := fmt.Sprintf("DELETE FROM ORACLEDATABASE.productioncard WHERE contractnbr = %s", contractnbr)
	fmt.Println("Query:", query)

	// Execute the DELETE query
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("productioncard with contractnbr:", contractnbr, "deleted successfully")
}

func UpdateProductionInDb(productioncards models.Productioncard) {
	db := ConnectToDb()
	fmt.Println("\nin GetProductionCardFromDb ", db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query := `UPDATE ORACLEDATABASE.productioncard SET Contractnbr = :1, Requestdate = :2, Status = :3, Statusdate = :4, Statusby = :5,
	Searchcode = :6, Cardcount = :7, Jobname = :8, Producedby = :9, Produceddate = :10,
	Scheduleddate = :11, Cardtemplatecode = :12, Groupnbr = :13, Suffixnbr = :14, Matrldist = :15, Trancd = :16,
	Reasoncd = :17, Reptype = :18, Litcode = :19  WHERE Contractnbr = :1`

	fmt.Println("Query:", query)

	_, err := db.ExecContext(ctx, query,
		productioncards.Contractnbr, productioncards.Requestdate, productioncards.Status,
		productioncards.Statusdate, productioncards.Statusby, productioncards.Searchcode,
		productioncards.Cardcount, productioncards.Jobname, productioncards.Producedby,
		productioncards.Produceddate, productioncards.Scheduleddate, productioncards.Cardtemplatecode,
		productioncards.Groupnbr, productioncards.Suffixnbr, productioncards.Matrldist,
		productioncards.Trancd, productioncards.Reasoncd, productioncards.Reptype,
		productioncards.Litcode)
	if err != nil {
		fmt.Println("Error in executing query")
		log.Fatal(err)
	}
	fmt.Println("Record updated successfully.")
}

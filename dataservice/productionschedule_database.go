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

func GetProductionScheduleFromDb(productionschedules []models.Productionschedule) []models.Productionschedule {

	db := ConnectToDb()
	fmt.Println("\nin GetProductionScheduleFromDb ", db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rows, err := db.QueryContext(ctx, "SELECT * FROM A1B2C#.productionschedule")
	fmt.Println(rows, err)
	if err != nil {
		fmt.Println("-------------------", err)
		log.Fatal(err)
	}
	//defer rows.Close()
	for rows.Next() {
		var a models.Productionschedule
		err := rows.Scan(&a.Proddate, &a.Jobname, &a.Jobqueue, &a.Fullpathname, &a.Cardcount,
			&a.Carriercount, &a.Requestdate, &a.Sentdate, &a.Completedate, &a.Issues,
			&a.Status, &a.Filesequence, &a.Completedby, &a.Veridiedby, &a.Notes, &a.Embossedcards,
			&a.Embossedcarriers, &a.Heldcards, &a.Heldcarriers)

		if err != nil {
			log.Fatal(err)
		}
		productionschedules = append(productionschedules, a)
		fmt.Println(productionschedules)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return productionschedules
}

func GetProductionscheduleByJobnameFromDb(jobname string) (models.Productionschedule, error) {
	fmt.Println("----In DB--------")
	db := ConnectToDb()
	fmt.Println("\nin GetProductionscheduleByJobnameFromDb ", db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var productionschedule models.Productionschedule
	query := `SELECT * FROM A1B2C#.productionschedule WHERE Jobname = :2`
	fmt.Println("Query:", query)

	// Prepare and execute the query
	rows, err := db.QueryContext(ctx, query, jobname)
	if err != nil {
		return productionschedule, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(
			&productionschedule.Proddate, &productionschedule.Jobname, &productionschedule.Jobqueue,
			&productionschedule.Fullpathname, &productionschedule.Cardcount, &productionschedule.Carriercount,
			&productionschedule.Requestdate, &productionschedule.Sentdate, &productionschedule.Completedate,
			&productionschedule.Issues, &productionschedule.Status, &productionschedule.Filesequence,
			&productionschedule.Completedby, &productionschedule.Veridiedby, &productionschedule.Notes,
			&productionschedule.Embossedcards, &productionschedule.Embossedcarriers, &productionschedule.Heldcards,
			&productionschedule.Heldcarriers,
		)

		if err != nil {
			return productionschedule, err
		}
		fmt.Println("Got a particular production schedule from the table successfully")
	}

	return productionschedule, nil
}

func PostAddProductionscheduleToDb(productionschedule models.Productionschedule) {
	db := ConnectToDb()
	fmt.Println("\nin PostAddAlbumToDb ", db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query := `INSERT INTO A1B2C#.Productionschedule (Proddate, Jobname, Jobqueue, Fullpathname, Cardcount,
		Carriercount, Requestdate, Sentdate, Completedate, Issues,
		Status, Filesequence, Completedby, Veridiedby, Notes, Embossedcards,
		Embossedcarriers, Heldcards, Heldcarriers) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)`

	fmt.Println("QUERY:", query)

	_, err := db.ExecContext(ctx, query,
		&productionschedule.Proddate, &productionschedule.Jobname, &productionschedule.Jobqueue,
		&productionschedule.Fullpathname, &productionschedule.Cardcount, &productionschedule.Carriercount,
		&productionschedule.Requestdate, &productionschedule.Sentdate, &productionschedule.Completedate,
		&productionschedule.Issues, &productionschedule.Status, &productionschedule.Filesequence,
		&productionschedule.Completedby, &productionschedule.Veridiedby, &productionschedule.Notes,
		&productionschedule.Embossedcards, &productionschedule.Embossedcarriers, &productionschedule.Heldcards,
		&productionschedule.Heldcarriers)
	if err != nil {
		fmt.Println("Error in executing query")
		log.Fatal(err)
	}
	fmt.Println("Added New Production card successfully")
}

func DeletescheduleAlbumFromDb(jobname string) {
	db := ConnectToDb()
	fmt.Println("\nin GetProductionScheduleFromDb ", db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query := fmt.Sprintf("DELETE FROM A1B2C#.Productionschedule WHERE jobname = '%s'", jobname)
	fmt.Println("Query:", query)

	// Execute the DELETE query
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("productionschedule with jobname:", jobname, "deleted successfully")
}

func UpdateProductionscheduleInDb(productionschedule models.Productionschedule) {
	db := ConnectToDb()
	fmt.Println("\nin GetProductionScheduleFromDb ", db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query := `UPDATE A1B2C#.Productionschedule SET Proddate = :1, Jobname = :2, Jobqueue = :3, Fullpathname = :4, Cardcount = :5,
	Carriercount = :6, Requestdate = :7, Sentdate = :8, Completedate = :9, Issues = :10,
	Status = :11, Filesequence = :12, Completedby = :13, Veridiedby = :14, Notes = :15, Embossedcards = :16,
	Embossedcarriers = :17, Heldcards = :18, Heldcarriers = :19  WHERE Jobname = :2`

	fmt.Println("Query:", query)

	_, err := db.ExecContext(ctx, query,
		&productionschedule.Proddate, &productionschedule.Jobname, &productionschedule.Jobqueue,
		&productionschedule.Fullpathname, &productionschedule.Cardcount, &productionschedule.Carriercount,
		&productionschedule.Requestdate, &productionschedule.Sentdate, &productionschedule.Completedate,
		&productionschedule.Issues, &productionschedule.Status, &productionschedule.Filesequence,
		&productionschedule.Completedby, &productionschedule.Veridiedby, &productionschedule.Notes,
		&productionschedule.Embossedcards, &productionschedule.Embossedcarriers, &productionschedule.Heldcards,
		&productionschedule.Heldcarriers)
	if err != nil {
		fmt.Println("Error in executing query")
		log.Fatal(err)
	}
	fmt.Println("Record updated successfully.")
}

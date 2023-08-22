package dataservice

import (
	"db/models"
	"fmt"
	"log"

	_ "github.com/sijms/go-ora/v2"
	"golang.org/x/net/context"
)

// var localDB = map[string]string{
// 	"service":  "xe",
// 	"username": "system",
// 	"server":   "localhost",
// 	"port":     "1521",
// 	"password": "varshith",
// }

// func ConnectToDb() *sql.DB {
// 	connectionString := "oracle://" + localDB["username"] + ":" + localDB["password"] + "@" + localDB["server"] + ":" + localDB["port"] + "/" + localDB["service"]
// 	db, err := sql.Open("oracle", connectionString)
// 	if err != nil {
// 		log.Fatal("Error connecting to the database:", err)
// 	}
// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatal("Error pinging the database:", err)
// 	}

// 	fmt.Println("Connected to the Oracle database!")
// 	fmt.Println("\n------", db)
// 	return db
// }

func GetCardDataFromDb(cards []models.Carddetails) []models.Carddetails {

	db := ConnectToDb()
	fmt.Println("\nin GetCardDataFromDb ", db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rows, err := db.QueryContext(ctx, "SELECT * FROM carddata.carddetails")
	fmt.Println(rows, err)
	if err != nil {
		fmt.Println("-------------------", err)
		log.Fatal(err)
	}
	//defer rows.Close()
	for rows.Next() {
		var a models.Carddetails

		err := rows.Scan(
			&a.Cardrevisionid,
			&a.Cardtemplatecode,
			&a.Comments,
			&a.Stockcode,
			&a.Effectivedate,
			&a.Carriertemplatecode,
			&a.Enddate,
			&a.Status,
			&a.Statusdate,
			&a.Creationdate,
			&a.Creationuserid,
			&a.Lastupdates,
			&a.Lastupdateuserid,
		)

		if err != nil {
			log.Fatal(err)
		}
		cards = append(cards, a)
		fmt.Println(cards)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return cards
}

func GetCardDataByCardrevisionidFromDb(Cardrevisionid int) (models.Carddetails, error) {
	fmt.Println("----In DB--------")
	db := ConnectToDb()
	fmt.Println("\nin GetCardDataByCardrevisionidFromDb ", db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var card models.Carddetails
	// query := `SELECT * FROM carddata.productioncard WHERE Cardrevisionid = ?`
	query := fmt.Sprintf("SELECT * FROM carddata.carddetails WHERE Cardrevisionid = %d", Cardrevisionid)
	fmt.Println("Query:", query)

	// Prepare and execute the query
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Error here <-------")
		return card, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(
			&card.Cardrevisionid,
			&card.Cardtemplatecode,
			&card.Comments,
			&card.Stockcode,
			&card.Effectivedate,
			&card.Carriertemplatecode,
			&card.Enddate,
			&card.Status,
			&card.Statusdate,
			&card.Creationdate,
			&card.Creationuserid,
			&card.Lastupdates,
			&card.Lastupdateuserid,
		)
		if err != nil {
			return card, err
		}
		fmt.Println("Got a particular card from the table successfully")
	}

	return card, nil
}

func PostAddCardToDb(cards models.Carddetails) {
	fmt.Println("-----------------------")
	db := ConnectToDb()
	fmt.Println("\nin PostAddAlbumToDb ", db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query := `INSERT INTO carddata.carddetails (CARDREVISIONID, CARDTEMPLATECODE, COMMENTS, STOCKCODE, EFFECTIVEDATE, 
		CARRIERTEMPLATECODE, ENDDATE, STATUS, STATUSDATE, CREATIONDATE, CREATIONUSERID, LASTUPDATES, LASTUPDATEUSERID)
	 VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13)`

	fmt.Println("QUERY:", query)

	_, err := db.ExecContext(ctx, query,
		&cards.Cardrevisionid,
		&cards.Cardtemplatecode,
		&cards.Comments,
		&cards.Stockcode,
		&cards.Effectivedate,
		&cards.Carriertemplatecode,
		&cards.Enddate,
		&cards.Status,
		&cards.Statusdate,
		&cards.Creationdate,
		&cards.Creationuserid,
		&cards.Lastupdates,
		&cards.Lastupdateuserid)
	if err != nil {
		fmt.Println("Error in executing query")
		log.Fatal(err)
	}
	fmt.Println("Added New card successfully")
}

func UpdateCardInDb(cards models.Carddetails) {
	db := ConnectToDb()
	fmt.Println("\nin GetCardDataFromDb ", db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query := `UPDATE carddata.carddetails SET CARDREVISIONID = :1, CARDTEMPLATECODE = :2, COMMENTS = :3, STOCKCODE = :4,
	EFFECTIVEDATE = :5, CARRIERTEMPLATECODE = :6, ENDDATE = :7, STATUS = :8,
	STATUSDATE = :9, CREATIONDATE = :10, CREATIONUSERID = :11, LASTUPDATES = :12,
	LASTUPDATEUSERID = :13 WHERE CARDREVISIONID = :1`

	fmt.Println("Query:", query)

	_, err := db.ExecContext(ctx, query,
		cards.Cardrevisionid,
		cards.Cardtemplatecode,
		cards.Comments,
		cards.Stockcode,
		cards.Effectivedate,
		cards.Carriertemplatecode,
		cards.Enddate,
		cards.Status,
		cards.Statusdate,
		cards.Creationdate,
		cards.Creationuserid,
		cards.Lastupdates,
		cards.Lastupdateuserid)
	if err != nil {
		fmt.Println("Error in executing query")
		log.Fatal(err)
	}
	fmt.Println("Record updated successfully.")
}

func DeleteCardFromDb(cardrevisionid int) {
	db := ConnectToDb()
	fmt.Println("\nin GetProductionCardFromDb ", db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query := fmt.Sprintf("DELETE FROM carddata.carddetails WHERE cardrevisionid = %d", cardrevisionid)
	fmt.Println("Query:", query)

	// Execute the DELETE query
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("productioncard with cardrevisionid:", cardrevisionid, "deleted successfully")
}

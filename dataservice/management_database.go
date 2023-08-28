package dataservice

import (
	"fmt"
	"go_API_CBCardDev/models"
	"log"

	_ "github.com/sijms/go-ora/v2"
	"golang.org/x/net/context"
)

func GetManagementFromDb(management []models.Management) []models.Management {

	db := ConnectToDb()
	fmt.Println("in GetManagementFromDb ", db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rows, err := db.QueryContext(ctx, "SELECT * FROM MY_DATABASE.stock_management")
	fmt.Println(&rows, err, "iM HERE")
	if err != nil {
		fmt.Println("-------------------", err)
		log.Fatal(err)
	}
	//defer rows.Close()
	for rows.Next() {
		var a models.Management
		err := rows.Scan(&a.StockCode, &a.EffectiveDate, &a.CarrierTemplateCode, &a.EndDate, &a.Status, &a.StatusDate,
			&a.CreationDate, &a.CreationUserID, &a.LastUpdates, &a.LastUpdatesUserID)

		if err != nil {
			log.Fatal(err)
		}

		management = append(management, a)
		fmt.Println(management)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return management
}

func GetManagementBystockcodeFromDb(stockcode string) (models.Management, error) {
	fmt.Println("----In DB--------")
	db := ConnectToDb()
	fmt.Println("\nin GetManagementBystockcodeFromDb ", db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var Management models.Management
	query := `SELECT * FROM MY_DATABASE.stock_management  WHERE stockcode = :1`
	fmt.Println("Query:", query)

	// Prepare and execute the query
	rows, err := db.QueryContext(ctx, query, stockcode)
	if err != nil {
		return Management, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(
			&Management.StockCode, &Management.EffectiveDate, &Management.CarrierTemplateCode,
			&Management.EndDate, &Management.Status, &Management.StatusDate,
			&Management.CreationDate, &Management.CreationUserID, &Management.LastUpdates,
			&Management.LastUpdatesUserID,
		)
		if err != nil {
			return Management, err
		}
		fmt.Println("Got a particular production card from the table successfully")
	}

	return Management, nil
}

func PostAddManagementToDb(Management models.Management) {
	db := ConnectToDb()
	fmt.Println("\nin PostAddAlbumToDb ", db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query := `INSERT INTO MY_DATABASE.stock_management  (Stockcode, Effective_date, Carrier_template_code, End_date, Status, Status_date, 
		Creation_Date, Creation_User_ID, LastUpdates, LastUpdates_User_ID) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10)`

	fmt.Println("QUERY:", query)

	_, err := db.ExecContext(ctx, query,
		&Management.StockCode, &Management.EffectiveDate, &Management.CarrierTemplateCode,
		&Management.EndDate, &Management.Status, &Management.StatusDate,
		&Management.CreationDate, &Management.CreationUserID, &Management.LastUpdates,
		&Management.LastUpdatesUserID)
	if err != nil {
		fmt.Println("Error in executing query")
		log.Fatal(err)
	}
	fmt.Println("Added New Management successfully")
}

func DeleteManagementFromDb(stockcode string) {
	db := ConnectToDb()
	fmt.Println("\nin GetManagementFromDb ", db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query := fmt.Sprintf("DELETE FROM MY_DATABASE.stock_management WHERE stockcode = %s", stockcode)
	fmt.Println("Query:", query)

	// Execute the DELETE query
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Management with stockcode:", stockcode, "deleted successfully")
}

func UpdateManagementInDb(Management models.Management) {
	db := ConnectToDb()
	fmt.Println("\nin GetManagementFromDb ", db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query := `UPDATE MY_DATABASE.stock_management  SET Stockcode=:1, Effective_date=:2, Carrier_template_code=:3, End_date=:4, Status=:5, Status_date=:6, 
	Creation_Date=:7, Creation_User_ID=:8, LastUpdates=:9, LastUpdates_User_ID=:10 WHERE stockcode = :1`

	fmt.Println("Query:", query)

	_, err := db.ExecContext(ctx, query,
		&Management.StockCode, &Management.EffectiveDate, &Management.CarrierTemplateCode,
		&Management.EndDate, &Management.Status, &Management.StatusDate,
		&Management.CreationDate, &Management.CreationUserID, &Management.LastUpdates,
		&Management.LastUpdatesUserID)
	if err != nil {
		fmt.Println("Error in executing query")
		log.Fatal(err)
	}
	fmt.Println("Record updated successfully.")
}

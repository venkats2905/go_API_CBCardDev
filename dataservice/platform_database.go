package dataservice

import (
	"db/models"
	//"database/sql"
	"fmt"
	"log"

	_ "github.com/sijms/go-ora/v2"
	"golang.org/x/net/context"
)

func GetPlatformFromDb(platforms []models.Platform) []models.Platform {

	db := ConnectToDb()
	fmt.Println("\nin GetPlatformFromDb ", db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rows, err := db.QueryContext(ctx, "SELECT * FROM Mounika_ORACLE_DB.Platform")
	fmt.Println(rows, err)
	if err != nil {
		fmt.Println("-------------------", err)
		log.Fatal(err)
	}
	//defer rows.Close()
	for rows.Next() {
		var a models.Platform
		err := rows.Scan(
			&a.Idtype,
			&a.Priority,
			&a.Description,
			&a.Plasticstock,
			&a.Carrierstock,
			&a.Envelopestock,
			&a.Setup,
			&a.Fmodule1,
			&a.Fmodule2,
			&a.Fmodule3,
			&a.Bmodule1,
			&a.Bmodule2,
			&a.Jobgroup,
			&a.Defaultqueue,
			&a.Platformcode,
		)

		if err != nil {
			log.Fatal(err)
		}
		platforms = append(platforms, a)
		fmt.Println(platforms)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return platforms
}

func GetPlatformByIdtypeFromDb(idtype int) (models.Platform, error) {
	fmt.Println("----In DB--------")
	db := ConnectToDb()
	fmt.Println("\nin GetPlatformByIdtypeFromDb ", db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var platform models.Platform
	query := `SELECT * FROM Mounika_ORACLE_DB.platform WHERE idtype = :1`
	fmt.Println("Query:", query)

	// Prepare and execute the query
	rows, err := db.QueryContext(ctx, query, idtype)
	if err != nil {
		return platform, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(
			&platform.Idtype,
			&platform.Priority,
			&platform.Description,
			&platform.Plasticstock,
			&platform.Carrierstock,
			&platform.Envelopestock,
			&platform.Setup,
			&platform.Fmodule1,
			&platform.Fmodule2,
			&platform.Fmodule3,
			&platform.Bmodule1,
			&platform.Bmodule2,
			&platform.Jobgroup,
			&platform.Defaultqueue,
			&platform.Platformcode,
		)

		if err != nil {
			return platform, err
		}
		fmt.Println("Got a particular platform from the table successfully")
	}

	return platform, nil
}

func PostAddPlatformToDb(platforms models.Platform) {
	db := ConnectToDb()
	fmt.Println("\nin PostAddAlbumToDb ", db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query := `INSERT INTO Mounika_ORACLE_DB.platform(
		Idtype, Priority, Description, Plastickstock, carrierstock,
		Envelopestock, Setup, Fmodule1, Fmodule2,Fmodule3,Bmodule1,Bmodule2,Jobgroup,DefaultQueue,Platformcode
		) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15)`

	fmt.Println("QUERY:", query)

	_, err := db.ExecContext(ctx, query,
		platforms.Idtype, platforms.Priority, platforms.Description,
		platforms.Plasticstock, platforms.Carrierstock, platforms.Envelopestock,
		platforms.Setup, platforms.Fmodule1, platforms.Fmodule2,
		platforms.Fmodule3, platforms.Bmodule1, platforms.Bmodule2,
		platforms.Jobgroup, platforms.Defaultqueue, platforms.Platformcode)

	if err != nil {
		fmt.Println("Error in executing query")
		log.Fatal(err)
	}
	fmt.Println("Added New Platform successfully")
}

func DeletePlatformAlbumFromDb(idtype int) {
	db := ConnectToDb()
	fmt.Println("\nin GetPlatformFromDb ", db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query := fmt.Sprintf("DELETE FROM Mounika_ORACLE_DB.platform WHERE idtype = %d", idtype)
	fmt.Println("Query:", query)

	// Execute the DELETE query
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("platform with idtype:", idtype, "deleted successfully")
}

func UpdatePlatformInDb(platforms models.Platform) {
	db := ConnectToDb()
	fmt.Println("\nin GetPlatformFromDb ", db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query := `UPDATE Mounika_ORACLE_DB.platform SET
     Idtype= :1, Priority = :2, Descrition= :3,
   Plasticstock = :4,Carrierstock = :5, Envelopestock = :6,
    Setup = :7, Fmodule1 = :8, Fmodule2 = :9,
 Fmodule3 = :10, Bmodule1 = :11, Bmodule2 = :12,
    Jobgroup = :13, Defaultqueue = :14, Platformcode= :15,WHERE idtype = :1`

	fmt.Println("Query:", query)

	_, err := db.ExecContext(ctx, query, platforms.Idtype, platforms.Priority, platforms.Description,
		platforms.Plasticstock, platforms.Carrierstock, platforms.Envelopestock, platforms.Setup,
		platforms.Fmodule1, platforms.Fmodule2, platforms.Fmodule3,
		platforms.Bmodule1, platforms.Bmodule2, platforms.Jobgroup,
		platforms.Defaultqueue, platforms.Platformcode,
	)
	if err != nil {
		fmt.Println("Error in executing query")
		log.Fatal(err)
	}
	fmt.Println("Record updated successfully.")

}

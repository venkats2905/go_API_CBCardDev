package dataservice

import (
	"db/models"
	"fmt"
	"log"

	_ "github.com/sijms/go-ora/v2"
	"golang.org/x/net/context"
)

func GetGraphicCardFromDb(graphics []models.Graphic) []models.Graphic {

	db := ConnectToDb()
	fmt.Println("\nin GetGraphicCardFromDb ", db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rows, err := db.QueryContext(ctx, "SELECT * FROM MANI_ORACLE_DB.graphic")
	fmt.Println(rows, err)
	if err != nil {
		fmt.Println("-------------------", err)
		log.Fatal(err)
	}
	//defer rows.Close()
	for rows.Next() {
		var a models.Graphic
		err := rows.Scan(
			&a.GraphicName,
			&a.GraphicRevisionNbr,
			&a.GraphicDesc,
			&a.GraphicFileName,
			&a.GraphicImage,
			&a.DefaultTopPos,
			&a.DefaultLeftPos,
			&a.DefaultColorCode,
			&a.DefaultLocationCode,
			&a.DefaultHeight,
			&a.DefaultWidth,
			&a.EffectiveDate,
			&a.EndDate,
			&a.LastUpdateTs,
			&a.LastUpdateUserId,
			&a.VertOffset,
			&a.HorizOffset,
		)

		if err != nil {
			log.Fatal(err)
		}
		graphics = append(graphics, a)
		fmt.Println(graphics)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return graphics
}

func GetGraphicByGraphicrevisionnbrFromDb(graphicrevisionnbr int) (models.Graphic, error) {
	fmt.Println("----In DB--------")
	db := ConnectToDb()
	fmt.Println("\nin GetGraphicByGraphicrevisionnbrFromDb ", db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var graphiccard models.Graphic
	query := `SELECT * FROM MANI_ORACLE_DB.graphic WHERE graphicrevisionnbr = :1`
	fmt.Println("Query:", query)

	// Prepare and execute the query
	rows, err := db.QueryContext(ctx, query, graphicrevisionnbr)
	if err != nil {
		return graphiccard, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(
			&graphiccard.GraphicName,
			&graphiccard.GraphicRevisionNbr,
			&graphiccard.GraphicDesc,
			&graphiccard.GraphicFileName,
			&graphiccard.GraphicImage,
			&graphiccard.DefaultTopPos,
			&graphiccard.DefaultLeftPos,
			&graphiccard.DefaultColorCode,
			&graphiccard.DefaultLocationCode,
			&graphiccard.DefaultHeight,
			&graphiccard.DefaultWidth,
			&graphiccard.EffectiveDate,
			&graphiccard.EndDate,
			&graphiccard.LastUpdateTs,
			&graphiccard.LastUpdateUserId,
			&graphiccard.VertOffset,
			&graphiccard.HorizOffset,
		)

		if err != nil {
			return graphiccard, err
		}
		fmt.Println("Got a particular graphic card from the table successfully")
	}

	return graphiccard, nil
}

func PostAddGraphicToDb(graphics models.Graphic) {
	db := ConnectToDb()
	fmt.Println("\nin PostAddAlbumToDb ", db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query := `INSERT INTO MANI_ORACLE_DB.graphic (
		GraphicName, GraphicRevisionNbr, GraphicDesc, GraphicFileName, GraphicImage,
		DefaultTopPos, DefaultLeftPos, DefaultColorCode, DefaultLocationCode,
		DefaultHeight, DefaultWidth, EffectiveDate, EndDate, LastUpdateTs, LastUpdateUserId,
		VertoOffset, HorizOffset) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)`

	fmt.Println("QUERY:", query)

	_, err := db.ExecContext(ctx, query,
		graphics.GraphicName, graphics.GraphicRevisionNbr, graphics.GraphicDesc,
		graphics.GraphicFileName, graphics.GraphicImage, graphics.DefaultTopPos,
		graphics.DefaultLeftPos, graphics.DefaultColorCode, graphics.DefaultLocationCode,
		graphics.DefaultHeight, graphics.DefaultWidth, graphics.EffectiveDate,
		graphics.EndDate, graphics.LastUpdateTs, graphics.LastUpdateUserId,
		graphics.VertOffset, graphics.HorizOffset,
	)
	if err != nil {
		fmt.Println("Error in executing query")
		log.Fatal(err)
	}
	fmt.Println("Added New Graphic card successfully")
}

func DeleteGraphicAlbumFromDb(graphicrevisionnbr int) {
	db := ConnectToDb()
	fmt.Println("\nin GetGraphicCardFromDb ", db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query := fmt.Sprintf("DELETE FROM MANI_ORACLE_DB.graphic WHERE graphicrevisionnbr = %d", graphicrevisionnbr)
	fmt.Println("Query:", query)

	// Execute the DELETE query
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("graphiccard with graphicrevisionnbr:", graphicrevisionnbr, "deleted successfully")
}

func UpdateGraphicInDb(graphics models.Graphic) {
	db := ConnectToDb()
	fmt.Println("\nin GetGraphicCardFromDb ", db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query := `UPDATE MANI_ORACLE_DB.graphic SET
    GraphicName = :1, GraphicRevisionNbr = :2, GraphicDesc = :3,
    GraphicFileName = :4, GraphicImage = :5, DefaultTopPos = :6,
    DefaultLeftPos = :7, DefaultColorCode = :8, DefaultLocationCode = :9,
    DefaultHeight = :10, DefaultWidth = :11, EffectiveDate = :12,
    EndDate = :13, LastUpdateTs = :14, LastUpdateUserId = :15,
    VertoOffset = :16, HorizOffset = :17 WHERE GraphicRevisionNbr = :2`

	fmt.Println("Query:", query)

	_, err := db.ExecContext(ctx, query,
		graphics.GraphicName, graphics.GraphicRevisionNbr, graphics.GraphicDesc,
		graphics.GraphicFileName, graphics.GraphicImage, graphics.DefaultTopPos,
		graphics.DefaultLeftPos, graphics.DefaultColorCode, graphics.DefaultLocationCode,
		graphics.DefaultHeight, graphics.DefaultWidth, graphics.EffectiveDate,
		graphics.EndDate, graphics.LastUpdateTs, graphics.LastUpdateUserId,
		graphics.VertOffset, graphics.HorizOffset,
	)
	if err != nil {
		fmt.Println("Error in executing query")
		log.Fatal(err)
	}
	fmt.Println("Record updated successfully.")
}

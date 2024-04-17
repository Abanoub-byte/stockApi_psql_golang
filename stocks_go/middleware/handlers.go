package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/bob/stocks_go/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func createConnection() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")

	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("connected successfully to the postgres")
	return db

}
func GetStock(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("unable to convert the string into int. %v ", err)
	}

	stock, err := getStock(int64(id))
	if err != nil {
		log.Fatalf("unbale to get the stock %v ", err)
	}

	json.NewEncoder(res).Encode(stock)
}

func GetAllStock(res http.ResponseWriter, req *http.Request) {

	stocks, err := getAllStocks()

	if err != nil {
		log.Fatalf("unbable to get all the stocks %v ", err)
	}

	json.NewEncoder(res).Encode(stocks)
}

func CreateStock(res http.ResponseWriter, req *http.Request) {

	var stock models.Stock
	err := json.NewDecoder(req.Body).Decode(&stock)

	json.NewDecoder(req.Body).Decode(&stock)
	if err != nil {
		log.Fatalf("Unable to decode the request body %v", err)
	}
	insertID := insertStock(stock)

	ress := Response{
		ID:      insertID,
		Message: "stock created successfully",
	}
	json.NewEncoder(res).Encode(ress)

}

func UpdateStock(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int %v ", err)
	}
	var stock models.Stock

	err = json.NewDecoder(req.Body).Decode(&stock)

	if err != nil {
		log.Fatalf("unable to decode the request body %v ", err)
	}
	updatedRows := updateStock(int64(id), stock)

	msg := fmt.Sprintf("stock updated %v ", updatedRows)
	ress := Response{
		ID:      int64(id),
		Message: msg,
	}
	json.NewEncoder(res).Encode(ress)
}

func DeletStock(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("unable to get the id %v ", err)
	}
	deletedRows := deleteStock(int64(id))

	msg := fmt.Sprintf("stock deleted successfully. total rows/records %v", deletedRows)
	ress := Response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(res).Encode(ress)

}
func insertStock(stock models.Stock) int64 {

	db := createConnection()
	defer db.Close()

	sqlStatement := `INSERT INTO stocks(name, price, company) VALUES($1, $2, $3) RETURNING stock_id`

	var id int64
	err := db.QueryRow(sqlStatement, stock.Name, stock.Price, stock.Company).Scan(&id)

	if err != nil {
		log.Fatalf("error creating database connection: %v ", err)
	}

	fmt.Printf("inserted a single record %v", id)
	return id
}

func getStock(id int64) (models.Stock, error) {

	db := createConnection()
	defer db.Close()
	var stock models.Stock

	sqlStatement := `SELECT * FROM stocks WHERE stock_id=$1`

	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&stock.StockId, &stock.Name, &stock.Price, &stock.Company)
	switch err {
	case sql.ErrNoRows:
		fmt.Print("no rows were returned!")
		return stock, err
	case nil:
		return stock, nil
	default:
		return stock, err

	}
}

func updateStock(id int64, stock models.Stock) int64 {

	db := createConnection()

	defer db.Close()
	sqlStatement := `UPDATE stocks SET name =$2, price =$3, company = $4 WHERE stock_id =$1`
	res, err := db.Exec(sqlStatement, id, stock.Name, stock.Price, stock.Company)

	if err != nil {
		log.Fatalf("unable to execute the query %v ", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("errors while checking the affectd rows. %v ", err)

	}
	fmt.Printf("total rows affected %v ", rowsAffected)
	return rowsAffected

}

func deleteStock(id int64) int64 {

	db := createConnection()
	defer db.Close()

	sqlStatement := `DELETE FROM stocks WHERE stock_id=$1`
	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		log.Fatalf("Unable to execute query %v", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("total rows affected %v ", rowsAffected)
	return rowsAffected
}

func getAllStocks() ([]models.Stock, error) {

	db := createConnection()
	defer db.Close()
	var stocks []models.Stock

	sqlStatement := `SELECT * FROM stocks`
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("unable to excute query, %v ", err)
	}

	defer rows.Close()

	for rows.Next() {

		var stock models.Stock
		err = rows.Scan(&stock.StockId, &stock.Name, &stock.Price, &stock.Company)

		if err != nil {
			log.Fatalf("unable to scan the row %v", err)
		}
		stocks = append(stocks, stock)

	}
	return stocks, err
}

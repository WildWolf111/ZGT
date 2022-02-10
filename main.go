package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

var (
	port string = "8080"
)

//Struct for representation total slice
// First Level ob JSON object Parsing
type Companies struct {
	Companies []Company `json:"companies"`
}

type Warehouses struct {
	Warehouses []Warehouse `json:"warehouses"`
}

type WarehousesCells struct {
	WarehousesCells []WarehouseCell `json:"warehouses_cells"`
}

//Internal user representation
//Second level of object JSON parsin
type Company struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	INN  uint64 `json:"inn"`
	KPP  uint64 `json:"kpp"`
}

type Warehouse struct {
	ID      uint64  `json:"id"`
	Name    string  `json:"name"`
	Slug    string  `json:"slug"`
	Company Company `json:"company"`
	Address string  `json:"kpp"`
}

type Stok struct {
	ID            uint64        `json:"id"`
	Sender        Company       `json:"sender"`
	Recipient     Company       `json:"recipient"`
	Product       Product       `json:"product"`
	Quantity      uint64        `json:"inn"`
	WarehouseCell WarehouseCell `json:"warehouse_cell"`
	GTD           GTD           `json:"gtd"`
}
type Product struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type WarehouseCell struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"inn"`
	Warehouse Warehouse `json:"warehouse"`
}

type GTD struct {
	ID      uint64  `json:"id"`
	Country Country `json:"country"`
	Number  string  `json:"slug"`
}

type Country struct {
	ID      uint64 `json:"id"`
	Code    uint64 `json:"name"`
	Country string `json:"slug"`
}
type ErrorMessage struct {
	Message string `json:"message"`
}

var (
	db []Company
)

//Функция для распечатывания User
func PrintCompany(c *Company) {
	fmt.Printf("ID: %d\n", c.ID)
	fmt.Printf("Name: %s\n", c.Name)
	fmt.Printf("Slug: %s\n", c.Slug)
	fmt.Printf("INN: %d\n", c.INN)
	fmt.Printf("KPP: %d\n", c.KPP)
}

func PrintWarehoouse(w *Warehouse) {
	fmt.Printf("ID: %d\n", w.ID)
	fmt.Printf("Name: %s\n", w.Name)
	fmt.Printf("Slug: %s\n", w.Slug)
	fmt.Printf("Company: %s,%d,%s\n", w.Company.Name, w.Company.ID, w.Company.Slug)
	fmt.Printf("Address: %s\n", w.Address)
}

func FindCompanyById(id uint64) (Company, bool) {
	var company Company
	var found bool
	for _, p := range db {
		if p.ID == id {
			company = p
			found = true
			break
		}
	}

	return company, found
}

func SearchCompanies(str string) ([]Company, bool) {
	var companies []Company
	var found bool

	var sel := "SELECT * FROM companies WHERE slug LIKE str"
	result = db.query(sel)

	for _, p := range db {
		if strings.Contains(p.Slug, str) {
			companies = append(companies, p)
			found = true
		}
	}

	return companies, found
}

//1. Рассмотрим процесс десериализации (то есть когда из последовательности в объект)

func GetAllCompanies(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(db)
}

func GetFoundCompany(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)

	if vars["str"] == "" {
		writer.WriteHeader(400)
		error := ErrorMessage{
			Message: "Not string fo search",
		}
		json.NewEncoder(writer).Encode(error)
		return
	}

	companies, ok := SearchCompanies(vars["str"])

	if ok {
		writer.WriteHeader(200)
		json.NewEncoder(writer).Encode(companies)
		return
	} else {
		writer.WriteHeader(404)
		error := ErrorMessage{
			Message: "Companies not found",
		}
		json.NewEncoder(writer).Encode(error)
		return
	}
}

func GetCompanyById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	id, err := strconv.ParseUint(vars["id"], 0, 64)
	if err != nil {
		writer.WriteHeader(400)
		error := ErrorMessage{
			Message: "Id is not integer",
		}
		json.NewEncoder(writer).Encode(error)
		return
	}

	company, ok := FindCompanyById(id)

	if ok {
		writer.WriteHeader(200)
		json.NewEncoder(writer).Encode(company)
		return
	} else {
		writer.WriteHeader(404)
		error := ErrorMessage{
			Message: "Company not found",
		}
		json.NewEncoder(writer).Encode(error)
		return
	}

}

func init() {

	var companies Companies
	jsonFile, err := os.Open("companies.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	// Вычитываем содержимое jsonFile в ВИДЕ ПОСЛЕДОВАТЕЛЬНОСТИ БАЙТ!
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	// Теперь задача - перенести все из byteValue в users - это и есть десериализация!
	json.Unmarshal(byteValue, &companies)
	for _, c := range companies.Companies {
		db = append(db, c)

		//PrintCompany(&c)
	}
}
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/companies", GetAllCompanies).Methods("GET")
	router.HandleFunc("/companies/{id}", GetCompanyById).Methods("GET")
	router.HandleFunc("/companies/found/{str}", GetFoundCompany).Methods("GET")
	log.Println("Router gonfigured sucsessfully! Let's Go! ")
	log.Fatal(http.ListenAndServe(":"+port, router))

}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//Struct for representation total slice
// First Level ob JSON object Parsing
type Companies struct {
	Companies []Company `json:"companies"`
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

type Warehouses struct {
	ID         uint64 `json:"id"`
	Name       string `json:"name"`
	Slug       string `json:"slug"`
	company_id uint64 `json:"inn"`
	Address    string `json:"kpp"`
}

type Stoks struct {
	ID                   uint64 `json:"id"`
	Company_aender_id    uint64 `json:"name"`
	Compani_recioient_id uint64 `json:"slug"`
	Product_id           uint64 `json:"inn"`
	Quantity             uint64 `json:"inn"`
	Warehouse_cell_id    uint64 `json:"inn"`
	GTD_id               uint64 `json:"kpp"`
}

type Warehouse_cells struct {
	ID           uint64 `json:"id"`
	Name         string `json:"name"`
	Slug         string `json:"inn"`
	Warehouse_id uint64 `json:"inn"`
}

type GTDs struct {
	ID         uint64 `json:"id"`
	Country_id uint64 `json:"name"`
	Number     string `json:"slug1"`
}

type Countries struct {
	ID      uint64 `json:"id"`
	Code    uint64 `json:"name"`
	Country string `json:"slug"`
}

//Функция для распечатывания User
func PrintUserCompany(c *Company) {
	fmt.Printf("ID: %d\n", c.ID)
	fmt.Printf("Name: %s\n", c.Name)
	fmt.Printf("Slug: %s\n", c.Slug)
	fmt.Printf("INN: %d\n", c.INN)
	fmt.Printf("KPP: %d\n", c.KPP)
}

//1. Рассмотрим процесс десериализации (то есть когда из последовательности в объект)
func main() {
	//1. Создадим файл дескриптор
	jsonFile, err := os.Open("companies.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	fmt.Println("File descriptor successfully created!")

	//2. Теперь десериализуем содержимое jsonFile в экземпляр Go
	// Инициализируем экземпляр Users
	var companies Companies

	// Вычитываем содержимое jsonFile в ВИДЕ ПОСЛЕДОВАТЕЛЬНОСТИ БАЙТ!
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	// Теперь задача - перенести все из byteValue в users - это и есть десериализация!
	json.Unmarshal(byteValue, &companies)
	for _, c := range companies.Companies {
		fmt.Println("================================")
		PrintUserCompany(&c)
	}
}

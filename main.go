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
	var warehouses Warehouses
	// Вычитываем содержимое jsonFile в ВИДЕ ПОСЛЕДОВАТЕЛЬНОСТИ БАЙТ!
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	// Теперь задача - перенести все из byteValue в users - это и есть десериализация!
	json.Unmarshal(byteValue, &companies)
	for _, c := range companies.Companies {
		fmt.Println("================================")
		PrintCompany(&c)
	}
	json.Unmarshal(byteValue, &warehouses)
	for _, w := range warehouses.Warehouses {
		fmt.Println("================================")
		PrintWarehoouse(&w)
	}
}

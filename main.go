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
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	INN  int    `json:"inn"`
	KPP  int    `json:"kpp"`
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

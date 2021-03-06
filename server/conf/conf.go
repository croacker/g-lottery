package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"../commonutils"
)

//Имя файла конфигурации
const fileName = "conf.json"

//Configuration Конфигурация приложения
type Configuration struct {
	Port string
	ClientFolder string
	Sqlite              SqliteConfig
}

//SqliteConfig конфигурация sqlite
type SqliteConfig struct {
	DbPath string
}

var configuration *Configuration

//Get Получить конфигурацию
func Get() *Configuration {
	if configuration == nil {
		config, err := load()
		if err != nil {
			config = makeDefault()
		}
		configuration = config
	}
	return configuration
}

//Загрузить конфигурацию
func load() (*Configuration, error) {
	var err error
	var configuration *Configuration

	dat, err := ioutil.ReadFile(fileName)
	handleError(err)

	err = json.Unmarshal(dat, &configuration)
	handleError(err)

	return configuration, err
}

//Создать конфигурацию поумолчанию
func makeDefault() *Configuration {
	fullFileName := commonutils.CurrentFolder() + "/" + fileName
	fmt.Println("Make default configuration", fullFileName)

	configuration := Configuration{
		Port: "3000",
		ClientFolder: "../client/dist",
		Sqlite: SqliteConfig{
			DbPath: "./db/h-accounting.db",
		},
	}
	b, err := json.Marshal(configuration)
	if err != nil {
		handleError(err)
	}

	ioutil.WriteFile(fullFileName, b, 0777)
	return &configuration
}

//Обработать ошибку
func handleError(err error) {
	if err != nil {
		log.Println(err)
	}
}

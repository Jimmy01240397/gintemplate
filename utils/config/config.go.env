package config

import (
    "os"
    "log"
    "github.com/joho/godotenv"
)

var Port string
var Secret string
var Sessionname string
var DBname string

func init() {
    err := godotenv.Load()
    if err != nil {
        log.Panicln("Error loading .env file")
    }
    Port = os.Getenv("PORT")
    Secret = os.Getenv("SECRET")
    Sessionname = os.Getenv("SESSIONNAME")
    DBname = os.Getenv("DBNAME")
}

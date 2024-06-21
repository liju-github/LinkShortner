package main

import (
	"linkshortner/internal/config"
	database "linkshortner/internal/db"
)

func init()  {
	config.LoadENV()
	database.ConnectDB()
}

func main()  {
	ServerStart()
}


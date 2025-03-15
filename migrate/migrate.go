package main

import (
	"fmt"
	"os"

	"github.com/DaffaKhayru/gofiber-starter-pack/config"
	"github.com/DaffaKhayru/gofiber-starter-pack/models"
)

func main() {
	args := os.Args[1]

	if(args == "up") {
		config.DB.AutoMigrate(&models.User{})
		fmt.Println("Migrate successfully")
	}else if(args == "down") {
		config.DB.Migrator().DropTable(&models.User{})
		fmt.Println("Drop table succesfully")
	}
}
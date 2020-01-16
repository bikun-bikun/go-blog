package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "root:mysql@tcp(127.0.0.1:3306)/blog")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	db.LogMode(true)

	db.AutoMigrate(&User{}, &Profile{}, &Repository{})

	var u = User{
		Name: "たろう",
		Profile: Profile{
			Repositories: []Repository{
				{
					Name:     "gorm",
					Language: "go",
				},
			},
		},
	}

	db.Create(&u)
}

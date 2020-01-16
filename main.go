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

	//ログを標準出力で出力してくれる。
	//デバッグ実装のときは入れておくと何を発行しているかわかりやすいと思う。
	db.LogMode(true)

	//AutoMigration機能、デバッグ実装なのでこれでMigrationしているけど本実装ではきちんとしたMigrationツールを使ったほうが良さそう。
	db.AutoMigrate(&User{}, &Profile{}, &Repository{})

	var u = User{
		Name: "taro",
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

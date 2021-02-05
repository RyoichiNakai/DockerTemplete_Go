package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin" // ginフレームワーク
	_ "github.com/go-sql-driver/mysql" // mysql用ドライバー
	"github.com/jinzhu/gorm" // gorm
	"github.com/joho/godotenv"
)

type User struct {
	gorm.Model
	NickName string `json:"nickName"`
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASSWORD")
	PROTOCOL := "tcp(db:3306)"
	DBNAME := os.Getenv("MYSQL_DATABASE")
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func dbInit() {
	db := gormConnect()
	defer db.Close() // コネクション解放

	db.AutoMigrate(&User{})
}

func dbInsert(nickname string) {
	db := gormConnect()
	defer db.Close()

	db.Create(&User{NickName: nickname})
}

func dbGetAll() []User {
	db := gormConnect()
	defer db.Close()

	var users []User
	// 取得した情報は引数で与えたモデルに格納される
	db.Find(&users)
	return users
}

func main() {
	// 環境変数ファイルの読み込み
	err := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		fmt.Println(err)
	}

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "dadsadasdsa",
		})
	})

	router.GET("/users", func (c *gin.Context)  {
		users := dbGetAll()
		c.JSON(200, gin.H{
			"users": users,
		})
	})

	router.POST("/users/:nickname", func (c *gin.Context)  {
		nickname := c.Param("nickname")
		dbInsert(nickname)
	})

	router.Run(":8080")
}
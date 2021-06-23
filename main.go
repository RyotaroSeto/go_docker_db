package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	// データベース
	Dialect = "mysql"

	// ユーザー名
	DBUser = "homestead"

	// パスワード
	DBPass = "secret"

	// プロトコル
	DBProtocol = "tcp(127.0.0.1:4306)"

	// DB名
	DBName = "homestead"
)

type User struct {
	gorm.Model
	Name string `gorm:"size:255"`
	Age  int
	Sex  string `gorm:"size:255"`
}

type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (u User) String() string {
	return fmt.Sprintf("%s(%d)", u.Name, u.Age)
}

func connectGorm() *gorm.DB {
	connectTemplate := "%s:%s@%s/%s"
	connect := fmt.Sprintf(connectTemplate, DBUser, DBPass, DBProtocol, DBName)
	db, err := gorm.Open(Dialect, connect)

	if err != nil {
		log.Println(err.Error())
	}

	return db
}

func insert(users []User, db *gorm.DB) {
	for _, user := range users {
		db.NewRecord(user)
		db.Create(&user)
	}
}

func findAll(db *gorm.DB) []User {
	var allUsers []User
	db.Find(&allUsers)
	return allUsers
}

func firstUserByID(db *gorm.DB) User {
	var firstUser User
	db.First(&firstUser)
	return firstUser
}

func findByID(db *gorm.DB, id int) User {
	var user User
	db.First(&user, id)
	return user
}

func findByName(db *gorm.DB, name string) []User {
	var user []User
	db.Where("name = ?", name).Find(&user)
	return user
}

func main() {
	// engine := gin.Default()
	// engine.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "hello world",
	// 	})
	// })
	// engine.Run(":3000")
	db := connectGorm()
	defer db.Close()
	// テーブルが存在しない時に対象のテーブルを作成
	db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&User{})

	// user1 := User{Name: "山田太郎", Age: 25, Sex: "男"}
	// user2 := User{Name: "田中花子", Age: 22, Sex: "女"}
	// insertUsers := []User{user1, user2}
	// insert(insertUsers, db)

	// fmt.Println(findAll(db))
	// fmt.Println(firstUserByID(db))
	// fmt.Println(findByID(db, 2)) // 田中花子(22)
	// 存在しない場合は初期値が設定される
	// fmt.Println(findByID(db, 100)) // (0)
	// fmt.Println(findByName(db, "山田太郎"))
}

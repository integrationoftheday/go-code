package config

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db * gorm.DB
)

func Connect(){
	d, err := gorm.Open("mysql", "root:Dhanvilla@13@tcp(127.0.0.1:3306)/simplerest?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        panic(err)
    }

    // set up database configuration
    d.DB().SetMaxIdleConns(10)
    d.DB().SetMaxOpenConns(100)

    // assign db instance to global variable
    db = d
}

func GetDB() *gorm.DB{
	return db
}




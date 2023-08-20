package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/go-tutorial/tutorial/order/models"
	"github.com/stretchr/testify/assert"
)


func TestMysqlFormat(t *testing.T) {

	o := engineOption{
		Host	: "localhost",
		Port	: 1234,
		User	: "gorm",
		Passwd	: "gorm1234",
		Dbname	: "gorm",
	}

	dsn := formatTemplate(mysqlTemplate, &o)
	assert.Equal(t, dsn, "gorm:gorm1234@tcp(localhost:1234)/gorm?parseTime=true")
}



func TestPostgresFormat(t *testing.T) {

	o := engineOption{
		Host	: "localhost",
		Port	: 1234,
		User	: "gorm",
		Passwd	: "gorm1234",
		Dbname	: "gorm",
	}

	dsn := formatTemplate(postgresTemplate, &o)
	assert.Equal(t, dsn, "host=localhost user=gorm password=gorm1234 dbname=gorm port=1234 sslmode=disable TimeZone=Asia/Seoul")
}

func TestFactory(t *testing.T) {

	e, _ := getEngineFactory("sqlite")
	assert.Equal(t, e.getName(), "sqlite")


	m, _ := getEngineFactory("mysql")
	assert.Equal(t, m.getName(), "mysql")


	p, _ := getEngineFactory("postgres")
	assert.Equal(t, p.getName(), "postgres")
}
//

func TestFactoryCreate(t *testing.T) {

	o := engineOption{
		Host	: "127.0.0.1",
		Port	: 3306,
		User	: "test",
		Passwd	: "1234",
		Dbname	: "test",
	}
	db, _ := New("mysql", &o)
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.Order{})

	type Request struct {

		Name string
		Price float32
		Code string
	}

	//r := Request{
	//	Name: "ptch",
	//	Price: 12345,
	//	Code: "C013",
	//}



	//db.Create(&models.User{Name: "sunny"})
	//db.Create(&models.Product{Name: r.Name, Price: r.Price, Code: &r.Code})
	//db.Create(&models.Order{ProductID: 1})
	//
	var user models.User
	//
	db.Find(&user, 1)
	db.Preload("Orders").Find(&user)
	db.Preload("Product").Find(&user.Orders)


	//
	h, err := json.Marshal(user)

	fmt.Println(err)
	fmt.Println(string(h))


}
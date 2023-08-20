package orm

import (
	"fmt"
	orm "github.com/go-tutorial/tutorial/orm/models"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"path"
	"runtime"
)

const (

	FILENAME = "config"
	CFG_FILE_TYPE = "json"
)

type Context struct {
	User *orm.User
}



func optionSetUp(option *engineOption, name string, cfgType string, path string)  error {

	viper.SetConfigName(name)
	viper.SetConfigType(cfgType)
	viper.AddConfigPath(path)

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		return fmt.Errorf("fatal error config file: %w", err)
	}
	viper.Unmarshal(option)
	return nil
}


func initDB() (*gorm.DB, error) {

	_, filename, _, _ := runtime.Caller(0)
	configDirPath := path.Dir(filename)
	err := optionSetUp(&option, FILENAME, CFG_FILE_TYPE, configDirPath)

	if err != nil {
		return nil, err
	}

	// db type
	// mysql or postgres or sqlite
	dbType := viper.GetString("type")

	gormDB, err := New(dbType, &option)
	if err != nil {
		return nil, err
	}

	// init database tables
	if err := gormDB.AutoMigrate(&orm.User{}, &orm.Product{}, &orm.Order{}); err != nil {
		return nil, err
	}

	user := orm.User{
			Name: "dragon",
		}


	products := []*orm.Product{
		{
			Name: "Iphone",
			Price:  128,
			Code: "CD01",
		},
		{
			Name: "MacBook",
			Price:  128,
			Code: "CD02",
		},
		{
			Name: "Imac",
			Price:  128,
			Code: "CD03",
		},
	}

	orders := []*orm.Order{
		{
			UserId: 1,
			ProductID: 1,
		},
		{
			UserId: 1,
			ProductID: 2,
		},
	}

	// init users
	gormDB.Create(&user)

	// init Product
	gormDB.Create(products)

	// init order
	gormDB.Create(orders)


	return gormDB, nil
}

func dropTables(d *gorm.DB) error {

	var migrator gorm.Migrator
	migrator = d.Migrator()
	if err := migrator.DropTable(&orm.User{}, &orm.Product{}, &orm.Order{}); err !=nil {
		return err
	}
	return nil
}


func getUser(db *gorm.DB, context *Context)  error {
	result := db.First(context.User)
	return result.Error
}

func T()  error {

	db, err := initDB()
	if err != nil {
		return err
	}
	defer dropTables(db)

	ctx := Context{
		User: &orm.User{},
	}
	if err := getUser(db, &ctx); err != nil {
		return err
	}
	db.Preload("Orders").Find(ctx.User)
	db.Preload("Product").Find(&ctx.User.Orders)
	return nil
}


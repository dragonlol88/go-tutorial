package orm

import (
	"fmt"
	"strings"
	"text/template"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)



var (

	enginesMap		 = make(map[string]engineFactory)

	engines 		 = []engineFactory{sqlLightEngineFactory{}, mysqlEngineFactory{}, postgresEngineFactory{}}

	mysqlTemplate    = "{{.User}}:{{.Passwd}}@tcp({{.Host}}:{{.Port}})/{{.Dbname}}?parseTime=true"
	postgresTemplate = "host={{.Host}} user={{.User}} password={{.Passwd}} dbname={{.Dbname}} port={{.Port}} sslmode=disable TimeZone=Asia/Seoul"

	formatTemplate = func (s string, option *engineOption) string {
		t, b := new(template.Template), new(strings.Builder)
		template.Must(t.Parse(s)).Execute(b, option)
		return b.String()
	}

	option engineOption

)



type engineOption struct {
	Dbname 	string
	Host 	string
	Port 	int
	User 	string
	Passwd 	string
}


type engineFactory interface {
	create(option *engineOption) (*gorm.DB, error)
	getName() string
}

type (
	sqlLightEngineFactory struct {
		engineOption
	}

	mysqlEngineFactory struct {
		engineOption
	}

	postgresEngineFactory struct {
		engineOption
	}
)


func (e sqlLightEngineFactory) getName() string {
	return "sqlite"
}

func (e sqlLightEngineFactory) create(option *engineOption) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(option.Dbname + ".orm"), &gorm.Config{SkipDefaultTransaction: true,})
}



func (e mysqlEngineFactory) getName() string {
	return "mysql"
}

func (e mysqlEngineFactory) create(option *engineOption) (*gorm.DB, error) {

	dsn := formatTemplate(mysqlTemplate, option)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{SkipDefaultTransaction: true,})
}


func (e postgresEngineFactory) getName() string {
	return "postgres"
}

func (e postgresEngineFactory) create(option *engineOption) (*gorm.DB, error) {

	dsn := formatTemplate(postgresTemplate, option)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true,})
}


func getEngineFactory(dial string) (engineFactory, error) {

	f, ok := enginesMap[dial]
	if ok {
		return f, nil
	}

	for _, engine := range engines {
		if dial == engine.getName() {
			enginesMap[dial] = engine
			return engine, nil
		}
	}
	return nil, fmt.Errorf("`%s` engine is not exist.", dial)
}


func New(dial string, option *engineOption)  (*gorm.DB, error) {

	f, err := getEngineFactory(dial)
	if err != nil {
		return nil, err
	}
	return f.create(option)
}
package main

import (
	"fmt"
	"github.com/go-tutorial/tutorial/atlas/models"
	"io"
	"os"

	_ "ariga.io/atlas-go-sdk/recordriver"
	"ariga.io/atlas-provider-gorm/gormschema"
)

func main() {

	// keep model load's order because of reference(foreign key)
	stmts, err := gormschema.New("mysql").Load(
		&models.User{},
		&models.Product{},
		&models.Order{},
		)

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}

	io.WriteString(os.Stdout, stmts)
}

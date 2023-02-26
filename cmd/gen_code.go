package main

import (
	"MyProject/config"
	"MyProject/entity"

	"gorm.io/gen"
)

func genCode() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./dao",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	db := config.LoadDB()
	g.UseDB(db) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `entity.User` following conventions
	g.ApplyBasic(entity.User{})

	// Generate the code
	g.Execute()

}

func main() {
	genCode()
}

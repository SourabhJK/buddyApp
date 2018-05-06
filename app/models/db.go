package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/revel/revel"
	"github.com/go-gorp/gorp"

)



var (
	// Dbm is db handle
	Dbm *gorp.DbMap
)

func init() {
	initDB()
	updateModels()
}

func initDB() {
	var dbInfo string
	
		revel.INFO.Println("RUNNING IN PROD MODE. DATABASE WILL BE RDS")
		dbInfo = "host=buddyapp.cantk5mmoyg2.us-east-2.rds.amazonaws.com user=buddeInsta password=Ch!Buddesaheb2 " +
			"dbname=BuddyApp sslmode=disable"
	

	Db, err := sql.Open("mysql", dbInfo)
	if Db == nil || err != nil {
		revel.ERROR.Println("could not connect to mysql", dbInfo)
		panic(err)
	}
	Dbm = &gorp.DbMap{Db: Db, Dialect: gorp.MySQLDialect{}}
}

func updateModels() {
	/*
	setColumnSizes := func(t *gorp.TableMap, colSizes map[string]int) {
		for col, size := range colSizes {
			t.ColMap(col).MaxSize = size
		}
	}
	t := Dbm.AddTableWithName(Users{}, "users").SetKeys(true, "Id")
	t.ColMap("Password").Transient = true
	setColumnSizes(t, map[string]int{
		"Username": 50,
		"Name":     100,
	})
	*/
	
	t = Dbm.AddTableWithName(User{}, "report").SetKeys(true, "Id")

	

	Dbm.TraceOn("[gorp]", revel.INFO)
	err := Dbm.CreateTablesIfNotExists()
	revel.INFO.Println("Error in creating table:", err)

	//add some test content
	//bcryptPassword, _ := bcrypt.GenerateFromPassword([]byte("havells"), bcrypt.DefaultCost)
	//user := &User{0, "operator havells", "havells", "havells", bcryptPassword}
}

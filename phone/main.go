package main

import (
	//"database/sql"
	"fmt"
	"regexp"
	// we will not use any functions from this pkg but inside pq package this inits and registers driver
	//https://www.calhoun.io/why-we-import-sql-drivers-with-the-blank-identifier/
	_ "github.com/lib/pq"
	phonedb "gophercises/phone/db"
)

const (
	host="localhost"
	port = 5432
	user = "postgres"
	password = "vikas@iitb"
	dbname = "gophercises_phone"
)

func main()  {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable",host,port,user,password)

	must(phonedb.Reset("postgres",psqlInfo,dbname))

	psqlInfo = fmt.Sprintf("%s dbname=%s",psqlInfo,dbname)

	must(phonedb.Migrate("postgres",psqlInfo))

	db,err := phonedb.Open("postgres",psqlInfo)
	must(err)

	defer db.Close()

	err = db.Seed()
	must(err)

	phones, err := db.AllPhones()
	must(err)

	for _,p := range phones{
		fmt.Printf("Working on ... %v\n ",p)
		number := normalize(p.Number)
		if number != p.Number{
			fmt.Println("Updating or removing ...",number)
			existing,err := db.FindPhone(number)
			must(err)
			if existing != nil{
				// delete this number
				must(db.DeletePhone(p.ID))
			}else{
				//update the number
				p.Number=number
				must(db.UpdatePhone(&p))
			}
		}else {
			fmt.Println("No changes required")
		}
	}
}

func normalize(phone string) string  {
	// we want these - 0123456789
	re := regexp.MustCompile("[^0-9]")
	return re.ReplaceAllString(phone,"")
}

func must(err error){
	if err != nil{
		panic(err)
	}
}



package db

import (
	"database/sql"
	//"errors"
	"fmt"
	_ "github.com/lib/pq"
)

//phone represents the phone_numbers table in the DB
type Phone struct {
	ID int
	Number string
}


func Open(driverName, dataSource string) (*DB,error)  {
	db,err := sql.Open(driverName,dataSource)
	if err != nil{
		return nil, err
	}
	return  &DB{db}, nil
}

type DB struct {
	db *sql.DB
}

func (db *DB) Close()  error {
	return db.db.Close()
}

func (db *DB) Seed() error{
	data := []string{
		"1234567890",
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"123-456-7890",
		"1234567892",
		"(123)456-7892",
	}
	for _,number := range data{
		if _,err:= insertPhone(db.db,number); err != nil{
			return err
		}
	}
	return nil
}

func (db *DB) AllPhones() ([]Phone,error) {
	return allPhones(db.db)
}
func allPhones(db *sql.DB)([]Phone,error)  {
	rows,err := db.Query("select id,value from phone_numbers")
	if err != nil{
		return nil, err
	}
	defer rows.Close()
	var ret []Phone
	for rows.Next(){
		var p Phone
		if err := rows.Scan(&p.ID,&p.Number); err != nil{
			return nil, err
		}
		ret = append(ret, p)
	}
	if err := rows.Err();err != nil{
		return nil, err
	}
	return ret, nil
}

func Migrate(driverName,dataSource string) error  {
	db,err := sql.Open(driverName,dataSource)
	if err != nil{
		return err
	}
	if err := createPhoneNumberTable(db); err != nil{
		return err
	}
	return db.Close()
}
func createPhoneNumberTable(db *sql.DB) error  {
	statement := fmt.Sprintf(`
		create table if not exists phone_numbers (
		id serial,
		value VARCHAR(255)
)`)
	_, err := db.Exec(statement)
	return err
}

func Reset(driverName,dataSource,dbName string)error{
		db,err := sql.Open(driverName,dataSource)
		if err != nil{
			return err
		}
		err = resetdb(db,dbName)
		if err != nil{
			return err
		}
		return db.Close()
}

func resetdb(db *sql.DB, name string) error {
	_,err := db.Exec("drop database if exists "+name)
	if err != nil{
		return err
	}
	return createDB(db,name)
}
func createDB(db *sql.DB, name string)error  {
	_,err := db.Exec("create database "+ name)
	if err != nil{
		return err
	}
	return nil
}

func insertPhone(db *sql.DB, phone string) (int , error)  {
	statement := `insert into phone_numbers(value) values($1) returning id`
	var id int
	err := db.QueryRow(statement,phone).Scan(&id)
	if err != nil{
		return -1, err
	}
	return int(id), nil
}


// we dont use this right now
func getPhone(db *sql.DB, id int) (string,error){
	var number string
	err := db.QueryRow("select value from phone_numbers where id=$1",id).Scan(&number)
	if err != nil{
		return "", err
	}
	return number, nil
}


func (db *DB)DeletePhone(id int) error {
	statement := `delete from phone_numbers where id=$1`
	_,err := db.db.Exec(statement,id)
	return  err
}

func (db *DB)UpdatePhone(p *Phone) error  {
	statement := `Update phone_numbers set value=$2 where id=$1`
	_,err := db.db.Exec(statement,p.ID,p.Number)
	return err
}

func (db *DB)FindPhone(number string) (*Phone,error){
	var p Phone
	err := db.db.QueryRow("select * from phone_numbers where value=$1",number).Scan(&p.ID,&p.Number)
	if err != nil{
		if err == sql.ErrNoRows{
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

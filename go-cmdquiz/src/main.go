package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {

	var csvfilename string
	if len(os.Args) > 1 {
		csvfilename = os.Args[1]
	} else {
		fmt.Println("No filename provided as argument using default filename of prblems.csv")
		fmt.Println(os.Args)
		csvfilename = "problems.csv"
	}

	//open the file
	csvfile, err := os.Open(csvfilename)
	if err != nil {
		log.Fatalln("couldn't open csv file", err)
	}

	//Parse the file
	r := csv.NewReader(csvfile)

	var correctAnswerCount int
	var totalCount int

	//Iterate over record
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		totalCount++
		// fmt.Println("record[0] = ",record[0],"record[1] = ",record[1]);
		nbytes, err := fmt.Println("What is ", record[0], "?")
		if err != nil {
			log.Print("read", nbytes)
			log.Fatalln(err)
		}
		var userinput int
		fmt.Scanln(&userinput)
		answer, err := strconv.Atoi(record[1])
		if userinput == answer {
			correctAnswerCount++
		}
	}
	fmt.Println("total number of correct answers = ", correctAnswerCount, "/", totalCount)
}

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main()  {
	//fileName := "birthday_001.txt"
	//// => Birthday - 1 of 4.txt
	//newName,err := match(fileName,4)
	//if err != nil{
	//	fmt.Println("no match")
	//	os.Exit(1)
	//}
	//fmt.Println("new name : ",newName)
	dir := "sample"
	files,err := ioutil.ReadDir(dir)
	if err != nil{
		panic(err)
	}
	count := 0
	var toRename []string
	for _,file := range files{
		if file.IsDir(){
			//fmt.Println("Dir:",file.Name())
		}else{
			_,err := match(file.Name(),4)
			if err == nil{
				count++
				toRename = append(toRename,file.Name())
			}
			//fmt.Println("match : ",tmp,err)
		}
	}
	for _,origFileName := range toRename{
		origPath :=  filepath.Join(dir,origFileName)
		newFileName,err := match(origFileName,count)
		if err != nil{
			panic(err)
		}
		newPath := filepath.Join(dir,newFileName)
		fmt.Printf("mv %s => %s\n",origPath,newPath)
		err = os.Rename(origPath,newPath)
		if err != nil{
			panic(err)
		}
	}
}

//match returns a new modified filename
func match(fileName string,total int) (string,error)  {
	//"birthday" "001" "txt"
	pieces := strings.Split(fileName,".")
	ext := pieces[len(pieces)-1]
	tmp := strings.Join(pieces[0:len(pieces)-1],".")
	pieces = strings.Split(tmp,"_")
	name := strings.Join(pieces[0:len(pieces)-1],"_")
	number,err := strconv.Atoi(pieces[len(pieces)-1])
	if err != nil{
		return "",fmt.Errorf("%s didn't match our pattern",fileName)
	}
	return fmt.Sprintf("%s - %d of %d.%s",strings.Title(name),number,total,ext), nil
}
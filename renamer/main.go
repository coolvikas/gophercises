package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)
type file struct {
	name string
	path string
}
func main()  {
	//fileName := "birthday_001.txt"
	//// => Birthday - 1 of 4.txt
	//newName,err := match(fileName,4)
	//if err != nil{
	//	fmt.Println("no match")
	//	os.Exit(1)
	//}
	//fmt.Println("new name : ",newName)
	dir := "./sample"
	var toRename []file
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		//fmt.Println(path,info.IsDir())
		if info.IsDir(){
			return nil
		}
		//if info.IsDir() && path == "sample/nested"{
		//	return filepath.SkipDir
		//}
		if _,err := match(info.Name()); err == nil{
			toRename = append(toRename,file{
				name: info.Name(),
				path: path,
			})
		}
		return nil
	})
	for _,fileName := range toRename{
		fmt.Println(fileName)
	}

	for _,orig := range toRename{
		var n file
		var err error
		n.name, err = match(orig.name)
		if err != nil{
			fmt.Println("Error matching: ", orig.path,err.Error())
		}
		n.path = filepath.Join(dir,n.name)
		fmt.Printf("mv %s => %s \n",orig.path,n.path)

		err = os.Rename(orig.path,n.path)
		if err != nil{
			fmt.Println("Error renaming: ", orig.path,err.Error())
		}
	}
}

//match returns a new modified filename
func match(fileName string) (string,error)  {
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
	return fmt.Sprintf("%s - %d.%s",strings.Title(name),number,ext), nil
}
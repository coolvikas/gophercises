package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"link"
	"strings"
)

//var examplehtml=`
//<html>
//<body>
//  <h1>Hello!</h1>
//  <a href="/other-page">A link to
//another page</a>
//</body>
//</html>
//`
func main(){
	htmlfilename := flag.String("htmlfilename","ex2.html","pass html file name to parse hrefs")
	flag.Parse()
	//fmt.Println("reading from file ", *htmlfilename)


	htmlcontent,err := ioutil.ReadFile(*htmlfilename)
	if err != nil{
		panic(err)
	}

	r:= strings.NewReader(string(htmlcontent))

	links,err:= link.Parse(r)
	if err != nil{
		panic(err)
	}
	fmt.Printf("%+v",links)
}


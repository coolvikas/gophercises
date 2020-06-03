package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"link"
)
const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type loc struct {
	Value string `xml:"loc"`
}

type urlset struct {
	Urls []loc `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}
func main()  {
	urlFlag := flag.String("url","https://gophercises.com","url to build sitemap for" )
	maxDepth := flag.Int("depth",10,"max depth to traverse")
	flag.Parse()
	fmt.Println("url is ",*urlFlag)

	/*
	1. GET the webpage
	2. parse all the links on the page
	3. build proper urls with our links
	4. filter out any links w/ a diff domain
	5. Find all the pages (BFS)
	6. print out XML
	 */

	pages:= bfs(*urlFlag,*maxDepth)
	//for _,page := range pages{
	//	fmt.Println(page)
	//}
	var toXml urlset
	toXml = urlset{
		Xmlns: xmlns,
	}
	for _,page:= range pages{
		toXml.Urls = append(toXml.Urls,loc{page})
	}
	fmt.Print(xml.Header)
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("","  ")
	if err := enc.Encode(toXml); err !=nil{
		log.Fatal(err)
	}
	fmt.Println()


	//fmt.Println(links)
	/*
	/some-path => add domain to it ang goGET
	https://gophercises.com/some-path  =>  goGET
	#fragment => skip it
	mailto:jon@calhoun.io => skip it
	 */
}

func bfs(urlStr string,maxDepth int) []string{
	seen:= make(map[string]struct{})
	var q map[string]struct{}
	nq := map[string]struct{}{
		//key:value
		urlStr: struct{}{},
	}
	for i:=0;i<=maxDepth;i++{
		q,nq = nq, make(map[string]struct{})
		if len(q) == 0{
			break
		}
		for url,_ := range q{
			// if page is already seen, continue
			if _,ok := seen[url]; ok{
				continue
			}
			seen[url] = struct{}{}
			// get the next urls and put in nq queue
			for _,l := range get(url){
				nq[l] = struct{}{}
			}
		}
	}
	// return all the seen urls
	//var ret []string
	ret := make([]string,0,len(seen))
	for url,_ := range seen{
		ret = append(ret,url)
	}
	return ret

}

func get(urlStr string) []string{
	resp,err := http.Get(urlStr)
	if err != nil {
		return []string{}
	}
	defer resp.Body.Close()

	reqURL := resp.Request.URL
	baseURL := &url.URL{
		Scheme:     reqURL.Scheme,
		Host:       reqURL.Host,
	}
	base := baseURL.String()
	// return the filtered hrefs values with prefix base
	return filter(hrefs(resp.Body,base),withPrefix(base))

}
func hrefs(r io.Reader,base string) []string{

	links,err := link.Parse(r)
	if err != nil{
		log.Fatal(err)
	}

	var ret []string
	for _,l := range links{
		switch  {
		case strings.HasPrefix(l.Href,"/"):
			ret = append(ret,base+l.Href)
		case strings.HasPrefix(l.Href,"http"):
			ret = append(ret,l.Href)
		default:
			//fmt.Println("skipping .....",l.Href)
		}
	}
	return ret
}

func filter(links []string, keepFn func(string) bool) []string{
	var ret []string
	for _,l := range links{
			if keepFn(l){
				ret =append(ret,l)
		}
	}
	return  ret
}

func withPrefix(pfx string) func(string) bool {
	return func(l string) bool{
		return strings.HasPrefix(l,pfx)
	}
}
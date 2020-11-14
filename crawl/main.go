package main

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func GetHtml(url string) io.Reader {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("get err", err)
	}
	return resp.Body
}

//type Content struct {
//	HeadLine string
//	Time string
//	Tag string
//	body string
//}
//
//func store(){
//	var content map[string] Content
//	content = make(map[string] Content)
//	r := GetHtml("https://blog.lenconda.top/")
//	dom,err := goquery.NewDocumentFromReader(r)
//	if err!=nil {
//		log.Fatalln(err)
//	}
//	Dom := dom.Find("article>h2").Text()
//	content["1"]. = Dom
//
//}


func main(){
	file, err := os.Create("D:/crawl.txt")
	if err != nil {
		panic(err)

	}
	defer file.Close()
	for i:= 1;i<=7;i++{

		r := GetHtml("https://blog.lenconda.top/page/"+strconv.Itoa(i)+"/")
		dom,err := goquery.NewDocumentFromReader(r)
		if err!=nil{
			log.Fatalln(err)
		}
		Dom := dom.Find("article")

		str := Dom.Text()
		file.WriteString(str)
				println(str)


		



	}

}
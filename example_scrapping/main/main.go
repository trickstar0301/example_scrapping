package main

import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"net/http"
	"os"
	"io"
	"strconv"
	"github.com/signintech/gopdf"
	"log"
)

func main() {
	fmt.Print("*** scraping start ***\n")
	uri := "https://xxx"
	doc, err := goquery.NewDocument(uri)
	if err != nil {
		fmt.Print("url scarapping failed\n")
	}

	urlSlice := make([]string, 0)

	doc.Find(".slide_container").Each(func(_ int, s *goquery.Selection) {
		s.Find(".slide_image").Each(func(_ int, s *goquery.Selection) {
			tmp,_ := s.Attr("data-full")
			urlSlice = append(urlSlice, tmp)
		})
	})

	for i,_:=range(urlSlice){
		response, err := http.Get(urlSlice[i])
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()

		file, err := os.Create("../images/" + strconv.Itoa(i) + ".jpg")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		io.Copy(file, response.Body)
	}
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 576, H: 326}}) //595.28, 841.89 = A4

	for i,_:=range(urlSlice){
		pdf.AddPage()
		filename := ("../images/" + strconv.Itoa(i) + ".jpg")
		pdf.Image(filename, 0, 0, nil)
		log.Print(filename+"\r\n")
	}
	pdf.WritePdf("../output/image.pdf")


}

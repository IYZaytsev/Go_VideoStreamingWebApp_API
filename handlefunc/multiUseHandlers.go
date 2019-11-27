package handlefunc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//Page used for sending Json for index
type Page struct {
	Name       string
	VideoNames []string
}

//ReturnIndex returns the names of videos uploaded on the site
func ReturnIndex(w http.ResponseWriter, r *http.Request) {
	page := Page{}
	files, err := ioutil.ReadDir("./videos")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if strings.Contains(f.Name(), ".mp4") || strings.Contains(f.Name(), ".webm") {
			page.VideoNames = append(page.VideoNames, f.Name())
		}
	}

	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	page.Name = "index"
	if err := json.NewEncoder(w).Encode(page); err != nil {
		panic(err)
	}

}

//Watch serves the video
func Watch(w http.ResponseWriter, r *http.Request) {

	vidIndex, _ := strconv.Atoi(r.URL.Path[len("/watch/"):])
	page := Page{}
	files, err := ioutil.ReadDir("./videos")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if strings.Contains(f.Name(), ".mp4") || strings.Contains(f.Name(), ".webm") {
			page.VideoNames = append(page.VideoNames, f.Name())
		}

	}

	http.ServeFile(w, r, "./videos/"+page.VideoNames[vidIndex])

}

//ReceiveFile used to download files from sender
func ReceiveFile(w http.ResponseWriter, r *http.Request) {

	// in your case file would be fileupload
	file, header, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}

	fmt.Print("thread started")
	start := time.Now()
	var Buf bytes.Buffer
	//defer means this will execute when the function returns
	defer file.Close()

	//fmt.Printf("File name %s\n", header.Filename)

	io.Copy(&Buf, file)

	contents := Buf.Bytes()
	//fmt.Println(contents)
	ioutil.WriteFile("./videos/"+header.Filename, contents, 0644)

	Buf.Reset()

	fmt.Println("file Uploaded !!!!")
	fmt.Println(time.Since(start))

	return

}

func downloadFiles(file multipart.File, header *multipart.FileHeader) {
	fmt.Print("thread started")
	start := time.Now()
	var Buf bytes.Buffer
	//defer means this will execute when the function returns
	defer file.Close()

	//fmt.Printf("File name %s\n", header.Filename)

	io.Copy(&Buf, file)

	contents := Buf.Bytes()
	//fmt.Println(contents)
	ioutil.WriteFile("./videos/"+header.Filename, contents, 0644)

	Buf.Reset()

	fmt.Println("file Uploaded !!!!")
	fmt.Println(time.Since(start))

	return
}

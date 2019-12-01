package handlefunc

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

//Page used for sending Json for index
type Page struct {
	Name       string
	VideoNames []string
	Index      int
}

//ReturnIndex returns the names of videos uploaded on the site
func ReturnIndex(w http.ResponseWriter, r *http.Request) {
	page := Page{}
	files, err := ioutil.ReadDir("./videos")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if strings.Contains(f.Name(), ".mp4") {
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
		if strings.Contains(f.Name(), ".mp4") {
			page.VideoNames = append(page.VideoNames, f.Name())
		}

	}

	http.ServeFile(w, r, "./videos/"+page.VideoNames[vidIndex])

}

//ServeThumbnails serves the Jpg thumbnails for videos
func ServeThumbnails(w http.ResponseWriter, r *http.Request) {

	vidIndex, _ := strconv.Atoi(r.URL.Path[len("/image/"):])
	page := Page{}
	files, err := ioutil.ReadDir("./videos/thumbnails")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if strings.Contains(f.Name(), ".jpg") {
			page.VideoNames = append(page.VideoNames, f.Name())
		}

	}

	http.ServeFile(w, r, "./videos/thumbnails/"+page.VideoNames[vidIndex])

}

//ReceiveFile used to download files from sender
func ReceiveFile(w http.ResponseWriter, r *http.Request) {

	// in your case file would be fileupload
	file, header, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}

	var Buf bytes.Buffer
	//defer means this will execute when the function returns
	defer file.Close()

	//fmt.Printf("File name %s\n", header.Filename)

	io.Copy(&Buf, file)

	contents := Buf.Bytes()
	//fmt.Println(contents)
	ioutil.WriteFile("./videos/"+header.Filename, contents, 0644)

	Buf.Reset()
	//this bash command creates thumbnails for the uploaded video
	cmd := exec.Command("ffmpeg", "-i", "videos/"+header.Filename, "-ss", "00:00:03", "-vframes", "1", "-s", "480x320", "videos/thumbnails/"+strings.Replace(header.Filename, ".mp4", ".jpg", -1))
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	errr := cmd.Run()
	if errr != nil {
		return
	}
	return

}

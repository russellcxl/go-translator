package main

import (
	"flag"
	"fmt"
	"github.com/russellcxl/go-translator/config"
	"github.com/russellcxl/go-translator/pkg/logger"
	"github.com/russellcxl/go-translator/server"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	MAX_UPLOAD_SIZE = 1024 * 1024
)

var (
	Port string
)

func handleUpload(w http.ResponseWriter, r *http.Request) {
	log.Println("File Upload Endpoint Hit")

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 32 MB is the default used by FormFile()
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get a reference to the fileHeaders.
	// They are accessible only after ParseMultipartForm is called
	files := r.MultipartForm.File["myFiles"]

	for _, fileHeader := range files {
		// Restrict the size of each uploaded file to 1MB.
		// To prevent the aggregate size from exceeding
		// a specified value, use the http.MaxBytesReader() method
		// before calling ParseMultipartForm()
		if fileHeader.Size > MAX_UPLOAD_SIZE {
			http.Error(w, fmt.Sprintf("The uploaded image is too big: %s. Please use an image less than 1MB in size", fileHeader.Filename), http.StatusBadRequest)
			return
		}

		// Open the file
		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer file.Close()

		buff := make([]byte, 512)
		_, err = file.Read(buff)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		filetype := http.DetectContentType(buff)
		if filetype != "image/jpeg" && filetype != "image/png" {
			http.Error(w, "The provided file format is not allowed. Please upload a JPEG or PNG image", http.StatusBadRequest)
			return
		}

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = os.MkdirAll("images/input", os.ModePerm)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		f, err := os.Create(fmt.Sprintf("images/input/%s", fileHeader.Filename))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		defer f.Close()

		_, err = io.Copy(f, file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	htmlTemp := `
	<div>hello</div>
	<button>button</button>
`

	fmt.Fprintf(w, htmlTemp)
}

func init() {
	flag.StringVar(&Port, "port", "", "port for http")
	flag.Parse()
}

func main() {

	fileServer := http.FileServer(http.Dir("templates"))
	http.Handle("/", fileServer)
	http.HandleFunc("/upload", handleUpload)


	fmt.Printf("Starting server at port %s\n", Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", Port), nil); err != nil {
		log.Fatal(err)
	}

	clog := logger.NewLogger("./logs")
	clog.InitLogger()

	cfg, err := config.LoadConfig("./config/config.json")
	if err != nil {
		clog.Fatalf("failed to read config: %s\n", err.Error())
		return
	}

	err = server.NewTranslator(clog, cfg).Execute()
	if err != nil {
		clog.Fatalf(err.Error())
	}
}
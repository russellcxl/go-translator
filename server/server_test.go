package server

import (
	"bufio"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func Test_ReadTxtInBatches(t *testing.T) {
	// open input file
	fi, err := os.Open("test_input/test_1.txt")
	if err != nil {
		panic(err)
	}
	// close fi on exit and check for its returned error
	defer fi.Close()

	//var a string
	temp := make([]byte, 1024)
	r := bufio.NewReader(fi)
	for {
		n, err := r.Read(temp)
		if err != nil {
			fmt.Printf("failed to read further: %s\n", err.Error())
			break
		}
		fmt.Printf("\nbytes read: %d\n", n)
		fmt.Printf("---> %s\n", html.UnescapeString(strings.ReplaceAll(string(temp[:n]), "\"", "")))
		if n == 0 {
			return
		}
	}
}


func Test_GetAllFiles(t *testing.T) {

	files, err := ioutil.ReadDir("images")
	if err != nil {
		log.Fatalf("failed to read images: %s\n", err.Error())
	}
	for _, v := range files {
		fmt.Printf("read from %s; write to %s.txt\n", v.Name(), strings.TrimSuffix(v.Name(), filepath.Ext(v.Name())))
	}
}

func Test_WriteToFile(t *testing.T) {

	// open output file
	f, err := os.Create("test_output/output.txt")
	if err != nil {
		log.Fatalf("failed to create output file: %s\n", err.Error())
	}
	defer f.Close()

	text := `hello there kind sir, would you like a donut?`

	_, err = f.Write([]byte(text))
	if err != nil {
		log.Fatalf("failed to write to output file: %s\n", err.Error())
	}
}

func Test_ReadFromInputFilesAndWriteToOutput(t *testing.T) {
	files, err := ioutil.ReadDir("test_input")
	if err != nil {
		log.Fatalf("failed to read images: %s\n", err.Error())
	}
	for _, v := range files {
		fmt.Printf("will read from %s; will write to %s.txt\n", v.Name(), strings.TrimSuffix(v.Name(), filepath.Ext(v.Name())))

		// open input file
		fi, err := os.Open(fmt.Sprintf("test_input/%s", v.Name()))
		if err != nil {
			panic(err)
		}

		// prepare output file
		fo, err := os.Create(fmt.Sprintf("test_output/%s.txt", strings.TrimSuffix(v.Name(), filepath.Ext(v.Name()))))
		if err != nil {
			log.Fatalf("failed to create output file: %s\n", err.Error())
		}

		//var a string
		temp := make([]byte, 1024)
		r := bufio.NewReader(fi)
		for {
			n, err := r.Read(temp)
			if n == 0 {
				break
			}
			if err != nil {
				fmt.Printf("failed to read further: %s\n", err.Error())
				break
			}
			formattedText := fmt.Sprintf("%s", html.UnescapeString(strings.ReplaceAll(string(temp[:n]), "\"", "")))
			fmt.Printf("READING --> %s\n", formattedText)
			_, err = fo.Write([]byte(formattedText))
			if err != nil {
				log.Fatalf("failed to write to output file: %s\n", err.Error())
			}
		}

		fo.Close()
		fi.Close()

	}
}

func Test_ReadByLine(t *testing.T) {
	fi, err := os.Open("test_input/test_1.txt")
	if err != nil {
		panic(err)
	}
	// close fi on exit and check for its returned error
	defer fi.Close()

	s := bufio.NewScanner(fi)
	var finalText string
	for s.Scan() {
		finalText += s.Text()
	}
	fmt.Println(strings.ReplaceAll(finalText, "\"", ""))

}

func Test_ReadWithDelimiter(t *testing.T) {
	// open input file
	fi, err := os.Open(fmt.Sprintf("test_input/test_1.txt"))
	if err != nil {
		panic(err)
	}

	// prepare output file
	fo, err := os.Create("test_output/ReadBySentence.txt")
	if err != nil {
		log.Fatalf("failed to create output file: %s\n", err.Error())
	}

	defer fi.Close()
	defer fo.Close()

	r :=  bufio.NewReader(fi)
	for {
		b, err := r.ReadBytes('n')
		fmt.Printf("reading %d bytes ---> %s\n", len(b), strings.TrimSpace(strings.ReplaceAll(string(b), "\n", "")))
		if err != nil {
			fmt.Printf("failed to read further: %s\n", err.Error())
			break
		}
	}

}

func Test_MakeReadable(t *testing.T) {
	// open input file
	fi, err := os.Open(fmt.Sprintf("../images/output/IMG_20220607_190242_edit_193636695911077_resized_20220607_070514281.txt"))
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(fi)
	for s.Scan() {
		t2 := strings.Replace(s.Text(), `\n`, "\n", -1)
		lines := strings.Split(t2, "\n")
		for _, v := range lines {
			fmt.Println(v)
		}
	}

	fi2, err := os.Open(fmt.Sprintf("../images/output/IMG_20220607_190249_edit_193610537127227_resized_20220607_070514640.txt"))
	if err != nil {
		panic(err)
	}

	s2 := bufio.NewScanner(fi2)
	for s2.Scan() {
		t2 := strings.Replace(s2.Text(), `\n`, "\n", -1)
		lines := strings.Split(t2, "\n")
		for _, v := range lines {
			fmt.Println(v)
		}
	}

}
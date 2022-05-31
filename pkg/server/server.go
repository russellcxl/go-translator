package server

import (
	"bufio"
	"bytes"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Translator struct {
	InputDir   string
	OutputDir  string
	OutputLang string
}

func NewTranslator(in, out, lang string) *Translator {
	return &Translator{
		InputDir:   in,
		OutputDir:  out,
		OutputLang: lang,
	}
}

//TODO: add checks for non-images
func (t *Translator) Execute() error {
	files, err := ioutil.ReadDir(t.InputDir)
	if err != nil {
		log.Fatalf("failed to read from input directory: %s\n", err.Error())
		return err
	}

	for _, f := range files {
		log.Printf("will read from %s; will write to %s.txt\n", f.Name(), strings.TrimSuffix(f.Name(), filepath.Ext(f.Name())))

		in := path.Join("images", f.Name())
		outFile := fmt.Sprintf("%s.txt", strings.TrimSuffix(f.Name(), filepath.Ext(f.Name())))
		out := path.Join(t.OutputDir, outFile)

		if err = readAndTranslate(in, out, t.OutputLang); err != nil {
			return err
		}
	}
	return nil
}

func readAndTranslate(inPath, outPath, outputLang string) error {
	buf := new(bytes.Buffer)

	err := detectText(buf, inPath)
	if err != nil {
		log.Printf("failed to detect text: %s\n", err.Error())
		return err
	}

	// prepare output file
	fo, err := os.Create(outPath)
	if err != nil {
		log.Printf("failed to create output file: %s\n", err.Error())
		return err
	}
	defer fo.Close()

	// translate api cannot take in too many bytes at once
	temp := make([]byte, 2048)
	r := bufio.NewReader(buf)

	for {
		n, err := r.Read(temp)
		if n == 0 {
			break
		}
		if err != nil {
			log.Printf("failed to read: %s\n", err.Error())
			return err
		}

		// clean the string
		str := string(temp[:n])
		str = strings.ReplaceAll(str, `"`, "")
		str = strings.TrimSpace(str)

		// translate cleaned string
		translatedText, err := translateText(outputLang, str)
		if err != nil {
			log.Printf("failed to translate: %s\n", err.Error())
			return err
		}

		// print translated text
		fmt.Printf("\n%s\n", html.UnescapeString(translatedText))

		// write translated text to file
		_, err = fo.Write([]byte(html.UnescapeString(translatedText)))
		if err != nil {
			log.Printf("failed to write to output: %s\n", err.Error())
			return err
		}

	}

	return nil
}
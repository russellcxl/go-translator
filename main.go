package main

import (
	"bufio"
	"bytes"
	"cloud.google.com/go/translate"
	vision "cloud.google.com/go/vision/apiv1"
	"context"
	"fmt"
	"golang.org/x/text/language"
	"html"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	files, err := ioutil.ReadDir("images")
	if err != nil {
		log.Fatalf("failed to read images: %s\n", err.Error())
	}

	for _, f := range files {
		fmt.Printf("will read from %s; will write to %s.txt\n", f.Name(), strings.TrimSuffix(f.Name(), filepath.Ext(f.Name())))
		in := path.Join("images", f.Name())
		out := path.Join("output", fmt.Sprintf("%s.txt", strings.TrimSuffix(f.Name(), filepath.Ext(f.Name()))))
		fmt.Println(in, out)
		if err := readAndTranslate(in, out, "en"); err != nil {
			log.Fatalln(err.Error())
		}
	}
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

// detectText gets text from the Vision API for an image at the given file path.
func detectText(w io.Writer, file string) error {
	ctx := context.Background()

	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		return err
	}

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	image, err := vision.NewImageFromReader(f)
	if err != nil {
		return err
	}
	annotations, err := client.DetectTexts(ctx, image, nil, 10)
	if err != nil {
		return err
	}

	if len(annotations) == 0 {
		fmt.Fprintln(w, "No text found.")
	} else {
		fmt.Fprintln(w, "Text:")
		for _, annotation := range annotations {
			fmt.Fprintf(w, "%q\n", annotation.Description)
		}
	}

	return nil
}

func translateText(targetLanguage, text string) (string, error) {
	ctx := context.Background()

	lang, err := language.Parse(targetLanguage)
	if err != nil {
		return "", fmt.Errorf("language.Parse: %v", err)
	}

	client, err := translate.NewClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	resp, err := client.Translate(ctx, []string{text}, lang, nil)
	if err != nil {
		return "", fmt.Errorf("translate: %v", err)
	}
	if len(resp) == 0 {
		return "", fmt.Errorf("translate returned empty response to text: %s", text)
	}
	return resp[0].Text, nil
}

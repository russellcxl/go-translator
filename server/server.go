package server

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/russellcxl/go-translator/config"
	"github.com/russellcxl/go-translator/pkg/logger"
	"html"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Translator struct {
	logger logger.Logger
	config *config.Config
}

func NewTranslator(logger logger.Logger, cfg *config.Config) *Translator {
	return &Translator{
		logger: logger,
		config: cfg,
	}
}

func (t *Translator) Execute() error {
	clog := t.logger

	files, err := ioutil.ReadDir(t.config.Translator.InDir)
	if err != nil {
		clog.Errorf("failed to read from input directory: %s\n", err.Error())
		return err
	}

	for _, f := range files {
		fileExt := filepath.Ext(f.Name())
		if fileExt != ".jpg" && fileExt != ".jpeg" && fileExt != ".png" {
			clog.Errorf("file %s not the correct type\n", f.Name())
			continue
		}

		clog.Infof("will read from %s; will write to %s.txt\n", f.Name(), strings.TrimSuffix(f.Name(), filepath.Ext(f.Name())))

		in := path.Join(t.config.Translator.InDir, f.Name())
		outFile := fmt.Sprintf("%s.txt", strings.TrimSuffix(f.Name(), filepath.Ext(f.Name())))
		out := path.Join(t.config.Translator.OutDir, outFile)

		if err = t.readAndTranslate(in, out, t.config.Translator.OutLang); err != nil {
			return err
		}

		clog.Infof("successfully read and translated file: %s\n", f.Name())
	}
	return nil
}

func (t *Translator) readAndTranslate(inPath, outPath, outputLang string) error {
	clog := t.logger

	buf := new(bytes.Buffer)

	err := detectText(buf, inPath)
	if err != nil {
		clog.Errorf("failed to detect text: %s\n", err.Error())
		return err
	}

	// prepare output file
	fo, err := os.Create(outPath)
	if err != nil {
		clog.Errorf("failed to create output file: %s\n", err.Error())
		return err
	}
	defer fo.Close()

	r := bufio.NewReader(buf)

	// read until '\n' because for some reason, the file is read twice
	b, err := r.ReadBytes('\n')
	if err != nil {
		clog.Errorf("failed to read: %s\n", err.Error())
		return err
	}

	var currByte int
	batchSize := 1024
	for currByte <= len(b) {

		var str string
		if len(b) - currByte < batchSize {
			str = string(b[currByte:])
		} else {
			str = string(b[currByte:currByte+batchSize])
		}

		translatedText, err := t.translateText(outputLang, str)
		if err != nil {
			return err
		}

		fmt.Printf("\n%s\n", html.UnescapeString(translatedText))

		// write translated text to file
		_, err = fo.Write([]byte(html.UnescapeString(translatedText)))
		if err != nil {
			clog.Errorf("failed to write to output: %s\n", err.Error())
			return err
		}

		currByte+=batchSize
	}

	return nil
}

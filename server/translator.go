package server

import (
	"cloud.google.com/go/translate"
	"context"
	"fmt"
	"golang.org/x/text/language"
	"log"
)

func translateText(targetLanguage, text string) (string, error) {
	ctx := context.Background()

	lang, err := language.Parse(targetLanguage)
	if err != nil {
		log.Printf("failed language.Parse: %s\n", err.Error())
		return "", err
	}

	client, err := translate.NewClient(ctx)
	if err != nil {
		log.Printf("failed to start new translator client: %s\n", err.Error())
		return "", err
	}
	defer client.Close()

	resp, err := client.Translate(ctx, []string{text}, lang, nil)
	if err != nil {
		log.Printf("failed to translate: %s\n", err.Error())
		return "", err
	}
	if len(resp) == 0 {
		return "", fmt.Errorf("translate returned empty response to text: %s", text)
	}
	return resp[0].Text, nil
}

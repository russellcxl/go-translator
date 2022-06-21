# go-translator

### How to use:

1. `export GOOGLE_APPLICATION_CREDENTIALS="PATH_TO_CREDENTIALS_JSON"` ([ref](https://cloud.google.com/docs/authentication/production))
2. `go build .`
3. `./go-translator -port XXXX`
4. Go to localhost:XXXX and upload any image (jpeg, jpg, png) containing text in a non-english language
5. Find the translated text files in `./images/output`

### Features
1. Google Vision and Google Translate APIs
2. Custom logger found in `./pkg/logger`
3. Webpage that accepts and parses images
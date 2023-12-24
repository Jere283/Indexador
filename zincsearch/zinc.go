package zinc

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Config struct {
	BaseURL  string
	Index    string
	Username string
	Password string
}

// using the ZincSearch API to create a document and add it to the index
func CreateDocument(bodyQuery []byte, config Config) {
	requestURL := fmt.Sprintf("%s/api/%s/_doc", config.BaseURL, config.Index)
	req, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(bodyQuery))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(config.Username, config.Password)
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	log.Println(res.StatusCode)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}

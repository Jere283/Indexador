package zinc

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func CreateDocument(bodyQuery []byte, index string) {

	requestURL := fmt.Sprintf("http://localhost:%d/api/%s/_doc", 4080, index)
	req, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(bodyQuery))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("admin", "Complexpass#123")
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

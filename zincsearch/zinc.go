package zinc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Config struct {
	BaseURL  string
	Index    string
	Username string
	Password string
}

// the structure of the API response from the searchDocument() function
type Hit struct {
	Index  string  `json:"_index"`
	Type   string  `json:"_type"`
	ID     string  `json:"_id"`
	Score  float64 `json:"_score"`
	Source struct {
		Timestamp  string `json:"@timestamp"`
		Athlete    string `json:"Athlete"`
		City       string `json:"City"`
		Country    string `json:"Country"`
		Discipline string `json:"Discipline"`
		Event      string `json:"Event"`
		Gender     string `json:"Gender"`
		Medal      string `json:"Medal"`
		Season     string `json:"Season"`
		Sport      string `json:"Sport"`
		Year       int    `json:"Year"`
	} `json:"_source"`
}

type HitsResponse struct {
	Hits struct {
		Hits []Hit `json:"hits"`
	} `json:"hits"`
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

func SearchDocument(word string) HitsResponse {
	var query = fmt.Sprintf(`{
        "search_type": "match",
        "query":
        {
            "term": "%s",
            "start_time": "2021-06-02T14:28:31.894Z",
            "end_time": "2023-12-27T15:28:31.894Z"
        },
        "from": 0,
        "max_results": 20,
        "_source": []
    }`, word)
	req, err := http.NewRequest("POST", "http://localhost:4080/api/olympics/_search", strings.NewReader(query))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)

	var hitsResponse HitsResponse
	err = json.NewDecoder(resp.Body).Decode(&hitsResponse)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", hitsResponse)

	return hitsResponse
}

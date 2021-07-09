package services

import (
	"io/ioutil"
	"log"
	"net/http"
    "encoding/json"
)

type HealthService struct{}
type Status struct {
	Message string `json:"message"`
}

func (p *HealthService) CheckStatus(url string) Status {
	resp, err := http.Get(url)

	if err != nil {
		log.Println(err)
		s := Status{}
		s.Message = "down"
		return s
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		s := Status{}
		s.Message = "down"
		return s
	}

	sb := string(body)

	status := Status{}
	json.Unmarshal([]byte(sb), &status)
	
	return status
}
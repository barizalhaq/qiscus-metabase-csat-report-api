package entities

import (
	"bytes"
	"csat-report-webhook/viewmodels"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Metabase struct {
	email     string
	password  string
	baseURL   string
	sessionID string
}

func NewMetabase() *Metabase {
	return &Metabase{
		email:     os.Getenv("METABASE_EMAIL"),
		password:  os.Getenv("METABASE_PASSWORD"),
		baseURL:   os.Getenv("METABASE_BASE_URL"),
		sessionID: "",
	}
}

func (m *Metabase) Login() *Metabase {
	url := m.baseURL + "/api/session"
	reqBody, _ := json.Marshal(map[string]string{
		"username": m.email,
		"password": m.password,
	})

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatalf("An Error Occured: %v", err)
		return m
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("An Error Occured: %v", err)
		return m
	}

	var metabaseAuthResponse viewmodels.MetabaseAuthResponse
	json.Unmarshal(body, &metabaseAuthResponse)

	m.sessionID = metabaseAuthResponse.ID

	return m
}

func (m *Metabase) GetQuestionData(cardID int, jsonReqBody []byte) viewmodels.MetabaseDataResponse {
	url := fmt.Sprintf("%s/api/card/%v/query", m.baseURL, cardID)

	client := http.Client{}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonReqBody))
	if err != nil {
		log.Fatalf("An Error Occured: %v", err)
		return viewmodels.MetabaseDataResponse{}
	}

	request.Header.Set("X-Metabase-Session", m.sessionID)
	request.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(request)
	if err != nil {
		log.Fatalf("An Error Occured: %v", err)
		return viewmodels.MetabaseDataResponse{}
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("An Error Occured: %v", err)
		return viewmodels.MetabaseDataResponse{}
	}

	var dataResponse viewmodels.MetabaseDataResponse
	json.Unmarshal(body, &dataResponse)

	return dataResponse
}

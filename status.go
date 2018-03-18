package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type statusResponse struct {
	Message string
	Data    *Status
}

type Status struct {
	DatabaseVersion int    `json:"database_version"`
	DatabaseStatus  string `json:"database_status"`
	DatabaseType    string `json:"database_type"`
}

func (pg *PGApi) GetStatus() (*Status, error) {
	resp, err := pg.makeRequest(http.MethodGet, "/api/status")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)

	if resp.StatusCode == http.StatusNotFound {
		return nil, errors.New("instance doesn't support status api")
	}
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, errors.New("invalid credientials")
	}

	var statusResp statusResponse
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&statusResp); err != nil {
		return nil, err
	}

	status := statusResp.Data
	return status, nil
}

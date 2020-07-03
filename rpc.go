package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

func GetEpochStartHeight() (int, error) {

	requestBody, err := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      "dontcare",
		"method":  "validators",
		"params":  []*string{nil},
	})
	if err != nil {
		return 0, err
	}

	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Post(ENDPOINT, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	type resultStruct struct {
		Result struct {
			EpochStartHeight int `json:"epoch_start_height"`
		} `json:"result"`
	}
	// parse body
	result := &resultStruct{}
	json.Unmarshal(body, result)

	return result.Result.EpochStartHeight, nil

}

func getLatestBlockHeight() (int, error) {

	requestBody, err := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      "dontcare",
		"method":  "status",
		"params":  []string{},
	})
	if err != nil {
		return 0, err
	}

	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Post(ENDPOINT, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	type resultStruct struct {
		Result struct {
			SyncInfo struct {
				LatestBlockHeight int `json:"latest_block_height"`
			} `json:"sync_info"`
		} `json:"result"`
	}
	// parse body
	result := &resultStruct{}
	json.Unmarshal(body, result)

	return result.Result.SyncInfo.LatestBlockHeight, err
}

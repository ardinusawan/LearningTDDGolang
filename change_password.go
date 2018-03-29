package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"time"
	"log"
	"encoding/json"
	"bytes"
	"errors"
	// "reflect"
	// "io"
 	//"github.com/gorilla/mux"
	// "github.com/stretchr/testify/assert"
)

type auth struct {
	Username string `json:"username"`
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}


type GetChangePasswordRequest interface {
	FetchBytes(url string) ([]byte, error)
}

type LiveGetChangePasswordRequest struct {

}

type PostChangePasswordRequest interface {
	SendData(url string, data string) ([]byte, error)
}

type LivePostChangePasswordRequest struct {

}

func (LiveGetChangePasswordRequest) FetchBytes(url string) ([]byte, error) {
	authClient := http.Client {
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "this is unit test")

	res, getErr := authClient.Do(req)
	if getErr != nil {

		log.Fatal(getErr)
	}
	if res.StatusCode != http.StatusOK {
    	fmt.Errorf("server didn’t respond 200 OK: %s", res.Status)
    	return []byte(`{"message":res.StatusCode}`), errors.New(res.Status)

  	} 
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	return body, nil
}

func (LivePostChangePasswordRequest) SendData(url string, data string) []byte {
    jsonStr := []byte(data)
    req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonStr))
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
    return body
}


func GetChangePassword(getChangePasswordRequest GetChangePasswordRequest, url string) (string, error) {
	bodyBytes, err := getChangePasswordRequest.FetchBytes(url)
	if err != nil {
		return string(bodyBytes), err
	}
	authResult := auth{}

	jsonErr := json.Unmarshal(bodyBytes, &authResult)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return authResult.Username, nil
}

// func PostChangePassword(postChangePasswordRequest PostChangePasswordRequest) string {
// 	url := ""
// 	var jsonStr = `{"username":"Ardi Nusawan","oldPassword":"1","newPassword":"2"}`
// 	bodyBytes := postChangePasswordRequest.SendData(url, jsonStr)
// 	authResult := auth{}

// 	jsonErr := json.Unmarshal(bodyBytes, &authResult)
// 	if jsonErr != nil {
// 		log.Fatal(jsonErr)
// 	}
// 	var response map[string]interface{}
// 	if err := json.Unmarshal(bodyBytes, &response); err != nil {
//         panic(err)
//     }
//     fmt.Println(response["status"])
// 	return authResult.Username

// }




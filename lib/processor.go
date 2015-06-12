package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"encoding/json"
	"errors"
	"strconv"
)

type JSend struct {
	Status           string `json:"status"`
	Message          string `json:"message"`
	Data             json.RawMessage `json:"data"`
}

type Account struct {
	Id              string    `json:"id"`
	Email           string    `json:"email"`
	AccountLimit    string    `json:"accountLimit"`
	AccountBalance  string    `json:"accountBalance"`
}

type Transaction struct {
	Id              uint32    `json:"id"`
	AccountId       string    `json:"accountId"`
	Amount          float64   `json:"amount"`
}

var accountIRI string = "http://processor.maxxjs.com/accounts";
var transactionIRI string = "http://processor.maxxjs.com/transactions";

func PostAccount(email string) error {
	_, err := http.PostForm(accountIRI, url.Values{"email": {email}})
	if err != nil {
		return err
	}
	
	return nil
}

func PostTransaction(email string, amount float32) error {
	var j JSend
	
	stringAmount := strconv.FormatFloat(float64(amount), 'f', 2, 64)
	
	resp, err := http.PostForm(transactionIRI, url.Values{"email": {email}, "amount": {stringAmount}})
	if err != nil {
		return err
	}
	
	body, _ := ioutil.ReadAll(resp.Body)
	
	err = json.Unmarshal(body, &j)
	if err != nil {
		fmt.Println(err)
		return err
	}
	
	if j.Status == "fail" || j.Status == "error" {
		return errors.New(j.Message)
	}
	
	return nil
}

// Parses the standard return format that may or may not contain a data or message
func GetAccount(email string) (*Account, error ){
	var account *Account
	var accounts []*Account
	
	j, err := GetJSend(accountIRI + "/" + email)
	
	if err != nil {
		return account, err
	}
	
	if (len(j.Data) > 0) {
		err := json.Unmarshal(j.Data, &accounts)
		if err != nil {
			fmt.Println(err)
		}
		
		return (accounts)[0], nil
	}
	
	return account, nil
}

// Parses the standard return format that will contain a status and may contain either contain data or message
func GetJSend(iri string) (JSend, error) {
	resp, err := http.Get(iri)
	if err != nil {
		fmt.Println(err)
	}
	
	var j JSend
	
	body, _ := ioutil.ReadAll(resp.Body)
	
	err = json.Unmarshal(body, &j)
	if err != nil {
		fmt.Println(err)
		return j, err
	}
	
	if j.Status == "fail" || j.Status == "error" {
		return j, errors.New(j.Message)
	}
	
	return j, nil
}


package client

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

// GetAllItems returns a list of all Items in the DB
func GetAllItems(itemType string, port string, verbose bool) []byte {

	host := viper.GetString("Host")

	url := "http://" + host + ":" + port + "/api/v1/" + itemType

	if verbose {
		fmt.Println("GET: " + url)
	}

	resp, err := http.Get(url)
	if err != nil {
		// handle error
		fmt.Println("An error occurred")
		fmt.Println(err)
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)

	return data
}

// DeleteItem deletes the given item
func DeleteItem(id string, itemType string, port string, verbose bool) []byte {

	host := viper.GetString("Host")

	// Create client
	client := &http.Client{}

	url := "http://" + host + ":" + port + "/api/v1/" + itemType + "/id/" + id

	if verbose {
		fmt.Println("DELETE: " + url)
	}

	// call /Item/id/{id}
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	return respBody

}

// DeleteItemNoIDURL deletes the given item
func DeleteItemNoIDURL(id string, itemType string, port string, verbose bool) []byte {

	host := viper.GetString("Host")

	url := "http://" + host + ":" + port + "/api/v1/" + itemType + "/" + id

	if verbose {
		fmt.Println("DELETE: " + url)
	}

	// Create client
	client := &http.Client{}
	// call /Item/id/{id}
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	return respBody
}

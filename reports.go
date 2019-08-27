package main

import (
	// "bytes"
	// "encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/pquerna/ffjson/ffjson"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
)

type Reports struct {
	Product        string       `json:"product"`
	Manufacturer    string       `json:"manufacturer"`
	Category  string       `json:"category"`
	VideoTitle  string  `json:"videoTitle"`
	VideoCode string `json:"videoCode"`
	DateReleased string `json:"dateReleased"`
}



func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	spaceClient := http.Client{
		Timeout: time.Second *20, // Maximum of 2 secs
	}


	// req.Header.Add("category", "Energy Crisis")

	
	choice:=0
	fmt.Println("1] show all reports/n2]By category/n3]by date range")
	fmt.Println("Enter your choice")
	fmt.Scan(&choice)
	
	switch(choice) {
	case 1:
	
		apiURL:= "https://thereportoftheweek-api.herokuapp.com/reports"

		req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		logginghelper.LogError(err)
	}
	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		logginghelper.LogError(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		logginghelper.LogError(readErr)
	}

	result := []Reports{}
	uErr := ffjson.Unmarshal(body, &result)
	if uErr != nil {
		logginghelper.LogError(uErr)

	}
	fmt.Println(result)

		break
	case 2:
	
		fmt.Println("enter category name:")
		category:=""
		fmt.Scan(&category)
		apiURL:= "https://thereportoftheweek-api.herokuapp.com/reports"
		req, err := http.NewRequest(http.MethodGet, apiURL, nil)
		if err != nil {
		logginghelper.LogError(err)
		}
		req.Header.Add("category",category)
		res, getErr := spaceClient.Do(req)
		if getErr != nil {
		
		
			logginghelper.LogError(getErr)
		}
	
		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			logginghelper.LogError(readErr)
		}
	
		result := []Reports{}
		uErr := ffjson.Unmarshal(body, &result)
		if uErr != nil {
			logginghelper.LogError(uErr)
	
		}
		fmt.Println(result)

		break
	case 3:
		fmt.Println("Enter Start date of released in the form of yyyy-mm-dd")
		startDate:=""
		fmt.Scan(&startDate)
		fmt.Println("Enter Start date of released in the form of yyyy-mm-dd")
		endDate:=""
		fmt.Scan(&endDate)
		between:=startDate + "-" + endDate
		apiURL:= "https://thereportoftheweek-api.herokuapp.com/reports"
		req, err := http.NewRequest(http.MethodGet, apiURL, nil)
		if err != nil {
		logginghelper.LogError(err)
		}
		req.Header.Add("between",between)
		res, getErr := spaceClient.Do(req)
		if getErr != nil {
			logginghelper.LogError(getErr)
		}
	
		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			logginghelper.LogError(readErr)
		}
	
		result := []Reports{}
		uErr := ffjson.Unmarshal(body, &result)
		if uErr != nil {
			logginghelper.LogError(uErr)
	
		}
		fmt.Println(result)
		break
	default:
		fmt.Println("Invalid choice")
		break
	}

	// fmt.Println(result)
	// http.Get()

}


func main() {
	http.HandleFunc("/", ServeHTTP)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}


}

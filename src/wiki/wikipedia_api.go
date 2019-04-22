package wiki

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type SearchResults struct {
	ready   bool
	Query   string
	Results []Result
}

type Result struct {
	Name, Description, URL string
}

func (sr *SearchResults) UnmarshalJSON(bs []byte) error {
	array := []interface{}{}
	if err := json.Unmarshal(bs, &array); err != nil {
		return err
	}
	sr.Query = array[0].(string)
	for i := range array[1].([]interface{}) {
		sr.Results = append(sr.Results, Result{
			array[1].([]interface{})[i].(string),
			array[2].([]interface{})[i].(string),
			array[3].([]interface{})[i].(string),
		})
	}
	return nil
}

func wikipediaAPI(request string) (answer []string) {

	// Make a slice for 3 elements
	s := make([]string, 3)

	// Send request
	if response, err := http.Get(request); err != nil {
		s[0] = "Wikipedia is not respond!"
	} else {
		defer response.Body.Close()

		// Read response
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		// Send data to the structure
		sr := &SearchResults{}
		if err = json.Unmarshal([]byte(contents), sr); err != nil {
			s[0] = "Something going wrong, try to change your question"
		}

		// Check if string is not empty
		if !sr.ready {
			s[0] = "Something going wrong, try to change your question"
		}

		// Going thourgh our struvt and send data to slice with answer
		for i := range sr.Results {
			s[i] = sr.Results[i].URL
		}
	}

	return s
}

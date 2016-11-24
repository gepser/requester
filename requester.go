package requester

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Get => does an HTTP get request and expect a 200 status code
func Get(url string, target interface{}) error {
	var request *http.Request
	var response *http.Response
	var err error
	client := &http.Client{}

	request, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	response, err = client.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return fmt.Errorf("Got status code: %d", response.StatusCode)
	}

	if target == nil {
		return fmt.Errorf("Target shouldn't be nil: %v", target)
	}

	json.NewDecoder(response.Body).Decode(target)

	return nil
}

package monitor

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/nerodesu017/lambdalabs-sniper/src/constants"
	"github.com/nerodesu017/lambdalabs-sniper/src/custom_httpclient"
)

type json_resp struct {
	Data map[string]struct{
		Instance_type struct{
			// Name string `json:"string"` // we'll use Description for now
			// Price int `json:"price_cents_per_hour"`
			Description string `json:"description"`
		} `json:"instance_type"`
		AvailableInRegions []struct{
			Name string `json:"name"`
			Description string `json:"description"`
		} `json:"regions_with_capacity_available"`
	} `json:"data"`
}

func FindAvailable() ([]*constants.GPU, error){
	gpus := make([]*constants.GPU, 0)
	var err error

	req, err := http.NewRequest("GET", string(constants.INSTANCE_TYPES_URL), nil)
	if err != nil {
		return nil, fmt.Errorf("error when fetching instances: %v", err)
	}

	for key, value := range constants.BASE_HEADERS {
		req.Header.Set(key, value)
	}

	resp, err := custom_httpclient.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error when doing 'instance types' request: %v", err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error - status code is NOT OK")
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error - status code is NOT OK")
	}

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("error when decoding gzip: %v", err)
		}
	case "deflate":
		return nil, fmt.Errorf("deflate not implemented")
	case "br":
		return nil, fmt.Errorf("br not implemented")
	default:
		reader = resp.Body
	}

	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("error with reading from reader: %v", err)
	}

	var json_response json_resp
	err = json.Unmarshal([]byte(data), &json_response)
	if err != nil {
		return nil, fmt.Errorf("error when unmarshaling instance types: %v", err)
	}

	for _, instance := range json_response.Data {
		if len(instance.AvailableInRegions) == 0 {
			continue
		}

		gpus = append(gpus, &constants.GPU{
			Name: instance.Instance_type.Description,
		})
	}
	
	return gpus, nil
}
package logged_in

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/nerodesu017/lambdalabs-sniper/src/constants"
	"github.com/nerodesu017/lambdalabs-sniper/src/custom_httpclient"
)

// find out if you are logged in and in which account

type json_resp struct{
	Data struct {
		Identity struct {
			Email string `json:"user_email"`
		} `json:"identity"`
	} `json:"data"`
}

// returns string=email, error=nil
// also adds sessionid cookie to the cookie jar
func GetEmail(session_id string) (string, error) {
	var err error

	custom_httpclient.AddCookie("sessionid", session_id)

	req, err := http.NewRequest("GET", string(constants.SPA_INIT_INFO_URL), nil)
	if err != nil {
		return "", fmt.Errorf("error when fetching account: %v", err)
	}

	for key, value := range constants.BASE_HEADERS {
		req.Header.Set(key, value)
	}

	resp, err := custom_httpclient.HttpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error when doing 'account fetching' request: %v", err)
	}

	// set XSRF
	// constants.Add_Csrf(custom_httpclient.CookieJar)

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error - status code is NOT OK")
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error - status code is NOT OK")
	}

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return "", fmt.Errorf("error when decoding gzip: %v", err)
		}
	case "deflate":
		return "", fmt.Errorf("deflate not implemented")
	case "br":
		return "", fmt.Errorf("br not implemented")
	default:
		reader = resp.Body
	}

	data, err := io.ReadAll(reader)
	if err != nil {
		return "", fmt.Errorf("error with reading from reader: %v", err)
	}

	var json_response json_resp
	err = json.Unmarshal([]byte(data), &json_response)
	if err != nil {
		return "", fmt.Errorf(`error when unmarshaling user email: %v`, err)
	}

	return json_response.Data.Identity.Email, nil
}
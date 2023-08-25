package constants

import (
	"fmt"
	"log"
	"net/http/cookiejar"
	"net/url"
)

var BASE_HEADERS map[string]string

func init() {
	BASE_HEADERS = make(map[string]string)
	BASE_HEADERS["Accept"] = "*/*"
	BASE_HEADERS["Accept-Encoding"] = "gzip" // "gzip, deflate, br"
	BASE_HEADERS["Accept-Language"] = "en-US,en;q=0.7"
	BASE_HEADERS["Cache-Control"] = "no-cache"
	// base_headers["Cookie"] = "" // get a cookie jar
	BASE_HEADERS["Pragma"] = "no-cache"
	BASE_HEADERS["Referer"] = "https://cloud.lambdalabs.com/instances"
	BASE_HEADERS["Sec-Ch-Ua"] = "\"Chromium\";v=\"116\", \"Not)A;Brand\";v=\"24\", \"Brave\";v=\"116\""
	BASE_HEADERS["Sec-Ch-Ua-Mobile"] = "?0"
	BASE_HEADERS["Sec-Ch-Ua-Platform"] = "Windows"
	BASE_HEADERS["Sec-Fetch-Dest"] = "empty"
	BASE_HEADERS["Sec-Fetch-Mode"] = "cors"
	BASE_HEADERS["Sec-Fetch-Site"] = "same-origin"
	BASE_HEADERS["Sec-Gpc"] = "1"
	BASE_HEADERS["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36"
	// base_headers["X-Csrftoken"] = "" // you should add this after fetching it, best idea
}

// we shouldn't modify constants, but I gotta make it work
func Add_Csrf(jar *cookiejar.Jar) error {
	parsed_url, err := url.Parse(string(BASE_URL))
	if err != nil {
		return fmt.Errorf("error when parsing url: %v", err)
	}
	
	cookies := jar.Cookies(parsed_url)
	for _, cookie := range cookies {
		if cookie.Name != "csrftoken" {
			continue
		}

		log.Printf("Found CSRF Token!\n")
		BASE_HEADERS["X-Csrftoken"] = cookie.Value
		return nil
	}
	log.Printf("CSRF Token not found...\n")

	return nil
}
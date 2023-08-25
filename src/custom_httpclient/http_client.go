package custom_httpclient

import (
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"

	"github.com/nerodesu017/lambdalabs-sniper/src/constants"
)

var HttpClient *http.Client
var CookieJar *cookiejar.Jar
var err error

func init() {
	CookieJar, err = cookiejar.New(nil)

	if err != nil {
		log.Fatalf("error while creating cookie jar: %v\n", err)
	}

	HttpClient = &http.Client{
		Jar:     CookieJar,
		Timeout: time.Second * 5, // 5 second timeout
	}
}

func AddCookie(cookie_name string, cookie_value string) error {
	cookie := http.Cookie{
		Name:  cookie_name,
		Value: cookie_value,
		Path: string(constants.HOST_NAME),
	}

	parsed_url, err := url.Parse(string(constants.BASE_URL))
	if err != nil {
		return fmt.Errorf("error when parsing url: %v", err)
	}

	cookies := make([]*http.Cookie, 1)
	cookies[0] = &cookie
	HttpClient.Jar.SetCookies(parsed_url, cookies)

	return nil
}
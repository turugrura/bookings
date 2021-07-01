package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{name: "home", url: "/", method: "GET", params: []postData{}, expectedStatusCode: http.StatusOK},
	{name: "about", url: "/about", method: "GET", params: []postData{}, expectedStatusCode: http.StatusOK},
	{name: "gq", url: "/generals-quarters", method: "GET", params: []postData{}, expectedStatusCode: http.StatusOK},
	{name: "ms", url: "/majors-suite", method: "GET", params: []postData{}, expectedStatusCode: http.StatusOK},
	{name: "sa", url: "/search-availability", method: "GET", params: []postData{}, expectedStatusCode: http.StatusOK},
	{name: "rs", url: "/reservation-summary", method: "GET", params: []postData{}, expectedStatusCode: http.StatusOK},
	{name: "c", url: "/contact", method: "GET", params: []postData{}, expectedStatusCode: http.StatusOK},
	{name: "mr", url: "/make-reservation", method: "GET", params: []postData{}, expectedStatusCode: http.StatusOK},
	{name: "search-availability", url: "/search-availability", method: "POST", params: []postData{
		{key: "start", value: "2020-01-01"},
		{key: "end", value: "2020-01-07"},
	}, expectedStatusCode: http.StatusOK},
	{name: "search-availability-json", url: "/search-availability-json", method: "POST", params: []postData{
		{key: "start", value: "2020-01-01"},
		{key: "end", value: "2020-01-07"},
	}, expectedStatusCode: http.StatusOK},
	{name: "make-reservation", url: "/make-reservation", method: "POST", params: []postData{
		{key: "first_name", value: "john"},
		{key: "last_name", value: "naja"},
		{key: "email", value: "a@a.com"},
		{key: "phone", value: "01234567890"},
	}, expectedStatusCode: http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, test := range theTests {
		if test.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + test.url)
			if err != nil {
				t.Log(err)
				t.Error(err)
			}

			if resp.StatusCode != test.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", test.name, test.expectedStatusCode, resp.StatusCode)
			}
		} else {
			values := url.Values{}
			for _, x := range test.params {
				values.Add(x.key, x.value)
			}

			resp, err := ts.Client().PostForm(ts.URL+test.url, values)

			if err != nil {
				t.Log(err)
				t.Error(err)
			}

			if resp.StatusCode != test.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", test.name, test.expectedStatusCode, resp.StatusCode)
			}
		}
	}
}

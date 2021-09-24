package handlers

import (
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
	//{"Home", "/", "GET", []postData{}, http.StatusOK},
	//
	//{"about", "/about", "GET", []postData{}, http.StatusOK},
	//{"gq", "/generals-quarters", "GET", []postData{}, http.StatusOK},
	//{"ms", "/majors-suite", "GET", []postData{}, http.StatusOK},
	//{"sa", "/search-availability", "GET", []postData{}, http.StatusOK},
	//{"makeR", "/make-reservation", "GET", []postData{}, http.StatusOK},
	//{"Reservation Summary", "/reservation-summary", "GET", []postData{}, http.StatusOK},
	//{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	//{"post-search", "/search-availability-json", "POST", []postData{
	//	{key: "start", value: "2020-01-01"},
	//	{key: "end", value: "2020-01-01"},
	//}, http.StatusOK},
	//{"post-search-json", "/search-availability", "POST", []postData{
	//	{key: "start", value: "2020-01-01"},
	//	{key: "end", value: "2020-01-01"},
	//}, http.StatusOK},
	//{"make reservation post", "/make-reservation", "POST", []postData{
	//	{key: "first_name", value: "John"},
	//	{key: "last_name", value: "James"},
	//	{key: "email", value: "me@here.com"},
	//	{key: "phone", value: "555-555-555"},
	//}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()
	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			// check the response
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf(
					"for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
			//test for POST
		} else {
			values := url.Values{}
			for _, x := range e.params {
				values.Add(x.key, x.value)
			}
			resp, err := ts.Client().PostForm(ts.URL+e.url, values)

			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			// check the response
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		}
	}
}

func TestRepository_MakeReservations(t *testing.T) {

}

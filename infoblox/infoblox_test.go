package infoblox

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// client is the Infoblox client being tested.
	client *Client

	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server
)

// setup sets up a test HTTP server along with a infoblox.Client that is
// configured to talk to that test server.  Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup() {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	// infoblox client configured to use test server
	client = NewClient(nil)
	url, _ := url.Parse(server.URL)
	client.BaseURL = url
}

// teardown closes the test HTTP server.
func teardown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, expected string) {
	if actual := r.Method; actual != expected {
		t.Errorf("Request method: %v, expected %v", actual, expected)
	}
}

func testHeader(t *testing.T, r *http.Request, header, expected string) {
	if actual := r.Header.Get(header); actual != expected {
		t.Errorf("Header.Get(%q) returned %q, expected %q", header, actual, expected)
	}
}

func testURLParseError(t *testing.T, err error) {
	if err == nil {
		t.Error("Expected an error to be returned")
	}
	if err, ok := err.(*url.Error); !ok || err.Op != "parse" {
		t.Errorf("Expected URL Parse error, actual %+v", err)
	}
}

func testBody(t *testing.T, r *http.Request, expected string) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Errorf("Error reading request body: %v", err)
	}
	if actual := string(b); actual != expected {
		t.Errorf("request body is %s, expected %s", actual, expected)
	}
}

func testJSONMarshal(t *testing.T, v interface{}, expected string) {
	j, err := json.Marshal(v)
	if err != nil {
		t.Errorf("Unable to marshal JSON for %v", err)
	}

	w := new(bytes.Buffer)
	err = json.Compact(w, []byte(expected))
	if err != nil {
		t.Errorf("String is not valid json: %s", expected)
	}

	if w.String() != string(j) {
		t.Errorf("json.Marshal(%q), returned %s, expected %s", v, j, w)
	}

	// Now let's go the other direction and make sure things unmarshal correctly
	u := reflect.ValueOf(v).Interface()
	if err := json.Unmarshal([]byte(expected), u); err != nil {
		t.Errorf("Unable to unmarshal JSON for %v", expected)
	}

	if !reflect.DeepEqual(v, u) {
		t.Errorf("json.Unmarshal(%q) returned %s, expected %s", expected, u, v)
	}
}

func TestNewClient(t *testing.T) {
	c := NewClient(nil)

	if actual, expected := c.BaseURL.String(), defaultBaseURL; actual != expected {
		t.Errorf("NewClient Base URL is %s, expected %s", actual, expected)
	}
}

func TestNewRequest(t *testing.T) {
	c := NewClient(nil)

	inURL, outURL := "/wapi/v"+WAPIVersion+"/foo", defaultBaseURL+"wapi/v"+WAPIVersion+"/foo"
	inBody, outBody := &Host{Name: String("host01")}, `{"name":"host01"}`+"\n"
	req, _ := c.NewRequest("GET", inURL, inBody)

	// test that relative url is correct
	if actual, expected := req.URL.String(), outURL; actual != expected {
		t.Errorf("NewRequest(%q) URL is %v, expected %v", inURL, actual, expected)
	}

	// test that the body was properly JSON encoded
	body, _ := ioutil.ReadAll(req.Body)
	if actual, expected := string(body), outBody; actual != expected {
		t.Errorf("NewRequest(%q) Body is %v, expected %v", body, actual, expected)
	}
}

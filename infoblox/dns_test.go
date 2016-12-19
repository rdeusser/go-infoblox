package infoblox

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestHostService_Get_Host(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/wapi/v%s/record:host", WAPIVersion), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"name":"host01"}`)
	})

	host, _, err := client.DNS.GetHost("host01")
	if err != nil {
		t.Errorf("DNS.GetHost returned error: %v", err)
	}

	expected := &Host{Name: String("host01")}
	if !reflect.DeepEqual(host, expected) {
		t.Errorf("DNS.GetHost returned %#v, expected %#v", host, expected)
	}
}

func TestHostService_Create_Host(t *testing.T) {
	setup()
	defer teardown()

	input := &Host{Name: String("host01")}

	mux.HandleFunc(fmt.Sprintf("/wapi/v%s/record:host", WAPIVersion), func(w http.ResponseWriter, r *http.Request) {
		v := new(Host)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, expected %+v", v, input)
		}

		fmt.Fprint(w, `{"name":"host01"}`)
	})

	host, _, err := client.DNS.CreateHost(input)
	if err != nil {
		t.Errorf("DNS.CreateHost returned error: %v", err)
	}

	expected := &Host{Name: String("host01")}
	if !reflect.DeepEqual(host, expected) {
		t.Errorf("DNS.CreateHost returned %+v, expected %+v", host, expected)
	}
}

package infoblox

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestPTRRecordService_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("%s/record:ptr", versionedURL), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"name":"host01"}`)
	})

	p, _, err := client.PTR.Get("host01")
	if err != nil {
		t.Errorf("PTR.Get returned error: %v", err)
	}

	expected := &PTR{Name: String("host01")}
	if !reflect.DeepEqual(p, expected) {
		t.Errorf("PTR.Get returned %#v, expected %#v", p, expected)
	}
}

func TestPTRRecordService_Create(t *testing.T) {
	setup()
	defer teardown()

	input := &PTR{Name: String("host01")}

	mux.HandleFunc(fmt.Sprintf("%s/record:ptr", versionedURL), func(w http.ResponseWriter, r *http.Request) {
		v := new(PTR)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, expected %+v", v, input)
		}

		fmt.Fprint(w, `{"name":"host01"}`)
	})

	p, _, err := client.PTR.Create(input)
	if err != nil {
		t.Errorf("PTR.Create returned error: %v", err)
	}

	expected := &PTR{Name: String("host01")}
	if !reflect.DeepEqual(p, expected) {
		t.Errorf("PTR.Create returned %+v, expected %+v", p, expected)
	}
}

func TestPTRRecordService_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("%s/record:ptr", versionedURL), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.PTR.Delete("host01")
	if err != nil {
		t.Errorf("PTR.Delete returned %v", err)
	}
}

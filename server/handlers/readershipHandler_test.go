package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"prog2005assignment1/server/shared"
	"reflect"
	"testing"
)

func Test_getCountriesWithLanguageWithTwoLetterLanguageCode(t *testing.T) {
	type args struct {
		w    http.ResponseWriter
		code string
	}
	tests := []struct {
		name string
		args args
		want []shared.Country
	}{
		{name: "Valid language code", args: args{nil, "mh"}, want: []shared.Country{{Iso31661Alpha3: "MHL", Iso31661Alpha2: "MH", OfficialName: "Marshall Islands", RegionName: "Oceania", SubRegionName: "Micronesia", Language: "mh"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCountriesWithLanguageWithTwoLetterLanguageCode(tt.args.w, tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCountriesWithLanguageWithTwoLetterLanguageCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getReadership(t *testing.T) {
	type args struct {
		w       http.ResponseWriter
		country shared.Country
	}
	tests := []struct {
		name        string
		args        args
		wantAtleast int
	}{
		{name: "Valid country", args: args{nil, shared.Country{Iso31661Alpha3: "MHL", Iso31661Alpha2: "MH", OfficialName: "Marshall Islands", RegionName: "Oceania", SubRegionName: "Micronesia", Language: "mh"}}, wantAtleast: 40000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getReadership(tt.args.w, tt.args.country); got < tt.wantAtleast {
				t.Errorf("getReadership() = %v, want atleast %v", got, tt.wantAtleast)
			}
		})
	}
}

func Test_handleReadershipGetRequest(t *testing.T) {
	readershipURL := "http://localhost:" + shared.DefaultPort + shared.ReadershipPath + "no"

	req, err := http.NewRequest("GET", readershipURL, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleReadershipGetRequest)

	handler.ServeHTTP(rr, req)

	// Test the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Test the content type
	expectedContentType := "application/json"
	if contentType := rr.Header().Get("content-type"); contentType != expectedContentType {
		t.Errorf("handler returned unexpected content type: got %v want %v",
			contentType, expectedContentType)
	}

	// Expecting an array of three readership structs
	// Here there are too many changing variables to test for exact values
	// We can only test for the structure of the JSON

	// Decode the JSON
	decoder := json.NewDecoder(rr.Body)
	var readerships []shared.Readership
	err = decoder.Decode(&readerships)
	if err != nil {
		t.Errorf("Failed to decode JSON: %v", err)
	}

	for _, readership := range readerships {
		if readership.Country == "" {
			t.Errorf("Expected country, got: %v", readership.Country)
		} else if reflect.TypeOf(readership.Country).Kind() != reflect.String {
			t.Errorf("Expected string, got: %v", reflect.TypeOf(readership.Country).Kind())
		}

		if readership.Isocode == "" {
			t.Errorf("Expected isocode, got: %v", readership.Isocode)
		} else if reflect.TypeOf(readership.Isocode).Kind() != reflect.String {
			t.Errorf("Expected string, got: %v", reflect.TypeOf(readership.Isocode).Kind())
		}

		if reflect.TypeOf(readership.Books).Kind() != reflect.Int {
			t.Errorf("Expected int, got: %v", reflect.TypeOf(readership.Books).Kind())
		}

		if reflect.TypeOf(readership.Authors).Kind() != reflect.Int {
			t.Errorf("Expected int, got: %v", reflect.TypeOf(readership.Authors).Kind())
		}

		if reflect.TypeOf(readership.Readership).Kind() != reflect.Int {
			t.Errorf("Expected int, got: %v", reflect.TypeOf(readership.Readership).Kind())
		}
	}

}

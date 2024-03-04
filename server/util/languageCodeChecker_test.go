package util

import (
	"net/http"
	"testing"
)

// TestLanguageCodeChecker tests the LanguageCodeChecker function
func TestLanguageCodeChecker(t *testing.T) {
	type args struct {
		languageCode   string
		responseWriter http.ResponseWriter
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Valid language code", args{"en", nil}, true},
		{"Invalid language code", args{"eng", nil}, false},
		{"Invalid language code", args{"e", nil}, false},
		{"Invalid language code", args{"", nil}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LanguageCodeChecker(tt.args.languageCode, tt.args.responseWriter); got != tt.want {
				t.Errorf("LanguageCodeChecker() = %v, want %v", got, tt.want)
			}
		})
	}
}

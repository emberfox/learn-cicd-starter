// File: test_auth.go

package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
    tests := []struct {
        name     string
        headers   http.Header
        expected string
        err      bool
    }{
        {"Valid API Key",
            http.Header{"Authorization": []string{"ApiKey abc123"}},
            "abc123", false},
        {"Empty Authorization Header",
            http.Header{},
            "", true},
        {"Malformed Authorization Header",
            http.Header{"Authorization": []string{"InvalidToken"}},
            "", true},
        {"Missing Token in Valid Header",
            http.Header{"Authorization": []string{"ApiKey"}},
            "", true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := GetAPIKey(convertHeaders(tt.headers))
            
            if err != nil && !tt.err {
                t.Errorf("Expected no error, but got %v", err)
            }
            
            if err == nil && tt.err {
                t.Error("Expected error, but got none")
            }
            
            if result != tt.expected {
                t.Errorf("Expected '%s', but got '%s'", tt.expected, result)
            }
        })
    }
}

func convertHeaders(headers http.Header) http.Header {
	converted := make(http.Header)
	for k, v := range headers {
			converted[k] = v
	}
	return converted
}

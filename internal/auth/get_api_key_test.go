package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name       string
		authHeader string
		wantKey    string
		wantErr    string
	}{
		{
			name:       "valid API key",
			authHeader: "ApiKey test-api-key",
			wantKey:    "test-api-key",
		},
		{
			name:    "missing authorization header",
			wantErr: ErrNoAuthHeaderIncluded.Error(),
		},
		{
			name:       "wrong authorization scheme",
			authHeader: "Bearer test-api-key",
			wantErr:    "malformed authorization header",
		},
		{
			name:       "missing API key",
			authHeader: "ApiKey",
			wantErr:    "malformed authorization header",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//headers := make(http.Header)
			if tt.authHeader != "" {
				headers.Set("Authorization", tt.authHeader)
			}

			gotKey, err := GetAPIKey(headers)
			if tt.wantErr != "" {
				if err == nil {
					t.Fatalf("GetAPIKey() error = nil, want %q", tt.wantErr)
				}
				if err.Error() != tt.wantErr {
					t.Fatalf("GetAPIKey() error = %q, want %q", err, tt.wantErr)
				}
				if gotKey != "" {
					t.Errorf("GetAPIKey() key = %q, want empty key", gotKey)
				}
				return
			}

			if err != nil {
				t.Fatalf("GetAPIKey() unexpected error: %v", err)
			}
			if gotKey != tt.wantKey {
				t.Errorf("GetAPIKey() key = %q, want %q", gotKey, tt.wantKey)
			}
		})
	}
}

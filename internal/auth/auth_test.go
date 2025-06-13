package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headerRandom := http.Header{}
	headerRandom.Add("Authorization", "test test")

	headerValid := http.Header{}
	headerValid.Add("", "ApiKey qwe123456")

	tests := []struct {
		name    string
		header  http.Header
		apiKey  string
		wantErr bool
	}{
		{
			name:    "Empty Api key",
			header:  http.Header{},
			apiKey:  "",
			wantErr: true,
		},
		{
			name:    "Missing ApiKey name",
			header:  http.Header{"Authorization": []string{"test test"}},
			apiKey:  "",
			wantErr: true,
		},
		{
			name:    "Return Valid ApiKey",
			header:  http.Header{"Authorization": []string{"ApiKey qwe123456"}},
			apiKey:  "qwe123456",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiKey, err := GetAPIKey(tt.header)
			if apiKey != tt.apiKey {
				t.Errorf("GetAPIKey() wanted %s, got %s", tt.apiKey, apiKey)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

}

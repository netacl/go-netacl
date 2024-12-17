package netacl

import (
	"net/http"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func setTestKey(t *testing.T) {
	// load .env and grab environment variable
	err := godotenv.Load(".env")
	if err != nil {
		t.Fatalf("failed to get api key: %v", err)
	}
	if c, err = NewClient(os.Getenv("NETACL_SECRET")); err != nil {
		t.Fatalf("failed to set secret key: %v", err)
	}
	logger.EnableDebug()
}

func TestRequestWithAPIClient(t *testing.T) {
	setTestKey(t)

	c.Request("/", http.MethodGet, "application/json", nil)
}

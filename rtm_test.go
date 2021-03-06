package rtm

import (
	"context"
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var (
	Ctx       = context.Background()
	GetClient func(t testing.TB) *Client
)

func getCreds(t testing.TB) (key, secret, token string) {
	t.Helper()

	key, secret = os.Getenv("RTM_TEST_KEY"), os.Getenv("RTM_TEST_SECRET")
	if key == "" || secret == "" {
		t.Fatal("set environment variables RTM_TEST_KEY and RTM_TEST_SECRET to run integration tests")
	}

	token = os.Getenv("RTM_TEST_TOKEN")
	if token != "" {
		return
	}

	client := &Client{
		APIKey:      key,
		APISecret:   secret,
		DebugLogger: t.Log,
	}

	frob, err := client.Auth().GetFrob(Ctx)
	if err != nil {
		t.Fatal(err)
	}

	u := client.AuthenticationURL(Delete, frob)
	log.Printf("Visit this URL: %s", u)

	for i := 0; i < 3; i++ {
		token, _ = client.Auth().GetToken(Ctx, frob)
		if token != "" {
			break
		}
		time.Sleep(3 * time.Second)
	}
	if token == "" {
		t.Fatal("failed to get authentication token")
	}
	log.Printf("Set environment variable `RTM_TEST_TOKEN` to %q for faster tests.", token)
	return
}

func decodeFile(t testing.TB, filename string, v interface{}) {
	t.Helper()

	b, err := ioutil.ReadFile(filepath.Join("testdata", filename))
	require.NoError(t, err)
	err = xml.Unmarshal(b, v)
	require.NoError(t, err)
}

func TestMain(m *testing.M) {
	log.SetFlags(0)
	log.SetPrefix("testmain: ")

	var key, secret, token string
	var getCredsOnce sync.Once
	GetClient = func(t testing.TB) *Client {
		if testing.Short() {
			t.Skip("-short passed, skipping integration test")
		}

		getCredsOnce.Do(func() {
			key, secret, token = getCreds(t)
		})
		if token == "" {
			t.Skip("no authentication token")
		}

		return &Client{
			APIKey:      key,
			APISecret:   secret,
			AuthToken:   token,
			DebugLogger: t.Log,
		}
	}

	os.Exit(m.Run())
}

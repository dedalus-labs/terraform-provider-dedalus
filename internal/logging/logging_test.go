package logging

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-log/tflogtest"
)

func TestMiddlewareLogsOnlyRequestMetadata(t *testing.T) {
	t.Parallel()
	var output bytes.Buffer
	ctx := tflogtest.RootLogger(context.Background(), &output)
	req, err := http.NewRequest(http.MethodPost, "https://example.com/resources?token=query-secret", strings.NewReader(`{"password":"request-secret"}`))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer authorization-secret")
	req.Header.Set("X-API-Key", "api-key-secret")
	resp := &http.Response{
		StatusCode: http.StatusCreated,
		Status:     "201 Created",
		Header:     http.Header{"Set-Cookie": {"session=cookie-secret"}},
		Body:       io.NopCloser(strings.NewReader(`{"credential":"response-secret"}`)),
	}
	got, err := Middleware(ctx)(req, func(sent *http.Request) (*http.Response, error) {
		if sent != req {
			t.Fatal("request identity changed")
		}
		return resp, nil
	})
	if err != nil || got != resp {
		t.Fatalf("response changed: response=%p error=%v", got, err)
	}
	for _, secret := range []string{"query-secret", "request-secret", "authorization-secret", "api-key-secret", "cookie-secret", "response-secret"} {
		if strings.Contains(output.String(), secret) {
			t.Errorf("log contains credential fixture %q", secret)
		}
	}
	entries, err := tflogtest.MultilineJSONDecode(&output)
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 2 || entries[0]["method"] != http.MethodPost || entries[0]["path"] != "/resources" || entries[1]["status"] != float64(http.StatusCreated) {
		t.Fatalf("missing request metadata: %v", entries)
	}
}

type unreadableBody struct {
	reads int
}

func (body *unreadableBody) Read([]byte) (int, error) {
	body.reads++
	return 0, errors.New("body must remain unread")
}

func (*unreadableBody) Close() error { return nil }

func TestMiddlewarePreservesBodiesAndTransportErrors(t *testing.T) {
	t.Parallel()
	req, err := http.NewRequest(http.MethodPost, "https://example.com/resources", nil)
	if err != nil {
		t.Fatal(err)
	}
	requestBody, responseBody := new(unreadableBody), new(unreadableBody)
	req.Body = requestBody
	resp := &http.Response{StatusCode: http.StatusBadGateway, Body: responseBody}
	transportError := errors.New("transport failed")
	called := false
	got, err := Middleware(context.Background())(req, func(sent *http.Request) (*http.Response, error) {
		called = true
		if sent.Body != requestBody {
			t.Fatal("request body replaced")
		}
		return resp, transportError
	})
	if !called || got != resp || err != transportError {
		t.Fatalf("transport result changed: called=%v response=%p error=%v", called, got, err)
	}
	if requestBody.reads != 0 || responseBody.reads != 0 || got.Body != responseBody {
		t.Fatal("middleware consumed or replaced an HTTP body")
	}
}

package logging

import (
	"context"
	"net/http"

	"github.com/dedalus-labs/dedalus-go/option"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func Middleware(ctx context.Context) option.Middleware {
	return func(req *http.Request, next option.MiddlewareNext) (*http.Response, error) {
		// Credentials can appear in either headers or bodies, including API responses.
		// Reading bodies here also changes streaming and transport error behavior.
		if req != nil {
			tflog.Debug(ctx, "HTTP request", map[string]interface{}{
				"method": req.Method,
				"path":   req.URL.Path,
			})
		}
		resp, err := next(req)
		if resp != nil {
			tflog.Debug(ctx, "HTTP response", map[string]interface{}{"status": resp.StatusCode})
		}
		return resp, err
	}
}

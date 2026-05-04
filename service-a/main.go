package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	serviceBURL := app.Config.Get("SERVICE_B_URL")

	app.GET("/call-b", func(ctx *gofr.Context) (any, error) {
		if serviceBURL == "" {
			return nil, fmt.Errorf("SERVICE_B_URL is not set")
		}

		url := serviceBURL + "/hello"
		ctx.Logger.Infof("calling service-b at %s", url)

		req, err := http.NewRequestWithContext(ctx.Context, http.MethodGet, url, nil)
		if err != nil {
			return nil, err
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var bResp map[string]any
		if err := json.Unmarshal(body, &bResp); err != nil {
			return nil, err
		}

		return map[string]any{
			"service":          "service-a",
			"called":           url,
			"service_b_status": resp.StatusCode,
			"service_b_body":   bResp,
		}, nil
	})

	app.Run()
}

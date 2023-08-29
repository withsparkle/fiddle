package fetcher

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func New(client Client) Fetcher {
	return Fetcher{client}
}

type Fetcher struct{ c Client }

func (f Fetcher) Fetch(ctx context.Context, src string) (*http.Response, error) {
	u, err := url.Parse(src)
	if err != nil {
		return nil, fmt.Errorf("invalid url %q: %w", src, err)
	}

	resp, err := f.fetch(ctx, u.String())
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (f Fetcher) fetch(ctx context.Context, url string) (*http.Response, error) {
	resp, err := f.verify(ctx, url)
	if err != nil {
		return nil, err
	}

	req := resp.Request.Clone(ctx)
	req.Method = http.MethodGet

	resp, err = f.c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to send %s request: %w", req.Method, err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid status code: %d", resp.StatusCode)
	}
	if !isValid(resp.Header.Get("Content-Type")) {
		return nil, fmt.Errorf("invalid content type: %s", resp.Header.Get("Content-Type"))
	}
	return resp, nil
}

func (f Fetcher) verify(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodHead, url, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %w", err)
	}
	req.Header.Set("Accept", strings.Join(accepted[:], ","))
	req.Header.Set("User-Agent", userAgent)

	resp, err := f.c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to send %s request: %w", req.Method, err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid status code: %d", resp.StatusCode)
	}
	if !isValid(resp.Header.Get("Content-Type")) {
		return nil, fmt.Errorf("invalid content type: %s", resp.Header.Get("Content-Type"))
	}
	return resp, nil
}

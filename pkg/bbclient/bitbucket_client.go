/*
Copyright 2019 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package bbclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"knative.dev/pkg/logging"
)

const (
	defaultBaseUrl   = "https://api.bitbucket.org/2.0"
	defaultUserAgent = "go-bitbucket"
	mediaTypeJson    = "application/json"
)

// Hook struct to marshal/unmarshal BitBucket requests/responses.
type Hook struct {
	URL         string   `json:"url,omitempty"`
	Description string   `json:"description,omitempty"`
	Events      []string `json:"events,omitempty"`
	Active      bool     `json:"active,omitempty"`
	UUID        string   `json:"uuid,omitempty"`
}

// Client struct used to send http.Requests to BitBucket.
type Client struct {
	client    *http.Client
	logger    *zap.SugaredLogger
	username  string
	password  string
	baseUrl   *url.URL
	userAgent string
}

type User struct {
	Username string
	Password string
}

type userKey struct{}

type ChainableContext func(ctx context.Context) context.Context

func WithAuth(user string, password string) ChainableContext {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, userKey{}, User{
			Username: user,
			Password: password,
		})
	}
}

func fromContext(ctx context.Context) (*User, bool) {
	u, ok := ctx.Value(userKey{}).(*User)
	return u, ok
}

// NewClient creates a new Client for sending http.Requests to BitBucket.
func NewClient(ctx context.Context, user string, password string) *Client {
	logger := logging.FromContext(ctx)
	timeout := time.Duration(5 * time.Second)
	httpClient := http.Client{
		Timeout: timeout,
	}
	baseUrl, _ := url.Parse(defaultBaseUrl)
	return &Client{
		client:    &httpClient,
		logger:    logger,
		username:  user,
		password:  password,
		baseUrl:   baseUrl,
		userAgent: defaultUserAgent,
	}
}

type WebhookOptions struct {
	Uuid           string
	ConsumerKey    string
	ConsumerSecret string
	Domain         string
	Owner          string
	Repo           string
	Events         []string
}

// CreateHook creates a WebHook for 'owner' and 'repo'.
func (c *Client) CreateHook(options *WebhookOptions, hook *Hook) (*Hook, error) {
	requestBody, err := json.Marshal(hook)
	os.Stdout.Write(requestBody)

	var urlStr string
	if options.Repo == "" {
		// For every repo of the owner.
		urlStr = fmt.Sprintf("teams/%s/hooks", options.Owner)
	} else {
		// For a specific repo of the owner.
		urlStr = fmt.Sprintf("repositories/%s/%s/hooks", options.Owner, options.Repo)
	}

	err = c.doRequest("POST", urlStr, string(requestBody), hook)
	return hook, err
}

// DeleteHook deletes the WebHook 'hookUUID' previously registered for 'owner' and 'repo'.
func (c *Client) DeleteHook(owner, repo, hookUUID string) error {
	var urlStr string
	if repo == "" {
		urlStr = fmt.Sprintf("teams/%s/hooks/%s", owner, hookUUID)
	} else {
		urlStr = fmt.Sprintf("repositories/%s/%s/hooks/%s", owner, repo, hookUUID)
	}

	return c.doRequest("DELETE", urlStr, "", nil)
}

// doRequest performs an http.Request to BitBucket. If v is not nil, it attempts to unmarshal the response to
// that particular struct.
func (c *Client) doRequest(method, urlStr string, body string, v interface{}) error {
	u, err := c.baseUrl.Parse(fmt.Sprintf("%s/%s", defaultBaseUrl, urlStr))
	if err != nil {
		return err
	}

	c.logger.Infof("BitBucket Request URL %s", u.String())

	b := strings.NewReader(body)
	req, err := http.NewRequest(method, u.String(), b)
	// req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Content-Type", mediaTypeJson)
	// req.Header.Set("Accept", mediaTypeJson)
	req.SetBasicAuth(c.username, c.password)

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	// Just checking for 200, 201, and 204 status codes as those are the success status codes for creating and deleting hooks.
	if (resp.StatusCode != http.StatusOK) && (resp.StatusCode != http.StatusCreated) && (resp.StatusCode != http.StatusNoContent) {
		return fmt.Errorf("invalid status %q: %v", resp.Status, resp)
	}

	if v != nil {
		if resp.Body != nil {
			decErr := json.NewDecoder(resp.Body).Decode(v)
			if decErr != nil {
				err = decErr
			}
			return err
		}
	}
	return nil
}

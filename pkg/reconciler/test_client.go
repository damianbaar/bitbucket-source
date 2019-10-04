package reconciler

import (
	"context"
	"fmt"

	bbclient "github.com/nachocano/bitbucket-source/pkg/bbclient"
	"knative.dev/pkg/logging"
)

const (
	// Environment variable containing the HTTP port
	envPort = "PORT"

	// Environment variable containing BitBucket UUID.
	envUUID = "BITBUCKET_UUID"
)

func CallHook(key string, secret string, owner string, repo string, events []string) (*bbclient.Hook, error) {
	ctx := context.TODO()
	logger := logging.FromContext(ctx)
	options := &bbclient.WebhookOptions{
		Domain: "buildit.org",
		Owner:  owner,
		Repo:   repo,
		Events: events,
	}

	bbClient := bbclient.NewClient(ctx, key, secret)
	hookConfig := HookConfig(options)

	var h *bbclient.Hook
	h, err := bbClient.CreateHook(options, &hookConfig)

	if err != nil {
		return h, fmt.Errorf("failed to Create the BitBucket Webhook: %v", err)
	}
	logger.Infof("Created BitBucket WebHook: %+v", h)
	return h, nil
}

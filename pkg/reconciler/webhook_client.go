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

package reconciler

import (
	"context"
	"fmt"

	"github.com/nachocano/bitbucket-source/pkg/bbclient"
	"knative.dev/pkg/logging"
)

// type WebhookOptions = webhookOptions

type webhookClient interface {
	Create(ctx context.Context, options *bbclient.WebhookOptions) (string, error)
	Delete(ctx context.Context, options *bbclient.WebhookOptions) error
}

type bitBucketWebhookClient struct{}

func (client bitBucketWebhookClient) Create(ctx context.Context, options *bbclient.WebhookOptions) (string, error) {
	logger := logging.FromContext(ctx)

	logger.Info("Creating BitBucket WebHook")

	bbClient := createBitBucketClient(ctx, options)

	hook := hookConfig(options)

	var h *bbclient.Hook
	h, err := bbClient.CreateHook(options, &hook)

	if err != nil {
		return "", fmt.Errorf("failed to Create the BitBucket Webhook: %v", err)
	}
	logger.Infof("Created BitBucket WebHook: %+v", h)
	return h.UUID, nil
}

func (client bitBucketWebhookClient) Delete(ctx context.Context, options *bbclient.WebhookOptions) error {
	logger := logging.FromContext(ctx)

	logger.Info("Deleting BitBucket WebHook: %q", options.Uuid)

	bbClient := createBitBucketClient(ctx, options)

	err := bbClient.DeleteHook(options.Owner, options.Repo, options.Uuid)

	if err != nil {
		return fmt.Errorf("failed to Delete the BitBucket Webhook: %v", err)
	}

	logger.Infof("Deleted BitBucket Webhook: %s", options.Uuid)
	return nil
}

func createBitBucketClient(ctx context.Context, options *bbclient.WebhookOptions) *bbclient.Client {
	return bbclient.NewClient(ctx, options.ConsumerKey, options.ConsumerSecret)
}

func hookConfig(options *bbclient.WebhookOptions) bbclient.Hook {
	hook := bbclient.Hook{
		Description: "knative-sources",
		URL:         fmt.Sprintf("http://%s", options.Domain),
		Events:      options.Events,
		Active:      true,
	}
	return hook
}

var HookConfig = hookConfig

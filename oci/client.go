package oci

import (
	"context"
	"fmt"

	fclient "github.com/fluxcd/pkg/oci/client"
	"github.com/google/go-containerregistry/pkg/crane"
)

type Client struct {
	url    string
	client *fclient.Client
}

func NewClient(url string) *Client {
	options := []crane.Option{
		crane.WithUserAgent("ocm-controller/v1"),
	}
	client := fclient.NewClient(options)
	return &Client{
		url:    url,
		client: client,
	}
}

// Push takes a path, creates an archive of the files in it and pushes the content to the OCI registry.
func (c *Client) Push(ctx context.Context, artifactPath, sourcePath string, metadata fclient.Metadata) error {
	if err := c.client.Build(artifactPath, sourcePath, nil); err != nil {
		return fmt.Errorf("failed to create archive of the fetched artifacts: %w", err)
	}
	if _, err := c.client.Push(ctx, c.url, sourcePath, metadata, nil); err != nil {
		return fmt.Errorf("failed to push oci image: %w", err)
	}
	return nil
}

// Pull downloads an artifact from an OCI repository and extracts the content to the given directory.
func (c *Client) Pull(ctx context.Context, outDir string) (*fclient.Metadata, error) {
	metadata, err := c.client.Pull(ctx, c.url, outDir)
	if err != nil {
		return nil, fmt.Errorf("failed to pull oci image: %w", err)
	}
	return metadata, nil
}

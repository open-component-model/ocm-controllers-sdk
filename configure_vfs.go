// SPDX-FileCopyrightText: 2022 SAP SE or an SAP affiliate company and Open Component Model contributors.
//
// SPDX-License-Identifier: Apache-2.0

package ocmcontrollerssdk

import (
	"bytes"
	"context"
	"fmt"

	"github.com/containers/image/v5/pkg/compression"
	"github.com/mandelsoft/vfs/pkg/osfs"
	"github.com/mandelsoft/vfs/pkg/vfs"

	"github.com/open-component-model/ocm/pkg/contexts/ocm"
	ocmmeta "github.com/open-component-model/ocm/pkg/contexts/ocm/compdesc/meta/v1"
	"github.com/open-component-model/ocm/pkg/contexts/ocm/cpi"
	"github.com/open-component-model/ocm/pkg/utils/tarutils"
)

// ConfigureTemplateFilesystem takes a version access and a resource name and sets up a virtual filesystem to work with.
func ConfigureTemplateFilesystem(ctx context.Context, cv ocm.ComponentVersionAccess, resourceName string) (vfs.FileSystem, error) {
	// get the template
	_, templateBytes, err := GetResourceForComponentVersion(cv, resourceName)
	if err != nil {
		return nil, fmt.Errorf("template error: %w", err)
	}

	// setup virtual filesystem
	virtualFS, err := osfs.NewTempFileSystem()
	if err != nil {
		return nil, fmt.Errorf("fs error: %w", err)
	}

	// extract the template
	if err := tarutils.ExtractTarToFs(virtualFS, templateBytes); err != nil {
		return nil, fmt.Errorf("extract tar error: %w", err)
	}

	return virtualFS, nil
}

// GetResourceForComponentVersion returns the raw resource and the access for that resource.
func GetResourceForComponentVersion(cv ocm.ComponentVersionAccess, resourceName string) (ocm.ResourceAccess, *bytes.Buffer, error) {
	resource, err := cv.GetResource(ocmmeta.NewIdentity(resourceName))
	if err != nil {
		return nil, nil, err
	}

	rd, err := cpi.ResourceReader(resource)
	if err != nil {
		return nil, nil, err
	}
	defer rd.Close()

	decompress, _, err := compression.AutoDecompress(rd)
	if err != nil {
		return nil, nil, err
	}

	data := new(bytes.Buffer)
	if _, err := data.ReadFrom(decompress); err != nil {
		return nil, nil, err
	}

	return resource, data, nil
}

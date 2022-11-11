// SPDX-FileCopyrightText: 2022 SAP SE or an SAP affiliate company and Open Component Model contributors.
//
// SPDX-License-Identifier: Apache-2.0

package ocmcontrollerssdk

import (
	"fmt"
	"time"
)

// GetSnapshotName generates a snapshot name that should be used by all controllers.
func GetSnapshotName(repository, resourceName string) string {
	return fmt.Sprintf("%s/%s:%d", repository, resourceName, time.Now().Unix())
}

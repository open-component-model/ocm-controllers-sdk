package ocmcontrollerssdk

import (
	"fmt"

	"github.com/open-component-model/ocm/pkg/contexts/oci/repositories/ocireg"
	"github.com/open-component-model/ocm/pkg/contexts/ocm"
	"github.com/open-component-model/ocm/pkg/contexts/ocm/repositories/genericocireg"
)

func GetComponentVersion(ctx ocm.Context, session ocm.Session, repositoryURL, name, version string) (ocm.ComponentVersionAccess, error) {
	// configure the repository access
	repoSpec := genericocireg.NewRepositorySpec(ocireg.NewRepositorySpec(repositoryURL), nil)
	repo, err := session.LookupRepository(ctx, repoSpec)
	if err != nil {
		return nil, fmt.Errorf("repo error: %w", err)
	}

	// get the component version
	cv, err := session.LookupComponentVersion(repo, name, version)
	if err != nil {
		return nil, fmt.Errorf("component error: %w", err)
	}

	return cv, nil
}

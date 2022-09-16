package ocmcontrollerssdk

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// GetParentObject returns the first find in the owner references of a given object.
// T denotes the type that the user wants returned.
func GetParentObject[T client.Object](ctx context.Context, c client.Client, obj client.Object, kind, group string) (T, error) {
	var result T
	for _, ref := range obj.GetOwnerReferences() {
		if ref.Kind != kind {
			continue
		}

		gv, err := schema.ParseGroupVersion(ref.APIVersion)
		if err != nil {
			return result, err
		}

		if gv.Group != group {
			continue
		}

		key := client.ObjectKey{
			Namespace: obj.GetNamespace(),
			Name:      ref.Name,
		}

		if err := c.Get(ctx, key, result); err != nil {
			return result, fmt.Errorf("failed to get parent Source: %w", err)
		}

		return result, nil
	}

	return result, fmt.Errorf("parent not found")
}

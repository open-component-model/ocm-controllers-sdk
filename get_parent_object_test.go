package ocmcontrollerssdk

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestGetParentObject(t *testing.T) {
	pod := &corev1.Pod{}
	pod.OwnerReferences = append(pod.OwnerReferences, metav1.OwnerReference{
		APIVersion: "x-delivery.github.com/v1alpha1",
		Kind:       "kind",
		Name:       "my-name",
	})
	// get fake kube client
	fakeClient := fake.NewClientBuilder()
	client := fakeClient.WithObjects(&corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "my-name",
		},
	}).Build()

	result := &corev1.Pod{}
	err := GetParentObject(context.Background(), client, "kind", "x-delivery.github.com", pod, result)
	assert.NoError(t, err)
	assert.Equal(t, "my-name", result.Name)
}

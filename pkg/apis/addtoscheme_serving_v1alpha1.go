package apis

import (
	servingv1alpha1 "github.com/knative/serving/pkg/apis/serving/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes, servingv1alpha1.SchemeBuilder.AddToScheme)
}

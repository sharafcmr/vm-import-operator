package apis

import (
	"github.com/machacekondra/vm-import-operator/pkg/apis/v2v/v1alpha3"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes, v1alpha3.SchemeBuilder.AddToScheme)
}

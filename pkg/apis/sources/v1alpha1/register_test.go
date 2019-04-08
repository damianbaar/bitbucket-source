package v1alpha1

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func TestResource(t *testing.T) {
	want := schema.GroupResource{
		Group:    "sources.nachocano.org",
		Resource: "foo",
	}

	got := Resource("foo")

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("unexpected resource (-want, +got) = %v", diff)
	}
}

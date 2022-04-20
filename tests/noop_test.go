// This package depends on:
// * pulumi/pkg
// * pulumi-pascal/pkg
package tests

import (
	"testing"

	"github.com/pulumi/pulumi-pascal/pkg"
)

func TestNoOp(t *testing.T) {
	pkg.NoopUseImportsInPkg()
}

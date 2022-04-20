// This package depends on:
// * pulumi/pkg
// * pulumi-pascal/pkg
// * pulumi-pascal/pkg/codegen (as a peer/extension of the codegen package when under test)

package codegen

import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/pkg/v3/codegen/testing/test"

	"github.com/pulumi/pulumi-pascal/pkg"
)

func TestNoopUseImportsInCodegen(t *testing.T) {
	NoopUseImportsInCodegen()
	pkg.NoopUseImportsInPkg()
	noop(codegen.CleanDir)
	noop(schema.Alias{})
	noop(test.AwsSchema)
}

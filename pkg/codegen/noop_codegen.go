package codegen

import (
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"github.com/pulumi/pulumi/pkg/v3/codegen/hcl2/model"
	"github.com/pulumi/pulumi/pkg/v3/codegen/hcl2/model/format"
	"github.com/pulumi/pulumi/pkg/v3/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v3/codegen/pcl"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/pkg/v3/codegen/testing/test"
	"github.com/pulumi/pulumi/pkg/v3/codegen/testing/utils"
	"github.com/pulumi/pulumi/pkg/v3/resource/deploy/deploytest"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/cmdutil"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func noop(value interface{}) {}

func NoopUseImportsInCodegen() {
	noop(codegen.CleanDir)
	noop(model.Attribute{})
	noop(format.Formatter{})
	noop(syntax.AttributeTokens{})
	noop(pcl.Component{})
	noop(schema.PackageSpec{})
	noop(test.AwsSchema)
	noop(utils.AWS)
	noop(deploytest.Analyzer{})
	noop(diag.DefaultSink)
	noop(resource.Archive{})
	noop(plugin.AllPlugins)
	noop(tokens.AsName)
	noop(cmdutil.Emoji)
	noop(contract.Assert)
	noop(pulumi.AnyOutput{})
}

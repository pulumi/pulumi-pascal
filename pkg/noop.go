package pkg

import (
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/cmdutil"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/rpcutil"
	"github.com/pulumi/pulumi/sdk/v3/go/common/version"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func noop(value interface{}) {}

func NoopUseImportsInPkg() {
	noop(codegen.CleanDir)
	noop(schema.Alias{})
	noop(diag.Debug)
	noop(resource.Archive{})
	noop(plugin.AllPlugins)
	noop(cmdutil.ArgsFunc)
	noop(logging.CreateFilter)
	noop(rpcutil.GrpcChannelOptions)
	noop(version.Version)
	noop(pulumi.AdditionalSecretOutputs)
	noop(config.Config{})
}

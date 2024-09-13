package main

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/aws/awsstaticwebsite"
	"github.com/plantoncloud/pulumi-module-golang-commons/pkg/stackinput"
	"github.com/plantoncloud/s3-website-pulumi-module/pkg"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		stackInput := &awsstaticwebsite.AwsStaticWebsiteStackInput{}

		if err := stackinput.LoadStackInput(ctx, stackInput); err != nil {
			return errors.Wrap(err, "failed to load stack-input")
		}

		s := &pkg.ResourceStack{
			StackInput: stackInput,
		}

		return s.Resources(ctx)
	})
}
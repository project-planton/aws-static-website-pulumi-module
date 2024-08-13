package outputs

import (
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/aws/s3website"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
)

func PulumiOutputsToStackOutputsConverter(pulumiOutputs auto.OutputMap,
	input *s3website.S3WebsiteStackInput) *s3website.S3WebsiteStackOutputs {
	return &s3website.S3WebsiteStackOutputs{}
}

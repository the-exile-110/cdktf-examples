package config

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9"
)

var (
	Region        = "ap-northeast-1"
	BackendBucket = "tf-test-bucket"
	DefaultTags   = &map[string]*string{
		"Type": jsii.String("tf-default-tag"),
	}
)

var AwsProviderConfig = &aws.AwsProviderConfig{
	Region:      jsii.String(Region),
	DefaultTags: &aws.AwsProviderDefaultTags{Tags: DefaultTags},
}

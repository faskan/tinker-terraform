package test

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestMyFirstVpc(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		TerraformBinary: "tofu", // use OpenTofu instead of terraform
		TerraformDir:    "../",

		Vars: map[string]interface{}{
			"cidr_block": "10.0.0.0/16",
			"vpc_name":   "my-first-vpc",
		},
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	vpcID := terraform.Output(t, terraformOptions, "vpc_id")

	sess, err := session.NewSession(&aws.Config{Region: aws.String("eu-central-1")})
	assert.NoError(t, err)

	ec2Client := ec2.New(sess)

	describeVpcOutput, err := ec2Client.DescribeVpcs(&ec2.DescribeVpcsInput{
		VpcIds: []*string{aws.String(vpcID)},
	})
	assert.NoError(t, err)
	assert.Len(t, describeVpcOutput.Vpcs, 1)

	vpc := describeVpcOutput.Vpcs[0]
	assert.Equal(t, "10.0.0.0/16", *vpc.CidrBlock)

	// DNS support
	dnsSupport, err := ec2Client.DescribeVpcAttribute(&ec2.DescribeVpcAttributeInput{
		VpcId:     vpc.VpcId,
		Attribute: aws.String("enableDnsSupport"),
	})
	assert.NoError(t, err)
	assert.True(t, *dnsSupport.EnableDnsSupport.Value)

	// DNS hostnames
	dnsHostnames, err := ec2Client.DescribeVpcAttribute(&ec2.DescribeVpcAttributeInput{
		VpcId:     vpc.VpcId,
		Attribute: aws.String("enableDnsHostnames"),
	})
	assert.NoError(t, err)
	assert.True(t, *dnsHostnames.EnableDnsHostnames.Value)
}

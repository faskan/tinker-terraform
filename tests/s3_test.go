package test

import (
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTofuS3Bucket(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		TerraformBinary: "tofu", // ← This tells Terratest to use OpenTofu
		TerraformDir:    "../terraform",
		Vars: map[string]interface{}{
			"env": "test",
		},
	}

	terraform.InitAndApply(t, terraformOptions)
	defer terraform.Destroy(t, terraformOptions)

	bucketName := terraform.Output(t, terraformOptions, "bucket_name")
	assert.True(t, strings.HasPrefix(bucketName, "tofu-s3-bucket-"))
}

package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestEc2InstanceProvisioning(t *testing.T) {
   t.Parallel()
  approvedRegions := []string{"us-east-1"}
  awsRegion := aws.GetRandomRegion(t, approvedRegions, nil)
	expectedName := fmt.Sprintf("terratest-aws-example-%s", random.UniqueId())
	terraformOptions := &terraform.Options{
      // The path to where our Terraform code is located
      TerraformDir: "../",

      // Variables to pass to our Terraform code using -var options
      Vars: map[string]interface{}{
         "instance_name": expectedName,
         "test_label": "yes",
      },

      // Environment variables to set when running Terraform
      EnvVars: map[string]string{
         "AWS_DEFAULT_REGION": awsRegion,
      },
   }
  // At the end of the test, run `terraform destroy`
  defer terraform.Destroy(t, terraformOptions)
  // Run `terraform init` and `terraform apply`
  terraform.InitAndApply(t, terraformOptions)
  instanceIdbyTF := terraform.Output(t, terraformOptions, "instance_id")

  // let's check that the instance is actually there by looking for it with it's tag:
  tagName := "Name"
  instanceIdbyTag := aws.GetEc2InstanceIdsByTag(t, awsRegion, tagName, expectedName)

  // check it's the instance just created
  assert.Equal(t, instanceIdbyTag, instanceIdbyTF)
}

func TestEc2InstanceProvisioningWithRandomRegion(t *testing.T) {
	t.Parallel()
	CanonicalAccountId := "137112412989" //amazon Account id
	approvedRegions := []string{"us-east-1","us-west-1","us-west-2"}
	awsRegion := aws.GetRandomRegion(t, approvedRegions, nil)
  expectedName := fmt.Sprintf("terratest-aws-example-%s", random.UniqueId())
	amiFilters := map[string][]string{
      "architecture": []string{"x86_64"},
      "virtualization-type": []string{"hvm"},
   }
  amiId := aws.GetMostRecentAmiId(t, awsRegion, CanonicalAccountId, amiFilters)
  terraformOptions := &terraform.Options{
      // The path to where our Terraform code is located
      TerraformDir: "../",

      // Variables to pass to our Terraform code using -var options
      Vars: map[string]interface{}{
         "instance_name": expectedName,
         "test_label": "yes",
         "ami_id": amiId,
      },

      // Environment variables to set when running Terraform
      EnvVars: map[string]string{
         "AWS_DEFAULT_REGION": awsRegion,
      },
   }
  // At the end of the test, run `terraform destroy`
  defer terraform.Destroy(t, terraformOptions)
  // Run `terraform init` and `terraform apply`
  terraform.InitAndApply(t, terraformOptions)
  actualInstanceId := []string{terraform.Output(t, terraformOptions, "instance_id")}

  // let's check that the instance is actually there by looking for it with it's tag:
  tagName := "Name"
  exptectedInstanceId := aws.GetEc2InstanceIdsByTag(t, awsRegion, tagName, expectedName)

  // check it's the instance just created
  assert.Equal(t, exptectedInstanceId, actualInstanceId)
}

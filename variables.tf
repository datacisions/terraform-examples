# ---------------------------------------------------------------------------------------------------------------------
# OPTIONAL PARAMETERS
# These parameters have reasonable defaults.
# ---------------------------------------------------------------------------------------------------------------------
variable "region" {
  description = "AWS region to deploy the instances"
  default = "us-east-1"
}

variable "instance_name" {
  description = "The name tag to set for the EC2 Instance."
  default     = "testing-example-instance"
}

variable "test_label" {
  description = "Tag to identify the instance as test instance."
  default = "yes"
}

variable "ami_id" {
  description = "Id of the ami to deploy, depends on region."
  default = "ami-b70554c8" // for region us-east-1
}

# ---------------------------------------------------------------------------------------------------------------------
# DEPLOY AN EC2 INSTANCE RUNNING UBUNTU
provider "aws" {
  region = "us-east-1"
  profile= "admin"

}

resource "aws_instance" "example" {
  ami           = "${var.ami_id}"
  instance_type = "t2.micro"

  tags {
    Name = "${var.instance_name}"
    Test = "${var.test_label}"
  }
}

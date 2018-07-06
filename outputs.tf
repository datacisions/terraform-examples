output "instance_id" {
  value = "${aws_instance.example.id}"
}

output "instance_private_ip" {
  value = "${aws_instance.example.private_ip}"
}

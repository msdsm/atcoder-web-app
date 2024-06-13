variable "aws_region" {
  description = "AWS region"
  type        = string
  default     = "us-west-1"
}

variable "aws_user" {
  description = "use name"
  type        = string
  default     = "msd_user"
}

variable "name" {
  description = "name"
  type        = string
  default     = "test"
}

variable "vpc_cidr_block" {
  description = "cidr_block of vpc"
  type        = string
  default     = "10.0.0.0/16"
}

variable "subnet_cidr_blocks" {
  description = "cidr_block of subnet"
  type        = list(string)
  default     = ["10.0.10.0/24", "10.0.20.0/24"]
}

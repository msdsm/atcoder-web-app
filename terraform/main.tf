provider "aws" {
  profile = var.aws_user
  region  = var.aws_region
}


data "aws_availability_zones" "available" {
  state = "available"
}

data "aws_ami" "amazon_linux" {
  most_recent = true
  owners      = ["amazon"]

  filter {
    name   = "name"
    values = ["amzn2-ami-hvm-*-x86_64-gp2"]
  }
}

# ec2
resource "aws_instance" "test" {
  ami                    = data.aws_ami.amazon_linux.id
  instance_type          = "t2.micro"
  user_data              = file("init-script.sh")
  subnet_id              = aws_subnet.test.id
  vpc_security_group_ids = [aws_security_group.test-sg.id]

  tags = {
    Name = "${var.name}-instance"
  }
}

# セキュリティグループ
resource "aws_security_group" "test-sg" {
  name   = "${var.name}-sg"
  vpc_id = aws_vpc.test.id
  // pgadminで使用
  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  // APIで使用
  ingress {
    from_port   = 8888
    to_port     = 8888
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  // ssh接続用
  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# VPC
resource "aws_vpc" "test" {
  cidr_block = var.vpc_cidr_block
  tags = {
    Name = "${var.name}-vpc"
  }
}

# サブネット
resource "aws_subnet" "test" {
  vpc_id                  = aws_vpc.test.id
  cidr_block              = var.subnet_cidr_blocks[0]
  availability_zone       = data.aws_availability_zones.available.names[0]
  map_public_ip_on_launch = true

  tags = {
    Name = "${var.name}-subnet"
  }
}

# インターネットゲートウェイ
resource "aws_internet_gateway" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
    Name = "${var.name}-igw"
  }
}

# ルートテーブル
resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.test.id
  }

  tags = {
    Name = "${var.name}-route-table"
  }
}

resource "aws_route_table_association" "test" {
  subnet_id      = aws_subnet.test.id
  route_table_id = aws_route_table.test.id
}

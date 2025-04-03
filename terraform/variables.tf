variable "aws_region" {
  description = "AWS region to deploy resources"
  type        = string
  default     = "us-east-1"
}

variable "container_image" {
  description = "Docker image for the SimpleTimeService container (e.g., 'yourusername/simpletimeservice:latest')"
  type        = string
}

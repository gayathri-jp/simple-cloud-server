
## SimpleTimeService & AWS Infrastructure Deployment

This repository contains a minimalist web service called SimpleTimeService built with Go and a Terraform configuration that deploys the service on AWS.

### Prerequisites

Before deploying, ensure you have the following installed and configured:

 - Download and install Go if you plan to build locally.
 -  Install from Docker's official website.
-   Install Terraform.
-   Install and configure via AWS CLI installation.

### Creating AWS Access Keys  

 1. Log in to the AWS Management Console:
 2. Access the IAM Service: In the IAM dashboard, click on "Users" in the sidebar. Select your IAM user from the list.
 3. Create New Access Key: Go to the `Security credentials` tab. Under the "Access keys" section, click `Create access key`.

  

### AWS Authentication

Before deploying the infrastructure with Terraform, you need to authenticate to AWS. Here are several methods to do so

Configure via Command Line:

Run the following command and follow the prompts to enter your AWS Access Key ID, Secret Access Key, region, and output format:

    aws configure

This command creates or updates the ~/.aws/credentials and ~/.aws/config files with your credentials and region settings.

## Application: SimpleTimeService

### Overview :
The SimpleTimeService is a small web server written in Go that returns a JSON response containing the current timestamp and the visitor's IP address. When a request is made to the root path (/), the service responds with:

    {
    "timestamp": "<current date and time>",
    "ip": "<the IP address of the visitor>"
    }
    
### Building and Running Locally :
1. Build the Docker Image
Navigate to the app directory and build the Docker image using:

       docker build -t docgayathri123/simpletimeservice:latest .

2. Run the Container
Run the container with:

    `docker run -p 8080:8080 docgayathri123/simpletimeservice:latest`

3. Test the Service
Open your browser or use curl to access:

    `http://127.0.0.1:8080`

  
## Terraform Infrastructure Deployment
### Overview:
The Terraform configuration in the terraform folder sets up the following AWS resources:

- VPC: A Virtual Private Cloud with 2 public and 2 private subnets.
- ECS Cluster: An ECS cluster to host your container.
- Task Definition & Service: A Fargate task/service running the SimpleTimeService container in private subnets.
- Application Load Balancer (ALB): Deployed in public subnets to route traffic to the ECS service.

### Steps to Deploy :
1. Navigate to the Terraform Directory
`cd terraform`
2. Initialize Terraform
`terraform init`
3. Review the Terraform Plan
`terraform plan`
4. Apply the Terraform Configuration
`terraform apply`
5. Access the Service
Use the provided ALB DNS name in your browser or via curl to access the SimpleTimeService.

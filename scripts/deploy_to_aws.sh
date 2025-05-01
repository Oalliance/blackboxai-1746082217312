#!/bin/bash
# Script to build Docker image, push to AWS ECR, and deploy to ECS Fargate

set -e

AWS_REGION="us-east-1"
ECR_REPO_NAME="logistics-marketplace"
IMAGE_TAG="latest"
CLUSTER_NAME="logistics-marketplace-cluster"
SERVICE_NAME="logistics-marketplace-service"
TASK_DEFINITION_FAMILY="logistics-marketplace-task"

# Get AWS account ID
ACCOUNT_ID=$(aws sts get-caller-identity --query Account --output text)

# Authenticate Docker to ECR
aws ecr get-login-password --region $AWS_REGION | docker login --username AWS --password-stdin $ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com

# Create ECR repository if not exists
aws ecr describe-repositories --repository-names $ECR_REPO_NAME --region $AWS_REGION || \
aws ecr create-repository --repository-name $ECR_REPO_NAME --region $AWS_REGION

# Build Docker image
docker build -t $ECR_REPO_NAME:$IMAGE_TAG .

# Tag Docker image for ECR
docker tag $ECR_REPO_NAME:$IMAGE_TAG $ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com/$ECR_REPO_NAME:$IMAGE_TAG

# Push Docker image to ECR
docker push $ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com/$ECR_REPO_NAME:$IMAGE_TAG

# Register ECS task definition (assumes taskdef.json exists)
aws ecs register-task-definition --cli-input-json file://taskdef.json

# Create ECS cluster if not exists
aws ecs describe-clusters --clusters $CLUSTER_NAME --region $AWS_REGION | grep ACTIVE || \
aws ecs create-cluster --cluster-name $CLUSTER_NAME --region $AWS_REGION

# Create or update ECS service
SERVICE_EXISTS=$(aws ecs describe-services --cluster $CLUSTER_NAME --services $SERVICE_NAME --region $AWS_REGION | jq -r '.services | length')

if [ "$SERVICE_EXISTS" -eq "0" ]; then
  aws ecs create-service --cluster $CLUSTER_NAME --service-name $SERVICE_NAME --task-definition $TASK_DEFINITION_FAMILY --desired-count 1 --launch-type FARGATE --network-configuration "awsvpcConfiguration={subnets=[subnet-xxxxxx],securityGroups=[sg-xxxxxx],assignPublicIp=ENABLED}" --region $AWS_REGION
else
  aws ecs update-service --cluster $CLUSTER_NAME --service $SERVICE_NAME --task-definition $TASK_DEFINITION_FAMILY --region $AWS_REGION
fi

echo "Deployment to AWS ECS Fargate completed."

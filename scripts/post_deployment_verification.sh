#!/bin/bash
# Post-Deployment Verification Checklist and Health Checks for AWS ECS deployment

set -e

CLUSTER_NAME="logistics-marketplace-cluster"
SERVICE_NAME="logistics-marketplace-service"
AWS_REGION="us-east-1"

echo "Starting post-deployment verification..."

# Check ECS service status
SERVICE_STATUS=$(aws ecs describe-services --cluster $CLUSTER_NAME --services $SERVICE_NAME --region $AWS_REGION --query "services[0].status" --output text)
echo "ECS Service Status: $SERVICE_STATUS"

if [ "$SERVICE_STATUS" != "ACTIVE" ]; then
  echo "Error: ECS service is not active."
  exit 1
fi

# Check running task count
RUNNING_COUNT=$(aws ecs describe-services --cluster $CLUSTER_NAME --services $SERVICE_NAME --region $AWS_REGION --query "services[0].runningCount" --output text)
DESIRED_COUNT=$(aws ecs describe-services --cluster $CLUSTER_NAME --services $SERVICE_NAME --region $AWS_REGION --query "services[0].desiredCount" --output text)
echo "Running tasks: $RUNNING_COUNT / Desired tasks: $DESIRED_COUNT"

if [ "$RUNNING_COUNT" -lt "$DESIRED_COUNT" ]; then
  echo "Warning: Not all desired tasks are running."
fi

# Get task ARNs
TASK_ARNS=$(aws ecs list-tasks --cluster $CLUSTER_NAME --service-name $SERVICE_NAME --region $AWS_REGION --query "taskArns[]" --output text)

if [ -z "$TASK_ARNS" ]; then
  echo "Error: No running tasks found."
  exit 1
fi

# Check task health status
for TASK_ARN in $TASK_ARNS; do
  HEALTH_STATUS=$(aws ecs describe-tasks --cluster $CLUSTER_NAME --tasks $TASK_ARN --region $AWS_REGION --query "tasks[0].healthStatus" --output text)
  echo "Task $TASK_ARN health status: $HEALTH_STATUS"
  if [ "$HEALTH_STATUS" != "HEALTHY" ]; then
    echo "Warning: Task $TASK_ARN is not healthy."
  fi
done

echo "Post-deployment verification completed successfully."

#!/bin/bash
set -euo pipefail

AWS="aws --endpoint-url=http://localhost:4566 --region us-east-1"
TOPIC_NAME="post-processed"
QUEUES=("check-failed-queue")

echo "Creating SNS topic: $TOPIC_NAME"
TOPIC_ARN=$($AWS sns create-topic --name "$TOPIC_NAME" --query 'TopicArn' --output text)
echo "SNS Topic ARN: $TOPIC_ARN"
echo

for QUEUE_NAME in "${QUEUES[@]}"; do
  echo "Creating SQS queue: $QUEUE_NAME"
  QUEUE_URL=$($AWS sqs create-queue --queue-name "$QUEUE_NAME" --query 'QueueUrl' --output text)
  echo "Queue URL: $QUEUE_URL"

  echo "Getting ARN for queue: $QUEUE_NAME"
  QUEUE_ARN=$($AWS sqs get-queue-attributes --queue-url "$QUEUE_URL" --attribute-names QueueArn --query 'Attributes.QueueArn' --output text)
  echo "Queue ARN: $QUEUE_ARN"

  echo "Subscribing $QUEUE_NAME to SNS topic"
  SUBSCRIPTION_ARN=$($AWS sns subscribe --topic-arn "$TOPIC_ARN" --protocol sqs --notification-endpoint "$QUEUE_ARN" --query 'SubscriptionArn' --output text)
  echo "Subscription ARN: $SUBSCRIPTION_ARN"

  echo "Setting SQS policy to allow SNS to publish to $QUEUE_NAME"
  POLICY_FILE=$(mktemp)

  jq -n --arg queueArn "$QUEUE_ARN" --arg topicArn "$TOPIC_ARN" '{
    Version: "2012-10-17",
    Statement: [
      {
        Effect: "Allow",
        Principal: "*",
        Action: "SQS:SendMessage",
        Resource: $queueArn,
        Condition: {
          ArnEquals: {
            "aws:SourceArn": $topicArn
          }
        }
      }
    ]
  }' > "$POLICY_FILE"

  $AWS sqs set-queue-attributes --queue-url "$QUEUE_URL" --attributes "Policy=file://$POLICY_FILE"

  rm "$POLICY_FILE"

  echo "✅ $QUEUE_NAME setup complete."
  echo
done

echo "🎉 All queues and subscriptions are configured."

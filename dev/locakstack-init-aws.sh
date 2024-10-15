#!/bin/bash

# ref https://docs.localstack.cloud/references/init-hooks/

# SNS TOPIC

export AWS_DEFAULT_REGION=ap-northeast-1

export AWS_ACCESS_KEY_ID=dummy
export AWS_SECRET_ACCESS_KEY=dummy
export AWS_DEFAULT_REGION=ap-northeast-1
export AWS_ENDPOINT_URL=http://localhost:4566

echo "SNS topic Creating..."

awslocal sns create-topic --name ddd-sample-event-order-order_created

awslocal sns create-topic --name ddd-sample-event-order-order_approved

awslocal sns create-topic --name ddd-sample-event-order-order_rejected

awslocal sns create-topic --name ddd-sample-command-order-order_approve

awslocal sns create-topic --name ddd-sample-command-order-order_reject

awslocal sns create-topic --name ddd-sample-event-kitchen-ticket_created

awslocal sns create-topic --name ddd-sample-event-kitchen-ticket_creation_failed

awslocal sns create-topic --name ddd-sample-event-kitchen-ticket_approved

awslocal sns create-topic --name ddd-sample-event-kitchen-ticket_rejected

awslocal sns create-topic --name ddd-sample-command-kitchen-ticket_create

awslocal sns create-topic --name ddd-sample-command-kitchen-ticket_approve

awslocal sns create-topic --name ddd-sample-command-kitchen-ticket_reject

awslocal sns create-topic --name ddd-sample-event-billing-card_authorized

awslocal sns create-topic --name ddd-sample-event-billing-card_authorize_failed

awslocal sns create-topic --name ddd-sample-command-billing-card_authorize

# OrderAPI SQS

echo "OrderAPI SQS Creating..."

awslocal sqs create-queue --queue-name ddd-sample-order-event-queque

awslocal sqs create-queue --queue-name ddd-sample-order-command-queque

awslocal sqs create-queue --queue-name ddd-sample-order-reply-queque

awslocal sqs create-queue --queue-name ddd-sample-order-dead-letter-queque

awslocal sqs set-queue-attributes \
--queue-url http://sqs.ap-northeast-1.localhost.localstack.cloud:4566/000000000000/ddd-sample-order-event-queque \
--attributes '{
    "RedrivePolicy": "{\"deadLetterTargetArn\":\"arn:aws:sqs:ap-northeast-1:000000000000:ddd-sample-order-dead-letter-queque\",\"maxReceiveCount\":\"3\"}"
}'

awslocal sqs set-queue-attributes \
--queue-url http://sqs.ap-northeast-1.localhost.localstack.cloud:4566/000000000000/ddd-sample-order-command-queque \
--attributes '{
    "RedrivePolicy": "{\"deadLetterTargetArn\":\"arn:aws:sqs:ap-northeast-1:000000000000:ddd-sample-order-dead-letter-queque\",\"maxReceiveCount\":\"3\"}"
}'

awslocal sqs set-queue-attributes \
--queue-url http://sqs.ap-northeast-1.localhost.localstack.cloud:4566/000000000000/ddd-sample-order-reply-queque \
--attributes '{
    "RedrivePolicy": "{\"deadLetterTargetArn\":\"arn:aws:sqs:ap-northeast-1:000000000000:ddd-sample-order-dead-letter-queque\",\"maxReceiveCount\":\"3\"}"
}'

# Topic SubScribe

echo "OrderAPI SubScribe Creating..."

awslocal sns subscribe --topic-arn "arn:aws:sns:ap-northeast-1:000000000000:ddd-sample-event-order-order_created" --protocol sqs --notification-endpoint "arn:aws:sqs:ap-northeast-1:000000000000:ddd-sample-order-event-queque"

awslocal sns subscribe --topic-arn "arn:aws:sns:ap-northeast-1:000000000000:ddd-sample-event-order-order_approved" --protocol sqs --notification-endpoint "arn:aws:sqs:ap-northeast-1:000000000000:ddd-sample-order-event-queque"

awslocal sns subscribe --topic-arn "arn:aws:sns:ap-northeast-1:000000000000:ddd-sample-event-order-order_rejected" --protocol sqs --notification-endpoint "arn:aws:sqs:ap-northeast-1:000000000000:ddd-sample-order-event-queque"

awslocal sns subscribe --topic-arn "arn:aws:sns:ap-northeast-1:000000000000:ddd-sample-event-kitchen-ticket_created" --protocol sqs --notification-endpoint "arn:aws:sqs:ap-northeast-1:000000000000:ddd-sample-order-reply-queque"

awslocal sns subscribe --topic-arn "arn:aws:sns:ap-northeast-1:000000000000:ddd-sample-event-kitchen-ticket_creation_failed" --protocol sqs --notification-endpoint "arn:aws:sqs:ap-northeast-1:000000000000:ddd-sample-order-reply-queque"

awslocal sns subscribe --topic-arn "arn:aws:sns:ap-northeast-1:000000000000:ddd-sample-event-kitchen-ticket_approved" --protocol sqs --notification-endpoint "arn:aws:sqs:ap-northeast-1:000000000000:ddd-sample-order-reply-queque"

awslocal sns subscribe --topic-arn "arn:aws:sns:ap-northeast-1:000000000000:ddd-sample-event-kitchen-ticket_rejected" --protocol sqs --notification-endpoint "arn:aws:sqs:ap-northeast-1:000000000000:ddd-sample-order-reply-queque"

awslocal sns subscribe --topic-arn "arn:aws:sns:ap-northeast-1:000000000000:ddd-sample-event-billing-card_authorized" --protocol sqs --notification-endpoint "arn:aws:sqs:ap-northeast-1:000000000000:ddd-sample-order-reply-queque"

awslocal sns subscribe --topic-arn "arn:aws:sns:ap-northeast-1:000000000000:ddd-sample-event-billing-card_authorize_failed" --protocol sqs --notification-endpoint "arn:aws:sqs:ap-northeast-1:000000000000:ddd-sample-order-reply-queque"

awslocal sns subscribe --topic-arn "arn:aws:sns:ap-northeast-1:000000000000:ddd-sample-command-order-order_approve" --protocol sqs --notification-endpoint "arn:aws:sqs:ap-northeast-1:000000000000:ddd-sample-order-command-queque"

awslocal sns subscribe --topic-arn "arn:aws:sns:ap-northeast-1:000000000000:ddd-sample-command-order-order_reject" --protocol sqs --notification-endpoint "arn:aws:sqs:ap-northeast-1:000000000000:ddd-sample-order-command-queque"

# KitchenAPI SQS

echo "KitchenAPI SQS Creating..."

awslocal sqs create-queue --queue-name ddd-sample-kitchen-event-queque

awslocal sqs create-queue --queue-name ddd-sample-kitchen-command-queque

awslocal sqs create-queue --queue-name ddd-sample-kitchen-reply-queque

awslocal sqs create-queue --queue-name ddd-sample-kitchen-dead-letter-queque

awslocal sqs set-queue-attributes \
--queue-url http://sqs.ap-northeast-1.localhost.localstack.cloud:4566/000000000000/ddd-sample-kitchen-event-queque \
--attributes '{
    "RedrivePolicy": "{\"deadLetterTargetArn\":\"arn:aws:sqs:ap-northeast-1:000000000000:ddd-sample-kitchen-dead-letter-queque\",\"maxReceiveCount\":\"3\"}"
}'

awslocal sqs set-queue-attributes \
--queue-url http://sqs.ap-northeast-1.localhost.localstack.cloud:4566/000000000000/ddd-sample-kitchen-command-queque \
--attributes '{
    "RedrivePolicy": "{\"deadLetterTargetArn\":\"arn:aws:sqs:ap-northeast-1:000000000000:ddd-sample-kitchen-dead-letter-queque\",\"maxReceiveCount\":\"3\"}"
}'

awslocal sqs set-queue-attributes \
--queue-url http://sqs.ap-northeast-1.localhost.localstack.cloud:4566/000000000000/ddd-sample-kitchen-reply-queque \
--attributes '{
    "RedrivePolicy": "{\"deadLetterTargetArn\":\"arn:aws:sqs:ap-northeast-1:000000000000:ddd-sample-kitchen-dead-letter-queque\",\"maxReceiveCount\":\"3\"}"
}'

echo "KitchenAPI SubScribe Creating..."

awslocal sns subscribe --topic-arn "arn:aws:sns:ap-northeast-1:000000000000:ddd-sample-command-kitchen-ticket_create" --protocol sqs --notification-endpoint "arn:aws:sqs:ap-northeast-1:000000000000:ddd-sample-kitchen-command-queque"

awslocal sns subscribe --topic-arn "arn:aws:sns:ap-northeast-1:000000000000:ddd-sample-command-kitchen-ticket_approve" --protocol sqs --notification-endpoint "arn:aws:sqs:ap-northeast-1:000000000000:ddd-sample-kitchen-command-queque"

awslocal sns subscribe --topic-arn "arn:aws:sns:ap-northeast-1:000000000000:ddd-sample-command-kitchen-ticket_reject" --protocol sqs --notification-endpoint "arn:aws:sqs:ap-northeast-1:000000000000:ddd-sample-kitchen-command-queque"

# BillingAPI SQS

echo "BillingAPI SQS Creating..."

awslocal sqs create-queue --queue-name ddd-sample-billing-event-queque

awslocal sqs create-queue --queue-name ddd-sample-billing-command-queque

awslocal sqs create-queue --queue-name ddd-sample-billing-reply-queque

awslocal sqs create-queue --queue-name ddd-sample-billing-dead-letter-queque

awslocal sqs set-queue-attributes \
--queue-url http://sqs.ap-northeast-1.localhost.localstack.cloud:4566/000000000000/ddd-sample-billing-event-queque \
--attributes '{
    "RedrivePolicy": "{\"deadLetterTargetArn\":\"arn:aws:sqs:ap-northeast-1:000000000000:ddd-sample-billing-dead-letter-queque\",\"maxReceiveCount\":\"3\"}"
}'

awslocal sqs set-queue-attributes \
--queue-url http://sqs.ap-northeast-1.localhost.localstack.cloud:4566/000000000000/ddd-sample-billing-command-queque \
--attributes '{
    "RedrivePolicy": "{\"deadLetterTargetArn\":\"arn:aws:sqs:ap-northeast-1:000000000000:ddd-sample-billing-dead-letter-queque\",\"maxReceiveCount\":\"3\"}"
}'

awslocal sqs set-queue-attributes \
--queue-url http://sqs.ap-northeast-1.localhost.localstack.cloud:4566/000000000000/ddd-sample-billing-reply-queque \
--attributes '{
    "RedrivePolicy": "{\"deadLetterTargetArn\":\"arn:aws:sqs:ap-northeast-1:000000000000:ddd-sample-billing-dead-letter-queque\",\"maxReceiveCount\":\"3\"}"
}'

echo "BillingAPI SubScribe Creating..."

awslocal sns subscribe --topic-arn "arn:aws:sns:ap-northeast-1:000000000000:ddd-sample-command-billing-card_authorize" --protocol sqs --notification-endpoint "arn:aws:sqs:ap-northeast-1:000000000000:ddd-sample-billing-command-queque"

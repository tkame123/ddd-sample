# localstack

> > [!WARNING]
> DockerFileなどに後ほどまとめるまでの覚書メモ

# 操作方法（CLIメモ)

```
# 起動
localstack start -d

# 確認
localstack status services
```

https://docs.localstack.cloud/user-guide/aws/sns/

https://docs.localstack.cloud/user-guide/aws/sqs/

# 構成

## 共通

### SNS TOPIC

### Topic

```
awslocal sns create-topic --name ddd-sample-event-order-order_created

awslocal sns create-topic --name ddd-sample-event-order-order_approved

awslocal sns create-topic --name ddd-sample-event-order-order_rejected
```

```
awslocal sns create-topic --name ddd-sample-event-kitchen-ticket_created

awslocal sns create-topic --name ddd-sample-event-kitchen-ticket_creation_failed

awslocal sns create-topic --name ddd-sample-event-kitchen-ticket_approved

awslocal sns create-topic --name ddd-sample-event-kitchen-ticket_rejected

awslocal sns create-topic --name ddd-sample-command-kitchen-ticket_create

awslocal sns create-topic --name ddd-sample-command-kitchen-ticket_approve

awslocal sns create-topic --name ddd-sample-command-kitchen-ticket_reject
```

```
awslocal sns create-topic --name ddd-sample-event-billing-card_authorized

awslocal sns create-topic --name ddd-sample-event-billing-card_authorize_failed

awslocal sns create-topic --name ddd-sample-command-billing-card_authorize
```

## OrderAPI

### SQS

```
awslocal sqs create-queue --queue-name ddd-sample-order-event-queque

awslocal sqs create-queue --queue-name ddd-sample-order-command-queque

awslocal sqs create-queue --queue-name ddd-sample-order-reply-queque
```

### Topic SubScribe

```
awslocal sns subscribe --topic-arn "arn:aws:sns:ap-northeast-1:000000000000:ddd-sample-event-order-order_created" --protocol sqs --notification-endpoint "arn:aws:sqs:ap-northeast-1:000000000000:ddd-sample-order-event-queque"

awslocal sns subscribe --topic-arn "arn:aws:sns:ap-northeast-1:000000000000:ddd-sample-event-order-order_approved" --protocol sqs --notification-endpoint "arn:aws:sqs:ap-northeast-1:000000000000:ddd-sample-order-event-queque"

awslocal sns subscribe --topic-arn "arn:aws:sns:ap-northeast-1:000000000000:ddd-sample-event-order-order_rejected" --protocol sqs --notification-endpoint "arn:aws:sqs:ap-northeast-1:000000000000:ddd-sample-order-event-queque"

awslocal sns subscribe --topic-arn "arn:aws:sns:ap-northeast-1:000000000000:ddd-sample-event-kitchen-ticket_created" --protocol sqs --notification-endpoint "arn:aws:sqs:ap-northeast-1:000000000000:ddd-sample-order-reply-queque"

awslocal sns subscribe --topic-arn "arn:aws:sns:ap-northeast-1:000000000000:ddd-sample-event-kitchen-ticket_creation_failed" --protocol sqs --notification-endpoint "arn:aws:sqs:ap-northeast-1:000000000000:ddd-sample-order-reply-queque"

awslocal sns subscribe --topic-arn "arn:aws:sns:ap-northeast-1:000000000000:ddd-sample-event-kitchen-ticket_approved" --protocol sqs --notification-endpoint "arn:aws:sqs:ap-northeast-1:000000000000:ddd-sample-order-reply-queque"

awslocal sns subscribe --topic-arn "arn:aws:sns:ap-northeast-1:000000000000:ddd-sample-event-kitchen-ticket_rejected" --protocol sqs --notification-endpoint "arn:aws:sqs:ap-northeast-1:000000000000:ddd-sample-order-reply-queque"

awslocal sns subscribe --topic-arn "arn:aws:sns:ap-northeast-1:000000000000:ddd-sample-event-billing-card_authorized" --protocol sqs --notification-endpoint "arn:aws:sqs:ap-northeast-1:000000000000:ddd-sample-order-reply-queque"

awslocal sns subscribe --topic-arn "arn:aws:sns:ap-northeast-1:000000000000:ddd-sample-event-billing-card_authorize_failed" --protocol sqs --notification-endpoint "arn:aws:sqs:ap-northeast-1:000000000000:ddd-sample-order-reply-queque"

```

## KitchenAPI

### SQS
```
awslocal sqs create-queue --queue-name ddd-sample-kitchen-event-queque

awslocal sqs create-queue --queue-name ddd-sample-kitchen-command-queque

awslocal sqs create-queue --queue-name ddd-sample-kitchen-reply-queque
```

### Topic SubScribe

```
awslocal sns subscribe --topic-arn "arn:aws:sns:ap-northeast-1:000000000000:ddd-sample-command-kitchen-ticket_create" --protocol sqs --notification-endpoint "arn:aws:sqs:ap-northeast-1:000000000000:ddd-sample-kitchen-command-queque"

awslocal sns subscribe --topic-arn "arn:aws:sns:ap-northeast-1:000000000000:ddd-sample-command-kitchen-ticket_approve" --protocol sqs --notification-endpoint "arn:aws:sqs:ap-northeast-1:000000000000:ddd-sample-kitchen-command-queque"

awslocal sns subscribe --topic-arn "arn:aws:sns:ap-northeast-1:000000000000:ddd-sample-command-kitchen-ticket_reject" --protocol sqs --notification-endpoint "arn:aws:sqs:ap-northeast-1:000000000000:ddd-sample-kitchen-command-queque"
```

## BillingAPI

### SQS
```
awslocal sqs create-queue --queue-name ddd-sample-billing-event-queque

awslocal sqs create-queue --queue-name ddd-sample-billing-command-queque

awslocal sqs create-queue --queue-name ddd-sample-billing-reply-queque
```

### Topic SubScribe

```
awslocal sns subscribe --topic-arn "arn:aws:sns:ap-northeast-1:000000000000:ddd-sample-command-billing-card_authorize" --protocol sqs --notification-endpoint "arn:aws:sqs:ap-northeast-1:000000000000:queue-name ddd-sample-billing-command-queque"
```

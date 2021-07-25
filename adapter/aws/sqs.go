package aws

import (
	"context"
	"fmt"
	"sync"

	"app/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/pkg/errors"
)

type SQS struct {
	sqs     *sqs.SQS
	baseURL string
}

func NewSQS(sess *session.Session) *SQS {
	cfg := config.GetConfig()

	return &SQS{
		sqs:     sqs.New(sess),
		baseURL: cfg.AWS.SQS.BaseURL,
	}
}

func (q *SQS) url(key string) string {
	return fmt.Sprintf("%s/%s", q.baseURL, key)
}

func (q *SQS) Send(ctx context.Context, key string, data string) error {
	params := &sqs.SendMessageInput{
		MessageBody:  aws.String(data),
		QueueUrl:     aws.String(q.url(key)),
		DelaySeconds: aws.Int64(1),
	}

	res, err := q.sqs.SendMessageWithContext(ctx, params)
	if err != nil {
		return errors.Wrapf(err, "faield to send message to `%s`", key)
	}

	fmt.Printf("SQSMessageID: %s\n", *res.MessageId)

	return nil
}

func (q *SQS) Recieve(ctx context.Context, key string, f func(data string) error) error {
	url := q.url(key)

	params := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(url),
		MaxNumberOfMessages: aws.Int64(5),
		WaitTimeSeconds:     aws.Int64(15),
	}

	res, err := q.sqs.ReceiveMessage(params)
	if err != nil {
		return errors.Wrapf(err, "failed to receive messages key=%s", key)
	}

	if len(res.Messages) == 0 {
		return nil
	}

	var wg sync.WaitGroup
	for _, m := range res.Messages {
		wg.Add(1)

		go func(msg *sqs.Message) {
			defer wg.Done()

			_ = f(*msg.Body)

			params := &sqs.DeleteMessageInput{
				QueueUrl:      aws.String(url),
				ReceiptHandle: msg.ReceiptHandle,
			}
			_, err = q.sqs.DeleteMessageWithContext(ctx, params)

			if err != nil {
				fmt.Println(err) // TODO
			}
		}(m)
	}

	wg.Wait()

	return nil
}

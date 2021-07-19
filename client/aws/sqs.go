package aws

import (
	"context"
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/gami/layered_arch_example/config"
	"github.com/gami/layered_arch_example/domain"
	"github.com/gami/layered_arch_example/domain/user"
	"github.com/pkg/errors"
)

type SQS struct {
	sqs     *sqs.SQS
	baseUrl string
}

func NewSQS(sess *session.Session) domain.Messenger {
	cfg := config.GetConfig()
	return &SQS{
		sqs:     sqs.New(sess),
		baseUrl: cfg.AWS.SQSBaseURL,
	}
}

func (q *SQS) url(key string) string {
	return fmt.Sprintf("%s/%s", q.baseUrl, key)
}

func (q *SQS) Send(ctx context.Context, key string, data domain.Message) error {
	s, err := data.Marshal()
	if err != nil {
		return err
	}

	params := &sqs.SendMessageInput{
		MessageBody:  aws.String(s),
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

func (q *SQS) RecieveUser(ctx context.Context, key string, f func(msg domain.Message) error) error {
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

			u := &user.User{}

			err := u.Unmarshal(*msg.Body)
			if err != nil {
				fmt.Println(err) //TODO
				return
			}

			_ = f(u)

			params := &sqs.DeleteMessageInput{
				QueueUrl:      aws.String(url),
				ReceiptHandle: msg.ReceiptHandle,
			}
			_, err = q.sqs.DeleteMessageWithContext(ctx, params)
			if err != nil {
				fmt.Println(err) //TODO
			}
		}(m)
	}

	wg.Wait()

	return nil
}

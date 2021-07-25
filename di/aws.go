package di

import (
	"app/adapter/aws"
	"app/config"
)

func InjectSQS() *aws.SQS {
	cfg := config.GetConfig()
	sess := aws.Session(cfg.AWS.Region) // panic if credential is invalid.
	return aws.NewSQS(
		sess,
	)
}

package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/aws/aws-sdk-go/aws/session"
)

func Config(region string) *aws.Config {
	return aws.NewConfig().WithRegion(region).WithCredentials(generateCreds())
}

func Session(region string) *session.Session {
	return session.Must(session.NewSession(Config(region)))
}

// generateCreds returns AWS Credentials from Env.
// If set AWS_CONTAINER_CREDENTIALS_RELATIVE_URI in env and in ECS, retrieve one as ECS Task Role.
// If set AWS_ACCESS_ID / AWS_SECRET_KEY in env, generate from env.
func generateCreds() *credentials.Credentials {
	return credentials.NewChainCredentials(
		[]credentials.Provider{
			defaults.RemoteCredProvider(*defaults.Config(), defaults.Handlers()),
			&credentials.EnvProvider{},
		},
	)
}

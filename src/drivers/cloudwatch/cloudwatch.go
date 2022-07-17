package cloudwatch

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/sts"
)

type Cloudwatch struct {
	Session       *session.Session
	CloudWatchLog *cloudwatchlogs.CloudWatchLogs
}

func InitCloudwatch() *Cloudwatch {
	env := os.Getenv("ENVIRONMENT")
	group := os.Getenv("CLOUDWATCH_LOG_GROUP")
	keyId := os.Getenv("CLOUDWATCH_KEY_ID")
	secretKey := os.Getenv("CLOUDWATCH_SECRET_KEY")
	region := os.Getenv("CLOUDWATCH_REGION")

	if group == "" || keyId == "" || secretKey == "" || region == "" || env == "development" {
		return &Cloudwatch{
			Session:       nil,
			CloudWatchLog: nil,
		}
	}

	session, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			Credentials: credentials.NewStaticCredentials(keyId, secretKey, ""),
			Region:      aws.String(region),
		},
	})
	if err != nil {
		panic("Not going to be able to write to cloud watch if you cant create a session")
	}

	_, err = sts.New(session).GetCallerIdentity(&sts.GetCallerIdentityInput{})
	if err != nil {
		panic("Couldn't Validate our aws credentials")
	}

	cloudwatchLog := cloudwatchlogs.New(session)

	return &Cloudwatch{
		Session:       session,
		CloudWatchLog: cloudwatchLog,
	}
}

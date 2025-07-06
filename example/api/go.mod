module api

go 1.24.4

replace github.com/olbrichattila/evmagic => ./../../

replace app => ./../app/

require github.com/olbrichattila/evmagic v0.0.0-00010101000000-000000000000

require (
	app v0.0.0-00010101000000-000000000000 // indirect
	github.com/ThreeDotsLabs/watermill v1.4.7 // indirect
	github.com/ThreeDotsLabs/watermill-aws v1.0.0 // indirect
	github.com/aws/aws-sdk-go-v2 v1.36.5 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.36 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.36 // indirect
	github.com/aws/aws-sdk-go-v2/service/sns v1.31.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/sqs v1.38.8 // indirect
	github.com/aws/smithy-go v1.22.4 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/lithammer/shortuuid/v3 v3.0.7 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/samber/lo v1.51.0 // indirect
	golang.org/x/text v0.22.0 // indirect
)

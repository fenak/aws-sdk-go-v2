module github.com/aws/aws-sdk-go-v2/internal/protocoltest/restxmlwithnamespace

go 1.15

require (
	github.com/aws/aws-sdk-go-v2 v1.16.6
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.13
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.7
	github.com/aws/smithy-go v1.12.0
	github.com/google/go-cmp v0.5.8
)

replace github.com/aws/aws-sdk-go-v2 => ../../../

replace github.com/aws/aws-sdk-go-v2/internal/configsources => ../../../internal/configsources/

replace github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 => ../../../internal/endpoints/v2/

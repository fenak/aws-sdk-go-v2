// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates an Amazon VPC with a private hosted zone. To perform the association,
// the VPC and the private hosted zone must already exist. You can't convert a
// public hosted zone into a private hosted zone. If you want to associate a VPC
// that was created by using one Amazon Web Services account with a private hosted
// zone that was created by using a different account, the Amazon Web Services
// account that created the private hosted zone must first submit a
// CreateVPCAssociationAuthorization request. Then the account that created the VPC
// must submit an AssociateVPCWithHostedZone request.
func (c *Client) AssociateVPCWithHostedZone(ctx context.Context, params *AssociateVPCWithHostedZoneInput, optFns ...func(*Options)) (*AssociateVPCWithHostedZoneOutput, error) {
	if params == nil {
		params = &AssociateVPCWithHostedZoneInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateVPCWithHostedZone", params, optFns, c.addOperationAssociateVPCWithHostedZoneMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateVPCWithHostedZoneOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A complex type that contains information about the request to associate a VPC
// with a private hosted zone.
type AssociateVPCWithHostedZoneInput struct {

	// The ID of the private hosted zone that you want to associate an Amazon VPC with.
	// Note that you can't associate a VPC with a hosted zone that doesn't have an
	// existing VPC association.
	//
	// This member is required.
	HostedZoneId *string

	// A complex type that contains information about the VPC that you want to
	// associate with a private hosted zone.
	//
	// This member is required.
	VPC *types.VPC

	// Optional: A comment about the association request.
	Comment *string

	noSmithyDocumentSerde
}

// A complex type that contains the response information for the
// AssociateVPCWithHostedZone request.
type AssociateVPCWithHostedZoneOutput struct {

	// A complex type that describes the changes made to your hosted zone.
	//
	// This member is required.
	ChangeInfo *types.ChangeInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateVPCWithHostedZoneMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpAssociateVPCWithHostedZone{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpAssociateVPCWithHostedZone{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpAssociateVPCWithHostedZoneValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateVPCWithHostedZone(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addSanitizeURLMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opAssociateVPCWithHostedZone(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "AssociateVPCWithHostedZone",
	}
}

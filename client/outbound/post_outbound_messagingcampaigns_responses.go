// Code generated by go-swagger; DO NOT EDIT.

package outbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostOutboundMessagingcampaignsReader is a Reader for the PostOutboundMessagingcampaigns structure.
type PostOutboundMessagingcampaignsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOutboundMessagingcampaignsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostOutboundMessagingcampaignsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostOutboundMessagingcampaignsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostOutboundMessagingcampaignsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostOutboundMessagingcampaignsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostOutboundMessagingcampaignsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostOutboundMessagingcampaignsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostOutboundMessagingcampaignsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostOutboundMessagingcampaignsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostOutboundMessagingcampaignsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostOutboundMessagingcampaignsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostOutboundMessagingcampaignsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostOutboundMessagingcampaignsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostOutboundMessagingcampaignsOK creates a PostOutboundMessagingcampaignsOK with default headers values
func NewPostOutboundMessagingcampaignsOK() *PostOutboundMessagingcampaignsOK {
	return &PostOutboundMessagingcampaignsOK{}
}

/*
PostOutboundMessagingcampaignsOK describes a response with status code 200, with default header values.

successful operation
*/
type PostOutboundMessagingcampaignsOK struct {
	Payload *models.MessagingCampaign
}

// IsSuccess returns true when this post outbound messagingcampaigns o k response has a 2xx status code
func (o *PostOutboundMessagingcampaignsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post outbound messagingcampaigns o k response has a 3xx status code
func (o *PostOutboundMessagingcampaignsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound messagingcampaigns o k response has a 4xx status code
func (o *PostOutboundMessagingcampaignsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post outbound messagingcampaigns o k response has a 5xx status code
func (o *PostOutboundMessagingcampaignsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound messagingcampaigns o k response a status code equal to that given
func (o *PostOutboundMessagingcampaignsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostOutboundMessagingcampaignsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/messagingcampaigns][%d] postOutboundMessagingcampaignsOK  %+v", 200, o.Payload)
}

func (o *PostOutboundMessagingcampaignsOK) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/messagingcampaigns][%d] postOutboundMessagingcampaignsOK  %+v", 200, o.Payload)
}

func (o *PostOutboundMessagingcampaignsOK) GetPayload() *models.MessagingCampaign {
	return o.Payload
}

func (o *PostOutboundMessagingcampaignsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessagingCampaign)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundMessagingcampaignsBadRequest creates a PostOutboundMessagingcampaignsBadRequest with default headers values
func NewPostOutboundMessagingcampaignsBadRequest() *PostOutboundMessagingcampaignsBadRequest {
	return &PostOutboundMessagingcampaignsBadRequest{}
}

/*
PostOutboundMessagingcampaignsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostOutboundMessagingcampaignsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound messagingcampaigns bad request response has a 2xx status code
func (o *PostOutboundMessagingcampaignsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound messagingcampaigns bad request response has a 3xx status code
func (o *PostOutboundMessagingcampaignsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound messagingcampaigns bad request response has a 4xx status code
func (o *PostOutboundMessagingcampaignsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound messagingcampaigns bad request response has a 5xx status code
func (o *PostOutboundMessagingcampaignsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound messagingcampaigns bad request response a status code equal to that given
func (o *PostOutboundMessagingcampaignsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostOutboundMessagingcampaignsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/messagingcampaigns][%d] postOutboundMessagingcampaignsBadRequest  %+v", 400, o.Payload)
}

func (o *PostOutboundMessagingcampaignsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/messagingcampaigns][%d] postOutboundMessagingcampaignsBadRequest  %+v", 400, o.Payload)
}

func (o *PostOutboundMessagingcampaignsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundMessagingcampaignsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundMessagingcampaignsUnauthorized creates a PostOutboundMessagingcampaignsUnauthorized with default headers values
func NewPostOutboundMessagingcampaignsUnauthorized() *PostOutboundMessagingcampaignsUnauthorized {
	return &PostOutboundMessagingcampaignsUnauthorized{}
}

/*
PostOutboundMessagingcampaignsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostOutboundMessagingcampaignsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound messagingcampaigns unauthorized response has a 2xx status code
func (o *PostOutboundMessagingcampaignsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound messagingcampaigns unauthorized response has a 3xx status code
func (o *PostOutboundMessagingcampaignsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound messagingcampaigns unauthorized response has a 4xx status code
func (o *PostOutboundMessagingcampaignsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound messagingcampaigns unauthorized response has a 5xx status code
func (o *PostOutboundMessagingcampaignsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound messagingcampaigns unauthorized response a status code equal to that given
func (o *PostOutboundMessagingcampaignsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostOutboundMessagingcampaignsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/messagingcampaigns][%d] postOutboundMessagingcampaignsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostOutboundMessagingcampaignsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/messagingcampaigns][%d] postOutboundMessagingcampaignsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostOutboundMessagingcampaignsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundMessagingcampaignsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundMessagingcampaignsForbidden creates a PostOutboundMessagingcampaignsForbidden with default headers values
func NewPostOutboundMessagingcampaignsForbidden() *PostOutboundMessagingcampaignsForbidden {
	return &PostOutboundMessagingcampaignsForbidden{}
}

/*
PostOutboundMessagingcampaignsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostOutboundMessagingcampaignsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound messagingcampaigns forbidden response has a 2xx status code
func (o *PostOutboundMessagingcampaignsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound messagingcampaigns forbidden response has a 3xx status code
func (o *PostOutboundMessagingcampaignsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound messagingcampaigns forbidden response has a 4xx status code
func (o *PostOutboundMessagingcampaignsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound messagingcampaigns forbidden response has a 5xx status code
func (o *PostOutboundMessagingcampaignsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound messagingcampaigns forbidden response a status code equal to that given
func (o *PostOutboundMessagingcampaignsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostOutboundMessagingcampaignsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/messagingcampaigns][%d] postOutboundMessagingcampaignsForbidden  %+v", 403, o.Payload)
}

func (o *PostOutboundMessagingcampaignsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/messagingcampaigns][%d] postOutboundMessagingcampaignsForbidden  %+v", 403, o.Payload)
}

func (o *PostOutboundMessagingcampaignsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundMessagingcampaignsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundMessagingcampaignsNotFound creates a PostOutboundMessagingcampaignsNotFound with default headers values
func NewPostOutboundMessagingcampaignsNotFound() *PostOutboundMessagingcampaignsNotFound {
	return &PostOutboundMessagingcampaignsNotFound{}
}

/*
PostOutboundMessagingcampaignsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostOutboundMessagingcampaignsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound messagingcampaigns not found response has a 2xx status code
func (o *PostOutboundMessagingcampaignsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound messagingcampaigns not found response has a 3xx status code
func (o *PostOutboundMessagingcampaignsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound messagingcampaigns not found response has a 4xx status code
func (o *PostOutboundMessagingcampaignsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound messagingcampaigns not found response has a 5xx status code
func (o *PostOutboundMessagingcampaignsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound messagingcampaigns not found response a status code equal to that given
func (o *PostOutboundMessagingcampaignsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostOutboundMessagingcampaignsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/messagingcampaigns][%d] postOutboundMessagingcampaignsNotFound  %+v", 404, o.Payload)
}

func (o *PostOutboundMessagingcampaignsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/messagingcampaigns][%d] postOutboundMessagingcampaignsNotFound  %+v", 404, o.Payload)
}

func (o *PostOutboundMessagingcampaignsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundMessagingcampaignsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundMessagingcampaignsRequestTimeout creates a PostOutboundMessagingcampaignsRequestTimeout with default headers values
func NewPostOutboundMessagingcampaignsRequestTimeout() *PostOutboundMessagingcampaignsRequestTimeout {
	return &PostOutboundMessagingcampaignsRequestTimeout{}
}

/*
PostOutboundMessagingcampaignsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostOutboundMessagingcampaignsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound messagingcampaigns request timeout response has a 2xx status code
func (o *PostOutboundMessagingcampaignsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound messagingcampaigns request timeout response has a 3xx status code
func (o *PostOutboundMessagingcampaignsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound messagingcampaigns request timeout response has a 4xx status code
func (o *PostOutboundMessagingcampaignsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound messagingcampaigns request timeout response has a 5xx status code
func (o *PostOutboundMessagingcampaignsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound messagingcampaigns request timeout response a status code equal to that given
func (o *PostOutboundMessagingcampaignsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostOutboundMessagingcampaignsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/messagingcampaigns][%d] postOutboundMessagingcampaignsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostOutboundMessagingcampaignsRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/messagingcampaigns][%d] postOutboundMessagingcampaignsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostOutboundMessagingcampaignsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundMessagingcampaignsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundMessagingcampaignsRequestEntityTooLarge creates a PostOutboundMessagingcampaignsRequestEntityTooLarge with default headers values
func NewPostOutboundMessagingcampaignsRequestEntityTooLarge() *PostOutboundMessagingcampaignsRequestEntityTooLarge {
	return &PostOutboundMessagingcampaignsRequestEntityTooLarge{}
}

/*
PostOutboundMessagingcampaignsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostOutboundMessagingcampaignsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound messagingcampaigns request entity too large response has a 2xx status code
func (o *PostOutboundMessagingcampaignsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound messagingcampaigns request entity too large response has a 3xx status code
func (o *PostOutboundMessagingcampaignsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound messagingcampaigns request entity too large response has a 4xx status code
func (o *PostOutboundMessagingcampaignsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound messagingcampaigns request entity too large response has a 5xx status code
func (o *PostOutboundMessagingcampaignsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound messagingcampaigns request entity too large response a status code equal to that given
func (o *PostOutboundMessagingcampaignsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostOutboundMessagingcampaignsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/messagingcampaigns][%d] postOutboundMessagingcampaignsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostOutboundMessagingcampaignsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/messagingcampaigns][%d] postOutboundMessagingcampaignsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostOutboundMessagingcampaignsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundMessagingcampaignsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundMessagingcampaignsUnsupportedMediaType creates a PostOutboundMessagingcampaignsUnsupportedMediaType with default headers values
func NewPostOutboundMessagingcampaignsUnsupportedMediaType() *PostOutboundMessagingcampaignsUnsupportedMediaType {
	return &PostOutboundMessagingcampaignsUnsupportedMediaType{}
}

/*
PostOutboundMessagingcampaignsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostOutboundMessagingcampaignsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound messagingcampaigns unsupported media type response has a 2xx status code
func (o *PostOutboundMessagingcampaignsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound messagingcampaigns unsupported media type response has a 3xx status code
func (o *PostOutboundMessagingcampaignsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound messagingcampaigns unsupported media type response has a 4xx status code
func (o *PostOutboundMessagingcampaignsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound messagingcampaigns unsupported media type response has a 5xx status code
func (o *PostOutboundMessagingcampaignsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound messagingcampaigns unsupported media type response a status code equal to that given
func (o *PostOutboundMessagingcampaignsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostOutboundMessagingcampaignsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/messagingcampaigns][%d] postOutboundMessagingcampaignsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostOutboundMessagingcampaignsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/messagingcampaigns][%d] postOutboundMessagingcampaignsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostOutboundMessagingcampaignsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundMessagingcampaignsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundMessagingcampaignsTooManyRequests creates a PostOutboundMessagingcampaignsTooManyRequests with default headers values
func NewPostOutboundMessagingcampaignsTooManyRequests() *PostOutboundMessagingcampaignsTooManyRequests {
	return &PostOutboundMessagingcampaignsTooManyRequests{}
}

/*
PostOutboundMessagingcampaignsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostOutboundMessagingcampaignsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound messagingcampaigns too many requests response has a 2xx status code
func (o *PostOutboundMessagingcampaignsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound messagingcampaigns too many requests response has a 3xx status code
func (o *PostOutboundMessagingcampaignsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound messagingcampaigns too many requests response has a 4xx status code
func (o *PostOutboundMessagingcampaignsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound messagingcampaigns too many requests response has a 5xx status code
func (o *PostOutboundMessagingcampaignsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound messagingcampaigns too many requests response a status code equal to that given
func (o *PostOutboundMessagingcampaignsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostOutboundMessagingcampaignsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/messagingcampaigns][%d] postOutboundMessagingcampaignsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostOutboundMessagingcampaignsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/messagingcampaigns][%d] postOutboundMessagingcampaignsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostOutboundMessagingcampaignsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundMessagingcampaignsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundMessagingcampaignsInternalServerError creates a PostOutboundMessagingcampaignsInternalServerError with default headers values
func NewPostOutboundMessagingcampaignsInternalServerError() *PostOutboundMessagingcampaignsInternalServerError {
	return &PostOutboundMessagingcampaignsInternalServerError{}
}

/*
PostOutboundMessagingcampaignsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostOutboundMessagingcampaignsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound messagingcampaigns internal server error response has a 2xx status code
func (o *PostOutboundMessagingcampaignsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound messagingcampaigns internal server error response has a 3xx status code
func (o *PostOutboundMessagingcampaignsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound messagingcampaigns internal server error response has a 4xx status code
func (o *PostOutboundMessagingcampaignsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post outbound messagingcampaigns internal server error response has a 5xx status code
func (o *PostOutboundMessagingcampaignsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post outbound messagingcampaigns internal server error response a status code equal to that given
func (o *PostOutboundMessagingcampaignsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostOutboundMessagingcampaignsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/messagingcampaigns][%d] postOutboundMessagingcampaignsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostOutboundMessagingcampaignsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/messagingcampaigns][%d] postOutboundMessagingcampaignsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostOutboundMessagingcampaignsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundMessagingcampaignsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundMessagingcampaignsServiceUnavailable creates a PostOutboundMessagingcampaignsServiceUnavailable with default headers values
func NewPostOutboundMessagingcampaignsServiceUnavailable() *PostOutboundMessagingcampaignsServiceUnavailable {
	return &PostOutboundMessagingcampaignsServiceUnavailable{}
}

/*
PostOutboundMessagingcampaignsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostOutboundMessagingcampaignsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound messagingcampaigns service unavailable response has a 2xx status code
func (o *PostOutboundMessagingcampaignsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound messagingcampaigns service unavailable response has a 3xx status code
func (o *PostOutboundMessagingcampaignsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound messagingcampaigns service unavailable response has a 4xx status code
func (o *PostOutboundMessagingcampaignsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post outbound messagingcampaigns service unavailable response has a 5xx status code
func (o *PostOutboundMessagingcampaignsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post outbound messagingcampaigns service unavailable response a status code equal to that given
func (o *PostOutboundMessagingcampaignsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostOutboundMessagingcampaignsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/messagingcampaigns][%d] postOutboundMessagingcampaignsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostOutboundMessagingcampaignsServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/messagingcampaigns][%d] postOutboundMessagingcampaignsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostOutboundMessagingcampaignsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundMessagingcampaignsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundMessagingcampaignsGatewayTimeout creates a PostOutboundMessagingcampaignsGatewayTimeout with default headers values
func NewPostOutboundMessagingcampaignsGatewayTimeout() *PostOutboundMessagingcampaignsGatewayTimeout {
	return &PostOutboundMessagingcampaignsGatewayTimeout{}
}

/*
PostOutboundMessagingcampaignsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostOutboundMessagingcampaignsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound messagingcampaigns gateway timeout response has a 2xx status code
func (o *PostOutboundMessagingcampaignsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound messagingcampaigns gateway timeout response has a 3xx status code
func (o *PostOutboundMessagingcampaignsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound messagingcampaigns gateway timeout response has a 4xx status code
func (o *PostOutboundMessagingcampaignsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post outbound messagingcampaigns gateway timeout response has a 5xx status code
func (o *PostOutboundMessagingcampaignsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post outbound messagingcampaigns gateway timeout response a status code equal to that given
func (o *PostOutboundMessagingcampaignsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostOutboundMessagingcampaignsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/messagingcampaigns][%d] postOutboundMessagingcampaignsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostOutboundMessagingcampaignsGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/messagingcampaigns][%d] postOutboundMessagingcampaignsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostOutboundMessagingcampaignsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundMessagingcampaignsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

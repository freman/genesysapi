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

// GetOutboundSchedulesMessagingcampaignsReader is a Reader for the GetOutboundSchedulesMessagingcampaigns structure.
type GetOutboundSchedulesMessagingcampaignsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutboundSchedulesMessagingcampaignsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutboundSchedulesMessagingcampaignsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutboundSchedulesMessagingcampaignsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutboundSchedulesMessagingcampaignsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutboundSchedulesMessagingcampaignsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutboundSchedulesMessagingcampaignsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOutboundSchedulesMessagingcampaignsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOutboundSchedulesMessagingcampaignsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOutboundSchedulesMessagingcampaignsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOutboundSchedulesMessagingcampaignsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOutboundSchedulesMessagingcampaignsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOutboundSchedulesMessagingcampaignsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOutboundSchedulesMessagingcampaignsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOutboundSchedulesMessagingcampaignsOK creates a GetOutboundSchedulesMessagingcampaignsOK with default headers values
func NewGetOutboundSchedulesMessagingcampaignsOK() *GetOutboundSchedulesMessagingcampaignsOK {
	return &GetOutboundSchedulesMessagingcampaignsOK{}
}

/*
GetOutboundSchedulesMessagingcampaignsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetOutboundSchedulesMessagingcampaignsOK struct {
	Payload *models.MessagingCampaignScheduleEntityListing
}

// IsSuccess returns true when this get outbound schedules messagingcampaigns o k response has a 2xx status code
func (o *GetOutboundSchedulesMessagingcampaignsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get outbound schedules messagingcampaigns o k response has a 3xx status code
func (o *GetOutboundSchedulesMessagingcampaignsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules messagingcampaigns o k response has a 4xx status code
func (o *GetOutboundSchedulesMessagingcampaignsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound schedules messagingcampaigns o k response has a 5xx status code
func (o *GetOutboundSchedulesMessagingcampaignsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound schedules messagingcampaigns o k response a status code equal to that given
func (o *GetOutboundSchedulesMessagingcampaignsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOutboundSchedulesMessagingcampaignsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/messagingcampaigns][%d] getOutboundSchedulesMessagingcampaignsOK  %+v", 200, o.Payload)
}

func (o *GetOutboundSchedulesMessagingcampaignsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/messagingcampaigns][%d] getOutboundSchedulesMessagingcampaignsOK  %+v", 200, o.Payload)
}

func (o *GetOutboundSchedulesMessagingcampaignsOK) GetPayload() *models.MessagingCampaignScheduleEntityListing {
	return o.Payload
}

func (o *GetOutboundSchedulesMessagingcampaignsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessagingCampaignScheduleEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesMessagingcampaignsBadRequest creates a GetOutboundSchedulesMessagingcampaignsBadRequest with default headers values
func NewGetOutboundSchedulesMessagingcampaignsBadRequest() *GetOutboundSchedulesMessagingcampaignsBadRequest {
	return &GetOutboundSchedulesMessagingcampaignsBadRequest{}
}

/*
GetOutboundSchedulesMessagingcampaignsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundSchedulesMessagingcampaignsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules messagingcampaigns bad request response has a 2xx status code
func (o *GetOutboundSchedulesMessagingcampaignsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules messagingcampaigns bad request response has a 3xx status code
func (o *GetOutboundSchedulesMessagingcampaignsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules messagingcampaigns bad request response has a 4xx status code
func (o *GetOutboundSchedulesMessagingcampaignsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound schedules messagingcampaigns bad request response has a 5xx status code
func (o *GetOutboundSchedulesMessagingcampaignsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound schedules messagingcampaigns bad request response a status code equal to that given
func (o *GetOutboundSchedulesMessagingcampaignsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetOutboundSchedulesMessagingcampaignsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/messagingcampaigns][%d] getOutboundSchedulesMessagingcampaignsBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundSchedulesMessagingcampaignsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/messagingcampaigns][%d] getOutboundSchedulesMessagingcampaignsBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundSchedulesMessagingcampaignsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesMessagingcampaignsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesMessagingcampaignsUnauthorized creates a GetOutboundSchedulesMessagingcampaignsUnauthorized with default headers values
func NewGetOutboundSchedulesMessagingcampaignsUnauthorized() *GetOutboundSchedulesMessagingcampaignsUnauthorized {
	return &GetOutboundSchedulesMessagingcampaignsUnauthorized{}
}

/*
GetOutboundSchedulesMessagingcampaignsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundSchedulesMessagingcampaignsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules messagingcampaigns unauthorized response has a 2xx status code
func (o *GetOutboundSchedulesMessagingcampaignsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules messagingcampaigns unauthorized response has a 3xx status code
func (o *GetOutboundSchedulesMessagingcampaignsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules messagingcampaigns unauthorized response has a 4xx status code
func (o *GetOutboundSchedulesMessagingcampaignsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound schedules messagingcampaigns unauthorized response has a 5xx status code
func (o *GetOutboundSchedulesMessagingcampaignsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound schedules messagingcampaigns unauthorized response a status code equal to that given
func (o *GetOutboundSchedulesMessagingcampaignsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetOutboundSchedulesMessagingcampaignsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/messagingcampaigns][%d] getOutboundSchedulesMessagingcampaignsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundSchedulesMessagingcampaignsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/messagingcampaigns][%d] getOutboundSchedulesMessagingcampaignsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundSchedulesMessagingcampaignsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesMessagingcampaignsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesMessagingcampaignsForbidden creates a GetOutboundSchedulesMessagingcampaignsForbidden with default headers values
func NewGetOutboundSchedulesMessagingcampaignsForbidden() *GetOutboundSchedulesMessagingcampaignsForbidden {
	return &GetOutboundSchedulesMessagingcampaignsForbidden{}
}

/*
GetOutboundSchedulesMessagingcampaignsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundSchedulesMessagingcampaignsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules messagingcampaigns forbidden response has a 2xx status code
func (o *GetOutboundSchedulesMessagingcampaignsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules messagingcampaigns forbidden response has a 3xx status code
func (o *GetOutboundSchedulesMessagingcampaignsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules messagingcampaigns forbidden response has a 4xx status code
func (o *GetOutboundSchedulesMessagingcampaignsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound schedules messagingcampaigns forbidden response has a 5xx status code
func (o *GetOutboundSchedulesMessagingcampaignsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound schedules messagingcampaigns forbidden response a status code equal to that given
func (o *GetOutboundSchedulesMessagingcampaignsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetOutboundSchedulesMessagingcampaignsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/messagingcampaigns][%d] getOutboundSchedulesMessagingcampaignsForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundSchedulesMessagingcampaignsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/messagingcampaigns][%d] getOutboundSchedulesMessagingcampaignsForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundSchedulesMessagingcampaignsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesMessagingcampaignsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesMessagingcampaignsNotFound creates a GetOutboundSchedulesMessagingcampaignsNotFound with default headers values
func NewGetOutboundSchedulesMessagingcampaignsNotFound() *GetOutboundSchedulesMessagingcampaignsNotFound {
	return &GetOutboundSchedulesMessagingcampaignsNotFound{}
}

/*
GetOutboundSchedulesMessagingcampaignsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetOutboundSchedulesMessagingcampaignsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules messagingcampaigns not found response has a 2xx status code
func (o *GetOutboundSchedulesMessagingcampaignsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules messagingcampaigns not found response has a 3xx status code
func (o *GetOutboundSchedulesMessagingcampaignsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules messagingcampaigns not found response has a 4xx status code
func (o *GetOutboundSchedulesMessagingcampaignsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound schedules messagingcampaigns not found response has a 5xx status code
func (o *GetOutboundSchedulesMessagingcampaignsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound schedules messagingcampaigns not found response a status code equal to that given
func (o *GetOutboundSchedulesMessagingcampaignsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetOutboundSchedulesMessagingcampaignsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/messagingcampaigns][%d] getOutboundSchedulesMessagingcampaignsNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundSchedulesMessagingcampaignsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/messagingcampaigns][%d] getOutboundSchedulesMessagingcampaignsNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundSchedulesMessagingcampaignsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesMessagingcampaignsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesMessagingcampaignsRequestTimeout creates a GetOutboundSchedulesMessagingcampaignsRequestTimeout with default headers values
func NewGetOutboundSchedulesMessagingcampaignsRequestTimeout() *GetOutboundSchedulesMessagingcampaignsRequestTimeout {
	return &GetOutboundSchedulesMessagingcampaignsRequestTimeout{}
}

/*
GetOutboundSchedulesMessagingcampaignsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetOutboundSchedulesMessagingcampaignsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules messagingcampaigns request timeout response has a 2xx status code
func (o *GetOutboundSchedulesMessagingcampaignsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules messagingcampaigns request timeout response has a 3xx status code
func (o *GetOutboundSchedulesMessagingcampaignsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules messagingcampaigns request timeout response has a 4xx status code
func (o *GetOutboundSchedulesMessagingcampaignsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound schedules messagingcampaigns request timeout response has a 5xx status code
func (o *GetOutboundSchedulesMessagingcampaignsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound schedules messagingcampaigns request timeout response a status code equal to that given
func (o *GetOutboundSchedulesMessagingcampaignsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetOutboundSchedulesMessagingcampaignsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/messagingcampaigns][%d] getOutboundSchedulesMessagingcampaignsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundSchedulesMessagingcampaignsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/messagingcampaigns][%d] getOutboundSchedulesMessagingcampaignsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundSchedulesMessagingcampaignsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesMessagingcampaignsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesMessagingcampaignsRequestEntityTooLarge creates a GetOutboundSchedulesMessagingcampaignsRequestEntityTooLarge with default headers values
func NewGetOutboundSchedulesMessagingcampaignsRequestEntityTooLarge() *GetOutboundSchedulesMessagingcampaignsRequestEntityTooLarge {
	return &GetOutboundSchedulesMessagingcampaignsRequestEntityTooLarge{}
}

/*
GetOutboundSchedulesMessagingcampaignsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetOutboundSchedulesMessagingcampaignsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules messagingcampaigns request entity too large response has a 2xx status code
func (o *GetOutboundSchedulesMessagingcampaignsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules messagingcampaigns request entity too large response has a 3xx status code
func (o *GetOutboundSchedulesMessagingcampaignsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules messagingcampaigns request entity too large response has a 4xx status code
func (o *GetOutboundSchedulesMessagingcampaignsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound schedules messagingcampaigns request entity too large response has a 5xx status code
func (o *GetOutboundSchedulesMessagingcampaignsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound schedules messagingcampaigns request entity too large response a status code equal to that given
func (o *GetOutboundSchedulesMessagingcampaignsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetOutboundSchedulesMessagingcampaignsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/messagingcampaigns][%d] getOutboundSchedulesMessagingcampaignsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundSchedulesMessagingcampaignsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/messagingcampaigns][%d] getOutboundSchedulesMessagingcampaignsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundSchedulesMessagingcampaignsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesMessagingcampaignsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesMessagingcampaignsUnsupportedMediaType creates a GetOutboundSchedulesMessagingcampaignsUnsupportedMediaType with default headers values
func NewGetOutboundSchedulesMessagingcampaignsUnsupportedMediaType() *GetOutboundSchedulesMessagingcampaignsUnsupportedMediaType {
	return &GetOutboundSchedulesMessagingcampaignsUnsupportedMediaType{}
}

/*
GetOutboundSchedulesMessagingcampaignsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundSchedulesMessagingcampaignsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules messagingcampaigns unsupported media type response has a 2xx status code
func (o *GetOutboundSchedulesMessagingcampaignsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules messagingcampaigns unsupported media type response has a 3xx status code
func (o *GetOutboundSchedulesMessagingcampaignsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules messagingcampaigns unsupported media type response has a 4xx status code
func (o *GetOutboundSchedulesMessagingcampaignsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound schedules messagingcampaigns unsupported media type response has a 5xx status code
func (o *GetOutboundSchedulesMessagingcampaignsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound schedules messagingcampaigns unsupported media type response a status code equal to that given
func (o *GetOutboundSchedulesMessagingcampaignsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetOutboundSchedulesMessagingcampaignsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/messagingcampaigns][%d] getOutboundSchedulesMessagingcampaignsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundSchedulesMessagingcampaignsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/messagingcampaigns][%d] getOutboundSchedulesMessagingcampaignsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundSchedulesMessagingcampaignsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesMessagingcampaignsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesMessagingcampaignsTooManyRequests creates a GetOutboundSchedulesMessagingcampaignsTooManyRequests with default headers values
func NewGetOutboundSchedulesMessagingcampaignsTooManyRequests() *GetOutboundSchedulesMessagingcampaignsTooManyRequests {
	return &GetOutboundSchedulesMessagingcampaignsTooManyRequests{}
}

/*
GetOutboundSchedulesMessagingcampaignsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetOutboundSchedulesMessagingcampaignsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules messagingcampaigns too many requests response has a 2xx status code
func (o *GetOutboundSchedulesMessagingcampaignsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules messagingcampaigns too many requests response has a 3xx status code
func (o *GetOutboundSchedulesMessagingcampaignsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules messagingcampaigns too many requests response has a 4xx status code
func (o *GetOutboundSchedulesMessagingcampaignsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound schedules messagingcampaigns too many requests response has a 5xx status code
func (o *GetOutboundSchedulesMessagingcampaignsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound schedules messagingcampaigns too many requests response a status code equal to that given
func (o *GetOutboundSchedulesMessagingcampaignsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetOutboundSchedulesMessagingcampaignsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/messagingcampaigns][%d] getOutboundSchedulesMessagingcampaignsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundSchedulesMessagingcampaignsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/messagingcampaigns][%d] getOutboundSchedulesMessagingcampaignsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundSchedulesMessagingcampaignsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesMessagingcampaignsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesMessagingcampaignsInternalServerError creates a GetOutboundSchedulesMessagingcampaignsInternalServerError with default headers values
func NewGetOutboundSchedulesMessagingcampaignsInternalServerError() *GetOutboundSchedulesMessagingcampaignsInternalServerError {
	return &GetOutboundSchedulesMessagingcampaignsInternalServerError{}
}

/*
GetOutboundSchedulesMessagingcampaignsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundSchedulesMessagingcampaignsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules messagingcampaigns internal server error response has a 2xx status code
func (o *GetOutboundSchedulesMessagingcampaignsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules messagingcampaigns internal server error response has a 3xx status code
func (o *GetOutboundSchedulesMessagingcampaignsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules messagingcampaigns internal server error response has a 4xx status code
func (o *GetOutboundSchedulesMessagingcampaignsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound schedules messagingcampaigns internal server error response has a 5xx status code
func (o *GetOutboundSchedulesMessagingcampaignsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound schedules messagingcampaigns internal server error response a status code equal to that given
func (o *GetOutboundSchedulesMessagingcampaignsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetOutboundSchedulesMessagingcampaignsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/messagingcampaigns][%d] getOutboundSchedulesMessagingcampaignsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundSchedulesMessagingcampaignsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/messagingcampaigns][%d] getOutboundSchedulesMessagingcampaignsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundSchedulesMessagingcampaignsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesMessagingcampaignsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesMessagingcampaignsServiceUnavailable creates a GetOutboundSchedulesMessagingcampaignsServiceUnavailable with default headers values
func NewGetOutboundSchedulesMessagingcampaignsServiceUnavailable() *GetOutboundSchedulesMessagingcampaignsServiceUnavailable {
	return &GetOutboundSchedulesMessagingcampaignsServiceUnavailable{}
}

/*
GetOutboundSchedulesMessagingcampaignsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundSchedulesMessagingcampaignsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules messagingcampaigns service unavailable response has a 2xx status code
func (o *GetOutboundSchedulesMessagingcampaignsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules messagingcampaigns service unavailable response has a 3xx status code
func (o *GetOutboundSchedulesMessagingcampaignsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules messagingcampaigns service unavailable response has a 4xx status code
func (o *GetOutboundSchedulesMessagingcampaignsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound schedules messagingcampaigns service unavailable response has a 5xx status code
func (o *GetOutboundSchedulesMessagingcampaignsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound schedules messagingcampaigns service unavailable response a status code equal to that given
func (o *GetOutboundSchedulesMessagingcampaignsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetOutboundSchedulesMessagingcampaignsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/messagingcampaigns][%d] getOutboundSchedulesMessagingcampaignsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundSchedulesMessagingcampaignsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/messagingcampaigns][%d] getOutboundSchedulesMessagingcampaignsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundSchedulesMessagingcampaignsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesMessagingcampaignsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesMessagingcampaignsGatewayTimeout creates a GetOutboundSchedulesMessagingcampaignsGatewayTimeout with default headers values
func NewGetOutboundSchedulesMessagingcampaignsGatewayTimeout() *GetOutboundSchedulesMessagingcampaignsGatewayTimeout {
	return &GetOutboundSchedulesMessagingcampaignsGatewayTimeout{}
}

/*
GetOutboundSchedulesMessagingcampaignsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetOutboundSchedulesMessagingcampaignsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules messagingcampaigns gateway timeout response has a 2xx status code
func (o *GetOutboundSchedulesMessagingcampaignsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules messagingcampaigns gateway timeout response has a 3xx status code
func (o *GetOutboundSchedulesMessagingcampaignsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules messagingcampaigns gateway timeout response has a 4xx status code
func (o *GetOutboundSchedulesMessagingcampaignsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound schedules messagingcampaigns gateway timeout response has a 5xx status code
func (o *GetOutboundSchedulesMessagingcampaignsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound schedules messagingcampaigns gateway timeout response a status code equal to that given
func (o *GetOutboundSchedulesMessagingcampaignsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetOutboundSchedulesMessagingcampaignsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/messagingcampaigns][%d] getOutboundSchedulesMessagingcampaignsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundSchedulesMessagingcampaignsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/messagingcampaigns][%d] getOutboundSchedulesMessagingcampaignsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundSchedulesMessagingcampaignsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesMessagingcampaignsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

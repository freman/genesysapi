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

// PutOutboundCampaignAgentReader is a Reader for the PutOutboundCampaignAgent structure.
type PutOutboundCampaignAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutOutboundCampaignAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutOutboundCampaignAgentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutOutboundCampaignAgentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutOutboundCampaignAgentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutOutboundCampaignAgentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutOutboundCampaignAgentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutOutboundCampaignAgentRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutOutboundCampaignAgentRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutOutboundCampaignAgentUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutOutboundCampaignAgentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutOutboundCampaignAgentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutOutboundCampaignAgentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutOutboundCampaignAgentGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutOutboundCampaignAgentOK creates a PutOutboundCampaignAgentOK with default headers values
func NewPutOutboundCampaignAgentOK() *PutOutboundCampaignAgentOK {
	return &PutOutboundCampaignAgentOK{}
}

/*
PutOutboundCampaignAgentOK describes a response with status code 200, with default header values.

successful operation
*/
type PutOutboundCampaignAgentOK struct {
	Payload string
}

// IsSuccess returns true when this put outbound campaign agent o k response has a 2xx status code
func (o *PutOutboundCampaignAgentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put outbound campaign agent o k response has a 3xx status code
func (o *PutOutboundCampaignAgentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound campaign agent o k response has a 4xx status code
func (o *PutOutboundCampaignAgentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put outbound campaign agent o k response has a 5xx status code
func (o *PutOutboundCampaignAgentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put outbound campaign agent o k response a status code equal to that given
func (o *PutOutboundCampaignAgentOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutOutboundCampaignAgentOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/campaigns/{campaignId}/agents/{userId}][%d] putOutboundCampaignAgentOK  %+v", 200, o.Payload)
}

func (o *PutOutboundCampaignAgentOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/campaigns/{campaignId}/agents/{userId}][%d] putOutboundCampaignAgentOK  %+v", 200, o.Payload)
}

func (o *PutOutboundCampaignAgentOK) GetPayload() string {
	return o.Payload
}

func (o *PutOutboundCampaignAgentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundCampaignAgentBadRequest creates a PutOutboundCampaignAgentBadRequest with default headers values
func NewPutOutboundCampaignAgentBadRequest() *PutOutboundCampaignAgentBadRequest {
	return &PutOutboundCampaignAgentBadRequest{}
}

/*
PutOutboundCampaignAgentBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutOutboundCampaignAgentBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound campaign agent bad request response has a 2xx status code
func (o *PutOutboundCampaignAgentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound campaign agent bad request response has a 3xx status code
func (o *PutOutboundCampaignAgentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound campaign agent bad request response has a 4xx status code
func (o *PutOutboundCampaignAgentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put outbound campaign agent bad request response has a 5xx status code
func (o *PutOutboundCampaignAgentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put outbound campaign agent bad request response a status code equal to that given
func (o *PutOutboundCampaignAgentBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutOutboundCampaignAgentBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/campaigns/{campaignId}/agents/{userId}][%d] putOutboundCampaignAgentBadRequest  %+v", 400, o.Payload)
}

func (o *PutOutboundCampaignAgentBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/campaigns/{campaignId}/agents/{userId}][%d] putOutboundCampaignAgentBadRequest  %+v", 400, o.Payload)
}

func (o *PutOutboundCampaignAgentBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundCampaignAgentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundCampaignAgentUnauthorized creates a PutOutboundCampaignAgentUnauthorized with default headers values
func NewPutOutboundCampaignAgentUnauthorized() *PutOutboundCampaignAgentUnauthorized {
	return &PutOutboundCampaignAgentUnauthorized{}
}

/*
PutOutboundCampaignAgentUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PutOutboundCampaignAgentUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound campaign agent unauthorized response has a 2xx status code
func (o *PutOutboundCampaignAgentUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound campaign agent unauthorized response has a 3xx status code
func (o *PutOutboundCampaignAgentUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound campaign agent unauthorized response has a 4xx status code
func (o *PutOutboundCampaignAgentUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put outbound campaign agent unauthorized response has a 5xx status code
func (o *PutOutboundCampaignAgentUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put outbound campaign agent unauthorized response a status code equal to that given
func (o *PutOutboundCampaignAgentUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PutOutboundCampaignAgentUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/campaigns/{campaignId}/agents/{userId}][%d] putOutboundCampaignAgentUnauthorized  %+v", 401, o.Payload)
}

func (o *PutOutboundCampaignAgentUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/campaigns/{campaignId}/agents/{userId}][%d] putOutboundCampaignAgentUnauthorized  %+v", 401, o.Payload)
}

func (o *PutOutboundCampaignAgentUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundCampaignAgentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundCampaignAgentForbidden creates a PutOutboundCampaignAgentForbidden with default headers values
func NewPutOutboundCampaignAgentForbidden() *PutOutboundCampaignAgentForbidden {
	return &PutOutboundCampaignAgentForbidden{}
}

/*
PutOutboundCampaignAgentForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PutOutboundCampaignAgentForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound campaign agent forbidden response has a 2xx status code
func (o *PutOutboundCampaignAgentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound campaign agent forbidden response has a 3xx status code
func (o *PutOutboundCampaignAgentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound campaign agent forbidden response has a 4xx status code
func (o *PutOutboundCampaignAgentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put outbound campaign agent forbidden response has a 5xx status code
func (o *PutOutboundCampaignAgentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put outbound campaign agent forbidden response a status code equal to that given
func (o *PutOutboundCampaignAgentForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutOutboundCampaignAgentForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/campaigns/{campaignId}/agents/{userId}][%d] putOutboundCampaignAgentForbidden  %+v", 403, o.Payload)
}

func (o *PutOutboundCampaignAgentForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/campaigns/{campaignId}/agents/{userId}][%d] putOutboundCampaignAgentForbidden  %+v", 403, o.Payload)
}

func (o *PutOutboundCampaignAgentForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundCampaignAgentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundCampaignAgentNotFound creates a PutOutboundCampaignAgentNotFound with default headers values
func NewPutOutboundCampaignAgentNotFound() *PutOutboundCampaignAgentNotFound {
	return &PutOutboundCampaignAgentNotFound{}
}

/*
PutOutboundCampaignAgentNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PutOutboundCampaignAgentNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound campaign agent not found response has a 2xx status code
func (o *PutOutboundCampaignAgentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound campaign agent not found response has a 3xx status code
func (o *PutOutboundCampaignAgentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound campaign agent not found response has a 4xx status code
func (o *PutOutboundCampaignAgentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put outbound campaign agent not found response has a 5xx status code
func (o *PutOutboundCampaignAgentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put outbound campaign agent not found response a status code equal to that given
func (o *PutOutboundCampaignAgentNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutOutboundCampaignAgentNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/campaigns/{campaignId}/agents/{userId}][%d] putOutboundCampaignAgentNotFound  %+v", 404, o.Payload)
}

func (o *PutOutboundCampaignAgentNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/campaigns/{campaignId}/agents/{userId}][%d] putOutboundCampaignAgentNotFound  %+v", 404, o.Payload)
}

func (o *PutOutboundCampaignAgentNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundCampaignAgentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundCampaignAgentRequestTimeout creates a PutOutboundCampaignAgentRequestTimeout with default headers values
func NewPutOutboundCampaignAgentRequestTimeout() *PutOutboundCampaignAgentRequestTimeout {
	return &PutOutboundCampaignAgentRequestTimeout{}
}

/*
PutOutboundCampaignAgentRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutOutboundCampaignAgentRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound campaign agent request timeout response has a 2xx status code
func (o *PutOutboundCampaignAgentRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound campaign agent request timeout response has a 3xx status code
func (o *PutOutboundCampaignAgentRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound campaign agent request timeout response has a 4xx status code
func (o *PutOutboundCampaignAgentRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this put outbound campaign agent request timeout response has a 5xx status code
func (o *PutOutboundCampaignAgentRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this put outbound campaign agent request timeout response a status code equal to that given
func (o *PutOutboundCampaignAgentRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PutOutboundCampaignAgentRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/campaigns/{campaignId}/agents/{userId}][%d] putOutboundCampaignAgentRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutOutboundCampaignAgentRequestTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/campaigns/{campaignId}/agents/{userId}][%d] putOutboundCampaignAgentRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutOutboundCampaignAgentRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundCampaignAgentRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundCampaignAgentRequestEntityTooLarge creates a PutOutboundCampaignAgentRequestEntityTooLarge with default headers values
func NewPutOutboundCampaignAgentRequestEntityTooLarge() *PutOutboundCampaignAgentRequestEntityTooLarge {
	return &PutOutboundCampaignAgentRequestEntityTooLarge{}
}

/*
PutOutboundCampaignAgentRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PutOutboundCampaignAgentRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound campaign agent request entity too large response has a 2xx status code
func (o *PutOutboundCampaignAgentRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound campaign agent request entity too large response has a 3xx status code
func (o *PutOutboundCampaignAgentRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound campaign agent request entity too large response has a 4xx status code
func (o *PutOutboundCampaignAgentRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this put outbound campaign agent request entity too large response has a 5xx status code
func (o *PutOutboundCampaignAgentRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this put outbound campaign agent request entity too large response a status code equal to that given
func (o *PutOutboundCampaignAgentRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PutOutboundCampaignAgentRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/campaigns/{campaignId}/agents/{userId}][%d] putOutboundCampaignAgentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutOutboundCampaignAgentRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/campaigns/{campaignId}/agents/{userId}][%d] putOutboundCampaignAgentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutOutboundCampaignAgentRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundCampaignAgentRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundCampaignAgentUnsupportedMediaType creates a PutOutboundCampaignAgentUnsupportedMediaType with default headers values
func NewPutOutboundCampaignAgentUnsupportedMediaType() *PutOutboundCampaignAgentUnsupportedMediaType {
	return &PutOutboundCampaignAgentUnsupportedMediaType{}
}

/*
PutOutboundCampaignAgentUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutOutboundCampaignAgentUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound campaign agent unsupported media type response has a 2xx status code
func (o *PutOutboundCampaignAgentUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound campaign agent unsupported media type response has a 3xx status code
func (o *PutOutboundCampaignAgentUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound campaign agent unsupported media type response has a 4xx status code
func (o *PutOutboundCampaignAgentUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this put outbound campaign agent unsupported media type response has a 5xx status code
func (o *PutOutboundCampaignAgentUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this put outbound campaign agent unsupported media type response a status code equal to that given
func (o *PutOutboundCampaignAgentUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PutOutboundCampaignAgentUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/campaigns/{campaignId}/agents/{userId}][%d] putOutboundCampaignAgentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutOutboundCampaignAgentUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/campaigns/{campaignId}/agents/{userId}][%d] putOutboundCampaignAgentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutOutboundCampaignAgentUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundCampaignAgentUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundCampaignAgentTooManyRequests creates a PutOutboundCampaignAgentTooManyRequests with default headers values
func NewPutOutboundCampaignAgentTooManyRequests() *PutOutboundCampaignAgentTooManyRequests {
	return &PutOutboundCampaignAgentTooManyRequests{}
}

/*
PutOutboundCampaignAgentTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutOutboundCampaignAgentTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound campaign agent too many requests response has a 2xx status code
func (o *PutOutboundCampaignAgentTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound campaign agent too many requests response has a 3xx status code
func (o *PutOutboundCampaignAgentTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound campaign agent too many requests response has a 4xx status code
func (o *PutOutboundCampaignAgentTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this put outbound campaign agent too many requests response has a 5xx status code
func (o *PutOutboundCampaignAgentTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this put outbound campaign agent too many requests response a status code equal to that given
func (o *PutOutboundCampaignAgentTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PutOutboundCampaignAgentTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/campaigns/{campaignId}/agents/{userId}][%d] putOutboundCampaignAgentTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutOutboundCampaignAgentTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/campaigns/{campaignId}/agents/{userId}][%d] putOutboundCampaignAgentTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutOutboundCampaignAgentTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundCampaignAgentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundCampaignAgentInternalServerError creates a PutOutboundCampaignAgentInternalServerError with default headers values
func NewPutOutboundCampaignAgentInternalServerError() *PutOutboundCampaignAgentInternalServerError {
	return &PutOutboundCampaignAgentInternalServerError{}
}

/*
PutOutboundCampaignAgentInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutOutboundCampaignAgentInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound campaign agent internal server error response has a 2xx status code
func (o *PutOutboundCampaignAgentInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound campaign agent internal server error response has a 3xx status code
func (o *PutOutboundCampaignAgentInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound campaign agent internal server error response has a 4xx status code
func (o *PutOutboundCampaignAgentInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put outbound campaign agent internal server error response has a 5xx status code
func (o *PutOutboundCampaignAgentInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put outbound campaign agent internal server error response a status code equal to that given
func (o *PutOutboundCampaignAgentInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutOutboundCampaignAgentInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/campaigns/{campaignId}/agents/{userId}][%d] putOutboundCampaignAgentInternalServerError  %+v", 500, o.Payload)
}

func (o *PutOutboundCampaignAgentInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/campaigns/{campaignId}/agents/{userId}][%d] putOutboundCampaignAgentInternalServerError  %+v", 500, o.Payload)
}

func (o *PutOutboundCampaignAgentInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundCampaignAgentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundCampaignAgentServiceUnavailable creates a PutOutboundCampaignAgentServiceUnavailable with default headers values
func NewPutOutboundCampaignAgentServiceUnavailable() *PutOutboundCampaignAgentServiceUnavailable {
	return &PutOutboundCampaignAgentServiceUnavailable{}
}

/*
PutOutboundCampaignAgentServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutOutboundCampaignAgentServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound campaign agent service unavailable response has a 2xx status code
func (o *PutOutboundCampaignAgentServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound campaign agent service unavailable response has a 3xx status code
func (o *PutOutboundCampaignAgentServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound campaign agent service unavailable response has a 4xx status code
func (o *PutOutboundCampaignAgentServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this put outbound campaign agent service unavailable response has a 5xx status code
func (o *PutOutboundCampaignAgentServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this put outbound campaign agent service unavailable response a status code equal to that given
func (o *PutOutboundCampaignAgentServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PutOutboundCampaignAgentServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/campaigns/{campaignId}/agents/{userId}][%d] putOutboundCampaignAgentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutOutboundCampaignAgentServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/campaigns/{campaignId}/agents/{userId}][%d] putOutboundCampaignAgentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutOutboundCampaignAgentServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundCampaignAgentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundCampaignAgentGatewayTimeout creates a PutOutboundCampaignAgentGatewayTimeout with default headers values
func NewPutOutboundCampaignAgentGatewayTimeout() *PutOutboundCampaignAgentGatewayTimeout {
	return &PutOutboundCampaignAgentGatewayTimeout{}
}

/*
PutOutboundCampaignAgentGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PutOutboundCampaignAgentGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound campaign agent gateway timeout response has a 2xx status code
func (o *PutOutboundCampaignAgentGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound campaign agent gateway timeout response has a 3xx status code
func (o *PutOutboundCampaignAgentGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound campaign agent gateway timeout response has a 4xx status code
func (o *PutOutboundCampaignAgentGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this put outbound campaign agent gateway timeout response has a 5xx status code
func (o *PutOutboundCampaignAgentGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this put outbound campaign agent gateway timeout response a status code equal to that given
func (o *PutOutboundCampaignAgentGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PutOutboundCampaignAgentGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/campaigns/{campaignId}/agents/{userId}][%d] putOutboundCampaignAgentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutOutboundCampaignAgentGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/campaigns/{campaignId}/agents/{userId}][%d] putOutboundCampaignAgentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutOutboundCampaignAgentGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundCampaignAgentGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package architect

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetFlowsMilestonesReader is a Reader for the GetFlowsMilestones structure.
type GetFlowsMilestonesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFlowsMilestonesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFlowsMilestonesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFlowsMilestonesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFlowsMilestonesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFlowsMilestonesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFlowsMilestonesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetFlowsMilestonesMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetFlowsMilestonesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetFlowsMilestonesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetFlowsMilestonesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetFlowsMilestonesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFlowsMilestonesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetFlowsMilestonesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetFlowsMilestonesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFlowsMilestonesOK creates a GetFlowsMilestonesOK with default headers values
func NewGetFlowsMilestonesOK() *GetFlowsMilestonesOK {
	return &GetFlowsMilestonesOK{}
}

/*
GetFlowsMilestonesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetFlowsMilestonesOK struct {
	Payload *models.FlowMilestoneListing
}

// IsSuccess returns true when this get flows milestones o k response has a 2xx status code
func (o *GetFlowsMilestonesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get flows milestones o k response has a 3xx status code
func (o *GetFlowsMilestonesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows milestones o k response has a 4xx status code
func (o *GetFlowsMilestonesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get flows milestones o k response has a 5xx status code
func (o *GetFlowsMilestonesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get flows milestones o k response a status code equal to that given
func (o *GetFlowsMilestonesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetFlowsMilestonesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/milestones][%d] getFlowsMilestonesOK  %+v", 200, o.Payload)
}

func (o *GetFlowsMilestonesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/milestones][%d] getFlowsMilestonesOK  %+v", 200, o.Payload)
}

func (o *GetFlowsMilestonesOK) GetPayload() *models.FlowMilestoneListing {
	return o.Payload
}

func (o *GetFlowsMilestonesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FlowMilestoneListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsMilestonesBadRequest creates a GetFlowsMilestonesBadRequest with default headers values
func NewGetFlowsMilestonesBadRequest() *GetFlowsMilestonesBadRequest {
	return &GetFlowsMilestonesBadRequest{}
}

/*
GetFlowsMilestonesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetFlowsMilestonesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows milestones bad request response has a 2xx status code
func (o *GetFlowsMilestonesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows milestones bad request response has a 3xx status code
func (o *GetFlowsMilestonesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows milestones bad request response has a 4xx status code
func (o *GetFlowsMilestonesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flows milestones bad request response has a 5xx status code
func (o *GetFlowsMilestonesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get flows milestones bad request response a status code equal to that given
func (o *GetFlowsMilestonesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetFlowsMilestonesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/milestones][%d] getFlowsMilestonesBadRequest  %+v", 400, o.Payload)
}

func (o *GetFlowsMilestonesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/milestones][%d] getFlowsMilestonesBadRequest  %+v", 400, o.Payload)
}

func (o *GetFlowsMilestonesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsMilestonesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsMilestonesUnauthorized creates a GetFlowsMilestonesUnauthorized with default headers values
func NewGetFlowsMilestonesUnauthorized() *GetFlowsMilestonesUnauthorized {
	return &GetFlowsMilestonesUnauthorized{}
}

/*
GetFlowsMilestonesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetFlowsMilestonesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows milestones unauthorized response has a 2xx status code
func (o *GetFlowsMilestonesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows milestones unauthorized response has a 3xx status code
func (o *GetFlowsMilestonesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows milestones unauthorized response has a 4xx status code
func (o *GetFlowsMilestonesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flows milestones unauthorized response has a 5xx status code
func (o *GetFlowsMilestonesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get flows milestones unauthorized response a status code equal to that given
func (o *GetFlowsMilestonesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetFlowsMilestonesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/milestones][%d] getFlowsMilestonesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFlowsMilestonesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/milestones][%d] getFlowsMilestonesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFlowsMilestonesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsMilestonesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsMilestonesForbidden creates a GetFlowsMilestonesForbidden with default headers values
func NewGetFlowsMilestonesForbidden() *GetFlowsMilestonesForbidden {
	return &GetFlowsMilestonesForbidden{}
}

/*
GetFlowsMilestonesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetFlowsMilestonesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows milestones forbidden response has a 2xx status code
func (o *GetFlowsMilestonesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows milestones forbidden response has a 3xx status code
func (o *GetFlowsMilestonesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows milestones forbidden response has a 4xx status code
func (o *GetFlowsMilestonesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flows milestones forbidden response has a 5xx status code
func (o *GetFlowsMilestonesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get flows milestones forbidden response a status code equal to that given
func (o *GetFlowsMilestonesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetFlowsMilestonesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/milestones][%d] getFlowsMilestonesForbidden  %+v", 403, o.Payload)
}

func (o *GetFlowsMilestonesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/milestones][%d] getFlowsMilestonesForbidden  %+v", 403, o.Payload)
}

func (o *GetFlowsMilestonesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsMilestonesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsMilestonesNotFound creates a GetFlowsMilestonesNotFound with default headers values
func NewGetFlowsMilestonesNotFound() *GetFlowsMilestonesNotFound {
	return &GetFlowsMilestonesNotFound{}
}

/*
GetFlowsMilestonesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetFlowsMilestonesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows milestones not found response has a 2xx status code
func (o *GetFlowsMilestonesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows milestones not found response has a 3xx status code
func (o *GetFlowsMilestonesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows milestones not found response has a 4xx status code
func (o *GetFlowsMilestonesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flows milestones not found response has a 5xx status code
func (o *GetFlowsMilestonesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get flows milestones not found response a status code equal to that given
func (o *GetFlowsMilestonesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetFlowsMilestonesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/milestones][%d] getFlowsMilestonesNotFound  %+v", 404, o.Payload)
}

func (o *GetFlowsMilestonesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/milestones][%d] getFlowsMilestonesNotFound  %+v", 404, o.Payload)
}

func (o *GetFlowsMilestonesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsMilestonesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsMilestonesMethodNotAllowed creates a GetFlowsMilestonesMethodNotAllowed with default headers values
func NewGetFlowsMilestonesMethodNotAllowed() *GetFlowsMilestonesMethodNotAllowed {
	return &GetFlowsMilestonesMethodNotAllowed{}
}

/*
GetFlowsMilestonesMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type GetFlowsMilestonesMethodNotAllowed struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows milestones method not allowed response has a 2xx status code
func (o *GetFlowsMilestonesMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows milestones method not allowed response has a 3xx status code
func (o *GetFlowsMilestonesMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows milestones method not allowed response has a 4xx status code
func (o *GetFlowsMilestonesMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flows milestones method not allowed response has a 5xx status code
func (o *GetFlowsMilestonesMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this get flows milestones method not allowed response a status code equal to that given
func (o *GetFlowsMilestonesMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

func (o *GetFlowsMilestonesMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/milestones][%d] getFlowsMilestonesMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GetFlowsMilestonesMethodNotAllowed) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/milestones][%d] getFlowsMilestonesMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GetFlowsMilestonesMethodNotAllowed) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsMilestonesMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsMilestonesRequestTimeout creates a GetFlowsMilestonesRequestTimeout with default headers values
func NewGetFlowsMilestonesRequestTimeout() *GetFlowsMilestonesRequestTimeout {
	return &GetFlowsMilestonesRequestTimeout{}
}

/*
GetFlowsMilestonesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetFlowsMilestonesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows milestones request timeout response has a 2xx status code
func (o *GetFlowsMilestonesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows milestones request timeout response has a 3xx status code
func (o *GetFlowsMilestonesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows milestones request timeout response has a 4xx status code
func (o *GetFlowsMilestonesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flows milestones request timeout response has a 5xx status code
func (o *GetFlowsMilestonesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get flows milestones request timeout response a status code equal to that given
func (o *GetFlowsMilestonesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetFlowsMilestonesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/milestones][%d] getFlowsMilestonesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetFlowsMilestonesRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/milestones][%d] getFlowsMilestonesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetFlowsMilestonesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsMilestonesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsMilestonesRequestEntityTooLarge creates a GetFlowsMilestonesRequestEntityTooLarge with default headers values
func NewGetFlowsMilestonesRequestEntityTooLarge() *GetFlowsMilestonesRequestEntityTooLarge {
	return &GetFlowsMilestonesRequestEntityTooLarge{}
}

/*
GetFlowsMilestonesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetFlowsMilestonesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows milestones request entity too large response has a 2xx status code
func (o *GetFlowsMilestonesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows milestones request entity too large response has a 3xx status code
func (o *GetFlowsMilestonesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows milestones request entity too large response has a 4xx status code
func (o *GetFlowsMilestonesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flows milestones request entity too large response has a 5xx status code
func (o *GetFlowsMilestonesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get flows milestones request entity too large response a status code equal to that given
func (o *GetFlowsMilestonesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetFlowsMilestonesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/milestones][%d] getFlowsMilestonesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetFlowsMilestonesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/milestones][%d] getFlowsMilestonesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetFlowsMilestonesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsMilestonesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsMilestonesUnsupportedMediaType creates a GetFlowsMilestonesUnsupportedMediaType with default headers values
func NewGetFlowsMilestonesUnsupportedMediaType() *GetFlowsMilestonesUnsupportedMediaType {
	return &GetFlowsMilestonesUnsupportedMediaType{}
}

/*
GetFlowsMilestonesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetFlowsMilestonesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows milestones unsupported media type response has a 2xx status code
func (o *GetFlowsMilestonesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows milestones unsupported media type response has a 3xx status code
func (o *GetFlowsMilestonesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows milestones unsupported media type response has a 4xx status code
func (o *GetFlowsMilestonesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flows milestones unsupported media type response has a 5xx status code
func (o *GetFlowsMilestonesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get flows milestones unsupported media type response a status code equal to that given
func (o *GetFlowsMilestonesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetFlowsMilestonesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/milestones][%d] getFlowsMilestonesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetFlowsMilestonesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/milestones][%d] getFlowsMilestonesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetFlowsMilestonesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsMilestonesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsMilestonesTooManyRequests creates a GetFlowsMilestonesTooManyRequests with default headers values
func NewGetFlowsMilestonesTooManyRequests() *GetFlowsMilestonesTooManyRequests {
	return &GetFlowsMilestonesTooManyRequests{}
}

/*
GetFlowsMilestonesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetFlowsMilestonesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows milestones too many requests response has a 2xx status code
func (o *GetFlowsMilestonesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows milestones too many requests response has a 3xx status code
func (o *GetFlowsMilestonesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows milestones too many requests response has a 4xx status code
func (o *GetFlowsMilestonesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flows milestones too many requests response has a 5xx status code
func (o *GetFlowsMilestonesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get flows milestones too many requests response a status code equal to that given
func (o *GetFlowsMilestonesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetFlowsMilestonesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/milestones][%d] getFlowsMilestonesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetFlowsMilestonesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/milestones][%d] getFlowsMilestonesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetFlowsMilestonesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsMilestonesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsMilestonesInternalServerError creates a GetFlowsMilestonesInternalServerError with default headers values
func NewGetFlowsMilestonesInternalServerError() *GetFlowsMilestonesInternalServerError {
	return &GetFlowsMilestonesInternalServerError{}
}

/*
GetFlowsMilestonesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetFlowsMilestonesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows milestones internal server error response has a 2xx status code
func (o *GetFlowsMilestonesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows milestones internal server error response has a 3xx status code
func (o *GetFlowsMilestonesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows milestones internal server error response has a 4xx status code
func (o *GetFlowsMilestonesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get flows milestones internal server error response has a 5xx status code
func (o *GetFlowsMilestonesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get flows milestones internal server error response a status code equal to that given
func (o *GetFlowsMilestonesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetFlowsMilestonesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/milestones][%d] getFlowsMilestonesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFlowsMilestonesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/milestones][%d] getFlowsMilestonesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFlowsMilestonesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsMilestonesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsMilestonesServiceUnavailable creates a GetFlowsMilestonesServiceUnavailable with default headers values
func NewGetFlowsMilestonesServiceUnavailable() *GetFlowsMilestonesServiceUnavailable {
	return &GetFlowsMilestonesServiceUnavailable{}
}

/*
GetFlowsMilestonesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetFlowsMilestonesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows milestones service unavailable response has a 2xx status code
func (o *GetFlowsMilestonesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows milestones service unavailable response has a 3xx status code
func (o *GetFlowsMilestonesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows milestones service unavailable response has a 4xx status code
func (o *GetFlowsMilestonesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get flows milestones service unavailable response has a 5xx status code
func (o *GetFlowsMilestonesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get flows milestones service unavailable response a status code equal to that given
func (o *GetFlowsMilestonesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetFlowsMilestonesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/milestones][%d] getFlowsMilestonesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFlowsMilestonesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/milestones][%d] getFlowsMilestonesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFlowsMilestonesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsMilestonesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsMilestonesGatewayTimeout creates a GetFlowsMilestonesGatewayTimeout with default headers values
func NewGetFlowsMilestonesGatewayTimeout() *GetFlowsMilestonesGatewayTimeout {
	return &GetFlowsMilestonesGatewayTimeout{}
}

/*
GetFlowsMilestonesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetFlowsMilestonesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows milestones gateway timeout response has a 2xx status code
func (o *GetFlowsMilestonesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows milestones gateway timeout response has a 3xx status code
func (o *GetFlowsMilestonesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows milestones gateway timeout response has a 4xx status code
func (o *GetFlowsMilestonesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get flows milestones gateway timeout response has a 5xx status code
func (o *GetFlowsMilestonesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get flows milestones gateway timeout response a status code equal to that given
func (o *GetFlowsMilestonesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetFlowsMilestonesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/milestones][%d] getFlowsMilestonesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetFlowsMilestonesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/milestones][%d] getFlowsMilestonesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetFlowsMilestonesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsMilestonesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

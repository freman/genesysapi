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

// PostFlowsMilestonesReader is a Reader for the PostFlowsMilestones structure.
type PostFlowsMilestonesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostFlowsMilestonesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostFlowsMilestonesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostFlowsMilestonesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostFlowsMilestonesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostFlowsMilestonesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostFlowsMilestonesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPostFlowsMilestonesMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostFlowsMilestonesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostFlowsMilestonesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostFlowsMilestonesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostFlowsMilestonesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostFlowsMilestonesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostFlowsMilestonesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostFlowsMilestonesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostFlowsMilestonesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostFlowsMilestonesOK creates a PostFlowsMilestonesOK with default headers values
func NewPostFlowsMilestonesOK() *PostFlowsMilestonesOK {
	return &PostFlowsMilestonesOK{}
}

/*
PostFlowsMilestonesOK describes a response with status code 200, with default header values.

successful operation
*/
type PostFlowsMilestonesOK struct {
	Payload *models.FlowMilestone
}

// IsSuccess returns true when this post flows milestones o k response has a 2xx status code
func (o *PostFlowsMilestonesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post flows milestones o k response has a 3xx status code
func (o *PostFlowsMilestonesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post flows milestones o k response has a 4xx status code
func (o *PostFlowsMilestonesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post flows milestones o k response has a 5xx status code
func (o *PostFlowsMilestonesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post flows milestones o k response a status code equal to that given
func (o *PostFlowsMilestonesOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostFlowsMilestonesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/milestones][%d] postFlowsMilestonesOK  %+v", 200, o.Payload)
}

func (o *PostFlowsMilestonesOK) String() string {
	return fmt.Sprintf("[POST /api/v2/flows/milestones][%d] postFlowsMilestonesOK  %+v", 200, o.Payload)
}

func (o *PostFlowsMilestonesOK) GetPayload() *models.FlowMilestone {
	return o.Payload
}

func (o *PostFlowsMilestonesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FlowMilestone)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsMilestonesBadRequest creates a PostFlowsMilestonesBadRequest with default headers values
func NewPostFlowsMilestonesBadRequest() *PostFlowsMilestonesBadRequest {
	return &PostFlowsMilestonesBadRequest{}
}

/*
PostFlowsMilestonesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostFlowsMilestonesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post flows milestones bad request response has a 2xx status code
func (o *PostFlowsMilestonesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post flows milestones bad request response has a 3xx status code
func (o *PostFlowsMilestonesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post flows milestones bad request response has a 4xx status code
func (o *PostFlowsMilestonesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post flows milestones bad request response has a 5xx status code
func (o *PostFlowsMilestonesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post flows milestones bad request response a status code equal to that given
func (o *PostFlowsMilestonesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostFlowsMilestonesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/milestones][%d] postFlowsMilestonesBadRequest  %+v", 400, o.Payload)
}

func (o *PostFlowsMilestonesBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/flows/milestones][%d] postFlowsMilestonesBadRequest  %+v", 400, o.Payload)
}

func (o *PostFlowsMilestonesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsMilestonesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsMilestonesUnauthorized creates a PostFlowsMilestonesUnauthorized with default headers values
func NewPostFlowsMilestonesUnauthorized() *PostFlowsMilestonesUnauthorized {
	return &PostFlowsMilestonesUnauthorized{}
}

/*
PostFlowsMilestonesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostFlowsMilestonesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post flows milestones unauthorized response has a 2xx status code
func (o *PostFlowsMilestonesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post flows milestones unauthorized response has a 3xx status code
func (o *PostFlowsMilestonesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post flows milestones unauthorized response has a 4xx status code
func (o *PostFlowsMilestonesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post flows milestones unauthorized response has a 5xx status code
func (o *PostFlowsMilestonesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post flows milestones unauthorized response a status code equal to that given
func (o *PostFlowsMilestonesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostFlowsMilestonesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/milestones][%d] postFlowsMilestonesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostFlowsMilestonesUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/flows/milestones][%d] postFlowsMilestonesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostFlowsMilestonesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsMilestonesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsMilestonesForbidden creates a PostFlowsMilestonesForbidden with default headers values
func NewPostFlowsMilestonesForbidden() *PostFlowsMilestonesForbidden {
	return &PostFlowsMilestonesForbidden{}
}

/*
PostFlowsMilestonesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostFlowsMilestonesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post flows milestones forbidden response has a 2xx status code
func (o *PostFlowsMilestonesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post flows milestones forbidden response has a 3xx status code
func (o *PostFlowsMilestonesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post flows milestones forbidden response has a 4xx status code
func (o *PostFlowsMilestonesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post flows milestones forbidden response has a 5xx status code
func (o *PostFlowsMilestonesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post flows milestones forbidden response a status code equal to that given
func (o *PostFlowsMilestonesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostFlowsMilestonesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/milestones][%d] postFlowsMilestonesForbidden  %+v", 403, o.Payload)
}

func (o *PostFlowsMilestonesForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/flows/milestones][%d] postFlowsMilestonesForbidden  %+v", 403, o.Payload)
}

func (o *PostFlowsMilestonesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsMilestonesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsMilestonesNotFound creates a PostFlowsMilestonesNotFound with default headers values
func NewPostFlowsMilestonesNotFound() *PostFlowsMilestonesNotFound {
	return &PostFlowsMilestonesNotFound{}
}

/*
PostFlowsMilestonesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostFlowsMilestonesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post flows milestones not found response has a 2xx status code
func (o *PostFlowsMilestonesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post flows milestones not found response has a 3xx status code
func (o *PostFlowsMilestonesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post flows milestones not found response has a 4xx status code
func (o *PostFlowsMilestonesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post flows milestones not found response has a 5xx status code
func (o *PostFlowsMilestonesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post flows milestones not found response a status code equal to that given
func (o *PostFlowsMilestonesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostFlowsMilestonesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/milestones][%d] postFlowsMilestonesNotFound  %+v", 404, o.Payload)
}

func (o *PostFlowsMilestonesNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/flows/milestones][%d] postFlowsMilestonesNotFound  %+v", 404, o.Payload)
}

func (o *PostFlowsMilestonesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsMilestonesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsMilestonesMethodNotAllowed creates a PostFlowsMilestonesMethodNotAllowed with default headers values
func NewPostFlowsMilestonesMethodNotAllowed() *PostFlowsMilestonesMethodNotAllowed {
	return &PostFlowsMilestonesMethodNotAllowed{}
}

/*
PostFlowsMilestonesMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type PostFlowsMilestonesMethodNotAllowed struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post flows milestones method not allowed response has a 2xx status code
func (o *PostFlowsMilestonesMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post flows milestones method not allowed response has a 3xx status code
func (o *PostFlowsMilestonesMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post flows milestones method not allowed response has a 4xx status code
func (o *PostFlowsMilestonesMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this post flows milestones method not allowed response has a 5xx status code
func (o *PostFlowsMilestonesMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this post flows milestones method not allowed response a status code equal to that given
func (o *PostFlowsMilestonesMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

func (o *PostFlowsMilestonesMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/milestones][%d] postFlowsMilestonesMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *PostFlowsMilestonesMethodNotAllowed) String() string {
	return fmt.Sprintf("[POST /api/v2/flows/milestones][%d] postFlowsMilestonesMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *PostFlowsMilestonesMethodNotAllowed) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsMilestonesMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsMilestonesRequestTimeout creates a PostFlowsMilestonesRequestTimeout with default headers values
func NewPostFlowsMilestonesRequestTimeout() *PostFlowsMilestonesRequestTimeout {
	return &PostFlowsMilestonesRequestTimeout{}
}

/*
PostFlowsMilestonesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostFlowsMilestonesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post flows milestones request timeout response has a 2xx status code
func (o *PostFlowsMilestonesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post flows milestones request timeout response has a 3xx status code
func (o *PostFlowsMilestonesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post flows milestones request timeout response has a 4xx status code
func (o *PostFlowsMilestonesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post flows milestones request timeout response has a 5xx status code
func (o *PostFlowsMilestonesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post flows milestones request timeout response a status code equal to that given
func (o *PostFlowsMilestonesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostFlowsMilestonesRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/milestones][%d] postFlowsMilestonesRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostFlowsMilestonesRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/flows/milestones][%d] postFlowsMilestonesRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostFlowsMilestonesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsMilestonesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsMilestonesConflict creates a PostFlowsMilestonesConflict with default headers values
func NewPostFlowsMilestonesConflict() *PostFlowsMilestonesConflict {
	return &PostFlowsMilestonesConflict{}
}

/*
PostFlowsMilestonesConflict describes a response with status code 409, with default header values.

Conflict
*/
type PostFlowsMilestonesConflict struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post flows milestones conflict response has a 2xx status code
func (o *PostFlowsMilestonesConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post flows milestones conflict response has a 3xx status code
func (o *PostFlowsMilestonesConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post flows milestones conflict response has a 4xx status code
func (o *PostFlowsMilestonesConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this post flows milestones conflict response has a 5xx status code
func (o *PostFlowsMilestonesConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this post flows milestones conflict response a status code equal to that given
func (o *PostFlowsMilestonesConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PostFlowsMilestonesConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/milestones][%d] postFlowsMilestonesConflict  %+v", 409, o.Payload)
}

func (o *PostFlowsMilestonesConflict) String() string {
	return fmt.Sprintf("[POST /api/v2/flows/milestones][%d] postFlowsMilestonesConflict  %+v", 409, o.Payload)
}

func (o *PostFlowsMilestonesConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsMilestonesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsMilestonesRequestEntityTooLarge creates a PostFlowsMilestonesRequestEntityTooLarge with default headers values
func NewPostFlowsMilestonesRequestEntityTooLarge() *PostFlowsMilestonesRequestEntityTooLarge {
	return &PostFlowsMilestonesRequestEntityTooLarge{}
}

/*
PostFlowsMilestonesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostFlowsMilestonesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post flows milestones request entity too large response has a 2xx status code
func (o *PostFlowsMilestonesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post flows milestones request entity too large response has a 3xx status code
func (o *PostFlowsMilestonesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post flows milestones request entity too large response has a 4xx status code
func (o *PostFlowsMilestonesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post flows milestones request entity too large response has a 5xx status code
func (o *PostFlowsMilestonesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post flows milestones request entity too large response a status code equal to that given
func (o *PostFlowsMilestonesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostFlowsMilestonesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/milestones][%d] postFlowsMilestonesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostFlowsMilestonesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/flows/milestones][%d] postFlowsMilestonesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostFlowsMilestonesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsMilestonesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsMilestonesUnsupportedMediaType creates a PostFlowsMilestonesUnsupportedMediaType with default headers values
func NewPostFlowsMilestonesUnsupportedMediaType() *PostFlowsMilestonesUnsupportedMediaType {
	return &PostFlowsMilestonesUnsupportedMediaType{}
}

/*
PostFlowsMilestonesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostFlowsMilestonesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post flows milestones unsupported media type response has a 2xx status code
func (o *PostFlowsMilestonesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post flows milestones unsupported media type response has a 3xx status code
func (o *PostFlowsMilestonesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post flows milestones unsupported media type response has a 4xx status code
func (o *PostFlowsMilestonesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post flows milestones unsupported media type response has a 5xx status code
func (o *PostFlowsMilestonesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post flows milestones unsupported media type response a status code equal to that given
func (o *PostFlowsMilestonesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostFlowsMilestonesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/milestones][%d] postFlowsMilestonesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostFlowsMilestonesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/flows/milestones][%d] postFlowsMilestonesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostFlowsMilestonesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsMilestonesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsMilestonesTooManyRequests creates a PostFlowsMilestonesTooManyRequests with default headers values
func NewPostFlowsMilestonesTooManyRequests() *PostFlowsMilestonesTooManyRequests {
	return &PostFlowsMilestonesTooManyRequests{}
}

/*
PostFlowsMilestonesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostFlowsMilestonesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post flows milestones too many requests response has a 2xx status code
func (o *PostFlowsMilestonesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post flows milestones too many requests response has a 3xx status code
func (o *PostFlowsMilestonesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post flows milestones too many requests response has a 4xx status code
func (o *PostFlowsMilestonesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post flows milestones too many requests response has a 5xx status code
func (o *PostFlowsMilestonesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post flows milestones too many requests response a status code equal to that given
func (o *PostFlowsMilestonesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostFlowsMilestonesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/milestones][%d] postFlowsMilestonesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostFlowsMilestonesTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/flows/milestones][%d] postFlowsMilestonesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostFlowsMilestonesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsMilestonesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsMilestonesInternalServerError creates a PostFlowsMilestonesInternalServerError with default headers values
func NewPostFlowsMilestonesInternalServerError() *PostFlowsMilestonesInternalServerError {
	return &PostFlowsMilestonesInternalServerError{}
}

/*
PostFlowsMilestonesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostFlowsMilestonesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post flows milestones internal server error response has a 2xx status code
func (o *PostFlowsMilestonesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post flows milestones internal server error response has a 3xx status code
func (o *PostFlowsMilestonesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post flows milestones internal server error response has a 4xx status code
func (o *PostFlowsMilestonesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post flows milestones internal server error response has a 5xx status code
func (o *PostFlowsMilestonesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post flows milestones internal server error response a status code equal to that given
func (o *PostFlowsMilestonesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostFlowsMilestonesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/milestones][%d] postFlowsMilestonesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostFlowsMilestonesInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/flows/milestones][%d] postFlowsMilestonesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostFlowsMilestonesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsMilestonesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsMilestonesServiceUnavailable creates a PostFlowsMilestonesServiceUnavailable with default headers values
func NewPostFlowsMilestonesServiceUnavailable() *PostFlowsMilestonesServiceUnavailable {
	return &PostFlowsMilestonesServiceUnavailable{}
}

/*
PostFlowsMilestonesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostFlowsMilestonesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post flows milestones service unavailable response has a 2xx status code
func (o *PostFlowsMilestonesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post flows milestones service unavailable response has a 3xx status code
func (o *PostFlowsMilestonesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post flows milestones service unavailable response has a 4xx status code
func (o *PostFlowsMilestonesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post flows milestones service unavailable response has a 5xx status code
func (o *PostFlowsMilestonesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post flows milestones service unavailable response a status code equal to that given
func (o *PostFlowsMilestonesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostFlowsMilestonesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/milestones][%d] postFlowsMilestonesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostFlowsMilestonesServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/flows/milestones][%d] postFlowsMilestonesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostFlowsMilestonesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsMilestonesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsMilestonesGatewayTimeout creates a PostFlowsMilestonesGatewayTimeout with default headers values
func NewPostFlowsMilestonesGatewayTimeout() *PostFlowsMilestonesGatewayTimeout {
	return &PostFlowsMilestonesGatewayTimeout{}
}

/*
PostFlowsMilestonesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostFlowsMilestonesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post flows milestones gateway timeout response has a 2xx status code
func (o *PostFlowsMilestonesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post flows milestones gateway timeout response has a 3xx status code
func (o *PostFlowsMilestonesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post flows milestones gateway timeout response has a 4xx status code
func (o *PostFlowsMilestonesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post flows milestones gateway timeout response has a 5xx status code
func (o *PostFlowsMilestonesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post flows milestones gateway timeout response a status code equal to that given
func (o *PostFlowsMilestonesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostFlowsMilestonesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/milestones][%d] postFlowsMilestonesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostFlowsMilestonesGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/flows/milestones][%d] postFlowsMilestonesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostFlowsMilestonesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsMilestonesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

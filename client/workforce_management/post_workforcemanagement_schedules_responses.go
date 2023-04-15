// Code generated by go-swagger; DO NOT EDIT.

package workforce_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostWorkforcemanagementSchedulesReader is a Reader for the PostWorkforcemanagementSchedules structure.
type PostWorkforcemanagementSchedulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWorkforcemanagementSchedulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostWorkforcemanagementSchedulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostWorkforcemanagementSchedulesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostWorkforcemanagementSchedulesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostWorkforcemanagementSchedulesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostWorkforcemanagementSchedulesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostWorkforcemanagementSchedulesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostWorkforcemanagementSchedulesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostWorkforcemanagementSchedulesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostWorkforcemanagementSchedulesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostWorkforcemanagementSchedulesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostWorkforcemanagementSchedulesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostWorkforcemanagementSchedulesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostWorkforcemanagementSchedulesOK creates a PostWorkforcemanagementSchedulesOK with default headers values
func NewPostWorkforcemanagementSchedulesOK() *PostWorkforcemanagementSchedulesOK {
	return &PostWorkforcemanagementSchedulesOK{}
}

/*
PostWorkforcemanagementSchedulesOK describes a response with status code 200, with default header values.

successful operation
*/
type PostWorkforcemanagementSchedulesOK struct {
	Payload *models.UserScheduleContainer
}

// IsSuccess returns true when this post workforcemanagement schedules o k response has a 2xx status code
func (o *PostWorkforcemanagementSchedulesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post workforcemanagement schedules o k response has a 3xx status code
func (o *PostWorkforcemanagementSchedulesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement schedules o k response has a 4xx status code
func (o *PostWorkforcemanagementSchedulesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post workforcemanagement schedules o k response has a 5xx status code
func (o *PostWorkforcemanagementSchedulesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement schedules o k response a status code equal to that given
func (o *PostWorkforcemanagementSchedulesOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostWorkforcemanagementSchedulesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/schedules][%d] postWorkforcemanagementSchedulesOK  %+v", 200, o.Payload)
}

func (o *PostWorkforcemanagementSchedulesOK) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/schedules][%d] postWorkforcemanagementSchedulesOK  %+v", 200, o.Payload)
}

func (o *PostWorkforcemanagementSchedulesOK) GetPayload() *models.UserScheduleContainer {
	return o.Payload
}

func (o *PostWorkforcemanagementSchedulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserScheduleContainer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementSchedulesBadRequest creates a PostWorkforcemanagementSchedulesBadRequest with default headers values
func NewPostWorkforcemanagementSchedulesBadRequest() *PostWorkforcemanagementSchedulesBadRequest {
	return &PostWorkforcemanagementSchedulesBadRequest{}
}

/*
PostWorkforcemanagementSchedulesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostWorkforcemanagementSchedulesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement schedules bad request response has a 2xx status code
func (o *PostWorkforcemanagementSchedulesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement schedules bad request response has a 3xx status code
func (o *PostWorkforcemanagementSchedulesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement schedules bad request response has a 4xx status code
func (o *PostWorkforcemanagementSchedulesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement schedules bad request response has a 5xx status code
func (o *PostWorkforcemanagementSchedulesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement schedules bad request response a status code equal to that given
func (o *PostWorkforcemanagementSchedulesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostWorkforcemanagementSchedulesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/schedules][%d] postWorkforcemanagementSchedulesBadRequest  %+v", 400, o.Payload)
}

func (o *PostWorkforcemanagementSchedulesBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/schedules][%d] postWorkforcemanagementSchedulesBadRequest  %+v", 400, o.Payload)
}

func (o *PostWorkforcemanagementSchedulesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementSchedulesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementSchedulesUnauthorized creates a PostWorkforcemanagementSchedulesUnauthorized with default headers values
func NewPostWorkforcemanagementSchedulesUnauthorized() *PostWorkforcemanagementSchedulesUnauthorized {
	return &PostWorkforcemanagementSchedulesUnauthorized{}
}

/*
PostWorkforcemanagementSchedulesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostWorkforcemanagementSchedulesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement schedules unauthorized response has a 2xx status code
func (o *PostWorkforcemanagementSchedulesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement schedules unauthorized response has a 3xx status code
func (o *PostWorkforcemanagementSchedulesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement schedules unauthorized response has a 4xx status code
func (o *PostWorkforcemanagementSchedulesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement schedules unauthorized response has a 5xx status code
func (o *PostWorkforcemanagementSchedulesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement schedules unauthorized response a status code equal to that given
func (o *PostWorkforcemanagementSchedulesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostWorkforcemanagementSchedulesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/schedules][%d] postWorkforcemanagementSchedulesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostWorkforcemanagementSchedulesUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/schedules][%d] postWorkforcemanagementSchedulesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostWorkforcemanagementSchedulesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementSchedulesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementSchedulesForbidden creates a PostWorkforcemanagementSchedulesForbidden with default headers values
func NewPostWorkforcemanagementSchedulesForbidden() *PostWorkforcemanagementSchedulesForbidden {
	return &PostWorkforcemanagementSchedulesForbidden{}
}

/*
PostWorkforcemanagementSchedulesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostWorkforcemanagementSchedulesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement schedules forbidden response has a 2xx status code
func (o *PostWorkforcemanagementSchedulesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement schedules forbidden response has a 3xx status code
func (o *PostWorkforcemanagementSchedulesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement schedules forbidden response has a 4xx status code
func (o *PostWorkforcemanagementSchedulesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement schedules forbidden response has a 5xx status code
func (o *PostWorkforcemanagementSchedulesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement schedules forbidden response a status code equal to that given
func (o *PostWorkforcemanagementSchedulesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostWorkforcemanagementSchedulesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/schedules][%d] postWorkforcemanagementSchedulesForbidden  %+v", 403, o.Payload)
}

func (o *PostWorkforcemanagementSchedulesForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/schedules][%d] postWorkforcemanagementSchedulesForbidden  %+v", 403, o.Payload)
}

func (o *PostWorkforcemanagementSchedulesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementSchedulesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementSchedulesNotFound creates a PostWorkforcemanagementSchedulesNotFound with default headers values
func NewPostWorkforcemanagementSchedulesNotFound() *PostWorkforcemanagementSchedulesNotFound {
	return &PostWorkforcemanagementSchedulesNotFound{}
}

/*
PostWorkforcemanagementSchedulesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostWorkforcemanagementSchedulesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement schedules not found response has a 2xx status code
func (o *PostWorkforcemanagementSchedulesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement schedules not found response has a 3xx status code
func (o *PostWorkforcemanagementSchedulesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement schedules not found response has a 4xx status code
func (o *PostWorkforcemanagementSchedulesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement schedules not found response has a 5xx status code
func (o *PostWorkforcemanagementSchedulesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement schedules not found response a status code equal to that given
func (o *PostWorkforcemanagementSchedulesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostWorkforcemanagementSchedulesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/schedules][%d] postWorkforcemanagementSchedulesNotFound  %+v", 404, o.Payload)
}

func (o *PostWorkforcemanagementSchedulesNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/schedules][%d] postWorkforcemanagementSchedulesNotFound  %+v", 404, o.Payload)
}

func (o *PostWorkforcemanagementSchedulesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementSchedulesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementSchedulesRequestTimeout creates a PostWorkforcemanagementSchedulesRequestTimeout with default headers values
func NewPostWorkforcemanagementSchedulesRequestTimeout() *PostWorkforcemanagementSchedulesRequestTimeout {
	return &PostWorkforcemanagementSchedulesRequestTimeout{}
}

/*
PostWorkforcemanagementSchedulesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostWorkforcemanagementSchedulesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement schedules request timeout response has a 2xx status code
func (o *PostWorkforcemanagementSchedulesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement schedules request timeout response has a 3xx status code
func (o *PostWorkforcemanagementSchedulesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement schedules request timeout response has a 4xx status code
func (o *PostWorkforcemanagementSchedulesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement schedules request timeout response has a 5xx status code
func (o *PostWorkforcemanagementSchedulesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement schedules request timeout response a status code equal to that given
func (o *PostWorkforcemanagementSchedulesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostWorkforcemanagementSchedulesRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/schedules][%d] postWorkforcemanagementSchedulesRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostWorkforcemanagementSchedulesRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/schedules][%d] postWorkforcemanagementSchedulesRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostWorkforcemanagementSchedulesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementSchedulesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementSchedulesRequestEntityTooLarge creates a PostWorkforcemanagementSchedulesRequestEntityTooLarge with default headers values
func NewPostWorkforcemanagementSchedulesRequestEntityTooLarge() *PostWorkforcemanagementSchedulesRequestEntityTooLarge {
	return &PostWorkforcemanagementSchedulesRequestEntityTooLarge{}
}

/*
PostWorkforcemanagementSchedulesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostWorkforcemanagementSchedulesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement schedules request entity too large response has a 2xx status code
func (o *PostWorkforcemanagementSchedulesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement schedules request entity too large response has a 3xx status code
func (o *PostWorkforcemanagementSchedulesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement schedules request entity too large response has a 4xx status code
func (o *PostWorkforcemanagementSchedulesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement schedules request entity too large response has a 5xx status code
func (o *PostWorkforcemanagementSchedulesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement schedules request entity too large response a status code equal to that given
func (o *PostWorkforcemanagementSchedulesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostWorkforcemanagementSchedulesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/schedules][%d] postWorkforcemanagementSchedulesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostWorkforcemanagementSchedulesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/schedules][%d] postWorkforcemanagementSchedulesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostWorkforcemanagementSchedulesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementSchedulesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementSchedulesUnsupportedMediaType creates a PostWorkforcemanagementSchedulesUnsupportedMediaType with default headers values
func NewPostWorkforcemanagementSchedulesUnsupportedMediaType() *PostWorkforcemanagementSchedulesUnsupportedMediaType {
	return &PostWorkforcemanagementSchedulesUnsupportedMediaType{}
}

/*
PostWorkforcemanagementSchedulesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostWorkforcemanagementSchedulesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement schedules unsupported media type response has a 2xx status code
func (o *PostWorkforcemanagementSchedulesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement schedules unsupported media type response has a 3xx status code
func (o *PostWorkforcemanagementSchedulesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement schedules unsupported media type response has a 4xx status code
func (o *PostWorkforcemanagementSchedulesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement schedules unsupported media type response has a 5xx status code
func (o *PostWorkforcemanagementSchedulesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement schedules unsupported media type response a status code equal to that given
func (o *PostWorkforcemanagementSchedulesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostWorkforcemanagementSchedulesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/schedules][%d] postWorkforcemanagementSchedulesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostWorkforcemanagementSchedulesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/schedules][%d] postWorkforcemanagementSchedulesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostWorkforcemanagementSchedulesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementSchedulesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementSchedulesTooManyRequests creates a PostWorkforcemanagementSchedulesTooManyRequests with default headers values
func NewPostWorkforcemanagementSchedulesTooManyRequests() *PostWorkforcemanagementSchedulesTooManyRequests {
	return &PostWorkforcemanagementSchedulesTooManyRequests{}
}

/*
PostWorkforcemanagementSchedulesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostWorkforcemanagementSchedulesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement schedules too many requests response has a 2xx status code
func (o *PostWorkforcemanagementSchedulesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement schedules too many requests response has a 3xx status code
func (o *PostWorkforcemanagementSchedulesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement schedules too many requests response has a 4xx status code
func (o *PostWorkforcemanagementSchedulesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement schedules too many requests response has a 5xx status code
func (o *PostWorkforcemanagementSchedulesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement schedules too many requests response a status code equal to that given
func (o *PostWorkforcemanagementSchedulesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostWorkforcemanagementSchedulesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/schedules][%d] postWorkforcemanagementSchedulesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostWorkforcemanagementSchedulesTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/schedules][%d] postWorkforcemanagementSchedulesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostWorkforcemanagementSchedulesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementSchedulesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementSchedulesInternalServerError creates a PostWorkforcemanagementSchedulesInternalServerError with default headers values
func NewPostWorkforcemanagementSchedulesInternalServerError() *PostWorkforcemanagementSchedulesInternalServerError {
	return &PostWorkforcemanagementSchedulesInternalServerError{}
}

/*
PostWorkforcemanagementSchedulesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostWorkforcemanagementSchedulesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement schedules internal server error response has a 2xx status code
func (o *PostWorkforcemanagementSchedulesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement schedules internal server error response has a 3xx status code
func (o *PostWorkforcemanagementSchedulesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement schedules internal server error response has a 4xx status code
func (o *PostWorkforcemanagementSchedulesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post workforcemanagement schedules internal server error response has a 5xx status code
func (o *PostWorkforcemanagementSchedulesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post workforcemanagement schedules internal server error response a status code equal to that given
func (o *PostWorkforcemanagementSchedulesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostWorkforcemanagementSchedulesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/schedules][%d] postWorkforcemanagementSchedulesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostWorkforcemanagementSchedulesInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/schedules][%d] postWorkforcemanagementSchedulesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostWorkforcemanagementSchedulesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementSchedulesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementSchedulesServiceUnavailable creates a PostWorkforcemanagementSchedulesServiceUnavailable with default headers values
func NewPostWorkforcemanagementSchedulesServiceUnavailable() *PostWorkforcemanagementSchedulesServiceUnavailable {
	return &PostWorkforcemanagementSchedulesServiceUnavailable{}
}

/*
PostWorkforcemanagementSchedulesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostWorkforcemanagementSchedulesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement schedules service unavailable response has a 2xx status code
func (o *PostWorkforcemanagementSchedulesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement schedules service unavailable response has a 3xx status code
func (o *PostWorkforcemanagementSchedulesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement schedules service unavailable response has a 4xx status code
func (o *PostWorkforcemanagementSchedulesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post workforcemanagement schedules service unavailable response has a 5xx status code
func (o *PostWorkforcemanagementSchedulesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post workforcemanagement schedules service unavailable response a status code equal to that given
func (o *PostWorkforcemanagementSchedulesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostWorkforcemanagementSchedulesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/schedules][%d] postWorkforcemanagementSchedulesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostWorkforcemanagementSchedulesServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/schedules][%d] postWorkforcemanagementSchedulesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostWorkforcemanagementSchedulesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementSchedulesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementSchedulesGatewayTimeout creates a PostWorkforcemanagementSchedulesGatewayTimeout with default headers values
func NewPostWorkforcemanagementSchedulesGatewayTimeout() *PostWorkforcemanagementSchedulesGatewayTimeout {
	return &PostWorkforcemanagementSchedulesGatewayTimeout{}
}

/*
PostWorkforcemanagementSchedulesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostWorkforcemanagementSchedulesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement schedules gateway timeout response has a 2xx status code
func (o *PostWorkforcemanagementSchedulesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement schedules gateway timeout response has a 3xx status code
func (o *PostWorkforcemanagementSchedulesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement schedules gateway timeout response has a 4xx status code
func (o *PostWorkforcemanagementSchedulesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post workforcemanagement schedules gateway timeout response has a 5xx status code
func (o *PostWorkforcemanagementSchedulesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post workforcemanagement schedules gateway timeout response a status code equal to that given
func (o *PostWorkforcemanagementSchedulesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostWorkforcemanagementSchedulesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/schedules][%d] postWorkforcemanagementSchedulesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostWorkforcemanagementSchedulesGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/schedules][%d] postWorkforcemanagementSchedulesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostWorkforcemanagementSchedulesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementSchedulesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

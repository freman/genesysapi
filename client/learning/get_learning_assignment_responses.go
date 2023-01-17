// Code generated by go-swagger; DO NOT EDIT.

package learning

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetLearningAssignmentReader is a Reader for the GetLearningAssignment structure.
type GetLearningAssignmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLearningAssignmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLearningAssignmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLearningAssignmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLearningAssignmentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLearningAssignmentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLearningAssignmentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetLearningAssignmentRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetLearningAssignmentRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetLearningAssignmentUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetLearningAssignmentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLearningAssignmentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetLearningAssignmentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetLearningAssignmentGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLearningAssignmentOK creates a GetLearningAssignmentOK with default headers values
func NewGetLearningAssignmentOK() *GetLearningAssignmentOK {
	return &GetLearningAssignmentOK{}
}

/*
GetLearningAssignmentOK describes a response with status code 200, with default header values.

successful operation
*/
type GetLearningAssignmentOK struct {
	Payload *models.LearningAssignment
}

// IsSuccess returns true when this get learning assignment o k response has a 2xx status code
func (o *GetLearningAssignmentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get learning assignment o k response has a 3xx status code
func (o *GetLearningAssignmentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get learning assignment o k response has a 4xx status code
func (o *GetLearningAssignmentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get learning assignment o k response has a 5xx status code
func (o *GetLearningAssignmentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get learning assignment o k response a status code equal to that given
func (o *GetLearningAssignmentOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetLearningAssignmentOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/{assignmentId}][%d] getLearningAssignmentOK  %+v", 200, o.Payload)
}

func (o *GetLearningAssignmentOK) String() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/{assignmentId}][%d] getLearningAssignmentOK  %+v", 200, o.Payload)
}

func (o *GetLearningAssignmentOK) GetPayload() *models.LearningAssignment {
	return o.Payload
}

func (o *GetLearningAssignmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LearningAssignment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningAssignmentBadRequest creates a GetLearningAssignmentBadRequest with default headers values
func NewGetLearningAssignmentBadRequest() *GetLearningAssignmentBadRequest {
	return &GetLearningAssignmentBadRequest{}
}

/*
GetLearningAssignmentBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetLearningAssignmentBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get learning assignment bad request response has a 2xx status code
func (o *GetLearningAssignmentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get learning assignment bad request response has a 3xx status code
func (o *GetLearningAssignmentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get learning assignment bad request response has a 4xx status code
func (o *GetLearningAssignmentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get learning assignment bad request response has a 5xx status code
func (o *GetLearningAssignmentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get learning assignment bad request response a status code equal to that given
func (o *GetLearningAssignmentBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetLearningAssignmentBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/{assignmentId}][%d] getLearningAssignmentBadRequest  %+v", 400, o.Payload)
}

func (o *GetLearningAssignmentBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/{assignmentId}][%d] getLearningAssignmentBadRequest  %+v", 400, o.Payload)
}

func (o *GetLearningAssignmentBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningAssignmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningAssignmentUnauthorized creates a GetLearningAssignmentUnauthorized with default headers values
func NewGetLearningAssignmentUnauthorized() *GetLearningAssignmentUnauthorized {
	return &GetLearningAssignmentUnauthorized{}
}

/*
GetLearningAssignmentUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetLearningAssignmentUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get learning assignment unauthorized response has a 2xx status code
func (o *GetLearningAssignmentUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get learning assignment unauthorized response has a 3xx status code
func (o *GetLearningAssignmentUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get learning assignment unauthorized response has a 4xx status code
func (o *GetLearningAssignmentUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get learning assignment unauthorized response has a 5xx status code
func (o *GetLearningAssignmentUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get learning assignment unauthorized response a status code equal to that given
func (o *GetLearningAssignmentUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetLearningAssignmentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/{assignmentId}][%d] getLearningAssignmentUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLearningAssignmentUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/{assignmentId}][%d] getLearningAssignmentUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLearningAssignmentUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningAssignmentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningAssignmentForbidden creates a GetLearningAssignmentForbidden with default headers values
func NewGetLearningAssignmentForbidden() *GetLearningAssignmentForbidden {
	return &GetLearningAssignmentForbidden{}
}

/*
GetLearningAssignmentForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetLearningAssignmentForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get learning assignment forbidden response has a 2xx status code
func (o *GetLearningAssignmentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get learning assignment forbidden response has a 3xx status code
func (o *GetLearningAssignmentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get learning assignment forbidden response has a 4xx status code
func (o *GetLearningAssignmentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get learning assignment forbidden response has a 5xx status code
func (o *GetLearningAssignmentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get learning assignment forbidden response a status code equal to that given
func (o *GetLearningAssignmentForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetLearningAssignmentForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/{assignmentId}][%d] getLearningAssignmentForbidden  %+v", 403, o.Payload)
}

func (o *GetLearningAssignmentForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/{assignmentId}][%d] getLearningAssignmentForbidden  %+v", 403, o.Payload)
}

func (o *GetLearningAssignmentForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningAssignmentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningAssignmentNotFound creates a GetLearningAssignmentNotFound with default headers values
func NewGetLearningAssignmentNotFound() *GetLearningAssignmentNotFound {
	return &GetLearningAssignmentNotFound{}
}

/*
GetLearningAssignmentNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetLearningAssignmentNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get learning assignment not found response has a 2xx status code
func (o *GetLearningAssignmentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get learning assignment not found response has a 3xx status code
func (o *GetLearningAssignmentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get learning assignment not found response has a 4xx status code
func (o *GetLearningAssignmentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get learning assignment not found response has a 5xx status code
func (o *GetLearningAssignmentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get learning assignment not found response a status code equal to that given
func (o *GetLearningAssignmentNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetLearningAssignmentNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/{assignmentId}][%d] getLearningAssignmentNotFound  %+v", 404, o.Payload)
}

func (o *GetLearningAssignmentNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/{assignmentId}][%d] getLearningAssignmentNotFound  %+v", 404, o.Payload)
}

func (o *GetLearningAssignmentNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningAssignmentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningAssignmentRequestTimeout creates a GetLearningAssignmentRequestTimeout with default headers values
func NewGetLearningAssignmentRequestTimeout() *GetLearningAssignmentRequestTimeout {
	return &GetLearningAssignmentRequestTimeout{}
}

/*
GetLearningAssignmentRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetLearningAssignmentRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get learning assignment request timeout response has a 2xx status code
func (o *GetLearningAssignmentRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get learning assignment request timeout response has a 3xx status code
func (o *GetLearningAssignmentRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get learning assignment request timeout response has a 4xx status code
func (o *GetLearningAssignmentRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get learning assignment request timeout response has a 5xx status code
func (o *GetLearningAssignmentRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get learning assignment request timeout response a status code equal to that given
func (o *GetLearningAssignmentRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetLearningAssignmentRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/{assignmentId}][%d] getLearningAssignmentRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetLearningAssignmentRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/{assignmentId}][%d] getLearningAssignmentRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetLearningAssignmentRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningAssignmentRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningAssignmentRequestEntityTooLarge creates a GetLearningAssignmentRequestEntityTooLarge with default headers values
func NewGetLearningAssignmentRequestEntityTooLarge() *GetLearningAssignmentRequestEntityTooLarge {
	return &GetLearningAssignmentRequestEntityTooLarge{}
}

/*
GetLearningAssignmentRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetLearningAssignmentRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get learning assignment request entity too large response has a 2xx status code
func (o *GetLearningAssignmentRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get learning assignment request entity too large response has a 3xx status code
func (o *GetLearningAssignmentRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get learning assignment request entity too large response has a 4xx status code
func (o *GetLearningAssignmentRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get learning assignment request entity too large response has a 5xx status code
func (o *GetLearningAssignmentRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get learning assignment request entity too large response a status code equal to that given
func (o *GetLearningAssignmentRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetLearningAssignmentRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/{assignmentId}][%d] getLearningAssignmentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetLearningAssignmentRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/{assignmentId}][%d] getLearningAssignmentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetLearningAssignmentRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningAssignmentRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningAssignmentUnsupportedMediaType creates a GetLearningAssignmentUnsupportedMediaType with default headers values
func NewGetLearningAssignmentUnsupportedMediaType() *GetLearningAssignmentUnsupportedMediaType {
	return &GetLearningAssignmentUnsupportedMediaType{}
}

/*
GetLearningAssignmentUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetLearningAssignmentUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get learning assignment unsupported media type response has a 2xx status code
func (o *GetLearningAssignmentUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get learning assignment unsupported media type response has a 3xx status code
func (o *GetLearningAssignmentUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get learning assignment unsupported media type response has a 4xx status code
func (o *GetLearningAssignmentUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get learning assignment unsupported media type response has a 5xx status code
func (o *GetLearningAssignmentUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get learning assignment unsupported media type response a status code equal to that given
func (o *GetLearningAssignmentUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetLearningAssignmentUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/{assignmentId}][%d] getLearningAssignmentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetLearningAssignmentUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/{assignmentId}][%d] getLearningAssignmentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetLearningAssignmentUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningAssignmentUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningAssignmentTooManyRequests creates a GetLearningAssignmentTooManyRequests with default headers values
func NewGetLearningAssignmentTooManyRequests() *GetLearningAssignmentTooManyRequests {
	return &GetLearningAssignmentTooManyRequests{}
}

/*
GetLearningAssignmentTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetLearningAssignmentTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get learning assignment too many requests response has a 2xx status code
func (o *GetLearningAssignmentTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get learning assignment too many requests response has a 3xx status code
func (o *GetLearningAssignmentTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get learning assignment too many requests response has a 4xx status code
func (o *GetLearningAssignmentTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get learning assignment too many requests response has a 5xx status code
func (o *GetLearningAssignmentTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get learning assignment too many requests response a status code equal to that given
func (o *GetLearningAssignmentTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetLearningAssignmentTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/{assignmentId}][%d] getLearningAssignmentTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLearningAssignmentTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/{assignmentId}][%d] getLearningAssignmentTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLearningAssignmentTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningAssignmentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningAssignmentInternalServerError creates a GetLearningAssignmentInternalServerError with default headers values
func NewGetLearningAssignmentInternalServerError() *GetLearningAssignmentInternalServerError {
	return &GetLearningAssignmentInternalServerError{}
}

/*
GetLearningAssignmentInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetLearningAssignmentInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get learning assignment internal server error response has a 2xx status code
func (o *GetLearningAssignmentInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get learning assignment internal server error response has a 3xx status code
func (o *GetLearningAssignmentInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get learning assignment internal server error response has a 4xx status code
func (o *GetLearningAssignmentInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get learning assignment internal server error response has a 5xx status code
func (o *GetLearningAssignmentInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get learning assignment internal server error response a status code equal to that given
func (o *GetLearningAssignmentInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetLearningAssignmentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/{assignmentId}][%d] getLearningAssignmentInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLearningAssignmentInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/{assignmentId}][%d] getLearningAssignmentInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLearningAssignmentInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningAssignmentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningAssignmentServiceUnavailable creates a GetLearningAssignmentServiceUnavailable with default headers values
func NewGetLearningAssignmentServiceUnavailable() *GetLearningAssignmentServiceUnavailable {
	return &GetLearningAssignmentServiceUnavailable{}
}

/*
GetLearningAssignmentServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetLearningAssignmentServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get learning assignment service unavailable response has a 2xx status code
func (o *GetLearningAssignmentServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get learning assignment service unavailable response has a 3xx status code
func (o *GetLearningAssignmentServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get learning assignment service unavailable response has a 4xx status code
func (o *GetLearningAssignmentServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get learning assignment service unavailable response has a 5xx status code
func (o *GetLearningAssignmentServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get learning assignment service unavailable response a status code equal to that given
func (o *GetLearningAssignmentServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetLearningAssignmentServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/{assignmentId}][%d] getLearningAssignmentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLearningAssignmentServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/{assignmentId}][%d] getLearningAssignmentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLearningAssignmentServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningAssignmentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningAssignmentGatewayTimeout creates a GetLearningAssignmentGatewayTimeout with default headers values
func NewGetLearningAssignmentGatewayTimeout() *GetLearningAssignmentGatewayTimeout {
	return &GetLearningAssignmentGatewayTimeout{}
}

/*
GetLearningAssignmentGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetLearningAssignmentGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get learning assignment gateway timeout response has a 2xx status code
func (o *GetLearningAssignmentGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get learning assignment gateway timeout response has a 3xx status code
func (o *GetLearningAssignmentGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get learning assignment gateway timeout response has a 4xx status code
func (o *GetLearningAssignmentGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get learning assignment gateway timeout response has a 5xx status code
func (o *GetLearningAssignmentGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get learning assignment gateway timeout response a status code equal to that given
func (o *GetLearningAssignmentGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetLearningAssignmentGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/{assignmentId}][%d] getLearningAssignmentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetLearningAssignmentGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/{assignmentId}][%d] getLearningAssignmentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetLearningAssignmentGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningAssignmentGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

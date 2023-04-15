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

// GetLearningAssignmentsMeReader is a Reader for the GetLearningAssignmentsMe structure.
type GetLearningAssignmentsMeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLearningAssignmentsMeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLearningAssignmentsMeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLearningAssignmentsMeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLearningAssignmentsMeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLearningAssignmentsMeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLearningAssignmentsMeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetLearningAssignmentsMeRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetLearningAssignmentsMeRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetLearningAssignmentsMeUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetLearningAssignmentsMeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLearningAssignmentsMeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetLearningAssignmentsMeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetLearningAssignmentsMeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLearningAssignmentsMeOK creates a GetLearningAssignmentsMeOK with default headers values
func NewGetLearningAssignmentsMeOK() *GetLearningAssignmentsMeOK {
	return &GetLearningAssignmentsMeOK{}
}

/*
GetLearningAssignmentsMeOK describes a response with status code 200, with default header values.

successful operation
*/
type GetLearningAssignmentsMeOK struct {
	Payload *models.LearningAssignmentsDomainEntity
}

// IsSuccess returns true when this get learning assignments me o k response has a 2xx status code
func (o *GetLearningAssignmentsMeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get learning assignments me o k response has a 3xx status code
func (o *GetLearningAssignmentsMeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get learning assignments me o k response has a 4xx status code
func (o *GetLearningAssignmentsMeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get learning assignments me o k response has a 5xx status code
func (o *GetLearningAssignmentsMeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get learning assignments me o k response a status code equal to that given
func (o *GetLearningAssignmentsMeOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetLearningAssignmentsMeOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/me][%d] getLearningAssignmentsMeOK  %+v", 200, o.Payload)
}

func (o *GetLearningAssignmentsMeOK) String() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/me][%d] getLearningAssignmentsMeOK  %+v", 200, o.Payload)
}

func (o *GetLearningAssignmentsMeOK) GetPayload() *models.LearningAssignmentsDomainEntity {
	return o.Payload
}

func (o *GetLearningAssignmentsMeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LearningAssignmentsDomainEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningAssignmentsMeBadRequest creates a GetLearningAssignmentsMeBadRequest with default headers values
func NewGetLearningAssignmentsMeBadRequest() *GetLearningAssignmentsMeBadRequest {
	return &GetLearningAssignmentsMeBadRequest{}
}

/*
GetLearningAssignmentsMeBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetLearningAssignmentsMeBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get learning assignments me bad request response has a 2xx status code
func (o *GetLearningAssignmentsMeBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get learning assignments me bad request response has a 3xx status code
func (o *GetLearningAssignmentsMeBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get learning assignments me bad request response has a 4xx status code
func (o *GetLearningAssignmentsMeBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get learning assignments me bad request response has a 5xx status code
func (o *GetLearningAssignmentsMeBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get learning assignments me bad request response a status code equal to that given
func (o *GetLearningAssignmentsMeBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetLearningAssignmentsMeBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/me][%d] getLearningAssignmentsMeBadRequest  %+v", 400, o.Payload)
}

func (o *GetLearningAssignmentsMeBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/me][%d] getLearningAssignmentsMeBadRequest  %+v", 400, o.Payload)
}

func (o *GetLearningAssignmentsMeBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningAssignmentsMeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningAssignmentsMeUnauthorized creates a GetLearningAssignmentsMeUnauthorized with default headers values
func NewGetLearningAssignmentsMeUnauthorized() *GetLearningAssignmentsMeUnauthorized {
	return &GetLearningAssignmentsMeUnauthorized{}
}

/*
GetLearningAssignmentsMeUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetLearningAssignmentsMeUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get learning assignments me unauthorized response has a 2xx status code
func (o *GetLearningAssignmentsMeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get learning assignments me unauthorized response has a 3xx status code
func (o *GetLearningAssignmentsMeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get learning assignments me unauthorized response has a 4xx status code
func (o *GetLearningAssignmentsMeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get learning assignments me unauthorized response has a 5xx status code
func (o *GetLearningAssignmentsMeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get learning assignments me unauthorized response a status code equal to that given
func (o *GetLearningAssignmentsMeUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetLearningAssignmentsMeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/me][%d] getLearningAssignmentsMeUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLearningAssignmentsMeUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/me][%d] getLearningAssignmentsMeUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLearningAssignmentsMeUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningAssignmentsMeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningAssignmentsMeForbidden creates a GetLearningAssignmentsMeForbidden with default headers values
func NewGetLearningAssignmentsMeForbidden() *GetLearningAssignmentsMeForbidden {
	return &GetLearningAssignmentsMeForbidden{}
}

/*
GetLearningAssignmentsMeForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetLearningAssignmentsMeForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get learning assignments me forbidden response has a 2xx status code
func (o *GetLearningAssignmentsMeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get learning assignments me forbidden response has a 3xx status code
func (o *GetLearningAssignmentsMeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get learning assignments me forbidden response has a 4xx status code
func (o *GetLearningAssignmentsMeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get learning assignments me forbidden response has a 5xx status code
func (o *GetLearningAssignmentsMeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get learning assignments me forbidden response a status code equal to that given
func (o *GetLearningAssignmentsMeForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetLearningAssignmentsMeForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/me][%d] getLearningAssignmentsMeForbidden  %+v", 403, o.Payload)
}

func (o *GetLearningAssignmentsMeForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/me][%d] getLearningAssignmentsMeForbidden  %+v", 403, o.Payload)
}

func (o *GetLearningAssignmentsMeForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningAssignmentsMeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningAssignmentsMeNotFound creates a GetLearningAssignmentsMeNotFound with default headers values
func NewGetLearningAssignmentsMeNotFound() *GetLearningAssignmentsMeNotFound {
	return &GetLearningAssignmentsMeNotFound{}
}

/*
GetLearningAssignmentsMeNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetLearningAssignmentsMeNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get learning assignments me not found response has a 2xx status code
func (o *GetLearningAssignmentsMeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get learning assignments me not found response has a 3xx status code
func (o *GetLearningAssignmentsMeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get learning assignments me not found response has a 4xx status code
func (o *GetLearningAssignmentsMeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get learning assignments me not found response has a 5xx status code
func (o *GetLearningAssignmentsMeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get learning assignments me not found response a status code equal to that given
func (o *GetLearningAssignmentsMeNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetLearningAssignmentsMeNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/me][%d] getLearningAssignmentsMeNotFound  %+v", 404, o.Payload)
}

func (o *GetLearningAssignmentsMeNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/me][%d] getLearningAssignmentsMeNotFound  %+v", 404, o.Payload)
}

func (o *GetLearningAssignmentsMeNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningAssignmentsMeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningAssignmentsMeRequestTimeout creates a GetLearningAssignmentsMeRequestTimeout with default headers values
func NewGetLearningAssignmentsMeRequestTimeout() *GetLearningAssignmentsMeRequestTimeout {
	return &GetLearningAssignmentsMeRequestTimeout{}
}

/*
GetLearningAssignmentsMeRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetLearningAssignmentsMeRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get learning assignments me request timeout response has a 2xx status code
func (o *GetLearningAssignmentsMeRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get learning assignments me request timeout response has a 3xx status code
func (o *GetLearningAssignmentsMeRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get learning assignments me request timeout response has a 4xx status code
func (o *GetLearningAssignmentsMeRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get learning assignments me request timeout response has a 5xx status code
func (o *GetLearningAssignmentsMeRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get learning assignments me request timeout response a status code equal to that given
func (o *GetLearningAssignmentsMeRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetLearningAssignmentsMeRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/me][%d] getLearningAssignmentsMeRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetLearningAssignmentsMeRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/me][%d] getLearningAssignmentsMeRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetLearningAssignmentsMeRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningAssignmentsMeRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningAssignmentsMeRequestEntityTooLarge creates a GetLearningAssignmentsMeRequestEntityTooLarge with default headers values
func NewGetLearningAssignmentsMeRequestEntityTooLarge() *GetLearningAssignmentsMeRequestEntityTooLarge {
	return &GetLearningAssignmentsMeRequestEntityTooLarge{}
}

/*
GetLearningAssignmentsMeRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetLearningAssignmentsMeRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get learning assignments me request entity too large response has a 2xx status code
func (o *GetLearningAssignmentsMeRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get learning assignments me request entity too large response has a 3xx status code
func (o *GetLearningAssignmentsMeRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get learning assignments me request entity too large response has a 4xx status code
func (o *GetLearningAssignmentsMeRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get learning assignments me request entity too large response has a 5xx status code
func (o *GetLearningAssignmentsMeRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get learning assignments me request entity too large response a status code equal to that given
func (o *GetLearningAssignmentsMeRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetLearningAssignmentsMeRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/me][%d] getLearningAssignmentsMeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetLearningAssignmentsMeRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/me][%d] getLearningAssignmentsMeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetLearningAssignmentsMeRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningAssignmentsMeRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningAssignmentsMeUnsupportedMediaType creates a GetLearningAssignmentsMeUnsupportedMediaType with default headers values
func NewGetLearningAssignmentsMeUnsupportedMediaType() *GetLearningAssignmentsMeUnsupportedMediaType {
	return &GetLearningAssignmentsMeUnsupportedMediaType{}
}

/*
GetLearningAssignmentsMeUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetLearningAssignmentsMeUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get learning assignments me unsupported media type response has a 2xx status code
func (o *GetLearningAssignmentsMeUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get learning assignments me unsupported media type response has a 3xx status code
func (o *GetLearningAssignmentsMeUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get learning assignments me unsupported media type response has a 4xx status code
func (o *GetLearningAssignmentsMeUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get learning assignments me unsupported media type response has a 5xx status code
func (o *GetLearningAssignmentsMeUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get learning assignments me unsupported media type response a status code equal to that given
func (o *GetLearningAssignmentsMeUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetLearningAssignmentsMeUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/me][%d] getLearningAssignmentsMeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetLearningAssignmentsMeUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/me][%d] getLearningAssignmentsMeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetLearningAssignmentsMeUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningAssignmentsMeUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningAssignmentsMeTooManyRequests creates a GetLearningAssignmentsMeTooManyRequests with default headers values
func NewGetLearningAssignmentsMeTooManyRequests() *GetLearningAssignmentsMeTooManyRequests {
	return &GetLearningAssignmentsMeTooManyRequests{}
}

/*
GetLearningAssignmentsMeTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetLearningAssignmentsMeTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get learning assignments me too many requests response has a 2xx status code
func (o *GetLearningAssignmentsMeTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get learning assignments me too many requests response has a 3xx status code
func (o *GetLearningAssignmentsMeTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get learning assignments me too many requests response has a 4xx status code
func (o *GetLearningAssignmentsMeTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get learning assignments me too many requests response has a 5xx status code
func (o *GetLearningAssignmentsMeTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get learning assignments me too many requests response a status code equal to that given
func (o *GetLearningAssignmentsMeTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetLearningAssignmentsMeTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/me][%d] getLearningAssignmentsMeTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLearningAssignmentsMeTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/me][%d] getLearningAssignmentsMeTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLearningAssignmentsMeTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningAssignmentsMeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningAssignmentsMeInternalServerError creates a GetLearningAssignmentsMeInternalServerError with default headers values
func NewGetLearningAssignmentsMeInternalServerError() *GetLearningAssignmentsMeInternalServerError {
	return &GetLearningAssignmentsMeInternalServerError{}
}

/*
GetLearningAssignmentsMeInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetLearningAssignmentsMeInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get learning assignments me internal server error response has a 2xx status code
func (o *GetLearningAssignmentsMeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get learning assignments me internal server error response has a 3xx status code
func (o *GetLearningAssignmentsMeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get learning assignments me internal server error response has a 4xx status code
func (o *GetLearningAssignmentsMeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get learning assignments me internal server error response has a 5xx status code
func (o *GetLearningAssignmentsMeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get learning assignments me internal server error response a status code equal to that given
func (o *GetLearningAssignmentsMeInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetLearningAssignmentsMeInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/me][%d] getLearningAssignmentsMeInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLearningAssignmentsMeInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/me][%d] getLearningAssignmentsMeInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLearningAssignmentsMeInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningAssignmentsMeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningAssignmentsMeServiceUnavailable creates a GetLearningAssignmentsMeServiceUnavailable with default headers values
func NewGetLearningAssignmentsMeServiceUnavailable() *GetLearningAssignmentsMeServiceUnavailable {
	return &GetLearningAssignmentsMeServiceUnavailable{}
}

/*
GetLearningAssignmentsMeServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetLearningAssignmentsMeServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get learning assignments me service unavailable response has a 2xx status code
func (o *GetLearningAssignmentsMeServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get learning assignments me service unavailable response has a 3xx status code
func (o *GetLearningAssignmentsMeServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get learning assignments me service unavailable response has a 4xx status code
func (o *GetLearningAssignmentsMeServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get learning assignments me service unavailable response has a 5xx status code
func (o *GetLearningAssignmentsMeServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get learning assignments me service unavailable response a status code equal to that given
func (o *GetLearningAssignmentsMeServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetLearningAssignmentsMeServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/me][%d] getLearningAssignmentsMeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLearningAssignmentsMeServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/me][%d] getLearningAssignmentsMeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLearningAssignmentsMeServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningAssignmentsMeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningAssignmentsMeGatewayTimeout creates a GetLearningAssignmentsMeGatewayTimeout with default headers values
func NewGetLearningAssignmentsMeGatewayTimeout() *GetLearningAssignmentsMeGatewayTimeout {
	return &GetLearningAssignmentsMeGatewayTimeout{}
}

/*
GetLearningAssignmentsMeGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetLearningAssignmentsMeGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get learning assignments me gateway timeout response has a 2xx status code
func (o *GetLearningAssignmentsMeGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get learning assignments me gateway timeout response has a 3xx status code
func (o *GetLearningAssignmentsMeGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get learning assignments me gateway timeout response has a 4xx status code
func (o *GetLearningAssignmentsMeGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get learning assignments me gateway timeout response has a 5xx status code
func (o *GetLearningAssignmentsMeGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get learning assignments me gateway timeout response a status code equal to that given
func (o *GetLearningAssignmentsMeGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetLearningAssignmentsMeGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/me][%d] getLearningAssignmentsMeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetLearningAssignmentsMeGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/learning/assignments/me][%d] getLearningAssignmentsMeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetLearningAssignmentsMeGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningAssignmentsMeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

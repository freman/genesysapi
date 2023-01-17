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

// PostLearningAssignmentsBulkremoveReader is a Reader for the PostLearningAssignmentsBulkremove structure.
type PostLearningAssignmentsBulkremoveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLearningAssignmentsBulkremoveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLearningAssignmentsBulkremoveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostLearningAssignmentsBulkremoveBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostLearningAssignmentsBulkremoveUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostLearningAssignmentsBulkremoveForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostLearningAssignmentsBulkremoveNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostLearningAssignmentsBulkremoveRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostLearningAssignmentsBulkremoveRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostLearningAssignmentsBulkremoveUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostLearningAssignmentsBulkremoveTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostLearningAssignmentsBulkremoveInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostLearningAssignmentsBulkremoveServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostLearningAssignmentsBulkremoveGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostLearningAssignmentsBulkremoveOK creates a PostLearningAssignmentsBulkremoveOK with default headers values
func NewPostLearningAssignmentsBulkremoveOK() *PostLearningAssignmentsBulkremoveOK {
	return &PostLearningAssignmentsBulkremoveOK{}
}

/*
PostLearningAssignmentsBulkremoveOK describes a response with status code 200, with default header values.

successful operation
*/
type PostLearningAssignmentsBulkremoveOK struct {
	Payload *models.LearningAssignmentBulkRemoveResponse
}

// IsSuccess returns true when this post learning assignments bulkremove o k response has a 2xx status code
func (o *PostLearningAssignmentsBulkremoveOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post learning assignments bulkremove o k response has a 3xx status code
func (o *PostLearningAssignmentsBulkremoveOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post learning assignments bulkremove o k response has a 4xx status code
func (o *PostLearningAssignmentsBulkremoveOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post learning assignments bulkremove o k response has a 5xx status code
func (o *PostLearningAssignmentsBulkremoveOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post learning assignments bulkremove o k response a status code equal to that given
func (o *PostLearningAssignmentsBulkremoveOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostLearningAssignmentsBulkremoveOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/bulkremove][%d] postLearningAssignmentsBulkremoveOK  %+v", 200, o.Payload)
}

func (o *PostLearningAssignmentsBulkremoveOK) String() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/bulkremove][%d] postLearningAssignmentsBulkremoveOK  %+v", 200, o.Payload)
}

func (o *PostLearningAssignmentsBulkremoveOK) GetPayload() *models.LearningAssignmentBulkRemoveResponse {
	return o.Payload
}

func (o *PostLearningAssignmentsBulkremoveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LearningAssignmentBulkRemoveResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentsBulkremoveBadRequest creates a PostLearningAssignmentsBulkremoveBadRequest with default headers values
func NewPostLearningAssignmentsBulkremoveBadRequest() *PostLearningAssignmentsBulkremoveBadRequest {
	return &PostLearningAssignmentsBulkremoveBadRequest{}
}

/*
PostLearningAssignmentsBulkremoveBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostLearningAssignmentsBulkremoveBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post learning assignments bulkremove bad request response has a 2xx status code
func (o *PostLearningAssignmentsBulkremoveBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post learning assignments bulkremove bad request response has a 3xx status code
func (o *PostLearningAssignmentsBulkremoveBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post learning assignments bulkremove bad request response has a 4xx status code
func (o *PostLearningAssignmentsBulkremoveBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post learning assignments bulkremove bad request response has a 5xx status code
func (o *PostLearningAssignmentsBulkremoveBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post learning assignments bulkremove bad request response a status code equal to that given
func (o *PostLearningAssignmentsBulkremoveBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostLearningAssignmentsBulkremoveBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/bulkremove][%d] postLearningAssignmentsBulkremoveBadRequest  %+v", 400, o.Payload)
}

func (o *PostLearningAssignmentsBulkremoveBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/bulkremove][%d] postLearningAssignmentsBulkremoveBadRequest  %+v", 400, o.Payload)
}

func (o *PostLearningAssignmentsBulkremoveBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentsBulkremoveBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentsBulkremoveUnauthorized creates a PostLearningAssignmentsBulkremoveUnauthorized with default headers values
func NewPostLearningAssignmentsBulkremoveUnauthorized() *PostLearningAssignmentsBulkremoveUnauthorized {
	return &PostLearningAssignmentsBulkremoveUnauthorized{}
}

/*
PostLearningAssignmentsBulkremoveUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostLearningAssignmentsBulkremoveUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post learning assignments bulkremove unauthorized response has a 2xx status code
func (o *PostLearningAssignmentsBulkremoveUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post learning assignments bulkremove unauthorized response has a 3xx status code
func (o *PostLearningAssignmentsBulkremoveUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post learning assignments bulkremove unauthorized response has a 4xx status code
func (o *PostLearningAssignmentsBulkremoveUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post learning assignments bulkremove unauthorized response has a 5xx status code
func (o *PostLearningAssignmentsBulkremoveUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post learning assignments bulkremove unauthorized response a status code equal to that given
func (o *PostLearningAssignmentsBulkremoveUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostLearningAssignmentsBulkremoveUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/bulkremove][%d] postLearningAssignmentsBulkremoveUnauthorized  %+v", 401, o.Payload)
}

func (o *PostLearningAssignmentsBulkremoveUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/bulkremove][%d] postLearningAssignmentsBulkremoveUnauthorized  %+v", 401, o.Payload)
}

func (o *PostLearningAssignmentsBulkremoveUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentsBulkremoveUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentsBulkremoveForbidden creates a PostLearningAssignmentsBulkremoveForbidden with default headers values
func NewPostLearningAssignmentsBulkremoveForbidden() *PostLearningAssignmentsBulkremoveForbidden {
	return &PostLearningAssignmentsBulkremoveForbidden{}
}

/*
PostLearningAssignmentsBulkremoveForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostLearningAssignmentsBulkremoveForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post learning assignments bulkremove forbidden response has a 2xx status code
func (o *PostLearningAssignmentsBulkremoveForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post learning assignments bulkremove forbidden response has a 3xx status code
func (o *PostLearningAssignmentsBulkremoveForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post learning assignments bulkremove forbidden response has a 4xx status code
func (o *PostLearningAssignmentsBulkremoveForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post learning assignments bulkremove forbidden response has a 5xx status code
func (o *PostLearningAssignmentsBulkremoveForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post learning assignments bulkremove forbidden response a status code equal to that given
func (o *PostLearningAssignmentsBulkremoveForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostLearningAssignmentsBulkremoveForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/bulkremove][%d] postLearningAssignmentsBulkremoveForbidden  %+v", 403, o.Payload)
}

func (o *PostLearningAssignmentsBulkremoveForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/bulkremove][%d] postLearningAssignmentsBulkremoveForbidden  %+v", 403, o.Payload)
}

func (o *PostLearningAssignmentsBulkremoveForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentsBulkremoveForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentsBulkremoveNotFound creates a PostLearningAssignmentsBulkremoveNotFound with default headers values
func NewPostLearningAssignmentsBulkremoveNotFound() *PostLearningAssignmentsBulkremoveNotFound {
	return &PostLearningAssignmentsBulkremoveNotFound{}
}

/*
PostLearningAssignmentsBulkremoveNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostLearningAssignmentsBulkremoveNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post learning assignments bulkremove not found response has a 2xx status code
func (o *PostLearningAssignmentsBulkremoveNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post learning assignments bulkremove not found response has a 3xx status code
func (o *PostLearningAssignmentsBulkremoveNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post learning assignments bulkremove not found response has a 4xx status code
func (o *PostLearningAssignmentsBulkremoveNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post learning assignments bulkremove not found response has a 5xx status code
func (o *PostLearningAssignmentsBulkremoveNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post learning assignments bulkremove not found response a status code equal to that given
func (o *PostLearningAssignmentsBulkremoveNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostLearningAssignmentsBulkremoveNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/bulkremove][%d] postLearningAssignmentsBulkremoveNotFound  %+v", 404, o.Payload)
}

func (o *PostLearningAssignmentsBulkremoveNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/bulkremove][%d] postLearningAssignmentsBulkremoveNotFound  %+v", 404, o.Payload)
}

func (o *PostLearningAssignmentsBulkremoveNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentsBulkremoveNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentsBulkremoveRequestTimeout creates a PostLearningAssignmentsBulkremoveRequestTimeout with default headers values
func NewPostLearningAssignmentsBulkremoveRequestTimeout() *PostLearningAssignmentsBulkremoveRequestTimeout {
	return &PostLearningAssignmentsBulkremoveRequestTimeout{}
}

/*
PostLearningAssignmentsBulkremoveRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostLearningAssignmentsBulkremoveRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post learning assignments bulkremove request timeout response has a 2xx status code
func (o *PostLearningAssignmentsBulkremoveRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post learning assignments bulkremove request timeout response has a 3xx status code
func (o *PostLearningAssignmentsBulkremoveRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post learning assignments bulkremove request timeout response has a 4xx status code
func (o *PostLearningAssignmentsBulkremoveRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post learning assignments bulkremove request timeout response has a 5xx status code
func (o *PostLearningAssignmentsBulkremoveRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post learning assignments bulkremove request timeout response a status code equal to that given
func (o *PostLearningAssignmentsBulkremoveRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostLearningAssignmentsBulkremoveRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/bulkremove][%d] postLearningAssignmentsBulkremoveRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostLearningAssignmentsBulkremoveRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/bulkremove][%d] postLearningAssignmentsBulkremoveRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostLearningAssignmentsBulkremoveRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentsBulkremoveRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentsBulkremoveRequestEntityTooLarge creates a PostLearningAssignmentsBulkremoveRequestEntityTooLarge with default headers values
func NewPostLearningAssignmentsBulkremoveRequestEntityTooLarge() *PostLearningAssignmentsBulkremoveRequestEntityTooLarge {
	return &PostLearningAssignmentsBulkremoveRequestEntityTooLarge{}
}

/*
PostLearningAssignmentsBulkremoveRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostLearningAssignmentsBulkremoveRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post learning assignments bulkremove request entity too large response has a 2xx status code
func (o *PostLearningAssignmentsBulkremoveRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post learning assignments bulkremove request entity too large response has a 3xx status code
func (o *PostLearningAssignmentsBulkremoveRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post learning assignments bulkremove request entity too large response has a 4xx status code
func (o *PostLearningAssignmentsBulkremoveRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post learning assignments bulkremove request entity too large response has a 5xx status code
func (o *PostLearningAssignmentsBulkremoveRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post learning assignments bulkremove request entity too large response a status code equal to that given
func (o *PostLearningAssignmentsBulkremoveRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostLearningAssignmentsBulkremoveRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/bulkremove][%d] postLearningAssignmentsBulkremoveRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostLearningAssignmentsBulkremoveRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/bulkremove][%d] postLearningAssignmentsBulkremoveRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostLearningAssignmentsBulkremoveRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentsBulkremoveRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentsBulkremoveUnsupportedMediaType creates a PostLearningAssignmentsBulkremoveUnsupportedMediaType with default headers values
func NewPostLearningAssignmentsBulkremoveUnsupportedMediaType() *PostLearningAssignmentsBulkremoveUnsupportedMediaType {
	return &PostLearningAssignmentsBulkremoveUnsupportedMediaType{}
}

/*
PostLearningAssignmentsBulkremoveUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostLearningAssignmentsBulkremoveUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post learning assignments bulkremove unsupported media type response has a 2xx status code
func (o *PostLearningAssignmentsBulkremoveUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post learning assignments bulkremove unsupported media type response has a 3xx status code
func (o *PostLearningAssignmentsBulkremoveUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post learning assignments bulkremove unsupported media type response has a 4xx status code
func (o *PostLearningAssignmentsBulkremoveUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post learning assignments bulkremove unsupported media type response has a 5xx status code
func (o *PostLearningAssignmentsBulkremoveUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post learning assignments bulkremove unsupported media type response a status code equal to that given
func (o *PostLearningAssignmentsBulkremoveUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostLearningAssignmentsBulkremoveUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/bulkremove][%d] postLearningAssignmentsBulkremoveUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostLearningAssignmentsBulkremoveUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/bulkremove][%d] postLearningAssignmentsBulkremoveUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostLearningAssignmentsBulkremoveUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentsBulkremoveUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentsBulkremoveTooManyRequests creates a PostLearningAssignmentsBulkremoveTooManyRequests with default headers values
func NewPostLearningAssignmentsBulkremoveTooManyRequests() *PostLearningAssignmentsBulkremoveTooManyRequests {
	return &PostLearningAssignmentsBulkremoveTooManyRequests{}
}

/*
PostLearningAssignmentsBulkremoveTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostLearningAssignmentsBulkremoveTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post learning assignments bulkremove too many requests response has a 2xx status code
func (o *PostLearningAssignmentsBulkremoveTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post learning assignments bulkremove too many requests response has a 3xx status code
func (o *PostLearningAssignmentsBulkremoveTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post learning assignments bulkremove too many requests response has a 4xx status code
func (o *PostLearningAssignmentsBulkremoveTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post learning assignments bulkremove too many requests response has a 5xx status code
func (o *PostLearningAssignmentsBulkremoveTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post learning assignments bulkremove too many requests response a status code equal to that given
func (o *PostLearningAssignmentsBulkremoveTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostLearningAssignmentsBulkremoveTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/bulkremove][%d] postLearningAssignmentsBulkremoveTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostLearningAssignmentsBulkremoveTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/bulkremove][%d] postLearningAssignmentsBulkremoveTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostLearningAssignmentsBulkremoveTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentsBulkremoveTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentsBulkremoveInternalServerError creates a PostLearningAssignmentsBulkremoveInternalServerError with default headers values
func NewPostLearningAssignmentsBulkremoveInternalServerError() *PostLearningAssignmentsBulkremoveInternalServerError {
	return &PostLearningAssignmentsBulkremoveInternalServerError{}
}

/*
PostLearningAssignmentsBulkremoveInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostLearningAssignmentsBulkremoveInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post learning assignments bulkremove internal server error response has a 2xx status code
func (o *PostLearningAssignmentsBulkremoveInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post learning assignments bulkremove internal server error response has a 3xx status code
func (o *PostLearningAssignmentsBulkremoveInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post learning assignments bulkremove internal server error response has a 4xx status code
func (o *PostLearningAssignmentsBulkremoveInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post learning assignments bulkremove internal server error response has a 5xx status code
func (o *PostLearningAssignmentsBulkremoveInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post learning assignments bulkremove internal server error response a status code equal to that given
func (o *PostLearningAssignmentsBulkremoveInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostLearningAssignmentsBulkremoveInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/bulkremove][%d] postLearningAssignmentsBulkremoveInternalServerError  %+v", 500, o.Payload)
}

func (o *PostLearningAssignmentsBulkremoveInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/bulkremove][%d] postLearningAssignmentsBulkremoveInternalServerError  %+v", 500, o.Payload)
}

func (o *PostLearningAssignmentsBulkremoveInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentsBulkremoveInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentsBulkremoveServiceUnavailable creates a PostLearningAssignmentsBulkremoveServiceUnavailable with default headers values
func NewPostLearningAssignmentsBulkremoveServiceUnavailable() *PostLearningAssignmentsBulkremoveServiceUnavailable {
	return &PostLearningAssignmentsBulkremoveServiceUnavailable{}
}

/*
PostLearningAssignmentsBulkremoveServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostLearningAssignmentsBulkremoveServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post learning assignments bulkremove service unavailable response has a 2xx status code
func (o *PostLearningAssignmentsBulkremoveServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post learning assignments bulkremove service unavailable response has a 3xx status code
func (o *PostLearningAssignmentsBulkremoveServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post learning assignments bulkremove service unavailable response has a 4xx status code
func (o *PostLearningAssignmentsBulkremoveServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post learning assignments bulkremove service unavailable response has a 5xx status code
func (o *PostLearningAssignmentsBulkremoveServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post learning assignments bulkremove service unavailable response a status code equal to that given
func (o *PostLearningAssignmentsBulkremoveServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostLearningAssignmentsBulkremoveServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/bulkremove][%d] postLearningAssignmentsBulkremoveServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostLearningAssignmentsBulkremoveServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/bulkremove][%d] postLearningAssignmentsBulkremoveServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostLearningAssignmentsBulkremoveServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentsBulkremoveServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentsBulkremoveGatewayTimeout creates a PostLearningAssignmentsBulkremoveGatewayTimeout with default headers values
func NewPostLearningAssignmentsBulkremoveGatewayTimeout() *PostLearningAssignmentsBulkremoveGatewayTimeout {
	return &PostLearningAssignmentsBulkremoveGatewayTimeout{}
}

/*
PostLearningAssignmentsBulkremoveGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostLearningAssignmentsBulkremoveGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post learning assignments bulkremove gateway timeout response has a 2xx status code
func (o *PostLearningAssignmentsBulkremoveGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post learning assignments bulkremove gateway timeout response has a 3xx status code
func (o *PostLearningAssignmentsBulkremoveGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post learning assignments bulkremove gateway timeout response has a 4xx status code
func (o *PostLearningAssignmentsBulkremoveGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post learning assignments bulkremove gateway timeout response has a 5xx status code
func (o *PostLearningAssignmentsBulkremoveGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post learning assignments bulkremove gateway timeout response a status code equal to that given
func (o *PostLearningAssignmentsBulkremoveGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostLearningAssignmentsBulkremoveGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/bulkremove][%d] postLearningAssignmentsBulkremoveGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostLearningAssignmentsBulkremoveGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/bulkremove][%d] postLearningAssignmentsBulkremoveGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostLearningAssignmentsBulkremoveGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentsBulkremoveGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

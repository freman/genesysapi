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

// PostLearningAssignmentReassignReader is a Reader for the PostLearningAssignmentReassign structure.
type PostLearningAssignmentReassignReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLearningAssignmentReassignReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLearningAssignmentReassignOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostLearningAssignmentReassignBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostLearningAssignmentReassignUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostLearningAssignmentReassignForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostLearningAssignmentReassignNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostLearningAssignmentReassignRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostLearningAssignmentReassignConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostLearningAssignmentReassignRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostLearningAssignmentReassignUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostLearningAssignmentReassignTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostLearningAssignmentReassignInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostLearningAssignmentReassignServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostLearningAssignmentReassignGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostLearningAssignmentReassignOK creates a PostLearningAssignmentReassignOK with default headers values
func NewPostLearningAssignmentReassignOK() *PostLearningAssignmentReassignOK {
	return &PostLearningAssignmentReassignOK{}
}

/*
PostLearningAssignmentReassignOK describes a response with status code 200, with default header values.

Reassigned version of assignment which can be taken again
*/
type PostLearningAssignmentReassignOK struct {
	Payload *models.LearningAssignment
}

// IsSuccess returns true when this post learning assignment reassign o k response has a 2xx status code
func (o *PostLearningAssignmentReassignOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post learning assignment reassign o k response has a 3xx status code
func (o *PostLearningAssignmentReassignOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post learning assignment reassign o k response has a 4xx status code
func (o *PostLearningAssignmentReassignOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post learning assignment reassign o k response has a 5xx status code
func (o *PostLearningAssignmentReassignOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post learning assignment reassign o k response a status code equal to that given
func (o *PostLearningAssignmentReassignOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostLearningAssignmentReassignOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/{assignmentId}/reassign][%d] postLearningAssignmentReassignOK  %+v", 200, o.Payload)
}

func (o *PostLearningAssignmentReassignOK) String() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/{assignmentId}/reassign][%d] postLearningAssignmentReassignOK  %+v", 200, o.Payload)
}

func (o *PostLearningAssignmentReassignOK) GetPayload() *models.LearningAssignment {
	return o.Payload
}

func (o *PostLearningAssignmentReassignOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LearningAssignment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentReassignBadRequest creates a PostLearningAssignmentReassignBadRequest with default headers values
func NewPostLearningAssignmentReassignBadRequest() *PostLearningAssignmentReassignBadRequest {
	return &PostLearningAssignmentReassignBadRequest{}
}

/*
PostLearningAssignmentReassignBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostLearningAssignmentReassignBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post learning assignment reassign bad request response has a 2xx status code
func (o *PostLearningAssignmentReassignBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post learning assignment reassign bad request response has a 3xx status code
func (o *PostLearningAssignmentReassignBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post learning assignment reassign bad request response has a 4xx status code
func (o *PostLearningAssignmentReassignBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post learning assignment reassign bad request response has a 5xx status code
func (o *PostLearningAssignmentReassignBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post learning assignment reassign bad request response a status code equal to that given
func (o *PostLearningAssignmentReassignBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostLearningAssignmentReassignBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/{assignmentId}/reassign][%d] postLearningAssignmentReassignBadRequest  %+v", 400, o.Payload)
}

func (o *PostLearningAssignmentReassignBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/{assignmentId}/reassign][%d] postLearningAssignmentReassignBadRequest  %+v", 400, o.Payload)
}

func (o *PostLearningAssignmentReassignBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentReassignBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentReassignUnauthorized creates a PostLearningAssignmentReassignUnauthorized with default headers values
func NewPostLearningAssignmentReassignUnauthorized() *PostLearningAssignmentReassignUnauthorized {
	return &PostLearningAssignmentReassignUnauthorized{}
}

/*
PostLearningAssignmentReassignUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostLearningAssignmentReassignUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post learning assignment reassign unauthorized response has a 2xx status code
func (o *PostLearningAssignmentReassignUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post learning assignment reassign unauthorized response has a 3xx status code
func (o *PostLearningAssignmentReassignUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post learning assignment reassign unauthorized response has a 4xx status code
func (o *PostLearningAssignmentReassignUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post learning assignment reassign unauthorized response has a 5xx status code
func (o *PostLearningAssignmentReassignUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post learning assignment reassign unauthorized response a status code equal to that given
func (o *PostLearningAssignmentReassignUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostLearningAssignmentReassignUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/{assignmentId}/reassign][%d] postLearningAssignmentReassignUnauthorized  %+v", 401, o.Payload)
}

func (o *PostLearningAssignmentReassignUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/{assignmentId}/reassign][%d] postLearningAssignmentReassignUnauthorized  %+v", 401, o.Payload)
}

func (o *PostLearningAssignmentReassignUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentReassignUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentReassignForbidden creates a PostLearningAssignmentReassignForbidden with default headers values
func NewPostLearningAssignmentReassignForbidden() *PostLearningAssignmentReassignForbidden {
	return &PostLearningAssignmentReassignForbidden{}
}

/*
PostLearningAssignmentReassignForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostLearningAssignmentReassignForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post learning assignment reassign forbidden response has a 2xx status code
func (o *PostLearningAssignmentReassignForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post learning assignment reassign forbidden response has a 3xx status code
func (o *PostLearningAssignmentReassignForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post learning assignment reassign forbidden response has a 4xx status code
func (o *PostLearningAssignmentReassignForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post learning assignment reassign forbidden response has a 5xx status code
func (o *PostLearningAssignmentReassignForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post learning assignment reassign forbidden response a status code equal to that given
func (o *PostLearningAssignmentReassignForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostLearningAssignmentReassignForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/{assignmentId}/reassign][%d] postLearningAssignmentReassignForbidden  %+v", 403, o.Payload)
}

func (o *PostLearningAssignmentReassignForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/{assignmentId}/reassign][%d] postLearningAssignmentReassignForbidden  %+v", 403, o.Payload)
}

func (o *PostLearningAssignmentReassignForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentReassignForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentReassignNotFound creates a PostLearningAssignmentReassignNotFound with default headers values
func NewPostLearningAssignmentReassignNotFound() *PostLearningAssignmentReassignNotFound {
	return &PostLearningAssignmentReassignNotFound{}
}

/*
PostLearningAssignmentReassignNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostLearningAssignmentReassignNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post learning assignment reassign not found response has a 2xx status code
func (o *PostLearningAssignmentReassignNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post learning assignment reassign not found response has a 3xx status code
func (o *PostLearningAssignmentReassignNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post learning assignment reassign not found response has a 4xx status code
func (o *PostLearningAssignmentReassignNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post learning assignment reassign not found response has a 5xx status code
func (o *PostLearningAssignmentReassignNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post learning assignment reassign not found response a status code equal to that given
func (o *PostLearningAssignmentReassignNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostLearningAssignmentReassignNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/{assignmentId}/reassign][%d] postLearningAssignmentReassignNotFound  %+v", 404, o.Payload)
}

func (o *PostLearningAssignmentReassignNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/{assignmentId}/reassign][%d] postLearningAssignmentReassignNotFound  %+v", 404, o.Payload)
}

func (o *PostLearningAssignmentReassignNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentReassignNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentReassignRequestTimeout creates a PostLearningAssignmentReassignRequestTimeout with default headers values
func NewPostLearningAssignmentReassignRequestTimeout() *PostLearningAssignmentReassignRequestTimeout {
	return &PostLearningAssignmentReassignRequestTimeout{}
}

/*
PostLearningAssignmentReassignRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostLearningAssignmentReassignRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post learning assignment reassign request timeout response has a 2xx status code
func (o *PostLearningAssignmentReassignRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post learning assignment reassign request timeout response has a 3xx status code
func (o *PostLearningAssignmentReassignRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post learning assignment reassign request timeout response has a 4xx status code
func (o *PostLearningAssignmentReassignRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post learning assignment reassign request timeout response has a 5xx status code
func (o *PostLearningAssignmentReassignRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post learning assignment reassign request timeout response a status code equal to that given
func (o *PostLearningAssignmentReassignRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostLearningAssignmentReassignRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/{assignmentId}/reassign][%d] postLearningAssignmentReassignRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostLearningAssignmentReassignRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/{assignmentId}/reassign][%d] postLearningAssignmentReassignRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostLearningAssignmentReassignRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentReassignRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentReassignConflict creates a PostLearningAssignmentReassignConflict with default headers values
func NewPostLearningAssignmentReassignConflict() *PostLearningAssignmentReassignConflict {
	return &PostLearningAssignmentReassignConflict{}
}

/*
PostLearningAssignmentReassignConflict describes a response with status code 409, with default header values.

Conflict
*/
type PostLearningAssignmentReassignConflict struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post learning assignment reassign conflict response has a 2xx status code
func (o *PostLearningAssignmentReassignConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post learning assignment reassign conflict response has a 3xx status code
func (o *PostLearningAssignmentReassignConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post learning assignment reassign conflict response has a 4xx status code
func (o *PostLearningAssignmentReassignConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this post learning assignment reassign conflict response has a 5xx status code
func (o *PostLearningAssignmentReassignConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this post learning assignment reassign conflict response a status code equal to that given
func (o *PostLearningAssignmentReassignConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PostLearningAssignmentReassignConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/{assignmentId}/reassign][%d] postLearningAssignmentReassignConflict  %+v", 409, o.Payload)
}

func (o *PostLearningAssignmentReassignConflict) String() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/{assignmentId}/reassign][%d] postLearningAssignmentReassignConflict  %+v", 409, o.Payload)
}

func (o *PostLearningAssignmentReassignConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentReassignConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentReassignRequestEntityTooLarge creates a PostLearningAssignmentReassignRequestEntityTooLarge with default headers values
func NewPostLearningAssignmentReassignRequestEntityTooLarge() *PostLearningAssignmentReassignRequestEntityTooLarge {
	return &PostLearningAssignmentReassignRequestEntityTooLarge{}
}

/*
PostLearningAssignmentReassignRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostLearningAssignmentReassignRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post learning assignment reassign request entity too large response has a 2xx status code
func (o *PostLearningAssignmentReassignRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post learning assignment reassign request entity too large response has a 3xx status code
func (o *PostLearningAssignmentReassignRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post learning assignment reassign request entity too large response has a 4xx status code
func (o *PostLearningAssignmentReassignRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post learning assignment reassign request entity too large response has a 5xx status code
func (o *PostLearningAssignmentReassignRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post learning assignment reassign request entity too large response a status code equal to that given
func (o *PostLearningAssignmentReassignRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostLearningAssignmentReassignRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/{assignmentId}/reassign][%d] postLearningAssignmentReassignRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostLearningAssignmentReassignRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/{assignmentId}/reassign][%d] postLearningAssignmentReassignRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostLearningAssignmentReassignRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentReassignRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentReassignUnsupportedMediaType creates a PostLearningAssignmentReassignUnsupportedMediaType with default headers values
func NewPostLearningAssignmentReassignUnsupportedMediaType() *PostLearningAssignmentReassignUnsupportedMediaType {
	return &PostLearningAssignmentReassignUnsupportedMediaType{}
}

/*
PostLearningAssignmentReassignUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostLearningAssignmentReassignUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post learning assignment reassign unsupported media type response has a 2xx status code
func (o *PostLearningAssignmentReassignUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post learning assignment reassign unsupported media type response has a 3xx status code
func (o *PostLearningAssignmentReassignUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post learning assignment reassign unsupported media type response has a 4xx status code
func (o *PostLearningAssignmentReassignUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post learning assignment reassign unsupported media type response has a 5xx status code
func (o *PostLearningAssignmentReassignUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post learning assignment reassign unsupported media type response a status code equal to that given
func (o *PostLearningAssignmentReassignUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostLearningAssignmentReassignUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/{assignmentId}/reassign][%d] postLearningAssignmentReassignUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostLearningAssignmentReassignUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/{assignmentId}/reassign][%d] postLearningAssignmentReassignUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostLearningAssignmentReassignUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentReassignUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentReassignTooManyRequests creates a PostLearningAssignmentReassignTooManyRequests with default headers values
func NewPostLearningAssignmentReassignTooManyRequests() *PostLearningAssignmentReassignTooManyRequests {
	return &PostLearningAssignmentReassignTooManyRequests{}
}

/*
PostLearningAssignmentReassignTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostLearningAssignmentReassignTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post learning assignment reassign too many requests response has a 2xx status code
func (o *PostLearningAssignmentReassignTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post learning assignment reassign too many requests response has a 3xx status code
func (o *PostLearningAssignmentReassignTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post learning assignment reassign too many requests response has a 4xx status code
func (o *PostLearningAssignmentReassignTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post learning assignment reassign too many requests response has a 5xx status code
func (o *PostLearningAssignmentReassignTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post learning assignment reassign too many requests response a status code equal to that given
func (o *PostLearningAssignmentReassignTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostLearningAssignmentReassignTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/{assignmentId}/reassign][%d] postLearningAssignmentReassignTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostLearningAssignmentReassignTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/{assignmentId}/reassign][%d] postLearningAssignmentReassignTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostLearningAssignmentReassignTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentReassignTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentReassignInternalServerError creates a PostLearningAssignmentReassignInternalServerError with default headers values
func NewPostLearningAssignmentReassignInternalServerError() *PostLearningAssignmentReassignInternalServerError {
	return &PostLearningAssignmentReassignInternalServerError{}
}

/*
PostLearningAssignmentReassignInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostLearningAssignmentReassignInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post learning assignment reassign internal server error response has a 2xx status code
func (o *PostLearningAssignmentReassignInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post learning assignment reassign internal server error response has a 3xx status code
func (o *PostLearningAssignmentReassignInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post learning assignment reassign internal server error response has a 4xx status code
func (o *PostLearningAssignmentReassignInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post learning assignment reassign internal server error response has a 5xx status code
func (o *PostLearningAssignmentReassignInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post learning assignment reassign internal server error response a status code equal to that given
func (o *PostLearningAssignmentReassignInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostLearningAssignmentReassignInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/{assignmentId}/reassign][%d] postLearningAssignmentReassignInternalServerError  %+v", 500, o.Payload)
}

func (o *PostLearningAssignmentReassignInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/{assignmentId}/reassign][%d] postLearningAssignmentReassignInternalServerError  %+v", 500, o.Payload)
}

func (o *PostLearningAssignmentReassignInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentReassignInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentReassignServiceUnavailable creates a PostLearningAssignmentReassignServiceUnavailable with default headers values
func NewPostLearningAssignmentReassignServiceUnavailable() *PostLearningAssignmentReassignServiceUnavailable {
	return &PostLearningAssignmentReassignServiceUnavailable{}
}

/*
PostLearningAssignmentReassignServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostLearningAssignmentReassignServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post learning assignment reassign service unavailable response has a 2xx status code
func (o *PostLearningAssignmentReassignServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post learning assignment reassign service unavailable response has a 3xx status code
func (o *PostLearningAssignmentReassignServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post learning assignment reassign service unavailable response has a 4xx status code
func (o *PostLearningAssignmentReassignServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post learning assignment reassign service unavailable response has a 5xx status code
func (o *PostLearningAssignmentReassignServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post learning assignment reassign service unavailable response a status code equal to that given
func (o *PostLearningAssignmentReassignServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostLearningAssignmentReassignServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/{assignmentId}/reassign][%d] postLearningAssignmentReassignServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostLearningAssignmentReassignServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/{assignmentId}/reassign][%d] postLearningAssignmentReassignServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostLearningAssignmentReassignServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentReassignServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentReassignGatewayTimeout creates a PostLearningAssignmentReassignGatewayTimeout with default headers values
func NewPostLearningAssignmentReassignGatewayTimeout() *PostLearningAssignmentReassignGatewayTimeout {
	return &PostLearningAssignmentReassignGatewayTimeout{}
}

/*
PostLearningAssignmentReassignGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostLearningAssignmentReassignGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post learning assignment reassign gateway timeout response has a 2xx status code
func (o *PostLearningAssignmentReassignGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post learning assignment reassign gateway timeout response has a 3xx status code
func (o *PostLearningAssignmentReassignGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post learning assignment reassign gateway timeout response has a 4xx status code
func (o *PostLearningAssignmentReassignGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post learning assignment reassign gateway timeout response has a 5xx status code
func (o *PostLearningAssignmentReassignGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post learning assignment reassign gateway timeout response a status code equal to that given
func (o *PostLearningAssignmentReassignGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostLearningAssignmentReassignGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/{assignmentId}/reassign][%d] postLearningAssignmentReassignGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostLearningAssignmentReassignGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/{assignmentId}/reassign][%d] postLearningAssignmentReassignGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostLearningAssignmentReassignGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentReassignGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

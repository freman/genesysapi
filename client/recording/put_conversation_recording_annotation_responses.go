// Code generated by go-swagger; DO NOT EDIT.

package recording

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PutConversationRecordingAnnotationReader is a Reader for the PutConversationRecordingAnnotation structure.
type PutConversationRecordingAnnotationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutConversationRecordingAnnotationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutConversationRecordingAnnotationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutConversationRecordingAnnotationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutConversationRecordingAnnotationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutConversationRecordingAnnotationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutConversationRecordingAnnotationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutConversationRecordingAnnotationRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutConversationRecordingAnnotationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutConversationRecordingAnnotationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutConversationRecordingAnnotationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutConversationRecordingAnnotationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutConversationRecordingAnnotationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutConversationRecordingAnnotationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutConversationRecordingAnnotationOK creates a PutConversationRecordingAnnotationOK with default headers values
func NewPutConversationRecordingAnnotationOK() *PutConversationRecordingAnnotationOK {
	return &PutConversationRecordingAnnotationOK{}
}

/*
PutConversationRecordingAnnotationOK describes a response with status code 200, with default header values.

successful operation
*/
type PutConversationRecordingAnnotationOK struct {
	Payload *models.Annotation
}

// IsSuccess returns true when this put conversation recording annotation o k response has a 2xx status code
func (o *PutConversationRecordingAnnotationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put conversation recording annotation o k response has a 3xx status code
func (o *PutConversationRecordingAnnotationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversation recording annotation o k response has a 4xx status code
func (o *PutConversationRecordingAnnotationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put conversation recording annotation o k response has a 5xx status code
func (o *PutConversationRecordingAnnotationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversation recording annotation o k response a status code equal to that given
func (o *PutConversationRecordingAnnotationOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutConversationRecordingAnnotationOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations/{annotationId}][%d] putConversationRecordingAnnotationOK  %+v", 200, o.Payload)
}

func (o *PutConversationRecordingAnnotationOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations/{annotationId}][%d] putConversationRecordingAnnotationOK  %+v", 200, o.Payload)
}

func (o *PutConversationRecordingAnnotationOK) GetPayload() *models.Annotation {
	return o.Payload
}

func (o *PutConversationRecordingAnnotationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Annotation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationRecordingAnnotationBadRequest creates a PutConversationRecordingAnnotationBadRequest with default headers values
func NewPutConversationRecordingAnnotationBadRequest() *PutConversationRecordingAnnotationBadRequest {
	return &PutConversationRecordingAnnotationBadRequest{}
}

/*
PutConversationRecordingAnnotationBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutConversationRecordingAnnotationBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversation recording annotation bad request response has a 2xx status code
func (o *PutConversationRecordingAnnotationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversation recording annotation bad request response has a 3xx status code
func (o *PutConversationRecordingAnnotationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversation recording annotation bad request response has a 4xx status code
func (o *PutConversationRecordingAnnotationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversation recording annotation bad request response has a 5xx status code
func (o *PutConversationRecordingAnnotationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversation recording annotation bad request response a status code equal to that given
func (o *PutConversationRecordingAnnotationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutConversationRecordingAnnotationBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations/{annotationId}][%d] putConversationRecordingAnnotationBadRequest  %+v", 400, o.Payload)
}

func (o *PutConversationRecordingAnnotationBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations/{annotationId}][%d] putConversationRecordingAnnotationBadRequest  %+v", 400, o.Payload)
}

func (o *PutConversationRecordingAnnotationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationRecordingAnnotationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationRecordingAnnotationUnauthorized creates a PutConversationRecordingAnnotationUnauthorized with default headers values
func NewPutConversationRecordingAnnotationUnauthorized() *PutConversationRecordingAnnotationUnauthorized {
	return &PutConversationRecordingAnnotationUnauthorized{}
}

/*
PutConversationRecordingAnnotationUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PutConversationRecordingAnnotationUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversation recording annotation unauthorized response has a 2xx status code
func (o *PutConversationRecordingAnnotationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversation recording annotation unauthorized response has a 3xx status code
func (o *PutConversationRecordingAnnotationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversation recording annotation unauthorized response has a 4xx status code
func (o *PutConversationRecordingAnnotationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversation recording annotation unauthorized response has a 5xx status code
func (o *PutConversationRecordingAnnotationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversation recording annotation unauthorized response a status code equal to that given
func (o *PutConversationRecordingAnnotationUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PutConversationRecordingAnnotationUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations/{annotationId}][%d] putConversationRecordingAnnotationUnauthorized  %+v", 401, o.Payload)
}

func (o *PutConversationRecordingAnnotationUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations/{annotationId}][%d] putConversationRecordingAnnotationUnauthorized  %+v", 401, o.Payload)
}

func (o *PutConversationRecordingAnnotationUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationRecordingAnnotationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationRecordingAnnotationForbidden creates a PutConversationRecordingAnnotationForbidden with default headers values
func NewPutConversationRecordingAnnotationForbidden() *PutConversationRecordingAnnotationForbidden {
	return &PutConversationRecordingAnnotationForbidden{}
}

/*
PutConversationRecordingAnnotationForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PutConversationRecordingAnnotationForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversation recording annotation forbidden response has a 2xx status code
func (o *PutConversationRecordingAnnotationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversation recording annotation forbidden response has a 3xx status code
func (o *PutConversationRecordingAnnotationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversation recording annotation forbidden response has a 4xx status code
func (o *PutConversationRecordingAnnotationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversation recording annotation forbidden response has a 5xx status code
func (o *PutConversationRecordingAnnotationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversation recording annotation forbidden response a status code equal to that given
func (o *PutConversationRecordingAnnotationForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutConversationRecordingAnnotationForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations/{annotationId}][%d] putConversationRecordingAnnotationForbidden  %+v", 403, o.Payload)
}

func (o *PutConversationRecordingAnnotationForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations/{annotationId}][%d] putConversationRecordingAnnotationForbidden  %+v", 403, o.Payload)
}

func (o *PutConversationRecordingAnnotationForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationRecordingAnnotationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationRecordingAnnotationNotFound creates a PutConversationRecordingAnnotationNotFound with default headers values
func NewPutConversationRecordingAnnotationNotFound() *PutConversationRecordingAnnotationNotFound {
	return &PutConversationRecordingAnnotationNotFound{}
}

/*
PutConversationRecordingAnnotationNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PutConversationRecordingAnnotationNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversation recording annotation not found response has a 2xx status code
func (o *PutConversationRecordingAnnotationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversation recording annotation not found response has a 3xx status code
func (o *PutConversationRecordingAnnotationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversation recording annotation not found response has a 4xx status code
func (o *PutConversationRecordingAnnotationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversation recording annotation not found response has a 5xx status code
func (o *PutConversationRecordingAnnotationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversation recording annotation not found response a status code equal to that given
func (o *PutConversationRecordingAnnotationNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutConversationRecordingAnnotationNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations/{annotationId}][%d] putConversationRecordingAnnotationNotFound  %+v", 404, o.Payload)
}

func (o *PutConversationRecordingAnnotationNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations/{annotationId}][%d] putConversationRecordingAnnotationNotFound  %+v", 404, o.Payload)
}

func (o *PutConversationRecordingAnnotationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationRecordingAnnotationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationRecordingAnnotationRequestTimeout creates a PutConversationRecordingAnnotationRequestTimeout with default headers values
func NewPutConversationRecordingAnnotationRequestTimeout() *PutConversationRecordingAnnotationRequestTimeout {
	return &PutConversationRecordingAnnotationRequestTimeout{}
}

/*
PutConversationRecordingAnnotationRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutConversationRecordingAnnotationRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversation recording annotation request timeout response has a 2xx status code
func (o *PutConversationRecordingAnnotationRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversation recording annotation request timeout response has a 3xx status code
func (o *PutConversationRecordingAnnotationRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversation recording annotation request timeout response has a 4xx status code
func (o *PutConversationRecordingAnnotationRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversation recording annotation request timeout response has a 5xx status code
func (o *PutConversationRecordingAnnotationRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversation recording annotation request timeout response a status code equal to that given
func (o *PutConversationRecordingAnnotationRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PutConversationRecordingAnnotationRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations/{annotationId}][%d] putConversationRecordingAnnotationRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutConversationRecordingAnnotationRequestTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations/{annotationId}][%d] putConversationRecordingAnnotationRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutConversationRecordingAnnotationRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationRecordingAnnotationRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationRecordingAnnotationRequestEntityTooLarge creates a PutConversationRecordingAnnotationRequestEntityTooLarge with default headers values
func NewPutConversationRecordingAnnotationRequestEntityTooLarge() *PutConversationRecordingAnnotationRequestEntityTooLarge {
	return &PutConversationRecordingAnnotationRequestEntityTooLarge{}
}

/*
PutConversationRecordingAnnotationRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PutConversationRecordingAnnotationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversation recording annotation request entity too large response has a 2xx status code
func (o *PutConversationRecordingAnnotationRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversation recording annotation request entity too large response has a 3xx status code
func (o *PutConversationRecordingAnnotationRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversation recording annotation request entity too large response has a 4xx status code
func (o *PutConversationRecordingAnnotationRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversation recording annotation request entity too large response has a 5xx status code
func (o *PutConversationRecordingAnnotationRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversation recording annotation request entity too large response a status code equal to that given
func (o *PutConversationRecordingAnnotationRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PutConversationRecordingAnnotationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations/{annotationId}][%d] putConversationRecordingAnnotationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutConversationRecordingAnnotationRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations/{annotationId}][%d] putConversationRecordingAnnotationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutConversationRecordingAnnotationRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationRecordingAnnotationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationRecordingAnnotationUnsupportedMediaType creates a PutConversationRecordingAnnotationUnsupportedMediaType with default headers values
func NewPutConversationRecordingAnnotationUnsupportedMediaType() *PutConversationRecordingAnnotationUnsupportedMediaType {
	return &PutConversationRecordingAnnotationUnsupportedMediaType{}
}

/*
PutConversationRecordingAnnotationUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutConversationRecordingAnnotationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversation recording annotation unsupported media type response has a 2xx status code
func (o *PutConversationRecordingAnnotationUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversation recording annotation unsupported media type response has a 3xx status code
func (o *PutConversationRecordingAnnotationUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversation recording annotation unsupported media type response has a 4xx status code
func (o *PutConversationRecordingAnnotationUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversation recording annotation unsupported media type response has a 5xx status code
func (o *PutConversationRecordingAnnotationUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversation recording annotation unsupported media type response a status code equal to that given
func (o *PutConversationRecordingAnnotationUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PutConversationRecordingAnnotationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations/{annotationId}][%d] putConversationRecordingAnnotationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutConversationRecordingAnnotationUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations/{annotationId}][%d] putConversationRecordingAnnotationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutConversationRecordingAnnotationUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationRecordingAnnotationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationRecordingAnnotationTooManyRequests creates a PutConversationRecordingAnnotationTooManyRequests with default headers values
func NewPutConversationRecordingAnnotationTooManyRequests() *PutConversationRecordingAnnotationTooManyRequests {
	return &PutConversationRecordingAnnotationTooManyRequests{}
}

/*
PutConversationRecordingAnnotationTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutConversationRecordingAnnotationTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversation recording annotation too many requests response has a 2xx status code
func (o *PutConversationRecordingAnnotationTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversation recording annotation too many requests response has a 3xx status code
func (o *PutConversationRecordingAnnotationTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversation recording annotation too many requests response has a 4xx status code
func (o *PutConversationRecordingAnnotationTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversation recording annotation too many requests response has a 5xx status code
func (o *PutConversationRecordingAnnotationTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversation recording annotation too many requests response a status code equal to that given
func (o *PutConversationRecordingAnnotationTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PutConversationRecordingAnnotationTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations/{annotationId}][%d] putConversationRecordingAnnotationTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutConversationRecordingAnnotationTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations/{annotationId}][%d] putConversationRecordingAnnotationTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutConversationRecordingAnnotationTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationRecordingAnnotationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationRecordingAnnotationInternalServerError creates a PutConversationRecordingAnnotationInternalServerError with default headers values
func NewPutConversationRecordingAnnotationInternalServerError() *PutConversationRecordingAnnotationInternalServerError {
	return &PutConversationRecordingAnnotationInternalServerError{}
}

/*
PutConversationRecordingAnnotationInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutConversationRecordingAnnotationInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversation recording annotation internal server error response has a 2xx status code
func (o *PutConversationRecordingAnnotationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversation recording annotation internal server error response has a 3xx status code
func (o *PutConversationRecordingAnnotationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversation recording annotation internal server error response has a 4xx status code
func (o *PutConversationRecordingAnnotationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put conversation recording annotation internal server error response has a 5xx status code
func (o *PutConversationRecordingAnnotationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put conversation recording annotation internal server error response a status code equal to that given
func (o *PutConversationRecordingAnnotationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutConversationRecordingAnnotationInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations/{annotationId}][%d] putConversationRecordingAnnotationInternalServerError  %+v", 500, o.Payload)
}

func (o *PutConversationRecordingAnnotationInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations/{annotationId}][%d] putConversationRecordingAnnotationInternalServerError  %+v", 500, o.Payload)
}

func (o *PutConversationRecordingAnnotationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationRecordingAnnotationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationRecordingAnnotationServiceUnavailable creates a PutConversationRecordingAnnotationServiceUnavailable with default headers values
func NewPutConversationRecordingAnnotationServiceUnavailable() *PutConversationRecordingAnnotationServiceUnavailable {
	return &PutConversationRecordingAnnotationServiceUnavailable{}
}

/*
PutConversationRecordingAnnotationServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutConversationRecordingAnnotationServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversation recording annotation service unavailable response has a 2xx status code
func (o *PutConversationRecordingAnnotationServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversation recording annotation service unavailable response has a 3xx status code
func (o *PutConversationRecordingAnnotationServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversation recording annotation service unavailable response has a 4xx status code
func (o *PutConversationRecordingAnnotationServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this put conversation recording annotation service unavailable response has a 5xx status code
func (o *PutConversationRecordingAnnotationServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this put conversation recording annotation service unavailable response a status code equal to that given
func (o *PutConversationRecordingAnnotationServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PutConversationRecordingAnnotationServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations/{annotationId}][%d] putConversationRecordingAnnotationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutConversationRecordingAnnotationServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations/{annotationId}][%d] putConversationRecordingAnnotationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutConversationRecordingAnnotationServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationRecordingAnnotationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationRecordingAnnotationGatewayTimeout creates a PutConversationRecordingAnnotationGatewayTimeout with default headers values
func NewPutConversationRecordingAnnotationGatewayTimeout() *PutConversationRecordingAnnotationGatewayTimeout {
	return &PutConversationRecordingAnnotationGatewayTimeout{}
}

/*
PutConversationRecordingAnnotationGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PutConversationRecordingAnnotationGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversation recording annotation gateway timeout response has a 2xx status code
func (o *PutConversationRecordingAnnotationGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversation recording annotation gateway timeout response has a 3xx status code
func (o *PutConversationRecordingAnnotationGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversation recording annotation gateway timeout response has a 4xx status code
func (o *PutConversationRecordingAnnotationGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this put conversation recording annotation gateway timeout response has a 5xx status code
func (o *PutConversationRecordingAnnotationGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this put conversation recording annotation gateway timeout response a status code equal to that given
func (o *PutConversationRecordingAnnotationGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PutConversationRecordingAnnotationGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations/{annotationId}][%d] putConversationRecordingAnnotationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutConversationRecordingAnnotationGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations/{annotationId}][%d] putConversationRecordingAnnotationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutConversationRecordingAnnotationGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationRecordingAnnotationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

/*PutConversationRecordingAnnotationOK handles this case with default header values.

successful operation
*/
type PutConversationRecordingAnnotationOK struct {
	Payload *models.Annotation
}

func (o *PutConversationRecordingAnnotationOK) Error() string {
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

/*PutConversationRecordingAnnotationBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutConversationRecordingAnnotationBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutConversationRecordingAnnotationBadRequest) Error() string {
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

/*PutConversationRecordingAnnotationUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutConversationRecordingAnnotationUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutConversationRecordingAnnotationUnauthorized) Error() string {
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

/*PutConversationRecordingAnnotationForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutConversationRecordingAnnotationForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutConversationRecordingAnnotationForbidden) Error() string {
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

/*PutConversationRecordingAnnotationNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutConversationRecordingAnnotationNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutConversationRecordingAnnotationNotFound) Error() string {
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

// NewPutConversationRecordingAnnotationRequestEntityTooLarge creates a PutConversationRecordingAnnotationRequestEntityTooLarge with default headers values
func NewPutConversationRecordingAnnotationRequestEntityTooLarge() *PutConversationRecordingAnnotationRequestEntityTooLarge {
	return &PutConversationRecordingAnnotationRequestEntityTooLarge{}
}

/*PutConversationRecordingAnnotationRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PutConversationRecordingAnnotationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutConversationRecordingAnnotationRequestEntityTooLarge) Error() string {
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

/*PutConversationRecordingAnnotationUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutConversationRecordingAnnotationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutConversationRecordingAnnotationUnsupportedMediaType) Error() string {
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

/*PutConversationRecordingAnnotationTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PutConversationRecordingAnnotationTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutConversationRecordingAnnotationTooManyRequests) Error() string {
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

/*PutConversationRecordingAnnotationInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutConversationRecordingAnnotationInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutConversationRecordingAnnotationInternalServerError) Error() string {
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

/*PutConversationRecordingAnnotationServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutConversationRecordingAnnotationServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutConversationRecordingAnnotationServiceUnavailable) Error() string {
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

/*PutConversationRecordingAnnotationGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutConversationRecordingAnnotationGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutConversationRecordingAnnotationGatewayTimeout) Error() string {
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

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

// GetConversationRecordingAnnotationReader is a Reader for the GetConversationRecordingAnnotation structure.
type GetConversationRecordingAnnotationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationRecordingAnnotationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationRecordingAnnotationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationRecordingAnnotationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationRecordingAnnotationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationRecordingAnnotationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationRecordingAnnotationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetConversationRecordingAnnotationRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationRecordingAnnotationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationRecordingAnnotationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationRecordingAnnotationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationRecordingAnnotationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationRecordingAnnotationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationRecordingAnnotationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationRecordingAnnotationOK creates a GetConversationRecordingAnnotationOK with default headers values
func NewGetConversationRecordingAnnotationOK() *GetConversationRecordingAnnotationOK {
	return &GetConversationRecordingAnnotationOK{}
}

/*GetConversationRecordingAnnotationOK handles this case with default header values.

successful operation
*/
type GetConversationRecordingAnnotationOK struct {
	Payload *models.Annotation
}

func (o *GetConversationRecordingAnnotationOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations/{annotationId}][%d] getConversationRecordingAnnotationOK  %+v", 200, o.Payload)
}

func (o *GetConversationRecordingAnnotationOK) GetPayload() *models.Annotation {
	return o.Payload
}

func (o *GetConversationRecordingAnnotationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Annotation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingAnnotationBadRequest creates a GetConversationRecordingAnnotationBadRequest with default headers values
func NewGetConversationRecordingAnnotationBadRequest() *GetConversationRecordingAnnotationBadRequest {
	return &GetConversationRecordingAnnotationBadRequest{}
}

/*GetConversationRecordingAnnotationBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationRecordingAnnotationBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRecordingAnnotationBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations/{annotationId}][%d] getConversationRecordingAnnotationBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationRecordingAnnotationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingAnnotationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingAnnotationUnauthorized creates a GetConversationRecordingAnnotationUnauthorized with default headers values
func NewGetConversationRecordingAnnotationUnauthorized() *GetConversationRecordingAnnotationUnauthorized {
	return &GetConversationRecordingAnnotationUnauthorized{}
}

/*GetConversationRecordingAnnotationUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationRecordingAnnotationUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRecordingAnnotationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations/{annotationId}][%d] getConversationRecordingAnnotationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationRecordingAnnotationUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingAnnotationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingAnnotationForbidden creates a GetConversationRecordingAnnotationForbidden with default headers values
func NewGetConversationRecordingAnnotationForbidden() *GetConversationRecordingAnnotationForbidden {
	return &GetConversationRecordingAnnotationForbidden{}
}

/*GetConversationRecordingAnnotationForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationRecordingAnnotationForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRecordingAnnotationForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations/{annotationId}][%d] getConversationRecordingAnnotationForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationRecordingAnnotationForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingAnnotationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingAnnotationNotFound creates a GetConversationRecordingAnnotationNotFound with default headers values
func NewGetConversationRecordingAnnotationNotFound() *GetConversationRecordingAnnotationNotFound {
	return &GetConversationRecordingAnnotationNotFound{}
}

/*GetConversationRecordingAnnotationNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetConversationRecordingAnnotationNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRecordingAnnotationNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations/{annotationId}][%d] getConversationRecordingAnnotationNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationRecordingAnnotationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingAnnotationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingAnnotationRequestTimeout creates a GetConversationRecordingAnnotationRequestTimeout with default headers values
func NewGetConversationRecordingAnnotationRequestTimeout() *GetConversationRecordingAnnotationRequestTimeout {
	return &GetConversationRecordingAnnotationRequestTimeout{}
}

/*GetConversationRecordingAnnotationRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetConversationRecordingAnnotationRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRecordingAnnotationRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations/{annotationId}][%d] getConversationRecordingAnnotationRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationRecordingAnnotationRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingAnnotationRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingAnnotationRequestEntityTooLarge creates a GetConversationRecordingAnnotationRequestEntityTooLarge with default headers values
func NewGetConversationRecordingAnnotationRequestEntityTooLarge() *GetConversationRecordingAnnotationRequestEntityTooLarge {
	return &GetConversationRecordingAnnotationRequestEntityTooLarge{}
}

/*GetConversationRecordingAnnotationRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetConversationRecordingAnnotationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRecordingAnnotationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations/{annotationId}][%d] getConversationRecordingAnnotationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationRecordingAnnotationRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingAnnotationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingAnnotationUnsupportedMediaType creates a GetConversationRecordingAnnotationUnsupportedMediaType with default headers values
func NewGetConversationRecordingAnnotationUnsupportedMediaType() *GetConversationRecordingAnnotationUnsupportedMediaType {
	return &GetConversationRecordingAnnotationUnsupportedMediaType{}
}

/*GetConversationRecordingAnnotationUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationRecordingAnnotationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRecordingAnnotationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations/{annotationId}][%d] getConversationRecordingAnnotationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationRecordingAnnotationUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingAnnotationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingAnnotationTooManyRequests creates a GetConversationRecordingAnnotationTooManyRequests with default headers values
func NewGetConversationRecordingAnnotationTooManyRequests() *GetConversationRecordingAnnotationTooManyRequests {
	return &GetConversationRecordingAnnotationTooManyRequests{}
}

/*GetConversationRecordingAnnotationTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetConversationRecordingAnnotationTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRecordingAnnotationTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations/{annotationId}][%d] getConversationRecordingAnnotationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationRecordingAnnotationTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingAnnotationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingAnnotationInternalServerError creates a GetConversationRecordingAnnotationInternalServerError with default headers values
func NewGetConversationRecordingAnnotationInternalServerError() *GetConversationRecordingAnnotationInternalServerError {
	return &GetConversationRecordingAnnotationInternalServerError{}
}

/*GetConversationRecordingAnnotationInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationRecordingAnnotationInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRecordingAnnotationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations/{annotationId}][%d] getConversationRecordingAnnotationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationRecordingAnnotationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingAnnotationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingAnnotationServiceUnavailable creates a GetConversationRecordingAnnotationServiceUnavailable with default headers values
func NewGetConversationRecordingAnnotationServiceUnavailable() *GetConversationRecordingAnnotationServiceUnavailable {
	return &GetConversationRecordingAnnotationServiceUnavailable{}
}

/*GetConversationRecordingAnnotationServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationRecordingAnnotationServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRecordingAnnotationServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations/{annotationId}][%d] getConversationRecordingAnnotationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationRecordingAnnotationServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingAnnotationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingAnnotationGatewayTimeout creates a GetConversationRecordingAnnotationGatewayTimeout with default headers values
func NewGetConversationRecordingAnnotationGatewayTimeout() *GetConversationRecordingAnnotationGatewayTimeout {
	return &GetConversationRecordingAnnotationGatewayTimeout{}
}

/*GetConversationRecordingAnnotationGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetConversationRecordingAnnotationGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRecordingAnnotationGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations/{annotationId}][%d] getConversationRecordingAnnotationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationRecordingAnnotationGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingAnnotationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

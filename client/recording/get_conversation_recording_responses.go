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

// GetConversationRecordingReader is a Reader for the GetConversationRecording structure.
type GetConversationRecordingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationRecordingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewGetConversationRecordingAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationRecordingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationRecordingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationRecordingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationRecordingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationRecordingRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationRecordingUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationRecordingTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationRecordingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationRecordingServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationRecordingGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationRecordingAccepted creates a GetConversationRecordingAccepted with default headers values
func NewGetConversationRecordingAccepted() *GetConversationRecordingAccepted {
	return &GetConversationRecordingAccepted{}
}

/*GetConversationRecordingAccepted handles this case with default header values.

Success - recording is transcoding
*/
type GetConversationRecordingAccepted struct {
	Payload *models.Recording
}

func (o *GetConversationRecordingAccepted) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}][%d] getConversationRecordingAccepted  %+v", 202, o.Payload)
}

func (o *GetConversationRecordingAccepted) GetPayload() *models.Recording {
	return o.Payload
}

func (o *GetConversationRecordingAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Recording)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingBadRequest creates a GetConversationRecordingBadRequest with default headers values
func NewGetConversationRecordingBadRequest() *GetConversationRecordingBadRequest {
	return &GetConversationRecordingBadRequest{}
}

/*GetConversationRecordingBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationRecordingBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRecordingBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}][%d] getConversationRecordingBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationRecordingBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingUnauthorized creates a GetConversationRecordingUnauthorized with default headers values
func NewGetConversationRecordingUnauthorized() *GetConversationRecordingUnauthorized {
	return &GetConversationRecordingUnauthorized{}
}

/*GetConversationRecordingUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationRecordingUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRecordingUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}][%d] getConversationRecordingUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationRecordingUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingForbidden creates a GetConversationRecordingForbidden with default headers values
func NewGetConversationRecordingForbidden() *GetConversationRecordingForbidden {
	return &GetConversationRecordingForbidden{}
}

/*GetConversationRecordingForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationRecordingForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRecordingForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}][%d] getConversationRecordingForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationRecordingForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingNotFound creates a GetConversationRecordingNotFound with default headers values
func NewGetConversationRecordingNotFound() *GetConversationRecordingNotFound {
	return &GetConversationRecordingNotFound{}
}

/*GetConversationRecordingNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetConversationRecordingNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRecordingNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}][%d] getConversationRecordingNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationRecordingNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingRequestEntityTooLarge creates a GetConversationRecordingRequestEntityTooLarge with default headers values
func NewGetConversationRecordingRequestEntityTooLarge() *GetConversationRecordingRequestEntityTooLarge {
	return &GetConversationRecordingRequestEntityTooLarge{}
}

/*GetConversationRecordingRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetConversationRecordingRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRecordingRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}][%d] getConversationRecordingRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationRecordingRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingUnsupportedMediaType creates a GetConversationRecordingUnsupportedMediaType with default headers values
func NewGetConversationRecordingUnsupportedMediaType() *GetConversationRecordingUnsupportedMediaType {
	return &GetConversationRecordingUnsupportedMediaType{}
}

/*GetConversationRecordingUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationRecordingUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRecordingUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}][%d] getConversationRecordingUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationRecordingUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingTooManyRequests creates a GetConversationRecordingTooManyRequests with default headers values
func NewGetConversationRecordingTooManyRequests() *GetConversationRecordingTooManyRequests {
	return &GetConversationRecordingTooManyRequests{}
}

/*GetConversationRecordingTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetConversationRecordingTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRecordingTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}][%d] getConversationRecordingTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationRecordingTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingInternalServerError creates a GetConversationRecordingInternalServerError with default headers values
func NewGetConversationRecordingInternalServerError() *GetConversationRecordingInternalServerError {
	return &GetConversationRecordingInternalServerError{}
}

/*GetConversationRecordingInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationRecordingInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRecordingInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}][%d] getConversationRecordingInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationRecordingInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingServiceUnavailable creates a GetConversationRecordingServiceUnavailable with default headers values
func NewGetConversationRecordingServiceUnavailable() *GetConversationRecordingServiceUnavailable {
	return &GetConversationRecordingServiceUnavailable{}
}

/*GetConversationRecordingServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationRecordingServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRecordingServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}][%d] getConversationRecordingServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationRecordingServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingGatewayTimeout creates a GetConversationRecordingGatewayTimeout with default headers values
func NewGetConversationRecordingGatewayTimeout() *GetConversationRecordingGatewayTimeout {
	return &GetConversationRecordingGatewayTimeout{}
}

/*GetConversationRecordingGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetConversationRecordingGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRecordingGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}][%d] getConversationRecordingGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationRecordingGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

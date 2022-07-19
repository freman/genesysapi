// Code generated by go-swagger; DO NOT EDIT.

package conversations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PutConversationsCallbackRecordingstateReader is a Reader for the PutConversationsCallbackRecordingstate structure.
type PutConversationsCallbackRecordingstateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutConversationsCallbackRecordingstateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutConversationsCallbackRecordingstateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPutConversationsCallbackRecordingstateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutConversationsCallbackRecordingstateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutConversationsCallbackRecordingstateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutConversationsCallbackRecordingstateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutConversationsCallbackRecordingstateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutConversationsCallbackRecordingstateRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutConversationsCallbackRecordingstateRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutConversationsCallbackRecordingstateUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutConversationsCallbackRecordingstateTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutConversationsCallbackRecordingstateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutConversationsCallbackRecordingstateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutConversationsCallbackRecordingstateGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutConversationsCallbackRecordingstateOK creates a PutConversationsCallbackRecordingstateOK with default headers values
func NewPutConversationsCallbackRecordingstateOK() *PutConversationsCallbackRecordingstateOK {
	return &PutConversationsCallbackRecordingstateOK{}
}

/*PutConversationsCallbackRecordingstateOK handles this case with default header values.

successful operation
*/
type PutConversationsCallbackRecordingstateOK struct {
	Payload string
}

func (o *PutConversationsCallbackRecordingstateOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/callbacks/{conversationId}/recordingstate][%d] putConversationsCallbackRecordingstateOK  %+v", 200, o.Payload)
}

func (o *PutConversationsCallbackRecordingstateOK) GetPayload() string {
	return o.Payload
}

func (o *PutConversationsCallbackRecordingstateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsCallbackRecordingstateAccepted creates a PutConversationsCallbackRecordingstateAccepted with default headers values
func NewPutConversationsCallbackRecordingstateAccepted() *PutConversationsCallbackRecordingstateAccepted {
	return &PutConversationsCallbackRecordingstateAccepted{}
}

/*PutConversationsCallbackRecordingstateAccepted handles this case with default header values.

Accepted - when pausing or resuming recordings (Secure Pause)
*/
type PutConversationsCallbackRecordingstateAccepted struct {
	Payload string
}

func (o *PutConversationsCallbackRecordingstateAccepted) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/callbacks/{conversationId}/recordingstate][%d] putConversationsCallbackRecordingstateAccepted  %+v", 202, o.Payload)
}

func (o *PutConversationsCallbackRecordingstateAccepted) GetPayload() string {
	return o.Payload
}

func (o *PutConversationsCallbackRecordingstateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsCallbackRecordingstateBadRequest creates a PutConversationsCallbackRecordingstateBadRequest with default headers values
func NewPutConversationsCallbackRecordingstateBadRequest() *PutConversationsCallbackRecordingstateBadRequest {
	return &PutConversationsCallbackRecordingstateBadRequest{}
}

/*PutConversationsCallbackRecordingstateBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutConversationsCallbackRecordingstateBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutConversationsCallbackRecordingstateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/callbacks/{conversationId}/recordingstate][%d] putConversationsCallbackRecordingstateBadRequest  %+v", 400, o.Payload)
}

func (o *PutConversationsCallbackRecordingstateBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsCallbackRecordingstateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsCallbackRecordingstateUnauthorized creates a PutConversationsCallbackRecordingstateUnauthorized with default headers values
func NewPutConversationsCallbackRecordingstateUnauthorized() *PutConversationsCallbackRecordingstateUnauthorized {
	return &PutConversationsCallbackRecordingstateUnauthorized{}
}

/*PutConversationsCallbackRecordingstateUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutConversationsCallbackRecordingstateUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutConversationsCallbackRecordingstateUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/callbacks/{conversationId}/recordingstate][%d] putConversationsCallbackRecordingstateUnauthorized  %+v", 401, o.Payload)
}

func (o *PutConversationsCallbackRecordingstateUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsCallbackRecordingstateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsCallbackRecordingstateForbidden creates a PutConversationsCallbackRecordingstateForbidden with default headers values
func NewPutConversationsCallbackRecordingstateForbidden() *PutConversationsCallbackRecordingstateForbidden {
	return &PutConversationsCallbackRecordingstateForbidden{}
}

/*PutConversationsCallbackRecordingstateForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutConversationsCallbackRecordingstateForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutConversationsCallbackRecordingstateForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/callbacks/{conversationId}/recordingstate][%d] putConversationsCallbackRecordingstateForbidden  %+v", 403, o.Payload)
}

func (o *PutConversationsCallbackRecordingstateForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsCallbackRecordingstateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsCallbackRecordingstateNotFound creates a PutConversationsCallbackRecordingstateNotFound with default headers values
func NewPutConversationsCallbackRecordingstateNotFound() *PutConversationsCallbackRecordingstateNotFound {
	return &PutConversationsCallbackRecordingstateNotFound{}
}

/*PutConversationsCallbackRecordingstateNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutConversationsCallbackRecordingstateNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutConversationsCallbackRecordingstateNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/callbacks/{conversationId}/recordingstate][%d] putConversationsCallbackRecordingstateNotFound  %+v", 404, o.Payload)
}

func (o *PutConversationsCallbackRecordingstateNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsCallbackRecordingstateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsCallbackRecordingstateRequestTimeout creates a PutConversationsCallbackRecordingstateRequestTimeout with default headers values
func NewPutConversationsCallbackRecordingstateRequestTimeout() *PutConversationsCallbackRecordingstateRequestTimeout {
	return &PutConversationsCallbackRecordingstateRequestTimeout{}
}

/*PutConversationsCallbackRecordingstateRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutConversationsCallbackRecordingstateRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutConversationsCallbackRecordingstateRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/callbacks/{conversationId}/recordingstate][%d] putConversationsCallbackRecordingstateRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutConversationsCallbackRecordingstateRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsCallbackRecordingstateRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsCallbackRecordingstateRequestEntityTooLarge creates a PutConversationsCallbackRecordingstateRequestEntityTooLarge with default headers values
func NewPutConversationsCallbackRecordingstateRequestEntityTooLarge() *PutConversationsCallbackRecordingstateRequestEntityTooLarge {
	return &PutConversationsCallbackRecordingstateRequestEntityTooLarge{}
}

/*PutConversationsCallbackRecordingstateRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PutConversationsCallbackRecordingstateRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutConversationsCallbackRecordingstateRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/callbacks/{conversationId}/recordingstate][%d] putConversationsCallbackRecordingstateRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutConversationsCallbackRecordingstateRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsCallbackRecordingstateRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsCallbackRecordingstateUnsupportedMediaType creates a PutConversationsCallbackRecordingstateUnsupportedMediaType with default headers values
func NewPutConversationsCallbackRecordingstateUnsupportedMediaType() *PutConversationsCallbackRecordingstateUnsupportedMediaType {
	return &PutConversationsCallbackRecordingstateUnsupportedMediaType{}
}

/*PutConversationsCallbackRecordingstateUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutConversationsCallbackRecordingstateUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutConversationsCallbackRecordingstateUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/callbacks/{conversationId}/recordingstate][%d] putConversationsCallbackRecordingstateUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutConversationsCallbackRecordingstateUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsCallbackRecordingstateUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsCallbackRecordingstateTooManyRequests creates a PutConversationsCallbackRecordingstateTooManyRequests with default headers values
func NewPutConversationsCallbackRecordingstateTooManyRequests() *PutConversationsCallbackRecordingstateTooManyRequests {
	return &PutConversationsCallbackRecordingstateTooManyRequests{}
}

/*PutConversationsCallbackRecordingstateTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutConversationsCallbackRecordingstateTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutConversationsCallbackRecordingstateTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/callbacks/{conversationId}/recordingstate][%d] putConversationsCallbackRecordingstateTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutConversationsCallbackRecordingstateTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsCallbackRecordingstateTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsCallbackRecordingstateInternalServerError creates a PutConversationsCallbackRecordingstateInternalServerError with default headers values
func NewPutConversationsCallbackRecordingstateInternalServerError() *PutConversationsCallbackRecordingstateInternalServerError {
	return &PutConversationsCallbackRecordingstateInternalServerError{}
}

/*PutConversationsCallbackRecordingstateInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutConversationsCallbackRecordingstateInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutConversationsCallbackRecordingstateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/callbacks/{conversationId}/recordingstate][%d] putConversationsCallbackRecordingstateInternalServerError  %+v", 500, o.Payload)
}

func (o *PutConversationsCallbackRecordingstateInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsCallbackRecordingstateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsCallbackRecordingstateServiceUnavailable creates a PutConversationsCallbackRecordingstateServiceUnavailable with default headers values
func NewPutConversationsCallbackRecordingstateServiceUnavailable() *PutConversationsCallbackRecordingstateServiceUnavailable {
	return &PutConversationsCallbackRecordingstateServiceUnavailable{}
}

/*PutConversationsCallbackRecordingstateServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutConversationsCallbackRecordingstateServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutConversationsCallbackRecordingstateServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/callbacks/{conversationId}/recordingstate][%d] putConversationsCallbackRecordingstateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutConversationsCallbackRecordingstateServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsCallbackRecordingstateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsCallbackRecordingstateGatewayTimeout creates a PutConversationsCallbackRecordingstateGatewayTimeout with default headers values
func NewPutConversationsCallbackRecordingstateGatewayTimeout() *PutConversationsCallbackRecordingstateGatewayTimeout {
	return &PutConversationsCallbackRecordingstateGatewayTimeout{}
}

/*PutConversationsCallbackRecordingstateGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutConversationsCallbackRecordingstateGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutConversationsCallbackRecordingstateGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/callbacks/{conversationId}/recordingstate][%d] putConversationsCallbackRecordingstateGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutConversationsCallbackRecordingstateGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsCallbackRecordingstateGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

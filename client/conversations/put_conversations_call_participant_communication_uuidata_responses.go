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

// PutConversationsCallParticipantCommunicationUuidataReader is a Reader for the PutConversationsCallParticipantCommunicationUuidata structure.
type PutConversationsCallParticipantCommunicationUuidataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutConversationsCallParticipantCommunicationUuidataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutConversationsCallParticipantCommunicationUuidataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutConversationsCallParticipantCommunicationUuidataBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutConversationsCallParticipantCommunicationUuidataUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutConversationsCallParticipantCommunicationUuidataForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutConversationsCallParticipantCommunicationUuidataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutConversationsCallParticipantCommunicationUuidataRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutConversationsCallParticipantCommunicationUuidataRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutConversationsCallParticipantCommunicationUuidataUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutConversationsCallParticipantCommunicationUuidataTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutConversationsCallParticipantCommunicationUuidataInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutConversationsCallParticipantCommunicationUuidataServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutConversationsCallParticipantCommunicationUuidataGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutConversationsCallParticipantCommunicationUuidataOK creates a PutConversationsCallParticipantCommunicationUuidataOK with default headers values
func NewPutConversationsCallParticipantCommunicationUuidataOK() *PutConversationsCallParticipantCommunicationUuidataOK {
	return &PutConversationsCallParticipantCommunicationUuidataOK{}
}

/*PutConversationsCallParticipantCommunicationUuidataOK handles this case with default header values.

successful operation
*/
type PutConversationsCallParticipantCommunicationUuidataOK struct {
	Payload models.Empty
}

func (o *PutConversationsCallParticipantCommunicationUuidataOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/calls/{conversationId}/participants/{participantId}/communications/{communicationId}/uuidata][%d] putConversationsCallParticipantCommunicationUuidataOK  %+v", 200, o.Payload)
}

func (o *PutConversationsCallParticipantCommunicationUuidataOK) GetPayload() models.Empty {
	return o.Payload
}

func (o *PutConversationsCallParticipantCommunicationUuidataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsCallParticipantCommunicationUuidataBadRequest creates a PutConversationsCallParticipantCommunicationUuidataBadRequest with default headers values
func NewPutConversationsCallParticipantCommunicationUuidataBadRequest() *PutConversationsCallParticipantCommunicationUuidataBadRequest {
	return &PutConversationsCallParticipantCommunicationUuidataBadRequest{}
}

/*PutConversationsCallParticipantCommunicationUuidataBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutConversationsCallParticipantCommunicationUuidataBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutConversationsCallParticipantCommunicationUuidataBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/calls/{conversationId}/participants/{participantId}/communications/{communicationId}/uuidata][%d] putConversationsCallParticipantCommunicationUuidataBadRequest  %+v", 400, o.Payload)
}

func (o *PutConversationsCallParticipantCommunicationUuidataBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsCallParticipantCommunicationUuidataBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsCallParticipantCommunicationUuidataUnauthorized creates a PutConversationsCallParticipantCommunicationUuidataUnauthorized with default headers values
func NewPutConversationsCallParticipantCommunicationUuidataUnauthorized() *PutConversationsCallParticipantCommunicationUuidataUnauthorized {
	return &PutConversationsCallParticipantCommunicationUuidataUnauthorized{}
}

/*PutConversationsCallParticipantCommunicationUuidataUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutConversationsCallParticipantCommunicationUuidataUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutConversationsCallParticipantCommunicationUuidataUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/calls/{conversationId}/participants/{participantId}/communications/{communicationId}/uuidata][%d] putConversationsCallParticipantCommunicationUuidataUnauthorized  %+v", 401, o.Payload)
}

func (o *PutConversationsCallParticipantCommunicationUuidataUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsCallParticipantCommunicationUuidataUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsCallParticipantCommunicationUuidataForbidden creates a PutConversationsCallParticipantCommunicationUuidataForbidden with default headers values
func NewPutConversationsCallParticipantCommunicationUuidataForbidden() *PutConversationsCallParticipantCommunicationUuidataForbidden {
	return &PutConversationsCallParticipantCommunicationUuidataForbidden{}
}

/*PutConversationsCallParticipantCommunicationUuidataForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutConversationsCallParticipantCommunicationUuidataForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutConversationsCallParticipantCommunicationUuidataForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/calls/{conversationId}/participants/{participantId}/communications/{communicationId}/uuidata][%d] putConversationsCallParticipantCommunicationUuidataForbidden  %+v", 403, o.Payload)
}

func (o *PutConversationsCallParticipantCommunicationUuidataForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsCallParticipantCommunicationUuidataForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsCallParticipantCommunicationUuidataNotFound creates a PutConversationsCallParticipantCommunicationUuidataNotFound with default headers values
func NewPutConversationsCallParticipantCommunicationUuidataNotFound() *PutConversationsCallParticipantCommunicationUuidataNotFound {
	return &PutConversationsCallParticipantCommunicationUuidataNotFound{}
}

/*PutConversationsCallParticipantCommunicationUuidataNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutConversationsCallParticipantCommunicationUuidataNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutConversationsCallParticipantCommunicationUuidataNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/calls/{conversationId}/participants/{participantId}/communications/{communicationId}/uuidata][%d] putConversationsCallParticipantCommunicationUuidataNotFound  %+v", 404, o.Payload)
}

func (o *PutConversationsCallParticipantCommunicationUuidataNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsCallParticipantCommunicationUuidataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsCallParticipantCommunicationUuidataRequestTimeout creates a PutConversationsCallParticipantCommunicationUuidataRequestTimeout with default headers values
func NewPutConversationsCallParticipantCommunicationUuidataRequestTimeout() *PutConversationsCallParticipantCommunicationUuidataRequestTimeout {
	return &PutConversationsCallParticipantCommunicationUuidataRequestTimeout{}
}

/*PutConversationsCallParticipantCommunicationUuidataRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutConversationsCallParticipantCommunicationUuidataRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutConversationsCallParticipantCommunicationUuidataRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/calls/{conversationId}/participants/{participantId}/communications/{communicationId}/uuidata][%d] putConversationsCallParticipantCommunicationUuidataRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutConversationsCallParticipantCommunicationUuidataRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsCallParticipantCommunicationUuidataRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsCallParticipantCommunicationUuidataRequestEntityTooLarge creates a PutConversationsCallParticipantCommunicationUuidataRequestEntityTooLarge with default headers values
func NewPutConversationsCallParticipantCommunicationUuidataRequestEntityTooLarge() *PutConversationsCallParticipantCommunicationUuidataRequestEntityTooLarge {
	return &PutConversationsCallParticipantCommunicationUuidataRequestEntityTooLarge{}
}

/*PutConversationsCallParticipantCommunicationUuidataRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PutConversationsCallParticipantCommunicationUuidataRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutConversationsCallParticipantCommunicationUuidataRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/calls/{conversationId}/participants/{participantId}/communications/{communicationId}/uuidata][%d] putConversationsCallParticipantCommunicationUuidataRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutConversationsCallParticipantCommunicationUuidataRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsCallParticipantCommunicationUuidataRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsCallParticipantCommunicationUuidataUnsupportedMediaType creates a PutConversationsCallParticipantCommunicationUuidataUnsupportedMediaType with default headers values
func NewPutConversationsCallParticipantCommunicationUuidataUnsupportedMediaType() *PutConversationsCallParticipantCommunicationUuidataUnsupportedMediaType {
	return &PutConversationsCallParticipantCommunicationUuidataUnsupportedMediaType{}
}

/*PutConversationsCallParticipantCommunicationUuidataUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutConversationsCallParticipantCommunicationUuidataUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutConversationsCallParticipantCommunicationUuidataUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/calls/{conversationId}/participants/{participantId}/communications/{communicationId}/uuidata][%d] putConversationsCallParticipantCommunicationUuidataUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutConversationsCallParticipantCommunicationUuidataUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsCallParticipantCommunicationUuidataUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsCallParticipantCommunicationUuidataTooManyRequests creates a PutConversationsCallParticipantCommunicationUuidataTooManyRequests with default headers values
func NewPutConversationsCallParticipantCommunicationUuidataTooManyRequests() *PutConversationsCallParticipantCommunicationUuidataTooManyRequests {
	return &PutConversationsCallParticipantCommunicationUuidataTooManyRequests{}
}

/*PutConversationsCallParticipantCommunicationUuidataTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutConversationsCallParticipantCommunicationUuidataTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutConversationsCallParticipantCommunicationUuidataTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/calls/{conversationId}/participants/{participantId}/communications/{communicationId}/uuidata][%d] putConversationsCallParticipantCommunicationUuidataTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutConversationsCallParticipantCommunicationUuidataTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsCallParticipantCommunicationUuidataTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsCallParticipantCommunicationUuidataInternalServerError creates a PutConversationsCallParticipantCommunicationUuidataInternalServerError with default headers values
func NewPutConversationsCallParticipantCommunicationUuidataInternalServerError() *PutConversationsCallParticipantCommunicationUuidataInternalServerError {
	return &PutConversationsCallParticipantCommunicationUuidataInternalServerError{}
}

/*PutConversationsCallParticipantCommunicationUuidataInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutConversationsCallParticipantCommunicationUuidataInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutConversationsCallParticipantCommunicationUuidataInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/calls/{conversationId}/participants/{participantId}/communications/{communicationId}/uuidata][%d] putConversationsCallParticipantCommunicationUuidataInternalServerError  %+v", 500, o.Payload)
}

func (o *PutConversationsCallParticipantCommunicationUuidataInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsCallParticipantCommunicationUuidataInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsCallParticipantCommunicationUuidataServiceUnavailable creates a PutConversationsCallParticipantCommunicationUuidataServiceUnavailable with default headers values
func NewPutConversationsCallParticipantCommunicationUuidataServiceUnavailable() *PutConversationsCallParticipantCommunicationUuidataServiceUnavailable {
	return &PutConversationsCallParticipantCommunicationUuidataServiceUnavailable{}
}

/*PutConversationsCallParticipantCommunicationUuidataServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutConversationsCallParticipantCommunicationUuidataServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutConversationsCallParticipantCommunicationUuidataServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/calls/{conversationId}/participants/{participantId}/communications/{communicationId}/uuidata][%d] putConversationsCallParticipantCommunicationUuidataServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutConversationsCallParticipantCommunicationUuidataServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsCallParticipantCommunicationUuidataServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsCallParticipantCommunicationUuidataGatewayTimeout creates a PutConversationsCallParticipantCommunicationUuidataGatewayTimeout with default headers values
func NewPutConversationsCallParticipantCommunicationUuidataGatewayTimeout() *PutConversationsCallParticipantCommunicationUuidataGatewayTimeout {
	return &PutConversationsCallParticipantCommunicationUuidataGatewayTimeout{}
}

/*PutConversationsCallParticipantCommunicationUuidataGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutConversationsCallParticipantCommunicationUuidataGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutConversationsCallParticipantCommunicationUuidataGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/calls/{conversationId}/participants/{participantId}/communications/{communicationId}/uuidata][%d] putConversationsCallParticipantCommunicationUuidataGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutConversationsCallParticipantCommunicationUuidataGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsCallParticipantCommunicationUuidataGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

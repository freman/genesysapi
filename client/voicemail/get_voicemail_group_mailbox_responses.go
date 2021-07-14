// Code generated by go-swagger; DO NOT EDIT.

package voicemail

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetVoicemailGroupMailboxReader is a Reader for the GetVoicemailGroupMailbox structure.
type GetVoicemailGroupMailboxReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVoicemailGroupMailboxReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVoicemailGroupMailboxOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVoicemailGroupMailboxBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetVoicemailGroupMailboxUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetVoicemailGroupMailboxForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetVoicemailGroupMailboxNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetVoicemailGroupMailboxRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetVoicemailGroupMailboxRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetVoicemailGroupMailboxUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetVoicemailGroupMailboxTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetVoicemailGroupMailboxInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetVoicemailGroupMailboxServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetVoicemailGroupMailboxGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVoicemailGroupMailboxOK creates a GetVoicemailGroupMailboxOK with default headers values
func NewGetVoicemailGroupMailboxOK() *GetVoicemailGroupMailboxOK {
	return &GetVoicemailGroupMailboxOK{}
}

/*GetVoicemailGroupMailboxOK handles this case with default header values.

successful operation
*/
type GetVoicemailGroupMailboxOK struct {
	Payload *models.VoicemailMailboxInfo
}

func (o *GetVoicemailGroupMailboxOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/mailbox][%d] getVoicemailGroupMailboxOK  %+v", 200, o.Payload)
}

func (o *GetVoicemailGroupMailboxOK) GetPayload() *models.VoicemailMailboxInfo {
	return o.Payload
}

func (o *GetVoicemailGroupMailboxOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VoicemailMailboxInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailGroupMailboxBadRequest creates a GetVoicemailGroupMailboxBadRequest with default headers values
func NewGetVoicemailGroupMailboxBadRequest() *GetVoicemailGroupMailboxBadRequest {
	return &GetVoicemailGroupMailboxBadRequest{}
}

/*GetVoicemailGroupMailboxBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetVoicemailGroupMailboxBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetVoicemailGroupMailboxBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/mailbox][%d] getVoicemailGroupMailboxBadRequest  %+v", 400, o.Payload)
}

func (o *GetVoicemailGroupMailboxBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailGroupMailboxBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailGroupMailboxUnauthorized creates a GetVoicemailGroupMailboxUnauthorized with default headers values
func NewGetVoicemailGroupMailboxUnauthorized() *GetVoicemailGroupMailboxUnauthorized {
	return &GetVoicemailGroupMailboxUnauthorized{}
}

/*GetVoicemailGroupMailboxUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetVoicemailGroupMailboxUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetVoicemailGroupMailboxUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/mailbox][%d] getVoicemailGroupMailboxUnauthorized  %+v", 401, o.Payload)
}

func (o *GetVoicemailGroupMailboxUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailGroupMailboxUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailGroupMailboxForbidden creates a GetVoicemailGroupMailboxForbidden with default headers values
func NewGetVoicemailGroupMailboxForbidden() *GetVoicemailGroupMailboxForbidden {
	return &GetVoicemailGroupMailboxForbidden{}
}

/*GetVoicemailGroupMailboxForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetVoicemailGroupMailboxForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetVoicemailGroupMailboxForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/mailbox][%d] getVoicemailGroupMailboxForbidden  %+v", 403, o.Payload)
}

func (o *GetVoicemailGroupMailboxForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailGroupMailboxForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailGroupMailboxNotFound creates a GetVoicemailGroupMailboxNotFound with default headers values
func NewGetVoicemailGroupMailboxNotFound() *GetVoicemailGroupMailboxNotFound {
	return &GetVoicemailGroupMailboxNotFound{}
}

/*GetVoicemailGroupMailboxNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetVoicemailGroupMailboxNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetVoicemailGroupMailboxNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/mailbox][%d] getVoicemailGroupMailboxNotFound  %+v", 404, o.Payload)
}

func (o *GetVoicemailGroupMailboxNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailGroupMailboxNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailGroupMailboxRequestTimeout creates a GetVoicemailGroupMailboxRequestTimeout with default headers values
func NewGetVoicemailGroupMailboxRequestTimeout() *GetVoicemailGroupMailboxRequestTimeout {
	return &GetVoicemailGroupMailboxRequestTimeout{}
}

/*GetVoicemailGroupMailboxRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetVoicemailGroupMailboxRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetVoicemailGroupMailboxRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/mailbox][%d] getVoicemailGroupMailboxRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetVoicemailGroupMailboxRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailGroupMailboxRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailGroupMailboxRequestEntityTooLarge creates a GetVoicemailGroupMailboxRequestEntityTooLarge with default headers values
func NewGetVoicemailGroupMailboxRequestEntityTooLarge() *GetVoicemailGroupMailboxRequestEntityTooLarge {
	return &GetVoicemailGroupMailboxRequestEntityTooLarge{}
}

/*GetVoicemailGroupMailboxRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetVoicemailGroupMailboxRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetVoicemailGroupMailboxRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/mailbox][%d] getVoicemailGroupMailboxRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetVoicemailGroupMailboxRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailGroupMailboxRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailGroupMailboxUnsupportedMediaType creates a GetVoicemailGroupMailboxUnsupportedMediaType with default headers values
func NewGetVoicemailGroupMailboxUnsupportedMediaType() *GetVoicemailGroupMailboxUnsupportedMediaType {
	return &GetVoicemailGroupMailboxUnsupportedMediaType{}
}

/*GetVoicemailGroupMailboxUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetVoicemailGroupMailboxUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetVoicemailGroupMailboxUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/mailbox][%d] getVoicemailGroupMailboxUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetVoicemailGroupMailboxUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailGroupMailboxUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailGroupMailboxTooManyRequests creates a GetVoicemailGroupMailboxTooManyRequests with default headers values
func NewGetVoicemailGroupMailboxTooManyRequests() *GetVoicemailGroupMailboxTooManyRequests {
	return &GetVoicemailGroupMailboxTooManyRequests{}
}

/*GetVoicemailGroupMailboxTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetVoicemailGroupMailboxTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetVoicemailGroupMailboxTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/mailbox][%d] getVoicemailGroupMailboxTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetVoicemailGroupMailboxTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailGroupMailboxTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailGroupMailboxInternalServerError creates a GetVoicemailGroupMailboxInternalServerError with default headers values
func NewGetVoicemailGroupMailboxInternalServerError() *GetVoicemailGroupMailboxInternalServerError {
	return &GetVoicemailGroupMailboxInternalServerError{}
}

/*GetVoicemailGroupMailboxInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetVoicemailGroupMailboxInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetVoicemailGroupMailboxInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/mailbox][%d] getVoicemailGroupMailboxInternalServerError  %+v", 500, o.Payload)
}

func (o *GetVoicemailGroupMailboxInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailGroupMailboxInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailGroupMailboxServiceUnavailable creates a GetVoicemailGroupMailboxServiceUnavailable with default headers values
func NewGetVoicemailGroupMailboxServiceUnavailable() *GetVoicemailGroupMailboxServiceUnavailable {
	return &GetVoicemailGroupMailboxServiceUnavailable{}
}

/*GetVoicemailGroupMailboxServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetVoicemailGroupMailboxServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetVoicemailGroupMailboxServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/mailbox][%d] getVoicemailGroupMailboxServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetVoicemailGroupMailboxServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailGroupMailboxServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailGroupMailboxGatewayTimeout creates a GetVoicemailGroupMailboxGatewayTimeout with default headers values
func NewGetVoicemailGroupMailboxGatewayTimeout() *GetVoicemailGroupMailboxGatewayTimeout {
	return &GetVoicemailGroupMailboxGatewayTimeout{}
}

/*GetVoicemailGroupMailboxGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetVoicemailGroupMailboxGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetVoicemailGroupMailboxGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/mailbox][%d] getVoicemailGroupMailboxGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetVoicemailGroupMailboxGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailGroupMailboxGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

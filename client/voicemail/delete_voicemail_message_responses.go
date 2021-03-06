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

// DeleteVoicemailMessageReader is a Reader for the DeleteVoicemailMessage structure.
type DeleteVoicemailMessageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVoicemailMessageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteVoicemailMessageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteVoicemailMessageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteVoicemailMessageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteVoicemailMessageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteVoicemailMessageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteVoicemailMessageRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteVoicemailMessageUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteVoicemailMessageTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteVoicemailMessageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteVoicemailMessageServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteVoicemailMessageGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteVoicemailMessageOK creates a DeleteVoicemailMessageOK with default headers values
func NewDeleteVoicemailMessageOK() *DeleteVoicemailMessageOK {
	return &DeleteVoicemailMessageOK{}
}

/*DeleteVoicemailMessageOK handles this case with default header values.

Operation was successful.
*/
type DeleteVoicemailMessageOK struct {
}

func (o *DeleteVoicemailMessageOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/voicemail/messages/{messageId}][%d] deleteVoicemailMessageOK ", 200)
}

func (o *DeleteVoicemailMessageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteVoicemailMessageBadRequest creates a DeleteVoicemailMessageBadRequest with default headers values
func NewDeleteVoicemailMessageBadRequest() *DeleteVoicemailMessageBadRequest {
	return &DeleteVoicemailMessageBadRequest{}
}

/*DeleteVoicemailMessageBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteVoicemailMessageBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteVoicemailMessageBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/voicemail/messages/{messageId}][%d] deleteVoicemailMessageBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteVoicemailMessageBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteVoicemailMessageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVoicemailMessageUnauthorized creates a DeleteVoicemailMessageUnauthorized with default headers values
func NewDeleteVoicemailMessageUnauthorized() *DeleteVoicemailMessageUnauthorized {
	return &DeleteVoicemailMessageUnauthorized{}
}

/*DeleteVoicemailMessageUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteVoicemailMessageUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteVoicemailMessageUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/voicemail/messages/{messageId}][%d] deleteVoicemailMessageUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteVoicemailMessageUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteVoicemailMessageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVoicemailMessageForbidden creates a DeleteVoicemailMessageForbidden with default headers values
func NewDeleteVoicemailMessageForbidden() *DeleteVoicemailMessageForbidden {
	return &DeleteVoicemailMessageForbidden{}
}

/*DeleteVoicemailMessageForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteVoicemailMessageForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteVoicemailMessageForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/voicemail/messages/{messageId}][%d] deleteVoicemailMessageForbidden  %+v", 403, o.Payload)
}

func (o *DeleteVoicemailMessageForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteVoicemailMessageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVoicemailMessageNotFound creates a DeleteVoicemailMessageNotFound with default headers values
func NewDeleteVoicemailMessageNotFound() *DeleteVoicemailMessageNotFound {
	return &DeleteVoicemailMessageNotFound{}
}

/*DeleteVoicemailMessageNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteVoicemailMessageNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteVoicemailMessageNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/voicemail/messages/{messageId}][%d] deleteVoicemailMessageNotFound  %+v", 404, o.Payload)
}

func (o *DeleteVoicemailMessageNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteVoicemailMessageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVoicemailMessageRequestEntityTooLarge creates a DeleteVoicemailMessageRequestEntityTooLarge with default headers values
func NewDeleteVoicemailMessageRequestEntityTooLarge() *DeleteVoicemailMessageRequestEntityTooLarge {
	return &DeleteVoicemailMessageRequestEntityTooLarge{}
}

/*DeleteVoicemailMessageRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteVoicemailMessageRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteVoicemailMessageRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/voicemail/messages/{messageId}][%d] deleteVoicemailMessageRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteVoicemailMessageRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteVoicemailMessageRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVoicemailMessageUnsupportedMediaType creates a DeleteVoicemailMessageUnsupportedMediaType with default headers values
func NewDeleteVoicemailMessageUnsupportedMediaType() *DeleteVoicemailMessageUnsupportedMediaType {
	return &DeleteVoicemailMessageUnsupportedMediaType{}
}

/*DeleteVoicemailMessageUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteVoicemailMessageUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteVoicemailMessageUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/voicemail/messages/{messageId}][%d] deleteVoicemailMessageUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteVoicemailMessageUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteVoicemailMessageUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVoicemailMessageTooManyRequests creates a DeleteVoicemailMessageTooManyRequests with default headers values
func NewDeleteVoicemailMessageTooManyRequests() *DeleteVoicemailMessageTooManyRequests {
	return &DeleteVoicemailMessageTooManyRequests{}
}

/*DeleteVoicemailMessageTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteVoicemailMessageTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteVoicemailMessageTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/voicemail/messages/{messageId}][%d] deleteVoicemailMessageTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteVoicemailMessageTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteVoicemailMessageTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVoicemailMessageInternalServerError creates a DeleteVoicemailMessageInternalServerError with default headers values
func NewDeleteVoicemailMessageInternalServerError() *DeleteVoicemailMessageInternalServerError {
	return &DeleteVoicemailMessageInternalServerError{}
}

/*DeleteVoicemailMessageInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteVoicemailMessageInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteVoicemailMessageInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/voicemail/messages/{messageId}][%d] deleteVoicemailMessageInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteVoicemailMessageInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteVoicemailMessageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVoicemailMessageServiceUnavailable creates a DeleteVoicemailMessageServiceUnavailable with default headers values
func NewDeleteVoicemailMessageServiceUnavailable() *DeleteVoicemailMessageServiceUnavailable {
	return &DeleteVoicemailMessageServiceUnavailable{}
}

/*DeleteVoicemailMessageServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteVoicemailMessageServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteVoicemailMessageServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/voicemail/messages/{messageId}][%d] deleteVoicemailMessageServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteVoicemailMessageServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteVoicemailMessageServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVoicemailMessageGatewayTimeout creates a DeleteVoicemailMessageGatewayTimeout with default headers values
func NewDeleteVoicemailMessageGatewayTimeout() *DeleteVoicemailMessageGatewayTimeout {
	return &DeleteVoicemailMessageGatewayTimeout{}
}

/*DeleteVoicemailMessageGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteVoicemailMessageGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteVoicemailMessageGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/voicemail/messages/{messageId}][%d] deleteVoicemailMessageGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteVoicemailMessageGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteVoicemailMessageGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

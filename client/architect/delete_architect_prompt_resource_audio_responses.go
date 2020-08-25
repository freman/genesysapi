// Code generated by go-swagger; DO NOT EDIT.

package architect

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeleteArchitectPromptResourceAudioReader is a Reader for the DeleteArchitectPromptResourceAudio structure.
type DeleteArchitectPromptResourceAudioReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteArchitectPromptResourceAudioReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteArchitectPromptResourceAudioNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteArchitectPromptResourceAudioBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteArchitectPromptResourceAudioUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteArchitectPromptResourceAudioForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteArchitectPromptResourceAudioNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteArchitectPromptResourceAudioRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteArchitectPromptResourceAudioUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteArchitectPromptResourceAudioTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteArchitectPromptResourceAudioInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteArchitectPromptResourceAudioServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteArchitectPromptResourceAudioGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteArchitectPromptResourceAudioNoContent creates a DeleteArchitectPromptResourceAudioNoContent with default headers values
func NewDeleteArchitectPromptResourceAudioNoContent() *DeleteArchitectPromptResourceAudioNoContent {
	return &DeleteArchitectPromptResourceAudioNoContent{}
}

/*DeleteArchitectPromptResourceAudioNoContent handles this case with default header values.

Audio successfully deleted
*/
type DeleteArchitectPromptResourceAudioNoContent struct {
}

func (o *DeleteArchitectPromptResourceAudioNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/architect/prompts/{promptId}/resources/{languageCode}/audio][%d] deleteArchitectPromptResourceAudioNoContent ", 204)
}

func (o *DeleteArchitectPromptResourceAudioNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteArchitectPromptResourceAudioBadRequest creates a DeleteArchitectPromptResourceAudioBadRequest with default headers values
func NewDeleteArchitectPromptResourceAudioBadRequest() *DeleteArchitectPromptResourceAudioBadRequest {
	return &DeleteArchitectPromptResourceAudioBadRequest{}
}

/*DeleteArchitectPromptResourceAudioBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteArchitectPromptResourceAudioBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteArchitectPromptResourceAudioBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/architect/prompts/{promptId}/resources/{languageCode}/audio][%d] deleteArchitectPromptResourceAudioBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteArchitectPromptResourceAudioBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteArchitectPromptResourceAudioBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteArchitectPromptResourceAudioUnauthorized creates a DeleteArchitectPromptResourceAudioUnauthorized with default headers values
func NewDeleteArchitectPromptResourceAudioUnauthorized() *DeleteArchitectPromptResourceAudioUnauthorized {
	return &DeleteArchitectPromptResourceAudioUnauthorized{}
}

/*DeleteArchitectPromptResourceAudioUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteArchitectPromptResourceAudioUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteArchitectPromptResourceAudioUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/architect/prompts/{promptId}/resources/{languageCode}/audio][%d] deleteArchitectPromptResourceAudioUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteArchitectPromptResourceAudioUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteArchitectPromptResourceAudioUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteArchitectPromptResourceAudioForbidden creates a DeleteArchitectPromptResourceAudioForbidden with default headers values
func NewDeleteArchitectPromptResourceAudioForbidden() *DeleteArchitectPromptResourceAudioForbidden {
	return &DeleteArchitectPromptResourceAudioForbidden{}
}

/*DeleteArchitectPromptResourceAudioForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteArchitectPromptResourceAudioForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteArchitectPromptResourceAudioForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/architect/prompts/{promptId}/resources/{languageCode}/audio][%d] deleteArchitectPromptResourceAudioForbidden  %+v", 403, o.Payload)
}

func (o *DeleteArchitectPromptResourceAudioForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteArchitectPromptResourceAudioForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteArchitectPromptResourceAudioNotFound creates a DeleteArchitectPromptResourceAudioNotFound with default headers values
func NewDeleteArchitectPromptResourceAudioNotFound() *DeleteArchitectPromptResourceAudioNotFound {
	return &DeleteArchitectPromptResourceAudioNotFound{}
}

/*DeleteArchitectPromptResourceAudioNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteArchitectPromptResourceAudioNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteArchitectPromptResourceAudioNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/architect/prompts/{promptId}/resources/{languageCode}/audio][%d] deleteArchitectPromptResourceAudioNotFound  %+v", 404, o.Payload)
}

func (o *DeleteArchitectPromptResourceAudioNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteArchitectPromptResourceAudioNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteArchitectPromptResourceAudioRequestEntityTooLarge creates a DeleteArchitectPromptResourceAudioRequestEntityTooLarge with default headers values
func NewDeleteArchitectPromptResourceAudioRequestEntityTooLarge() *DeleteArchitectPromptResourceAudioRequestEntityTooLarge {
	return &DeleteArchitectPromptResourceAudioRequestEntityTooLarge{}
}

/*DeleteArchitectPromptResourceAudioRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteArchitectPromptResourceAudioRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteArchitectPromptResourceAudioRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/architect/prompts/{promptId}/resources/{languageCode}/audio][%d] deleteArchitectPromptResourceAudioRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteArchitectPromptResourceAudioRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteArchitectPromptResourceAudioRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteArchitectPromptResourceAudioUnsupportedMediaType creates a DeleteArchitectPromptResourceAudioUnsupportedMediaType with default headers values
func NewDeleteArchitectPromptResourceAudioUnsupportedMediaType() *DeleteArchitectPromptResourceAudioUnsupportedMediaType {
	return &DeleteArchitectPromptResourceAudioUnsupportedMediaType{}
}

/*DeleteArchitectPromptResourceAudioUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteArchitectPromptResourceAudioUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteArchitectPromptResourceAudioUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/architect/prompts/{promptId}/resources/{languageCode}/audio][%d] deleteArchitectPromptResourceAudioUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteArchitectPromptResourceAudioUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteArchitectPromptResourceAudioUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteArchitectPromptResourceAudioTooManyRequests creates a DeleteArchitectPromptResourceAudioTooManyRequests with default headers values
func NewDeleteArchitectPromptResourceAudioTooManyRequests() *DeleteArchitectPromptResourceAudioTooManyRequests {
	return &DeleteArchitectPromptResourceAudioTooManyRequests{}
}

/*DeleteArchitectPromptResourceAudioTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteArchitectPromptResourceAudioTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteArchitectPromptResourceAudioTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/architect/prompts/{promptId}/resources/{languageCode}/audio][%d] deleteArchitectPromptResourceAudioTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteArchitectPromptResourceAudioTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteArchitectPromptResourceAudioTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteArchitectPromptResourceAudioInternalServerError creates a DeleteArchitectPromptResourceAudioInternalServerError with default headers values
func NewDeleteArchitectPromptResourceAudioInternalServerError() *DeleteArchitectPromptResourceAudioInternalServerError {
	return &DeleteArchitectPromptResourceAudioInternalServerError{}
}

/*DeleteArchitectPromptResourceAudioInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteArchitectPromptResourceAudioInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteArchitectPromptResourceAudioInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/architect/prompts/{promptId}/resources/{languageCode}/audio][%d] deleteArchitectPromptResourceAudioInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteArchitectPromptResourceAudioInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteArchitectPromptResourceAudioInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteArchitectPromptResourceAudioServiceUnavailable creates a DeleteArchitectPromptResourceAudioServiceUnavailable with default headers values
func NewDeleteArchitectPromptResourceAudioServiceUnavailable() *DeleteArchitectPromptResourceAudioServiceUnavailable {
	return &DeleteArchitectPromptResourceAudioServiceUnavailable{}
}

/*DeleteArchitectPromptResourceAudioServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteArchitectPromptResourceAudioServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteArchitectPromptResourceAudioServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/architect/prompts/{promptId}/resources/{languageCode}/audio][%d] deleteArchitectPromptResourceAudioServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteArchitectPromptResourceAudioServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteArchitectPromptResourceAudioServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteArchitectPromptResourceAudioGatewayTimeout creates a DeleteArchitectPromptResourceAudioGatewayTimeout with default headers values
func NewDeleteArchitectPromptResourceAudioGatewayTimeout() *DeleteArchitectPromptResourceAudioGatewayTimeout {
	return &DeleteArchitectPromptResourceAudioGatewayTimeout{}
}

/*DeleteArchitectPromptResourceAudioGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteArchitectPromptResourceAudioGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteArchitectPromptResourceAudioGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/architect/prompts/{promptId}/resources/{languageCode}/audio][%d] deleteArchitectPromptResourceAudioGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteArchitectPromptResourceAudioGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteArchitectPromptResourceAudioGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
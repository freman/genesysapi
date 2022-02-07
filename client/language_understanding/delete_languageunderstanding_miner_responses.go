// Code generated by go-swagger; DO NOT EDIT.

package language_understanding

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeleteLanguageunderstandingMinerReader is a Reader for the DeleteLanguageunderstandingMiner structure.
type DeleteLanguageunderstandingMinerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLanguageunderstandingMinerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteLanguageunderstandingMinerNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteLanguageunderstandingMinerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteLanguageunderstandingMinerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteLanguageunderstandingMinerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteLanguageunderstandingMinerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteLanguageunderstandingMinerRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteLanguageunderstandingMinerRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteLanguageunderstandingMinerUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteLanguageunderstandingMinerTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteLanguageunderstandingMinerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteLanguageunderstandingMinerServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteLanguageunderstandingMinerGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteLanguageunderstandingMinerNoContent creates a DeleteLanguageunderstandingMinerNoContent with default headers values
func NewDeleteLanguageunderstandingMinerNoContent() *DeleteLanguageunderstandingMinerNoContent {
	return &DeleteLanguageunderstandingMinerNoContent{}
}

/*DeleteLanguageunderstandingMinerNoContent handles this case with default header values.

Miner deleted
*/
type DeleteLanguageunderstandingMinerNoContent struct {
}

func (o *DeleteLanguageunderstandingMinerNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/languageunderstanding/miners/{minerId}][%d] deleteLanguageunderstandingMinerNoContent ", 204)
}

func (o *DeleteLanguageunderstandingMinerNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLanguageunderstandingMinerBadRequest creates a DeleteLanguageunderstandingMinerBadRequest with default headers values
func NewDeleteLanguageunderstandingMinerBadRequest() *DeleteLanguageunderstandingMinerBadRequest {
	return &DeleteLanguageunderstandingMinerBadRequest{}
}

/*DeleteLanguageunderstandingMinerBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteLanguageunderstandingMinerBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteLanguageunderstandingMinerBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/languageunderstanding/miners/{minerId}][%d] deleteLanguageunderstandingMinerBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteLanguageunderstandingMinerBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLanguageunderstandingMinerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLanguageunderstandingMinerUnauthorized creates a DeleteLanguageunderstandingMinerUnauthorized with default headers values
func NewDeleteLanguageunderstandingMinerUnauthorized() *DeleteLanguageunderstandingMinerUnauthorized {
	return &DeleteLanguageunderstandingMinerUnauthorized{}
}

/*DeleteLanguageunderstandingMinerUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteLanguageunderstandingMinerUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteLanguageunderstandingMinerUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/languageunderstanding/miners/{minerId}][%d] deleteLanguageunderstandingMinerUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteLanguageunderstandingMinerUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLanguageunderstandingMinerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLanguageunderstandingMinerForbidden creates a DeleteLanguageunderstandingMinerForbidden with default headers values
func NewDeleteLanguageunderstandingMinerForbidden() *DeleteLanguageunderstandingMinerForbidden {
	return &DeleteLanguageunderstandingMinerForbidden{}
}

/*DeleteLanguageunderstandingMinerForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteLanguageunderstandingMinerForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteLanguageunderstandingMinerForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/languageunderstanding/miners/{minerId}][%d] deleteLanguageunderstandingMinerForbidden  %+v", 403, o.Payload)
}

func (o *DeleteLanguageunderstandingMinerForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLanguageunderstandingMinerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLanguageunderstandingMinerNotFound creates a DeleteLanguageunderstandingMinerNotFound with default headers values
func NewDeleteLanguageunderstandingMinerNotFound() *DeleteLanguageunderstandingMinerNotFound {
	return &DeleteLanguageunderstandingMinerNotFound{}
}

/*DeleteLanguageunderstandingMinerNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteLanguageunderstandingMinerNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteLanguageunderstandingMinerNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/languageunderstanding/miners/{minerId}][%d] deleteLanguageunderstandingMinerNotFound  %+v", 404, o.Payload)
}

func (o *DeleteLanguageunderstandingMinerNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLanguageunderstandingMinerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLanguageunderstandingMinerRequestTimeout creates a DeleteLanguageunderstandingMinerRequestTimeout with default headers values
func NewDeleteLanguageunderstandingMinerRequestTimeout() *DeleteLanguageunderstandingMinerRequestTimeout {
	return &DeleteLanguageunderstandingMinerRequestTimeout{}
}

/*DeleteLanguageunderstandingMinerRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteLanguageunderstandingMinerRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteLanguageunderstandingMinerRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/languageunderstanding/miners/{minerId}][%d] deleteLanguageunderstandingMinerRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteLanguageunderstandingMinerRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLanguageunderstandingMinerRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLanguageunderstandingMinerRequestEntityTooLarge creates a DeleteLanguageunderstandingMinerRequestEntityTooLarge with default headers values
func NewDeleteLanguageunderstandingMinerRequestEntityTooLarge() *DeleteLanguageunderstandingMinerRequestEntityTooLarge {
	return &DeleteLanguageunderstandingMinerRequestEntityTooLarge{}
}

/*DeleteLanguageunderstandingMinerRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteLanguageunderstandingMinerRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteLanguageunderstandingMinerRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/languageunderstanding/miners/{minerId}][%d] deleteLanguageunderstandingMinerRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteLanguageunderstandingMinerRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLanguageunderstandingMinerRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLanguageunderstandingMinerUnsupportedMediaType creates a DeleteLanguageunderstandingMinerUnsupportedMediaType with default headers values
func NewDeleteLanguageunderstandingMinerUnsupportedMediaType() *DeleteLanguageunderstandingMinerUnsupportedMediaType {
	return &DeleteLanguageunderstandingMinerUnsupportedMediaType{}
}

/*DeleteLanguageunderstandingMinerUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteLanguageunderstandingMinerUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteLanguageunderstandingMinerUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/languageunderstanding/miners/{minerId}][%d] deleteLanguageunderstandingMinerUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteLanguageunderstandingMinerUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLanguageunderstandingMinerUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLanguageunderstandingMinerTooManyRequests creates a DeleteLanguageunderstandingMinerTooManyRequests with default headers values
func NewDeleteLanguageunderstandingMinerTooManyRequests() *DeleteLanguageunderstandingMinerTooManyRequests {
	return &DeleteLanguageunderstandingMinerTooManyRequests{}
}

/*DeleteLanguageunderstandingMinerTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteLanguageunderstandingMinerTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteLanguageunderstandingMinerTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/languageunderstanding/miners/{minerId}][%d] deleteLanguageunderstandingMinerTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteLanguageunderstandingMinerTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLanguageunderstandingMinerTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLanguageunderstandingMinerInternalServerError creates a DeleteLanguageunderstandingMinerInternalServerError with default headers values
func NewDeleteLanguageunderstandingMinerInternalServerError() *DeleteLanguageunderstandingMinerInternalServerError {
	return &DeleteLanguageunderstandingMinerInternalServerError{}
}

/*DeleteLanguageunderstandingMinerInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteLanguageunderstandingMinerInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteLanguageunderstandingMinerInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/languageunderstanding/miners/{minerId}][%d] deleteLanguageunderstandingMinerInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteLanguageunderstandingMinerInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLanguageunderstandingMinerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLanguageunderstandingMinerServiceUnavailable creates a DeleteLanguageunderstandingMinerServiceUnavailable with default headers values
func NewDeleteLanguageunderstandingMinerServiceUnavailable() *DeleteLanguageunderstandingMinerServiceUnavailable {
	return &DeleteLanguageunderstandingMinerServiceUnavailable{}
}

/*DeleteLanguageunderstandingMinerServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteLanguageunderstandingMinerServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteLanguageunderstandingMinerServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/languageunderstanding/miners/{minerId}][%d] deleteLanguageunderstandingMinerServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteLanguageunderstandingMinerServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLanguageunderstandingMinerServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLanguageunderstandingMinerGatewayTimeout creates a DeleteLanguageunderstandingMinerGatewayTimeout with default headers values
func NewDeleteLanguageunderstandingMinerGatewayTimeout() *DeleteLanguageunderstandingMinerGatewayTimeout {
	return &DeleteLanguageunderstandingMinerGatewayTimeout{}
}

/*DeleteLanguageunderstandingMinerGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteLanguageunderstandingMinerGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteLanguageunderstandingMinerGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/languageunderstanding/miners/{minerId}][%d] deleteLanguageunderstandingMinerGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteLanguageunderstandingMinerGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLanguageunderstandingMinerGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
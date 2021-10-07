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

// GetLanguageunderstandingMinerIntentReader is a Reader for the GetLanguageunderstandingMinerIntent structure.
type GetLanguageunderstandingMinerIntentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLanguageunderstandingMinerIntentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLanguageunderstandingMinerIntentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLanguageunderstandingMinerIntentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLanguageunderstandingMinerIntentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLanguageunderstandingMinerIntentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLanguageunderstandingMinerIntentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetLanguageunderstandingMinerIntentRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetLanguageunderstandingMinerIntentRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetLanguageunderstandingMinerIntentUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetLanguageunderstandingMinerIntentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLanguageunderstandingMinerIntentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetLanguageunderstandingMinerIntentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetLanguageunderstandingMinerIntentGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLanguageunderstandingMinerIntentOK creates a GetLanguageunderstandingMinerIntentOK with default headers values
func NewGetLanguageunderstandingMinerIntentOK() *GetLanguageunderstandingMinerIntentOK {
	return &GetLanguageunderstandingMinerIntentOK{}
}

/*GetLanguageunderstandingMinerIntentOK handles this case with default header values.

successful operation
*/
type GetLanguageunderstandingMinerIntentOK struct {
	Payload *models.MinerIntent
}

func (o *GetLanguageunderstandingMinerIntentOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/intents/{intentId}][%d] getLanguageunderstandingMinerIntentOK  %+v", 200, o.Payload)
}

func (o *GetLanguageunderstandingMinerIntentOK) GetPayload() *models.MinerIntent {
	return o.Payload
}

func (o *GetLanguageunderstandingMinerIntentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MinerIntent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinerIntentBadRequest creates a GetLanguageunderstandingMinerIntentBadRequest with default headers values
func NewGetLanguageunderstandingMinerIntentBadRequest() *GetLanguageunderstandingMinerIntentBadRequest {
	return &GetLanguageunderstandingMinerIntentBadRequest{}
}

/*GetLanguageunderstandingMinerIntentBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetLanguageunderstandingMinerIntentBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingMinerIntentBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/intents/{intentId}][%d] getLanguageunderstandingMinerIntentBadRequest  %+v", 400, o.Payload)
}

func (o *GetLanguageunderstandingMinerIntentBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinerIntentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinerIntentUnauthorized creates a GetLanguageunderstandingMinerIntentUnauthorized with default headers values
func NewGetLanguageunderstandingMinerIntentUnauthorized() *GetLanguageunderstandingMinerIntentUnauthorized {
	return &GetLanguageunderstandingMinerIntentUnauthorized{}
}

/*GetLanguageunderstandingMinerIntentUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetLanguageunderstandingMinerIntentUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingMinerIntentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/intents/{intentId}][%d] getLanguageunderstandingMinerIntentUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLanguageunderstandingMinerIntentUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinerIntentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinerIntentForbidden creates a GetLanguageunderstandingMinerIntentForbidden with default headers values
func NewGetLanguageunderstandingMinerIntentForbidden() *GetLanguageunderstandingMinerIntentForbidden {
	return &GetLanguageunderstandingMinerIntentForbidden{}
}

/*GetLanguageunderstandingMinerIntentForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetLanguageunderstandingMinerIntentForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingMinerIntentForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/intents/{intentId}][%d] getLanguageunderstandingMinerIntentForbidden  %+v", 403, o.Payload)
}

func (o *GetLanguageunderstandingMinerIntentForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinerIntentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinerIntentNotFound creates a GetLanguageunderstandingMinerIntentNotFound with default headers values
func NewGetLanguageunderstandingMinerIntentNotFound() *GetLanguageunderstandingMinerIntentNotFound {
	return &GetLanguageunderstandingMinerIntentNotFound{}
}

/*GetLanguageunderstandingMinerIntentNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetLanguageunderstandingMinerIntentNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingMinerIntentNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/intents/{intentId}][%d] getLanguageunderstandingMinerIntentNotFound  %+v", 404, o.Payload)
}

func (o *GetLanguageunderstandingMinerIntentNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinerIntentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinerIntentRequestTimeout creates a GetLanguageunderstandingMinerIntentRequestTimeout with default headers values
func NewGetLanguageunderstandingMinerIntentRequestTimeout() *GetLanguageunderstandingMinerIntentRequestTimeout {
	return &GetLanguageunderstandingMinerIntentRequestTimeout{}
}

/*GetLanguageunderstandingMinerIntentRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetLanguageunderstandingMinerIntentRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingMinerIntentRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/intents/{intentId}][%d] getLanguageunderstandingMinerIntentRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetLanguageunderstandingMinerIntentRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinerIntentRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinerIntentRequestEntityTooLarge creates a GetLanguageunderstandingMinerIntentRequestEntityTooLarge with default headers values
func NewGetLanguageunderstandingMinerIntentRequestEntityTooLarge() *GetLanguageunderstandingMinerIntentRequestEntityTooLarge {
	return &GetLanguageunderstandingMinerIntentRequestEntityTooLarge{}
}

/*GetLanguageunderstandingMinerIntentRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetLanguageunderstandingMinerIntentRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingMinerIntentRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/intents/{intentId}][%d] getLanguageunderstandingMinerIntentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetLanguageunderstandingMinerIntentRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinerIntentRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinerIntentUnsupportedMediaType creates a GetLanguageunderstandingMinerIntentUnsupportedMediaType with default headers values
func NewGetLanguageunderstandingMinerIntentUnsupportedMediaType() *GetLanguageunderstandingMinerIntentUnsupportedMediaType {
	return &GetLanguageunderstandingMinerIntentUnsupportedMediaType{}
}

/*GetLanguageunderstandingMinerIntentUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetLanguageunderstandingMinerIntentUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingMinerIntentUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/intents/{intentId}][%d] getLanguageunderstandingMinerIntentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetLanguageunderstandingMinerIntentUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinerIntentUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinerIntentTooManyRequests creates a GetLanguageunderstandingMinerIntentTooManyRequests with default headers values
func NewGetLanguageunderstandingMinerIntentTooManyRequests() *GetLanguageunderstandingMinerIntentTooManyRequests {
	return &GetLanguageunderstandingMinerIntentTooManyRequests{}
}

/*GetLanguageunderstandingMinerIntentTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetLanguageunderstandingMinerIntentTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingMinerIntentTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/intents/{intentId}][%d] getLanguageunderstandingMinerIntentTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLanguageunderstandingMinerIntentTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinerIntentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinerIntentInternalServerError creates a GetLanguageunderstandingMinerIntentInternalServerError with default headers values
func NewGetLanguageunderstandingMinerIntentInternalServerError() *GetLanguageunderstandingMinerIntentInternalServerError {
	return &GetLanguageunderstandingMinerIntentInternalServerError{}
}

/*GetLanguageunderstandingMinerIntentInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetLanguageunderstandingMinerIntentInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingMinerIntentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/intents/{intentId}][%d] getLanguageunderstandingMinerIntentInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLanguageunderstandingMinerIntentInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinerIntentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinerIntentServiceUnavailable creates a GetLanguageunderstandingMinerIntentServiceUnavailable with default headers values
func NewGetLanguageunderstandingMinerIntentServiceUnavailable() *GetLanguageunderstandingMinerIntentServiceUnavailable {
	return &GetLanguageunderstandingMinerIntentServiceUnavailable{}
}

/*GetLanguageunderstandingMinerIntentServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetLanguageunderstandingMinerIntentServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingMinerIntentServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/intents/{intentId}][%d] getLanguageunderstandingMinerIntentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLanguageunderstandingMinerIntentServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinerIntentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinerIntentGatewayTimeout creates a GetLanguageunderstandingMinerIntentGatewayTimeout with default headers values
func NewGetLanguageunderstandingMinerIntentGatewayTimeout() *GetLanguageunderstandingMinerIntentGatewayTimeout {
	return &GetLanguageunderstandingMinerIntentGatewayTimeout{}
}

/*GetLanguageunderstandingMinerIntentGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetLanguageunderstandingMinerIntentGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingMinerIntentGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/intents/{intentId}][%d] getLanguageunderstandingMinerIntentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetLanguageunderstandingMinerIntentGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinerIntentGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

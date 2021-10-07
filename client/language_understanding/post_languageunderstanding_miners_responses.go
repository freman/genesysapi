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

// PostLanguageunderstandingMinersReader is a Reader for the PostLanguageunderstandingMiners structure.
type PostLanguageunderstandingMinersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLanguageunderstandingMinersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLanguageunderstandingMinersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewPostLanguageunderstandingMinersCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostLanguageunderstandingMinersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostLanguageunderstandingMinersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostLanguageunderstandingMinersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostLanguageunderstandingMinersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostLanguageunderstandingMinersRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostLanguageunderstandingMinersRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostLanguageunderstandingMinersUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostLanguageunderstandingMinersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostLanguageunderstandingMinersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostLanguageunderstandingMinersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostLanguageunderstandingMinersGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostLanguageunderstandingMinersOK creates a PostLanguageunderstandingMinersOK with default headers values
func NewPostLanguageunderstandingMinersOK() *PostLanguageunderstandingMinersOK {
	return &PostLanguageunderstandingMinersOK{}
}

/*PostLanguageunderstandingMinersOK handles this case with default header values.

successful operation
*/
type PostLanguageunderstandingMinersOK struct {
	Payload *models.Miner
}

func (o *PostLanguageunderstandingMinersOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners][%d] postLanguageunderstandingMinersOK  %+v", 200, o.Payload)
}

func (o *PostLanguageunderstandingMinersOK) GetPayload() *models.Miner {
	return o.Payload
}

func (o *PostLanguageunderstandingMinersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Miner)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinersCreated creates a PostLanguageunderstandingMinersCreated with default headers values
func NewPostLanguageunderstandingMinersCreated() *PostLanguageunderstandingMinersCreated {
	return &PostLanguageunderstandingMinersCreated{}
}

/*PostLanguageunderstandingMinersCreated handles this case with default header values.

Miner created successfully
*/
type PostLanguageunderstandingMinersCreated struct {
	Payload *models.Miner
}

func (o *PostLanguageunderstandingMinersCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners][%d] postLanguageunderstandingMinersCreated  %+v", 201, o.Payload)
}

func (o *PostLanguageunderstandingMinersCreated) GetPayload() *models.Miner {
	return o.Payload
}

func (o *PostLanguageunderstandingMinersCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Miner)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinersBadRequest creates a PostLanguageunderstandingMinersBadRequest with default headers values
func NewPostLanguageunderstandingMinersBadRequest() *PostLanguageunderstandingMinersBadRequest {
	return &PostLanguageunderstandingMinersBadRequest{}
}

/*PostLanguageunderstandingMinersBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostLanguageunderstandingMinersBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostLanguageunderstandingMinersBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners][%d] postLanguageunderstandingMinersBadRequest  %+v", 400, o.Payload)
}

func (o *PostLanguageunderstandingMinersBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinersUnauthorized creates a PostLanguageunderstandingMinersUnauthorized with default headers values
func NewPostLanguageunderstandingMinersUnauthorized() *PostLanguageunderstandingMinersUnauthorized {
	return &PostLanguageunderstandingMinersUnauthorized{}
}

/*PostLanguageunderstandingMinersUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostLanguageunderstandingMinersUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostLanguageunderstandingMinersUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners][%d] postLanguageunderstandingMinersUnauthorized  %+v", 401, o.Payload)
}

func (o *PostLanguageunderstandingMinersUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinersForbidden creates a PostLanguageunderstandingMinersForbidden with default headers values
func NewPostLanguageunderstandingMinersForbidden() *PostLanguageunderstandingMinersForbidden {
	return &PostLanguageunderstandingMinersForbidden{}
}

/*PostLanguageunderstandingMinersForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostLanguageunderstandingMinersForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostLanguageunderstandingMinersForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners][%d] postLanguageunderstandingMinersForbidden  %+v", 403, o.Payload)
}

func (o *PostLanguageunderstandingMinersForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinersNotFound creates a PostLanguageunderstandingMinersNotFound with default headers values
func NewPostLanguageunderstandingMinersNotFound() *PostLanguageunderstandingMinersNotFound {
	return &PostLanguageunderstandingMinersNotFound{}
}

/*PostLanguageunderstandingMinersNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostLanguageunderstandingMinersNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostLanguageunderstandingMinersNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners][%d] postLanguageunderstandingMinersNotFound  %+v", 404, o.Payload)
}

func (o *PostLanguageunderstandingMinersNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinersRequestTimeout creates a PostLanguageunderstandingMinersRequestTimeout with default headers values
func NewPostLanguageunderstandingMinersRequestTimeout() *PostLanguageunderstandingMinersRequestTimeout {
	return &PostLanguageunderstandingMinersRequestTimeout{}
}

/*PostLanguageunderstandingMinersRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostLanguageunderstandingMinersRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostLanguageunderstandingMinersRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners][%d] postLanguageunderstandingMinersRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostLanguageunderstandingMinersRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinersRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinersRequestEntityTooLarge creates a PostLanguageunderstandingMinersRequestEntityTooLarge with default headers values
func NewPostLanguageunderstandingMinersRequestEntityTooLarge() *PostLanguageunderstandingMinersRequestEntityTooLarge {
	return &PostLanguageunderstandingMinersRequestEntityTooLarge{}
}

/*PostLanguageunderstandingMinersRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostLanguageunderstandingMinersRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostLanguageunderstandingMinersRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners][%d] postLanguageunderstandingMinersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostLanguageunderstandingMinersRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinersRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinersUnsupportedMediaType creates a PostLanguageunderstandingMinersUnsupportedMediaType with default headers values
func NewPostLanguageunderstandingMinersUnsupportedMediaType() *PostLanguageunderstandingMinersUnsupportedMediaType {
	return &PostLanguageunderstandingMinersUnsupportedMediaType{}
}

/*PostLanguageunderstandingMinersUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostLanguageunderstandingMinersUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostLanguageunderstandingMinersUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners][%d] postLanguageunderstandingMinersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostLanguageunderstandingMinersUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinersUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinersTooManyRequests creates a PostLanguageunderstandingMinersTooManyRequests with default headers values
func NewPostLanguageunderstandingMinersTooManyRequests() *PostLanguageunderstandingMinersTooManyRequests {
	return &PostLanguageunderstandingMinersTooManyRequests{}
}

/*PostLanguageunderstandingMinersTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostLanguageunderstandingMinersTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostLanguageunderstandingMinersTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners][%d] postLanguageunderstandingMinersTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostLanguageunderstandingMinersTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinersInternalServerError creates a PostLanguageunderstandingMinersInternalServerError with default headers values
func NewPostLanguageunderstandingMinersInternalServerError() *PostLanguageunderstandingMinersInternalServerError {
	return &PostLanguageunderstandingMinersInternalServerError{}
}

/*PostLanguageunderstandingMinersInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostLanguageunderstandingMinersInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostLanguageunderstandingMinersInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners][%d] postLanguageunderstandingMinersInternalServerError  %+v", 500, o.Payload)
}

func (o *PostLanguageunderstandingMinersInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinersServiceUnavailable creates a PostLanguageunderstandingMinersServiceUnavailable with default headers values
func NewPostLanguageunderstandingMinersServiceUnavailable() *PostLanguageunderstandingMinersServiceUnavailable {
	return &PostLanguageunderstandingMinersServiceUnavailable{}
}

/*PostLanguageunderstandingMinersServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostLanguageunderstandingMinersServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostLanguageunderstandingMinersServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners][%d] postLanguageunderstandingMinersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostLanguageunderstandingMinersServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinersGatewayTimeout creates a PostLanguageunderstandingMinersGatewayTimeout with default headers values
func NewPostLanguageunderstandingMinersGatewayTimeout() *PostLanguageunderstandingMinersGatewayTimeout {
	return &PostLanguageunderstandingMinersGatewayTimeout{}
}

/*PostLanguageunderstandingMinersGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostLanguageunderstandingMinersGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostLanguageunderstandingMinersGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners][%d] postLanguageunderstandingMinersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostLanguageunderstandingMinersGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinersGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

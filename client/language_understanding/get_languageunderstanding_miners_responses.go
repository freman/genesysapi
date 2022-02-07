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

// GetLanguageunderstandingMinersReader is a Reader for the GetLanguageunderstandingMiners structure.
type GetLanguageunderstandingMinersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLanguageunderstandingMinersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLanguageunderstandingMinersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLanguageunderstandingMinersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLanguageunderstandingMinersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLanguageunderstandingMinersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLanguageunderstandingMinersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetLanguageunderstandingMinersRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetLanguageunderstandingMinersRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetLanguageunderstandingMinersUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetLanguageunderstandingMinersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLanguageunderstandingMinersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetLanguageunderstandingMinersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetLanguageunderstandingMinersGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLanguageunderstandingMinersOK creates a GetLanguageunderstandingMinersOK with default headers values
func NewGetLanguageunderstandingMinersOK() *GetLanguageunderstandingMinersOK {
	return &GetLanguageunderstandingMinersOK{}
}

/*GetLanguageunderstandingMinersOK handles this case with default header values.

successful operation
*/
type GetLanguageunderstandingMinersOK struct {
	Payload *models.MinerListing
}

func (o *GetLanguageunderstandingMinersOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners][%d] getLanguageunderstandingMinersOK  %+v", 200, o.Payload)
}

func (o *GetLanguageunderstandingMinersOK) GetPayload() *models.MinerListing {
	return o.Payload
}

func (o *GetLanguageunderstandingMinersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MinerListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinersBadRequest creates a GetLanguageunderstandingMinersBadRequest with default headers values
func NewGetLanguageunderstandingMinersBadRequest() *GetLanguageunderstandingMinersBadRequest {
	return &GetLanguageunderstandingMinersBadRequest{}
}

/*GetLanguageunderstandingMinersBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetLanguageunderstandingMinersBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingMinersBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners][%d] getLanguageunderstandingMinersBadRequest  %+v", 400, o.Payload)
}

func (o *GetLanguageunderstandingMinersBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinersUnauthorized creates a GetLanguageunderstandingMinersUnauthorized with default headers values
func NewGetLanguageunderstandingMinersUnauthorized() *GetLanguageunderstandingMinersUnauthorized {
	return &GetLanguageunderstandingMinersUnauthorized{}
}

/*GetLanguageunderstandingMinersUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetLanguageunderstandingMinersUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingMinersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners][%d] getLanguageunderstandingMinersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLanguageunderstandingMinersUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinersForbidden creates a GetLanguageunderstandingMinersForbidden with default headers values
func NewGetLanguageunderstandingMinersForbidden() *GetLanguageunderstandingMinersForbidden {
	return &GetLanguageunderstandingMinersForbidden{}
}

/*GetLanguageunderstandingMinersForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetLanguageunderstandingMinersForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingMinersForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners][%d] getLanguageunderstandingMinersForbidden  %+v", 403, o.Payload)
}

func (o *GetLanguageunderstandingMinersForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinersNotFound creates a GetLanguageunderstandingMinersNotFound with default headers values
func NewGetLanguageunderstandingMinersNotFound() *GetLanguageunderstandingMinersNotFound {
	return &GetLanguageunderstandingMinersNotFound{}
}

/*GetLanguageunderstandingMinersNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetLanguageunderstandingMinersNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingMinersNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners][%d] getLanguageunderstandingMinersNotFound  %+v", 404, o.Payload)
}

func (o *GetLanguageunderstandingMinersNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinersRequestTimeout creates a GetLanguageunderstandingMinersRequestTimeout with default headers values
func NewGetLanguageunderstandingMinersRequestTimeout() *GetLanguageunderstandingMinersRequestTimeout {
	return &GetLanguageunderstandingMinersRequestTimeout{}
}

/*GetLanguageunderstandingMinersRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetLanguageunderstandingMinersRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingMinersRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners][%d] getLanguageunderstandingMinersRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetLanguageunderstandingMinersRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinersRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinersRequestEntityTooLarge creates a GetLanguageunderstandingMinersRequestEntityTooLarge with default headers values
func NewGetLanguageunderstandingMinersRequestEntityTooLarge() *GetLanguageunderstandingMinersRequestEntityTooLarge {
	return &GetLanguageunderstandingMinersRequestEntityTooLarge{}
}

/*GetLanguageunderstandingMinersRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetLanguageunderstandingMinersRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingMinersRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners][%d] getLanguageunderstandingMinersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetLanguageunderstandingMinersRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinersRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinersUnsupportedMediaType creates a GetLanguageunderstandingMinersUnsupportedMediaType with default headers values
func NewGetLanguageunderstandingMinersUnsupportedMediaType() *GetLanguageunderstandingMinersUnsupportedMediaType {
	return &GetLanguageunderstandingMinersUnsupportedMediaType{}
}

/*GetLanguageunderstandingMinersUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetLanguageunderstandingMinersUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingMinersUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners][%d] getLanguageunderstandingMinersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetLanguageunderstandingMinersUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinersUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinersTooManyRequests creates a GetLanguageunderstandingMinersTooManyRequests with default headers values
func NewGetLanguageunderstandingMinersTooManyRequests() *GetLanguageunderstandingMinersTooManyRequests {
	return &GetLanguageunderstandingMinersTooManyRequests{}
}

/*GetLanguageunderstandingMinersTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetLanguageunderstandingMinersTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingMinersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners][%d] getLanguageunderstandingMinersTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLanguageunderstandingMinersTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinersInternalServerError creates a GetLanguageunderstandingMinersInternalServerError with default headers values
func NewGetLanguageunderstandingMinersInternalServerError() *GetLanguageunderstandingMinersInternalServerError {
	return &GetLanguageunderstandingMinersInternalServerError{}
}

/*GetLanguageunderstandingMinersInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetLanguageunderstandingMinersInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingMinersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners][%d] getLanguageunderstandingMinersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLanguageunderstandingMinersInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinersServiceUnavailable creates a GetLanguageunderstandingMinersServiceUnavailable with default headers values
func NewGetLanguageunderstandingMinersServiceUnavailable() *GetLanguageunderstandingMinersServiceUnavailable {
	return &GetLanguageunderstandingMinersServiceUnavailable{}
}

/*GetLanguageunderstandingMinersServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetLanguageunderstandingMinersServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingMinersServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners][%d] getLanguageunderstandingMinersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLanguageunderstandingMinersServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinersGatewayTimeout creates a GetLanguageunderstandingMinersGatewayTimeout with default headers values
func NewGetLanguageunderstandingMinersGatewayTimeout() *GetLanguageunderstandingMinersGatewayTimeout {
	return &GetLanguageunderstandingMinersGatewayTimeout{}
}

/*GetLanguageunderstandingMinersGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetLanguageunderstandingMinersGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingMinersGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners][%d] getLanguageunderstandingMinersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetLanguageunderstandingMinersGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinersGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
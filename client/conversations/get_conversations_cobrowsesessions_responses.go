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

// GetConversationsCobrowsesessionsReader is a Reader for the GetConversationsCobrowsesessions structure.
type GetConversationsCobrowsesessionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationsCobrowsesessionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationsCobrowsesessionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationsCobrowsesessionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationsCobrowsesessionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationsCobrowsesessionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationsCobrowsesessionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetConversationsCobrowsesessionsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationsCobrowsesessionsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationsCobrowsesessionsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationsCobrowsesessionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationsCobrowsesessionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationsCobrowsesessionsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationsCobrowsesessionsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationsCobrowsesessionsOK creates a GetConversationsCobrowsesessionsOK with default headers values
func NewGetConversationsCobrowsesessionsOK() *GetConversationsCobrowsesessionsOK {
	return &GetConversationsCobrowsesessionsOK{}
}

/*GetConversationsCobrowsesessionsOK handles this case with default header values.

successful operation
*/
type GetConversationsCobrowsesessionsOK struct {
	Payload *models.CobrowseConversationEntityListing
}

func (o *GetConversationsCobrowsesessionsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/cobrowsesessions][%d] getConversationsCobrowsesessionsOK  %+v", 200, o.Payload)
}

func (o *GetConversationsCobrowsesessionsOK) GetPayload() *models.CobrowseConversationEntityListing {
	return o.Payload
}

func (o *GetConversationsCobrowsesessionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CobrowseConversationEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCobrowsesessionsBadRequest creates a GetConversationsCobrowsesessionsBadRequest with default headers values
func NewGetConversationsCobrowsesessionsBadRequest() *GetConversationsCobrowsesessionsBadRequest {
	return &GetConversationsCobrowsesessionsBadRequest{}
}

/*GetConversationsCobrowsesessionsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationsCobrowsesessionsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCobrowsesessionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/cobrowsesessions][%d] getConversationsCobrowsesessionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsCobrowsesessionsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCobrowsesessionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCobrowsesessionsUnauthorized creates a GetConversationsCobrowsesessionsUnauthorized with default headers values
func NewGetConversationsCobrowsesessionsUnauthorized() *GetConversationsCobrowsesessionsUnauthorized {
	return &GetConversationsCobrowsesessionsUnauthorized{}
}

/*GetConversationsCobrowsesessionsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationsCobrowsesessionsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCobrowsesessionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/cobrowsesessions][%d] getConversationsCobrowsesessionsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsCobrowsesessionsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCobrowsesessionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCobrowsesessionsForbidden creates a GetConversationsCobrowsesessionsForbidden with default headers values
func NewGetConversationsCobrowsesessionsForbidden() *GetConversationsCobrowsesessionsForbidden {
	return &GetConversationsCobrowsesessionsForbidden{}
}

/*GetConversationsCobrowsesessionsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationsCobrowsesessionsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCobrowsesessionsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/cobrowsesessions][%d] getConversationsCobrowsesessionsForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsCobrowsesessionsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCobrowsesessionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCobrowsesessionsNotFound creates a GetConversationsCobrowsesessionsNotFound with default headers values
func NewGetConversationsCobrowsesessionsNotFound() *GetConversationsCobrowsesessionsNotFound {
	return &GetConversationsCobrowsesessionsNotFound{}
}

/*GetConversationsCobrowsesessionsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetConversationsCobrowsesessionsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCobrowsesessionsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/cobrowsesessions][%d] getConversationsCobrowsesessionsNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsCobrowsesessionsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCobrowsesessionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCobrowsesessionsRequestTimeout creates a GetConversationsCobrowsesessionsRequestTimeout with default headers values
func NewGetConversationsCobrowsesessionsRequestTimeout() *GetConversationsCobrowsesessionsRequestTimeout {
	return &GetConversationsCobrowsesessionsRequestTimeout{}
}

/*GetConversationsCobrowsesessionsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetConversationsCobrowsesessionsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCobrowsesessionsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/cobrowsesessions][%d] getConversationsCobrowsesessionsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsCobrowsesessionsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCobrowsesessionsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCobrowsesessionsRequestEntityTooLarge creates a GetConversationsCobrowsesessionsRequestEntityTooLarge with default headers values
func NewGetConversationsCobrowsesessionsRequestEntityTooLarge() *GetConversationsCobrowsesessionsRequestEntityTooLarge {
	return &GetConversationsCobrowsesessionsRequestEntityTooLarge{}
}

/*GetConversationsCobrowsesessionsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetConversationsCobrowsesessionsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCobrowsesessionsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/cobrowsesessions][%d] getConversationsCobrowsesessionsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsCobrowsesessionsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCobrowsesessionsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCobrowsesessionsUnsupportedMediaType creates a GetConversationsCobrowsesessionsUnsupportedMediaType with default headers values
func NewGetConversationsCobrowsesessionsUnsupportedMediaType() *GetConversationsCobrowsesessionsUnsupportedMediaType {
	return &GetConversationsCobrowsesessionsUnsupportedMediaType{}
}

/*GetConversationsCobrowsesessionsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationsCobrowsesessionsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCobrowsesessionsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/cobrowsesessions][%d] getConversationsCobrowsesessionsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsCobrowsesessionsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCobrowsesessionsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCobrowsesessionsTooManyRequests creates a GetConversationsCobrowsesessionsTooManyRequests with default headers values
func NewGetConversationsCobrowsesessionsTooManyRequests() *GetConversationsCobrowsesessionsTooManyRequests {
	return &GetConversationsCobrowsesessionsTooManyRequests{}
}

/*GetConversationsCobrowsesessionsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetConversationsCobrowsesessionsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCobrowsesessionsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/cobrowsesessions][%d] getConversationsCobrowsesessionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsCobrowsesessionsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCobrowsesessionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCobrowsesessionsInternalServerError creates a GetConversationsCobrowsesessionsInternalServerError with default headers values
func NewGetConversationsCobrowsesessionsInternalServerError() *GetConversationsCobrowsesessionsInternalServerError {
	return &GetConversationsCobrowsesessionsInternalServerError{}
}

/*GetConversationsCobrowsesessionsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationsCobrowsesessionsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCobrowsesessionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/cobrowsesessions][%d] getConversationsCobrowsesessionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsCobrowsesessionsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCobrowsesessionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCobrowsesessionsServiceUnavailable creates a GetConversationsCobrowsesessionsServiceUnavailable with default headers values
func NewGetConversationsCobrowsesessionsServiceUnavailable() *GetConversationsCobrowsesessionsServiceUnavailable {
	return &GetConversationsCobrowsesessionsServiceUnavailable{}
}

/*GetConversationsCobrowsesessionsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationsCobrowsesessionsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCobrowsesessionsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/cobrowsesessions][%d] getConversationsCobrowsesessionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsCobrowsesessionsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCobrowsesessionsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCobrowsesessionsGatewayTimeout creates a GetConversationsCobrowsesessionsGatewayTimeout with default headers values
func NewGetConversationsCobrowsesessionsGatewayTimeout() *GetConversationsCobrowsesessionsGatewayTimeout {
	return &GetConversationsCobrowsesessionsGatewayTimeout{}
}

/*GetConversationsCobrowsesessionsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetConversationsCobrowsesessionsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCobrowsesessionsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/cobrowsesessions][%d] getConversationsCobrowsesessionsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsCobrowsesessionsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCobrowsesessionsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// GetConversationsCallsHistoryReader is a Reader for the GetConversationsCallsHistory structure.
type GetConversationsCallsHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationsCallsHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationsCallsHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationsCallsHistoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationsCallsHistoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationsCallsHistoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationsCallsHistoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetConversationsCallsHistoryRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationsCallsHistoryRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationsCallsHistoryUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationsCallsHistoryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationsCallsHistoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationsCallsHistoryServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationsCallsHistoryGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationsCallsHistoryOK creates a GetConversationsCallsHistoryOK with default headers values
func NewGetConversationsCallsHistoryOK() *GetConversationsCallsHistoryOK {
	return &GetConversationsCallsHistoryOK{}
}

/*GetConversationsCallsHistoryOK handles this case with default header values.

successful operation
*/
type GetConversationsCallsHistoryOK struct {
	Payload *models.CallHistoryConversationEntityListing
}

func (o *GetConversationsCallsHistoryOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/history][%d] getConversationsCallsHistoryOK  %+v", 200, o.Payload)
}

func (o *GetConversationsCallsHistoryOK) GetPayload() *models.CallHistoryConversationEntityListing {
	return o.Payload
}

func (o *GetConversationsCallsHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CallHistoryConversationEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallsHistoryBadRequest creates a GetConversationsCallsHistoryBadRequest with default headers values
func NewGetConversationsCallsHistoryBadRequest() *GetConversationsCallsHistoryBadRequest {
	return &GetConversationsCallsHistoryBadRequest{}
}

/*GetConversationsCallsHistoryBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationsCallsHistoryBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCallsHistoryBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/history][%d] getConversationsCallsHistoryBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsCallsHistoryBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallsHistoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallsHistoryUnauthorized creates a GetConversationsCallsHistoryUnauthorized with default headers values
func NewGetConversationsCallsHistoryUnauthorized() *GetConversationsCallsHistoryUnauthorized {
	return &GetConversationsCallsHistoryUnauthorized{}
}

/*GetConversationsCallsHistoryUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationsCallsHistoryUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCallsHistoryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/history][%d] getConversationsCallsHistoryUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsCallsHistoryUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallsHistoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallsHistoryForbidden creates a GetConversationsCallsHistoryForbidden with default headers values
func NewGetConversationsCallsHistoryForbidden() *GetConversationsCallsHistoryForbidden {
	return &GetConversationsCallsHistoryForbidden{}
}

/*GetConversationsCallsHistoryForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationsCallsHistoryForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCallsHistoryForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/history][%d] getConversationsCallsHistoryForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsCallsHistoryForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallsHistoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallsHistoryNotFound creates a GetConversationsCallsHistoryNotFound with default headers values
func NewGetConversationsCallsHistoryNotFound() *GetConversationsCallsHistoryNotFound {
	return &GetConversationsCallsHistoryNotFound{}
}

/*GetConversationsCallsHistoryNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetConversationsCallsHistoryNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCallsHistoryNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/history][%d] getConversationsCallsHistoryNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsCallsHistoryNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallsHistoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallsHistoryRequestTimeout creates a GetConversationsCallsHistoryRequestTimeout with default headers values
func NewGetConversationsCallsHistoryRequestTimeout() *GetConversationsCallsHistoryRequestTimeout {
	return &GetConversationsCallsHistoryRequestTimeout{}
}

/*GetConversationsCallsHistoryRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetConversationsCallsHistoryRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCallsHistoryRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/history][%d] getConversationsCallsHistoryRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsCallsHistoryRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallsHistoryRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallsHistoryRequestEntityTooLarge creates a GetConversationsCallsHistoryRequestEntityTooLarge with default headers values
func NewGetConversationsCallsHistoryRequestEntityTooLarge() *GetConversationsCallsHistoryRequestEntityTooLarge {
	return &GetConversationsCallsHistoryRequestEntityTooLarge{}
}

/*GetConversationsCallsHistoryRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetConversationsCallsHistoryRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCallsHistoryRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/history][%d] getConversationsCallsHistoryRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsCallsHistoryRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallsHistoryRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallsHistoryUnsupportedMediaType creates a GetConversationsCallsHistoryUnsupportedMediaType with default headers values
func NewGetConversationsCallsHistoryUnsupportedMediaType() *GetConversationsCallsHistoryUnsupportedMediaType {
	return &GetConversationsCallsHistoryUnsupportedMediaType{}
}

/*GetConversationsCallsHistoryUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationsCallsHistoryUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCallsHistoryUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/history][%d] getConversationsCallsHistoryUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsCallsHistoryUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallsHistoryUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallsHistoryTooManyRequests creates a GetConversationsCallsHistoryTooManyRequests with default headers values
func NewGetConversationsCallsHistoryTooManyRequests() *GetConversationsCallsHistoryTooManyRequests {
	return &GetConversationsCallsHistoryTooManyRequests{}
}

/*GetConversationsCallsHistoryTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetConversationsCallsHistoryTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCallsHistoryTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/history][%d] getConversationsCallsHistoryTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsCallsHistoryTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallsHistoryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallsHistoryInternalServerError creates a GetConversationsCallsHistoryInternalServerError with default headers values
func NewGetConversationsCallsHistoryInternalServerError() *GetConversationsCallsHistoryInternalServerError {
	return &GetConversationsCallsHistoryInternalServerError{}
}

/*GetConversationsCallsHistoryInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationsCallsHistoryInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCallsHistoryInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/history][%d] getConversationsCallsHistoryInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsCallsHistoryInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallsHistoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallsHistoryServiceUnavailable creates a GetConversationsCallsHistoryServiceUnavailable with default headers values
func NewGetConversationsCallsHistoryServiceUnavailable() *GetConversationsCallsHistoryServiceUnavailable {
	return &GetConversationsCallsHistoryServiceUnavailable{}
}

/*GetConversationsCallsHistoryServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationsCallsHistoryServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCallsHistoryServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/history][%d] getConversationsCallsHistoryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsCallsHistoryServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallsHistoryServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallsHistoryGatewayTimeout creates a GetConversationsCallsHistoryGatewayTimeout with default headers values
func NewGetConversationsCallsHistoryGatewayTimeout() *GetConversationsCallsHistoryGatewayTimeout {
	return &GetConversationsCallsHistoryGatewayTimeout{}
}

/*GetConversationsCallsHistoryGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetConversationsCallsHistoryGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCallsHistoryGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/history][%d] getConversationsCallsHistoryGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsCallsHistoryGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallsHistoryGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

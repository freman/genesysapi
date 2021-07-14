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

// GetConversationsCallsMaximumconferencepartiesReader is a Reader for the GetConversationsCallsMaximumconferenceparties structure.
type GetConversationsCallsMaximumconferencepartiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationsCallsMaximumconferencepartiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationsCallsMaximumconferencepartiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationsCallsMaximumconferencepartiesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationsCallsMaximumconferencepartiesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationsCallsMaximumconferencepartiesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationsCallsMaximumconferencepartiesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetConversationsCallsMaximumconferencepartiesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationsCallsMaximumconferencepartiesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationsCallsMaximumconferencepartiesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationsCallsMaximumconferencepartiesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationsCallsMaximumconferencepartiesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationsCallsMaximumconferencepartiesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationsCallsMaximumconferencepartiesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationsCallsMaximumconferencepartiesOK creates a GetConversationsCallsMaximumconferencepartiesOK with default headers values
func NewGetConversationsCallsMaximumconferencepartiesOK() *GetConversationsCallsMaximumconferencepartiesOK {
	return &GetConversationsCallsMaximumconferencepartiesOK{}
}

/*GetConversationsCallsMaximumconferencepartiesOK handles this case with default header values.

successful operation
*/
type GetConversationsCallsMaximumconferencepartiesOK struct {
	Payload *models.MaxParticipants
}

func (o *GetConversationsCallsMaximumconferencepartiesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/maximumconferenceparties][%d] getConversationsCallsMaximumconferencepartiesOK  %+v", 200, o.Payload)
}

func (o *GetConversationsCallsMaximumconferencepartiesOK) GetPayload() *models.MaxParticipants {
	return o.Payload
}

func (o *GetConversationsCallsMaximumconferencepartiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MaxParticipants)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallsMaximumconferencepartiesBadRequest creates a GetConversationsCallsMaximumconferencepartiesBadRequest with default headers values
func NewGetConversationsCallsMaximumconferencepartiesBadRequest() *GetConversationsCallsMaximumconferencepartiesBadRequest {
	return &GetConversationsCallsMaximumconferencepartiesBadRequest{}
}

/*GetConversationsCallsMaximumconferencepartiesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationsCallsMaximumconferencepartiesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCallsMaximumconferencepartiesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/maximumconferenceparties][%d] getConversationsCallsMaximumconferencepartiesBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsCallsMaximumconferencepartiesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallsMaximumconferencepartiesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallsMaximumconferencepartiesUnauthorized creates a GetConversationsCallsMaximumconferencepartiesUnauthorized with default headers values
func NewGetConversationsCallsMaximumconferencepartiesUnauthorized() *GetConversationsCallsMaximumconferencepartiesUnauthorized {
	return &GetConversationsCallsMaximumconferencepartiesUnauthorized{}
}

/*GetConversationsCallsMaximumconferencepartiesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationsCallsMaximumconferencepartiesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCallsMaximumconferencepartiesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/maximumconferenceparties][%d] getConversationsCallsMaximumconferencepartiesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsCallsMaximumconferencepartiesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallsMaximumconferencepartiesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallsMaximumconferencepartiesForbidden creates a GetConversationsCallsMaximumconferencepartiesForbidden with default headers values
func NewGetConversationsCallsMaximumconferencepartiesForbidden() *GetConversationsCallsMaximumconferencepartiesForbidden {
	return &GetConversationsCallsMaximumconferencepartiesForbidden{}
}

/*GetConversationsCallsMaximumconferencepartiesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationsCallsMaximumconferencepartiesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCallsMaximumconferencepartiesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/maximumconferenceparties][%d] getConversationsCallsMaximumconferencepartiesForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsCallsMaximumconferencepartiesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallsMaximumconferencepartiesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallsMaximumconferencepartiesNotFound creates a GetConversationsCallsMaximumconferencepartiesNotFound with default headers values
func NewGetConversationsCallsMaximumconferencepartiesNotFound() *GetConversationsCallsMaximumconferencepartiesNotFound {
	return &GetConversationsCallsMaximumconferencepartiesNotFound{}
}

/*GetConversationsCallsMaximumconferencepartiesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetConversationsCallsMaximumconferencepartiesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCallsMaximumconferencepartiesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/maximumconferenceparties][%d] getConversationsCallsMaximumconferencepartiesNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsCallsMaximumconferencepartiesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallsMaximumconferencepartiesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallsMaximumconferencepartiesRequestTimeout creates a GetConversationsCallsMaximumconferencepartiesRequestTimeout with default headers values
func NewGetConversationsCallsMaximumconferencepartiesRequestTimeout() *GetConversationsCallsMaximumconferencepartiesRequestTimeout {
	return &GetConversationsCallsMaximumconferencepartiesRequestTimeout{}
}

/*GetConversationsCallsMaximumconferencepartiesRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetConversationsCallsMaximumconferencepartiesRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCallsMaximumconferencepartiesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/maximumconferenceparties][%d] getConversationsCallsMaximumconferencepartiesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsCallsMaximumconferencepartiesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallsMaximumconferencepartiesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallsMaximumconferencepartiesRequestEntityTooLarge creates a GetConversationsCallsMaximumconferencepartiesRequestEntityTooLarge with default headers values
func NewGetConversationsCallsMaximumconferencepartiesRequestEntityTooLarge() *GetConversationsCallsMaximumconferencepartiesRequestEntityTooLarge {
	return &GetConversationsCallsMaximumconferencepartiesRequestEntityTooLarge{}
}

/*GetConversationsCallsMaximumconferencepartiesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetConversationsCallsMaximumconferencepartiesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCallsMaximumconferencepartiesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/maximumconferenceparties][%d] getConversationsCallsMaximumconferencepartiesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsCallsMaximumconferencepartiesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallsMaximumconferencepartiesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallsMaximumconferencepartiesUnsupportedMediaType creates a GetConversationsCallsMaximumconferencepartiesUnsupportedMediaType with default headers values
func NewGetConversationsCallsMaximumconferencepartiesUnsupportedMediaType() *GetConversationsCallsMaximumconferencepartiesUnsupportedMediaType {
	return &GetConversationsCallsMaximumconferencepartiesUnsupportedMediaType{}
}

/*GetConversationsCallsMaximumconferencepartiesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationsCallsMaximumconferencepartiesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCallsMaximumconferencepartiesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/maximumconferenceparties][%d] getConversationsCallsMaximumconferencepartiesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsCallsMaximumconferencepartiesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallsMaximumconferencepartiesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallsMaximumconferencepartiesTooManyRequests creates a GetConversationsCallsMaximumconferencepartiesTooManyRequests with default headers values
func NewGetConversationsCallsMaximumconferencepartiesTooManyRequests() *GetConversationsCallsMaximumconferencepartiesTooManyRequests {
	return &GetConversationsCallsMaximumconferencepartiesTooManyRequests{}
}

/*GetConversationsCallsMaximumconferencepartiesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetConversationsCallsMaximumconferencepartiesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCallsMaximumconferencepartiesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/maximumconferenceparties][%d] getConversationsCallsMaximumconferencepartiesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsCallsMaximumconferencepartiesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallsMaximumconferencepartiesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallsMaximumconferencepartiesInternalServerError creates a GetConversationsCallsMaximumconferencepartiesInternalServerError with default headers values
func NewGetConversationsCallsMaximumconferencepartiesInternalServerError() *GetConversationsCallsMaximumconferencepartiesInternalServerError {
	return &GetConversationsCallsMaximumconferencepartiesInternalServerError{}
}

/*GetConversationsCallsMaximumconferencepartiesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationsCallsMaximumconferencepartiesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCallsMaximumconferencepartiesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/maximumconferenceparties][%d] getConversationsCallsMaximumconferencepartiesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsCallsMaximumconferencepartiesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallsMaximumconferencepartiesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallsMaximumconferencepartiesServiceUnavailable creates a GetConversationsCallsMaximumconferencepartiesServiceUnavailable with default headers values
func NewGetConversationsCallsMaximumconferencepartiesServiceUnavailable() *GetConversationsCallsMaximumconferencepartiesServiceUnavailable {
	return &GetConversationsCallsMaximumconferencepartiesServiceUnavailable{}
}

/*GetConversationsCallsMaximumconferencepartiesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationsCallsMaximumconferencepartiesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCallsMaximumconferencepartiesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/maximumconferenceparties][%d] getConversationsCallsMaximumconferencepartiesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsCallsMaximumconferencepartiesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallsMaximumconferencepartiesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallsMaximumconferencepartiesGatewayTimeout creates a GetConversationsCallsMaximumconferencepartiesGatewayTimeout with default headers values
func NewGetConversationsCallsMaximumconferencepartiesGatewayTimeout() *GetConversationsCallsMaximumconferencepartiesGatewayTimeout {
	return &GetConversationsCallsMaximumconferencepartiesGatewayTimeout{}
}

/*GetConversationsCallsMaximumconferencepartiesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetConversationsCallsMaximumconferencepartiesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCallsMaximumconferencepartiesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/maximumconferenceparties][%d] getConversationsCallsMaximumconferencepartiesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsCallsMaximumconferencepartiesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallsMaximumconferencepartiesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

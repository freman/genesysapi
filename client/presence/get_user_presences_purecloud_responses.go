// Code generated by go-swagger; DO NOT EDIT.

package presence

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetUserPresencesPurecloudReader is a Reader for the GetUserPresencesPurecloud structure.
type GetUserPresencesPurecloudReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserPresencesPurecloudReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserPresencesPurecloudOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserPresencesPurecloudBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUserPresencesPurecloudUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserPresencesPurecloudForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserPresencesPurecloudNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetUserPresencesPurecloudRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetUserPresencesPurecloudUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetUserPresencesPurecloudTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserPresencesPurecloudInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUserPresencesPurecloudServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUserPresencesPurecloudGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserPresencesPurecloudOK creates a GetUserPresencesPurecloudOK with default headers values
func NewGetUserPresencesPurecloudOK() *GetUserPresencesPurecloudOK {
	return &GetUserPresencesPurecloudOK{}
}

/*GetUserPresencesPurecloudOK handles this case with default header values.

successful operation
*/
type GetUserPresencesPurecloudOK struct {
	Payload *models.UserPresence
}

func (o *GetUserPresencesPurecloudOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/presences/purecloud][%d] getUserPresencesPurecloudOK  %+v", 200, o.Payload)
}

func (o *GetUserPresencesPurecloudOK) GetPayload() *models.UserPresence {
	return o.Payload
}

func (o *GetUserPresencesPurecloudOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserPresence)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserPresencesPurecloudBadRequest creates a GetUserPresencesPurecloudBadRequest with default headers values
func NewGetUserPresencesPurecloudBadRequest() *GetUserPresencesPurecloudBadRequest {
	return &GetUserPresencesPurecloudBadRequest{}
}

/*GetUserPresencesPurecloudBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetUserPresencesPurecloudBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetUserPresencesPurecloudBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/presences/purecloud][%d] getUserPresencesPurecloudBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserPresencesPurecloudBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserPresencesPurecloudBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserPresencesPurecloudUnauthorized creates a GetUserPresencesPurecloudUnauthorized with default headers values
func NewGetUserPresencesPurecloudUnauthorized() *GetUserPresencesPurecloudUnauthorized {
	return &GetUserPresencesPurecloudUnauthorized{}
}

/*GetUserPresencesPurecloudUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetUserPresencesPurecloudUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetUserPresencesPurecloudUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/presences/purecloud][%d] getUserPresencesPurecloudUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserPresencesPurecloudUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserPresencesPurecloudUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserPresencesPurecloudForbidden creates a GetUserPresencesPurecloudForbidden with default headers values
func NewGetUserPresencesPurecloudForbidden() *GetUserPresencesPurecloudForbidden {
	return &GetUserPresencesPurecloudForbidden{}
}

/*GetUserPresencesPurecloudForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetUserPresencesPurecloudForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetUserPresencesPurecloudForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/presences/purecloud][%d] getUserPresencesPurecloudForbidden  %+v", 403, o.Payload)
}

func (o *GetUserPresencesPurecloudForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserPresencesPurecloudForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserPresencesPurecloudNotFound creates a GetUserPresencesPurecloudNotFound with default headers values
func NewGetUserPresencesPurecloudNotFound() *GetUserPresencesPurecloudNotFound {
	return &GetUserPresencesPurecloudNotFound{}
}

/*GetUserPresencesPurecloudNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetUserPresencesPurecloudNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetUserPresencesPurecloudNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/presences/purecloud][%d] getUserPresencesPurecloudNotFound  %+v", 404, o.Payload)
}

func (o *GetUserPresencesPurecloudNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserPresencesPurecloudNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserPresencesPurecloudRequestEntityTooLarge creates a GetUserPresencesPurecloudRequestEntityTooLarge with default headers values
func NewGetUserPresencesPurecloudRequestEntityTooLarge() *GetUserPresencesPurecloudRequestEntityTooLarge {
	return &GetUserPresencesPurecloudRequestEntityTooLarge{}
}

/*GetUserPresencesPurecloudRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetUserPresencesPurecloudRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetUserPresencesPurecloudRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/presences/purecloud][%d] getUserPresencesPurecloudRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetUserPresencesPurecloudRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserPresencesPurecloudRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserPresencesPurecloudUnsupportedMediaType creates a GetUserPresencesPurecloudUnsupportedMediaType with default headers values
func NewGetUserPresencesPurecloudUnsupportedMediaType() *GetUserPresencesPurecloudUnsupportedMediaType {
	return &GetUserPresencesPurecloudUnsupportedMediaType{}
}

/*GetUserPresencesPurecloudUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetUserPresencesPurecloudUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetUserPresencesPurecloudUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/presences/purecloud][%d] getUserPresencesPurecloudUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetUserPresencesPurecloudUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserPresencesPurecloudUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserPresencesPurecloudTooManyRequests creates a GetUserPresencesPurecloudTooManyRequests with default headers values
func NewGetUserPresencesPurecloudTooManyRequests() *GetUserPresencesPurecloudTooManyRequests {
	return &GetUserPresencesPurecloudTooManyRequests{}
}

/*GetUserPresencesPurecloudTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetUserPresencesPurecloudTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetUserPresencesPurecloudTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/presences/purecloud][%d] getUserPresencesPurecloudTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUserPresencesPurecloudTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserPresencesPurecloudTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserPresencesPurecloudInternalServerError creates a GetUserPresencesPurecloudInternalServerError with default headers values
func NewGetUserPresencesPurecloudInternalServerError() *GetUserPresencesPurecloudInternalServerError {
	return &GetUserPresencesPurecloudInternalServerError{}
}

/*GetUserPresencesPurecloudInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetUserPresencesPurecloudInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetUserPresencesPurecloudInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/presences/purecloud][%d] getUserPresencesPurecloudInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserPresencesPurecloudInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserPresencesPurecloudInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserPresencesPurecloudServiceUnavailable creates a GetUserPresencesPurecloudServiceUnavailable with default headers values
func NewGetUserPresencesPurecloudServiceUnavailable() *GetUserPresencesPurecloudServiceUnavailable {
	return &GetUserPresencesPurecloudServiceUnavailable{}
}

/*GetUserPresencesPurecloudServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetUserPresencesPurecloudServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetUserPresencesPurecloudServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/presences/purecloud][%d] getUserPresencesPurecloudServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUserPresencesPurecloudServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserPresencesPurecloudServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserPresencesPurecloudGatewayTimeout creates a GetUserPresencesPurecloudGatewayTimeout with default headers values
func NewGetUserPresencesPurecloudGatewayTimeout() *GetUserPresencesPurecloudGatewayTimeout {
	return &GetUserPresencesPurecloudGatewayTimeout{}
}

/*GetUserPresencesPurecloudGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetUserPresencesPurecloudGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetUserPresencesPurecloudGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/presences/purecloud][%d] getUserPresencesPurecloudGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUserPresencesPurecloudGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserPresencesPurecloudGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package gamification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetGamificationProfilesUsersMeReader is a Reader for the GetGamificationProfilesUsersMe structure.
type GetGamificationProfilesUsersMeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGamificationProfilesUsersMeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGamificationProfilesUsersMeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGamificationProfilesUsersMeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetGamificationProfilesUsersMeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGamificationProfilesUsersMeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGamificationProfilesUsersMeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetGamificationProfilesUsersMeRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetGamificationProfilesUsersMeRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetGamificationProfilesUsersMeUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetGamificationProfilesUsersMeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGamificationProfilesUsersMeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetGamificationProfilesUsersMeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetGamificationProfilesUsersMeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGamificationProfilesUsersMeOK creates a GetGamificationProfilesUsersMeOK with default headers values
func NewGetGamificationProfilesUsersMeOK() *GetGamificationProfilesUsersMeOK {
	return &GetGamificationProfilesUsersMeOK{}
}

/*GetGamificationProfilesUsersMeOK handles this case with default header values.

successful operation
*/
type GetGamificationProfilesUsersMeOK struct {
	Payload *models.PerformanceProfile
}

func (o *GetGamificationProfilesUsersMeOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/users/me][%d] getGamificationProfilesUsersMeOK  %+v", 200, o.Payload)
}

func (o *GetGamificationProfilesUsersMeOK) GetPayload() *models.PerformanceProfile {
	return o.Payload
}

func (o *GetGamificationProfilesUsersMeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PerformanceProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfilesUsersMeBadRequest creates a GetGamificationProfilesUsersMeBadRequest with default headers values
func NewGetGamificationProfilesUsersMeBadRequest() *GetGamificationProfilesUsersMeBadRequest {
	return &GetGamificationProfilesUsersMeBadRequest{}
}

/*GetGamificationProfilesUsersMeBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetGamificationProfilesUsersMeBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfilesUsersMeBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/users/me][%d] getGamificationProfilesUsersMeBadRequest  %+v", 400, o.Payload)
}

func (o *GetGamificationProfilesUsersMeBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfilesUsersMeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfilesUsersMeUnauthorized creates a GetGamificationProfilesUsersMeUnauthorized with default headers values
func NewGetGamificationProfilesUsersMeUnauthorized() *GetGamificationProfilesUsersMeUnauthorized {
	return &GetGamificationProfilesUsersMeUnauthorized{}
}

/*GetGamificationProfilesUsersMeUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetGamificationProfilesUsersMeUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfilesUsersMeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/users/me][%d] getGamificationProfilesUsersMeUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGamificationProfilesUsersMeUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfilesUsersMeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfilesUsersMeForbidden creates a GetGamificationProfilesUsersMeForbidden with default headers values
func NewGetGamificationProfilesUsersMeForbidden() *GetGamificationProfilesUsersMeForbidden {
	return &GetGamificationProfilesUsersMeForbidden{}
}

/*GetGamificationProfilesUsersMeForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetGamificationProfilesUsersMeForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfilesUsersMeForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/users/me][%d] getGamificationProfilesUsersMeForbidden  %+v", 403, o.Payload)
}

func (o *GetGamificationProfilesUsersMeForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfilesUsersMeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfilesUsersMeNotFound creates a GetGamificationProfilesUsersMeNotFound with default headers values
func NewGetGamificationProfilesUsersMeNotFound() *GetGamificationProfilesUsersMeNotFound {
	return &GetGamificationProfilesUsersMeNotFound{}
}

/*GetGamificationProfilesUsersMeNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetGamificationProfilesUsersMeNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfilesUsersMeNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/users/me][%d] getGamificationProfilesUsersMeNotFound  %+v", 404, o.Payload)
}

func (o *GetGamificationProfilesUsersMeNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfilesUsersMeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfilesUsersMeRequestTimeout creates a GetGamificationProfilesUsersMeRequestTimeout with default headers values
func NewGetGamificationProfilesUsersMeRequestTimeout() *GetGamificationProfilesUsersMeRequestTimeout {
	return &GetGamificationProfilesUsersMeRequestTimeout{}
}

/*GetGamificationProfilesUsersMeRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetGamificationProfilesUsersMeRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfilesUsersMeRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/users/me][%d] getGamificationProfilesUsersMeRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetGamificationProfilesUsersMeRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfilesUsersMeRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfilesUsersMeRequestEntityTooLarge creates a GetGamificationProfilesUsersMeRequestEntityTooLarge with default headers values
func NewGetGamificationProfilesUsersMeRequestEntityTooLarge() *GetGamificationProfilesUsersMeRequestEntityTooLarge {
	return &GetGamificationProfilesUsersMeRequestEntityTooLarge{}
}

/*GetGamificationProfilesUsersMeRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetGamificationProfilesUsersMeRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfilesUsersMeRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/users/me][%d] getGamificationProfilesUsersMeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetGamificationProfilesUsersMeRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfilesUsersMeRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfilesUsersMeUnsupportedMediaType creates a GetGamificationProfilesUsersMeUnsupportedMediaType with default headers values
func NewGetGamificationProfilesUsersMeUnsupportedMediaType() *GetGamificationProfilesUsersMeUnsupportedMediaType {
	return &GetGamificationProfilesUsersMeUnsupportedMediaType{}
}

/*GetGamificationProfilesUsersMeUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetGamificationProfilesUsersMeUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfilesUsersMeUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/users/me][%d] getGamificationProfilesUsersMeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetGamificationProfilesUsersMeUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfilesUsersMeUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfilesUsersMeTooManyRequests creates a GetGamificationProfilesUsersMeTooManyRequests with default headers values
func NewGetGamificationProfilesUsersMeTooManyRequests() *GetGamificationProfilesUsersMeTooManyRequests {
	return &GetGamificationProfilesUsersMeTooManyRequests{}
}

/*GetGamificationProfilesUsersMeTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetGamificationProfilesUsersMeTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfilesUsersMeTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/users/me][%d] getGamificationProfilesUsersMeTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetGamificationProfilesUsersMeTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfilesUsersMeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfilesUsersMeInternalServerError creates a GetGamificationProfilesUsersMeInternalServerError with default headers values
func NewGetGamificationProfilesUsersMeInternalServerError() *GetGamificationProfilesUsersMeInternalServerError {
	return &GetGamificationProfilesUsersMeInternalServerError{}
}

/*GetGamificationProfilesUsersMeInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetGamificationProfilesUsersMeInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfilesUsersMeInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/users/me][%d] getGamificationProfilesUsersMeInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGamificationProfilesUsersMeInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfilesUsersMeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfilesUsersMeServiceUnavailable creates a GetGamificationProfilesUsersMeServiceUnavailable with default headers values
func NewGetGamificationProfilesUsersMeServiceUnavailable() *GetGamificationProfilesUsersMeServiceUnavailable {
	return &GetGamificationProfilesUsersMeServiceUnavailable{}
}

/*GetGamificationProfilesUsersMeServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetGamificationProfilesUsersMeServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfilesUsersMeServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/users/me][%d] getGamificationProfilesUsersMeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetGamificationProfilesUsersMeServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfilesUsersMeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfilesUsersMeGatewayTimeout creates a GetGamificationProfilesUsersMeGatewayTimeout with default headers values
func NewGetGamificationProfilesUsersMeGatewayTimeout() *GetGamificationProfilesUsersMeGatewayTimeout {
	return &GetGamificationProfilesUsersMeGatewayTimeout{}
}

/*GetGamificationProfilesUsersMeGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetGamificationProfilesUsersMeGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfilesUsersMeGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/users/me][%d] getGamificationProfilesUsersMeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetGamificationProfilesUsersMeGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfilesUsersMeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

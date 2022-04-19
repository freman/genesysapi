// Code generated by go-swagger; DO NOT EDIT.

package authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetAuthorizationDivisionspermittedMeReader is a Reader for the GetAuthorizationDivisionspermittedMe structure.
type GetAuthorizationDivisionspermittedMeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuthorizationDivisionspermittedMeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAuthorizationDivisionspermittedMeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAuthorizationDivisionspermittedMeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAuthorizationDivisionspermittedMeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAuthorizationDivisionspermittedMeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAuthorizationDivisionspermittedMeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetAuthorizationDivisionspermittedMeRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetAuthorizationDivisionspermittedMeRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetAuthorizationDivisionspermittedMeUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAuthorizationDivisionspermittedMeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAuthorizationDivisionspermittedMeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAuthorizationDivisionspermittedMeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAuthorizationDivisionspermittedMeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAuthorizationDivisionspermittedMeOK creates a GetAuthorizationDivisionspermittedMeOK with default headers values
func NewGetAuthorizationDivisionspermittedMeOK() *GetAuthorizationDivisionspermittedMeOK {
	return &GetAuthorizationDivisionspermittedMeOK{}
}

/*GetAuthorizationDivisionspermittedMeOK handles this case with default header values.

successful operation
*/
type GetAuthorizationDivisionspermittedMeOK struct {
	Payload []*models.AuthzDivision
}

func (o *GetAuthorizationDivisionspermittedMeOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/divisionspermitted/me][%d] getAuthorizationDivisionspermittedMeOK  %+v", 200, o.Payload)
}

func (o *GetAuthorizationDivisionspermittedMeOK) GetPayload() []*models.AuthzDivision {
	return o.Payload
}

func (o *GetAuthorizationDivisionspermittedMeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationDivisionspermittedMeBadRequest creates a GetAuthorizationDivisionspermittedMeBadRequest with default headers values
func NewGetAuthorizationDivisionspermittedMeBadRequest() *GetAuthorizationDivisionspermittedMeBadRequest {
	return &GetAuthorizationDivisionspermittedMeBadRequest{}
}

/*GetAuthorizationDivisionspermittedMeBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAuthorizationDivisionspermittedMeBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationDivisionspermittedMeBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/divisionspermitted/me][%d] getAuthorizationDivisionspermittedMeBadRequest  %+v", 400, o.Payload)
}

func (o *GetAuthorizationDivisionspermittedMeBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationDivisionspermittedMeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationDivisionspermittedMeUnauthorized creates a GetAuthorizationDivisionspermittedMeUnauthorized with default headers values
func NewGetAuthorizationDivisionspermittedMeUnauthorized() *GetAuthorizationDivisionspermittedMeUnauthorized {
	return &GetAuthorizationDivisionspermittedMeUnauthorized{}
}

/*GetAuthorizationDivisionspermittedMeUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAuthorizationDivisionspermittedMeUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationDivisionspermittedMeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/divisionspermitted/me][%d] getAuthorizationDivisionspermittedMeUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAuthorizationDivisionspermittedMeUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationDivisionspermittedMeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationDivisionspermittedMeForbidden creates a GetAuthorizationDivisionspermittedMeForbidden with default headers values
func NewGetAuthorizationDivisionspermittedMeForbidden() *GetAuthorizationDivisionspermittedMeForbidden {
	return &GetAuthorizationDivisionspermittedMeForbidden{}
}

/*GetAuthorizationDivisionspermittedMeForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetAuthorizationDivisionspermittedMeForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationDivisionspermittedMeForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/divisionspermitted/me][%d] getAuthorizationDivisionspermittedMeForbidden  %+v", 403, o.Payload)
}

func (o *GetAuthorizationDivisionspermittedMeForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationDivisionspermittedMeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationDivisionspermittedMeNotFound creates a GetAuthorizationDivisionspermittedMeNotFound with default headers values
func NewGetAuthorizationDivisionspermittedMeNotFound() *GetAuthorizationDivisionspermittedMeNotFound {
	return &GetAuthorizationDivisionspermittedMeNotFound{}
}

/*GetAuthorizationDivisionspermittedMeNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetAuthorizationDivisionspermittedMeNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationDivisionspermittedMeNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/divisionspermitted/me][%d] getAuthorizationDivisionspermittedMeNotFound  %+v", 404, o.Payload)
}

func (o *GetAuthorizationDivisionspermittedMeNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationDivisionspermittedMeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationDivisionspermittedMeRequestTimeout creates a GetAuthorizationDivisionspermittedMeRequestTimeout with default headers values
func NewGetAuthorizationDivisionspermittedMeRequestTimeout() *GetAuthorizationDivisionspermittedMeRequestTimeout {
	return &GetAuthorizationDivisionspermittedMeRequestTimeout{}
}

/*GetAuthorizationDivisionspermittedMeRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetAuthorizationDivisionspermittedMeRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationDivisionspermittedMeRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/divisionspermitted/me][%d] getAuthorizationDivisionspermittedMeRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAuthorizationDivisionspermittedMeRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationDivisionspermittedMeRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationDivisionspermittedMeRequestEntityTooLarge creates a GetAuthorizationDivisionspermittedMeRequestEntityTooLarge with default headers values
func NewGetAuthorizationDivisionspermittedMeRequestEntityTooLarge() *GetAuthorizationDivisionspermittedMeRequestEntityTooLarge {
	return &GetAuthorizationDivisionspermittedMeRequestEntityTooLarge{}
}

/*GetAuthorizationDivisionspermittedMeRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetAuthorizationDivisionspermittedMeRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationDivisionspermittedMeRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/divisionspermitted/me][%d] getAuthorizationDivisionspermittedMeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAuthorizationDivisionspermittedMeRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationDivisionspermittedMeRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationDivisionspermittedMeUnsupportedMediaType creates a GetAuthorizationDivisionspermittedMeUnsupportedMediaType with default headers values
func NewGetAuthorizationDivisionspermittedMeUnsupportedMediaType() *GetAuthorizationDivisionspermittedMeUnsupportedMediaType {
	return &GetAuthorizationDivisionspermittedMeUnsupportedMediaType{}
}

/*GetAuthorizationDivisionspermittedMeUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAuthorizationDivisionspermittedMeUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationDivisionspermittedMeUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/divisionspermitted/me][%d] getAuthorizationDivisionspermittedMeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAuthorizationDivisionspermittedMeUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationDivisionspermittedMeUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationDivisionspermittedMeTooManyRequests creates a GetAuthorizationDivisionspermittedMeTooManyRequests with default headers values
func NewGetAuthorizationDivisionspermittedMeTooManyRequests() *GetAuthorizationDivisionspermittedMeTooManyRequests {
	return &GetAuthorizationDivisionspermittedMeTooManyRequests{}
}

/*GetAuthorizationDivisionspermittedMeTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetAuthorizationDivisionspermittedMeTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationDivisionspermittedMeTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/divisionspermitted/me][%d] getAuthorizationDivisionspermittedMeTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAuthorizationDivisionspermittedMeTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationDivisionspermittedMeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationDivisionspermittedMeInternalServerError creates a GetAuthorizationDivisionspermittedMeInternalServerError with default headers values
func NewGetAuthorizationDivisionspermittedMeInternalServerError() *GetAuthorizationDivisionspermittedMeInternalServerError {
	return &GetAuthorizationDivisionspermittedMeInternalServerError{}
}

/*GetAuthorizationDivisionspermittedMeInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAuthorizationDivisionspermittedMeInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationDivisionspermittedMeInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/divisionspermitted/me][%d] getAuthorizationDivisionspermittedMeInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAuthorizationDivisionspermittedMeInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationDivisionspermittedMeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationDivisionspermittedMeServiceUnavailable creates a GetAuthorizationDivisionspermittedMeServiceUnavailable with default headers values
func NewGetAuthorizationDivisionspermittedMeServiceUnavailable() *GetAuthorizationDivisionspermittedMeServiceUnavailable {
	return &GetAuthorizationDivisionspermittedMeServiceUnavailable{}
}

/*GetAuthorizationDivisionspermittedMeServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAuthorizationDivisionspermittedMeServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationDivisionspermittedMeServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/divisionspermitted/me][%d] getAuthorizationDivisionspermittedMeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAuthorizationDivisionspermittedMeServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationDivisionspermittedMeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationDivisionspermittedMeGatewayTimeout creates a GetAuthorizationDivisionspermittedMeGatewayTimeout with default headers values
func NewGetAuthorizationDivisionspermittedMeGatewayTimeout() *GetAuthorizationDivisionspermittedMeGatewayTimeout {
	return &GetAuthorizationDivisionspermittedMeGatewayTimeout{}
}

/*GetAuthorizationDivisionspermittedMeGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetAuthorizationDivisionspermittedMeGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationDivisionspermittedMeGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/divisionspermitted/me][%d] getAuthorizationDivisionspermittedMeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAuthorizationDivisionspermittedMeGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationDivisionspermittedMeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package o_auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetOauthClientsReader is a Reader for the GetOauthClients structure.
type GetOauthClientsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOauthClientsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOauthClientsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOauthClientsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOauthClientsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOauthClientsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOauthClientsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOauthClientsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOauthClientsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOauthClientsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOauthClientsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOauthClientsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOauthClientsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOauthClientsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOauthClientsOK creates a GetOauthClientsOK with default headers values
func NewGetOauthClientsOK() *GetOauthClientsOK {
	return &GetOauthClientsOK{}
}

/*GetOauthClientsOK handles this case with default header values.

successful operation
*/
type GetOauthClientsOK struct {
	Payload *models.OAuthClientEntityListing
}

func (o *GetOauthClientsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/oauth/clients][%d] getOauthClientsOK  %+v", 200, o.Payload)
}

func (o *GetOauthClientsOK) GetPayload() *models.OAuthClientEntityListing {
	return o.Payload
}

func (o *GetOauthClientsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OAuthClientEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOauthClientsBadRequest creates a GetOauthClientsBadRequest with default headers values
func NewGetOauthClientsBadRequest() *GetOauthClientsBadRequest {
	return &GetOauthClientsBadRequest{}
}

/*GetOauthClientsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOauthClientsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetOauthClientsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/oauth/clients][%d] getOauthClientsBadRequest  %+v", 400, o.Payload)
}

func (o *GetOauthClientsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOauthClientsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOauthClientsUnauthorized creates a GetOauthClientsUnauthorized with default headers values
func NewGetOauthClientsUnauthorized() *GetOauthClientsUnauthorized {
	return &GetOauthClientsUnauthorized{}
}

/*GetOauthClientsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOauthClientsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetOauthClientsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/oauth/clients][%d] getOauthClientsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOauthClientsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOauthClientsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOauthClientsForbidden creates a GetOauthClientsForbidden with default headers values
func NewGetOauthClientsForbidden() *GetOauthClientsForbidden {
	return &GetOauthClientsForbidden{}
}

/*GetOauthClientsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetOauthClientsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetOauthClientsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/oauth/clients][%d] getOauthClientsForbidden  %+v", 403, o.Payload)
}

func (o *GetOauthClientsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOauthClientsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOauthClientsNotFound creates a GetOauthClientsNotFound with default headers values
func NewGetOauthClientsNotFound() *GetOauthClientsNotFound {
	return &GetOauthClientsNotFound{}
}

/*GetOauthClientsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetOauthClientsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetOauthClientsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/oauth/clients][%d] getOauthClientsNotFound  %+v", 404, o.Payload)
}

func (o *GetOauthClientsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOauthClientsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOauthClientsRequestTimeout creates a GetOauthClientsRequestTimeout with default headers values
func NewGetOauthClientsRequestTimeout() *GetOauthClientsRequestTimeout {
	return &GetOauthClientsRequestTimeout{}
}

/*GetOauthClientsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetOauthClientsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOauthClientsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/oauth/clients][%d] getOauthClientsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOauthClientsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOauthClientsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOauthClientsRequestEntityTooLarge creates a GetOauthClientsRequestEntityTooLarge with default headers values
func NewGetOauthClientsRequestEntityTooLarge() *GetOauthClientsRequestEntityTooLarge {
	return &GetOauthClientsRequestEntityTooLarge{}
}

/*GetOauthClientsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetOauthClientsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetOauthClientsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/oauth/clients][%d] getOauthClientsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOauthClientsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOauthClientsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOauthClientsUnsupportedMediaType creates a GetOauthClientsUnsupportedMediaType with default headers values
func NewGetOauthClientsUnsupportedMediaType() *GetOauthClientsUnsupportedMediaType {
	return &GetOauthClientsUnsupportedMediaType{}
}

/*GetOauthClientsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOauthClientsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetOauthClientsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/oauth/clients][%d] getOauthClientsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOauthClientsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOauthClientsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOauthClientsTooManyRequests creates a GetOauthClientsTooManyRequests with default headers values
func NewGetOauthClientsTooManyRequests() *GetOauthClientsTooManyRequests {
	return &GetOauthClientsTooManyRequests{}
}

/*GetOauthClientsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetOauthClientsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetOauthClientsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/oauth/clients][%d] getOauthClientsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOauthClientsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOauthClientsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOauthClientsInternalServerError creates a GetOauthClientsInternalServerError with default headers values
func NewGetOauthClientsInternalServerError() *GetOauthClientsInternalServerError {
	return &GetOauthClientsInternalServerError{}
}

/*GetOauthClientsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOauthClientsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetOauthClientsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/oauth/clients][%d] getOauthClientsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOauthClientsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOauthClientsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOauthClientsServiceUnavailable creates a GetOauthClientsServiceUnavailable with default headers values
func NewGetOauthClientsServiceUnavailable() *GetOauthClientsServiceUnavailable {
	return &GetOauthClientsServiceUnavailable{}
}

/*GetOauthClientsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOauthClientsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetOauthClientsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/oauth/clients][%d] getOauthClientsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOauthClientsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOauthClientsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOauthClientsGatewayTimeout creates a GetOauthClientsGatewayTimeout with default headers values
func NewGetOauthClientsGatewayTimeout() *GetOauthClientsGatewayTimeout {
	return &GetOauthClientsGatewayTimeout{}
}

/*GetOauthClientsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetOauthClientsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOauthClientsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/oauth/clients][%d] getOauthClientsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOauthClientsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOauthClientsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

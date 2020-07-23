// Code generated by go-swagger; DO NOT EDIT.

package identity_provider

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetIdentityprovidersReader is a Reader for the GetIdentityproviders structure.
type GetIdentityprovidersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIdentityprovidersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIdentityprovidersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIdentityprovidersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetIdentityprovidersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIdentityprovidersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIdentityprovidersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetIdentityprovidersRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetIdentityprovidersUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetIdentityprovidersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIdentityprovidersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetIdentityprovidersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetIdentityprovidersGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIdentityprovidersOK creates a GetIdentityprovidersOK with default headers values
func NewGetIdentityprovidersOK() *GetIdentityprovidersOK {
	return &GetIdentityprovidersOK{}
}

/*GetIdentityprovidersOK handles this case with default header values.

successful operation
*/
type GetIdentityprovidersOK struct {
	Payload *models.OAuthProviderEntityListing
}

func (o *GetIdentityprovidersOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders][%d] getIdentityprovidersOK  %+v", 200, o.Payload)
}

func (o *GetIdentityprovidersOK) GetPayload() *models.OAuthProviderEntityListing {
	return o.Payload
}

func (o *GetIdentityprovidersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OAuthProviderEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersBadRequest creates a GetIdentityprovidersBadRequest with default headers values
func NewGetIdentityprovidersBadRequest() *GetIdentityprovidersBadRequest {
	return &GetIdentityprovidersBadRequest{}
}

/*GetIdentityprovidersBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetIdentityprovidersBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders][%d] getIdentityprovidersBadRequest  %+v", 400, o.Payload)
}

func (o *GetIdentityprovidersBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersUnauthorized creates a GetIdentityprovidersUnauthorized with default headers values
func NewGetIdentityprovidersUnauthorized() *GetIdentityprovidersUnauthorized {
	return &GetIdentityprovidersUnauthorized{}
}

/*GetIdentityprovidersUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetIdentityprovidersUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders][%d] getIdentityprovidersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIdentityprovidersUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersForbidden creates a GetIdentityprovidersForbidden with default headers values
func NewGetIdentityprovidersForbidden() *GetIdentityprovidersForbidden {
	return &GetIdentityprovidersForbidden{}
}

/*GetIdentityprovidersForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetIdentityprovidersForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders][%d] getIdentityprovidersForbidden  %+v", 403, o.Payload)
}

func (o *GetIdentityprovidersForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersNotFound creates a GetIdentityprovidersNotFound with default headers values
func NewGetIdentityprovidersNotFound() *GetIdentityprovidersNotFound {
	return &GetIdentityprovidersNotFound{}
}

/*GetIdentityprovidersNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetIdentityprovidersNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders][%d] getIdentityprovidersNotFound  %+v", 404, o.Payload)
}

func (o *GetIdentityprovidersNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersRequestEntityTooLarge creates a GetIdentityprovidersRequestEntityTooLarge with default headers values
func NewGetIdentityprovidersRequestEntityTooLarge() *GetIdentityprovidersRequestEntityTooLarge {
	return &GetIdentityprovidersRequestEntityTooLarge{}
}

/*GetIdentityprovidersRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetIdentityprovidersRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders][%d] getIdentityprovidersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetIdentityprovidersRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersUnsupportedMediaType creates a GetIdentityprovidersUnsupportedMediaType with default headers values
func NewGetIdentityprovidersUnsupportedMediaType() *GetIdentityprovidersUnsupportedMediaType {
	return &GetIdentityprovidersUnsupportedMediaType{}
}

/*GetIdentityprovidersUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetIdentityprovidersUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders][%d] getIdentityprovidersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetIdentityprovidersUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersTooManyRequests creates a GetIdentityprovidersTooManyRequests with default headers values
func NewGetIdentityprovidersTooManyRequests() *GetIdentityprovidersTooManyRequests {
	return &GetIdentityprovidersTooManyRequests{}
}

/*GetIdentityprovidersTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetIdentityprovidersTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders][%d] getIdentityprovidersTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIdentityprovidersTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersInternalServerError creates a GetIdentityprovidersInternalServerError with default headers values
func NewGetIdentityprovidersInternalServerError() *GetIdentityprovidersInternalServerError {
	return &GetIdentityprovidersInternalServerError{}
}

/*GetIdentityprovidersInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetIdentityprovidersInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders][%d] getIdentityprovidersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIdentityprovidersInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersServiceUnavailable creates a GetIdentityprovidersServiceUnavailable with default headers values
func NewGetIdentityprovidersServiceUnavailable() *GetIdentityprovidersServiceUnavailable {
	return &GetIdentityprovidersServiceUnavailable{}
}

/*GetIdentityprovidersServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetIdentityprovidersServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders][%d] getIdentityprovidersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIdentityprovidersServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersGatewayTimeout creates a GetIdentityprovidersGatewayTimeout with default headers values
func NewGetIdentityprovidersGatewayTimeout() *GetIdentityprovidersGatewayTimeout {
	return &GetIdentityprovidersGatewayTimeout{}
}

/*GetIdentityprovidersGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetIdentityprovidersGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders][%d] getIdentityprovidersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIdentityprovidersGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

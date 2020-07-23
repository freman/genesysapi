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

// GetIdentityprovidersAdfsReader is a Reader for the GetIdentityprovidersAdfs structure.
type GetIdentityprovidersAdfsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIdentityprovidersAdfsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIdentityprovidersAdfsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIdentityprovidersAdfsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetIdentityprovidersAdfsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIdentityprovidersAdfsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIdentityprovidersAdfsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetIdentityprovidersAdfsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetIdentityprovidersAdfsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetIdentityprovidersAdfsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIdentityprovidersAdfsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetIdentityprovidersAdfsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetIdentityprovidersAdfsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIdentityprovidersAdfsOK creates a GetIdentityprovidersAdfsOK with default headers values
func NewGetIdentityprovidersAdfsOK() *GetIdentityprovidersAdfsOK {
	return &GetIdentityprovidersAdfsOK{}
}

/*GetIdentityprovidersAdfsOK handles this case with default header values.

successful operation
*/
type GetIdentityprovidersAdfsOK struct {
	Payload *models.ADFS
}

func (o *GetIdentityprovidersAdfsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/adfs][%d] getIdentityprovidersAdfsOK  %+v", 200, o.Payload)
}

func (o *GetIdentityprovidersAdfsOK) GetPayload() *models.ADFS {
	return o.Payload
}

func (o *GetIdentityprovidersAdfsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ADFS)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersAdfsBadRequest creates a GetIdentityprovidersAdfsBadRequest with default headers values
func NewGetIdentityprovidersAdfsBadRequest() *GetIdentityprovidersAdfsBadRequest {
	return &GetIdentityprovidersAdfsBadRequest{}
}

/*GetIdentityprovidersAdfsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetIdentityprovidersAdfsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersAdfsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/adfs][%d] getIdentityprovidersAdfsBadRequest  %+v", 400, o.Payload)
}

func (o *GetIdentityprovidersAdfsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersAdfsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersAdfsUnauthorized creates a GetIdentityprovidersAdfsUnauthorized with default headers values
func NewGetIdentityprovidersAdfsUnauthorized() *GetIdentityprovidersAdfsUnauthorized {
	return &GetIdentityprovidersAdfsUnauthorized{}
}

/*GetIdentityprovidersAdfsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetIdentityprovidersAdfsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersAdfsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/adfs][%d] getIdentityprovidersAdfsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIdentityprovidersAdfsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersAdfsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersAdfsForbidden creates a GetIdentityprovidersAdfsForbidden with default headers values
func NewGetIdentityprovidersAdfsForbidden() *GetIdentityprovidersAdfsForbidden {
	return &GetIdentityprovidersAdfsForbidden{}
}

/*GetIdentityprovidersAdfsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetIdentityprovidersAdfsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersAdfsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/adfs][%d] getIdentityprovidersAdfsForbidden  %+v", 403, o.Payload)
}

func (o *GetIdentityprovidersAdfsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersAdfsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersAdfsNotFound creates a GetIdentityprovidersAdfsNotFound with default headers values
func NewGetIdentityprovidersAdfsNotFound() *GetIdentityprovidersAdfsNotFound {
	return &GetIdentityprovidersAdfsNotFound{}
}

/*GetIdentityprovidersAdfsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetIdentityprovidersAdfsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersAdfsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/adfs][%d] getIdentityprovidersAdfsNotFound  %+v", 404, o.Payload)
}

func (o *GetIdentityprovidersAdfsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersAdfsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersAdfsRequestEntityTooLarge creates a GetIdentityprovidersAdfsRequestEntityTooLarge with default headers values
func NewGetIdentityprovidersAdfsRequestEntityTooLarge() *GetIdentityprovidersAdfsRequestEntityTooLarge {
	return &GetIdentityprovidersAdfsRequestEntityTooLarge{}
}

/*GetIdentityprovidersAdfsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetIdentityprovidersAdfsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersAdfsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/adfs][%d] getIdentityprovidersAdfsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetIdentityprovidersAdfsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersAdfsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersAdfsUnsupportedMediaType creates a GetIdentityprovidersAdfsUnsupportedMediaType with default headers values
func NewGetIdentityprovidersAdfsUnsupportedMediaType() *GetIdentityprovidersAdfsUnsupportedMediaType {
	return &GetIdentityprovidersAdfsUnsupportedMediaType{}
}

/*GetIdentityprovidersAdfsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetIdentityprovidersAdfsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersAdfsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/adfs][%d] getIdentityprovidersAdfsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetIdentityprovidersAdfsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersAdfsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersAdfsTooManyRequests creates a GetIdentityprovidersAdfsTooManyRequests with default headers values
func NewGetIdentityprovidersAdfsTooManyRequests() *GetIdentityprovidersAdfsTooManyRequests {
	return &GetIdentityprovidersAdfsTooManyRequests{}
}

/*GetIdentityprovidersAdfsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetIdentityprovidersAdfsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersAdfsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/adfs][%d] getIdentityprovidersAdfsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIdentityprovidersAdfsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersAdfsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersAdfsInternalServerError creates a GetIdentityprovidersAdfsInternalServerError with default headers values
func NewGetIdentityprovidersAdfsInternalServerError() *GetIdentityprovidersAdfsInternalServerError {
	return &GetIdentityprovidersAdfsInternalServerError{}
}

/*GetIdentityprovidersAdfsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetIdentityprovidersAdfsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersAdfsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/adfs][%d] getIdentityprovidersAdfsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIdentityprovidersAdfsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersAdfsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersAdfsServiceUnavailable creates a GetIdentityprovidersAdfsServiceUnavailable with default headers values
func NewGetIdentityprovidersAdfsServiceUnavailable() *GetIdentityprovidersAdfsServiceUnavailable {
	return &GetIdentityprovidersAdfsServiceUnavailable{}
}

/*GetIdentityprovidersAdfsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetIdentityprovidersAdfsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersAdfsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/adfs][%d] getIdentityprovidersAdfsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIdentityprovidersAdfsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersAdfsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersAdfsGatewayTimeout creates a GetIdentityprovidersAdfsGatewayTimeout with default headers values
func NewGetIdentityprovidersAdfsGatewayTimeout() *GetIdentityprovidersAdfsGatewayTimeout {
	return &GetIdentityprovidersAdfsGatewayTimeout{}
}

/*GetIdentityprovidersAdfsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetIdentityprovidersAdfsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersAdfsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/adfs][%d] getIdentityprovidersAdfsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIdentityprovidersAdfsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersAdfsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

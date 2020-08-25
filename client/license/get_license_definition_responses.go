// Code generated by go-swagger; DO NOT EDIT.

package license

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetLicenseDefinitionReader is a Reader for the GetLicenseDefinition structure.
type GetLicenseDefinitionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLicenseDefinitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLicenseDefinitionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLicenseDefinitionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLicenseDefinitionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLicenseDefinitionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLicenseDefinitionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetLicenseDefinitionRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetLicenseDefinitionUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetLicenseDefinitionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLicenseDefinitionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetLicenseDefinitionServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetLicenseDefinitionGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLicenseDefinitionOK creates a GetLicenseDefinitionOK with default headers values
func NewGetLicenseDefinitionOK() *GetLicenseDefinitionOK {
	return &GetLicenseDefinitionOK{}
}

/*GetLicenseDefinitionOK handles this case with default header values.

successful operation
*/
type GetLicenseDefinitionOK struct {
	Payload *models.LicenseDefinition
}

func (o *GetLicenseDefinitionOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions/{licenseId}][%d] getLicenseDefinitionOK  %+v", 200, o.Payload)
}

func (o *GetLicenseDefinitionOK) GetPayload() *models.LicenseDefinition {
	return o.Payload
}

func (o *GetLicenseDefinitionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LicenseDefinition)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseDefinitionBadRequest creates a GetLicenseDefinitionBadRequest with default headers values
func NewGetLicenseDefinitionBadRequest() *GetLicenseDefinitionBadRequest {
	return &GetLicenseDefinitionBadRequest{}
}

/*GetLicenseDefinitionBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetLicenseDefinitionBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetLicenseDefinitionBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions/{licenseId}][%d] getLicenseDefinitionBadRequest  %+v", 400, o.Payload)
}

func (o *GetLicenseDefinitionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseDefinitionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseDefinitionUnauthorized creates a GetLicenseDefinitionUnauthorized with default headers values
func NewGetLicenseDefinitionUnauthorized() *GetLicenseDefinitionUnauthorized {
	return &GetLicenseDefinitionUnauthorized{}
}

/*GetLicenseDefinitionUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetLicenseDefinitionUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetLicenseDefinitionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions/{licenseId}][%d] getLicenseDefinitionUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLicenseDefinitionUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseDefinitionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseDefinitionForbidden creates a GetLicenseDefinitionForbidden with default headers values
func NewGetLicenseDefinitionForbidden() *GetLicenseDefinitionForbidden {
	return &GetLicenseDefinitionForbidden{}
}

/*GetLicenseDefinitionForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetLicenseDefinitionForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetLicenseDefinitionForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions/{licenseId}][%d] getLicenseDefinitionForbidden  %+v", 403, o.Payload)
}

func (o *GetLicenseDefinitionForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseDefinitionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseDefinitionNotFound creates a GetLicenseDefinitionNotFound with default headers values
func NewGetLicenseDefinitionNotFound() *GetLicenseDefinitionNotFound {
	return &GetLicenseDefinitionNotFound{}
}

/*GetLicenseDefinitionNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetLicenseDefinitionNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetLicenseDefinitionNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions/{licenseId}][%d] getLicenseDefinitionNotFound  %+v", 404, o.Payload)
}

func (o *GetLicenseDefinitionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseDefinitionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseDefinitionRequestEntityTooLarge creates a GetLicenseDefinitionRequestEntityTooLarge with default headers values
func NewGetLicenseDefinitionRequestEntityTooLarge() *GetLicenseDefinitionRequestEntityTooLarge {
	return &GetLicenseDefinitionRequestEntityTooLarge{}
}

/*GetLicenseDefinitionRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetLicenseDefinitionRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetLicenseDefinitionRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions/{licenseId}][%d] getLicenseDefinitionRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetLicenseDefinitionRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseDefinitionRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseDefinitionUnsupportedMediaType creates a GetLicenseDefinitionUnsupportedMediaType with default headers values
func NewGetLicenseDefinitionUnsupportedMediaType() *GetLicenseDefinitionUnsupportedMediaType {
	return &GetLicenseDefinitionUnsupportedMediaType{}
}

/*GetLicenseDefinitionUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetLicenseDefinitionUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetLicenseDefinitionUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions/{licenseId}][%d] getLicenseDefinitionUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetLicenseDefinitionUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseDefinitionUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseDefinitionTooManyRequests creates a GetLicenseDefinitionTooManyRequests with default headers values
func NewGetLicenseDefinitionTooManyRequests() *GetLicenseDefinitionTooManyRequests {
	return &GetLicenseDefinitionTooManyRequests{}
}

/*GetLicenseDefinitionTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetLicenseDefinitionTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetLicenseDefinitionTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions/{licenseId}][%d] getLicenseDefinitionTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLicenseDefinitionTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseDefinitionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseDefinitionInternalServerError creates a GetLicenseDefinitionInternalServerError with default headers values
func NewGetLicenseDefinitionInternalServerError() *GetLicenseDefinitionInternalServerError {
	return &GetLicenseDefinitionInternalServerError{}
}

/*GetLicenseDefinitionInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetLicenseDefinitionInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetLicenseDefinitionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions/{licenseId}][%d] getLicenseDefinitionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLicenseDefinitionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseDefinitionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseDefinitionServiceUnavailable creates a GetLicenseDefinitionServiceUnavailable with default headers values
func NewGetLicenseDefinitionServiceUnavailable() *GetLicenseDefinitionServiceUnavailable {
	return &GetLicenseDefinitionServiceUnavailable{}
}

/*GetLicenseDefinitionServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetLicenseDefinitionServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetLicenseDefinitionServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions/{licenseId}][%d] getLicenseDefinitionServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLicenseDefinitionServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseDefinitionServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseDefinitionGatewayTimeout creates a GetLicenseDefinitionGatewayTimeout with default headers values
func NewGetLicenseDefinitionGatewayTimeout() *GetLicenseDefinitionGatewayTimeout {
	return &GetLicenseDefinitionGatewayTimeout{}
}

/*GetLicenseDefinitionGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetLicenseDefinitionGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetLicenseDefinitionGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions/{licenseId}][%d] getLicenseDefinitionGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetLicenseDefinitionGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseDefinitionGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
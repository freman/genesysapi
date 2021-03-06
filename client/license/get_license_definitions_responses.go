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

// GetLicenseDefinitionsReader is a Reader for the GetLicenseDefinitions structure.
type GetLicenseDefinitionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLicenseDefinitionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLicenseDefinitionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLicenseDefinitionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLicenseDefinitionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLicenseDefinitionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLicenseDefinitionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetLicenseDefinitionsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetLicenseDefinitionsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetLicenseDefinitionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLicenseDefinitionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetLicenseDefinitionsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetLicenseDefinitionsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLicenseDefinitionsOK creates a GetLicenseDefinitionsOK with default headers values
func NewGetLicenseDefinitionsOK() *GetLicenseDefinitionsOK {
	return &GetLicenseDefinitionsOK{}
}

/*GetLicenseDefinitionsOK handles this case with default header values.

successful operation
*/
type GetLicenseDefinitionsOK struct {
	Payload []*models.LicenseDefinition
}

func (o *GetLicenseDefinitionsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions][%d] getLicenseDefinitionsOK  %+v", 200, o.Payload)
}

func (o *GetLicenseDefinitionsOK) GetPayload() []*models.LicenseDefinition {
	return o.Payload
}

func (o *GetLicenseDefinitionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseDefinitionsBadRequest creates a GetLicenseDefinitionsBadRequest with default headers values
func NewGetLicenseDefinitionsBadRequest() *GetLicenseDefinitionsBadRequest {
	return &GetLicenseDefinitionsBadRequest{}
}

/*GetLicenseDefinitionsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetLicenseDefinitionsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetLicenseDefinitionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions][%d] getLicenseDefinitionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetLicenseDefinitionsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseDefinitionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseDefinitionsUnauthorized creates a GetLicenseDefinitionsUnauthorized with default headers values
func NewGetLicenseDefinitionsUnauthorized() *GetLicenseDefinitionsUnauthorized {
	return &GetLicenseDefinitionsUnauthorized{}
}

/*GetLicenseDefinitionsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetLicenseDefinitionsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetLicenseDefinitionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions][%d] getLicenseDefinitionsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLicenseDefinitionsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseDefinitionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseDefinitionsForbidden creates a GetLicenseDefinitionsForbidden with default headers values
func NewGetLicenseDefinitionsForbidden() *GetLicenseDefinitionsForbidden {
	return &GetLicenseDefinitionsForbidden{}
}

/*GetLicenseDefinitionsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetLicenseDefinitionsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetLicenseDefinitionsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions][%d] getLicenseDefinitionsForbidden  %+v", 403, o.Payload)
}

func (o *GetLicenseDefinitionsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseDefinitionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseDefinitionsNotFound creates a GetLicenseDefinitionsNotFound with default headers values
func NewGetLicenseDefinitionsNotFound() *GetLicenseDefinitionsNotFound {
	return &GetLicenseDefinitionsNotFound{}
}

/*GetLicenseDefinitionsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetLicenseDefinitionsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetLicenseDefinitionsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions][%d] getLicenseDefinitionsNotFound  %+v", 404, o.Payload)
}

func (o *GetLicenseDefinitionsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseDefinitionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseDefinitionsRequestEntityTooLarge creates a GetLicenseDefinitionsRequestEntityTooLarge with default headers values
func NewGetLicenseDefinitionsRequestEntityTooLarge() *GetLicenseDefinitionsRequestEntityTooLarge {
	return &GetLicenseDefinitionsRequestEntityTooLarge{}
}

/*GetLicenseDefinitionsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetLicenseDefinitionsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetLicenseDefinitionsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions][%d] getLicenseDefinitionsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetLicenseDefinitionsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseDefinitionsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseDefinitionsUnsupportedMediaType creates a GetLicenseDefinitionsUnsupportedMediaType with default headers values
func NewGetLicenseDefinitionsUnsupportedMediaType() *GetLicenseDefinitionsUnsupportedMediaType {
	return &GetLicenseDefinitionsUnsupportedMediaType{}
}

/*GetLicenseDefinitionsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetLicenseDefinitionsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetLicenseDefinitionsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions][%d] getLicenseDefinitionsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetLicenseDefinitionsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseDefinitionsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseDefinitionsTooManyRequests creates a GetLicenseDefinitionsTooManyRequests with default headers values
func NewGetLicenseDefinitionsTooManyRequests() *GetLicenseDefinitionsTooManyRequests {
	return &GetLicenseDefinitionsTooManyRequests{}
}

/*GetLicenseDefinitionsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetLicenseDefinitionsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetLicenseDefinitionsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions][%d] getLicenseDefinitionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLicenseDefinitionsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseDefinitionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseDefinitionsInternalServerError creates a GetLicenseDefinitionsInternalServerError with default headers values
func NewGetLicenseDefinitionsInternalServerError() *GetLicenseDefinitionsInternalServerError {
	return &GetLicenseDefinitionsInternalServerError{}
}

/*GetLicenseDefinitionsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetLicenseDefinitionsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetLicenseDefinitionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions][%d] getLicenseDefinitionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLicenseDefinitionsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseDefinitionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseDefinitionsServiceUnavailable creates a GetLicenseDefinitionsServiceUnavailable with default headers values
func NewGetLicenseDefinitionsServiceUnavailable() *GetLicenseDefinitionsServiceUnavailable {
	return &GetLicenseDefinitionsServiceUnavailable{}
}

/*GetLicenseDefinitionsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetLicenseDefinitionsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetLicenseDefinitionsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions][%d] getLicenseDefinitionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLicenseDefinitionsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseDefinitionsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseDefinitionsGatewayTimeout creates a GetLicenseDefinitionsGatewayTimeout with default headers values
func NewGetLicenseDefinitionsGatewayTimeout() *GetLicenseDefinitionsGatewayTimeout {
	return &GetLicenseDefinitionsGatewayTimeout{}
}

/*GetLicenseDefinitionsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetLicenseDefinitionsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetLicenseDefinitionsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions][%d] getLicenseDefinitionsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetLicenseDefinitionsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseDefinitionsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

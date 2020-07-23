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

// GetLicenseToggleReader is a Reader for the GetLicenseToggle structure.
type GetLicenseToggleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLicenseToggleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLicenseToggleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLicenseToggleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLicenseToggleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLicenseToggleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLicenseToggleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetLicenseToggleRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetLicenseToggleUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetLicenseToggleTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLicenseToggleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetLicenseToggleServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetLicenseToggleGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLicenseToggleOK creates a GetLicenseToggleOK with default headers values
func NewGetLicenseToggleOK() *GetLicenseToggleOK {
	return &GetLicenseToggleOK{}
}

/*GetLicenseToggleOK handles this case with default header values.

successful operation
*/
type GetLicenseToggleOK struct {
	Payload *models.LicenseOrgToggle
}

func (o *GetLicenseToggleOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/toggles/{featureName}][%d] getLicenseToggleOK  %+v", 200, o.Payload)
}

func (o *GetLicenseToggleOK) GetPayload() *models.LicenseOrgToggle {
	return o.Payload
}

func (o *GetLicenseToggleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LicenseOrgToggle)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseToggleBadRequest creates a GetLicenseToggleBadRequest with default headers values
func NewGetLicenseToggleBadRequest() *GetLicenseToggleBadRequest {
	return &GetLicenseToggleBadRequest{}
}

/*GetLicenseToggleBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetLicenseToggleBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetLicenseToggleBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/toggles/{featureName}][%d] getLicenseToggleBadRequest  %+v", 400, o.Payload)
}

func (o *GetLicenseToggleBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseToggleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseToggleUnauthorized creates a GetLicenseToggleUnauthorized with default headers values
func NewGetLicenseToggleUnauthorized() *GetLicenseToggleUnauthorized {
	return &GetLicenseToggleUnauthorized{}
}

/*GetLicenseToggleUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetLicenseToggleUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetLicenseToggleUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/toggles/{featureName}][%d] getLicenseToggleUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLicenseToggleUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseToggleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseToggleForbidden creates a GetLicenseToggleForbidden with default headers values
func NewGetLicenseToggleForbidden() *GetLicenseToggleForbidden {
	return &GetLicenseToggleForbidden{}
}

/*GetLicenseToggleForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetLicenseToggleForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetLicenseToggleForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/toggles/{featureName}][%d] getLicenseToggleForbidden  %+v", 403, o.Payload)
}

func (o *GetLicenseToggleForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseToggleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseToggleNotFound creates a GetLicenseToggleNotFound with default headers values
func NewGetLicenseToggleNotFound() *GetLicenseToggleNotFound {
	return &GetLicenseToggleNotFound{}
}

/*GetLicenseToggleNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetLicenseToggleNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetLicenseToggleNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/toggles/{featureName}][%d] getLicenseToggleNotFound  %+v", 404, o.Payload)
}

func (o *GetLicenseToggleNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseToggleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseToggleRequestEntityTooLarge creates a GetLicenseToggleRequestEntityTooLarge with default headers values
func NewGetLicenseToggleRequestEntityTooLarge() *GetLicenseToggleRequestEntityTooLarge {
	return &GetLicenseToggleRequestEntityTooLarge{}
}

/*GetLicenseToggleRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetLicenseToggleRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetLicenseToggleRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/toggles/{featureName}][%d] getLicenseToggleRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetLicenseToggleRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseToggleRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseToggleUnsupportedMediaType creates a GetLicenseToggleUnsupportedMediaType with default headers values
func NewGetLicenseToggleUnsupportedMediaType() *GetLicenseToggleUnsupportedMediaType {
	return &GetLicenseToggleUnsupportedMediaType{}
}

/*GetLicenseToggleUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetLicenseToggleUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetLicenseToggleUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/toggles/{featureName}][%d] getLicenseToggleUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetLicenseToggleUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseToggleUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseToggleTooManyRequests creates a GetLicenseToggleTooManyRequests with default headers values
func NewGetLicenseToggleTooManyRequests() *GetLicenseToggleTooManyRequests {
	return &GetLicenseToggleTooManyRequests{}
}

/*GetLicenseToggleTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetLicenseToggleTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetLicenseToggleTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/toggles/{featureName}][%d] getLicenseToggleTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLicenseToggleTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseToggleTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseToggleInternalServerError creates a GetLicenseToggleInternalServerError with default headers values
func NewGetLicenseToggleInternalServerError() *GetLicenseToggleInternalServerError {
	return &GetLicenseToggleInternalServerError{}
}

/*GetLicenseToggleInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetLicenseToggleInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetLicenseToggleInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/toggles/{featureName}][%d] getLicenseToggleInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLicenseToggleInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseToggleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseToggleServiceUnavailable creates a GetLicenseToggleServiceUnavailable with default headers values
func NewGetLicenseToggleServiceUnavailable() *GetLicenseToggleServiceUnavailable {
	return &GetLicenseToggleServiceUnavailable{}
}

/*GetLicenseToggleServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetLicenseToggleServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetLicenseToggleServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/toggles/{featureName}][%d] getLicenseToggleServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLicenseToggleServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseToggleServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseToggleGatewayTimeout creates a GetLicenseToggleGatewayTimeout with default headers values
func NewGetLicenseToggleGatewayTimeout() *GetLicenseToggleGatewayTimeout {
	return &GetLicenseToggleGatewayTimeout{}
}

/*GetLicenseToggleGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetLicenseToggleGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetLicenseToggleGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/toggles/{featureName}][%d] getLicenseToggleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetLicenseToggleGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseToggleGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

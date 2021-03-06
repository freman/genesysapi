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

// PostLicenseToggleReader is a Reader for the PostLicenseToggle structure.
type PostLicenseToggleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLicenseToggleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLicenseToggleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostLicenseToggleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostLicenseToggleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostLicenseToggleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostLicenseToggleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostLicenseToggleRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostLicenseToggleUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostLicenseToggleTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostLicenseToggleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostLicenseToggleServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostLicenseToggleGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostLicenseToggleOK creates a PostLicenseToggleOK with default headers values
func NewPostLicenseToggleOK() *PostLicenseToggleOK {
	return &PostLicenseToggleOK{}
}

/*PostLicenseToggleOK handles this case with default header values.

successful operation
*/
type PostLicenseToggleOK struct {
	Payload *models.LicenseOrgToggle
}

func (o *PostLicenseToggleOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/toggles/{featureName}][%d] postLicenseToggleOK  %+v", 200, o.Payload)
}

func (o *PostLicenseToggleOK) GetPayload() *models.LicenseOrgToggle {
	return o.Payload
}

func (o *PostLicenseToggleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LicenseOrgToggle)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseToggleBadRequest creates a PostLicenseToggleBadRequest with default headers values
func NewPostLicenseToggleBadRequest() *PostLicenseToggleBadRequest {
	return &PostLicenseToggleBadRequest{}
}

/*PostLicenseToggleBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostLicenseToggleBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostLicenseToggleBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/toggles/{featureName}][%d] postLicenseToggleBadRequest  %+v", 400, o.Payload)
}

func (o *PostLicenseToggleBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseToggleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseToggleUnauthorized creates a PostLicenseToggleUnauthorized with default headers values
func NewPostLicenseToggleUnauthorized() *PostLicenseToggleUnauthorized {
	return &PostLicenseToggleUnauthorized{}
}

/*PostLicenseToggleUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostLicenseToggleUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostLicenseToggleUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/toggles/{featureName}][%d] postLicenseToggleUnauthorized  %+v", 401, o.Payload)
}

func (o *PostLicenseToggleUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseToggleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseToggleForbidden creates a PostLicenseToggleForbidden with default headers values
func NewPostLicenseToggleForbidden() *PostLicenseToggleForbidden {
	return &PostLicenseToggleForbidden{}
}

/*PostLicenseToggleForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostLicenseToggleForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostLicenseToggleForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/toggles/{featureName}][%d] postLicenseToggleForbidden  %+v", 403, o.Payload)
}

func (o *PostLicenseToggleForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseToggleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseToggleNotFound creates a PostLicenseToggleNotFound with default headers values
func NewPostLicenseToggleNotFound() *PostLicenseToggleNotFound {
	return &PostLicenseToggleNotFound{}
}

/*PostLicenseToggleNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostLicenseToggleNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostLicenseToggleNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/toggles/{featureName}][%d] postLicenseToggleNotFound  %+v", 404, o.Payload)
}

func (o *PostLicenseToggleNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseToggleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseToggleRequestEntityTooLarge creates a PostLicenseToggleRequestEntityTooLarge with default headers values
func NewPostLicenseToggleRequestEntityTooLarge() *PostLicenseToggleRequestEntityTooLarge {
	return &PostLicenseToggleRequestEntityTooLarge{}
}

/*PostLicenseToggleRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostLicenseToggleRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostLicenseToggleRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/toggles/{featureName}][%d] postLicenseToggleRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostLicenseToggleRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseToggleRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseToggleUnsupportedMediaType creates a PostLicenseToggleUnsupportedMediaType with default headers values
func NewPostLicenseToggleUnsupportedMediaType() *PostLicenseToggleUnsupportedMediaType {
	return &PostLicenseToggleUnsupportedMediaType{}
}

/*PostLicenseToggleUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostLicenseToggleUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostLicenseToggleUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/toggles/{featureName}][%d] postLicenseToggleUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostLicenseToggleUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseToggleUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseToggleTooManyRequests creates a PostLicenseToggleTooManyRequests with default headers values
func NewPostLicenseToggleTooManyRequests() *PostLicenseToggleTooManyRequests {
	return &PostLicenseToggleTooManyRequests{}
}

/*PostLicenseToggleTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostLicenseToggleTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostLicenseToggleTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/toggles/{featureName}][%d] postLicenseToggleTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostLicenseToggleTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseToggleTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseToggleInternalServerError creates a PostLicenseToggleInternalServerError with default headers values
func NewPostLicenseToggleInternalServerError() *PostLicenseToggleInternalServerError {
	return &PostLicenseToggleInternalServerError{}
}

/*PostLicenseToggleInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostLicenseToggleInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostLicenseToggleInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/toggles/{featureName}][%d] postLicenseToggleInternalServerError  %+v", 500, o.Payload)
}

func (o *PostLicenseToggleInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseToggleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseToggleServiceUnavailable creates a PostLicenseToggleServiceUnavailable with default headers values
func NewPostLicenseToggleServiceUnavailable() *PostLicenseToggleServiceUnavailable {
	return &PostLicenseToggleServiceUnavailable{}
}

/*PostLicenseToggleServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostLicenseToggleServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostLicenseToggleServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/toggles/{featureName}][%d] postLicenseToggleServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostLicenseToggleServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseToggleServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseToggleGatewayTimeout creates a PostLicenseToggleGatewayTimeout with default headers values
func NewPostLicenseToggleGatewayTimeout() *PostLicenseToggleGatewayTimeout {
	return &PostLicenseToggleGatewayTimeout{}
}

/*PostLicenseToggleGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostLicenseToggleGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostLicenseToggleGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/toggles/{featureName}][%d] postLicenseToggleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostLicenseToggleGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseToggleGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// PostLicenseOrganizationReader is a Reader for the PostLicenseOrganization structure.
type PostLicenseOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLicenseOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLicenseOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostLicenseOrganizationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostLicenseOrganizationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostLicenseOrganizationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostLicenseOrganizationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostLicenseOrganizationRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostLicenseOrganizationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostLicenseOrganizationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostLicenseOrganizationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostLicenseOrganizationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostLicenseOrganizationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostLicenseOrganizationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostLicenseOrganizationOK creates a PostLicenseOrganizationOK with default headers values
func NewPostLicenseOrganizationOK() *PostLicenseOrganizationOK {
	return &PostLicenseOrganizationOK{}
}

/*PostLicenseOrganizationOK handles this case with default header values.

successful operation
*/
type PostLicenseOrganizationOK struct {
	Payload []*models.LicenseUpdateStatus
}

func (o *PostLicenseOrganizationOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/organization][%d] postLicenseOrganizationOK  %+v", 200, o.Payload)
}

func (o *PostLicenseOrganizationOK) GetPayload() []*models.LicenseUpdateStatus {
	return o.Payload
}

func (o *PostLicenseOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseOrganizationBadRequest creates a PostLicenseOrganizationBadRequest with default headers values
func NewPostLicenseOrganizationBadRequest() *PostLicenseOrganizationBadRequest {
	return &PostLicenseOrganizationBadRequest{}
}

/*PostLicenseOrganizationBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostLicenseOrganizationBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostLicenseOrganizationBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/organization][%d] postLicenseOrganizationBadRequest  %+v", 400, o.Payload)
}

func (o *PostLicenseOrganizationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseOrganizationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseOrganizationUnauthorized creates a PostLicenseOrganizationUnauthorized with default headers values
func NewPostLicenseOrganizationUnauthorized() *PostLicenseOrganizationUnauthorized {
	return &PostLicenseOrganizationUnauthorized{}
}

/*PostLicenseOrganizationUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostLicenseOrganizationUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostLicenseOrganizationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/organization][%d] postLicenseOrganizationUnauthorized  %+v", 401, o.Payload)
}

func (o *PostLicenseOrganizationUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseOrganizationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseOrganizationForbidden creates a PostLicenseOrganizationForbidden with default headers values
func NewPostLicenseOrganizationForbidden() *PostLicenseOrganizationForbidden {
	return &PostLicenseOrganizationForbidden{}
}

/*PostLicenseOrganizationForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostLicenseOrganizationForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostLicenseOrganizationForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/organization][%d] postLicenseOrganizationForbidden  %+v", 403, o.Payload)
}

func (o *PostLicenseOrganizationForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseOrganizationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseOrganizationNotFound creates a PostLicenseOrganizationNotFound with default headers values
func NewPostLicenseOrganizationNotFound() *PostLicenseOrganizationNotFound {
	return &PostLicenseOrganizationNotFound{}
}

/*PostLicenseOrganizationNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostLicenseOrganizationNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostLicenseOrganizationNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/organization][%d] postLicenseOrganizationNotFound  %+v", 404, o.Payload)
}

func (o *PostLicenseOrganizationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseOrganizationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseOrganizationRequestTimeout creates a PostLicenseOrganizationRequestTimeout with default headers values
func NewPostLicenseOrganizationRequestTimeout() *PostLicenseOrganizationRequestTimeout {
	return &PostLicenseOrganizationRequestTimeout{}
}

/*PostLicenseOrganizationRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostLicenseOrganizationRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostLicenseOrganizationRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/organization][%d] postLicenseOrganizationRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostLicenseOrganizationRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseOrganizationRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseOrganizationRequestEntityTooLarge creates a PostLicenseOrganizationRequestEntityTooLarge with default headers values
func NewPostLicenseOrganizationRequestEntityTooLarge() *PostLicenseOrganizationRequestEntityTooLarge {
	return &PostLicenseOrganizationRequestEntityTooLarge{}
}

/*PostLicenseOrganizationRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostLicenseOrganizationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostLicenseOrganizationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/organization][%d] postLicenseOrganizationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostLicenseOrganizationRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseOrganizationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseOrganizationUnsupportedMediaType creates a PostLicenseOrganizationUnsupportedMediaType with default headers values
func NewPostLicenseOrganizationUnsupportedMediaType() *PostLicenseOrganizationUnsupportedMediaType {
	return &PostLicenseOrganizationUnsupportedMediaType{}
}

/*PostLicenseOrganizationUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostLicenseOrganizationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostLicenseOrganizationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/organization][%d] postLicenseOrganizationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostLicenseOrganizationUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseOrganizationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseOrganizationTooManyRequests creates a PostLicenseOrganizationTooManyRequests with default headers values
func NewPostLicenseOrganizationTooManyRequests() *PostLicenseOrganizationTooManyRequests {
	return &PostLicenseOrganizationTooManyRequests{}
}

/*PostLicenseOrganizationTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostLicenseOrganizationTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostLicenseOrganizationTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/organization][%d] postLicenseOrganizationTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostLicenseOrganizationTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseOrganizationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseOrganizationInternalServerError creates a PostLicenseOrganizationInternalServerError with default headers values
func NewPostLicenseOrganizationInternalServerError() *PostLicenseOrganizationInternalServerError {
	return &PostLicenseOrganizationInternalServerError{}
}

/*PostLicenseOrganizationInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostLicenseOrganizationInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostLicenseOrganizationInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/organization][%d] postLicenseOrganizationInternalServerError  %+v", 500, o.Payload)
}

func (o *PostLicenseOrganizationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseOrganizationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseOrganizationServiceUnavailable creates a PostLicenseOrganizationServiceUnavailable with default headers values
func NewPostLicenseOrganizationServiceUnavailable() *PostLicenseOrganizationServiceUnavailable {
	return &PostLicenseOrganizationServiceUnavailable{}
}

/*PostLicenseOrganizationServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostLicenseOrganizationServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostLicenseOrganizationServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/organization][%d] postLicenseOrganizationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostLicenseOrganizationServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseOrganizationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseOrganizationGatewayTimeout creates a PostLicenseOrganizationGatewayTimeout with default headers values
func NewPostLicenseOrganizationGatewayTimeout() *PostLicenseOrganizationGatewayTimeout {
	return &PostLicenseOrganizationGatewayTimeout{}
}

/*PostLicenseOrganizationGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostLicenseOrganizationGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostLicenseOrganizationGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/organization][%d] postLicenseOrganizationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostLicenseOrganizationGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseOrganizationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

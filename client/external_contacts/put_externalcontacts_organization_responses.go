// Code generated by go-swagger; DO NOT EDIT.

package external_contacts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PutExternalcontactsOrganizationReader is a Reader for the PutExternalcontactsOrganization structure.
type PutExternalcontactsOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutExternalcontactsOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutExternalcontactsOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutExternalcontactsOrganizationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutExternalcontactsOrganizationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutExternalcontactsOrganizationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutExternalcontactsOrganizationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutExternalcontactsOrganizationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutExternalcontactsOrganizationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPutExternalcontactsOrganizationUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutExternalcontactsOrganizationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutExternalcontactsOrganizationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutExternalcontactsOrganizationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutExternalcontactsOrganizationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutExternalcontactsOrganizationOK creates a PutExternalcontactsOrganizationOK with default headers values
func NewPutExternalcontactsOrganizationOK() *PutExternalcontactsOrganizationOK {
	return &PutExternalcontactsOrganizationOK{}
}

/*PutExternalcontactsOrganizationOK handles this case with default header values.

successful operation
*/
type PutExternalcontactsOrganizationOK struct {
	Payload *models.ExternalOrganization
}

func (o *PutExternalcontactsOrganizationOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/externalcontacts/organizations/{externalOrganizationId}][%d] putExternalcontactsOrganizationOK  %+v", 200, o.Payload)
}

func (o *PutExternalcontactsOrganizationOK) GetPayload() *models.ExternalOrganization {
	return o.Payload
}

func (o *PutExternalcontactsOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExternalOrganization)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutExternalcontactsOrganizationBadRequest creates a PutExternalcontactsOrganizationBadRequest with default headers values
func NewPutExternalcontactsOrganizationBadRequest() *PutExternalcontactsOrganizationBadRequest {
	return &PutExternalcontactsOrganizationBadRequest{}
}

/*PutExternalcontactsOrganizationBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutExternalcontactsOrganizationBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutExternalcontactsOrganizationBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/externalcontacts/organizations/{externalOrganizationId}][%d] putExternalcontactsOrganizationBadRequest  %+v", 400, o.Payload)
}

func (o *PutExternalcontactsOrganizationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutExternalcontactsOrganizationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutExternalcontactsOrganizationUnauthorized creates a PutExternalcontactsOrganizationUnauthorized with default headers values
func NewPutExternalcontactsOrganizationUnauthorized() *PutExternalcontactsOrganizationUnauthorized {
	return &PutExternalcontactsOrganizationUnauthorized{}
}

/*PutExternalcontactsOrganizationUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutExternalcontactsOrganizationUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutExternalcontactsOrganizationUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/externalcontacts/organizations/{externalOrganizationId}][%d] putExternalcontactsOrganizationUnauthorized  %+v", 401, o.Payload)
}

func (o *PutExternalcontactsOrganizationUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutExternalcontactsOrganizationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutExternalcontactsOrganizationForbidden creates a PutExternalcontactsOrganizationForbidden with default headers values
func NewPutExternalcontactsOrganizationForbidden() *PutExternalcontactsOrganizationForbidden {
	return &PutExternalcontactsOrganizationForbidden{}
}

/*PutExternalcontactsOrganizationForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutExternalcontactsOrganizationForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutExternalcontactsOrganizationForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/externalcontacts/organizations/{externalOrganizationId}][%d] putExternalcontactsOrganizationForbidden  %+v", 403, o.Payload)
}

func (o *PutExternalcontactsOrganizationForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutExternalcontactsOrganizationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutExternalcontactsOrganizationNotFound creates a PutExternalcontactsOrganizationNotFound with default headers values
func NewPutExternalcontactsOrganizationNotFound() *PutExternalcontactsOrganizationNotFound {
	return &PutExternalcontactsOrganizationNotFound{}
}

/*PutExternalcontactsOrganizationNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutExternalcontactsOrganizationNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutExternalcontactsOrganizationNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/externalcontacts/organizations/{externalOrganizationId}][%d] putExternalcontactsOrganizationNotFound  %+v", 404, o.Payload)
}

func (o *PutExternalcontactsOrganizationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutExternalcontactsOrganizationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutExternalcontactsOrganizationRequestEntityTooLarge creates a PutExternalcontactsOrganizationRequestEntityTooLarge with default headers values
func NewPutExternalcontactsOrganizationRequestEntityTooLarge() *PutExternalcontactsOrganizationRequestEntityTooLarge {
	return &PutExternalcontactsOrganizationRequestEntityTooLarge{}
}

/*PutExternalcontactsOrganizationRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PutExternalcontactsOrganizationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutExternalcontactsOrganizationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/externalcontacts/organizations/{externalOrganizationId}][%d] putExternalcontactsOrganizationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutExternalcontactsOrganizationRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutExternalcontactsOrganizationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutExternalcontactsOrganizationUnsupportedMediaType creates a PutExternalcontactsOrganizationUnsupportedMediaType with default headers values
func NewPutExternalcontactsOrganizationUnsupportedMediaType() *PutExternalcontactsOrganizationUnsupportedMediaType {
	return &PutExternalcontactsOrganizationUnsupportedMediaType{}
}

/*PutExternalcontactsOrganizationUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutExternalcontactsOrganizationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutExternalcontactsOrganizationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/externalcontacts/organizations/{externalOrganizationId}][%d] putExternalcontactsOrganizationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutExternalcontactsOrganizationUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutExternalcontactsOrganizationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutExternalcontactsOrganizationUnprocessableEntity creates a PutExternalcontactsOrganizationUnprocessableEntity with default headers values
func NewPutExternalcontactsOrganizationUnprocessableEntity() *PutExternalcontactsOrganizationUnprocessableEntity {
	return &PutExternalcontactsOrganizationUnprocessableEntity{}
}

/*PutExternalcontactsOrganizationUnprocessableEntity handles this case with default header values.

PutExternalcontactsOrganizationUnprocessableEntity put externalcontacts organization unprocessable entity
*/
type PutExternalcontactsOrganizationUnprocessableEntity struct {
	Payload *models.ErrorBody
}

func (o *PutExternalcontactsOrganizationUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /api/v2/externalcontacts/organizations/{externalOrganizationId}][%d] putExternalcontactsOrganizationUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PutExternalcontactsOrganizationUnprocessableEntity) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutExternalcontactsOrganizationUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutExternalcontactsOrganizationTooManyRequests creates a PutExternalcontactsOrganizationTooManyRequests with default headers values
func NewPutExternalcontactsOrganizationTooManyRequests() *PutExternalcontactsOrganizationTooManyRequests {
	return &PutExternalcontactsOrganizationTooManyRequests{}
}

/*PutExternalcontactsOrganizationTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PutExternalcontactsOrganizationTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutExternalcontactsOrganizationTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/externalcontacts/organizations/{externalOrganizationId}][%d] putExternalcontactsOrganizationTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutExternalcontactsOrganizationTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutExternalcontactsOrganizationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutExternalcontactsOrganizationInternalServerError creates a PutExternalcontactsOrganizationInternalServerError with default headers values
func NewPutExternalcontactsOrganizationInternalServerError() *PutExternalcontactsOrganizationInternalServerError {
	return &PutExternalcontactsOrganizationInternalServerError{}
}

/*PutExternalcontactsOrganizationInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutExternalcontactsOrganizationInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutExternalcontactsOrganizationInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/externalcontacts/organizations/{externalOrganizationId}][%d] putExternalcontactsOrganizationInternalServerError  %+v", 500, o.Payload)
}

func (o *PutExternalcontactsOrganizationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutExternalcontactsOrganizationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutExternalcontactsOrganizationServiceUnavailable creates a PutExternalcontactsOrganizationServiceUnavailable with default headers values
func NewPutExternalcontactsOrganizationServiceUnavailable() *PutExternalcontactsOrganizationServiceUnavailable {
	return &PutExternalcontactsOrganizationServiceUnavailable{}
}

/*PutExternalcontactsOrganizationServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutExternalcontactsOrganizationServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutExternalcontactsOrganizationServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/externalcontacts/organizations/{externalOrganizationId}][%d] putExternalcontactsOrganizationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutExternalcontactsOrganizationServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutExternalcontactsOrganizationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutExternalcontactsOrganizationGatewayTimeout creates a PutExternalcontactsOrganizationGatewayTimeout with default headers values
func NewPutExternalcontactsOrganizationGatewayTimeout() *PutExternalcontactsOrganizationGatewayTimeout {
	return &PutExternalcontactsOrganizationGatewayTimeout{}
}

/*PutExternalcontactsOrganizationGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutExternalcontactsOrganizationGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutExternalcontactsOrganizationGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/externalcontacts/organizations/{externalOrganizationId}][%d] putExternalcontactsOrganizationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutExternalcontactsOrganizationGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutExternalcontactsOrganizationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
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

// GetExternalcontactsOrganizationContactsReader is a Reader for the GetExternalcontactsOrganizationContacts structure.
type GetExternalcontactsOrganizationContactsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExternalcontactsOrganizationContactsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExternalcontactsOrganizationContactsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetExternalcontactsOrganizationContactsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetExternalcontactsOrganizationContactsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetExternalcontactsOrganizationContactsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetExternalcontactsOrganizationContactsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetExternalcontactsOrganizationContactsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetExternalcontactsOrganizationContactsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetExternalcontactsOrganizationContactsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetExternalcontactsOrganizationContactsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetExternalcontactsOrganizationContactsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetExternalcontactsOrganizationContactsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetExternalcontactsOrganizationContactsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetExternalcontactsOrganizationContactsOK creates a GetExternalcontactsOrganizationContactsOK with default headers values
func NewGetExternalcontactsOrganizationContactsOK() *GetExternalcontactsOrganizationContactsOK {
	return &GetExternalcontactsOrganizationContactsOK{}
}

/*GetExternalcontactsOrganizationContactsOK handles this case with default header values.

successful operation
*/
type GetExternalcontactsOrganizationContactsOK struct {
	Payload *models.ContactListing
}

func (o *GetExternalcontactsOrganizationContactsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/contacts][%d] getExternalcontactsOrganizationContactsOK  %+v", 200, o.Payload)
}

func (o *GetExternalcontactsOrganizationContactsOK) GetPayload() *models.ContactListing {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationContactsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContactListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationContactsBadRequest creates a GetExternalcontactsOrganizationContactsBadRequest with default headers values
func NewGetExternalcontactsOrganizationContactsBadRequest() *GetExternalcontactsOrganizationContactsBadRequest {
	return &GetExternalcontactsOrganizationContactsBadRequest{}
}

/*GetExternalcontactsOrganizationContactsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetExternalcontactsOrganizationContactsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsOrganizationContactsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/contacts][%d] getExternalcontactsOrganizationContactsBadRequest  %+v", 400, o.Payload)
}

func (o *GetExternalcontactsOrganizationContactsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationContactsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationContactsUnauthorized creates a GetExternalcontactsOrganizationContactsUnauthorized with default headers values
func NewGetExternalcontactsOrganizationContactsUnauthorized() *GetExternalcontactsOrganizationContactsUnauthorized {
	return &GetExternalcontactsOrganizationContactsUnauthorized{}
}

/*GetExternalcontactsOrganizationContactsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetExternalcontactsOrganizationContactsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsOrganizationContactsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/contacts][%d] getExternalcontactsOrganizationContactsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetExternalcontactsOrganizationContactsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationContactsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationContactsForbidden creates a GetExternalcontactsOrganizationContactsForbidden with default headers values
func NewGetExternalcontactsOrganizationContactsForbidden() *GetExternalcontactsOrganizationContactsForbidden {
	return &GetExternalcontactsOrganizationContactsForbidden{}
}

/*GetExternalcontactsOrganizationContactsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetExternalcontactsOrganizationContactsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsOrganizationContactsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/contacts][%d] getExternalcontactsOrganizationContactsForbidden  %+v", 403, o.Payload)
}

func (o *GetExternalcontactsOrganizationContactsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationContactsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationContactsNotFound creates a GetExternalcontactsOrganizationContactsNotFound with default headers values
func NewGetExternalcontactsOrganizationContactsNotFound() *GetExternalcontactsOrganizationContactsNotFound {
	return &GetExternalcontactsOrganizationContactsNotFound{}
}

/*GetExternalcontactsOrganizationContactsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetExternalcontactsOrganizationContactsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsOrganizationContactsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/contacts][%d] getExternalcontactsOrganizationContactsNotFound  %+v", 404, o.Payload)
}

func (o *GetExternalcontactsOrganizationContactsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationContactsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationContactsRequestTimeout creates a GetExternalcontactsOrganizationContactsRequestTimeout with default headers values
func NewGetExternalcontactsOrganizationContactsRequestTimeout() *GetExternalcontactsOrganizationContactsRequestTimeout {
	return &GetExternalcontactsOrganizationContactsRequestTimeout{}
}

/*GetExternalcontactsOrganizationContactsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetExternalcontactsOrganizationContactsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsOrganizationContactsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/contacts][%d] getExternalcontactsOrganizationContactsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetExternalcontactsOrganizationContactsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationContactsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationContactsRequestEntityTooLarge creates a GetExternalcontactsOrganizationContactsRequestEntityTooLarge with default headers values
func NewGetExternalcontactsOrganizationContactsRequestEntityTooLarge() *GetExternalcontactsOrganizationContactsRequestEntityTooLarge {
	return &GetExternalcontactsOrganizationContactsRequestEntityTooLarge{}
}

/*GetExternalcontactsOrganizationContactsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetExternalcontactsOrganizationContactsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsOrganizationContactsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/contacts][%d] getExternalcontactsOrganizationContactsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetExternalcontactsOrganizationContactsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationContactsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationContactsUnsupportedMediaType creates a GetExternalcontactsOrganizationContactsUnsupportedMediaType with default headers values
func NewGetExternalcontactsOrganizationContactsUnsupportedMediaType() *GetExternalcontactsOrganizationContactsUnsupportedMediaType {
	return &GetExternalcontactsOrganizationContactsUnsupportedMediaType{}
}

/*GetExternalcontactsOrganizationContactsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetExternalcontactsOrganizationContactsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsOrganizationContactsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/contacts][%d] getExternalcontactsOrganizationContactsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetExternalcontactsOrganizationContactsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationContactsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationContactsTooManyRequests creates a GetExternalcontactsOrganizationContactsTooManyRequests with default headers values
func NewGetExternalcontactsOrganizationContactsTooManyRequests() *GetExternalcontactsOrganizationContactsTooManyRequests {
	return &GetExternalcontactsOrganizationContactsTooManyRequests{}
}

/*GetExternalcontactsOrganizationContactsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetExternalcontactsOrganizationContactsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsOrganizationContactsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/contacts][%d] getExternalcontactsOrganizationContactsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetExternalcontactsOrganizationContactsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationContactsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationContactsInternalServerError creates a GetExternalcontactsOrganizationContactsInternalServerError with default headers values
func NewGetExternalcontactsOrganizationContactsInternalServerError() *GetExternalcontactsOrganizationContactsInternalServerError {
	return &GetExternalcontactsOrganizationContactsInternalServerError{}
}

/*GetExternalcontactsOrganizationContactsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetExternalcontactsOrganizationContactsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsOrganizationContactsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/contacts][%d] getExternalcontactsOrganizationContactsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetExternalcontactsOrganizationContactsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationContactsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationContactsServiceUnavailable creates a GetExternalcontactsOrganizationContactsServiceUnavailable with default headers values
func NewGetExternalcontactsOrganizationContactsServiceUnavailable() *GetExternalcontactsOrganizationContactsServiceUnavailable {
	return &GetExternalcontactsOrganizationContactsServiceUnavailable{}
}

/*GetExternalcontactsOrganizationContactsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetExternalcontactsOrganizationContactsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsOrganizationContactsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/contacts][%d] getExternalcontactsOrganizationContactsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetExternalcontactsOrganizationContactsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationContactsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationContactsGatewayTimeout creates a GetExternalcontactsOrganizationContactsGatewayTimeout with default headers values
func NewGetExternalcontactsOrganizationContactsGatewayTimeout() *GetExternalcontactsOrganizationContactsGatewayTimeout {
	return &GetExternalcontactsOrganizationContactsGatewayTimeout{}
}

/*GetExternalcontactsOrganizationContactsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetExternalcontactsOrganizationContactsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsOrganizationContactsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/contacts][%d] getExternalcontactsOrganizationContactsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetExternalcontactsOrganizationContactsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationContactsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// GetExternalcontactsScanContactsReader is a Reader for the GetExternalcontactsScanContacts structure.
type GetExternalcontactsScanContactsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExternalcontactsScanContactsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExternalcontactsScanContactsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetExternalcontactsScanContactsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetExternalcontactsScanContactsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetExternalcontactsScanContactsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetExternalcontactsScanContactsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetExternalcontactsScanContactsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetExternalcontactsScanContactsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetExternalcontactsScanContactsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetExternalcontactsScanContactsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetExternalcontactsScanContactsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetExternalcontactsScanContactsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetExternalcontactsScanContactsOK creates a GetExternalcontactsScanContactsOK with default headers values
func NewGetExternalcontactsScanContactsOK() *GetExternalcontactsScanContactsOK {
	return &GetExternalcontactsScanContactsOK{}
}

/*GetExternalcontactsScanContactsOK handles this case with default header values.

successful operation
*/
type GetExternalcontactsScanContactsOK struct {
	Payload *models.CursorContactListing
}

func (o *GetExternalcontactsScanContactsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/contacts][%d] getExternalcontactsScanContactsOK  %+v", 200, o.Payload)
}

func (o *GetExternalcontactsScanContactsOK) GetPayload() *models.CursorContactListing {
	return o.Payload
}

func (o *GetExternalcontactsScanContactsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CursorContactListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsScanContactsBadRequest creates a GetExternalcontactsScanContactsBadRequest with default headers values
func NewGetExternalcontactsScanContactsBadRequest() *GetExternalcontactsScanContactsBadRequest {
	return &GetExternalcontactsScanContactsBadRequest{}
}

/*GetExternalcontactsScanContactsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetExternalcontactsScanContactsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsScanContactsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/contacts][%d] getExternalcontactsScanContactsBadRequest  %+v", 400, o.Payload)
}

func (o *GetExternalcontactsScanContactsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsScanContactsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsScanContactsUnauthorized creates a GetExternalcontactsScanContactsUnauthorized with default headers values
func NewGetExternalcontactsScanContactsUnauthorized() *GetExternalcontactsScanContactsUnauthorized {
	return &GetExternalcontactsScanContactsUnauthorized{}
}

/*GetExternalcontactsScanContactsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetExternalcontactsScanContactsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsScanContactsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/contacts][%d] getExternalcontactsScanContactsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetExternalcontactsScanContactsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsScanContactsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsScanContactsForbidden creates a GetExternalcontactsScanContactsForbidden with default headers values
func NewGetExternalcontactsScanContactsForbidden() *GetExternalcontactsScanContactsForbidden {
	return &GetExternalcontactsScanContactsForbidden{}
}

/*GetExternalcontactsScanContactsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetExternalcontactsScanContactsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsScanContactsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/contacts][%d] getExternalcontactsScanContactsForbidden  %+v", 403, o.Payload)
}

func (o *GetExternalcontactsScanContactsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsScanContactsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsScanContactsNotFound creates a GetExternalcontactsScanContactsNotFound with default headers values
func NewGetExternalcontactsScanContactsNotFound() *GetExternalcontactsScanContactsNotFound {
	return &GetExternalcontactsScanContactsNotFound{}
}

/*GetExternalcontactsScanContactsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetExternalcontactsScanContactsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsScanContactsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/contacts][%d] getExternalcontactsScanContactsNotFound  %+v", 404, o.Payload)
}

func (o *GetExternalcontactsScanContactsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsScanContactsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsScanContactsRequestEntityTooLarge creates a GetExternalcontactsScanContactsRequestEntityTooLarge with default headers values
func NewGetExternalcontactsScanContactsRequestEntityTooLarge() *GetExternalcontactsScanContactsRequestEntityTooLarge {
	return &GetExternalcontactsScanContactsRequestEntityTooLarge{}
}

/*GetExternalcontactsScanContactsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetExternalcontactsScanContactsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsScanContactsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/contacts][%d] getExternalcontactsScanContactsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetExternalcontactsScanContactsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsScanContactsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsScanContactsUnsupportedMediaType creates a GetExternalcontactsScanContactsUnsupportedMediaType with default headers values
func NewGetExternalcontactsScanContactsUnsupportedMediaType() *GetExternalcontactsScanContactsUnsupportedMediaType {
	return &GetExternalcontactsScanContactsUnsupportedMediaType{}
}

/*GetExternalcontactsScanContactsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetExternalcontactsScanContactsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsScanContactsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/contacts][%d] getExternalcontactsScanContactsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetExternalcontactsScanContactsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsScanContactsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsScanContactsTooManyRequests creates a GetExternalcontactsScanContactsTooManyRequests with default headers values
func NewGetExternalcontactsScanContactsTooManyRequests() *GetExternalcontactsScanContactsTooManyRequests {
	return &GetExternalcontactsScanContactsTooManyRequests{}
}

/*GetExternalcontactsScanContactsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetExternalcontactsScanContactsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsScanContactsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/contacts][%d] getExternalcontactsScanContactsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetExternalcontactsScanContactsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsScanContactsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsScanContactsInternalServerError creates a GetExternalcontactsScanContactsInternalServerError with default headers values
func NewGetExternalcontactsScanContactsInternalServerError() *GetExternalcontactsScanContactsInternalServerError {
	return &GetExternalcontactsScanContactsInternalServerError{}
}

/*GetExternalcontactsScanContactsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetExternalcontactsScanContactsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsScanContactsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/contacts][%d] getExternalcontactsScanContactsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetExternalcontactsScanContactsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsScanContactsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsScanContactsServiceUnavailable creates a GetExternalcontactsScanContactsServiceUnavailable with default headers values
func NewGetExternalcontactsScanContactsServiceUnavailable() *GetExternalcontactsScanContactsServiceUnavailable {
	return &GetExternalcontactsScanContactsServiceUnavailable{}
}

/*GetExternalcontactsScanContactsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetExternalcontactsScanContactsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsScanContactsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/contacts][%d] getExternalcontactsScanContactsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetExternalcontactsScanContactsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsScanContactsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsScanContactsGatewayTimeout creates a GetExternalcontactsScanContactsGatewayTimeout with default headers values
func NewGetExternalcontactsScanContactsGatewayTimeout() *GetExternalcontactsScanContactsGatewayTimeout {
	return &GetExternalcontactsScanContactsGatewayTimeout{}
}

/*GetExternalcontactsScanContactsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetExternalcontactsScanContactsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsScanContactsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/contacts][%d] getExternalcontactsScanContactsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetExternalcontactsScanContactsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsScanContactsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
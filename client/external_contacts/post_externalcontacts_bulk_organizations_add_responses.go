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

// PostExternalcontactsBulkOrganizationsAddReader is a Reader for the PostExternalcontactsBulkOrganizationsAdd structure.
type PostExternalcontactsBulkOrganizationsAddReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostExternalcontactsBulkOrganizationsAddReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostExternalcontactsBulkOrganizationsAddOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostExternalcontactsBulkOrganizationsAddBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostExternalcontactsBulkOrganizationsAddUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostExternalcontactsBulkOrganizationsAddForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostExternalcontactsBulkOrganizationsAddNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostExternalcontactsBulkOrganizationsAddRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostExternalcontactsBulkOrganizationsAddUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPostExternalcontactsBulkOrganizationsAddUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostExternalcontactsBulkOrganizationsAddTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostExternalcontactsBulkOrganizationsAddInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostExternalcontactsBulkOrganizationsAddServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostExternalcontactsBulkOrganizationsAddGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostExternalcontactsBulkOrganizationsAddOK creates a PostExternalcontactsBulkOrganizationsAddOK with default headers values
func NewPostExternalcontactsBulkOrganizationsAddOK() *PostExternalcontactsBulkOrganizationsAddOK {
	return &PostExternalcontactsBulkOrganizationsAddOK{}
}

/*PostExternalcontactsBulkOrganizationsAddOK handles this case with default header values.

successful operation
*/
type PostExternalcontactsBulkOrganizationsAddOK struct {
	Payload *models.BulkOrganizationsResponse
}

func (o *PostExternalcontactsBulkOrganizationsAddOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations/add][%d] postExternalcontactsBulkOrganizationsAddOK  %+v", 200, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsAddOK) GetPayload() *models.BulkOrganizationsResponse {
	return o.Payload
}

func (o *PostExternalcontactsBulkOrganizationsAddOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BulkOrganizationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkOrganizationsAddBadRequest creates a PostExternalcontactsBulkOrganizationsAddBadRequest with default headers values
func NewPostExternalcontactsBulkOrganizationsAddBadRequest() *PostExternalcontactsBulkOrganizationsAddBadRequest {
	return &PostExternalcontactsBulkOrganizationsAddBadRequest{}
}

/*PostExternalcontactsBulkOrganizationsAddBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostExternalcontactsBulkOrganizationsAddBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkOrganizationsAddBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations/add][%d] postExternalcontactsBulkOrganizationsAddBadRequest  %+v", 400, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsAddBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkOrganizationsAddBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkOrganizationsAddUnauthorized creates a PostExternalcontactsBulkOrganizationsAddUnauthorized with default headers values
func NewPostExternalcontactsBulkOrganizationsAddUnauthorized() *PostExternalcontactsBulkOrganizationsAddUnauthorized {
	return &PostExternalcontactsBulkOrganizationsAddUnauthorized{}
}

/*PostExternalcontactsBulkOrganizationsAddUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostExternalcontactsBulkOrganizationsAddUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkOrganizationsAddUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations/add][%d] postExternalcontactsBulkOrganizationsAddUnauthorized  %+v", 401, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsAddUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkOrganizationsAddUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkOrganizationsAddForbidden creates a PostExternalcontactsBulkOrganizationsAddForbidden with default headers values
func NewPostExternalcontactsBulkOrganizationsAddForbidden() *PostExternalcontactsBulkOrganizationsAddForbidden {
	return &PostExternalcontactsBulkOrganizationsAddForbidden{}
}

/*PostExternalcontactsBulkOrganizationsAddForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostExternalcontactsBulkOrganizationsAddForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkOrganizationsAddForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations/add][%d] postExternalcontactsBulkOrganizationsAddForbidden  %+v", 403, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsAddForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkOrganizationsAddForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkOrganizationsAddNotFound creates a PostExternalcontactsBulkOrganizationsAddNotFound with default headers values
func NewPostExternalcontactsBulkOrganizationsAddNotFound() *PostExternalcontactsBulkOrganizationsAddNotFound {
	return &PostExternalcontactsBulkOrganizationsAddNotFound{}
}

/*PostExternalcontactsBulkOrganizationsAddNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostExternalcontactsBulkOrganizationsAddNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkOrganizationsAddNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations/add][%d] postExternalcontactsBulkOrganizationsAddNotFound  %+v", 404, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsAddNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkOrganizationsAddNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkOrganizationsAddRequestEntityTooLarge creates a PostExternalcontactsBulkOrganizationsAddRequestEntityTooLarge with default headers values
func NewPostExternalcontactsBulkOrganizationsAddRequestEntityTooLarge() *PostExternalcontactsBulkOrganizationsAddRequestEntityTooLarge {
	return &PostExternalcontactsBulkOrganizationsAddRequestEntityTooLarge{}
}

/*PostExternalcontactsBulkOrganizationsAddRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostExternalcontactsBulkOrganizationsAddRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkOrganizationsAddRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations/add][%d] postExternalcontactsBulkOrganizationsAddRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsAddRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkOrganizationsAddRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkOrganizationsAddUnsupportedMediaType creates a PostExternalcontactsBulkOrganizationsAddUnsupportedMediaType with default headers values
func NewPostExternalcontactsBulkOrganizationsAddUnsupportedMediaType() *PostExternalcontactsBulkOrganizationsAddUnsupportedMediaType {
	return &PostExternalcontactsBulkOrganizationsAddUnsupportedMediaType{}
}

/*PostExternalcontactsBulkOrganizationsAddUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostExternalcontactsBulkOrganizationsAddUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkOrganizationsAddUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations/add][%d] postExternalcontactsBulkOrganizationsAddUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsAddUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkOrganizationsAddUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkOrganizationsAddUnprocessableEntity creates a PostExternalcontactsBulkOrganizationsAddUnprocessableEntity with default headers values
func NewPostExternalcontactsBulkOrganizationsAddUnprocessableEntity() *PostExternalcontactsBulkOrganizationsAddUnprocessableEntity {
	return &PostExternalcontactsBulkOrganizationsAddUnprocessableEntity{}
}

/*PostExternalcontactsBulkOrganizationsAddUnprocessableEntity handles this case with default header values.

PostExternalcontactsBulkOrganizationsAddUnprocessableEntity post externalcontacts bulk organizations add unprocessable entity
*/
type PostExternalcontactsBulkOrganizationsAddUnprocessableEntity struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkOrganizationsAddUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations/add][%d] postExternalcontactsBulkOrganizationsAddUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsAddUnprocessableEntity) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkOrganizationsAddUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkOrganizationsAddTooManyRequests creates a PostExternalcontactsBulkOrganizationsAddTooManyRequests with default headers values
func NewPostExternalcontactsBulkOrganizationsAddTooManyRequests() *PostExternalcontactsBulkOrganizationsAddTooManyRequests {
	return &PostExternalcontactsBulkOrganizationsAddTooManyRequests{}
}

/*PostExternalcontactsBulkOrganizationsAddTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostExternalcontactsBulkOrganizationsAddTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkOrganizationsAddTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations/add][%d] postExternalcontactsBulkOrganizationsAddTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsAddTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkOrganizationsAddTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkOrganizationsAddInternalServerError creates a PostExternalcontactsBulkOrganizationsAddInternalServerError with default headers values
func NewPostExternalcontactsBulkOrganizationsAddInternalServerError() *PostExternalcontactsBulkOrganizationsAddInternalServerError {
	return &PostExternalcontactsBulkOrganizationsAddInternalServerError{}
}

/*PostExternalcontactsBulkOrganizationsAddInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostExternalcontactsBulkOrganizationsAddInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkOrganizationsAddInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations/add][%d] postExternalcontactsBulkOrganizationsAddInternalServerError  %+v", 500, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsAddInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkOrganizationsAddInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkOrganizationsAddServiceUnavailable creates a PostExternalcontactsBulkOrganizationsAddServiceUnavailable with default headers values
func NewPostExternalcontactsBulkOrganizationsAddServiceUnavailable() *PostExternalcontactsBulkOrganizationsAddServiceUnavailable {
	return &PostExternalcontactsBulkOrganizationsAddServiceUnavailable{}
}

/*PostExternalcontactsBulkOrganizationsAddServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostExternalcontactsBulkOrganizationsAddServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkOrganizationsAddServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations/add][%d] postExternalcontactsBulkOrganizationsAddServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsAddServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkOrganizationsAddServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkOrganizationsAddGatewayTimeout creates a PostExternalcontactsBulkOrganizationsAddGatewayTimeout with default headers values
func NewPostExternalcontactsBulkOrganizationsAddGatewayTimeout() *PostExternalcontactsBulkOrganizationsAddGatewayTimeout {
	return &PostExternalcontactsBulkOrganizationsAddGatewayTimeout{}
}

/*PostExternalcontactsBulkOrganizationsAddGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostExternalcontactsBulkOrganizationsAddGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkOrganizationsAddGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations/add][%d] postExternalcontactsBulkOrganizationsAddGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsAddGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkOrganizationsAddGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
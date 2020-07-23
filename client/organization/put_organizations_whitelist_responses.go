// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PutOrganizationsWhitelistReader is a Reader for the PutOrganizationsWhitelist structure.
type PutOrganizationsWhitelistReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutOrganizationsWhitelistReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutOrganizationsWhitelistOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutOrganizationsWhitelistBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutOrganizationsWhitelistUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutOrganizationsWhitelistForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutOrganizationsWhitelistNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutOrganizationsWhitelistRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutOrganizationsWhitelistUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutOrganizationsWhitelistTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutOrganizationsWhitelistInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutOrganizationsWhitelistServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutOrganizationsWhitelistGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutOrganizationsWhitelistOK creates a PutOrganizationsWhitelistOK with default headers values
func NewPutOrganizationsWhitelistOK() *PutOrganizationsWhitelistOK {
	return &PutOrganizationsWhitelistOK{}
}

/*PutOrganizationsWhitelistOK handles this case with default header values.

successful operation
*/
type PutOrganizationsWhitelistOK struct {
	Payload *models.OrgWhitelistSettings
}

func (o *PutOrganizationsWhitelistOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/organizations/whitelist][%d] putOrganizationsWhitelistOK  %+v", 200, o.Payload)
}

func (o *PutOrganizationsWhitelistOK) GetPayload() *models.OrgWhitelistSettings {
	return o.Payload
}

func (o *PutOrganizationsWhitelistOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrgWhitelistSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrganizationsWhitelistBadRequest creates a PutOrganizationsWhitelistBadRequest with default headers values
func NewPutOrganizationsWhitelistBadRequest() *PutOrganizationsWhitelistBadRequest {
	return &PutOrganizationsWhitelistBadRequest{}
}

/*PutOrganizationsWhitelistBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutOrganizationsWhitelistBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutOrganizationsWhitelistBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/organizations/whitelist][%d] putOrganizationsWhitelistBadRequest  %+v", 400, o.Payload)
}

func (o *PutOrganizationsWhitelistBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOrganizationsWhitelistBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrganizationsWhitelistUnauthorized creates a PutOrganizationsWhitelistUnauthorized with default headers values
func NewPutOrganizationsWhitelistUnauthorized() *PutOrganizationsWhitelistUnauthorized {
	return &PutOrganizationsWhitelistUnauthorized{}
}

/*PutOrganizationsWhitelistUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutOrganizationsWhitelistUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutOrganizationsWhitelistUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/organizations/whitelist][%d] putOrganizationsWhitelistUnauthorized  %+v", 401, o.Payload)
}

func (o *PutOrganizationsWhitelistUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOrganizationsWhitelistUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrganizationsWhitelistForbidden creates a PutOrganizationsWhitelistForbidden with default headers values
func NewPutOrganizationsWhitelistForbidden() *PutOrganizationsWhitelistForbidden {
	return &PutOrganizationsWhitelistForbidden{}
}

/*PutOrganizationsWhitelistForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutOrganizationsWhitelistForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutOrganizationsWhitelistForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/organizations/whitelist][%d] putOrganizationsWhitelistForbidden  %+v", 403, o.Payload)
}

func (o *PutOrganizationsWhitelistForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOrganizationsWhitelistForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrganizationsWhitelistNotFound creates a PutOrganizationsWhitelistNotFound with default headers values
func NewPutOrganizationsWhitelistNotFound() *PutOrganizationsWhitelistNotFound {
	return &PutOrganizationsWhitelistNotFound{}
}

/*PutOrganizationsWhitelistNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutOrganizationsWhitelistNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutOrganizationsWhitelistNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/organizations/whitelist][%d] putOrganizationsWhitelistNotFound  %+v", 404, o.Payload)
}

func (o *PutOrganizationsWhitelistNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOrganizationsWhitelistNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrganizationsWhitelistRequestEntityTooLarge creates a PutOrganizationsWhitelistRequestEntityTooLarge with default headers values
func NewPutOrganizationsWhitelistRequestEntityTooLarge() *PutOrganizationsWhitelistRequestEntityTooLarge {
	return &PutOrganizationsWhitelistRequestEntityTooLarge{}
}

/*PutOrganizationsWhitelistRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PutOrganizationsWhitelistRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutOrganizationsWhitelistRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/organizations/whitelist][%d] putOrganizationsWhitelistRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutOrganizationsWhitelistRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOrganizationsWhitelistRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrganizationsWhitelistUnsupportedMediaType creates a PutOrganizationsWhitelistUnsupportedMediaType with default headers values
func NewPutOrganizationsWhitelistUnsupportedMediaType() *PutOrganizationsWhitelistUnsupportedMediaType {
	return &PutOrganizationsWhitelistUnsupportedMediaType{}
}

/*PutOrganizationsWhitelistUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutOrganizationsWhitelistUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutOrganizationsWhitelistUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/organizations/whitelist][%d] putOrganizationsWhitelistUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutOrganizationsWhitelistUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOrganizationsWhitelistUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrganizationsWhitelistTooManyRequests creates a PutOrganizationsWhitelistTooManyRequests with default headers values
func NewPutOrganizationsWhitelistTooManyRequests() *PutOrganizationsWhitelistTooManyRequests {
	return &PutOrganizationsWhitelistTooManyRequests{}
}

/*PutOrganizationsWhitelistTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PutOrganizationsWhitelistTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutOrganizationsWhitelistTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/organizations/whitelist][%d] putOrganizationsWhitelistTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutOrganizationsWhitelistTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOrganizationsWhitelistTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrganizationsWhitelistInternalServerError creates a PutOrganizationsWhitelistInternalServerError with default headers values
func NewPutOrganizationsWhitelistInternalServerError() *PutOrganizationsWhitelistInternalServerError {
	return &PutOrganizationsWhitelistInternalServerError{}
}

/*PutOrganizationsWhitelistInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutOrganizationsWhitelistInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutOrganizationsWhitelistInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/organizations/whitelist][%d] putOrganizationsWhitelistInternalServerError  %+v", 500, o.Payload)
}

func (o *PutOrganizationsWhitelistInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOrganizationsWhitelistInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrganizationsWhitelistServiceUnavailable creates a PutOrganizationsWhitelistServiceUnavailable with default headers values
func NewPutOrganizationsWhitelistServiceUnavailable() *PutOrganizationsWhitelistServiceUnavailable {
	return &PutOrganizationsWhitelistServiceUnavailable{}
}

/*PutOrganizationsWhitelistServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutOrganizationsWhitelistServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutOrganizationsWhitelistServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/organizations/whitelist][%d] putOrganizationsWhitelistServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutOrganizationsWhitelistServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOrganizationsWhitelistServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrganizationsWhitelistGatewayTimeout creates a PutOrganizationsWhitelistGatewayTimeout with default headers values
func NewPutOrganizationsWhitelistGatewayTimeout() *PutOrganizationsWhitelistGatewayTimeout {
	return &PutOrganizationsWhitelistGatewayTimeout{}
}

/*PutOrganizationsWhitelistGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutOrganizationsWhitelistGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutOrganizationsWhitelistGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/organizations/whitelist][%d] putOrganizationsWhitelistGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutOrganizationsWhitelistGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOrganizationsWhitelistGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

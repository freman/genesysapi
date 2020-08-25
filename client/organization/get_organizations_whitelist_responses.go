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

// GetOrganizationsWhitelistReader is a Reader for the GetOrganizationsWhitelist structure.
type GetOrganizationsWhitelistReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationsWhitelistReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationsWhitelistOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOrganizationsWhitelistBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOrganizationsWhitelistUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrganizationsWhitelistForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrganizationsWhitelistNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOrganizationsWhitelistRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOrganizationsWhitelistUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOrganizationsWhitelistTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOrganizationsWhitelistInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOrganizationsWhitelistServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOrganizationsWhitelistGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrganizationsWhitelistOK creates a GetOrganizationsWhitelistOK with default headers values
func NewGetOrganizationsWhitelistOK() *GetOrganizationsWhitelistOK {
	return &GetOrganizationsWhitelistOK{}
}

/*GetOrganizationsWhitelistOK handles this case with default header values.

successful operation
*/
type GetOrganizationsWhitelistOK struct {
	Payload *models.OrgWhitelistSettings
}

func (o *GetOrganizationsWhitelistOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/whitelist][%d] getOrganizationsWhitelistOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationsWhitelistOK) GetPayload() *models.OrgWhitelistSettings {
	return o.Payload
}

func (o *GetOrganizationsWhitelistOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrgWhitelistSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsWhitelistBadRequest creates a GetOrganizationsWhitelistBadRequest with default headers values
func NewGetOrganizationsWhitelistBadRequest() *GetOrganizationsWhitelistBadRequest {
	return &GetOrganizationsWhitelistBadRequest{}
}

/*GetOrganizationsWhitelistBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOrganizationsWhitelistBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetOrganizationsWhitelistBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/whitelist][%d] getOrganizationsWhitelistBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrganizationsWhitelistBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsWhitelistBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsWhitelistUnauthorized creates a GetOrganizationsWhitelistUnauthorized with default headers values
func NewGetOrganizationsWhitelistUnauthorized() *GetOrganizationsWhitelistUnauthorized {
	return &GetOrganizationsWhitelistUnauthorized{}
}

/*GetOrganizationsWhitelistUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOrganizationsWhitelistUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetOrganizationsWhitelistUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/whitelist][%d] getOrganizationsWhitelistUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOrganizationsWhitelistUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsWhitelistUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsWhitelistForbidden creates a GetOrganizationsWhitelistForbidden with default headers values
func NewGetOrganizationsWhitelistForbidden() *GetOrganizationsWhitelistForbidden {
	return &GetOrganizationsWhitelistForbidden{}
}

/*GetOrganizationsWhitelistForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetOrganizationsWhitelistForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetOrganizationsWhitelistForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/whitelist][%d] getOrganizationsWhitelistForbidden  %+v", 403, o.Payload)
}

func (o *GetOrganizationsWhitelistForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsWhitelistForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsWhitelistNotFound creates a GetOrganizationsWhitelistNotFound with default headers values
func NewGetOrganizationsWhitelistNotFound() *GetOrganizationsWhitelistNotFound {
	return &GetOrganizationsWhitelistNotFound{}
}

/*GetOrganizationsWhitelistNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetOrganizationsWhitelistNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetOrganizationsWhitelistNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/whitelist][%d] getOrganizationsWhitelistNotFound  %+v", 404, o.Payload)
}

func (o *GetOrganizationsWhitelistNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsWhitelistNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsWhitelistRequestEntityTooLarge creates a GetOrganizationsWhitelistRequestEntityTooLarge with default headers values
func NewGetOrganizationsWhitelistRequestEntityTooLarge() *GetOrganizationsWhitelistRequestEntityTooLarge {
	return &GetOrganizationsWhitelistRequestEntityTooLarge{}
}

/*GetOrganizationsWhitelistRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetOrganizationsWhitelistRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetOrganizationsWhitelistRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/whitelist][%d] getOrganizationsWhitelistRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOrganizationsWhitelistRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsWhitelistRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsWhitelistUnsupportedMediaType creates a GetOrganizationsWhitelistUnsupportedMediaType with default headers values
func NewGetOrganizationsWhitelistUnsupportedMediaType() *GetOrganizationsWhitelistUnsupportedMediaType {
	return &GetOrganizationsWhitelistUnsupportedMediaType{}
}

/*GetOrganizationsWhitelistUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOrganizationsWhitelistUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetOrganizationsWhitelistUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/whitelist][%d] getOrganizationsWhitelistUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOrganizationsWhitelistUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsWhitelistUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsWhitelistTooManyRequests creates a GetOrganizationsWhitelistTooManyRequests with default headers values
func NewGetOrganizationsWhitelistTooManyRequests() *GetOrganizationsWhitelistTooManyRequests {
	return &GetOrganizationsWhitelistTooManyRequests{}
}

/*GetOrganizationsWhitelistTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetOrganizationsWhitelistTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetOrganizationsWhitelistTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/whitelist][%d] getOrganizationsWhitelistTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOrganizationsWhitelistTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsWhitelistTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsWhitelistInternalServerError creates a GetOrganizationsWhitelistInternalServerError with default headers values
func NewGetOrganizationsWhitelistInternalServerError() *GetOrganizationsWhitelistInternalServerError {
	return &GetOrganizationsWhitelistInternalServerError{}
}

/*GetOrganizationsWhitelistInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOrganizationsWhitelistInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetOrganizationsWhitelistInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/whitelist][%d] getOrganizationsWhitelistInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrganizationsWhitelistInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsWhitelistInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsWhitelistServiceUnavailable creates a GetOrganizationsWhitelistServiceUnavailable with default headers values
func NewGetOrganizationsWhitelistServiceUnavailable() *GetOrganizationsWhitelistServiceUnavailable {
	return &GetOrganizationsWhitelistServiceUnavailable{}
}

/*GetOrganizationsWhitelistServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOrganizationsWhitelistServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetOrganizationsWhitelistServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/whitelist][%d] getOrganizationsWhitelistServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOrganizationsWhitelistServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsWhitelistServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsWhitelistGatewayTimeout creates a GetOrganizationsWhitelistGatewayTimeout with default headers values
func NewGetOrganizationsWhitelistGatewayTimeout() *GetOrganizationsWhitelistGatewayTimeout {
	return &GetOrganizationsWhitelistGatewayTimeout{}
}

/*GetOrganizationsWhitelistGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetOrganizationsWhitelistGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOrganizationsWhitelistGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/whitelist][%d] getOrganizationsWhitelistGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOrganizationsWhitelistGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsWhitelistGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
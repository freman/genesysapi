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

/*
GetExternalcontactsOrganizationContactsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetExternalcontactsOrganizationContactsOK struct {
	Payload *models.ContactListing
}

// IsSuccess returns true when this get externalcontacts organization contacts o k response has a 2xx status code
func (o *GetExternalcontactsOrganizationContactsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get externalcontacts organization contacts o k response has a 3xx status code
func (o *GetExternalcontactsOrganizationContactsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts organization contacts o k response has a 4xx status code
func (o *GetExternalcontactsOrganizationContactsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get externalcontacts organization contacts o k response has a 5xx status code
func (o *GetExternalcontactsOrganizationContactsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts organization contacts o k response a status code equal to that given
func (o *GetExternalcontactsOrganizationContactsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetExternalcontactsOrganizationContactsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/contacts][%d] getExternalcontactsOrganizationContactsOK  %+v", 200, o.Payload)
}

func (o *GetExternalcontactsOrganizationContactsOK) String() string {
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

/*
GetExternalcontactsOrganizationContactsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetExternalcontactsOrganizationContactsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts organization contacts bad request response has a 2xx status code
func (o *GetExternalcontactsOrganizationContactsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts organization contacts bad request response has a 3xx status code
func (o *GetExternalcontactsOrganizationContactsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts organization contacts bad request response has a 4xx status code
func (o *GetExternalcontactsOrganizationContactsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts organization contacts bad request response has a 5xx status code
func (o *GetExternalcontactsOrganizationContactsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts organization contacts bad request response a status code equal to that given
func (o *GetExternalcontactsOrganizationContactsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetExternalcontactsOrganizationContactsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/contacts][%d] getExternalcontactsOrganizationContactsBadRequest  %+v", 400, o.Payload)
}

func (o *GetExternalcontactsOrganizationContactsBadRequest) String() string {
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

/*
GetExternalcontactsOrganizationContactsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetExternalcontactsOrganizationContactsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts organization contacts unauthorized response has a 2xx status code
func (o *GetExternalcontactsOrganizationContactsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts organization contacts unauthorized response has a 3xx status code
func (o *GetExternalcontactsOrganizationContactsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts organization contacts unauthorized response has a 4xx status code
func (o *GetExternalcontactsOrganizationContactsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts organization contacts unauthorized response has a 5xx status code
func (o *GetExternalcontactsOrganizationContactsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts organization contacts unauthorized response a status code equal to that given
func (o *GetExternalcontactsOrganizationContactsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetExternalcontactsOrganizationContactsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/contacts][%d] getExternalcontactsOrganizationContactsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetExternalcontactsOrganizationContactsUnauthorized) String() string {
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

/*
GetExternalcontactsOrganizationContactsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetExternalcontactsOrganizationContactsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts organization contacts forbidden response has a 2xx status code
func (o *GetExternalcontactsOrganizationContactsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts organization contacts forbidden response has a 3xx status code
func (o *GetExternalcontactsOrganizationContactsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts organization contacts forbidden response has a 4xx status code
func (o *GetExternalcontactsOrganizationContactsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts organization contacts forbidden response has a 5xx status code
func (o *GetExternalcontactsOrganizationContactsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts organization contacts forbidden response a status code equal to that given
func (o *GetExternalcontactsOrganizationContactsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetExternalcontactsOrganizationContactsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/contacts][%d] getExternalcontactsOrganizationContactsForbidden  %+v", 403, o.Payload)
}

func (o *GetExternalcontactsOrganizationContactsForbidden) String() string {
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

/*
GetExternalcontactsOrganizationContactsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetExternalcontactsOrganizationContactsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts organization contacts not found response has a 2xx status code
func (o *GetExternalcontactsOrganizationContactsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts organization contacts not found response has a 3xx status code
func (o *GetExternalcontactsOrganizationContactsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts organization contacts not found response has a 4xx status code
func (o *GetExternalcontactsOrganizationContactsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts organization contacts not found response has a 5xx status code
func (o *GetExternalcontactsOrganizationContactsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts organization contacts not found response a status code equal to that given
func (o *GetExternalcontactsOrganizationContactsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetExternalcontactsOrganizationContactsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/contacts][%d] getExternalcontactsOrganizationContactsNotFound  %+v", 404, o.Payload)
}

func (o *GetExternalcontactsOrganizationContactsNotFound) String() string {
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

/*
GetExternalcontactsOrganizationContactsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetExternalcontactsOrganizationContactsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts organization contacts request timeout response has a 2xx status code
func (o *GetExternalcontactsOrganizationContactsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts organization contacts request timeout response has a 3xx status code
func (o *GetExternalcontactsOrganizationContactsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts organization contacts request timeout response has a 4xx status code
func (o *GetExternalcontactsOrganizationContactsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts organization contacts request timeout response has a 5xx status code
func (o *GetExternalcontactsOrganizationContactsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts organization contacts request timeout response a status code equal to that given
func (o *GetExternalcontactsOrganizationContactsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetExternalcontactsOrganizationContactsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/contacts][%d] getExternalcontactsOrganizationContactsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetExternalcontactsOrganizationContactsRequestTimeout) String() string {
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

/*
GetExternalcontactsOrganizationContactsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetExternalcontactsOrganizationContactsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts organization contacts request entity too large response has a 2xx status code
func (o *GetExternalcontactsOrganizationContactsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts organization contacts request entity too large response has a 3xx status code
func (o *GetExternalcontactsOrganizationContactsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts organization contacts request entity too large response has a 4xx status code
func (o *GetExternalcontactsOrganizationContactsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts organization contacts request entity too large response has a 5xx status code
func (o *GetExternalcontactsOrganizationContactsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts organization contacts request entity too large response a status code equal to that given
func (o *GetExternalcontactsOrganizationContactsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetExternalcontactsOrganizationContactsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/contacts][%d] getExternalcontactsOrganizationContactsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetExternalcontactsOrganizationContactsRequestEntityTooLarge) String() string {
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

/*
GetExternalcontactsOrganizationContactsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetExternalcontactsOrganizationContactsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts organization contacts unsupported media type response has a 2xx status code
func (o *GetExternalcontactsOrganizationContactsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts organization contacts unsupported media type response has a 3xx status code
func (o *GetExternalcontactsOrganizationContactsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts organization contacts unsupported media type response has a 4xx status code
func (o *GetExternalcontactsOrganizationContactsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts organization contacts unsupported media type response has a 5xx status code
func (o *GetExternalcontactsOrganizationContactsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts organization contacts unsupported media type response a status code equal to that given
func (o *GetExternalcontactsOrganizationContactsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetExternalcontactsOrganizationContactsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/contacts][%d] getExternalcontactsOrganizationContactsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetExternalcontactsOrganizationContactsUnsupportedMediaType) String() string {
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

/*
GetExternalcontactsOrganizationContactsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetExternalcontactsOrganizationContactsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts organization contacts too many requests response has a 2xx status code
func (o *GetExternalcontactsOrganizationContactsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts organization contacts too many requests response has a 3xx status code
func (o *GetExternalcontactsOrganizationContactsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts organization contacts too many requests response has a 4xx status code
func (o *GetExternalcontactsOrganizationContactsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts organization contacts too many requests response has a 5xx status code
func (o *GetExternalcontactsOrganizationContactsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts organization contacts too many requests response a status code equal to that given
func (o *GetExternalcontactsOrganizationContactsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetExternalcontactsOrganizationContactsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/contacts][%d] getExternalcontactsOrganizationContactsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetExternalcontactsOrganizationContactsTooManyRequests) String() string {
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

/*
GetExternalcontactsOrganizationContactsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetExternalcontactsOrganizationContactsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts organization contacts internal server error response has a 2xx status code
func (o *GetExternalcontactsOrganizationContactsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts organization contacts internal server error response has a 3xx status code
func (o *GetExternalcontactsOrganizationContactsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts organization contacts internal server error response has a 4xx status code
func (o *GetExternalcontactsOrganizationContactsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get externalcontacts organization contacts internal server error response has a 5xx status code
func (o *GetExternalcontactsOrganizationContactsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get externalcontacts organization contacts internal server error response a status code equal to that given
func (o *GetExternalcontactsOrganizationContactsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetExternalcontactsOrganizationContactsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/contacts][%d] getExternalcontactsOrganizationContactsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetExternalcontactsOrganizationContactsInternalServerError) String() string {
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

/*
GetExternalcontactsOrganizationContactsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetExternalcontactsOrganizationContactsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts organization contacts service unavailable response has a 2xx status code
func (o *GetExternalcontactsOrganizationContactsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts organization contacts service unavailable response has a 3xx status code
func (o *GetExternalcontactsOrganizationContactsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts organization contacts service unavailable response has a 4xx status code
func (o *GetExternalcontactsOrganizationContactsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get externalcontacts organization contacts service unavailable response has a 5xx status code
func (o *GetExternalcontactsOrganizationContactsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get externalcontacts organization contacts service unavailable response a status code equal to that given
func (o *GetExternalcontactsOrganizationContactsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetExternalcontactsOrganizationContactsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/contacts][%d] getExternalcontactsOrganizationContactsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetExternalcontactsOrganizationContactsServiceUnavailable) String() string {
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

/*
GetExternalcontactsOrganizationContactsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetExternalcontactsOrganizationContactsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts organization contacts gateway timeout response has a 2xx status code
func (o *GetExternalcontactsOrganizationContactsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts organization contacts gateway timeout response has a 3xx status code
func (o *GetExternalcontactsOrganizationContactsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts organization contacts gateway timeout response has a 4xx status code
func (o *GetExternalcontactsOrganizationContactsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get externalcontacts organization contacts gateway timeout response has a 5xx status code
func (o *GetExternalcontactsOrganizationContactsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get externalcontacts organization contacts gateway timeout response a status code equal to that given
func (o *GetExternalcontactsOrganizationContactsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetExternalcontactsOrganizationContactsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/contacts][%d] getExternalcontactsOrganizationContactsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetExternalcontactsOrganizationContactsGatewayTimeout) String() string {
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

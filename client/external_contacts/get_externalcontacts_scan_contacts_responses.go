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
	case 408:
		result := NewGetExternalcontactsScanContactsRequestTimeout()
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
	case 422:
		result := NewGetExternalcontactsScanContactsUnprocessableEntity()
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

/*
GetExternalcontactsScanContactsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetExternalcontactsScanContactsOK struct {
	Payload *models.CursorContactListing
}

// IsSuccess returns true when this get externalcontacts scan contacts o k response has a 2xx status code
func (o *GetExternalcontactsScanContactsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get externalcontacts scan contacts o k response has a 3xx status code
func (o *GetExternalcontactsScanContactsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts scan contacts o k response has a 4xx status code
func (o *GetExternalcontactsScanContactsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get externalcontacts scan contacts o k response has a 5xx status code
func (o *GetExternalcontactsScanContactsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts scan contacts o k response a status code equal to that given
func (o *GetExternalcontactsScanContactsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetExternalcontactsScanContactsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/contacts][%d] getExternalcontactsScanContactsOK  %+v", 200, o.Payload)
}

func (o *GetExternalcontactsScanContactsOK) String() string {
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

/*
GetExternalcontactsScanContactsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetExternalcontactsScanContactsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts scan contacts bad request response has a 2xx status code
func (o *GetExternalcontactsScanContactsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts scan contacts bad request response has a 3xx status code
func (o *GetExternalcontactsScanContactsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts scan contacts bad request response has a 4xx status code
func (o *GetExternalcontactsScanContactsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts scan contacts bad request response has a 5xx status code
func (o *GetExternalcontactsScanContactsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts scan contacts bad request response a status code equal to that given
func (o *GetExternalcontactsScanContactsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetExternalcontactsScanContactsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/contacts][%d] getExternalcontactsScanContactsBadRequest  %+v", 400, o.Payload)
}

func (o *GetExternalcontactsScanContactsBadRequest) String() string {
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

/*
GetExternalcontactsScanContactsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetExternalcontactsScanContactsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts scan contacts unauthorized response has a 2xx status code
func (o *GetExternalcontactsScanContactsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts scan contacts unauthorized response has a 3xx status code
func (o *GetExternalcontactsScanContactsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts scan contacts unauthorized response has a 4xx status code
func (o *GetExternalcontactsScanContactsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts scan contacts unauthorized response has a 5xx status code
func (o *GetExternalcontactsScanContactsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts scan contacts unauthorized response a status code equal to that given
func (o *GetExternalcontactsScanContactsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetExternalcontactsScanContactsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/contacts][%d] getExternalcontactsScanContactsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetExternalcontactsScanContactsUnauthorized) String() string {
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

/*
GetExternalcontactsScanContactsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetExternalcontactsScanContactsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts scan contacts forbidden response has a 2xx status code
func (o *GetExternalcontactsScanContactsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts scan contacts forbidden response has a 3xx status code
func (o *GetExternalcontactsScanContactsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts scan contacts forbidden response has a 4xx status code
func (o *GetExternalcontactsScanContactsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts scan contacts forbidden response has a 5xx status code
func (o *GetExternalcontactsScanContactsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts scan contacts forbidden response a status code equal to that given
func (o *GetExternalcontactsScanContactsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetExternalcontactsScanContactsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/contacts][%d] getExternalcontactsScanContactsForbidden  %+v", 403, o.Payload)
}

func (o *GetExternalcontactsScanContactsForbidden) String() string {
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

/*
GetExternalcontactsScanContactsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetExternalcontactsScanContactsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts scan contacts not found response has a 2xx status code
func (o *GetExternalcontactsScanContactsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts scan contacts not found response has a 3xx status code
func (o *GetExternalcontactsScanContactsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts scan contacts not found response has a 4xx status code
func (o *GetExternalcontactsScanContactsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts scan contacts not found response has a 5xx status code
func (o *GetExternalcontactsScanContactsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts scan contacts not found response a status code equal to that given
func (o *GetExternalcontactsScanContactsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetExternalcontactsScanContactsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/contacts][%d] getExternalcontactsScanContactsNotFound  %+v", 404, o.Payload)
}

func (o *GetExternalcontactsScanContactsNotFound) String() string {
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

// NewGetExternalcontactsScanContactsRequestTimeout creates a GetExternalcontactsScanContactsRequestTimeout with default headers values
func NewGetExternalcontactsScanContactsRequestTimeout() *GetExternalcontactsScanContactsRequestTimeout {
	return &GetExternalcontactsScanContactsRequestTimeout{}
}

/*
GetExternalcontactsScanContactsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetExternalcontactsScanContactsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts scan contacts request timeout response has a 2xx status code
func (o *GetExternalcontactsScanContactsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts scan contacts request timeout response has a 3xx status code
func (o *GetExternalcontactsScanContactsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts scan contacts request timeout response has a 4xx status code
func (o *GetExternalcontactsScanContactsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts scan contacts request timeout response has a 5xx status code
func (o *GetExternalcontactsScanContactsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts scan contacts request timeout response a status code equal to that given
func (o *GetExternalcontactsScanContactsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetExternalcontactsScanContactsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/contacts][%d] getExternalcontactsScanContactsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetExternalcontactsScanContactsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/contacts][%d] getExternalcontactsScanContactsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetExternalcontactsScanContactsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsScanContactsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*
GetExternalcontactsScanContactsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetExternalcontactsScanContactsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts scan contacts request entity too large response has a 2xx status code
func (o *GetExternalcontactsScanContactsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts scan contacts request entity too large response has a 3xx status code
func (o *GetExternalcontactsScanContactsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts scan contacts request entity too large response has a 4xx status code
func (o *GetExternalcontactsScanContactsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts scan contacts request entity too large response has a 5xx status code
func (o *GetExternalcontactsScanContactsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts scan contacts request entity too large response a status code equal to that given
func (o *GetExternalcontactsScanContactsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetExternalcontactsScanContactsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/contacts][%d] getExternalcontactsScanContactsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetExternalcontactsScanContactsRequestEntityTooLarge) String() string {
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

/*
GetExternalcontactsScanContactsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetExternalcontactsScanContactsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts scan contacts unsupported media type response has a 2xx status code
func (o *GetExternalcontactsScanContactsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts scan contacts unsupported media type response has a 3xx status code
func (o *GetExternalcontactsScanContactsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts scan contacts unsupported media type response has a 4xx status code
func (o *GetExternalcontactsScanContactsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts scan contacts unsupported media type response has a 5xx status code
func (o *GetExternalcontactsScanContactsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts scan contacts unsupported media type response a status code equal to that given
func (o *GetExternalcontactsScanContactsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetExternalcontactsScanContactsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/contacts][%d] getExternalcontactsScanContactsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetExternalcontactsScanContactsUnsupportedMediaType) String() string {
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

// NewGetExternalcontactsScanContactsUnprocessableEntity creates a GetExternalcontactsScanContactsUnprocessableEntity with default headers values
func NewGetExternalcontactsScanContactsUnprocessableEntity() *GetExternalcontactsScanContactsUnprocessableEntity {
	return &GetExternalcontactsScanContactsUnprocessableEntity{}
}

/*
GetExternalcontactsScanContactsUnprocessableEntity describes a response with status code 422, with default header values.

GetExternalcontactsScanContactsUnprocessableEntity get externalcontacts scan contacts unprocessable entity
*/
type GetExternalcontactsScanContactsUnprocessableEntity struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts scan contacts unprocessable entity response has a 2xx status code
func (o *GetExternalcontactsScanContactsUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts scan contacts unprocessable entity response has a 3xx status code
func (o *GetExternalcontactsScanContactsUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts scan contacts unprocessable entity response has a 4xx status code
func (o *GetExternalcontactsScanContactsUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts scan contacts unprocessable entity response has a 5xx status code
func (o *GetExternalcontactsScanContactsUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts scan contacts unprocessable entity response a status code equal to that given
func (o *GetExternalcontactsScanContactsUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *GetExternalcontactsScanContactsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/contacts][%d] getExternalcontactsScanContactsUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetExternalcontactsScanContactsUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/contacts][%d] getExternalcontactsScanContactsUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetExternalcontactsScanContactsUnprocessableEntity) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsScanContactsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*
GetExternalcontactsScanContactsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetExternalcontactsScanContactsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts scan contacts too many requests response has a 2xx status code
func (o *GetExternalcontactsScanContactsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts scan contacts too many requests response has a 3xx status code
func (o *GetExternalcontactsScanContactsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts scan contacts too many requests response has a 4xx status code
func (o *GetExternalcontactsScanContactsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts scan contacts too many requests response has a 5xx status code
func (o *GetExternalcontactsScanContactsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts scan contacts too many requests response a status code equal to that given
func (o *GetExternalcontactsScanContactsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetExternalcontactsScanContactsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/contacts][%d] getExternalcontactsScanContactsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetExternalcontactsScanContactsTooManyRequests) String() string {
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

/*
GetExternalcontactsScanContactsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetExternalcontactsScanContactsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts scan contacts internal server error response has a 2xx status code
func (o *GetExternalcontactsScanContactsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts scan contacts internal server error response has a 3xx status code
func (o *GetExternalcontactsScanContactsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts scan contacts internal server error response has a 4xx status code
func (o *GetExternalcontactsScanContactsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get externalcontacts scan contacts internal server error response has a 5xx status code
func (o *GetExternalcontactsScanContactsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get externalcontacts scan contacts internal server error response a status code equal to that given
func (o *GetExternalcontactsScanContactsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetExternalcontactsScanContactsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/contacts][%d] getExternalcontactsScanContactsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetExternalcontactsScanContactsInternalServerError) String() string {
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

/*
GetExternalcontactsScanContactsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetExternalcontactsScanContactsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts scan contacts service unavailable response has a 2xx status code
func (o *GetExternalcontactsScanContactsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts scan contacts service unavailable response has a 3xx status code
func (o *GetExternalcontactsScanContactsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts scan contacts service unavailable response has a 4xx status code
func (o *GetExternalcontactsScanContactsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get externalcontacts scan contacts service unavailable response has a 5xx status code
func (o *GetExternalcontactsScanContactsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get externalcontacts scan contacts service unavailable response a status code equal to that given
func (o *GetExternalcontactsScanContactsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetExternalcontactsScanContactsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/contacts][%d] getExternalcontactsScanContactsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetExternalcontactsScanContactsServiceUnavailable) String() string {
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

/*
GetExternalcontactsScanContactsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetExternalcontactsScanContactsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts scan contacts gateway timeout response has a 2xx status code
func (o *GetExternalcontactsScanContactsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts scan contacts gateway timeout response has a 3xx status code
func (o *GetExternalcontactsScanContactsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts scan contacts gateway timeout response has a 4xx status code
func (o *GetExternalcontactsScanContactsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get externalcontacts scan contacts gateway timeout response has a 5xx status code
func (o *GetExternalcontactsScanContactsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get externalcontacts scan contacts gateway timeout response a status code equal to that given
func (o *GetExternalcontactsScanContactsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetExternalcontactsScanContactsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/contacts][%d] getExternalcontactsScanContactsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetExternalcontactsScanContactsGatewayTimeout) String() string {
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

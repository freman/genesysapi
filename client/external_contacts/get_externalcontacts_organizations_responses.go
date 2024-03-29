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

// GetExternalcontactsOrganizationsReader is a Reader for the GetExternalcontactsOrganizations structure.
type GetExternalcontactsOrganizationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExternalcontactsOrganizationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExternalcontactsOrganizationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetExternalcontactsOrganizationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetExternalcontactsOrganizationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetExternalcontactsOrganizationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetExternalcontactsOrganizationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetExternalcontactsOrganizationsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetExternalcontactsOrganizationsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetExternalcontactsOrganizationsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetExternalcontactsOrganizationsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetExternalcontactsOrganizationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetExternalcontactsOrganizationsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetExternalcontactsOrganizationsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetExternalcontactsOrganizationsOK creates a GetExternalcontactsOrganizationsOK with default headers values
func NewGetExternalcontactsOrganizationsOK() *GetExternalcontactsOrganizationsOK {
	return &GetExternalcontactsOrganizationsOK{}
}

/*
GetExternalcontactsOrganizationsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetExternalcontactsOrganizationsOK struct {
	Payload *models.ExternalOrganizationListing
}

// IsSuccess returns true when this get externalcontacts organizations o k response has a 2xx status code
func (o *GetExternalcontactsOrganizationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get externalcontacts organizations o k response has a 3xx status code
func (o *GetExternalcontactsOrganizationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts organizations o k response has a 4xx status code
func (o *GetExternalcontactsOrganizationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get externalcontacts organizations o k response has a 5xx status code
func (o *GetExternalcontactsOrganizationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts organizations o k response a status code equal to that given
func (o *GetExternalcontactsOrganizationsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetExternalcontactsOrganizationsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations][%d] getExternalcontactsOrganizationsOK  %+v", 200, o.Payload)
}

func (o *GetExternalcontactsOrganizationsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations][%d] getExternalcontactsOrganizationsOK  %+v", 200, o.Payload)
}

func (o *GetExternalcontactsOrganizationsOK) GetPayload() *models.ExternalOrganizationListing {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExternalOrganizationListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationsBadRequest creates a GetExternalcontactsOrganizationsBadRequest with default headers values
func NewGetExternalcontactsOrganizationsBadRequest() *GetExternalcontactsOrganizationsBadRequest {
	return &GetExternalcontactsOrganizationsBadRequest{}
}

/*
GetExternalcontactsOrganizationsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetExternalcontactsOrganizationsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts organizations bad request response has a 2xx status code
func (o *GetExternalcontactsOrganizationsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts organizations bad request response has a 3xx status code
func (o *GetExternalcontactsOrganizationsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts organizations bad request response has a 4xx status code
func (o *GetExternalcontactsOrganizationsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts organizations bad request response has a 5xx status code
func (o *GetExternalcontactsOrganizationsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts organizations bad request response a status code equal to that given
func (o *GetExternalcontactsOrganizationsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetExternalcontactsOrganizationsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations][%d] getExternalcontactsOrganizationsBadRequest  %+v", 400, o.Payload)
}

func (o *GetExternalcontactsOrganizationsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations][%d] getExternalcontactsOrganizationsBadRequest  %+v", 400, o.Payload)
}

func (o *GetExternalcontactsOrganizationsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationsUnauthorized creates a GetExternalcontactsOrganizationsUnauthorized with default headers values
func NewGetExternalcontactsOrganizationsUnauthorized() *GetExternalcontactsOrganizationsUnauthorized {
	return &GetExternalcontactsOrganizationsUnauthorized{}
}

/*
GetExternalcontactsOrganizationsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetExternalcontactsOrganizationsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts organizations unauthorized response has a 2xx status code
func (o *GetExternalcontactsOrganizationsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts organizations unauthorized response has a 3xx status code
func (o *GetExternalcontactsOrganizationsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts organizations unauthorized response has a 4xx status code
func (o *GetExternalcontactsOrganizationsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts organizations unauthorized response has a 5xx status code
func (o *GetExternalcontactsOrganizationsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts organizations unauthorized response a status code equal to that given
func (o *GetExternalcontactsOrganizationsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetExternalcontactsOrganizationsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations][%d] getExternalcontactsOrganizationsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetExternalcontactsOrganizationsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations][%d] getExternalcontactsOrganizationsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetExternalcontactsOrganizationsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationsForbidden creates a GetExternalcontactsOrganizationsForbidden with default headers values
func NewGetExternalcontactsOrganizationsForbidden() *GetExternalcontactsOrganizationsForbidden {
	return &GetExternalcontactsOrganizationsForbidden{}
}

/*
GetExternalcontactsOrganizationsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetExternalcontactsOrganizationsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts organizations forbidden response has a 2xx status code
func (o *GetExternalcontactsOrganizationsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts organizations forbidden response has a 3xx status code
func (o *GetExternalcontactsOrganizationsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts organizations forbidden response has a 4xx status code
func (o *GetExternalcontactsOrganizationsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts organizations forbidden response has a 5xx status code
func (o *GetExternalcontactsOrganizationsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts organizations forbidden response a status code equal to that given
func (o *GetExternalcontactsOrganizationsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetExternalcontactsOrganizationsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations][%d] getExternalcontactsOrganizationsForbidden  %+v", 403, o.Payload)
}

func (o *GetExternalcontactsOrganizationsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations][%d] getExternalcontactsOrganizationsForbidden  %+v", 403, o.Payload)
}

func (o *GetExternalcontactsOrganizationsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationsNotFound creates a GetExternalcontactsOrganizationsNotFound with default headers values
func NewGetExternalcontactsOrganizationsNotFound() *GetExternalcontactsOrganizationsNotFound {
	return &GetExternalcontactsOrganizationsNotFound{}
}

/*
GetExternalcontactsOrganizationsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetExternalcontactsOrganizationsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts organizations not found response has a 2xx status code
func (o *GetExternalcontactsOrganizationsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts organizations not found response has a 3xx status code
func (o *GetExternalcontactsOrganizationsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts organizations not found response has a 4xx status code
func (o *GetExternalcontactsOrganizationsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts organizations not found response has a 5xx status code
func (o *GetExternalcontactsOrganizationsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts organizations not found response a status code equal to that given
func (o *GetExternalcontactsOrganizationsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetExternalcontactsOrganizationsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations][%d] getExternalcontactsOrganizationsNotFound  %+v", 404, o.Payload)
}

func (o *GetExternalcontactsOrganizationsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations][%d] getExternalcontactsOrganizationsNotFound  %+v", 404, o.Payload)
}

func (o *GetExternalcontactsOrganizationsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationsRequestTimeout creates a GetExternalcontactsOrganizationsRequestTimeout with default headers values
func NewGetExternalcontactsOrganizationsRequestTimeout() *GetExternalcontactsOrganizationsRequestTimeout {
	return &GetExternalcontactsOrganizationsRequestTimeout{}
}

/*
GetExternalcontactsOrganizationsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetExternalcontactsOrganizationsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts organizations request timeout response has a 2xx status code
func (o *GetExternalcontactsOrganizationsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts organizations request timeout response has a 3xx status code
func (o *GetExternalcontactsOrganizationsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts organizations request timeout response has a 4xx status code
func (o *GetExternalcontactsOrganizationsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts organizations request timeout response has a 5xx status code
func (o *GetExternalcontactsOrganizationsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts organizations request timeout response a status code equal to that given
func (o *GetExternalcontactsOrganizationsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetExternalcontactsOrganizationsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations][%d] getExternalcontactsOrganizationsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetExternalcontactsOrganizationsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations][%d] getExternalcontactsOrganizationsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetExternalcontactsOrganizationsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationsRequestEntityTooLarge creates a GetExternalcontactsOrganizationsRequestEntityTooLarge with default headers values
func NewGetExternalcontactsOrganizationsRequestEntityTooLarge() *GetExternalcontactsOrganizationsRequestEntityTooLarge {
	return &GetExternalcontactsOrganizationsRequestEntityTooLarge{}
}

/*
GetExternalcontactsOrganizationsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetExternalcontactsOrganizationsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts organizations request entity too large response has a 2xx status code
func (o *GetExternalcontactsOrganizationsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts organizations request entity too large response has a 3xx status code
func (o *GetExternalcontactsOrganizationsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts organizations request entity too large response has a 4xx status code
func (o *GetExternalcontactsOrganizationsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts organizations request entity too large response has a 5xx status code
func (o *GetExternalcontactsOrganizationsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts organizations request entity too large response a status code equal to that given
func (o *GetExternalcontactsOrganizationsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetExternalcontactsOrganizationsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations][%d] getExternalcontactsOrganizationsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetExternalcontactsOrganizationsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations][%d] getExternalcontactsOrganizationsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetExternalcontactsOrganizationsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationsUnsupportedMediaType creates a GetExternalcontactsOrganizationsUnsupportedMediaType with default headers values
func NewGetExternalcontactsOrganizationsUnsupportedMediaType() *GetExternalcontactsOrganizationsUnsupportedMediaType {
	return &GetExternalcontactsOrganizationsUnsupportedMediaType{}
}

/*
GetExternalcontactsOrganizationsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetExternalcontactsOrganizationsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts organizations unsupported media type response has a 2xx status code
func (o *GetExternalcontactsOrganizationsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts organizations unsupported media type response has a 3xx status code
func (o *GetExternalcontactsOrganizationsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts organizations unsupported media type response has a 4xx status code
func (o *GetExternalcontactsOrganizationsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts organizations unsupported media type response has a 5xx status code
func (o *GetExternalcontactsOrganizationsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts organizations unsupported media type response a status code equal to that given
func (o *GetExternalcontactsOrganizationsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetExternalcontactsOrganizationsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations][%d] getExternalcontactsOrganizationsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetExternalcontactsOrganizationsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations][%d] getExternalcontactsOrganizationsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetExternalcontactsOrganizationsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationsTooManyRequests creates a GetExternalcontactsOrganizationsTooManyRequests with default headers values
func NewGetExternalcontactsOrganizationsTooManyRequests() *GetExternalcontactsOrganizationsTooManyRequests {
	return &GetExternalcontactsOrganizationsTooManyRequests{}
}

/*
GetExternalcontactsOrganizationsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetExternalcontactsOrganizationsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts organizations too many requests response has a 2xx status code
func (o *GetExternalcontactsOrganizationsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts organizations too many requests response has a 3xx status code
func (o *GetExternalcontactsOrganizationsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts organizations too many requests response has a 4xx status code
func (o *GetExternalcontactsOrganizationsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts organizations too many requests response has a 5xx status code
func (o *GetExternalcontactsOrganizationsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts organizations too many requests response a status code equal to that given
func (o *GetExternalcontactsOrganizationsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetExternalcontactsOrganizationsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations][%d] getExternalcontactsOrganizationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetExternalcontactsOrganizationsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations][%d] getExternalcontactsOrganizationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetExternalcontactsOrganizationsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationsInternalServerError creates a GetExternalcontactsOrganizationsInternalServerError with default headers values
func NewGetExternalcontactsOrganizationsInternalServerError() *GetExternalcontactsOrganizationsInternalServerError {
	return &GetExternalcontactsOrganizationsInternalServerError{}
}

/*
GetExternalcontactsOrganizationsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetExternalcontactsOrganizationsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts organizations internal server error response has a 2xx status code
func (o *GetExternalcontactsOrganizationsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts organizations internal server error response has a 3xx status code
func (o *GetExternalcontactsOrganizationsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts organizations internal server error response has a 4xx status code
func (o *GetExternalcontactsOrganizationsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get externalcontacts organizations internal server error response has a 5xx status code
func (o *GetExternalcontactsOrganizationsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get externalcontacts organizations internal server error response a status code equal to that given
func (o *GetExternalcontactsOrganizationsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetExternalcontactsOrganizationsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations][%d] getExternalcontactsOrganizationsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetExternalcontactsOrganizationsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations][%d] getExternalcontactsOrganizationsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetExternalcontactsOrganizationsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationsServiceUnavailable creates a GetExternalcontactsOrganizationsServiceUnavailable with default headers values
func NewGetExternalcontactsOrganizationsServiceUnavailable() *GetExternalcontactsOrganizationsServiceUnavailable {
	return &GetExternalcontactsOrganizationsServiceUnavailable{}
}

/*
GetExternalcontactsOrganizationsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetExternalcontactsOrganizationsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts organizations service unavailable response has a 2xx status code
func (o *GetExternalcontactsOrganizationsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts organizations service unavailable response has a 3xx status code
func (o *GetExternalcontactsOrganizationsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts organizations service unavailable response has a 4xx status code
func (o *GetExternalcontactsOrganizationsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get externalcontacts organizations service unavailable response has a 5xx status code
func (o *GetExternalcontactsOrganizationsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get externalcontacts organizations service unavailable response a status code equal to that given
func (o *GetExternalcontactsOrganizationsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetExternalcontactsOrganizationsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations][%d] getExternalcontactsOrganizationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetExternalcontactsOrganizationsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations][%d] getExternalcontactsOrganizationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetExternalcontactsOrganizationsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationsGatewayTimeout creates a GetExternalcontactsOrganizationsGatewayTimeout with default headers values
func NewGetExternalcontactsOrganizationsGatewayTimeout() *GetExternalcontactsOrganizationsGatewayTimeout {
	return &GetExternalcontactsOrganizationsGatewayTimeout{}
}

/*
GetExternalcontactsOrganizationsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetExternalcontactsOrganizationsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts organizations gateway timeout response has a 2xx status code
func (o *GetExternalcontactsOrganizationsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts organizations gateway timeout response has a 3xx status code
func (o *GetExternalcontactsOrganizationsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts organizations gateway timeout response has a 4xx status code
func (o *GetExternalcontactsOrganizationsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get externalcontacts organizations gateway timeout response has a 5xx status code
func (o *GetExternalcontactsOrganizationsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get externalcontacts organizations gateway timeout response a status code equal to that given
func (o *GetExternalcontactsOrganizationsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetExternalcontactsOrganizationsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations][%d] getExternalcontactsOrganizationsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetExternalcontactsOrganizationsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations][%d] getExternalcontactsOrganizationsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetExternalcontactsOrganizationsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

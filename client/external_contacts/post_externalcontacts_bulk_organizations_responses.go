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

// PostExternalcontactsBulkOrganizationsReader is a Reader for the PostExternalcontactsBulkOrganizations structure.
type PostExternalcontactsBulkOrganizationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostExternalcontactsBulkOrganizationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostExternalcontactsBulkOrganizationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostExternalcontactsBulkOrganizationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostExternalcontactsBulkOrganizationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostExternalcontactsBulkOrganizationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostExternalcontactsBulkOrganizationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostExternalcontactsBulkOrganizationsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostExternalcontactsBulkOrganizationsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostExternalcontactsBulkOrganizationsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostExternalcontactsBulkOrganizationsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostExternalcontactsBulkOrganizationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostExternalcontactsBulkOrganizationsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostExternalcontactsBulkOrganizationsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostExternalcontactsBulkOrganizationsOK creates a PostExternalcontactsBulkOrganizationsOK with default headers values
func NewPostExternalcontactsBulkOrganizationsOK() *PostExternalcontactsBulkOrganizationsOK {
	return &PostExternalcontactsBulkOrganizationsOK{}
}

/*
PostExternalcontactsBulkOrganizationsOK describes a response with status code 200, with default header values.

successful operation
*/
type PostExternalcontactsBulkOrganizationsOK struct {
	Payload *models.BulkFetchOrganizationsResponse
}

// IsSuccess returns true when this post externalcontacts bulk organizations o k response has a 2xx status code
func (o *PostExternalcontactsBulkOrganizationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post externalcontacts bulk organizations o k response has a 3xx status code
func (o *PostExternalcontactsBulkOrganizationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk organizations o k response has a 4xx status code
func (o *PostExternalcontactsBulkOrganizationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post externalcontacts bulk organizations o k response has a 5xx status code
func (o *PostExternalcontactsBulkOrganizationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk organizations o k response a status code equal to that given
func (o *PostExternalcontactsBulkOrganizationsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostExternalcontactsBulkOrganizationsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations][%d] postExternalcontactsBulkOrganizationsOK  %+v", 200, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsOK) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations][%d] postExternalcontactsBulkOrganizationsOK  %+v", 200, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsOK) GetPayload() *models.BulkFetchOrganizationsResponse {
	return o.Payload
}

func (o *PostExternalcontactsBulkOrganizationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BulkFetchOrganizationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkOrganizationsBadRequest creates a PostExternalcontactsBulkOrganizationsBadRequest with default headers values
func NewPostExternalcontactsBulkOrganizationsBadRequest() *PostExternalcontactsBulkOrganizationsBadRequest {
	return &PostExternalcontactsBulkOrganizationsBadRequest{}
}

/*
PostExternalcontactsBulkOrganizationsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostExternalcontactsBulkOrganizationsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk organizations bad request response has a 2xx status code
func (o *PostExternalcontactsBulkOrganizationsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk organizations bad request response has a 3xx status code
func (o *PostExternalcontactsBulkOrganizationsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk organizations bad request response has a 4xx status code
func (o *PostExternalcontactsBulkOrganizationsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk organizations bad request response has a 5xx status code
func (o *PostExternalcontactsBulkOrganizationsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk organizations bad request response a status code equal to that given
func (o *PostExternalcontactsBulkOrganizationsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostExternalcontactsBulkOrganizationsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations][%d] postExternalcontactsBulkOrganizationsBadRequest  %+v", 400, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations][%d] postExternalcontactsBulkOrganizationsBadRequest  %+v", 400, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkOrganizationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkOrganizationsUnauthorized creates a PostExternalcontactsBulkOrganizationsUnauthorized with default headers values
func NewPostExternalcontactsBulkOrganizationsUnauthorized() *PostExternalcontactsBulkOrganizationsUnauthorized {
	return &PostExternalcontactsBulkOrganizationsUnauthorized{}
}

/*
PostExternalcontactsBulkOrganizationsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostExternalcontactsBulkOrganizationsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk organizations unauthorized response has a 2xx status code
func (o *PostExternalcontactsBulkOrganizationsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk organizations unauthorized response has a 3xx status code
func (o *PostExternalcontactsBulkOrganizationsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk organizations unauthorized response has a 4xx status code
func (o *PostExternalcontactsBulkOrganizationsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk organizations unauthorized response has a 5xx status code
func (o *PostExternalcontactsBulkOrganizationsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk organizations unauthorized response a status code equal to that given
func (o *PostExternalcontactsBulkOrganizationsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostExternalcontactsBulkOrganizationsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations][%d] postExternalcontactsBulkOrganizationsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations][%d] postExternalcontactsBulkOrganizationsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkOrganizationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkOrganizationsForbidden creates a PostExternalcontactsBulkOrganizationsForbidden with default headers values
func NewPostExternalcontactsBulkOrganizationsForbidden() *PostExternalcontactsBulkOrganizationsForbidden {
	return &PostExternalcontactsBulkOrganizationsForbidden{}
}

/*
PostExternalcontactsBulkOrganizationsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostExternalcontactsBulkOrganizationsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk organizations forbidden response has a 2xx status code
func (o *PostExternalcontactsBulkOrganizationsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk organizations forbidden response has a 3xx status code
func (o *PostExternalcontactsBulkOrganizationsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk organizations forbidden response has a 4xx status code
func (o *PostExternalcontactsBulkOrganizationsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk organizations forbidden response has a 5xx status code
func (o *PostExternalcontactsBulkOrganizationsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk organizations forbidden response a status code equal to that given
func (o *PostExternalcontactsBulkOrganizationsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostExternalcontactsBulkOrganizationsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations][%d] postExternalcontactsBulkOrganizationsForbidden  %+v", 403, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations][%d] postExternalcontactsBulkOrganizationsForbidden  %+v", 403, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkOrganizationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkOrganizationsNotFound creates a PostExternalcontactsBulkOrganizationsNotFound with default headers values
func NewPostExternalcontactsBulkOrganizationsNotFound() *PostExternalcontactsBulkOrganizationsNotFound {
	return &PostExternalcontactsBulkOrganizationsNotFound{}
}

/*
PostExternalcontactsBulkOrganizationsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostExternalcontactsBulkOrganizationsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk organizations not found response has a 2xx status code
func (o *PostExternalcontactsBulkOrganizationsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk organizations not found response has a 3xx status code
func (o *PostExternalcontactsBulkOrganizationsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk organizations not found response has a 4xx status code
func (o *PostExternalcontactsBulkOrganizationsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk organizations not found response has a 5xx status code
func (o *PostExternalcontactsBulkOrganizationsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk organizations not found response a status code equal to that given
func (o *PostExternalcontactsBulkOrganizationsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostExternalcontactsBulkOrganizationsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations][%d] postExternalcontactsBulkOrganizationsNotFound  %+v", 404, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations][%d] postExternalcontactsBulkOrganizationsNotFound  %+v", 404, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkOrganizationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkOrganizationsRequestTimeout creates a PostExternalcontactsBulkOrganizationsRequestTimeout with default headers values
func NewPostExternalcontactsBulkOrganizationsRequestTimeout() *PostExternalcontactsBulkOrganizationsRequestTimeout {
	return &PostExternalcontactsBulkOrganizationsRequestTimeout{}
}

/*
PostExternalcontactsBulkOrganizationsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostExternalcontactsBulkOrganizationsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk organizations request timeout response has a 2xx status code
func (o *PostExternalcontactsBulkOrganizationsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk organizations request timeout response has a 3xx status code
func (o *PostExternalcontactsBulkOrganizationsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk organizations request timeout response has a 4xx status code
func (o *PostExternalcontactsBulkOrganizationsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk organizations request timeout response has a 5xx status code
func (o *PostExternalcontactsBulkOrganizationsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk organizations request timeout response a status code equal to that given
func (o *PostExternalcontactsBulkOrganizationsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostExternalcontactsBulkOrganizationsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations][%d] postExternalcontactsBulkOrganizationsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations][%d] postExternalcontactsBulkOrganizationsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkOrganizationsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkOrganizationsRequestEntityTooLarge creates a PostExternalcontactsBulkOrganizationsRequestEntityTooLarge with default headers values
func NewPostExternalcontactsBulkOrganizationsRequestEntityTooLarge() *PostExternalcontactsBulkOrganizationsRequestEntityTooLarge {
	return &PostExternalcontactsBulkOrganizationsRequestEntityTooLarge{}
}

/*
PostExternalcontactsBulkOrganizationsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostExternalcontactsBulkOrganizationsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk organizations request entity too large response has a 2xx status code
func (o *PostExternalcontactsBulkOrganizationsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk organizations request entity too large response has a 3xx status code
func (o *PostExternalcontactsBulkOrganizationsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk organizations request entity too large response has a 4xx status code
func (o *PostExternalcontactsBulkOrganizationsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk organizations request entity too large response has a 5xx status code
func (o *PostExternalcontactsBulkOrganizationsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk organizations request entity too large response a status code equal to that given
func (o *PostExternalcontactsBulkOrganizationsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostExternalcontactsBulkOrganizationsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations][%d] postExternalcontactsBulkOrganizationsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations][%d] postExternalcontactsBulkOrganizationsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkOrganizationsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkOrganizationsUnsupportedMediaType creates a PostExternalcontactsBulkOrganizationsUnsupportedMediaType with default headers values
func NewPostExternalcontactsBulkOrganizationsUnsupportedMediaType() *PostExternalcontactsBulkOrganizationsUnsupportedMediaType {
	return &PostExternalcontactsBulkOrganizationsUnsupportedMediaType{}
}

/*
PostExternalcontactsBulkOrganizationsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostExternalcontactsBulkOrganizationsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk organizations unsupported media type response has a 2xx status code
func (o *PostExternalcontactsBulkOrganizationsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk organizations unsupported media type response has a 3xx status code
func (o *PostExternalcontactsBulkOrganizationsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk organizations unsupported media type response has a 4xx status code
func (o *PostExternalcontactsBulkOrganizationsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk organizations unsupported media type response has a 5xx status code
func (o *PostExternalcontactsBulkOrganizationsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk organizations unsupported media type response a status code equal to that given
func (o *PostExternalcontactsBulkOrganizationsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostExternalcontactsBulkOrganizationsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations][%d] postExternalcontactsBulkOrganizationsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations][%d] postExternalcontactsBulkOrganizationsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkOrganizationsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkOrganizationsTooManyRequests creates a PostExternalcontactsBulkOrganizationsTooManyRequests with default headers values
func NewPostExternalcontactsBulkOrganizationsTooManyRequests() *PostExternalcontactsBulkOrganizationsTooManyRequests {
	return &PostExternalcontactsBulkOrganizationsTooManyRequests{}
}

/*
PostExternalcontactsBulkOrganizationsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostExternalcontactsBulkOrganizationsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk organizations too many requests response has a 2xx status code
func (o *PostExternalcontactsBulkOrganizationsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk organizations too many requests response has a 3xx status code
func (o *PostExternalcontactsBulkOrganizationsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk organizations too many requests response has a 4xx status code
func (o *PostExternalcontactsBulkOrganizationsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk organizations too many requests response has a 5xx status code
func (o *PostExternalcontactsBulkOrganizationsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk organizations too many requests response a status code equal to that given
func (o *PostExternalcontactsBulkOrganizationsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostExternalcontactsBulkOrganizationsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations][%d] postExternalcontactsBulkOrganizationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations][%d] postExternalcontactsBulkOrganizationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkOrganizationsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkOrganizationsInternalServerError creates a PostExternalcontactsBulkOrganizationsInternalServerError with default headers values
func NewPostExternalcontactsBulkOrganizationsInternalServerError() *PostExternalcontactsBulkOrganizationsInternalServerError {
	return &PostExternalcontactsBulkOrganizationsInternalServerError{}
}

/*
PostExternalcontactsBulkOrganizationsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostExternalcontactsBulkOrganizationsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk organizations internal server error response has a 2xx status code
func (o *PostExternalcontactsBulkOrganizationsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk organizations internal server error response has a 3xx status code
func (o *PostExternalcontactsBulkOrganizationsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk organizations internal server error response has a 4xx status code
func (o *PostExternalcontactsBulkOrganizationsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post externalcontacts bulk organizations internal server error response has a 5xx status code
func (o *PostExternalcontactsBulkOrganizationsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post externalcontacts bulk organizations internal server error response a status code equal to that given
func (o *PostExternalcontactsBulkOrganizationsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostExternalcontactsBulkOrganizationsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations][%d] postExternalcontactsBulkOrganizationsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations][%d] postExternalcontactsBulkOrganizationsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkOrganizationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkOrganizationsServiceUnavailable creates a PostExternalcontactsBulkOrganizationsServiceUnavailable with default headers values
func NewPostExternalcontactsBulkOrganizationsServiceUnavailable() *PostExternalcontactsBulkOrganizationsServiceUnavailable {
	return &PostExternalcontactsBulkOrganizationsServiceUnavailable{}
}

/*
PostExternalcontactsBulkOrganizationsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostExternalcontactsBulkOrganizationsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk organizations service unavailable response has a 2xx status code
func (o *PostExternalcontactsBulkOrganizationsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk organizations service unavailable response has a 3xx status code
func (o *PostExternalcontactsBulkOrganizationsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk organizations service unavailable response has a 4xx status code
func (o *PostExternalcontactsBulkOrganizationsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post externalcontacts bulk organizations service unavailable response has a 5xx status code
func (o *PostExternalcontactsBulkOrganizationsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post externalcontacts bulk organizations service unavailable response a status code equal to that given
func (o *PostExternalcontactsBulkOrganizationsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostExternalcontactsBulkOrganizationsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations][%d] postExternalcontactsBulkOrganizationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations][%d] postExternalcontactsBulkOrganizationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkOrganizationsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkOrganizationsGatewayTimeout creates a PostExternalcontactsBulkOrganizationsGatewayTimeout with default headers values
func NewPostExternalcontactsBulkOrganizationsGatewayTimeout() *PostExternalcontactsBulkOrganizationsGatewayTimeout {
	return &PostExternalcontactsBulkOrganizationsGatewayTimeout{}
}

/*
PostExternalcontactsBulkOrganizationsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostExternalcontactsBulkOrganizationsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk organizations gateway timeout response has a 2xx status code
func (o *PostExternalcontactsBulkOrganizationsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk organizations gateway timeout response has a 3xx status code
func (o *PostExternalcontactsBulkOrganizationsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk organizations gateway timeout response has a 4xx status code
func (o *PostExternalcontactsBulkOrganizationsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post externalcontacts bulk organizations gateway timeout response has a 5xx status code
func (o *PostExternalcontactsBulkOrganizationsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post externalcontacts bulk organizations gateway timeout response a status code equal to that given
func (o *PostExternalcontactsBulkOrganizationsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostExternalcontactsBulkOrganizationsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations][%d] postExternalcontactsBulkOrganizationsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations][%d] postExternalcontactsBulkOrganizationsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkOrganizationsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

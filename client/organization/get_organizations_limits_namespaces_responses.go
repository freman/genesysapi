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

// GetOrganizationsLimitsNamespacesReader is a Reader for the GetOrganizationsLimitsNamespaces structure.
type GetOrganizationsLimitsNamespacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationsLimitsNamespacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationsLimitsNamespacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOrganizationsLimitsNamespacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOrganizationsLimitsNamespacesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrganizationsLimitsNamespacesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrganizationsLimitsNamespacesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOrganizationsLimitsNamespacesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOrganizationsLimitsNamespacesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOrganizationsLimitsNamespacesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOrganizationsLimitsNamespacesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOrganizationsLimitsNamespacesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOrganizationsLimitsNamespacesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOrganizationsLimitsNamespacesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrganizationsLimitsNamespacesOK creates a GetOrganizationsLimitsNamespacesOK with default headers values
func NewGetOrganizationsLimitsNamespacesOK() *GetOrganizationsLimitsNamespacesOK {
	return &GetOrganizationsLimitsNamespacesOK{}
}

/*
GetOrganizationsLimitsNamespacesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetOrganizationsLimitsNamespacesOK struct {
	Payload models.PagedNamespaceListing
}

// IsSuccess returns true when this get organizations limits namespaces o k response has a 2xx status code
func (o *GetOrganizationsLimitsNamespacesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organizations limits namespaces o k response has a 3xx status code
func (o *GetOrganizationsLimitsNamespacesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organizations limits namespaces o k response has a 4xx status code
func (o *GetOrganizationsLimitsNamespacesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organizations limits namespaces o k response has a 5xx status code
func (o *GetOrganizationsLimitsNamespacesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organizations limits namespaces o k response a status code equal to that given
func (o *GetOrganizationsLimitsNamespacesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOrganizationsLimitsNamespacesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces][%d] getOrganizationsLimitsNamespacesOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationsLimitsNamespacesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces][%d] getOrganizationsLimitsNamespacesOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationsLimitsNamespacesOK) GetPayload() models.PagedNamespaceListing {
	return o.Payload
}

func (o *GetOrganizationsLimitsNamespacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsNamespacesBadRequest creates a GetOrganizationsLimitsNamespacesBadRequest with default headers values
func NewGetOrganizationsLimitsNamespacesBadRequest() *GetOrganizationsLimitsNamespacesBadRequest {
	return &GetOrganizationsLimitsNamespacesBadRequest{}
}

/*
GetOrganizationsLimitsNamespacesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOrganizationsLimitsNamespacesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get organizations limits namespaces bad request response has a 2xx status code
func (o *GetOrganizationsLimitsNamespacesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organizations limits namespaces bad request response has a 3xx status code
func (o *GetOrganizationsLimitsNamespacesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organizations limits namespaces bad request response has a 4xx status code
func (o *GetOrganizationsLimitsNamespacesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organizations limits namespaces bad request response has a 5xx status code
func (o *GetOrganizationsLimitsNamespacesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get organizations limits namespaces bad request response a status code equal to that given
func (o *GetOrganizationsLimitsNamespacesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetOrganizationsLimitsNamespacesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces][%d] getOrganizationsLimitsNamespacesBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrganizationsLimitsNamespacesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces][%d] getOrganizationsLimitsNamespacesBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrganizationsLimitsNamespacesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsNamespacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsNamespacesUnauthorized creates a GetOrganizationsLimitsNamespacesUnauthorized with default headers values
func NewGetOrganizationsLimitsNamespacesUnauthorized() *GetOrganizationsLimitsNamespacesUnauthorized {
	return &GetOrganizationsLimitsNamespacesUnauthorized{}
}

/*
GetOrganizationsLimitsNamespacesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOrganizationsLimitsNamespacesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get organizations limits namespaces unauthorized response has a 2xx status code
func (o *GetOrganizationsLimitsNamespacesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organizations limits namespaces unauthorized response has a 3xx status code
func (o *GetOrganizationsLimitsNamespacesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organizations limits namespaces unauthorized response has a 4xx status code
func (o *GetOrganizationsLimitsNamespacesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organizations limits namespaces unauthorized response has a 5xx status code
func (o *GetOrganizationsLimitsNamespacesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get organizations limits namespaces unauthorized response a status code equal to that given
func (o *GetOrganizationsLimitsNamespacesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetOrganizationsLimitsNamespacesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces][%d] getOrganizationsLimitsNamespacesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOrganizationsLimitsNamespacesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces][%d] getOrganizationsLimitsNamespacesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOrganizationsLimitsNamespacesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsNamespacesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsNamespacesForbidden creates a GetOrganizationsLimitsNamespacesForbidden with default headers values
func NewGetOrganizationsLimitsNamespacesForbidden() *GetOrganizationsLimitsNamespacesForbidden {
	return &GetOrganizationsLimitsNamespacesForbidden{}
}

/*
GetOrganizationsLimitsNamespacesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetOrganizationsLimitsNamespacesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get organizations limits namespaces forbidden response has a 2xx status code
func (o *GetOrganizationsLimitsNamespacesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organizations limits namespaces forbidden response has a 3xx status code
func (o *GetOrganizationsLimitsNamespacesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organizations limits namespaces forbidden response has a 4xx status code
func (o *GetOrganizationsLimitsNamespacesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organizations limits namespaces forbidden response has a 5xx status code
func (o *GetOrganizationsLimitsNamespacesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get organizations limits namespaces forbidden response a status code equal to that given
func (o *GetOrganizationsLimitsNamespacesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetOrganizationsLimitsNamespacesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces][%d] getOrganizationsLimitsNamespacesForbidden  %+v", 403, o.Payload)
}

func (o *GetOrganizationsLimitsNamespacesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces][%d] getOrganizationsLimitsNamespacesForbidden  %+v", 403, o.Payload)
}

func (o *GetOrganizationsLimitsNamespacesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsNamespacesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsNamespacesNotFound creates a GetOrganizationsLimitsNamespacesNotFound with default headers values
func NewGetOrganizationsLimitsNamespacesNotFound() *GetOrganizationsLimitsNamespacesNotFound {
	return &GetOrganizationsLimitsNamespacesNotFound{}
}

/*
GetOrganizationsLimitsNamespacesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetOrganizationsLimitsNamespacesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get organizations limits namespaces not found response has a 2xx status code
func (o *GetOrganizationsLimitsNamespacesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organizations limits namespaces not found response has a 3xx status code
func (o *GetOrganizationsLimitsNamespacesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organizations limits namespaces not found response has a 4xx status code
func (o *GetOrganizationsLimitsNamespacesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organizations limits namespaces not found response has a 5xx status code
func (o *GetOrganizationsLimitsNamespacesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get organizations limits namespaces not found response a status code equal to that given
func (o *GetOrganizationsLimitsNamespacesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetOrganizationsLimitsNamespacesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces][%d] getOrganizationsLimitsNamespacesNotFound  %+v", 404, o.Payload)
}

func (o *GetOrganizationsLimitsNamespacesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces][%d] getOrganizationsLimitsNamespacesNotFound  %+v", 404, o.Payload)
}

func (o *GetOrganizationsLimitsNamespacesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsNamespacesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsNamespacesRequestTimeout creates a GetOrganizationsLimitsNamespacesRequestTimeout with default headers values
func NewGetOrganizationsLimitsNamespacesRequestTimeout() *GetOrganizationsLimitsNamespacesRequestTimeout {
	return &GetOrganizationsLimitsNamespacesRequestTimeout{}
}

/*
GetOrganizationsLimitsNamespacesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetOrganizationsLimitsNamespacesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get organizations limits namespaces request timeout response has a 2xx status code
func (o *GetOrganizationsLimitsNamespacesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organizations limits namespaces request timeout response has a 3xx status code
func (o *GetOrganizationsLimitsNamespacesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organizations limits namespaces request timeout response has a 4xx status code
func (o *GetOrganizationsLimitsNamespacesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organizations limits namespaces request timeout response has a 5xx status code
func (o *GetOrganizationsLimitsNamespacesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get organizations limits namespaces request timeout response a status code equal to that given
func (o *GetOrganizationsLimitsNamespacesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetOrganizationsLimitsNamespacesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces][%d] getOrganizationsLimitsNamespacesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOrganizationsLimitsNamespacesRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces][%d] getOrganizationsLimitsNamespacesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOrganizationsLimitsNamespacesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsNamespacesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsNamespacesRequestEntityTooLarge creates a GetOrganizationsLimitsNamespacesRequestEntityTooLarge with default headers values
func NewGetOrganizationsLimitsNamespacesRequestEntityTooLarge() *GetOrganizationsLimitsNamespacesRequestEntityTooLarge {
	return &GetOrganizationsLimitsNamespacesRequestEntityTooLarge{}
}

/*
GetOrganizationsLimitsNamespacesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetOrganizationsLimitsNamespacesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get organizations limits namespaces request entity too large response has a 2xx status code
func (o *GetOrganizationsLimitsNamespacesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organizations limits namespaces request entity too large response has a 3xx status code
func (o *GetOrganizationsLimitsNamespacesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organizations limits namespaces request entity too large response has a 4xx status code
func (o *GetOrganizationsLimitsNamespacesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organizations limits namespaces request entity too large response has a 5xx status code
func (o *GetOrganizationsLimitsNamespacesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get organizations limits namespaces request entity too large response a status code equal to that given
func (o *GetOrganizationsLimitsNamespacesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetOrganizationsLimitsNamespacesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces][%d] getOrganizationsLimitsNamespacesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOrganizationsLimitsNamespacesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces][%d] getOrganizationsLimitsNamespacesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOrganizationsLimitsNamespacesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsNamespacesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsNamespacesUnsupportedMediaType creates a GetOrganizationsLimitsNamespacesUnsupportedMediaType with default headers values
func NewGetOrganizationsLimitsNamespacesUnsupportedMediaType() *GetOrganizationsLimitsNamespacesUnsupportedMediaType {
	return &GetOrganizationsLimitsNamespacesUnsupportedMediaType{}
}

/*
GetOrganizationsLimitsNamespacesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOrganizationsLimitsNamespacesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get organizations limits namespaces unsupported media type response has a 2xx status code
func (o *GetOrganizationsLimitsNamespacesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organizations limits namespaces unsupported media type response has a 3xx status code
func (o *GetOrganizationsLimitsNamespacesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organizations limits namespaces unsupported media type response has a 4xx status code
func (o *GetOrganizationsLimitsNamespacesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organizations limits namespaces unsupported media type response has a 5xx status code
func (o *GetOrganizationsLimitsNamespacesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get organizations limits namespaces unsupported media type response a status code equal to that given
func (o *GetOrganizationsLimitsNamespacesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetOrganizationsLimitsNamespacesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces][%d] getOrganizationsLimitsNamespacesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOrganizationsLimitsNamespacesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces][%d] getOrganizationsLimitsNamespacesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOrganizationsLimitsNamespacesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsNamespacesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsNamespacesTooManyRequests creates a GetOrganizationsLimitsNamespacesTooManyRequests with default headers values
func NewGetOrganizationsLimitsNamespacesTooManyRequests() *GetOrganizationsLimitsNamespacesTooManyRequests {
	return &GetOrganizationsLimitsNamespacesTooManyRequests{}
}

/*
GetOrganizationsLimitsNamespacesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetOrganizationsLimitsNamespacesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get organizations limits namespaces too many requests response has a 2xx status code
func (o *GetOrganizationsLimitsNamespacesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organizations limits namespaces too many requests response has a 3xx status code
func (o *GetOrganizationsLimitsNamespacesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organizations limits namespaces too many requests response has a 4xx status code
func (o *GetOrganizationsLimitsNamespacesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organizations limits namespaces too many requests response has a 5xx status code
func (o *GetOrganizationsLimitsNamespacesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get organizations limits namespaces too many requests response a status code equal to that given
func (o *GetOrganizationsLimitsNamespacesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetOrganizationsLimitsNamespacesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces][%d] getOrganizationsLimitsNamespacesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOrganizationsLimitsNamespacesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces][%d] getOrganizationsLimitsNamespacesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOrganizationsLimitsNamespacesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsNamespacesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsNamespacesInternalServerError creates a GetOrganizationsLimitsNamespacesInternalServerError with default headers values
func NewGetOrganizationsLimitsNamespacesInternalServerError() *GetOrganizationsLimitsNamespacesInternalServerError {
	return &GetOrganizationsLimitsNamespacesInternalServerError{}
}

/*
GetOrganizationsLimitsNamespacesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOrganizationsLimitsNamespacesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get organizations limits namespaces internal server error response has a 2xx status code
func (o *GetOrganizationsLimitsNamespacesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organizations limits namespaces internal server error response has a 3xx status code
func (o *GetOrganizationsLimitsNamespacesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organizations limits namespaces internal server error response has a 4xx status code
func (o *GetOrganizationsLimitsNamespacesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organizations limits namespaces internal server error response has a 5xx status code
func (o *GetOrganizationsLimitsNamespacesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get organizations limits namespaces internal server error response a status code equal to that given
func (o *GetOrganizationsLimitsNamespacesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetOrganizationsLimitsNamespacesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces][%d] getOrganizationsLimitsNamespacesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrganizationsLimitsNamespacesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces][%d] getOrganizationsLimitsNamespacesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrganizationsLimitsNamespacesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsNamespacesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsNamespacesServiceUnavailable creates a GetOrganizationsLimitsNamespacesServiceUnavailable with default headers values
func NewGetOrganizationsLimitsNamespacesServiceUnavailable() *GetOrganizationsLimitsNamespacesServiceUnavailable {
	return &GetOrganizationsLimitsNamespacesServiceUnavailable{}
}

/*
GetOrganizationsLimitsNamespacesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOrganizationsLimitsNamespacesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get organizations limits namespaces service unavailable response has a 2xx status code
func (o *GetOrganizationsLimitsNamespacesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organizations limits namespaces service unavailable response has a 3xx status code
func (o *GetOrganizationsLimitsNamespacesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organizations limits namespaces service unavailable response has a 4xx status code
func (o *GetOrganizationsLimitsNamespacesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organizations limits namespaces service unavailable response has a 5xx status code
func (o *GetOrganizationsLimitsNamespacesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get organizations limits namespaces service unavailable response a status code equal to that given
func (o *GetOrganizationsLimitsNamespacesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetOrganizationsLimitsNamespacesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces][%d] getOrganizationsLimitsNamespacesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOrganizationsLimitsNamespacesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces][%d] getOrganizationsLimitsNamespacesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOrganizationsLimitsNamespacesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsNamespacesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsNamespacesGatewayTimeout creates a GetOrganizationsLimitsNamespacesGatewayTimeout with default headers values
func NewGetOrganizationsLimitsNamespacesGatewayTimeout() *GetOrganizationsLimitsNamespacesGatewayTimeout {
	return &GetOrganizationsLimitsNamespacesGatewayTimeout{}
}

/*
GetOrganizationsLimitsNamespacesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetOrganizationsLimitsNamespacesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get organizations limits namespaces gateway timeout response has a 2xx status code
func (o *GetOrganizationsLimitsNamespacesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organizations limits namespaces gateway timeout response has a 3xx status code
func (o *GetOrganizationsLimitsNamespacesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organizations limits namespaces gateway timeout response has a 4xx status code
func (o *GetOrganizationsLimitsNamespacesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organizations limits namespaces gateway timeout response has a 5xx status code
func (o *GetOrganizationsLimitsNamespacesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get organizations limits namespaces gateway timeout response a status code equal to that given
func (o *GetOrganizationsLimitsNamespacesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetOrganizationsLimitsNamespacesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces][%d] getOrganizationsLimitsNamespacesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOrganizationsLimitsNamespacesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces][%d] getOrganizationsLimitsNamespacesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOrganizationsLimitsNamespacesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsNamespacesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

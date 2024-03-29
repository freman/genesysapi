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

// GetOrganizationsLimitsNamespaceReader is a Reader for the GetOrganizationsLimitsNamespace structure.
type GetOrganizationsLimitsNamespaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationsLimitsNamespaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationsLimitsNamespaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOrganizationsLimitsNamespaceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOrganizationsLimitsNamespaceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrganizationsLimitsNamespaceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrganizationsLimitsNamespaceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOrganizationsLimitsNamespaceRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOrganizationsLimitsNamespaceRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOrganizationsLimitsNamespaceUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOrganizationsLimitsNamespaceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOrganizationsLimitsNamespaceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOrganizationsLimitsNamespaceServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOrganizationsLimitsNamespaceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrganizationsLimitsNamespaceOK creates a GetOrganizationsLimitsNamespaceOK with default headers values
func NewGetOrganizationsLimitsNamespaceOK() *GetOrganizationsLimitsNamespaceOK {
	return &GetOrganizationsLimitsNamespaceOK{}
}

/*
GetOrganizationsLimitsNamespaceOK describes a response with status code 200, with default header values.

successful operation
*/
type GetOrganizationsLimitsNamespaceOK struct {
	Payload *models.LimitsEntityListing
}

// IsSuccess returns true when this get organizations limits namespace o k response has a 2xx status code
func (o *GetOrganizationsLimitsNamespaceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organizations limits namespace o k response has a 3xx status code
func (o *GetOrganizationsLimitsNamespaceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organizations limits namespace o k response has a 4xx status code
func (o *GetOrganizationsLimitsNamespaceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organizations limits namespace o k response has a 5xx status code
func (o *GetOrganizationsLimitsNamespaceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organizations limits namespace o k response a status code equal to that given
func (o *GetOrganizationsLimitsNamespaceOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOrganizationsLimitsNamespaceOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces/{namespaceName}][%d] getOrganizationsLimitsNamespaceOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationsLimitsNamespaceOK) String() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces/{namespaceName}][%d] getOrganizationsLimitsNamespaceOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationsLimitsNamespaceOK) GetPayload() *models.LimitsEntityListing {
	return o.Payload
}

func (o *GetOrganizationsLimitsNamespaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LimitsEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsNamespaceBadRequest creates a GetOrganizationsLimitsNamespaceBadRequest with default headers values
func NewGetOrganizationsLimitsNamespaceBadRequest() *GetOrganizationsLimitsNamespaceBadRequest {
	return &GetOrganizationsLimitsNamespaceBadRequest{}
}

/*
GetOrganizationsLimitsNamespaceBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOrganizationsLimitsNamespaceBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get organizations limits namespace bad request response has a 2xx status code
func (o *GetOrganizationsLimitsNamespaceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organizations limits namespace bad request response has a 3xx status code
func (o *GetOrganizationsLimitsNamespaceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organizations limits namespace bad request response has a 4xx status code
func (o *GetOrganizationsLimitsNamespaceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organizations limits namespace bad request response has a 5xx status code
func (o *GetOrganizationsLimitsNamespaceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get organizations limits namespace bad request response a status code equal to that given
func (o *GetOrganizationsLimitsNamespaceBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetOrganizationsLimitsNamespaceBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces/{namespaceName}][%d] getOrganizationsLimitsNamespaceBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrganizationsLimitsNamespaceBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces/{namespaceName}][%d] getOrganizationsLimitsNamespaceBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrganizationsLimitsNamespaceBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsNamespaceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsNamespaceUnauthorized creates a GetOrganizationsLimitsNamespaceUnauthorized with default headers values
func NewGetOrganizationsLimitsNamespaceUnauthorized() *GetOrganizationsLimitsNamespaceUnauthorized {
	return &GetOrganizationsLimitsNamespaceUnauthorized{}
}

/*
GetOrganizationsLimitsNamespaceUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOrganizationsLimitsNamespaceUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get organizations limits namespace unauthorized response has a 2xx status code
func (o *GetOrganizationsLimitsNamespaceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organizations limits namespace unauthorized response has a 3xx status code
func (o *GetOrganizationsLimitsNamespaceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organizations limits namespace unauthorized response has a 4xx status code
func (o *GetOrganizationsLimitsNamespaceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organizations limits namespace unauthorized response has a 5xx status code
func (o *GetOrganizationsLimitsNamespaceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get organizations limits namespace unauthorized response a status code equal to that given
func (o *GetOrganizationsLimitsNamespaceUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetOrganizationsLimitsNamespaceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces/{namespaceName}][%d] getOrganizationsLimitsNamespaceUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOrganizationsLimitsNamespaceUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces/{namespaceName}][%d] getOrganizationsLimitsNamespaceUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOrganizationsLimitsNamespaceUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsNamespaceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsNamespaceForbidden creates a GetOrganizationsLimitsNamespaceForbidden with default headers values
func NewGetOrganizationsLimitsNamespaceForbidden() *GetOrganizationsLimitsNamespaceForbidden {
	return &GetOrganizationsLimitsNamespaceForbidden{}
}

/*
GetOrganizationsLimitsNamespaceForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetOrganizationsLimitsNamespaceForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get organizations limits namespace forbidden response has a 2xx status code
func (o *GetOrganizationsLimitsNamespaceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organizations limits namespace forbidden response has a 3xx status code
func (o *GetOrganizationsLimitsNamespaceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organizations limits namespace forbidden response has a 4xx status code
func (o *GetOrganizationsLimitsNamespaceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organizations limits namespace forbidden response has a 5xx status code
func (o *GetOrganizationsLimitsNamespaceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get organizations limits namespace forbidden response a status code equal to that given
func (o *GetOrganizationsLimitsNamespaceForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetOrganizationsLimitsNamespaceForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces/{namespaceName}][%d] getOrganizationsLimitsNamespaceForbidden  %+v", 403, o.Payload)
}

func (o *GetOrganizationsLimitsNamespaceForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces/{namespaceName}][%d] getOrganizationsLimitsNamespaceForbidden  %+v", 403, o.Payload)
}

func (o *GetOrganizationsLimitsNamespaceForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsNamespaceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsNamespaceNotFound creates a GetOrganizationsLimitsNamespaceNotFound with default headers values
func NewGetOrganizationsLimitsNamespaceNotFound() *GetOrganizationsLimitsNamespaceNotFound {
	return &GetOrganizationsLimitsNamespaceNotFound{}
}

/*
GetOrganizationsLimitsNamespaceNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetOrganizationsLimitsNamespaceNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get organizations limits namespace not found response has a 2xx status code
func (o *GetOrganizationsLimitsNamespaceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organizations limits namespace not found response has a 3xx status code
func (o *GetOrganizationsLimitsNamespaceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organizations limits namespace not found response has a 4xx status code
func (o *GetOrganizationsLimitsNamespaceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organizations limits namespace not found response has a 5xx status code
func (o *GetOrganizationsLimitsNamespaceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get organizations limits namespace not found response a status code equal to that given
func (o *GetOrganizationsLimitsNamespaceNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetOrganizationsLimitsNamespaceNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces/{namespaceName}][%d] getOrganizationsLimitsNamespaceNotFound  %+v", 404, o.Payload)
}

func (o *GetOrganizationsLimitsNamespaceNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces/{namespaceName}][%d] getOrganizationsLimitsNamespaceNotFound  %+v", 404, o.Payload)
}

func (o *GetOrganizationsLimitsNamespaceNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsNamespaceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsNamespaceRequestTimeout creates a GetOrganizationsLimitsNamespaceRequestTimeout with default headers values
func NewGetOrganizationsLimitsNamespaceRequestTimeout() *GetOrganizationsLimitsNamespaceRequestTimeout {
	return &GetOrganizationsLimitsNamespaceRequestTimeout{}
}

/*
GetOrganizationsLimitsNamespaceRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetOrganizationsLimitsNamespaceRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get organizations limits namespace request timeout response has a 2xx status code
func (o *GetOrganizationsLimitsNamespaceRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organizations limits namespace request timeout response has a 3xx status code
func (o *GetOrganizationsLimitsNamespaceRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organizations limits namespace request timeout response has a 4xx status code
func (o *GetOrganizationsLimitsNamespaceRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organizations limits namespace request timeout response has a 5xx status code
func (o *GetOrganizationsLimitsNamespaceRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get organizations limits namespace request timeout response a status code equal to that given
func (o *GetOrganizationsLimitsNamespaceRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetOrganizationsLimitsNamespaceRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces/{namespaceName}][%d] getOrganizationsLimitsNamespaceRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOrganizationsLimitsNamespaceRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces/{namespaceName}][%d] getOrganizationsLimitsNamespaceRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOrganizationsLimitsNamespaceRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsNamespaceRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsNamespaceRequestEntityTooLarge creates a GetOrganizationsLimitsNamespaceRequestEntityTooLarge with default headers values
func NewGetOrganizationsLimitsNamespaceRequestEntityTooLarge() *GetOrganizationsLimitsNamespaceRequestEntityTooLarge {
	return &GetOrganizationsLimitsNamespaceRequestEntityTooLarge{}
}

/*
GetOrganizationsLimitsNamespaceRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetOrganizationsLimitsNamespaceRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get organizations limits namespace request entity too large response has a 2xx status code
func (o *GetOrganizationsLimitsNamespaceRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organizations limits namespace request entity too large response has a 3xx status code
func (o *GetOrganizationsLimitsNamespaceRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organizations limits namespace request entity too large response has a 4xx status code
func (o *GetOrganizationsLimitsNamespaceRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organizations limits namespace request entity too large response has a 5xx status code
func (o *GetOrganizationsLimitsNamespaceRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get organizations limits namespace request entity too large response a status code equal to that given
func (o *GetOrganizationsLimitsNamespaceRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetOrganizationsLimitsNamespaceRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces/{namespaceName}][%d] getOrganizationsLimitsNamespaceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOrganizationsLimitsNamespaceRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces/{namespaceName}][%d] getOrganizationsLimitsNamespaceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOrganizationsLimitsNamespaceRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsNamespaceRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsNamespaceUnsupportedMediaType creates a GetOrganizationsLimitsNamespaceUnsupportedMediaType with default headers values
func NewGetOrganizationsLimitsNamespaceUnsupportedMediaType() *GetOrganizationsLimitsNamespaceUnsupportedMediaType {
	return &GetOrganizationsLimitsNamespaceUnsupportedMediaType{}
}

/*
GetOrganizationsLimitsNamespaceUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOrganizationsLimitsNamespaceUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get organizations limits namespace unsupported media type response has a 2xx status code
func (o *GetOrganizationsLimitsNamespaceUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organizations limits namespace unsupported media type response has a 3xx status code
func (o *GetOrganizationsLimitsNamespaceUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organizations limits namespace unsupported media type response has a 4xx status code
func (o *GetOrganizationsLimitsNamespaceUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organizations limits namespace unsupported media type response has a 5xx status code
func (o *GetOrganizationsLimitsNamespaceUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get organizations limits namespace unsupported media type response a status code equal to that given
func (o *GetOrganizationsLimitsNamespaceUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetOrganizationsLimitsNamespaceUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces/{namespaceName}][%d] getOrganizationsLimitsNamespaceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOrganizationsLimitsNamespaceUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces/{namespaceName}][%d] getOrganizationsLimitsNamespaceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOrganizationsLimitsNamespaceUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsNamespaceUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsNamespaceTooManyRequests creates a GetOrganizationsLimitsNamespaceTooManyRequests with default headers values
func NewGetOrganizationsLimitsNamespaceTooManyRequests() *GetOrganizationsLimitsNamespaceTooManyRequests {
	return &GetOrganizationsLimitsNamespaceTooManyRequests{}
}

/*
GetOrganizationsLimitsNamespaceTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetOrganizationsLimitsNamespaceTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get organizations limits namespace too many requests response has a 2xx status code
func (o *GetOrganizationsLimitsNamespaceTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organizations limits namespace too many requests response has a 3xx status code
func (o *GetOrganizationsLimitsNamespaceTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organizations limits namespace too many requests response has a 4xx status code
func (o *GetOrganizationsLimitsNamespaceTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organizations limits namespace too many requests response has a 5xx status code
func (o *GetOrganizationsLimitsNamespaceTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get organizations limits namespace too many requests response a status code equal to that given
func (o *GetOrganizationsLimitsNamespaceTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetOrganizationsLimitsNamespaceTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces/{namespaceName}][%d] getOrganizationsLimitsNamespaceTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOrganizationsLimitsNamespaceTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces/{namespaceName}][%d] getOrganizationsLimitsNamespaceTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOrganizationsLimitsNamespaceTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsNamespaceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsNamespaceInternalServerError creates a GetOrganizationsLimitsNamespaceInternalServerError with default headers values
func NewGetOrganizationsLimitsNamespaceInternalServerError() *GetOrganizationsLimitsNamespaceInternalServerError {
	return &GetOrganizationsLimitsNamespaceInternalServerError{}
}

/*
GetOrganizationsLimitsNamespaceInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOrganizationsLimitsNamespaceInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get organizations limits namespace internal server error response has a 2xx status code
func (o *GetOrganizationsLimitsNamespaceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organizations limits namespace internal server error response has a 3xx status code
func (o *GetOrganizationsLimitsNamespaceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organizations limits namespace internal server error response has a 4xx status code
func (o *GetOrganizationsLimitsNamespaceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organizations limits namespace internal server error response has a 5xx status code
func (o *GetOrganizationsLimitsNamespaceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get organizations limits namespace internal server error response a status code equal to that given
func (o *GetOrganizationsLimitsNamespaceInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetOrganizationsLimitsNamespaceInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces/{namespaceName}][%d] getOrganizationsLimitsNamespaceInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrganizationsLimitsNamespaceInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces/{namespaceName}][%d] getOrganizationsLimitsNamespaceInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrganizationsLimitsNamespaceInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsNamespaceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsNamespaceServiceUnavailable creates a GetOrganizationsLimitsNamespaceServiceUnavailable with default headers values
func NewGetOrganizationsLimitsNamespaceServiceUnavailable() *GetOrganizationsLimitsNamespaceServiceUnavailable {
	return &GetOrganizationsLimitsNamespaceServiceUnavailable{}
}

/*
GetOrganizationsLimitsNamespaceServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOrganizationsLimitsNamespaceServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get organizations limits namespace service unavailable response has a 2xx status code
func (o *GetOrganizationsLimitsNamespaceServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organizations limits namespace service unavailable response has a 3xx status code
func (o *GetOrganizationsLimitsNamespaceServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organizations limits namespace service unavailable response has a 4xx status code
func (o *GetOrganizationsLimitsNamespaceServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organizations limits namespace service unavailable response has a 5xx status code
func (o *GetOrganizationsLimitsNamespaceServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get organizations limits namespace service unavailable response a status code equal to that given
func (o *GetOrganizationsLimitsNamespaceServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetOrganizationsLimitsNamespaceServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces/{namespaceName}][%d] getOrganizationsLimitsNamespaceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOrganizationsLimitsNamespaceServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces/{namespaceName}][%d] getOrganizationsLimitsNamespaceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOrganizationsLimitsNamespaceServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsNamespaceServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsNamespaceGatewayTimeout creates a GetOrganizationsLimitsNamespaceGatewayTimeout with default headers values
func NewGetOrganizationsLimitsNamespaceGatewayTimeout() *GetOrganizationsLimitsNamespaceGatewayTimeout {
	return &GetOrganizationsLimitsNamespaceGatewayTimeout{}
}

/*
GetOrganizationsLimitsNamespaceGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetOrganizationsLimitsNamespaceGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get organizations limits namespace gateway timeout response has a 2xx status code
func (o *GetOrganizationsLimitsNamespaceGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organizations limits namespace gateway timeout response has a 3xx status code
func (o *GetOrganizationsLimitsNamespaceGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organizations limits namespace gateway timeout response has a 4xx status code
func (o *GetOrganizationsLimitsNamespaceGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organizations limits namespace gateway timeout response has a 5xx status code
func (o *GetOrganizationsLimitsNamespaceGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get organizations limits namespace gateway timeout response a status code equal to that given
func (o *GetOrganizationsLimitsNamespaceGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetOrganizationsLimitsNamespaceGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces/{namespaceName}][%d] getOrganizationsLimitsNamespaceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOrganizationsLimitsNamespaceGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/namespaces/{namespaceName}][%d] getOrganizationsLimitsNamespaceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOrganizationsLimitsNamespaceGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsNamespaceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

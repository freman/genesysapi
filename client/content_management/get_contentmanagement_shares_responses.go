// Code generated by go-swagger; DO NOT EDIT.

package content_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetContentmanagementSharesReader is a Reader for the GetContentmanagementShares structure.
type GetContentmanagementSharesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetContentmanagementSharesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetContentmanagementSharesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetContentmanagementSharesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetContentmanagementSharesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetContentmanagementSharesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetContentmanagementSharesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetContentmanagementSharesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetContentmanagementSharesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetContentmanagementSharesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetContentmanagementSharesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetContentmanagementSharesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetContentmanagementSharesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetContentmanagementSharesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetContentmanagementSharesOK creates a GetContentmanagementSharesOK with default headers values
func NewGetContentmanagementSharesOK() *GetContentmanagementSharesOK {
	return &GetContentmanagementSharesOK{}
}

/*
GetContentmanagementSharesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetContentmanagementSharesOK struct {
	Payload *models.ShareEntityListing
}

// IsSuccess returns true when this get contentmanagement shares o k response has a 2xx status code
func (o *GetContentmanagementSharesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get contentmanagement shares o k response has a 3xx status code
func (o *GetContentmanagementSharesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement shares o k response has a 4xx status code
func (o *GetContentmanagementSharesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get contentmanagement shares o k response has a 5xx status code
func (o *GetContentmanagementSharesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement shares o k response a status code equal to that given
func (o *GetContentmanagementSharesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetContentmanagementSharesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shares][%d] getContentmanagementSharesOK  %+v", 200, o.Payload)
}

func (o *GetContentmanagementSharesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shares][%d] getContentmanagementSharesOK  %+v", 200, o.Payload)
}

func (o *GetContentmanagementSharesOK) GetPayload() *models.ShareEntityListing {
	return o.Payload
}

func (o *GetContentmanagementSharesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ShareEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSharesBadRequest creates a GetContentmanagementSharesBadRequest with default headers values
func NewGetContentmanagementSharesBadRequest() *GetContentmanagementSharesBadRequest {
	return &GetContentmanagementSharesBadRequest{}
}

/*
GetContentmanagementSharesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetContentmanagementSharesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement shares bad request response has a 2xx status code
func (o *GetContentmanagementSharesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement shares bad request response has a 3xx status code
func (o *GetContentmanagementSharesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement shares bad request response has a 4xx status code
func (o *GetContentmanagementSharesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get contentmanagement shares bad request response has a 5xx status code
func (o *GetContentmanagementSharesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement shares bad request response a status code equal to that given
func (o *GetContentmanagementSharesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetContentmanagementSharesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shares][%d] getContentmanagementSharesBadRequest  %+v", 400, o.Payload)
}

func (o *GetContentmanagementSharesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shares][%d] getContentmanagementSharesBadRequest  %+v", 400, o.Payload)
}

func (o *GetContentmanagementSharesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSharesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSharesUnauthorized creates a GetContentmanagementSharesUnauthorized with default headers values
func NewGetContentmanagementSharesUnauthorized() *GetContentmanagementSharesUnauthorized {
	return &GetContentmanagementSharesUnauthorized{}
}

/*
GetContentmanagementSharesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetContentmanagementSharesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement shares unauthorized response has a 2xx status code
func (o *GetContentmanagementSharesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement shares unauthorized response has a 3xx status code
func (o *GetContentmanagementSharesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement shares unauthorized response has a 4xx status code
func (o *GetContentmanagementSharesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get contentmanagement shares unauthorized response has a 5xx status code
func (o *GetContentmanagementSharesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement shares unauthorized response a status code equal to that given
func (o *GetContentmanagementSharesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetContentmanagementSharesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shares][%d] getContentmanagementSharesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetContentmanagementSharesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shares][%d] getContentmanagementSharesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetContentmanagementSharesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSharesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSharesForbidden creates a GetContentmanagementSharesForbidden with default headers values
func NewGetContentmanagementSharesForbidden() *GetContentmanagementSharesForbidden {
	return &GetContentmanagementSharesForbidden{}
}

/*
GetContentmanagementSharesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetContentmanagementSharesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement shares forbidden response has a 2xx status code
func (o *GetContentmanagementSharesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement shares forbidden response has a 3xx status code
func (o *GetContentmanagementSharesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement shares forbidden response has a 4xx status code
func (o *GetContentmanagementSharesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get contentmanagement shares forbidden response has a 5xx status code
func (o *GetContentmanagementSharesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement shares forbidden response a status code equal to that given
func (o *GetContentmanagementSharesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetContentmanagementSharesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shares][%d] getContentmanagementSharesForbidden  %+v", 403, o.Payload)
}

func (o *GetContentmanagementSharesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shares][%d] getContentmanagementSharesForbidden  %+v", 403, o.Payload)
}

func (o *GetContentmanagementSharesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSharesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSharesNotFound creates a GetContentmanagementSharesNotFound with default headers values
func NewGetContentmanagementSharesNotFound() *GetContentmanagementSharesNotFound {
	return &GetContentmanagementSharesNotFound{}
}

/*
GetContentmanagementSharesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetContentmanagementSharesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement shares not found response has a 2xx status code
func (o *GetContentmanagementSharesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement shares not found response has a 3xx status code
func (o *GetContentmanagementSharesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement shares not found response has a 4xx status code
func (o *GetContentmanagementSharesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get contentmanagement shares not found response has a 5xx status code
func (o *GetContentmanagementSharesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement shares not found response a status code equal to that given
func (o *GetContentmanagementSharesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetContentmanagementSharesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shares][%d] getContentmanagementSharesNotFound  %+v", 404, o.Payload)
}

func (o *GetContentmanagementSharesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shares][%d] getContentmanagementSharesNotFound  %+v", 404, o.Payload)
}

func (o *GetContentmanagementSharesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSharesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSharesRequestTimeout creates a GetContentmanagementSharesRequestTimeout with default headers values
func NewGetContentmanagementSharesRequestTimeout() *GetContentmanagementSharesRequestTimeout {
	return &GetContentmanagementSharesRequestTimeout{}
}

/*
GetContentmanagementSharesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetContentmanagementSharesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement shares request timeout response has a 2xx status code
func (o *GetContentmanagementSharesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement shares request timeout response has a 3xx status code
func (o *GetContentmanagementSharesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement shares request timeout response has a 4xx status code
func (o *GetContentmanagementSharesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get contentmanagement shares request timeout response has a 5xx status code
func (o *GetContentmanagementSharesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement shares request timeout response a status code equal to that given
func (o *GetContentmanagementSharesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetContentmanagementSharesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shares][%d] getContentmanagementSharesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetContentmanagementSharesRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shares][%d] getContentmanagementSharesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetContentmanagementSharesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSharesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSharesRequestEntityTooLarge creates a GetContentmanagementSharesRequestEntityTooLarge with default headers values
func NewGetContentmanagementSharesRequestEntityTooLarge() *GetContentmanagementSharesRequestEntityTooLarge {
	return &GetContentmanagementSharesRequestEntityTooLarge{}
}

/*
GetContentmanagementSharesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetContentmanagementSharesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement shares request entity too large response has a 2xx status code
func (o *GetContentmanagementSharesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement shares request entity too large response has a 3xx status code
func (o *GetContentmanagementSharesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement shares request entity too large response has a 4xx status code
func (o *GetContentmanagementSharesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get contentmanagement shares request entity too large response has a 5xx status code
func (o *GetContentmanagementSharesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement shares request entity too large response a status code equal to that given
func (o *GetContentmanagementSharesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetContentmanagementSharesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shares][%d] getContentmanagementSharesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetContentmanagementSharesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shares][%d] getContentmanagementSharesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetContentmanagementSharesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSharesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSharesUnsupportedMediaType creates a GetContentmanagementSharesUnsupportedMediaType with default headers values
func NewGetContentmanagementSharesUnsupportedMediaType() *GetContentmanagementSharesUnsupportedMediaType {
	return &GetContentmanagementSharesUnsupportedMediaType{}
}

/*
GetContentmanagementSharesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetContentmanagementSharesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement shares unsupported media type response has a 2xx status code
func (o *GetContentmanagementSharesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement shares unsupported media type response has a 3xx status code
func (o *GetContentmanagementSharesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement shares unsupported media type response has a 4xx status code
func (o *GetContentmanagementSharesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get contentmanagement shares unsupported media type response has a 5xx status code
func (o *GetContentmanagementSharesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement shares unsupported media type response a status code equal to that given
func (o *GetContentmanagementSharesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetContentmanagementSharesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shares][%d] getContentmanagementSharesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetContentmanagementSharesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shares][%d] getContentmanagementSharesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetContentmanagementSharesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSharesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSharesTooManyRequests creates a GetContentmanagementSharesTooManyRequests with default headers values
func NewGetContentmanagementSharesTooManyRequests() *GetContentmanagementSharesTooManyRequests {
	return &GetContentmanagementSharesTooManyRequests{}
}

/*
GetContentmanagementSharesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetContentmanagementSharesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement shares too many requests response has a 2xx status code
func (o *GetContentmanagementSharesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement shares too many requests response has a 3xx status code
func (o *GetContentmanagementSharesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement shares too many requests response has a 4xx status code
func (o *GetContentmanagementSharesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get contentmanagement shares too many requests response has a 5xx status code
func (o *GetContentmanagementSharesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement shares too many requests response a status code equal to that given
func (o *GetContentmanagementSharesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetContentmanagementSharesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shares][%d] getContentmanagementSharesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetContentmanagementSharesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shares][%d] getContentmanagementSharesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetContentmanagementSharesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSharesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSharesInternalServerError creates a GetContentmanagementSharesInternalServerError with default headers values
func NewGetContentmanagementSharesInternalServerError() *GetContentmanagementSharesInternalServerError {
	return &GetContentmanagementSharesInternalServerError{}
}

/*
GetContentmanagementSharesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetContentmanagementSharesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement shares internal server error response has a 2xx status code
func (o *GetContentmanagementSharesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement shares internal server error response has a 3xx status code
func (o *GetContentmanagementSharesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement shares internal server error response has a 4xx status code
func (o *GetContentmanagementSharesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get contentmanagement shares internal server error response has a 5xx status code
func (o *GetContentmanagementSharesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get contentmanagement shares internal server error response a status code equal to that given
func (o *GetContentmanagementSharesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetContentmanagementSharesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shares][%d] getContentmanagementSharesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetContentmanagementSharesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shares][%d] getContentmanagementSharesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetContentmanagementSharesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSharesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSharesServiceUnavailable creates a GetContentmanagementSharesServiceUnavailable with default headers values
func NewGetContentmanagementSharesServiceUnavailable() *GetContentmanagementSharesServiceUnavailable {
	return &GetContentmanagementSharesServiceUnavailable{}
}

/*
GetContentmanagementSharesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetContentmanagementSharesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement shares service unavailable response has a 2xx status code
func (o *GetContentmanagementSharesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement shares service unavailable response has a 3xx status code
func (o *GetContentmanagementSharesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement shares service unavailable response has a 4xx status code
func (o *GetContentmanagementSharesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get contentmanagement shares service unavailable response has a 5xx status code
func (o *GetContentmanagementSharesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get contentmanagement shares service unavailable response a status code equal to that given
func (o *GetContentmanagementSharesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetContentmanagementSharesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shares][%d] getContentmanagementSharesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetContentmanagementSharesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shares][%d] getContentmanagementSharesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetContentmanagementSharesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSharesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSharesGatewayTimeout creates a GetContentmanagementSharesGatewayTimeout with default headers values
func NewGetContentmanagementSharesGatewayTimeout() *GetContentmanagementSharesGatewayTimeout {
	return &GetContentmanagementSharesGatewayTimeout{}
}

/*
GetContentmanagementSharesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetContentmanagementSharesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement shares gateway timeout response has a 2xx status code
func (o *GetContentmanagementSharesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement shares gateway timeout response has a 3xx status code
func (o *GetContentmanagementSharesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement shares gateway timeout response has a 4xx status code
func (o *GetContentmanagementSharesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get contentmanagement shares gateway timeout response has a 5xx status code
func (o *GetContentmanagementSharesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get contentmanagement shares gateway timeout response a status code equal to that given
func (o *GetContentmanagementSharesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetContentmanagementSharesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shares][%d] getContentmanagementSharesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetContentmanagementSharesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shares][%d] getContentmanagementSharesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetContentmanagementSharesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSharesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

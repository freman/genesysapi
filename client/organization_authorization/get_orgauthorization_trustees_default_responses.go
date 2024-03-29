// Code generated by go-swagger; DO NOT EDIT.

package organization_authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetOrgauthorizationTrusteesDefaultReader is a Reader for the GetOrgauthorizationTrusteesDefault structure.
type GetOrgauthorizationTrusteesDefaultReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrgauthorizationTrusteesDefaultReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrgauthorizationTrusteesDefaultOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOrgauthorizationTrusteesDefaultBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOrgauthorizationTrusteesDefaultUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrgauthorizationTrusteesDefaultForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrgauthorizationTrusteesDefaultNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOrgauthorizationTrusteesDefaultRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOrgauthorizationTrusteesDefaultRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOrgauthorizationTrusteesDefaultUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOrgauthorizationTrusteesDefaultTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOrgauthorizationTrusteesDefaultInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOrgauthorizationTrusteesDefaultServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOrgauthorizationTrusteesDefaultGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrgauthorizationTrusteesDefaultOK creates a GetOrgauthorizationTrusteesDefaultOK with default headers values
func NewGetOrgauthorizationTrusteesDefaultOK() *GetOrgauthorizationTrusteesDefaultOK {
	return &GetOrgauthorizationTrusteesDefaultOK{}
}

/*
GetOrgauthorizationTrusteesDefaultOK describes a response with status code 200, with default header values.

successful operation
*/
type GetOrgauthorizationTrusteesDefaultOK struct {
	Payload *models.Trustee
}

// IsSuccess returns true when this get orgauthorization trustees default o k response has a 2xx status code
func (o *GetOrgauthorizationTrusteesDefaultOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get orgauthorization trustees default o k response has a 3xx status code
func (o *GetOrgauthorizationTrusteesDefaultOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustees default o k response has a 4xx status code
func (o *GetOrgauthorizationTrusteesDefaultOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get orgauthorization trustees default o k response has a 5xx status code
func (o *GetOrgauthorizationTrusteesDefaultOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustees default o k response a status code equal to that given
func (o *GetOrgauthorizationTrusteesDefaultOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOrgauthorizationTrusteesDefaultOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/default][%d] getOrgauthorizationTrusteesDefaultOK  %+v", 200, o.Payload)
}

func (o *GetOrgauthorizationTrusteesDefaultOK) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/default][%d] getOrgauthorizationTrusteesDefaultOK  %+v", 200, o.Payload)
}

func (o *GetOrgauthorizationTrusteesDefaultOK) GetPayload() *models.Trustee {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteesDefaultOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Trustee)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteesDefaultBadRequest creates a GetOrgauthorizationTrusteesDefaultBadRequest with default headers values
func NewGetOrgauthorizationTrusteesDefaultBadRequest() *GetOrgauthorizationTrusteesDefaultBadRequest {
	return &GetOrgauthorizationTrusteesDefaultBadRequest{}
}

/*
GetOrgauthorizationTrusteesDefaultBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOrgauthorizationTrusteesDefaultBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustees default bad request response has a 2xx status code
func (o *GetOrgauthorizationTrusteesDefaultBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustees default bad request response has a 3xx status code
func (o *GetOrgauthorizationTrusteesDefaultBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustees default bad request response has a 4xx status code
func (o *GetOrgauthorizationTrusteesDefaultBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get orgauthorization trustees default bad request response has a 5xx status code
func (o *GetOrgauthorizationTrusteesDefaultBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustees default bad request response a status code equal to that given
func (o *GetOrgauthorizationTrusteesDefaultBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetOrgauthorizationTrusteesDefaultBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/default][%d] getOrgauthorizationTrusteesDefaultBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrgauthorizationTrusteesDefaultBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/default][%d] getOrgauthorizationTrusteesDefaultBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrgauthorizationTrusteesDefaultBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteesDefaultBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteesDefaultUnauthorized creates a GetOrgauthorizationTrusteesDefaultUnauthorized with default headers values
func NewGetOrgauthorizationTrusteesDefaultUnauthorized() *GetOrgauthorizationTrusteesDefaultUnauthorized {
	return &GetOrgauthorizationTrusteesDefaultUnauthorized{}
}

/*
GetOrgauthorizationTrusteesDefaultUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOrgauthorizationTrusteesDefaultUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustees default unauthorized response has a 2xx status code
func (o *GetOrgauthorizationTrusteesDefaultUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustees default unauthorized response has a 3xx status code
func (o *GetOrgauthorizationTrusteesDefaultUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustees default unauthorized response has a 4xx status code
func (o *GetOrgauthorizationTrusteesDefaultUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get orgauthorization trustees default unauthorized response has a 5xx status code
func (o *GetOrgauthorizationTrusteesDefaultUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustees default unauthorized response a status code equal to that given
func (o *GetOrgauthorizationTrusteesDefaultUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetOrgauthorizationTrusteesDefaultUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/default][%d] getOrgauthorizationTrusteesDefaultUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOrgauthorizationTrusteesDefaultUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/default][%d] getOrgauthorizationTrusteesDefaultUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOrgauthorizationTrusteesDefaultUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteesDefaultUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteesDefaultForbidden creates a GetOrgauthorizationTrusteesDefaultForbidden with default headers values
func NewGetOrgauthorizationTrusteesDefaultForbidden() *GetOrgauthorizationTrusteesDefaultForbidden {
	return &GetOrgauthorizationTrusteesDefaultForbidden{}
}

/*
GetOrgauthorizationTrusteesDefaultForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetOrgauthorizationTrusteesDefaultForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustees default forbidden response has a 2xx status code
func (o *GetOrgauthorizationTrusteesDefaultForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustees default forbidden response has a 3xx status code
func (o *GetOrgauthorizationTrusteesDefaultForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustees default forbidden response has a 4xx status code
func (o *GetOrgauthorizationTrusteesDefaultForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get orgauthorization trustees default forbidden response has a 5xx status code
func (o *GetOrgauthorizationTrusteesDefaultForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustees default forbidden response a status code equal to that given
func (o *GetOrgauthorizationTrusteesDefaultForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetOrgauthorizationTrusteesDefaultForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/default][%d] getOrgauthorizationTrusteesDefaultForbidden  %+v", 403, o.Payload)
}

func (o *GetOrgauthorizationTrusteesDefaultForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/default][%d] getOrgauthorizationTrusteesDefaultForbidden  %+v", 403, o.Payload)
}

func (o *GetOrgauthorizationTrusteesDefaultForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteesDefaultForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteesDefaultNotFound creates a GetOrgauthorizationTrusteesDefaultNotFound with default headers values
func NewGetOrgauthorizationTrusteesDefaultNotFound() *GetOrgauthorizationTrusteesDefaultNotFound {
	return &GetOrgauthorizationTrusteesDefaultNotFound{}
}

/*
GetOrgauthorizationTrusteesDefaultNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetOrgauthorizationTrusteesDefaultNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustees default not found response has a 2xx status code
func (o *GetOrgauthorizationTrusteesDefaultNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustees default not found response has a 3xx status code
func (o *GetOrgauthorizationTrusteesDefaultNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustees default not found response has a 4xx status code
func (o *GetOrgauthorizationTrusteesDefaultNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get orgauthorization trustees default not found response has a 5xx status code
func (o *GetOrgauthorizationTrusteesDefaultNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustees default not found response a status code equal to that given
func (o *GetOrgauthorizationTrusteesDefaultNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetOrgauthorizationTrusteesDefaultNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/default][%d] getOrgauthorizationTrusteesDefaultNotFound  %+v", 404, o.Payload)
}

func (o *GetOrgauthorizationTrusteesDefaultNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/default][%d] getOrgauthorizationTrusteesDefaultNotFound  %+v", 404, o.Payload)
}

func (o *GetOrgauthorizationTrusteesDefaultNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteesDefaultNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteesDefaultRequestTimeout creates a GetOrgauthorizationTrusteesDefaultRequestTimeout with default headers values
func NewGetOrgauthorizationTrusteesDefaultRequestTimeout() *GetOrgauthorizationTrusteesDefaultRequestTimeout {
	return &GetOrgauthorizationTrusteesDefaultRequestTimeout{}
}

/*
GetOrgauthorizationTrusteesDefaultRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetOrgauthorizationTrusteesDefaultRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustees default request timeout response has a 2xx status code
func (o *GetOrgauthorizationTrusteesDefaultRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustees default request timeout response has a 3xx status code
func (o *GetOrgauthorizationTrusteesDefaultRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustees default request timeout response has a 4xx status code
func (o *GetOrgauthorizationTrusteesDefaultRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get orgauthorization trustees default request timeout response has a 5xx status code
func (o *GetOrgauthorizationTrusteesDefaultRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustees default request timeout response a status code equal to that given
func (o *GetOrgauthorizationTrusteesDefaultRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetOrgauthorizationTrusteesDefaultRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/default][%d] getOrgauthorizationTrusteesDefaultRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOrgauthorizationTrusteesDefaultRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/default][%d] getOrgauthorizationTrusteesDefaultRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOrgauthorizationTrusteesDefaultRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteesDefaultRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteesDefaultRequestEntityTooLarge creates a GetOrgauthorizationTrusteesDefaultRequestEntityTooLarge with default headers values
func NewGetOrgauthorizationTrusteesDefaultRequestEntityTooLarge() *GetOrgauthorizationTrusteesDefaultRequestEntityTooLarge {
	return &GetOrgauthorizationTrusteesDefaultRequestEntityTooLarge{}
}

/*
GetOrgauthorizationTrusteesDefaultRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetOrgauthorizationTrusteesDefaultRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustees default request entity too large response has a 2xx status code
func (o *GetOrgauthorizationTrusteesDefaultRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustees default request entity too large response has a 3xx status code
func (o *GetOrgauthorizationTrusteesDefaultRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustees default request entity too large response has a 4xx status code
func (o *GetOrgauthorizationTrusteesDefaultRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get orgauthorization trustees default request entity too large response has a 5xx status code
func (o *GetOrgauthorizationTrusteesDefaultRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustees default request entity too large response a status code equal to that given
func (o *GetOrgauthorizationTrusteesDefaultRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetOrgauthorizationTrusteesDefaultRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/default][%d] getOrgauthorizationTrusteesDefaultRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOrgauthorizationTrusteesDefaultRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/default][%d] getOrgauthorizationTrusteesDefaultRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOrgauthorizationTrusteesDefaultRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteesDefaultRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteesDefaultUnsupportedMediaType creates a GetOrgauthorizationTrusteesDefaultUnsupportedMediaType with default headers values
func NewGetOrgauthorizationTrusteesDefaultUnsupportedMediaType() *GetOrgauthorizationTrusteesDefaultUnsupportedMediaType {
	return &GetOrgauthorizationTrusteesDefaultUnsupportedMediaType{}
}

/*
GetOrgauthorizationTrusteesDefaultUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOrgauthorizationTrusteesDefaultUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustees default unsupported media type response has a 2xx status code
func (o *GetOrgauthorizationTrusteesDefaultUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustees default unsupported media type response has a 3xx status code
func (o *GetOrgauthorizationTrusteesDefaultUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustees default unsupported media type response has a 4xx status code
func (o *GetOrgauthorizationTrusteesDefaultUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get orgauthorization trustees default unsupported media type response has a 5xx status code
func (o *GetOrgauthorizationTrusteesDefaultUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustees default unsupported media type response a status code equal to that given
func (o *GetOrgauthorizationTrusteesDefaultUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetOrgauthorizationTrusteesDefaultUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/default][%d] getOrgauthorizationTrusteesDefaultUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOrgauthorizationTrusteesDefaultUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/default][%d] getOrgauthorizationTrusteesDefaultUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOrgauthorizationTrusteesDefaultUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteesDefaultUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteesDefaultTooManyRequests creates a GetOrgauthorizationTrusteesDefaultTooManyRequests with default headers values
func NewGetOrgauthorizationTrusteesDefaultTooManyRequests() *GetOrgauthorizationTrusteesDefaultTooManyRequests {
	return &GetOrgauthorizationTrusteesDefaultTooManyRequests{}
}

/*
GetOrgauthorizationTrusteesDefaultTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetOrgauthorizationTrusteesDefaultTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustees default too many requests response has a 2xx status code
func (o *GetOrgauthorizationTrusteesDefaultTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustees default too many requests response has a 3xx status code
func (o *GetOrgauthorizationTrusteesDefaultTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustees default too many requests response has a 4xx status code
func (o *GetOrgauthorizationTrusteesDefaultTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get orgauthorization trustees default too many requests response has a 5xx status code
func (o *GetOrgauthorizationTrusteesDefaultTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustees default too many requests response a status code equal to that given
func (o *GetOrgauthorizationTrusteesDefaultTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetOrgauthorizationTrusteesDefaultTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/default][%d] getOrgauthorizationTrusteesDefaultTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOrgauthorizationTrusteesDefaultTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/default][%d] getOrgauthorizationTrusteesDefaultTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOrgauthorizationTrusteesDefaultTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteesDefaultTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteesDefaultInternalServerError creates a GetOrgauthorizationTrusteesDefaultInternalServerError with default headers values
func NewGetOrgauthorizationTrusteesDefaultInternalServerError() *GetOrgauthorizationTrusteesDefaultInternalServerError {
	return &GetOrgauthorizationTrusteesDefaultInternalServerError{}
}

/*
GetOrgauthorizationTrusteesDefaultInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOrgauthorizationTrusteesDefaultInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustees default internal server error response has a 2xx status code
func (o *GetOrgauthorizationTrusteesDefaultInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustees default internal server error response has a 3xx status code
func (o *GetOrgauthorizationTrusteesDefaultInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustees default internal server error response has a 4xx status code
func (o *GetOrgauthorizationTrusteesDefaultInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get orgauthorization trustees default internal server error response has a 5xx status code
func (o *GetOrgauthorizationTrusteesDefaultInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get orgauthorization trustees default internal server error response a status code equal to that given
func (o *GetOrgauthorizationTrusteesDefaultInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetOrgauthorizationTrusteesDefaultInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/default][%d] getOrgauthorizationTrusteesDefaultInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrgauthorizationTrusteesDefaultInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/default][%d] getOrgauthorizationTrusteesDefaultInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrgauthorizationTrusteesDefaultInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteesDefaultInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteesDefaultServiceUnavailable creates a GetOrgauthorizationTrusteesDefaultServiceUnavailable with default headers values
func NewGetOrgauthorizationTrusteesDefaultServiceUnavailable() *GetOrgauthorizationTrusteesDefaultServiceUnavailable {
	return &GetOrgauthorizationTrusteesDefaultServiceUnavailable{}
}

/*
GetOrgauthorizationTrusteesDefaultServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOrgauthorizationTrusteesDefaultServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustees default service unavailable response has a 2xx status code
func (o *GetOrgauthorizationTrusteesDefaultServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustees default service unavailable response has a 3xx status code
func (o *GetOrgauthorizationTrusteesDefaultServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustees default service unavailable response has a 4xx status code
func (o *GetOrgauthorizationTrusteesDefaultServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get orgauthorization trustees default service unavailable response has a 5xx status code
func (o *GetOrgauthorizationTrusteesDefaultServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get orgauthorization trustees default service unavailable response a status code equal to that given
func (o *GetOrgauthorizationTrusteesDefaultServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetOrgauthorizationTrusteesDefaultServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/default][%d] getOrgauthorizationTrusteesDefaultServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOrgauthorizationTrusteesDefaultServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/default][%d] getOrgauthorizationTrusteesDefaultServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOrgauthorizationTrusteesDefaultServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteesDefaultServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteesDefaultGatewayTimeout creates a GetOrgauthorizationTrusteesDefaultGatewayTimeout with default headers values
func NewGetOrgauthorizationTrusteesDefaultGatewayTimeout() *GetOrgauthorizationTrusteesDefaultGatewayTimeout {
	return &GetOrgauthorizationTrusteesDefaultGatewayTimeout{}
}

/*
GetOrgauthorizationTrusteesDefaultGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetOrgauthorizationTrusteesDefaultGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustees default gateway timeout response has a 2xx status code
func (o *GetOrgauthorizationTrusteesDefaultGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustees default gateway timeout response has a 3xx status code
func (o *GetOrgauthorizationTrusteesDefaultGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustees default gateway timeout response has a 4xx status code
func (o *GetOrgauthorizationTrusteesDefaultGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get orgauthorization trustees default gateway timeout response has a 5xx status code
func (o *GetOrgauthorizationTrusteesDefaultGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get orgauthorization trustees default gateway timeout response a status code equal to that given
func (o *GetOrgauthorizationTrusteesDefaultGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetOrgauthorizationTrusteesDefaultGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/default][%d] getOrgauthorizationTrusteesDefaultGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOrgauthorizationTrusteesDefaultGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/default][%d] getOrgauthorizationTrusteesDefaultGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOrgauthorizationTrusteesDefaultGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteesDefaultGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

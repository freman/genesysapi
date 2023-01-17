// Code generated by go-swagger; DO NOT EDIT.

package s_c_i_m

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetScimV2ServiceproviderconfigReader is a Reader for the GetScimV2Serviceproviderconfig structure.
type GetScimV2ServiceproviderconfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScimV2ServiceproviderconfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScimV2ServiceproviderconfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetScimV2ServiceproviderconfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetScimV2ServiceproviderconfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetScimV2ServiceproviderconfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetScimV2ServiceproviderconfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetScimV2ServiceproviderconfigRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetScimV2ServiceproviderconfigRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetScimV2ServiceproviderconfigUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetScimV2ServiceproviderconfigTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetScimV2ServiceproviderconfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetScimV2ServiceproviderconfigServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetScimV2ServiceproviderconfigGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetScimV2ServiceproviderconfigOK creates a GetScimV2ServiceproviderconfigOK with default headers values
func NewGetScimV2ServiceproviderconfigOK() *GetScimV2ServiceproviderconfigOK {
	return &GetScimV2ServiceproviderconfigOK{}
}

/*
GetScimV2ServiceproviderconfigOK describes a response with status code 200, with default header values.

successful operation
*/
type GetScimV2ServiceproviderconfigOK struct {
	Payload *models.ScimServiceProviderConfig
}

// IsSuccess returns true when this get scim v2 serviceproviderconfig o k response has a 2xx status code
func (o *GetScimV2ServiceproviderconfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get scim v2 serviceproviderconfig o k response has a 3xx status code
func (o *GetScimV2ServiceproviderconfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scim v2 serviceproviderconfig o k response has a 4xx status code
func (o *GetScimV2ServiceproviderconfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get scim v2 serviceproviderconfig o k response has a 5xx status code
func (o *GetScimV2ServiceproviderconfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get scim v2 serviceproviderconfig o k response a status code equal to that given
func (o *GetScimV2ServiceproviderconfigOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetScimV2ServiceproviderconfigOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/serviceproviderconfig][%d] getScimV2ServiceproviderconfigOK  %+v", 200, o.Payload)
}

func (o *GetScimV2ServiceproviderconfigOK) String() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/serviceproviderconfig][%d] getScimV2ServiceproviderconfigOK  %+v", 200, o.Payload)
}

func (o *GetScimV2ServiceproviderconfigOK) GetPayload() *models.ScimServiceProviderConfig {
	return o.Payload
}

func (o *GetScimV2ServiceproviderconfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScimServiceProviderConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2ServiceproviderconfigBadRequest creates a GetScimV2ServiceproviderconfigBadRequest with default headers values
func NewGetScimV2ServiceproviderconfigBadRequest() *GetScimV2ServiceproviderconfigBadRequest {
	return &GetScimV2ServiceproviderconfigBadRequest{}
}

/*
GetScimV2ServiceproviderconfigBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetScimV2ServiceproviderconfigBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scim v2 serviceproviderconfig bad request response has a 2xx status code
func (o *GetScimV2ServiceproviderconfigBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scim v2 serviceproviderconfig bad request response has a 3xx status code
func (o *GetScimV2ServiceproviderconfigBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scim v2 serviceproviderconfig bad request response has a 4xx status code
func (o *GetScimV2ServiceproviderconfigBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scim v2 serviceproviderconfig bad request response has a 5xx status code
func (o *GetScimV2ServiceproviderconfigBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get scim v2 serviceproviderconfig bad request response a status code equal to that given
func (o *GetScimV2ServiceproviderconfigBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetScimV2ServiceproviderconfigBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/serviceproviderconfig][%d] getScimV2ServiceproviderconfigBadRequest  %+v", 400, o.Payload)
}

func (o *GetScimV2ServiceproviderconfigBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/serviceproviderconfig][%d] getScimV2ServiceproviderconfigBadRequest  %+v", 400, o.Payload)
}

func (o *GetScimV2ServiceproviderconfigBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2ServiceproviderconfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2ServiceproviderconfigUnauthorized creates a GetScimV2ServiceproviderconfigUnauthorized with default headers values
func NewGetScimV2ServiceproviderconfigUnauthorized() *GetScimV2ServiceproviderconfigUnauthorized {
	return &GetScimV2ServiceproviderconfigUnauthorized{}
}

/*
GetScimV2ServiceproviderconfigUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetScimV2ServiceproviderconfigUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scim v2 serviceproviderconfig unauthorized response has a 2xx status code
func (o *GetScimV2ServiceproviderconfigUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scim v2 serviceproviderconfig unauthorized response has a 3xx status code
func (o *GetScimV2ServiceproviderconfigUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scim v2 serviceproviderconfig unauthorized response has a 4xx status code
func (o *GetScimV2ServiceproviderconfigUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scim v2 serviceproviderconfig unauthorized response has a 5xx status code
func (o *GetScimV2ServiceproviderconfigUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get scim v2 serviceproviderconfig unauthorized response a status code equal to that given
func (o *GetScimV2ServiceproviderconfigUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetScimV2ServiceproviderconfigUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/serviceproviderconfig][%d] getScimV2ServiceproviderconfigUnauthorized  %+v", 401, o.Payload)
}

func (o *GetScimV2ServiceproviderconfigUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/serviceproviderconfig][%d] getScimV2ServiceproviderconfigUnauthorized  %+v", 401, o.Payload)
}

func (o *GetScimV2ServiceproviderconfigUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2ServiceproviderconfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2ServiceproviderconfigForbidden creates a GetScimV2ServiceproviderconfigForbidden with default headers values
func NewGetScimV2ServiceproviderconfigForbidden() *GetScimV2ServiceproviderconfigForbidden {
	return &GetScimV2ServiceproviderconfigForbidden{}
}

/*
GetScimV2ServiceproviderconfigForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetScimV2ServiceproviderconfigForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scim v2 serviceproviderconfig forbidden response has a 2xx status code
func (o *GetScimV2ServiceproviderconfigForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scim v2 serviceproviderconfig forbidden response has a 3xx status code
func (o *GetScimV2ServiceproviderconfigForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scim v2 serviceproviderconfig forbidden response has a 4xx status code
func (o *GetScimV2ServiceproviderconfigForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scim v2 serviceproviderconfig forbidden response has a 5xx status code
func (o *GetScimV2ServiceproviderconfigForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get scim v2 serviceproviderconfig forbidden response a status code equal to that given
func (o *GetScimV2ServiceproviderconfigForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetScimV2ServiceproviderconfigForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/serviceproviderconfig][%d] getScimV2ServiceproviderconfigForbidden  %+v", 403, o.Payload)
}

func (o *GetScimV2ServiceproviderconfigForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/serviceproviderconfig][%d] getScimV2ServiceproviderconfigForbidden  %+v", 403, o.Payload)
}

func (o *GetScimV2ServiceproviderconfigForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2ServiceproviderconfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2ServiceproviderconfigNotFound creates a GetScimV2ServiceproviderconfigNotFound with default headers values
func NewGetScimV2ServiceproviderconfigNotFound() *GetScimV2ServiceproviderconfigNotFound {
	return &GetScimV2ServiceproviderconfigNotFound{}
}

/*
GetScimV2ServiceproviderconfigNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetScimV2ServiceproviderconfigNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scim v2 serviceproviderconfig not found response has a 2xx status code
func (o *GetScimV2ServiceproviderconfigNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scim v2 serviceproviderconfig not found response has a 3xx status code
func (o *GetScimV2ServiceproviderconfigNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scim v2 serviceproviderconfig not found response has a 4xx status code
func (o *GetScimV2ServiceproviderconfigNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scim v2 serviceproviderconfig not found response has a 5xx status code
func (o *GetScimV2ServiceproviderconfigNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get scim v2 serviceproviderconfig not found response a status code equal to that given
func (o *GetScimV2ServiceproviderconfigNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetScimV2ServiceproviderconfigNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/serviceproviderconfig][%d] getScimV2ServiceproviderconfigNotFound  %+v", 404, o.Payload)
}

func (o *GetScimV2ServiceproviderconfigNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/serviceproviderconfig][%d] getScimV2ServiceproviderconfigNotFound  %+v", 404, o.Payload)
}

func (o *GetScimV2ServiceproviderconfigNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2ServiceproviderconfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2ServiceproviderconfigRequestTimeout creates a GetScimV2ServiceproviderconfigRequestTimeout with default headers values
func NewGetScimV2ServiceproviderconfigRequestTimeout() *GetScimV2ServiceproviderconfigRequestTimeout {
	return &GetScimV2ServiceproviderconfigRequestTimeout{}
}

/*
GetScimV2ServiceproviderconfigRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetScimV2ServiceproviderconfigRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scim v2 serviceproviderconfig request timeout response has a 2xx status code
func (o *GetScimV2ServiceproviderconfigRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scim v2 serviceproviderconfig request timeout response has a 3xx status code
func (o *GetScimV2ServiceproviderconfigRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scim v2 serviceproviderconfig request timeout response has a 4xx status code
func (o *GetScimV2ServiceproviderconfigRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scim v2 serviceproviderconfig request timeout response has a 5xx status code
func (o *GetScimV2ServiceproviderconfigRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get scim v2 serviceproviderconfig request timeout response a status code equal to that given
func (o *GetScimV2ServiceproviderconfigRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetScimV2ServiceproviderconfigRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/serviceproviderconfig][%d] getScimV2ServiceproviderconfigRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetScimV2ServiceproviderconfigRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/serviceproviderconfig][%d] getScimV2ServiceproviderconfigRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetScimV2ServiceproviderconfigRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2ServiceproviderconfigRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2ServiceproviderconfigRequestEntityTooLarge creates a GetScimV2ServiceproviderconfigRequestEntityTooLarge with default headers values
func NewGetScimV2ServiceproviderconfigRequestEntityTooLarge() *GetScimV2ServiceproviderconfigRequestEntityTooLarge {
	return &GetScimV2ServiceproviderconfigRequestEntityTooLarge{}
}

/*
GetScimV2ServiceproviderconfigRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetScimV2ServiceproviderconfigRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scim v2 serviceproviderconfig request entity too large response has a 2xx status code
func (o *GetScimV2ServiceproviderconfigRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scim v2 serviceproviderconfig request entity too large response has a 3xx status code
func (o *GetScimV2ServiceproviderconfigRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scim v2 serviceproviderconfig request entity too large response has a 4xx status code
func (o *GetScimV2ServiceproviderconfigRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scim v2 serviceproviderconfig request entity too large response has a 5xx status code
func (o *GetScimV2ServiceproviderconfigRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get scim v2 serviceproviderconfig request entity too large response a status code equal to that given
func (o *GetScimV2ServiceproviderconfigRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetScimV2ServiceproviderconfigRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/serviceproviderconfig][%d] getScimV2ServiceproviderconfigRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetScimV2ServiceproviderconfigRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/serviceproviderconfig][%d] getScimV2ServiceproviderconfigRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetScimV2ServiceproviderconfigRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2ServiceproviderconfigRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2ServiceproviderconfigUnsupportedMediaType creates a GetScimV2ServiceproviderconfigUnsupportedMediaType with default headers values
func NewGetScimV2ServiceproviderconfigUnsupportedMediaType() *GetScimV2ServiceproviderconfigUnsupportedMediaType {
	return &GetScimV2ServiceproviderconfigUnsupportedMediaType{}
}

/*
GetScimV2ServiceproviderconfigUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetScimV2ServiceproviderconfigUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scim v2 serviceproviderconfig unsupported media type response has a 2xx status code
func (o *GetScimV2ServiceproviderconfigUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scim v2 serviceproviderconfig unsupported media type response has a 3xx status code
func (o *GetScimV2ServiceproviderconfigUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scim v2 serviceproviderconfig unsupported media type response has a 4xx status code
func (o *GetScimV2ServiceproviderconfigUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scim v2 serviceproviderconfig unsupported media type response has a 5xx status code
func (o *GetScimV2ServiceproviderconfigUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get scim v2 serviceproviderconfig unsupported media type response a status code equal to that given
func (o *GetScimV2ServiceproviderconfigUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetScimV2ServiceproviderconfigUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/serviceproviderconfig][%d] getScimV2ServiceproviderconfigUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetScimV2ServiceproviderconfigUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/serviceproviderconfig][%d] getScimV2ServiceproviderconfigUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetScimV2ServiceproviderconfigUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2ServiceproviderconfigUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2ServiceproviderconfigTooManyRequests creates a GetScimV2ServiceproviderconfigTooManyRequests with default headers values
func NewGetScimV2ServiceproviderconfigTooManyRequests() *GetScimV2ServiceproviderconfigTooManyRequests {
	return &GetScimV2ServiceproviderconfigTooManyRequests{}
}

/*
GetScimV2ServiceproviderconfigTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetScimV2ServiceproviderconfigTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scim v2 serviceproviderconfig too many requests response has a 2xx status code
func (o *GetScimV2ServiceproviderconfigTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scim v2 serviceproviderconfig too many requests response has a 3xx status code
func (o *GetScimV2ServiceproviderconfigTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scim v2 serviceproviderconfig too many requests response has a 4xx status code
func (o *GetScimV2ServiceproviderconfigTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scim v2 serviceproviderconfig too many requests response has a 5xx status code
func (o *GetScimV2ServiceproviderconfigTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get scim v2 serviceproviderconfig too many requests response a status code equal to that given
func (o *GetScimV2ServiceproviderconfigTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetScimV2ServiceproviderconfigTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/serviceproviderconfig][%d] getScimV2ServiceproviderconfigTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetScimV2ServiceproviderconfigTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/serviceproviderconfig][%d] getScimV2ServiceproviderconfigTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetScimV2ServiceproviderconfigTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2ServiceproviderconfigTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2ServiceproviderconfigInternalServerError creates a GetScimV2ServiceproviderconfigInternalServerError with default headers values
func NewGetScimV2ServiceproviderconfigInternalServerError() *GetScimV2ServiceproviderconfigInternalServerError {
	return &GetScimV2ServiceproviderconfigInternalServerError{}
}

/*
GetScimV2ServiceproviderconfigInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetScimV2ServiceproviderconfigInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scim v2 serviceproviderconfig internal server error response has a 2xx status code
func (o *GetScimV2ServiceproviderconfigInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scim v2 serviceproviderconfig internal server error response has a 3xx status code
func (o *GetScimV2ServiceproviderconfigInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scim v2 serviceproviderconfig internal server error response has a 4xx status code
func (o *GetScimV2ServiceproviderconfigInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get scim v2 serviceproviderconfig internal server error response has a 5xx status code
func (o *GetScimV2ServiceproviderconfigInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get scim v2 serviceproviderconfig internal server error response a status code equal to that given
func (o *GetScimV2ServiceproviderconfigInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetScimV2ServiceproviderconfigInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/serviceproviderconfig][%d] getScimV2ServiceproviderconfigInternalServerError  %+v", 500, o.Payload)
}

func (o *GetScimV2ServiceproviderconfigInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/serviceproviderconfig][%d] getScimV2ServiceproviderconfigInternalServerError  %+v", 500, o.Payload)
}

func (o *GetScimV2ServiceproviderconfigInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2ServiceproviderconfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2ServiceproviderconfigServiceUnavailable creates a GetScimV2ServiceproviderconfigServiceUnavailable with default headers values
func NewGetScimV2ServiceproviderconfigServiceUnavailable() *GetScimV2ServiceproviderconfigServiceUnavailable {
	return &GetScimV2ServiceproviderconfigServiceUnavailable{}
}

/*
GetScimV2ServiceproviderconfigServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetScimV2ServiceproviderconfigServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scim v2 serviceproviderconfig service unavailable response has a 2xx status code
func (o *GetScimV2ServiceproviderconfigServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scim v2 serviceproviderconfig service unavailable response has a 3xx status code
func (o *GetScimV2ServiceproviderconfigServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scim v2 serviceproviderconfig service unavailable response has a 4xx status code
func (o *GetScimV2ServiceproviderconfigServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get scim v2 serviceproviderconfig service unavailable response has a 5xx status code
func (o *GetScimV2ServiceproviderconfigServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get scim v2 serviceproviderconfig service unavailable response a status code equal to that given
func (o *GetScimV2ServiceproviderconfigServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetScimV2ServiceproviderconfigServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/serviceproviderconfig][%d] getScimV2ServiceproviderconfigServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetScimV2ServiceproviderconfigServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/serviceproviderconfig][%d] getScimV2ServiceproviderconfigServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetScimV2ServiceproviderconfigServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2ServiceproviderconfigServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2ServiceproviderconfigGatewayTimeout creates a GetScimV2ServiceproviderconfigGatewayTimeout with default headers values
func NewGetScimV2ServiceproviderconfigGatewayTimeout() *GetScimV2ServiceproviderconfigGatewayTimeout {
	return &GetScimV2ServiceproviderconfigGatewayTimeout{}
}

/*
GetScimV2ServiceproviderconfigGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetScimV2ServiceproviderconfigGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scim v2 serviceproviderconfig gateway timeout response has a 2xx status code
func (o *GetScimV2ServiceproviderconfigGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scim v2 serviceproviderconfig gateway timeout response has a 3xx status code
func (o *GetScimV2ServiceproviderconfigGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scim v2 serviceproviderconfig gateway timeout response has a 4xx status code
func (o *GetScimV2ServiceproviderconfigGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get scim v2 serviceproviderconfig gateway timeout response has a 5xx status code
func (o *GetScimV2ServiceproviderconfigGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get scim v2 serviceproviderconfig gateway timeout response a status code equal to that given
func (o *GetScimV2ServiceproviderconfigGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetScimV2ServiceproviderconfigGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/serviceproviderconfig][%d] getScimV2ServiceproviderconfigGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetScimV2ServiceproviderconfigGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/serviceproviderconfig][%d] getScimV2ServiceproviderconfigGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetScimV2ServiceproviderconfigGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2ServiceproviderconfigGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package outbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetOutboundCallabletimesetReader is a Reader for the GetOutboundCallabletimeset structure.
type GetOutboundCallabletimesetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutboundCallabletimesetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutboundCallabletimesetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutboundCallabletimesetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutboundCallabletimesetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutboundCallabletimesetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutboundCallabletimesetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOutboundCallabletimesetRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOutboundCallabletimesetRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOutboundCallabletimesetUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOutboundCallabletimesetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOutboundCallabletimesetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOutboundCallabletimesetServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOutboundCallabletimesetGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOutboundCallabletimesetOK creates a GetOutboundCallabletimesetOK with default headers values
func NewGetOutboundCallabletimesetOK() *GetOutboundCallabletimesetOK {
	return &GetOutboundCallabletimesetOK{}
}

/*
GetOutboundCallabletimesetOK describes a response with status code 200, with default header values.

successful operation
*/
type GetOutboundCallabletimesetOK struct {
	Payload *models.CallableTimeSet
}

// IsSuccess returns true when this get outbound callabletimeset o k response has a 2xx status code
func (o *GetOutboundCallabletimesetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get outbound callabletimeset o k response has a 3xx status code
func (o *GetOutboundCallabletimesetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound callabletimeset o k response has a 4xx status code
func (o *GetOutboundCallabletimesetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound callabletimeset o k response has a 5xx status code
func (o *GetOutboundCallabletimesetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound callabletimeset o k response a status code equal to that given
func (o *GetOutboundCallabletimesetOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOutboundCallabletimesetOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callabletimesets/{callableTimeSetId}][%d] getOutboundCallabletimesetOK  %+v", 200, o.Payload)
}

func (o *GetOutboundCallabletimesetOK) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callabletimesets/{callableTimeSetId}][%d] getOutboundCallabletimesetOK  %+v", 200, o.Payload)
}

func (o *GetOutboundCallabletimesetOK) GetPayload() *models.CallableTimeSet {
	return o.Payload
}

func (o *GetOutboundCallabletimesetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CallableTimeSet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCallabletimesetBadRequest creates a GetOutboundCallabletimesetBadRequest with default headers values
func NewGetOutboundCallabletimesetBadRequest() *GetOutboundCallabletimesetBadRequest {
	return &GetOutboundCallabletimesetBadRequest{}
}

/*
GetOutboundCallabletimesetBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundCallabletimesetBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound callabletimeset bad request response has a 2xx status code
func (o *GetOutboundCallabletimesetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound callabletimeset bad request response has a 3xx status code
func (o *GetOutboundCallabletimesetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound callabletimeset bad request response has a 4xx status code
func (o *GetOutboundCallabletimesetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound callabletimeset bad request response has a 5xx status code
func (o *GetOutboundCallabletimesetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound callabletimeset bad request response a status code equal to that given
func (o *GetOutboundCallabletimesetBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetOutboundCallabletimesetBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callabletimesets/{callableTimeSetId}][%d] getOutboundCallabletimesetBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundCallabletimesetBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callabletimesets/{callableTimeSetId}][%d] getOutboundCallabletimesetBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundCallabletimesetBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCallabletimesetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCallabletimesetUnauthorized creates a GetOutboundCallabletimesetUnauthorized with default headers values
func NewGetOutboundCallabletimesetUnauthorized() *GetOutboundCallabletimesetUnauthorized {
	return &GetOutboundCallabletimesetUnauthorized{}
}

/*
GetOutboundCallabletimesetUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundCallabletimesetUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound callabletimeset unauthorized response has a 2xx status code
func (o *GetOutboundCallabletimesetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound callabletimeset unauthorized response has a 3xx status code
func (o *GetOutboundCallabletimesetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound callabletimeset unauthorized response has a 4xx status code
func (o *GetOutboundCallabletimesetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound callabletimeset unauthorized response has a 5xx status code
func (o *GetOutboundCallabletimesetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound callabletimeset unauthorized response a status code equal to that given
func (o *GetOutboundCallabletimesetUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetOutboundCallabletimesetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callabletimesets/{callableTimeSetId}][%d] getOutboundCallabletimesetUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundCallabletimesetUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callabletimesets/{callableTimeSetId}][%d] getOutboundCallabletimesetUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundCallabletimesetUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCallabletimesetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCallabletimesetForbidden creates a GetOutboundCallabletimesetForbidden with default headers values
func NewGetOutboundCallabletimesetForbidden() *GetOutboundCallabletimesetForbidden {
	return &GetOutboundCallabletimesetForbidden{}
}

/*
GetOutboundCallabletimesetForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundCallabletimesetForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound callabletimeset forbidden response has a 2xx status code
func (o *GetOutboundCallabletimesetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound callabletimeset forbidden response has a 3xx status code
func (o *GetOutboundCallabletimesetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound callabletimeset forbidden response has a 4xx status code
func (o *GetOutboundCallabletimesetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound callabletimeset forbidden response has a 5xx status code
func (o *GetOutboundCallabletimesetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound callabletimeset forbidden response a status code equal to that given
func (o *GetOutboundCallabletimesetForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetOutboundCallabletimesetForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callabletimesets/{callableTimeSetId}][%d] getOutboundCallabletimesetForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundCallabletimesetForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callabletimesets/{callableTimeSetId}][%d] getOutboundCallabletimesetForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundCallabletimesetForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCallabletimesetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCallabletimesetNotFound creates a GetOutboundCallabletimesetNotFound with default headers values
func NewGetOutboundCallabletimesetNotFound() *GetOutboundCallabletimesetNotFound {
	return &GetOutboundCallabletimesetNotFound{}
}

/*
GetOutboundCallabletimesetNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetOutboundCallabletimesetNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound callabletimeset not found response has a 2xx status code
func (o *GetOutboundCallabletimesetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound callabletimeset not found response has a 3xx status code
func (o *GetOutboundCallabletimesetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound callabletimeset not found response has a 4xx status code
func (o *GetOutboundCallabletimesetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound callabletimeset not found response has a 5xx status code
func (o *GetOutboundCallabletimesetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound callabletimeset not found response a status code equal to that given
func (o *GetOutboundCallabletimesetNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetOutboundCallabletimesetNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callabletimesets/{callableTimeSetId}][%d] getOutboundCallabletimesetNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundCallabletimesetNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callabletimesets/{callableTimeSetId}][%d] getOutboundCallabletimesetNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundCallabletimesetNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCallabletimesetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCallabletimesetRequestTimeout creates a GetOutboundCallabletimesetRequestTimeout with default headers values
func NewGetOutboundCallabletimesetRequestTimeout() *GetOutboundCallabletimesetRequestTimeout {
	return &GetOutboundCallabletimesetRequestTimeout{}
}

/*
GetOutboundCallabletimesetRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetOutboundCallabletimesetRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound callabletimeset request timeout response has a 2xx status code
func (o *GetOutboundCallabletimesetRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound callabletimeset request timeout response has a 3xx status code
func (o *GetOutboundCallabletimesetRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound callabletimeset request timeout response has a 4xx status code
func (o *GetOutboundCallabletimesetRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound callabletimeset request timeout response has a 5xx status code
func (o *GetOutboundCallabletimesetRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound callabletimeset request timeout response a status code equal to that given
func (o *GetOutboundCallabletimesetRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetOutboundCallabletimesetRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callabletimesets/{callableTimeSetId}][%d] getOutboundCallabletimesetRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundCallabletimesetRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callabletimesets/{callableTimeSetId}][%d] getOutboundCallabletimesetRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundCallabletimesetRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCallabletimesetRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCallabletimesetRequestEntityTooLarge creates a GetOutboundCallabletimesetRequestEntityTooLarge with default headers values
func NewGetOutboundCallabletimesetRequestEntityTooLarge() *GetOutboundCallabletimesetRequestEntityTooLarge {
	return &GetOutboundCallabletimesetRequestEntityTooLarge{}
}

/*
GetOutboundCallabletimesetRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetOutboundCallabletimesetRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound callabletimeset request entity too large response has a 2xx status code
func (o *GetOutboundCallabletimesetRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound callabletimeset request entity too large response has a 3xx status code
func (o *GetOutboundCallabletimesetRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound callabletimeset request entity too large response has a 4xx status code
func (o *GetOutboundCallabletimesetRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound callabletimeset request entity too large response has a 5xx status code
func (o *GetOutboundCallabletimesetRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound callabletimeset request entity too large response a status code equal to that given
func (o *GetOutboundCallabletimesetRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetOutboundCallabletimesetRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callabletimesets/{callableTimeSetId}][%d] getOutboundCallabletimesetRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundCallabletimesetRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callabletimesets/{callableTimeSetId}][%d] getOutboundCallabletimesetRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundCallabletimesetRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCallabletimesetRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCallabletimesetUnsupportedMediaType creates a GetOutboundCallabletimesetUnsupportedMediaType with default headers values
func NewGetOutboundCallabletimesetUnsupportedMediaType() *GetOutboundCallabletimesetUnsupportedMediaType {
	return &GetOutboundCallabletimesetUnsupportedMediaType{}
}

/*
GetOutboundCallabletimesetUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundCallabletimesetUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound callabletimeset unsupported media type response has a 2xx status code
func (o *GetOutboundCallabletimesetUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound callabletimeset unsupported media type response has a 3xx status code
func (o *GetOutboundCallabletimesetUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound callabletimeset unsupported media type response has a 4xx status code
func (o *GetOutboundCallabletimesetUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound callabletimeset unsupported media type response has a 5xx status code
func (o *GetOutboundCallabletimesetUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound callabletimeset unsupported media type response a status code equal to that given
func (o *GetOutboundCallabletimesetUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetOutboundCallabletimesetUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callabletimesets/{callableTimeSetId}][%d] getOutboundCallabletimesetUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundCallabletimesetUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callabletimesets/{callableTimeSetId}][%d] getOutboundCallabletimesetUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundCallabletimesetUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCallabletimesetUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCallabletimesetTooManyRequests creates a GetOutboundCallabletimesetTooManyRequests with default headers values
func NewGetOutboundCallabletimesetTooManyRequests() *GetOutboundCallabletimesetTooManyRequests {
	return &GetOutboundCallabletimesetTooManyRequests{}
}

/*
GetOutboundCallabletimesetTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetOutboundCallabletimesetTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound callabletimeset too many requests response has a 2xx status code
func (o *GetOutboundCallabletimesetTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound callabletimeset too many requests response has a 3xx status code
func (o *GetOutboundCallabletimesetTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound callabletimeset too many requests response has a 4xx status code
func (o *GetOutboundCallabletimesetTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound callabletimeset too many requests response has a 5xx status code
func (o *GetOutboundCallabletimesetTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound callabletimeset too many requests response a status code equal to that given
func (o *GetOutboundCallabletimesetTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetOutboundCallabletimesetTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callabletimesets/{callableTimeSetId}][%d] getOutboundCallabletimesetTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundCallabletimesetTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callabletimesets/{callableTimeSetId}][%d] getOutboundCallabletimesetTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundCallabletimesetTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCallabletimesetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCallabletimesetInternalServerError creates a GetOutboundCallabletimesetInternalServerError with default headers values
func NewGetOutboundCallabletimesetInternalServerError() *GetOutboundCallabletimesetInternalServerError {
	return &GetOutboundCallabletimesetInternalServerError{}
}

/*
GetOutboundCallabletimesetInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundCallabletimesetInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound callabletimeset internal server error response has a 2xx status code
func (o *GetOutboundCallabletimesetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound callabletimeset internal server error response has a 3xx status code
func (o *GetOutboundCallabletimesetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound callabletimeset internal server error response has a 4xx status code
func (o *GetOutboundCallabletimesetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound callabletimeset internal server error response has a 5xx status code
func (o *GetOutboundCallabletimesetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound callabletimeset internal server error response a status code equal to that given
func (o *GetOutboundCallabletimesetInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetOutboundCallabletimesetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callabletimesets/{callableTimeSetId}][%d] getOutboundCallabletimesetInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundCallabletimesetInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callabletimesets/{callableTimeSetId}][%d] getOutboundCallabletimesetInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundCallabletimesetInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCallabletimesetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCallabletimesetServiceUnavailable creates a GetOutboundCallabletimesetServiceUnavailable with default headers values
func NewGetOutboundCallabletimesetServiceUnavailable() *GetOutboundCallabletimesetServiceUnavailable {
	return &GetOutboundCallabletimesetServiceUnavailable{}
}

/*
GetOutboundCallabletimesetServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundCallabletimesetServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound callabletimeset service unavailable response has a 2xx status code
func (o *GetOutboundCallabletimesetServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound callabletimeset service unavailable response has a 3xx status code
func (o *GetOutboundCallabletimesetServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound callabletimeset service unavailable response has a 4xx status code
func (o *GetOutboundCallabletimesetServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound callabletimeset service unavailable response has a 5xx status code
func (o *GetOutboundCallabletimesetServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound callabletimeset service unavailable response a status code equal to that given
func (o *GetOutboundCallabletimesetServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetOutboundCallabletimesetServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callabletimesets/{callableTimeSetId}][%d] getOutboundCallabletimesetServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundCallabletimesetServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callabletimesets/{callableTimeSetId}][%d] getOutboundCallabletimesetServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundCallabletimesetServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCallabletimesetServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCallabletimesetGatewayTimeout creates a GetOutboundCallabletimesetGatewayTimeout with default headers values
func NewGetOutboundCallabletimesetGatewayTimeout() *GetOutboundCallabletimesetGatewayTimeout {
	return &GetOutboundCallabletimesetGatewayTimeout{}
}

/*
GetOutboundCallabletimesetGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetOutboundCallabletimesetGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound callabletimeset gateway timeout response has a 2xx status code
func (o *GetOutboundCallabletimesetGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound callabletimeset gateway timeout response has a 3xx status code
func (o *GetOutboundCallabletimesetGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound callabletimeset gateway timeout response has a 4xx status code
func (o *GetOutboundCallabletimesetGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound callabletimeset gateway timeout response has a 5xx status code
func (o *GetOutboundCallabletimesetGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound callabletimeset gateway timeout response a status code equal to that given
func (o *GetOutboundCallabletimesetGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetOutboundCallabletimesetGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callabletimesets/{callableTimeSetId}][%d] getOutboundCallabletimesetGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundCallabletimesetGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callabletimesets/{callableTimeSetId}][%d] getOutboundCallabletimesetGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundCallabletimesetGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCallabletimesetGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

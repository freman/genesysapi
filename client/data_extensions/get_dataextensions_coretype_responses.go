// Code generated by go-swagger; DO NOT EDIT.

package data_extensions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetDataextensionsCoretypeReader is a Reader for the GetDataextensionsCoretype structure.
type GetDataextensionsCoretypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDataextensionsCoretypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDataextensionsCoretypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDataextensionsCoretypeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDataextensionsCoretypeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDataextensionsCoretypeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDataextensionsCoretypeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetDataextensionsCoretypeRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetDataextensionsCoretypeRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetDataextensionsCoretypeUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetDataextensionsCoretypeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDataextensionsCoretypeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetDataextensionsCoretypeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetDataextensionsCoretypeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDataextensionsCoretypeOK creates a GetDataextensionsCoretypeOK with default headers values
func NewGetDataextensionsCoretypeOK() *GetDataextensionsCoretypeOK {
	return &GetDataextensionsCoretypeOK{}
}

/*
GetDataextensionsCoretypeOK describes a response with status code 200, with default header values.

successful operation
*/
type GetDataextensionsCoretypeOK struct {
	Payload *models.Coretype
}

// IsSuccess returns true when this get dataextensions coretype o k response has a 2xx status code
func (o *GetDataextensionsCoretypeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get dataextensions coretype o k response has a 3xx status code
func (o *GetDataextensionsCoretypeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dataextensions coretype o k response has a 4xx status code
func (o *GetDataextensionsCoretypeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get dataextensions coretype o k response has a 5xx status code
func (o *GetDataextensionsCoretypeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get dataextensions coretype o k response a status code equal to that given
func (o *GetDataextensionsCoretypeOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetDataextensionsCoretypeOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/dataextensions/coretypes/{coretypeName}][%d] getDataextensionsCoretypeOK  %+v", 200, o.Payload)
}

func (o *GetDataextensionsCoretypeOK) String() string {
	return fmt.Sprintf("[GET /api/v2/dataextensions/coretypes/{coretypeName}][%d] getDataextensionsCoretypeOK  %+v", 200, o.Payload)
}

func (o *GetDataextensionsCoretypeOK) GetPayload() *models.Coretype {
	return o.Payload
}

func (o *GetDataextensionsCoretypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Coretype)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataextensionsCoretypeBadRequest creates a GetDataextensionsCoretypeBadRequest with default headers values
func NewGetDataextensionsCoretypeBadRequest() *GetDataextensionsCoretypeBadRequest {
	return &GetDataextensionsCoretypeBadRequest{}
}

/*
GetDataextensionsCoretypeBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetDataextensionsCoretypeBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get dataextensions coretype bad request response has a 2xx status code
func (o *GetDataextensionsCoretypeBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get dataextensions coretype bad request response has a 3xx status code
func (o *GetDataextensionsCoretypeBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dataextensions coretype bad request response has a 4xx status code
func (o *GetDataextensionsCoretypeBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get dataextensions coretype bad request response has a 5xx status code
func (o *GetDataextensionsCoretypeBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get dataextensions coretype bad request response a status code equal to that given
func (o *GetDataextensionsCoretypeBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetDataextensionsCoretypeBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/dataextensions/coretypes/{coretypeName}][%d] getDataextensionsCoretypeBadRequest  %+v", 400, o.Payload)
}

func (o *GetDataextensionsCoretypeBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/dataextensions/coretypes/{coretypeName}][%d] getDataextensionsCoretypeBadRequest  %+v", 400, o.Payload)
}

func (o *GetDataextensionsCoretypeBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDataextensionsCoretypeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataextensionsCoretypeUnauthorized creates a GetDataextensionsCoretypeUnauthorized with default headers values
func NewGetDataextensionsCoretypeUnauthorized() *GetDataextensionsCoretypeUnauthorized {
	return &GetDataextensionsCoretypeUnauthorized{}
}

/*
GetDataextensionsCoretypeUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetDataextensionsCoretypeUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get dataextensions coretype unauthorized response has a 2xx status code
func (o *GetDataextensionsCoretypeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get dataextensions coretype unauthorized response has a 3xx status code
func (o *GetDataextensionsCoretypeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dataextensions coretype unauthorized response has a 4xx status code
func (o *GetDataextensionsCoretypeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get dataextensions coretype unauthorized response has a 5xx status code
func (o *GetDataextensionsCoretypeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get dataextensions coretype unauthorized response a status code equal to that given
func (o *GetDataextensionsCoretypeUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetDataextensionsCoretypeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/dataextensions/coretypes/{coretypeName}][%d] getDataextensionsCoretypeUnauthorized  %+v", 401, o.Payload)
}

func (o *GetDataextensionsCoretypeUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/dataextensions/coretypes/{coretypeName}][%d] getDataextensionsCoretypeUnauthorized  %+v", 401, o.Payload)
}

func (o *GetDataextensionsCoretypeUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDataextensionsCoretypeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataextensionsCoretypeForbidden creates a GetDataextensionsCoretypeForbidden with default headers values
func NewGetDataextensionsCoretypeForbidden() *GetDataextensionsCoretypeForbidden {
	return &GetDataextensionsCoretypeForbidden{}
}

/*
GetDataextensionsCoretypeForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetDataextensionsCoretypeForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get dataextensions coretype forbidden response has a 2xx status code
func (o *GetDataextensionsCoretypeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get dataextensions coretype forbidden response has a 3xx status code
func (o *GetDataextensionsCoretypeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dataextensions coretype forbidden response has a 4xx status code
func (o *GetDataextensionsCoretypeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get dataextensions coretype forbidden response has a 5xx status code
func (o *GetDataextensionsCoretypeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get dataextensions coretype forbidden response a status code equal to that given
func (o *GetDataextensionsCoretypeForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetDataextensionsCoretypeForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/dataextensions/coretypes/{coretypeName}][%d] getDataextensionsCoretypeForbidden  %+v", 403, o.Payload)
}

func (o *GetDataextensionsCoretypeForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/dataextensions/coretypes/{coretypeName}][%d] getDataextensionsCoretypeForbidden  %+v", 403, o.Payload)
}

func (o *GetDataextensionsCoretypeForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDataextensionsCoretypeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataextensionsCoretypeNotFound creates a GetDataextensionsCoretypeNotFound with default headers values
func NewGetDataextensionsCoretypeNotFound() *GetDataextensionsCoretypeNotFound {
	return &GetDataextensionsCoretypeNotFound{}
}

/*
GetDataextensionsCoretypeNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetDataextensionsCoretypeNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get dataextensions coretype not found response has a 2xx status code
func (o *GetDataextensionsCoretypeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get dataextensions coretype not found response has a 3xx status code
func (o *GetDataextensionsCoretypeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dataextensions coretype not found response has a 4xx status code
func (o *GetDataextensionsCoretypeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get dataextensions coretype not found response has a 5xx status code
func (o *GetDataextensionsCoretypeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get dataextensions coretype not found response a status code equal to that given
func (o *GetDataextensionsCoretypeNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetDataextensionsCoretypeNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/dataextensions/coretypes/{coretypeName}][%d] getDataextensionsCoretypeNotFound  %+v", 404, o.Payload)
}

func (o *GetDataextensionsCoretypeNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/dataextensions/coretypes/{coretypeName}][%d] getDataextensionsCoretypeNotFound  %+v", 404, o.Payload)
}

func (o *GetDataextensionsCoretypeNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDataextensionsCoretypeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataextensionsCoretypeRequestTimeout creates a GetDataextensionsCoretypeRequestTimeout with default headers values
func NewGetDataextensionsCoretypeRequestTimeout() *GetDataextensionsCoretypeRequestTimeout {
	return &GetDataextensionsCoretypeRequestTimeout{}
}

/*
GetDataextensionsCoretypeRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetDataextensionsCoretypeRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get dataextensions coretype request timeout response has a 2xx status code
func (o *GetDataextensionsCoretypeRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get dataextensions coretype request timeout response has a 3xx status code
func (o *GetDataextensionsCoretypeRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dataextensions coretype request timeout response has a 4xx status code
func (o *GetDataextensionsCoretypeRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get dataextensions coretype request timeout response has a 5xx status code
func (o *GetDataextensionsCoretypeRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get dataextensions coretype request timeout response a status code equal to that given
func (o *GetDataextensionsCoretypeRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetDataextensionsCoretypeRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/dataextensions/coretypes/{coretypeName}][%d] getDataextensionsCoretypeRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetDataextensionsCoretypeRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/dataextensions/coretypes/{coretypeName}][%d] getDataextensionsCoretypeRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetDataextensionsCoretypeRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDataextensionsCoretypeRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataextensionsCoretypeRequestEntityTooLarge creates a GetDataextensionsCoretypeRequestEntityTooLarge with default headers values
func NewGetDataextensionsCoretypeRequestEntityTooLarge() *GetDataextensionsCoretypeRequestEntityTooLarge {
	return &GetDataextensionsCoretypeRequestEntityTooLarge{}
}

/*
GetDataextensionsCoretypeRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetDataextensionsCoretypeRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get dataextensions coretype request entity too large response has a 2xx status code
func (o *GetDataextensionsCoretypeRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get dataextensions coretype request entity too large response has a 3xx status code
func (o *GetDataextensionsCoretypeRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dataextensions coretype request entity too large response has a 4xx status code
func (o *GetDataextensionsCoretypeRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get dataextensions coretype request entity too large response has a 5xx status code
func (o *GetDataextensionsCoretypeRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get dataextensions coretype request entity too large response a status code equal to that given
func (o *GetDataextensionsCoretypeRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetDataextensionsCoretypeRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/dataextensions/coretypes/{coretypeName}][%d] getDataextensionsCoretypeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetDataextensionsCoretypeRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/dataextensions/coretypes/{coretypeName}][%d] getDataextensionsCoretypeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetDataextensionsCoretypeRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDataextensionsCoretypeRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataextensionsCoretypeUnsupportedMediaType creates a GetDataextensionsCoretypeUnsupportedMediaType with default headers values
func NewGetDataextensionsCoretypeUnsupportedMediaType() *GetDataextensionsCoretypeUnsupportedMediaType {
	return &GetDataextensionsCoretypeUnsupportedMediaType{}
}

/*
GetDataextensionsCoretypeUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetDataextensionsCoretypeUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get dataextensions coretype unsupported media type response has a 2xx status code
func (o *GetDataextensionsCoretypeUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get dataextensions coretype unsupported media type response has a 3xx status code
func (o *GetDataextensionsCoretypeUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dataextensions coretype unsupported media type response has a 4xx status code
func (o *GetDataextensionsCoretypeUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get dataextensions coretype unsupported media type response has a 5xx status code
func (o *GetDataextensionsCoretypeUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get dataextensions coretype unsupported media type response a status code equal to that given
func (o *GetDataextensionsCoretypeUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetDataextensionsCoretypeUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/dataextensions/coretypes/{coretypeName}][%d] getDataextensionsCoretypeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetDataextensionsCoretypeUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/dataextensions/coretypes/{coretypeName}][%d] getDataextensionsCoretypeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetDataextensionsCoretypeUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDataextensionsCoretypeUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataextensionsCoretypeTooManyRequests creates a GetDataextensionsCoretypeTooManyRequests with default headers values
func NewGetDataextensionsCoretypeTooManyRequests() *GetDataextensionsCoretypeTooManyRequests {
	return &GetDataextensionsCoretypeTooManyRequests{}
}

/*
GetDataextensionsCoretypeTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetDataextensionsCoretypeTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get dataextensions coretype too many requests response has a 2xx status code
func (o *GetDataextensionsCoretypeTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get dataextensions coretype too many requests response has a 3xx status code
func (o *GetDataextensionsCoretypeTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dataextensions coretype too many requests response has a 4xx status code
func (o *GetDataextensionsCoretypeTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get dataextensions coretype too many requests response has a 5xx status code
func (o *GetDataextensionsCoretypeTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get dataextensions coretype too many requests response a status code equal to that given
func (o *GetDataextensionsCoretypeTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetDataextensionsCoretypeTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/dataextensions/coretypes/{coretypeName}][%d] getDataextensionsCoretypeTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetDataextensionsCoretypeTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/dataextensions/coretypes/{coretypeName}][%d] getDataextensionsCoretypeTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetDataextensionsCoretypeTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDataextensionsCoretypeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataextensionsCoretypeInternalServerError creates a GetDataextensionsCoretypeInternalServerError with default headers values
func NewGetDataextensionsCoretypeInternalServerError() *GetDataextensionsCoretypeInternalServerError {
	return &GetDataextensionsCoretypeInternalServerError{}
}

/*
GetDataextensionsCoretypeInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetDataextensionsCoretypeInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get dataextensions coretype internal server error response has a 2xx status code
func (o *GetDataextensionsCoretypeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get dataextensions coretype internal server error response has a 3xx status code
func (o *GetDataextensionsCoretypeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dataextensions coretype internal server error response has a 4xx status code
func (o *GetDataextensionsCoretypeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get dataextensions coretype internal server error response has a 5xx status code
func (o *GetDataextensionsCoretypeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get dataextensions coretype internal server error response a status code equal to that given
func (o *GetDataextensionsCoretypeInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetDataextensionsCoretypeInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/dataextensions/coretypes/{coretypeName}][%d] getDataextensionsCoretypeInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDataextensionsCoretypeInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/dataextensions/coretypes/{coretypeName}][%d] getDataextensionsCoretypeInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDataextensionsCoretypeInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDataextensionsCoretypeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataextensionsCoretypeServiceUnavailable creates a GetDataextensionsCoretypeServiceUnavailable with default headers values
func NewGetDataextensionsCoretypeServiceUnavailable() *GetDataextensionsCoretypeServiceUnavailable {
	return &GetDataextensionsCoretypeServiceUnavailable{}
}

/*
GetDataextensionsCoretypeServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetDataextensionsCoretypeServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get dataextensions coretype service unavailable response has a 2xx status code
func (o *GetDataextensionsCoretypeServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get dataextensions coretype service unavailable response has a 3xx status code
func (o *GetDataextensionsCoretypeServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dataextensions coretype service unavailable response has a 4xx status code
func (o *GetDataextensionsCoretypeServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get dataextensions coretype service unavailable response has a 5xx status code
func (o *GetDataextensionsCoretypeServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get dataextensions coretype service unavailable response a status code equal to that given
func (o *GetDataextensionsCoretypeServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetDataextensionsCoretypeServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/dataextensions/coretypes/{coretypeName}][%d] getDataextensionsCoretypeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetDataextensionsCoretypeServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/dataextensions/coretypes/{coretypeName}][%d] getDataextensionsCoretypeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetDataextensionsCoretypeServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDataextensionsCoretypeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataextensionsCoretypeGatewayTimeout creates a GetDataextensionsCoretypeGatewayTimeout with default headers values
func NewGetDataextensionsCoretypeGatewayTimeout() *GetDataextensionsCoretypeGatewayTimeout {
	return &GetDataextensionsCoretypeGatewayTimeout{}
}

/*
GetDataextensionsCoretypeGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetDataextensionsCoretypeGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get dataextensions coretype gateway timeout response has a 2xx status code
func (o *GetDataextensionsCoretypeGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get dataextensions coretype gateway timeout response has a 3xx status code
func (o *GetDataextensionsCoretypeGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dataextensions coretype gateway timeout response has a 4xx status code
func (o *GetDataextensionsCoretypeGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get dataextensions coretype gateway timeout response has a 5xx status code
func (o *GetDataextensionsCoretypeGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get dataextensions coretype gateway timeout response a status code equal to that given
func (o *GetDataextensionsCoretypeGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetDataextensionsCoretypeGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/dataextensions/coretypes/{coretypeName}][%d] getDataextensionsCoretypeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetDataextensionsCoretypeGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/dataextensions/coretypes/{coretypeName}][%d] getDataextensionsCoretypeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetDataextensionsCoretypeGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDataextensionsCoretypeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

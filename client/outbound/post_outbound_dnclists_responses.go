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

// PostOutboundDnclistsReader is a Reader for the PostOutboundDnclists structure.
type PostOutboundDnclistsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOutboundDnclistsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostOutboundDnclistsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostOutboundDnclistsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostOutboundDnclistsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostOutboundDnclistsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostOutboundDnclistsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostOutboundDnclistsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostOutboundDnclistsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostOutboundDnclistsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostOutboundDnclistsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostOutboundDnclistsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewPostOutboundDnclistsNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostOutboundDnclistsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostOutboundDnclistsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostOutboundDnclistsOK creates a PostOutboundDnclistsOK with default headers values
func NewPostOutboundDnclistsOK() *PostOutboundDnclistsOK {
	return &PostOutboundDnclistsOK{}
}

/*
PostOutboundDnclistsOK describes a response with status code 200, with default header values.

successful operation
*/
type PostOutboundDnclistsOK struct {
	Payload *models.DncList
}

// IsSuccess returns true when this post outbound dnclists o k response has a 2xx status code
func (o *PostOutboundDnclistsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post outbound dnclists o k response has a 3xx status code
func (o *PostOutboundDnclistsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound dnclists o k response has a 4xx status code
func (o *PostOutboundDnclistsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post outbound dnclists o k response has a 5xx status code
func (o *PostOutboundDnclistsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound dnclists o k response a status code equal to that given
func (o *PostOutboundDnclistsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostOutboundDnclistsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists][%d] postOutboundDnclistsOK  %+v", 200, o.Payload)
}

func (o *PostOutboundDnclistsOK) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists][%d] postOutboundDnclistsOK  %+v", 200, o.Payload)
}

func (o *PostOutboundDnclistsOK) GetPayload() *models.DncList {
	return o.Payload
}

func (o *PostOutboundDnclistsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DncList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundDnclistsBadRequest creates a PostOutboundDnclistsBadRequest with default headers values
func NewPostOutboundDnclistsBadRequest() *PostOutboundDnclistsBadRequest {
	return &PostOutboundDnclistsBadRequest{}
}

/*
PostOutboundDnclistsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostOutboundDnclistsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound dnclists bad request response has a 2xx status code
func (o *PostOutboundDnclistsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound dnclists bad request response has a 3xx status code
func (o *PostOutboundDnclistsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound dnclists bad request response has a 4xx status code
func (o *PostOutboundDnclistsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound dnclists bad request response has a 5xx status code
func (o *PostOutboundDnclistsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound dnclists bad request response a status code equal to that given
func (o *PostOutboundDnclistsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostOutboundDnclistsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists][%d] postOutboundDnclistsBadRequest  %+v", 400, o.Payload)
}

func (o *PostOutboundDnclistsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists][%d] postOutboundDnclistsBadRequest  %+v", 400, o.Payload)
}

func (o *PostOutboundDnclistsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundDnclistsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundDnclistsUnauthorized creates a PostOutboundDnclistsUnauthorized with default headers values
func NewPostOutboundDnclistsUnauthorized() *PostOutboundDnclistsUnauthorized {
	return &PostOutboundDnclistsUnauthorized{}
}

/*
PostOutboundDnclistsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostOutboundDnclistsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound dnclists unauthorized response has a 2xx status code
func (o *PostOutboundDnclistsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound dnclists unauthorized response has a 3xx status code
func (o *PostOutboundDnclistsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound dnclists unauthorized response has a 4xx status code
func (o *PostOutboundDnclistsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound dnclists unauthorized response has a 5xx status code
func (o *PostOutboundDnclistsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound dnclists unauthorized response a status code equal to that given
func (o *PostOutboundDnclistsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostOutboundDnclistsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists][%d] postOutboundDnclistsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostOutboundDnclistsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists][%d] postOutboundDnclistsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostOutboundDnclistsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundDnclistsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundDnclistsForbidden creates a PostOutboundDnclistsForbidden with default headers values
func NewPostOutboundDnclistsForbidden() *PostOutboundDnclistsForbidden {
	return &PostOutboundDnclistsForbidden{}
}

/*
PostOutboundDnclistsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostOutboundDnclistsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound dnclists forbidden response has a 2xx status code
func (o *PostOutboundDnclistsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound dnclists forbidden response has a 3xx status code
func (o *PostOutboundDnclistsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound dnclists forbidden response has a 4xx status code
func (o *PostOutboundDnclistsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound dnclists forbidden response has a 5xx status code
func (o *PostOutboundDnclistsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound dnclists forbidden response a status code equal to that given
func (o *PostOutboundDnclistsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostOutboundDnclistsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists][%d] postOutboundDnclistsForbidden  %+v", 403, o.Payload)
}

func (o *PostOutboundDnclistsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists][%d] postOutboundDnclistsForbidden  %+v", 403, o.Payload)
}

func (o *PostOutboundDnclistsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundDnclistsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundDnclistsNotFound creates a PostOutboundDnclistsNotFound with default headers values
func NewPostOutboundDnclistsNotFound() *PostOutboundDnclistsNotFound {
	return &PostOutboundDnclistsNotFound{}
}

/*
PostOutboundDnclistsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostOutboundDnclistsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound dnclists not found response has a 2xx status code
func (o *PostOutboundDnclistsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound dnclists not found response has a 3xx status code
func (o *PostOutboundDnclistsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound dnclists not found response has a 4xx status code
func (o *PostOutboundDnclistsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound dnclists not found response has a 5xx status code
func (o *PostOutboundDnclistsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound dnclists not found response a status code equal to that given
func (o *PostOutboundDnclistsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostOutboundDnclistsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists][%d] postOutboundDnclistsNotFound  %+v", 404, o.Payload)
}

func (o *PostOutboundDnclistsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists][%d] postOutboundDnclistsNotFound  %+v", 404, o.Payload)
}

func (o *PostOutboundDnclistsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundDnclistsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundDnclistsRequestTimeout creates a PostOutboundDnclistsRequestTimeout with default headers values
func NewPostOutboundDnclistsRequestTimeout() *PostOutboundDnclistsRequestTimeout {
	return &PostOutboundDnclistsRequestTimeout{}
}

/*
PostOutboundDnclistsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostOutboundDnclistsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound dnclists request timeout response has a 2xx status code
func (o *PostOutboundDnclistsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound dnclists request timeout response has a 3xx status code
func (o *PostOutboundDnclistsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound dnclists request timeout response has a 4xx status code
func (o *PostOutboundDnclistsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound dnclists request timeout response has a 5xx status code
func (o *PostOutboundDnclistsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound dnclists request timeout response a status code equal to that given
func (o *PostOutboundDnclistsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostOutboundDnclistsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists][%d] postOutboundDnclistsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostOutboundDnclistsRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists][%d] postOutboundDnclistsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostOutboundDnclistsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundDnclistsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundDnclistsRequestEntityTooLarge creates a PostOutboundDnclistsRequestEntityTooLarge with default headers values
func NewPostOutboundDnclistsRequestEntityTooLarge() *PostOutboundDnclistsRequestEntityTooLarge {
	return &PostOutboundDnclistsRequestEntityTooLarge{}
}

/*
PostOutboundDnclistsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostOutboundDnclistsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound dnclists request entity too large response has a 2xx status code
func (o *PostOutboundDnclistsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound dnclists request entity too large response has a 3xx status code
func (o *PostOutboundDnclistsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound dnclists request entity too large response has a 4xx status code
func (o *PostOutboundDnclistsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound dnclists request entity too large response has a 5xx status code
func (o *PostOutboundDnclistsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound dnclists request entity too large response a status code equal to that given
func (o *PostOutboundDnclistsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostOutboundDnclistsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists][%d] postOutboundDnclistsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostOutboundDnclistsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists][%d] postOutboundDnclistsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostOutboundDnclistsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundDnclistsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundDnclistsUnsupportedMediaType creates a PostOutboundDnclistsUnsupportedMediaType with default headers values
func NewPostOutboundDnclistsUnsupportedMediaType() *PostOutboundDnclistsUnsupportedMediaType {
	return &PostOutboundDnclistsUnsupportedMediaType{}
}

/*
PostOutboundDnclistsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostOutboundDnclistsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound dnclists unsupported media type response has a 2xx status code
func (o *PostOutboundDnclistsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound dnclists unsupported media type response has a 3xx status code
func (o *PostOutboundDnclistsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound dnclists unsupported media type response has a 4xx status code
func (o *PostOutboundDnclistsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound dnclists unsupported media type response has a 5xx status code
func (o *PostOutboundDnclistsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound dnclists unsupported media type response a status code equal to that given
func (o *PostOutboundDnclistsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostOutboundDnclistsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists][%d] postOutboundDnclistsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostOutboundDnclistsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists][%d] postOutboundDnclistsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostOutboundDnclistsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundDnclistsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundDnclistsTooManyRequests creates a PostOutboundDnclistsTooManyRequests with default headers values
func NewPostOutboundDnclistsTooManyRequests() *PostOutboundDnclistsTooManyRequests {
	return &PostOutboundDnclistsTooManyRequests{}
}

/*
PostOutboundDnclistsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostOutboundDnclistsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound dnclists too many requests response has a 2xx status code
func (o *PostOutboundDnclistsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound dnclists too many requests response has a 3xx status code
func (o *PostOutboundDnclistsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound dnclists too many requests response has a 4xx status code
func (o *PostOutboundDnclistsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound dnclists too many requests response has a 5xx status code
func (o *PostOutboundDnclistsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound dnclists too many requests response a status code equal to that given
func (o *PostOutboundDnclistsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostOutboundDnclistsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists][%d] postOutboundDnclistsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostOutboundDnclistsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists][%d] postOutboundDnclistsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostOutboundDnclistsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundDnclistsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundDnclistsInternalServerError creates a PostOutboundDnclistsInternalServerError with default headers values
func NewPostOutboundDnclistsInternalServerError() *PostOutboundDnclistsInternalServerError {
	return &PostOutboundDnclistsInternalServerError{}
}

/*
PostOutboundDnclistsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostOutboundDnclistsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound dnclists internal server error response has a 2xx status code
func (o *PostOutboundDnclistsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound dnclists internal server error response has a 3xx status code
func (o *PostOutboundDnclistsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound dnclists internal server error response has a 4xx status code
func (o *PostOutboundDnclistsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post outbound dnclists internal server error response has a 5xx status code
func (o *PostOutboundDnclistsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post outbound dnclists internal server error response a status code equal to that given
func (o *PostOutboundDnclistsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostOutboundDnclistsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists][%d] postOutboundDnclistsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostOutboundDnclistsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists][%d] postOutboundDnclistsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostOutboundDnclistsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundDnclistsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundDnclistsNotImplemented creates a PostOutboundDnclistsNotImplemented with default headers values
func NewPostOutboundDnclistsNotImplemented() *PostOutboundDnclistsNotImplemented {
	return &PostOutboundDnclistsNotImplemented{}
}

/*
PostOutboundDnclistsNotImplemented describes a response with status code 501, with default header values.

Not Implemented
*/
type PostOutboundDnclistsNotImplemented struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound dnclists not implemented response has a 2xx status code
func (o *PostOutboundDnclistsNotImplemented) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound dnclists not implemented response has a 3xx status code
func (o *PostOutboundDnclistsNotImplemented) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound dnclists not implemented response has a 4xx status code
func (o *PostOutboundDnclistsNotImplemented) IsClientError() bool {
	return false
}

// IsServerError returns true when this post outbound dnclists not implemented response has a 5xx status code
func (o *PostOutboundDnclistsNotImplemented) IsServerError() bool {
	return true
}

// IsCode returns true when this post outbound dnclists not implemented response a status code equal to that given
func (o *PostOutboundDnclistsNotImplemented) IsCode(code int) bool {
	return code == 501
}

func (o *PostOutboundDnclistsNotImplemented) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists][%d] postOutboundDnclistsNotImplemented  %+v", 501, o.Payload)
}

func (o *PostOutboundDnclistsNotImplemented) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists][%d] postOutboundDnclistsNotImplemented  %+v", 501, o.Payload)
}

func (o *PostOutboundDnclistsNotImplemented) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundDnclistsNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundDnclistsServiceUnavailable creates a PostOutboundDnclistsServiceUnavailable with default headers values
func NewPostOutboundDnclistsServiceUnavailable() *PostOutboundDnclistsServiceUnavailable {
	return &PostOutboundDnclistsServiceUnavailable{}
}

/*
PostOutboundDnclistsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostOutboundDnclistsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound dnclists service unavailable response has a 2xx status code
func (o *PostOutboundDnclistsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound dnclists service unavailable response has a 3xx status code
func (o *PostOutboundDnclistsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound dnclists service unavailable response has a 4xx status code
func (o *PostOutboundDnclistsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post outbound dnclists service unavailable response has a 5xx status code
func (o *PostOutboundDnclistsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post outbound dnclists service unavailable response a status code equal to that given
func (o *PostOutboundDnclistsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostOutboundDnclistsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists][%d] postOutboundDnclistsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostOutboundDnclistsServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists][%d] postOutboundDnclistsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostOutboundDnclistsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundDnclistsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundDnclistsGatewayTimeout creates a PostOutboundDnclistsGatewayTimeout with default headers values
func NewPostOutboundDnclistsGatewayTimeout() *PostOutboundDnclistsGatewayTimeout {
	return &PostOutboundDnclistsGatewayTimeout{}
}

/*
PostOutboundDnclistsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostOutboundDnclistsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound dnclists gateway timeout response has a 2xx status code
func (o *PostOutboundDnclistsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound dnclists gateway timeout response has a 3xx status code
func (o *PostOutboundDnclistsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound dnclists gateway timeout response has a 4xx status code
func (o *PostOutboundDnclistsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post outbound dnclists gateway timeout response has a 5xx status code
func (o *PostOutboundDnclistsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post outbound dnclists gateway timeout response a status code equal to that given
func (o *PostOutboundDnclistsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostOutboundDnclistsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists][%d] postOutboundDnclistsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostOutboundDnclistsGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists][%d] postOutboundDnclistsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostOutboundDnclistsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundDnclistsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

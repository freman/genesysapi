// Code generated by go-swagger; DO NOT EDIT.

package journey

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetJourneyActiontargetReader is a Reader for the GetJourneyActiontarget structure.
type GetJourneyActiontargetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetJourneyActiontargetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetJourneyActiontargetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetJourneyActiontargetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetJourneyActiontargetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetJourneyActiontargetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetJourneyActiontargetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetJourneyActiontargetRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetJourneyActiontargetRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetJourneyActiontargetUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetJourneyActiontargetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetJourneyActiontargetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetJourneyActiontargetServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetJourneyActiontargetGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetJourneyActiontargetOK creates a GetJourneyActiontargetOK with default headers values
func NewGetJourneyActiontargetOK() *GetJourneyActiontargetOK {
	return &GetJourneyActiontargetOK{}
}

/*
GetJourneyActiontargetOK describes a response with status code 200, with default header values.

successful operation
*/
type GetJourneyActiontargetOK struct {
	Payload *models.ActionTarget
}

// IsSuccess returns true when this get journey actiontarget o k response has a 2xx status code
func (o *GetJourneyActiontargetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get journey actiontarget o k response has a 3xx status code
func (o *GetJourneyActiontargetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actiontarget o k response has a 4xx status code
func (o *GetJourneyActiontargetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get journey actiontarget o k response has a 5xx status code
func (o *GetJourneyActiontargetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey actiontarget o k response a status code equal to that given
func (o *GetJourneyActiontargetOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetJourneyActiontargetOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets/{actionTargetId}][%d] getJourneyActiontargetOK  %+v", 200, o.Payload)
}

func (o *GetJourneyActiontargetOK) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets/{actionTargetId}][%d] getJourneyActiontargetOK  %+v", 200, o.Payload)
}

func (o *GetJourneyActiontargetOK) GetPayload() *models.ActionTarget {
	return o.Payload
}

func (o *GetJourneyActiontargetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActionTarget)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActiontargetBadRequest creates a GetJourneyActiontargetBadRequest with default headers values
func NewGetJourneyActiontargetBadRequest() *GetJourneyActiontargetBadRequest {
	return &GetJourneyActiontargetBadRequest{}
}

/*
GetJourneyActiontargetBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetJourneyActiontargetBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actiontarget bad request response has a 2xx status code
func (o *GetJourneyActiontargetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actiontarget bad request response has a 3xx status code
func (o *GetJourneyActiontargetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actiontarget bad request response has a 4xx status code
func (o *GetJourneyActiontargetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey actiontarget bad request response has a 5xx status code
func (o *GetJourneyActiontargetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey actiontarget bad request response a status code equal to that given
func (o *GetJourneyActiontargetBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetJourneyActiontargetBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets/{actionTargetId}][%d] getJourneyActiontargetBadRequest  %+v", 400, o.Payload)
}

func (o *GetJourneyActiontargetBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets/{actionTargetId}][%d] getJourneyActiontargetBadRequest  %+v", 400, o.Payload)
}

func (o *GetJourneyActiontargetBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActiontargetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActiontargetUnauthorized creates a GetJourneyActiontargetUnauthorized with default headers values
func NewGetJourneyActiontargetUnauthorized() *GetJourneyActiontargetUnauthorized {
	return &GetJourneyActiontargetUnauthorized{}
}

/*
GetJourneyActiontargetUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetJourneyActiontargetUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actiontarget unauthorized response has a 2xx status code
func (o *GetJourneyActiontargetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actiontarget unauthorized response has a 3xx status code
func (o *GetJourneyActiontargetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actiontarget unauthorized response has a 4xx status code
func (o *GetJourneyActiontargetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey actiontarget unauthorized response has a 5xx status code
func (o *GetJourneyActiontargetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey actiontarget unauthorized response a status code equal to that given
func (o *GetJourneyActiontargetUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetJourneyActiontargetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets/{actionTargetId}][%d] getJourneyActiontargetUnauthorized  %+v", 401, o.Payload)
}

func (o *GetJourneyActiontargetUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets/{actionTargetId}][%d] getJourneyActiontargetUnauthorized  %+v", 401, o.Payload)
}

func (o *GetJourneyActiontargetUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActiontargetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActiontargetForbidden creates a GetJourneyActiontargetForbidden with default headers values
func NewGetJourneyActiontargetForbidden() *GetJourneyActiontargetForbidden {
	return &GetJourneyActiontargetForbidden{}
}

/*
GetJourneyActiontargetForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetJourneyActiontargetForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actiontarget forbidden response has a 2xx status code
func (o *GetJourneyActiontargetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actiontarget forbidden response has a 3xx status code
func (o *GetJourneyActiontargetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actiontarget forbidden response has a 4xx status code
func (o *GetJourneyActiontargetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey actiontarget forbidden response has a 5xx status code
func (o *GetJourneyActiontargetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey actiontarget forbidden response a status code equal to that given
func (o *GetJourneyActiontargetForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetJourneyActiontargetForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets/{actionTargetId}][%d] getJourneyActiontargetForbidden  %+v", 403, o.Payload)
}

func (o *GetJourneyActiontargetForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets/{actionTargetId}][%d] getJourneyActiontargetForbidden  %+v", 403, o.Payload)
}

func (o *GetJourneyActiontargetForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActiontargetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActiontargetNotFound creates a GetJourneyActiontargetNotFound with default headers values
func NewGetJourneyActiontargetNotFound() *GetJourneyActiontargetNotFound {
	return &GetJourneyActiontargetNotFound{}
}

/*
GetJourneyActiontargetNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetJourneyActiontargetNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actiontarget not found response has a 2xx status code
func (o *GetJourneyActiontargetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actiontarget not found response has a 3xx status code
func (o *GetJourneyActiontargetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actiontarget not found response has a 4xx status code
func (o *GetJourneyActiontargetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey actiontarget not found response has a 5xx status code
func (o *GetJourneyActiontargetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey actiontarget not found response a status code equal to that given
func (o *GetJourneyActiontargetNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetJourneyActiontargetNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets/{actionTargetId}][%d] getJourneyActiontargetNotFound  %+v", 404, o.Payload)
}

func (o *GetJourneyActiontargetNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets/{actionTargetId}][%d] getJourneyActiontargetNotFound  %+v", 404, o.Payload)
}

func (o *GetJourneyActiontargetNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActiontargetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActiontargetRequestTimeout creates a GetJourneyActiontargetRequestTimeout with default headers values
func NewGetJourneyActiontargetRequestTimeout() *GetJourneyActiontargetRequestTimeout {
	return &GetJourneyActiontargetRequestTimeout{}
}

/*
GetJourneyActiontargetRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetJourneyActiontargetRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actiontarget request timeout response has a 2xx status code
func (o *GetJourneyActiontargetRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actiontarget request timeout response has a 3xx status code
func (o *GetJourneyActiontargetRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actiontarget request timeout response has a 4xx status code
func (o *GetJourneyActiontargetRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey actiontarget request timeout response has a 5xx status code
func (o *GetJourneyActiontargetRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey actiontarget request timeout response a status code equal to that given
func (o *GetJourneyActiontargetRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetJourneyActiontargetRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets/{actionTargetId}][%d] getJourneyActiontargetRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetJourneyActiontargetRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets/{actionTargetId}][%d] getJourneyActiontargetRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetJourneyActiontargetRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActiontargetRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActiontargetRequestEntityTooLarge creates a GetJourneyActiontargetRequestEntityTooLarge with default headers values
func NewGetJourneyActiontargetRequestEntityTooLarge() *GetJourneyActiontargetRequestEntityTooLarge {
	return &GetJourneyActiontargetRequestEntityTooLarge{}
}

/*
GetJourneyActiontargetRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetJourneyActiontargetRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actiontarget request entity too large response has a 2xx status code
func (o *GetJourneyActiontargetRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actiontarget request entity too large response has a 3xx status code
func (o *GetJourneyActiontargetRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actiontarget request entity too large response has a 4xx status code
func (o *GetJourneyActiontargetRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey actiontarget request entity too large response has a 5xx status code
func (o *GetJourneyActiontargetRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey actiontarget request entity too large response a status code equal to that given
func (o *GetJourneyActiontargetRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetJourneyActiontargetRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets/{actionTargetId}][%d] getJourneyActiontargetRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetJourneyActiontargetRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets/{actionTargetId}][%d] getJourneyActiontargetRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetJourneyActiontargetRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActiontargetRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActiontargetUnsupportedMediaType creates a GetJourneyActiontargetUnsupportedMediaType with default headers values
func NewGetJourneyActiontargetUnsupportedMediaType() *GetJourneyActiontargetUnsupportedMediaType {
	return &GetJourneyActiontargetUnsupportedMediaType{}
}

/*
GetJourneyActiontargetUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetJourneyActiontargetUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actiontarget unsupported media type response has a 2xx status code
func (o *GetJourneyActiontargetUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actiontarget unsupported media type response has a 3xx status code
func (o *GetJourneyActiontargetUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actiontarget unsupported media type response has a 4xx status code
func (o *GetJourneyActiontargetUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey actiontarget unsupported media type response has a 5xx status code
func (o *GetJourneyActiontargetUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey actiontarget unsupported media type response a status code equal to that given
func (o *GetJourneyActiontargetUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetJourneyActiontargetUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets/{actionTargetId}][%d] getJourneyActiontargetUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetJourneyActiontargetUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets/{actionTargetId}][%d] getJourneyActiontargetUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetJourneyActiontargetUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActiontargetUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActiontargetTooManyRequests creates a GetJourneyActiontargetTooManyRequests with default headers values
func NewGetJourneyActiontargetTooManyRequests() *GetJourneyActiontargetTooManyRequests {
	return &GetJourneyActiontargetTooManyRequests{}
}

/*
GetJourneyActiontargetTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetJourneyActiontargetTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actiontarget too many requests response has a 2xx status code
func (o *GetJourneyActiontargetTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actiontarget too many requests response has a 3xx status code
func (o *GetJourneyActiontargetTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actiontarget too many requests response has a 4xx status code
func (o *GetJourneyActiontargetTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey actiontarget too many requests response has a 5xx status code
func (o *GetJourneyActiontargetTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey actiontarget too many requests response a status code equal to that given
func (o *GetJourneyActiontargetTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetJourneyActiontargetTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets/{actionTargetId}][%d] getJourneyActiontargetTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetJourneyActiontargetTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets/{actionTargetId}][%d] getJourneyActiontargetTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetJourneyActiontargetTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActiontargetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActiontargetInternalServerError creates a GetJourneyActiontargetInternalServerError with default headers values
func NewGetJourneyActiontargetInternalServerError() *GetJourneyActiontargetInternalServerError {
	return &GetJourneyActiontargetInternalServerError{}
}

/*
GetJourneyActiontargetInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetJourneyActiontargetInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actiontarget internal server error response has a 2xx status code
func (o *GetJourneyActiontargetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actiontarget internal server error response has a 3xx status code
func (o *GetJourneyActiontargetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actiontarget internal server error response has a 4xx status code
func (o *GetJourneyActiontargetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get journey actiontarget internal server error response has a 5xx status code
func (o *GetJourneyActiontargetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get journey actiontarget internal server error response a status code equal to that given
func (o *GetJourneyActiontargetInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetJourneyActiontargetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets/{actionTargetId}][%d] getJourneyActiontargetInternalServerError  %+v", 500, o.Payload)
}

func (o *GetJourneyActiontargetInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets/{actionTargetId}][%d] getJourneyActiontargetInternalServerError  %+v", 500, o.Payload)
}

func (o *GetJourneyActiontargetInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActiontargetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActiontargetServiceUnavailable creates a GetJourneyActiontargetServiceUnavailable with default headers values
func NewGetJourneyActiontargetServiceUnavailable() *GetJourneyActiontargetServiceUnavailable {
	return &GetJourneyActiontargetServiceUnavailable{}
}

/*
GetJourneyActiontargetServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetJourneyActiontargetServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actiontarget service unavailable response has a 2xx status code
func (o *GetJourneyActiontargetServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actiontarget service unavailable response has a 3xx status code
func (o *GetJourneyActiontargetServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actiontarget service unavailable response has a 4xx status code
func (o *GetJourneyActiontargetServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get journey actiontarget service unavailable response has a 5xx status code
func (o *GetJourneyActiontargetServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get journey actiontarget service unavailable response a status code equal to that given
func (o *GetJourneyActiontargetServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetJourneyActiontargetServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets/{actionTargetId}][%d] getJourneyActiontargetServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetJourneyActiontargetServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets/{actionTargetId}][%d] getJourneyActiontargetServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetJourneyActiontargetServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActiontargetServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActiontargetGatewayTimeout creates a GetJourneyActiontargetGatewayTimeout with default headers values
func NewGetJourneyActiontargetGatewayTimeout() *GetJourneyActiontargetGatewayTimeout {
	return &GetJourneyActiontargetGatewayTimeout{}
}

/*
GetJourneyActiontargetGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetJourneyActiontargetGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actiontarget gateway timeout response has a 2xx status code
func (o *GetJourneyActiontargetGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actiontarget gateway timeout response has a 3xx status code
func (o *GetJourneyActiontargetGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actiontarget gateway timeout response has a 4xx status code
func (o *GetJourneyActiontargetGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get journey actiontarget gateway timeout response has a 5xx status code
func (o *GetJourneyActiontargetGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get journey actiontarget gateway timeout response a status code equal to that given
func (o *GetJourneyActiontargetGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetJourneyActiontargetGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets/{actionTargetId}][%d] getJourneyActiontargetGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetJourneyActiontargetGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets/{actionTargetId}][%d] getJourneyActiontargetGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetJourneyActiontargetGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActiontargetGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

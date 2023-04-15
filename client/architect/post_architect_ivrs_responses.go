// Code generated by go-swagger; DO NOT EDIT.

package architect

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostArchitectIvrsReader is a Reader for the PostArchitectIvrs structure.
type PostArchitectIvrsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostArchitectIvrsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostArchitectIvrsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostArchitectIvrsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostArchitectIvrsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostArchitectIvrsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostArchitectIvrsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostArchitectIvrsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostArchitectIvrsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostArchitectIvrsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostArchitectIvrsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostArchitectIvrsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostArchitectIvrsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostArchitectIvrsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostArchitectIvrsOK creates a PostArchitectIvrsOK with default headers values
func NewPostArchitectIvrsOK() *PostArchitectIvrsOK {
	return &PostArchitectIvrsOK{}
}

/*
PostArchitectIvrsOK describes a response with status code 200, with default header values.

successful operation
*/
type PostArchitectIvrsOK struct {
	Payload *models.IVR
}

// IsSuccess returns true when this post architect ivrs o k response has a 2xx status code
func (o *PostArchitectIvrsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post architect ivrs o k response has a 3xx status code
func (o *PostArchitectIvrsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect ivrs o k response has a 4xx status code
func (o *PostArchitectIvrsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post architect ivrs o k response has a 5xx status code
func (o *PostArchitectIvrsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect ivrs o k response a status code equal to that given
func (o *PostArchitectIvrsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostArchitectIvrsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/ivrs][%d] postArchitectIvrsOK  %+v", 200, o.Payload)
}

func (o *PostArchitectIvrsOK) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/ivrs][%d] postArchitectIvrsOK  %+v", 200, o.Payload)
}

func (o *PostArchitectIvrsOK) GetPayload() *models.IVR {
	return o.Payload
}

func (o *PostArchitectIvrsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IVR)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectIvrsBadRequest creates a PostArchitectIvrsBadRequest with default headers values
func NewPostArchitectIvrsBadRequest() *PostArchitectIvrsBadRequest {
	return &PostArchitectIvrsBadRequest{}
}

/*
PostArchitectIvrsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostArchitectIvrsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect ivrs bad request response has a 2xx status code
func (o *PostArchitectIvrsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect ivrs bad request response has a 3xx status code
func (o *PostArchitectIvrsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect ivrs bad request response has a 4xx status code
func (o *PostArchitectIvrsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post architect ivrs bad request response has a 5xx status code
func (o *PostArchitectIvrsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect ivrs bad request response a status code equal to that given
func (o *PostArchitectIvrsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostArchitectIvrsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/ivrs][%d] postArchitectIvrsBadRequest  %+v", 400, o.Payload)
}

func (o *PostArchitectIvrsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/ivrs][%d] postArchitectIvrsBadRequest  %+v", 400, o.Payload)
}

func (o *PostArchitectIvrsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectIvrsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectIvrsUnauthorized creates a PostArchitectIvrsUnauthorized with default headers values
func NewPostArchitectIvrsUnauthorized() *PostArchitectIvrsUnauthorized {
	return &PostArchitectIvrsUnauthorized{}
}

/*
PostArchitectIvrsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostArchitectIvrsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect ivrs unauthorized response has a 2xx status code
func (o *PostArchitectIvrsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect ivrs unauthorized response has a 3xx status code
func (o *PostArchitectIvrsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect ivrs unauthorized response has a 4xx status code
func (o *PostArchitectIvrsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post architect ivrs unauthorized response has a 5xx status code
func (o *PostArchitectIvrsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect ivrs unauthorized response a status code equal to that given
func (o *PostArchitectIvrsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostArchitectIvrsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/ivrs][%d] postArchitectIvrsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostArchitectIvrsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/ivrs][%d] postArchitectIvrsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostArchitectIvrsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectIvrsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectIvrsForbidden creates a PostArchitectIvrsForbidden with default headers values
func NewPostArchitectIvrsForbidden() *PostArchitectIvrsForbidden {
	return &PostArchitectIvrsForbidden{}
}

/*
PostArchitectIvrsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostArchitectIvrsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect ivrs forbidden response has a 2xx status code
func (o *PostArchitectIvrsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect ivrs forbidden response has a 3xx status code
func (o *PostArchitectIvrsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect ivrs forbidden response has a 4xx status code
func (o *PostArchitectIvrsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post architect ivrs forbidden response has a 5xx status code
func (o *PostArchitectIvrsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect ivrs forbidden response a status code equal to that given
func (o *PostArchitectIvrsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostArchitectIvrsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/ivrs][%d] postArchitectIvrsForbidden  %+v", 403, o.Payload)
}

func (o *PostArchitectIvrsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/ivrs][%d] postArchitectIvrsForbidden  %+v", 403, o.Payload)
}

func (o *PostArchitectIvrsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectIvrsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectIvrsNotFound creates a PostArchitectIvrsNotFound with default headers values
func NewPostArchitectIvrsNotFound() *PostArchitectIvrsNotFound {
	return &PostArchitectIvrsNotFound{}
}

/*
PostArchitectIvrsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostArchitectIvrsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect ivrs not found response has a 2xx status code
func (o *PostArchitectIvrsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect ivrs not found response has a 3xx status code
func (o *PostArchitectIvrsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect ivrs not found response has a 4xx status code
func (o *PostArchitectIvrsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post architect ivrs not found response has a 5xx status code
func (o *PostArchitectIvrsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect ivrs not found response a status code equal to that given
func (o *PostArchitectIvrsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostArchitectIvrsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/ivrs][%d] postArchitectIvrsNotFound  %+v", 404, o.Payload)
}

func (o *PostArchitectIvrsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/ivrs][%d] postArchitectIvrsNotFound  %+v", 404, o.Payload)
}

func (o *PostArchitectIvrsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectIvrsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectIvrsRequestTimeout creates a PostArchitectIvrsRequestTimeout with default headers values
func NewPostArchitectIvrsRequestTimeout() *PostArchitectIvrsRequestTimeout {
	return &PostArchitectIvrsRequestTimeout{}
}

/*
PostArchitectIvrsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostArchitectIvrsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect ivrs request timeout response has a 2xx status code
func (o *PostArchitectIvrsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect ivrs request timeout response has a 3xx status code
func (o *PostArchitectIvrsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect ivrs request timeout response has a 4xx status code
func (o *PostArchitectIvrsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post architect ivrs request timeout response has a 5xx status code
func (o *PostArchitectIvrsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect ivrs request timeout response a status code equal to that given
func (o *PostArchitectIvrsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostArchitectIvrsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/ivrs][%d] postArchitectIvrsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostArchitectIvrsRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/ivrs][%d] postArchitectIvrsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostArchitectIvrsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectIvrsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectIvrsRequestEntityTooLarge creates a PostArchitectIvrsRequestEntityTooLarge with default headers values
func NewPostArchitectIvrsRequestEntityTooLarge() *PostArchitectIvrsRequestEntityTooLarge {
	return &PostArchitectIvrsRequestEntityTooLarge{}
}

/*
PostArchitectIvrsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostArchitectIvrsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect ivrs request entity too large response has a 2xx status code
func (o *PostArchitectIvrsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect ivrs request entity too large response has a 3xx status code
func (o *PostArchitectIvrsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect ivrs request entity too large response has a 4xx status code
func (o *PostArchitectIvrsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post architect ivrs request entity too large response has a 5xx status code
func (o *PostArchitectIvrsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect ivrs request entity too large response a status code equal to that given
func (o *PostArchitectIvrsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostArchitectIvrsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/ivrs][%d] postArchitectIvrsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostArchitectIvrsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/ivrs][%d] postArchitectIvrsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostArchitectIvrsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectIvrsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectIvrsUnsupportedMediaType creates a PostArchitectIvrsUnsupportedMediaType with default headers values
func NewPostArchitectIvrsUnsupportedMediaType() *PostArchitectIvrsUnsupportedMediaType {
	return &PostArchitectIvrsUnsupportedMediaType{}
}

/*
PostArchitectIvrsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostArchitectIvrsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect ivrs unsupported media type response has a 2xx status code
func (o *PostArchitectIvrsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect ivrs unsupported media type response has a 3xx status code
func (o *PostArchitectIvrsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect ivrs unsupported media type response has a 4xx status code
func (o *PostArchitectIvrsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post architect ivrs unsupported media type response has a 5xx status code
func (o *PostArchitectIvrsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect ivrs unsupported media type response a status code equal to that given
func (o *PostArchitectIvrsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostArchitectIvrsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/ivrs][%d] postArchitectIvrsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostArchitectIvrsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/ivrs][%d] postArchitectIvrsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostArchitectIvrsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectIvrsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectIvrsTooManyRequests creates a PostArchitectIvrsTooManyRequests with default headers values
func NewPostArchitectIvrsTooManyRequests() *PostArchitectIvrsTooManyRequests {
	return &PostArchitectIvrsTooManyRequests{}
}

/*
PostArchitectIvrsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostArchitectIvrsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect ivrs too many requests response has a 2xx status code
func (o *PostArchitectIvrsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect ivrs too many requests response has a 3xx status code
func (o *PostArchitectIvrsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect ivrs too many requests response has a 4xx status code
func (o *PostArchitectIvrsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post architect ivrs too many requests response has a 5xx status code
func (o *PostArchitectIvrsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect ivrs too many requests response a status code equal to that given
func (o *PostArchitectIvrsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostArchitectIvrsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/ivrs][%d] postArchitectIvrsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostArchitectIvrsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/ivrs][%d] postArchitectIvrsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostArchitectIvrsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectIvrsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectIvrsInternalServerError creates a PostArchitectIvrsInternalServerError with default headers values
func NewPostArchitectIvrsInternalServerError() *PostArchitectIvrsInternalServerError {
	return &PostArchitectIvrsInternalServerError{}
}

/*
PostArchitectIvrsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostArchitectIvrsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect ivrs internal server error response has a 2xx status code
func (o *PostArchitectIvrsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect ivrs internal server error response has a 3xx status code
func (o *PostArchitectIvrsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect ivrs internal server error response has a 4xx status code
func (o *PostArchitectIvrsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post architect ivrs internal server error response has a 5xx status code
func (o *PostArchitectIvrsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post architect ivrs internal server error response a status code equal to that given
func (o *PostArchitectIvrsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostArchitectIvrsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/ivrs][%d] postArchitectIvrsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostArchitectIvrsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/ivrs][%d] postArchitectIvrsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostArchitectIvrsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectIvrsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectIvrsServiceUnavailable creates a PostArchitectIvrsServiceUnavailable with default headers values
func NewPostArchitectIvrsServiceUnavailable() *PostArchitectIvrsServiceUnavailable {
	return &PostArchitectIvrsServiceUnavailable{}
}

/*
PostArchitectIvrsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostArchitectIvrsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect ivrs service unavailable response has a 2xx status code
func (o *PostArchitectIvrsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect ivrs service unavailable response has a 3xx status code
func (o *PostArchitectIvrsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect ivrs service unavailable response has a 4xx status code
func (o *PostArchitectIvrsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post architect ivrs service unavailable response has a 5xx status code
func (o *PostArchitectIvrsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post architect ivrs service unavailable response a status code equal to that given
func (o *PostArchitectIvrsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostArchitectIvrsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/ivrs][%d] postArchitectIvrsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostArchitectIvrsServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/ivrs][%d] postArchitectIvrsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostArchitectIvrsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectIvrsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectIvrsGatewayTimeout creates a PostArchitectIvrsGatewayTimeout with default headers values
func NewPostArchitectIvrsGatewayTimeout() *PostArchitectIvrsGatewayTimeout {
	return &PostArchitectIvrsGatewayTimeout{}
}

/*
PostArchitectIvrsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostArchitectIvrsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect ivrs gateway timeout response has a 2xx status code
func (o *PostArchitectIvrsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect ivrs gateway timeout response has a 3xx status code
func (o *PostArchitectIvrsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect ivrs gateway timeout response has a 4xx status code
func (o *PostArchitectIvrsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post architect ivrs gateway timeout response has a 5xx status code
func (o *PostArchitectIvrsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post architect ivrs gateway timeout response a status code equal to that given
func (o *PostArchitectIvrsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostArchitectIvrsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/ivrs][%d] postArchitectIvrsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostArchitectIvrsGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/ivrs][%d] postArchitectIvrsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostArchitectIvrsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectIvrsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

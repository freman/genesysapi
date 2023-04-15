// Code generated by go-swagger; DO NOT EDIT.

package audit

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostAuditsQueryRealtimeReader is a Reader for the PostAuditsQueryRealtime structure.
type PostAuditsQueryRealtimeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAuditsQueryRealtimeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAuditsQueryRealtimeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAuditsQueryRealtimeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostAuditsQueryRealtimeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostAuditsQueryRealtimeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAuditsQueryRealtimeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostAuditsQueryRealtimeRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostAuditsQueryRealtimeRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostAuditsQueryRealtimeUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostAuditsQueryRealtimeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAuditsQueryRealtimeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostAuditsQueryRealtimeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostAuditsQueryRealtimeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAuditsQueryRealtimeOK creates a PostAuditsQueryRealtimeOK with default headers values
func NewPostAuditsQueryRealtimeOK() *PostAuditsQueryRealtimeOK {
	return &PostAuditsQueryRealtimeOK{}
}

/*
PostAuditsQueryRealtimeOK describes a response with status code 200, with default header values.

successful operation
*/
type PostAuditsQueryRealtimeOK struct {
	Payload *models.AuditRealtimeQueryResultsResponse
}

// IsSuccess returns true when this post audits query realtime o k response has a 2xx status code
func (o *PostAuditsQueryRealtimeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post audits query realtime o k response has a 3xx status code
func (o *PostAuditsQueryRealtimeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post audits query realtime o k response has a 4xx status code
func (o *PostAuditsQueryRealtimeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post audits query realtime o k response has a 5xx status code
func (o *PostAuditsQueryRealtimeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post audits query realtime o k response a status code equal to that given
func (o *PostAuditsQueryRealtimeOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostAuditsQueryRealtimeOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/audits/query/realtime][%d] postAuditsQueryRealtimeOK  %+v", 200, o.Payload)
}

func (o *PostAuditsQueryRealtimeOK) String() string {
	return fmt.Sprintf("[POST /api/v2/audits/query/realtime][%d] postAuditsQueryRealtimeOK  %+v", 200, o.Payload)
}

func (o *PostAuditsQueryRealtimeOK) GetPayload() *models.AuditRealtimeQueryResultsResponse {
	return o.Payload
}

func (o *PostAuditsQueryRealtimeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuditRealtimeQueryResultsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuditsQueryRealtimeBadRequest creates a PostAuditsQueryRealtimeBadRequest with default headers values
func NewPostAuditsQueryRealtimeBadRequest() *PostAuditsQueryRealtimeBadRequest {
	return &PostAuditsQueryRealtimeBadRequest{}
}

/*
PostAuditsQueryRealtimeBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostAuditsQueryRealtimeBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post audits query realtime bad request response has a 2xx status code
func (o *PostAuditsQueryRealtimeBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post audits query realtime bad request response has a 3xx status code
func (o *PostAuditsQueryRealtimeBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post audits query realtime bad request response has a 4xx status code
func (o *PostAuditsQueryRealtimeBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post audits query realtime bad request response has a 5xx status code
func (o *PostAuditsQueryRealtimeBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post audits query realtime bad request response a status code equal to that given
func (o *PostAuditsQueryRealtimeBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostAuditsQueryRealtimeBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/audits/query/realtime][%d] postAuditsQueryRealtimeBadRequest  %+v", 400, o.Payload)
}

func (o *PostAuditsQueryRealtimeBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/audits/query/realtime][%d] postAuditsQueryRealtimeBadRequest  %+v", 400, o.Payload)
}

func (o *PostAuditsQueryRealtimeBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuditsQueryRealtimeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuditsQueryRealtimeUnauthorized creates a PostAuditsQueryRealtimeUnauthorized with default headers values
func NewPostAuditsQueryRealtimeUnauthorized() *PostAuditsQueryRealtimeUnauthorized {
	return &PostAuditsQueryRealtimeUnauthorized{}
}

/*
PostAuditsQueryRealtimeUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostAuditsQueryRealtimeUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post audits query realtime unauthorized response has a 2xx status code
func (o *PostAuditsQueryRealtimeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post audits query realtime unauthorized response has a 3xx status code
func (o *PostAuditsQueryRealtimeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post audits query realtime unauthorized response has a 4xx status code
func (o *PostAuditsQueryRealtimeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post audits query realtime unauthorized response has a 5xx status code
func (o *PostAuditsQueryRealtimeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post audits query realtime unauthorized response a status code equal to that given
func (o *PostAuditsQueryRealtimeUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostAuditsQueryRealtimeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/audits/query/realtime][%d] postAuditsQueryRealtimeUnauthorized  %+v", 401, o.Payload)
}

func (o *PostAuditsQueryRealtimeUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/audits/query/realtime][%d] postAuditsQueryRealtimeUnauthorized  %+v", 401, o.Payload)
}

func (o *PostAuditsQueryRealtimeUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuditsQueryRealtimeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuditsQueryRealtimeForbidden creates a PostAuditsQueryRealtimeForbidden with default headers values
func NewPostAuditsQueryRealtimeForbidden() *PostAuditsQueryRealtimeForbidden {
	return &PostAuditsQueryRealtimeForbidden{}
}

/*
PostAuditsQueryRealtimeForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostAuditsQueryRealtimeForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post audits query realtime forbidden response has a 2xx status code
func (o *PostAuditsQueryRealtimeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post audits query realtime forbidden response has a 3xx status code
func (o *PostAuditsQueryRealtimeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post audits query realtime forbidden response has a 4xx status code
func (o *PostAuditsQueryRealtimeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post audits query realtime forbidden response has a 5xx status code
func (o *PostAuditsQueryRealtimeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post audits query realtime forbidden response a status code equal to that given
func (o *PostAuditsQueryRealtimeForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostAuditsQueryRealtimeForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/audits/query/realtime][%d] postAuditsQueryRealtimeForbidden  %+v", 403, o.Payload)
}

func (o *PostAuditsQueryRealtimeForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/audits/query/realtime][%d] postAuditsQueryRealtimeForbidden  %+v", 403, o.Payload)
}

func (o *PostAuditsQueryRealtimeForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuditsQueryRealtimeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuditsQueryRealtimeNotFound creates a PostAuditsQueryRealtimeNotFound with default headers values
func NewPostAuditsQueryRealtimeNotFound() *PostAuditsQueryRealtimeNotFound {
	return &PostAuditsQueryRealtimeNotFound{}
}

/*
PostAuditsQueryRealtimeNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostAuditsQueryRealtimeNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post audits query realtime not found response has a 2xx status code
func (o *PostAuditsQueryRealtimeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post audits query realtime not found response has a 3xx status code
func (o *PostAuditsQueryRealtimeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post audits query realtime not found response has a 4xx status code
func (o *PostAuditsQueryRealtimeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post audits query realtime not found response has a 5xx status code
func (o *PostAuditsQueryRealtimeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post audits query realtime not found response a status code equal to that given
func (o *PostAuditsQueryRealtimeNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostAuditsQueryRealtimeNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/audits/query/realtime][%d] postAuditsQueryRealtimeNotFound  %+v", 404, o.Payload)
}

func (o *PostAuditsQueryRealtimeNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/audits/query/realtime][%d] postAuditsQueryRealtimeNotFound  %+v", 404, o.Payload)
}

func (o *PostAuditsQueryRealtimeNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuditsQueryRealtimeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuditsQueryRealtimeRequestTimeout creates a PostAuditsQueryRealtimeRequestTimeout with default headers values
func NewPostAuditsQueryRealtimeRequestTimeout() *PostAuditsQueryRealtimeRequestTimeout {
	return &PostAuditsQueryRealtimeRequestTimeout{}
}

/*
PostAuditsQueryRealtimeRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostAuditsQueryRealtimeRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post audits query realtime request timeout response has a 2xx status code
func (o *PostAuditsQueryRealtimeRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post audits query realtime request timeout response has a 3xx status code
func (o *PostAuditsQueryRealtimeRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post audits query realtime request timeout response has a 4xx status code
func (o *PostAuditsQueryRealtimeRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post audits query realtime request timeout response has a 5xx status code
func (o *PostAuditsQueryRealtimeRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post audits query realtime request timeout response a status code equal to that given
func (o *PostAuditsQueryRealtimeRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostAuditsQueryRealtimeRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/audits/query/realtime][%d] postAuditsQueryRealtimeRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostAuditsQueryRealtimeRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/audits/query/realtime][%d] postAuditsQueryRealtimeRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostAuditsQueryRealtimeRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuditsQueryRealtimeRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuditsQueryRealtimeRequestEntityTooLarge creates a PostAuditsQueryRealtimeRequestEntityTooLarge with default headers values
func NewPostAuditsQueryRealtimeRequestEntityTooLarge() *PostAuditsQueryRealtimeRequestEntityTooLarge {
	return &PostAuditsQueryRealtimeRequestEntityTooLarge{}
}

/*
PostAuditsQueryRealtimeRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostAuditsQueryRealtimeRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post audits query realtime request entity too large response has a 2xx status code
func (o *PostAuditsQueryRealtimeRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post audits query realtime request entity too large response has a 3xx status code
func (o *PostAuditsQueryRealtimeRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post audits query realtime request entity too large response has a 4xx status code
func (o *PostAuditsQueryRealtimeRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post audits query realtime request entity too large response has a 5xx status code
func (o *PostAuditsQueryRealtimeRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post audits query realtime request entity too large response a status code equal to that given
func (o *PostAuditsQueryRealtimeRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostAuditsQueryRealtimeRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/audits/query/realtime][%d] postAuditsQueryRealtimeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostAuditsQueryRealtimeRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/audits/query/realtime][%d] postAuditsQueryRealtimeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostAuditsQueryRealtimeRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuditsQueryRealtimeRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuditsQueryRealtimeUnsupportedMediaType creates a PostAuditsQueryRealtimeUnsupportedMediaType with default headers values
func NewPostAuditsQueryRealtimeUnsupportedMediaType() *PostAuditsQueryRealtimeUnsupportedMediaType {
	return &PostAuditsQueryRealtimeUnsupportedMediaType{}
}

/*
PostAuditsQueryRealtimeUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostAuditsQueryRealtimeUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post audits query realtime unsupported media type response has a 2xx status code
func (o *PostAuditsQueryRealtimeUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post audits query realtime unsupported media type response has a 3xx status code
func (o *PostAuditsQueryRealtimeUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post audits query realtime unsupported media type response has a 4xx status code
func (o *PostAuditsQueryRealtimeUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post audits query realtime unsupported media type response has a 5xx status code
func (o *PostAuditsQueryRealtimeUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post audits query realtime unsupported media type response a status code equal to that given
func (o *PostAuditsQueryRealtimeUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostAuditsQueryRealtimeUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/audits/query/realtime][%d] postAuditsQueryRealtimeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostAuditsQueryRealtimeUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/audits/query/realtime][%d] postAuditsQueryRealtimeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostAuditsQueryRealtimeUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuditsQueryRealtimeUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuditsQueryRealtimeTooManyRequests creates a PostAuditsQueryRealtimeTooManyRequests with default headers values
func NewPostAuditsQueryRealtimeTooManyRequests() *PostAuditsQueryRealtimeTooManyRequests {
	return &PostAuditsQueryRealtimeTooManyRequests{}
}

/*
PostAuditsQueryRealtimeTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostAuditsQueryRealtimeTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post audits query realtime too many requests response has a 2xx status code
func (o *PostAuditsQueryRealtimeTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post audits query realtime too many requests response has a 3xx status code
func (o *PostAuditsQueryRealtimeTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post audits query realtime too many requests response has a 4xx status code
func (o *PostAuditsQueryRealtimeTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post audits query realtime too many requests response has a 5xx status code
func (o *PostAuditsQueryRealtimeTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post audits query realtime too many requests response a status code equal to that given
func (o *PostAuditsQueryRealtimeTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostAuditsQueryRealtimeTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/audits/query/realtime][%d] postAuditsQueryRealtimeTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostAuditsQueryRealtimeTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/audits/query/realtime][%d] postAuditsQueryRealtimeTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostAuditsQueryRealtimeTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuditsQueryRealtimeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuditsQueryRealtimeInternalServerError creates a PostAuditsQueryRealtimeInternalServerError with default headers values
func NewPostAuditsQueryRealtimeInternalServerError() *PostAuditsQueryRealtimeInternalServerError {
	return &PostAuditsQueryRealtimeInternalServerError{}
}

/*
PostAuditsQueryRealtimeInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostAuditsQueryRealtimeInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post audits query realtime internal server error response has a 2xx status code
func (o *PostAuditsQueryRealtimeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post audits query realtime internal server error response has a 3xx status code
func (o *PostAuditsQueryRealtimeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post audits query realtime internal server error response has a 4xx status code
func (o *PostAuditsQueryRealtimeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post audits query realtime internal server error response has a 5xx status code
func (o *PostAuditsQueryRealtimeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post audits query realtime internal server error response a status code equal to that given
func (o *PostAuditsQueryRealtimeInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostAuditsQueryRealtimeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/audits/query/realtime][%d] postAuditsQueryRealtimeInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAuditsQueryRealtimeInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/audits/query/realtime][%d] postAuditsQueryRealtimeInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAuditsQueryRealtimeInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuditsQueryRealtimeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuditsQueryRealtimeServiceUnavailable creates a PostAuditsQueryRealtimeServiceUnavailable with default headers values
func NewPostAuditsQueryRealtimeServiceUnavailable() *PostAuditsQueryRealtimeServiceUnavailable {
	return &PostAuditsQueryRealtimeServiceUnavailable{}
}

/*
PostAuditsQueryRealtimeServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostAuditsQueryRealtimeServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post audits query realtime service unavailable response has a 2xx status code
func (o *PostAuditsQueryRealtimeServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post audits query realtime service unavailable response has a 3xx status code
func (o *PostAuditsQueryRealtimeServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post audits query realtime service unavailable response has a 4xx status code
func (o *PostAuditsQueryRealtimeServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post audits query realtime service unavailable response has a 5xx status code
func (o *PostAuditsQueryRealtimeServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post audits query realtime service unavailable response a status code equal to that given
func (o *PostAuditsQueryRealtimeServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostAuditsQueryRealtimeServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/audits/query/realtime][%d] postAuditsQueryRealtimeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostAuditsQueryRealtimeServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/audits/query/realtime][%d] postAuditsQueryRealtimeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostAuditsQueryRealtimeServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuditsQueryRealtimeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuditsQueryRealtimeGatewayTimeout creates a PostAuditsQueryRealtimeGatewayTimeout with default headers values
func NewPostAuditsQueryRealtimeGatewayTimeout() *PostAuditsQueryRealtimeGatewayTimeout {
	return &PostAuditsQueryRealtimeGatewayTimeout{}
}

/*
PostAuditsQueryRealtimeGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostAuditsQueryRealtimeGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post audits query realtime gateway timeout response has a 2xx status code
func (o *PostAuditsQueryRealtimeGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post audits query realtime gateway timeout response has a 3xx status code
func (o *PostAuditsQueryRealtimeGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post audits query realtime gateway timeout response has a 4xx status code
func (o *PostAuditsQueryRealtimeGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post audits query realtime gateway timeout response has a 5xx status code
func (o *PostAuditsQueryRealtimeGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post audits query realtime gateway timeout response a status code equal to that given
func (o *PostAuditsQueryRealtimeGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostAuditsQueryRealtimeGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/audits/query/realtime][%d] postAuditsQueryRealtimeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostAuditsQueryRealtimeGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/audits/query/realtime][%d] postAuditsQueryRealtimeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostAuditsQueryRealtimeGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuditsQueryRealtimeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

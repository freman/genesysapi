// Code generated by go-swagger; DO NOT EDIT.

package search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostSearchSuggestReader is a Reader for the PostSearchSuggest structure.
type PostSearchSuggestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSearchSuggestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostSearchSuggestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostSearchSuggestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostSearchSuggestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostSearchSuggestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostSearchSuggestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostSearchSuggestRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostSearchSuggestRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostSearchSuggestUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostSearchSuggestTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostSearchSuggestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostSearchSuggestServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostSearchSuggestGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostSearchSuggestOK creates a PostSearchSuggestOK with default headers values
func NewPostSearchSuggestOK() *PostSearchSuggestOK {
	return &PostSearchSuggestOK{}
}

/*
PostSearchSuggestOK describes a response with status code 200, with default header values.

successful operation
*/
type PostSearchSuggestOK struct {
	Payload *models.JSONNodeSearchResponse
}

// IsSuccess returns true when this post search suggest o k response has a 2xx status code
func (o *PostSearchSuggestOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post search suggest o k response has a 3xx status code
func (o *PostSearchSuggestOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post search suggest o k response has a 4xx status code
func (o *PostSearchSuggestOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post search suggest o k response has a 5xx status code
func (o *PostSearchSuggestOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post search suggest o k response a status code equal to that given
func (o *PostSearchSuggestOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostSearchSuggestOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/search/suggest][%d] postSearchSuggestOK  %+v", 200, o.Payload)
}

func (o *PostSearchSuggestOK) String() string {
	return fmt.Sprintf("[POST /api/v2/search/suggest][%d] postSearchSuggestOK  %+v", 200, o.Payload)
}

func (o *PostSearchSuggestOK) GetPayload() *models.JSONNodeSearchResponse {
	return o.Payload
}

func (o *PostSearchSuggestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONNodeSearchResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSearchSuggestBadRequest creates a PostSearchSuggestBadRequest with default headers values
func NewPostSearchSuggestBadRequest() *PostSearchSuggestBadRequest {
	return &PostSearchSuggestBadRequest{}
}

/*
PostSearchSuggestBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostSearchSuggestBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post search suggest bad request response has a 2xx status code
func (o *PostSearchSuggestBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post search suggest bad request response has a 3xx status code
func (o *PostSearchSuggestBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post search suggest bad request response has a 4xx status code
func (o *PostSearchSuggestBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post search suggest bad request response has a 5xx status code
func (o *PostSearchSuggestBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post search suggest bad request response a status code equal to that given
func (o *PostSearchSuggestBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostSearchSuggestBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/search/suggest][%d] postSearchSuggestBadRequest  %+v", 400, o.Payload)
}

func (o *PostSearchSuggestBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/search/suggest][%d] postSearchSuggestBadRequest  %+v", 400, o.Payload)
}

func (o *PostSearchSuggestBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSearchSuggestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSearchSuggestUnauthorized creates a PostSearchSuggestUnauthorized with default headers values
func NewPostSearchSuggestUnauthorized() *PostSearchSuggestUnauthorized {
	return &PostSearchSuggestUnauthorized{}
}

/*
PostSearchSuggestUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostSearchSuggestUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post search suggest unauthorized response has a 2xx status code
func (o *PostSearchSuggestUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post search suggest unauthorized response has a 3xx status code
func (o *PostSearchSuggestUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post search suggest unauthorized response has a 4xx status code
func (o *PostSearchSuggestUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post search suggest unauthorized response has a 5xx status code
func (o *PostSearchSuggestUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post search suggest unauthorized response a status code equal to that given
func (o *PostSearchSuggestUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostSearchSuggestUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/search/suggest][%d] postSearchSuggestUnauthorized  %+v", 401, o.Payload)
}

func (o *PostSearchSuggestUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/search/suggest][%d] postSearchSuggestUnauthorized  %+v", 401, o.Payload)
}

func (o *PostSearchSuggestUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSearchSuggestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSearchSuggestForbidden creates a PostSearchSuggestForbidden with default headers values
func NewPostSearchSuggestForbidden() *PostSearchSuggestForbidden {
	return &PostSearchSuggestForbidden{}
}

/*
PostSearchSuggestForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostSearchSuggestForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post search suggest forbidden response has a 2xx status code
func (o *PostSearchSuggestForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post search suggest forbidden response has a 3xx status code
func (o *PostSearchSuggestForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post search suggest forbidden response has a 4xx status code
func (o *PostSearchSuggestForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post search suggest forbidden response has a 5xx status code
func (o *PostSearchSuggestForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post search suggest forbidden response a status code equal to that given
func (o *PostSearchSuggestForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostSearchSuggestForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/search/suggest][%d] postSearchSuggestForbidden  %+v", 403, o.Payload)
}

func (o *PostSearchSuggestForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/search/suggest][%d] postSearchSuggestForbidden  %+v", 403, o.Payload)
}

func (o *PostSearchSuggestForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSearchSuggestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSearchSuggestNotFound creates a PostSearchSuggestNotFound with default headers values
func NewPostSearchSuggestNotFound() *PostSearchSuggestNotFound {
	return &PostSearchSuggestNotFound{}
}

/*
PostSearchSuggestNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostSearchSuggestNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post search suggest not found response has a 2xx status code
func (o *PostSearchSuggestNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post search suggest not found response has a 3xx status code
func (o *PostSearchSuggestNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post search suggest not found response has a 4xx status code
func (o *PostSearchSuggestNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post search suggest not found response has a 5xx status code
func (o *PostSearchSuggestNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post search suggest not found response a status code equal to that given
func (o *PostSearchSuggestNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostSearchSuggestNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/search/suggest][%d] postSearchSuggestNotFound  %+v", 404, o.Payload)
}

func (o *PostSearchSuggestNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/search/suggest][%d] postSearchSuggestNotFound  %+v", 404, o.Payload)
}

func (o *PostSearchSuggestNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSearchSuggestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSearchSuggestRequestTimeout creates a PostSearchSuggestRequestTimeout with default headers values
func NewPostSearchSuggestRequestTimeout() *PostSearchSuggestRequestTimeout {
	return &PostSearchSuggestRequestTimeout{}
}

/*
PostSearchSuggestRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostSearchSuggestRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post search suggest request timeout response has a 2xx status code
func (o *PostSearchSuggestRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post search suggest request timeout response has a 3xx status code
func (o *PostSearchSuggestRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post search suggest request timeout response has a 4xx status code
func (o *PostSearchSuggestRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post search suggest request timeout response has a 5xx status code
func (o *PostSearchSuggestRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post search suggest request timeout response a status code equal to that given
func (o *PostSearchSuggestRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostSearchSuggestRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/search/suggest][%d] postSearchSuggestRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostSearchSuggestRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/search/suggest][%d] postSearchSuggestRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostSearchSuggestRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSearchSuggestRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSearchSuggestRequestEntityTooLarge creates a PostSearchSuggestRequestEntityTooLarge with default headers values
func NewPostSearchSuggestRequestEntityTooLarge() *PostSearchSuggestRequestEntityTooLarge {
	return &PostSearchSuggestRequestEntityTooLarge{}
}

/*
PostSearchSuggestRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostSearchSuggestRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post search suggest request entity too large response has a 2xx status code
func (o *PostSearchSuggestRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post search suggest request entity too large response has a 3xx status code
func (o *PostSearchSuggestRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post search suggest request entity too large response has a 4xx status code
func (o *PostSearchSuggestRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post search suggest request entity too large response has a 5xx status code
func (o *PostSearchSuggestRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post search suggest request entity too large response a status code equal to that given
func (o *PostSearchSuggestRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostSearchSuggestRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/search/suggest][%d] postSearchSuggestRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostSearchSuggestRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/search/suggest][%d] postSearchSuggestRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostSearchSuggestRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSearchSuggestRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSearchSuggestUnsupportedMediaType creates a PostSearchSuggestUnsupportedMediaType with default headers values
func NewPostSearchSuggestUnsupportedMediaType() *PostSearchSuggestUnsupportedMediaType {
	return &PostSearchSuggestUnsupportedMediaType{}
}

/*
PostSearchSuggestUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostSearchSuggestUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post search suggest unsupported media type response has a 2xx status code
func (o *PostSearchSuggestUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post search suggest unsupported media type response has a 3xx status code
func (o *PostSearchSuggestUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post search suggest unsupported media type response has a 4xx status code
func (o *PostSearchSuggestUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post search suggest unsupported media type response has a 5xx status code
func (o *PostSearchSuggestUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post search suggest unsupported media type response a status code equal to that given
func (o *PostSearchSuggestUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostSearchSuggestUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/search/suggest][%d] postSearchSuggestUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostSearchSuggestUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/search/suggest][%d] postSearchSuggestUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostSearchSuggestUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSearchSuggestUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSearchSuggestTooManyRequests creates a PostSearchSuggestTooManyRequests with default headers values
func NewPostSearchSuggestTooManyRequests() *PostSearchSuggestTooManyRequests {
	return &PostSearchSuggestTooManyRequests{}
}

/*
PostSearchSuggestTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostSearchSuggestTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post search suggest too many requests response has a 2xx status code
func (o *PostSearchSuggestTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post search suggest too many requests response has a 3xx status code
func (o *PostSearchSuggestTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post search suggest too many requests response has a 4xx status code
func (o *PostSearchSuggestTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post search suggest too many requests response has a 5xx status code
func (o *PostSearchSuggestTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post search suggest too many requests response a status code equal to that given
func (o *PostSearchSuggestTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostSearchSuggestTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/search/suggest][%d] postSearchSuggestTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostSearchSuggestTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/search/suggest][%d] postSearchSuggestTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostSearchSuggestTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSearchSuggestTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSearchSuggestInternalServerError creates a PostSearchSuggestInternalServerError with default headers values
func NewPostSearchSuggestInternalServerError() *PostSearchSuggestInternalServerError {
	return &PostSearchSuggestInternalServerError{}
}

/*
PostSearchSuggestInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostSearchSuggestInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post search suggest internal server error response has a 2xx status code
func (o *PostSearchSuggestInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post search suggest internal server error response has a 3xx status code
func (o *PostSearchSuggestInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post search suggest internal server error response has a 4xx status code
func (o *PostSearchSuggestInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post search suggest internal server error response has a 5xx status code
func (o *PostSearchSuggestInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post search suggest internal server error response a status code equal to that given
func (o *PostSearchSuggestInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostSearchSuggestInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/search/suggest][%d] postSearchSuggestInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSearchSuggestInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/search/suggest][%d] postSearchSuggestInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSearchSuggestInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSearchSuggestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSearchSuggestServiceUnavailable creates a PostSearchSuggestServiceUnavailable with default headers values
func NewPostSearchSuggestServiceUnavailable() *PostSearchSuggestServiceUnavailable {
	return &PostSearchSuggestServiceUnavailable{}
}

/*
PostSearchSuggestServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostSearchSuggestServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post search suggest service unavailable response has a 2xx status code
func (o *PostSearchSuggestServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post search suggest service unavailable response has a 3xx status code
func (o *PostSearchSuggestServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post search suggest service unavailable response has a 4xx status code
func (o *PostSearchSuggestServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post search suggest service unavailable response has a 5xx status code
func (o *PostSearchSuggestServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post search suggest service unavailable response a status code equal to that given
func (o *PostSearchSuggestServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostSearchSuggestServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/search/suggest][%d] postSearchSuggestServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostSearchSuggestServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/search/suggest][%d] postSearchSuggestServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostSearchSuggestServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSearchSuggestServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSearchSuggestGatewayTimeout creates a PostSearchSuggestGatewayTimeout with default headers values
func NewPostSearchSuggestGatewayTimeout() *PostSearchSuggestGatewayTimeout {
	return &PostSearchSuggestGatewayTimeout{}
}

/*
PostSearchSuggestGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostSearchSuggestGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post search suggest gateway timeout response has a 2xx status code
func (o *PostSearchSuggestGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post search suggest gateway timeout response has a 3xx status code
func (o *PostSearchSuggestGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post search suggest gateway timeout response has a 4xx status code
func (o *PostSearchSuggestGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post search suggest gateway timeout response has a 5xx status code
func (o *PostSearchSuggestGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post search suggest gateway timeout response a status code equal to that given
func (o *PostSearchSuggestGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostSearchSuggestGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/search/suggest][%d] postSearchSuggestGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostSearchSuggestGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/search/suggest][%d] postSearchSuggestGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostSearchSuggestGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSearchSuggestGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

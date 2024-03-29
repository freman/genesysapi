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

// GetSearchSuggestReader is a Reader for the GetSearchSuggest structure.
type GetSearchSuggestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSearchSuggestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSearchSuggestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSearchSuggestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetSearchSuggestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSearchSuggestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSearchSuggestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetSearchSuggestRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetSearchSuggestRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetSearchSuggestUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetSearchSuggestTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSearchSuggestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetSearchSuggestServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetSearchSuggestGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSearchSuggestOK creates a GetSearchSuggestOK with default headers values
func NewGetSearchSuggestOK() *GetSearchSuggestOK {
	return &GetSearchSuggestOK{}
}

/*
GetSearchSuggestOK describes a response with status code 200, with default header values.

successful operation
*/
type GetSearchSuggestOK struct {
	Payload *models.JSONNodeSearchResponse
}

// IsSuccess returns true when this get search suggest o k response has a 2xx status code
func (o *GetSearchSuggestOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get search suggest o k response has a 3xx status code
func (o *GetSearchSuggestOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get search suggest o k response has a 4xx status code
func (o *GetSearchSuggestOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get search suggest o k response has a 5xx status code
func (o *GetSearchSuggestOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get search suggest o k response a status code equal to that given
func (o *GetSearchSuggestOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetSearchSuggestOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/search/suggest][%d] getSearchSuggestOK  %+v", 200, o.Payload)
}

func (o *GetSearchSuggestOK) String() string {
	return fmt.Sprintf("[GET /api/v2/search/suggest][%d] getSearchSuggestOK  %+v", 200, o.Payload)
}

func (o *GetSearchSuggestOK) GetPayload() *models.JSONNodeSearchResponse {
	return o.Payload
}

func (o *GetSearchSuggestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONNodeSearchResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSearchSuggestBadRequest creates a GetSearchSuggestBadRequest with default headers values
func NewGetSearchSuggestBadRequest() *GetSearchSuggestBadRequest {
	return &GetSearchSuggestBadRequest{}
}

/*
GetSearchSuggestBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetSearchSuggestBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get search suggest bad request response has a 2xx status code
func (o *GetSearchSuggestBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get search suggest bad request response has a 3xx status code
func (o *GetSearchSuggestBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get search suggest bad request response has a 4xx status code
func (o *GetSearchSuggestBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get search suggest bad request response has a 5xx status code
func (o *GetSearchSuggestBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get search suggest bad request response a status code equal to that given
func (o *GetSearchSuggestBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetSearchSuggestBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/search/suggest][%d] getSearchSuggestBadRequest  %+v", 400, o.Payload)
}

func (o *GetSearchSuggestBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/search/suggest][%d] getSearchSuggestBadRequest  %+v", 400, o.Payload)
}

func (o *GetSearchSuggestBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSearchSuggestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSearchSuggestUnauthorized creates a GetSearchSuggestUnauthorized with default headers values
func NewGetSearchSuggestUnauthorized() *GetSearchSuggestUnauthorized {
	return &GetSearchSuggestUnauthorized{}
}

/*
GetSearchSuggestUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetSearchSuggestUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get search suggest unauthorized response has a 2xx status code
func (o *GetSearchSuggestUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get search suggest unauthorized response has a 3xx status code
func (o *GetSearchSuggestUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get search suggest unauthorized response has a 4xx status code
func (o *GetSearchSuggestUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get search suggest unauthorized response has a 5xx status code
func (o *GetSearchSuggestUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get search suggest unauthorized response a status code equal to that given
func (o *GetSearchSuggestUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetSearchSuggestUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/search/suggest][%d] getSearchSuggestUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSearchSuggestUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/search/suggest][%d] getSearchSuggestUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSearchSuggestUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSearchSuggestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSearchSuggestForbidden creates a GetSearchSuggestForbidden with default headers values
func NewGetSearchSuggestForbidden() *GetSearchSuggestForbidden {
	return &GetSearchSuggestForbidden{}
}

/*
GetSearchSuggestForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetSearchSuggestForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get search suggest forbidden response has a 2xx status code
func (o *GetSearchSuggestForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get search suggest forbidden response has a 3xx status code
func (o *GetSearchSuggestForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get search suggest forbidden response has a 4xx status code
func (o *GetSearchSuggestForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get search suggest forbidden response has a 5xx status code
func (o *GetSearchSuggestForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get search suggest forbidden response a status code equal to that given
func (o *GetSearchSuggestForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetSearchSuggestForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/search/suggest][%d] getSearchSuggestForbidden  %+v", 403, o.Payload)
}

func (o *GetSearchSuggestForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/search/suggest][%d] getSearchSuggestForbidden  %+v", 403, o.Payload)
}

func (o *GetSearchSuggestForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSearchSuggestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSearchSuggestNotFound creates a GetSearchSuggestNotFound with default headers values
func NewGetSearchSuggestNotFound() *GetSearchSuggestNotFound {
	return &GetSearchSuggestNotFound{}
}

/*
GetSearchSuggestNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetSearchSuggestNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get search suggest not found response has a 2xx status code
func (o *GetSearchSuggestNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get search suggest not found response has a 3xx status code
func (o *GetSearchSuggestNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get search suggest not found response has a 4xx status code
func (o *GetSearchSuggestNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get search suggest not found response has a 5xx status code
func (o *GetSearchSuggestNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get search suggest not found response a status code equal to that given
func (o *GetSearchSuggestNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetSearchSuggestNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/search/suggest][%d] getSearchSuggestNotFound  %+v", 404, o.Payload)
}

func (o *GetSearchSuggestNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/search/suggest][%d] getSearchSuggestNotFound  %+v", 404, o.Payload)
}

func (o *GetSearchSuggestNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSearchSuggestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSearchSuggestRequestTimeout creates a GetSearchSuggestRequestTimeout with default headers values
func NewGetSearchSuggestRequestTimeout() *GetSearchSuggestRequestTimeout {
	return &GetSearchSuggestRequestTimeout{}
}

/*
GetSearchSuggestRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetSearchSuggestRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get search suggest request timeout response has a 2xx status code
func (o *GetSearchSuggestRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get search suggest request timeout response has a 3xx status code
func (o *GetSearchSuggestRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get search suggest request timeout response has a 4xx status code
func (o *GetSearchSuggestRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get search suggest request timeout response has a 5xx status code
func (o *GetSearchSuggestRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get search suggest request timeout response a status code equal to that given
func (o *GetSearchSuggestRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetSearchSuggestRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/search/suggest][%d] getSearchSuggestRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetSearchSuggestRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/search/suggest][%d] getSearchSuggestRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetSearchSuggestRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSearchSuggestRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSearchSuggestRequestEntityTooLarge creates a GetSearchSuggestRequestEntityTooLarge with default headers values
func NewGetSearchSuggestRequestEntityTooLarge() *GetSearchSuggestRequestEntityTooLarge {
	return &GetSearchSuggestRequestEntityTooLarge{}
}

/*
GetSearchSuggestRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetSearchSuggestRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get search suggest request entity too large response has a 2xx status code
func (o *GetSearchSuggestRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get search suggest request entity too large response has a 3xx status code
func (o *GetSearchSuggestRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get search suggest request entity too large response has a 4xx status code
func (o *GetSearchSuggestRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get search suggest request entity too large response has a 5xx status code
func (o *GetSearchSuggestRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get search suggest request entity too large response a status code equal to that given
func (o *GetSearchSuggestRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetSearchSuggestRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/search/suggest][%d] getSearchSuggestRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetSearchSuggestRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/search/suggest][%d] getSearchSuggestRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetSearchSuggestRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSearchSuggestRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSearchSuggestUnsupportedMediaType creates a GetSearchSuggestUnsupportedMediaType with default headers values
func NewGetSearchSuggestUnsupportedMediaType() *GetSearchSuggestUnsupportedMediaType {
	return &GetSearchSuggestUnsupportedMediaType{}
}

/*
GetSearchSuggestUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetSearchSuggestUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get search suggest unsupported media type response has a 2xx status code
func (o *GetSearchSuggestUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get search suggest unsupported media type response has a 3xx status code
func (o *GetSearchSuggestUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get search suggest unsupported media type response has a 4xx status code
func (o *GetSearchSuggestUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get search suggest unsupported media type response has a 5xx status code
func (o *GetSearchSuggestUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get search suggest unsupported media type response a status code equal to that given
func (o *GetSearchSuggestUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetSearchSuggestUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/search/suggest][%d] getSearchSuggestUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetSearchSuggestUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/search/suggest][%d] getSearchSuggestUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetSearchSuggestUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSearchSuggestUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSearchSuggestTooManyRequests creates a GetSearchSuggestTooManyRequests with default headers values
func NewGetSearchSuggestTooManyRequests() *GetSearchSuggestTooManyRequests {
	return &GetSearchSuggestTooManyRequests{}
}

/*
GetSearchSuggestTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetSearchSuggestTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get search suggest too many requests response has a 2xx status code
func (o *GetSearchSuggestTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get search suggest too many requests response has a 3xx status code
func (o *GetSearchSuggestTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get search suggest too many requests response has a 4xx status code
func (o *GetSearchSuggestTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get search suggest too many requests response has a 5xx status code
func (o *GetSearchSuggestTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get search suggest too many requests response a status code equal to that given
func (o *GetSearchSuggestTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetSearchSuggestTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/search/suggest][%d] getSearchSuggestTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSearchSuggestTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/search/suggest][%d] getSearchSuggestTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSearchSuggestTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSearchSuggestTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSearchSuggestInternalServerError creates a GetSearchSuggestInternalServerError with default headers values
func NewGetSearchSuggestInternalServerError() *GetSearchSuggestInternalServerError {
	return &GetSearchSuggestInternalServerError{}
}

/*
GetSearchSuggestInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetSearchSuggestInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get search suggest internal server error response has a 2xx status code
func (o *GetSearchSuggestInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get search suggest internal server error response has a 3xx status code
func (o *GetSearchSuggestInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get search suggest internal server error response has a 4xx status code
func (o *GetSearchSuggestInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get search suggest internal server error response has a 5xx status code
func (o *GetSearchSuggestInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get search suggest internal server error response a status code equal to that given
func (o *GetSearchSuggestInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetSearchSuggestInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/search/suggest][%d] getSearchSuggestInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSearchSuggestInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/search/suggest][%d] getSearchSuggestInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSearchSuggestInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSearchSuggestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSearchSuggestServiceUnavailable creates a GetSearchSuggestServiceUnavailable with default headers values
func NewGetSearchSuggestServiceUnavailable() *GetSearchSuggestServiceUnavailable {
	return &GetSearchSuggestServiceUnavailable{}
}

/*
GetSearchSuggestServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetSearchSuggestServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get search suggest service unavailable response has a 2xx status code
func (o *GetSearchSuggestServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get search suggest service unavailable response has a 3xx status code
func (o *GetSearchSuggestServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get search suggest service unavailable response has a 4xx status code
func (o *GetSearchSuggestServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get search suggest service unavailable response has a 5xx status code
func (o *GetSearchSuggestServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get search suggest service unavailable response a status code equal to that given
func (o *GetSearchSuggestServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetSearchSuggestServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/search/suggest][%d] getSearchSuggestServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetSearchSuggestServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/search/suggest][%d] getSearchSuggestServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetSearchSuggestServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSearchSuggestServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSearchSuggestGatewayTimeout creates a GetSearchSuggestGatewayTimeout with default headers values
func NewGetSearchSuggestGatewayTimeout() *GetSearchSuggestGatewayTimeout {
	return &GetSearchSuggestGatewayTimeout{}
}

/*
GetSearchSuggestGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetSearchSuggestGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get search suggest gateway timeout response has a 2xx status code
func (o *GetSearchSuggestGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get search suggest gateway timeout response has a 3xx status code
func (o *GetSearchSuggestGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get search suggest gateway timeout response has a 4xx status code
func (o *GetSearchSuggestGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get search suggest gateway timeout response has a 5xx status code
func (o *GetSearchSuggestGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get search suggest gateway timeout response a status code equal to that given
func (o *GetSearchSuggestGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetSearchSuggestGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/search/suggest][%d] getSearchSuggestGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetSearchSuggestGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/search/suggest][%d] getSearchSuggestGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetSearchSuggestGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSearchSuggestGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

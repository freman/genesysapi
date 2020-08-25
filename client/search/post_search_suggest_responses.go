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

/*PostSearchSuggestOK handles this case with default header values.

successful operation
*/
type PostSearchSuggestOK struct {
	Payload *models.JSONNodeSearchResponse
}

func (o *PostSearchSuggestOK) Error() string {
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

/*PostSearchSuggestBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostSearchSuggestBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostSearchSuggestBadRequest) Error() string {
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

/*PostSearchSuggestUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostSearchSuggestUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostSearchSuggestUnauthorized) Error() string {
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

/*PostSearchSuggestForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostSearchSuggestForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostSearchSuggestForbidden) Error() string {
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

/*PostSearchSuggestNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostSearchSuggestNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostSearchSuggestNotFound) Error() string {
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

// NewPostSearchSuggestRequestEntityTooLarge creates a PostSearchSuggestRequestEntityTooLarge with default headers values
func NewPostSearchSuggestRequestEntityTooLarge() *PostSearchSuggestRequestEntityTooLarge {
	return &PostSearchSuggestRequestEntityTooLarge{}
}

/*PostSearchSuggestRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostSearchSuggestRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostSearchSuggestRequestEntityTooLarge) Error() string {
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

/*PostSearchSuggestUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostSearchSuggestUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostSearchSuggestUnsupportedMediaType) Error() string {
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

/*PostSearchSuggestTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostSearchSuggestTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostSearchSuggestTooManyRequests) Error() string {
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

/*PostSearchSuggestInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostSearchSuggestInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostSearchSuggestInternalServerError) Error() string {
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

/*PostSearchSuggestServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostSearchSuggestServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostSearchSuggestServiceUnavailable) Error() string {
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

/*PostSearchSuggestGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostSearchSuggestGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostSearchSuggestGatewayTimeout) Error() string {
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
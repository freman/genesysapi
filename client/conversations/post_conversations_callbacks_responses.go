// Code generated by go-swagger; DO NOT EDIT.

package conversations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostConversationsCallbacksReader is a Reader for the PostConversationsCallbacks structure.
type PostConversationsCallbacksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostConversationsCallbacksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostConversationsCallbacksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPostConversationsCallbacksAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostConversationsCallbacksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostConversationsCallbacksUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostConversationsCallbacksForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostConversationsCallbacksNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostConversationsCallbacksRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostConversationsCallbacksRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostConversationsCallbacksUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostConversationsCallbacksTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostConversationsCallbacksInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostConversationsCallbacksServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostConversationsCallbacksGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostConversationsCallbacksOK creates a PostConversationsCallbacksOK with default headers values
func NewPostConversationsCallbacksOK() *PostConversationsCallbacksOK {
	return &PostConversationsCallbacksOK{}
}

/*PostConversationsCallbacksOK handles this case with default header values.

successful operation
*/
type PostConversationsCallbacksOK struct {
	Payload *models.CreateCallbackResponse
}

func (o *PostConversationsCallbacksOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/callbacks][%d] postConversationsCallbacksOK  %+v", 200, o.Payload)
}

func (o *PostConversationsCallbacksOK) GetPayload() *models.CreateCallbackResponse {
	return o.Payload
}

func (o *PostConversationsCallbacksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateCallbackResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsCallbacksAccepted creates a PostConversationsCallbacksAccepted with default headers values
func NewPostConversationsCallbacksAccepted() *PostConversationsCallbacksAccepted {
	return &PostConversationsCallbacksAccepted{}
}

/*PostConversationsCallbacksAccepted handles this case with default header values.

Accepted - Creating and Processing a Callback
*/
type PostConversationsCallbacksAccepted struct {
}

func (o *PostConversationsCallbacksAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/callbacks][%d] postConversationsCallbacksAccepted ", 202)
}

func (o *PostConversationsCallbacksAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostConversationsCallbacksBadRequest creates a PostConversationsCallbacksBadRequest with default headers values
func NewPostConversationsCallbacksBadRequest() *PostConversationsCallbacksBadRequest {
	return &PostConversationsCallbacksBadRequest{}
}

/*PostConversationsCallbacksBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostConversationsCallbacksBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsCallbacksBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/callbacks][%d] postConversationsCallbacksBadRequest  %+v", 400, o.Payload)
}

func (o *PostConversationsCallbacksBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallbacksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsCallbacksUnauthorized creates a PostConversationsCallbacksUnauthorized with default headers values
func NewPostConversationsCallbacksUnauthorized() *PostConversationsCallbacksUnauthorized {
	return &PostConversationsCallbacksUnauthorized{}
}

/*PostConversationsCallbacksUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostConversationsCallbacksUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsCallbacksUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/callbacks][%d] postConversationsCallbacksUnauthorized  %+v", 401, o.Payload)
}

func (o *PostConversationsCallbacksUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallbacksUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsCallbacksForbidden creates a PostConversationsCallbacksForbidden with default headers values
func NewPostConversationsCallbacksForbidden() *PostConversationsCallbacksForbidden {
	return &PostConversationsCallbacksForbidden{}
}

/*PostConversationsCallbacksForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostConversationsCallbacksForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsCallbacksForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/callbacks][%d] postConversationsCallbacksForbidden  %+v", 403, o.Payload)
}

func (o *PostConversationsCallbacksForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallbacksForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsCallbacksNotFound creates a PostConversationsCallbacksNotFound with default headers values
func NewPostConversationsCallbacksNotFound() *PostConversationsCallbacksNotFound {
	return &PostConversationsCallbacksNotFound{}
}

/*PostConversationsCallbacksNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostConversationsCallbacksNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsCallbacksNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/callbacks][%d] postConversationsCallbacksNotFound  %+v", 404, o.Payload)
}

func (o *PostConversationsCallbacksNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallbacksNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsCallbacksRequestTimeout creates a PostConversationsCallbacksRequestTimeout with default headers values
func NewPostConversationsCallbacksRequestTimeout() *PostConversationsCallbacksRequestTimeout {
	return &PostConversationsCallbacksRequestTimeout{}
}

/*PostConversationsCallbacksRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostConversationsCallbacksRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsCallbacksRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/callbacks][%d] postConversationsCallbacksRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostConversationsCallbacksRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallbacksRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsCallbacksRequestEntityTooLarge creates a PostConversationsCallbacksRequestEntityTooLarge with default headers values
func NewPostConversationsCallbacksRequestEntityTooLarge() *PostConversationsCallbacksRequestEntityTooLarge {
	return &PostConversationsCallbacksRequestEntityTooLarge{}
}

/*PostConversationsCallbacksRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostConversationsCallbacksRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsCallbacksRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/callbacks][%d] postConversationsCallbacksRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostConversationsCallbacksRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallbacksRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsCallbacksUnsupportedMediaType creates a PostConversationsCallbacksUnsupportedMediaType with default headers values
func NewPostConversationsCallbacksUnsupportedMediaType() *PostConversationsCallbacksUnsupportedMediaType {
	return &PostConversationsCallbacksUnsupportedMediaType{}
}

/*PostConversationsCallbacksUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostConversationsCallbacksUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsCallbacksUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/callbacks][%d] postConversationsCallbacksUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostConversationsCallbacksUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallbacksUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsCallbacksTooManyRequests creates a PostConversationsCallbacksTooManyRequests with default headers values
func NewPostConversationsCallbacksTooManyRequests() *PostConversationsCallbacksTooManyRequests {
	return &PostConversationsCallbacksTooManyRequests{}
}

/*PostConversationsCallbacksTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostConversationsCallbacksTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsCallbacksTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/callbacks][%d] postConversationsCallbacksTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostConversationsCallbacksTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallbacksTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsCallbacksInternalServerError creates a PostConversationsCallbacksInternalServerError with default headers values
func NewPostConversationsCallbacksInternalServerError() *PostConversationsCallbacksInternalServerError {
	return &PostConversationsCallbacksInternalServerError{}
}

/*PostConversationsCallbacksInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostConversationsCallbacksInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsCallbacksInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/callbacks][%d] postConversationsCallbacksInternalServerError  %+v", 500, o.Payload)
}

func (o *PostConversationsCallbacksInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallbacksInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsCallbacksServiceUnavailable creates a PostConversationsCallbacksServiceUnavailable with default headers values
func NewPostConversationsCallbacksServiceUnavailable() *PostConversationsCallbacksServiceUnavailable {
	return &PostConversationsCallbacksServiceUnavailable{}
}

/*PostConversationsCallbacksServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostConversationsCallbacksServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsCallbacksServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/callbacks][%d] postConversationsCallbacksServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostConversationsCallbacksServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallbacksServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsCallbacksGatewayTimeout creates a PostConversationsCallbacksGatewayTimeout with default headers values
func NewPostConversationsCallbacksGatewayTimeout() *PostConversationsCallbacksGatewayTimeout {
	return &PostConversationsCallbacksGatewayTimeout{}
}

/*PostConversationsCallbacksGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostConversationsCallbacksGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsCallbacksGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/callbacks][%d] postConversationsCallbacksGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostConversationsCallbacksGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallbacksGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

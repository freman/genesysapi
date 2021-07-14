// Code generated by go-swagger; DO NOT EDIT.

package quality

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostQualityConversationEvaluationsReader is a Reader for the PostQualityConversationEvaluations structure.
type PostQualityConversationEvaluationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostQualityConversationEvaluationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostQualityConversationEvaluationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostQualityConversationEvaluationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostQualityConversationEvaluationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostQualityConversationEvaluationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostQualityConversationEvaluationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostQualityConversationEvaluationsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostQualityConversationEvaluationsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostQualityConversationEvaluationsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostQualityConversationEvaluationsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostQualityConversationEvaluationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostQualityConversationEvaluationsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostQualityConversationEvaluationsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostQualityConversationEvaluationsOK creates a PostQualityConversationEvaluationsOK with default headers values
func NewPostQualityConversationEvaluationsOK() *PostQualityConversationEvaluationsOK {
	return &PostQualityConversationEvaluationsOK{}
}

/*PostQualityConversationEvaluationsOK handles this case with default header values.

successful operation
*/
type PostQualityConversationEvaluationsOK struct {
	Payload *models.Evaluation
}

func (o *PostQualityConversationEvaluationsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/conversations/{conversationId}/evaluations][%d] postQualityConversationEvaluationsOK  %+v", 200, o.Payload)
}

func (o *PostQualityConversationEvaluationsOK) GetPayload() *models.Evaluation {
	return o.Payload
}

func (o *PostQualityConversationEvaluationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Evaluation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityConversationEvaluationsBadRequest creates a PostQualityConversationEvaluationsBadRequest with default headers values
func NewPostQualityConversationEvaluationsBadRequest() *PostQualityConversationEvaluationsBadRequest {
	return &PostQualityConversationEvaluationsBadRequest{}
}

/*PostQualityConversationEvaluationsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostQualityConversationEvaluationsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostQualityConversationEvaluationsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/conversations/{conversationId}/evaluations][%d] postQualityConversationEvaluationsBadRequest  %+v", 400, o.Payload)
}

func (o *PostQualityConversationEvaluationsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityConversationEvaluationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityConversationEvaluationsUnauthorized creates a PostQualityConversationEvaluationsUnauthorized with default headers values
func NewPostQualityConversationEvaluationsUnauthorized() *PostQualityConversationEvaluationsUnauthorized {
	return &PostQualityConversationEvaluationsUnauthorized{}
}

/*PostQualityConversationEvaluationsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostQualityConversationEvaluationsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostQualityConversationEvaluationsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/conversations/{conversationId}/evaluations][%d] postQualityConversationEvaluationsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostQualityConversationEvaluationsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityConversationEvaluationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityConversationEvaluationsForbidden creates a PostQualityConversationEvaluationsForbidden with default headers values
func NewPostQualityConversationEvaluationsForbidden() *PostQualityConversationEvaluationsForbidden {
	return &PostQualityConversationEvaluationsForbidden{}
}

/*PostQualityConversationEvaluationsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostQualityConversationEvaluationsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostQualityConversationEvaluationsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/conversations/{conversationId}/evaluations][%d] postQualityConversationEvaluationsForbidden  %+v", 403, o.Payload)
}

func (o *PostQualityConversationEvaluationsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityConversationEvaluationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityConversationEvaluationsNotFound creates a PostQualityConversationEvaluationsNotFound with default headers values
func NewPostQualityConversationEvaluationsNotFound() *PostQualityConversationEvaluationsNotFound {
	return &PostQualityConversationEvaluationsNotFound{}
}

/*PostQualityConversationEvaluationsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostQualityConversationEvaluationsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostQualityConversationEvaluationsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/conversations/{conversationId}/evaluations][%d] postQualityConversationEvaluationsNotFound  %+v", 404, o.Payload)
}

func (o *PostQualityConversationEvaluationsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityConversationEvaluationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityConversationEvaluationsRequestTimeout creates a PostQualityConversationEvaluationsRequestTimeout with default headers values
func NewPostQualityConversationEvaluationsRequestTimeout() *PostQualityConversationEvaluationsRequestTimeout {
	return &PostQualityConversationEvaluationsRequestTimeout{}
}

/*PostQualityConversationEvaluationsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostQualityConversationEvaluationsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostQualityConversationEvaluationsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/conversations/{conversationId}/evaluations][%d] postQualityConversationEvaluationsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostQualityConversationEvaluationsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityConversationEvaluationsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityConversationEvaluationsRequestEntityTooLarge creates a PostQualityConversationEvaluationsRequestEntityTooLarge with default headers values
func NewPostQualityConversationEvaluationsRequestEntityTooLarge() *PostQualityConversationEvaluationsRequestEntityTooLarge {
	return &PostQualityConversationEvaluationsRequestEntityTooLarge{}
}

/*PostQualityConversationEvaluationsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostQualityConversationEvaluationsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostQualityConversationEvaluationsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/conversations/{conversationId}/evaluations][%d] postQualityConversationEvaluationsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostQualityConversationEvaluationsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityConversationEvaluationsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityConversationEvaluationsUnsupportedMediaType creates a PostQualityConversationEvaluationsUnsupportedMediaType with default headers values
func NewPostQualityConversationEvaluationsUnsupportedMediaType() *PostQualityConversationEvaluationsUnsupportedMediaType {
	return &PostQualityConversationEvaluationsUnsupportedMediaType{}
}

/*PostQualityConversationEvaluationsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostQualityConversationEvaluationsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostQualityConversationEvaluationsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/conversations/{conversationId}/evaluations][%d] postQualityConversationEvaluationsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostQualityConversationEvaluationsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityConversationEvaluationsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityConversationEvaluationsTooManyRequests creates a PostQualityConversationEvaluationsTooManyRequests with default headers values
func NewPostQualityConversationEvaluationsTooManyRequests() *PostQualityConversationEvaluationsTooManyRequests {
	return &PostQualityConversationEvaluationsTooManyRequests{}
}

/*PostQualityConversationEvaluationsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostQualityConversationEvaluationsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostQualityConversationEvaluationsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/conversations/{conversationId}/evaluations][%d] postQualityConversationEvaluationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostQualityConversationEvaluationsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityConversationEvaluationsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityConversationEvaluationsInternalServerError creates a PostQualityConversationEvaluationsInternalServerError with default headers values
func NewPostQualityConversationEvaluationsInternalServerError() *PostQualityConversationEvaluationsInternalServerError {
	return &PostQualityConversationEvaluationsInternalServerError{}
}

/*PostQualityConversationEvaluationsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostQualityConversationEvaluationsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostQualityConversationEvaluationsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/conversations/{conversationId}/evaluations][%d] postQualityConversationEvaluationsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostQualityConversationEvaluationsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityConversationEvaluationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityConversationEvaluationsServiceUnavailable creates a PostQualityConversationEvaluationsServiceUnavailable with default headers values
func NewPostQualityConversationEvaluationsServiceUnavailable() *PostQualityConversationEvaluationsServiceUnavailable {
	return &PostQualityConversationEvaluationsServiceUnavailable{}
}

/*PostQualityConversationEvaluationsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostQualityConversationEvaluationsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostQualityConversationEvaluationsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/conversations/{conversationId}/evaluations][%d] postQualityConversationEvaluationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostQualityConversationEvaluationsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityConversationEvaluationsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityConversationEvaluationsGatewayTimeout creates a PostQualityConversationEvaluationsGatewayTimeout with default headers values
func NewPostQualityConversationEvaluationsGatewayTimeout() *PostQualityConversationEvaluationsGatewayTimeout {
	return &PostQualityConversationEvaluationsGatewayTimeout{}
}

/*PostQualityConversationEvaluationsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostQualityConversationEvaluationsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostQualityConversationEvaluationsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/conversations/{conversationId}/evaluations][%d] postQualityConversationEvaluationsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostQualityConversationEvaluationsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityConversationEvaluationsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// PostConversationParticipantDigitsReader is a Reader for the PostConversationParticipantDigits structure.
type PostConversationParticipantDigitsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostConversationParticipantDigitsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPostConversationParticipantDigitsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostConversationParticipantDigitsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostConversationParticipantDigitsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostConversationParticipantDigitsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostConversationParticipantDigitsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostConversationParticipantDigitsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostConversationParticipantDigitsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostConversationParticipantDigitsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostConversationParticipantDigitsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostConversationParticipantDigitsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostConversationParticipantDigitsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostConversationParticipantDigitsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostConversationParticipantDigitsAccepted creates a PostConversationParticipantDigitsAccepted with default headers values
func NewPostConversationParticipantDigitsAccepted() *PostConversationParticipantDigitsAccepted {
	return &PostConversationParticipantDigitsAccepted{}
}

/*PostConversationParticipantDigitsAccepted handles this case with default header values.

Accepted
*/
type PostConversationParticipantDigitsAccepted struct {
}

func (o *PostConversationParticipantDigitsAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/participants/{participantId}/digits][%d] postConversationParticipantDigitsAccepted ", 202)
}

func (o *PostConversationParticipantDigitsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostConversationParticipantDigitsBadRequest creates a PostConversationParticipantDigitsBadRequest with default headers values
func NewPostConversationParticipantDigitsBadRequest() *PostConversationParticipantDigitsBadRequest {
	return &PostConversationParticipantDigitsBadRequest{}
}

/*PostConversationParticipantDigitsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostConversationParticipantDigitsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostConversationParticipantDigitsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/participants/{participantId}/digits][%d] postConversationParticipantDigitsBadRequest  %+v", 400, o.Payload)
}

func (o *PostConversationParticipantDigitsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationParticipantDigitsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationParticipantDigitsUnauthorized creates a PostConversationParticipantDigitsUnauthorized with default headers values
func NewPostConversationParticipantDigitsUnauthorized() *PostConversationParticipantDigitsUnauthorized {
	return &PostConversationParticipantDigitsUnauthorized{}
}

/*PostConversationParticipantDigitsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostConversationParticipantDigitsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostConversationParticipantDigitsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/participants/{participantId}/digits][%d] postConversationParticipantDigitsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostConversationParticipantDigitsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationParticipantDigitsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationParticipantDigitsForbidden creates a PostConversationParticipantDigitsForbidden with default headers values
func NewPostConversationParticipantDigitsForbidden() *PostConversationParticipantDigitsForbidden {
	return &PostConversationParticipantDigitsForbidden{}
}

/*PostConversationParticipantDigitsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostConversationParticipantDigitsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostConversationParticipantDigitsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/participants/{participantId}/digits][%d] postConversationParticipantDigitsForbidden  %+v", 403, o.Payload)
}

func (o *PostConversationParticipantDigitsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationParticipantDigitsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationParticipantDigitsNotFound creates a PostConversationParticipantDigitsNotFound with default headers values
func NewPostConversationParticipantDigitsNotFound() *PostConversationParticipantDigitsNotFound {
	return &PostConversationParticipantDigitsNotFound{}
}

/*PostConversationParticipantDigitsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostConversationParticipantDigitsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostConversationParticipantDigitsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/participants/{participantId}/digits][%d] postConversationParticipantDigitsNotFound  %+v", 404, o.Payload)
}

func (o *PostConversationParticipantDigitsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationParticipantDigitsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationParticipantDigitsRequestTimeout creates a PostConversationParticipantDigitsRequestTimeout with default headers values
func NewPostConversationParticipantDigitsRequestTimeout() *PostConversationParticipantDigitsRequestTimeout {
	return &PostConversationParticipantDigitsRequestTimeout{}
}

/*PostConversationParticipantDigitsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostConversationParticipantDigitsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostConversationParticipantDigitsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/participants/{participantId}/digits][%d] postConversationParticipantDigitsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostConversationParticipantDigitsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationParticipantDigitsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationParticipantDigitsRequestEntityTooLarge creates a PostConversationParticipantDigitsRequestEntityTooLarge with default headers values
func NewPostConversationParticipantDigitsRequestEntityTooLarge() *PostConversationParticipantDigitsRequestEntityTooLarge {
	return &PostConversationParticipantDigitsRequestEntityTooLarge{}
}

/*PostConversationParticipantDigitsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostConversationParticipantDigitsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostConversationParticipantDigitsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/participants/{participantId}/digits][%d] postConversationParticipantDigitsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostConversationParticipantDigitsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationParticipantDigitsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationParticipantDigitsUnsupportedMediaType creates a PostConversationParticipantDigitsUnsupportedMediaType with default headers values
func NewPostConversationParticipantDigitsUnsupportedMediaType() *PostConversationParticipantDigitsUnsupportedMediaType {
	return &PostConversationParticipantDigitsUnsupportedMediaType{}
}

/*PostConversationParticipantDigitsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostConversationParticipantDigitsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostConversationParticipantDigitsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/participants/{participantId}/digits][%d] postConversationParticipantDigitsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostConversationParticipantDigitsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationParticipantDigitsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationParticipantDigitsTooManyRequests creates a PostConversationParticipantDigitsTooManyRequests with default headers values
func NewPostConversationParticipantDigitsTooManyRequests() *PostConversationParticipantDigitsTooManyRequests {
	return &PostConversationParticipantDigitsTooManyRequests{}
}

/*PostConversationParticipantDigitsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostConversationParticipantDigitsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostConversationParticipantDigitsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/participants/{participantId}/digits][%d] postConversationParticipantDigitsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostConversationParticipantDigitsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationParticipantDigitsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationParticipantDigitsInternalServerError creates a PostConversationParticipantDigitsInternalServerError with default headers values
func NewPostConversationParticipantDigitsInternalServerError() *PostConversationParticipantDigitsInternalServerError {
	return &PostConversationParticipantDigitsInternalServerError{}
}

/*PostConversationParticipantDigitsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostConversationParticipantDigitsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostConversationParticipantDigitsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/participants/{participantId}/digits][%d] postConversationParticipantDigitsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostConversationParticipantDigitsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationParticipantDigitsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationParticipantDigitsServiceUnavailable creates a PostConversationParticipantDigitsServiceUnavailable with default headers values
func NewPostConversationParticipantDigitsServiceUnavailable() *PostConversationParticipantDigitsServiceUnavailable {
	return &PostConversationParticipantDigitsServiceUnavailable{}
}

/*PostConversationParticipantDigitsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostConversationParticipantDigitsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostConversationParticipantDigitsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/participants/{participantId}/digits][%d] postConversationParticipantDigitsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostConversationParticipantDigitsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationParticipantDigitsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationParticipantDigitsGatewayTimeout creates a PostConversationParticipantDigitsGatewayTimeout with default headers values
func NewPostConversationParticipantDigitsGatewayTimeout() *PostConversationParticipantDigitsGatewayTimeout {
	return &PostConversationParticipantDigitsGatewayTimeout{}
}

/*PostConversationParticipantDigitsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostConversationParticipantDigitsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostConversationParticipantDigitsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/participants/{participantId}/digits][%d] postConversationParticipantDigitsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostConversationParticipantDigitsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationParticipantDigitsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

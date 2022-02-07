// Code generated by go-swagger; DO NOT EDIT.

package speech_and_text_analytics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostSpeechandtextanalyticsSentimentfeedbackReader is a Reader for the PostSpeechandtextanalyticsSentimentfeedback structure.
type PostSpeechandtextanalyticsSentimentfeedbackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSpeechandtextanalyticsSentimentfeedbackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostSpeechandtextanalyticsSentimentfeedbackCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostSpeechandtextanalyticsSentimentfeedbackBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostSpeechandtextanalyticsSentimentfeedbackUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostSpeechandtextanalyticsSentimentfeedbackForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostSpeechandtextanalyticsSentimentfeedbackNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostSpeechandtextanalyticsSentimentfeedbackRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostSpeechandtextanalyticsSentimentfeedbackRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostSpeechandtextanalyticsSentimentfeedbackUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPostSpeechandtextanalyticsSentimentfeedbackUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostSpeechandtextanalyticsSentimentfeedbackTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostSpeechandtextanalyticsSentimentfeedbackInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostSpeechandtextanalyticsSentimentfeedbackServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostSpeechandtextanalyticsSentimentfeedbackGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostSpeechandtextanalyticsSentimentfeedbackCreated creates a PostSpeechandtextanalyticsSentimentfeedbackCreated with default headers values
func NewPostSpeechandtextanalyticsSentimentfeedbackCreated() *PostSpeechandtextanalyticsSentimentfeedbackCreated {
	return &PostSpeechandtextanalyticsSentimentfeedbackCreated{}
}

/*PostSpeechandtextanalyticsSentimentfeedbackCreated handles this case with default header values.

Created
*/
type PostSpeechandtextanalyticsSentimentfeedbackCreated struct {
	Payload *models.SentimentFeedback
}

func (o *PostSpeechandtextanalyticsSentimentfeedbackCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/sentimentfeedback][%d] postSpeechandtextanalyticsSentimentfeedbackCreated  %+v", 201, o.Payload)
}

func (o *PostSpeechandtextanalyticsSentimentfeedbackCreated) GetPayload() *models.SentimentFeedback {
	return o.Payload
}

func (o *PostSpeechandtextanalyticsSentimentfeedbackCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SentimentFeedback)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSpeechandtextanalyticsSentimentfeedbackBadRequest creates a PostSpeechandtextanalyticsSentimentfeedbackBadRequest with default headers values
func NewPostSpeechandtextanalyticsSentimentfeedbackBadRequest() *PostSpeechandtextanalyticsSentimentfeedbackBadRequest {
	return &PostSpeechandtextanalyticsSentimentfeedbackBadRequest{}
}

/*PostSpeechandtextanalyticsSentimentfeedbackBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostSpeechandtextanalyticsSentimentfeedbackBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostSpeechandtextanalyticsSentimentfeedbackBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/sentimentfeedback][%d] postSpeechandtextanalyticsSentimentfeedbackBadRequest  %+v", 400, o.Payload)
}

func (o *PostSpeechandtextanalyticsSentimentfeedbackBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSpeechandtextanalyticsSentimentfeedbackBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSpeechandtextanalyticsSentimentfeedbackUnauthorized creates a PostSpeechandtextanalyticsSentimentfeedbackUnauthorized with default headers values
func NewPostSpeechandtextanalyticsSentimentfeedbackUnauthorized() *PostSpeechandtextanalyticsSentimentfeedbackUnauthorized {
	return &PostSpeechandtextanalyticsSentimentfeedbackUnauthorized{}
}

/*PostSpeechandtextanalyticsSentimentfeedbackUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostSpeechandtextanalyticsSentimentfeedbackUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostSpeechandtextanalyticsSentimentfeedbackUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/sentimentfeedback][%d] postSpeechandtextanalyticsSentimentfeedbackUnauthorized  %+v", 401, o.Payload)
}

func (o *PostSpeechandtextanalyticsSentimentfeedbackUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSpeechandtextanalyticsSentimentfeedbackUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSpeechandtextanalyticsSentimentfeedbackForbidden creates a PostSpeechandtextanalyticsSentimentfeedbackForbidden with default headers values
func NewPostSpeechandtextanalyticsSentimentfeedbackForbidden() *PostSpeechandtextanalyticsSentimentfeedbackForbidden {
	return &PostSpeechandtextanalyticsSentimentfeedbackForbidden{}
}

/*PostSpeechandtextanalyticsSentimentfeedbackForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostSpeechandtextanalyticsSentimentfeedbackForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostSpeechandtextanalyticsSentimentfeedbackForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/sentimentfeedback][%d] postSpeechandtextanalyticsSentimentfeedbackForbidden  %+v", 403, o.Payload)
}

func (o *PostSpeechandtextanalyticsSentimentfeedbackForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSpeechandtextanalyticsSentimentfeedbackForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSpeechandtextanalyticsSentimentfeedbackNotFound creates a PostSpeechandtextanalyticsSentimentfeedbackNotFound with default headers values
func NewPostSpeechandtextanalyticsSentimentfeedbackNotFound() *PostSpeechandtextanalyticsSentimentfeedbackNotFound {
	return &PostSpeechandtextanalyticsSentimentfeedbackNotFound{}
}

/*PostSpeechandtextanalyticsSentimentfeedbackNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostSpeechandtextanalyticsSentimentfeedbackNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostSpeechandtextanalyticsSentimentfeedbackNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/sentimentfeedback][%d] postSpeechandtextanalyticsSentimentfeedbackNotFound  %+v", 404, o.Payload)
}

func (o *PostSpeechandtextanalyticsSentimentfeedbackNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSpeechandtextanalyticsSentimentfeedbackNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSpeechandtextanalyticsSentimentfeedbackRequestTimeout creates a PostSpeechandtextanalyticsSentimentfeedbackRequestTimeout with default headers values
func NewPostSpeechandtextanalyticsSentimentfeedbackRequestTimeout() *PostSpeechandtextanalyticsSentimentfeedbackRequestTimeout {
	return &PostSpeechandtextanalyticsSentimentfeedbackRequestTimeout{}
}

/*PostSpeechandtextanalyticsSentimentfeedbackRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostSpeechandtextanalyticsSentimentfeedbackRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostSpeechandtextanalyticsSentimentfeedbackRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/sentimentfeedback][%d] postSpeechandtextanalyticsSentimentfeedbackRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostSpeechandtextanalyticsSentimentfeedbackRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSpeechandtextanalyticsSentimentfeedbackRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSpeechandtextanalyticsSentimentfeedbackRequestEntityTooLarge creates a PostSpeechandtextanalyticsSentimentfeedbackRequestEntityTooLarge with default headers values
func NewPostSpeechandtextanalyticsSentimentfeedbackRequestEntityTooLarge() *PostSpeechandtextanalyticsSentimentfeedbackRequestEntityTooLarge {
	return &PostSpeechandtextanalyticsSentimentfeedbackRequestEntityTooLarge{}
}

/*PostSpeechandtextanalyticsSentimentfeedbackRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostSpeechandtextanalyticsSentimentfeedbackRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostSpeechandtextanalyticsSentimentfeedbackRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/sentimentfeedback][%d] postSpeechandtextanalyticsSentimentfeedbackRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostSpeechandtextanalyticsSentimentfeedbackRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSpeechandtextanalyticsSentimentfeedbackRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSpeechandtextanalyticsSentimentfeedbackUnsupportedMediaType creates a PostSpeechandtextanalyticsSentimentfeedbackUnsupportedMediaType with default headers values
func NewPostSpeechandtextanalyticsSentimentfeedbackUnsupportedMediaType() *PostSpeechandtextanalyticsSentimentfeedbackUnsupportedMediaType {
	return &PostSpeechandtextanalyticsSentimentfeedbackUnsupportedMediaType{}
}

/*PostSpeechandtextanalyticsSentimentfeedbackUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostSpeechandtextanalyticsSentimentfeedbackUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostSpeechandtextanalyticsSentimentfeedbackUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/sentimentfeedback][%d] postSpeechandtextanalyticsSentimentfeedbackUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostSpeechandtextanalyticsSentimentfeedbackUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSpeechandtextanalyticsSentimentfeedbackUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSpeechandtextanalyticsSentimentfeedbackUnprocessableEntity creates a PostSpeechandtextanalyticsSentimentfeedbackUnprocessableEntity with default headers values
func NewPostSpeechandtextanalyticsSentimentfeedbackUnprocessableEntity() *PostSpeechandtextanalyticsSentimentfeedbackUnprocessableEntity {
	return &PostSpeechandtextanalyticsSentimentfeedbackUnprocessableEntity{}
}

/*PostSpeechandtextanalyticsSentimentfeedbackUnprocessableEntity handles this case with default header values.

PostSpeechandtextanalyticsSentimentfeedbackUnprocessableEntity post speechandtextanalytics sentimentfeedback unprocessable entity
*/
type PostSpeechandtextanalyticsSentimentfeedbackUnprocessableEntity struct {
	Payload *models.ErrorBody
}

func (o *PostSpeechandtextanalyticsSentimentfeedbackUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/sentimentfeedback][%d] postSpeechandtextanalyticsSentimentfeedbackUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostSpeechandtextanalyticsSentimentfeedbackUnprocessableEntity) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSpeechandtextanalyticsSentimentfeedbackUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSpeechandtextanalyticsSentimentfeedbackTooManyRequests creates a PostSpeechandtextanalyticsSentimentfeedbackTooManyRequests with default headers values
func NewPostSpeechandtextanalyticsSentimentfeedbackTooManyRequests() *PostSpeechandtextanalyticsSentimentfeedbackTooManyRequests {
	return &PostSpeechandtextanalyticsSentimentfeedbackTooManyRequests{}
}

/*PostSpeechandtextanalyticsSentimentfeedbackTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostSpeechandtextanalyticsSentimentfeedbackTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostSpeechandtextanalyticsSentimentfeedbackTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/sentimentfeedback][%d] postSpeechandtextanalyticsSentimentfeedbackTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostSpeechandtextanalyticsSentimentfeedbackTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSpeechandtextanalyticsSentimentfeedbackTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSpeechandtextanalyticsSentimentfeedbackInternalServerError creates a PostSpeechandtextanalyticsSentimentfeedbackInternalServerError with default headers values
func NewPostSpeechandtextanalyticsSentimentfeedbackInternalServerError() *PostSpeechandtextanalyticsSentimentfeedbackInternalServerError {
	return &PostSpeechandtextanalyticsSentimentfeedbackInternalServerError{}
}

/*PostSpeechandtextanalyticsSentimentfeedbackInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostSpeechandtextanalyticsSentimentfeedbackInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostSpeechandtextanalyticsSentimentfeedbackInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/sentimentfeedback][%d] postSpeechandtextanalyticsSentimentfeedbackInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSpeechandtextanalyticsSentimentfeedbackInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSpeechandtextanalyticsSentimentfeedbackInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSpeechandtextanalyticsSentimentfeedbackServiceUnavailable creates a PostSpeechandtextanalyticsSentimentfeedbackServiceUnavailable with default headers values
func NewPostSpeechandtextanalyticsSentimentfeedbackServiceUnavailable() *PostSpeechandtextanalyticsSentimentfeedbackServiceUnavailable {
	return &PostSpeechandtextanalyticsSentimentfeedbackServiceUnavailable{}
}

/*PostSpeechandtextanalyticsSentimentfeedbackServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostSpeechandtextanalyticsSentimentfeedbackServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostSpeechandtextanalyticsSentimentfeedbackServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/sentimentfeedback][%d] postSpeechandtextanalyticsSentimentfeedbackServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostSpeechandtextanalyticsSentimentfeedbackServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSpeechandtextanalyticsSentimentfeedbackServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSpeechandtextanalyticsSentimentfeedbackGatewayTimeout creates a PostSpeechandtextanalyticsSentimentfeedbackGatewayTimeout with default headers values
func NewPostSpeechandtextanalyticsSentimentfeedbackGatewayTimeout() *PostSpeechandtextanalyticsSentimentfeedbackGatewayTimeout {
	return &PostSpeechandtextanalyticsSentimentfeedbackGatewayTimeout{}
}

/*PostSpeechandtextanalyticsSentimentfeedbackGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostSpeechandtextanalyticsSentimentfeedbackGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostSpeechandtextanalyticsSentimentfeedbackGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/sentimentfeedback][%d] postSpeechandtextanalyticsSentimentfeedbackGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostSpeechandtextanalyticsSentimentfeedbackGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSpeechandtextanalyticsSentimentfeedbackGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
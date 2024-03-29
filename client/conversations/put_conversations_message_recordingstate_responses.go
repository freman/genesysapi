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

// PutConversationsMessageRecordingstateReader is a Reader for the PutConversationsMessageRecordingstate structure.
type PutConversationsMessageRecordingstateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutConversationsMessageRecordingstateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutConversationsMessageRecordingstateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPutConversationsMessageRecordingstateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutConversationsMessageRecordingstateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutConversationsMessageRecordingstateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutConversationsMessageRecordingstateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutConversationsMessageRecordingstateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutConversationsMessageRecordingstateRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutConversationsMessageRecordingstateRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutConversationsMessageRecordingstateUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutConversationsMessageRecordingstateTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutConversationsMessageRecordingstateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutConversationsMessageRecordingstateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutConversationsMessageRecordingstateGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutConversationsMessageRecordingstateOK creates a PutConversationsMessageRecordingstateOK with default headers values
func NewPutConversationsMessageRecordingstateOK() *PutConversationsMessageRecordingstateOK {
	return &PutConversationsMessageRecordingstateOK{}
}

/*
PutConversationsMessageRecordingstateOK describes a response with status code 200, with default header values.

successful operation
*/
type PutConversationsMessageRecordingstateOK struct {
	Payload string
}

// IsSuccess returns true when this put conversations message recordingstate o k response has a 2xx status code
func (o *PutConversationsMessageRecordingstateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put conversations message recordingstate o k response has a 3xx status code
func (o *PutConversationsMessageRecordingstateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations message recordingstate o k response has a 4xx status code
func (o *PutConversationsMessageRecordingstateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put conversations message recordingstate o k response has a 5xx status code
func (o *PutConversationsMessageRecordingstateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversations message recordingstate o k response a status code equal to that given
func (o *PutConversationsMessageRecordingstateOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutConversationsMessageRecordingstateOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/messages/{conversationId}/recordingstate][%d] putConversationsMessageRecordingstateOK  %+v", 200, o.Payload)
}

func (o *PutConversationsMessageRecordingstateOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/messages/{conversationId}/recordingstate][%d] putConversationsMessageRecordingstateOK  %+v", 200, o.Payload)
}

func (o *PutConversationsMessageRecordingstateOK) GetPayload() string {
	return o.Payload
}

func (o *PutConversationsMessageRecordingstateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsMessageRecordingstateAccepted creates a PutConversationsMessageRecordingstateAccepted with default headers values
func NewPutConversationsMessageRecordingstateAccepted() *PutConversationsMessageRecordingstateAccepted {
	return &PutConversationsMessageRecordingstateAccepted{}
}

/*
PutConversationsMessageRecordingstateAccepted describes a response with status code 202, with default header values.

Accepted - when pausing or resuming recordings (Secure Pause)
*/
type PutConversationsMessageRecordingstateAccepted struct {
	Payload string
}

// IsSuccess returns true when this put conversations message recordingstate accepted response has a 2xx status code
func (o *PutConversationsMessageRecordingstateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put conversations message recordingstate accepted response has a 3xx status code
func (o *PutConversationsMessageRecordingstateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations message recordingstate accepted response has a 4xx status code
func (o *PutConversationsMessageRecordingstateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this put conversations message recordingstate accepted response has a 5xx status code
func (o *PutConversationsMessageRecordingstateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversations message recordingstate accepted response a status code equal to that given
func (o *PutConversationsMessageRecordingstateAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *PutConversationsMessageRecordingstateAccepted) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/messages/{conversationId}/recordingstate][%d] putConversationsMessageRecordingstateAccepted  %+v", 202, o.Payload)
}

func (o *PutConversationsMessageRecordingstateAccepted) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/messages/{conversationId}/recordingstate][%d] putConversationsMessageRecordingstateAccepted  %+v", 202, o.Payload)
}

func (o *PutConversationsMessageRecordingstateAccepted) GetPayload() string {
	return o.Payload
}

func (o *PutConversationsMessageRecordingstateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsMessageRecordingstateBadRequest creates a PutConversationsMessageRecordingstateBadRequest with default headers values
func NewPutConversationsMessageRecordingstateBadRequest() *PutConversationsMessageRecordingstateBadRequest {
	return &PutConversationsMessageRecordingstateBadRequest{}
}

/*
PutConversationsMessageRecordingstateBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutConversationsMessageRecordingstateBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations message recordingstate bad request response has a 2xx status code
func (o *PutConversationsMessageRecordingstateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations message recordingstate bad request response has a 3xx status code
func (o *PutConversationsMessageRecordingstateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations message recordingstate bad request response has a 4xx status code
func (o *PutConversationsMessageRecordingstateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversations message recordingstate bad request response has a 5xx status code
func (o *PutConversationsMessageRecordingstateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversations message recordingstate bad request response a status code equal to that given
func (o *PutConversationsMessageRecordingstateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutConversationsMessageRecordingstateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/messages/{conversationId}/recordingstate][%d] putConversationsMessageRecordingstateBadRequest  %+v", 400, o.Payload)
}

func (o *PutConversationsMessageRecordingstateBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/messages/{conversationId}/recordingstate][%d] putConversationsMessageRecordingstateBadRequest  %+v", 400, o.Payload)
}

func (o *PutConversationsMessageRecordingstateBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsMessageRecordingstateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsMessageRecordingstateUnauthorized creates a PutConversationsMessageRecordingstateUnauthorized with default headers values
func NewPutConversationsMessageRecordingstateUnauthorized() *PutConversationsMessageRecordingstateUnauthorized {
	return &PutConversationsMessageRecordingstateUnauthorized{}
}

/*
PutConversationsMessageRecordingstateUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PutConversationsMessageRecordingstateUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations message recordingstate unauthorized response has a 2xx status code
func (o *PutConversationsMessageRecordingstateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations message recordingstate unauthorized response has a 3xx status code
func (o *PutConversationsMessageRecordingstateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations message recordingstate unauthorized response has a 4xx status code
func (o *PutConversationsMessageRecordingstateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversations message recordingstate unauthorized response has a 5xx status code
func (o *PutConversationsMessageRecordingstateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversations message recordingstate unauthorized response a status code equal to that given
func (o *PutConversationsMessageRecordingstateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PutConversationsMessageRecordingstateUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/messages/{conversationId}/recordingstate][%d] putConversationsMessageRecordingstateUnauthorized  %+v", 401, o.Payload)
}

func (o *PutConversationsMessageRecordingstateUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/messages/{conversationId}/recordingstate][%d] putConversationsMessageRecordingstateUnauthorized  %+v", 401, o.Payload)
}

func (o *PutConversationsMessageRecordingstateUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsMessageRecordingstateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsMessageRecordingstateForbidden creates a PutConversationsMessageRecordingstateForbidden with default headers values
func NewPutConversationsMessageRecordingstateForbidden() *PutConversationsMessageRecordingstateForbidden {
	return &PutConversationsMessageRecordingstateForbidden{}
}

/*
PutConversationsMessageRecordingstateForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PutConversationsMessageRecordingstateForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations message recordingstate forbidden response has a 2xx status code
func (o *PutConversationsMessageRecordingstateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations message recordingstate forbidden response has a 3xx status code
func (o *PutConversationsMessageRecordingstateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations message recordingstate forbidden response has a 4xx status code
func (o *PutConversationsMessageRecordingstateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversations message recordingstate forbidden response has a 5xx status code
func (o *PutConversationsMessageRecordingstateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversations message recordingstate forbidden response a status code equal to that given
func (o *PutConversationsMessageRecordingstateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutConversationsMessageRecordingstateForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/messages/{conversationId}/recordingstate][%d] putConversationsMessageRecordingstateForbidden  %+v", 403, o.Payload)
}

func (o *PutConversationsMessageRecordingstateForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/messages/{conversationId}/recordingstate][%d] putConversationsMessageRecordingstateForbidden  %+v", 403, o.Payload)
}

func (o *PutConversationsMessageRecordingstateForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsMessageRecordingstateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsMessageRecordingstateNotFound creates a PutConversationsMessageRecordingstateNotFound with default headers values
func NewPutConversationsMessageRecordingstateNotFound() *PutConversationsMessageRecordingstateNotFound {
	return &PutConversationsMessageRecordingstateNotFound{}
}

/*
PutConversationsMessageRecordingstateNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PutConversationsMessageRecordingstateNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations message recordingstate not found response has a 2xx status code
func (o *PutConversationsMessageRecordingstateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations message recordingstate not found response has a 3xx status code
func (o *PutConversationsMessageRecordingstateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations message recordingstate not found response has a 4xx status code
func (o *PutConversationsMessageRecordingstateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversations message recordingstate not found response has a 5xx status code
func (o *PutConversationsMessageRecordingstateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversations message recordingstate not found response a status code equal to that given
func (o *PutConversationsMessageRecordingstateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutConversationsMessageRecordingstateNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/messages/{conversationId}/recordingstate][%d] putConversationsMessageRecordingstateNotFound  %+v", 404, o.Payload)
}

func (o *PutConversationsMessageRecordingstateNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/messages/{conversationId}/recordingstate][%d] putConversationsMessageRecordingstateNotFound  %+v", 404, o.Payload)
}

func (o *PutConversationsMessageRecordingstateNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsMessageRecordingstateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsMessageRecordingstateRequestTimeout creates a PutConversationsMessageRecordingstateRequestTimeout with default headers values
func NewPutConversationsMessageRecordingstateRequestTimeout() *PutConversationsMessageRecordingstateRequestTimeout {
	return &PutConversationsMessageRecordingstateRequestTimeout{}
}

/*
PutConversationsMessageRecordingstateRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutConversationsMessageRecordingstateRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations message recordingstate request timeout response has a 2xx status code
func (o *PutConversationsMessageRecordingstateRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations message recordingstate request timeout response has a 3xx status code
func (o *PutConversationsMessageRecordingstateRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations message recordingstate request timeout response has a 4xx status code
func (o *PutConversationsMessageRecordingstateRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversations message recordingstate request timeout response has a 5xx status code
func (o *PutConversationsMessageRecordingstateRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversations message recordingstate request timeout response a status code equal to that given
func (o *PutConversationsMessageRecordingstateRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PutConversationsMessageRecordingstateRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/messages/{conversationId}/recordingstate][%d] putConversationsMessageRecordingstateRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutConversationsMessageRecordingstateRequestTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/messages/{conversationId}/recordingstate][%d] putConversationsMessageRecordingstateRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutConversationsMessageRecordingstateRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsMessageRecordingstateRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsMessageRecordingstateRequestEntityTooLarge creates a PutConversationsMessageRecordingstateRequestEntityTooLarge with default headers values
func NewPutConversationsMessageRecordingstateRequestEntityTooLarge() *PutConversationsMessageRecordingstateRequestEntityTooLarge {
	return &PutConversationsMessageRecordingstateRequestEntityTooLarge{}
}

/*
PutConversationsMessageRecordingstateRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PutConversationsMessageRecordingstateRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations message recordingstate request entity too large response has a 2xx status code
func (o *PutConversationsMessageRecordingstateRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations message recordingstate request entity too large response has a 3xx status code
func (o *PutConversationsMessageRecordingstateRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations message recordingstate request entity too large response has a 4xx status code
func (o *PutConversationsMessageRecordingstateRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversations message recordingstate request entity too large response has a 5xx status code
func (o *PutConversationsMessageRecordingstateRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversations message recordingstate request entity too large response a status code equal to that given
func (o *PutConversationsMessageRecordingstateRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PutConversationsMessageRecordingstateRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/messages/{conversationId}/recordingstate][%d] putConversationsMessageRecordingstateRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutConversationsMessageRecordingstateRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/messages/{conversationId}/recordingstate][%d] putConversationsMessageRecordingstateRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutConversationsMessageRecordingstateRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsMessageRecordingstateRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsMessageRecordingstateUnsupportedMediaType creates a PutConversationsMessageRecordingstateUnsupportedMediaType with default headers values
func NewPutConversationsMessageRecordingstateUnsupportedMediaType() *PutConversationsMessageRecordingstateUnsupportedMediaType {
	return &PutConversationsMessageRecordingstateUnsupportedMediaType{}
}

/*
PutConversationsMessageRecordingstateUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutConversationsMessageRecordingstateUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations message recordingstate unsupported media type response has a 2xx status code
func (o *PutConversationsMessageRecordingstateUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations message recordingstate unsupported media type response has a 3xx status code
func (o *PutConversationsMessageRecordingstateUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations message recordingstate unsupported media type response has a 4xx status code
func (o *PutConversationsMessageRecordingstateUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversations message recordingstate unsupported media type response has a 5xx status code
func (o *PutConversationsMessageRecordingstateUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversations message recordingstate unsupported media type response a status code equal to that given
func (o *PutConversationsMessageRecordingstateUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PutConversationsMessageRecordingstateUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/messages/{conversationId}/recordingstate][%d] putConversationsMessageRecordingstateUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutConversationsMessageRecordingstateUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/messages/{conversationId}/recordingstate][%d] putConversationsMessageRecordingstateUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutConversationsMessageRecordingstateUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsMessageRecordingstateUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsMessageRecordingstateTooManyRequests creates a PutConversationsMessageRecordingstateTooManyRequests with default headers values
func NewPutConversationsMessageRecordingstateTooManyRequests() *PutConversationsMessageRecordingstateTooManyRequests {
	return &PutConversationsMessageRecordingstateTooManyRequests{}
}

/*
PutConversationsMessageRecordingstateTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutConversationsMessageRecordingstateTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations message recordingstate too many requests response has a 2xx status code
func (o *PutConversationsMessageRecordingstateTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations message recordingstate too many requests response has a 3xx status code
func (o *PutConversationsMessageRecordingstateTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations message recordingstate too many requests response has a 4xx status code
func (o *PutConversationsMessageRecordingstateTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversations message recordingstate too many requests response has a 5xx status code
func (o *PutConversationsMessageRecordingstateTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversations message recordingstate too many requests response a status code equal to that given
func (o *PutConversationsMessageRecordingstateTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PutConversationsMessageRecordingstateTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/messages/{conversationId}/recordingstate][%d] putConversationsMessageRecordingstateTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutConversationsMessageRecordingstateTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/messages/{conversationId}/recordingstate][%d] putConversationsMessageRecordingstateTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutConversationsMessageRecordingstateTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsMessageRecordingstateTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsMessageRecordingstateInternalServerError creates a PutConversationsMessageRecordingstateInternalServerError with default headers values
func NewPutConversationsMessageRecordingstateInternalServerError() *PutConversationsMessageRecordingstateInternalServerError {
	return &PutConversationsMessageRecordingstateInternalServerError{}
}

/*
PutConversationsMessageRecordingstateInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutConversationsMessageRecordingstateInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations message recordingstate internal server error response has a 2xx status code
func (o *PutConversationsMessageRecordingstateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations message recordingstate internal server error response has a 3xx status code
func (o *PutConversationsMessageRecordingstateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations message recordingstate internal server error response has a 4xx status code
func (o *PutConversationsMessageRecordingstateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put conversations message recordingstate internal server error response has a 5xx status code
func (o *PutConversationsMessageRecordingstateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put conversations message recordingstate internal server error response a status code equal to that given
func (o *PutConversationsMessageRecordingstateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutConversationsMessageRecordingstateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/messages/{conversationId}/recordingstate][%d] putConversationsMessageRecordingstateInternalServerError  %+v", 500, o.Payload)
}

func (o *PutConversationsMessageRecordingstateInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/messages/{conversationId}/recordingstate][%d] putConversationsMessageRecordingstateInternalServerError  %+v", 500, o.Payload)
}

func (o *PutConversationsMessageRecordingstateInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsMessageRecordingstateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsMessageRecordingstateServiceUnavailable creates a PutConversationsMessageRecordingstateServiceUnavailable with default headers values
func NewPutConversationsMessageRecordingstateServiceUnavailable() *PutConversationsMessageRecordingstateServiceUnavailable {
	return &PutConversationsMessageRecordingstateServiceUnavailable{}
}

/*
PutConversationsMessageRecordingstateServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutConversationsMessageRecordingstateServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations message recordingstate service unavailable response has a 2xx status code
func (o *PutConversationsMessageRecordingstateServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations message recordingstate service unavailable response has a 3xx status code
func (o *PutConversationsMessageRecordingstateServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations message recordingstate service unavailable response has a 4xx status code
func (o *PutConversationsMessageRecordingstateServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this put conversations message recordingstate service unavailable response has a 5xx status code
func (o *PutConversationsMessageRecordingstateServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this put conversations message recordingstate service unavailable response a status code equal to that given
func (o *PutConversationsMessageRecordingstateServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PutConversationsMessageRecordingstateServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/messages/{conversationId}/recordingstate][%d] putConversationsMessageRecordingstateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutConversationsMessageRecordingstateServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/messages/{conversationId}/recordingstate][%d] putConversationsMessageRecordingstateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutConversationsMessageRecordingstateServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsMessageRecordingstateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsMessageRecordingstateGatewayTimeout creates a PutConversationsMessageRecordingstateGatewayTimeout with default headers values
func NewPutConversationsMessageRecordingstateGatewayTimeout() *PutConversationsMessageRecordingstateGatewayTimeout {
	return &PutConversationsMessageRecordingstateGatewayTimeout{}
}

/*
PutConversationsMessageRecordingstateGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PutConversationsMessageRecordingstateGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations message recordingstate gateway timeout response has a 2xx status code
func (o *PutConversationsMessageRecordingstateGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations message recordingstate gateway timeout response has a 3xx status code
func (o *PutConversationsMessageRecordingstateGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations message recordingstate gateway timeout response has a 4xx status code
func (o *PutConversationsMessageRecordingstateGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this put conversations message recordingstate gateway timeout response has a 5xx status code
func (o *PutConversationsMessageRecordingstateGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this put conversations message recordingstate gateway timeout response a status code equal to that given
func (o *PutConversationsMessageRecordingstateGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PutConversationsMessageRecordingstateGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/messages/{conversationId}/recordingstate][%d] putConversationsMessageRecordingstateGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutConversationsMessageRecordingstateGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/messages/{conversationId}/recordingstate][%d] putConversationsMessageRecordingstateGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutConversationsMessageRecordingstateGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsMessageRecordingstateGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

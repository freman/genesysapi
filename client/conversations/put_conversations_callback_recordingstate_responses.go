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

// PutConversationsCallbackRecordingstateReader is a Reader for the PutConversationsCallbackRecordingstate structure.
type PutConversationsCallbackRecordingstateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutConversationsCallbackRecordingstateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutConversationsCallbackRecordingstateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPutConversationsCallbackRecordingstateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutConversationsCallbackRecordingstateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutConversationsCallbackRecordingstateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutConversationsCallbackRecordingstateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutConversationsCallbackRecordingstateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutConversationsCallbackRecordingstateRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutConversationsCallbackRecordingstateRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutConversationsCallbackRecordingstateUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutConversationsCallbackRecordingstateTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutConversationsCallbackRecordingstateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutConversationsCallbackRecordingstateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutConversationsCallbackRecordingstateGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutConversationsCallbackRecordingstateOK creates a PutConversationsCallbackRecordingstateOK with default headers values
func NewPutConversationsCallbackRecordingstateOK() *PutConversationsCallbackRecordingstateOK {
	return &PutConversationsCallbackRecordingstateOK{}
}

/*
PutConversationsCallbackRecordingstateOK describes a response with status code 200, with default header values.

successful operation
*/
type PutConversationsCallbackRecordingstateOK struct {
	Payload string
}

// IsSuccess returns true when this put conversations callback recordingstate o k response has a 2xx status code
func (o *PutConversationsCallbackRecordingstateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put conversations callback recordingstate o k response has a 3xx status code
func (o *PutConversationsCallbackRecordingstateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations callback recordingstate o k response has a 4xx status code
func (o *PutConversationsCallbackRecordingstateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put conversations callback recordingstate o k response has a 5xx status code
func (o *PutConversationsCallbackRecordingstateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversations callback recordingstate o k response a status code equal to that given
func (o *PutConversationsCallbackRecordingstateOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutConversationsCallbackRecordingstateOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/callbacks/{conversationId}/recordingstate][%d] putConversationsCallbackRecordingstateOK  %+v", 200, o.Payload)
}

func (o *PutConversationsCallbackRecordingstateOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/callbacks/{conversationId}/recordingstate][%d] putConversationsCallbackRecordingstateOK  %+v", 200, o.Payload)
}

func (o *PutConversationsCallbackRecordingstateOK) GetPayload() string {
	return o.Payload
}

func (o *PutConversationsCallbackRecordingstateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsCallbackRecordingstateAccepted creates a PutConversationsCallbackRecordingstateAccepted with default headers values
func NewPutConversationsCallbackRecordingstateAccepted() *PutConversationsCallbackRecordingstateAccepted {
	return &PutConversationsCallbackRecordingstateAccepted{}
}

/*
PutConversationsCallbackRecordingstateAccepted describes a response with status code 202, with default header values.

Accepted - when pausing or resuming recordings (Secure Pause)
*/
type PutConversationsCallbackRecordingstateAccepted struct {
	Payload string
}

// IsSuccess returns true when this put conversations callback recordingstate accepted response has a 2xx status code
func (o *PutConversationsCallbackRecordingstateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put conversations callback recordingstate accepted response has a 3xx status code
func (o *PutConversationsCallbackRecordingstateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations callback recordingstate accepted response has a 4xx status code
func (o *PutConversationsCallbackRecordingstateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this put conversations callback recordingstate accepted response has a 5xx status code
func (o *PutConversationsCallbackRecordingstateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversations callback recordingstate accepted response a status code equal to that given
func (o *PutConversationsCallbackRecordingstateAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *PutConversationsCallbackRecordingstateAccepted) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/callbacks/{conversationId}/recordingstate][%d] putConversationsCallbackRecordingstateAccepted  %+v", 202, o.Payload)
}

func (o *PutConversationsCallbackRecordingstateAccepted) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/callbacks/{conversationId}/recordingstate][%d] putConversationsCallbackRecordingstateAccepted  %+v", 202, o.Payload)
}

func (o *PutConversationsCallbackRecordingstateAccepted) GetPayload() string {
	return o.Payload
}

func (o *PutConversationsCallbackRecordingstateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsCallbackRecordingstateBadRequest creates a PutConversationsCallbackRecordingstateBadRequest with default headers values
func NewPutConversationsCallbackRecordingstateBadRequest() *PutConversationsCallbackRecordingstateBadRequest {
	return &PutConversationsCallbackRecordingstateBadRequest{}
}

/*
PutConversationsCallbackRecordingstateBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutConversationsCallbackRecordingstateBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations callback recordingstate bad request response has a 2xx status code
func (o *PutConversationsCallbackRecordingstateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations callback recordingstate bad request response has a 3xx status code
func (o *PutConversationsCallbackRecordingstateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations callback recordingstate bad request response has a 4xx status code
func (o *PutConversationsCallbackRecordingstateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversations callback recordingstate bad request response has a 5xx status code
func (o *PutConversationsCallbackRecordingstateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversations callback recordingstate bad request response a status code equal to that given
func (o *PutConversationsCallbackRecordingstateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutConversationsCallbackRecordingstateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/callbacks/{conversationId}/recordingstate][%d] putConversationsCallbackRecordingstateBadRequest  %+v", 400, o.Payload)
}

func (o *PutConversationsCallbackRecordingstateBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/callbacks/{conversationId}/recordingstate][%d] putConversationsCallbackRecordingstateBadRequest  %+v", 400, o.Payload)
}

func (o *PutConversationsCallbackRecordingstateBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsCallbackRecordingstateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsCallbackRecordingstateUnauthorized creates a PutConversationsCallbackRecordingstateUnauthorized with default headers values
func NewPutConversationsCallbackRecordingstateUnauthorized() *PutConversationsCallbackRecordingstateUnauthorized {
	return &PutConversationsCallbackRecordingstateUnauthorized{}
}

/*
PutConversationsCallbackRecordingstateUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PutConversationsCallbackRecordingstateUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations callback recordingstate unauthorized response has a 2xx status code
func (o *PutConversationsCallbackRecordingstateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations callback recordingstate unauthorized response has a 3xx status code
func (o *PutConversationsCallbackRecordingstateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations callback recordingstate unauthorized response has a 4xx status code
func (o *PutConversationsCallbackRecordingstateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversations callback recordingstate unauthorized response has a 5xx status code
func (o *PutConversationsCallbackRecordingstateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversations callback recordingstate unauthorized response a status code equal to that given
func (o *PutConversationsCallbackRecordingstateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PutConversationsCallbackRecordingstateUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/callbacks/{conversationId}/recordingstate][%d] putConversationsCallbackRecordingstateUnauthorized  %+v", 401, o.Payload)
}

func (o *PutConversationsCallbackRecordingstateUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/callbacks/{conversationId}/recordingstate][%d] putConversationsCallbackRecordingstateUnauthorized  %+v", 401, o.Payload)
}

func (o *PutConversationsCallbackRecordingstateUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsCallbackRecordingstateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsCallbackRecordingstateForbidden creates a PutConversationsCallbackRecordingstateForbidden with default headers values
func NewPutConversationsCallbackRecordingstateForbidden() *PutConversationsCallbackRecordingstateForbidden {
	return &PutConversationsCallbackRecordingstateForbidden{}
}

/*
PutConversationsCallbackRecordingstateForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PutConversationsCallbackRecordingstateForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations callback recordingstate forbidden response has a 2xx status code
func (o *PutConversationsCallbackRecordingstateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations callback recordingstate forbidden response has a 3xx status code
func (o *PutConversationsCallbackRecordingstateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations callback recordingstate forbidden response has a 4xx status code
func (o *PutConversationsCallbackRecordingstateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversations callback recordingstate forbidden response has a 5xx status code
func (o *PutConversationsCallbackRecordingstateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversations callback recordingstate forbidden response a status code equal to that given
func (o *PutConversationsCallbackRecordingstateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutConversationsCallbackRecordingstateForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/callbacks/{conversationId}/recordingstate][%d] putConversationsCallbackRecordingstateForbidden  %+v", 403, o.Payload)
}

func (o *PutConversationsCallbackRecordingstateForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/callbacks/{conversationId}/recordingstate][%d] putConversationsCallbackRecordingstateForbidden  %+v", 403, o.Payload)
}

func (o *PutConversationsCallbackRecordingstateForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsCallbackRecordingstateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsCallbackRecordingstateNotFound creates a PutConversationsCallbackRecordingstateNotFound with default headers values
func NewPutConversationsCallbackRecordingstateNotFound() *PutConversationsCallbackRecordingstateNotFound {
	return &PutConversationsCallbackRecordingstateNotFound{}
}

/*
PutConversationsCallbackRecordingstateNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PutConversationsCallbackRecordingstateNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations callback recordingstate not found response has a 2xx status code
func (o *PutConversationsCallbackRecordingstateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations callback recordingstate not found response has a 3xx status code
func (o *PutConversationsCallbackRecordingstateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations callback recordingstate not found response has a 4xx status code
func (o *PutConversationsCallbackRecordingstateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversations callback recordingstate not found response has a 5xx status code
func (o *PutConversationsCallbackRecordingstateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversations callback recordingstate not found response a status code equal to that given
func (o *PutConversationsCallbackRecordingstateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutConversationsCallbackRecordingstateNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/callbacks/{conversationId}/recordingstate][%d] putConversationsCallbackRecordingstateNotFound  %+v", 404, o.Payload)
}

func (o *PutConversationsCallbackRecordingstateNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/callbacks/{conversationId}/recordingstate][%d] putConversationsCallbackRecordingstateNotFound  %+v", 404, o.Payload)
}

func (o *PutConversationsCallbackRecordingstateNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsCallbackRecordingstateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsCallbackRecordingstateRequestTimeout creates a PutConversationsCallbackRecordingstateRequestTimeout with default headers values
func NewPutConversationsCallbackRecordingstateRequestTimeout() *PutConversationsCallbackRecordingstateRequestTimeout {
	return &PutConversationsCallbackRecordingstateRequestTimeout{}
}

/*
PutConversationsCallbackRecordingstateRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutConversationsCallbackRecordingstateRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations callback recordingstate request timeout response has a 2xx status code
func (o *PutConversationsCallbackRecordingstateRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations callback recordingstate request timeout response has a 3xx status code
func (o *PutConversationsCallbackRecordingstateRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations callback recordingstate request timeout response has a 4xx status code
func (o *PutConversationsCallbackRecordingstateRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversations callback recordingstate request timeout response has a 5xx status code
func (o *PutConversationsCallbackRecordingstateRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversations callback recordingstate request timeout response a status code equal to that given
func (o *PutConversationsCallbackRecordingstateRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PutConversationsCallbackRecordingstateRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/callbacks/{conversationId}/recordingstate][%d] putConversationsCallbackRecordingstateRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutConversationsCallbackRecordingstateRequestTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/callbacks/{conversationId}/recordingstate][%d] putConversationsCallbackRecordingstateRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutConversationsCallbackRecordingstateRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsCallbackRecordingstateRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsCallbackRecordingstateRequestEntityTooLarge creates a PutConversationsCallbackRecordingstateRequestEntityTooLarge with default headers values
func NewPutConversationsCallbackRecordingstateRequestEntityTooLarge() *PutConversationsCallbackRecordingstateRequestEntityTooLarge {
	return &PutConversationsCallbackRecordingstateRequestEntityTooLarge{}
}

/*
PutConversationsCallbackRecordingstateRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PutConversationsCallbackRecordingstateRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations callback recordingstate request entity too large response has a 2xx status code
func (o *PutConversationsCallbackRecordingstateRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations callback recordingstate request entity too large response has a 3xx status code
func (o *PutConversationsCallbackRecordingstateRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations callback recordingstate request entity too large response has a 4xx status code
func (o *PutConversationsCallbackRecordingstateRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversations callback recordingstate request entity too large response has a 5xx status code
func (o *PutConversationsCallbackRecordingstateRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversations callback recordingstate request entity too large response a status code equal to that given
func (o *PutConversationsCallbackRecordingstateRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PutConversationsCallbackRecordingstateRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/callbacks/{conversationId}/recordingstate][%d] putConversationsCallbackRecordingstateRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutConversationsCallbackRecordingstateRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/callbacks/{conversationId}/recordingstate][%d] putConversationsCallbackRecordingstateRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutConversationsCallbackRecordingstateRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsCallbackRecordingstateRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsCallbackRecordingstateUnsupportedMediaType creates a PutConversationsCallbackRecordingstateUnsupportedMediaType with default headers values
func NewPutConversationsCallbackRecordingstateUnsupportedMediaType() *PutConversationsCallbackRecordingstateUnsupportedMediaType {
	return &PutConversationsCallbackRecordingstateUnsupportedMediaType{}
}

/*
PutConversationsCallbackRecordingstateUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutConversationsCallbackRecordingstateUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations callback recordingstate unsupported media type response has a 2xx status code
func (o *PutConversationsCallbackRecordingstateUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations callback recordingstate unsupported media type response has a 3xx status code
func (o *PutConversationsCallbackRecordingstateUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations callback recordingstate unsupported media type response has a 4xx status code
func (o *PutConversationsCallbackRecordingstateUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversations callback recordingstate unsupported media type response has a 5xx status code
func (o *PutConversationsCallbackRecordingstateUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversations callback recordingstate unsupported media type response a status code equal to that given
func (o *PutConversationsCallbackRecordingstateUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PutConversationsCallbackRecordingstateUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/callbacks/{conversationId}/recordingstate][%d] putConversationsCallbackRecordingstateUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutConversationsCallbackRecordingstateUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/callbacks/{conversationId}/recordingstate][%d] putConversationsCallbackRecordingstateUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutConversationsCallbackRecordingstateUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsCallbackRecordingstateUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsCallbackRecordingstateTooManyRequests creates a PutConversationsCallbackRecordingstateTooManyRequests with default headers values
func NewPutConversationsCallbackRecordingstateTooManyRequests() *PutConversationsCallbackRecordingstateTooManyRequests {
	return &PutConversationsCallbackRecordingstateTooManyRequests{}
}

/*
PutConversationsCallbackRecordingstateTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutConversationsCallbackRecordingstateTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations callback recordingstate too many requests response has a 2xx status code
func (o *PutConversationsCallbackRecordingstateTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations callback recordingstate too many requests response has a 3xx status code
func (o *PutConversationsCallbackRecordingstateTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations callback recordingstate too many requests response has a 4xx status code
func (o *PutConversationsCallbackRecordingstateTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversations callback recordingstate too many requests response has a 5xx status code
func (o *PutConversationsCallbackRecordingstateTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversations callback recordingstate too many requests response a status code equal to that given
func (o *PutConversationsCallbackRecordingstateTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PutConversationsCallbackRecordingstateTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/callbacks/{conversationId}/recordingstate][%d] putConversationsCallbackRecordingstateTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutConversationsCallbackRecordingstateTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/callbacks/{conversationId}/recordingstate][%d] putConversationsCallbackRecordingstateTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutConversationsCallbackRecordingstateTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsCallbackRecordingstateTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsCallbackRecordingstateInternalServerError creates a PutConversationsCallbackRecordingstateInternalServerError with default headers values
func NewPutConversationsCallbackRecordingstateInternalServerError() *PutConversationsCallbackRecordingstateInternalServerError {
	return &PutConversationsCallbackRecordingstateInternalServerError{}
}

/*
PutConversationsCallbackRecordingstateInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutConversationsCallbackRecordingstateInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations callback recordingstate internal server error response has a 2xx status code
func (o *PutConversationsCallbackRecordingstateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations callback recordingstate internal server error response has a 3xx status code
func (o *PutConversationsCallbackRecordingstateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations callback recordingstate internal server error response has a 4xx status code
func (o *PutConversationsCallbackRecordingstateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put conversations callback recordingstate internal server error response has a 5xx status code
func (o *PutConversationsCallbackRecordingstateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put conversations callback recordingstate internal server error response a status code equal to that given
func (o *PutConversationsCallbackRecordingstateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutConversationsCallbackRecordingstateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/callbacks/{conversationId}/recordingstate][%d] putConversationsCallbackRecordingstateInternalServerError  %+v", 500, o.Payload)
}

func (o *PutConversationsCallbackRecordingstateInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/callbacks/{conversationId}/recordingstate][%d] putConversationsCallbackRecordingstateInternalServerError  %+v", 500, o.Payload)
}

func (o *PutConversationsCallbackRecordingstateInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsCallbackRecordingstateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsCallbackRecordingstateServiceUnavailable creates a PutConversationsCallbackRecordingstateServiceUnavailable with default headers values
func NewPutConversationsCallbackRecordingstateServiceUnavailable() *PutConversationsCallbackRecordingstateServiceUnavailable {
	return &PutConversationsCallbackRecordingstateServiceUnavailable{}
}

/*
PutConversationsCallbackRecordingstateServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutConversationsCallbackRecordingstateServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations callback recordingstate service unavailable response has a 2xx status code
func (o *PutConversationsCallbackRecordingstateServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations callback recordingstate service unavailable response has a 3xx status code
func (o *PutConversationsCallbackRecordingstateServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations callback recordingstate service unavailable response has a 4xx status code
func (o *PutConversationsCallbackRecordingstateServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this put conversations callback recordingstate service unavailable response has a 5xx status code
func (o *PutConversationsCallbackRecordingstateServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this put conversations callback recordingstate service unavailable response a status code equal to that given
func (o *PutConversationsCallbackRecordingstateServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PutConversationsCallbackRecordingstateServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/callbacks/{conversationId}/recordingstate][%d] putConversationsCallbackRecordingstateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutConversationsCallbackRecordingstateServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/callbacks/{conversationId}/recordingstate][%d] putConversationsCallbackRecordingstateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutConversationsCallbackRecordingstateServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsCallbackRecordingstateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsCallbackRecordingstateGatewayTimeout creates a PutConversationsCallbackRecordingstateGatewayTimeout with default headers values
func NewPutConversationsCallbackRecordingstateGatewayTimeout() *PutConversationsCallbackRecordingstateGatewayTimeout {
	return &PutConversationsCallbackRecordingstateGatewayTimeout{}
}

/*
PutConversationsCallbackRecordingstateGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PutConversationsCallbackRecordingstateGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations callback recordingstate gateway timeout response has a 2xx status code
func (o *PutConversationsCallbackRecordingstateGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations callback recordingstate gateway timeout response has a 3xx status code
func (o *PutConversationsCallbackRecordingstateGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations callback recordingstate gateway timeout response has a 4xx status code
func (o *PutConversationsCallbackRecordingstateGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this put conversations callback recordingstate gateway timeout response has a 5xx status code
func (o *PutConversationsCallbackRecordingstateGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this put conversations callback recordingstate gateway timeout response a status code equal to that given
func (o *PutConversationsCallbackRecordingstateGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PutConversationsCallbackRecordingstateGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/callbacks/{conversationId}/recordingstate][%d] putConversationsCallbackRecordingstateGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutConversationsCallbackRecordingstateGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/callbacks/{conversationId}/recordingstate][%d] putConversationsCallbackRecordingstateGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutConversationsCallbackRecordingstateGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsCallbackRecordingstateGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

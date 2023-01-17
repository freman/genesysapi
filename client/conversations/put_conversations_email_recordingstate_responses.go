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

// PutConversationsEmailRecordingstateReader is a Reader for the PutConversationsEmailRecordingstate structure.
type PutConversationsEmailRecordingstateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutConversationsEmailRecordingstateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutConversationsEmailRecordingstateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPutConversationsEmailRecordingstateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutConversationsEmailRecordingstateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutConversationsEmailRecordingstateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutConversationsEmailRecordingstateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutConversationsEmailRecordingstateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutConversationsEmailRecordingstateRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutConversationsEmailRecordingstateRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutConversationsEmailRecordingstateUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutConversationsEmailRecordingstateTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutConversationsEmailRecordingstateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutConversationsEmailRecordingstateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutConversationsEmailRecordingstateGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutConversationsEmailRecordingstateOK creates a PutConversationsEmailRecordingstateOK with default headers values
func NewPutConversationsEmailRecordingstateOK() *PutConversationsEmailRecordingstateOK {
	return &PutConversationsEmailRecordingstateOK{}
}

/*
PutConversationsEmailRecordingstateOK describes a response with status code 200, with default header values.

successful operation
*/
type PutConversationsEmailRecordingstateOK struct {
	Payload string
}

// IsSuccess returns true when this put conversations email recordingstate o k response has a 2xx status code
func (o *PutConversationsEmailRecordingstateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put conversations email recordingstate o k response has a 3xx status code
func (o *PutConversationsEmailRecordingstateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations email recordingstate o k response has a 4xx status code
func (o *PutConversationsEmailRecordingstateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put conversations email recordingstate o k response has a 5xx status code
func (o *PutConversationsEmailRecordingstateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversations email recordingstate o k response a status code equal to that given
func (o *PutConversationsEmailRecordingstateOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutConversationsEmailRecordingstateOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/emails/{conversationId}/recordingstate][%d] putConversationsEmailRecordingstateOK  %+v", 200, o.Payload)
}

func (o *PutConversationsEmailRecordingstateOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/emails/{conversationId}/recordingstate][%d] putConversationsEmailRecordingstateOK  %+v", 200, o.Payload)
}

func (o *PutConversationsEmailRecordingstateOK) GetPayload() string {
	return o.Payload
}

func (o *PutConversationsEmailRecordingstateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsEmailRecordingstateAccepted creates a PutConversationsEmailRecordingstateAccepted with default headers values
func NewPutConversationsEmailRecordingstateAccepted() *PutConversationsEmailRecordingstateAccepted {
	return &PutConversationsEmailRecordingstateAccepted{}
}

/*
PutConversationsEmailRecordingstateAccepted describes a response with status code 202, with default header values.

Accepted - when pausing or resuming recordings (Secure Pause)
*/
type PutConversationsEmailRecordingstateAccepted struct {
	Payload string
}

// IsSuccess returns true when this put conversations email recordingstate accepted response has a 2xx status code
func (o *PutConversationsEmailRecordingstateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put conversations email recordingstate accepted response has a 3xx status code
func (o *PutConversationsEmailRecordingstateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations email recordingstate accepted response has a 4xx status code
func (o *PutConversationsEmailRecordingstateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this put conversations email recordingstate accepted response has a 5xx status code
func (o *PutConversationsEmailRecordingstateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversations email recordingstate accepted response a status code equal to that given
func (o *PutConversationsEmailRecordingstateAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *PutConversationsEmailRecordingstateAccepted) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/emails/{conversationId}/recordingstate][%d] putConversationsEmailRecordingstateAccepted  %+v", 202, o.Payload)
}

func (o *PutConversationsEmailRecordingstateAccepted) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/emails/{conversationId}/recordingstate][%d] putConversationsEmailRecordingstateAccepted  %+v", 202, o.Payload)
}

func (o *PutConversationsEmailRecordingstateAccepted) GetPayload() string {
	return o.Payload
}

func (o *PutConversationsEmailRecordingstateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsEmailRecordingstateBadRequest creates a PutConversationsEmailRecordingstateBadRequest with default headers values
func NewPutConversationsEmailRecordingstateBadRequest() *PutConversationsEmailRecordingstateBadRequest {
	return &PutConversationsEmailRecordingstateBadRequest{}
}

/*
PutConversationsEmailRecordingstateBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutConversationsEmailRecordingstateBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations email recordingstate bad request response has a 2xx status code
func (o *PutConversationsEmailRecordingstateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations email recordingstate bad request response has a 3xx status code
func (o *PutConversationsEmailRecordingstateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations email recordingstate bad request response has a 4xx status code
func (o *PutConversationsEmailRecordingstateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversations email recordingstate bad request response has a 5xx status code
func (o *PutConversationsEmailRecordingstateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversations email recordingstate bad request response a status code equal to that given
func (o *PutConversationsEmailRecordingstateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutConversationsEmailRecordingstateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/emails/{conversationId}/recordingstate][%d] putConversationsEmailRecordingstateBadRequest  %+v", 400, o.Payload)
}

func (o *PutConversationsEmailRecordingstateBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/emails/{conversationId}/recordingstate][%d] putConversationsEmailRecordingstateBadRequest  %+v", 400, o.Payload)
}

func (o *PutConversationsEmailRecordingstateBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsEmailRecordingstateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsEmailRecordingstateUnauthorized creates a PutConversationsEmailRecordingstateUnauthorized with default headers values
func NewPutConversationsEmailRecordingstateUnauthorized() *PutConversationsEmailRecordingstateUnauthorized {
	return &PutConversationsEmailRecordingstateUnauthorized{}
}

/*
PutConversationsEmailRecordingstateUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PutConversationsEmailRecordingstateUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations email recordingstate unauthorized response has a 2xx status code
func (o *PutConversationsEmailRecordingstateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations email recordingstate unauthorized response has a 3xx status code
func (o *PutConversationsEmailRecordingstateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations email recordingstate unauthorized response has a 4xx status code
func (o *PutConversationsEmailRecordingstateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversations email recordingstate unauthorized response has a 5xx status code
func (o *PutConversationsEmailRecordingstateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversations email recordingstate unauthorized response a status code equal to that given
func (o *PutConversationsEmailRecordingstateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PutConversationsEmailRecordingstateUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/emails/{conversationId}/recordingstate][%d] putConversationsEmailRecordingstateUnauthorized  %+v", 401, o.Payload)
}

func (o *PutConversationsEmailRecordingstateUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/emails/{conversationId}/recordingstate][%d] putConversationsEmailRecordingstateUnauthorized  %+v", 401, o.Payload)
}

func (o *PutConversationsEmailRecordingstateUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsEmailRecordingstateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsEmailRecordingstateForbidden creates a PutConversationsEmailRecordingstateForbidden with default headers values
func NewPutConversationsEmailRecordingstateForbidden() *PutConversationsEmailRecordingstateForbidden {
	return &PutConversationsEmailRecordingstateForbidden{}
}

/*
PutConversationsEmailRecordingstateForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PutConversationsEmailRecordingstateForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations email recordingstate forbidden response has a 2xx status code
func (o *PutConversationsEmailRecordingstateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations email recordingstate forbidden response has a 3xx status code
func (o *PutConversationsEmailRecordingstateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations email recordingstate forbidden response has a 4xx status code
func (o *PutConversationsEmailRecordingstateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversations email recordingstate forbidden response has a 5xx status code
func (o *PutConversationsEmailRecordingstateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversations email recordingstate forbidden response a status code equal to that given
func (o *PutConversationsEmailRecordingstateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutConversationsEmailRecordingstateForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/emails/{conversationId}/recordingstate][%d] putConversationsEmailRecordingstateForbidden  %+v", 403, o.Payload)
}

func (o *PutConversationsEmailRecordingstateForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/emails/{conversationId}/recordingstate][%d] putConversationsEmailRecordingstateForbidden  %+v", 403, o.Payload)
}

func (o *PutConversationsEmailRecordingstateForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsEmailRecordingstateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsEmailRecordingstateNotFound creates a PutConversationsEmailRecordingstateNotFound with default headers values
func NewPutConversationsEmailRecordingstateNotFound() *PutConversationsEmailRecordingstateNotFound {
	return &PutConversationsEmailRecordingstateNotFound{}
}

/*
PutConversationsEmailRecordingstateNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PutConversationsEmailRecordingstateNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations email recordingstate not found response has a 2xx status code
func (o *PutConversationsEmailRecordingstateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations email recordingstate not found response has a 3xx status code
func (o *PutConversationsEmailRecordingstateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations email recordingstate not found response has a 4xx status code
func (o *PutConversationsEmailRecordingstateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversations email recordingstate not found response has a 5xx status code
func (o *PutConversationsEmailRecordingstateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversations email recordingstate not found response a status code equal to that given
func (o *PutConversationsEmailRecordingstateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutConversationsEmailRecordingstateNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/emails/{conversationId}/recordingstate][%d] putConversationsEmailRecordingstateNotFound  %+v", 404, o.Payload)
}

func (o *PutConversationsEmailRecordingstateNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/emails/{conversationId}/recordingstate][%d] putConversationsEmailRecordingstateNotFound  %+v", 404, o.Payload)
}

func (o *PutConversationsEmailRecordingstateNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsEmailRecordingstateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsEmailRecordingstateRequestTimeout creates a PutConversationsEmailRecordingstateRequestTimeout with default headers values
func NewPutConversationsEmailRecordingstateRequestTimeout() *PutConversationsEmailRecordingstateRequestTimeout {
	return &PutConversationsEmailRecordingstateRequestTimeout{}
}

/*
PutConversationsEmailRecordingstateRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutConversationsEmailRecordingstateRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations email recordingstate request timeout response has a 2xx status code
func (o *PutConversationsEmailRecordingstateRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations email recordingstate request timeout response has a 3xx status code
func (o *PutConversationsEmailRecordingstateRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations email recordingstate request timeout response has a 4xx status code
func (o *PutConversationsEmailRecordingstateRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversations email recordingstate request timeout response has a 5xx status code
func (o *PutConversationsEmailRecordingstateRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversations email recordingstate request timeout response a status code equal to that given
func (o *PutConversationsEmailRecordingstateRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PutConversationsEmailRecordingstateRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/emails/{conversationId}/recordingstate][%d] putConversationsEmailRecordingstateRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutConversationsEmailRecordingstateRequestTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/emails/{conversationId}/recordingstate][%d] putConversationsEmailRecordingstateRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutConversationsEmailRecordingstateRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsEmailRecordingstateRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsEmailRecordingstateRequestEntityTooLarge creates a PutConversationsEmailRecordingstateRequestEntityTooLarge with default headers values
func NewPutConversationsEmailRecordingstateRequestEntityTooLarge() *PutConversationsEmailRecordingstateRequestEntityTooLarge {
	return &PutConversationsEmailRecordingstateRequestEntityTooLarge{}
}

/*
PutConversationsEmailRecordingstateRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PutConversationsEmailRecordingstateRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations email recordingstate request entity too large response has a 2xx status code
func (o *PutConversationsEmailRecordingstateRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations email recordingstate request entity too large response has a 3xx status code
func (o *PutConversationsEmailRecordingstateRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations email recordingstate request entity too large response has a 4xx status code
func (o *PutConversationsEmailRecordingstateRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversations email recordingstate request entity too large response has a 5xx status code
func (o *PutConversationsEmailRecordingstateRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversations email recordingstate request entity too large response a status code equal to that given
func (o *PutConversationsEmailRecordingstateRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PutConversationsEmailRecordingstateRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/emails/{conversationId}/recordingstate][%d] putConversationsEmailRecordingstateRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutConversationsEmailRecordingstateRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/emails/{conversationId}/recordingstate][%d] putConversationsEmailRecordingstateRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutConversationsEmailRecordingstateRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsEmailRecordingstateRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsEmailRecordingstateUnsupportedMediaType creates a PutConversationsEmailRecordingstateUnsupportedMediaType with default headers values
func NewPutConversationsEmailRecordingstateUnsupportedMediaType() *PutConversationsEmailRecordingstateUnsupportedMediaType {
	return &PutConversationsEmailRecordingstateUnsupportedMediaType{}
}

/*
PutConversationsEmailRecordingstateUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutConversationsEmailRecordingstateUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations email recordingstate unsupported media type response has a 2xx status code
func (o *PutConversationsEmailRecordingstateUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations email recordingstate unsupported media type response has a 3xx status code
func (o *PutConversationsEmailRecordingstateUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations email recordingstate unsupported media type response has a 4xx status code
func (o *PutConversationsEmailRecordingstateUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversations email recordingstate unsupported media type response has a 5xx status code
func (o *PutConversationsEmailRecordingstateUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversations email recordingstate unsupported media type response a status code equal to that given
func (o *PutConversationsEmailRecordingstateUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PutConversationsEmailRecordingstateUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/emails/{conversationId}/recordingstate][%d] putConversationsEmailRecordingstateUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutConversationsEmailRecordingstateUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/emails/{conversationId}/recordingstate][%d] putConversationsEmailRecordingstateUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutConversationsEmailRecordingstateUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsEmailRecordingstateUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsEmailRecordingstateTooManyRequests creates a PutConversationsEmailRecordingstateTooManyRequests with default headers values
func NewPutConversationsEmailRecordingstateTooManyRequests() *PutConversationsEmailRecordingstateTooManyRequests {
	return &PutConversationsEmailRecordingstateTooManyRequests{}
}

/*
PutConversationsEmailRecordingstateTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutConversationsEmailRecordingstateTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations email recordingstate too many requests response has a 2xx status code
func (o *PutConversationsEmailRecordingstateTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations email recordingstate too many requests response has a 3xx status code
func (o *PutConversationsEmailRecordingstateTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations email recordingstate too many requests response has a 4xx status code
func (o *PutConversationsEmailRecordingstateTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversations email recordingstate too many requests response has a 5xx status code
func (o *PutConversationsEmailRecordingstateTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversations email recordingstate too many requests response a status code equal to that given
func (o *PutConversationsEmailRecordingstateTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PutConversationsEmailRecordingstateTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/emails/{conversationId}/recordingstate][%d] putConversationsEmailRecordingstateTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutConversationsEmailRecordingstateTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/emails/{conversationId}/recordingstate][%d] putConversationsEmailRecordingstateTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutConversationsEmailRecordingstateTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsEmailRecordingstateTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsEmailRecordingstateInternalServerError creates a PutConversationsEmailRecordingstateInternalServerError with default headers values
func NewPutConversationsEmailRecordingstateInternalServerError() *PutConversationsEmailRecordingstateInternalServerError {
	return &PutConversationsEmailRecordingstateInternalServerError{}
}

/*
PutConversationsEmailRecordingstateInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutConversationsEmailRecordingstateInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations email recordingstate internal server error response has a 2xx status code
func (o *PutConversationsEmailRecordingstateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations email recordingstate internal server error response has a 3xx status code
func (o *PutConversationsEmailRecordingstateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations email recordingstate internal server error response has a 4xx status code
func (o *PutConversationsEmailRecordingstateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put conversations email recordingstate internal server error response has a 5xx status code
func (o *PutConversationsEmailRecordingstateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put conversations email recordingstate internal server error response a status code equal to that given
func (o *PutConversationsEmailRecordingstateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutConversationsEmailRecordingstateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/emails/{conversationId}/recordingstate][%d] putConversationsEmailRecordingstateInternalServerError  %+v", 500, o.Payload)
}

func (o *PutConversationsEmailRecordingstateInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/emails/{conversationId}/recordingstate][%d] putConversationsEmailRecordingstateInternalServerError  %+v", 500, o.Payload)
}

func (o *PutConversationsEmailRecordingstateInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsEmailRecordingstateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsEmailRecordingstateServiceUnavailable creates a PutConversationsEmailRecordingstateServiceUnavailable with default headers values
func NewPutConversationsEmailRecordingstateServiceUnavailable() *PutConversationsEmailRecordingstateServiceUnavailable {
	return &PutConversationsEmailRecordingstateServiceUnavailable{}
}

/*
PutConversationsEmailRecordingstateServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutConversationsEmailRecordingstateServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations email recordingstate service unavailable response has a 2xx status code
func (o *PutConversationsEmailRecordingstateServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations email recordingstate service unavailable response has a 3xx status code
func (o *PutConversationsEmailRecordingstateServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations email recordingstate service unavailable response has a 4xx status code
func (o *PutConversationsEmailRecordingstateServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this put conversations email recordingstate service unavailable response has a 5xx status code
func (o *PutConversationsEmailRecordingstateServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this put conversations email recordingstate service unavailable response a status code equal to that given
func (o *PutConversationsEmailRecordingstateServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PutConversationsEmailRecordingstateServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/emails/{conversationId}/recordingstate][%d] putConversationsEmailRecordingstateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutConversationsEmailRecordingstateServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/emails/{conversationId}/recordingstate][%d] putConversationsEmailRecordingstateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutConversationsEmailRecordingstateServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsEmailRecordingstateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsEmailRecordingstateGatewayTimeout creates a PutConversationsEmailRecordingstateGatewayTimeout with default headers values
func NewPutConversationsEmailRecordingstateGatewayTimeout() *PutConversationsEmailRecordingstateGatewayTimeout {
	return &PutConversationsEmailRecordingstateGatewayTimeout{}
}

/*
PutConversationsEmailRecordingstateGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PutConversationsEmailRecordingstateGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations email recordingstate gateway timeout response has a 2xx status code
func (o *PutConversationsEmailRecordingstateGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations email recordingstate gateway timeout response has a 3xx status code
func (o *PutConversationsEmailRecordingstateGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations email recordingstate gateway timeout response has a 4xx status code
func (o *PutConversationsEmailRecordingstateGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this put conversations email recordingstate gateway timeout response has a 5xx status code
func (o *PutConversationsEmailRecordingstateGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this put conversations email recordingstate gateway timeout response a status code equal to that given
func (o *PutConversationsEmailRecordingstateGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PutConversationsEmailRecordingstateGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/emails/{conversationId}/recordingstate][%d] putConversationsEmailRecordingstateGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutConversationsEmailRecordingstateGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/emails/{conversationId}/recordingstate][%d] putConversationsEmailRecordingstateGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutConversationsEmailRecordingstateGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsEmailRecordingstateGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

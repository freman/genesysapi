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

// PutConversationsVideoRecordingstateReader is a Reader for the PutConversationsVideoRecordingstate structure.
type PutConversationsVideoRecordingstateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutConversationsVideoRecordingstateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutConversationsVideoRecordingstateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPutConversationsVideoRecordingstateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutConversationsVideoRecordingstateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutConversationsVideoRecordingstateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutConversationsVideoRecordingstateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutConversationsVideoRecordingstateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutConversationsVideoRecordingstateRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutConversationsVideoRecordingstateRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutConversationsVideoRecordingstateUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutConversationsVideoRecordingstateTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutConversationsVideoRecordingstateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutConversationsVideoRecordingstateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutConversationsVideoRecordingstateGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutConversationsVideoRecordingstateOK creates a PutConversationsVideoRecordingstateOK with default headers values
func NewPutConversationsVideoRecordingstateOK() *PutConversationsVideoRecordingstateOK {
	return &PutConversationsVideoRecordingstateOK{}
}

/*
PutConversationsVideoRecordingstateOK describes a response with status code 200, with default header values.

successful operation
*/
type PutConversationsVideoRecordingstateOK struct {
	Payload string
}

// IsSuccess returns true when this put conversations video recordingstate o k response has a 2xx status code
func (o *PutConversationsVideoRecordingstateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put conversations video recordingstate o k response has a 3xx status code
func (o *PutConversationsVideoRecordingstateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations video recordingstate o k response has a 4xx status code
func (o *PutConversationsVideoRecordingstateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put conversations video recordingstate o k response has a 5xx status code
func (o *PutConversationsVideoRecordingstateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversations video recordingstate o k response a status code equal to that given
func (o *PutConversationsVideoRecordingstateOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutConversationsVideoRecordingstateOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/videos/{conversationId}/recordingstate][%d] putConversationsVideoRecordingstateOK  %+v", 200, o.Payload)
}

func (o *PutConversationsVideoRecordingstateOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/videos/{conversationId}/recordingstate][%d] putConversationsVideoRecordingstateOK  %+v", 200, o.Payload)
}

func (o *PutConversationsVideoRecordingstateOK) GetPayload() string {
	return o.Payload
}

func (o *PutConversationsVideoRecordingstateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsVideoRecordingstateAccepted creates a PutConversationsVideoRecordingstateAccepted with default headers values
func NewPutConversationsVideoRecordingstateAccepted() *PutConversationsVideoRecordingstateAccepted {
	return &PutConversationsVideoRecordingstateAccepted{}
}

/*
PutConversationsVideoRecordingstateAccepted describes a response with status code 202, with default header values.

Accepted - when pausing or resuming recordings (Secure Pause)
*/
type PutConversationsVideoRecordingstateAccepted struct {
	Payload string
}

// IsSuccess returns true when this put conversations video recordingstate accepted response has a 2xx status code
func (o *PutConversationsVideoRecordingstateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put conversations video recordingstate accepted response has a 3xx status code
func (o *PutConversationsVideoRecordingstateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations video recordingstate accepted response has a 4xx status code
func (o *PutConversationsVideoRecordingstateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this put conversations video recordingstate accepted response has a 5xx status code
func (o *PutConversationsVideoRecordingstateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversations video recordingstate accepted response a status code equal to that given
func (o *PutConversationsVideoRecordingstateAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *PutConversationsVideoRecordingstateAccepted) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/videos/{conversationId}/recordingstate][%d] putConversationsVideoRecordingstateAccepted  %+v", 202, o.Payload)
}

func (o *PutConversationsVideoRecordingstateAccepted) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/videos/{conversationId}/recordingstate][%d] putConversationsVideoRecordingstateAccepted  %+v", 202, o.Payload)
}

func (o *PutConversationsVideoRecordingstateAccepted) GetPayload() string {
	return o.Payload
}

func (o *PutConversationsVideoRecordingstateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsVideoRecordingstateBadRequest creates a PutConversationsVideoRecordingstateBadRequest with default headers values
func NewPutConversationsVideoRecordingstateBadRequest() *PutConversationsVideoRecordingstateBadRequest {
	return &PutConversationsVideoRecordingstateBadRequest{}
}

/*
PutConversationsVideoRecordingstateBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutConversationsVideoRecordingstateBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations video recordingstate bad request response has a 2xx status code
func (o *PutConversationsVideoRecordingstateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations video recordingstate bad request response has a 3xx status code
func (o *PutConversationsVideoRecordingstateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations video recordingstate bad request response has a 4xx status code
func (o *PutConversationsVideoRecordingstateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversations video recordingstate bad request response has a 5xx status code
func (o *PutConversationsVideoRecordingstateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversations video recordingstate bad request response a status code equal to that given
func (o *PutConversationsVideoRecordingstateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutConversationsVideoRecordingstateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/videos/{conversationId}/recordingstate][%d] putConversationsVideoRecordingstateBadRequest  %+v", 400, o.Payload)
}

func (o *PutConversationsVideoRecordingstateBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/videos/{conversationId}/recordingstate][%d] putConversationsVideoRecordingstateBadRequest  %+v", 400, o.Payload)
}

func (o *PutConversationsVideoRecordingstateBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsVideoRecordingstateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsVideoRecordingstateUnauthorized creates a PutConversationsVideoRecordingstateUnauthorized with default headers values
func NewPutConversationsVideoRecordingstateUnauthorized() *PutConversationsVideoRecordingstateUnauthorized {
	return &PutConversationsVideoRecordingstateUnauthorized{}
}

/*
PutConversationsVideoRecordingstateUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PutConversationsVideoRecordingstateUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations video recordingstate unauthorized response has a 2xx status code
func (o *PutConversationsVideoRecordingstateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations video recordingstate unauthorized response has a 3xx status code
func (o *PutConversationsVideoRecordingstateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations video recordingstate unauthorized response has a 4xx status code
func (o *PutConversationsVideoRecordingstateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversations video recordingstate unauthorized response has a 5xx status code
func (o *PutConversationsVideoRecordingstateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversations video recordingstate unauthorized response a status code equal to that given
func (o *PutConversationsVideoRecordingstateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PutConversationsVideoRecordingstateUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/videos/{conversationId}/recordingstate][%d] putConversationsVideoRecordingstateUnauthorized  %+v", 401, o.Payload)
}

func (o *PutConversationsVideoRecordingstateUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/videos/{conversationId}/recordingstate][%d] putConversationsVideoRecordingstateUnauthorized  %+v", 401, o.Payload)
}

func (o *PutConversationsVideoRecordingstateUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsVideoRecordingstateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsVideoRecordingstateForbidden creates a PutConversationsVideoRecordingstateForbidden with default headers values
func NewPutConversationsVideoRecordingstateForbidden() *PutConversationsVideoRecordingstateForbidden {
	return &PutConversationsVideoRecordingstateForbidden{}
}

/*
PutConversationsVideoRecordingstateForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PutConversationsVideoRecordingstateForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations video recordingstate forbidden response has a 2xx status code
func (o *PutConversationsVideoRecordingstateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations video recordingstate forbidden response has a 3xx status code
func (o *PutConversationsVideoRecordingstateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations video recordingstate forbidden response has a 4xx status code
func (o *PutConversationsVideoRecordingstateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversations video recordingstate forbidden response has a 5xx status code
func (o *PutConversationsVideoRecordingstateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversations video recordingstate forbidden response a status code equal to that given
func (o *PutConversationsVideoRecordingstateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutConversationsVideoRecordingstateForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/videos/{conversationId}/recordingstate][%d] putConversationsVideoRecordingstateForbidden  %+v", 403, o.Payload)
}

func (o *PutConversationsVideoRecordingstateForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/videos/{conversationId}/recordingstate][%d] putConversationsVideoRecordingstateForbidden  %+v", 403, o.Payload)
}

func (o *PutConversationsVideoRecordingstateForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsVideoRecordingstateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsVideoRecordingstateNotFound creates a PutConversationsVideoRecordingstateNotFound with default headers values
func NewPutConversationsVideoRecordingstateNotFound() *PutConversationsVideoRecordingstateNotFound {
	return &PutConversationsVideoRecordingstateNotFound{}
}

/*
PutConversationsVideoRecordingstateNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PutConversationsVideoRecordingstateNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations video recordingstate not found response has a 2xx status code
func (o *PutConversationsVideoRecordingstateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations video recordingstate not found response has a 3xx status code
func (o *PutConversationsVideoRecordingstateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations video recordingstate not found response has a 4xx status code
func (o *PutConversationsVideoRecordingstateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversations video recordingstate not found response has a 5xx status code
func (o *PutConversationsVideoRecordingstateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversations video recordingstate not found response a status code equal to that given
func (o *PutConversationsVideoRecordingstateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutConversationsVideoRecordingstateNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/videos/{conversationId}/recordingstate][%d] putConversationsVideoRecordingstateNotFound  %+v", 404, o.Payload)
}

func (o *PutConversationsVideoRecordingstateNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/videos/{conversationId}/recordingstate][%d] putConversationsVideoRecordingstateNotFound  %+v", 404, o.Payload)
}

func (o *PutConversationsVideoRecordingstateNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsVideoRecordingstateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsVideoRecordingstateRequestTimeout creates a PutConversationsVideoRecordingstateRequestTimeout with default headers values
func NewPutConversationsVideoRecordingstateRequestTimeout() *PutConversationsVideoRecordingstateRequestTimeout {
	return &PutConversationsVideoRecordingstateRequestTimeout{}
}

/*
PutConversationsVideoRecordingstateRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutConversationsVideoRecordingstateRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations video recordingstate request timeout response has a 2xx status code
func (o *PutConversationsVideoRecordingstateRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations video recordingstate request timeout response has a 3xx status code
func (o *PutConversationsVideoRecordingstateRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations video recordingstate request timeout response has a 4xx status code
func (o *PutConversationsVideoRecordingstateRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversations video recordingstate request timeout response has a 5xx status code
func (o *PutConversationsVideoRecordingstateRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversations video recordingstate request timeout response a status code equal to that given
func (o *PutConversationsVideoRecordingstateRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PutConversationsVideoRecordingstateRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/videos/{conversationId}/recordingstate][%d] putConversationsVideoRecordingstateRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutConversationsVideoRecordingstateRequestTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/videos/{conversationId}/recordingstate][%d] putConversationsVideoRecordingstateRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutConversationsVideoRecordingstateRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsVideoRecordingstateRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsVideoRecordingstateRequestEntityTooLarge creates a PutConversationsVideoRecordingstateRequestEntityTooLarge with default headers values
func NewPutConversationsVideoRecordingstateRequestEntityTooLarge() *PutConversationsVideoRecordingstateRequestEntityTooLarge {
	return &PutConversationsVideoRecordingstateRequestEntityTooLarge{}
}

/*
PutConversationsVideoRecordingstateRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PutConversationsVideoRecordingstateRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations video recordingstate request entity too large response has a 2xx status code
func (o *PutConversationsVideoRecordingstateRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations video recordingstate request entity too large response has a 3xx status code
func (o *PutConversationsVideoRecordingstateRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations video recordingstate request entity too large response has a 4xx status code
func (o *PutConversationsVideoRecordingstateRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversations video recordingstate request entity too large response has a 5xx status code
func (o *PutConversationsVideoRecordingstateRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversations video recordingstate request entity too large response a status code equal to that given
func (o *PutConversationsVideoRecordingstateRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PutConversationsVideoRecordingstateRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/videos/{conversationId}/recordingstate][%d] putConversationsVideoRecordingstateRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutConversationsVideoRecordingstateRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/videos/{conversationId}/recordingstate][%d] putConversationsVideoRecordingstateRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutConversationsVideoRecordingstateRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsVideoRecordingstateRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsVideoRecordingstateUnsupportedMediaType creates a PutConversationsVideoRecordingstateUnsupportedMediaType with default headers values
func NewPutConversationsVideoRecordingstateUnsupportedMediaType() *PutConversationsVideoRecordingstateUnsupportedMediaType {
	return &PutConversationsVideoRecordingstateUnsupportedMediaType{}
}

/*
PutConversationsVideoRecordingstateUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutConversationsVideoRecordingstateUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations video recordingstate unsupported media type response has a 2xx status code
func (o *PutConversationsVideoRecordingstateUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations video recordingstate unsupported media type response has a 3xx status code
func (o *PutConversationsVideoRecordingstateUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations video recordingstate unsupported media type response has a 4xx status code
func (o *PutConversationsVideoRecordingstateUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversations video recordingstate unsupported media type response has a 5xx status code
func (o *PutConversationsVideoRecordingstateUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversations video recordingstate unsupported media type response a status code equal to that given
func (o *PutConversationsVideoRecordingstateUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PutConversationsVideoRecordingstateUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/videos/{conversationId}/recordingstate][%d] putConversationsVideoRecordingstateUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutConversationsVideoRecordingstateUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/videos/{conversationId}/recordingstate][%d] putConversationsVideoRecordingstateUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutConversationsVideoRecordingstateUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsVideoRecordingstateUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsVideoRecordingstateTooManyRequests creates a PutConversationsVideoRecordingstateTooManyRequests with default headers values
func NewPutConversationsVideoRecordingstateTooManyRequests() *PutConversationsVideoRecordingstateTooManyRequests {
	return &PutConversationsVideoRecordingstateTooManyRequests{}
}

/*
PutConversationsVideoRecordingstateTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutConversationsVideoRecordingstateTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations video recordingstate too many requests response has a 2xx status code
func (o *PutConversationsVideoRecordingstateTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations video recordingstate too many requests response has a 3xx status code
func (o *PutConversationsVideoRecordingstateTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations video recordingstate too many requests response has a 4xx status code
func (o *PutConversationsVideoRecordingstateTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversations video recordingstate too many requests response has a 5xx status code
func (o *PutConversationsVideoRecordingstateTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversations video recordingstate too many requests response a status code equal to that given
func (o *PutConversationsVideoRecordingstateTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PutConversationsVideoRecordingstateTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/videos/{conversationId}/recordingstate][%d] putConversationsVideoRecordingstateTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutConversationsVideoRecordingstateTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/videos/{conversationId}/recordingstate][%d] putConversationsVideoRecordingstateTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutConversationsVideoRecordingstateTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsVideoRecordingstateTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsVideoRecordingstateInternalServerError creates a PutConversationsVideoRecordingstateInternalServerError with default headers values
func NewPutConversationsVideoRecordingstateInternalServerError() *PutConversationsVideoRecordingstateInternalServerError {
	return &PutConversationsVideoRecordingstateInternalServerError{}
}

/*
PutConversationsVideoRecordingstateInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutConversationsVideoRecordingstateInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations video recordingstate internal server error response has a 2xx status code
func (o *PutConversationsVideoRecordingstateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations video recordingstate internal server error response has a 3xx status code
func (o *PutConversationsVideoRecordingstateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations video recordingstate internal server error response has a 4xx status code
func (o *PutConversationsVideoRecordingstateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put conversations video recordingstate internal server error response has a 5xx status code
func (o *PutConversationsVideoRecordingstateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put conversations video recordingstate internal server error response a status code equal to that given
func (o *PutConversationsVideoRecordingstateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutConversationsVideoRecordingstateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/videos/{conversationId}/recordingstate][%d] putConversationsVideoRecordingstateInternalServerError  %+v", 500, o.Payload)
}

func (o *PutConversationsVideoRecordingstateInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/videos/{conversationId}/recordingstate][%d] putConversationsVideoRecordingstateInternalServerError  %+v", 500, o.Payload)
}

func (o *PutConversationsVideoRecordingstateInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsVideoRecordingstateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsVideoRecordingstateServiceUnavailable creates a PutConversationsVideoRecordingstateServiceUnavailable with default headers values
func NewPutConversationsVideoRecordingstateServiceUnavailable() *PutConversationsVideoRecordingstateServiceUnavailable {
	return &PutConversationsVideoRecordingstateServiceUnavailable{}
}

/*
PutConversationsVideoRecordingstateServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutConversationsVideoRecordingstateServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations video recordingstate service unavailable response has a 2xx status code
func (o *PutConversationsVideoRecordingstateServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations video recordingstate service unavailable response has a 3xx status code
func (o *PutConversationsVideoRecordingstateServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations video recordingstate service unavailable response has a 4xx status code
func (o *PutConversationsVideoRecordingstateServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this put conversations video recordingstate service unavailable response has a 5xx status code
func (o *PutConversationsVideoRecordingstateServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this put conversations video recordingstate service unavailable response a status code equal to that given
func (o *PutConversationsVideoRecordingstateServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PutConversationsVideoRecordingstateServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/videos/{conversationId}/recordingstate][%d] putConversationsVideoRecordingstateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutConversationsVideoRecordingstateServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/videos/{conversationId}/recordingstate][%d] putConversationsVideoRecordingstateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutConversationsVideoRecordingstateServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsVideoRecordingstateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationsVideoRecordingstateGatewayTimeout creates a PutConversationsVideoRecordingstateGatewayTimeout with default headers values
func NewPutConversationsVideoRecordingstateGatewayTimeout() *PutConversationsVideoRecordingstateGatewayTimeout {
	return &PutConversationsVideoRecordingstateGatewayTimeout{}
}

/*
PutConversationsVideoRecordingstateGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PutConversationsVideoRecordingstateGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversations video recordingstate gateway timeout response has a 2xx status code
func (o *PutConversationsVideoRecordingstateGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversations video recordingstate gateway timeout response has a 3xx status code
func (o *PutConversationsVideoRecordingstateGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversations video recordingstate gateway timeout response has a 4xx status code
func (o *PutConversationsVideoRecordingstateGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this put conversations video recordingstate gateway timeout response has a 5xx status code
func (o *PutConversationsVideoRecordingstateGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this put conversations video recordingstate gateway timeout response a status code equal to that given
func (o *PutConversationsVideoRecordingstateGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PutConversationsVideoRecordingstateGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/videos/{conversationId}/recordingstate][%d] putConversationsVideoRecordingstateGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutConversationsVideoRecordingstateGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/videos/{conversationId}/recordingstate][%d] putConversationsVideoRecordingstateGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutConversationsVideoRecordingstateGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationsVideoRecordingstateGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

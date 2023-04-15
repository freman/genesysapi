// Code generated by go-swagger; DO NOT EDIT.

package recording

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetRecordingRecordingkeysRotationscheduleReader is a Reader for the GetRecordingRecordingkeysRotationschedule structure.
type GetRecordingRecordingkeysRotationscheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRecordingRecordingkeysRotationscheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRecordingRecordingkeysRotationscheduleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRecordingRecordingkeysRotationscheduleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRecordingRecordingkeysRotationscheduleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRecordingRecordingkeysRotationscheduleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRecordingRecordingkeysRotationscheduleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetRecordingRecordingkeysRotationscheduleRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRecordingRecordingkeysRotationscheduleRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRecordingRecordingkeysRotationscheduleUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRecordingRecordingkeysRotationscheduleTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRecordingRecordingkeysRotationscheduleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRecordingRecordingkeysRotationscheduleServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRecordingRecordingkeysRotationscheduleGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRecordingRecordingkeysRotationscheduleOK creates a GetRecordingRecordingkeysRotationscheduleOK with default headers values
func NewGetRecordingRecordingkeysRotationscheduleOK() *GetRecordingRecordingkeysRotationscheduleOK {
	return &GetRecordingRecordingkeysRotationscheduleOK{}
}

/*
GetRecordingRecordingkeysRotationscheduleOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRecordingRecordingkeysRotationscheduleOK struct {
	Payload *models.KeyRotationSchedule
}

// IsSuccess returns true when this get recording recordingkeys rotationschedule o k response has a 2xx status code
func (o *GetRecordingRecordingkeysRotationscheduleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get recording recordingkeys rotationschedule o k response has a 3xx status code
func (o *GetRecordingRecordingkeysRotationscheduleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording recordingkeys rotationschedule o k response has a 4xx status code
func (o *GetRecordingRecordingkeysRotationscheduleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get recording recordingkeys rotationschedule o k response has a 5xx status code
func (o *GetRecordingRecordingkeysRotationscheduleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording recordingkeys rotationschedule o k response a status code equal to that given
func (o *GetRecordingRecordingkeysRotationscheduleOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRecordingRecordingkeysRotationscheduleOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/recordingkeys/rotationschedule][%d] getRecordingRecordingkeysRotationscheduleOK  %+v", 200, o.Payload)
}

func (o *GetRecordingRecordingkeysRotationscheduleOK) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/recordingkeys/rotationschedule][%d] getRecordingRecordingkeysRotationscheduleOK  %+v", 200, o.Payload)
}

func (o *GetRecordingRecordingkeysRotationscheduleOK) GetPayload() *models.KeyRotationSchedule {
	return o.Payload
}

func (o *GetRecordingRecordingkeysRotationscheduleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KeyRotationSchedule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingRecordingkeysRotationscheduleBadRequest creates a GetRecordingRecordingkeysRotationscheduleBadRequest with default headers values
func NewGetRecordingRecordingkeysRotationscheduleBadRequest() *GetRecordingRecordingkeysRotationscheduleBadRequest {
	return &GetRecordingRecordingkeysRotationscheduleBadRequest{}
}

/*
GetRecordingRecordingkeysRotationscheduleBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRecordingRecordingkeysRotationscheduleBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording recordingkeys rotationschedule bad request response has a 2xx status code
func (o *GetRecordingRecordingkeysRotationscheduleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording recordingkeys rotationschedule bad request response has a 3xx status code
func (o *GetRecordingRecordingkeysRotationscheduleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording recordingkeys rotationschedule bad request response has a 4xx status code
func (o *GetRecordingRecordingkeysRotationscheduleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get recording recordingkeys rotationschedule bad request response has a 5xx status code
func (o *GetRecordingRecordingkeysRotationscheduleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording recordingkeys rotationschedule bad request response a status code equal to that given
func (o *GetRecordingRecordingkeysRotationscheduleBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetRecordingRecordingkeysRotationscheduleBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/recordingkeys/rotationschedule][%d] getRecordingRecordingkeysRotationscheduleBadRequest  %+v", 400, o.Payload)
}

func (o *GetRecordingRecordingkeysRotationscheduleBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/recordingkeys/rotationschedule][%d] getRecordingRecordingkeysRotationscheduleBadRequest  %+v", 400, o.Payload)
}

func (o *GetRecordingRecordingkeysRotationscheduleBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingRecordingkeysRotationscheduleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingRecordingkeysRotationscheduleUnauthorized creates a GetRecordingRecordingkeysRotationscheduleUnauthorized with default headers values
func NewGetRecordingRecordingkeysRotationscheduleUnauthorized() *GetRecordingRecordingkeysRotationscheduleUnauthorized {
	return &GetRecordingRecordingkeysRotationscheduleUnauthorized{}
}

/*
GetRecordingRecordingkeysRotationscheduleUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRecordingRecordingkeysRotationscheduleUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording recordingkeys rotationschedule unauthorized response has a 2xx status code
func (o *GetRecordingRecordingkeysRotationscheduleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording recordingkeys rotationschedule unauthorized response has a 3xx status code
func (o *GetRecordingRecordingkeysRotationscheduleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording recordingkeys rotationschedule unauthorized response has a 4xx status code
func (o *GetRecordingRecordingkeysRotationscheduleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get recording recordingkeys rotationschedule unauthorized response has a 5xx status code
func (o *GetRecordingRecordingkeysRotationscheduleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording recordingkeys rotationschedule unauthorized response a status code equal to that given
func (o *GetRecordingRecordingkeysRotationscheduleUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetRecordingRecordingkeysRotationscheduleUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/recordingkeys/rotationschedule][%d] getRecordingRecordingkeysRotationscheduleUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRecordingRecordingkeysRotationscheduleUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/recordingkeys/rotationschedule][%d] getRecordingRecordingkeysRotationscheduleUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRecordingRecordingkeysRotationscheduleUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingRecordingkeysRotationscheduleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingRecordingkeysRotationscheduleForbidden creates a GetRecordingRecordingkeysRotationscheduleForbidden with default headers values
func NewGetRecordingRecordingkeysRotationscheduleForbidden() *GetRecordingRecordingkeysRotationscheduleForbidden {
	return &GetRecordingRecordingkeysRotationscheduleForbidden{}
}

/*
GetRecordingRecordingkeysRotationscheduleForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetRecordingRecordingkeysRotationscheduleForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording recordingkeys rotationschedule forbidden response has a 2xx status code
func (o *GetRecordingRecordingkeysRotationscheduleForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording recordingkeys rotationschedule forbidden response has a 3xx status code
func (o *GetRecordingRecordingkeysRotationscheduleForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording recordingkeys rotationschedule forbidden response has a 4xx status code
func (o *GetRecordingRecordingkeysRotationscheduleForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get recording recordingkeys rotationschedule forbidden response has a 5xx status code
func (o *GetRecordingRecordingkeysRotationscheduleForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording recordingkeys rotationschedule forbidden response a status code equal to that given
func (o *GetRecordingRecordingkeysRotationscheduleForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetRecordingRecordingkeysRotationscheduleForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/recordingkeys/rotationschedule][%d] getRecordingRecordingkeysRotationscheduleForbidden  %+v", 403, o.Payload)
}

func (o *GetRecordingRecordingkeysRotationscheduleForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/recordingkeys/rotationschedule][%d] getRecordingRecordingkeysRotationscheduleForbidden  %+v", 403, o.Payload)
}

func (o *GetRecordingRecordingkeysRotationscheduleForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingRecordingkeysRotationscheduleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingRecordingkeysRotationscheduleNotFound creates a GetRecordingRecordingkeysRotationscheduleNotFound with default headers values
func NewGetRecordingRecordingkeysRotationscheduleNotFound() *GetRecordingRecordingkeysRotationscheduleNotFound {
	return &GetRecordingRecordingkeysRotationscheduleNotFound{}
}

/*
GetRecordingRecordingkeysRotationscheduleNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetRecordingRecordingkeysRotationscheduleNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording recordingkeys rotationschedule not found response has a 2xx status code
func (o *GetRecordingRecordingkeysRotationscheduleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording recordingkeys rotationschedule not found response has a 3xx status code
func (o *GetRecordingRecordingkeysRotationscheduleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording recordingkeys rotationschedule not found response has a 4xx status code
func (o *GetRecordingRecordingkeysRotationscheduleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get recording recordingkeys rotationschedule not found response has a 5xx status code
func (o *GetRecordingRecordingkeysRotationscheduleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording recordingkeys rotationschedule not found response a status code equal to that given
func (o *GetRecordingRecordingkeysRotationscheduleNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetRecordingRecordingkeysRotationscheduleNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/recordingkeys/rotationschedule][%d] getRecordingRecordingkeysRotationscheduleNotFound  %+v", 404, o.Payload)
}

func (o *GetRecordingRecordingkeysRotationscheduleNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/recordingkeys/rotationschedule][%d] getRecordingRecordingkeysRotationscheduleNotFound  %+v", 404, o.Payload)
}

func (o *GetRecordingRecordingkeysRotationscheduleNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingRecordingkeysRotationscheduleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingRecordingkeysRotationscheduleRequestTimeout creates a GetRecordingRecordingkeysRotationscheduleRequestTimeout with default headers values
func NewGetRecordingRecordingkeysRotationscheduleRequestTimeout() *GetRecordingRecordingkeysRotationscheduleRequestTimeout {
	return &GetRecordingRecordingkeysRotationscheduleRequestTimeout{}
}

/*
GetRecordingRecordingkeysRotationscheduleRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetRecordingRecordingkeysRotationscheduleRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording recordingkeys rotationschedule request timeout response has a 2xx status code
func (o *GetRecordingRecordingkeysRotationscheduleRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording recordingkeys rotationschedule request timeout response has a 3xx status code
func (o *GetRecordingRecordingkeysRotationscheduleRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording recordingkeys rotationschedule request timeout response has a 4xx status code
func (o *GetRecordingRecordingkeysRotationscheduleRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get recording recordingkeys rotationschedule request timeout response has a 5xx status code
func (o *GetRecordingRecordingkeysRotationscheduleRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording recordingkeys rotationschedule request timeout response a status code equal to that given
func (o *GetRecordingRecordingkeysRotationscheduleRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetRecordingRecordingkeysRotationscheduleRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/recordingkeys/rotationschedule][%d] getRecordingRecordingkeysRotationscheduleRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRecordingRecordingkeysRotationscheduleRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/recordingkeys/rotationschedule][%d] getRecordingRecordingkeysRotationscheduleRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRecordingRecordingkeysRotationscheduleRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingRecordingkeysRotationscheduleRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingRecordingkeysRotationscheduleRequestEntityTooLarge creates a GetRecordingRecordingkeysRotationscheduleRequestEntityTooLarge with default headers values
func NewGetRecordingRecordingkeysRotationscheduleRequestEntityTooLarge() *GetRecordingRecordingkeysRotationscheduleRequestEntityTooLarge {
	return &GetRecordingRecordingkeysRotationscheduleRequestEntityTooLarge{}
}

/*
GetRecordingRecordingkeysRotationscheduleRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetRecordingRecordingkeysRotationscheduleRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording recordingkeys rotationschedule request entity too large response has a 2xx status code
func (o *GetRecordingRecordingkeysRotationscheduleRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording recordingkeys rotationschedule request entity too large response has a 3xx status code
func (o *GetRecordingRecordingkeysRotationscheduleRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording recordingkeys rotationschedule request entity too large response has a 4xx status code
func (o *GetRecordingRecordingkeysRotationscheduleRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get recording recordingkeys rotationschedule request entity too large response has a 5xx status code
func (o *GetRecordingRecordingkeysRotationscheduleRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording recordingkeys rotationschedule request entity too large response a status code equal to that given
func (o *GetRecordingRecordingkeysRotationscheduleRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetRecordingRecordingkeysRotationscheduleRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/recordingkeys/rotationschedule][%d] getRecordingRecordingkeysRotationscheduleRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRecordingRecordingkeysRotationscheduleRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/recordingkeys/rotationschedule][%d] getRecordingRecordingkeysRotationscheduleRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRecordingRecordingkeysRotationscheduleRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingRecordingkeysRotationscheduleRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingRecordingkeysRotationscheduleUnsupportedMediaType creates a GetRecordingRecordingkeysRotationscheduleUnsupportedMediaType with default headers values
func NewGetRecordingRecordingkeysRotationscheduleUnsupportedMediaType() *GetRecordingRecordingkeysRotationscheduleUnsupportedMediaType {
	return &GetRecordingRecordingkeysRotationscheduleUnsupportedMediaType{}
}

/*
GetRecordingRecordingkeysRotationscheduleUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRecordingRecordingkeysRotationscheduleUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording recordingkeys rotationschedule unsupported media type response has a 2xx status code
func (o *GetRecordingRecordingkeysRotationscheduleUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording recordingkeys rotationschedule unsupported media type response has a 3xx status code
func (o *GetRecordingRecordingkeysRotationscheduleUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording recordingkeys rotationschedule unsupported media type response has a 4xx status code
func (o *GetRecordingRecordingkeysRotationscheduleUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get recording recordingkeys rotationschedule unsupported media type response has a 5xx status code
func (o *GetRecordingRecordingkeysRotationscheduleUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording recordingkeys rotationschedule unsupported media type response a status code equal to that given
func (o *GetRecordingRecordingkeysRotationscheduleUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetRecordingRecordingkeysRotationscheduleUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/recordingkeys/rotationschedule][%d] getRecordingRecordingkeysRotationscheduleUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRecordingRecordingkeysRotationscheduleUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/recordingkeys/rotationschedule][%d] getRecordingRecordingkeysRotationscheduleUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRecordingRecordingkeysRotationscheduleUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingRecordingkeysRotationscheduleUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingRecordingkeysRotationscheduleTooManyRequests creates a GetRecordingRecordingkeysRotationscheduleTooManyRequests with default headers values
func NewGetRecordingRecordingkeysRotationscheduleTooManyRequests() *GetRecordingRecordingkeysRotationscheduleTooManyRequests {
	return &GetRecordingRecordingkeysRotationscheduleTooManyRequests{}
}

/*
GetRecordingRecordingkeysRotationscheduleTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetRecordingRecordingkeysRotationscheduleTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording recordingkeys rotationschedule too many requests response has a 2xx status code
func (o *GetRecordingRecordingkeysRotationscheduleTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording recordingkeys rotationschedule too many requests response has a 3xx status code
func (o *GetRecordingRecordingkeysRotationscheduleTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording recordingkeys rotationschedule too many requests response has a 4xx status code
func (o *GetRecordingRecordingkeysRotationscheduleTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get recording recordingkeys rotationschedule too many requests response has a 5xx status code
func (o *GetRecordingRecordingkeysRotationscheduleTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording recordingkeys rotationschedule too many requests response a status code equal to that given
func (o *GetRecordingRecordingkeysRotationscheduleTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetRecordingRecordingkeysRotationscheduleTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/recordingkeys/rotationschedule][%d] getRecordingRecordingkeysRotationscheduleTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRecordingRecordingkeysRotationscheduleTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/recordingkeys/rotationschedule][%d] getRecordingRecordingkeysRotationscheduleTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRecordingRecordingkeysRotationscheduleTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingRecordingkeysRotationscheduleTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingRecordingkeysRotationscheduleInternalServerError creates a GetRecordingRecordingkeysRotationscheduleInternalServerError with default headers values
func NewGetRecordingRecordingkeysRotationscheduleInternalServerError() *GetRecordingRecordingkeysRotationscheduleInternalServerError {
	return &GetRecordingRecordingkeysRotationscheduleInternalServerError{}
}

/*
GetRecordingRecordingkeysRotationscheduleInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRecordingRecordingkeysRotationscheduleInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording recordingkeys rotationschedule internal server error response has a 2xx status code
func (o *GetRecordingRecordingkeysRotationscheduleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording recordingkeys rotationschedule internal server error response has a 3xx status code
func (o *GetRecordingRecordingkeysRotationscheduleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording recordingkeys rotationschedule internal server error response has a 4xx status code
func (o *GetRecordingRecordingkeysRotationscheduleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get recording recordingkeys rotationschedule internal server error response has a 5xx status code
func (o *GetRecordingRecordingkeysRotationscheduleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get recording recordingkeys rotationschedule internal server error response a status code equal to that given
func (o *GetRecordingRecordingkeysRotationscheduleInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetRecordingRecordingkeysRotationscheduleInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/recordingkeys/rotationschedule][%d] getRecordingRecordingkeysRotationscheduleInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRecordingRecordingkeysRotationscheduleInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/recordingkeys/rotationschedule][%d] getRecordingRecordingkeysRotationscheduleInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRecordingRecordingkeysRotationscheduleInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingRecordingkeysRotationscheduleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingRecordingkeysRotationscheduleServiceUnavailable creates a GetRecordingRecordingkeysRotationscheduleServiceUnavailable with default headers values
func NewGetRecordingRecordingkeysRotationscheduleServiceUnavailable() *GetRecordingRecordingkeysRotationscheduleServiceUnavailable {
	return &GetRecordingRecordingkeysRotationscheduleServiceUnavailable{}
}

/*
GetRecordingRecordingkeysRotationscheduleServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRecordingRecordingkeysRotationscheduleServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording recordingkeys rotationschedule service unavailable response has a 2xx status code
func (o *GetRecordingRecordingkeysRotationscheduleServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording recordingkeys rotationschedule service unavailable response has a 3xx status code
func (o *GetRecordingRecordingkeysRotationscheduleServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording recordingkeys rotationschedule service unavailable response has a 4xx status code
func (o *GetRecordingRecordingkeysRotationscheduleServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get recording recordingkeys rotationschedule service unavailable response has a 5xx status code
func (o *GetRecordingRecordingkeysRotationscheduleServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get recording recordingkeys rotationschedule service unavailable response a status code equal to that given
func (o *GetRecordingRecordingkeysRotationscheduleServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetRecordingRecordingkeysRotationscheduleServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/recordingkeys/rotationschedule][%d] getRecordingRecordingkeysRotationscheduleServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRecordingRecordingkeysRotationscheduleServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/recordingkeys/rotationschedule][%d] getRecordingRecordingkeysRotationscheduleServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRecordingRecordingkeysRotationscheduleServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingRecordingkeysRotationscheduleServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingRecordingkeysRotationscheduleGatewayTimeout creates a GetRecordingRecordingkeysRotationscheduleGatewayTimeout with default headers values
func NewGetRecordingRecordingkeysRotationscheduleGatewayTimeout() *GetRecordingRecordingkeysRotationscheduleGatewayTimeout {
	return &GetRecordingRecordingkeysRotationscheduleGatewayTimeout{}
}

/*
GetRecordingRecordingkeysRotationscheduleGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetRecordingRecordingkeysRotationscheduleGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording recordingkeys rotationschedule gateway timeout response has a 2xx status code
func (o *GetRecordingRecordingkeysRotationscheduleGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording recordingkeys rotationschedule gateway timeout response has a 3xx status code
func (o *GetRecordingRecordingkeysRotationscheduleGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording recordingkeys rotationschedule gateway timeout response has a 4xx status code
func (o *GetRecordingRecordingkeysRotationscheduleGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get recording recordingkeys rotationschedule gateway timeout response has a 5xx status code
func (o *GetRecordingRecordingkeysRotationscheduleGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get recording recordingkeys rotationschedule gateway timeout response a status code equal to that given
func (o *GetRecordingRecordingkeysRotationscheduleGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetRecordingRecordingkeysRotationscheduleGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/recordingkeys/rotationschedule][%d] getRecordingRecordingkeysRotationscheduleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRecordingRecordingkeysRotationscheduleGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/recordingkeys/rotationschedule][%d] getRecordingRecordingkeysRotationscheduleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRecordingRecordingkeysRotationscheduleGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingRecordingkeysRotationscheduleGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

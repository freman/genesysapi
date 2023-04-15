// Code generated by go-swagger; DO NOT EDIT.

package outbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetOutboundSchedulesSequenceReader is a Reader for the GetOutboundSchedulesSequence structure.
type GetOutboundSchedulesSequenceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutboundSchedulesSequenceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutboundSchedulesSequenceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutboundSchedulesSequenceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutboundSchedulesSequenceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutboundSchedulesSequenceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutboundSchedulesSequenceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOutboundSchedulesSequenceRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOutboundSchedulesSequenceRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOutboundSchedulesSequenceUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOutboundSchedulesSequenceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOutboundSchedulesSequenceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOutboundSchedulesSequenceServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOutboundSchedulesSequenceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOutboundSchedulesSequenceOK creates a GetOutboundSchedulesSequenceOK with default headers values
func NewGetOutboundSchedulesSequenceOK() *GetOutboundSchedulesSequenceOK {
	return &GetOutboundSchedulesSequenceOK{}
}

/*
GetOutboundSchedulesSequenceOK describes a response with status code 200, with default header values.

successful operation
*/
type GetOutboundSchedulesSequenceOK struct {
	Payload *models.SequenceSchedule
}

// IsSuccess returns true when this get outbound schedules sequence o k response has a 2xx status code
func (o *GetOutboundSchedulesSequenceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get outbound schedules sequence o k response has a 3xx status code
func (o *GetOutboundSchedulesSequenceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules sequence o k response has a 4xx status code
func (o *GetOutboundSchedulesSequenceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound schedules sequence o k response has a 5xx status code
func (o *GetOutboundSchedulesSequenceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound schedules sequence o k response a status code equal to that given
func (o *GetOutboundSchedulesSequenceOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOutboundSchedulesSequenceOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences/{sequenceId}][%d] getOutboundSchedulesSequenceOK  %+v", 200, o.Payload)
}

func (o *GetOutboundSchedulesSequenceOK) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences/{sequenceId}][%d] getOutboundSchedulesSequenceOK  %+v", 200, o.Payload)
}

func (o *GetOutboundSchedulesSequenceOK) GetPayload() *models.SequenceSchedule {
	return o.Payload
}

func (o *GetOutboundSchedulesSequenceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SequenceSchedule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesSequenceBadRequest creates a GetOutboundSchedulesSequenceBadRequest with default headers values
func NewGetOutboundSchedulesSequenceBadRequest() *GetOutboundSchedulesSequenceBadRequest {
	return &GetOutboundSchedulesSequenceBadRequest{}
}

/*
GetOutboundSchedulesSequenceBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundSchedulesSequenceBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules sequence bad request response has a 2xx status code
func (o *GetOutboundSchedulesSequenceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules sequence bad request response has a 3xx status code
func (o *GetOutboundSchedulesSequenceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules sequence bad request response has a 4xx status code
func (o *GetOutboundSchedulesSequenceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound schedules sequence bad request response has a 5xx status code
func (o *GetOutboundSchedulesSequenceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound schedules sequence bad request response a status code equal to that given
func (o *GetOutboundSchedulesSequenceBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetOutboundSchedulesSequenceBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences/{sequenceId}][%d] getOutboundSchedulesSequenceBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundSchedulesSequenceBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences/{sequenceId}][%d] getOutboundSchedulesSequenceBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundSchedulesSequenceBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesSequenceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesSequenceUnauthorized creates a GetOutboundSchedulesSequenceUnauthorized with default headers values
func NewGetOutboundSchedulesSequenceUnauthorized() *GetOutboundSchedulesSequenceUnauthorized {
	return &GetOutboundSchedulesSequenceUnauthorized{}
}

/*
GetOutboundSchedulesSequenceUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundSchedulesSequenceUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules sequence unauthorized response has a 2xx status code
func (o *GetOutboundSchedulesSequenceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules sequence unauthorized response has a 3xx status code
func (o *GetOutboundSchedulesSequenceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules sequence unauthorized response has a 4xx status code
func (o *GetOutboundSchedulesSequenceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound schedules sequence unauthorized response has a 5xx status code
func (o *GetOutboundSchedulesSequenceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound schedules sequence unauthorized response a status code equal to that given
func (o *GetOutboundSchedulesSequenceUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetOutboundSchedulesSequenceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences/{sequenceId}][%d] getOutboundSchedulesSequenceUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundSchedulesSequenceUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences/{sequenceId}][%d] getOutboundSchedulesSequenceUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundSchedulesSequenceUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesSequenceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesSequenceForbidden creates a GetOutboundSchedulesSequenceForbidden with default headers values
func NewGetOutboundSchedulesSequenceForbidden() *GetOutboundSchedulesSequenceForbidden {
	return &GetOutboundSchedulesSequenceForbidden{}
}

/*
GetOutboundSchedulesSequenceForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundSchedulesSequenceForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules sequence forbidden response has a 2xx status code
func (o *GetOutboundSchedulesSequenceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules sequence forbidden response has a 3xx status code
func (o *GetOutboundSchedulesSequenceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules sequence forbidden response has a 4xx status code
func (o *GetOutboundSchedulesSequenceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound schedules sequence forbidden response has a 5xx status code
func (o *GetOutboundSchedulesSequenceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound schedules sequence forbidden response a status code equal to that given
func (o *GetOutboundSchedulesSequenceForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetOutboundSchedulesSequenceForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences/{sequenceId}][%d] getOutboundSchedulesSequenceForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundSchedulesSequenceForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences/{sequenceId}][%d] getOutboundSchedulesSequenceForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundSchedulesSequenceForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesSequenceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesSequenceNotFound creates a GetOutboundSchedulesSequenceNotFound with default headers values
func NewGetOutboundSchedulesSequenceNotFound() *GetOutboundSchedulesSequenceNotFound {
	return &GetOutboundSchedulesSequenceNotFound{}
}

/*
GetOutboundSchedulesSequenceNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetOutboundSchedulesSequenceNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules sequence not found response has a 2xx status code
func (o *GetOutboundSchedulesSequenceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules sequence not found response has a 3xx status code
func (o *GetOutboundSchedulesSequenceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules sequence not found response has a 4xx status code
func (o *GetOutboundSchedulesSequenceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound schedules sequence not found response has a 5xx status code
func (o *GetOutboundSchedulesSequenceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound schedules sequence not found response a status code equal to that given
func (o *GetOutboundSchedulesSequenceNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetOutboundSchedulesSequenceNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences/{sequenceId}][%d] getOutboundSchedulesSequenceNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundSchedulesSequenceNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences/{sequenceId}][%d] getOutboundSchedulesSequenceNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundSchedulesSequenceNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesSequenceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesSequenceRequestTimeout creates a GetOutboundSchedulesSequenceRequestTimeout with default headers values
func NewGetOutboundSchedulesSequenceRequestTimeout() *GetOutboundSchedulesSequenceRequestTimeout {
	return &GetOutboundSchedulesSequenceRequestTimeout{}
}

/*
GetOutboundSchedulesSequenceRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetOutboundSchedulesSequenceRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules sequence request timeout response has a 2xx status code
func (o *GetOutboundSchedulesSequenceRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules sequence request timeout response has a 3xx status code
func (o *GetOutboundSchedulesSequenceRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules sequence request timeout response has a 4xx status code
func (o *GetOutboundSchedulesSequenceRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound schedules sequence request timeout response has a 5xx status code
func (o *GetOutboundSchedulesSequenceRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound schedules sequence request timeout response a status code equal to that given
func (o *GetOutboundSchedulesSequenceRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetOutboundSchedulesSequenceRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences/{sequenceId}][%d] getOutboundSchedulesSequenceRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundSchedulesSequenceRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences/{sequenceId}][%d] getOutboundSchedulesSequenceRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundSchedulesSequenceRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesSequenceRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesSequenceRequestEntityTooLarge creates a GetOutboundSchedulesSequenceRequestEntityTooLarge with default headers values
func NewGetOutboundSchedulesSequenceRequestEntityTooLarge() *GetOutboundSchedulesSequenceRequestEntityTooLarge {
	return &GetOutboundSchedulesSequenceRequestEntityTooLarge{}
}

/*
GetOutboundSchedulesSequenceRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetOutboundSchedulesSequenceRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules sequence request entity too large response has a 2xx status code
func (o *GetOutboundSchedulesSequenceRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules sequence request entity too large response has a 3xx status code
func (o *GetOutboundSchedulesSequenceRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules sequence request entity too large response has a 4xx status code
func (o *GetOutboundSchedulesSequenceRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound schedules sequence request entity too large response has a 5xx status code
func (o *GetOutboundSchedulesSequenceRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound schedules sequence request entity too large response a status code equal to that given
func (o *GetOutboundSchedulesSequenceRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetOutboundSchedulesSequenceRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences/{sequenceId}][%d] getOutboundSchedulesSequenceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundSchedulesSequenceRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences/{sequenceId}][%d] getOutboundSchedulesSequenceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundSchedulesSequenceRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesSequenceRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesSequenceUnsupportedMediaType creates a GetOutboundSchedulesSequenceUnsupportedMediaType with default headers values
func NewGetOutboundSchedulesSequenceUnsupportedMediaType() *GetOutboundSchedulesSequenceUnsupportedMediaType {
	return &GetOutboundSchedulesSequenceUnsupportedMediaType{}
}

/*
GetOutboundSchedulesSequenceUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundSchedulesSequenceUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules sequence unsupported media type response has a 2xx status code
func (o *GetOutboundSchedulesSequenceUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules sequence unsupported media type response has a 3xx status code
func (o *GetOutboundSchedulesSequenceUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules sequence unsupported media type response has a 4xx status code
func (o *GetOutboundSchedulesSequenceUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound schedules sequence unsupported media type response has a 5xx status code
func (o *GetOutboundSchedulesSequenceUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound schedules sequence unsupported media type response a status code equal to that given
func (o *GetOutboundSchedulesSequenceUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetOutboundSchedulesSequenceUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences/{sequenceId}][%d] getOutboundSchedulesSequenceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundSchedulesSequenceUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences/{sequenceId}][%d] getOutboundSchedulesSequenceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundSchedulesSequenceUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesSequenceUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesSequenceTooManyRequests creates a GetOutboundSchedulesSequenceTooManyRequests with default headers values
func NewGetOutboundSchedulesSequenceTooManyRequests() *GetOutboundSchedulesSequenceTooManyRequests {
	return &GetOutboundSchedulesSequenceTooManyRequests{}
}

/*
GetOutboundSchedulesSequenceTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetOutboundSchedulesSequenceTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules sequence too many requests response has a 2xx status code
func (o *GetOutboundSchedulesSequenceTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules sequence too many requests response has a 3xx status code
func (o *GetOutboundSchedulesSequenceTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules sequence too many requests response has a 4xx status code
func (o *GetOutboundSchedulesSequenceTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound schedules sequence too many requests response has a 5xx status code
func (o *GetOutboundSchedulesSequenceTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound schedules sequence too many requests response a status code equal to that given
func (o *GetOutboundSchedulesSequenceTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetOutboundSchedulesSequenceTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences/{sequenceId}][%d] getOutboundSchedulesSequenceTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundSchedulesSequenceTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences/{sequenceId}][%d] getOutboundSchedulesSequenceTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundSchedulesSequenceTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesSequenceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesSequenceInternalServerError creates a GetOutboundSchedulesSequenceInternalServerError with default headers values
func NewGetOutboundSchedulesSequenceInternalServerError() *GetOutboundSchedulesSequenceInternalServerError {
	return &GetOutboundSchedulesSequenceInternalServerError{}
}

/*
GetOutboundSchedulesSequenceInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundSchedulesSequenceInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules sequence internal server error response has a 2xx status code
func (o *GetOutboundSchedulesSequenceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules sequence internal server error response has a 3xx status code
func (o *GetOutboundSchedulesSequenceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules sequence internal server error response has a 4xx status code
func (o *GetOutboundSchedulesSequenceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound schedules sequence internal server error response has a 5xx status code
func (o *GetOutboundSchedulesSequenceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound schedules sequence internal server error response a status code equal to that given
func (o *GetOutboundSchedulesSequenceInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetOutboundSchedulesSequenceInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences/{sequenceId}][%d] getOutboundSchedulesSequenceInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundSchedulesSequenceInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences/{sequenceId}][%d] getOutboundSchedulesSequenceInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundSchedulesSequenceInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesSequenceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesSequenceServiceUnavailable creates a GetOutboundSchedulesSequenceServiceUnavailable with default headers values
func NewGetOutboundSchedulesSequenceServiceUnavailable() *GetOutboundSchedulesSequenceServiceUnavailable {
	return &GetOutboundSchedulesSequenceServiceUnavailable{}
}

/*
GetOutboundSchedulesSequenceServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundSchedulesSequenceServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules sequence service unavailable response has a 2xx status code
func (o *GetOutboundSchedulesSequenceServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules sequence service unavailable response has a 3xx status code
func (o *GetOutboundSchedulesSequenceServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules sequence service unavailable response has a 4xx status code
func (o *GetOutboundSchedulesSequenceServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound schedules sequence service unavailable response has a 5xx status code
func (o *GetOutboundSchedulesSequenceServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound schedules sequence service unavailable response a status code equal to that given
func (o *GetOutboundSchedulesSequenceServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetOutboundSchedulesSequenceServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences/{sequenceId}][%d] getOutboundSchedulesSequenceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundSchedulesSequenceServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences/{sequenceId}][%d] getOutboundSchedulesSequenceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundSchedulesSequenceServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesSequenceServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesSequenceGatewayTimeout creates a GetOutboundSchedulesSequenceGatewayTimeout with default headers values
func NewGetOutboundSchedulesSequenceGatewayTimeout() *GetOutboundSchedulesSequenceGatewayTimeout {
	return &GetOutboundSchedulesSequenceGatewayTimeout{}
}

/*
GetOutboundSchedulesSequenceGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetOutboundSchedulesSequenceGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules sequence gateway timeout response has a 2xx status code
func (o *GetOutboundSchedulesSequenceGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules sequence gateway timeout response has a 3xx status code
func (o *GetOutboundSchedulesSequenceGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules sequence gateway timeout response has a 4xx status code
func (o *GetOutboundSchedulesSequenceGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound schedules sequence gateway timeout response has a 5xx status code
func (o *GetOutboundSchedulesSequenceGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound schedules sequence gateway timeout response a status code equal to that given
func (o *GetOutboundSchedulesSequenceGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetOutboundSchedulesSequenceGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences/{sequenceId}][%d] getOutboundSchedulesSequenceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundSchedulesSequenceGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences/{sequenceId}][%d] getOutboundSchedulesSequenceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundSchedulesSequenceGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesSequenceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

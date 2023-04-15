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

// GetOutboundSequenceReader is a Reader for the GetOutboundSequence structure.
type GetOutboundSequenceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutboundSequenceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutboundSequenceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutboundSequenceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutboundSequenceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutboundSequenceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutboundSequenceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOutboundSequenceRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOutboundSequenceRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOutboundSequenceUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOutboundSequenceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOutboundSequenceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOutboundSequenceServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOutboundSequenceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOutboundSequenceOK creates a GetOutboundSequenceOK with default headers values
func NewGetOutboundSequenceOK() *GetOutboundSequenceOK {
	return &GetOutboundSequenceOK{}
}

/*
GetOutboundSequenceOK describes a response with status code 200, with default header values.

successful operation
*/
type GetOutboundSequenceOK struct {
	Payload *models.CampaignSequence
}

// IsSuccess returns true when this get outbound sequence o k response has a 2xx status code
func (o *GetOutboundSequenceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get outbound sequence o k response has a 3xx status code
func (o *GetOutboundSequenceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound sequence o k response has a 4xx status code
func (o *GetOutboundSequenceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound sequence o k response has a 5xx status code
func (o *GetOutboundSequenceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound sequence o k response a status code equal to that given
func (o *GetOutboundSequenceOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOutboundSequenceOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/sequences/{sequenceId}][%d] getOutboundSequenceOK  %+v", 200, o.Payload)
}

func (o *GetOutboundSequenceOK) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/sequences/{sequenceId}][%d] getOutboundSequenceOK  %+v", 200, o.Payload)
}

func (o *GetOutboundSequenceOK) GetPayload() *models.CampaignSequence {
	return o.Payload
}

func (o *GetOutboundSequenceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CampaignSequence)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSequenceBadRequest creates a GetOutboundSequenceBadRequest with default headers values
func NewGetOutboundSequenceBadRequest() *GetOutboundSequenceBadRequest {
	return &GetOutboundSequenceBadRequest{}
}

/*
GetOutboundSequenceBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundSequenceBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound sequence bad request response has a 2xx status code
func (o *GetOutboundSequenceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound sequence bad request response has a 3xx status code
func (o *GetOutboundSequenceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound sequence bad request response has a 4xx status code
func (o *GetOutboundSequenceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound sequence bad request response has a 5xx status code
func (o *GetOutboundSequenceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound sequence bad request response a status code equal to that given
func (o *GetOutboundSequenceBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetOutboundSequenceBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/sequences/{sequenceId}][%d] getOutboundSequenceBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundSequenceBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/sequences/{sequenceId}][%d] getOutboundSequenceBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundSequenceBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSequenceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSequenceUnauthorized creates a GetOutboundSequenceUnauthorized with default headers values
func NewGetOutboundSequenceUnauthorized() *GetOutboundSequenceUnauthorized {
	return &GetOutboundSequenceUnauthorized{}
}

/*
GetOutboundSequenceUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundSequenceUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound sequence unauthorized response has a 2xx status code
func (o *GetOutboundSequenceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound sequence unauthorized response has a 3xx status code
func (o *GetOutboundSequenceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound sequence unauthorized response has a 4xx status code
func (o *GetOutboundSequenceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound sequence unauthorized response has a 5xx status code
func (o *GetOutboundSequenceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound sequence unauthorized response a status code equal to that given
func (o *GetOutboundSequenceUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetOutboundSequenceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/sequences/{sequenceId}][%d] getOutboundSequenceUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundSequenceUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/sequences/{sequenceId}][%d] getOutboundSequenceUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundSequenceUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSequenceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSequenceForbidden creates a GetOutboundSequenceForbidden with default headers values
func NewGetOutboundSequenceForbidden() *GetOutboundSequenceForbidden {
	return &GetOutboundSequenceForbidden{}
}

/*
GetOutboundSequenceForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundSequenceForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound sequence forbidden response has a 2xx status code
func (o *GetOutboundSequenceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound sequence forbidden response has a 3xx status code
func (o *GetOutboundSequenceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound sequence forbidden response has a 4xx status code
func (o *GetOutboundSequenceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound sequence forbidden response has a 5xx status code
func (o *GetOutboundSequenceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound sequence forbidden response a status code equal to that given
func (o *GetOutboundSequenceForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetOutboundSequenceForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/sequences/{sequenceId}][%d] getOutboundSequenceForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundSequenceForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/sequences/{sequenceId}][%d] getOutboundSequenceForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundSequenceForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSequenceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSequenceNotFound creates a GetOutboundSequenceNotFound with default headers values
func NewGetOutboundSequenceNotFound() *GetOutboundSequenceNotFound {
	return &GetOutboundSequenceNotFound{}
}

/*
GetOutboundSequenceNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetOutboundSequenceNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound sequence not found response has a 2xx status code
func (o *GetOutboundSequenceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound sequence not found response has a 3xx status code
func (o *GetOutboundSequenceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound sequence not found response has a 4xx status code
func (o *GetOutboundSequenceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound sequence not found response has a 5xx status code
func (o *GetOutboundSequenceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound sequence not found response a status code equal to that given
func (o *GetOutboundSequenceNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetOutboundSequenceNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/sequences/{sequenceId}][%d] getOutboundSequenceNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundSequenceNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/sequences/{sequenceId}][%d] getOutboundSequenceNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundSequenceNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSequenceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSequenceRequestTimeout creates a GetOutboundSequenceRequestTimeout with default headers values
func NewGetOutboundSequenceRequestTimeout() *GetOutboundSequenceRequestTimeout {
	return &GetOutboundSequenceRequestTimeout{}
}

/*
GetOutboundSequenceRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetOutboundSequenceRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound sequence request timeout response has a 2xx status code
func (o *GetOutboundSequenceRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound sequence request timeout response has a 3xx status code
func (o *GetOutboundSequenceRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound sequence request timeout response has a 4xx status code
func (o *GetOutboundSequenceRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound sequence request timeout response has a 5xx status code
func (o *GetOutboundSequenceRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound sequence request timeout response a status code equal to that given
func (o *GetOutboundSequenceRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetOutboundSequenceRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/sequences/{sequenceId}][%d] getOutboundSequenceRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundSequenceRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/sequences/{sequenceId}][%d] getOutboundSequenceRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundSequenceRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSequenceRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSequenceRequestEntityTooLarge creates a GetOutboundSequenceRequestEntityTooLarge with default headers values
func NewGetOutboundSequenceRequestEntityTooLarge() *GetOutboundSequenceRequestEntityTooLarge {
	return &GetOutboundSequenceRequestEntityTooLarge{}
}

/*
GetOutboundSequenceRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetOutboundSequenceRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound sequence request entity too large response has a 2xx status code
func (o *GetOutboundSequenceRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound sequence request entity too large response has a 3xx status code
func (o *GetOutboundSequenceRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound sequence request entity too large response has a 4xx status code
func (o *GetOutboundSequenceRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound sequence request entity too large response has a 5xx status code
func (o *GetOutboundSequenceRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound sequence request entity too large response a status code equal to that given
func (o *GetOutboundSequenceRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetOutboundSequenceRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/sequences/{sequenceId}][%d] getOutboundSequenceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundSequenceRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/sequences/{sequenceId}][%d] getOutboundSequenceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundSequenceRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSequenceRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSequenceUnsupportedMediaType creates a GetOutboundSequenceUnsupportedMediaType with default headers values
func NewGetOutboundSequenceUnsupportedMediaType() *GetOutboundSequenceUnsupportedMediaType {
	return &GetOutboundSequenceUnsupportedMediaType{}
}

/*
GetOutboundSequenceUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundSequenceUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound sequence unsupported media type response has a 2xx status code
func (o *GetOutboundSequenceUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound sequence unsupported media type response has a 3xx status code
func (o *GetOutboundSequenceUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound sequence unsupported media type response has a 4xx status code
func (o *GetOutboundSequenceUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound sequence unsupported media type response has a 5xx status code
func (o *GetOutboundSequenceUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound sequence unsupported media type response a status code equal to that given
func (o *GetOutboundSequenceUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetOutboundSequenceUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/sequences/{sequenceId}][%d] getOutboundSequenceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundSequenceUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/sequences/{sequenceId}][%d] getOutboundSequenceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundSequenceUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSequenceUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSequenceTooManyRequests creates a GetOutboundSequenceTooManyRequests with default headers values
func NewGetOutboundSequenceTooManyRequests() *GetOutboundSequenceTooManyRequests {
	return &GetOutboundSequenceTooManyRequests{}
}

/*
GetOutboundSequenceTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetOutboundSequenceTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound sequence too many requests response has a 2xx status code
func (o *GetOutboundSequenceTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound sequence too many requests response has a 3xx status code
func (o *GetOutboundSequenceTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound sequence too many requests response has a 4xx status code
func (o *GetOutboundSequenceTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound sequence too many requests response has a 5xx status code
func (o *GetOutboundSequenceTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound sequence too many requests response a status code equal to that given
func (o *GetOutboundSequenceTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetOutboundSequenceTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/sequences/{sequenceId}][%d] getOutboundSequenceTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundSequenceTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/sequences/{sequenceId}][%d] getOutboundSequenceTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundSequenceTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSequenceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSequenceInternalServerError creates a GetOutboundSequenceInternalServerError with default headers values
func NewGetOutboundSequenceInternalServerError() *GetOutboundSequenceInternalServerError {
	return &GetOutboundSequenceInternalServerError{}
}

/*
GetOutboundSequenceInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundSequenceInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound sequence internal server error response has a 2xx status code
func (o *GetOutboundSequenceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound sequence internal server error response has a 3xx status code
func (o *GetOutboundSequenceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound sequence internal server error response has a 4xx status code
func (o *GetOutboundSequenceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound sequence internal server error response has a 5xx status code
func (o *GetOutboundSequenceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound sequence internal server error response a status code equal to that given
func (o *GetOutboundSequenceInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetOutboundSequenceInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/sequences/{sequenceId}][%d] getOutboundSequenceInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundSequenceInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/sequences/{sequenceId}][%d] getOutboundSequenceInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundSequenceInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSequenceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSequenceServiceUnavailable creates a GetOutboundSequenceServiceUnavailable with default headers values
func NewGetOutboundSequenceServiceUnavailable() *GetOutboundSequenceServiceUnavailable {
	return &GetOutboundSequenceServiceUnavailable{}
}

/*
GetOutboundSequenceServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundSequenceServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound sequence service unavailable response has a 2xx status code
func (o *GetOutboundSequenceServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound sequence service unavailable response has a 3xx status code
func (o *GetOutboundSequenceServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound sequence service unavailable response has a 4xx status code
func (o *GetOutboundSequenceServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound sequence service unavailable response has a 5xx status code
func (o *GetOutboundSequenceServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound sequence service unavailable response a status code equal to that given
func (o *GetOutboundSequenceServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetOutboundSequenceServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/sequences/{sequenceId}][%d] getOutboundSequenceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundSequenceServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/sequences/{sequenceId}][%d] getOutboundSequenceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundSequenceServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSequenceServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSequenceGatewayTimeout creates a GetOutboundSequenceGatewayTimeout with default headers values
func NewGetOutboundSequenceGatewayTimeout() *GetOutboundSequenceGatewayTimeout {
	return &GetOutboundSequenceGatewayTimeout{}
}

/*
GetOutboundSequenceGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetOutboundSequenceGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound sequence gateway timeout response has a 2xx status code
func (o *GetOutboundSequenceGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound sequence gateway timeout response has a 3xx status code
func (o *GetOutboundSequenceGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound sequence gateway timeout response has a 4xx status code
func (o *GetOutboundSequenceGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound sequence gateway timeout response has a 5xx status code
func (o *GetOutboundSequenceGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound sequence gateway timeout response a status code equal to that given
func (o *GetOutboundSequenceGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetOutboundSequenceGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/sequences/{sequenceId}][%d] getOutboundSequenceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundSequenceGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/sequences/{sequenceId}][%d] getOutboundSequenceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundSequenceGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSequenceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

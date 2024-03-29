// Code generated by go-swagger; DO NOT EDIT.

package journey

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetJourneyOutcomeReader is a Reader for the GetJourneyOutcome structure.
type GetJourneyOutcomeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetJourneyOutcomeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetJourneyOutcomeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetJourneyOutcomeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetJourneyOutcomeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetJourneyOutcomeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetJourneyOutcomeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetJourneyOutcomeRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetJourneyOutcomeRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetJourneyOutcomeUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetJourneyOutcomeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetJourneyOutcomeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetJourneyOutcomeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetJourneyOutcomeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetJourneyOutcomeOK creates a GetJourneyOutcomeOK with default headers values
func NewGetJourneyOutcomeOK() *GetJourneyOutcomeOK {
	return &GetJourneyOutcomeOK{}
}

/*
GetJourneyOutcomeOK describes a response with status code 200, with default header values.

successful operation
*/
type GetJourneyOutcomeOK struct {
	Payload *models.Outcome
}

// IsSuccess returns true when this get journey outcome o k response has a 2xx status code
func (o *GetJourneyOutcomeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get journey outcome o k response has a 3xx status code
func (o *GetJourneyOutcomeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey outcome o k response has a 4xx status code
func (o *GetJourneyOutcomeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get journey outcome o k response has a 5xx status code
func (o *GetJourneyOutcomeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey outcome o k response a status code equal to that given
func (o *GetJourneyOutcomeOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetJourneyOutcomeOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/{outcomeId}][%d] getJourneyOutcomeOK  %+v", 200, o.Payload)
}

func (o *GetJourneyOutcomeOK) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/{outcomeId}][%d] getJourneyOutcomeOK  %+v", 200, o.Payload)
}

func (o *GetJourneyOutcomeOK) GetPayload() *models.Outcome {
	return o.Payload
}

func (o *GetJourneyOutcomeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Outcome)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyOutcomeBadRequest creates a GetJourneyOutcomeBadRequest with default headers values
func NewGetJourneyOutcomeBadRequest() *GetJourneyOutcomeBadRequest {
	return &GetJourneyOutcomeBadRequest{}
}

/*
GetJourneyOutcomeBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetJourneyOutcomeBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey outcome bad request response has a 2xx status code
func (o *GetJourneyOutcomeBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey outcome bad request response has a 3xx status code
func (o *GetJourneyOutcomeBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey outcome bad request response has a 4xx status code
func (o *GetJourneyOutcomeBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey outcome bad request response has a 5xx status code
func (o *GetJourneyOutcomeBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey outcome bad request response a status code equal to that given
func (o *GetJourneyOutcomeBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetJourneyOutcomeBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/{outcomeId}][%d] getJourneyOutcomeBadRequest  %+v", 400, o.Payload)
}

func (o *GetJourneyOutcomeBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/{outcomeId}][%d] getJourneyOutcomeBadRequest  %+v", 400, o.Payload)
}

func (o *GetJourneyOutcomeBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyOutcomeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyOutcomeUnauthorized creates a GetJourneyOutcomeUnauthorized with default headers values
func NewGetJourneyOutcomeUnauthorized() *GetJourneyOutcomeUnauthorized {
	return &GetJourneyOutcomeUnauthorized{}
}

/*
GetJourneyOutcomeUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetJourneyOutcomeUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey outcome unauthorized response has a 2xx status code
func (o *GetJourneyOutcomeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey outcome unauthorized response has a 3xx status code
func (o *GetJourneyOutcomeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey outcome unauthorized response has a 4xx status code
func (o *GetJourneyOutcomeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey outcome unauthorized response has a 5xx status code
func (o *GetJourneyOutcomeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey outcome unauthorized response a status code equal to that given
func (o *GetJourneyOutcomeUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetJourneyOutcomeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/{outcomeId}][%d] getJourneyOutcomeUnauthorized  %+v", 401, o.Payload)
}

func (o *GetJourneyOutcomeUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/{outcomeId}][%d] getJourneyOutcomeUnauthorized  %+v", 401, o.Payload)
}

func (o *GetJourneyOutcomeUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyOutcomeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyOutcomeForbidden creates a GetJourneyOutcomeForbidden with default headers values
func NewGetJourneyOutcomeForbidden() *GetJourneyOutcomeForbidden {
	return &GetJourneyOutcomeForbidden{}
}

/*
GetJourneyOutcomeForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetJourneyOutcomeForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey outcome forbidden response has a 2xx status code
func (o *GetJourneyOutcomeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey outcome forbidden response has a 3xx status code
func (o *GetJourneyOutcomeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey outcome forbidden response has a 4xx status code
func (o *GetJourneyOutcomeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey outcome forbidden response has a 5xx status code
func (o *GetJourneyOutcomeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey outcome forbidden response a status code equal to that given
func (o *GetJourneyOutcomeForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetJourneyOutcomeForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/{outcomeId}][%d] getJourneyOutcomeForbidden  %+v", 403, o.Payload)
}

func (o *GetJourneyOutcomeForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/{outcomeId}][%d] getJourneyOutcomeForbidden  %+v", 403, o.Payload)
}

func (o *GetJourneyOutcomeForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyOutcomeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyOutcomeNotFound creates a GetJourneyOutcomeNotFound with default headers values
func NewGetJourneyOutcomeNotFound() *GetJourneyOutcomeNotFound {
	return &GetJourneyOutcomeNotFound{}
}

/*
GetJourneyOutcomeNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetJourneyOutcomeNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey outcome not found response has a 2xx status code
func (o *GetJourneyOutcomeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey outcome not found response has a 3xx status code
func (o *GetJourneyOutcomeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey outcome not found response has a 4xx status code
func (o *GetJourneyOutcomeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey outcome not found response has a 5xx status code
func (o *GetJourneyOutcomeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey outcome not found response a status code equal to that given
func (o *GetJourneyOutcomeNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetJourneyOutcomeNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/{outcomeId}][%d] getJourneyOutcomeNotFound  %+v", 404, o.Payload)
}

func (o *GetJourneyOutcomeNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/{outcomeId}][%d] getJourneyOutcomeNotFound  %+v", 404, o.Payload)
}

func (o *GetJourneyOutcomeNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyOutcomeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyOutcomeRequestTimeout creates a GetJourneyOutcomeRequestTimeout with default headers values
func NewGetJourneyOutcomeRequestTimeout() *GetJourneyOutcomeRequestTimeout {
	return &GetJourneyOutcomeRequestTimeout{}
}

/*
GetJourneyOutcomeRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetJourneyOutcomeRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey outcome request timeout response has a 2xx status code
func (o *GetJourneyOutcomeRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey outcome request timeout response has a 3xx status code
func (o *GetJourneyOutcomeRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey outcome request timeout response has a 4xx status code
func (o *GetJourneyOutcomeRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey outcome request timeout response has a 5xx status code
func (o *GetJourneyOutcomeRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey outcome request timeout response a status code equal to that given
func (o *GetJourneyOutcomeRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetJourneyOutcomeRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/{outcomeId}][%d] getJourneyOutcomeRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetJourneyOutcomeRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/{outcomeId}][%d] getJourneyOutcomeRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetJourneyOutcomeRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyOutcomeRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyOutcomeRequestEntityTooLarge creates a GetJourneyOutcomeRequestEntityTooLarge with default headers values
func NewGetJourneyOutcomeRequestEntityTooLarge() *GetJourneyOutcomeRequestEntityTooLarge {
	return &GetJourneyOutcomeRequestEntityTooLarge{}
}

/*
GetJourneyOutcomeRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetJourneyOutcomeRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey outcome request entity too large response has a 2xx status code
func (o *GetJourneyOutcomeRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey outcome request entity too large response has a 3xx status code
func (o *GetJourneyOutcomeRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey outcome request entity too large response has a 4xx status code
func (o *GetJourneyOutcomeRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey outcome request entity too large response has a 5xx status code
func (o *GetJourneyOutcomeRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey outcome request entity too large response a status code equal to that given
func (o *GetJourneyOutcomeRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetJourneyOutcomeRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/{outcomeId}][%d] getJourneyOutcomeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetJourneyOutcomeRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/{outcomeId}][%d] getJourneyOutcomeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetJourneyOutcomeRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyOutcomeRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyOutcomeUnsupportedMediaType creates a GetJourneyOutcomeUnsupportedMediaType with default headers values
func NewGetJourneyOutcomeUnsupportedMediaType() *GetJourneyOutcomeUnsupportedMediaType {
	return &GetJourneyOutcomeUnsupportedMediaType{}
}

/*
GetJourneyOutcomeUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetJourneyOutcomeUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey outcome unsupported media type response has a 2xx status code
func (o *GetJourneyOutcomeUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey outcome unsupported media type response has a 3xx status code
func (o *GetJourneyOutcomeUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey outcome unsupported media type response has a 4xx status code
func (o *GetJourneyOutcomeUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey outcome unsupported media type response has a 5xx status code
func (o *GetJourneyOutcomeUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey outcome unsupported media type response a status code equal to that given
func (o *GetJourneyOutcomeUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetJourneyOutcomeUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/{outcomeId}][%d] getJourneyOutcomeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetJourneyOutcomeUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/{outcomeId}][%d] getJourneyOutcomeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetJourneyOutcomeUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyOutcomeUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyOutcomeTooManyRequests creates a GetJourneyOutcomeTooManyRequests with default headers values
func NewGetJourneyOutcomeTooManyRequests() *GetJourneyOutcomeTooManyRequests {
	return &GetJourneyOutcomeTooManyRequests{}
}

/*
GetJourneyOutcomeTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetJourneyOutcomeTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey outcome too many requests response has a 2xx status code
func (o *GetJourneyOutcomeTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey outcome too many requests response has a 3xx status code
func (o *GetJourneyOutcomeTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey outcome too many requests response has a 4xx status code
func (o *GetJourneyOutcomeTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey outcome too many requests response has a 5xx status code
func (o *GetJourneyOutcomeTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey outcome too many requests response a status code equal to that given
func (o *GetJourneyOutcomeTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetJourneyOutcomeTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/{outcomeId}][%d] getJourneyOutcomeTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetJourneyOutcomeTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/{outcomeId}][%d] getJourneyOutcomeTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetJourneyOutcomeTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyOutcomeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyOutcomeInternalServerError creates a GetJourneyOutcomeInternalServerError with default headers values
func NewGetJourneyOutcomeInternalServerError() *GetJourneyOutcomeInternalServerError {
	return &GetJourneyOutcomeInternalServerError{}
}

/*
GetJourneyOutcomeInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetJourneyOutcomeInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey outcome internal server error response has a 2xx status code
func (o *GetJourneyOutcomeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey outcome internal server error response has a 3xx status code
func (o *GetJourneyOutcomeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey outcome internal server error response has a 4xx status code
func (o *GetJourneyOutcomeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get journey outcome internal server error response has a 5xx status code
func (o *GetJourneyOutcomeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get journey outcome internal server error response a status code equal to that given
func (o *GetJourneyOutcomeInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetJourneyOutcomeInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/{outcomeId}][%d] getJourneyOutcomeInternalServerError  %+v", 500, o.Payload)
}

func (o *GetJourneyOutcomeInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/{outcomeId}][%d] getJourneyOutcomeInternalServerError  %+v", 500, o.Payload)
}

func (o *GetJourneyOutcomeInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyOutcomeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyOutcomeServiceUnavailable creates a GetJourneyOutcomeServiceUnavailable with default headers values
func NewGetJourneyOutcomeServiceUnavailable() *GetJourneyOutcomeServiceUnavailable {
	return &GetJourneyOutcomeServiceUnavailable{}
}

/*
GetJourneyOutcomeServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetJourneyOutcomeServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey outcome service unavailable response has a 2xx status code
func (o *GetJourneyOutcomeServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey outcome service unavailable response has a 3xx status code
func (o *GetJourneyOutcomeServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey outcome service unavailable response has a 4xx status code
func (o *GetJourneyOutcomeServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get journey outcome service unavailable response has a 5xx status code
func (o *GetJourneyOutcomeServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get journey outcome service unavailable response a status code equal to that given
func (o *GetJourneyOutcomeServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetJourneyOutcomeServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/{outcomeId}][%d] getJourneyOutcomeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetJourneyOutcomeServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/{outcomeId}][%d] getJourneyOutcomeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetJourneyOutcomeServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyOutcomeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyOutcomeGatewayTimeout creates a GetJourneyOutcomeGatewayTimeout with default headers values
func NewGetJourneyOutcomeGatewayTimeout() *GetJourneyOutcomeGatewayTimeout {
	return &GetJourneyOutcomeGatewayTimeout{}
}

/*
GetJourneyOutcomeGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetJourneyOutcomeGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey outcome gateway timeout response has a 2xx status code
func (o *GetJourneyOutcomeGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey outcome gateway timeout response has a 3xx status code
func (o *GetJourneyOutcomeGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey outcome gateway timeout response has a 4xx status code
func (o *GetJourneyOutcomeGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get journey outcome gateway timeout response has a 5xx status code
func (o *GetJourneyOutcomeGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get journey outcome gateway timeout response a status code equal to that given
func (o *GetJourneyOutcomeGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetJourneyOutcomeGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/{outcomeId}][%d] getJourneyOutcomeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetJourneyOutcomeGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/{outcomeId}][%d] getJourneyOutcomeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetJourneyOutcomeGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyOutcomeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

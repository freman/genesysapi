// Code generated by go-swagger; DO NOT EDIT.

package architect

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PutFlowsOutcomeReader is a Reader for the PutFlowsOutcome structure.
type PutFlowsOutcomeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutFlowsOutcomeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutFlowsOutcomeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutFlowsOutcomeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutFlowsOutcomeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutFlowsOutcomeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutFlowsOutcomeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPutFlowsOutcomeMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutFlowsOutcomeRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutFlowsOutcomeConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutFlowsOutcomeRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutFlowsOutcomeUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutFlowsOutcomeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutFlowsOutcomeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutFlowsOutcomeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutFlowsOutcomeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutFlowsOutcomeOK creates a PutFlowsOutcomeOK with default headers values
func NewPutFlowsOutcomeOK() *PutFlowsOutcomeOK {
	return &PutFlowsOutcomeOK{}
}

/*
PutFlowsOutcomeOK describes a response with status code 200, with default header values.

successful operation
*/
type PutFlowsOutcomeOK struct {
	Payload *models.Operation
}

// IsSuccess returns true when this put flows outcome o k response has a 2xx status code
func (o *PutFlowsOutcomeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put flows outcome o k response has a 3xx status code
func (o *PutFlowsOutcomeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put flows outcome o k response has a 4xx status code
func (o *PutFlowsOutcomeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put flows outcome o k response has a 5xx status code
func (o *PutFlowsOutcomeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put flows outcome o k response a status code equal to that given
func (o *PutFlowsOutcomeOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutFlowsOutcomeOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/outcomes/{flowOutcomeId}][%d] putFlowsOutcomeOK  %+v", 200, o.Payload)
}

func (o *PutFlowsOutcomeOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/flows/outcomes/{flowOutcomeId}][%d] putFlowsOutcomeOK  %+v", 200, o.Payload)
}

func (o *PutFlowsOutcomeOK) GetPayload() *models.Operation {
	return o.Payload
}

func (o *PutFlowsOutcomeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Operation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowsOutcomeBadRequest creates a PutFlowsOutcomeBadRequest with default headers values
func NewPutFlowsOutcomeBadRequest() *PutFlowsOutcomeBadRequest {
	return &PutFlowsOutcomeBadRequest{}
}

/*
PutFlowsOutcomeBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutFlowsOutcomeBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put flows outcome bad request response has a 2xx status code
func (o *PutFlowsOutcomeBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put flows outcome bad request response has a 3xx status code
func (o *PutFlowsOutcomeBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put flows outcome bad request response has a 4xx status code
func (o *PutFlowsOutcomeBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put flows outcome bad request response has a 5xx status code
func (o *PutFlowsOutcomeBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put flows outcome bad request response a status code equal to that given
func (o *PutFlowsOutcomeBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutFlowsOutcomeBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/outcomes/{flowOutcomeId}][%d] putFlowsOutcomeBadRequest  %+v", 400, o.Payload)
}

func (o *PutFlowsOutcomeBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v2/flows/outcomes/{flowOutcomeId}][%d] putFlowsOutcomeBadRequest  %+v", 400, o.Payload)
}

func (o *PutFlowsOutcomeBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowsOutcomeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowsOutcomeUnauthorized creates a PutFlowsOutcomeUnauthorized with default headers values
func NewPutFlowsOutcomeUnauthorized() *PutFlowsOutcomeUnauthorized {
	return &PutFlowsOutcomeUnauthorized{}
}

/*
PutFlowsOutcomeUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PutFlowsOutcomeUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put flows outcome unauthorized response has a 2xx status code
func (o *PutFlowsOutcomeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put flows outcome unauthorized response has a 3xx status code
func (o *PutFlowsOutcomeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put flows outcome unauthorized response has a 4xx status code
func (o *PutFlowsOutcomeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put flows outcome unauthorized response has a 5xx status code
func (o *PutFlowsOutcomeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put flows outcome unauthorized response a status code equal to that given
func (o *PutFlowsOutcomeUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PutFlowsOutcomeUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/outcomes/{flowOutcomeId}][%d] putFlowsOutcomeUnauthorized  %+v", 401, o.Payload)
}

func (o *PutFlowsOutcomeUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/flows/outcomes/{flowOutcomeId}][%d] putFlowsOutcomeUnauthorized  %+v", 401, o.Payload)
}

func (o *PutFlowsOutcomeUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowsOutcomeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowsOutcomeForbidden creates a PutFlowsOutcomeForbidden with default headers values
func NewPutFlowsOutcomeForbidden() *PutFlowsOutcomeForbidden {
	return &PutFlowsOutcomeForbidden{}
}

/*
PutFlowsOutcomeForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PutFlowsOutcomeForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put flows outcome forbidden response has a 2xx status code
func (o *PutFlowsOutcomeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put flows outcome forbidden response has a 3xx status code
func (o *PutFlowsOutcomeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put flows outcome forbidden response has a 4xx status code
func (o *PutFlowsOutcomeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put flows outcome forbidden response has a 5xx status code
func (o *PutFlowsOutcomeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put flows outcome forbidden response a status code equal to that given
func (o *PutFlowsOutcomeForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutFlowsOutcomeForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/outcomes/{flowOutcomeId}][%d] putFlowsOutcomeForbidden  %+v", 403, o.Payload)
}

func (o *PutFlowsOutcomeForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/flows/outcomes/{flowOutcomeId}][%d] putFlowsOutcomeForbidden  %+v", 403, o.Payload)
}

func (o *PutFlowsOutcomeForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowsOutcomeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowsOutcomeNotFound creates a PutFlowsOutcomeNotFound with default headers values
func NewPutFlowsOutcomeNotFound() *PutFlowsOutcomeNotFound {
	return &PutFlowsOutcomeNotFound{}
}

/*
PutFlowsOutcomeNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PutFlowsOutcomeNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put flows outcome not found response has a 2xx status code
func (o *PutFlowsOutcomeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put flows outcome not found response has a 3xx status code
func (o *PutFlowsOutcomeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put flows outcome not found response has a 4xx status code
func (o *PutFlowsOutcomeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put flows outcome not found response has a 5xx status code
func (o *PutFlowsOutcomeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put flows outcome not found response a status code equal to that given
func (o *PutFlowsOutcomeNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutFlowsOutcomeNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/outcomes/{flowOutcomeId}][%d] putFlowsOutcomeNotFound  %+v", 404, o.Payload)
}

func (o *PutFlowsOutcomeNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v2/flows/outcomes/{flowOutcomeId}][%d] putFlowsOutcomeNotFound  %+v", 404, o.Payload)
}

func (o *PutFlowsOutcomeNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowsOutcomeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowsOutcomeMethodNotAllowed creates a PutFlowsOutcomeMethodNotAllowed with default headers values
func NewPutFlowsOutcomeMethodNotAllowed() *PutFlowsOutcomeMethodNotAllowed {
	return &PutFlowsOutcomeMethodNotAllowed{}
}

/*
PutFlowsOutcomeMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type PutFlowsOutcomeMethodNotAllowed struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put flows outcome method not allowed response has a 2xx status code
func (o *PutFlowsOutcomeMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put flows outcome method not allowed response has a 3xx status code
func (o *PutFlowsOutcomeMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put flows outcome method not allowed response has a 4xx status code
func (o *PutFlowsOutcomeMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this put flows outcome method not allowed response has a 5xx status code
func (o *PutFlowsOutcomeMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this put flows outcome method not allowed response a status code equal to that given
func (o *PutFlowsOutcomeMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

func (o *PutFlowsOutcomeMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/outcomes/{flowOutcomeId}][%d] putFlowsOutcomeMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *PutFlowsOutcomeMethodNotAllowed) String() string {
	return fmt.Sprintf("[PUT /api/v2/flows/outcomes/{flowOutcomeId}][%d] putFlowsOutcomeMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *PutFlowsOutcomeMethodNotAllowed) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowsOutcomeMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowsOutcomeRequestTimeout creates a PutFlowsOutcomeRequestTimeout with default headers values
func NewPutFlowsOutcomeRequestTimeout() *PutFlowsOutcomeRequestTimeout {
	return &PutFlowsOutcomeRequestTimeout{}
}

/*
PutFlowsOutcomeRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutFlowsOutcomeRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put flows outcome request timeout response has a 2xx status code
func (o *PutFlowsOutcomeRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put flows outcome request timeout response has a 3xx status code
func (o *PutFlowsOutcomeRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put flows outcome request timeout response has a 4xx status code
func (o *PutFlowsOutcomeRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this put flows outcome request timeout response has a 5xx status code
func (o *PutFlowsOutcomeRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this put flows outcome request timeout response a status code equal to that given
func (o *PutFlowsOutcomeRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PutFlowsOutcomeRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/outcomes/{flowOutcomeId}][%d] putFlowsOutcomeRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutFlowsOutcomeRequestTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/flows/outcomes/{flowOutcomeId}][%d] putFlowsOutcomeRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutFlowsOutcomeRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowsOutcomeRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowsOutcomeConflict creates a PutFlowsOutcomeConflict with default headers values
func NewPutFlowsOutcomeConflict() *PutFlowsOutcomeConflict {
	return &PutFlowsOutcomeConflict{}
}

/*
PutFlowsOutcomeConflict describes a response with status code 409, with default header values.

Conflict
*/
type PutFlowsOutcomeConflict struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put flows outcome conflict response has a 2xx status code
func (o *PutFlowsOutcomeConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put flows outcome conflict response has a 3xx status code
func (o *PutFlowsOutcomeConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put flows outcome conflict response has a 4xx status code
func (o *PutFlowsOutcomeConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this put flows outcome conflict response has a 5xx status code
func (o *PutFlowsOutcomeConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this put flows outcome conflict response a status code equal to that given
func (o *PutFlowsOutcomeConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PutFlowsOutcomeConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/outcomes/{flowOutcomeId}][%d] putFlowsOutcomeConflict  %+v", 409, o.Payload)
}

func (o *PutFlowsOutcomeConflict) String() string {
	return fmt.Sprintf("[PUT /api/v2/flows/outcomes/{flowOutcomeId}][%d] putFlowsOutcomeConflict  %+v", 409, o.Payload)
}

func (o *PutFlowsOutcomeConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowsOutcomeConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowsOutcomeRequestEntityTooLarge creates a PutFlowsOutcomeRequestEntityTooLarge with default headers values
func NewPutFlowsOutcomeRequestEntityTooLarge() *PutFlowsOutcomeRequestEntityTooLarge {
	return &PutFlowsOutcomeRequestEntityTooLarge{}
}

/*
PutFlowsOutcomeRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PutFlowsOutcomeRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put flows outcome request entity too large response has a 2xx status code
func (o *PutFlowsOutcomeRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put flows outcome request entity too large response has a 3xx status code
func (o *PutFlowsOutcomeRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put flows outcome request entity too large response has a 4xx status code
func (o *PutFlowsOutcomeRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this put flows outcome request entity too large response has a 5xx status code
func (o *PutFlowsOutcomeRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this put flows outcome request entity too large response a status code equal to that given
func (o *PutFlowsOutcomeRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PutFlowsOutcomeRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/outcomes/{flowOutcomeId}][%d] putFlowsOutcomeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutFlowsOutcomeRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /api/v2/flows/outcomes/{flowOutcomeId}][%d] putFlowsOutcomeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutFlowsOutcomeRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowsOutcomeRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowsOutcomeUnsupportedMediaType creates a PutFlowsOutcomeUnsupportedMediaType with default headers values
func NewPutFlowsOutcomeUnsupportedMediaType() *PutFlowsOutcomeUnsupportedMediaType {
	return &PutFlowsOutcomeUnsupportedMediaType{}
}

/*
PutFlowsOutcomeUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutFlowsOutcomeUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put flows outcome unsupported media type response has a 2xx status code
func (o *PutFlowsOutcomeUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put flows outcome unsupported media type response has a 3xx status code
func (o *PutFlowsOutcomeUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put flows outcome unsupported media type response has a 4xx status code
func (o *PutFlowsOutcomeUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this put flows outcome unsupported media type response has a 5xx status code
func (o *PutFlowsOutcomeUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this put flows outcome unsupported media type response a status code equal to that given
func (o *PutFlowsOutcomeUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PutFlowsOutcomeUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/outcomes/{flowOutcomeId}][%d] putFlowsOutcomeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutFlowsOutcomeUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /api/v2/flows/outcomes/{flowOutcomeId}][%d] putFlowsOutcomeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutFlowsOutcomeUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowsOutcomeUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowsOutcomeTooManyRequests creates a PutFlowsOutcomeTooManyRequests with default headers values
func NewPutFlowsOutcomeTooManyRequests() *PutFlowsOutcomeTooManyRequests {
	return &PutFlowsOutcomeTooManyRequests{}
}

/*
PutFlowsOutcomeTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutFlowsOutcomeTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put flows outcome too many requests response has a 2xx status code
func (o *PutFlowsOutcomeTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put flows outcome too many requests response has a 3xx status code
func (o *PutFlowsOutcomeTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put flows outcome too many requests response has a 4xx status code
func (o *PutFlowsOutcomeTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this put flows outcome too many requests response has a 5xx status code
func (o *PutFlowsOutcomeTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this put flows outcome too many requests response a status code equal to that given
func (o *PutFlowsOutcomeTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PutFlowsOutcomeTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/outcomes/{flowOutcomeId}][%d] putFlowsOutcomeTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutFlowsOutcomeTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /api/v2/flows/outcomes/{flowOutcomeId}][%d] putFlowsOutcomeTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutFlowsOutcomeTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowsOutcomeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowsOutcomeInternalServerError creates a PutFlowsOutcomeInternalServerError with default headers values
func NewPutFlowsOutcomeInternalServerError() *PutFlowsOutcomeInternalServerError {
	return &PutFlowsOutcomeInternalServerError{}
}

/*
PutFlowsOutcomeInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutFlowsOutcomeInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put flows outcome internal server error response has a 2xx status code
func (o *PutFlowsOutcomeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put flows outcome internal server error response has a 3xx status code
func (o *PutFlowsOutcomeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put flows outcome internal server error response has a 4xx status code
func (o *PutFlowsOutcomeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put flows outcome internal server error response has a 5xx status code
func (o *PutFlowsOutcomeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put flows outcome internal server error response a status code equal to that given
func (o *PutFlowsOutcomeInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutFlowsOutcomeInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/outcomes/{flowOutcomeId}][%d] putFlowsOutcomeInternalServerError  %+v", 500, o.Payload)
}

func (o *PutFlowsOutcomeInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v2/flows/outcomes/{flowOutcomeId}][%d] putFlowsOutcomeInternalServerError  %+v", 500, o.Payload)
}

func (o *PutFlowsOutcomeInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowsOutcomeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowsOutcomeServiceUnavailable creates a PutFlowsOutcomeServiceUnavailable with default headers values
func NewPutFlowsOutcomeServiceUnavailable() *PutFlowsOutcomeServiceUnavailable {
	return &PutFlowsOutcomeServiceUnavailable{}
}

/*
PutFlowsOutcomeServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutFlowsOutcomeServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put flows outcome service unavailable response has a 2xx status code
func (o *PutFlowsOutcomeServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put flows outcome service unavailable response has a 3xx status code
func (o *PutFlowsOutcomeServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put flows outcome service unavailable response has a 4xx status code
func (o *PutFlowsOutcomeServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this put flows outcome service unavailable response has a 5xx status code
func (o *PutFlowsOutcomeServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this put flows outcome service unavailable response a status code equal to that given
func (o *PutFlowsOutcomeServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PutFlowsOutcomeServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/outcomes/{flowOutcomeId}][%d] putFlowsOutcomeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutFlowsOutcomeServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /api/v2/flows/outcomes/{flowOutcomeId}][%d] putFlowsOutcomeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutFlowsOutcomeServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowsOutcomeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowsOutcomeGatewayTimeout creates a PutFlowsOutcomeGatewayTimeout with default headers values
func NewPutFlowsOutcomeGatewayTimeout() *PutFlowsOutcomeGatewayTimeout {
	return &PutFlowsOutcomeGatewayTimeout{}
}

/*
PutFlowsOutcomeGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PutFlowsOutcomeGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put flows outcome gateway timeout response has a 2xx status code
func (o *PutFlowsOutcomeGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put flows outcome gateway timeout response has a 3xx status code
func (o *PutFlowsOutcomeGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put flows outcome gateway timeout response has a 4xx status code
func (o *PutFlowsOutcomeGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this put flows outcome gateway timeout response has a 5xx status code
func (o *PutFlowsOutcomeGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this put flows outcome gateway timeout response a status code equal to that given
func (o *PutFlowsOutcomeGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PutFlowsOutcomeGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/outcomes/{flowOutcomeId}][%d] putFlowsOutcomeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutFlowsOutcomeGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/flows/outcomes/{flowOutcomeId}][%d] putFlowsOutcomeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutFlowsOutcomeGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowsOutcomeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

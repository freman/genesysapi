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

// PutFlowReader is a Reader for the PutFlow structure.
type PutFlowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutFlowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutFlowOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutFlowBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutFlowUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutFlowForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutFlowNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPutFlowMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutFlowRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutFlowConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPutFlowGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutFlowRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutFlowUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutFlowTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutFlowInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutFlowServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutFlowGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutFlowOK creates a PutFlowOK with default headers values
func NewPutFlowOK() *PutFlowOK {
	return &PutFlowOK{}
}

/*
PutFlowOK describes a response with status code 200, with default header values.

successful operation
*/
type PutFlowOK struct {
	Payload *models.Flow
}

// IsSuccess returns true when this put flow o k response has a 2xx status code
func (o *PutFlowOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put flow o k response has a 3xx status code
func (o *PutFlowOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put flow o k response has a 4xx status code
func (o *PutFlowOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put flow o k response has a 5xx status code
func (o *PutFlowOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put flow o k response a status code equal to that given
func (o *PutFlowOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutFlowOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowOK  %+v", 200, o.Payload)
}

func (o *PutFlowOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowOK  %+v", 200, o.Payload)
}

func (o *PutFlowOK) GetPayload() *models.Flow {
	return o.Payload
}

func (o *PutFlowOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Flow)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowBadRequest creates a PutFlowBadRequest with default headers values
func NewPutFlowBadRequest() *PutFlowBadRequest {
	return &PutFlowBadRequest{}
}

/*
PutFlowBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutFlowBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put flow bad request response has a 2xx status code
func (o *PutFlowBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put flow bad request response has a 3xx status code
func (o *PutFlowBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put flow bad request response has a 4xx status code
func (o *PutFlowBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put flow bad request response has a 5xx status code
func (o *PutFlowBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put flow bad request response a status code equal to that given
func (o *PutFlowBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutFlowBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowBadRequest  %+v", 400, o.Payload)
}

func (o *PutFlowBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowBadRequest  %+v", 400, o.Payload)
}

func (o *PutFlowBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowUnauthorized creates a PutFlowUnauthorized with default headers values
func NewPutFlowUnauthorized() *PutFlowUnauthorized {
	return &PutFlowUnauthorized{}
}

/*
PutFlowUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PutFlowUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put flow unauthorized response has a 2xx status code
func (o *PutFlowUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put flow unauthorized response has a 3xx status code
func (o *PutFlowUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put flow unauthorized response has a 4xx status code
func (o *PutFlowUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put flow unauthorized response has a 5xx status code
func (o *PutFlowUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put flow unauthorized response a status code equal to that given
func (o *PutFlowUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PutFlowUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowUnauthorized  %+v", 401, o.Payload)
}

func (o *PutFlowUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowUnauthorized  %+v", 401, o.Payload)
}

func (o *PutFlowUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowForbidden creates a PutFlowForbidden with default headers values
func NewPutFlowForbidden() *PutFlowForbidden {
	return &PutFlowForbidden{}
}

/*
PutFlowForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PutFlowForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put flow forbidden response has a 2xx status code
func (o *PutFlowForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put flow forbidden response has a 3xx status code
func (o *PutFlowForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put flow forbidden response has a 4xx status code
func (o *PutFlowForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put flow forbidden response has a 5xx status code
func (o *PutFlowForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put flow forbidden response a status code equal to that given
func (o *PutFlowForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutFlowForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowForbidden  %+v", 403, o.Payload)
}

func (o *PutFlowForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowForbidden  %+v", 403, o.Payload)
}

func (o *PutFlowForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowNotFound creates a PutFlowNotFound with default headers values
func NewPutFlowNotFound() *PutFlowNotFound {
	return &PutFlowNotFound{}
}

/*
PutFlowNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PutFlowNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put flow not found response has a 2xx status code
func (o *PutFlowNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put flow not found response has a 3xx status code
func (o *PutFlowNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put flow not found response has a 4xx status code
func (o *PutFlowNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put flow not found response has a 5xx status code
func (o *PutFlowNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put flow not found response a status code equal to that given
func (o *PutFlowNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutFlowNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowNotFound  %+v", 404, o.Payload)
}

func (o *PutFlowNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowNotFound  %+v", 404, o.Payload)
}

func (o *PutFlowNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowMethodNotAllowed creates a PutFlowMethodNotAllowed with default headers values
func NewPutFlowMethodNotAllowed() *PutFlowMethodNotAllowed {
	return &PutFlowMethodNotAllowed{}
}

/*
PutFlowMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type PutFlowMethodNotAllowed struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put flow method not allowed response has a 2xx status code
func (o *PutFlowMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put flow method not allowed response has a 3xx status code
func (o *PutFlowMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put flow method not allowed response has a 4xx status code
func (o *PutFlowMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this put flow method not allowed response has a 5xx status code
func (o *PutFlowMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this put flow method not allowed response a status code equal to that given
func (o *PutFlowMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

func (o *PutFlowMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *PutFlowMethodNotAllowed) String() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *PutFlowMethodNotAllowed) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowRequestTimeout creates a PutFlowRequestTimeout with default headers values
func NewPutFlowRequestTimeout() *PutFlowRequestTimeout {
	return &PutFlowRequestTimeout{}
}

/*
PutFlowRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutFlowRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put flow request timeout response has a 2xx status code
func (o *PutFlowRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put flow request timeout response has a 3xx status code
func (o *PutFlowRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put flow request timeout response has a 4xx status code
func (o *PutFlowRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this put flow request timeout response has a 5xx status code
func (o *PutFlowRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this put flow request timeout response a status code equal to that given
func (o *PutFlowRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PutFlowRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutFlowRequestTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutFlowRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowConflict creates a PutFlowConflict with default headers values
func NewPutFlowConflict() *PutFlowConflict {
	return &PutFlowConflict{}
}

/*
PutFlowConflict describes a response with status code 409, with default header values.

Conflict
*/
type PutFlowConflict struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put flow conflict response has a 2xx status code
func (o *PutFlowConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put flow conflict response has a 3xx status code
func (o *PutFlowConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put flow conflict response has a 4xx status code
func (o *PutFlowConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this put flow conflict response has a 5xx status code
func (o *PutFlowConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this put flow conflict response a status code equal to that given
func (o *PutFlowConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PutFlowConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowConflict  %+v", 409, o.Payload)
}

func (o *PutFlowConflict) String() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowConflict  %+v", 409, o.Payload)
}

func (o *PutFlowConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowGone creates a PutFlowGone with default headers values
func NewPutFlowGone() *PutFlowGone {
	return &PutFlowGone{}
}

/*
PutFlowGone describes a response with status code 410, with default header values.

Gone
*/
type PutFlowGone struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put flow gone response has a 2xx status code
func (o *PutFlowGone) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put flow gone response has a 3xx status code
func (o *PutFlowGone) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put flow gone response has a 4xx status code
func (o *PutFlowGone) IsClientError() bool {
	return true
}

// IsServerError returns true when this put flow gone response has a 5xx status code
func (o *PutFlowGone) IsServerError() bool {
	return false
}

// IsCode returns true when this put flow gone response a status code equal to that given
func (o *PutFlowGone) IsCode(code int) bool {
	return code == 410
}

func (o *PutFlowGone) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowGone  %+v", 410, o.Payload)
}

func (o *PutFlowGone) String() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowGone  %+v", 410, o.Payload)
}

func (o *PutFlowGone) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowRequestEntityTooLarge creates a PutFlowRequestEntityTooLarge with default headers values
func NewPutFlowRequestEntityTooLarge() *PutFlowRequestEntityTooLarge {
	return &PutFlowRequestEntityTooLarge{}
}

/*
PutFlowRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PutFlowRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put flow request entity too large response has a 2xx status code
func (o *PutFlowRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put flow request entity too large response has a 3xx status code
func (o *PutFlowRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put flow request entity too large response has a 4xx status code
func (o *PutFlowRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this put flow request entity too large response has a 5xx status code
func (o *PutFlowRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this put flow request entity too large response a status code equal to that given
func (o *PutFlowRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PutFlowRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutFlowRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutFlowRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowUnsupportedMediaType creates a PutFlowUnsupportedMediaType with default headers values
func NewPutFlowUnsupportedMediaType() *PutFlowUnsupportedMediaType {
	return &PutFlowUnsupportedMediaType{}
}

/*
PutFlowUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutFlowUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put flow unsupported media type response has a 2xx status code
func (o *PutFlowUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put flow unsupported media type response has a 3xx status code
func (o *PutFlowUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put flow unsupported media type response has a 4xx status code
func (o *PutFlowUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this put flow unsupported media type response has a 5xx status code
func (o *PutFlowUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this put flow unsupported media type response a status code equal to that given
func (o *PutFlowUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PutFlowUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutFlowUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutFlowUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowTooManyRequests creates a PutFlowTooManyRequests with default headers values
func NewPutFlowTooManyRequests() *PutFlowTooManyRequests {
	return &PutFlowTooManyRequests{}
}

/*
PutFlowTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutFlowTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put flow too many requests response has a 2xx status code
func (o *PutFlowTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put flow too many requests response has a 3xx status code
func (o *PutFlowTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put flow too many requests response has a 4xx status code
func (o *PutFlowTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this put flow too many requests response has a 5xx status code
func (o *PutFlowTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this put flow too many requests response a status code equal to that given
func (o *PutFlowTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PutFlowTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutFlowTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutFlowTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowInternalServerError creates a PutFlowInternalServerError with default headers values
func NewPutFlowInternalServerError() *PutFlowInternalServerError {
	return &PutFlowInternalServerError{}
}

/*
PutFlowInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutFlowInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put flow internal server error response has a 2xx status code
func (o *PutFlowInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put flow internal server error response has a 3xx status code
func (o *PutFlowInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put flow internal server error response has a 4xx status code
func (o *PutFlowInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put flow internal server error response has a 5xx status code
func (o *PutFlowInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put flow internal server error response a status code equal to that given
func (o *PutFlowInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutFlowInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowInternalServerError  %+v", 500, o.Payload)
}

func (o *PutFlowInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowInternalServerError  %+v", 500, o.Payload)
}

func (o *PutFlowInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowServiceUnavailable creates a PutFlowServiceUnavailable with default headers values
func NewPutFlowServiceUnavailable() *PutFlowServiceUnavailable {
	return &PutFlowServiceUnavailable{}
}

/*
PutFlowServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutFlowServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put flow service unavailable response has a 2xx status code
func (o *PutFlowServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put flow service unavailable response has a 3xx status code
func (o *PutFlowServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put flow service unavailable response has a 4xx status code
func (o *PutFlowServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this put flow service unavailable response has a 5xx status code
func (o *PutFlowServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this put flow service unavailable response a status code equal to that given
func (o *PutFlowServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PutFlowServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutFlowServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutFlowServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowGatewayTimeout creates a PutFlowGatewayTimeout with default headers values
func NewPutFlowGatewayTimeout() *PutFlowGatewayTimeout {
	return &PutFlowGatewayTimeout{}
}

/*
PutFlowGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PutFlowGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put flow gateway timeout response has a 2xx status code
func (o *PutFlowGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put flow gateway timeout response has a 3xx status code
func (o *PutFlowGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put flow gateway timeout response has a 4xx status code
func (o *PutFlowGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this put flow gateway timeout response has a 5xx status code
func (o *PutFlowGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this put flow gateway timeout response a status code equal to that given
func (o *PutFlowGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PutFlowGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutFlowGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutFlowGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

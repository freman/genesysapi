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

// DeleteFlowsReader is a Reader for the DeleteFlows structure.
type DeleteFlowsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFlowsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteFlowsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteFlowsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteFlowsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteFlowsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteFlowsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteFlowsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteFlowsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteFlowsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteFlowsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteFlowsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteFlowsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteFlowsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteFlowsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteFlowsOK creates a DeleteFlowsOK with default headers values
func NewDeleteFlowsOK() *DeleteFlowsOK {
	return &DeleteFlowsOK{}
}

/*
DeleteFlowsOK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteFlowsOK struct {
	Payload *models.Operation
}

// IsSuccess returns true when this delete flows o k response has a 2xx status code
func (o *DeleteFlowsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete flows o k response has a 3xx status code
func (o *DeleteFlowsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete flows o k response has a 4xx status code
func (o *DeleteFlowsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete flows o k response has a 5xx status code
func (o *DeleteFlowsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete flows o k response a status code equal to that given
func (o *DeleteFlowsOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteFlowsOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows][%d] deleteFlowsOK  %+v", 200, o.Payload)
}

func (o *DeleteFlowsOK) String() string {
	return fmt.Sprintf("[DELETE /api/v2/flows][%d] deleteFlowsOK  %+v", 200, o.Payload)
}

func (o *DeleteFlowsOK) GetPayload() *models.Operation {
	return o.Payload
}

func (o *DeleteFlowsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Operation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsBadRequest creates a DeleteFlowsBadRequest with default headers values
func NewDeleteFlowsBadRequest() *DeleteFlowsBadRequest {
	return &DeleteFlowsBadRequest{}
}

/*
DeleteFlowsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteFlowsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete flows bad request response has a 2xx status code
func (o *DeleteFlowsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete flows bad request response has a 3xx status code
func (o *DeleteFlowsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete flows bad request response has a 4xx status code
func (o *DeleteFlowsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete flows bad request response has a 5xx status code
func (o *DeleteFlowsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete flows bad request response a status code equal to that given
func (o *DeleteFlowsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteFlowsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows][%d] deleteFlowsBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteFlowsBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/flows][%d] deleteFlowsBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteFlowsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsUnauthorized creates a DeleteFlowsUnauthorized with default headers values
func NewDeleteFlowsUnauthorized() *DeleteFlowsUnauthorized {
	return &DeleteFlowsUnauthorized{}
}

/*
DeleteFlowsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteFlowsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete flows unauthorized response has a 2xx status code
func (o *DeleteFlowsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete flows unauthorized response has a 3xx status code
func (o *DeleteFlowsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete flows unauthorized response has a 4xx status code
func (o *DeleteFlowsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete flows unauthorized response has a 5xx status code
func (o *DeleteFlowsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete flows unauthorized response a status code equal to that given
func (o *DeleteFlowsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteFlowsUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows][%d] deleteFlowsUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteFlowsUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/flows][%d] deleteFlowsUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteFlowsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsForbidden creates a DeleteFlowsForbidden with default headers values
func NewDeleteFlowsForbidden() *DeleteFlowsForbidden {
	return &DeleteFlowsForbidden{}
}

/*
DeleteFlowsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteFlowsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete flows forbidden response has a 2xx status code
func (o *DeleteFlowsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete flows forbidden response has a 3xx status code
func (o *DeleteFlowsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete flows forbidden response has a 4xx status code
func (o *DeleteFlowsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete flows forbidden response has a 5xx status code
func (o *DeleteFlowsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete flows forbidden response a status code equal to that given
func (o *DeleteFlowsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteFlowsForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows][%d] deleteFlowsForbidden  %+v", 403, o.Payload)
}

func (o *DeleteFlowsForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/flows][%d] deleteFlowsForbidden  %+v", 403, o.Payload)
}

func (o *DeleteFlowsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsNotFound creates a DeleteFlowsNotFound with default headers values
func NewDeleteFlowsNotFound() *DeleteFlowsNotFound {
	return &DeleteFlowsNotFound{}
}

/*
DeleteFlowsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteFlowsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete flows not found response has a 2xx status code
func (o *DeleteFlowsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete flows not found response has a 3xx status code
func (o *DeleteFlowsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete flows not found response has a 4xx status code
func (o *DeleteFlowsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete flows not found response has a 5xx status code
func (o *DeleteFlowsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete flows not found response a status code equal to that given
func (o *DeleteFlowsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteFlowsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows][%d] deleteFlowsNotFound  %+v", 404, o.Payload)
}

func (o *DeleteFlowsNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/flows][%d] deleteFlowsNotFound  %+v", 404, o.Payload)
}

func (o *DeleteFlowsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsRequestTimeout creates a DeleteFlowsRequestTimeout with default headers values
func NewDeleteFlowsRequestTimeout() *DeleteFlowsRequestTimeout {
	return &DeleteFlowsRequestTimeout{}
}

/*
DeleteFlowsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteFlowsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete flows request timeout response has a 2xx status code
func (o *DeleteFlowsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete flows request timeout response has a 3xx status code
func (o *DeleteFlowsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete flows request timeout response has a 4xx status code
func (o *DeleteFlowsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete flows request timeout response has a 5xx status code
func (o *DeleteFlowsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete flows request timeout response a status code equal to that given
func (o *DeleteFlowsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteFlowsRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows][%d] deleteFlowsRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteFlowsRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/flows][%d] deleteFlowsRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteFlowsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsConflict creates a DeleteFlowsConflict with default headers values
func NewDeleteFlowsConflict() *DeleteFlowsConflict {
	return &DeleteFlowsConflict{}
}

/*
DeleteFlowsConflict describes a response with status code 409, with default header values.

Conflict
*/
type DeleteFlowsConflict struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete flows conflict response has a 2xx status code
func (o *DeleteFlowsConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete flows conflict response has a 3xx status code
func (o *DeleteFlowsConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete flows conflict response has a 4xx status code
func (o *DeleteFlowsConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete flows conflict response has a 5xx status code
func (o *DeleteFlowsConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this delete flows conflict response a status code equal to that given
func (o *DeleteFlowsConflict) IsCode(code int) bool {
	return code == 409
}

func (o *DeleteFlowsConflict) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows][%d] deleteFlowsConflict  %+v", 409, o.Payload)
}

func (o *DeleteFlowsConflict) String() string {
	return fmt.Sprintf("[DELETE /api/v2/flows][%d] deleteFlowsConflict  %+v", 409, o.Payload)
}

func (o *DeleteFlowsConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsRequestEntityTooLarge creates a DeleteFlowsRequestEntityTooLarge with default headers values
func NewDeleteFlowsRequestEntityTooLarge() *DeleteFlowsRequestEntityTooLarge {
	return &DeleteFlowsRequestEntityTooLarge{}
}

/*
DeleteFlowsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type DeleteFlowsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete flows request entity too large response has a 2xx status code
func (o *DeleteFlowsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete flows request entity too large response has a 3xx status code
func (o *DeleteFlowsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete flows request entity too large response has a 4xx status code
func (o *DeleteFlowsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete flows request entity too large response has a 5xx status code
func (o *DeleteFlowsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete flows request entity too large response a status code equal to that given
func (o *DeleteFlowsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteFlowsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows][%d] deleteFlowsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteFlowsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/flows][%d] deleteFlowsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteFlowsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsUnsupportedMediaType creates a DeleteFlowsUnsupportedMediaType with default headers values
func NewDeleteFlowsUnsupportedMediaType() *DeleteFlowsUnsupportedMediaType {
	return &DeleteFlowsUnsupportedMediaType{}
}

/*
DeleteFlowsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteFlowsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete flows unsupported media type response has a 2xx status code
func (o *DeleteFlowsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete flows unsupported media type response has a 3xx status code
func (o *DeleteFlowsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete flows unsupported media type response has a 4xx status code
func (o *DeleteFlowsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete flows unsupported media type response has a 5xx status code
func (o *DeleteFlowsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete flows unsupported media type response a status code equal to that given
func (o *DeleteFlowsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteFlowsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows][%d] deleteFlowsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteFlowsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/flows][%d] deleteFlowsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteFlowsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsTooManyRequests creates a DeleteFlowsTooManyRequests with default headers values
func NewDeleteFlowsTooManyRequests() *DeleteFlowsTooManyRequests {
	return &DeleteFlowsTooManyRequests{}
}

/*
DeleteFlowsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteFlowsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete flows too many requests response has a 2xx status code
func (o *DeleteFlowsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete flows too many requests response has a 3xx status code
func (o *DeleteFlowsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete flows too many requests response has a 4xx status code
func (o *DeleteFlowsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete flows too many requests response has a 5xx status code
func (o *DeleteFlowsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete flows too many requests response a status code equal to that given
func (o *DeleteFlowsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteFlowsTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows][%d] deleteFlowsTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteFlowsTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/flows][%d] deleteFlowsTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteFlowsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsInternalServerError creates a DeleteFlowsInternalServerError with default headers values
func NewDeleteFlowsInternalServerError() *DeleteFlowsInternalServerError {
	return &DeleteFlowsInternalServerError{}
}

/*
DeleteFlowsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteFlowsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete flows internal server error response has a 2xx status code
func (o *DeleteFlowsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete flows internal server error response has a 3xx status code
func (o *DeleteFlowsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete flows internal server error response has a 4xx status code
func (o *DeleteFlowsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete flows internal server error response has a 5xx status code
func (o *DeleteFlowsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete flows internal server error response a status code equal to that given
func (o *DeleteFlowsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteFlowsInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows][%d] deleteFlowsInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteFlowsInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/flows][%d] deleteFlowsInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteFlowsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsServiceUnavailable creates a DeleteFlowsServiceUnavailable with default headers values
func NewDeleteFlowsServiceUnavailable() *DeleteFlowsServiceUnavailable {
	return &DeleteFlowsServiceUnavailable{}
}

/*
DeleteFlowsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteFlowsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete flows service unavailable response has a 2xx status code
func (o *DeleteFlowsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete flows service unavailable response has a 3xx status code
func (o *DeleteFlowsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete flows service unavailable response has a 4xx status code
func (o *DeleteFlowsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete flows service unavailable response has a 5xx status code
func (o *DeleteFlowsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete flows service unavailable response a status code equal to that given
func (o *DeleteFlowsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteFlowsServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows][%d] deleteFlowsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteFlowsServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/flows][%d] deleteFlowsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteFlowsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsGatewayTimeout creates a DeleteFlowsGatewayTimeout with default headers values
func NewDeleteFlowsGatewayTimeout() *DeleteFlowsGatewayTimeout {
	return &DeleteFlowsGatewayTimeout{}
}

/*
DeleteFlowsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteFlowsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete flows gateway timeout response has a 2xx status code
func (o *DeleteFlowsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete flows gateway timeout response has a 3xx status code
func (o *DeleteFlowsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete flows gateway timeout response has a 4xx status code
func (o *DeleteFlowsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete flows gateway timeout response has a 5xx status code
func (o *DeleteFlowsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete flows gateway timeout response a status code equal to that given
func (o *DeleteFlowsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteFlowsGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows][%d] deleteFlowsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteFlowsGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/flows][%d] deleteFlowsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteFlowsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

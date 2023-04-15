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

// DeleteFlowsMilestoneReader is a Reader for the DeleteFlowsMilestone structure.
type DeleteFlowsMilestoneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFlowsMilestoneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteFlowsMilestoneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteFlowsMilestoneNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteFlowsMilestoneBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteFlowsMilestoneUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteFlowsMilestoneForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteFlowsMilestoneNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewDeleteFlowsMilestoneMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteFlowsMilestoneRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteFlowsMilestoneConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteFlowsMilestoneRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteFlowsMilestoneUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteFlowsMilestoneTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteFlowsMilestoneInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteFlowsMilestoneServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteFlowsMilestoneGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteFlowsMilestoneOK creates a DeleteFlowsMilestoneOK with default headers values
func NewDeleteFlowsMilestoneOK() *DeleteFlowsMilestoneOK {
	return &DeleteFlowsMilestoneOK{}
}

/*
DeleteFlowsMilestoneOK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteFlowsMilestoneOK struct {
	Payload models.Empty
}

// IsSuccess returns true when this delete flows milestone o k response has a 2xx status code
func (o *DeleteFlowsMilestoneOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete flows milestone o k response has a 3xx status code
func (o *DeleteFlowsMilestoneOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete flows milestone o k response has a 4xx status code
func (o *DeleteFlowsMilestoneOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete flows milestone o k response has a 5xx status code
func (o *DeleteFlowsMilestoneOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete flows milestone o k response a status code equal to that given
func (o *DeleteFlowsMilestoneOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteFlowsMilestoneOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneOK  %+v", 200, o.Payload)
}

func (o *DeleteFlowsMilestoneOK) String() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneOK  %+v", 200, o.Payload)
}

func (o *DeleteFlowsMilestoneOK) GetPayload() models.Empty {
	return o.Payload
}

func (o *DeleteFlowsMilestoneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsMilestoneNoContent creates a DeleteFlowsMilestoneNoContent with default headers values
func NewDeleteFlowsMilestoneNoContent() *DeleteFlowsMilestoneNoContent {
	return &DeleteFlowsMilestoneNoContent{}
}

/*
DeleteFlowsMilestoneNoContent describes a response with status code 204, with default header values.

Delete was successful.
*/
type DeleteFlowsMilestoneNoContent struct {
}

// IsSuccess returns true when this delete flows milestone no content response has a 2xx status code
func (o *DeleteFlowsMilestoneNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete flows milestone no content response has a 3xx status code
func (o *DeleteFlowsMilestoneNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete flows milestone no content response has a 4xx status code
func (o *DeleteFlowsMilestoneNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete flows milestone no content response has a 5xx status code
func (o *DeleteFlowsMilestoneNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete flows milestone no content response a status code equal to that given
func (o *DeleteFlowsMilestoneNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteFlowsMilestoneNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneNoContent ", 204)
}

func (o *DeleteFlowsMilestoneNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneNoContent ", 204)
}

func (o *DeleteFlowsMilestoneNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFlowsMilestoneBadRequest creates a DeleteFlowsMilestoneBadRequest with default headers values
func NewDeleteFlowsMilestoneBadRequest() *DeleteFlowsMilestoneBadRequest {
	return &DeleteFlowsMilestoneBadRequest{}
}

/*
DeleteFlowsMilestoneBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteFlowsMilestoneBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete flows milestone bad request response has a 2xx status code
func (o *DeleteFlowsMilestoneBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete flows milestone bad request response has a 3xx status code
func (o *DeleteFlowsMilestoneBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete flows milestone bad request response has a 4xx status code
func (o *DeleteFlowsMilestoneBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete flows milestone bad request response has a 5xx status code
func (o *DeleteFlowsMilestoneBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete flows milestone bad request response a status code equal to that given
func (o *DeleteFlowsMilestoneBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteFlowsMilestoneBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteFlowsMilestoneBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteFlowsMilestoneBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsMilestoneBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsMilestoneUnauthorized creates a DeleteFlowsMilestoneUnauthorized with default headers values
func NewDeleteFlowsMilestoneUnauthorized() *DeleteFlowsMilestoneUnauthorized {
	return &DeleteFlowsMilestoneUnauthorized{}
}

/*
DeleteFlowsMilestoneUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteFlowsMilestoneUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete flows milestone unauthorized response has a 2xx status code
func (o *DeleteFlowsMilestoneUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete flows milestone unauthorized response has a 3xx status code
func (o *DeleteFlowsMilestoneUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete flows milestone unauthorized response has a 4xx status code
func (o *DeleteFlowsMilestoneUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete flows milestone unauthorized response has a 5xx status code
func (o *DeleteFlowsMilestoneUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete flows milestone unauthorized response a status code equal to that given
func (o *DeleteFlowsMilestoneUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteFlowsMilestoneUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteFlowsMilestoneUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteFlowsMilestoneUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsMilestoneUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsMilestoneForbidden creates a DeleteFlowsMilestoneForbidden with default headers values
func NewDeleteFlowsMilestoneForbidden() *DeleteFlowsMilestoneForbidden {
	return &DeleteFlowsMilestoneForbidden{}
}

/*
DeleteFlowsMilestoneForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteFlowsMilestoneForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete flows milestone forbidden response has a 2xx status code
func (o *DeleteFlowsMilestoneForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete flows milestone forbidden response has a 3xx status code
func (o *DeleteFlowsMilestoneForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete flows milestone forbidden response has a 4xx status code
func (o *DeleteFlowsMilestoneForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete flows milestone forbidden response has a 5xx status code
func (o *DeleteFlowsMilestoneForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete flows milestone forbidden response a status code equal to that given
func (o *DeleteFlowsMilestoneForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteFlowsMilestoneForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneForbidden  %+v", 403, o.Payload)
}

func (o *DeleteFlowsMilestoneForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneForbidden  %+v", 403, o.Payload)
}

func (o *DeleteFlowsMilestoneForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsMilestoneForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsMilestoneNotFound creates a DeleteFlowsMilestoneNotFound with default headers values
func NewDeleteFlowsMilestoneNotFound() *DeleteFlowsMilestoneNotFound {
	return &DeleteFlowsMilestoneNotFound{}
}

/*
DeleteFlowsMilestoneNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteFlowsMilestoneNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete flows milestone not found response has a 2xx status code
func (o *DeleteFlowsMilestoneNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete flows milestone not found response has a 3xx status code
func (o *DeleteFlowsMilestoneNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete flows milestone not found response has a 4xx status code
func (o *DeleteFlowsMilestoneNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete flows milestone not found response has a 5xx status code
func (o *DeleteFlowsMilestoneNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete flows milestone not found response a status code equal to that given
func (o *DeleteFlowsMilestoneNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteFlowsMilestoneNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneNotFound  %+v", 404, o.Payload)
}

func (o *DeleteFlowsMilestoneNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneNotFound  %+v", 404, o.Payload)
}

func (o *DeleteFlowsMilestoneNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsMilestoneNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsMilestoneMethodNotAllowed creates a DeleteFlowsMilestoneMethodNotAllowed with default headers values
func NewDeleteFlowsMilestoneMethodNotAllowed() *DeleteFlowsMilestoneMethodNotAllowed {
	return &DeleteFlowsMilestoneMethodNotAllowed{}
}

/*
DeleteFlowsMilestoneMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type DeleteFlowsMilestoneMethodNotAllowed struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete flows milestone method not allowed response has a 2xx status code
func (o *DeleteFlowsMilestoneMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete flows milestone method not allowed response has a 3xx status code
func (o *DeleteFlowsMilestoneMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete flows milestone method not allowed response has a 4xx status code
func (o *DeleteFlowsMilestoneMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete flows milestone method not allowed response has a 5xx status code
func (o *DeleteFlowsMilestoneMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this delete flows milestone method not allowed response a status code equal to that given
func (o *DeleteFlowsMilestoneMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

func (o *DeleteFlowsMilestoneMethodNotAllowed) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *DeleteFlowsMilestoneMethodNotAllowed) String() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *DeleteFlowsMilestoneMethodNotAllowed) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsMilestoneMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsMilestoneRequestTimeout creates a DeleteFlowsMilestoneRequestTimeout with default headers values
func NewDeleteFlowsMilestoneRequestTimeout() *DeleteFlowsMilestoneRequestTimeout {
	return &DeleteFlowsMilestoneRequestTimeout{}
}

/*
DeleteFlowsMilestoneRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteFlowsMilestoneRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete flows milestone request timeout response has a 2xx status code
func (o *DeleteFlowsMilestoneRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete flows milestone request timeout response has a 3xx status code
func (o *DeleteFlowsMilestoneRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete flows milestone request timeout response has a 4xx status code
func (o *DeleteFlowsMilestoneRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete flows milestone request timeout response has a 5xx status code
func (o *DeleteFlowsMilestoneRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete flows milestone request timeout response a status code equal to that given
func (o *DeleteFlowsMilestoneRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteFlowsMilestoneRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteFlowsMilestoneRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteFlowsMilestoneRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsMilestoneRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsMilestoneConflict creates a DeleteFlowsMilestoneConflict with default headers values
func NewDeleteFlowsMilestoneConflict() *DeleteFlowsMilestoneConflict {
	return &DeleteFlowsMilestoneConflict{}
}

/*
DeleteFlowsMilestoneConflict describes a response with status code 409, with default header values.

Conflict
*/
type DeleteFlowsMilestoneConflict struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete flows milestone conflict response has a 2xx status code
func (o *DeleteFlowsMilestoneConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete flows milestone conflict response has a 3xx status code
func (o *DeleteFlowsMilestoneConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete flows milestone conflict response has a 4xx status code
func (o *DeleteFlowsMilestoneConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete flows milestone conflict response has a 5xx status code
func (o *DeleteFlowsMilestoneConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this delete flows milestone conflict response a status code equal to that given
func (o *DeleteFlowsMilestoneConflict) IsCode(code int) bool {
	return code == 409
}

func (o *DeleteFlowsMilestoneConflict) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneConflict  %+v", 409, o.Payload)
}

func (o *DeleteFlowsMilestoneConflict) String() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneConflict  %+v", 409, o.Payload)
}

func (o *DeleteFlowsMilestoneConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsMilestoneConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsMilestoneRequestEntityTooLarge creates a DeleteFlowsMilestoneRequestEntityTooLarge with default headers values
func NewDeleteFlowsMilestoneRequestEntityTooLarge() *DeleteFlowsMilestoneRequestEntityTooLarge {
	return &DeleteFlowsMilestoneRequestEntityTooLarge{}
}

/*
DeleteFlowsMilestoneRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type DeleteFlowsMilestoneRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete flows milestone request entity too large response has a 2xx status code
func (o *DeleteFlowsMilestoneRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete flows milestone request entity too large response has a 3xx status code
func (o *DeleteFlowsMilestoneRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete flows milestone request entity too large response has a 4xx status code
func (o *DeleteFlowsMilestoneRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete flows milestone request entity too large response has a 5xx status code
func (o *DeleteFlowsMilestoneRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete flows milestone request entity too large response a status code equal to that given
func (o *DeleteFlowsMilestoneRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteFlowsMilestoneRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteFlowsMilestoneRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteFlowsMilestoneRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsMilestoneRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsMilestoneUnsupportedMediaType creates a DeleteFlowsMilestoneUnsupportedMediaType with default headers values
func NewDeleteFlowsMilestoneUnsupportedMediaType() *DeleteFlowsMilestoneUnsupportedMediaType {
	return &DeleteFlowsMilestoneUnsupportedMediaType{}
}

/*
DeleteFlowsMilestoneUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteFlowsMilestoneUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete flows milestone unsupported media type response has a 2xx status code
func (o *DeleteFlowsMilestoneUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete flows milestone unsupported media type response has a 3xx status code
func (o *DeleteFlowsMilestoneUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete flows milestone unsupported media type response has a 4xx status code
func (o *DeleteFlowsMilestoneUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete flows milestone unsupported media type response has a 5xx status code
func (o *DeleteFlowsMilestoneUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete flows milestone unsupported media type response a status code equal to that given
func (o *DeleteFlowsMilestoneUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteFlowsMilestoneUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteFlowsMilestoneUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteFlowsMilestoneUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsMilestoneUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsMilestoneTooManyRequests creates a DeleteFlowsMilestoneTooManyRequests with default headers values
func NewDeleteFlowsMilestoneTooManyRequests() *DeleteFlowsMilestoneTooManyRequests {
	return &DeleteFlowsMilestoneTooManyRequests{}
}

/*
DeleteFlowsMilestoneTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteFlowsMilestoneTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete flows milestone too many requests response has a 2xx status code
func (o *DeleteFlowsMilestoneTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete flows milestone too many requests response has a 3xx status code
func (o *DeleteFlowsMilestoneTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete flows milestone too many requests response has a 4xx status code
func (o *DeleteFlowsMilestoneTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete flows milestone too many requests response has a 5xx status code
func (o *DeleteFlowsMilestoneTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete flows milestone too many requests response a status code equal to that given
func (o *DeleteFlowsMilestoneTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteFlowsMilestoneTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteFlowsMilestoneTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteFlowsMilestoneTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsMilestoneTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsMilestoneInternalServerError creates a DeleteFlowsMilestoneInternalServerError with default headers values
func NewDeleteFlowsMilestoneInternalServerError() *DeleteFlowsMilestoneInternalServerError {
	return &DeleteFlowsMilestoneInternalServerError{}
}

/*
DeleteFlowsMilestoneInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteFlowsMilestoneInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete flows milestone internal server error response has a 2xx status code
func (o *DeleteFlowsMilestoneInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete flows milestone internal server error response has a 3xx status code
func (o *DeleteFlowsMilestoneInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete flows milestone internal server error response has a 4xx status code
func (o *DeleteFlowsMilestoneInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete flows milestone internal server error response has a 5xx status code
func (o *DeleteFlowsMilestoneInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete flows milestone internal server error response a status code equal to that given
func (o *DeleteFlowsMilestoneInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteFlowsMilestoneInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteFlowsMilestoneInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteFlowsMilestoneInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsMilestoneInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsMilestoneServiceUnavailable creates a DeleteFlowsMilestoneServiceUnavailable with default headers values
func NewDeleteFlowsMilestoneServiceUnavailable() *DeleteFlowsMilestoneServiceUnavailable {
	return &DeleteFlowsMilestoneServiceUnavailable{}
}

/*
DeleteFlowsMilestoneServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteFlowsMilestoneServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete flows milestone service unavailable response has a 2xx status code
func (o *DeleteFlowsMilestoneServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete flows milestone service unavailable response has a 3xx status code
func (o *DeleteFlowsMilestoneServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete flows milestone service unavailable response has a 4xx status code
func (o *DeleteFlowsMilestoneServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete flows milestone service unavailable response has a 5xx status code
func (o *DeleteFlowsMilestoneServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete flows milestone service unavailable response a status code equal to that given
func (o *DeleteFlowsMilestoneServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteFlowsMilestoneServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteFlowsMilestoneServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteFlowsMilestoneServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsMilestoneServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsMilestoneGatewayTimeout creates a DeleteFlowsMilestoneGatewayTimeout with default headers values
func NewDeleteFlowsMilestoneGatewayTimeout() *DeleteFlowsMilestoneGatewayTimeout {
	return &DeleteFlowsMilestoneGatewayTimeout{}
}

/*
DeleteFlowsMilestoneGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteFlowsMilestoneGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete flows milestone gateway timeout response has a 2xx status code
func (o *DeleteFlowsMilestoneGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete flows milestone gateway timeout response has a 3xx status code
func (o *DeleteFlowsMilestoneGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete flows milestone gateway timeout response has a 4xx status code
func (o *DeleteFlowsMilestoneGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete flows milestone gateway timeout response has a 5xx status code
func (o *DeleteFlowsMilestoneGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete flows milestone gateway timeout response a status code equal to that given
func (o *DeleteFlowsMilestoneGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteFlowsMilestoneGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteFlowsMilestoneGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteFlowsMilestoneGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsMilestoneGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

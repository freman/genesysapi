// Code generated by go-swagger; DO NOT EDIT.

package workforce_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeleteWorkforcemanagementManagementunitReader is a Reader for the DeleteWorkforcemanagementManagementunit structure.
type DeleteWorkforcemanagementManagementunitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteWorkforcemanagementManagementunitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteWorkforcemanagementManagementunitNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteWorkforcemanagementManagementunitBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteWorkforcemanagementManagementunitUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteWorkforcemanagementManagementunitForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteWorkforcemanagementManagementunitNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteWorkforcemanagementManagementunitRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteWorkforcemanagementManagementunitRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteWorkforcemanagementManagementunitUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteWorkforcemanagementManagementunitTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteWorkforcemanagementManagementunitInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteWorkforcemanagementManagementunitServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteWorkforcemanagementManagementunitGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteWorkforcemanagementManagementunitNoContent creates a DeleteWorkforcemanagementManagementunitNoContent with default headers values
func NewDeleteWorkforcemanagementManagementunitNoContent() *DeleteWorkforcemanagementManagementunitNoContent {
	return &DeleteWorkforcemanagementManagementunitNoContent{}
}

/*
DeleteWorkforcemanagementManagementunitNoContent describes a response with status code 204, with default header values.

The management unit was successfully deleted
*/
type DeleteWorkforcemanagementManagementunitNoContent struct {
}

// IsSuccess returns true when this delete workforcemanagement managementunit no content response has a 2xx status code
func (o *DeleteWorkforcemanagementManagementunitNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete workforcemanagement managementunit no content response has a 3xx status code
func (o *DeleteWorkforcemanagementManagementunitNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete workforcemanagement managementunit no content response has a 4xx status code
func (o *DeleteWorkforcemanagementManagementunitNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete workforcemanagement managementunit no content response has a 5xx status code
func (o *DeleteWorkforcemanagementManagementunitNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete workforcemanagement managementunit no content response a status code equal to that given
func (o *DeleteWorkforcemanagementManagementunitNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteWorkforcemanagementManagementunitNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] deleteWorkforcemanagementManagementunitNoContent ", 204)
}

func (o *DeleteWorkforcemanagementManagementunitNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] deleteWorkforcemanagementManagementunitNoContent ", 204)
}

func (o *DeleteWorkforcemanagementManagementunitNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteWorkforcemanagementManagementunitBadRequest creates a DeleteWorkforcemanagementManagementunitBadRequest with default headers values
func NewDeleteWorkforcemanagementManagementunitBadRequest() *DeleteWorkforcemanagementManagementunitBadRequest {
	return &DeleteWorkforcemanagementManagementunitBadRequest{}
}

/*
DeleteWorkforcemanagementManagementunitBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteWorkforcemanagementManagementunitBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete workforcemanagement managementunit bad request response has a 2xx status code
func (o *DeleteWorkforcemanagementManagementunitBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete workforcemanagement managementunit bad request response has a 3xx status code
func (o *DeleteWorkforcemanagementManagementunitBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete workforcemanagement managementunit bad request response has a 4xx status code
func (o *DeleteWorkforcemanagementManagementunitBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete workforcemanagement managementunit bad request response has a 5xx status code
func (o *DeleteWorkforcemanagementManagementunitBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete workforcemanagement managementunit bad request response a status code equal to that given
func (o *DeleteWorkforcemanagementManagementunitBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteWorkforcemanagementManagementunitBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] deleteWorkforcemanagementManagementunitBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] deleteWorkforcemanagementManagementunitBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementManagementunitBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementManagementunitUnauthorized creates a DeleteWorkforcemanagementManagementunitUnauthorized with default headers values
func NewDeleteWorkforcemanagementManagementunitUnauthorized() *DeleteWorkforcemanagementManagementunitUnauthorized {
	return &DeleteWorkforcemanagementManagementunitUnauthorized{}
}

/*
DeleteWorkforcemanagementManagementunitUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteWorkforcemanagementManagementunitUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete workforcemanagement managementunit unauthorized response has a 2xx status code
func (o *DeleteWorkforcemanagementManagementunitUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete workforcemanagement managementunit unauthorized response has a 3xx status code
func (o *DeleteWorkforcemanagementManagementunitUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete workforcemanagement managementunit unauthorized response has a 4xx status code
func (o *DeleteWorkforcemanagementManagementunitUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete workforcemanagement managementunit unauthorized response has a 5xx status code
func (o *DeleteWorkforcemanagementManagementunitUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete workforcemanagement managementunit unauthorized response a status code equal to that given
func (o *DeleteWorkforcemanagementManagementunitUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteWorkforcemanagementManagementunitUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] deleteWorkforcemanagementManagementunitUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] deleteWorkforcemanagementManagementunitUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementManagementunitUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementManagementunitForbidden creates a DeleteWorkforcemanagementManagementunitForbidden with default headers values
func NewDeleteWorkforcemanagementManagementunitForbidden() *DeleteWorkforcemanagementManagementunitForbidden {
	return &DeleteWorkforcemanagementManagementunitForbidden{}
}

/*
DeleteWorkforcemanagementManagementunitForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteWorkforcemanagementManagementunitForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete workforcemanagement managementunit forbidden response has a 2xx status code
func (o *DeleteWorkforcemanagementManagementunitForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete workforcemanagement managementunit forbidden response has a 3xx status code
func (o *DeleteWorkforcemanagementManagementunitForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete workforcemanagement managementunit forbidden response has a 4xx status code
func (o *DeleteWorkforcemanagementManagementunitForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete workforcemanagement managementunit forbidden response has a 5xx status code
func (o *DeleteWorkforcemanagementManagementunitForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete workforcemanagement managementunit forbidden response a status code equal to that given
func (o *DeleteWorkforcemanagementManagementunitForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteWorkforcemanagementManagementunitForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] deleteWorkforcemanagementManagementunitForbidden  %+v", 403, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] deleteWorkforcemanagementManagementunitForbidden  %+v", 403, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementManagementunitForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementManagementunitNotFound creates a DeleteWorkforcemanagementManagementunitNotFound with default headers values
func NewDeleteWorkforcemanagementManagementunitNotFound() *DeleteWorkforcemanagementManagementunitNotFound {
	return &DeleteWorkforcemanagementManagementunitNotFound{}
}

/*
DeleteWorkforcemanagementManagementunitNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteWorkforcemanagementManagementunitNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete workforcemanagement managementunit not found response has a 2xx status code
func (o *DeleteWorkforcemanagementManagementunitNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete workforcemanagement managementunit not found response has a 3xx status code
func (o *DeleteWorkforcemanagementManagementunitNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete workforcemanagement managementunit not found response has a 4xx status code
func (o *DeleteWorkforcemanagementManagementunitNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete workforcemanagement managementunit not found response has a 5xx status code
func (o *DeleteWorkforcemanagementManagementunitNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete workforcemanagement managementunit not found response a status code equal to that given
func (o *DeleteWorkforcemanagementManagementunitNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteWorkforcemanagementManagementunitNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] deleteWorkforcemanagementManagementunitNotFound  %+v", 404, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] deleteWorkforcemanagementManagementunitNotFound  %+v", 404, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementManagementunitNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementManagementunitRequestTimeout creates a DeleteWorkforcemanagementManagementunitRequestTimeout with default headers values
func NewDeleteWorkforcemanagementManagementunitRequestTimeout() *DeleteWorkforcemanagementManagementunitRequestTimeout {
	return &DeleteWorkforcemanagementManagementunitRequestTimeout{}
}

/*
DeleteWorkforcemanagementManagementunitRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteWorkforcemanagementManagementunitRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete workforcemanagement managementunit request timeout response has a 2xx status code
func (o *DeleteWorkforcemanagementManagementunitRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete workforcemanagement managementunit request timeout response has a 3xx status code
func (o *DeleteWorkforcemanagementManagementunitRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete workforcemanagement managementunit request timeout response has a 4xx status code
func (o *DeleteWorkforcemanagementManagementunitRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete workforcemanagement managementunit request timeout response has a 5xx status code
func (o *DeleteWorkforcemanagementManagementunitRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete workforcemanagement managementunit request timeout response a status code equal to that given
func (o *DeleteWorkforcemanagementManagementunitRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteWorkforcemanagementManagementunitRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] deleteWorkforcemanagementManagementunitRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] deleteWorkforcemanagementManagementunitRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementManagementunitRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementManagementunitRequestEntityTooLarge creates a DeleteWorkforcemanagementManagementunitRequestEntityTooLarge with default headers values
func NewDeleteWorkforcemanagementManagementunitRequestEntityTooLarge() *DeleteWorkforcemanagementManagementunitRequestEntityTooLarge {
	return &DeleteWorkforcemanagementManagementunitRequestEntityTooLarge{}
}

/*
DeleteWorkforcemanagementManagementunitRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type DeleteWorkforcemanagementManagementunitRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete workforcemanagement managementunit request entity too large response has a 2xx status code
func (o *DeleteWorkforcemanagementManagementunitRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete workforcemanagement managementunit request entity too large response has a 3xx status code
func (o *DeleteWorkforcemanagementManagementunitRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete workforcemanagement managementunit request entity too large response has a 4xx status code
func (o *DeleteWorkforcemanagementManagementunitRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete workforcemanagement managementunit request entity too large response has a 5xx status code
func (o *DeleteWorkforcemanagementManagementunitRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete workforcemanagement managementunit request entity too large response a status code equal to that given
func (o *DeleteWorkforcemanagementManagementunitRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteWorkforcemanagementManagementunitRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] deleteWorkforcemanagementManagementunitRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] deleteWorkforcemanagementManagementunitRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementManagementunitRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementManagementunitUnsupportedMediaType creates a DeleteWorkforcemanagementManagementunitUnsupportedMediaType with default headers values
func NewDeleteWorkforcemanagementManagementunitUnsupportedMediaType() *DeleteWorkforcemanagementManagementunitUnsupportedMediaType {
	return &DeleteWorkforcemanagementManagementunitUnsupportedMediaType{}
}

/*
DeleteWorkforcemanagementManagementunitUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteWorkforcemanagementManagementunitUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete workforcemanagement managementunit unsupported media type response has a 2xx status code
func (o *DeleteWorkforcemanagementManagementunitUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete workforcemanagement managementunit unsupported media type response has a 3xx status code
func (o *DeleteWorkforcemanagementManagementunitUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete workforcemanagement managementunit unsupported media type response has a 4xx status code
func (o *DeleteWorkforcemanagementManagementunitUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete workforcemanagement managementunit unsupported media type response has a 5xx status code
func (o *DeleteWorkforcemanagementManagementunitUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete workforcemanagement managementunit unsupported media type response a status code equal to that given
func (o *DeleteWorkforcemanagementManagementunitUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteWorkforcemanagementManagementunitUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] deleteWorkforcemanagementManagementunitUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] deleteWorkforcemanagementManagementunitUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementManagementunitUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementManagementunitTooManyRequests creates a DeleteWorkforcemanagementManagementunitTooManyRequests with default headers values
func NewDeleteWorkforcemanagementManagementunitTooManyRequests() *DeleteWorkforcemanagementManagementunitTooManyRequests {
	return &DeleteWorkforcemanagementManagementunitTooManyRequests{}
}

/*
DeleteWorkforcemanagementManagementunitTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteWorkforcemanagementManagementunitTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete workforcemanagement managementunit too many requests response has a 2xx status code
func (o *DeleteWorkforcemanagementManagementunitTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete workforcemanagement managementunit too many requests response has a 3xx status code
func (o *DeleteWorkforcemanagementManagementunitTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete workforcemanagement managementunit too many requests response has a 4xx status code
func (o *DeleteWorkforcemanagementManagementunitTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete workforcemanagement managementunit too many requests response has a 5xx status code
func (o *DeleteWorkforcemanagementManagementunitTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete workforcemanagement managementunit too many requests response a status code equal to that given
func (o *DeleteWorkforcemanagementManagementunitTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteWorkforcemanagementManagementunitTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] deleteWorkforcemanagementManagementunitTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] deleteWorkforcemanagementManagementunitTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementManagementunitTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementManagementunitInternalServerError creates a DeleteWorkforcemanagementManagementunitInternalServerError with default headers values
func NewDeleteWorkforcemanagementManagementunitInternalServerError() *DeleteWorkforcemanagementManagementunitInternalServerError {
	return &DeleteWorkforcemanagementManagementunitInternalServerError{}
}

/*
DeleteWorkforcemanagementManagementunitInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteWorkforcemanagementManagementunitInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete workforcemanagement managementunit internal server error response has a 2xx status code
func (o *DeleteWorkforcemanagementManagementunitInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete workforcemanagement managementunit internal server error response has a 3xx status code
func (o *DeleteWorkforcemanagementManagementunitInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete workforcemanagement managementunit internal server error response has a 4xx status code
func (o *DeleteWorkforcemanagementManagementunitInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete workforcemanagement managementunit internal server error response has a 5xx status code
func (o *DeleteWorkforcemanagementManagementunitInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete workforcemanagement managementunit internal server error response a status code equal to that given
func (o *DeleteWorkforcemanagementManagementunitInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteWorkforcemanagementManagementunitInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] deleteWorkforcemanagementManagementunitInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] deleteWorkforcemanagementManagementunitInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementManagementunitInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementManagementunitServiceUnavailable creates a DeleteWorkforcemanagementManagementunitServiceUnavailable with default headers values
func NewDeleteWorkforcemanagementManagementunitServiceUnavailable() *DeleteWorkforcemanagementManagementunitServiceUnavailable {
	return &DeleteWorkforcemanagementManagementunitServiceUnavailable{}
}

/*
DeleteWorkforcemanagementManagementunitServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteWorkforcemanagementManagementunitServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete workforcemanagement managementunit service unavailable response has a 2xx status code
func (o *DeleteWorkforcemanagementManagementunitServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete workforcemanagement managementunit service unavailable response has a 3xx status code
func (o *DeleteWorkforcemanagementManagementunitServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete workforcemanagement managementunit service unavailable response has a 4xx status code
func (o *DeleteWorkforcemanagementManagementunitServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete workforcemanagement managementunit service unavailable response has a 5xx status code
func (o *DeleteWorkforcemanagementManagementunitServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete workforcemanagement managementunit service unavailable response a status code equal to that given
func (o *DeleteWorkforcemanagementManagementunitServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteWorkforcemanagementManagementunitServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] deleteWorkforcemanagementManagementunitServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] deleteWorkforcemanagementManagementunitServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementManagementunitServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementManagementunitGatewayTimeout creates a DeleteWorkforcemanagementManagementunitGatewayTimeout with default headers values
func NewDeleteWorkforcemanagementManagementunitGatewayTimeout() *DeleteWorkforcemanagementManagementunitGatewayTimeout {
	return &DeleteWorkforcemanagementManagementunitGatewayTimeout{}
}

/*
DeleteWorkforcemanagementManagementunitGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteWorkforcemanagementManagementunitGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete workforcemanagement managementunit gateway timeout response has a 2xx status code
func (o *DeleteWorkforcemanagementManagementunitGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete workforcemanagement managementunit gateway timeout response has a 3xx status code
func (o *DeleteWorkforcemanagementManagementunitGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete workforcemanagement managementunit gateway timeout response has a 4xx status code
func (o *DeleteWorkforcemanagementManagementunitGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete workforcemanagement managementunit gateway timeout response has a 5xx status code
func (o *DeleteWorkforcemanagementManagementunitGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete workforcemanagement managementunit gateway timeout response a status code equal to that given
func (o *DeleteWorkforcemanagementManagementunitGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteWorkforcemanagementManagementunitGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] deleteWorkforcemanagementManagementunitGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] deleteWorkforcemanagementManagementunitGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementManagementunitGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

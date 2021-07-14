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

// DeleteWorkforcemanagementBusinessunitActivitycodeReader is a Reader for the DeleteWorkforcemanagementBusinessunitActivitycode structure.
type DeleteWorkforcemanagementBusinessunitActivitycodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteWorkforcemanagementBusinessunitActivitycodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteWorkforcemanagementBusinessunitActivitycodeNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteWorkforcemanagementBusinessunitActivitycodeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteWorkforcemanagementBusinessunitActivitycodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteWorkforcemanagementBusinessunitActivitycodeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteWorkforcemanagementBusinessunitActivitycodeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteWorkforcemanagementBusinessunitActivitycodeRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteWorkforcemanagementBusinessunitActivitycodeRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteWorkforcemanagementBusinessunitActivitycodeUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteWorkforcemanagementBusinessunitActivitycodeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteWorkforcemanagementBusinessunitActivitycodeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteWorkforcemanagementBusinessunitActivitycodeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteWorkforcemanagementBusinessunitActivitycodeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteWorkforcemanagementBusinessunitActivitycodeNoContent creates a DeleteWorkforcemanagementBusinessunitActivitycodeNoContent with default headers values
func NewDeleteWorkforcemanagementBusinessunitActivitycodeNoContent() *DeleteWorkforcemanagementBusinessunitActivitycodeNoContent {
	return &DeleteWorkforcemanagementBusinessunitActivitycodeNoContent{}
}

/*DeleteWorkforcemanagementBusinessunitActivitycodeNoContent handles this case with default header values.

The activity code was deleted successfully
*/
type DeleteWorkforcemanagementBusinessunitActivitycodeNoContent struct {
}

func (o *DeleteWorkforcemanagementBusinessunitActivitycodeNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}][%d] deleteWorkforcemanagementBusinessunitActivitycodeNoContent ", 204)
}

func (o *DeleteWorkforcemanagementBusinessunitActivitycodeNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteWorkforcemanagementBusinessunitActivitycodeBadRequest creates a DeleteWorkforcemanagementBusinessunitActivitycodeBadRequest with default headers values
func NewDeleteWorkforcemanagementBusinessunitActivitycodeBadRequest() *DeleteWorkforcemanagementBusinessunitActivitycodeBadRequest {
	return &DeleteWorkforcemanagementBusinessunitActivitycodeBadRequest{}
}

/*DeleteWorkforcemanagementBusinessunitActivitycodeBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteWorkforcemanagementBusinessunitActivitycodeBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementBusinessunitActivitycodeBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}][%d] deleteWorkforcemanagementBusinessunitActivitycodeBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteWorkforcemanagementBusinessunitActivitycodeBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementBusinessunitActivitycodeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementBusinessunitActivitycodeUnauthorized creates a DeleteWorkforcemanagementBusinessunitActivitycodeUnauthorized with default headers values
func NewDeleteWorkforcemanagementBusinessunitActivitycodeUnauthorized() *DeleteWorkforcemanagementBusinessunitActivitycodeUnauthorized {
	return &DeleteWorkforcemanagementBusinessunitActivitycodeUnauthorized{}
}

/*DeleteWorkforcemanagementBusinessunitActivitycodeUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteWorkforcemanagementBusinessunitActivitycodeUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementBusinessunitActivitycodeUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}][%d] deleteWorkforcemanagementBusinessunitActivitycodeUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteWorkforcemanagementBusinessunitActivitycodeUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementBusinessunitActivitycodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementBusinessunitActivitycodeForbidden creates a DeleteWorkforcemanagementBusinessunitActivitycodeForbidden with default headers values
func NewDeleteWorkforcemanagementBusinessunitActivitycodeForbidden() *DeleteWorkforcemanagementBusinessunitActivitycodeForbidden {
	return &DeleteWorkforcemanagementBusinessunitActivitycodeForbidden{}
}

/*DeleteWorkforcemanagementBusinessunitActivitycodeForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteWorkforcemanagementBusinessunitActivitycodeForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementBusinessunitActivitycodeForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}][%d] deleteWorkforcemanagementBusinessunitActivitycodeForbidden  %+v", 403, o.Payload)
}

func (o *DeleteWorkforcemanagementBusinessunitActivitycodeForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementBusinessunitActivitycodeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementBusinessunitActivitycodeNotFound creates a DeleteWorkforcemanagementBusinessunitActivitycodeNotFound with default headers values
func NewDeleteWorkforcemanagementBusinessunitActivitycodeNotFound() *DeleteWorkforcemanagementBusinessunitActivitycodeNotFound {
	return &DeleteWorkforcemanagementBusinessunitActivitycodeNotFound{}
}

/*DeleteWorkforcemanagementBusinessunitActivitycodeNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteWorkforcemanagementBusinessunitActivitycodeNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementBusinessunitActivitycodeNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}][%d] deleteWorkforcemanagementBusinessunitActivitycodeNotFound  %+v", 404, o.Payload)
}

func (o *DeleteWorkforcemanagementBusinessunitActivitycodeNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementBusinessunitActivitycodeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementBusinessunitActivitycodeRequestTimeout creates a DeleteWorkforcemanagementBusinessunitActivitycodeRequestTimeout with default headers values
func NewDeleteWorkforcemanagementBusinessunitActivitycodeRequestTimeout() *DeleteWorkforcemanagementBusinessunitActivitycodeRequestTimeout {
	return &DeleteWorkforcemanagementBusinessunitActivitycodeRequestTimeout{}
}

/*DeleteWorkforcemanagementBusinessunitActivitycodeRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteWorkforcemanagementBusinessunitActivitycodeRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementBusinessunitActivitycodeRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}][%d] deleteWorkforcemanagementBusinessunitActivitycodeRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteWorkforcemanagementBusinessunitActivitycodeRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementBusinessunitActivitycodeRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementBusinessunitActivitycodeRequestEntityTooLarge creates a DeleteWorkforcemanagementBusinessunitActivitycodeRequestEntityTooLarge with default headers values
func NewDeleteWorkforcemanagementBusinessunitActivitycodeRequestEntityTooLarge() *DeleteWorkforcemanagementBusinessunitActivitycodeRequestEntityTooLarge {
	return &DeleteWorkforcemanagementBusinessunitActivitycodeRequestEntityTooLarge{}
}

/*DeleteWorkforcemanagementBusinessunitActivitycodeRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteWorkforcemanagementBusinessunitActivitycodeRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementBusinessunitActivitycodeRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}][%d] deleteWorkforcemanagementBusinessunitActivitycodeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteWorkforcemanagementBusinessunitActivitycodeRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementBusinessunitActivitycodeRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementBusinessunitActivitycodeUnsupportedMediaType creates a DeleteWorkforcemanagementBusinessunitActivitycodeUnsupportedMediaType with default headers values
func NewDeleteWorkforcemanagementBusinessunitActivitycodeUnsupportedMediaType() *DeleteWorkforcemanagementBusinessunitActivitycodeUnsupportedMediaType {
	return &DeleteWorkforcemanagementBusinessunitActivitycodeUnsupportedMediaType{}
}

/*DeleteWorkforcemanagementBusinessunitActivitycodeUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteWorkforcemanagementBusinessunitActivitycodeUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementBusinessunitActivitycodeUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}][%d] deleteWorkforcemanagementBusinessunitActivitycodeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteWorkforcemanagementBusinessunitActivitycodeUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementBusinessunitActivitycodeUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementBusinessunitActivitycodeTooManyRequests creates a DeleteWorkforcemanagementBusinessunitActivitycodeTooManyRequests with default headers values
func NewDeleteWorkforcemanagementBusinessunitActivitycodeTooManyRequests() *DeleteWorkforcemanagementBusinessunitActivitycodeTooManyRequests {
	return &DeleteWorkforcemanagementBusinessunitActivitycodeTooManyRequests{}
}

/*DeleteWorkforcemanagementBusinessunitActivitycodeTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteWorkforcemanagementBusinessunitActivitycodeTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementBusinessunitActivitycodeTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}][%d] deleteWorkforcemanagementBusinessunitActivitycodeTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteWorkforcemanagementBusinessunitActivitycodeTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementBusinessunitActivitycodeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementBusinessunitActivitycodeInternalServerError creates a DeleteWorkforcemanagementBusinessunitActivitycodeInternalServerError with default headers values
func NewDeleteWorkforcemanagementBusinessunitActivitycodeInternalServerError() *DeleteWorkforcemanagementBusinessunitActivitycodeInternalServerError {
	return &DeleteWorkforcemanagementBusinessunitActivitycodeInternalServerError{}
}

/*DeleteWorkforcemanagementBusinessunitActivitycodeInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteWorkforcemanagementBusinessunitActivitycodeInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementBusinessunitActivitycodeInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}][%d] deleteWorkforcemanagementBusinessunitActivitycodeInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteWorkforcemanagementBusinessunitActivitycodeInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementBusinessunitActivitycodeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementBusinessunitActivitycodeServiceUnavailable creates a DeleteWorkforcemanagementBusinessunitActivitycodeServiceUnavailable with default headers values
func NewDeleteWorkforcemanagementBusinessunitActivitycodeServiceUnavailable() *DeleteWorkforcemanagementBusinessunitActivitycodeServiceUnavailable {
	return &DeleteWorkforcemanagementBusinessunitActivitycodeServiceUnavailable{}
}

/*DeleteWorkforcemanagementBusinessunitActivitycodeServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteWorkforcemanagementBusinessunitActivitycodeServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementBusinessunitActivitycodeServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}][%d] deleteWorkforcemanagementBusinessunitActivitycodeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteWorkforcemanagementBusinessunitActivitycodeServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementBusinessunitActivitycodeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementBusinessunitActivitycodeGatewayTimeout creates a DeleteWorkforcemanagementBusinessunitActivitycodeGatewayTimeout with default headers values
func NewDeleteWorkforcemanagementBusinessunitActivitycodeGatewayTimeout() *DeleteWorkforcemanagementBusinessunitActivitycodeGatewayTimeout {
	return &DeleteWorkforcemanagementBusinessunitActivitycodeGatewayTimeout{}
}

/*DeleteWorkforcemanagementBusinessunitActivitycodeGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteWorkforcemanagementBusinessunitActivitycodeGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementBusinessunitActivitycodeGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}][%d] deleteWorkforcemanagementBusinessunitActivitycodeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteWorkforcemanagementBusinessunitActivitycodeGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementBusinessunitActivitycodeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

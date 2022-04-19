// Code generated by go-swagger; DO NOT EDIT.

package response_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PutResponsemanagementResponseReader is a Reader for the PutResponsemanagementResponse structure.
type PutResponsemanagementResponseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutResponsemanagementResponseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutResponsemanagementResponseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutResponsemanagementResponseBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutResponsemanagementResponseUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutResponsemanagementResponseForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutResponsemanagementResponseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutResponsemanagementResponseRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutResponsemanagementResponseConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutResponsemanagementResponseRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutResponsemanagementResponseUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutResponsemanagementResponseTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutResponsemanagementResponseInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutResponsemanagementResponseServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutResponsemanagementResponseGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutResponsemanagementResponseOK creates a PutResponsemanagementResponseOK with default headers values
func NewPutResponsemanagementResponseOK() *PutResponsemanagementResponseOK {
	return &PutResponsemanagementResponseOK{}
}

/*PutResponsemanagementResponseOK handles this case with default header values.

successful operation
*/
type PutResponsemanagementResponseOK struct {
	Payload *models.Response
}

func (o *PutResponsemanagementResponseOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/responsemanagement/responses/{responseId}][%d] putResponsemanagementResponseOK  %+v", 200, o.Payload)
}

func (o *PutResponsemanagementResponseOK) GetPayload() *models.Response {
	return o.Payload
}

func (o *PutResponsemanagementResponseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutResponsemanagementResponseBadRequest creates a PutResponsemanagementResponseBadRequest with default headers values
func NewPutResponsemanagementResponseBadRequest() *PutResponsemanagementResponseBadRequest {
	return &PutResponsemanagementResponseBadRequest{}
}

/*PutResponsemanagementResponseBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutResponsemanagementResponseBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutResponsemanagementResponseBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/responsemanagement/responses/{responseId}][%d] putResponsemanagementResponseBadRequest  %+v", 400, o.Payload)
}

func (o *PutResponsemanagementResponseBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutResponsemanagementResponseBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutResponsemanagementResponseUnauthorized creates a PutResponsemanagementResponseUnauthorized with default headers values
func NewPutResponsemanagementResponseUnauthorized() *PutResponsemanagementResponseUnauthorized {
	return &PutResponsemanagementResponseUnauthorized{}
}

/*PutResponsemanagementResponseUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutResponsemanagementResponseUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutResponsemanagementResponseUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/responsemanagement/responses/{responseId}][%d] putResponsemanagementResponseUnauthorized  %+v", 401, o.Payload)
}

func (o *PutResponsemanagementResponseUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutResponsemanagementResponseUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutResponsemanagementResponseForbidden creates a PutResponsemanagementResponseForbidden with default headers values
func NewPutResponsemanagementResponseForbidden() *PutResponsemanagementResponseForbidden {
	return &PutResponsemanagementResponseForbidden{}
}

/*PutResponsemanagementResponseForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutResponsemanagementResponseForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutResponsemanagementResponseForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/responsemanagement/responses/{responseId}][%d] putResponsemanagementResponseForbidden  %+v", 403, o.Payload)
}

func (o *PutResponsemanagementResponseForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutResponsemanagementResponseForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutResponsemanagementResponseNotFound creates a PutResponsemanagementResponseNotFound with default headers values
func NewPutResponsemanagementResponseNotFound() *PutResponsemanagementResponseNotFound {
	return &PutResponsemanagementResponseNotFound{}
}

/*PutResponsemanagementResponseNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutResponsemanagementResponseNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutResponsemanagementResponseNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/responsemanagement/responses/{responseId}][%d] putResponsemanagementResponseNotFound  %+v", 404, o.Payload)
}

func (o *PutResponsemanagementResponseNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutResponsemanagementResponseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutResponsemanagementResponseRequestTimeout creates a PutResponsemanagementResponseRequestTimeout with default headers values
func NewPutResponsemanagementResponseRequestTimeout() *PutResponsemanagementResponseRequestTimeout {
	return &PutResponsemanagementResponseRequestTimeout{}
}

/*PutResponsemanagementResponseRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutResponsemanagementResponseRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutResponsemanagementResponseRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/responsemanagement/responses/{responseId}][%d] putResponsemanagementResponseRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutResponsemanagementResponseRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutResponsemanagementResponseRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutResponsemanagementResponseConflict creates a PutResponsemanagementResponseConflict with default headers values
func NewPutResponsemanagementResponseConflict() *PutResponsemanagementResponseConflict {
	return &PutResponsemanagementResponseConflict{}
}

/*PutResponsemanagementResponseConflict handles this case with default header values.

Resource conflict - Unexpected version was provided
*/
type PutResponsemanagementResponseConflict struct {
}

func (o *PutResponsemanagementResponseConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v2/responsemanagement/responses/{responseId}][%d] putResponsemanagementResponseConflict ", 409)
}

func (o *PutResponsemanagementResponseConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutResponsemanagementResponseRequestEntityTooLarge creates a PutResponsemanagementResponseRequestEntityTooLarge with default headers values
func NewPutResponsemanagementResponseRequestEntityTooLarge() *PutResponsemanagementResponseRequestEntityTooLarge {
	return &PutResponsemanagementResponseRequestEntityTooLarge{}
}

/*PutResponsemanagementResponseRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PutResponsemanagementResponseRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutResponsemanagementResponseRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/responsemanagement/responses/{responseId}][%d] putResponsemanagementResponseRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutResponsemanagementResponseRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutResponsemanagementResponseRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutResponsemanagementResponseUnsupportedMediaType creates a PutResponsemanagementResponseUnsupportedMediaType with default headers values
func NewPutResponsemanagementResponseUnsupportedMediaType() *PutResponsemanagementResponseUnsupportedMediaType {
	return &PutResponsemanagementResponseUnsupportedMediaType{}
}

/*PutResponsemanagementResponseUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutResponsemanagementResponseUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutResponsemanagementResponseUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/responsemanagement/responses/{responseId}][%d] putResponsemanagementResponseUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutResponsemanagementResponseUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutResponsemanagementResponseUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutResponsemanagementResponseTooManyRequests creates a PutResponsemanagementResponseTooManyRequests with default headers values
func NewPutResponsemanagementResponseTooManyRequests() *PutResponsemanagementResponseTooManyRequests {
	return &PutResponsemanagementResponseTooManyRequests{}
}

/*PutResponsemanagementResponseTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutResponsemanagementResponseTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutResponsemanagementResponseTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/responsemanagement/responses/{responseId}][%d] putResponsemanagementResponseTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutResponsemanagementResponseTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutResponsemanagementResponseTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutResponsemanagementResponseInternalServerError creates a PutResponsemanagementResponseInternalServerError with default headers values
func NewPutResponsemanagementResponseInternalServerError() *PutResponsemanagementResponseInternalServerError {
	return &PutResponsemanagementResponseInternalServerError{}
}

/*PutResponsemanagementResponseInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutResponsemanagementResponseInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutResponsemanagementResponseInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/responsemanagement/responses/{responseId}][%d] putResponsemanagementResponseInternalServerError  %+v", 500, o.Payload)
}

func (o *PutResponsemanagementResponseInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutResponsemanagementResponseInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutResponsemanagementResponseServiceUnavailable creates a PutResponsemanagementResponseServiceUnavailable with default headers values
func NewPutResponsemanagementResponseServiceUnavailable() *PutResponsemanagementResponseServiceUnavailable {
	return &PutResponsemanagementResponseServiceUnavailable{}
}

/*PutResponsemanagementResponseServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutResponsemanagementResponseServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutResponsemanagementResponseServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/responsemanagement/responses/{responseId}][%d] putResponsemanagementResponseServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutResponsemanagementResponseServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutResponsemanagementResponseServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutResponsemanagementResponseGatewayTimeout creates a PutResponsemanagementResponseGatewayTimeout with default headers values
func NewPutResponsemanagementResponseGatewayTimeout() *PutResponsemanagementResponseGatewayTimeout {
	return &PutResponsemanagementResponseGatewayTimeout{}
}

/*PutResponsemanagementResponseGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutResponsemanagementResponseGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutResponsemanagementResponseGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/responsemanagement/responses/{responseId}][%d] putResponsemanagementResponseGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutResponsemanagementResponseGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutResponsemanagementResponseGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

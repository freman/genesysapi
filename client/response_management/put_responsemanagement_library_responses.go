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

// PutResponsemanagementLibraryReader is a Reader for the PutResponsemanagementLibrary structure.
type PutResponsemanagementLibraryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutResponsemanagementLibraryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutResponsemanagementLibraryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutResponsemanagementLibraryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutResponsemanagementLibraryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutResponsemanagementLibraryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutResponsemanagementLibraryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutResponsemanagementLibraryRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutResponsemanagementLibraryConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutResponsemanagementLibraryRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutResponsemanagementLibraryUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutResponsemanagementLibraryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutResponsemanagementLibraryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutResponsemanagementLibraryServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutResponsemanagementLibraryGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutResponsemanagementLibraryOK creates a PutResponsemanagementLibraryOK with default headers values
func NewPutResponsemanagementLibraryOK() *PutResponsemanagementLibraryOK {
	return &PutResponsemanagementLibraryOK{}
}

/*PutResponsemanagementLibraryOK handles this case with default header values.

successful operation
*/
type PutResponsemanagementLibraryOK struct {
	Payload *models.Library
}

func (o *PutResponsemanagementLibraryOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/responsemanagement/libraries/{libraryId}][%d] putResponsemanagementLibraryOK  %+v", 200, o.Payload)
}

func (o *PutResponsemanagementLibraryOK) GetPayload() *models.Library {
	return o.Payload
}

func (o *PutResponsemanagementLibraryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Library)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutResponsemanagementLibraryBadRequest creates a PutResponsemanagementLibraryBadRequest with default headers values
func NewPutResponsemanagementLibraryBadRequest() *PutResponsemanagementLibraryBadRequest {
	return &PutResponsemanagementLibraryBadRequest{}
}

/*PutResponsemanagementLibraryBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutResponsemanagementLibraryBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutResponsemanagementLibraryBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/responsemanagement/libraries/{libraryId}][%d] putResponsemanagementLibraryBadRequest  %+v", 400, o.Payload)
}

func (o *PutResponsemanagementLibraryBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutResponsemanagementLibraryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutResponsemanagementLibraryUnauthorized creates a PutResponsemanagementLibraryUnauthorized with default headers values
func NewPutResponsemanagementLibraryUnauthorized() *PutResponsemanagementLibraryUnauthorized {
	return &PutResponsemanagementLibraryUnauthorized{}
}

/*PutResponsemanagementLibraryUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutResponsemanagementLibraryUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutResponsemanagementLibraryUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/responsemanagement/libraries/{libraryId}][%d] putResponsemanagementLibraryUnauthorized  %+v", 401, o.Payload)
}

func (o *PutResponsemanagementLibraryUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutResponsemanagementLibraryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutResponsemanagementLibraryForbidden creates a PutResponsemanagementLibraryForbidden with default headers values
func NewPutResponsemanagementLibraryForbidden() *PutResponsemanagementLibraryForbidden {
	return &PutResponsemanagementLibraryForbidden{}
}

/*PutResponsemanagementLibraryForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutResponsemanagementLibraryForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutResponsemanagementLibraryForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/responsemanagement/libraries/{libraryId}][%d] putResponsemanagementLibraryForbidden  %+v", 403, o.Payload)
}

func (o *PutResponsemanagementLibraryForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutResponsemanagementLibraryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutResponsemanagementLibraryNotFound creates a PutResponsemanagementLibraryNotFound with default headers values
func NewPutResponsemanagementLibraryNotFound() *PutResponsemanagementLibraryNotFound {
	return &PutResponsemanagementLibraryNotFound{}
}

/*PutResponsemanagementLibraryNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutResponsemanagementLibraryNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutResponsemanagementLibraryNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/responsemanagement/libraries/{libraryId}][%d] putResponsemanagementLibraryNotFound  %+v", 404, o.Payload)
}

func (o *PutResponsemanagementLibraryNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutResponsemanagementLibraryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutResponsemanagementLibraryRequestTimeout creates a PutResponsemanagementLibraryRequestTimeout with default headers values
func NewPutResponsemanagementLibraryRequestTimeout() *PutResponsemanagementLibraryRequestTimeout {
	return &PutResponsemanagementLibraryRequestTimeout{}
}

/*PutResponsemanagementLibraryRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutResponsemanagementLibraryRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutResponsemanagementLibraryRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/responsemanagement/libraries/{libraryId}][%d] putResponsemanagementLibraryRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutResponsemanagementLibraryRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutResponsemanagementLibraryRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutResponsemanagementLibraryConflict creates a PutResponsemanagementLibraryConflict with default headers values
func NewPutResponsemanagementLibraryConflict() *PutResponsemanagementLibraryConflict {
	return &PutResponsemanagementLibraryConflict{}
}

/*PutResponsemanagementLibraryConflict handles this case with default header values.

Resource conflict - Unexpected version was provided
*/
type PutResponsemanagementLibraryConflict struct {
}

func (o *PutResponsemanagementLibraryConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v2/responsemanagement/libraries/{libraryId}][%d] putResponsemanagementLibraryConflict ", 409)
}

func (o *PutResponsemanagementLibraryConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutResponsemanagementLibraryRequestEntityTooLarge creates a PutResponsemanagementLibraryRequestEntityTooLarge with default headers values
func NewPutResponsemanagementLibraryRequestEntityTooLarge() *PutResponsemanagementLibraryRequestEntityTooLarge {
	return &PutResponsemanagementLibraryRequestEntityTooLarge{}
}

/*PutResponsemanagementLibraryRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PutResponsemanagementLibraryRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutResponsemanagementLibraryRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/responsemanagement/libraries/{libraryId}][%d] putResponsemanagementLibraryRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutResponsemanagementLibraryRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutResponsemanagementLibraryRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutResponsemanagementLibraryUnsupportedMediaType creates a PutResponsemanagementLibraryUnsupportedMediaType with default headers values
func NewPutResponsemanagementLibraryUnsupportedMediaType() *PutResponsemanagementLibraryUnsupportedMediaType {
	return &PutResponsemanagementLibraryUnsupportedMediaType{}
}

/*PutResponsemanagementLibraryUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutResponsemanagementLibraryUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutResponsemanagementLibraryUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/responsemanagement/libraries/{libraryId}][%d] putResponsemanagementLibraryUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutResponsemanagementLibraryUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutResponsemanagementLibraryUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutResponsemanagementLibraryTooManyRequests creates a PutResponsemanagementLibraryTooManyRequests with default headers values
func NewPutResponsemanagementLibraryTooManyRequests() *PutResponsemanagementLibraryTooManyRequests {
	return &PutResponsemanagementLibraryTooManyRequests{}
}

/*PutResponsemanagementLibraryTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutResponsemanagementLibraryTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutResponsemanagementLibraryTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/responsemanagement/libraries/{libraryId}][%d] putResponsemanagementLibraryTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutResponsemanagementLibraryTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutResponsemanagementLibraryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutResponsemanagementLibraryInternalServerError creates a PutResponsemanagementLibraryInternalServerError with default headers values
func NewPutResponsemanagementLibraryInternalServerError() *PutResponsemanagementLibraryInternalServerError {
	return &PutResponsemanagementLibraryInternalServerError{}
}

/*PutResponsemanagementLibraryInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutResponsemanagementLibraryInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutResponsemanagementLibraryInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/responsemanagement/libraries/{libraryId}][%d] putResponsemanagementLibraryInternalServerError  %+v", 500, o.Payload)
}

func (o *PutResponsemanagementLibraryInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutResponsemanagementLibraryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutResponsemanagementLibraryServiceUnavailable creates a PutResponsemanagementLibraryServiceUnavailable with default headers values
func NewPutResponsemanagementLibraryServiceUnavailable() *PutResponsemanagementLibraryServiceUnavailable {
	return &PutResponsemanagementLibraryServiceUnavailable{}
}

/*PutResponsemanagementLibraryServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutResponsemanagementLibraryServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutResponsemanagementLibraryServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/responsemanagement/libraries/{libraryId}][%d] putResponsemanagementLibraryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutResponsemanagementLibraryServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutResponsemanagementLibraryServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutResponsemanagementLibraryGatewayTimeout creates a PutResponsemanagementLibraryGatewayTimeout with default headers values
func NewPutResponsemanagementLibraryGatewayTimeout() *PutResponsemanagementLibraryGatewayTimeout {
	return &PutResponsemanagementLibraryGatewayTimeout{}
}

/*PutResponsemanagementLibraryGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutResponsemanagementLibraryGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutResponsemanagementLibraryGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/responsemanagement/libraries/{libraryId}][%d] putResponsemanagementLibraryGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutResponsemanagementLibraryGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutResponsemanagementLibraryGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

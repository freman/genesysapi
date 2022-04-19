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

// PostFlowVersionsReader is a Reader for the PostFlowVersions structure.
type PostFlowVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostFlowVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostFlowVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostFlowVersionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostFlowVersionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostFlowVersionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostFlowVersionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPostFlowVersionsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostFlowVersionsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostFlowVersionsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPostFlowVersionsGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostFlowVersionsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostFlowVersionsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostFlowVersionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostFlowVersionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostFlowVersionsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostFlowVersionsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostFlowVersionsOK creates a PostFlowVersionsOK with default headers values
func NewPostFlowVersionsOK() *PostFlowVersionsOK {
	return &PostFlowVersionsOK{}
}

/*PostFlowVersionsOK handles this case with default header values.

successful operation
*/
type PostFlowVersionsOK struct {
	Payload *models.FlowVersion
}

func (o *PostFlowVersionsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/{flowId}/versions][%d] postFlowVersionsOK  %+v", 200, o.Payload)
}

func (o *PostFlowVersionsOK) GetPayload() *models.FlowVersion {
	return o.Payload
}

func (o *PostFlowVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FlowVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowVersionsBadRequest creates a PostFlowVersionsBadRequest with default headers values
func NewPostFlowVersionsBadRequest() *PostFlowVersionsBadRequest {
	return &PostFlowVersionsBadRequest{}
}

/*PostFlowVersionsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostFlowVersionsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostFlowVersionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/{flowId}/versions][%d] postFlowVersionsBadRequest  %+v", 400, o.Payload)
}

func (o *PostFlowVersionsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowVersionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowVersionsUnauthorized creates a PostFlowVersionsUnauthorized with default headers values
func NewPostFlowVersionsUnauthorized() *PostFlowVersionsUnauthorized {
	return &PostFlowVersionsUnauthorized{}
}

/*PostFlowVersionsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostFlowVersionsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostFlowVersionsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/{flowId}/versions][%d] postFlowVersionsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostFlowVersionsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowVersionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowVersionsForbidden creates a PostFlowVersionsForbidden with default headers values
func NewPostFlowVersionsForbidden() *PostFlowVersionsForbidden {
	return &PostFlowVersionsForbidden{}
}

/*PostFlowVersionsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostFlowVersionsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostFlowVersionsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/{flowId}/versions][%d] postFlowVersionsForbidden  %+v", 403, o.Payload)
}

func (o *PostFlowVersionsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowVersionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowVersionsNotFound creates a PostFlowVersionsNotFound with default headers values
func NewPostFlowVersionsNotFound() *PostFlowVersionsNotFound {
	return &PostFlowVersionsNotFound{}
}

/*PostFlowVersionsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostFlowVersionsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostFlowVersionsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/{flowId}/versions][%d] postFlowVersionsNotFound  %+v", 404, o.Payload)
}

func (o *PostFlowVersionsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowVersionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowVersionsMethodNotAllowed creates a PostFlowVersionsMethodNotAllowed with default headers values
func NewPostFlowVersionsMethodNotAllowed() *PostFlowVersionsMethodNotAllowed {
	return &PostFlowVersionsMethodNotAllowed{}
}

/*PostFlowVersionsMethodNotAllowed handles this case with default header values.

Method Not Allowed
*/
type PostFlowVersionsMethodNotAllowed struct {
	Payload *models.ErrorBody
}

func (o *PostFlowVersionsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/{flowId}/versions][%d] postFlowVersionsMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *PostFlowVersionsMethodNotAllowed) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowVersionsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowVersionsRequestTimeout creates a PostFlowVersionsRequestTimeout with default headers values
func NewPostFlowVersionsRequestTimeout() *PostFlowVersionsRequestTimeout {
	return &PostFlowVersionsRequestTimeout{}
}

/*PostFlowVersionsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostFlowVersionsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostFlowVersionsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/{flowId}/versions][%d] postFlowVersionsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostFlowVersionsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowVersionsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowVersionsConflict creates a PostFlowVersionsConflict with default headers values
func NewPostFlowVersionsConflict() *PostFlowVersionsConflict {
	return &PostFlowVersionsConflict{}
}

/*PostFlowVersionsConflict handles this case with default header values.

Conflict
*/
type PostFlowVersionsConflict struct {
	Payload *models.ErrorBody
}

func (o *PostFlowVersionsConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/{flowId}/versions][%d] postFlowVersionsConflict  %+v", 409, o.Payload)
}

func (o *PostFlowVersionsConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowVersionsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowVersionsGone creates a PostFlowVersionsGone with default headers values
func NewPostFlowVersionsGone() *PostFlowVersionsGone {
	return &PostFlowVersionsGone{}
}

/*PostFlowVersionsGone handles this case with default header values.

Gone
*/
type PostFlowVersionsGone struct {
	Payload *models.ErrorBody
}

func (o *PostFlowVersionsGone) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/{flowId}/versions][%d] postFlowVersionsGone  %+v", 410, o.Payload)
}

func (o *PostFlowVersionsGone) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowVersionsGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowVersionsRequestEntityTooLarge creates a PostFlowVersionsRequestEntityTooLarge with default headers values
func NewPostFlowVersionsRequestEntityTooLarge() *PostFlowVersionsRequestEntityTooLarge {
	return &PostFlowVersionsRequestEntityTooLarge{}
}

/*PostFlowVersionsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostFlowVersionsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostFlowVersionsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/{flowId}/versions][%d] postFlowVersionsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostFlowVersionsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowVersionsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowVersionsUnsupportedMediaType creates a PostFlowVersionsUnsupportedMediaType with default headers values
func NewPostFlowVersionsUnsupportedMediaType() *PostFlowVersionsUnsupportedMediaType {
	return &PostFlowVersionsUnsupportedMediaType{}
}

/*PostFlowVersionsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostFlowVersionsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostFlowVersionsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/{flowId}/versions][%d] postFlowVersionsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostFlowVersionsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowVersionsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowVersionsTooManyRequests creates a PostFlowVersionsTooManyRequests with default headers values
func NewPostFlowVersionsTooManyRequests() *PostFlowVersionsTooManyRequests {
	return &PostFlowVersionsTooManyRequests{}
}

/*PostFlowVersionsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostFlowVersionsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostFlowVersionsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/{flowId}/versions][%d] postFlowVersionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostFlowVersionsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowVersionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowVersionsInternalServerError creates a PostFlowVersionsInternalServerError with default headers values
func NewPostFlowVersionsInternalServerError() *PostFlowVersionsInternalServerError {
	return &PostFlowVersionsInternalServerError{}
}

/*PostFlowVersionsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostFlowVersionsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostFlowVersionsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/{flowId}/versions][%d] postFlowVersionsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostFlowVersionsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowVersionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowVersionsServiceUnavailable creates a PostFlowVersionsServiceUnavailable with default headers values
func NewPostFlowVersionsServiceUnavailable() *PostFlowVersionsServiceUnavailable {
	return &PostFlowVersionsServiceUnavailable{}
}

/*PostFlowVersionsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostFlowVersionsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostFlowVersionsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/{flowId}/versions][%d] postFlowVersionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostFlowVersionsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowVersionsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowVersionsGatewayTimeout creates a PostFlowVersionsGatewayTimeout with default headers values
func NewPostFlowVersionsGatewayTimeout() *PostFlowVersionsGatewayTimeout {
	return &PostFlowVersionsGatewayTimeout{}
}

/*PostFlowVersionsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostFlowVersionsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostFlowVersionsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/{flowId}/versions][%d] postFlowVersionsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostFlowVersionsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowVersionsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

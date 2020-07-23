// Code generated by go-swagger; DO NOT EDIT.

package routing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeleteRoutingUserUtilizationReader is a Reader for the DeleteRoutingUserUtilization structure.
type DeleteRoutingUserUtilizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRoutingUserUtilizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRoutingUserUtilizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteRoutingUserUtilizationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteRoutingUserUtilizationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteRoutingUserUtilizationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRoutingUserUtilizationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteRoutingUserUtilizationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteRoutingUserUtilizationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteRoutingUserUtilizationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteRoutingUserUtilizationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteRoutingUserUtilizationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteRoutingUserUtilizationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRoutingUserUtilizationOK creates a DeleteRoutingUserUtilizationOK with default headers values
func NewDeleteRoutingUserUtilizationOK() *DeleteRoutingUserUtilizationOK {
	return &DeleteRoutingUserUtilizationOK{}
}

/*DeleteRoutingUserUtilizationOK handles this case with default header values.

Operation was successful.
*/
type DeleteRoutingUserUtilizationOK struct {
}

func (o *DeleteRoutingUserUtilizationOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/users/{userId}/utilization][%d] deleteRoutingUserUtilizationOK ", 200)
}

func (o *DeleteRoutingUserUtilizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRoutingUserUtilizationBadRequest creates a DeleteRoutingUserUtilizationBadRequest with default headers values
func NewDeleteRoutingUserUtilizationBadRequest() *DeleteRoutingUserUtilizationBadRequest {
	return &DeleteRoutingUserUtilizationBadRequest{}
}

/*DeleteRoutingUserUtilizationBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteRoutingUserUtilizationBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingUserUtilizationBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/users/{userId}/utilization][%d] deleteRoutingUserUtilizationBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteRoutingUserUtilizationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingUserUtilizationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingUserUtilizationUnauthorized creates a DeleteRoutingUserUtilizationUnauthorized with default headers values
func NewDeleteRoutingUserUtilizationUnauthorized() *DeleteRoutingUserUtilizationUnauthorized {
	return &DeleteRoutingUserUtilizationUnauthorized{}
}

/*DeleteRoutingUserUtilizationUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteRoutingUserUtilizationUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingUserUtilizationUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/users/{userId}/utilization][%d] deleteRoutingUserUtilizationUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRoutingUserUtilizationUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingUserUtilizationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingUserUtilizationForbidden creates a DeleteRoutingUserUtilizationForbidden with default headers values
func NewDeleteRoutingUserUtilizationForbidden() *DeleteRoutingUserUtilizationForbidden {
	return &DeleteRoutingUserUtilizationForbidden{}
}

/*DeleteRoutingUserUtilizationForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteRoutingUserUtilizationForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingUserUtilizationForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/users/{userId}/utilization][%d] deleteRoutingUserUtilizationForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRoutingUserUtilizationForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingUserUtilizationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingUserUtilizationNotFound creates a DeleteRoutingUserUtilizationNotFound with default headers values
func NewDeleteRoutingUserUtilizationNotFound() *DeleteRoutingUserUtilizationNotFound {
	return &DeleteRoutingUserUtilizationNotFound{}
}

/*DeleteRoutingUserUtilizationNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteRoutingUserUtilizationNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingUserUtilizationNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/users/{userId}/utilization][%d] deleteRoutingUserUtilizationNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRoutingUserUtilizationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingUserUtilizationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingUserUtilizationRequestEntityTooLarge creates a DeleteRoutingUserUtilizationRequestEntityTooLarge with default headers values
func NewDeleteRoutingUserUtilizationRequestEntityTooLarge() *DeleteRoutingUserUtilizationRequestEntityTooLarge {
	return &DeleteRoutingUserUtilizationRequestEntityTooLarge{}
}

/*DeleteRoutingUserUtilizationRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteRoutingUserUtilizationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingUserUtilizationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/users/{userId}/utilization][%d] deleteRoutingUserUtilizationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteRoutingUserUtilizationRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingUserUtilizationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingUserUtilizationUnsupportedMediaType creates a DeleteRoutingUserUtilizationUnsupportedMediaType with default headers values
func NewDeleteRoutingUserUtilizationUnsupportedMediaType() *DeleteRoutingUserUtilizationUnsupportedMediaType {
	return &DeleteRoutingUserUtilizationUnsupportedMediaType{}
}

/*DeleteRoutingUserUtilizationUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteRoutingUserUtilizationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingUserUtilizationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/users/{userId}/utilization][%d] deleteRoutingUserUtilizationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteRoutingUserUtilizationUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingUserUtilizationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingUserUtilizationTooManyRequests creates a DeleteRoutingUserUtilizationTooManyRequests with default headers values
func NewDeleteRoutingUserUtilizationTooManyRequests() *DeleteRoutingUserUtilizationTooManyRequests {
	return &DeleteRoutingUserUtilizationTooManyRequests{}
}

/*DeleteRoutingUserUtilizationTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteRoutingUserUtilizationTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingUserUtilizationTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/users/{userId}/utilization][%d] deleteRoutingUserUtilizationTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteRoutingUserUtilizationTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingUserUtilizationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingUserUtilizationInternalServerError creates a DeleteRoutingUserUtilizationInternalServerError with default headers values
func NewDeleteRoutingUserUtilizationInternalServerError() *DeleteRoutingUserUtilizationInternalServerError {
	return &DeleteRoutingUserUtilizationInternalServerError{}
}

/*DeleteRoutingUserUtilizationInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteRoutingUserUtilizationInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingUserUtilizationInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/users/{userId}/utilization][%d] deleteRoutingUserUtilizationInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRoutingUserUtilizationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingUserUtilizationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingUserUtilizationServiceUnavailable creates a DeleteRoutingUserUtilizationServiceUnavailable with default headers values
func NewDeleteRoutingUserUtilizationServiceUnavailable() *DeleteRoutingUserUtilizationServiceUnavailable {
	return &DeleteRoutingUserUtilizationServiceUnavailable{}
}

/*DeleteRoutingUserUtilizationServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteRoutingUserUtilizationServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingUserUtilizationServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/users/{userId}/utilization][%d] deleteRoutingUserUtilizationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteRoutingUserUtilizationServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingUserUtilizationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingUserUtilizationGatewayTimeout creates a DeleteRoutingUserUtilizationGatewayTimeout with default headers values
func NewDeleteRoutingUserUtilizationGatewayTimeout() *DeleteRoutingUserUtilizationGatewayTimeout {
	return &DeleteRoutingUserUtilizationGatewayTimeout{}
}

/*DeleteRoutingUserUtilizationGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteRoutingUserUtilizationGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingUserUtilizationGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/users/{userId}/utilization][%d] deleteRoutingUserUtilizationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteRoutingUserUtilizationGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingUserUtilizationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

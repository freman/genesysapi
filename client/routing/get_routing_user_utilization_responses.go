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

// GetRoutingUserUtilizationReader is a Reader for the GetRoutingUserUtilization structure.
type GetRoutingUserUtilizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoutingUserUtilizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoutingUserUtilizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRoutingUserUtilizationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRoutingUserUtilizationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRoutingUserUtilizationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRoutingUserUtilizationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRoutingUserUtilizationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRoutingUserUtilizationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRoutingUserUtilizationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRoutingUserUtilizationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRoutingUserUtilizationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRoutingUserUtilizationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRoutingUserUtilizationOK creates a GetRoutingUserUtilizationOK with default headers values
func NewGetRoutingUserUtilizationOK() *GetRoutingUserUtilizationOK {
	return &GetRoutingUserUtilizationOK{}
}

/*GetRoutingUserUtilizationOK handles this case with default header values.

successful operation
*/
type GetRoutingUserUtilizationOK struct {
	Payload *models.Utilization
}

func (o *GetRoutingUserUtilizationOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/users/{userId}/utilization][%d] getRoutingUserUtilizationOK  %+v", 200, o.Payload)
}

func (o *GetRoutingUserUtilizationOK) GetPayload() *models.Utilization {
	return o.Payload
}

func (o *GetRoutingUserUtilizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Utilization)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingUserUtilizationBadRequest creates a GetRoutingUserUtilizationBadRequest with default headers values
func NewGetRoutingUserUtilizationBadRequest() *GetRoutingUserUtilizationBadRequest {
	return &GetRoutingUserUtilizationBadRequest{}
}

/*GetRoutingUserUtilizationBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRoutingUserUtilizationBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingUserUtilizationBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/users/{userId}/utilization][%d] getRoutingUserUtilizationBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingUserUtilizationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingUserUtilizationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingUserUtilizationUnauthorized creates a GetRoutingUserUtilizationUnauthorized with default headers values
func NewGetRoutingUserUtilizationUnauthorized() *GetRoutingUserUtilizationUnauthorized {
	return &GetRoutingUserUtilizationUnauthorized{}
}

/*GetRoutingUserUtilizationUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRoutingUserUtilizationUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingUserUtilizationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/users/{userId}/utilization][%d] getRoutingUserUtilizationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingUserUtilizationUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingUserUtilizationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingUserUtilizationForbidden creates a GetRoutingUserUtilizationForbidden with default headers values
func NewGetRoutingUserUtilizationForbidden() *GetRoutingUserUtilizationForbidden {
	return &GetRoutingUserUtilizationForbidden{}
}

/*GetRoutingUserUtilizationForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetRoutingUserUtilizationForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingUserUtilizationForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/users/{userId}/utilization][%d] getRoutingUserUtilizationForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingUserUtilizationForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingUserUtilizationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingUserUtilizationNotFound creates a GetRoutingUserUtilizationNotFound with default headers values
func NewGetRoutingUserUtilizationNotFound() *GetRoutingUserUtilizationNotFound {
	return &GetRoutingUserUtilizationNotFound{}
}

/*GetRoutingUserUtilizationNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetRoutingUserUtilizationNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingUserUtilizationNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/users/{userId}/utilization][%d] getRoutingUserUtilizationNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingUserUtilizationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingUserUtilizationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingUserUtilizationRequestEntityTooLarge creates a GetRoutingUserUtilizationRequestEntityTooLarge with default headers values
func NewGetRoutingUserUtilizationRequestEntityTooLarge() *GetRoutingUserUtilizationRequestEntityTooLarge {
	return &GetRoutingUserUtilizationRequestEntityTooLarge{}
}

/*GetRoutingUserUtilizationRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetRoutingUserUtilizationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingUserUtilizationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/users/{userId}/utilization][%d] getRoutingUserUtilizationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingUserUtilizationRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingUserUtilizationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingUserUtilizationUnsupportedMediaType creates a GetRoutingUserUtilizationUnsupportedMediaType with default headers values
func NewGetRoutingUserUtilizationUnsupportedMediaType() *GetRoutingUserUtilizationUnsupportedMediaType {
	return &GetRoutingUserUtilizationUnsupportedMediaType{}
}

/*GetRoutingUserUtilizationUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRoutingUserUtilizationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingUserUtilizationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/users/{userId}/utilization][%d] getRoutingUserUtilizationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingUserUtilizationUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingUserUtilizationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingUserUtilizationTooManyRequests creates a GetRoutingUserUtilizationTooManyRequests with default headers values
func NewGetRoutingUserUtilizationTooManyRequests() *GetRoutingUserUtilizationTooManyRequests {
	return &GetRoutingUserUtilizationTooManyRequests{}
}

/*GetRoutingUserUtilizationTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetRoutingUserUtilizationTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingUserUtilizationTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/users/{userId}/utilization][%d] getRoutingUserUtilizationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingUserUtilizationTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingUserUtilizationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingUserUtilizationInternalServerError creates a GetRoutingUserUtilizationInternalServerError with default headers values
func NewGetRoutingUserUtilizationInternalServerError() *GetRoutingUserUtilizationInternalServerError {
	return &GetRoutingUserUtilizationInternalServerError{}
}

/*GetRoutingUserUtilizationInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRoutingUserUtilizationInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingUserUtilizationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/users/{userId}/utilization][%d] getRoutingUserUtilizationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingUserUtilizationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingUserUtilizationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingUserUtilizationServiceUnavailable creates a GetRoutingUserUtilizationServiceUnavailable with default headers values
func NewGetRoutingUserUtilizationServiceUnavailable() *GetRoutingUserUtilizationServiceUnavailable {
	return &GetRoutingUserUtilizationServiceUnavailable{}
}

/*GetRoutingUserUtilizationServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRoutingUserUtilizationServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingUserUtilizationServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/users/{userId}/utilization][%d] getRoutingUserUtilizationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingUserUtilizationServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingUserUtilizationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingUserUtilizationGatewayTimeout creates a GetRoutingUserUtilizationGatewayTimeout with default headers values
func NewGetRoutingUserUtilizationGatewayTimeout() *GetRoutingUserUtilizationGatewayTimeout {
	return &GetRoutingUserUtilizationGatewayTimeout{}
}

/*GetRoutingUserUtilizationGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetRoutingUserUtilizationGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingUserUtilizationGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/users/{userId}/utilization][%d] getRoutingUserUtilizationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingUserUtilizationGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingUserUtilizationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

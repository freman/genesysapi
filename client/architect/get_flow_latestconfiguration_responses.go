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

// GetFlowLatestconfigurationReader is a Reader for the GetFlowLatestconfiguration structure.
type GetFlowLatestconfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFlowLatestconfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFlowLatestconfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFlowLatestconfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFlowLatestconfigurationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFlowLatestconfigurationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFlowLatestconfigurationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetFlowLatestconfigurationMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetFlowLatestconfigurationRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetFlowLatestconfigurationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetFlowLatestconfigurationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetFlowLatestconfigurationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFlowLatestconfigurationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetFlowLatestconfigurationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetFlowLatestconfigurationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFlowLatestconfigurationOK creates a GetFlowLatestconfigurationOK with default headers values
func NewGetFlowLatestconfigurationOK() *GetFlowLatestconfigurationOK {
	return &GetFlowLatestconfigurationOK{}
}

/*GetFlowLatestconfigurationOK handles this case with default header values.

successful operation
*/
type GetFlowLatestconfigurationOK struct {
	Payload interface{}
}

func (o *GetFlowLatestconfigurationOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/latestconfiguration][%d] getFlowLatestconfigurationOK  %+v", 200, o.Payload)
}

func (o *GetFlowLatestconfigurationOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetFlowLatestconfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowLatestconfigurationBadRequest creates a GetFlowLatestconfigurationBadRequest with default headers values
func NewGetFlowLatestconfigurationBadRequest() *GetFlowLatestconfigurationBadRequest {
	return &GetFlowLatestconfigurationBadRequest{}
}

/*GetFlowLatestconfigurationBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetFlowLatestconfigurationBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetFlowLatestconfigurationBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/latestconfiguration][%d] getFlowLatestconfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *GetFlowLatestconfigurationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowLatestconfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowLatestconfigurationUnauthorized creates a GetFlowLatestconfigurationUnauthorized with default headers values
func NewGetFlowLatestconfigurationUnauthorized() *GetFlowLatestconfigurationUnauthorized {
	return &GetFlowLatestconfigurationUnauthorized{}
}

/*GetFlowLatestconfigurationUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetFlowLatestconfigurationUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetFlowLatestconfigurationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/latestconfiguration][%d] getFlowLatestconfigurationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFlowLatestconfigurationUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowLatestconfigurationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowLatestconfigurationForbidden creates a GetFlowLatestconfigurationForbidden with default headers values
func NewGetFlowLatestconfigurationForbidden() *GetFlowLatestconfigurationForbidden {
	return &GetFlowLatestconfigurationForbidden{}
}

/*GetFlowLatestconfigurationForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetFlowLatestconfigurationForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetFlowLatestconfigurationForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/latestconfiguration][%d] getFlowLatestconfigurationForbidden  %+v", 403, o.Payload)
}

func (o *GetFlowLatestconfigurationForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowLatestconfigurationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowLatestconfigurationNotFound creates a GetFlowLatestconfigurationNotFound with default headers values
func NewGetFlowLatestconfigurationNotFound() *GetFlowLatestconfigurationNotFound {
	return &GetFlowLatestconfigurationNotFound{}
}

/*GetFlowLatestconfigurationNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetFlowLatestconfigurationNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetFlowLatestconfigurationNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/latestconfiguration][%d] getFlowLatestconfigurationNotFound  %+v", 404, o.Payload)
}

func (o *GetFlowLatestconfigurationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowLatestconfigurationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowLatestconfigurationMethodNotAllowed creates a GetFlowLatestconfigurationMethodNotAllowed with default headers values
func NewGetFlowLatestconfigurationMethodNotAllowed() *GetFlowLatestconfigurationMethodNotAllowed {
	return &GetFlowLatestconfigurationMethodNotAllowed{}
}

/*GetFlowLatestconfigurationMethodNotAllowed handles this case with default header values.

Method Not Allowed
*/
type GetFlowLatestconfigurationMethodNotAllowed struct {
	Payload *models.ErrorBody
}

func (o *GetFlowLatestconfigurationMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/latestconfiguration][%d] getFlowLatestconfigurationMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GetFlowLatestconfigurationMethodNotAllowed) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowLatestconfigurationMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowLatestconfigurationRequestTimeout creates a GetFlowLatestconfigurationRequestTimeout with default headers values
func NewGetFlowLatestconfigurationRequestTimeout() *GetFlowLatestconfigurationRequestTimeout {
	return &GetFlowLatestconfigurationRequestTimeout{}
}

/*GetFlowLatestconfigurationRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetFlowLatestconfigurationRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetFlowLatestconfigurationRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/latestconfiguration][%d] getFlowLatestconfigurationRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetFlowLatestconfigurationRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowLatestconfigurationRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowLatestconfigurationRequestEntityTooLarge creates a GetFlowLatestconfigurationRequestEntityTooLarge with default headers values
func NewGetFlowLatestconfigurationRequestEntityTooLarge() *GetFlowLatestconfigurationRequestEntityTooLarge {
	return &GetFlowLatestconfigurationRequestEntityTooLarge{}
}

/*GetFlowLatestconfigurationRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetFlowLatestconfigurationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetFlowLatestconfigurationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/latestconfiguration][%d] getFlowLatestconfigurationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetFlowLatestconfigurationRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowLatestconfigurationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowLatestconfigurationUnsupportedMediaType creates a GetFlowLatestconfigurationUnsupportedMediaType with default headers values
func NewGetFlowLatestconfigurationUnsupportedMediaType() *GetFlowLatestconfigurationUnsupportedMediaType {
	return &GetFlowLatestconfigurationUnsupportedMediaType{}
}

/*GetFlowLatestconfigurationUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetFlowLatestconfigurationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetFlowLatestconfigurationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/latestconfiguration][%d] getFlowLatestconfigurationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetFlowLatestconfigurationUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowLatestconfigurationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowLatestconfigurationTooManyRequests creates a GetFlowLatestconfigurationTooManyRequests with default headers values
func NewGetFlowLatestconfigurationTooManyRequests() *GetFlowLatestconfigurationTooManyRequests {
	return &GetFlowLatestconfigurationTooManyRequests{}
}

/*GetFlowLatestconfigurationTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetFlowLatestconfigurationTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetFlowLatestconfigurationTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/latestconfiguration][%d] getFlowLatestconfigurationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetFlowLatestconfigurationTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowLatestconfigurationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowLatestconfigurationInternalServerError creates a GetFlowLatestconfigurationInternalServerError with default headers values
func NewGetFlowLatestconfigurationInternalServerError() *GetFlowLatestconfigurationInternalServerError {
	return &GetFlowLatestconfigurationInternalServerError{}
}

/*GetFlowLatestconfigurationInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetFlowLatestconfigurationInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetFlowLatestconfigurationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/latestconfiguration][%d] getFlowLatestconfigurationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFlowLatestconfigurationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowLatestconfigurationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowLatestconfigurationServiceUnavailable creates a GetFlowLatestconfigurationServiceUnavailable with default headers values
func NewGetFlowLatestconfigurationServiceUnavailable() *GetFlowLatestconfigurationServiceUnavailable {
	return &GetFlowLatestconfigurationServiceUnavailable{}
}

/*GetFlowLatestconfigurationServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetFlowLatestconfigurationServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetFlowLatestconfigurationServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/latestconfiguration][%d] getFlowLatestconfigurationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFlowLatestconfigurationServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowLatestconfigurationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowLatestconfigurationGatewayTimeout creates a GetFlowLatestconfigurationGatewayTimeout with default headers values
func NewGetFlowLatestconfigurationGatewayTimeout() *GetFlowLatestconfigurationGatewayTimeout {
	return &GetFlowLatestconfigurationGatewayTimeout{}
}

/*GetFlowLatestconfigurationGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetFlowLatestconfigurationGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetFlowLatestconfigurationGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/latestconfiguration][%d] getFlowLatestconfigurationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetFlowLatestconfigurationGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowLatestconfigurationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

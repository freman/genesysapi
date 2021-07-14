// Code generated by go-swagger; DO NOT EDIT.

package general_data_protection_regulation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetGdprRequestReader is a Reader for the GetGdprRequest structure.
type GetGdprRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGdprRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGdprRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGdprRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetGdprRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGdprRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGdprRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetGdprRequestRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetGdprRequestRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetGdprRequestUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetGdprRequestTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGdprRequestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetGdprRequestServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetGdprRequestGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGdprRequestOK creates a GetGdprRequestOK with default headers values
func NewGetGdprRequestOK() *GetGdprRequestOK {
	return &GetGdprRequestOK{}
}

/*GetGdprRequestOK handles this case with default header values.

successful operation
*/
type GetGdprRequestOK struct {
	Payload *models.GDPRRequest
}

func (o *GetGdprRequestOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/gdpr/requests/{requestId}][%d] getGdprRequestOK  %+v", 200, o.Payload)
}

func (o *GetGdprRequestOK) GetPayload() *models.GDPRRequest {
	return o.Payload
}

func (o *GetGdprRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GDPRRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGdprRequestBadRequest creates a GetGdprRequestBadRequest with default headers values
func NewGetGdprRequestBadRequest() *GetGdprRequestBadRequest {
	return &GetGdprRequestBadRequest{}
}

/*GetGdprRequestBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetGdprRequestBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetGdprRequestBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/gdpr/requests/{requestId}][%d] getGdprRequestBadRequest  %+v", 400, o.Payload)
}

func (o *GetGdprRequestBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGdprRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGdprRequestUnauthorized creates a GetGdprRequestUnauthorized with default headers values
func NewGetGdprRequestUnauthorized() *GetGdprRequestUnauthorized {
	return &GetGdprRequestUnauthorized{}
}

/*GetGdprRequestUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetGdprRequestUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetGdprRequestUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/gdpr/requests/{requestId}][%d] getGdprRequestUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGdprRequestUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGdprRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGdprRequestForbidden creates a GetGdprRequestForbidden with default headers values
func NewGetGdprRequestForbidden() *GetGdprRequestForbidden {
	return &GetGdprRequestForbidden{}
}

/*GetGdprRequestForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetGdprRequestForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetGdprRequestForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/gdpr/requests/{requestId}][%d] getGdprRequestForbidden  %+v", 403, o.Payload)
}

func (o *GetGdprRequestForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGdprRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGdprRequestNotFound creates a GetGdprRequestNotFound with default headers values
func NewGetGdprRequestNotFound() *GetGdprRequestNotFound {
	return &GetGdprRequestNotFound{}
}

/*GetGdprRequestNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetGdprRequestNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetGdprRequestNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/gdpr/requests/{requestId}][%d] getGdprRequestNotFound  %+v", 404, o.Payload)
}

func (o *GetGdprRequestNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGdprRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGdprRequestRequestTimeout creates a GetGdprRequestRequestTimeout with default headers values
func NewGetGdprRequestRequestTimeout() *GetGdprRequestRequestTimeout {
	return &GetGdprRequestRequestTimeout{}
}

/*GetGdprRequestRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetGdprRequestRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetGdprRequestRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/gdpr/requests/{requestId}][%d] getGdprRequestRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetGdprRequestRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGdprRequestRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGdprRequestRequestEntityTooLarge creates a GetGdprRequestRequestEntityTooLarge with default headers values
func NewGetGdprRequestRequestEntityTooLarge() *GetGdprRequestRequestEntityTooLarge {
	return &GetGdprRequestRequestEntityTooLarge{}
}

/*GetGdprRequestRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetGdprRequestRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetGdprRequestRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/gdpr/requests/{requestId}][%d] getGdprRequestRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetGdprRequestRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGdprRequestRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGdprRequestUnsupportedMediaType creates a GetGdprRequestUnsupportedMediaType with default headers values
func NewGetGdprRequestUnsupportedMediaType() *GetGdprRequestUnsupportedMediaType {
	return &GetGdprRequestUnsupportedMediaType{}
}

/*GetGdprRequestUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetGdprRequestUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetGdprRequestUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/gdpr/requests/{requestId}][%d] getGdprRequestUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetGdprRequestUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGdprRequestUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGdprRequestTooManyRequests creates a GetGdprRequestTooManyRequests with default headers values
func NewGetGdprRequestTooManyRequests() *GetGdprRequestTooManyRequests {
	return &GetGdprRequestTooManyRequests{}
}

/*GetGdprRequestTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetGdprRequestTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetGdprRequestTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/gdpr/requests/{requestId}][%d] getGdprRequestTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetGdprRequestTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGdprRequestTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGdprRequestInternalServerError creates a GetGdprRequestInternalServerError with default headers values
func NewGetGdprRequestInternalServerError() *GetGdprRequestInternalServerError {
	return &GetGdprRequestInternalServerError{}
}

/*GetGdprRequestInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetGdprRequestInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetGdprRequestInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/gdpr/requests/{requestId}][%d] getGdprRequestInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGdprRequestInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGdprRequestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGdprRequestServiceUnavailable creates a GetGdprRequestServiceUnavailable with default headers values
func NewGetGdprRequestServiceUnavailable() *GetGdprRequestServiceUnavailable {
	return &GetGdprRequestServiceUnavailable{}
}

/*GetGdprRequestServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetGdprRequestServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetGdprRequestServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/gdpr/requests/{requestId}][%d] getGdprRequestServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetGdprRequestServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGdprRequestServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGdprRequestGatewayTimeout creates a GetGdprRequestGatewayTimeout with default headers values
func NewGetGdprRequestGatewayTimeout() *GetGdprRequestGatewayTimeout {
	return &GetGdprRequestGatewayTimeout{}
}

/*GetGdprRequestGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetGdprRequestGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetGdprRequestGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/gdpr/requests/{requestId}][%d] getGdprRequestGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetGdprRequestGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGdprRequestGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

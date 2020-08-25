// Code generated by go-swagger; DO NOT EDIT.

package telephony_providers_edge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetTelephonyProvidersEdgesDidpoolReader is a Reader for the GetTelephonyProvidersEdgesDidpool structure.
type GetTelephonyProvidersEdgesDidpoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonyProvidersEdgesDidpoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonyProvidersEdgesDidpoolOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonyProvidersEdgesDidpoolBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonyProvidersEdgesDidpoolUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonyProvidersEdgesDidpoolForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonyProvidersEdgesDidpoolNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonyProvidersEdgesDidpoolRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonyProvidersEdgesDidpoolUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonyProvidersEdgesDidpoolTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonyProvidersEdgesDidpoolInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonyProvidersEdgesDidpoolServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonyProvidersEdgesDidpoolGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonyProvidersEdgesDidpoolOK creates a GetTelephonyProvidersEdgesDidpoolOK with default headers values
func NewGetTelephonyProvidersEdgesDidpoolOK() *GetTelephonyProvidersEdgesDidpoolOK {
	return &GetTelephonyProvidersEdgesDidpoolOK{}
}

/*GetTelephonyProvidersEdgesDidpoolOK handles this case with default header values.

successful operation
*/
type GetTelephonyProvidersEdgesDidpoolOK struct {
	Payload *models.DIDPool
}

func (o *GetTelephonyProvidersEdgesDidpoolOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] getTelephonyProvidersEdgesDidpoolOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolOK) GetPayload() *models.DIDPool {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidpoolOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DIDPool)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidpoolBadRequest creates a GetTelephonyProvidersEdgesDidpoolBadRequest with default headers values
func NewGetTelephonyProvidersEdgesDidpoolBadRequest() *GetTelephonyProvidersEdgesDidpoolBadRequest {
	return &GetTelephonyProvidersEdgesDidpoolBadRequest{}
}

/*GetTelephonyProvidersEdgesDidpoolBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyProvidersEdgesDidpoolBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesDidpoolBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] getTelephonyProvidersEdgesDidpoolBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidpoolBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidpoolUnauthorized creates a GetTelephonyProvidersEdgesDidpoolUnauthorized with default headers values
func NewGetTelephonyProvidersEdgesDidpoolUnauthorized() *GetTelephonyProvidersEdgesDidpoolUnauthorized {
	return &GetTelephonyProvidersEdgesDidpoolUnauthorized{}
}

/*GetTelephonyProvidersEdgesDidpoolUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyProvidersEdgesDidpoolUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesDidpoolUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] getTelephonyProvidersEdgesDidpoolUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidpoolUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidpoolForbidden creates a GetTelephonyProvidersEdgesDidpoolForbidden with default headers values
func NewGetTelephonyProvidersEdgesDidpoolForbidden() *GetTelephonyProvidersEdgesDidpoolForbidden {
	return &GetTelephonyProvidersEdgesDidpoolForbidden{}
}

/*GetTelephonyProvidersEdgesDidpoolForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyProvidersEdgesDidpoolForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesDidpoolForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] getTelephonyProvidersEdgesDidpoolForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidpoolForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidpoolNotFound creates a GetTelephonyProvidersEdgesDidpoolNotFound with default headers values
func NewGetTelephonyProvidersEdgesDidpoolNotFound() *GetTelephonyProvidersEdgesDidpoolNotFound {
	return &GetTelephonyProvidersEdgesDidpoolNotFound{}
}

/*GetTelephonyProvidersEdgesDidpoolNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetTelephonyProvidersEdgesDidpoolNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesDidpoolNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] getTelephonyProvidersEdgesDidpoolNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidpoolNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidpoolRequestEntityTooLarge creates a GetTelephonyProvidersEdgesDidpoolRequestEntityTooLarge with default headers values
func NewGetTelephonyProvidersEdgesDidpoolRequestEntityTooLarge() *GetTelephonyProvidersEdgesDidpoolRequestEntityTooLarge {
	return &GetTelephonyProvidersEdgesDidpoolRequestEntityTooLarge{}
}

/*GetTelephonyProvidersEdgesDidpoolRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetTelephonyProvidersEdgesDidpoolRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesDidpoolRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] getTelephonyProvidersEdgesDidpoolRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidpoolRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidpoolUnsupportedMediaType creates a GetTelephonyProvidersEdgesDidpoolUnsupportedMediaType with default headers values
func NewGetTelephonyProvidersEdgesDidpoolUnsupportedMediaType() *GetTelephonyProvidersEdgesDidpoolUnsupportedMediaType {
	return &GetTelephonyProvidersEdgesDidpoolUnsupportedMediaType{}
}

/*GetTelephonyProvidersEdgesDidpoolUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyProvidersEdgesDidpoolUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesDidpoolUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] getTelephonyProvidersEdgesDidpoolUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidpoolUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidpoolTooManyRequests creates a GetTelephonyProvidersEdgesDidpoolTooManyRequests with default headers values
func NewGetTelephonyProvidersEdgesDidpoolTooManyRequests() *GetTelephonyProvidersEdgesDidpoolTooManyRequests {
	return &GetTelephonyProvidersEdgesDidpoolTooManyRequests{}
}

/*GetTelephonyProvidersEdgesDidpoolTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetTelephonyProvidersEdgesDidpoolTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesDidpoolTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] getTelephonyProvidersEdgesDidpoolTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidpoolTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidpoolInternalServerError creates a GetTelephonyProvidersEdgesDidpoolInternalServerError with default headers values
func NewGetTelephonyProvidersEdgesDidpoolInternalServerError() *GetTelephonyProvidersEdgesDidpoolInternalServerError {
	return &GetTelephonyProvidersEdgesDidpoolInternalServerError{}
}

/*GetTelephonyProvidersEdgesDidpoolInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyProvidersEdgesDidpoolInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesDidpoolInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] getTelephonyProvidersEdgesDidpoolInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidpoolInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidpoolServiceUnavailable creates a GetTelephonyProvidersEdgesDidpoolServiceUnavailable with default headers values
func NewGetTelephonyProvidersEdgesDidpoolServiceUnavailable() *GetTelephonyProvidersEdgesDidpoolServiceUnavailable {
	return &GetTelephonyProvidersEdgesDidpoolServiceUnavailable{}
}

/*GetTelephonyProvidersEdgesDidpoolServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyProvidersEdgesDidpoolServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesDidpoolServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] getTelephonyProvidersEdgesDidpoolServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidpoolServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidpoolGatewayTimeout creates a GetTelephonyProvidersEdgesDidpoolGatewayTimeout with default headers values
func NewGetTelephonyProvidersEdgesDidpoolGatewayTimeout() *GetTelephonyProvidersEdgesDidpoolGatewayTimeout {
	return &GetTelephonyProvidersEdgesDidpoolGatewayTimeout{}
}

/*GetTelephonyProvidersEdgesDidpoolGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetTelephonyProvidersEdgesDidpoolGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesDidpoolGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] getTelephonyProvidersEdgesDidpoolGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidpoolGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
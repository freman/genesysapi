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

// GetTelephonyProvidersEdgesExtensionsReader is a Reader for the GetTelephonyProvidersEdgesExtensions structure.
type GetTelephonyProvidersEdgesExtensionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonyProvidersEdgesExtensionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonyProvidersEdgesExtensionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonyProvidersEdgesExtensionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonyProvidersEdgesExtensionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonyProvidersEdgesExtensionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonyProvidersEdgesExtensionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetTelephonyProvidersEdgesExtensionsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonyProvidersEdgesExtensionsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonyProvidersEdgesExtensionsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonyProvidersEdgesExtensionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonyProvidersEdgesExtensionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonyProvidersEdgesExtensionsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonyProvidersEdgesExtensionsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonyProvidersEdgesExtensionsOK creates a GetTelephonyProvidersEdgesExtensionsOK with default headers values
func NewGetTelephonyProvidersEdgesExtensionsOK() *GetTelephonyProvidersEdgesExtensionsOK {
	return &GetTelephonyProvidersEdgesExtensionsOK{}
}

/*GetTelephonyProvidersEdgesExtensionsOK handles this case with default header values.

successful operation
*/
type GetTelephonyProvidersEdgesExtensionsOK struct {
	Payload *models.ExtensionEntityListing
}

func (o *GetTelephonyProvidersEdgesExtensionsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/extensions][%d] getTelephonyProvidersEdgesExtensionsOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgesExtensionsOK) GetPayload() *models.ExtensionEntityListing {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesExtensionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExtensionEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesExtensionsBadRequest creates a GetTelephonyProvidersEdgesExtensionsBadRequest with default headers values
func NewGetTelephonyProvidersEdgesExtensionsBadRequest() *GetTelephonyProvidersEdgesExtensionsBadRequest {
	return &GetTelephonyProvidersEdgesExtensionsBadRequest{}
}

/*GetTelephonyProvidersEdgesExtensionsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyProvidersEdgesExtensionsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesExtensionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/extensions][%d] getTelephonyProvidersEdgesExtensionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgesExtensionsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesExtensionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesExtensionsUnauthorized creates a GetTelephonyProvidersEdgesExtensionsUnauthorized with default headers values
func NewGetTelephonyProvidersEdgesExtensionsUnauthorized() *GetTelephonyProvidersEdgesExtensionsUnauthorized {
	return &GetTelephonyProvidersEdgesExtensionsUnauthorized{}
}

/*GetTelephonyProvidersEdgesExtensionsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyProvidersEdgesExtensionsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesExtensionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/extensions][%d] getTelephonyProvidersEdgesExtensionsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgesExtensionsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesExtensionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesExtensionsForbidden creates a GetTelephonyProvidersEdgesExtensionsForbidden with default headers values
func NewGetTelephonyProvidersEdgesExtensionsForbidden() *GetTelephonyProvidersEdgesExtensionsForbidden {
	return &GetTelephonyProvidersEdgesExtensionsForbidden{}
}

/*GetTelephonyProvidersEdgesExtensionsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyProvidersEdgesExtensionsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesExtensionsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/extensions][%d] getTelephonyProvidersEdgesExtensionsForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgesExtensionsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesExtensionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesExtensionsNotFound creates a GetTelephonyProvidersEdgesExtensionsNotFound with default headers values
func NewGetTelephonyProvidersEdgesExtensionsNotFound() *GetTelephonyProvidersEdgesExtensionsNotFound {
	return &GetTelephonyProvidersEdgesExtensionsNotFound{}
}

/*GetTelephonyProvidersEdgesExtensionsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetTelephonyProvidersEdgesExtensionsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesExtensionsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/extensions][%d] getTelephonyProvidersEdgesExtensionsNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgesExtensionsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesExtensionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesExtensionsRequestTimeout creates a GetTelephonyProvidersEdgesExtensionsRequestTimeout with default headers values
func NewGetTelephonyProvidersEdgesExtensionsRequestTimeout() *GetTelephonyProvidersEdgesExtensionsRequestTimeout {
	return &GetTelephonyProvidersEdgesExtensionsRequestTimeout{}
}

/*GetTelephonyProvidersEdgesExtensionsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetTelephonyProvidersEdgesExtensionsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesExtensionsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/extensions][%d] getTelephonyProvidersEdgesExtensionsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTelephonyProvidersEdgesExtensionsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesExtensionsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesExtensionsRequestEntityTooLarge creates a GetTelephonyProvidersEdgesExtensionsRequestEntityTooLarge with default headers values
func NewGetTelephonyProvidersEdgesExtensionsRequestEntityTooLarge() *GetTelephonyProvidersEdgesExtensionsRequestEntityTooLarge {
	return &GetTelephonyProvidersEdgesExtensionsRequestEntityTooLarge{}
}

/*GetTelephonyProvidersEdgesExtensionsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetTelephonyProvidersEdgesExtensionsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesExtensionsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/extensions][%d] getTelephonyProvidersEdgesExtensionsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgesExtensionsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesExtensionsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesExtensionsUnsupportedMediaType creates a GetTelephonyProvidersEdgesExtensionsUnsupportedMediaType with default headers values
func NewGetTelephonyProvidersEdgesExtensionsUnsupportedMediaType() *GetTelephonyProvidersEdgesExtensionsUnsupportedMediaType {
	return &GetTelephonyProvidersEdgesExtensionsUnsupportedMediaType{}
}

/*GetTelephonyProvidersEdgesExtensionsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyProvidersEdgesExtensionsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesExtensionsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/extensions][%d] getTelephonyProvidersEdgesExtensionsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgesExtensionsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesExtensionsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesExtensionsTooManyRequests creates a GetTelephonyProvidersEdgesExtensionsTooManyRequests with default headers values
func NewGetTelephonyProvidersEdgesExtensionsTooManyRequests() *GetTelephonyProvidersEdgesExtensionsTooManyRequests {
	return &GetTelephonyProvidersEdgesExtensionsTooManyRequests{}
}

/*GetTelephonyProvidersEdgesExtensionsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetTelephonyProvidersEdgesExtensionsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesExtensionsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/extensions][%d] getTelephonyProvidersEdgesExtensionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgesExtensionsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesExtensionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesExtensionsInternalServerError creates a GetTelephonyProvidersEdgesExtensionsInternalServerError with default headers values
func NewGetTelephonyProvidersEdgesExtensionsInternalServerError() *GetTelephonyProvidersEdgesExtensionsInternalServerError {
	return &GetTelephonyProvidersEdgesExtensionsInternalServerError{}
}

/*GetTelephonyProvidersEdgesExtensionsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyProvidersEdgesExtensionsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesExtensionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/extensions][%d] getTelephonyProvidersEdgesExtensionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgesExtensionsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesExtensionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesExtensionsServiceUnavailable creates a GetTelephonyProvidersEdgesExtensionsServiceUnavailable with default headers values
func NewGetTelephonyProvidersEdgesExtensionsServiceUnavailable() *GetTelephonyProvidersEdgesExtensionsServiceUnavailable {
	return &GetTelephonyProvidersEdgesExtensionsServiceUnavailable{}
}

/*GetTelephonyProvidersEdgesExtensionsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyProvidersEdgesExtensionsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesExtensionsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/extensions][%d] getTelephonyProvidersEdgesExtensionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgesExtensionsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesExtensionsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesExtensionsGatewayTimeout creates a GetTelephonyProvidersEdgesExtensionsGatewayTimeout with default headers values
func NewGetTelephonyProvidersEdgesExtensionsGatewayTimeout() *GetTelephonyProvidersEdgesExtensionsGatewayTimeout {
	return &GetTelephonyProvidersEdgesExtensionsGatewayTimeout{}
}

/*GetTelephonyProvidersEdgesExtensionsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetTelephonyProvidersEdgesExtensionsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesExtensionsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/extensions][%d] getTelephonyProvidersEdgesExtensionsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgesExtensionsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesExtensionsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

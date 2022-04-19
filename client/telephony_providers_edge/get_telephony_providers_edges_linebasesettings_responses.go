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

// GetTelephonyProvidersEdgesLinebasesettingsReader is a Reader for the GetTelephonyProvidersEdgesLinebasesettings structure.
type GetTelephonyProvidersEdgesLinebasesettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonyProvidersEdgesLinebasesettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonyProvidersEdgesLinebasesettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonyProvidersEdgesLinebasesettingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonyProvidersEdgesLinebasesettingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonyProvidersEdgesLinebasesettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonyProvidersEdgesLinebasesettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetTelephonyProvidersEdgesLinebasesettingsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonyProvidersEdgesLinebasesettingsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonyProvidersEdgesLinebasesettingsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonyProvidersEdgesLinebasesettingsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonyProvidersEdgesLinebasesettingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonyProvidersEdgesLinebasesettingsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonyProvidersEdgesLinebasesettingsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonyProvidersEdgesLinebasesettingsOK creates a GetTelephonyProvidersEdgesLinebasesettingsOK with default headers values
func NewGetTelephonyProvidersEdgesLinebasesettingsOK() *GetTelephonyProvidersEdgesLinebasesettingsOK {
	return &GetTelephonyProvidersEdgesLinebasesettingsOK{}
}

/*GetTelephonyProvidersEdgesLinebasesettingsOK handles this case with default header values.

successful operation
*/
type GetTelephonyProvidersEdgesLinebasesettingsOK struct {
	Payload *models.LineBaseEntityListing
}

func (o *GetTelephonyProvidersEdgesLinebasesettingsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/linebasesettings][%d] getTelephonyProvidersEdgesLinebasesettingsOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgesLinebasesettingsOK) GetPayload() *models.LineBaseEntityListing {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesLinebasesettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LineBaseEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesLinebasesettingsBadRequest creates a GetTelephonyProvidersEdgesLinebasesettingsBadRequest with default headers values
func NewGetTelephonyProvidersEdgesLinebasesettingsBadRequest() *GetTelephonyProvidersEdgesLinebasesettingsBadRequest {
	return &GetTelephonyProvidersEdgesLinebasesettingsBadRequest{}
}

/*GetTelephonyProvidersEdgesLinebasesettingsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyProvidersEdgesLinebasesettingsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesLinebasesettingsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/linebasesettings][%d] getTelephonyProvidersEdgesLinebasesettingsBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgesLinebasesettingsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesLinebasesettingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesLinebasesettingsUnauthorized creates a GetTelephonyProvidersEdgesLinebasesettingsUnauthorized with default headers values
func NewGetTelephonyProvidersEdgesLinebasesettingsUnauthorized() *GetTelephonyProvidersEdgesLinebasesettingsUnauthorized {
	return &GetTelephonyProvidersEdgesLinebasesettingsUnauthorized{}
}

/*GetTelephonyProvidersEdgesLinebasesettingsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyProvidersEdgesLinebasesettingsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesLinebasesettingsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/linebasesettings][%d] getTelephonyProvidersEdgesLinebasesettingsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgesLinebasesettingsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesLinebasesettingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesLinebasesettingsForbidden creates a GetTelephonyProvidersEdgesLinebasesettingsForbidden with default headers values
func NewGetTelephonyProvidersEdgesLinebasesettingsForbidden() *GetTelephonyProvidersEdgesLinebasesettingsForbidden {
	return &GetTelephonyProvidersEdgesLinebasesettingsForbidden{}
}

/*GetTelephonyProvidersEdgesLinebasesettingsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyProvidersEdgesLinebasesettingsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesLinebasesettingsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/linebasesettings][%d] getTelephonyProvidersEdgesLinebasesettingsForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgesLinebasesettingsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesLinebasesettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesLinebasesettingsNotFound creates a GetTelephonyProvidersEdgesLinebasesettingsNotFound with default headers values
func NewGetTelephonyProvidersEdgesLinebasesettingsNotFound() *GetTelephonyProvidersEdgesLinebasesettingsNotFound {
	return &GetTelephonyProvidersEdgesLinebasesettingsNotFound{}
}

/*GetTelephonyProvidersEdgesLinebasesettingsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetTelephonyProvidersEdgesLinebasesettingsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesLinebasesettingsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/linebasesettings][%d] getTelephonyProvidersEdgesLinebasesettingsNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgesLinebasesettingsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesLinebasesettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesLinebasesettingsRequestTimeout creates a GetTelephonyProvidersEdgesLinebasesettingsRequestTimeout with default headers values
func NewGetTelephonyProvidersEdgesLinebasesettingsRequestTimeout() *GetTelephonyProvidersEdgesLinebasesettingsRequestTimeout {
	return &GetTelephonyProvidersEdgesLinebasesettingsRequestTimeout{}
}

/*GetTelephonyProvidersEdgesLinebasesettingsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetTelephonyProvidersEdgesLinebasesettingsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesLinebasesettingsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/linebasesettings][%d] getTelephonyProvidersEdgesLinebasesettingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTelephonyProvidersEdgesLinebasesettingsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesLinebasesettingsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesLinebasesettingsRequestEntityTooLarge creates a GetTelephonyProvidersEdgesLinebasesettingsRequestEntityTooLarge with default headers values
func NewGetTelephonyProvidersEdgesLinebasesettingsRequestEntityTooLarge() *GetTelephonyProvidersEdgesLinebasesettingsRequestEntityTooLarge {
	return &GetTelephonyProvidersEdgesLinebasesettingsRequestEntityTooLarge{}
}

/*GetTelephonyProvidersEdgesLinebasesettingsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetTelephonyProvidersEdgesLinebasesettingsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesLinebasesettingsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/linebasesettings][%d] getTelephonyProvidersEdgesLinebasesettingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgesLinebasesettingsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesLinebasesettingsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesLinebasesettingsUnsupportedMediaType creates a GetTelephonyProvidersEdgesLinebasesettingsUnsupportedMediaType with default headers values
func NewGetTelephonyProvidersEdgesLinebasesettingsUnsupportedMediaType() *GetTelephonyProvidersEdgesLinebasesettingsUnsupportedMediaType {
	return &GetTelephonyProvidersEdgesLinebasesettingsUnsupportedMediaType{}
}

/*GetTelephonyProvidersEdgesLinebasesettingsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyProvidersEdgesLinebasesettingsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesLinebasesettingsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/linebasesettings][%d] getTelephonyProvidersEdgesLinebasesettingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgesLinebasesettingsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesLinebasesettingsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesLinebasesettingsTooManyRequests creates a GetTelephonyProvidersEdgesLinebasesettingsTooManyRequests with default headers values
func NewGetTelephonyProvidersEdgesLinebasesettingsTooManyRequests() *GetTelephonyProvidersEdgesLinebasesettingsTooManyRequests {
	return &GetTelephonyProvidersEdgesLinebasesettingsTooManyRequests{}
}

/*GetTelephonyProvidersEdgesLinebasesettingsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetTelephonyProvidersEdgesLinebasesettingsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesLinebasesettingsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/linebasesettings][%d] getTelephonyProvidersEdgesLinebasesettingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgesLinebasesettingsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesLinebasesettingsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesLinebasesettingsInternalServerError creates a GetTelephonyProvidersEdgesLinebasesettingsInternalServerError with default headers values
func NewGetTelephonyProvidersEdgesLinebasesettingsInternalServerError() *GetTelephonyProvidersEdgesLinebasesettingsInternalServerError {
	return &GetTelephonyProvidersEdgesLinebasesettingsInternalServerError{}
}

/*GetTelephonyProvidersEdgesLinebasesettingsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyProvidersEdgesLinebasesettingsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesLinebasesettingsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/linebasesettings][%d] getTelephonyProvidersEdgesLinebasesettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgesLinebasesettingsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesLinebasesettingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesLinebasesettingsServiceUnavailable creates a GetTelephonyProvidersEdgesLinebasesettingsServiceUnavailable with default headers values
func NewGetTelephonyProvidersEdgesLinebasesettingsServiceUnavailable() *GetTelephonyProvidersEdgesLinebasesettingsServiceUnavailable {
	return &GetTelephonyProvidersEdgesLinebasesettingsServiceUnavailable{}
}

/*GetTelephonyProvidersEdgesLinebasesettingsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyProvidersEdgesLinebasesettingsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesLinebasesettingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/linebasesettings][%d] getTelephonyProvidersEdgesLinebasesettingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgesLinebasesettingsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesLinebasesettingsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesLinebasesettingsGatewayTimeout creates a GetTelephonyProvidersEdgesLinebasesettingsGatewayTimeout with default headers values
func NewGetTelephonyProvidersEdgesLinebasesettingsGatewayTimeout() *GetTelephonyProvidersEdgesLinebasesettingsGatewayTimeout {
	return &GetTelephonyProvidersEdgesLinebasesettingsGatewayTimeout{}
}

/*GetTelephonyProvidersEdgesLinebasesettingsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetTelephonyProvidersEdgesLinebasesettingsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesLinebasesettingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/linebasesettings][%d] getTelephonyProvidersEdgesLinebasesettingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgesLinebasesettingsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesLinebasesettingsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// GetTelephonyProvidersEdgesPhonebasesettingsReader is a Reader for the GetTelephonyProvidersEdgesPhonebasesettings structure.
type GetTelephonyProvidersEdgesPhonebasesettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonyProvidersEdgesPhonebasesettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonyProvidersEdgesPhonebasesettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonyProvidersEdgesPhonebasesettingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonyProvidersEdgesPhonebasesettingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonyProvidersEdgesPhonebasesettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonyProvidersEdgesPhonebasesettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetTelephonyProvidersEdgesPhonebasesettingsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonyProvidersEdgesPhonebasesettingsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonyProvidersEdgesPhonebasesettingsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonyProvidersEdgesPhonebasesettingsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonyProvidersEdgesPhonebasesettingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonyProvidersEdgesPhonebasesettingsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonyProvidersEdgesPhonebasesettingsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonyProvidersEdgesPhonebasesettingsOK creates a GetTelephonyProvidersEdgesPhonebasesettingsOK with default headers values
func NewGetTelephonyProvidersEdgesPhonebasesettingsOK() *GetTelephonyProvidersEdgesPhonebasesettingsOK {
	return &GetTelephonyProvidersEdgesPhonebasesettingsOK{}
}

/*GetTelephonyProvidersEdgesPhonebasesettingsOK handles this case with default header values.

successful operation
*/
type GetTelephonyProvidersEdgesPhonebasesettingsOK struct {
	Payload *models.PhoneBaseEntityListing
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/phonebasesettings][%d] getTelephonyProvidersEdgesPhonebasesettingsOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsOK) GetPayload() *models.PhoneBaseEntityListing {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PhoneBaseEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhonebasesettingsBadRequest creates a GetTelephonyProvidersEdgesPhonebasesettingsBadRequest with default headers values
func NewGetTelephonyProvidersEdgesPhonebasesettingsBadRequest() *GetTelephonyProvidersEdgesPhonebasesettingsBadRequest {
	return &GetTelephonyProvidersEdgesPhonebasesettingsBadRequest{}
}

/*GetTelephonyProvidersEdgesPhonebasesettingsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyProvidersEdgesPhonebasesettingsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/phonebasesettings][%d] getTelephonyProvidersEdgesPhonebasesettingsBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhonebasesettingsUnauthorized creates a GetTelephonyProvidersEdgesPhonebasesettingsUnauthorized with default headers values
func NewGetTelephonyProvidersEdgesPhonebasesettingsUnauthorized() *GetTelephonyProvidersEdgesPhonebasesettingsUnauthorized {
	return &GetTelephonyProvidersEdgesPhonebasesettingsUnauthorized{}
}

/*GetTelephonyProvidersEdgesPhonebasesettingsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyProvidersEdgesPhonebasesettingsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/phonebasesettings][%d] getTelephonyProvidersEdgesPhonebasesettingsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhonebasesettingsForbidden creates a GetTelephonyProvidersEdgesPhonebasesettingsForbidden with default headers values
func NewGetTelephonyProvidersEdgesPhonebasesettingsForbidden() *GetTelephonyProvidersEdgesPhonebasesettingsForbidden {
	return &GetTelephonyProvidersEdgesPhonebasesettingsForbidden{}
}

/*GetTelephonyProvidersEdgesPhonebasesettingsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyProvidersEdgesPhonebasesettingsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/phonebasesettings][%d] getTelephonyProvidersEdgesPhonebasesettingsForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhonebasesettingsNotFound creates a GetTelephonyProvidersEdgesPhonebasesettingsNotFound with default headers values
func NewGetTelephonyProvidersEdgesPhonebasesettingsNotFound() *GetTelephonyProvidersEdgesPhonebasesettingsNotFound {
	return &GetTelephonyProvidersEdgesPhonebasesettingsNotFound{}
}

/*GetTelephonyProvidersEdgesPhonebasesettingsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetTelephonyProvidersEdgesPhonebasesettingsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/phonebasesettings][%d] getTelephonyProvidersEdgesPhonebasesettingsNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhonebasesettingsRequestTimeout creates a GetTelephonyProvidersEdgesPhonebasesettingsRequestTimeout with default headers values
func NewGetTelephonyProvidersEdgesPhonebasesettingsRequestTimeout() *GetTelephonyProvidersEdgesPhonebasesettingsRequestTimeout {
	return &GetTelephonyProvidersEdgesPhonebasesettingsRequestTimeout{}
}

/*GetTelephonyProvidersEdgesPhonebasesettingsRequestTimeout handles this case with default header values.

Request Timeout
*/
type GetTelephonyProvidersEdgesPhonebasesettingsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/phonebasesettings][%d] getTelephonyProvidersEdgesPhonebasesettingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhonebasesettingsRequestEntityTooLarge creates a GetTelephonyProvidersEdgesPhonebasesettingsRequestEntityTooLarge with default headers values
func NewGetTelephonyProvidersEdgesPhonebasesettingsRequestEntityTooLarge() *GetTelephonyProvidersEdgesPhonebasesettingsRequestEntityTooLarge {
	return &GetTelephonyProvidersEdgesPhonebasesettingsRequestEntityTooLarge{}
}

/*GetTelephonyProvidersEdgesPhonebasesettingsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetTelephonyProvidersEdgesPhonebasesettingsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/phonebasesettings][%d] getTelephonyProvidersEdgesPhonebasesettingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhonebasesettingsUnsupportedMediaType creates a GetTelephonyProvidersEdgesPhonebasesettingsUnsupportedMediaType with default headers values
func NewGetTelephonyProvidersEdgesPhonebasesettingsUnsupportedMediaType() *GetTelephonyProvidersEdgesPhonebasesettingsUnsupportedMediaType {
	return &GetTelephonyProvidersEdgesPhonebasesettingsUnsupportedMediaType{}
}

/*GetTelephonyProvidersEdgesPhonebasesettingsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyProvidersEdgesPhonebasesettingsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/phonebasesettings][%d] getTelephonyProvidersEdgesPhonebasesettingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhonebasesettingsTooManyRequests creates a GetTelephonyProvidersEdgesPhonebasesettingsTooManyRequests with default headers values
func NewGetTelephonyProvidersEdgesPhonebasesettingsTooManyRequests() *GetTelephonyProvidersEdgesPhonebasesettingsTooManyRequests {
	return &GetTelephonyProvidersEdgesPhonebasesettingsTooManyRequests{}
}

/*GetTelephonyProvidersEdgesPhonebasesettingsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetTelephonyProvidersEdgesPhonebasesettingsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/phonebasesettings][%d] getTelephonyProvidersEdgesPhonebasesettingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhonebasesettingsInternalServerError creates a GetTelephonyProvidersEdgesPhonebasesettingsInternalServerError with default headers values
func NewGetTelephonyProvidersEdgesPhonebasesettingsInternalServerError() *GetTelephonyProvidersEdgesPhonebasesettingsInternalServerError {
	return &GetTelephonyProvidersEdgesPhonebasesettingsInternalServerError{}
}

/*GetTelephonyProvidersEdgesPhonebasesettingsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyProvidersEdgesPhonebasesettingsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/phonebasesettings][%d] getTelephonyProvidersEdgesPhonebasesettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhonebasesettingsServiceUnavailable creates a GetTelephonyProvidersEdgesPhonebasesettingsServiceUnavailable with default headers values
func NewGetTelephonyProvidersEdgesPhonebasesettingsServiceUnavailable() *GetTelephonyProvidersEdgesPhonebasesettingsServiceUnavailable {
	return &GetTelephonyProvidersEdgesPhonebasesettingsServiceUnavailable{}
}

/*GetTelephonyProvidersEdgesPhonebasesettingsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyProvidersEdgesPhonebasesettingsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/phonebasesettings][%d] getTelephonyProvidersEdgesPhonebasesettingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhonebasesettingsGatewayTimeout creates a GetTelephonyProvidersEdgesPhonebasesettingsGatewayTimeout with default headers values
func NewGetTelephonyProvidersEdgesPhonebasesettingsGatewayTimeout() *GetTelephonyProvidersEdgesPhonebasesettingsGatewayTimeout {
	return &GetTelephonyProvidersEdgesPhonebasesettingsGatewayTimeout{}
}

/*GetTelephonyProvidersEdgesPhonebasesettingsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetTelephonyProvidersEdgesPhonebasesettingsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/phonebasesettings][%d] getTelephonyProvidersEdgesPhonebasesettingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
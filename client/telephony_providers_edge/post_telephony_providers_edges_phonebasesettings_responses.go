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

// PostTelephonyProvidersEdgesPhonebasesettingsReader is a Reader for the PostTelephonyProvidersEdgesPhonebasesettings structure.
type PostTelephonyProvidersEdgesPhonebasesettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTelephonyProvidersEdgesPhonebasesettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostTelephonyProvidersEdgesPhonebasesettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostTelephonyProvidersEdgesPhonebasesettingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostTelephonyProvidersEdgesPhonebasesettingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostTelephonyProvidersEdgesPhonebasesettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostTelephonyProvidersEdgesPhonebasesettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostTelephonyProvidersEdgesPhonebasesettingsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostTelephonyProvidersEdgesPhonebasesettingsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostTelephonyProvidersEdgesPhonebasesettingsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostTelephonyProvidersEdgesPhonebasesettingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostTelephonyProvidersEdgesPhonebasesettingsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostTelephonyProvidersEdgesPhonebasesettingsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostTelephonyProvidersEdgesPhonebasesettingsOK creates a PostTelephonyProvidersEdgesPhonebasesettingsOK with default headers values
func NewPostTelephonyProvidersEdgesPhonebasesettingsOK() *PostTelephonyProvidersEdgesPhonebasesettingsOK {
	return &PostTelephonyProvidersEdgesPhonebasesettingsOK{}
}

/*PostTelephonyProvidersEdgesPhonebasesettingsOK handles this case with default header values.

successful operation
*/
type PostTelephonyProvidersEdgesPhonebasesettingsOK struct {
	Payload *models.PhoneBase
}

func (o *PostTelephonyProvidersEdgesPhonebasesettingsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/phonebasesettings][%d] postTelephonyProvidersEdgesPhonebasesettingsOK  %+v", 200, o.Payload)
}

func (o *PostTelephonyProvidersEdgesPhonebasesettingsOK) GetPayload() *models.PhoneBase {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesPhonebasesettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PhoneBase)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesPhonebasesettingsBadRequest creates a PostTelephonyProvidersEdgesPhonebasesettingsBadRequest with default headers values
func NewPostTelephonyProvidersEdgesPhonebasesettingsBadRequest() *PostTelephonyProvidersEdgesPhonebasesettingsBadRequest {
	return &PostTelephonyProvidersEdgesPhonebasesettingsBadRequest{}
}

/*PostTelephonyProvidersEdgesPhonebasesettingsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostTelephonyProvidersEdgesPhonebasesettingsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesPhonebasesettingsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/phonebasesettings][%d] postTelephonyProvidersEdgesPhonebasesettingsBadRequest  %+v", 400, o.Payload)
}

func (o *PostTelephonyProvidersEdgesPhonebasesettingsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesPhonebasesettingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesPhonebasesettingsUnauthorized creates a PostTelephonyProvidersEdgesPhonebasesettingsUnauthorized with default headers values
func NewPostTelephonyProvidersEdgesPhonebasesettingsUnauthorized() *PostTelephonyProvidersEdgesPhonebasesettingsUnauthorized {
	return &PostTelephonyProvidersEdgesPhonebasesettingsUnauthorized{}
}

/*PostTelephonyProvidersEdgesPhonebasesettingsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostTelephonyProvidersEdgesPhonebasesettingsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesPhonebasesettingsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/phonebasesettings][%d] postTelephonyProvidersEdgesPhonebasesettingsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostTelephonyProvidersEdgesPhonebasesettingsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesPhonebasesettingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesPhonebasesettingsForbidden creates a PostTelephonyProvidersEdgesPhonebasesettingsForbidden with default headers values
func NewPostTelephonyProvidersEdgesPhonebasesettingsForbidden() *PostTelephonyProvidersEdgesPhonebasesettingsForbidden {
	return &PostTelephonyProvidersEdgesPhonebasesettingsForbidden{}
}

/*PostTelephonyProvidersEdgesPhonebasesettingsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostTelephonyProvidersEdgesPhonebasesettingsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesPhonebasesettingsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/phonebasesettings][%d] postTelephonyProvidersEdgesPhonebasesettingsForbidden  %+v", 403, o.Payload)
}

func (o *PostTelephonyProvidersEdgesPhonebasesettingsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesPhonebasesettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesPhonebasesettingsNotFound creates a PostTelephonyProvidersEdgesPhonebasesettingsNotFound with default headers values
func NewPostTelephonyProvidersEdgesPhonebasesettingsNotFound() *PostTelephonyProvidersEdgesPhonebasesettingsNotFound {
	return &PostTelephonyProvidersEdgesPhonebasesettingsNotFound{}
}

/*PostTelephonyProvidersEdgesPhonebasesettingsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostTelephonyProvidersEdgesPhonebasesettingsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesPhonebasesettingsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/phonebasesettings][%d] postTelephonyProvidersEdgesPhonebasesettingsNotFound  %+v", 404, o.Payload)
}

func (o *PostTelephonyProvidersEdgesPhonebasesettingsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesPhonebasesettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesPhonebasesettingsRequestEntityTooLarge creates a PostTelephonyProvidersEdgesPhonebasesettingsRequestEntityTooLarge with default headers values
func NewPostTelephonyProvidersEdgesPhonebasesettingsRequestEntityTooLarge() *PostTelephonyProvidersEdgesPhonebasesettingsRequestEntityTooLarge {
	return &PostTelephonyProvidersEdgesPhonebasesettingsRequestEntityTooLarge{}
}

/*PostTelephonyProvidersEdgesPhonebasesettingsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostTelephonyProvidersEdgesPhonebasesettingsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesPhonebasesettingsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/phonebasesettings][%d] postTelephonyProvidersEdgesPhonebasesettingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostTelephonyProvidersEdgesPhonebasesettingsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesPhonebasesettingsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesPhonebasesettingsUnsupportedMediaType creates a PostTelephonyProvidersEdgesPhonebasesettingsUnsupportedMediaType with default headers values
func NewPostTelephonyProvidersEdgesPhonebasesettingsUnsupportedMediaType() *PostTelephonyProvidersEdgesPhonebasesettingsUnsupportedMediaType {
	return &PostTelephonyProvidersEdgesPhonebasesettingsUnsupportedMediaType{}
}

/*PostTelephonyProvidersEdgesPhonebasesettingsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostTelephonyProvidersEdgesPhonebasesettingsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesPhonebasesettingsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/phonebasesettings][%d] postTelephonyProvidersEdgesPhonebasesettingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostTelephonyProvidersEdgesPhonebasesettingsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesPhonebasesettingsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesPhonebasesettingsTooManyRequests creates a PostTelephonyProvidersEdgesPhonebasesettingsTooManyRequests with default headers values
func NewPostTelephonyProvidersEdgesPhonebasesettingsTooManyRequests() *PostTelephonyProvidersEdgesPhonebasesettingsTooManyRequests {
	return &PostTelephonyProvidersEdgesPhonebasesettingsTooManyRequests{}
}

/*PostTelephonyProvidersEdgesPhonebasesettingsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostTelephonyProvidersEdgesPhonebasesettingsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesPhonebasesettingsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/phonebasesettings][%d] postTelephonyProvidersEdgesPhonebasesettingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostTelephonyProvidersEdgesPhonebasesettingsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesPhonebasesettingsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesPhonebasesettingsInternalServerError creates a PostTelephonyProvidersEdgesPhonebasesettingsInternalServerError with default headers values
func NewPostTelephonyProvidersEdgesPhonebasesettingsInternalServerError() *PostTelephonyProvidersEdgesPhonebasesettingsInternalServerError {
	return &PostTelephonyProvidersEdgesPhonebasesettingsInternalServerError{}
}

/*PostTelephonyProvidersEdgesPhonebasesettingsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostTelephonyProvidersEdgesPhonebasesettingsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesPhonebasesettingsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/phonebasesettings][%d] postTelephonyProvidersEdgesPhonebasesettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostTelephonyProvidersEdgesPhonebasesettingsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesPhonebasesettingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesPhonebasesettingsServiceUnavailable creates a PostTelephonyProvidersEdgesPhonebasesettingsServiceUnavailable with default headers values
func NewPostTelephonyProvidersEdgesPhonebasesettingsServiceUnavailable() *PostTelephonyProvidersEdgesPhonebasesettingsServiceUnavailable {
	return &PostTelephonyProvidersEdgesPhonebasesettingsServiceUnavailable{}
}

/*PostTelephonyProvidersEdgesPhonebasesettingsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostTelephonyProvidersEdgesPhonebasesettingsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesPhonebasesettingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/phonebasesettings][%d] postTelephonyProvidersEdgesPhonebasesettingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostTelephonyProvidersEdgesPhonebasesettingsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesPhonebasesettingsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesPhonebasesettingsGatewayTimeout creates a PostTelephonyProvidersEdgesPhonebasesettingsGatewayTimeout with default headers values
func NewPostTelephonyProvidersEdgesPhonebasesettingsGatewayTimeout() *PostTelephonyProvidersEdgesPhonebasesettingsGatewayTimeout {
	return &PostTelephonyProvidersEdgesPhonebasesettingsGatewayTimeout{}
}

/*PostTelephonyProvidersEdgesPhonebasesettingsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostTelephonyProvidersEdgesPhonebasesettingsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesPhonebasesettingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/phonebasesettings][%d] postTelephonyProvidersEdgesPhonebasesettingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostTelephonyProvidersEdgesPhonebasesettingsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesPhonebasesettingsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

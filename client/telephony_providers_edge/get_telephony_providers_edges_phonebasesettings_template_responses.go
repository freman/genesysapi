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

// GetTelephonyProvidersEdgesPhonebasesettingsTemplateReader is a Reader for the GetTelephonyProvidersEdgesPhonebasesettingsTemplate structure.
type GetTelephonyProvidersEdgesPhonebasesettingsTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonyProvidersEdgesPhonebasesettingsTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonyProvidersEdgesPhonebasesettingsTemplateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonyProvidersEdgesPhonebasesettingsTemplateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonyProvidersEdgesPhonebasesettingsTemplateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonyProvidersEdgesPhonebasesettingsTemplateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonyProvidersEdgesPhonebasesettingsTemplateRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonyProvidersEdgesPhonebasesettingsTemplateUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonyProvidersEdgesPhonebasesettingsTemplateTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonyProvidersEdgesPhonebasesettingsTemplateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonyProvidersEdgesPhonebasesettingsTemplateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonyProvidersEdgesPhonebasesettingsTemplateGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonyProvidersEdgesPhonebasesettingsTemplateOK creates a GetTelephonyProvidersEdgesPhonebasesettingsTemplateOK with default headers values
func NewGetTelephonyProvidersEdgesPhonebasesettingsTemplateOK() *GetTelephonyProvidersEdgesPhonebasesettingsTemplateOK {
	return &GetTelephonyProvidersEdgesPhonebasesettingsTemplateOK{}
}

/*GetTelephonyProvidersEdgesPhonebasesettingsTemplateOK handles this case with default header values.

successful operation
*/
type GetTelephonyProvidersEdgesPhonebasesettingsTemplateOK struct {
	Payload *models.PhoneBase
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/phonebasesettings/template][%d] getTelephonyProvidersEdgesPhonebasesettingsTemplateOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateOK) GetPayload() *models.PhoneBase {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PhoneBase)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhonebasesettingsTemplateBadRequest creates a GetTelephonyProvidersEdgesPhonebasesettingsTemplateBadRequest with default headers values
func NewGetTelephonyProvidersEdgesPhonebasesettingsTemplateBadRequest() *GetTelephonyProvidersEdgesPhonebasesettingsTemplateBadRequest {
	return &GetTelephonyProvidersEdgesPhonebasesettingsTemplateBadRequest{}
}

/*GetTelephonyProvidersEdgesPhonebasesettingsTemplateBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyProvidersEdgesPhonebasesettingsTemplateBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/phonebasesettings/template][%d] getTelephonyProvidersEdgesPhonebasesettingsTemplateBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhonebasesettingsTemplateUnauthorized creates a GetTelephonyProvidersEdgesPhonebasesettingsTemplateUnauthorized with default headers values
func NewGetTelephonyProvidersEdgesPhonebasesettingsTemplateUnauthorized() *GetTelephonyProvidersEdgesPhonebasesettingsTemplateUnauthorized {
	return &GetTelephonyProvidersEdgesPhonebasesettingsTemplateUnauthorized{}
}

/*GetTelephonyProvidersEdgesPhonebasesettingsTemplateUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyProvidersEdgesPhonebasesettingsTemplateUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/phonebasesettings/template][%d] getTelephonyProvidersEdgesPhonebasesettingsTemplateUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhonebasesettingsTemplateForbidden creates a GetTelephonyProvidersEdgesPhonebasesettingsTemplateForbidden with default headers values
func NewGetTelephonyProvidersEdgesPhonebasesettingsTemplateForbidden() *GetTelephonyProvidersEdgesPhonebasesettingsTemplateForbidden {
	return &GetTelephonyProvidersEdgesPhonebasesettingsTemplateForbidden{}
}

/*GetTelephonyProvidersEdgesPhonebasesettingsTemplateForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyProvidersEdgesPhonebasesettingsTemplateForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/phonebasesettings/template][%d] getTelephonyProvidersEdgesPhonebasesettingsTemplateForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhonebasesettingsTemplateNotFound creates a GetTelephonyProvidersEdgesPhonebasesettingsTemplateNotFound with default headers values
func NewGetTelephonyProvidersEdgesPhonebasesettingsTemplateNotFound() *GetTelephonyProvidersEdgesPhonebasesettingsTemplateNotFound {
	return &GetTelephonyProvidersEdgesPhonebasesettingsTemplateNotFound{}
}

/*GetTelephonyProvidersEdgesPhonebasesettingsTemplateNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetTelephonyProvidersEdgesPhonebasesettingsTemplateNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/phonebasesettings/template][%d] getTelephonyProvidersEdgesPhonebasesettingsTemplateNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhonebasesettingsTemplateRequestEntityTooLarge creates a GetTelephonyProvidersEdgesPhonebasesettingsTemplateRequestEntityTooLarge with default headers values
func NewGetTelephonyProvidersEdgesPhonebasesettingsTemplateRequestEntityTooLarge() *GetTelephonyProvidersEdgesPhonebasesettingsTemplateRequestEntityTooLarge {
	return &GetTelephonyProvidersEdgesPhonebasesettingsTemplateRequestEntityTooLarge{}
}

/*GetTelephonyProvidersEdgesPhonebasesettingsTemplateRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetTelephonyProvidersEdgesPhonebasesettingsTemplateRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/phonebasesettings/template][%d] getTelephonyProvidersEdgesPhonebasesettingsTemplateRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhonebasesettingsTemplateUnsupportedMediaType creates a GetTelephonyProvidersEdgesPhonebasesettingsTemplateUnsupportedMediaType with default headers values
func NewGetTelephonyProvidersEdgesPhonebasesettingsTemplateUnsupportedMediaType() *GetTelephonyProvidersEdgesPhonebasesettingsTemplateUnsupportedMediaType {
	return &GetTelephonyProvidersEdgesPhonebasesettingsTemplateUnsupportedMediaType{}
}

/*GetTelephonyProvidersEdgesPhonebasesettingsTemplateUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyProvidersEdgesPhonebasesettingsTemplateUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/phonebasesettings/template][%d] getTelephonyProvidersEdgesPhonebasesettingsTemplateUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhonebasesettingsTemplateTooManyRequests creates a GetTelephonyProvidersEdgesPhonebasesettingsTemplateTooManyRequests with default headers values
func NewGetTelephonyProvidersEdgesPhonebasesettingsTemplateTooManyRequests() *GetTelephonyProvidersEdgesPhonebasesettingsTemplateTooManyRequests {
	return &GetTelephonyProvidersEdgesPhonebasesettingsTemplateTooManyRequests{}
}

/*GetTelephonyProvidersEdgesPhonebasesettingsTemplateTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetTelephonyProvidersEdgesPhonebasesettingsTemplateTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/phonebasesettings/template][%d] getTelephonyProvidersEdgesPhonebasesettingsTemplateTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhonebasesettingsTemplateInternalServerError creates a GetTelephonyProvidersEdgesPhonebasesettingsTemplateInternalServerError with default headers values
func NewGetTelephonyProvidersEdgesPhonebasesettingsTemplateInternalServerError() *GetTelephonyProvidersEdgesPhonebasesettingsTemplateInternalServerError {
	return &GetTelephonyProvidersEdgesPhonebasesettingsTemplateInternalServerError{}
}

/*GetTelephonyProvidersEdgesPhonebasesettingsTemplateInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyProvidersEdgesPhonebasesettingsTemplateInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/phonebasesettings/template][%d] getTelephonyProvidersEdgesPhonebasesettingsTemplateInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhonebasesettingsTemplateServiceUnavailable creates a GetTelephonyProvidersEdgesPhonebasesettingsTemplateServiceUnavailable with default headers values
func NewGetTelephonyProvidersEdgesPhonebasesettingsTemplateServiceUnavailable() *GetTelephonyProvidersEdgesPhonebasesettingsTemplateServiceUnavailable {
	return &GetTelephonyProvidersEdgesPhonebasesettingsTemplateServiceUnavailable{}
}

/*GetTelephonyProvidersEdgesPhonebasesettingsTemplateServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyProvidersEdgesPhonebasesettingsTemplateServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/phonebasesettings/template][%d] getTelephonyProvidersEdgesPhonebasesettingsTemplateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhonebasesettingsTemplateGatewayTimeout creates a GetTelephonyProvidersEdgesPhonebasesettingsTemplateGatewayTimeout with default headers values
func NewGetTelephonyProvidersEdgesPhonebasesettingsTemplateGatewayTimeout() *GetTelephonyProvidersEdgesPhonebasesettingsTemplateGatewayTimeout {
	return &GetTelephonyProvidersEdgesPhonebasesettingsTemplateGatewayTimeout{}
}

/*GetTelephonyProvidersEdgesPhonebasesettingsTemplateGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetTelephonyProvidersEdgesPhonebasesettingsTemplateGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/phonebasesettings/template][%d] getTelephonyProvidersEdgesPhonebasesettingsTemplateGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

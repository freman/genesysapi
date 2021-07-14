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

// GetTelephonyProvidersEdgesSiteReader is a Reader for the GetTelephonyProvidersEdgesSite structure.
type GetTelephonyProvidersEdgesSiteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonyProvidersEdgesSiteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonyProvidersEdgesSiteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonyProvidersEdgesSiteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonyProvidersEdgesSiteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonyProvidersEdgesSiteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonyProvidersEdgesSiteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetTelephonyProvidersEdgesSiteRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonyProvidersEdgesSiteRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonyProvidersEdgesSiteUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonyProvidersEdgesSiteTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonyProvidersEdgesSiteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonyProvidersEdgesSiteServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonyProvidersEdgesSiteGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonyProvidersEdgesSiteOK creates a GetTelephonyProvidersEdgesSiteOK with default headers values
func NewGetTelephonyProvidersEdgesSiteOK() *GetTelephonyProvidersEdgesSiteOK {
	return &GetTelephonyProvidersEdgesSiteOK{}
}

/*GetTelephonyProvidersEdgesSiteOK handles this case with default header values.

successful operation
*/
type GetTelephonyProvidersEdgesSiteOK struct {
	Payload *models.Site
}

func (o *GetTelephonyProvidersEdgesSiteOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}][%d] getTelephonyProvidersEdgesSiteOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteOK) GetPayload() *models.Site {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Site)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteBadRequest creates a GetTelephonyProvidersEdgesSiteBadRequest with default headers values
func NewGetTelephonyProvidersEdgesSiteBadRequest() *GetTelephonyProvidersEdgesSiteBadRequest {
	return &GetTelephonyProvidersEdgesSiteBadRequest{}
}

/*GetTelephonyProvidersEdgesSiteBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyProvidersEdgesSiteBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}][%d] getTelephonyProvidersEdgesSiteBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteUnauthorized creates a GetTelephonyProvidersEdgesSiteUnauthorized with default headers values
func NewGetTelephonyProvidersEdgesSiteUnauthorized() *GetTelephonyProvidersEdgesSiteUnauthorized {
	return &GetTelephonyProvidersEdgesSiteUnauthorized{}
}

/*GetTelephonyProvidersEdgesSiteUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyProvidersEdgesSiteUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}][%d] getTelephonyProvidersEdgesSiteUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteForbidden creates a GetTelephonyProvidersEdgesSiteForbidden with default headers values
func NewGetTelephonyProvidersEdgesSiteForbidden() *GetTelephonyProvidersEdgesSiteForbidden {
	return &GetTelephonyProvidersEdgesSiteForbidden{}
}

/*GetTelephonyProvidersEdgesSiteForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyProvidersEdgesSiteForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}][%d] getTelephonyProvidersEdgesSiteForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteNotFound creates a GetTelephonyProvidersEdgesSiteNotFound with default headers values
func NewGetTelephonyProvidersEdgesSiteNotFound() *GetTelephonyProvidersEdgesSiteNotFound {
	return &GetTelephonyProvidersEdgesSiteNotFound{}
}

/*GetTelephonyProvidersEdgesSiteNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetTelephonyProvidersEdgesSiteNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}][%d] getTelephonyProvidersEdgesSiteNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteRequestTimeout creates a GetTelephonyProvidersEdgesSiteRequestTimeout with default headers values
func NewGetTelephonyProvidersEdgesSiteRequestTimeout() *GetTelephonyProvidersEdgesSiteRequestTimeout {
	return &GetTelephonyProvidersEdgesSiteRequestTimeout{}
}

/*GetTelephonyProvidersEdgesSiteRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetTelephonyProvidersEdgesSiteRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}][%d] getTelephonyProvidersEdgesSiteRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteRequestEntityTooLarge creates a GetTelephonyProvidersEdgesSiteRequestEntityTooLarge with default headers values
func NewGetTelephonyProvidersEdgesSiteRequestEntityTooLarge() *GetTelephonyProvidersEdgesSiteRequestEntityTooLarge {
	return &GetTelephonyProvidersEdgesSiteRequestEntityTooLarge{}
}

/*GetTelephonyProvidersEdgesSiteRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetTelephonyProvidersEdgesSiteRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}][%d] getTelephonyProvidersEdgesSiteRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteUnsupportedMediaType creates a GetTelephonyProvidersEdgesSiteUnsupportedMediaType with default headers values
func NewGetTelephonyProvidersEdgesSiteUnsupportedMediaType() *GetTelephonyProvidersEdgesSiteUnsupportedMediaType {
	return &GetTelephonyProvidersEdgesSiteUnsupportedMediaType{}
}

/*GetTelephonyProvidersEdgesSiteUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyProvidersEdgesSiteUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}][%d] getTelephonyProvidersEdgesSiteUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteTooManyRequests creates a GetTelephonyProvidersEdgesSiteTooManyRequests with default headers values
func NewGetTelephonyProvidersEdgesSiteTooManyRequests() *GetTelephonyProvidersEdgesSiteTooManyRequests {
	return &GetTelephonyProvidersEdgesSiteTooManyRequests{}
}

/*GetTelephonyProvidersEdgesSiteTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetTelephonyProvidersEdgesSiteTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}][%d] getTelephonyProvidersEdgesSiteTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteInternalServerError creates a GetTelephonyProvidersEdgesSiteInternalServerError with default headers values
func NewGetTelephonyProvidersEdgesSiteInternalServerError() *GetTelephonyProvidersEdgesSiteInternalServerError {
	return &GetTelephonyProvidersEdgesSiteInternalServerError{}
}

/*GetTelephonyProvidersEdgesSiteInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyProvidersEdgesSiteInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}][%d] getTelephonyProvidersEdgesSiteInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteServiceUnavailable creates a GetTelephonyProvidersEdgesSiteServiceUnavailable with default headers values
func NewGetTelephonyProvidersEdgesSiteServiceUnavailable() *GetTelephonyProvidersEdgesSiteServiceUnavailable {
	return &GetTelephonyProvidersEdgesSiteServiceUnavailable{}
}

/*GetTelephonyProvidersEdgesSiteServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyProvidersEdgesSiteServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}][%d] getTelephonyProvidersEdgesSiteServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteGatewayTimeout creates a GetTelephonyProvidersEdgesSiteGatewayTimeout with default headers values
func NewGetTelephonyProvidersEdgesSiteGatewayTimeout() *GetTelephonyProvidersEdgesSiteGatewayTimeout {
	return &GetTelephonyProvidersEdgesSiteGatewayTimeout{}
}

/*GetTelephonyProvidersEdgesSiteGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetTelephonyProvidersEdgesSiteGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}][%d] getTelephonyProvidersEdgesSiteGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

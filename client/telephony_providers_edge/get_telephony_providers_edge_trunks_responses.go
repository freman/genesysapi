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

// GetTelephonyProvidersEdgeTrunksReader is a Reader for the GetTelephonyProvidersEdgeTrunks structure.
type GetTelephonyProvidersEdgeTrunksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonyProvidersEdgeTrunksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonyProvidersEdgeTrunksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonyProvidersEdgeTrunksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonyProvidersEdgeTrunksUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonyProvidersEdgeTrunksForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonyProvidersEdgeTrunksNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetTelephonyProvidersEdgeTrunksRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonyProvidersEdgeTrunksRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonyProvidersEdgeTrunksUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonyProvidersEdgeTrunksTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonyProvidersEdgeTrunksInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonyProvidersEdgeTrunksServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonyProvidersEdgeTrunksGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonyProvidersEdgeTrunksOK creates a GetTelephonyProvidersEdgeTrunksOK with default headers values
func NewGetTelephonyProvidersEdgeTrunksOK() *GetTelephonyProvidersEdgeTrunksOK {
	return &GetTelephonyProvidersEdgeTrunksOK{}
}

/*GetTelephonyProvidersEdgeTrunksOK handles this case with default header values.

successful operation
*/
type GetTelephonyProvidersEdgeTrunksOK struct {
	Payload *models.TrunkEntityListing
}

func (o *GetTelephonyProvidersEdgeTrunksOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/trunks][%d] getTelephonyProvidersEdgeTrunksOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgeTrunksOK) GetPayload() *models.TrunkEntityListing {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeTrunksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TrunkEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeTrunksBadRequest creates a GetTelephonyProvidersEdgeTrunksBadRequest with default headers values
func NewGetTelephonyProvidersEdgeTrunksBadRequest() *GetTelephonyProvidersEdgeTrunksBadRequest {
	return &GetTelephonyProvidersEdgeTrunksBadRequest{}
}

/*GetTelephonyProvidersEdgeTrunksBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyProvidersEdgeTrunksBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeTrunksBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/trunks][%d] getTelephonyProvidersEdgeTrunksBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgeTrunksBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeTrunksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeTrunksUnauthorized creates a GetTelephonyProvidersEdgeTrunksUnauthorized with default headers values
func NewGetTelephonyProvidersEdgeTrunksUnauthorized() *GetTelephonyProvidersEdgeTrunksUnauthorized {
	return &GetTelephonyProvidersEdgeTrunksUnauthorized{}
}

/*GetTelephonyProvidersEdgeTrunksUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyProvidersEdgeTrunksUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeTrunksUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/trunks][%d] getTelephonyProvidersEdgeTrunksUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgeTrunksUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeTrunksUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeTrunksForbidden creates a GetTelephonyProvidersEdgeTrunksForbidden with default headers values
func NewGetTelephonyProvidersEdgeTrunksForbidden() *GetTelephonyProvidersEdgeTrunksForbidden {
	return &GetTelephonyProvidersEdgeTrunksForbidden{}
}

/*GetTelephonyProvidersEdgeTrunksForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyProvidersEdgeTrunksForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeTrunksForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/trunks][%d] getTelephonyProvidersEdgeTrunksForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgeTrunksForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeTrunksForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeTrunksNotFound creates a GetTelephonyProvidersEdgeTrunksNotFound with default headers values
func NewGetTelephonyProvidersEdgeTrunksNotFound() *GetTelephonyProvidersEdgeTrunksNotFound {
	return &GetTelephonyProvidersEdgeTrunksNotFound{}
}

/*GetTelephonyProvidersEdgeTrunksNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetTelephonyProvidersEdgeTrunksNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeTrunksNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/trunks][%d] getTelephonyProvidersEdgeTrunksNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgeTrunksNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeTrunksNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeTrunksRequestTimeout creates a GetTelephonyProvidersEdgeTrunksRequestTimeout with default headers values
func NewGetTelephonyProvidersEdgeTrunksRequestTimeout() *GetTelephonyProvidersEdgeTrunksRequestTimeout {
	return &GetTelephonyProvidersEdgeTrunksRequestTimeout{}
}

/*GetTelephonyProvidersEdgeTrunksRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetTelephonyProvidersEdgeTrunksRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeTrunksRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/trunks][%d] getTelephonyProvidersEdgeTrunksRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTelephonyProvidersEdgeTrunksRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeTrunksRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeTrunksRequestEntityTooLarge creates a GetTelephonyProvidersEdgeTrunksRequestEntityTooLarge with default headers values
func NewGetTelephonyProvidersEdgeTrunksRequestEntityTooLarge() *GetTelephonyProvidersEdgeTrunksRequestEntityTooLarge {
	return &GetTelephonyProvidersEdgeTrunksRequestEntityTooLarge{}
}

/*GetTelephonyProvidersEdgeTrunksRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetTelephonyProvidersEdgeTrunksRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeTrunksRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/trunks][%d] getTelephonyProvidersEdgeTrunksRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgeTrunksRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeTrunksRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeTrunksUnsupportedMediaType creates a GetTelephonyProvidersEdgeTrunksUnsupportedMediaType with default headers values
func NewGetTelephonyProvidersEdgeTrunksUnsupportedMediaType() *GetTelephonyProvidersEdgeTrunksUnsupportedMediaType {
	return &GetTelephonyProvidersEdgeTrunksUnsupportedMediaType{}
}

/*GetTelephonyProvidersEdgeTrunksUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyProvidersEdgeTrunksUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeTrunksUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/trunks][%d] getTelephonyProvidersEdgeTrunksUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgeTrunksUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeTrunksUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeTrunksTooManyRequests creates a GetTelephonyProvidersEdgeTrunksTooManyRequests with default headers values
func NewGetTelephonyProvidersEdgeTrunksTooManyRequests() *GetTelephonyProvidersEdgeTrunksTooManyRequests {
	return &GetTelephonyProvidersEdgeTrunksTooManyRequests{}
}

/*GetTelephonyProvidersEdgeTrunksTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetTelephonyProvidersEdgeTrunksTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeTrunksTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/trunks][%d] getTelephonyProvidersEdgeTrunksTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgeTrunksTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeTrunksTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeTrunksInternalServerError creates a GetTelephonyProvidersEdgeTrunksInternalServerError with default headers values
func NewGetTelephonyProvidersEdgeTrunksInternalServerError() *GetTelephonyProvidersEdgeTrunksInternalServerError {
	return &GetTelephonyProvidersEdgeTrunksInternalServerError{}
}

/*GetTelephonyProvidersEdgeTrunksInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyProvidersEdgeTrunksInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeTrunksInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/trunks][%d] getTelephonyProvidersEdgeTrunksInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgeTrunksInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeTrunksInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeTrunksServiceUnavailable creates a GetTelephonyProvidersEdgeTrunksServiceUnavailable with default headers values
func NewGetTelephonyProvidersEdgeTrunksServiceUnavailable() *GetTelephonyProvidersEdgeTrunksServiceUnavailable {
	return &GetTelephonyProvidersEdgeTrunksServiceUnavailable{}
}

/*GetTelephonyProvidersEdgeTrunksServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyProvidersEdgeTrunksServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeTrunksServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/trunks][%d] getTelephonyProvidersEdgeTrunksServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgeTrunksServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeTrunksServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeTrunksGatewayTimeout creates a GetTelephonyProvidersEdgeTrunksGatewayTimeout with default headers values
func NewGetTelephonyProvidersEdgeTrunksGatewayTimeout() *GetTelephonyProvidersEdgeTrunksGatewayTimeout {
	return &GetTelephonyProvidersEdgeTrunksGatewayTimeout{}
}

/*GetTelephonyProvidersEdgeTrunksGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetTelephonyProvidersEdgeTrunksGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeTrunksGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/trunks][%d] getTelephonyProvidersEdgeTrunksGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgeTrunksGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeTrunksGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

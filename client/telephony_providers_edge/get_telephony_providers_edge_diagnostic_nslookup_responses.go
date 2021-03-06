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

// GetTelephonyProvidersEdgeDiagnosticNslookupReader is a Reader for the GetTelephonyProvidersEdgeDiagnosticNslookup structure.
type GetTelephonyProvidersEdgeDiagnosticNslookupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonyProvidersEdgeDiagnosticNslookupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonyProvidersEdgeDiagnosticNslookupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonyProvidersEdgeDiagnosticNslookupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonyProvidersEdgeDiagnosticNslookupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonyProvidersEdgeDiagnosticNslookupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonyProvidersEdgeDiagnosticNslookupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonyProvidersEdgeDiagnosticNslookupRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonyProvidersEdgeDiagnosticNslookupUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonyProvidersEdgeDiagnosticNslookupTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonyProvidersEdgeDiagnosticNslookupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonyProvidersEdgeDiagnosticNslookupServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonyProvidersEdgeDiagnosticNslookupGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonyProvidersEdgeDiagnosticNslookupOK creates a GetTelephonyProvidersEdgeDiagnosticNslookupOK with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticNslookupOK() *GetTelephonyProvidersEdgeDiagnosticNslookupOK {
	return &GetTelephonyProvidersEdgeDiagnosticNslookupOK{}
}

/*GetTelephonyProvidersEdgeDiagnosticNslookupOK handles this case with default header values.

Request to get network diagnostic was successful.
*/
type GetTelephonyProvidersEdgeDiagnosticNslookupOK struct {
	Payload *models.EdgeNetworkDiagnosticResponse
}

func (o *GetTelephonyProvidersEdgeDiagnosticNslookupOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/nslookup][%d] getTelephonyProvidersEdgeDiagnosticNslookupOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticNslookupOK) GetPayload() *models.EdgeNetworkDiagnosticResponse {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticNslookupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EdgeNetworkDiagnosticResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeDiagnosticNslookupBadRequest creates a GetTelephonyProvidersEdgeDiagnosticNslookupBadRequest with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticNslookupBadRequest() *GetTelephonyProvidersEdgeDiagnosticNslookupBadRequest {
	return &GetTelephonyProvidersEdgeDiagnosticNslookupBadRequest{}
}

/*GetTelephonyProvidersEdgeDiagnosticNslookupBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyProvidersEdgeDiagnosticNslookupBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeDiagnosticNslookupBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/nslookup][%d] getTelephonyProvidersEdgeDiagnosticNslookupBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticNslookupBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticNslookupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeDiagnosticNslookupUnauthorized creates a GetTelephonyProvidersEdgeDiagnosticNslookupUnauthorized with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticNslookupUnauthorized() *GetTelephonyProvidersEdgeDiagnosticNslookupUnauthorized {
	return &GetTelephonyProvidersEdgeDiagnosticNslookupUnauthorized{}
}

/*GetTelephonyProvidersEdgeDiagnosticNslookupUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyProvidersEdgeDiagnosticNslookupUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeDiagnosticNslookupUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/nslookup][%d] getTelephonyProvidersEdgeDiagnosticNslookupUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticNslookupUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticNslookupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeDiagnosticNslookupForbidden creates a GetTelephonyProvidersEdgeDiagnosticNslookupForbidden with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticNslookupForbidden() *GetTelephonyProvidersEdgeDiagnosticNslookupForbidden {
	return &GetTelephonyProvidersEdgeDiagnosticNslookupForbidden{}
}

/*GetTelephonyProvidersEdgeDiagnosticNslookupForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyProvidersEdgeDiagnosticNslookupForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeDiagnosticNslookupForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/nslookup][%d] getTelephonyProvidersEdgeDiagnosticNslookupForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticNslookupForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticNslookupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeDiagnosticNslookupNotFound creates a GetTelephonyProvidersEdgeDiagnosticNslookupNotFound with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticNslookupNotFound() *GetTelephonyProvidersEdgeDiagnosticNslookupNotFound {
	return &GetTelephonyProvidersEdgeDiagnosticNslookupNotFound{}
}

/*GetTelephonyProvidersEdgeDiagnosticNslookupNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetTelephonyProvidersEdgeDiagnosticNslookupNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeDiagnosticNslookupNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/nslookup][%d] getTelephonyProvidersEdgeDiagnosticNslookupNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticNslookupNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticNslookupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeDiagnosticNslookupRequestEntityTooLarge creates a GetTelephonyProvidersEdgeDiagnosticNslookupRequestEntityTooLarge with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticNslookupRequestEntityTooLarge() *GetTelephonyProvidersEdgeDiagnosticNslookupRequestEntityTooLarge {
	return &GetTelephonyProvidersEdgeDiagnosticNslookupRequestEntityTooLarge{}
}

/*GetTelephonyProvidersEdgeDiagnosticNslookupRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetTelephonyProvidersEdgeDiagnosticNslookupRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeDiagnosticNslookupRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/nslookup][%d] getTelephonyProvidersEdgeDiagnosticNslookupRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticNslookupRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticNslookupRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeDiagnosticNslookupUnsupportedMediaType creates a GetTelephonyProvidersEdgeDiagnosticNslookupUnsupportedMediaType with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticNslookupUnsupportedMediaType() *GetTelephonyProvidersEdgeDiagnosticNslookupUnsupportedMediaType {
	return &GetTelephonyProvidersEdgeDiagnosticNslookupUnsupportedMediaType{}
}

/*GetTelephonyProvidersEdgeDiagnosticNslookupUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyProvidersEdgeDiagnosticNslookupUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeDiagnosticNslookupUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/nslookup][%d] getTelephonyProvidersEdgeDiagnosticNslookupUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticNslookupUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticNslookupUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeDiagnosticNslookupTooManyRequests creates a GetTelephonyProvidersEdgeDiagnosticNslookupTooManyRequests with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticNslookupTooManyRequests() *GetTelephonyProvidersEdgeDiagnosticNslookupTooManyRequests {
	return &GetTelephonyProvidersEdgeDiagnosticNslookupTooManyRequests{}
}

/*GetTelephonyProvidersEdgeDiagnosticNslookupTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetTelephonyProvidersEdgeDiagnosticNslookupTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeDiagnosticNslookupTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/nslookup][%d] getTelephonyProvidersEdgeDiagnosticNslookupTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticNslookupTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticNslookupTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeDiagnosticNslookupInternalServerError creates a GetTelephonyProvidersEdgeDiagnosticNslookupInternalServerError with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticNslookupInternalServerError() *GetTelephonyProvidersEdgeDiagnosticNslookupInternalServerError {
	return &GetTelephonyProvidersEdgeDiagnosticNslookupInternalServerError{}
}

/*GetTelephonyProvidersEdgeDiagnosticNslookupInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyProvidersEdgeDiagnosticNslookupInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeDiagnosticNslookupInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/nslookup][%d] getTelephonyProvidersEdgeDiagnosticNslookupInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticNslookupInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticNslookupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeDiagnosticNslookupServiceUnavailable creates a GetTelephonyProvidersEdgeDiagnosticNslookupServiceUnavailable with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticNslookupServiceUnavailable() *GetTelephonyProvidersEdgeDiagnosticNslookupServiceUnavailable {
	return &GetTelephonyProvidersEdgeDiagnosticNslookupServiceUnavailable{}
}

/*GetTelephonyProvidersEdgeDiagnosticNslookupServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyProvidersEdgeDiagnosticNslookupServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeDiagnosticNslookupServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/nslookup][%d] getTelephonyProvidersEdgeDiagnosticNslookupServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticNslookupServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticNslookupServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeDiagnosticNslookupGatewayTimeout creates a GetTelephonyProvidersEdgeDiagnosticNslookupGatewayTimeout with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticNslookupGatewayTimeout() *GetTelephonyProvidersEdgeDiagnosticNslookupGatewayTimeout {
	return &GetTelephonyProvidersEdgeDiagnosticNslookupGatewayTimeout{}
}

/*GetTelephonyProvidersEdgeDiagnosticNslookupGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetTelephonyProvidersEdgeDiagnosticNslookupGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeDiagnosticNslookupGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/nslookup][%d] getTelephonyProvidersEdgeDiagnosticNslookupGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticNslookupGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticNslookupGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

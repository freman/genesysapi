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

// GetTelephonyProvidersEdgeDiagnosticPingReader is a Reader for the GetTelephonyProvidersEdgeDiagnosticPing structure.
type GetTelephonyProvidersEdgeDiagnosticPingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonyProvidersEdgeDiagnosticPingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonyProvidersEdgeDiagnosticPingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonyProvidersEdgeDiagnosticPingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonyProvidersEdgeDiagnosticPingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonyProvidersEdgeDiagnosticPingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonyProvidersEdgeDiagnosticPingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetTelephonyProvidersEdgeDiagnosticPingRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonyProvidersEdgeDiagnosticPingRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonyProvidersEdgeDiagnosticPingUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonyProvidersEdgeDiagnosticPingTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonyProvidersEdgeDiagnosticPingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonyProvidersEdgeDiagnosticPingServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonyProvidersEdgeDiagnosticPingGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonyProvidersEdgeDiagnosticPingOK creates a GetTelephonyProvidersEdgeDiagnosticPingOK with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticPingOK() *GetTelephonyProvidersEdgeDiagnosticPingOK {
	return &GetTelephonyProvidersEdgeDiagnosticPingOK{}
}

/*GetTelephonyProvidersEdgeDiagnosticPingOK handles this case with default header values.

Request to get network diagnostic was successful.
*/
type GetTelephonyProvidersEdgeDiagnosticPingOK struct {
	Payload *models.EdgeNetworkDiagnosticResponse
}

func (o *GetTelephonyProvidersEdgeDiagnosticPingOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/ping][%d] getTelephonyProvidersEdgeDiagnosticPingOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticPingOK) GetPayload() *models.EdgeNetworkDiagnosticResponse {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticPingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EdgeNetworkDiagnosticResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeDiagnosticPingBadRequest creates a GetTelephonyProvidersEdgeDiagnosticPingBadRequest with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticPingBadRequest() *GetTelephonyProvidersEdgeDiagnosticPingBadRequest {
	return &GetTelephonyProvidersEdgeDiagnosticPingBadRequest{}
}

/*GetTelephonyProvidersEdgeDiagnosticPingBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyProvidersEdgeDiagnosticPingBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeDiagnosticPingBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/ping][%d] getTelephonyProvidersEdgeDiagnosticPingBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticPingBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticPingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeDiagnosticPingUnauthorized creates a GetTelephonyProvidersEdgeDiagnosticPingUnauthorized with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticPingUnauthorized() *GetTelephonyProvidersEdgeDiagnosticPingUnauthorized {
	return &GetTelephonyProvidersEdgeDiagnosticPingUnauthorized{}
}

/*GetTelephonyProvidersEdgeDiagnosticPingUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyProvidersEdgeDiagnosticPingUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeDiagnosticPingUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/ping][%d] getTelephonyProvidersEdgeDiagnosticPingUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticPingUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticPingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeDiagnosticPingForbidden creates a GetTelephonyProvidersEdgeDiagnosticPingForbidden with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticPingForbidden() *GetTelephonyProvidersEdgeDiagnosticPingForbidden {
	return &GetTelephonyProvidersEdgeDiagnosticPingForbidden{}
}

/*GetTelephonyProvidersEdgeDiagnosticPingForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyProvidersEdgeDiagnosticPingForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeDiagnosticPingForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/ping][%d] getTelephonyProvidersEdgeDiagnosticPingForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticPingForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticPingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeDiagnosticPingNotFound creates a GetTelephonyProvidersEdgeDiagnosticPingNotFound with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticPingNotFound() *GetTelephonyProvidersEdgeDiagnosticPingNotFound {
	return &GetTelephonyProvidersEdgeDiagnosticPingNotFound{}
}

/*GetTelephonyProvidersEdgeDiagnosticPingNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetTelephonyProvidersEdgeDiagnosticPingNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeDiagnosticPingNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/ping][%d] getTelephonyProvidersEdgeDiagnosticPingNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticPingNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticPingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeDiagnosticPingRequestTimeout creates a GetTelephonyProvidersEdgeDiagnosticPingRequestTimeout with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticPingRequestTimeout() *GetTelephonyProvidersEdgeDiagnosticPingRequestTimeout {
	return &GetTelephonyProvidersEdgeDiagnosticPingRequestTimeout{}
}

/*GetTelephonyProvidersEdgeDiagnosticPingRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetTelephonyProvidersEdgeDiagnosticPingRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeDiagnosticPingRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/ping][%d] getTelephonyProvidersEdgeDiagnosticPingRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticPingRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticPingRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeDiagnosticPingRequestEntityTooLarge creates a GetTelephonyProvidersEdgeDiagnosticPingRequestEntityTooLarge with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticPingRequestEntityTooLarge() *GetTelephonyProvidersEdgeDiagnosticPingRequestEntityTooLarge {
	return &GetTelephonyProvidersEdgeDiagnosticPingRequestEntityTooLarge{}
}

/*GetTelephonyProvidersEdgeDiagnosticPingRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetTelephonyProvidersEdgeDiagnosticPingRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeDiagnosticPingRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/ping][%d] getTelephonyProvidersEdgeDiagnosticPingRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticPingRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticPingRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeDiagnosticPingUnsupportedMediaType creates a GetTelephonyProvidersEdgeDiagnosticPingUnsupportedMediaType with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticPingUnsupportedMediaType() *GetTelephonyProvidersEdgeDiagnosticPingUnsupportedMediaType {
	return &GetTelephonyProvidersEdgeDiagnosticPingUnsupportedMediaType{}
}

/*GetTelephonyProvidersEdgeDiagnosticPingUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyProvidersEdgeDiagnosticPingUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeDiagnosticPingUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/ping][%d] getTelephonyProvidersEdgeDiagnosticPingUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticPingUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticPingUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeDiagnosticPingTooManyRequests creates a GetTelephonyProvidersEdgeDiagnosticPingTooManyRequests with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticPingTooManyRequests() *GetTelephonyProvidersEdgeDiagnosticPingTooManyRequests {
	return &GetTelephonyProvidersEdgeDiagnosticPingTooManyRequests{}
}

/*GetTelephonyProvidersEdgeDiagnosticPingTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetTelephonyProvidersEdgeDiagnosticPingTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeDiagnosticPingTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/ping][%d] getTelephonyProvidersEdgeDiagnosticPingTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticPingTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticPingTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeDiagnosticPingInternalServerError creates a GetTelephonyProvidersEdgeDiagnosticPingInternalServerError with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticPingInternalServerError() *GetTelephonyProvidersEdgeDiagnosticPingInternalServerError {
	return &GetTelephonyProvidersEdgeDiagnosticPingInternalServerError{}
}

/*GetTelephonyProvidersEdgeDiagnosticPingInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyProvidersEdgeDiagnosticPingInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeDiagnosticPingInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/ping][%d] getTelephonyProvidersEdgeDiagnosticPingInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticPingInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticPingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeDiagnosticPingServiceUnavailable creates a GetTelephonyProvidersEdgeDiagnosticPingServiceUnavailable with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticPingServiceUnavailable() *GetTelephonyProvidersEdgeDiagnosticPingServiceUnavailable {
	return &GetTelephonyProvidersEdgeDiagnosticPingServiceUnavailable{}
}

/*GetTelephonyProvidersEdgeDiagnosticPingServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyProvidersEdgeDiagnosticPingServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeDiagnosticPingServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/ping][%d] getTelephonyProvidersEdgeDiagnosticPingServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticPingServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticPingServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeDiagnosticPingGatewayTimeout creates a GetTelephonyProvidersEdgeDiagnosticPingGatewayTimeout with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticPingGatewayTimeout() *GetTelephonyProvidersEdgeDiagnosticPingGatewayTimeout {
	return &GetTelephonyProvidersEdgeDiagnosticPingGatewayTimeout{}
}

/*GetTelephonyProvidersEdgeDiagnosticPingGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetTelephonyProvidersEdgeDiagnosticPingGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeDiagnosticPingGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/ping][%d] getTelephonyProvidersEdgeDiagnosticPingGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticPingGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticPingGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

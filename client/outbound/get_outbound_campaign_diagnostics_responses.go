// Code generated by go-swagger; DO NOT EDIT.

package outbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetOutboundCampaignDiagnosticsReader is a Reader for the GetOutboundCampaignDiagnostics structure.
type GetOutboundCampaignDiagnosticsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutboundCampaignDiagnosticsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutboundCampaignDiagnosticsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutboundCampaignDiagnosticsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutboundCampaignDiagnosticsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutboundCampaignDiagnosticsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutboundCampaignDiagnosticsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOutboundCampaignDiagnosticsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOutboundCampaignDiagnosticsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOutboundCampaignDiagnosticsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOutboundCampaignDiagnosticsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOutboundCampaignDiagnosticsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOutboundCampaignDiagnosticsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOutboundCampaignDiagnosticsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOutboundCampaignDiagnosticsOK creates a GetOutboundCampaignDiagnosticsOK with default headers values
func NewGetOutboundCampaignDiagnosticsOK() *GetOutboundCampaignDiagnosticsOK {
	return &GetOutboundCampaignDiagnosticsOK{}
}

/*GetOutboundCampaignDiagnosticsOK handles this case with default header values.

successful operation
*/
type GetOutboundCampaignDiagnosticsOK struct {
	Payload *models.CampaignDiagnostics
}

func (o *GetOutboundCampaignDiagnosticsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/diagnostics][%d] getOutboundCampaignDiagnosticsOK  %+v", 200, o.Payload)
}

func (o *GetOutboundCampaignDiagnosticsOK) GetPayload() *models.CampaignDiagnostics {
	return o.Payload
}

func (o *GetOutboundCampaignDiagnosticsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CampaignDiagnostics)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignDiagnosticsBadRequest creates a GetOutboundCampaignDiagnosticsBadRequest with default headers values
func NewGetOutboundCampaignDiagnosticsBadRequest() *GetOutboundCampaignDiagnosticsBadRequest {
	return &GetOutboundCampaignDiagnosticsBadRequest{}
}

/*GetOutboundCampaignDiagnosticsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundCampaignDiagnosticsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignDiagnosticsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/diagnostics][%d] getOutboundCampaignDiagnosticsBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundCampaignDiagnosticsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignDiagnosticsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignDiagnosticsUnauthorized creates a GetOutboundCampaignDiagnosticsUnauthorized with default headers values
func NewGetOutboundCampaignDiagnosticsUnauthorized() *GetOutboundCampaignDiagnosticsUnauthorized {
	return &GetOutboundCampaignDiagnosticsUnauthorized{}
}

/*GetOutboundCampaignDiagnosticsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundCampaignDiagnosticsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignDiagnosticsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/diagnostics][%d] getOutboundCampaignDiagnosticsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundCampaignDiagnosticsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignDiagnosticsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignDiagnosticsForbidden creates a GetOutboundCampaignDiagnosticsForbidden with default headers values
func NewGetOutboundCampaignDiagnosticsForbidden() *GetOutboundCampaignDiagnosticsForbidden {
	return &GetOutboundCampaignDiagnosticsForbidden{}
}

/*GetOutboundCampaignDiagnosticsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundCampaignDiagnosticsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignDiagnosticsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/diagnostics][%d] getOutboundCampaignDiagnosticsForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundCampaignDiagnosticsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignDiagnosticsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignDiagnosticsNotFound creates a GetOutboundCampaignDiagnosticsNotFound with default headers values
func NewGetOutboundCampaignDiagnosticsNotFound() *GetOutboundCampaignDiagnosticsNotFound {
	return &GetOutboundCampaignDiagnosticsNotFound{}
}

/*GetOutboundCampaignDiagnosticsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetOutboundCampaignDiagnosticsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignDiagnosticsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/diagnostics][%d] getOutboundCampaignDiagnosticsNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundCampaignDiagnosticsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignDiagnosticsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignDiagnosticsRequestTimeout creates a GetOutboundCampaignDiagnosticsRequestTimeout with default headers values
func NewGetOutboundCampaignDiagnosticsRequestTimeout() *GetOutboundCampaignDiagnosticsRequestTimeout {
	return &GetOutboundCampaignDiagnosticsRequestTimeout{}
}

/*GetOutboundCampaignDiagnosticsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetOutboundCampaignDiagnosticsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignDiagnosticsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/diagnostics][%d] getOutboundCampaignDiagnosticsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundCampaignDiagnosticsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignDiagnosticsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignDiagnosticsRequestEntityTooLarge creates a GetOutboundCampaignDiagnosticsRequestEntityTooLarge with default headers values
func NewGetOutboundCampaignDiagnosticsRequestEntityTooLarge() *GetOutboundCampaignDiagnosticsRequestEntityTooLarge {
	return &GetOutboundCampaignDiagnosticsRequestEntityTooLarge{}
}

/*GetOutboundCampaignDiagnosticsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetOutboundCampaignDiagnosticsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignDiagnosticsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/diagnostics][%d] getOutboundCampaignDiagnosticsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundCampaignDiagnosticsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignDiagnosticsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignDiagnosticsUnsupportedMediaType creates a GetOutboundCampaignDiagnosticsUnsupportedMediaType with default headers values
func NewGetOutboundCampaignDiagnosticsUnsupportedMediaType() *GetOutboundCampaignDiagnosticsUnsupportedMediaType {
	return &GetOutboundCampaignDiagnosticsUnsupportedMediaType{}
}

/*GetOutboundCampaignDiagnosticsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundCampaignDiagnosticsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignDiagnosticsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/diagnostics][%d] getOutboundCampaignDiagnosticsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundCampaignDiagnosticsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignDiagnosticsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignDiagnosticsTooManyRequests creates a GetOutboundCampaignDiagnosticsTooManyRequests with default headers values
func NewGetOutboundCampaignDiagnosticsTooManyRequests() *GetOutboundCampaignDiagnosticsTooManyRequests {
	return &GetOutboundCampaignDiagnosticsTooManyRequests{}
}

/*GetOutboundCampaignDiagnosticsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetOutboundCampaignDiagnosticsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignDiagnosticsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/diagnostics][%d] getOutboundCampaignDiagnosticsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundCampaignDiagnosticsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignDiagnosticsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignDiagnosticsInternalServerError creates a GetOutboundCampaignDiagnosticsInternalServerError with default headers values
func NewGetOutboundCampaignDiagnosticsInternalServerError() *GetOutboundCampaignDiagnosticsInternalServerError {
	return &GetOutboundCampaignDiagnosticsInternalServerError{}
}

/*GetOutboundCampaignDiagnosticsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundCampaignDiagnosticsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignDiagnosticsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/diagnostics][%d] getOutboundCampaignDiagnosticsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundCampaignDiagnosticsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignDiagnosticsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignDiagnosticsServiceUnavailable creates a GetOutboundCampaignDiagnosticsServiceUnavailable with default headers values
func NewGetOutboundCampaignDiagnosticsServiceUnavailable() *GetOutboundCampaignDiagnosticsServiceUnavailable {
	return &GetOutboundCampaignDiagnosticsServiceUnavailable{}
}

/*GetOutboundCampaignDiagnosticsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundCampaignDiagnosticsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignDiagnosticsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/diagnostics][%d] getOutboundCampaignDiagnosticsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundCampaignDiagnosticsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignDiagnosticsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignDiagnosticsGatewayTimeout creates a GetOutboundCampaignDiagnosticsGatewayTimeout with default headers values
func NewGetOutboundCampaignDiagnosticsGatewayTimeout() *GetOutboundCampaignDiagnosticsGatewayTimeout {
	return &GetOutboundCampaignDiagnosticsGatewayTimeout{}
}

/*GetOutboundCampaignDiagnosticsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetOutboundCampaignDiagnosticsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignDiagnosticsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/diagnostics][%d] getOutboundCampaignDiagnosticsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundCampaignDiagnosticsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignDiagnosticsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

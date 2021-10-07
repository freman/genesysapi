// Code generated by go-swagger; DO NOT EDIT.

package analytics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetAnalyticsBotflowReportingturnsReader is a Reader for the GetAnalyticsBotflowReportingturns structure.
type GetAnalyticsBotflowReportingturnsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAnalyticsBotflowReportingturnsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAnalyticsBotflowReportingturnsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAnalyticsBotflowReportingturnsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAnalyticsBotflowReportingturnsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAnalyticsBotflowReportingturnsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAnalyticsBotflowReportingturnsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetAnalyticsBotflowReportingturnsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetAnalyticsBotflowReportingturnsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetAnalyticsBotflowReportingturnsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAnalyticsBotflowReportingturnsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAnalyticsBotflowReportingturnsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAnalyticsBotflowReportingturnsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAnalyticsBotflowReportingturnsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAnalyticsBotflowReportingturnsOK creates a GetAnalyticsBotflowReportingturnsOK with default headers values
func NewGetAnalyticsBotflowReportingturnsOK() *GetAnalyticsBotflowReportingturnsOK {
	return &GetAnalyticsBotflowReportingturnsOK{}
}

/*GetAnalyticsBotflowReportingturnsOK handles this case with default header values.

successful operation
*/
type GetAnalyticsBotflowReportingturnsOK struct {
	Payload *models.ReportingTurnsResponse
}

func (o *GetAnalyticsBotflowReportingturnsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/botflows/{botFlowId}/reportingturns][%d] getAnalyticsBotflowReportingturnsOK  %+v", 200, o.Payload)
}

func (o *GetAnalyticsBotflowReportingturnsOK) GetPayload() *models.ReportingTurnsResponse {
	return o.Payload
}

func (o *GetAnalyticsBotflowReportingturnsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReportingTurnsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsBotflowReportingturnsBadRequest creates a GetAnalyticsBotflowReportingturnsBadRequest with default headers values
func NewGetAnalyticsBotflowReportingturnsBadRequest() *GetAnalyticsBotflowReportingturnsBadRequest {
	return &GetAnalyticsBotflowReportingturnsBadRequest{}
}

/*GetAnalyticsBotflowReportingturnsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAnalyticsBotflowReportingturnsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsBotflowReportingturnsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/botflows/{botFlowId}/reportingturns][%d] getAnalyticsBotflowReportingturnsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAnalyticsBotflowReportingturnsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsBotflowReportingturnsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsBotflowReportingturnsUnauthorized creates a GetAnalyticsBotflowReportingturnsUnauthorized with default headers values
func NewGetAnalyticsBotflowReportingturnsUnauthorized() *GetAnalyticsBotflowReportingturnsUnauthorized {
	return &GetAnalyticsBotflowReportingturnsUnauthorized{}
}

/*GetAnalyticsBotflowReportingturnsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAnalyticsBotflowReportingturnsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsBotflowReportingturnsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/botflows/{botFlowId}/reportingturns][%d] getAnalyticsBotflowReportingturnsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAnalyticsBotflowReportingturnsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsBotflowReportingturnsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsBotflowReportingturnsForbidden creates a GetAnalyticsBotflowReportingturnsForbidden with default headers values
func NewGetAnalyticsBotflowReportingturnsForbidden() *GetAnalyticsBotflowReportingturnsForbidden {
	return &GetAnalyticsBotflowReportingturnsForbidden{}
}

/*GetAnalyticsBotflowReportingturnsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetAnalyticsBotflowReportingturnsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsBotflowReportingturnsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/botflows/{botFlowId}/reportingturns][%d] getAnalyticsBotflowReportingturnsForbidden  %+v", 403, o.Payload)
}

func (o *GetAnalyticsBotflowReportingturnsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsBotflowReportingturnsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsBotflowReportingturnsNotFound creates a GetAnalyticsBotflowReportingturnsNotFound with default headers values
func NewGetAnalyticsBotflowReportingturnsNotFound() *GetAnalyticsBotflowReportingturnsNotFound {
	return &GetAnalyticsBotflowReportingturnsNotFound{}
}

/*GetAnalyticsBotflowReportingturnsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetAnalyticsBotflowReportingturnsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsBotflowReportingturnsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/botflows/{botFlowId}/reportingturns][%d] getAnalyticsBotflowReportingturnsNotFound  %+v", 404, o.Payload)
}

func (o *GetAnalyticsBotflowReportingturnsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsBotflowReportingturnsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsBotflowReportingturnsRequestTimeout creates a GetAnalyticsBotflowReportingturnsRequestTimeout with default headers values
func NewGetAnalyticsBotflowReportingturnsRequestTimeout() *GetAnalyticsBotflowReportingturnsRequestTimeout {
	return &GetAnalyticsBotflowReportingturnsRequestTimeout{}
}

/*GetAnalyticsBotflowReportingturnsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetAnalyticsBotflowReportingturnsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsBotflowReportingturnsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/botflows/{botFlowId}/reportingturns][%d] getAnalyticsBotflowReportingturnsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAnalyticsBotflowReportingturnsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsBotflowReportingturnsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsBotflowReportingturnsRequestEntityTooLarge creates a GetAnalyticsBotflowReportingturnsRequestEntityTooLarge with default headers values
func NewGetAnalyticsBotflowReportingturnsRequestEntityTooLarge() *GetAnalyticsBotflowReportingturnsRequestEntityTooLarge {
	return &GetAnalyticsBotflowReportingturnsRequestEntityTooLarge{}
}

/*GetAnalyticsBotflowReportingturnsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetAnalyticsBotflowReportingturnsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsBotflowReportingturnsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/botflows/{botFlowId}/reportingturns][%d] getAnalyticsBotflowReportingturnsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAnalyticsBotflowReportingturnsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsBotflowReportingturnsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsBotflowReportingturnsUnsupportedMediaType creates a GetAnalyticsBotflowReportingturnsUnsupportedMediaType with default headers values
func NewGetAnalyticsBotflowReportingturnsUnsupportedMediaType() *GetAnalyticsBotflowReportingturnsUnsupportedMediaType {
	return &GetAnalyticsBotflowReportingturnsUnsupportedMediaType{}
}

/*GetAnalyticsBotflowReportingturnsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAnalyticsBotflowReportingturnsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsBotflowReportingturnsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/botflows/{botFlowId}/reportingturns][%d] getAnalyticsBotflowReportingturnsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAnalyticsBotflowReportingturnsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsBotflowReportingturnsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsBotflowReportingturnsTooManyRequests creates a GetAnalyticsBotflowReportingturnsTooManyRequests with default headers values
func NewGetAnalyticsBotflowReportingturnsTooManyRequests() *GetAnalyticsBotflowReportingturnsTooManyRequests {
	return &GetAnalyticsBotflowReportingturnsTooManyRequests{}
}

/*GetAnalyticsBotflowReportingturnsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetAnalyticsBotflowReportingturnsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsBotflowReportingturnsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/botflows/{botFlowId}/reportingturns][%d] getAnalyticsBotflowReportingturnsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAnalyticsBotflowReportingturnsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsBotflowReportingturnsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsBotflowReportingturnsInternalServerError creates a GetAnalyticsBotflowReportingturnsInternalServerError with default headers values
func NewGetAnalyticsBotflowReportingturnsInternalServerError() *GetAnalyticsBotflowReportingturnsInternalServerError {
	return &GetAnalyticsBotflowReportingturnsInternalServerError{}
}

/*GetAnalyticsBotflowReportingturnsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAnalyticsBotflowReportingturnsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsBotflowReportingturnsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/botflows/{botFlowId}/reportingturns][%d] getAnalyticsBotflowReportingturnsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAnalyticsBotflowReportingturnsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsBotflowReportingturnsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsBotflowReportingturnsServiceUnavailable creates a GetAnalyticsBotflowReportingturnsServiceUnavailable with default headers values
func NewGetAnalyticsBotflowReportingturnsServiceUnavailable() *GetAnalyticsBotflowReportingturnsServiceUnavailable {
	return &GetAnalyticsBotflowReportingturnsServiceUnavailable{}
}

/*GetAnalyticsBotflowReportingturnsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAnalyticsBotflowReportingturnsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsBotflowReportingturnsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/botflows/{botFlowId}/reportingturns][%d] getAnalyticsBotflowReportingturnsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAnalyticsBotflowReportingturnsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsBotflowReportingturnsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsBotflowReportingturnsGatewayTimeout creates a GetAnalyticsBotflowReportingturnsGatewayTimeout with default headers values
func NewGetAnalyticsBotflowReportingturnsGatewayTimeout() *GetAnalyticsBotflowReportingturnsGatewayTimeout {
	return &GetAnalyticsBotflowReportingturnsGatewayTimeout{}
}

/*GetAnalyticsBotflowReportingturnsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetAnalyticsBotflowReportingturnsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsBotflowReportingturnsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/botflows/{botFlowId}/reportingturns][%d] getAnalyticsBotflowReportingturnsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAnalyticsBotflowReportingturnsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsBotflowReportingturnsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

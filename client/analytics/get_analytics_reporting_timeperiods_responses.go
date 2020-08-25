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

// GetAnalyticsReportingTimeperiodsReader is a Reader for the GetAnalyticsReportingTimeperiods structure.
type GetAnalyticsReportingTimeperiodsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAnalyticsReportingTimeperiodsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAnalyticsReportingTimeperiodsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAnalyticsReportingTimeperiodsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAnalyticsReportingTimeperiodsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAnalyticsReportingTimeperiodsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAnalyticsReportingTimeperiodsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetAnalyticsReportingTimeperiodsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetAnalyticsReportingTimeperiodsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAnalyticsReportingTimeperiodsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAnalyticsReportingTimeperiodsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAnalyticsReportingTimeperiodsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAnalyticsReportingTimeperiodsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAnalyticsReportingTimeperiodsOK creates a GetAnalyticsReportingTimeperiodsOK with default headers values
func NewGetAnalyticsReportingTimeperiodsOK() *GetAnalyticsReportingTimeperiodsOK {
	return &GetAnalyticsReportingTimeperiodsOK{}
}

/*GetAnalyticsReportingTimeperiodsOK handles this case with default header values.

successful operation
*/
type GetAnalyticsReportingTimeperiodsOK struct {
	Payload []string
}

func (o *GetAnalyticsReportingTimeperiodsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/timeperiods][%d] getAnalyticsReportingTimeperiodsOK  %+v", 200, o.Payload)
}

func (o *GetAnalyticsReportingTimeperiodsOK) GetPayload() []string {
	return o.Payload
}

func (o *GetAnalyticsReportingTimeperiodsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingTimeperiodsBadRequest creates a GetAnalyticsReportingTimeperiodsBadRequest with default headers values
func NewGetAnalyticsReportingTimeperiodsBadRequest() *GetAnalyticsReportingTimeperiodsBadRequest {
	return &GetAnalyticsReportingTimeperiodsBadRequest{}
}

/*GetAnalyticsReportingTimeperiodsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAnalyticsReportingTimeperiodsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingTimeperiodsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/timeperiods][%d] getAnalyticsReportingTimeperiodsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAnalyticsReportingTimeperiodsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingTimeperiodsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingTimeperiodsUnauthorized creates a GetAnalyticsReportingTimeperiodsUnauthorized with default headers values
func NewGetAnalyticsReportingTimeperiodsUnauthorized() *GetAnalyticsReportingTimeperiodsUnauthorized {
	return &GetAnalyticsReportingTimeperiodsUnauthorized{}
}

/*GetAnalyticsReportingTimeperiodsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAnalyticsReportingTimeperiodsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingTimeperiodsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/timeperiods][%d] getAnalyticsReportingTimeperiodsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAnalyticsReportingTimeperiodsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingTimeperiodsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingTimeperiodsForbidden creates a GetAnalyticsReportingTimeperiodsForbidden with default headers values
func NewGetAnalyticsReportingTimeperiodsForbidden() *GetAnalyticsReportingTimeperiodsForbidden {
	return &GetAnalyticsReportingTimeperiodsForbidden{}
}

/*GetAnalyticsReportingTimeperiodsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetAnalyticsReportingTimeperiodsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingTimeperiodsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/timeperiods][%d] getAnalyticsReportingTimeperiodsForbidden  %+v", 403, o.Payload)
}

func (o *GetAnalyticsReportingTimeperiodsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingTimeperiodsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingTimeperiodsNotFound creates a GetAnalyticsReportingTimeperiodsNotFound with default headers values
func NewGetAnalyticsReportingTimeperiodsNotFound() *GetAnalyticsReportingTimeperiodsNotFound {
	return &GetAnalyticsReportingTimeperiodsNotFound{}
}

/*GetAnalyticsReportingTimeperiodsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetAnalyticsReportingTimeperiodsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingTimeperiodsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/timeperiods][%d] getAnalyticsReportingTimeperiodsNotFound  %+v", 404, o.Payload)
}

func (o *GetAnalyticsReportingTimeperiodsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingTimeperiodsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingTimeperiodsRequestEntityTooLarge creates a GetAnalyticsReportingTimeperiodsRequestEntityTooLarge with default headers values
func NewGetAnalyticsReportingTimeperiodsRequestEntityTooLarge() *GetAnalyticsReportingTimeperiodsRequestEntityTooLarge {
	return &GetAnalyticsReportingTimeperiodsRequestEntityTooLarge{}
}

/*GetAnalyticsReportingTimeperiodsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetAnalyticsReportingTimeperiodsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingTimeperiodsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/timeperiods][%d] getAnalyticsReportingTimeperiodsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAnalyticsReportingTimeperiodsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingTimeperiodsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingTimeperiodsUnsupportedMediaType creates a GetAnalyticsReportingTimeperiodsUnsupportedMediaType with default headers values
func NewGetAnalyticsReportingTimeperiodsUnsupportedMediaType() *GetAnalyticsReportingTimeperiodsUnsupportedMediaType {
	return &GetAnalyticsReportingTimeperiodsUnsupportedMediaType{}
}

/*GetAnalyticsReportingTimeperiodsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAnalyticsReportingTimeperiodsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingTimeperiodsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/timeperiods][%d] getAnalyticsReportingTimeperiodsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAnalyticsReportingTimeperiodsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingTimeperiodsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingTimeperiodsTooManyRequests creates a GetAnalyticsReportingTimeperiodsTooManyRequests with default headers values
func NewGetAnalyticsReportingTimeperiodsTooManyRequests() *GetAnalyticsReportingTimeperiodsTooManyRequests {
	return &GetAnalyticsReportingTimeperiodsTooManyRequests{}
}

/*GetAnalyticsReportingTimeperiodsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetAnalyticsReportingTimeperiodsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingTimeperiodsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/timeperiods][%d] getAnalyticsReportingTimeperiodsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAnalyticsReportingTimeperiodsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingTimeperiodsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingTimeperiodsInternalServerError creates a GetAnalyticsReportingTimeperiodsInternalServerError with default headers values
func NewGetAnalyticsReportingTimeperiodsInternalServerError() *GetAnalyticsReportingTimeperiodsInternalServerError {
	return &GetAnalyticsReportingTimeperiodsInternalServerError{}
}

/*GetAnalyticsReportingTimeperiodsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAnalyticsReportingTimeperiodsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingTimeperiodsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/timeperiods][%d] getAnalyticsReportingTimeperiodsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAnalyticsReportingTimeperiodsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingTimeperiodsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingTimeperiodsServiceUnavailable creates a GetAnalyticsReportingTimeperiodsServiceUnavailable with default headers values
func NewGetAnalyticsReportingTimeperiodsServiceUnavailable() *GetAnalyticsReportingTimeperiodsServiceUnavailable {
	return &GetAnalyticsReportingTimeperiodsServiceUnavailable{}
}

/*GetAnalyticsReportingTimeperiodsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAnalyticsReportingTimeperiodsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingTimeperiodsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/timeperiods][%d] getAnalyticsReportingTimeperiodsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAnalyticsReportingTimeperiodsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingTimeperiodsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingTimeperiodsGatewayTimeout creates a GetAnalyticsReportingTimeperiodsGatewayTimeout with default headers values
func NewGetAnalyticsReportingTimeperiodsGatewayTimeout() *GetAnalyticsReportingTimeperiodsGatewayTimeout {
	return &GetAnalyticsReportingTimeperiodsGatewayTimeout{}
}

/*GetAnalyticsReportingTimeperiodsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetAnalyticsReportingTimeperiodsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingTimeperiodsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/timeperiods][%d] getAnalyticsReportingTimeperiodsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAnalyticsReportingTimeperiodsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingTimeperiodsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
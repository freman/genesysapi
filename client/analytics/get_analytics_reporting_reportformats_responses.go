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

// GetAnalyticsReportingReportformatsReader is a Reader for the GetAnalyticsReportingReportformats structure.
type GetAnalyticsReportingReportformatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAnalyticsReportingReportformatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAnalyticsReportingReportformatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAnalyticsReportingReportformatsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAnalyticsReportingReportformatsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAnalyticsReportingReportformatsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAnalyticsReportingReportformatsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetAnalyticsReportingReportformatsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetAnalyticsReportingReportformatsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAnalyticsReportingReportformatsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAnalyticsReportingReportformatsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAnalyticsReportingReportformatsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAnalyticsReportingReportformatsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAnalyticsReportingReportformatsOK creates a GetAnalyticsReportingReportformatsOK with default headers values
func NewGetAnalyticsReportingReportformatsOK() *GetAnalyticsReportingReportformatsOK {
	return &GetAnalyticsReportingReportformatsOK{}
}

/*GetAnalyticsReportingReportformatsOK handles this case with default header values.

successful operation
*/
type GetAnalyticsReportingReportformatsOK struct {
	Payload []string
}

func (o *GetAnalyticsReportingReportformatsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/reportformats][%d] getAnalyticsReportingReportformatsOK  %+v", 200, o.Payload)
}

func (o *GetAnalyticsReportingReportformatsOK) GetPayload() []string {
	return o.Payload
}

func (o *GetAnalyticsReportingReportformatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingReportformatsBadRequest creates a GetAnalyticsReportingReportformatsBadRequest with default headers values
func NewGetAnalyticsReportingReportformatsBadRequest() *GetAnalyticsReportingReportformatsBadRequest {
	return &GetAnalyticsReportingReportformatsBadRequest{}
}

/*GetAnalyticsReportingReportformatsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAnalyticsReportingReportformatsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingReportformatsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/reportformats][%d] getAnalyticsReportingReportformatsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAnalyticsReportingReportformatsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingReportformatsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingReportformatsUnauthorized creates a GetAnalyticsReportingReportformatsUnauthorized with default headers values
func NewGetAnalyticsReportingReportformatsUnauthorized() *GetAnalyticsReportingReportformatsUnauthorized {
	return &GetAnalyticsReportingReportformatsUnauthorized{}
}

/*GetAnalyticsReportingReportformatsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAnalyticsReportingReportformatsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingReportformatsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/reportformats][%d] getAnalyticsReportingReportformatsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAnalyticsReportingReportformatsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingReportformatsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingReportformatsForbidden creates a GetAnalyticsReportingReportformatsForbidden with default headers values
func NewGetAnalyticsReportingReportformatsForbidden() *GetAnalyticsReportingReportformatsForbidden {
	return &GetAnalyticsReportingReportformatsForbidden{}
}

/*GetAnalyticsReportingReportformatsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetAnalyticsReportingReportformatsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingReportformatsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/reportformats][%d] getAnalyticsReportingReportformatsForbidden  %+v", 403, o.Payload)
}

func (o *GetAnalyticsReportingReportformatsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingReportformatsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingReportformatsNotFound creates a GetAnalyticsReportingReportformatsNotFound with default headers values
func NewGetAnalyticsReportingReportformatsNotFound() *GetAnalyticsReportingReportformatsNotFound {
	return &GetAnalyticsReportingReportformatsNotFound{}
}

/*GetAnalyticsReportingReportformatsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetAnalyticsReportingReportformatsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingReportformatsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/reportformats][%d] getAnalyticsReportingReportformatsNotFound  %+v", 404, o.Payload)
}

func (o *GetAnalyticsReportingReportformatsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingReportformatsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingReportformatsRequestEntityTooLarge creates a GetAnalyticsReportingReportformatsRequestEntityTooLarge with default headers values
func NewGetAnalyticsReportingReportformatsRequestEntityTooLarge() *GetAnalyticsReportingReportformatsRequestEntityTooLarge {
	return &GetAnalyticsReportingReportformatsRequestEntityTooLarge{}
}

/*GetAnalyticsReportingReportformatsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetAnalyticsReportingReportformatsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingReportformatsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/reportformats][%d] getAnalyticsReportingReportformatsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAnalyticsReportingReportformatsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingReportformatsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingReportformatsUnsupportedMediaType creates a GetAnalyticsReportingReportformatsUnsupportedMediaType with default headers values
func NewGetAnalyticsReportingReportformatsUnsupportedMediaType() *GetAnalyticsReportingReportformatsUnsupportedMediaType {
	return &GetAnalyticsReportingReportformatsUnsupportedMediaType{}
}

/*GetAnalyticsReportingReportformatsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAnalyticsReportingReportformatsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingReportformatsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/reportformats][%d] getAnalyticsReportingReportformatsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAnalyticsReportingReportformatsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingReportformatsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingReportformatsTooManyRequests creates a GetAnalyticsReportingReportformatsTooManyRequests with default headers values
func NewGetAnalyticsReportingReportformatsTooManyRequests() *GetAnalyticsReportingReportformatsTooManyRequests {
	return &GetAnalyticsReportingReportformatsTooManyRequests{}
}

/*GetAnalyticsReportingReportformatsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetAnalyticsReportingReportformatsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingReportformatsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/reportformats][%d] getAnalyticsReportingReportformatsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAnalyticsReportingReportformatsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingReportformatsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingReportformatsInternalServerError creates a GetAnalyticsReportingReportformatsInternalServerError with default headers values
func NewGetAnalyticsReportingReportformatsInternalServerError() *GetAnalyticsReportingReportformatsInternalServerError {
	return &GetAnalyticsReportingReportformatsInternalServerError{}
}

/*GetAnalyticsReportingReportformatsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAnalyticsReportingReportformatsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingReportformatsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/reportformats][%d] getAnalyticsReportingReportformatsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAnalyticsReportingReportformatsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingReportformatsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingReportformatsServiceUnavailable creates a GetAnalyticsReportingReportformatsServiceUnavailable with default headers values
func NewGetAnalyticsReportingReportformatsServiceUnavailable() *GetAnalyticsReportingReportformatsServiceUnavailable {
	return &GetAnalyticsReportingReportformatsServiceUnavailable{}
}

/*GetAnalyticsReportingReportformatsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAnalyticsReportingReportformatsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingReportformatsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/reportformats][%d] getAnalyticsReportingReportformatsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAnalyticsReportingReportformatsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingReportformatsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingReportformatsGatewayTimeout creates a GetAnalyticsReportingReportformatsGatewayTimeout with default headers values
func NewGetAnalyticsReportingReportformatsGatewayTimeout() *GetAnalyticsReportingReportformatsGatewayTimeout {
	return &GetAnalyticsReportingReportformatsGatewayTimeout{}
}

/*GetAnalyticsReportingReportformatsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetAnalyticsReportingReportformatsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingReportformatsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/reportformats][%d] getAnalyticsReportingReportformatsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAnalyticsReportingReportformatsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingReportformatsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
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

// PostAnalyticsReportingExportsReader is a Reader for the PostAnalyticsReportingExports structure.
type PostAnalyticsReportingExportsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAnalyticsReportingExportsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAnalyticsReportingExportsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAnalyticsReportingExportsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostAnalyticsReportingExportsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostAnalyticsReportingExportsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAnalyticsReportingExportsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostAnalyticsReportingExportsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostAnalyticsReportingExportsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostAnalyticsReportingExportsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAnalyticsReportingExportsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostAnalyticsReportingExportsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostAnalyticsReportingExportsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAnalyticsReportingExportsOK creates a PostAnalyticsReportingExportsOK with default headers values
func NewPostAnalyticsReportingExportsOK() *PostAnalyticsReportingExportsOK {
	return &PostAnalyticsReportingExportsOK{}
}

/*PostAnalyticsReportingExportsOK handles this case with default header values.

successful operation
*/
type PostAnalyticsReportingExportsOK struct {
	Payload *models.ReportingExportJobResponse
}

func (o *PostAnalyticsReportingExportsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/reporting/exports][%d] postAnalyticsReportingExportsOK  %+v", 200, o.Payload)
}

func (o *PostAnalyticsReportingExportsOK) GetPayload() *models.ReportingExportJobResponse {
	return o.Payload
}

func (o *PostAnalyticsReportingExportsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReportingExportJobResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsReportingExportsBadRequest creates a PostAnalyticsReportingExportsBadRequest with default headers values
func NewPostAnalyticsReportingExportsBadRequest() *PostAnalyticsReportingExportsBadRequest {
	return &PostAnalyticsReportingExportsBadRequest{}
}

/*PostAnalyticsReportingExportsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostAnalyticsReportingExportsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsReportingExportsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/reporting/exports][%d] postAnalyticsReportingExportsBadRequest  %+v", 400, o.Payload)
}

func (o *PostAnalyticsReportingExportsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsReportingExportsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsReportingExportsUnauthorized creates a PostAnalyticsReportingExportsUnauthorized with default headers values
func NewPostAnalyticsReportingExportsUnauthorized() *PostAnalyticsReportingExportsUnauthorized {
	return &PostAnalyticsReportingExportsUnauthorized{}
}

/*PostAnalyticsReportingExportsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostAnalyticsReportingExportsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsReportingExportsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/reporting/exports][%d] postAnalyticsReportingExportsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostAnalyticsReportingExportsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsReportingExportsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsReportingExportsForbidden creates a PostAnalyticsReportingExportsForbidden with default headers values
func NewPostAnalyticsReportingExportsForbidden() *PostAnalyticsReportingExportsForbidden {
	return &PostAnalyticsReportingExportsForbidden{}
}

/*PostAnalyticsReportingExportsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostAnalyticsReportingExportsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsReportingExportsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/reporting/exports][%d] postAnalyticsReportingExportsForbidden  %+v", 403, o.Payload)
}

func (o *PostAnalyticsReportingExportsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsReportingExportsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsReportingExportsNotFound creates a PostAnalyticsReportingExportsNotFound with default headers values
func NewPostAnalyticsReportingExportsNotFound() *PostAnalyticsReportingExportsNotFound {
	return &PostAnalyticsReportingExportsNotFound{}
}

/*PostAnalyticsReportingExportsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostAnalyticsReportingExportsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsReportingExportsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/reporting/exports][%d] postAnalyticsReportingExportsNotFound  %+v", 404, o.Payload)
}

func (o *PostAnalyticsReportingExportsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsReportingExportsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsReportingExportsRequestEntityTooLarge creates a PostAnalyticsReportingExportsRequestEntityTooLarge with default headers values
func NewPostAnalyticsReportingExportsRequestEntityTooLarge() *PostAnalyticsReportingExportsRequestEntityTooLarge {
	return &PostAnalyticsReportingExportsRequestEntityTooLarge{}
}

/*PostAnalyticsReportingExportsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostAnalyticsReportingExportsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsReportingExportsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/reporting/exports][%d] postAnalyticsReportingExportsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostAnalyticsReportingExportsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsReportingExportsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsReportingExportsUnsupportedMediaType creates a PostAnalyticsReportingExportsUnsupportedMediaType with default headers values
func NewPostAnalyticsReportingExportsUnsupportedMediaType() *PostAnalyticsReportingExportsUnsupportedMediaType {
	return &PostAnalyticsReportingExportsUnsupportedMediaType{}
}

/*PostAnalyticsReportingExportsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostAnalyticsReportingExportsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsReportingExportsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/reporting/exports][%d] postAnalyticsReportingExportsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostAnalyticsReportingExportsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsReportingExportsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsReportingExportsTooManyRequests creates a PostAnalyticsReportingExportsTooManyRequests with default headers values
func NewPostAnalyticsReportingExportsTooManyRequests() *PostAnalyticsReportingExportsTooManyRequests {
	return &PostAnalyticsReportingExportsTooManyRequests{}
}

/*PostAnalyticsReportingExportsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostAnalyticsReportingExportsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsReportingExportsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/reporting/exports][%d] postAnalyticsReportingExportsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostAnalyticsReportingExportsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsReportingExportsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsReportingExportsInternalServerError creates a PostAnalyticsReportingExportsInternalServerError with default headers values
func NewPostAnalyticsReportingExportsInternalServerError() *PostAnalyticsReportingExportsInternalServerError {
	return &PostAnalyticsReportingExportsInternalServerError{}
}

/*PostAnalyticsReportingExportsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostAnalyticsReportingExportsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsReportingExportsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/reporting/exports][%d] postAnalyticsReportingExportsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAnalyticsReportingExportsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsReportingExportsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsReportingExportsServiceUnavailable creates a PostAnalyticsReportingExportsServiceUnavailable with default headers values
func NewPostAnalyticsReportingExportsServiceUnavailable() *PostAnalyticsReportingExportsServiceUnavailable {
	return &PostAnalyticsReportingExportsServiceUnavailable{}
}

/*PostAnalyticsReportingExportsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostAnalyticsReportingExportsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsReportingExportsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/reporting/exports][%d] postAnalyticsReportingExportsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostAnalyticsReportingExportsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsReportingExportsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsReportingExportsGatewayTimeout creates a PostAnalyticsReportingExportsGatewayTimeout with default headers values
func NewPostAnalyticsReportingExportsGatewayTimeout() *PostAnalyticsReportingExportsGatewayTimeout {
	return &PostAnalyticsReportingExportsGatewayTimeout{}
}

/*PostAnalyticsReportingExportsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostAnalyticsReportingExportsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsReportingExportsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/reporting/exports][%d] postAnalyticsReportingExportsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostAnalyticsReportingExportsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsReportingExportsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

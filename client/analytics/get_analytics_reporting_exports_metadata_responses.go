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

// GetAnalyticsReportingExportsMetadataReader is a Reader for the GetAnalyticsReportingExportsMetadata structure.
type GetAnalyticsReportingExportsMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAnalyticsReportingExportsMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAnalyticsReportingExportsMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAnalyticsReportingExportsMetadataBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAnalyticsReportingExportsMetadataUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAnalyticsReportingExportsMetadataForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAnalyticsReportingExportsMetadataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetAnalyticsReportingExportsMetadataRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetAnalyticsReportingExportsMetadataRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetAnalyticsReportingExportsMetadataUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAnalyticsReportingExportsMetadataTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAnalyticsReportingExportsMetadataInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAnalyticsReportingExportsMetadataServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAnalyticsReportingExportsMetadataGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAnalyticsReportingExportsMetadataOK creates a GetAnalyticsReportingExportsMetadataOK with default headers values
func NewGetAnalyticsReportingExportsMetadataOK() *GetAnalyticsReportingExportsMetadataOK {
	return &GetAnalyticsReportingExportsMetadataOK{}
}

/*GetAnalyticsReportingExportsMetadataOK handles this case with default header values.

successful operation
*/
type GetAnalyticsReportingExportsMetadataOK struct {
	Payload *models.ReportingExportMetadataJobListing
}

func (o *GetAnalyticsReportingExportsMetadataOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/exports/metadata][%d] getAnalyticsReportingExportsMetadataOK  %+v", 200, o.Payload)
}

func (o *GetAnalyticsReportingExportsMetadataOK) GetPayload() *models.ReportingExportMetadataJobListing {
	return o.Payload
}

func (o *GetAnalyticsReportingExportsMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReportingExportMetadataJobListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingExportsMetadataBadRequest creates a GetAnalyticsReportingExportsMetadataBadRequest with default headers values
func NewGetAnalyticsReportingExportsMetadataBadRequest() *GetAnalyticsReportingExportsMetadataBadRequest {
	return &GetAnalyticsReportingExportsMetadataBadRequest{}
}

/*GetAnalyticsReportingExportsMetadataBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAnalyticsReportingExportsMetadataBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingExportsMetadataBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/exports/metadata][%d] getAnalyticsReportingExportsMetadataBadRequest  %+v", 400, o.Payload)
}

func (o *GetAnalyticsReportingExportsMetadataBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingExportsMetadataBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingExportsMetadataUnauthorized creates a GetAnalyticsReportingExportsMetadataUnauthorized with default headers values
func NewGetAnalyticsReportingExportsMetadataUnauthorized() *GetAnalyticsReportingExportsMetadataUnauthorized {
	return &GetAnalyticsReportingExportsMetadataUnauthorized{}
}

/*GetAnalyticsReportingExportsMetadataUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAnalyticsReportingExportsMetadataUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingExportsMetadataUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/exports/metadata][%d] getAnalyticsReportingExportsMetadataUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAnalyticsReportingExportsMetadataUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingExportsMetadataUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingExportsMetadataForbidden creates a GetAnalyticsReportingExportsMetadataForbidden with default headers values
func NewGetAnalyticsReportingExportsMetadataForbidden() *GetAnalyticsReportingExportsMetadataForbidden {
	return &GetAnalyticsReportingExportsMetadataForbidden{}
}

/*GetAnalyticsReportingExportsMetadataForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetAnalyticsReportingExportsMetadataForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingExportsMetadataForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/exports/metadata][%d] getAnalyticsReportingExportsMetadataForbidden  %+v", 403, o.Payload)
}

func (o *GetAnalyticsReportingExportsMetadataForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingExportsMetadataForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingExportsMetadataNotFound creates a GetAnalyticsReportingExportsMetadataNotFound with default headers values
func NewGetAnalyticsReportingExportsMetadataNotFound() *GetAnalyticsReportingExportsMetadataNotFound {
	return &GetAnalyticsReportingExportsMetadataNotFound{}
}

/*GetAnalyticsReportingExportsMetadataNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetAnalyticsReportingExportsMetadataNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingExportsMetadataNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/exports/metadata][%d] getAnalyticsReportingExportsMetadataNotFound  %+v", 404, o.Payload)
}

func (o *GetAnalyticsReportingExportsMetadataNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingExportsMetadataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingExportsMetadataRequestTimeout creates a GetAnalyticsReportingExportsMetadataRequestTimeout with default headers values
func NewGetAnalyticsReportingExportsMetadataRequestTimeout() *GetAnalyticsReportingExportsMetadataRequestTimeout {
	return &GetAnalyticsReportingExportsMetadataRequestTimeout{}
}

/*GetAnalyticsReportingExportsMetadataRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetAnalyticsReportingExportsMetadataRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingExportsMetadataRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/exports/metadata][%d] getAnalyticsReportingExportsMetadataRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAnalyticsReportingExportsMetadataRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingExportsMetadataRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingExportsMetadataRequestEntityTooLarge creates a GetAnalyticsReportingExportsMetadataRequestEntityTooLarge with default headers values
func NewGetAnalyticsReportingExportsMetadataRequestEntityTooLarge() *GetAnalyticsReportingExportsMetadataRequestEntityTooLarge {
	return &GetAnalyticsReportingExportsMetadataRequestEntityTooLarge{}
}

/*GetAnalyticsReportingExportsMetadataRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetAnalyticsReportingExportsMetadataRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingExportsMetadataRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/exports/metadata][%d] getAnalyticsReportingExportsMetadataRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAnalyticsReportingExportsMetadataRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingExportsMetadataRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingExportsMetadataUnsupportedMediaType creates a GetAnalyticsReportingExportsMetadataUnsupportedMediaType with default headers values
func NewGetAnalyticsReportingExportsMetadataUnsupportedMediaType() *GetAnalyticsReportingExportsMetadataUnsupportedMediaType {
	return &GetAnalyticsReportingExportsMetadataUnsupportedMediaType{}
}

/*GetAnalyticsReportingExportsMetadataUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAnalyticsReportingExportsMetadataUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingExportsMetadataUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/exports/metadata][%d] getAnalyticsReportingExportsMetadataUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAnalyticsReportingExportsMetadataUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingExportsMetadataUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingExportsMetadataTooManyRequests creates a GetAnalyticsReportingExportsMetadataTooManyRequests with default headers values
func NewGetAnalyticsReportingExportsMetadataTooManyRequests() *GetAnalyticsReportingExportsMetadataTooManyRequests {
	return &GetAnalyticsReportingExportsMetadataTooManyRequests{}
}

/*GetAnalyticsReportingExportsMetadataTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetAnalyticsReportingExportsMetadataTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingExportsMetadataTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/exports/metadata][%d] getAnalyticsReportingExportsMetadataTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAnalyticsReportingExportsMetadataTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingExportsMetadataTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingExportsMetadataInternalServerError creates a GetAnalyticsReportingExportsMetadataInternalServerError with default headers values
func NewGetAnalyticsReportingExportsMetadataInternalServerError() *GetAnalyticsReportingExportsMetadataInternalServerError {
	return &GetAnalyticsReportingExportsMetadataInternalServerError{}
}

/*GetAnalyticsReportingExportsMetadataInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAnalyticsReportingExportsMetadataInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingExportsMetadataInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/exports/metadata][%d] getAnalyticsReportingExportsMetadataInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAnalyticsReportingExportsMetadataInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingExportsMetadataInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingExportsMetadataServiceUnavailable creates a GetAnalyticsReportingExportsMetadataServiceUnavailable with default headers values
func NewGetAnalyticsReportingExportsMetadataServiceUnavailable() *GetAnalyticsReportingExportsMetadataServiceUnavailable {
	return &GetAnalyticsReportingExportsMetadataServiceUnavailable{}
}

/*GetAnalyticsReportingExportsMetadataServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAnalyticsReportingExportsMetadataServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingExportsMetadataServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/exports/metadata][%d] getAnalyticsReportingExportsMetadataServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAnalyticsReportingExportsMetadataServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingExportsMetadataServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingExportsMetadataGatewayTimeout creates a GetAnalyticsReportingExportsMetadataGatewayTimeout with default headers values
func NewGetAnalyticsReportingExportsMetadataGatewayTimeout() *GetAnalyticsReportingExportsMetadataGatewayTimeout {
	return &GetAnalyticsReportingExportsMetadataGatewayTimeout{}
}

/*GetAnalyticsReportingExportsMetadataGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetAnalyticsReportingExportsMetadataGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingExportsMetadataGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/exports/metadata][%d] getAnalyticsReportingExportsMetadataGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAnalyticsReportingExportsMetadataGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingExportsMetadataGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

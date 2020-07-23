// Code generated by go-swagger; DO NOT EDIT.

package billing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetBillingReportsBillableusageReader is a Reader for the GetBillingReportsBillableusage structure.
type GetBillingReportsBillableusageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBillingReportsBillableusageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBillingReportsBillableusageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetBillingReportsBillableusageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetBillingReportsBillableusageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetBillingReportsBillableusageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetBillingReportsBillableusageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetBillingReportsBillableusageRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetBillingReportsBillableusageUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetBillingReportsBillableusageTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetBillingReportsBillableusageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetBillingReportsBillableusageServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetBillingReportsBillableusageGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBillingReportsBillableusageOK creates a GetBillingReportsBillableusageOK with default headers values
func NewGetBillingReportsBillableusageOK() *GetBillingReportsBillableusageOK {
	return &GetBillingReportsBillableusageOK{}
}

/*GetBillingReportsBillableusageOK handles this case with default header values.

successful operation
*/
type GetBillingReportsBillableusageOK struct {
	Payload *models.BillingUsageReport
}

func (o *GetBillingReportsBillableusageOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/billing/reports/billableusage][%d] getBillingReportsBillableusageOK  %+v", 200, o.Payload)
}

func (o *GetBillingReportsBillableusageOK) GetPayload() *models.BillingUsageReport {
	return o.Payload
}

func (o *GetBillingReportsBillableusageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BillingUsageReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBillingReportsBillableusageBadRequest creates a GetBillingReportsBillableusageBadRequest with default headers values
func NewGetBillingReportsBillableusageBadRequest() *GetBillingReportsBillableusageBadRequest {
	return &GetBillingReportsBillableusageBadRequest{}
}

/*GetBillingReportsBillableusageBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetBillingReportsBillableusageBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetBillingReportsBillableusageBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/billing/reports/billableusage][%d] getBillingReportsBillableusageBadRequest  %+v", 400, o.Payload)
}

func (o *GetBillingReportsBillableusageBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetBillingReportsBillableusageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBillingReportsBillableusageUnauthorized creates a GetBillingReportsBillableusageUnauthorized with default headers values
func NewGetBillingReportsBillableusageUnauthorized() *GetBillingReportsBillableusageUnauthorized {
	return &GetBillingReportsBillableusageUnauthorized{}
}

/*GetBillingReportsBillableusageUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetBillingReportsBillableusageUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetBillingReportsBillableusageUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/billing/reports/billableusage][%d] getBillingReportsBillableusageUnauthorized  %+v", 401, o.Payload)
}

func (o *GetBillingReportsBillableusageUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetBillingReportsBillableusageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBillingReportsBillableusageForbidden creates a GetBillingReportsBillableusageForbidden with default headers values
func NewGetBillingReportsBillableusageForbidden() *GetBillingReportsBillableusageForbidden {
	return &GetBillingReportsBillableusageForbidden{}
}

/*GetBillingReportsBillableusageForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetBillingReportsBillableusageForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetBillingReportsBillableusageForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/billing/reports/billableusage][%d] getBillingReportsBillableusageForbidden  %+v", 403, o.Payload)
}

func (o *GetBillingReportsBillableusageForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetBillingReportsBillableusageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBillingReportsBillableusageNotFound creates a GetBillingReportsBillableusageNotFound with default headers values
func NewGetBillingReportsBillableusageNotFound() *GetBillingReportsBillableusageNotFound {
	return &GetBillingReportsBillableusageNotFound{}
}

/*GetBillingReportsBillableusageNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetBillingReportsBillableusageNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetBillingReportsBillableusageNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/billing/reports/billableusage][%d] getBillingReportsBillableusageNotFound  %+v", 404, o.Payload)
}

func (o *GetBillingReportsBillableusageNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetBillingReportsBillableusageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBillingReportsBillableusageRequestEntityTooLarge creates a GetBillingReportsBillableusageRequestEntityTooLarge with default headers values
func NewGetBillingReportsBillableusageRequestEntityTooLarge() *GetBillingReportsBillableusageRequestEntityTooLarge {
	return &GetBillingReportsBillableusageRequestEntityTooLarge{}
}

/*GetBillingReportsBillableusageRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetBillingReportsBillableusageRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetBillingReportsBillableusageRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/billing/reports/billableusage][%d] getBillingReportsBillableusageRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetBillingReportsBillableusageRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetBillingReportsBillableusageRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBillingReportsBillableusageUnsupportedMediaType creates a GetBillingReportsBillableusageUnsupportedMediaType with default headers values
func NewGetBillingReportsBillableusageUnsupportedMediaType() *GetBillingReportsBillableusageUnsupportedMediaType {
	return &GetBillingReportsBillableusageUnsupportedMediaType{}
}

/*GetBillingReportsBillableusageUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetBillingReportsBillableusageUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetBillingReportsBillableusageUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/billing/reports/billableusage][%d] getBillingReportsBillableusageUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetBillingReportsBillableusageUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetBillingReportsBillableusageUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBillingReportsBillableusageTooManyRequests creates a GetBillingReportsBillableusageTooManyRequests with default headers values
func NewGetBillingReportsBillableusageTooManyRequests() *GetBillingReportsBillableusageTooManyRequests {
	return &GetBillingReportsBillableusageTooManyRequests{}
}

/*GetBillingReportsBillableusageTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetBillingReportsBillableusageTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetBillingReportsBillableusageTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/billing/reports/billableusage][%d] getBillingReportsBillableusageTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetBillingReportsBillableusageTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetBillingReportsBillableusageTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBillingReportsBillableusageInternalServerError creates a GetBillingReportsBillableusageInternalServerError with default headers values
func NewGetBillingReportsBillableusageInternalServerError() *GetBillingReportsBillableusageInternalServerError {
	return &GetBillingReportsBillableusageInternalServerError{}
}

/*GetBillingReportsBillableusageInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetBillingReportsBillableusageInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetBillingReportsBillableusageInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/billing/reports/billableusage][%d] getBillingReportsBillableusageInternalServerError  %+v", 500, o.Payload)
}

func (o *GetBillingReportsBillableusageInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetBillingReportsBillableusageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBillingReportsBillableusageServiceUnavailable creates a GetBillingReportsBillableusageServiceUnavailable with default headers values
func NewGetBillingReportsBillableusageServiceUnavailable() *GetBillingReportsBillableusageServiceUnavailable {
	return &GetBillingReportsBillableusageServiceUnavailable{}
}

/*GetBillingReportsBillableusageServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetBillingReportsBillableusageServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetBillingReportsBillableusageServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/billing/reports/billableusage][%d] getBillingReportsBillableusageServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetBillingReportsBillableusageServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetBillingReportsBillableusageServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBillingReportsBillableusageGatewayTimeout creates a GetBillingReportsBillableusageGatewayTimeout with default headers values
func NewGetBillingReportsBillableusageGatewayTimeout() *GetBillingReportsBillableusageGatewayTimeout {
	return &GetBillingReportsBillableusageGatewayTimeout{}
}

/*GetBillingReportsBillableusageGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetBillingReportsBillableusageGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetBillingReportsBillableusageGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/billing/reports/billableusage][%d] getBillingReportsBillableusageGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetBillingReportsBillableusageGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetBillingReportsBillableusageGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

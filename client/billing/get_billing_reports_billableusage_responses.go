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
	case 408:
		result := NewGetBillingReportsBillableusageRequestTimeout()
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

/*
GetBillingReportsBillableusageOK describes a response with status code 200, with default header values.

successful operation
*/
type GetBillingReportsBillableusageOK struct {
	Payload *models.BillingUsageReport
}

// IsSuccess returns true when this get billing reports billableusage o k response has a 2xx status code
func (o *GetBillingReportsBillableusageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get billing reports billableusage o k response has a 3xx status code
func (o *GetBillingReportsBillableusageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get billing reports billableusage o k response has a 4xx status code
func (o *GetBillingReportsBillableusageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get billing reports billableusage o k response has a 5xx status code
func (o *GetBillingReportsBillableusageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get billing reports billableusage o k response a status code equal to that given
func (o *GetBillingReportsBillableusageOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetBillingReportsBillableusageOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/billing/reports/billableusage][%d] getBillingReportsBillableusageOK  %+v", 200, o.Payload)
}

func (o *GetBillingReportsBillableusageOK) String() string {
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

/*
GetBillingReportsBillableusageBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetBillingReportsBillableusageBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get billing reports billableusage bad request response has a 2xx status code
func (o *GetBillingReportsBillableusageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get billing reports billableusage bad request response has a 3xx status code
func (o *GetBillingReportsBillableusageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get billing reports billableusage bad request response has a 4xx status code
func (o *GetBillingReportsBillableusageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get billing reports billableusage bad request response has a 5xx status code
func (o *GetBillingReportsBillableusageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get billing reports billableusage bad request response a status code equal to that given
func (o *GetBillingReportsBillableusageBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetBillingReportsBillableusageBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/billing/reports/billableusage][%d] getBillingReportsBillableusageBadRequest  %+v", 400, o.Payload)
}

func (o *GetBillingReportsBillableusageBadRequest) String() string {
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

/*
GetBillingReportsBillableusageUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetBillingReportsBillableusageUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get billing reports billableusage unauthorized response has a 2xx status code
func (o *GetBillingReportsBillableusageUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get billing reports billableusage unauthorized response has a 3xx status code
func (o *GetBillingReportsBillableusageUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get billing reports billableusage unauthorized response has a 4xx status code
func (o *GetBillingReportsBillableusageUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get billing reports billableusage unauthorized response has a 5xx status code
func (o *GetBillingReportsBillableusageUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get billing reports billableusage unauthorized response a status code equal to that given
func (o *GetBillingReportsBillableusageUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetBillingReportsBillableusageUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/billing/reports/billableusage][%d] getBillingReportsBillableusageUnauthorized  %+v", 401, o.Payload)
}

func (o *GetBillingReportsBillableusageUnauthorized) String() string {
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

/*
GetBillingReportsBillableusageForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetBillingReportsBillableusageForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get billing reports billableusage forbidden response has a 2xx status code
func (o *GetBillingReportsBillableusageForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get billing reports billableusage forbidden response has a 3xx status code
func (o *GetBillingReportsBillableusageForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get billing reports billableusage forbidden response has a 4xx status code
func (o *GetBillingReportsBillableusageForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get billing reports billableusage forbidden response has a 5xx status code
func (o *GetBillingReportsBillableusageForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get billing reports billableusage forbidden response a status code equal to that given
func (o *GetBillingReportsBillableusageForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetBillingReportsBillableusageForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/billing/reports/billableusage][%d] getBillingReportsBillableusageForbidden  %+v", 403, o.Payload)
}

func (o *GetBillingReportsBillableusageForbidden) String() string {
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

/*
GetBillingReportsBillableusageNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetBillingReportsBillableusageNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get billing reports billableusage not found response has a 2xx status code
func (o *GetBillingReportsBillableusageNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get billing reports billableusage not found response has a 3xx status code
func (o *GetBillingReportsBillableusageNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get billing reports billableusage not found response has a 4xx status code
func (o *GetBillingReportsBillableusageNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get billing reports billableusage not found response has a 5xx status code
func (o *GetBillingReportsBillableusageNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get billing reports billableusage not found response a status code equal to that given
func (o *GetBillingReportsBillableusageNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetBillingReportsBillableusageNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/billing/reports/billableusage][%d] getBillingReportsBillableusageNotFound  %+v", 404, o.Payload)
}

func (o *GetBillingReportsBillableusageNotFound) String() string {
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

// NewGetBillingReportsBillableusageRequestTimeout creates a GetBillingReportsBillableusageRequestTimeout with default headers values
func NewGetBillingReportsBillableusageRequestTimeout() *GetBillingReportsBillableusageRequestTimeout {
	return &GetBillingReportsBillableusageRequestTimeout{}
}

/*
GetBillingReportsBillableusageRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetBillingReportsBillableusageRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get billing reports billableusage request timeout response has a 2xx status code
func (o *GetBillingReportsBillableusageRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get billing reports billableusage request timeout response has a 3xx status code
func (o *GetBillingReportsBillableusageRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get billing reports billableusage request timeout response has a 4xx status code
func (o *GetBillingReportsBillableusageRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get billing reports billableusage request timeout response has a 5xx status code
func (o *GetBillingReportsBillableusageRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get billing reports billableusage request timeout response a status code equal to that given
func (o *GetBillingReportsBillableusageRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetBillingReportsBillableusageRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/billing/reports/billableusage][%d] getBillingReportsBillableusageRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetBillingReportsBillableusageRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/billing/reports/billableusage][%d] getBillingReportsBillableusageRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetBillingReportsBillableusageRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetBillingReportsBillableusageRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*
GetBillingReportsBillableusageRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetBillingReportsBillableusageRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get billing reports billableusage request entity too large response has a 2xx status code
func (o *GetBillingReportsBillableusageRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get billing reports billableusage request entity too large response has a 3xx status code
func (o *GetBillingReportsBillableusageRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get billing reports billableusage request entity too large response has a 4xx status code
func (o *GetBillingReportsBillableusageRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get billing reports billableusage request entity too large response has a 5xx status code
func (o *GetBillingReportsBillableusageRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get billing reports billableusage request entity too large response a status code equal to that given
func (o *GetBillingReportsBillableusageRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetBillingReportsBillableusageRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/billing/reports/billableusage][%d] getBillingReportsBillableusageRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetBillingReportsBillableusageRequestEntityTooLarge) String() string {
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

/*
GetBillingReportsBillableusageUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetBillingReportsBillableusageUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get billing reports billableusage unsupported media type response has a 2xx status code
func (o *GetBillingReportsBillableusageUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get billing reports billableusage unsupported media type response has a 3xx status code
func (o *GetBillingReportsBillableusageUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get billing reports billableusage unsupported media type response has a 4xx status code
func (o *GetBillingReportsBillableusageUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get billing reports billableusage unsupported media type response has a 5xx status code
func (o *GetBillingReportsBillableusageUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get billing reports billableusage unsupported media type response a status code equal to that given
func (o *GetBillingReportsBillableusageUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetBillingReportsBillableusageUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/billing/reports/billableusage][%d] getBillingReportsBillableusageUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetBillingReportsBillableusageUnsupportedMediaType) String() string {
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

/*
GetBillingReportsBillableusageTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetBillingReportsBillableusageTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get billing reports billableusage too many requests response has a 2xx status code
func (o *GetBillingReportsBillableusageTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get billing reports billableusage too many requests response has a 3xx status code
func (o *GetBillingReportsBillableusageTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get billing reports billableusage too many requests response has a 4xx status code
func (o *GetBillingReportsBillableusageTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get billing reports billableusage too many requests response has a 5xx status code
func (o *GetBillingReportsBillableusageTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get billing reports billableusage too many requests response a status code equal to that given
func (o *GetBillingReportsBillableusageTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetBillingReportsBillableusageTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/billing/reports/billableusage][%d] getBillingReportsBillableusageTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetBillingReportsBillableusageTooManyRequests) String() string {
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

/*
GetBillingReportsBillableusageInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetBillingReportsBillableusageInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get billing reports billableusage internal server error response has a 2xx status code
func (o *GetBillingReportsBillableusageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get billing reports billableusage internal server error response has a 3xx status code
func (o *GetBillingReportsBillableusageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get billing reports billableusage internal server error response has a 4xx status code
func (o *GetBillingReportsBillableusageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get billing reports billableusage internal server error response has a 5xx status code
func (o *GetBillingReportsBillableusageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get billing reports billableusage internal server error response a status code equal to that given
func (o *GetBillingReportsBillableusageInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetBillingReportsBillableusageInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/billing/reports/billableusage][%d] getBillingReportsBillableusageInternalServerError  %+v", 500, o.Payload)
}

func (o *GetBillingReportsBillableusageInternalServerError) String() string {
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

/*
GetBillingReportsBillableusageServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetBillingReportsBillableusageServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get billing reports billableusage service unavailable response has a 2xx status code
func (o *GetBillingReportsBillableusageServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get billing reports billableusage service unavailable response has a 3xx status code
func (o *GetBillingReportsBillableusageServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get billing reports billableusage service unavailable response has a 4xx status code
func (o *GetBillingReportsBillableusageServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get billing reports billableusage service unavailable response has a 5xx status code
func (o *GetBillingReportsBillableusageServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get billing reports billableusage service unavailable response a status code equal to that given
func (o *GetBillingReportsBillableusageServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetBillingReportsBillableusageServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/billing/reports/billableusage][%d] getBillingReportsBillableusageServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetBillingReportsBillableusageServiceUnavailable) String() string {
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

/*
GetBillingReportsBillableusageGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetBillingReportsBillableusageGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get billing reports billableusage gateway timeout response has a 2xx status code
func (o *GetBillingReportsBillableusageGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get billing reports billableusage gateway timeout response has a 3xx status code
func (o *GetBillingReportsBillableusageGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get billing reports billableusage gateway timeout response has a 4xx status code
func (o *GetBillingReportsBillableusageGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get billing reports billableusage gateway timeout response has a 5xx status code
func (o *GetBillingReportsBillableusageGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get billing reports billableusage gateway timeout response a status code equal to that given
func (o *GetBillingReportsBillableusageGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetBillingReportsBillableusageGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/billing/reports/billableusage][%d] getBillingReportsBillableusageGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetBillingReportsBillableusageGatewayTimeout) String() string {
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

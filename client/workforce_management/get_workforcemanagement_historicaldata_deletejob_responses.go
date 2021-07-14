// Code generated by go-swagger; DO NOT EDIT.

package workforce_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetWorkforcemanagementHistoricaldataDeletejobReader is a Reader for the GetWorkforcemanagementHistoricaldataDeletejob structure.
type GetWorkforcemanagementHistoricaldataDeletejobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkforcemanagementHistoricaldataDeletejobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkforcemanagementHistoricaldataDeletejobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkforcemanagementHistoricaldataDeletejobBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWorkforcemanagementHistoricaldataDeletejobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkforcemanagementHistoricaldataDeletejobForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkforcemanagementHistoricaldataDeletejobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetWorkforcemanagementHistoricaldataDeletejobRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWorkforcemanagementHistoricaldataDeletejobRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWorkforcemanagementHistoricaldataDeletejobUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWorkforcemanagementHistoricaldataDeletejobTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkforcemanagementHistoricaldataDeletejobInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWorkforcemanagementHistoricaldataDeletejobServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWorkforcemanagementHistoricaldataDeletejobGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkforcemanagementHistoricaldataDeletejobOK creates a GetWorkforcemanagementHistoricaldataDeletejobOK with default headers values
func NewGetWorkforcemanagementHistoricaldataDeletejobOK() *GetWorkforcemanagementHistoricaldataDeletejobOK {
	return &GetWorkforcemanagementHistoricaldataDeletejobOK{}
}

/*GetWorkforcemanagementHistoricaldataDeletejobOK handles this case with default header values.

successful operation
*/
type GetWorkforcemanagementHistoricaldataDeletejobOK struct {
	Payload *models.HistoricalImportDeleteJobResponse
}

func (o *GetWorkforcemanagementHistoricaldataDeletejobOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/historicaldata/deletejob][%d] getWorkforcemanagementHistoricaldataDeletejobOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementHistoricaldataDeletejobOK) GetPayload() *models.HistoricalImportDeleteJobResponse {
	return o.Payload
}

func (o *GetWorkforcemanagementHistoricaldataDeletejobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HistoricalImportDeleteJobResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementHistoricaldataDeletejobBadRequest creates a GetWorkforcemanagementHistoricaldataDeletejobBadRequest with default headers values
func NewGetWorkforcemanagementHistoricaldataDeletejobBadRequest() *GetWorkforcemanagementHistoricaldataDeletejobBadRequest {
	return &GetWorkforcemanagementHistoricaldataDeletejobBadRequest{}
}

/*GetWorkforcemanagementHistoricaldataDeletejobBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWorkforcemanagementHistoricaldataDeletejobBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementHistoricaldataDeletejobBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/historicaldata/deletejob][%d] getWorkforcemanagementHistoricaldataDeletejobBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementHistoricaldataDeletejobBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementHistoricaldataDeletejobBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementHistoricaldataDeletejobUnauthorized creates a GetWorkforcemanagementHistoricaldataDeletejobUnauthorized with default headers values
func NewGetWorkforcemanagementHistoricaldataDeletejobUnauthorized() *GetWorkforcemanagementHistoricaldataDeletejobUnauthorized {
	return &GetWorkforcemanagementHistoricaldataDeletejobUnauthorized{}
}

/*GetWorkforcemanagementHistoricaldataDeletejobUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWorkforcemanagementHistoricaldataDeletejobUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementHistoricaldataDeletejobUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/historicaldata/deletejob][%d] getWorkforcemanagementHistoricaldataDeletejobUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementHistoricaldataDeletejobUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementHistoricaldataDeletejobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementHistoricaldataDeletejobForbidden creates a GetWorkforcemanagementHistoricaldataDeletejobForbidden with default headers values
func NewGetWorkforcemanagementHistoricaldataDeletejobForbidden() *GetWorkforcemanagementHistoricaldataDeletejobForbidden {
	return &GetWorkforcemanagementHistoricaldataDeletejobForbidden{}
}

/*GetWorkforcemanagementHistoricaldataDeletejobForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetWorkforcemanagementHistoricaldataDeletejobForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementHistoricaldataDeletejobForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/historicaldata/deletejob][%d] getWorkforcemanagementHistoricaldataDeletejobForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementHistoricaldataDeletejobForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementHistoricaldataDeletejobForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementHistoricaldataDeletejobNotFound creates a GetWorkforcemanagementHistoricaldataDeletejobNotFound with default headers values
func NewGetWorkforcemanagementHistoricaldataDeletejobNotFound() *GetWorkforcemanagementHistoricaldataDeletejobNotFound {
	return &GetWorkforcemanagementHistoricaldataDeletejobNotFound{}
}

/*GetWorkforcemanagementHistoricaldataDeletejobNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetWorkforcemanagementHistoricaldataDeletejobNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementHistoricaldataDeletejobNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/historicaldata/deletejob][%d] getWorkforcemanagementHistoricaldataDeletejobNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementHistoricaldataDeletejobNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementHistoricaldataDeletejobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementHistoricaldataDeletejobRequestTimeout creates a GetWorkforcemanagementHistoricaldataDeletejobRequestTimeout with default headers values
func NewGetWorkforcemanagementHistoricaldataDeletejobRequestTimeout() *GetWorkforcemanagementHistoricaldataDeletejobRequestTimeout {
	return &GetWorkforcemanagementHistoricaldataDeletejobRequestTimeout{}
}

/*GetWorkforcemanagementHistoricaldataDeletejobRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetWorkforcemanagementHistoricaldataDeletejobRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementHistoricaldataDeletejobRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/historicaldata/deletejob][%d] getWorkforcemanagementHistoricaldataDeletejobRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementHistoricaldataDeletejobRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementHistoricaldataDeletejobRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementHistoricaldataDeletejobRequestEntityTooLarge creates a GetWorkforcemanagementHistoricaldataDeletejobRequestEntityTooLarge with default headers values
func NewGetWorkforcemanagementHistoricaldataDeletejobRequestEntityTooLarge() *GetWorkforcemanagementHistoricaldataDeletejobRequestEntityTooLarge {
	return &GetWorkforcemanagementHistoricaldataDeletejobRequestEntityTooLarge{}
}

/*GetWorkforcemanagementHistoricaldataDeletejobRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetWorkforcemanagementHistoricaldataDeletejobRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementHistoricaldataDeletejobRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/historicaldata/deletejob][%d] getWorkforcemanagementHistoricaldataDeletejobRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementHistoricaldataDeletejobRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementHistoricaldataDeletejobRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementHistoricaldataDeletejobUnsupportedMediaType creates a GetWorkforcemanagementHistoricaldataDeletejobUnsupportedMediaType with default headers values
func NewGetWorkforcemanagementHistoricaldataDeletejobUnsupportedMediaType() *GetWorkforcemanagementHistoricaldataDeletejobUnsupportedMediaType {
	return &GetWorkforcemanagementHistoricaldataDeletejobUnsupportedMediaType{}
}

/*GetWorkforcemanagementHistoricaldataDeletejobUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWorkforcemanagementHistoricaldataDeletejobUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementHistoricaldataDeletejobUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/historicaldata/deletejob][%d] getWorkforcemanagementHistoricaldataDeletejobUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementHistoricaldataDeletejobUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementHistoricaldataDeletejobUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementHistoricaldataDeletejobTooManyRequests creates a GetWorkforcemanagementHistoricaldataDeletejobTooManyRequests with default headers values
func NewGetWorkforcemanagementHistoricaldataDeletejobTooManyRequests() *GetWorkforcemanagementHistoricaldataDeletejobTooManyRequests {
	return &GetWorkforcemanagementHistoricaldataDeletejobTooManyRequests{}
}

/*GetWorkforcemanagementHistoricaldataDeletejobTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetWorkforcemanagementHistoricaldataDeletejobTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementHistoricaldataDeletejobTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/historicaldata/deletejob][%d] getWorkforcemanagementHistoricaldataDeletejobTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementHistoricaldataDeletejobTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementHistoricaldataDeletejobTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementHistoricaldataDeletejobInternalServerError creates a GetWorkforcemanagementHistoricaldataDeletejobInternalServerError with default headers values
func NewGetWorkforcemanagementHistoricaldataDeletejobInternalServerError() *GetWorkforcemanagementHistoricaldataDeletejobInternalServerError {
	return &GetWorkforcemanagementHistoricaldataDeletejobInternalServerError{}
}

/*GetWorkforcemanagementHistoricaldataDeletejobInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWorkforcemanagementHistoricaldataDeletejobInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementHistoricaldataDeletejobInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/historicaldata/deletejob][%d] getWorkforcemanagementHistoricaldataDeletejobInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementHistoricaldataDeletejobInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementHistoricaldataDeletejobInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementHistoricaldataDeletejobServiceUnavailable creates a GetWorkforcemanagementHistoricaldataDeletejobServiceUnavailable with default headers values
func NewGetWorkforcemanagementHistoricaldataDeletejobServiceUnavailable() *GetWorkforcemanagementHistoricaldataDeletejobServiceUnavailable {
	return &GetWorkforcemanagementHistoricaldataDeletejobServiceUnavailable{}
}

/*GetWorkforcemanagementHistoricaldataDeletejobServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWorkforcemanagementHistoricaldataDeletejobServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementHistoricaldataDeletejobServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/historicaldata/deletejob][%d] getWorkforcemanagementHistoricaldataDeletejobServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementHistoricaldataDeletejobServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementHistoricaldataDeletejobServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementHistoricaldataDeletejobGatewayTimeout creates a GetWorkforcemanagementHistoricaldataDeletejobGatewayTimeout with default headers values
func NewGetWorkforcemanagementHistoricaldataDeletejobGatewayTimeout() *GetWorkforcemanagementHistoricaldataDeletejobGatewayTimeout {
	return &GetWorkforcemanagementHistoricaldataDeletejobGatewayTimeout{}
}

/*GetWorkforcemanagementHistoricaldataDeletejobGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetWorkforcemanagementHistoricaldataDeletejobGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementHistoricaldataDeletejobGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/historicaldata/deletejob][%d] getWorkforcemanagementHistoricaldataDeletejobGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementHistoricaldataDeletejobGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementHistoricaldataDeletejobGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
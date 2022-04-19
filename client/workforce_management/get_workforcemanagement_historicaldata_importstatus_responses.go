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

// GetWorkforcemanagementHistoricaldataImportstatusReader is a Reader for the GetWorkforcemanagementHistoricaldataImportstatus structure.
type GetWorkforcemanagementHistoricaldataImportstatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkforcemanagementHistoricaldataImportstatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkforcemanagementHistoricaldataImportstatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkforcemanagementHistoricaldataImportstatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWorkforcemanagementHistoricaldataImportstatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkforcemanagementHistoricaldataImportstatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkforcemanagementHistoricaldataImportstatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetWorkforcemanagementHistoricaldataImportstatusRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWorkforcemanagementHistoricaldataImportstatusRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWorkforcemanagementHistoricaldataImportstatusUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWorkforcemanagementHistoricaldataImportstatusTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkforcemanagementHistoricaldataImportstatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWorkforcemanagementHistoricaldataImportstatusServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWorkforcemanagementHistoricaldataImportstatusGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkforcemanagementHistoricaldataImportstatusOK creates a GetWorkforcemanagementHistoricaldataImportstatusOK with default headers values
func NewGetWorkforcemanagementHistoricaldataImportstatusOK() *GetWorkforcemanagementHistoricaldataImportstatusOK {
	return &GetWorkforcemanagementHistoricaldataImportstatusOK{}
}

/*GetWorkforcemanagementHistoricaldataImportstatusOK handles this case with default header values.

successful operation
*/
type GetWorkforcemanagementHistoricaldataImportstatusOK struct {
	Payload *models.HistoricalImportStatusListing
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/historicaldata/importstatus][%d] getWorkforcemanagementHistoricaldataImportstatusOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusOK) GetPayload() *models.HistoricalImportStatusListing {
	return o.Payload
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HistoricalImportStatusListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementHistoricaldataImportstatusBadRequest creates a GetWorkforcemanagementHistoricaldataImportstatusBadRequest with default headers values
func NewGetWorkforcemanagementHistoricaldataImportstatusBadRequest() *GetWorkforcemanagementHistoricaldataImportstatusBadRequest {
	return &GetWorkforcemanagementHistoricaldataImportstatusBadRequest{}
}

/*GetWorkforcemanagementHistoricaldataImportstatusBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWorkforcemanagementHistoricaldataImportstatusBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/historicaldata/importstatus][%d] getWorkforcemanagementHistoricaldataImportstatusBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementHistoricaldataImportstatusUnauthorized creates a GetWorkforcemanagementHistoricaldataImportstatusUnauthorized with default headers values
func NewGetWorkforcemanagementHistoricaldataImportstatusUnauthorized() *GetWorkforcemanagementHistoricaldataImportstatusUnauthorized {
	return &GetWorkforcemanagementHistoricaldataImportstatusUnauthorized{}
}

/*GetWorkforcemanagementHistoricaldataImportstatusUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWorkforcemanagementHistoricaldataImportstatusUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/historicaldata/importstatus][%d] getWorkforcemanagementHistoricaldataImportstatusUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementHistoricaldataImportstatusForbidden creates a GetWorkforcemanagementHistoricaldataImportstatusForbidden with default headers values
func NewGetWorkforcemanagementHistoricaldataImportstatusForbidden() *GetWorkforcemanagementHistoricaldataImportstatusForbidden {
	return &GetWorkforcemanagementHistoricaldataImportstatusForbidden{}
}

/*GetWorkforcemanagementHistoricaldataImportstatusForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetWorkforcemanagementHistoricaldataImportstatusForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/historicaldata/importstatus][%d] getWorkforcemanagementHistoricaldataImportstatusForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementHistoricaldataImportstatusNotFound creates a GetWorkforcemanagementHistoricaldataImportstatusNotFound with default headers values
func NewGetWorkforcemanagementHistoricaldataImportstatusNotFound() *GetWorkforcemanagementHistoricaldataImportstatusNotFound {
	return &GetWorkforcemanagementHistoricaldataImportstatusNotFound{}
}

/*GetWorkforcemanagementHistoricaldataImportstatusNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetWorkforcemanagementHistoricaldataImportstatusNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/historicaldata/importstatus][%d] getWorkforcemanagementHistoricaldataImportstatusNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementHistoricaldataImportstatusRequestTimeout creates a GetWorkforcemanagementHistoricaldataImportstatusRequestTimeout with default headers values
func NewGetWorkforcemanagementHistoricaldataImportstatusRequestTimeout() *GetWorkforcemanagementHistoricaldataImportstatusRequestTimeout {
	return &GetWorkforcemanagementHistoricaldataImportstatusRequestTimeout{}
}

/*GetWorkforcemanagementHistoricaldataImportstatusRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetWorkforcemanagementHistoricaldataImportstatusRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/historicaldata/importstatus][%d] getWorkforcemanagementHistoricaldataImportstatusRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementHistoricaldataImportstatusRequestEntityTooLarge creates a GetWorkforcemanagementHistoricaldataImportstatusRequestEntityTooLarge with default headers values
func NewGetWorkforcemanagementHistoricaldataImportstatusRequestEntityTooLarge() *GetWorkforcemanagementHistoricaldataImportstatusRequestEntityTooLarge {
	return &GetWorkforcemanagementHistoricaldataImportstatusRequestEntityTooLarge{}
}

/*GetWorkforcemanagementHistoricaldataImportstatusRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetWorkforcemanagementHistoricaldataImportstatusRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/historicaldata/importstatus][%d] getWorkforcemanagementHistoricaldataImportstatusRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementHistoricaldataImportstatusUnsupportedMediaType creates a GetWorkforcemanagementHistoricaldataImportstatusUnsupportedMediaType with default headers values
func NewGetWorkforcemanagementHistoricaldataImportstatusUnsupportedMediaType() *GetWorkforcemanagementHistoricaldataImportstatusUnsupportedMediaType {
	return &GetWorkforcemanagementHistoricaldataImportstatusUnsupportedMediaType{}
}

/*GetWorkforcemanagementHistoricaldataImportstatusUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWorkforcemanagementHistoricaldataImportstatusUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/historicaldata/importstatus][%d] getWorkforcemanagementHistoricaldataImportstatusUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementHistoricaldataImportstatusTooManyRequests creates a GetWorkforcemanagementHistoricaldataImportstatusTooManyRequests with default headers values
func NewGetWorkforcemanagementHistoricaldataImportstatusTooManyRequests() *GetWorkforcemanagementHistoricaldataImportstatusTooManyRequests {
	return &GetWorkforcemanagementHistoricaldataImportstatusTooManyRequests{}
}

/*GetWorkforcemanagementHistoricaldataImportstatusTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetWorkforcemanagementHistoricaldataImportstatusTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/historicaldata/importstatus][%d] getWorkforcemanagementHistoricaldataImportstatusTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementHistoricaldataImportstatusInternalServerError creates a GetWorkforcemanagementHistoricaldataImportstatusInternalServerError with default headers values
func NewGetWorkforcemanagementHistoricaldataImportstatusInternalServerError() *GetWorkforcemanagementHistoricaldataImportstatusInternalServerError {
	return &GetWorkforcemanagementHistoricaldataImportstatusInternalServerError{}
}

/*GetWorkforcemanagementHistoricaldataImportstatusInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWorkforcemanagementHistoricaldataImportstatusInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/historicaldata/importstatus][%d] getWorkforcemanagementHistoricaldataImportstatusInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementHistoricaldataImportstatusServiceUnavailable creates a GetWorkforcemanagementHistoricaldataImportstatusServiceUnavailable with default headers values
func NewGetWorkforcemanagementHistoricaldataImportstatusServiceUnavailable() *GetWorkforcemanagementHistoricaldataImportstatusServiceUnavailable {
	return &GetWorkforcemanagementHistoricaldataImportstatusServiceUnavailable{}
}

/*GetWorkforcemanagementHistoricaldataImportstatusServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWorkforcemanagementHistoricaldataImportstatusServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/historicaldata/importstatus][%d] getWorkforcemanagementHistoricaldataImportstatusServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementHistoricaldataImportstatusGatewayTimeout creates a GetWorkforcemanagementHistoricaldataImportstatusGatewayTimeout with default headers values
func NewGetWorkforcemanagementHistoricaldataImportstatusGatewayTimeout() *GetWorkforcemanagementHistoricaldataImportstatusGatewayTimeout {
	return &GetWorkforcemanagementHistoricaldataImportstatusGatewayTimeout{}
}

/*GetWorkforcemanagementHistoricaldataImportstatusGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetWorkforcemanagementHistoricaldataImportstatusGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/historicaldata/importstatus][%d] getWorkforcemanagementHistoricaldataImportstatusGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

/*
GetWorkforcemanagementHistoricaldataImportstatusOK describes a response with status code 200, with default header values.

successful operation
*/
type GetWorkforcemanagementHistoricaldataImportstatusOK struct {
	Payload *models.HistoricalImportStatusListing
}

// IsSuccess returns true when this get workforcemanagement historicaldata importstatus o k response has a 2xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get workforcemanagement historicaldata importstatus o k response has a 3xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement historicaldata importstatus o k response has a 4xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement historicaldata importstatus o k response has a 5xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement historicaldata importstatus o k response a status code equal to that given
func (o *GetWorkforcemanagementHistoricaldataImportstatusOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/historicaldata/importstatus][%d] getWorkforcemanagementHistoricaldataImportstatusOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusOK) String() string {
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

/*
GetWorkforcemanagementHistoricaldataImportstatusBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWorkforcemanagementHistoricaldataImportstatusBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement historicaldata importstatus bad request response has a 2xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement historicaldata importstatus bad request response has a 3xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement historicaldata importstatus bad request response has a 4xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement historicaldata importstatus bad request response has a 5xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement historicaldata importstatus bad request response a status code equal to that given
func (o *GetWorkforcemanagementHistoricaldataImportstatusBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/historicaldata/importstatus][%d] getWorkforcemanagementHistoricaldataImportstatusBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusBadRequest) String() string {
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

/*
GetWorkforcemanagementHistoricaldataImportstatusUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWorkforcemanagementHistoricaldataImportstatusUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement historicaldata importstatus unauthorized response has a 2xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement historicaldata importstatus unauthorized response has a 3xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement historicaldata importstatus unauthorized response has a 4xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement historicaldata importstatus unauthorized response has a 5xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement historicaldata importstatus unauthorized response a status code equal to that given
func (o *GetWorkforcemanagementHistoricaldataImportstatusUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/historicaldata/importstatus][%d] getWorkforcemanagementHistoricaldataImportstatusUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusUnauthorized) String() string {
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

/*
GetWorkforcemanagementHistoricaldataImportstatusForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetWorkforcemanagementHistoricaldataImportstatusForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement historicaldata importstatus forbidden response has a 2xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement historicaldata importstatus forbidden response has a 3xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement historicaldata importstatus forbidden response has a 4xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement historicaldata importstatus forbidden response has a 5xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement historicaldata importstatus forbidden response a status code equal to that given
func (o *GetWorkforcemanagementHistoricaldataImportstatusForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/historicaldata/importstatus][%d] getWorkforcemanagementHistoricaldataImportstatusForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusForbidden) String() string {
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

/*
GetWorkforcemanagementHistoricaldataImportstatusNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetWorkforcemanagementHistoricaldataImportstatusNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement historicaldata importstatus not found response has a 2xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement historicaldata importstatus not found response has a 3xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement historicaldata importstatus not found response has a 4xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement historicaldata importstatus not found response has a 5xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement historicaldata importstatus not found response a status code equal to that given
func (o *GetWorkforcemanagementHistoricaldataImportstatusNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/historicaldata/importstatus][%d] getWorkforcemanagementHistoricaldataImportstatusNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusNotFound) String() string {
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

/*
GetWorkforcemanagementHistoricaldataImportstatusRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetWorkforcemanagementHistoricaldataImportstatusRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement historicaldata importstatus request timeout response has a 2xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement historicaldata importstatus request timeout response has a 3xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement historicaldata importstatus request timeout response has a 4xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement historicaldata importstatus request timeout response has a 5xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement historicaldata importstatus request timeout response a status code equal to that given
func (o *GetWorkforcemanagementHistoricaldataImportstatusRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/historicaldata/importstatus][%d] getWorkforcemanagementHistoricaldataImportstatusRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusRequestTimeout) String() string {
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

/*
GetWorkforcemanagementHistoricaldataImportstatusRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetWorkforcemanagementHistoricaldataImportstatusRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement historicaldata importstatus request entity too large response has a 2xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement historicaldata importstatus request entity too large response has a 3xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement historicaldata importstatus request entity too large response has a 4xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement historicaldata importstatus request entity too large response has a 5xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement historicaldata importstatus request entity too large response a status code equal to that given
func (o *GetWorkforcemanagementHistoricaldataImportstatusRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/historicaldata/importstatus][%d] getWorkforcemanagementHistoricaldataImportstatusRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusRequestEntityTooLarge) String() string {
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

/*
GetWorkforcemanagementHistoricaldataImportstatusUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWorkforcemanagementHistoricaldataImportstatusUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement historicaldata importstatus unsupported media type response has a 2xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement historicaldata importstatus unsupported media type response has a 3xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement historicaldata importstatus unsupported media type response has a 4xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement historicaldata importstatus unsupported media type response has a 5xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement historicaldata importstatus unsupported media type response a status code equal to that given
func (o *GetWorkforcemanagementHistoricaldataImportstatusUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/historicaldata/importstatus][%d] getWorkforcemanagementHistoricaldataImportstatusUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusUnsupportedMediaType) String() string {
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

/*
GetWorkforcemanagementHistoricaldataImportstatusTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetWorkforcemanagementHistoricaldataImportstatusTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement historicaldata importstatus too many requests response has a 2xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement historicaldata importstatus too many requests response has a 3xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement historicaldata importstatus too many requests response has a 4xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement historicaldata importstatus too many requests response has a 5xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement historicaldata importstatus too many requests response a status code equal to that given
func (o *GetWorkforcemanagementHistoricaldataImportstatusTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/historicaldata/importstatus][%d] getWorkforcemanagementHistoricaldataImportstatusTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusTooManyRequests) String() string {
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

/*
GetWorkforcemanagementHistoricaldataImportstatusInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWorkforcemanagementHistoricaldataImportstatusInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement historicaldata importstatus internal server error response has a 2xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement historicaldata importstatus internal server error response has a 3xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement historicaldata importstatus internal server error response has a 4xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement historicaldata importstatus internal server error response has a 5xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement historicaldata importstatus internal server error response a status code equal to that given
func (o *GetWorkforcemanagementHistoricaldataImportstatusInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/historicaldata/importstatus][%d] getWorkforcemanagementHistoricaldataImportstatusInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusInternalServerError) String() string {
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

/*
GetWorkforcemanagementHistoricaldataImportstatusServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWorkforcemanagementHistoricaldataImportstatusServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement historicaldata importstatus service unavailable response has a 2xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement historicaldata importstatus service unavailable response has a 3xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement historicaldata importstatus service unavailable response has a 4xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement historicaldata importstatus service unavailable response has a 5xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement historicaldata importstatus service unavailable response a status code equal to that given
func (o *GetWorkforcemanagementHistoricaldataImportstatusServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/historicaldata/importstatus][%d] getWorkforcemanagementHistoricaldataImportstatusServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusServiceUnavailable) String() string {
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

/*
GetWorkforcemanagementHistoricaldataImportstatusGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetWorkforcemanagementHistoricaldataImportstatusGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement historicaldata importstatus gateway timeout response has a 2xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement historicaldata importstatus gateway timeout response has a 3xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement historicaldata importstatus gateway timeout response has a 4xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement historicaldata importstatus gateway timeout response has a 5xx status code
func (o *GetWorkforcemanagementHistoricaldataImportstatusGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement historicaldata importstatus gateway timeout response a status code equal to that given
func (o *GetWorkforcemanagementHistoricaldataImportstatusGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/historicaldata/importstatus][%d] getWorkforcemanagementHistoricaldataImportstatusGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementHistoricaldataImportstatusGatewayTimeout) String() string {
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

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

// GetWorkforcemanagementAdherenceHistoricalBulkJobReader is a Reader for the GetWorkforcemanagementAdherenceHistoricalBulkJob structure.
type GetWorkforcemanagementAdherenceHistoricalBulkJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkforcemanagementAdherenceHistoricalBulkJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkforcemanagementAdherenceHistoricalBulkJobBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWorkforcemanagementAdherenceHistoricalBulkJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkforcemanagementAdherenceHistoricalBulkJobForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkforcemanagementAdherenceHistoricalBulkJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetWorkforcemanagementAdherenceHistoricalBulkJobRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWorkforcemanagementAdherenceHistoricalBulkJobRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWorkforcemanagementAdherenceHistoricalBulkJobUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWorkforcemanagementAdherenceHistoricalBulkJobTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkforcemanagementAdherenceHistoricalBulkJobInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWorkforcemanagementAdherenceHistoricalBulkJobServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWorkforcemanagementAdherenceHistoricalBulkJobGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkforcemanagementAdherenceHistoricalBulkJobOK creates a GetWorkforcemanagementAdherenceHistoricalBulkJobOK with default headers values
func NewGetWorkforcemanagementAdherenceHistoricalBulkJobOK() *GetWorkforcemanagementAdherenceHistoricalBulkJobOK {
	return &GetWorkforcemanagementAdherenceHistoricalBulkJobOK{}
}

/*
GetWorkforcemanagementAdherenceHistoricalBulkJobOK describes a response with status code 200, with default header values.

successful operation
*/
type GetWorkforcemanagementAdherenceHistoricalBulkJobOK struct {
	Payload *models.WfmHistoricalAdherenceBulkResponse
}

// IsSuccess returns true when this get workforcemanagement adherence historical bulk job o k response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get workforcemanagement adherence historical bulk job o k response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence historical bulk job o k response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement adherence historical bulk job o k response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adherence historical bulk job o k response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/bulk/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalBulkJobOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobOK) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/bulk/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalBulkJobOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobOK) GetPayload() *models.WfmHistoricalAdherenceBulkResponse {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WfmHistoricalAdherenceBulkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceHistoricalBulkJobBadRequest creates a GetWorkforcemanagementAdherenceHistoricalBulkJobBadRequest with default headers values
func NewGetWorkforcemanagementAdherenceHistoricalBulkJobBadRequest() *GetWorkforcemanagementAdherenceHistoricalBulkJobBadRequest {
	return &GetWorkforcemanagementAdherenceHistoricalBulkJobBadRequest{}
}

/*
GetWorkforcemanagementAdherenceHistoricalBulkJobBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWorkforcemanagementAdherenceHistoricalBulkJobBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence historical bulk job bad request response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence historical bulk job bad request response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence historical bulk job bad request response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement adherence historical bulk job bad request response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adherence historical bulk job bad request response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/bulk/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalBulkJobBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/bulk/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalBulkJobBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceHistoricalBulkJobUnauthorized creates a GetWorkforcemanagementAdherenceHistoricalBulkJobUnauthorized with default headers values
func NewGetWorkforcemanagementAdherenceHistoricalBulkJobUnauthorized() *GetWorkforcemanagementAdherenceHistoricalBulkJobUnauthorized {
	return &GetWorkforcemanagementAdherenceHistoricalBulkJobUnauthorized{}
}

/*
GetWorkforcemanagementAdherenceHistoricalBulkJobUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWorkforcemanagementAdherenceHistoricalBulkJobUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence historical bulk job unauthorized response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence historical bulk job unauthorized response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence historical bulk job unauthorized response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement adherence historical bulk job unauthorized response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adherence historical bulk job unauthorized response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/bulk/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalBulkJobUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/bulk/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalBulkJobUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceHistoricalBulkJobForbidden creates a GetWorkforcemanagementAdherenceHistoricalBulkJobForbidden with default headers values
func NewGetWorkforcemanagementAdherenceHistoricalBulkJobForbidden() *GetWorkforcemanagementAdherenceHistoricalBulkJobForbidden {
	return &GetWorkforcemanagementAdherenceHistoricalBulkJobForbidden{}
}

/*
GetWorkforcemanagementAdherenceHistoricalBulkJobForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetWorkforcemanagementAdherenceHistoricalBulkJobForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence historical bulk job forbidden response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence historical bulk job forbidden response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence historical bulk job forbidden response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement adherence historical bulk job forbidden response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adherence historical bulk job forbidden response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/bulk/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalBulkJobForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/bulk/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalBulkJobForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceHistoricalBulkJobNotFound creates a GetWorkforcemanagementAdherenceHistoricalBulkJobNotFound with default headers values
func NewGetWorkforcemanagementAdherenceHistoricalBulkJobNotFound() *GetWorkforcemanagementAdherenceHistoricalBulkJobNotFound {
	return &GetWorkforcemanagementAdherenceHistoricalBulkJobNotFound{}
}

/*
GetWorkforcemanagementAdherenceHistoricalBulkJobNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetWorkforcemanagementAdherenceHistoricalBulkJobNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence historical bulk job not found response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence historical bulk job not found response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence historical bulk job not found response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement adherence historical bulk job not found response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adherence historical bulk job not found response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/bulk/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalBulkJobNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/bulk/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalBulkJobNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceHistoricalBulkJobRequestTimeout creates a GetWorkforcemanagementAdherenceHistoricalBulkJobRequestTimeout with default headers values
func NewGetWorkforcemanagementAdherenceHistoricalBulkJobRequestTimeout() *GetWorkforcemanagementAdherenceHistoricalBulkJobRequestTimeout {
	return &GetWorkforcemanagementAdherenceHistoricalBulkJobRequestTimeout{}
}

/*
GetWorkforcemanagementAdherenceHistoricalBulkJobRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetWorkforcemanagementAdherenceHistoricalBulkJobRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence historical bulk job request timeout response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence historical bulk job request timeout response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence historical bulk job request timeout response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement adherence historical bulk job request timeout response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adherence historical bulk job request timeout response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/bulk/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalBulkJobRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/bulk/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalBulkJobRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceHistoricalBulkJobRequestEntityTooLarge creates a GetWorkforcemanagementAdherenceHistoricalBulkJobRequestEntityTooLarge with default headers values
func NewGetWorkforcemanagementAdherenceHistoricalBulkJobRequestEntityTooLarge() *GetWorkforcemanagementAdherenceHistoricalBulkJobRequestEntityTooLarge {
	return &GetWorkforcemanagementAdherenceHistoricalBulkJobRequestEntityTooLarge{}
}

/*
GetWorkforcemanagementAdherenceHistoricalBulkJobRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetWorkforcemanagementAdherenceHistoricalBulkJobRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence historical bulk job request entity too large response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence historical bulk job request entity too large response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence historical bulk job request entity too large response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement adherence historical bulk job request entity too large response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adherence historical bulk job request entity too large response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/bulk/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalBulkJobRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/bulk/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalBulkJobRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceHistoricalBulkJobUnsupportedMediaType creates a GetWorkforcemanagementAdherenceHistoricalBulkJobUnsupportedMediaType with default headers values
func NewGetWorkforcemanagementAdherenceHistoricalBulkJobUnsupportedMediaType() *GetWorkforcemanagementAdherenceHistoricalBulkJobUnsupportedMediaType {
	return &GetWorkforcemanagementAdherenceHistoricalBulkJobUnsupportedMediaType{}
}

/*
GetWorkforcemanagementAdherenceHistoricalBulkJobUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWorkforcemanagementAdherenceHistoricalBulkJobUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence historical bulk job unsupported media type response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence historical bulk job unsupported media type response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence historical bulk job unsupported media type response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement adherence historical bulk job unsupported media type response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adherence historical bulk job unsupported media type response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/bulk/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalBulkJobUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/bulk/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalBulkJobUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceHistoricalBulkJobTooManyRequests creates a GetWorkforcemanagementAdherenceHistoricalBulkJobTooManyRequests with default headers values
func NewGetWorkforcemanagementAdherenceHistoricalBulkJobTooManyRequests() *GetWorkforcemanagementAdherenceHistoricalBulkJobTooManyRequests {
	return &GetWorkforcemanagementAdherenceHistoricalBulkJobTooManyRequests{}
}

/*
GetWorkforcemanagementAdherenceHistoricalBulkJobTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetWorkforcemanagementAdherenceHistoricalBulkJobTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence historical bulk job too many requests response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence historical bulk job too many requests response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence historical bulk job too many requests response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement adherence historical bulk job too many requests response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adherence historical bulk job too many requests response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/bulk/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalBulkJobTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/bulk/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalBulkJobTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceHistoricalBulkJobInternalServerError creates a GetWorkforcemanagementAdherenceHistoricalBulkJobInternalServerError with default headers values
func NewGetWorkforcemanagementAdherenceHistoricalBulkJobInternalServerError() *GetWorkforcemanagementAdherenceHistoricalBulkJobInternalServerError {
	return &GetWorkforcemanagementAdherenceHistoricalBulkJobInternalServerError{}
}

/*
GetWorkforcemanagementAdherenceHistoricalBulkJobInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWorkforcemanagementAdherenceHistoricalBulkJobInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence historical bulk job internal server error response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence historical bulk job internal server error response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence historical bulk job internal server error response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement adherence historical bulk job internal server error response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement adherence historical bulk job internal server error response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/bulk/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalBulkJobInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/bulk/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalBulkJobInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceHistoricalBulkJobServiceUnavailable creates a GetWorkforcemanagementAdherenceHistoricalBulkJobServiceUnavailable with default headers values
func NewGetWorkforcemanagementAdherenceHistoricalBulkJobServiceUnavailable() *GetWorkforcemanagementAdherenceHistoricalBulkJobServiceUnavailable {
	return &GetWorkforcemanagementAdherenceHistoricalBulkJobServiceUnavailable{}
}

/*
GetWorkforcemanagementAdherenceHistoricalBulkJobServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWorkforcemanagementAdherenceHistoricalBulkJobServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence historical bulk job service unavailable response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence historical bulk job service unavailable response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence historical bulk job service unavailable response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement adherence historical bulk job service unavailable response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement adherence historical bulk job service unavailable response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/bulk/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalBulkJobServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/bulk/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalBulkJobServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceHistoricalBulkJobGatewayTimeout creates a GetWorkforcemanagementAdherenceHistoricalBulkJobGatewayTimeout with default headers values
func NewGetWorkforcemanagementAdherenceHistoricalBulkJobGatewayTimeout() *GetWorkforcemanagementAdherenceHistoricalBulkJobGatewayTimeout {
	return &GetWorkforcemanagementAdherenceHistoricalBulkJobGatewayTimeout{}
}

/*
GetWorkforcemanagementAdherenceHistoricalBulkJobGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetWorkforcemanagementAdherenceHistoricalBulkJobGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence historical bulk job gateway timeout response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence historical bulk job gateway timeout response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence historical bulk job gateway timeout response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement adherence historical bulk job gateway timeout response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement adherence historical bulk job gateway timeout response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/bulk/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalBulkJobGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/bulk/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalBulkJobGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
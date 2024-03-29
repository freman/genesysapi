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

// GetWorkforcemanagementAdherenceHistoricalJobReader is a Reader for the GetWorkforcemanagementAdherenceHistoricalJob structure.
type GetWorkforcemanagementAdherenceHistoricalJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkforcemanagementAdherenceHistoricalJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkforcemanagementAdherenceHistoricalJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkforcemanagementAdherenceHistoricalJobBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWorkforcemanagementAdherenceHistoricalJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkforcemanagementAdherenceHistoricalJobForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkforcemanagementAdherenceHistoricalJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetWorkforcemanagementAdherenceHistoricalJobRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWorkforcemanagementAdherenceHistoricalJobRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWorkforcemanagementAdherenceHistoricalJobUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWorkforcemanagementAdherenceHistoricalJobTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkforcemanagementAdherenceHistoricalJobInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWorkforcemanagementAdherenceHistoricalJobServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWorkforcemanagementAdherenceHistoricalJobGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkforcemanagementAdherenceHistoricalJobOK creates a GetWorkforcemanagementAdherenceHistoricalJobOK with default headers values
func NewGetWorkforcemanagementAdherenceHistoricalJobOK() *GetWorkforcemanagementAdherenceHistoricalJobOK {
	return &GetWorkforcemanagementAdherenceHistoricalJobOK{}
}

/*
GetWorkforcemanagementAdherenceHistoricalJobOK describes a response with status code 200, with default header values.

Successful Operation
*/
type GetWorkforcemanagementAdherenceHistoricalJobOK struct {
	Payload *models.WfmHistoricalAdherenceResponse
}

// IsSuccess returns true when this get workforcemanagement adherence historical job o k response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get workforcemanagement adherence historical job o k response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence historical job o k response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement adherence historical job o k response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adherence historical job o k response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceHistoricalJobOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalJobOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobOK) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalJobOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobOK) GetPayload() *models.WfmHistoricalAdherenceResponse {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WfmHistoricalAdherenceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceHistoricalJobBadRequest creates a GetWorkforcemanagementAdherenceHistoricalJobBadRequest with default headers values
func NewGetWorkforcemanagementAdherenceHistoricalJobBadRequest() *GetWorkforcemanagementAdherenceHistoricalJobBadRequest {
	return &GetWorkforcemanagementAdherenceHistoricalJobBadRequest{}
}

/*
GetWorkforcemanagementAdherenceHistoricalJobBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWorkforcemanagementAdherenceHistoricalJobBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence historical job bad request response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence historical job bad request response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence historical job bad request response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement adherence historical job bad request response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adherence historical job bad request response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceHistoricalJobBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalJobBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalJobBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceHistoricalJobUnauthorized creates a GetWorkforcemanagementAdherenceHistoricalJobUnauthorized with default headers values
func NewGetWorkforcemanagementAdherenceHistoricalJobUnauthorized() *GetWorkforcemanagementAdherenceHistoricalJobUnauthorized {
	return &GetWorkforcemanagementAdherenceHistoricalJobUnauthorized{}
}

/*
GetWorkforcemanagementAdherenceHistoricalJobUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWorkforcemanagementAdherenceHistoricalJobUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence historical job unauthorized response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence historical job unauthorized response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence historical job unauthorized response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement adherence historical job unauthorized response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adherence historical job unauthorized response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceHistoricalJobUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalJobUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalJobUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceHistoricalJobForbidden creates a GetWorkforcemanagementAdherenceHistoricalJobForbidden with default headers values
func NewGetWorkforcemanagementAdherenceHistoricalJobForbidden() *GetWorkforcemanagementAdherenceHistoricalJobForbidden {
	return &GetWorkforcemanagementAdherenceHistoricalJobForbidden{}
}

/*
GetWorkforcemanagementAdherenceHistoricalJobForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetWorkforcemanagementAdherenceHistoricalJobForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence historical job forbidden response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence historical job forbidden response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence historical job forbidden response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement adherence historical job forbidden response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adherence historical job forbidden response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceHistoricalJobForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalJobForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalJobForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceHistoricalJobNotFound creates a GetWorkforcemanagementAdherenceHistoricalJobNotFound with default headers values
func NewGetWorkforcemanagementAdherenceHistoricalJobNotFound() *GetWorkforcemanagementAdherenceHistoricalJobNotFound {
	return &GetWorkforcemanagementAdherenceHistoricalJobNotFound{}
}

/*
GetWorkforcemanagementAdherenceHistoricalJobNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetWorkforcemanagementAdherenceHistoricalJobNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence historical job not found response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence historical job not found response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence historical job not found response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement adherence historical job not found response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adherence historical job not found response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceHistoricalJobNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalJobNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalJobNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceHistoricalJobRequestTimeout creates a GetWorkforcemanagementAdherenceHistoricalJobRequestTimeout with default headers values
func NewGetWorkforcemanagementAdherenceHistoricalJobRequestTimeout() *GetWorkforcemanagementAdherenceHistoricalJobRequestTimeout {
	return &GetWorkforcemanagementAdherenceHistoricalJobRequestTimeout{}
}

/*
GetWorkforcemanagementAdherenceHistoricalJobRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetWorkforcemanagementAdherenceHistoricalJobRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence historical job request timeout response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence historical job request timeout response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence historical job request timeout response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement adherence historical job request timeout response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adherence historical job request timeout response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceHistoricalJobRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalJobRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalJobRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceHistoricalJobRequestEntityTooLarge creates a GetWorkforcemanagementAdherenceHistoricalJobRequestEntityTooLarge with default headers values
func NewGetWorkforcemanagementAdherenceHistoricalJobRequestEntityTooLarge() *GetWorkforcemanagementAdherenceHistoricalJobRequestEntityTooLarge {
	return &GetWorkforcemanagementAdherenceHistoricalJobRequestEntityTooLarge{}
}

/*
GetWorkforcemanagementAdherenceHistoricalJobRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetWorkforcemanagementAdherenceHistoricalJobRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence historical job request entity too large response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence historical job request entity too large response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence historical job request entity too large response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement adherence historical job request entity too large response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adherence historical job request entity too large response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceHistoricalJobRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalJobRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalJobRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceHistoricalJobUnsupportedMediaType creates a GetWorkforcemanagementAdherenceHistoricalJobUnsupportedMediaType with default headers values
func NewGetWorkforcemanagementAdherenceHistoricalJobUnsupportedMediaType() *GetWorkforcemanagementAdherenceHistoricalJobUnsupportedMediaType {
	return &GetWorkforcemanagementAdherenceHistoricalJobUnsupportedMediaType{}
}

/*
GetWorkforcemanagementAdherenceHistoricalJobUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWorkforcemanagementAdherenceHistoricalJobUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence historical job unsupported media type response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence historical job unsupported media type response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence historical job unsupported media type response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement adherence historical job unsupported media type response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adherence historical job unsupported media type response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceHistoricalJobUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalJobUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalJobUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceHistoricalJobTooManyRequests creates a GetWorkforcemanagementAdherenceHistoricalJobTooManyRequests with default headers values
func NewGetWorkforcemanagementAdherenceHistoricalJobTooManyRequests() *GetWorkforcemanagementAdherenceHistoricalJobTooManyRequests {
	return &GetWorkforcemanagementAdherenceHistoricalJobTooManyRequests{}
}

/*
GetWorkforcemanagementAdherenceHistoricalJobTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetWorkforcemanagementAdherenceHistoricalJobTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence historical job too many requests response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence historical job too many requests response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence historical job too many requests response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement adherence historical job too many requests response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adherence historical job too many requests response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceHistoricalJobTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalJobTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalJobTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceHistoricalJobInternalServerError creates a GetWorkforcemanagementAdherenceHistoricalJobInternalServerError with default headers values
func NewGetWorkforcemanagementAdherenceHistoricalJobInternalServerError() *GetWorkforcemanagementAdherenceHistoricalJobInternalServerError {
	return &GetWorkforcemanagementAdherenceHistoricalJobInternalServerError{}
}

/*
GetWorkforcemanagementAdherenceHistoricalJobInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWorkforcemanagementAdherenceHistoricalJobInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence historical job internal server error response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence historical job internal server error response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence historical job internal server error response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement adherence historical job internal server error response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement adherence historical job internal server error response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceHistoricalJobInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalJobInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalJobInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceHistoricalJobServiceUnavailable creates a GetWorkforcemanagementAdherenceHistoricalJobServiceUnavailable with default headers values
func NewGetWorkforcemanagementAdherenceHistoricalJobServiceUnavailable() *GetWorkforcemanagementAdherenceHistoricalJobServiceUnavailable {
	return &GetWorkforcemanagementAdherenceHistoricalJobServiceUnavailable{}
}

/*
GetWorkforcemanagementAdherenceHistoricalJobServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWorkforcemanagementAdherenceHistoricalJobServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence historical job service unavailable response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence historical job service unavailable response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence historical job service unavailable response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement adherence historical job service unavailable response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement adherence historical job service unavailable response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceHistoricalJobServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalJobServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalJobServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceHistoricalJobGatewayTimeout creates a GetWorkforcemanagementAdherenceHistoricalJobGatewayTimeout with default headers values
func NewGetWorkforcemanagementAdherenceHistoricalJobGatewayTimeout() *GetWorkforcemanagementAdherenceHistoricalJobGatewayTimeout {
	return &GetWorkforcemanagementAdherenceHistoricalJobGatewayTimeout{}
}

/*
GetWorkforcemanagementAdherenceHistoricalJobGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetWorkforcemanagementAdherenceHistoricalJobGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence historical job gateway timeout response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence historical job gateway timeout response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence historical job gateway timeout response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement adherence historical job gateway timeout response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceHistoricalJobGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement adherence historical job gateway timeout response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceHistoricalJobGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalJobGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/historical/jobs/{jobId}][%d] getWorkforcemanagementAdherenceHistoricalJobGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceHistoricalJobGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

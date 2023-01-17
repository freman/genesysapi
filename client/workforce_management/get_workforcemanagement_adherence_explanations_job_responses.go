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

// GetWorkforcemanagementAdherenceExplanationsJobReader is a Reader for the GetWorkforcemanagementAdherenceExplanationsJob structure.
type GetWorkforcemanagementAdherenceExplanationsJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkforcemanagementAdherenceExplanationsJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkforcemanagementAdherenceExplanationsJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkforcemanagementAdherenceExplanationsJobBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWorkforcemanagementAdherenceExplanationsJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkforcemanagementAdherenceExplanationsJobForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkforcemanagementAdherenceExplanationsJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetWorkforcemanagementAdherenceExplanationsJobRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWorkforcemanagementAdherenceExplanationsJobRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWorkforcemanagementAdherenceExplanationsJobUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWorkforcemanagementAdherenceExplanationsJobTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkforcemanagementAdherenceExplanationsJobInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWorkforcemanagementAdherenceExplanationsJobServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWorkforcemanagementAdherenceExplanationsJobGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkforcemanagementAdherenceExplanationsJobOK creates a GetWorkforcemanagementAdherenceExplanationsJobOK with default headers values
func NewGetWorkforcemanagementAdherenceExplanationsJobOK() *GetWorkforcemanagementAdherenceExplanationsJobOK {
	return &GetWorkforcemanagementAdherenceExplanationsJobOK{}
}

/*
GetWorkforcemanagementAdherenceExplanationsJobOK describes a response with status code 200, with default header values.

successful operation
*/
type GetWorkforcemanagementAdherenceExplanationsJobOK struct {
	Payload *models.AdherenceExplanationJob
}

// IsSuccess returns true when this get workforcemanagement adherence explanations job o k response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get workforcemanagement adherence explanations job o k response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence explanations job o k response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement adherence explanations job o k response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adherence explanations job o k response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceExplanationsJobOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/jobs/{jobId}][%d] getWorkforcemanagementAdherenceExplanationsJobOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobOK) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/jobs/{jobId}][%d] getWorkforcemanagementAdherenceExplanationsJobOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobOK) GetPayload() *models.AdherenceExplanationJob {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AdherenceExplanationJob)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceExplanationsJobBadRequest creates a GetWorkforcemanagementAdherenceExplanationsJobBadRequest with default headers values
func NewGetWorkforcemanagementAdherenceExplanationsJobBadRequest() *GetWorkforcemanagementAdherenceExplanationsJobBadRequest {
	return &GetWorkforcemanagementAdherenceExplanationsJobBadRequest{}
}

/*
GetWorkforcemanagementAdherenceExplanationsJobBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWorkforcemanagementAdherenceExplanationsJobBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence explanations job bad request response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence explanations job bad request response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence explanations job bad request response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement adherence explanations job bad request response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adherence explanations job bad request response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceExplanationsJobBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/jobs/{jobId}][%d] getWorkforcemanagementAdherenceExplanationsJobBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/jobs/{jobId}][%d] getWorkforcemanagementAdherenceExplanationsJobBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceExplanationsJobUnauthorized creates a GetWorkforcemanagementAdherenceExplanationsJobUnauthorized with default headers values
func NewGetWorkforcemanagementAdherenceExplanationsJobUnauthorized() *GetWorkforcemanagementAdherenceExplanationsJobUnauthorized {
	return &GetWorkforcemanagementAdherenceExplanationsJobUnauthorized{}
}

/*
GetWorkforcemanagementAdherenceExplanationsJobUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWorkforcemanagementAdherenceExplanationsJobUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence explanations job unauthorized response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence explanations job unauthorized response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence explanations job unauthorized response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement adherence explanations job unauthorized response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adherence explanations job unauthorized response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceExplanationsJobUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/jobs/{jobId}][%d] getWorkforcemanagementAdherenceExplanationsJobUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/jobs/{jobId}][%d] getWorkforcemanagementAdherenceExplanationsJobUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceExplanationsJobForbidden creates a GetWorkforcemanagementAdherenceExplanationsJobForbidden with default headers values
func NewGetWorkforcemanagementAdherenceExplanationsJobForbidden() *GetWorkforcemanagementAdherenceExplanationsJobForbidden {
	return &GetWorkforcemanagementAdherenceExplanationsJobForbidden{}
}

/*
GetWorkforcemanagementAdherenceExplanationsJobForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetWorkforcemanagementAdherenceExplanationsJobForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence explanations job forbidden response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence explanations job forbidden response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence explanations job forbidden response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement adherence explanations job forbidden response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adherence explanations job forbidden response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceExplanationsJobForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/jobs/{jobId}][%d] getWorkforcemanagementAdherenceExplanationsJobForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/jobs/{jobId}][%d] getWorkforcemanagementAdherenceExplanationsJobForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceExplanationsJobNotFound creates a GetWorkforcemanagementAdherenceExplanationsJobNotFound with default headers values
func NewGetWorkforcemanagementAdherenceExplanationsJobNotFound() *GetWorkforcemanagementAdherenceExplanationsJobNotFound {
	return &GetWorkforcemanagementAdherenceExplanationsJobNotFound{}
}

/*
GetWorkforcemanagementAdherenceExplanationsJobNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetWorkforcemanagementAdherenceExplanationsJobNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence explanations job not found response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence explanations job not found response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence explanations job not found response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement adherence explanations job not found response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adherence explanations job not found response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceExplanationsJobNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/jobs/{jobId}][%d] getWorkforcemanagementAdherenceExplanationsJobNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/jobs/{jobId}][%d] getWorkforcemanagementAdherenceExplanationsJobNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceExplanationsJobRequestTimeout creates a GetWorkforcemanagementAdherenceExplanationsJobRequestTimeout with default headers values
func NewGetWorkforcemanagementAdherenceExplanationsJobRequestTimeout() *GetWorkforcemanagementAdherenceExplanationsJobRequestTimeout {
	return &GetWorkforcemanagementAdherenceExplanationsJobRequestTimeout{}
}

/*
GetWorkforcemanagementAdherenceExplanationsJobRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetWorkforcemanagementAdherenceExplanationsJobRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence explanations job request timeout response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence explanations job request timeout response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence explanations job request timeout response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement adherence explanations job request timeout response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adherence explanations job request timeout response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceExplanationsJobRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/jobs/{jobId}][%d] getWorkforcemanagementAdherenceExplanationsJobRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/jobs/{jobId}][%d] getWorkforcemanagementAdherenceExplanationsJobRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceExplanationsJobRequestEntityTooLarge creates a GetWorkforcemanagementAdherenceExplanationsJobRequestEntityTooLarge with default headers values
func NewGetWorkforcemanagementAdherenceExplanationsJobRequestEntityTooLarge() *GetWorkforcemanagementAdherenceExplanationsJobRequestEntityTooLarge {
	return &GetWorkforcemanagementAdherenceExplanationsJobRequestEntityTooLarge{}
}

/*
GetWorkforcemanagementAdherenceExplanationsJobRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetWorkforcemanagementAdherenceExplanationsJobRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence explanations job request entity too large response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence explanations job request entity too large response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence explanations job request entity too large response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement adherence explanations job request entity too large response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adherence explanations job request entity too large response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceExplanationsJobRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/jobs/{jobId}][%d] getWorkforcemanagementAdherenceExplanationsJobRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/jobs/{jobId}][%d] getWorkforcemanagementAdherenceExplanationsJobRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceExplanationsJobUnsupportedMediaType creates a GetWorkforcemanagementAdherenceExplanationsJobUnsupportedMediaType with default headers values
func NewGetWorkforcemanagementAdherenceExplanationsJobUnsupportedMediaType() *GetWorkforcemanagementAdherenceExplanationsJobUnsupportedMediaType {
	return &GetWorkforcemanagementAdherenceExplanationsJobUnsupportedMediaType{}
}

/*
GetWorkforcemanagementAdherenceExplanationsJobUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWorkforcemanagementAdherenceExplanationsJobUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence explanations job unsupported media type response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence explanations job unsupported media type response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence explanations job unsupported media type response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement adherence explanations job unsupported media type response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adherence explanations job unsupported media type response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceExplanationsJobUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/jobs/{jobId}][%d] getWorkforcemanagementAdherenceExplanationsJobUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/jobs/{jobId}][%d] getWorkforcemanagementAdherenceExplanationsJobUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceExplanationsJobTooManyRequests creates a GetWorkforcemanagementAdherenceExplanationsJobTooManyRequests with default headers values
func NewGetWorkforcemanagementAdherenceExplanationsJobTooManyRequests() *GetWorkforcemanagementAdherenceExplanationsJobTooManyRequests {
	return &GetWorkforcemanagementAdherenceExplanationsJobTooManyRequests{}
}

/*
GetWorkforcemanagementAdherenceExplanationsJobTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetWorkforcemanagementAdherenceExplanationsJobTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence explanations job too many requests response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence explanations job too many requests response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence explanations job too many requests response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement adherence explanations job too many requests response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adherence explanations job too many requests response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceExplanationsJobTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/jobs/{jobId}][%d] getWorkforcemanagementAdherenceExplanationsJobTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/jobs/{jobId}][%d] getWorkforcemanagementAdherenceExplanationsJobTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceExplanationsJobInternalServerError creates a GetWorkforcemanagementAdherenceExplanationsJobInternalServerError with default headers values
func NewGetWorkforcemanagementAdherenceExplanationsJobInternalServerError() *GetWorkforcemanagementAdherenceExplanationsJobInternalServerError {
	return &GetWorkforcemanagementAdherenceExplanationsJobInternalServerError{}
}

/*
GetWorkforcemanagementAdherenceExplanationsJobInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWorkforcemanagementAdherenceExplanationsJobInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence explanations job internal server error response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence explanations job internal server error response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence explanations job internal server error response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement adherence explanations job internal server error response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement adherence explanations job internal server error response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceExplanationsJobInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/jobs/{jobId}][%d] getWorkforcemanagementAdherenceExplanationsJobInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/jobs/{jobId}][%d] getWorkforcemanagementAdherenceExplanationsJobInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceExplanationsJobServiceUnavailable creates a GetWorkforcemanagementAdherenceExplanationsJobServiceUnavailable with default headers values
func NewGetWorkforcemanagementAdherenceExplanationsJobServiceUnavailable() *GetWorkforcemanagementAdherenceExplanationsJobServiceUnavailable {
	return &GetWorkforcemanagementAdherenceExplanationsJobServiceUnavailable{}
}

/*
GetWorkforcemanagementAdherenceExplanationsJobServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWorkforcemanagementAdherenceExplanationsJobServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence explanations job service unavailable response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence explanations job service unavailable response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence explanations job service unavailable response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement adherence explanations job service unavailable response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement adherence explanations job service unavailable response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceExplanationsJobServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/jobs/{jobId}][%d] getWorkforcemanagementAdherenceExplanationsJobServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/jobs/{jobId}][%d] getWorkforcemanagementAdherenceExplanationsJobServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceExplanationsJobGatewayTimeout creates a GetWorkforcemanagementAdherenceExplanationsJobGatewayTimeout with default headers values
func NewGetWorkforcemanagementAdherenceExplanationsJobGatewayTimeout() *GetWorkforcemanagementAdherenceExplanationsJobGatewayTimeout {
	return &GetWorkforcemanagementAdherenceExplanationsJobGatewayTimeout{}
}

/*
GetWorkforcemanagementAdherenceExplanationsJobGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetWorkforcemanagementAdherenceExplanationsJobGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence explanations job gateway timeout response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence explanations job gateway timeout response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence explanations job gateway timeout response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement adherence explanations job gateway timeout response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceExplanationsJobGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement adherence explanations job gateway timeout response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceExplanationsJobGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/jobs/{jobId}][%d] getWorkforcemanagementAdherenceExplanationsJobGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/jobs/{jobId}][%d] getWorkforcemanagementAdherenceExplanationsJobGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceExplanationsJobGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
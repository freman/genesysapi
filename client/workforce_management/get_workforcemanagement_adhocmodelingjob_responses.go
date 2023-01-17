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

// GetWorkforcemanagementAdhocmodelingjobReader is a Reader for the GetWorkforcemanagementAdhocmodelingjob structure.
type GetWorkforcemanagementAdhocmodelingjobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkforcemanagementAdhocmodelingjobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkforcemanagementAdhocmodelingjobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkforcemanagementAdhocmodelingjobBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWorkforcemanagementAdhocmodelingjobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkforcemanagementAdhocmodelingjobForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkforcemanagementAdhocmodelingjobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetWorkforcemanagementAdhocmodelingjobRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWorkforcemanagementAdhocmodelingjobRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWorkforcemanagementAdhocmodelingjobUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWorkforcemanagementAdhocmodelingjobTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkforcemanagementAdhocmodelingjobInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWorkforcemanagementAdhocmodelingjobServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWorkforcemanagementAdhocmodelingjobGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkforcemanagementAdhocmodelingjobOK creates a GetWorkforcemanagementAdhocmodelingjobOK with default headers values
func NewGetWorkforcemanagementAdhocmodelingjobOK() *GetWorkforcemanagementAdhocmodelingjobOK {
	return &GetWorkforcemanagementAdhocmodelingjobOK{}
}

/*
GetWorkforcemanagementAdhocmodelingjobOK describes a response with status code 200, with default header values.

successful operation
*/
type GetWorkforcemanagementAdhocmodelingjobOK struct {
	Payload *models.ModelingStatusResponse
}

// IsSuccess returns true when this get workforcemanagement adhocmodelingjob o k response has a 2xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get workforcemanagement adhocmodelingjob o k response has a 3xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adhocmodelingjob o k response has a 4xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement adhocmodelingjob o k response has a 5xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adhocmodelingjob o k response a status code equal to that given
func (o *GetWorkforcemanagementAdhocmodelingjobOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetWorkforcemanagementAdhocmodelingjobOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adhocmodelingjobs/{jobId}][%d] getWorkforcemanagementAdhocmodelingjobOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementAdhocmodelingjobOK) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adhocmodelingjobs/{jobId}][%d] getWorkforcemanagementAdhocmodelingjobOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementAdhocmodelingjobOK) GetPayload() *models.ModelingStatusResponse {
	return o.Payload
}

func (o *GetWorkforcemanagementAdhocmodelingjobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelingStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdhocmodelingjobBadRequest creates a GetWorkforcemanagementAdhocmodelingjobBadRequest with default headers values
func NewGetWorkforcemanagementAdhocmodelingjobBadRequest() *GetWorkforcemanagementAdhocmodelingjobBadRequest {
	return &GetWorkforcemanagementAdhocmodelingjobBadRequest{}
}

/*
GetWorkforcemanagementAdhocmodelingjobBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWorkforcemanagementAdhocmodelingjobBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adhocmodelingjob bad request response has a 2xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adhocmodelingjob bad request response has a 3xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adhocmodelingjob bad request response has a 4xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement adhocmodelingjob bad request response has a 5xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adhocmodelingjob bad request response a status code equal to that given
func (o *GetWorkforcemanagementAdhocmodelingjobBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetWorkforcemanagementAdhocmodelingjobBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adhocmodelingjobs/{jobId}][%d] getWorkforcemanagementAdhocmodelingjobBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementAdhocmodelingjobBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adhocmodelingjobs/{jobId}][%d] getWorkforcemanagementAdhocmodelingjobBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementAdhocmodelingjobBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdhocmodelingjobBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdhocmodelingjobUnauthorized creates a GetWorkforcemanagementAdhocmodelingjobUnauthorized with default headers values
func NewGetWorkforcemanagementAdhocmodelingjobUnauthorized() *GetWorkforcemanagementAdhocmodelingjobUnauthorized {
	return &GetWorkforcemanagementAdhocmodelingjobUnauthorized{}
}

/*
GetWorkforcemanagementAdhocmodelingjobUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWorkforcemanagementAdhocmodelingjobUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adhocmodelingjob unauthorized response has a 2xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adhocmodelingjob unauthorized response has a 3xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adhocmodelingjob unauthorized response has a 4xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement adhocmodelingjob unauthorized response has a 5xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adhocmodelingjob unauthorized response a status code equal to that given
func (o *GetWorkforcemanagementAdhocmodelingjobUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetWorkforcemanagementAdhocmodelingjobUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adhocmodelingjobs/{jobId}][%d] getWorkforcemanagementAdhocmodelingjobUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementAdhocmodelingjobUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adhocmodelingjobs/{jobId}][%d] getWorkforcemanagementAdhocmodelingjobUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementAdhocmodelingjobUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdhocmodelingjobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdhocmodelingjobForbidden creates a GetWorkforcemanagementAdhocmodelingjobForbidden with default headers values
func NewGetWorkforcemanagementAdhocmodelingjobForbidden() *GetWorkforcemanagementAdhocmodelingjobForbidden {
	return &GetWorkforcemanagementAdhocmodelingjobForbidden{}
}

/*
GetWorkforcemanagementAdhocmodelingjobForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetWorkforcemanagementAdhocmodelingjobForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adhocmodelingjob forbidden response has a 2xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adhocmodelingjob forbidden response has a 3xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adhocmodelingjob forbidden response has a 4xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement adhocmodelingjob forbidden response has a 5xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adhocmodelingjob forbidden response a status code equal to that given
func (o *GetWorkforcemanagementAdhocmodelingjobForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetWorkforcemanagementAdhocmodelingjobForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adhocmodelingjobs/{jobId}][%d] getWorkforcemanagementAdhocmodelingjobForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementAdhocmodelingjobForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adhocmodelingjobs/{jobId}][%d] getWorkforcemanagementAdhocmodelingjobForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementAdhocmodelingjobForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdhocmodelingjobForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdhocmodelingjobNotFound creates a GetWorkforcemanagementAdhocmodelingjobNotFound with default headers values
func NewGetWorkforcemanagementAdhocmodelingjobNotFound() *GetWorkforcemanagementAdhocmodelingjobNotFound {
	return &GetWorkforcemanagementAdhocmodelingjobNotFound{}
}

/*
GetWorkforcemanagementAdhocmodelingjobNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetWorkforcemanagementAdhocmodelingjobNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adhocmodelingjob not found response has a 2xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adhocmodelingjob not found response has a 3xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adhocmodelingjob not found response has a 4xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement adhocmodelingjob not found response has a 5xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adhocmodelingjob not found response a status code equal to that given
func (o *GetWorkforcemanagementAdhocmodelingjobNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetWorkforcemanagementAdhocmodelingjobNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adhocmodelingjobs/{jobId}][%d] getWorkforcemanagementAdhocmodelingjobNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementAdhocmodelingjobNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adhocmodelingjobs/{jobId}][%d] getWorkforcemanagementAdhocmodelingjobNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementAdhocmodelingjobNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdhocmodelingjobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdhocmodelingjobRequestTimeout creates a GetWorkforcemanagementAdhocmodelingjobRequestTimeout with default headers values
func NewGetWorkforcemanagementAdhocmodelingjobRequestTimeout() *GetWorkforcemanagementAdhocmodelingjobRequestTimeout {
	return &GetWorkforcemanagementAdhocmodelingjobRequestTimeout{}
}

/*
GetWorkforcemanagementAdhocmodelingjobRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetWorkforcemanagementAdhocmodelingjobRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adhocmodelingjob request timeout response has a 2xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adhocmodelingjob request timeout response has a 3xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adhocmodelingjob request timeout response has a 4xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement adhocmodelingjob request timeout response has a 5xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adhocmodelingjob request timeout response a status code equal to that given
func (o *GetWorkforcemanagementAdhocmodelingjobRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetWorkforcemanagementAdhocmodelingjobRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adhocmodelingjobs/{jobId}][%d] getWorkforcemanagementAdhocmodelingjobRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementAdhocmodelingjobRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adhocmodelingjobs/{jobId}][%d] getWorkforcemanagementAdhocmodelingjobRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementAdhocmodelingjobRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdhocmodelingjobRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdhocmodelingjobRequestEntityTooLarge creates a GetWorkforcemanagementAdhocmodelingjobRequestEntityTooLarge with default headers values
func NewGetWorkforcemanagementAdhocmodelingjobRequestEntityTooLarge() *GetWorkforcemanagementAdhocmodelingjobRequestEntityTooLarge {
	return &GetWorkforcemanagementAdhocmodelingjobRequestEntityTooLarge{}
}

/*
GetWorkforcemanagementAdhocmodelingjobRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetWorkforcemanagementAdhocmodelingjobRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adhocmodelingjob request entity too large response has a 2xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adhocmodelingjob request entity too large response has a 3xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adhocmodelingjob request entity too large response has a 4xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement adhocmodelingjob request entity too large response has a 5xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adhocmodelingjob request entity too large response a status code equal to that given
func (o *GetWorkforcemanagementAdhocmodelingjobRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetWorkforcemanagementAdhocmodelingjobRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adhocmodelingjobs/{jobId}][%d] getWorkforcemanagementAdhocmodelingjobRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementAdhocmodelingjobRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adhocmodelingjobs/{jobId}][%d] getWorkforcemanagementAdhocmodelingjobRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementAdhocmodelingjobRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdhocmodelingjobRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdhocmodelingjobUnsupportedMediaType creates a GetWorkforcemanagementAdhocmodelingjobUnsupportedMediaType with default headers values
func NewGetWorkforcemanagementAdhocmodelingjobUnsupportedMediaType() *GetWorkforcemanagementAdhocmodelingjobUnsupportedMediaType {
	return &GetWorkforcemanagementAdhocmodelingjobUnsupportedMediaType{}
}

/*
GetWorkforcemanagementAdhocmodelingjobUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWorkforcemanagementAdhocmodelingjobUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adhocmodelingjob unsupported media type response has a 2xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adhocmodelingjob unsupported media type response has a 3xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adhocmodelingjob unsupported media type response has a 4xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement adhocmodelingjob unsupported media type response has a 5xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adhocmodelingjob unsupported media type response a status code equal to that given
func (o *GetWorkforcemanagementAdhocmodelingjobUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetWorkforcemanagementAdhocmodelingjobUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adhocmodelingjobs/{jobId}][%d] getWorkforcemanagementAdhocmodelingjobUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementAdhocmodelingjobUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adhocmodelingjobs/{jobId}][%d] getWorkforcemanagementAdhocmodelingjobUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementAdhocmodelingjobUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdhocmodelingjobUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdhocmodelingjobTooManyRequests creates a GetWorkforcemanagementAdhocmodelingjobTooManyRequests with default headers values
func NewGetWorkforcemanagementAdhocmodelingjobTooManyRequests() *GetWorkforcemanagementAdhocmodelingjobTooManyRequests {
	return &GetWorkforcemanagementAdhocmodelingjobTooManyRequests{}
}

/*
GetWorkforcemanagementAdhocmodelingjobTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetWorkforcemanagementAdhocmodelingjobTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adhocmodelingjob too many requests response has a 2xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adhocmodelingjob too many requests response has a 3xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adhocmodelingjob too many requests response has a 4xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement adhocmodelingjob too many requests response has a 5xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adhocmodelingjob too many requests response a status code equal to that given
func (o *GetWorkforcemanagementAdhocmodelingjobTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetWorkforcemanagementAdhocmodelingjobTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adhocmodelingjobs/{jobId}][%d] getWorkforcemanagementAdhocmodelingjobTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementAdhocmodelingjobTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adhocmodelingjobs/{jobId}][%d] getWorkforcemanagementAdhocmodelingjobTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementAdhocmodelingjobTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdhocmodelingjobTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdhocmodelingjobInternalServerError creates a GetWorkforcemanagementAdhocmodelingjobInternalServerError with default headers values
func NewGetWorkforcemanagementAdhocmodelingjobInternalServerError() *GetWorkforcemanagementAdhocmodelingjobInternalServerError {
	return &GetWorkforcemanagementAdhocmodelingjobInternalServerError{}
}

/*
GetWorkforcemanagementAdhocmodelingjobInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWorkforcemanagementAdhocmodelingjobInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adhocmodelingjob internal server error response has a 2xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adhocmodelingjob internal server error response has a 3xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adhocmodelingjob internal server error response has a 4xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement adhocmodelingjob internal server error response has a 5xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement adhocmodelingjob internal server error response a status code equal to that given
func (o *GetWorkforcemanagementAdhocmodelingjobInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetWorkforcemanagementAdhocmodelingjobInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adhocmodelingjobs/{jobId}][%d] getWorkforcemanagementAdhocmodelingjobInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementAdhocmodelingjobInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adhocmodelingjobs/{jobId}][%d] getWorkforcemanagementAdhocmodelingjobInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementAdhocmodelingjobInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdhocmodelingjobInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdhocmodelingjobServiceUnavailable creates a GetWorkforcemanagementAdhocmodelingjobServiceUnavailable with default headers values
func NewGetWorkforcemanagementAdhocmodelingjobServiceUnavailable() *GetWorkforcemanagementAdhocmodelingjobServiceUnavailable {
	return &GetWorkforcemanagementAdhocmodelingjobServiceUnavailable{}
}

/*
GetWorkforcemanagementAdhocmodelingjobServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWorkforcemanagementAdhocmodelingjobServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adhocmodelingjob service unavailable response has a 2xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adhocmodelingjob service unavailable response has a 3xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adhocmodelingjob service unavailable response has a 4xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement adhocmodelingjob service unavailable response has a 5xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement adhocmodelingjob service unavailable response a status code equal to that given
func (o *GetWorkforcemanagementAdhocmodelingjobServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetWorkforcemanagementAdhocmodelingjobServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adhocmodelingjobs/{jobId}][%d] getWorkforcemanagementAdhocmodelingjobServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementAdhocmodelingjobServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adhocmodelingjobs/{jobId}][%d] getWorkforcemanagementAdhocmodelingjobServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementAdhocmodelingjobServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdhocmodelingjobServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdhocmodelingjobGatewayTimeout creates a GetWorkforcemanagementAdhocmodelingjobGatewayTimeout with default headers values
func NewGetWorkforcemanagementAdhocmodelingjobGatewayTimeout() *GetWorkforcemanagementAdhocmodelingjobGatewayTimeout {
	return &GetWorkforcemanagementAdhocmodelingjobGatewayTimeout{}
}

/*
GetWorkforcemanagementAdhocmodelingjobGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetWorkforcemanagementAdhocmodelingjobGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adhocmodelingjob gateway timeout response has a 2xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adhocmodelingjob gateway timeout response has a 3xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adhocmodelingjob gateway timeout response has a 4xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement adhocmodelingjob gateway timeout response has a 5xx status code
func (o *GetWorkforcemanagementAdhocmodelingjobGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement adhocmodelingjob gateway timeout response a status code equal to that given
func (o *GetWorkforcemanagementAdhocmodelingjobGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetWorkforcemanagementAdhocmodelingjobGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adhocmodelingjobs/{jobId}][%d] getWorkforcemanagementAdhocmodelingjobGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementAdhocmodelingjobGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adhocmodelingjobs/{jobId}][%d] getWorkforcemanagementAdhocmodelingjobGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementAdhocmodelingjobGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdhocmodelingjobGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package architect

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetFlowsDatatableExportJobReader is a Reader for the GetFlowsDatatableExportJob structure.
type GetFlowsDatatableExportJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFlowsDatatableExportJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFlowsDatatableExportJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewGetFlowsDatatableExportJobAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFlowsDatatableExportJobBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFlowsDatatableExportJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFlowsDatatableExportJobForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFlowsDatatableExportJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetFlowsDatatableExportJobRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetFlowsDatatableExportJobRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetFlowsDatatableExportJobUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetFlowsDatatableExportJobTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFlowsDatatableExportJobInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetFlowsDatatableExportJobServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetFlowsDatatableExportJobGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFlowsDatatableExportJobOK creates a GetFlowsDatatableExportJobOK with default headers values
func NewGetFlowsDatatableExportJobOK() *GetFlowsDatatableExportJobOK {
	return &GetFlowsDatatableExportJobOK{}
}

/*
GetFlowsDatatableExportJobOK describes a response with status code 200, with default header values.

successful operation
*/
type GetFlowsDatatableExportJobOK struct {
	Payload *models.DataTableExportJob
}

// IsSuccess returns true when this get flows datatable export job o k response has a 2xx status code
func (o *GetFlowsDatatableExportJobOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get flows datatable export job o k response has a 3xx status code
func (o *GetFlowsDatatableExportJobOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows datatable export job o k response has a 4xx status code
func (o *GetFlowsDatatableExportJobOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get flows datatable export job o k response has a 5xx status code
func (o *GetFlowsDatatableExportJobOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get flows datatable export job o k response a status code equal to that given
func (o *GetFlowsDatatableExportJobOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetFlowsDatatableExportJobOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}/export/jobs/{exportJobId}][%d] getFlowsDatatableExportJobOK  %+v", 200, o.Payload)
}

func (o *GetFlowsDatatableExportJobOK) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}/export/jobs/{exportJobId}][%d] getFlowsDatatableExportJobOK  %+v", 200, o.Payload)
}

func (o *GetFlowsDatatableExportJobOK) GetPayload() *models.DataTableExportJob {
	return o.Payload
}

func (o *GetFlowsDatatableExportJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataTableExportJob)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableExportJobAccepted creates a GetFlowsDatatableExportJobAccepted with default headers values
func NewGetFlowsDatatableExportJobAccepted() *GetFlowsDatatableExportJobAccepted {
	return &GetFlowsDatatableExportJobAccepted{}
}

/*
GetFlowsDatatableExportJobAccepted describes a response with status code 202, with default header values.

Request Accepted
*/
type GetFlowsDatatableExportJobAccepted struct {
	Payload *models.DataTableExportJob
}

// IsSuccess returns true when this get flows datatable export job accepted response has a 2xx status code
func (o *GetFlowsDatatableExportJobAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get flows datatable export job accepted response has a 3xx status code
func (o *GetFlowsDatatableExportJobAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows datatable export job accepted response has a 4xx status code
func (o *GetFlowsDatatableExportJobAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this get flows datatable export job accepted response has a 5xx status code
func (o *GetFlowsDatatableExportJobAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this get flows datatable export job accepted response a status code equal to that given
func (o *GetFlowsDatatableExportJobAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *GetFlowsDatatableExportJobAccepted) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}/export/jobs/{exportJobId}][%d] getFlowsDatatableExportJobAccepted  %+v", 202, o.Payload)
}

func (o *GetFlowsDatatableExportJobAccepted) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}/export/jobs/{exportJobId}][%d] getFlowsDatatableExportJobAccepted  %+v", 202, o.Payload)
}

func (o *GetFlowsDatatableExportJobAccepted) GetPayload() *models.DataTableExportJob {
	return o.Payload
}

func (o *GetFlowsDatatableExportJobAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataTableExportJob)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableExportJobBadRequest creates a GetFlowsDatatableExportJobBadRequest with default headers values
func NewGetFlowsDatatableExportJobBadRequest() *GetFlowsDatatableExportJobBadRequest {
	return &GetFlowsDatatableExportJobBadRequest{}
}

/*
GetFlowsDatatableExportJobBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetFlowsDatatableExportJobBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows datatable export job bad request response has a 2xx status code
func (o *GetFlowsDatatableExportJobBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows datatable export job bad request response has a 3xx status code
func (o *GetFlowsDatatableExportJobBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows datatable export job bad request response has a 4xx status code
func (o *GetFlowsDatatableExportJobBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flows datatable export job bad request response has a 5xx status code
func (o *GetFlowsDatatableExportJobBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get flows datatable export job bad request response a status code equal to that given
func (o *GetFlowsDatatableExportJobBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetFlowsDatatableExportJobBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}/export/jobs/{exportJobId}][%d] getFlowsDatatableExportJobBadRequest  %+v", 400, o.Payload)
}

func (o *GetFlowsDatatableExportJobBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}/export/jobs/{exportJobId}][%d] getFlowsDatatableExportJobBadRequest  %+v", 400, o.Payload)
}

func (o *GetFlowsDatatableExportJobBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatableExportJobBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableExportJobUnauthorized creates a GetFlowsDatatableExportJobUnauthorized with default headers values
func NewGetFlowsDatatableExportJobUnauthorized() *GetFlowsDatatableExportJobUnauthorized {
	return &GetFlowsDatatableExportJobUnauthorized{}
}

/*
GetFlowsDatatableExportJobUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetFlowsDatatableExportJobUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows datatable export job unauthorized response has a 2xx status code
func (o *GetFlowsDatatableExportJobUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows datatable export job unauthorized response has a 3xx status code
func (o *GetFlowsDatatableExportJobUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows datatable export job unauthorized response has a 4xx status code
func (o *GetFlowsDatatableExportJobUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flows datatable export job unauthorized response has a 5xx status code
func (o *GetFlowsDatatableExportJobUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get flows datatable export job unauthorized response a status code equal to that given
func (o *GetFlowsDatatableExportJobUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetFlowsDatatableExportJobUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}/export/jobs/{exportJobId}][%d] getFlowsDatatableExportJobUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFlowsDatatableExportJobUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}/export/jobs/{exportJobId}][%d] getFlowsDatatableExportJobUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFlowsDatatableExportJobUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatableExportJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableExportJobForbidden creates a GetFlowsDatatableExportJobForbidden with default headers values
func NewGetFlowsDatatableExportJobForbidden() *GetFlowsDatatableExportJobForbidden {
	return &GetFlowsDatatableExportJobForbidden{}
}

/*
GetFlowsDatatableExportJobForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetFlowsDatatableExportJobForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows datatable export job forbidden response has a 2xx status code
func (o *GetFlowsDatatableExportJobForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows datatable export job forbidden response has a 3xx status code
func (o *GetFlowsDatatableExportJobForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows datatable export job forbidden response has a 4xx status code
func (o *GetFlowsDatatableExportJobForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flows datatable export job forbidden response has a 5xx status code
func (o *GetFlowsDatatableExportJobForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get flows datatable export job forbidden response a status code equal to that given
func (o *GetFlowsDatatableExportJobForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetFlowsDatatableExportJobForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}/export/jobs/{exportJobId}][%d] getFlowsDatatableExportJobForbidden  %+v", 403, o.Payload)
}

func (o *GetFlowsDatatableExportJobForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}/export/jobs/{exportJobId}][%d] getFlowsDatatableExportJobForbidden  %+v", 403, o.Payload)
}

func (o *GetFlowsDatatableExportJobForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatableExportJobForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableExportJobNotFound creates a GetFlowsDatatableExportJobNotFound with default headers values
func NewGetFlowsDatatableExportJobNotFound() *GetFlowsDatatableExportJobNotFound {
	return &GetFlowsDatatableExportJobNotFound{}
}

/*
GetFlowsDatatableExportJobNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetFlowsDatatableExportJobNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows datatable export job not found response has a 2xx status code
func (o *GetFlowsDatatableExportJobNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows datatable export job not found response has a 3xx status code
func (o *GetFlowsDatatableExportJobNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows datatable export job not found response has a 4xx status code
func (o *GetFlowsDatatableExportJobNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flows datatable export job not found response has a 5xx status code
func (o *GetFlowsDatatableExportJobNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get flows datatable export job not found response a status code equal to that given
func (o *GetFlowsDatatableExportJobNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetFlowsDatatableExportJobNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}/export/jobs/{exportJobId}][%d] getFlowsDatatableExportJobNotFound  %+v", 404, o.Payload)
}

func (o *GetFlowsDatatableExportJobNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}/export/jobs/{exportJobId}][%d] getFlowsDatatableExportJobNotFound  %+v", 404, o.Payload)
}

func (o *GetFlowsDatatableExportJobNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatableExportJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableExportJobRequestTimeout creates a GetFlowsDatatableExportJobRequestTimeout with default headers values
func NewGetFlowsDatatableExportJobRequestTimeout() *GetFlowsDatatableExportJobRequestTimeout {
	return &GetFlowsDatatableExportJobRequestTimeout{}
}

/*
GetFlowsDatatableExportJobRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetFlowsDatatableExportJobRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows datatable export job request timeout response has a 2xx status code
func (o *GetFlowsDatatableExportJobRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows datatable export job request timeout response has a 3xx status code
func (o *GetFlowsDatatableExportJobRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows datatable export job request timeout response has a 4xx status code
func (o *GetFlowsDatatableExportJobRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flows datatable export job request timeout response has a 5xx status code
func (o *GetFlowsDatatableExportJobRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get flows datatable export job request timeout response a status code equal to that given
func (o *GetFlowsDatatableExportJobRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetFlowsDatatableExportJobRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}/export/jobs/{exportJobId}][%d] getFlowsDatatableExportJobRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetFlowsDatatableExportJobRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}/export/jobs/{exportJobId}][%d] getFlowsDatatableExportJobRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetFlowsDatatableExportJobRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatableExportJobRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableExportJobRequestEntityTooLarge creates a GetFlowsDatatableExportJobRequestEntityTooLarge with default headers values
func NewGetFlowsDatatableExportJobRequestEntityTooLarge() *GetFlowsDatatableExportJobRequestEntityTooLarge {
	return &GetFlowsDatatableExportJobRequestEntityTooLarge{}
}

/*
GetFlowsDatatableExportJobRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetFlowsDatatableExportJobRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows datatable export job request entity too large response has a 2xx status code
func (o *GetFlowsDatatableExportJobRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows datatable export job request entity too large response has a 3xx status code
func (o *GetFlowsDatatableExportJobRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows datatable export job request entity too large response has a 4xx status code
func (o *GetFlowsDatatableExportJobRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flows datatable export job request entity too large response has a 5xx status code
func (o *GetFlowsDatatableExportJobRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get flows datatable export job request entity too large response a status code equal to that given
func (o *GetFlowsDatatableExportJobRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetFlowsDatatableExportJobRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}/export/jobs/{exportJobId}][%d] getFlowsDatatableExportJobRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetFlowsDatatableExportJobRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}/export/jobs/{exportJobId}][%d] getFlowsDatatableExportJobRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetFlowsDatatableExportJobRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatableExportJobRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableExportJobUnsupportedMediaType creates a GetFlowsDatatableExportJobUnsupportedMediaType with default headers values
func NewGetFlowsDatatableExportJobUnsupportedMediaType() *GetFlowsDatatableExportJobUnsupportedMediaType {
	return &GetFlowsDatatableExportJobUnsupportedMediaType{}
}

/*
GetFlowsDatatableExportJobUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetFlowsDatatableExportJobUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows datatable export job unsupported media type response has a 2xx status code
func (o *GetFlowsDatatableExportJobUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows datatable export job unsupported media type response has a 3xx status code
func (o *GetFlowsDatatableExportJobUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows datatable export job unsupported media type response has a 4xx status code
func (o *GetFlowsDatatableExportJobUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flows datatable export job unsupported media type response has a 5xx status code
func (o *GetFlowsDatatableExportJobUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get flows datatable export job unsupported media type response a status code equal to that given
func (o *GetFlowsDatatableExportJobUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetFlowsDatatableExportJobUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}/export/jobs/{exportJobId}][%d] getFlowsDatatableExportJobUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetFlowsDatatableExportJobUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}/export/jobs/{exportJobId}][%d] getFlowsDatatableExportJobUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetFlowsDatatableExportJobUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatableExportJobUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableExportJobTooManyRequests creates a GetFlowsDatatableExportJobTooManyRequests with default headers values
func NewGetFlowsDatatableExportJobTooManyRequests() *GetFlowsDatatableExportJobTooManyRequests {
	return &GetFlowsDatatableExportJobTooManyRequests{}
}

/*
GetFlowsDatatableExportJobTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetFlowsDatatableExportJobTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows datatable export job too many requests response has a 2xx status code
func (o *GetFlowsDatatableExportJobTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows datatable export job too many requests response has a 3xx status code
func (o *GetFlowsDatatableExportJobTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows datatable export job too many requests response has a 4xx status code
func (o *GetFlowsDatatableExportJobTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flows datatable export job too many requests response has a 5xx status code
func (o *GetFlowsDatatableExportJobTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get flows datatable export job too many requests response a status code equal to that given
func (o *GetFlowsDatatableExportJobTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetFlowsDatatableExportJobTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}/export/jobs/{exportJobId}][%d] getFlowsDatatableExportJobTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetFlowsDatatableExportJobTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}/export/jobs/{exportJobId}][%d] getFlowsDatatableExportJobTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetFlowsDatatableExportJobTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatableExportJobTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableExportJobInternalServerError creates a GetFlowsDatatableExportJobInternalServerError with default headers values
func NewGetFlowsDatatableExportJobInternalServerError() *GetFlowsDatatableExportJobInternalServerError {
	return &GetFlowsDatatableExportJobInternalServerError{}
}

/*
GetFlowsDatatableExportJobInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetFlowsDatatableExportJobInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows datatable export job internal server error response has a 2xx status code
func (o *GetFlowsDatatableExportJobInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows datatable export job internal server error response has a 3xx status code
func (o *GetFlowsDatatableExportJobInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows datatable export job internal server error response has a 4xx status code
func (o *GetFlowsDatatableExportJobInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get flows datatable export job internal server error response has a 5xx status code
func (o *GetFlowsDatatableExportJobInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get flows datatable export job internal server error response a status code equal to that given
func (o *GetFlowsDatatableExportJobInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetFlowsDatatableExportJobInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}/export/jobs/{exportJobId}][%d] getFlowsDatatableExportJobInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFlowsDatatableExportJobInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}/export/jobs/{exportJobId}][%d] getFlowsDatatableExportJobInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFlowsDatatableExportJobInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatableExportJobInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableExportJobServiceUnavailable creates a GetFlowsDatatableExportJobServiceUnavailable with default headers values
func NewGetFlowsDatatableExportJobServiceUnavailable() *GetFlowsDatatableExportJobServiceUnavailable {
	return &GetFlowsDatatableExportJobServiceUnavailable{}
}

/*
GetFlowsDatatableExportJobServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetFlowsDatatableExportJobServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows datatable export job service unavailable response has a 2xx status code
func (o *GetFlowsDatatableExportJobServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows datatable export job service unavailable response has a 3xx status code
func (o *GetFlowsDatatableExportJobServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows datatable export job service unavailable response has a 4xx status code
func (o *GetFlowsDatatableExportJobServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get flows datatable export job service unavailable response has a 5xx status code
func (o *GetFlowsDatatableExportJobServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get flows datatable export job service unavailable response a status code equal to that given
func (o *GetFlowsDatatableExportJobServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetFlowsDatatableExportJobServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}/export/jobs/{exportJobId}][%d] getFlowsDatatableExportJobServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFlowsDatatableExportJobServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}/export/jobs/{exportJobId}][%d] getFlowsDatatableExportJobServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFlowsDatatableExportJobServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatableExportJobServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableExportJobGatewayTimeout creates a GetFlowsDatatableExportJobGatewayTimeout with default headers values
func NewGetFlowsDatatableExportJobGatewayTimeout() *GetFlowsDatatableExportJobGatewayTimeout {
	return &GetFlowsDatatableExportJobGatewayTimeout{}
}

/*
GetFlowsDatatableExportJobGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetFlowsDatatableExportJobGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows datatable export job gateway timeout response has a 2xx status code
func (o *GetFlowsDatatableExportJobGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows datatable export job gateway timeout response has a 3xx status code
func (o *GetFlowsDatatableExportJobGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows datatable export job gateway timeout response has a 4xx status code
func (o *GetFlowsDatatableExportJobGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get flows datatable export job gateway timeout response has a 5xx status code
func (o *GetFlowsDatatableExportJobGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get flows datatable export job gateway timeout response a status code equal to that given
func (o *GetFlowsDatatableExportJobGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetFlowsDatatableExportJobGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}/export/jobs/{exportJobId}][%d] getFlowsDatatableExportJobGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetFlowsDatatableExportJobGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}/export/jobs/{exportJobId}][%d] getFlowsDatatableExportJobGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetFlowsDatatableExportJobGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatableExportJobGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

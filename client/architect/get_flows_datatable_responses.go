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

// GetFlowsDatatableReader is a Reader for the GetFlowsDatatable structure.
type GetFlowsDatatableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFlowsDatatableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFlowsDatatableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFlowsDatatableBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFlowsDatatableUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFlowsDatatableForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFlowsDatatableNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetFlowsDatatableRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetFlowsDatatableRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetFlowsDatatableUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetFlowsDatatableTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFlowsDatatableInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetFlowsDatatableServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetFlowsDatatableGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFlowsDatatableOK creates a GetFlowsDatatableOK with default headers values
func NewGetFlowsDatatableOK() *GetFlowsDatatableOK {
	return &GetFlowsDatatableOK{}
}

/*
GetFlowsDatatableOK describes a response with status code 200, with default header values.

successful operation
*/
type GetFlowsDatatableOK struct {
	Payload *models.DataTable
}

// IsSuccess returns true when this get flows datatable o k response has a 2xx status code
func (o *GetFlowsDatatableOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get flows datatable o k response has a 3xx status code
func (o *GetFlowsDatatableOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows datatable o k response has a 4xx status code
func (o *GetFlowsDatatableOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get flows datatable o k response has a 5xx status code
func (o *GetFlowsDatatableOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get flows datatable o k response a status code equal to that given
func (o *GetFlowsDatatableOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetFlowsDatatableOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}][%d] getFlowsDatatableOK  %+v", 200, o.Payload)
}

func (o *GetFlowsDatatableOK) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}][%d] getFlowsDatatableOK  %+v", 200, o.Payload)
}

func (o *GetFlowsDatatableOK) GetPayload() *models.DataTable {
	return o.Payload
}

func (o *GetFlowsDatatableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataTable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableBadRequest creates a GetFlowsDatatableBadRequest with default headers values
func NewGetFlowsDatatableBadRequest() *GetFlowsDatatableBadRequest {
	return &GetFlowsDatatableBadRequest{}
}

/*
GetFlowsDatatableBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetFlowsDatatableBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows datatable bad request response has a 2xx status code
func (o *GetFlowsDatatableBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows datatable bad request response has a 3xx status code
func (o *GetFlowsDatatableBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows datatable bad request response has a 4xx status code
func (o *GetFlowsDatatableBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flows datatable bad request response has a 5xx status code
func (o *GetFlowsDatatableBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get flows datatable bad request response a status code equal to that given
func (o *GetFlowsDatatableBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetFlowsDatatableBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}][%d] getFlowsDatatableBadRequest  %+v", 400, o.Payload)
}

func (o *GetFlowsDatatableBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}][%d] getFlowsDatatableBadRequest  %+v", 400, o.Payload)
}

func (o *GetFlowsDatatableBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatableBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableUnauthorized creates a GetFlowsDatatableUnauthorized with default headers values
func NewGetFlowsDatatableUnauthorized() *GetFlowsDatatableUnauthorized {
	return &GetFlowsDatatableUnauthorized{}
}

/*
GetFlowsDatatableUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetFlowsDatatableUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows datatable unauthorized response has a 2xx status code
func (o *GetFlowsDatatableUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows datatable unauthorized response has a 3xx status code
func (o *GetFlowsDatatableUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows datatable unauthorized response has a 4xx status code
func (o *GetFlowsDatatableUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flows datatable unauthorized response has a 5xx status code
func (o *GetFlowsDatatableUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get flows datatable unauthorized response a status code equal to that given
func (o *GetFlowsDatatableUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetFlowsDatatableUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}][%d] getFlowsDatatableUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFlowsDatatableUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}][%d] getFlowsDatatableUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFlowsDatatableUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatableUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableForbidden creates a GetFlowsDatatableForbidden with default headers values
func NewGetFlowsDatatableForbidden() *GetFlowsDatatableForbidden {
	return &GetFlowsDatatableForbidden{}
}

/*
GetFlowsDatatableForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetFlowsDatatableForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows datatable forbidden response has a 2xx status code
func (o *GetFlowsDatatableForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows datatable forbidden response has a 3xx status code
func (o *GetFlowsDatatableForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows datatable forbidden response has a 4xx status code
func (o *GetFlowsDatatableForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flows datatable forbidden response has a 5xx status code
func (o *GetFlowsDatatableForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get flows datatable forbidden response a status code equal to that given
func (o *GetFlowsDatatableForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetFlowsDatatableForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}][%d] getFlowsDatatableForbidden  %+v", 403, o.Payload)
}

func (o *GetFlowsDatatableForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}][%d] getFlowsDatatableForbidden  %+v", 403, o.Payload)
}

func (o *GetFlowsDatatableForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatableForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableNotFound creates a GetFlowsDatatableNotFound with default headers values
func NewGetFlowsDatatableNotFound() *GetFlowsDatatableNotFound {
	return &GetFlowsDatatableNotFound{}
}

/*
GetFlowsDatatableNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetFlowsDatatableNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows datatable not found response has a 2xx status code
func (o *GetFlowsDatatableNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows datatable not found response has a 3xx status code
func (o *GetFlowsDatatableNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows datatable not found response has a 4xx status code
func (o *GetFlowsDatatableNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flows datatable not found response has a 5xx status code
func (o *GetFlowsDatatableNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get flows datatable not found response a status code equal to that given
func (o *GetFlowsDatatableNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetFlowsDatatableNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}][%d] getFlowsDatatableNotFound  %+v", 404, o.Payload)
}

func (o *GetFlowsDatatableNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}][%d] getFlowsDatatableNotFound  %+v", 404, o.Payload)
}

func (o *GetFlowsDatatableNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatableNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableRequestTimeout creates a GetFlowsDatatableRequestTimeout with default headers values
func NewGetFlowsDatatableRequestTimeout() *GetFlowsDatatableRequestTimeout {
	return &GetFlowsDatatableRequestTimeout{}
}

/*
GetFlowsDatatableRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetFlowsDatatableRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows datatable request timeout response has a 2xx status code
func (o *GetFlowsDatatableRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows datatable request timeout response has a 3xx status code
func (o *GetFlowsDatatableRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows datatable request timeout response has a 4xx status code
func (o *GetFlowsDatatableRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flows datatable request timeout response has a 5xx status code
func (o *GetFlowsDatatableRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get flows datatable request timeout response a status code equal to that given
func (o *GetFlowsDatatableRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetFlowsDatatableRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}][%d] getFlowsDatatableRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetFlowsDatatableRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}][%d] getFlowsDatatableRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetFlowsDatatableRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatableRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableRequestEntityTooLarge creates a GetFlowsDatatableRequestEntityTooLarge with default headers values
func NewGetFlowsDatatableRequestEntityTooLarge() *GetFlowsDatatableRequestEntityTooLarge {
	return &GetFlowsDatatableRequestEntityTooLarge{}
}

/*
GetFlowsDatatableRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetFlowsDatatableRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows datatable request entity too large response has a 2xx status code
func (o *GetFlowsDatatableRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows datatable request entity too large response has a 3xx status code
func (o *GetFlowsDatatableRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows datatable request entity too large response has a 4xx status code
func (o *GetFlowsDatatableRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flows datatable request entity too large response has a 5xx status code
func (o *GetFlowsDatatableRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get flows datatable request entity too large response a status code equal to that given
func (o *GetFlowsDatatableRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetFlowsDatatableRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}][%d] getFlowsDatatableRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetFlowsDatatableRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}][%d] getFlowsDatatableRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetFlowsDatatableRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatableRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableUnsupportedMediaType creates a GetFlowsDatatableUnsupportedMediaType with default headers values
func NewGetFlowsDatatableUnsupportedMediaType() *GetFlowsDatatableUnsupportedMediaType {
	return &GetFlowsDatatableUnsupportedMediaType{}
}

/*
GetFlowsDatatableUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetFlowsDatatableUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows datatable unsupported media type response has a 2xx status code
func (o *GetFlowsDatatableUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows datatable unsupported media type response has a 3xx status code
func (o *GetFlowsDatatableUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows datatable unsupported media type response has a 4xx status code
func (o *GetFlowsDatatableUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flows datatable unsupported media type response has a 5xx status code
func (o *GetFlowsDatatableUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get flows datatable unsupported media type response a status code equal to that given
func (o *GetFlowsDatatableUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetFlowsDatatableUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}][%d] getFlowsDatatableUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetFlowsDatatableUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}][%d] getFlowsDatatableUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetFlowsDatatableUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatableUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableTooManyRequests creates a GetFlowsDatatableTooManyRequests with default headers values
func NewGetFlowsDatatableTooManyRequests() *GetFlowsDatatableTooManyRequests {
	return &GetFlowsDatatableTooManyRequests{}
}

/*
GetFlowsDatatableTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetFlowsDatatableTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows datatable too many requests response has a 2xx status code
func (o *GetFlowsDatatableTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows datatable too many requests response has a 3xx status code
func (o *GetFlowsDatatableTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows datatable too many requests response has a 4xx status code
func (o *GetFlowsDatatableTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flows datatable too many requests response has a 5xx status code
func (o *GetFlowsDatatableTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get flows datatable too many requests response a status code equal to that given
func (o *GetFlowsDatatableTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetFlowsDatatableTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}][%d] getFlowsDatatableTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetFlowsDatatableTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}][%d] getFlowsDatatableTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetFlowsDatatableTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatableTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableInternalServerError creates a GetFlowsDatatableInternalServerError with default headers values
func NewGetFlowsDatatableInternalServerError() *GetFlowsDatatableInternalServerError {
	return &GetFlowsDatatableInternalServerError{}
}

/*
GetFlowsDatatableInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetFlowsDatatableInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows datatable internal server error response has a 2xx status code
func (o *GetFlowsDatatableInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows datatable internal server error response has a 3xx status code
func (o *GetFlowsDatatableInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows datatable internal server error response has a 4xx status code
func (o *GetFlowsDatatableInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get flows datatable internal server error response has a 5xx status code
func (o *GetFlowsDatatableInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get flows datatable internal server error response a status code equal to that given
func (o *GetFlowsDatatableInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetFlowsDatatableInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}][%d] getFlowsDatatableInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFlowsDatatableInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}][%d] getFlowsDatatableInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFlowsDatatableInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatableInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableServiceUnavailable creates a GetFlowsDatatableServiceUnavailable with default headers values
func NewGetFlowsDatatableServiceUnavailable() *GetFlowsDatatableServiceUnavailable {
	return &GetFlowsDatatableServiceUnavailable{}
}

/*
GetFlowsDatatableServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetFlowsDatatableServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows datatable service unavailable response has a 2xx status code
func (o *GetFlowsDatatableServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows datatable service unavailable response has a 3xx status code
func (o *GetFlowsDatatableServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows datatable service unavailable response has a 4xx status code
func (o *GetFlowsDatatableServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get flows datatable service unavailable response has a 5xx status code
func (o *GetFlowsDatatableServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get flows datatable service unavailable response a status code equal to that given
func (o *GetFlowsDatatableServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetFlowsDatatableServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}][%d] getFlowsDatatableServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFlowsDatatableServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}][%d] getFlowsDatatableServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFlowsDatatableServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatableServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableGatewayTimeout creates a GetFlowsDatatableGatewayTimeout with default headers values
func NewGetFlowsDatatableGatewayTimeout() *GetFlowsDatatableGatewayTimeout {
	return &GetFlowsDatatableGatewayTimeout{}
}

/*
GetFlowsDatatableGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetFlowsDatatableGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows datatable gateway timeout response has a 2xx status code
func (o *GetFlowsDatatableGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows datatable gateway timeout response has a 3xx status code
func (o *GetFlowsDatatableGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows datatable gateway timeout response has a 4xx status code
func (o *GetFlowsDatatableGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get flows datatable gateway timeout response has a 5xx status code
func (o *GetFlowsDatatableGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get flows datatable gateway timeout response a status code equal to that given
func (o *GetFlowsDatatableGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetFlowsDatatableGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}][%d] getFlowsDatatableGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetFlowsDatatableGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}][%d] getFlowsDatatableGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetFlowsDatatableGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatableGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

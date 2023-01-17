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

// GetFlowsDatatablesDivisionviewsReader is a Reader for the GetFlowsDatatablesDivisionviews structure.
type GetFlowsDatatablesDivisionviewsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFlowsDatatablesDivisionviewsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFlowsDatatablesDivisionviewsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFlowsDatatablesDivisionviewsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFlowsDatatablesDivisionviewsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFlowsDatatablesDivisionviewsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFlowsDatatablesDivisionviewsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetFlowsDatatablesDivisionviewsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetFlowsDatatablesDivisionviewsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetFlowsDatatablesDivisionviewsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetFlowsDatatablesDivisionviewsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFlowsDatatablesDivisionviewsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetFlowsDatatablesDivisionviewsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetFlowsDatatablesDivisionviewsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFlowsDatatablesDivisionviewsOK creates a GetFlowsDatatablesDivisionviewsOK with default headers values
func NewGetFlowsDatatablesDivisionviewsOK() *GetFlowsDatatablesDivisionviewsOK {
	return &GetFlowsDatatablesDivisionviewsOK{}
}

/*
GetFlowsDatatablesDivisionviewsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetFlowsDatatablesDivisionviewsOK struct {
	Payload *models.DataTablesDomainEntityListing
}

// IsSuccess returns true when this get flows datatables divisionviews o k response has a 2xx status code
func (o *GetFlowsDatatablesDivisionviewsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get flows datatables divisionviews o k response has a 3xx status code
func (o *GetFlowsDatatablesDivisionviewsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows datatables divisionviews o k response has a 4xx status code
func (o *GetFlowsDatatablesDivisionviewsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get flows datatables divisionviews o k response has a 5xx status code
func (o *GetFlowsDatatablesDivisionviewsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get flows datatables divisionviews o k response a status code equal to that given
func (o *GetFlowsDatatablesDivisionviewsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetFlowsDatatablesDivisionviewsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/divisionviews][%d] getFlowsDatatablesDivisionviewsOK  %+v", 200, o.Payload)
}

func (o *GetFlowsDatatablesDivisionviewsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/divisionviews][%d] getFlowsDatatablesDivisionviewsOK  %+v", 200, o.Payload)
}

func (o *GetFlowsDatatablesDivisionviewsOK) GetPayload() *models.DataTablesDomainEntityListing {
	return o.Payload
}

func (o *GetFlowsDatatablesDivisionviewsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataTablesDomainEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatablesDivisionviewsBadRequest creates a GetFlowsDatatablesDivisionviewsBadRequest with default headers values
func NewGetFlowsDatatablesDivisionviewsBadRequest() *GetFlowsDatatablesDivisionviewsBadRequest {
	return &GetFlowsDatatablesDivisionviewsBadRequest{}
}

/*
GetFlowsDatatablesDivisionviewsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetFlowsDatatablesDivisionviewsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows datatables divisionviews bad request response has a 2xx status code
func (o *GetFlowsDatatablesDivisionviewsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows datatables divisionviews bad request response has a 3xx status code
func (o *GetFlowsDatatablesDivisionviewsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows datatables divisionviews bad request response has a 4xx status code
func (o *GetFlowsDatatablesDivisionviewsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flows datatables divisionviews bad request response has a 5xx status code
func (o *GetFlowsDatatablesDivisionviewsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get flows datatables divisionviews bad request response a status code equal to that given
func (o *GetFlowsDatatablesDivisionviewsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetFlowsDatatablesDivisionviewsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/divisionviews][%d] getFlowsDatatablesDivisionviewsBadRequest  %+v", 400, o.Payload)
}

func (o *GetFlowsDatatablesDivisionviewsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/divisionviews][%d] getFlowsDatatablesDivisionviewsBadRequest  %+v", 400, o.Payload)
}

func (o *GetFlowsDatatablesDivisionviewsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatablesDivisionviewsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatablesDivisionviewsUnauthorized creates a GetFlowsDatatablesDivisionviewsUnauthorized with default headers values
func NewGetFlowsDatatablesDivisionviewsUnauthorized() *GetFlowsDatatablesDivisionviewsUnauthorized {
	return &GetFlowsDatatablesDivisionviewsUnauthorized{}
}

/*
GetFlowsDatatablesDivisionviewsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetFlowsDatatablesDivisionviewsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows datatables divisionviews unauthorized response has a 2xx status code
func (o *GetFlowsDatatablesDivisionviewsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows datatables divisionviews unauthorized response has a 3xx status code
func (o *GetFlowsDatatablesDivisionviewsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows datatables divisionviews unauthorized response has a 4xx status code
func (o *GetFlowsDatatablesDivisionviewsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flows datatables divisionviews unauthorized response has a 5xx status code
func (o *GetFlowsDatatablesDivisionviewsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get flows datatables divisionviews unauthorized response a status code equal to that given
func (o *GetFlowsDatatablesDivisionviewsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetFlowsDatatablesDivisionviewsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/divisionviews][%d] getFlowsDatatablesDivisionviewsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFlowsDatatablesDivisionviewsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/divisionviews][%d] getFlowsDatatablesDivisionviewsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFlowsDatatablesDivisionviewsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatablesDivisionviewsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatablesDivisionviewsForbidden creates a GetFlowsDatatablesDivisionviewsForbidden with default headers values
func NewGetFlowsDatatablesDivisionviewsForbidden() *GetFlowsDatatablesDivisionviewsForbidden {
	return &GetFlowsDatatablesDivisionviewsForbidden{}
}

/*
GetFlowsDatatablesDivisionviewsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetFlowsDatatablesDivisionviewsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows datatables divisionviews forbidden response has a 2xx status code
func (o *GetFlowsDatatablesDivisionviewsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows datatables divisionviews forbidden response has a 3xx status code
func (o *GetFlowsDatatablesDivisionviewsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows datatables divisionviews forbidden response has a 4xx status code
func (o *GetFlowsDatatablesDivisionviewsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flows datatables divisionviews forbidden response has a 5xx status code
func (o *GetFlowsDatatablesDivisionviewsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get flows datatables divisionviews forbidden response a status code equal to that given
func (o *GetFlowsDatatablesDivisionviewsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetFlowsDatatablesDivisionviewsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/divisionviews][%d] getFlowsDatatablesDivisionviewsForbidden  %+v", 403, o.Payload)
}

func (o *GetFlowsDatatablesDivisionviewsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/divisionviews][%d] getFlowsDatatablesDivisionviewsForbidden  %+v", 403, o.Payload)
}

func (o *GetFlowsDatatablesDivisionviewsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatablesDivisionviewsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatablesDivisionviewsNotFound creates a GetFlowsDatatablesDivisionviewsNotFound with default headers values
func NewGetFlowsDatatablesDivisionviewsNotFound() *GetFlowsDatatablesDivisionviewsNotFound {
	return &GetFlowsDatatablesDivisionviewsNotFound{}
}

/*
GetFlowsDatatablesDivisionviewsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetFlowsDatatablesDivisionviewsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows datatables divisionviews not found response has a 2xx status code
func (o *GetFlowsDatatablesDivisionviewsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows datatables divisionviews not found response has a 3xx status code
func (o *GetFlowsDatatablesDivisionviewsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows datatables divisionviews not found response has a 4xx status code
func (o *GetFlowsDatatablesDivisionviewsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flows datatables divisionviews not found response has a 5xx status code
func (o *GetFlowsDatatablesDivisionviewsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get flows datatables divisionviews not found response a status code equal to that given
func (o *GetFlowsDatatablesDivisionviewsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetFlowsDatatablesDivisionviewsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/divisionviews][%d] getFlowsDatatablesDivisionviewsNotFound  %+v", 404, o.Payload)
}

func (o *GetFlowsDatatablesDivisionviewsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/divisionviews][%d] getFlowsDatatablesDivisionviewsNotFound  %+v", 404, o.Payload)
}

func (o *GetFlowsDatatablesDivisionviewsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatablesDivisionviewsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatablesDivisionviewsRequestTimeout creates a GetFlowsDatatablesDivisionviewsRequestTimeout with default headers values
func NewGetFlowsDatatablesDivisionviewsRequestTimeout() *GetFlowsDatatablesDivisionviewsRequestTimeout {
	return &GetFlowsDatatablesDivisionviewsRequestTimeout{}
}

/*
GetFlowsDatatablesDivisionviewsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetFlowsDatatablesDivisionviewsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows datatables divisionviews request timeout response has a 2xx status code
func (o *GetFlowsDatatablesDivisionviewsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows datatables divisionviews request timeout response has a 3xx status code
func (o *GetFlowsDatatablesDivisionviewsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows datatables divisionviews request timeout response has a 4xx status code
func (o *GetFlowsDatatablesDivisionviewsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flows datatables divisionviews request timeout response has a 5xx status code
func (o *GetFlowsDatatablesDivisionviewsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get flows datatables divisionviews request timeout response a status code equal to that given
func (o *GetFlowsDatatablesDivisionviewsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetFlowsDatatablesDivisionviewsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/divisionviews][%d] getFlowsDatatablesDivisionviewsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetFlowsDatatablesDivisionviewsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/divisionviews][%d] getFlowsDatatablesDivisionviewsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetFlowsDatatablesDivisionviewsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatablesDivisionviewsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatablesDivisionviewsRequestEntityTooLarge creates a GetFlowsDatatablesDivisionviewsRequestEntityTooLarge with default headers values
func NewGetFlowsDatatablesDivisionviewsRequestEntityTooLarge() *GetFlowsDatatablesDivisionviewsRequestEntityTooLarge {
	return &GetFlowsDatatablesDivisionviewsRequestEntityTooLarge{}
}

/*
GetFlowsDatatablesDivisionviewsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetFlowsDatatablesDivisionviewsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows datatables divisionviews request entity too large response has a 2xx status code
func (o *GetFlowsDatatablesDivisionviewsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows datatables divisionviews request entity too large response has a 3xx status code
func (o *GetFlowsDatatablesDivisionviewsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows datatables divisionviews request entity too large response has a 4xx status code
func (o *GetFlowsDatatablesDivisionviewsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flows datatables divisionviews request entity too large response has a 5xx status code
func (o *GetFlowsDatatablesDivisionviewsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get flows datatables divisionviews request entity too large response a status code equal to that given
func (o *GetFlowsDatatablesDivisionviewsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetFlowsDatatablesDivisionviewsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/divisionviews][%d] getFlowsDatatablesDivisionviewsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetFlowsDatatablesDivisionviewsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/divisionviews][%d] getFlowsDatatablesDivisionviewsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetFlowsDatatablesDivisionviewsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatablesDivisionviewsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatablesDivisionviewsUnsupportedMediaType creates a GetFlowsDatatablesDivisionviewsUnsupportedMediaType with default headers values
func NewGetFlowsDatatablesDivisionviewsUnsupportedMediaType() *GetFlowsDatatablesDivisionviewsUnsupportedMediaType {
	return &GetFlowsDatatablesDivisionviewsUnsupportedMediaType{}
}

/*
GetFlowsDatatablesDivisionviewsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetFlowsDatatablesDivisionviewsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows datatables divisionviews unsupported media type response has a 2xx status code
func (o *GetFlowsDatatablesDivisionviewsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows datatables divisionviews unsupported media type response has a 3xx status code
func (o *GetFlowsDatatablesDivisionviewsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows datatables divisionviews unsupported media type response has a 4xx status code
func (o *GetFlowsDatatablesDivisionviewsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flows datatables divisionviews unsupported media type response has a 5xx status code
func (o *GetFlowsDatatablesDivisionviewsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get flows datatables divisionviews unsupported media type response a status code equal to that given
func (o *GetFlowsDatatablesDivisionviewsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetFlowsDatatablesDivisionviewsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/divisionviews][%d] getFlowsDatatablesDivisionviewsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetFlowsDatatablesDivisionviewsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/divisionviews][%d] getFlowsDatatablesDivisionviewsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetFlowsDatatablesDivisionviewsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatablesDivisionviewsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatablesDivisionviewsTooManyRequests creates a GetFlowsDatatablesDivisionviewsTooManyRequests with default headers values
func NewGetFlowsDatatablesDivisionviewsTooManyRequests() *GetFlowsDatatablesDivisionviewsTooManyRequests {
	return &GetFlowsDatatablesDivisionviewsTooManyRequests{}
}

/*
GetFlowsDatatablesDivisionviewsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetFlowsDatatablesDivisionviewsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows datatables divisionviews too many requests response has a 2xx status code
func (o *GetFlowsDatatablesDivisionviewsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows datatables divisionviews too many requests response has a 3xx status code
func (o *GetFlowsDatatablesDivisionviewsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows datatables divisionviews too many requests response has a 4xx status code
func (o *GetFlowsDatatablesDivisionviewsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flows datatables divisionviews too many requests response has a 5xx status code
func (o *GetFlowsDatatablesDivisionviewsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get flows datatables divisionviews too many requests response a status code equal to that given
func (o *GetFlowsDatatablesDivisionviewsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetFlowsDatatablesDivisionviewsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/divisionviews][%d] getFlowsDatatablesDivisionviewsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetFlowsDatatablesDivisionviewsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/divisionviews][%d] getFlowsDatatablesDivisionviewsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetFlowsDatatablesDivisionviewsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatablesDivisionviewsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatablesDivisionviewsInternalServerError creates a GetFlowsDatatablesDivisionviewsInternalServerError with default headers values
func NewGetFlowsDatatablesDivisionviewsInternalServerError() *GetFlowsDatatablesDivisionviewsInternalServerError {
	return &GetFlowsDatatablesDivisionviewsInternalServerError{}
}

/*
GetFlowsDatatablesDivisionviewsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetFlowsDatatablesDivisionviewsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows datatables divisionviews internal server error response has a 2xx status code
func (o *GetFlowsDatatablesDivisionviewsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows datatables divisionviews internal server error response has a 3xx status code
func (o *GetFlowsDatatablesDivisionviewsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows datatables divisionviews internal server error response has a 4xx status code
func (o *GetFlowsDatatablesDivisionviewsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get flows datatables divisionviews internal server error response has a 5xx status code
func (o *GetFlowsDatatablesDivisionviewsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get flows datatables divisionviews internal server error response a status code equal to that given
func (o *GetFlowsDatatablesDivisionviewsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetFlowsDatatablesDivisionviewsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/divisionviews][%d] getFlowsDatatablesDivisionviewsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFlowsDatatablesDivisionviewsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/divisionviews][%d] getFlowsDatatablesDivisionviewsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFlowsDatatablesDivisionviewsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatablesDivisionviewsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatablesDivisionviewsServiceUnavailable creates a GetFlowsDatatablesDivisionviewsServiceUnavailable with default headers values
func NewGetFlowsDatatablesDivisionviewsServiceUnavailable() *GetFlowsDatatablesDivisionviewsServiceUnavailable {
	return &GetFlowsDatatablesDivisionviewsServiceUnavailable{}
}

/*
GetFlowsDatatablesDivisionviewsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetFlowsDatatablesDivisionviewsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows datatables divisionviews service unavailable response has a 2xx status code
func (o *GetFlowsDatatablesDivisionviewsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows datatables divisionviews service unavailable response has a 3xx status code
func (o *GetFlowsDatatablesDivisionviewsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows datatables divisionviews service unavailable response has a 4xx status code
func (o *GetFlowsDatatablesDivisionviewsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get flows datatables divisionviews service unavailable response has a 5xx status code
func (o *GetFlowsDatatablesDivisionviewsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get flows datatables divisionviews service unavailable response a status code equal to that given
func (o *GetFlowsDatatablesDivisionviewsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetFlowsDatatablesDivisionviewsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/divisionviews][%d] getFlowsDatatablesDivisionviewsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFlowsDatatablesDivisionviewsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/divisionviews][%d] getFlowsDatatablesDivisionviewsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFlowsDatatablesDivisionviewsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatablesDivisionviewsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatablesDivisionviewsGatewayTimeout creates a GetFlowsDatatablesDivisionviewsGatewayTimeout with default headers values
func NewGetFlowsDatatablesDivisionviewsGatewayTimeout() *GetFlowsDatatablesDivisionviewsGatewayTimeout {
	return &GetFlowsDatatablesDivisionviewsGatewayTimeout{}
}

/*
GetFlowsDatatablesDivisionviewsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetFlowsDatatablesDivisionviewsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flows datatables divisionviews gateway timeout response has a 2xx status code
func (o *GetFlowsDatatablesDivisionviewsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flows datatables divisionviews gateway timeout response has a 3xx status code
func (o *GetFlowsDatatablesDivisionviewsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flows datatables divisionviews gateway timeout response has a 4xx status code
func (o *GetFlowsDatatablesDivisionviewsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get flows datatables divisionviews gateway timeout response has a 5xx status code
func (o *GetFlowsDatatablesDivisionviewsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get flows datatables divisionviews gateway timeout response a status code equal to that given
func (o *GetFlowsDatatablesDivisionviewsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetFlowsDatatablesDivisionviewsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/divisionviews][%d] getFlowsDatatablesDivisionviewsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetFlowsDatatablesDivisionviewsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/divisionviews][%d] getFlowsDatatablesDivisionviewsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetFlowsDatatablesDivisionviewsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatablesDivisionviewsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// GetWorkforcemanagementBusinessunitActivitycodeReader is a Reader for the GetWorkforcemanagementBusinessunitActivitycode structure.
type GetWorkforcemanagementBusinessunitActivitycodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkforcemanagementBusinessunitActivitycodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkforcemanagementBusinessunitActivitycodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkforcemanagementBusinessunitActivitycodeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWorkforcemanagementBusinessunitActivitycodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkforcemanagementBusinessunitActivitycodeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkforcemanagementBusinessunitActivitycodeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetWorkforcemanagementBusinessunitActivitycodeRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWorkforcemanagementBusinessunitActivitycodeRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWorkforcemanagementBusinessunitActivitycodeUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWorkforcemanagementBusinessunitActivitycodeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkforcemanagementBusinessunitActivitycodeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWorkforcemanagementBusinessunitActivitycodeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWorkforcemanagementBusinessunitActivitycodeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkforcemanagementBusinessunitActivitycodeOK creates a GetWorkforcemanagementBusinessunitActivitycodeOK with default headers values
func NewGetWorkforcemanagementBusinessunitActivitycodeOK() *GetWorkforcemanagementBusinessunitActivitycodeOK {
	return &GetWorkforcemanagementBusinessunitActivitycodeOK{}
}

/*GetWorkforcemanagementBusinessunitActivitycodeOK handles this case with default header values.

successful operation
*/
type GetWorkforcemanagementBusinessunitActivitycodeOK struct {
	Payload *models.BusinessUnitActivityCode
}

func (o *GetWorkforcemanagementBusinessunitActivitycodeOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}][%d] getWorkforcemanagementBusinessunitActivitycodeOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitActivitycodeOK) GetPayload() *models.BusinessUnitActivityCode {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitActivitycodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BusinessUnitActivityCode)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitActivitycodeBadRequest creates a GetWorkforcemanagementBusinessunitActivitycodeBadRequest with default headers values
func NewGetWorkforcemanagementBusinessunitActivitycodeBadRequest() *GetWorkforcemanagementBusinessunitActivitycodeBadRequest {
	return &GetWorkforcemanagementBusinessunitActivitycodeBadRequest{}
}

/*GetWorkforcemanagementBusinessunitActivitycodeBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWorkforcemanagementBusinessunitActivitycodeBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitActivitycodeBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}][%d] getWorkforcemanagementBusinessunitActivitycodeBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitActivitycodeBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitActivitycodeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitActivitycodeUnauthorized creates a GetWorkforcemanagementBusinessunitActivitycodeUnauthorized with default headers values
func NewGetWorkforcemanagementBusinessunitActivitycodeUnauthorized() *GetWorkforcemanagementBusinessunitActivitycodeUnauthorized {
	return &GetWorkforcemanagementBusinessunitActivitycodeUnauthorized{}
}

/*GetWorkforcemanagementBusinessunitActivitycodeUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWorkforcemanagementBusinessunitActivitycodeUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitActivitycodeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}][%d] getWorkforcemanagementBusinessunitActivitycodeUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitActivitycodeUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitActivitycodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitActivitycodeForbidden creates a GetWorkforcemanagementBusinessunitActivitycodeForbidden with default headers values
func NewGetWorkforcemanagementBusinessunitActivitycodeForbidden() *GetWorkforcemanagementBusinessunitActivitycodeForbidden {
	return &GetWorkforcemanagementBusinessunitActivitycodeForbidden{}
}

/*GetWorkforcemanagementBusinessunitActivitycodeForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetWorkforcemanagementBusinessunitActivitycodeForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitActivitycodeForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}][%d] getWorkforcemanagementBusinessunitActivitycodeForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitActivitycodeForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitActivitycodeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitActivitycodeNotFound creates a GetWorkforcemanagementBusinessunitActivitycodeNotFound with default headers values
func NewGetWorkforcemanagementBusinessunitActivitycodeNotFound() *GetWorkforcemanagementBusinessunitActivitycodeNotFound {
	return &GetWorkforcemanagementBusinessunitActivitycodeNotFound{}
}

/*GetWorkforcemanagementBusinessunitActivitycodeNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetWorkforcemanagementBusinessunitActivitycodeNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitActivitycodeNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}][%d] getWorkforcemanagementBusinessunitActivitycodeNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitActivitycodeNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitActivitycodeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitActivitycodeRequestTimeout creates a GetWorkforcemanagementBusinessunitActivitycodeRequestTimeout with default headers values
func NewGetWorkforcemanagementBusinessunitActivitycodeRequestTimeout() *GetWorkforcemanagementBusinessunitActivitycodeRequestTimeout {
	return &GetWorkforcemanagementBusinessunitActivitycodeRequestTimeout{}
}

/*GetWorkforcemanagementBusinessunitActivitycodeRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetWorkforcemanagementBusinessunitActivitycodeRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitActivitycodeRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}][%d] getWorkforcemanagementBusinessunitActivitycodeRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitActivitycodeRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitActivitycodeRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitActivitycodeRequestEntityTooLarge creates a GetWorkforcemanagementBusinessunitActivitycodeRequestEntityTooLarge with default headers values
func NewGetWorkforcemanagementBusinessunitActivitycodeRequestEntityTooLarge() *GetWorkforcemanagementBusinessunitActivitycodeRequestEntityTooLarge {
	return &GetWorkforcemanagementBusinessunitActivitycodeRequestEntityTooLarge{}
}

/*GetWorkforcemanagementBusinessunitActivitycodeRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetWorkforcemanagementBusinessunitActivitycodeRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitActivitycodeRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}][%d] getWorkforcemanagementBusinessunitActivitycodeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitActivitycodeRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitActivitycodeRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitActivitycodeUnsupportedMediaType creates a GetWorkforcemanagementBusinessunitActivitycodeUnsupportedMediaType with default headers values
func NewGetWorkforcemanagementBusinessunitActivitycodeUnsupportedMediaType() *GetWorkforcemanagementBusinessunitActivitycodeUnsupportedMediaType {
	return &GetWorkforcemanagementBusinessunitActivitycodeUnsupportedMediaType{}
}

/*GetWorkforcemanagementBusinessunitActivitycodeUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWorkforcemanagementBusinessunitActivitycodeUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitActivitycodeUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}][%d] getWorkforcemanagementBusinessunitActivitycodeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitActivitycodeUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitActivitycodeUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitActivitycodeTooManyRequests creates a GetWorkforcemanagementBusinessunitActivitycodeTooManyRequests with default headers values
func NewGetWorkforcemanagementBusinessunitActivitycodeTooManyRequests() *GetWorkforcemanagementBusinessunitActivitycodeTooManyRequests {
	return &GetWorkforcemanagementBusinessunitActivitycodeTooManyRequests{}
}

/*GetWorkforcemanagementBusinessunitActivitycodeTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetWorkforcemanagementBusinessunitActivitycodeTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitActivitycodeTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}][%d] getWorkforcemanagementBusinessunitActivitycodeTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitActivitycodeTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitActivitycodeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitActivitycodeInternalServerError creates a GetWorkforcemanagementBusinessunitActivitycodeInternalServerError with default headers values
func NewGetWorkforcemanagementBusinessunitActivitycodeInternalServerError() *GetWorkforcemanagementBusinessunitActivitycodeInternalServerError {
	return &GetWorkforcemanagementBusinessunitActivitycodeInternalServerError{}
}

/*GetWorkforcemanagementBusinessunitActivitycodeInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWorkforcemanagementBusinessunitActivitycodeInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitActivitycodeInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}][%d] getWorkforcemanagementBusinessunitActivitycodeInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitActivitycodeInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitActivitycodeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitActivitycodeServiceUnavailable creates a GetWorkforcemanagementBusinessunitActivitycodeServiceUnavailable with default headers values
func NewGetWorkforcemanagementBusinessunitActivitycodeServiceUnavailable() *GetWorkforcemanagementBusinessunitActivitycodeServiceUnavailable {
	return &GetWorkforcemanagementBusinessunitActivitycodeServiceUnavailable{}
}

/*GetWorkforcemanagementBusinessunitActivitycodeServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWorkforcemanagementBusinessunitActivitycodeServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitActivitycodeServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}][%d] getWorkforcemanagementBusinessunitActivitycodeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitActivitycodeServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitActivitycodeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitActivitycodeGatewayTimeout creates a GetWorkforcemanagementBusinessunitActivitycodeGatewayTimeout with default headers values
func NewGetWorkforcemanagementBusinessunitActivitycodeGatewayTimeout() *GetWorkforcemanagementBusinessunitActivitycodeGatewayTimeout {
	return &GetWorkforcemanagementBusinessunitActivitycodeGatewayTimeout{}
}

/*GetWorkforcemanagementBusinessunitActivitycodeGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetWorkforcemanagementBusinessunitActivitycodeGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitActivitycodeGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}][%d] getWorkforcemanagementBusinessunitActivitycodeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitActivitycodeGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitActivitycodeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

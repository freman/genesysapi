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

// GetWorkforcemanagementAgentsMeManagementunitReader is a Reader for the GetWorkforcemanagementAgentsMeManagementunit structure.
type GetWorkforcemanagementAgentsMeManagementunitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkforcemanagementAgentsMeManagementunitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkforcemanagementAgentsMeManagementunitOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkforcemanagementAgentsMeManagementunitBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWorkforcemanagementAgentsMeManagementunitUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkforcemanagementAgentsMeManagementunitForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkforcemanagementAgentsMeManagementunitNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetWorkforcemanagementAgentsMeManagementunitRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWorkforcemanagementAgentsMeManagementunitRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWorkforcemanagementAgentsMeManagementunitUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWorkforcemanagementAgentsMeManagementunitTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkforcemanagementAgentsMeManagementunitInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWorkforcemanagementAgentsMeManagementunitServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWorkforcemanagementAgentsMeManagementunitGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkforcemanagementAgentsMeManagementunitOK creates a GetWorkforcemanagementAgentsMeManagementunitOK with default headers values
func NewGetWorkforcemanagementAgentsMeManagementunitOK() *GetWorkforcemanagementAgentsMeManagementunitOK {
	return &GetWorkforcemanagementAgentsMeManagementunitOK{}
}

/*
GetWorkforcemanagementAgentsMeManagementunitOK describes a response with status code 200, with default header values.

successful operation
*/
type GetWorkforcemanagementAgentsMeManagementunitOK struct {
	Payload *models.AgentManagementUnitReference
}

// IsSuccess returns true when this get workforcemanagement agents me managementunit o k response has a 2xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get workforcemanagement agents me managementunit o k response has a 3xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement agents me managementunit o k response has a 4xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement agents me managementunit o k response has a 5xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement agents me managementunit o k response a status code equal to that given
func (o *GetWorkforcemanagementAgentsMeManagementunitOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetWorkforcemanagementAgentsMeManagementunitOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/me/managementunit][%d] getWorkforcemanagementAgentsMeManagementunitOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementAgentsMeManagementunitOK) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/me/managementunit][%d] getWorkforcemanagementAgentsMeManagementunitOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementAgentsMeManagementunitOK) GetPayload() *models.AgentManagementUnitReference {
	return o.Payload
}

func (o *GetWorkforcemanagementAgentsMeManagementunitOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AgentManagementUnitReference)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAgentsMeManagementunitBadRequest creates a GetWorkforcemanagementAgentsMeManagementunitBadRequest with default headers values
func NewGetWorkforcemanagementAgentsMeManagementunitBadRequest() *GetWorkforcemanagementAgentsMeManagementunitBadRequest {
	return &GetWorkforcemanagementAgentsMeManagementunitBadRequest{}
}

/*
GetWorkforcemanagementAgentsMeManagementunitBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWorkforcemanagementAgentsMeManagementunitBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement agents me managementunit bad request response has a 2xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement agents me managementunit bad request response has a 3xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement agents me managementunit bad request response has a 4xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement agents me managementunit bad request response has a 5xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement agents me managementunit bad request response a status code equal to that given
func (o *GetWorkforcemanagementAgentsMeManagementunitBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetWorkforcemanagementAgentsMeManagementunitBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/me/managementunit][%d] getWorkforcemanagementAgentsMeManagementunitBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementAgentsMeManagementunitBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/me/managementunit][%d] getWorkforcemanagementAgentsMeManagementunitBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementAgentsMeManagementunitBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAgentsMeManagementunitBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAgentsMeManagementunitUnauthorized creates a GetWorkforcemanagementAgentsMeManagementunitUnauthorized with default headers values
func NewGetWorkforcemanagementAgentsMeManagementunitUnauthorized() *GetWorkforcemanagementAgentsMeManagementunitUnauthorized {
	return &GetWorkforcemanagementAgentsMeManagementunitUnauthorized{}
}

/*
GetWorkforcemanagementAgentsMeManagementunitUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWorkforcemanagementAgentsMeManagementunitUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement agents me managementunit unauthorized response has a 2xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement agents me managementunit unauthorized response has a 3xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement agents me managementunit unauthorized response has a 4xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement agents me managementunit unauthorized response has a 5xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement agents me managementunit unauthorized response a status code equal to that given
func (o *GetWorkforcemanagementAgentsMeManagementunitUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetWorkforcemanagementAgentsMeManagementunitUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/me/managementunit][%d] getWorkforcemanagementAgentsMeManagementunitUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementAgentsMeManagementunitUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/me/managementunit][%d] getWorkforcemanagementAgentsMeManagementunitUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementAgentsMeManagementunitUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAgentsMeManagementunitUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAgentsMeManagementunitForbidden creates a GetWorkforcemanagementAgentsMeManagementunitForbidden with default headers values
func NewGetWorkforcemanagementAgentsMeManagementunitForbidden() *GetWorkforcemanagementAgentsMeManagementunitForbidden {
	return &GetWorkforcemanagementAgentsMeManagementunitForbidden{}
}

/*
GetWorkforcemanagementAgentsMeManagementunitForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetWorkforcemanagementAgentsMeManagementunitForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement agents me managementunit forbidden response has a 2xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement agents me managementunit forbidden response has a 3xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement agents me managementunit forbidden response has a 4xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement agents me managementunit forbidden response has a 5xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement agents me managementunit forbidden response a status code equal to that given
func (o *GetWorkforcemanagementAgentsMeManagementunitForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetWorkforcemanagementAgentsMeManagementunitForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/me/managementunit][%d] getWorkforcemanagementAgentsMeManagementunitForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementAgentsMeManagementunitForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/me/managementunit][%d] getWorkforcemanagementAgentsMeManagementunitForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementAgentsMeManagementunitForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAgentsMeManagementunitForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAgentsMeManagementunitNotFound creates a GetWorkforcemanagementAgentsMeManagementunitNotFound with default headers values
func NewGetWorkforcemanagementAgentsMeManagementunitNotFound() *GetWorkforcemanagementAgentsMeManagementunitNotFound {
	return &GetWorkforcemanagementAgentsMeManagementunitNotFound{}
}

/*
GetWorkforcemanagementAgentsMeManagementunitNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetWorkforcemanagementAgentsMeManagementunitNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement agents me managementunit not found response has a 2xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement agents me managementunit not found response has a 3xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement agents me managementunit not found response has a 4xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement agents me managementunit not found response has a 5xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement agents me managementunit not found response a status code equal to that given
func (o *GetWorkforcemanagementAgentsMeManagementunitNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetWorkforcemanagementAgentsMeManagementunitNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/me/managementunit][%d] getWorkforcemanagementAgentsMeManagementunitNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementAgentsMeManagementunitNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/me/managementunit][%d] getWorkforcemanagementAgentsMeManagementunitNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementAgentsMeManagementunitNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAgentsMeManagementunitNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAgentsMeManagementunitRequestTimeout creates a GetWorkforcemanagementAgentsMeManagementunitRequestTimeout with default headers values
func NewGetWorkforcemanagementAgentsMeManagementunitRequestTimeout() *GetWorkforcemanagementAgentsMeManagementunitRequestTimeout {
	return &GetWorkforcemanagementAgentsMeManagementunitRequestTimeout{}
}

/*
GetWorkforcemanagementAgentsMeManagementunitRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetWorkforcemanagementAgentsMeManagementunitRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement agents me managementunit request timeout response has a 2xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement agents me managementunit request timeout response has a 3xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement agents me managementunit request timeout response has a 4xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement agents me managementunit request timeout response has a 5xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement agents me managementunit request timeout response a status code equal to that given
func (o *GetWorkforcemanagementAgentsMeManagementunitRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetWorkforcemanagementAgentsMeManagementunitRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/me/managementunit][%d] getWorkforcemanagementAgentsMeManagementunitRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementAgentsMeManagementunitRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/me/managementunit][%d] getWorkforcemanagementAgentsMeManagementunitRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementAgentsMeManagementunitRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAgentsMeManagementunitRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAgentsMeManagementunitRequestEntityTooLarge creates a GetWorkforcemanagementAgentsMeManagementunitRequestEntityTooLarge with default headers values
func NewGetWorkforcemanagementAgentsMeManagementunitRequestEntityTooLarge() *GetWorkforcemanagementAgentsMeManagementunitRequestEntityTooLarge {
	return &GetWorkforcemanagementAgentsMeManagementunitRequestEntityTooLarge{}
}

/*
GetWorkforcemanagementAgentsMeManagementunitRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetWorkforcemanagementAgentsMeManagementunitRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement agents me managementunit request entity too large response has a 2xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement agents me managementunit request entity too large response has a 3xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement agents me managementunit request entity too large response has a 4xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement agents me managementunit request entity too large response has a 5xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement agents me managementunit request entity too large response a status code equal to that given
func (o *GetWorkforcemanagementAgentsMeManagementunitRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetWorkforcemanagementAgentsMeManagementunitRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/me/managementunit][%d] getWorkforcemanagementAgentsMeManagementunitRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementAgentsMeManagementunitRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/me/managementunit][%d] getWorkforcemanagementAgentsMeManagementunitRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementAgentsMeManagementunitRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAgentsMeManagementunitRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAgentsMeManagementunitUnsupportedMediaType creates a GetWorkforcemanagementAgentsMeManagementunitUnsupportedMediaType with default headers values
func NewGetWorkforcemanagementAgentsMeManagementunitUnsupportedMediaType() *GetWorkforcemanagementAgentsMeManagementunitUnsupportedMediaType {
	return &GetWorkforcemanagementAgentsMeManagementunitUnsupportedMediaType{}
}

/*
GetWorkforcemanagementAgentsMeManagementunitUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWorkforcemanagementAgentsMeManagementunitUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement agents me managementunit unsupported media type response has a 2xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement agents me managementunit unsupported media type response has a 3xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement agents me managementunit unsupported media type response has a 4xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement agents me managementunit unsupported media type response has a 5xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement agents me managementunit unsupported media type response a status code equal to that given
func (o *GetWorkforcemanagementAgentsMeManagementunitUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetWorkforcemanagementAgentsMeManagementunitUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/me/managementunit][%d] getWorkforcemanagementAgentsMeManagementunitUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementAgentsMeManagementunitUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/me/managementunit][%d] getWorkforcemanagementAgentsMeManagementunitUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementAgentsMeManagementunitUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAgentsMeManagementunitUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAgentsMeManagementunitTooManyRequests creates a GetWorkforcemanagementAgentsMeManagementunitTooManyRequests with default headers values
func NewGetWorkforcemanagementAgentsMeManagementunitTooManyRequests() *GetWorkforcemanagementAgentsMeManagementunitTooManyRequests {
	return &GetWorkforcemanagementAgentsMeManagementunitTooManyRequests{}
}

/*
GetWorkforcemanagementAgentsMeManagementunitTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetWorkforcemanagementAgentsMeManagementunitTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement agents me managementunit too many requests response has a 2xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement agents me managementunit too many requests response has a 3xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement agents me managementunit too many requests response has a 4xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement agents me managementunit too many requests response has a 5xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement agents me managementunit too many requests response a status code equal to that given
func (o *GetWorkforcemanagementAgentsMeManagementunitTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetWorkforcemanagementAgentsMeManagementunitTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/me/managementunit][%d] getWorkforcemanagementAgentsMeManagementunitTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementAgentsMeManagementunitTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/me/managementunit][%d] getWorkforcemanagementAgentsMeManagementunitTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementAgentsMeManagementunitTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAgentsMeManagementunitTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAgentsMeManagementunitInternalServerError creates a GetWorkforcemanagementAgentsMeManagementunitInternalServerError with default headers values
func NewGetWorkforcemanagementAgentsMeManagementunitInternalServerError() *GetWorkforcemanagementAgentsMeManagementunitInternalServerError {
	return &GetWorkforcemanagementAgentsMeManagementunitInternalServerError{}
}

/*
GetWorkforcemanagementAgentsMeManagementunitInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWorkforcemanagementAgentsMeManagementunitInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement agents me managementunit internal server error response has a 2xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement agents me managementunit internal server error response has a 3xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement agents me managementunit internal server error response has a 4xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement agents me managementunit internal server error response has a 5xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement agents me managementunit internal server error response a status code equal to that given
func (o *GetWorkforcemanagementAgentsMeManagementunitInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetWorkforcemanagementAgentsMeManagementunitInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/me/managementunit][%d] getWorkforcemanagementAgentsMeManagementunitInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementAgentsMeManagementunitInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/me/managementunit][%d] getWorkforcemanagementAgentsMeManagementunitInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementAgentsMeManagementunitInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAgentsMeManagementunitInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAgentsMeManagementunitServiceUnavailable creates a GetWorkforcemanagementAgentsMeManagementunitServiceUnavailable with default headers values
func NewGetWorkforcemanagementAgentsMeManagementunitServiceUnavailable() *GetWorkforcemanagementAgentsMeManagementunitServiceUnavailable {
	return &GetWorkforcemanagementAgentsMeManagementunitServiceUnavailable{}
}

/*
GetWorkforcemanagementAgentsMeManagementunitServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWorkforcemanagementAgentsMeManagementunitServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement agents me managementunit service unavailable response has a 2xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement agents me managementunit service unavailable response has a 3xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement agents me managementunit service unavailable response has a 4xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement agents me managementunit service unavailable response has a 5xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement agents me managementunit service unavailable response a status code equal to that given
func (o *GetWorkforcemanagementAgentsMeManagementunitServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetWorkforcemanagementAgentsMeManagementunitServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/me/managementunit][%d] getWorkforcemanagementAgentsMeManagementunitServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementAgentsMeManagementunitServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/me/managementunit][%d] getWorkforcemanagementAgentsMeManagementunitServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementAgentsMeManagementunitServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAgentsMeManagementunitServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAgentsMeManagementunitGatewayTimeout creates a GetWorkforcemanagementAgentsMeManagementunitGatewayTimeout with default headers values
func NewGetWorkforcemanagementAgentsMeManagementunitGatewayTimeout() *GetWorkforcemanagementAgentsMeManagementunitGatewayTimeout {
	return &GetWorkforcemanagementAgentsMeManagementunitGatewayTimeout{}
}

/*
GetWorkforcemanagementAgentsMeManagementunitGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetWorkforcemanagementAgentsMeManagementunitGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement agents me managementunit gateway timeout response has a 2xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement agents me managementunit gateway timeout response has a 3xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement agents me managementunit gateway timeout response has a 4xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement agents me managementunit gateway timeout response has a 5xx status code
func (o *GetWorkforcemanagementAgentsMeManagementunitGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement agents me managementunit gateway timeout response a status code equal to that given
func (o *GetWorkforcemanagementAgentsMeManagementunitGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetWorkforcemanagementAgentsMeManagementunitGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/me/managementunit][%d] getWorkforcemanagementAgentsMeManagementunitGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementAgentsMeManagementunitGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/me/managementunit][%d] getWorkforcemanagementAgentsMeManagementunitGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementAgentsMeManagementunitGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAgentsMeManagementunitGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

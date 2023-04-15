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

// GetWorkforcemanagementAgentAdherenceExplanationReader is a Reader for the GetWorkforcemanagementAgentAdherenceExplanation structure.
type GetWorkforcemanagementAgentAdherenceExplanationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkforcemanagementAgentAdherenceExplanationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkforcemanagementAgentAdherenceExplanationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkforcemanagementAgentAdherenceExplanationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWorkforcemanagementAgentAdherenceExplanationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkforcemanagementAgentAdherenceExplanationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkforcemanagementAgentAdherenceExplanationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetWorkforcemanagementAgentAdherenceExplanationRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWorkforcemanagementAgentAdherenceExplanationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWorkforcemanagementAgentAdherenceExplanationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWorkforcemanagementAgentAdherenceExplanationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkforcemanagementAgentAdherenceExplanationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWorkforcemanagementAgentAdherenceExplanationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWorkforcemanagementAgentAdherenceExplanationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkforcemanagementAgentAdherenceExplanationOK creates a GetWorkforcemanagementAgentAdherenceExplanationOK with default headers values
func NewGetWorkforcemanagementAgentAdherenceExplanationOK() *GetWorkforcemanagementAgentAdherenceExplanationOK {
	return &GetWorkforcemanagementAgentAdherenceExplanationOK{}
}

/*
GetWorkforcemanagementAgentAdherenceExplanationOK describes a response with status code 200, with default header values.

successful operation
*/
type GetWorkforcemanagementAgentAdherenceExplanationOK struct {
	Payload *models.AdherenceExplanationResponse
}

// IsSuccess returns true when this get workforcemanagement agent adherence explanation o k response has a 2xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get workforcemanagement agent adherence explanation o k response has a 3xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement agent adherence explanation o k response has a 4xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement agent adherence explanation o k response has a 5xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement agent adherence explanation o k response a status code equal to that given
func (o *GetWorkforcemanagementAgentAdherenceExplanationOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/{agentId}/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAgentAdherenceExplanationOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationOK) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/{agentId}/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAgentAdherenceExplanationOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationOK) GetPayload() *models.AdherenceExplanationResponse {
	return o.Payload
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AdherenceExplanationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAgentAdherenceExplanationBadRequest creates a GetWorkforcemanagementAgentAdherenceExplanationBadRequest with default headers values
func NewGetWorkforcemanagementAgentAdherenceExplanationBadRequest() *GetWorkforcemanagementAgentAdherenceExplanationBadRequest {
	return &GetWorkforcemanagementAgentAdherenceExplanationBadRequest{}
}

/*
GetWorkforcemanagementAgentAdherenceExplanationBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWorkforcemanagementAgentAdherenceExplanationBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement agent adherence explanation bad request response has a 2xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement agent adherence explanation bad request response has a 3xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement agent adherence explanation bad request response has a 4xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement agent adherence explanation bad request response has a 5xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement agent adherence explanation bad request response a status code equal to that given
func (o *GetWorkforcemanagementAgentAdherenceExplanationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/{agentId}/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAgentAdherenceExplanationBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/{agentId}/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAgentAdherenceExplanationBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAgentAdherenceExplanationUnauthorized creates a GetWorkforcemanagementAgentAdherenceExplanationUnauthorized with default headers values
func NewGetWorkforcemanagementAgentAdherenceExplanationUnauthorized() *GetWorkforcemanagementAgentAdherenceExplanationUnauthorized {
	return &GetWorkforcemanagementAgentAdherenceExplanationUnauthorized{}
}

/*
GetWorkforcemanagementAgentAdherenceExplanationUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWorkforcemanagementAgentAdherenceExplanationUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement agent adherence explanation unauthorized response has a 2xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement agent adherence explanation unauthorized response has a 3xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement agent adherence explanation unauthorized response has a 4xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement agent adherence explanation unauthorized response has a 5xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement agent adherence explanation unauthorized response a status code equal to that given
func (o *GetWorkforcemanagementAgentAdherenceExplanationUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/{agentId}/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAgentAdherenceExplanationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/{agentId}/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAgentAdherenceExplanationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAgentAdherenceExplanationForbidden creates a GetWorkforcemanagementAgentAdherenceExplanationForbidden with default headers values
func NewGetWorkforcemanagementAgentAdherenceExplanationForbidden() *GetWorkforcemanagementAgentAdherenceExplanationForbidden {
	return &GetWorkforcemanagementAgentAdherenceExplanationForbidden{}
}

/*
GetWorkforcemanagementAgentAdherenceExplanationForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetWorkforcemanagementAgentAdherenceExplanationForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement agent adherence explanation forbidden response has a 2xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement agent adherence explanation forbidden response has a 3xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement agent adherence explanation forbidden response has a 4xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement agent adherence explanation forbidden response has a 5xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement agent adherence explanation forbidden response a status code equal to that given
func (o *GetWorkforcemanagementAgentAdherenceExplanationForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/{agentId}/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAgentAdherenceExplanationForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/{agentId}/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAgentAdherenceExplanationForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAgentAdherenceExplanationNotFound creates a GetWorkforcemanagementAgentAdherenceExplanationNotFound with default headers values
func NewGetWorkforcemanagementAgentAdherenceExplanationNotFound() *GetWorkforcemanagementAgentAdherenceExplanationNotFound {
	return &GetWorkforcemanagementAgentAdherenceExplanationNotFound{}
}

/*
GetWorkforcemanagementAgentAdherenceExplanationNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetWorkforcemanagementAgentAdherenceExplanationNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement agent adherence explanation not found response has a 2xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement agent adherence explanation not found response has a 3xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement agent adherence explanation not found response has a 4xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement agent adherence explanation not found response has a 5xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement agent adherence explanation not found response a status code equal to that given
func (o *GetWorkforcemanagementAgentAdherenceExplanationNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/{agentId}/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAgentAdherenceExplanationNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/{agentId}/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAgentAdherenceExplanationNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAgentAdherenceExplanationRequestTimeout creates a GetWorkforcemanagementAgentAdherenceExplanationRequestTimeout with default headers values
func NewGetWorkforcemanagementAgentAdherenceExplanationRequestTimeout() *GetWorkforcemanagementAgentAdherenceExplanationRequestTimeout {
	return &GetWorkforcemanagementAgentAdherenceExplanationRequestTimeout{}
}

/*
GetWorkforcemanagementAgentAdherenceExplanationRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetWorkforcemanagementAgentAdherenceExplanationRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement agent adherence explanation request timeout response has a 2xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement agent adherence explanation request timeout response has a 3xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement agent adherence explanation request timeout response has a 4xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement agent adherence explanation request timeout response has a 5xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement agent adherence explanation request timeout response a status code equal to that given
func (o *GetWorkforcemanagementAgentAdherenceExplanationRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/{agentId}/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAgentAdherenceExplanationRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/{agentId}/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAgentAdherenceExplanationRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAgentAdherenceExplanationRequestEntityTooLarge creates a GetWorkforcemanagementAgentAdherenceExplanationRequestEntityTooLarge with default headers values
func NewGetWorkforcemanagementAgentAdherenceExplanationRequestEntityTooLarge() *GetWorkforcemanagementAgentAdherenceExplanationRequestEntityTooLarge {
	return &GetWorkforcemanagementAgentAdherenceExplanationRequestEntityTooLarge{}
}

/*
GetWorkforcemanagementAgentAdherenceExplanationRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetWorkforcemanagementAgentAdherenceExplanationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement agent adherence explanation request entity too large response has a 2xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement agent adherence explanation request entity too large response has a 3xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement agent adherence explanation request entity too large response has a 4xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement agent adherence explanation request entity too large response has a 5xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement agent adherence explanation request entity too large response a status code equal to that given
func (o *GetWorkforcemanagementAgentAdherenceExplanationRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/{agentId}/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAgentAdherenceExplanationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/{agentId}/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAgentAdherenceExplanationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAgentAdherenceExplanationUnsupportedMediaType creates a GetWorkforcemanagementAgentAdherenceExplanationUnsupportedMediaType with default headers values
func NewGetWorkforcemanagementAgentAdherenceExplanationUnsupportedMediaType() *GetWorkforcemanagementAgentAdherenceExplanationUnsupportedMediaType {
	return &GetWorkforcemanagementAgentAdherenceExplanationUnsupportedMediaType{}
}

/*
GetWorkforcemanagementAgentAdherenceExplanationUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWorkforcemanagementAgentAdherenceExplanationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement agent adherence explanation unsupported media type response has a 2xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement agent adherence explanation unsupported media type response has a 3xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement agent adherence explanation unsupported media type response has a 4xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement agent adherence explanation unsupported media type response has a 5xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement agent adherence explanation unsupported media type response a status code equal to that given
func (o *GetWorkforcemanagementAgentAdherenceExplanationUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/{agentId}/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAgentAdherenceExplanationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/{agentId}/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAgentAdherenceExplanationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAgentAdherenceExplanationTooManyRequests creates a GetWorkforcemanagementAgentAdherenceExplanationTooManyRequests with default headers values
func NewGetWorkforcemanagementAgentAdherenceExplanationTooManyRequests() *GetWorkforcemanagementAgentAdherenceExplanationTooManyRequests {
	return &GetWorkforcemanagementAgentAdherenceExplanationTooManyRequests{}
}

/*
GetWorkforcemanagementAgentAdherenceExplanationTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetWorkforcemanagementAgentAdherenceExplanationTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement agent adherence explanation too many requests response has a 2xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement agent adherence explanation too many requests response has a 3xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement agent adherence explanation too many requests response has a 4xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement agent adherence explanation too many requests response has a 5xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement agent adherence explanation too many requests response a status code equal to that given
func (o *GetWorkforcemanagementAgentAdherenceExplanationTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/{agentId}/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAgentAdherenceExplanationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/{agentId}/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAgentAdherenceExplanationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAgentAdherenceExplanationInternalServerError creates a GetWorkforcemanagementAgentAdherenceExplanationInternalServerError with default headers values
func NewGetWorkforcemanagementAgentAdherenceExplanationInternalServerError() *GetWorkforcemanagementAgentAdherenceExplanationInternalServerError {
	return &GetWorkforcemanagementAgentAdherenceExplanationInternalServerError{}
}

/*
GetWorkforcemanagementAgentAdherenceExplanationInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWorkforcemanagementAgentAdherenceExplanationInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement agent adherence explanation internal server error response has a 2xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement agent adherence explanation internal server error response has a 3xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement agent adherence explanation internal server error response has a 4xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement agent adherence explanation internal server error response has a 5xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement agent adherence explanation internal server error response a status code equal to that given
func (o *GetWorkforcemanagementAgentAdherenceExplanationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/{agentId}/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAgentAdherenceExplanationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/{agentId}/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAgentAdherenceExplanationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAgentAdherenceExplanationServiceUnavailable creates a GetWorkforcemanagementAgentAdherenceExplanationServiceUnavailable with default headers values
func NewGetWorkforcemanagementAgentAdherenceExplanationServiceUnavailable() *GetWorkforcemanagementAgentAdherenceExplanationServiceUnavailable {
	return &GetWorkforcemanagementAgentAdherenceExplanationServiceUnavailable{}
}

/*
GetWorkforcemanagementAgentAdherenceExplanationServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWorkforcemanagementAgentAdherenceExplanationServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement agent adherence explanation service unavailable response has a 2xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement agent adherence explanation service unavailable response has a 3xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement agent adherence explanation service unavailable response has a 4xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement agent adherence explanation service unavailable response has a 5xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement agent adherence explanation service unavailable response a status code equal to that given
func (o *GetWorkforcemanagementAgentAdherenceExplanationServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/{agentId}/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAgentAdherenceExplanationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/{agentId}/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAgentAdherenceExplanationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAgentAdherenceExplanationGatewayTimeout creates a GetWorkforcemanagementAgentAdherenceExplanationGatewayTimeout with default headers values
func NewGetWorkforcemanagementAgentAdherenceExplanationGatewayTimeout() *GetWorkforcemanagementAgentAdherenceExplanationGatewayTimeout {
	return &GetWorkforcemanagementAgentAdherenceExplanationGatewayTimeout{}
}

/*
GetWorkforcemanagementAgentAdherenceExplanationGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetWorkforcemanagementAgentAdherenceExplanationGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement agent adherence explanation gateway timeout response has a 2xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement agent adherence explanation gateway timeout response has a 3xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement agent adherence explanation gateway timeout response has a 4xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement agent adherence explanation gateway timeout response has a 5xx status code
func (o *GetWorkforcemanagementAgentAdherenceExplanationGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement agent adherence explanation gateway timeout response a status code equal to that given
func (o *GetWorkforcemanagementAgentAdherenceExplanationGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/{agentId}/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAgentAdherenceExplanationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/{agentId}/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAgentAdherenceExplanationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAgentAdherenceExplanationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

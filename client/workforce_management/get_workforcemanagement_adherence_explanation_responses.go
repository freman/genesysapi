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

// GetWorkforcemanagementAdherenceExplanationReader is a Reader for the GetWorkforcemanagementAdherenceExplanation structure.
type GetWorkforcemanagementAdherenceExplanationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkforcemanagementAdherenceExplanationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkforcemanagementAdherenceExplanationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkforcemanagementAdherenceExplanationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWorkforcemanagementAdherenceExplanationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkforcemanagementAdherenceExplanationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkforcemanagementAdherenceExplanationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetWorkforcemanagementAdherenceExplanationRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWorkforcemanagementAdherenceExplanationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWorkforcemanagementAdherenceExplanationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWorkforcemanagementAdherenceExplanationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkforcemanagementAdherenceExplanationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWorkforcemanagementAdherenceExplanationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWorkforcemanagementAdherenceExplanationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkforcemanagementAdherenceExplanationOK creates a GetWorkforcemanagementAdherenceExplanationOK with default headers values
func NewGetWorkforcemanagementAdherenceExplanationOK() *GetWorkforcemanagementAdherenceExplanationOK {
	return &GetWorkforcemanagementAdherenceExplanationOK{}
}

/*
GetWorkforcemanagementAdherenceExplanationOK describes a response with status code 200, with default header values.

successful operation
*/
type GetWorkforcemanagementAdherenceExplanationOK struct {
	Payload *models.AdherenceExplanationResponse
}

// IsSuccess returns true when this get workforcemanagement adherence explanation o k response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceExplanationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get workforcemanagement adherence explanation o k response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceExplanationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence explanation o k response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceExplanationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement adherence explanation o k response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceExplanationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adherence explanation o k response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceExplanationOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetWorkforcemanagementAdherenceExplanationOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAdherenceExplanationOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationOK) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAdherenceExplanationOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationOK) GetPayload() *models.AdherenceExplanationResponse {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceExplanationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AdherenceExplanationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceExplanationBadRequest creates a GetWorkforcemanagementAdherenceExplanationBadRequest with default headers values
func NewGetWorkforcemanagementAdherenceExplanationBadRequest() *GetWorkforcemanagementAdherenceExplanationBadRequest {
	return &GetWorkforcemanagementAdherenceExplanationBadRequest{}
}

/*
GetWorkforcemanagementAdherenceExplanationBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWorkforcemanagementAdherenceExplanationBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence explanation bad request response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceExplanationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence explanation bad request response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceExplanationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence explanation bad request response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceExplanationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement adherence explanation bad request response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceExplanationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adherence explanation bad request response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceExplanationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetWorkforcemanagementAdherenceExplanationBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAdherenceExplanationBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAdherenceExplanationBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceExplanationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceExplanationUnauthorized creates a GetWorkforcemanagementAdherenceExplanationUnauthorized with default headers values
func NewGetWorkforcemanagementAdherenceExplanationUnauthorized() *GetWorkforcemanagementAdherenceExplanationUnauthorized {
	return &GetWorkforcemanagementAdherenceExplanationUnauthorized{}
}

/*
GetWorkforcemanagementAdherenceExplanationUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWorkforcemanagementAdherenceExplanationUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence explanation unauthorized response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceExplanationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence explanation unauthorized response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceExplanationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence explanation unauthorized response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceExplanationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement adherence explanation unauthorized response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceExplanationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adherence explanation unauthorized response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceExplanationUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetWorkforcemanagementAdherenceExplanationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAdherenceExplanationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAdherenceExplanationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceExplanationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceExplanationForbidden creates a GetWorkforcemanagementAdherenceExplanationForbidden with default headers values
func NewGetWorkforcemanagementAdherenceExplanationForbidden() *GetWorkforcemanagementAdherenceExplanationForbidden {
	return &GetWorkforcemanagementAdherenceExplanationForbidden{}
}

/*
GetWorkforcemanagementAdherenceExplanationForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetWorkforcemanagementAdherenceExplanationForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence explanation forbidden response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceExplanationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence explanation forbidden response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceExplanationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence explanation forbidden response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceExplanationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement adherence explanation forbidden response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceExplanationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adherence explanation forbidden response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceExplanationForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetWorkforcemanagementAdherenceExplanationForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAdherenceExplanationForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAdherenceExplanationForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceExplanationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceExplanationNotFound creates a GetWorkforcemanagementAdherenceExplanationNotFound with default headers values
func NewGetWorkforcemanagementAdherenceExplanationNotFound() *GetWorkforcemanagementAdherenceExplanationNotFound {
	return &GetWorkforcemanagementAdherenceExplanationNotFound{}
}

/*
GetWorkforcemanagementAdherenceExplanationNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetWorkforcemanagementAdherenceExplanationNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence explanation not found response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceExplanationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence explanation not found response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceExplanationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence explanation not found response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceExplanationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement adherence explanation not found response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceExplanationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adherence explanation not found response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceExplanationNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetWorkforcemanagementAdherenceExplanationNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAdherenceExplanationNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAdherenceExplanationNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceExplanationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceExplanationRequestTimeout creates a GetWorkforcemanagementAdherenceExplanationRequestTimeout with default headers values
func NewGetWorkforcemanagementAdherenceExplanationRequestTimeout() *GetWorkforcemanagementAdherenceExplanationRequestTimeout {
	return &GetWorkforcemanagementAdherenceExplanationRequestTimeout{}
}

/*
GetWorkforcemanagementAdherenceExplanationRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetWorkforcemanagementAdherenceExplanationRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence explanation request timeout response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceExplanationRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence explanation request timeout response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceExplanationRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence explanation request timeout response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceExplanationRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement adherence explanation request timeout response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceExplanationRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adherence explanation request timeout response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceExplanationRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetWorkforcemanagementAdherenceExplanationRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAdherenceExplanationRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAdherenceExplanationRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceExplanationRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceExplanationRequestEntityTooLarge creates a GetWorkforcemanagementAdherenceExplanationRequestEntityTooLarge with default headers values
func NewGetWorkforcemanagementAdherenceExplanationRequestEntityTooLarge() *GetWorkforcemanagementAdherenceExplanationRequestEntityTooLarge {
	return &GetWorkforcemanagementAdherenceExplanationRequestEntityTooLarge{}
}

/*
GetWorkforcemanagementAdherenceExplanationRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetWorkforcemanagementAdherenceExplanationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence explanation request entity too large response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceExplanationRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence explanation request entity too large response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceExplanationRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence explanation request entity too large response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceExplanationRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement adherence explanation request entity too large response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceExplanationRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adherence explanation request entity too large response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceExplanationRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetWorkforcemanagementAdherenceExplanationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAdherenceExplanationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAdherenceExplanationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceExplanationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceExplanationUnsupportedMediaType creates a GetWorkforcemanagementAdherenceExplanationUnsupportedMediaType with default headers values
func NewGetWorkforcemanagementAdherenceExplanationUnsupportedMediaType() *GetWorkforcemanagementAdherenceExplanationUnsupportedMediaType {
	return &GetWorkforcemanagementAdherenceExplanationUnsupportedMediaType{}
}

/*
GetWorkforcemanagementAdherenceExplanationUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWorkforcemanagementAdherenceExplanationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence explanation unsupported media type response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceExplanationUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence explanation unsupported media type response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceExplanationUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence explanation unsupported media type response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceExplanationUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement adherence explanation unsupported media type response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceExplanationUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adherence explanation unsupported media type response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceExplanationUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetWorkforcemanagementAdherenceExplanationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAdherenceExplanationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAdherenceExplanationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceExplanationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceExplanationTooManyRequests creates a GetWorkforcemanagementAdherenceExplanationTooManyRequests with default headers values
func NewGetWorkforcemanagementAdherenceExplanationTooManyRequests() *GetWorkforcemanagementAdherenceExplanationTooManyRequests {
	return &GetWorkforcemanagementAdherenceExplanationTooManyRequests{}
}

/*
GetWorkforcemanagementAdherenceExplanationTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetWorkforcemanagementAdherenceExplanationTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence explanation too many requests response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceExplanationTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence explanation too many requests response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceExplanationTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence explanation too many requests response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceExplanationTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement adherence explanation too many requests response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceExplanationTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement adherence explanation too many requests response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceExplanationTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetWorkforcemanagementAdherenceExplanationTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAdherenceExplanationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAdherenceExplanationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceExplanationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceExplanationInternalServerError creates a GetWorkforcemanagementAdherenceExplanationInternalServerError with default headers values
func NewGetWorkforcemanagementAdherenceExplanationInternalServerError() *GetWorkforcemanagementAdherenceExplanationInternalServerError {
	return &GetWorkforcemanagementAdherenceExplanationInternalServerError{}
}

/*
GetWorkforcemanagementAdherenceExplanationInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWorkforcemanagementAdherenceExplanationInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence explanation internal server error response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceExplanationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence explanation internal server error response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceExplanationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence explanation internal server error response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceExplanationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement adherence explanation internal server error response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceExplanationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement adherence explanation internal server error response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceExplanationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetWorkforcemanagementAdherenceExplanationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAdherenceExplanationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAdherenceExplanationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceExplanationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceExplanationServiceUnavailable creates a GetWorkforcemanagementAdherenceExplanationServiceUnavailable with default headers values
func NewGetWorkforcemanagementAdherenceExplanationServiceUnavailable() *GetWorkforcemanagementAdherenceExplanationServiceUnavailable {
	return &GetWorkforcemanagementAdherenceExplanationServiceUnavailable{}
}

/*
GetWorkforcemanagementAdherenceExplanationServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWorkforcemanagementAdherenceExplanationServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence explanation service unavailable response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceExplanationServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence explanation service unavailable response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceExplanationServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence explanation service unavailable response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceExplanationServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement adherence explanation service unavailable response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceExplanationServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement adherence explanation service unavailable response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceExplanationServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetWorkforcemanagementAdherenceExplanationServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAdherenceExplanationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAdherenceExplanationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceExplanationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceExplanationGatewayTimeout creates a GetWorkforcemanagementAdherenceExplanationGatewayTimeout with default headers values
func NewGetWorkforcemanagementAdherenceExplanationGatewayTimeout() *GetWorkforcemanagementAdherenceExplanationGatewayTimeout {
	return &GetWorkforcemanagementAdherenceExplanationGatewayTimeout{}
}

/*
GetWorkforcemanagementAdherenceExplanationGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetWorkforcemanagementAdherenceExplanationGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement adherence explanation gateway timeout response has a 2xx status code
func (o *GetWorkforcemanagementAdherenceExplanationGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement adherence explanation gateway timeout response has a 3xx status code
func (o *GetWorkforcemanagementAdherenceExplanationGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement adherence explanation gateway timeout response has a 4xx status code
func (o *GetWorkforcemanagementAdherenceExplanationGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement adherence explanation gateway timeout response has a 5xx status code
func (o *GetWorkforcemanagementAdherenceExplanationGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement adherence explanation gateway timeout response a status code equal to that given
func (o *GetWorkforcemanagementAdherenceExplanationGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetWorkforcemanagementAdherenceExplanationGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAdherenceExplanationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence/explanations/{explanationId}][%d] getWorkforcemanagementAdherenceExplanationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceExplanationGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceExplanationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

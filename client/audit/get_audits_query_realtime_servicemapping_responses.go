// Code generated by go-swagger; DO NOT EDIT.

package audit

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetAuditsQueryRealtimeServicemappingReader is a Reader for the GetAuditsQueryRealtimeServicemapping structure.
type GetAuditsQueryRealtimeServicemappingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuditsQueryRealtimeServicemappingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAuditsQueryRealtimeServicemappingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAuditsQueryRealtimeServicemappingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAuditsQueryRealtimeServicemappingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAuditsQueryRealtimeServicemappingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAuditsQueryRealtimeServicemappingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetAuditsQueryRealtimeServicemappingRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetAuditsQueryRealtimeServicemappingRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetAuditsQueryRealtimeServicemappingUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAuditsQueryRealtimeServicemappingTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAuditsQueryRealtimeServicemappingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAuditsQueryRealtimeServicemappingServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAuditsQueryRealtimeServicemappingGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAuditsQueryRealtimeServicemappingOK creates a GetAuditsQueryRealtimeServicemappingOK with default headers values
func NewGetAuditsQueryRealtimeServicemappingOK() *GetAuditsQueryRealtimeServicemappingOK {
	return &GetAuditsQueryRealtimeServicemappingOK{}
}

/*
GetAuditsQueryRealtimeServicemappingOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAuditsQueryRealtimeServicemappingOK struct {
	Payload *models.AuditQueryServiceMapping
}

// IsSuccess returns true when this get audits query realtime servicemapping o k response has a 2xx status code
func (o *GetAuditsQueryRealtimeServicemappingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get audits query realtime servicemapping o k response has a 3xx status code
func (o *GetAuditsQueryRealtimeServicemappingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get audits query realtime servicemapping o k response has a 4xx status code
func (o *GetAuditsQueryRealtimeServicemappingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get audits query realtime servicemapping o k response has a 5xx status code
func (o *GetAuditsQueryRealtimeServicemappingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get audits query realtime servicemapping o k response a status code equal to that given
func (o *GetAuditsQueryRealtimeServicemappingOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAuditsQueryRealtimeServicemappingOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/realtime/servicemapping][%d] getAuditsQueryRealtimeServicemappingOK  %+v", 200, o.Payload)
}

func (o *GetAuditsQueryRealtimeServicemappingOK) String() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/realtime/servicemapping][%d] getAuditsQueryRealtimeServicemappingOK  %+v", 200, o.Payload)
}

func (o *GetAuditsQueryRealtimeServicemappingOK) GetPayload() *models.AuditQueryServiceMapping {
	return o.Payload
}

func (o *GetAuditsQueryRealtimeServicemappingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuditQueryServiceMapping)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditsQueryRealtimeServicemappingBadRequest creates a GetAuditsQueryRealtimeServicemappingBadRequest with default headers values
func NewGetAuditsQueryRealtimeServicemappingBadRequest() *GetAuditsQueryRealtimeServicemappingBadRequest {
	return &GetAuditsQueryRealtimeServicemappingBadRequest{}
}

/*
GetAuditsQueryRealtimeServicemappingBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAuditsQueryRealtimeServicemappingBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get audits query realtime servicemapping bad request response has a 2xx status code
func (o *GetAuditsQueryRealtimeServicemappingBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get audits query realtime servicemapping bad request response has a 3xx status code
func (o *GetAuditsQueryRealtimeServicemappingBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get audits query realtime servicemapping bad request response has a 4xx status code
func (o *GetAuditsQueryRealtimeServicemappingBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get audits query realtime servicemapping bad request response has a 5xx status code
func (o *GetAuditsQueryRealtimeServicemappingBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get audits query realtime servicemapping bad request response a status code equal to that given
func (o *GetAuditsQueryRealtimeServicemappingBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetAuditsQueryRealtimeServicemappingBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/realtime/servicemapping][%d] getAuditsQueryRealtimeServicemappingBadRequest  %+v", 400, o.Payload)
}

func (o *GetAuditsQueryRealtimeServicemappingBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/realtime/servicemapping][%d] getAuditsQueryRealtimeServicemappingBadRequest  %+v", 400, o.Payload)
}

func (o *GetAuditsQueryRealtimeServicemappingBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuditsQueryRealtimeServicemappingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditsQueryRealtimeServicemappingUnauthorized creates a GetAuditsQueryRealtimeServicemappingUnauthorized with default headers values
func NewGetAuditsQueryRealtimeServicemappingUnauthorized() *GetAuditsQueryRealtimeServicemappingUnauthorized {
	return &GetAuditsQueryRealtimeServicemappingUnauthorized{}
}

/*
GetAuditsQueryRealtimeServicemappingUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAuditsQueryRealtimeServicemappingUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get audits query realtime servicemapping unauthorized response has a 2xx status code
func (o *GetAuditsQueryRealtimeServicemappingUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get audits query realtime servicemapping unauthorized response has a 3xx status code
func (o *GetAuditsQueryRealtimeServicemappingUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get audits query realtime servicemapping unauthorized response has a 4xx status code
func (o *GetAuditsQueryRealtimeServicemappingUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get audits query realtime servicemapping unauthorized response has a 5xx status code
func (o *GetAuditsQueryRealtimeServicemappingUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get audits query realtime servicemapping unauthorized response a status code equal to that given
func (o *GetAuditsQueryRealtimeServicemappingUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetAuditsQueryRealtimeServicemappingUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/realtime/servicemapping][%d] getAuditsQueryRealtimeServicemappingUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAuditsQueryRealtimeServicemappingUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/realtime/servicemapping][%d] getAuditsQueryRealtimeServicemappingUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAuditsQueryRealtimeServicemappingUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuditsQueryRealtimeServicemappingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditsQueryRealtimeServicemappingForbidden creates a GetAuditsQueryRealtimeServicemappingForbidden with default headers values
func NewGetAuditsQueryRealtimeServicemappingForbidden() *GetAuditsQueryRealtimeServicemappingForbidden {
	return &GetAuditsQueryRealtimeServicemappingForbidden{}
}

/*
GetAuditsQueryRealtimeServicemappingForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetAuditsQueryRealtimeServicemappingForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get audits query realtime servicemapping forbidden response has a 2xx status code
func (o *GetAuditsQueryRealtimeServicemappingForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get audits query realtime servicemapping forbidden response has a 3xx status code
func (o *GetAuditsQueryRealtimeServicemappingForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get audits query realtime servicemapping forbidden response has a 4xx status code
func (o *GetAuditsQueryRealtimeServicemappingForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get audits query realtime servicemapping forbidden response has a 5xx status code
func (o *GetAuditsQueryRealtimeServicemappingForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get audits query realtime servicemapping forbidden response a status code equal to that given
func (o *GetAuditsQueryRealtimeServicemappingForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetAuditsQueryRealtimeServicemappingForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/realtime/servicemapping][%d] getAuditsQueryRealtimeServicemappingForbidden  %+v", 403, o.Payload)
}

func (o *GetAuditsQueryRealtimeServicemappingForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/realtime/servicemapping][%d] getAuditsQueryRealtimeServicemappingForbidden  %+v", 403, o.Payload)
}

func (o *GetAuditsQueryRealtimeServicemappingForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuditsQueryRealtimeServicemappingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditsQueryRealtimeServicemappingNotFound creates a GetAuditsQueryRealtimeServicemappingNotFound with default headers values
func NewGetAuditsQueryRealtimeServicemappingNotFound() *GetAuditsQueryRealtimeServicemappingNotFound {
	return &GetAuditsQueryRealtimeServicemappingNotFound{}
}

/*
GetAuditsQueryRealtimeServicemappingNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetAuditsQueryRealtimeServicemappingNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get audits query realtime servicemapping not found response has a 2xx status code
func (o *GetAuditsQueryRealtimeServicemappingNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get audits query realtime servicemapping not found response has a 3xx status code
func (o *GetAuditsQueryRealtimeServicemappingNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get audits query realtime servicemapping not found response has a 4xx status code
func (o *GetAuditsQueryRealtimeServicemappingNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get audits query realtime servicemapping not found response has a 5xx status code
func (o *GetAuditsQueryRealtimeServicemappingNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get audits query realtime servicemapping not found response a status code equal to that given
func (o *GetAuditsQueryRealtimeServicemappingNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetAuditsQueryRealtimeServicemappingNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/realtime/servicemapping][%d] getAuditsQueryRealtimeServicemappingNotFound  %+v", 404, o.Payload)
}

func (o *GetAuditsQueryRealtimeServicemappingNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/realtime/servicemapping][%d] getAuditsQueryRealtimeServicemappingNotFound  %+v", 404, o.Payload)
}

func (o *GetAuditsQueryRealtimeServicemappingNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuditsQueryRealtimeServicemappingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditsQueryRealtimeServicemappingRequestTimeout creates a GetAuditsQueryRealtimeServicemappingRequestTimeout with default headers values
func NewGetAuditsQueryRealtimeServicemappingRequestTimeout() *GetAuditsQueryRealtimeServicemappingRequestTimeout {
	return &GetAuditsQueryRealtimeServicemappingRequestTimeout{}
}

/*
GetAuditsQueryRealtimeServicemappingRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetAuditsQueryRealtimeServicemappingRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get audits query realtime servicemapping request timeout response has a 2xx status code
func (o *GetAuditsQueryRealtimeServicemappingRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get audits query realtime servicemapping request timeout response has a 3xx status code
func (o *GetAuditsQueryRealtimeServicemappingRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get audits query realtime servicemapping request timeout response has a 4xx status code
func (o *GetAuditsQueryRealtimeServicemappingRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get audits query realtime servicemapping request timeout response has a 5xx status code
func (o *GetAuditsQueryRealtimeServicemappingRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get audits query realtime servicemapping request timeout response a status code equal to that given
func (o *GetAuditsQueryRealtimeServicemappingRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetAuditsQueryRealtimeServicemappingRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/realtime/servicemapping][%d] getAuditsQueryRealtimeServicemappingRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAuditsQueryRealtimeServicemappingRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/realtime/servicemapping][%d] getAuditsQueryRealtimeServicemappingRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAuditsQueryRealtimeServicemappingRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuditsQueryRealtimeServicemappingRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditsQueryRealtimeServicemappingRequestEntityTooLarge creates a GetAuditsQueryRealtimeServicemappingRequestEntityTooLarge with default headers values
func NewGetAuditsQueryRealtimeServicemappingRequestEntityTooLarge() *GetAuditsQueryRealtimeServicemappingRequestEntityTooLarge {
	return &GetAuditsQueryRealtimeServicemappingRequestEntityTooLarge{}
}

/*
GetAuditsQueryRealtimeServicemappingRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetAuditsQueryRealtimeServicemappingRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get audits query realtime servicemapping request entity too large response has a 2xx status code
func (o *GetAuditsQueryRealtimeServicemappingRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get audits query realtime servicemapping request entity too large response has a 3xx status code
func (o *GetAuditsQueryRealtimeServicemappingRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get audits query realtime servicemapping request entity too large response has a 4xx status code
func (o *GetAuditsQueryRealtimeServicemappingRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get audits query realtime servicemapping request entity too large response has a 5xx status code
func (o *GetAuditsQueryRealtimeServicemappingRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get audits query realtime servicemapping request entity too large response a status code equal to that given
func (o *GetAuditsQueryRealtimeServicemappingRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetAuditsQueryRealtimeServicemappingRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/realtime/servicemapping][%d] getAuditsQueryRealtimeServicemappingRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAuditsQueryRealtimeServicemappingRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/realtime/servicemapping][%d] getAuditsQueryRealtimeServicemappingRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAuditsQueryRealtimeServicemappingRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuditsQueryRealtimeServicemappingRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditsQueryRealtimeServicemappingUnsupportedMediaType creates a GetAuditsQueryRealtimeServicemappingUnsupportedMediaType with default headers values
func NewGetAuditsQueryRealtimeServicemappingUnsupportedMediaType() *GetAuditsQueryRealtimeServicemappingUnsupportedMediaType {
	return &GetAuditsQueryRealtimeServicemappingUnsupportedMediaType{}
}

/*
GetAuditsQueryRealtimeServicemappingUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAuditsQueryRealtimeServicemappingUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get audits query realtime servicemapping unsupported media type response has a 2xx status code
func (o *GetAuditsQueryRealtimeServicemappingUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get audits query realtime servicemapping unsupported media type response has a 3xx status code
func (o *GetAuditsQueryRealtimeServicemappingUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get audits query realtime servicemapping unsupported media type response has a 4xx status code
func (o *GetAuditsQueryRealtimeServicemappingUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get audits query realtime servicemapping unsupported media type response has a 5xx status code
func (o *GetAuditsQueryRealtimeServicemappingUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get audits query realtime servicemapping unsupported media type response a status code equal to that given
func (o *GetAuditsQueryRealtimeServicemappingUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetAuditsQueryRealtimeServicemappingUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/realtime/servicemapping][%d] getAuditsQueryRealtimeServicemappingUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAuditsQueryRealtimeServicemappingUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/realtime/servicemapping][%d] getAuditsQueryRealtimeServicemappingUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAuditsQueryRealtimeServicemappingUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuditsQueryRealtimeServicemappingUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditsQueryRealtimeServicemappingTooManyRequests creates a GetAuditsQueryRealtimeServicemappingTooManyRequests with default headers values
func NewGetAuditsQueryRealtimeServicemappingTooManyRequests() *GetAuditsQueryRealtimeServicemappingTooManyRequests {
	return &GetAuditsQueryRealtimeServicemappingTooManyRequests{}
}

/*
GetAuditsQueryRealtimeServicemappingTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetAuditsQueryRealtimeServicemappingTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get audits query realtime servicemapping too many requests response has a 2xx status code
func (o *GetAuditsQueryRealtimeServicemappingTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get audits query realtime servicemapping too many requests response has a 3xx status code
func (o *GetAuditsQueryRealtimeServicemappingTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get audits query realtime servicemapping too many requests response has a 4xx status code
func (o *GetAuditsQueryRealtimeServicemappingTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get audits query realtime servicemapping too many requests response has a 5xx status code
func (o *GetAuditsQueryRealtimeServicemappingTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get audits query realtime servicemapping too many requests response a status code equal to that given
func (o *GetAuditsQueryRealtimeServicemappingTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetAuditsQueryRealtimeServicemappingTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/realtime/servicemapping][%d] getAuditsQueryRealtimeServicemappingTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAuditsQueryRealtimeServicemappingTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/realtime/servicemapping][%d] getAuditsQueryRealtimeServicemappingTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAuditsQueryRealtimeServicemappingTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuditsQueryRealtimeServicemappingTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditsQueryRealtimeServicemappingInternalServerError creates a GetAuditsQueryRealtimeServicemappingInternalServerError with default headers values
func NewGetAuditsQueryRealtimeServicemappingInternalServerError() *GetAuditsQueryRealtimeServicemappingInternalServerError {
	return &GetAuditsQueryRealtimeServicemappingInternalServerError{}
}

/*
GetAuditsQueryRealtimeServicemappingInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAuditsQueryRealtimeServicemappingInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get audits query realtime servicemapping internal server error response has a 2xx status code
func (o *GetAuditsQueryRealtimeServicemappingInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get audits query realtime servicemapping internal server error response has a 3xx status code
func (o *GetAuditsQueryRealtimeServicemappingInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get audits query realtime servicemapping internal server error response has a 4xx status code
func (o *GetAuditsQueryRealtimeServicemappingInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get audits query realtime servicemapping internal server error response has a 5xx status code
func (o *GetAuditsQueryRealtimeServicemappingInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get audits query realtime servicemapping internal server error response a status code equal to that given
func (o *GetAuditsQueryRealtimeServicemappingInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetAuditsQueryRealtimeServicemappingInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/realtime/servicemapping][%d] getAuditsQueryRealtimeServicemappingInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAuditsQueryRealtimeServicemappingInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/realtime/servicemapping][%d] getAuditsQueryRealtimeServicemappingInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAuditsQueryRealtimeServicemappingInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuditsQueryRealtimeServicemappingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditsQueryRealtimeServicemappingServiceUnavailable creates a GetAuditsQueryRealtimeServicemappingServiceUnavailable with default headers values
func NewGetAuditsQueryRealtimeServicemappingServiceUnavailable() *GetAuditsQueryRealtimeServicemappingServiceUnavailable {
	return &GetAuditsQueryRealtimeServicemappingServiceUnavailable{}
}

/*
GetAuditsQueryRealtimeServicemappingServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAuditsQueryRealtimeServicemappingServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get audits query realtime servicemapping service unavailable response has a 2xx status code
func (o *GetAuditsQueryRealtimeServicemappingServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get audits query realtime servicemapping service unavailable response has a 3xx status code
func (o *GetAuditsQueryRealtimeServicemappingServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get audits query realtime servicemapping service unavailable response has a 4xx status code
func (o *GetAuditsQueryRealtimeServicemappingServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get audits query realtime servicemapping service unavailable response has a 5xx status code
func (o *GetAuditsQueryRealtimeServicemappingServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get audits query realtime servicemapping service unavailable response a status code equal to that given
func (o *GetAuditsQueryRealtimeServicemappingServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetAuditsQueryRealtimeServicemappingServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/realtime/servicemapping][%d] getAuditsQueryRealtimeServicemappingServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAuditsQueryRealtimeServicemappingServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/realtime/servicemapping][%d] getAuditsQueryRealtimeServicemappingServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAuditsQueryRealtimeServicemappingServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuditsQueryRealtimeServicemappingServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditsQueryRealtimeServicemappingGatewayTimeout creates a GetAuditsQueryRealtimeServicemappingGatewayTimeout with default headers values
func NewGetAuditsQueryRealtimeServicemappingGatewayTimeout() *GetAuditsQueryRealtimeServicemappingGatewayTimeout {
	return &GetAuditsQueryRealtimeServicemappingGatewayTimeout{}
}

/*
GetAuditsQueryRealtimeServicemappingGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetAuditsQueryRealtimeServicemappingGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get audits query realtime servicemapping gateway timeout response has a 2xx status code
func (o *GetAuditsQueryRealtimeServicemappingGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get audits query realtime servicemapping gateway timeout response has a 3xx status code
func (o *GetAuditsQueryRealtimeServicemappingGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get audits query realtime servicemapping gateway timeout response has a 4xx status code
func (o *GetAuditsQueryRealtimeServicemappingGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get audits query realtime servicemapping gateway timeout response has a 5xx status code
func (o *GetAuditsQueryRealtimeServicemappingGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get audits query realtime servicemapping gateway timeout response a status code equal to that given
func (o *GetAuditsQueryRealtimeServicemappingGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetAuditsQueryRealtimeServicemappingGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/realtime/servicemapping][%d] getAuditsQueryRealtimeServicemappingGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAuditsQueryRealtimeServicemappingGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/realtime/servicemapping][%d] getAuditsQueryRealtimeServicemappingGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAuditsQueryRealtimeServicemappingGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuditsQueryRealtimeServicemappingGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

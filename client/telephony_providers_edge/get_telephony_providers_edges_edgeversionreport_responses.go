// Code generated by go-swagger; DO NOT EDIT.

package telephony_providers_edge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetTelephonyProvidersEdgesEdgeversionreportReader is a Reader for the GetTelephonyProvidersEdgesEdgeversionreport structure.
type GetTelephonyProvidersEdgesEdgeversionreportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonyProvidersEdgesEdgeversionreportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonyProvidersEdgesEdgeversionreportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonyProvidersEdgesEdgeversionreportBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonyProvidersEdgesEdgeversionreportUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonyProvidersEdgesEdgeversionreportForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonyProvidersEdgesEdgeversionreportNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetTelephonyProvidersEdgesEdgeversionreportRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonyProvidersEdgesEdgeversionreportRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonyProvidersEdgesEdgeversionreportUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonyProvidersEdgesEdgeversionreportTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonyProvidersEdgesEdgeversionreportInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonyProvidersEdgesEdgeversionreportServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonyProvidersEdgesEdgeversionreportGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonyProvidersEdgesEdgeversionreportOK creates a GetTelephonyProvidersEdgesEdgeversionreportOK with default headers values
func NewGetTelephonyProvidersEdgesEdgeversionreportOK() *GetTelephonyProvidersEdgesEdgeversionreportOK {
	return &GetTelephonyProvidersEdgesEdgeversionreportOK{}
}

/*
GetTelephonyProvidersEdgesEdgeversionreportOK describes a response with status code 200, with default header values.

successful operation
*/
type GetTelephonyProvidersEdgesEdgeversionreportOK struct {
	Payload *models.EdgeVersionReport
}

// IsSuccess returns true when this get telephony providers edges edgeversionreport o k response has a 2xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get telephony providers edges edgeversionreport o k response has a 3xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges edgeversionreport o k response has a 4xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edges edgeversionreport o k response has a 5xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges edgeversionreport o k response a status code equal to that given
func (o *GetTelephonyProvidersEdgesEdgeversionreportOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgeversionreport][%d] getTelephonyProvidersEdgesEdgeversionreportOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportOK) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgeversionreport][%d] getTelephonyProvidersEdgesEdgeversionreportOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportOK) GetPayload() *models.EdgeVersionReport {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EdgeVersionReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesEdgeversionreportBadRequest creates a GetTelephonyProvidersEdgesEdgeversionreportBadRequest with default headers values
func NewGetTelephonyProvidersEdgesEdgeversionreportBadRequest() *GetTelephonyProvidersEdgesEdgeversionreportBadRequest {
	return &GetTelephonyProvidersEdgesEdgeversionreportBadRequest{}
}

/*
GetTelephonyProvidersEdgesEdgeversionreportBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyProvidersEdgesEdgeversionreportBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges edgeversionreport bad request response has a 2xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges edgeversionreport bad request response has a 3xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges edgeversionreport bad request response has a 4xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges edgeversionreport bad request response has a 5xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges edgeversionreport bad request response a status code equal to that given
func (o *GetTelephonyProvidersEdgesEdgeversionreportBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgeversionreport][%d] getTelephonyProvidersEdgesEdgeversionreportBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgeversionreport][%d] getTelephonyProvidersEdgesEdgeversionreportBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesEdgeversionreportUnauthorized creates a GetTelephonyProvidersEdgesEdgeversionreportUnauthorized with default headers values
func NewGetTelephonyProvidersEdgesEdgeversionreportUnauthorized() *GetTelephonyProvidersEdgesEdgeversionreportUnauthorized {
	return &GetTelephonyProvidersEdgesEdgeversionreportUnauthorized{}
}

/*
GetTelephonyProvidersEdgesEdgeversionreportUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyProvidersEdgesEdgeversionreportUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges edgeversionreport unauthorized response has a 2xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges edgeversionreport unauthorized response has a 3xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges edgeversionreport unauthorized response has a 4xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges edgeversionreport unauthorized response has a 5xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges edgeversionreport unauthorized response a status code equal to that given
func (o *GetTelephonyProvidersEdgesEdgeversionreportUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgeversionreport][%d] getTelephonyProvidersEdgesEdgeversionreportUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgeversionreport][%d] getTelephonyProvidersEdgesEdgeversionreportUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesEdgeversionreportForbidden creates a GetTelephonyProvidersEdgesEdgeversionreportForbidden with default headers values
func NewGetTelephonyProvidersEdgesEdgeversionreportForbidden() *GetTelephonyProvidersEdgesEdgeversionreportForbidden {
	return &GetTelephonyProvidersEdgesEdgeversionreportForbidden{}
}

/*
GetTelephonyProvidersEdgesEdgeversionreportForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyProvidersEdgesEdgeversionreportForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges edgeversionreport forbidden response has a 2xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges edgeversionreport forbidden response has a 3xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges edgeversionreport forbidden response has a 4xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges edgeversionreport forbidden response has a 5xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges edgeversionreport forbidden response a status code equal to that given
func (o *GetTelephonyProvidersEdgesEdgeversionreportForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgeversionreport][%d] getTelephonyProvidersEdgesEdgeversionreportForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgeversionreport][%d] getTelephonyProvidersEdgesEdgeversionreportForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesEdgeversionreportNotFound creates a GetTelephonyProvidersEdgesEdgeversionreportNotFound with default headers values
func NewGetTelephonyProvidersEdgesEdgeversionreportNotFound() *GetTelephonyProvidersEdgesEdgeversionreportNotFound {
	return &GetTelephonyProvidersEdgesEdgeversionreportNotFound{}
}

/*
GetTelephonyProvidersEdgesEdgeversionreportNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetTelephonyProvidersEdgesEdgeversionreportNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges edgeversionreport not found response has a 2xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges edgeversionreport not found response has a 3xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges edgeversionreport not found response has a 4xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges edgeversionreport not found response has a 5xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges edgeversionreport not found response a status code equal to that given
func (o *GetTelephonyProvidersEdgesEdgeversionreportNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgeversionreport][%d] getTelephonyProvidersEdgesEdgeversionreportNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgeversionreport][%d] getTelephonyProvidersEdgesEdgeversionreportNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesEdgeversionreportRequestTimeout creates a GetTelephonyProvidersEdgesEdgeversionreportRequestTimeout with default headers values
func NewGetTelephonyProvidersEdgesEdgeversionreportRequestTimeout() *GetTelephonyProvidersEdgesEdgeversionreportRequestTimeout {
	return &GetTelephonyProvidersEdgesEdgeversionreportRequestTimeout{}
}

/*
GetTelephonyProvidersEdgesEdgeversionreportRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetTelephonyProvidersEdgesEdgeversionreportRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges edgeversionreport request timeout response has a 2xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges edgeversionreport request timeout response has a 3xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges edgeversionreport request timeout response has a 4xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges edgeversionreport request timeout response has a 5xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges edgeversionreport request timeout response a status code equal to that given
func (o *GetTelephonyProvidersEdgesEdgeversionreportRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgeversionreport][%d] getTelephonyProvidersEdgesEdgeversionreportRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgeversionreport][%d] getTelephonyProvidersEdgesEdgeversionreportRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesEdgeversionreportRequestEntityTooLarge creates a GetTelephonyProvidersEdgesEdgeversionreportRequestEntityTooLarge with default headers values
func NewGetTelephonyProvidersEdgesEdgeversionreportRequestEntityTooLarge() *GetTelephonyProvidersEdgesEdgeversionreportRequestEntityTooLarge {
	return &GetTelephonyProvidersEdgesEdgeversionreportRequestEntityTooLarge{}
}

/*
GetTelephonyProvidersEdgesEdgeversionreportRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetTelephonyProvidersEdgesEdgeversionreportRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges edgeversionreport request entity too large response has a 2xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges edgeversionreport request entity too large response has a 3xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges edgeversionreport request entity too large response has a 4xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges edgeversionreport request entity too large response has a 5xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges edgeversionreport request entity too large response a status code equal to that given
func (o *GetTelephonyProvidersEdgesEdgeversionreportRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgeversionreport][%d] getTelephonyProvidersEdgesEdgeversionreportRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgeversionreport][%d] getTelephonyProvidersEdgesEdgeversionreportRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesEdgeversionreportUnsupportedMediaType creates a GetTelephonyProvidersEdgesEdgeversionreportUnsupportedMediaType with default headers values
func NewGetTelephonyProvidersEdgesEdgeversionreportUnsupportedMediaType() *GetTelephonyProvidersEdgesEdgeversionreportUnsupportedMediaType {
	return &GetTelephonyProvidersEdgesEdgeversionreportUnsupportedMediaType{}
}

/*
GetTelephonyProvidersEdgesEdgeversionreportUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyProvidersEdgesEdgeversionreportUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges edgeversionreport unsupported media type response has a 2xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges edgeversionreport unsupported media type response has a 3xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges edgeversionreport unsupported media type response has a 4xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges edgeversionreport unsupported media type response has a 5xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges edgeversionreport unsupported media type response a status code equal to that given
func (o *GetTelephonyProvidersEdgesEdgeversionreportUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgeversionreport][%d] getTelephonyProvidersEdgesEdgeversionreportUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgeversionreport][%d] getTelephonyProvidersEdgesEdgeversionreportUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesEdgeversionreportTooManyRequests creates a GetTelephonyProvidersEdgesEdgeversionreportTooManyRequests with default headers values
func NewGetTelephonyProvidersEdgesEdgeversionreportTooManyRequests() *GetTelephonyProvidersEdgesEdgeversionreportTooManyRequests {
	return &GetTelephonyProvidersEdgesEdgeversionreportTooManyRequests{}
}

/*
GetTelephonyProvidersEdgesEdgeversionreportTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetTelephonyProvidersEdgesEdgeversionreportTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges edgeversionreport too many requests response has a 2xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges edgeversionreport too many requests response has a 3xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges edgeversionreport too many requests response has a 4xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges edgeversionreport too many requests response has a 5xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges edgeversionreport too many requests response a status code equal to that given
func (o *GetTelephonyProvidersEdgesEdgeversionreportTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgeversionreport][%d] getTelephonyProvidersEdgesEdgeversionreportTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgeversionreport][%d] getTelephonyProvidersEdgesEdgeversionreportTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesEdgeversionreportInternalServerError creates a GetTelephonyProvidersEdgesEdgeversionreportInternalServerError with default headers values
func NewGetTelephonyProvidersEdgesEdgeversionreportInternalServerError() *GetTelephonyProvidersEdgesEdgeversionreportInternalServerError {
	return &GetTelephonyProvidersEdgesEdgeversionreportInternalServerError{}
}

/*
GetTelephonyProvidersEdgesEdgeversionreportInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyProvidersEdgesEdgeversionreportInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges edgeversionreport internal server error response has a 2xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges edgeversionreport internal server error response has a 3xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges edgeversionreport internal server error response has a 4xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edges edgeversionreport internal server error response has a 5xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get telephony providers edges edgeversionreport internal server error response a status code equal to that given
func (o *GetTelephonyProvidersEdgesEdgeversionreportInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgeversionreport][%d] getTelephonyProvidersEdgesEdgeversionreportInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgeversionreport][%d] getTelephonyProvidersEdgesEdgeversionreportInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesEdgeversionreportServiceUnavailable creates a GetTelephonyProvidersEdgesEdgeversionreportServiceUnavailable with default headers values
func NewGetTelephonyProvidersEdgesEdgeversionreportServiceUnavailable() *GetTelephonyProvidersEdgesEdgeversionreportServiceUnavailable {
	return &GetTelephonyProvidersEdgesEdgeversionreportServiceUnavailable{}
}

/*
GetTelephonyProvidersEdgesEdgeversionreportServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyProvidersEdgesEdgeversionreportServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges edgeversionreport service unavailable response has a 2xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges edgeversionreport service unavailable response has a 3xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges edgeversionreport service unavailable response has a 4xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edges edgeversionreport service unavailable response has a 5xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get telephony providers edges edgeversionreport service unavailable response a status code equal to that given
func (o *GetTelephonyProvidersEdgesEdgeversionreportServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgeversionreport][%d] getTelephonyProvidersEdgesEdgeversionreportServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgeversionreport][%d] getTelephonyProvidersEdgesEdgeversionreportServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesEdgeversionreportGatewayTimeout creates a GetTelephonyProvidersEdgesEdgeversionreportGatewayTimeout with default headers values
func NewGetTelephonyProvidersEdgesEdgeversionreportGatewayTimeout() *GetTelephonyProvidersEdgesEdgeversionreportGatewayTimeout {
	return &GetTelephonyProvidersEdgesEdgeversionreportGatewayTimeout{}
}

/*
GetTelephonyProvidersEdgesEdgeversionreportGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetTelephonyProvidersEdgesEdgeversionreportGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges edgeversionreport gateway timeout response has a 2xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges edgeversionreport gateway timeout response has a 3xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges edgeversionreport gateway timeout response has a 4xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edges edgeversionreport gateway timeout response has a 5xx status code
func (o *GetTelephonyProvidersEdgesEdgeversionreportGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get telephony providers edges edgeversionreport gateway timeout response a status code equal to that given
func (o *GetTelephonyProvidersEdgesEdgeversionreportGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgeversionreport][%d] getTelephonyProvidersEdgesEdgeversionreportGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgeversionreport][%d] getTelephonyProvidersEdgesEdgeversionreportGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// GetTelephonyProvidersEdgesDidpoolsDidsReader is a Reader for the GetTelephonyProvidersEdgesDidpoolsDids structure.
type GetTelephonyProvidersEdgesDidpoolsDidsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonyProvidersEdgesDidpoolsDidsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonyProvidersEdgesDidpoolsDidsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonyProvidersEdgesDidpoolsDidsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonyProvidersEdgesDidpoolsDidsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonyProvidersEdgesDidpoolsDidsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonyProvidersEdgesDidpoolsDidsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetTelephonyProvidersEdgesDidpoolsDidsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonyProvidersEdgesDidpoolsDidsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonyProvidersEdgesDidpoolsDidsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonyProvidersEdgesDidpoolsDidsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonyProvidersEdgesDidpoolsDidsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonyProvidersEdgesDidpoolsDidsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonyProvidersEdgesDidpoolsDidsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonyProvidersEdgesDidpoolsDidsOK creates a GetTelephonyProvidersEdgesDidpoolsDidsOK with default headers values
func NewGetTelephonyProvidersEdgesDidpoolsDidsOK() *GetTelephonyProvidersEdgesDidpoolsDidsOK {
	return &GetTelephonyProvidersEdgesDidpoolsDidsOK{}
}

/*
GetTelephonyProvidersEdgesDidpoolsDidsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetTelephonyProvidersEdgesDidpoolsDidsOK struct {
	Payload *models.DIDNumberEntityListing
}

// IsSuccess returns true when this get telephony providers edges didpools dids o k response has a 2xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get telephony providers edges didpools dids o k response has a 3xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges didpools dids o k response has a 4xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edges didpools dids o k response has a 5xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges didpools dids o k response a status code equal to that given
func (o *GetTelephonyProvidersEdgesDidpoolsDidsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools/dids][%d] getTelephonyProvidersEdgesDidpoolsDidsOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools/dids][%d] getTelephonyProvidersEdgesDidpoolsDidsOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsOK) GetPayload() *models.DIDNumberEntityListing {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DIDNumberEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidpoolsDidsBadRequest creates a GetTelephonyProvidersEdgesDidpoolsDidsBadRequest with default headers values
func NewGetTelephonyProvidersEdgesDidpoolsDidsBadRequest() *GetTelephonyProvidersEdgesDidpoolsDidsBadRequest {
	return &GetTelephonyProvidersEdgesDidpoolsDidsBadRequest{}
}

/*
GetTelephonyProvidersEdgesDidpoolsDidsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyProvidersEdgesDidpoolsDidsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges didpools dids bad request response has a 2xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges didpools dids bad request response has a 3xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges didpools dids bad request response has a 4xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges didpools dids bad request response has a 5xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges didpools dids bad request response a status code equal to that given
func (o *GetTelephonyProvidersEdgesDidpoolsDidsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools/dids][%d] getTelephonyProvidersEdgesDidpoolsDidsBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools/dids][%d] getTelephonyProvidersEdgesDidpoolsDidsBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidpoolsDidsUnauthorized creates a GetTelephonyProvidersEdgesDidpoolsDidsUnauthorized with default headers values
func NewGetTelephonyProvidersEdgesDidpoolsDidsUnauthorized() *GetTelephonyProvidersEdgesDidpoolsDidsUnauthorized {
	return &GetTelephonyProvidersEdgesDidpoolsDidsUnauthorized{}
}

/*
GetTelephonyProvidersEdgesDidpoolsDidsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyProvidersEdgesDidpoolsDidsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges didpools dids unauthorized response has a 2xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges didpools dids unauthorized response has a 3xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges didpools dids unauthorized response has a 4xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges didpools dids unauthorized response has a 5xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges didpools dids unauthorized response a status code equal to that given
func (o *GetTelephonyProvidersEdgesDidpoolsDidsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools/dids][%d] getTelephonyProvidersEdgesDidpoolsDidsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools/dids][%d] getTelephonyProvidersEdgesDidpoolsDidsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidpoolsDidsForbidden creates a GetTelephonyProvidersEdgesDidpoolsDidsForbidden with default headers values
func NewGetTelephonyProvidersEdgesDidpoolsDidsForbidden() *GetTelephonyProvidersEdgesDidpoolsDidsForbidden {
	return &GetTelephonyProvidersEdgesDidpoolsDidsForbidden{}
}

/*
GetTelephonyProvidersEdgesDidpoolsDidsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyProvidersEdgesDidpoolsDidsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges didpools dids forbidden response has a 2xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges didpools dids forbidden response has a 3xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges didpools dids forbidden response has a 4xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges didpools dids forbidden response has a 5xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges didpools dids forbidden response a status code equal to that given
func (o *GetTelephonyProvidersEdgesDidpoolsDidsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools/dids][%d] getTelephonyProvidersEdgesDidpoolsDidsForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools/dids][%d] getTelephonyProvidersEdgesDidpoolsDidsForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidpoolsDidsNotFound creates a GetTelephonyProvidersEdgesDidpoolsDidsNotFound with default headers values
func NewGetTelephonyProvidersEdgesDidpoolsDidsNotFound() *GetTelephonyProvidersEdgesDidpoolsDidsNotFound {
	return &GetTelephonyProvidersEdgesDidpoolsDidsNotFound{}
}

/*
GetTelephonyProvidersEdgesDidpoolsDidsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetTelephonyProvidersEdgesDidpoolsDidsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges didpools dids not found response has a 2xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges didpools dids not found response has a 3xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges didpools dids not found response has a 4xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges didpools dids not found response has a 5xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges didpools dids not found response a status code equal to that given
func (o *GetTelephonyProvidersEdgesDidpoolsDidsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools/dids][%d] getTelephonyProvidersEdgesDidpoolsDidsNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools/dids][%d] getTelephonyProvidersEdgesDidpoolsDidsNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidpoolsDidsRequestTimeout creates a GetTelephonyProvidersEdgesDidpoolsDidsRequestTimeout with default headers values
func NewGetTelephonyProvidersEdgesDidpoolsDidsRequestTimeout() *GetTelephonyProvidersEdgesDidpoolsDidsRequestTimeout {
	return &GetTelephonyProvidersEdgesDidpoolsDidsRequestTimeout{}
}

/*
GetTelephonyProvidersEdgesDidpoolsDidsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetTelephonyProvidersEdgesDidpoolsDidsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges didpools dids request timeout response has a 2xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges didpools dids request timeout response has a 3xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges didpools dids request timeout response has a 4xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges didpools dids request timeout response has a 5xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges didpools dids request timeout response a status code equal to that given
func (o *GetTelephonyProvidersEdgesDidpoolsDidsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools/dids][%d] getTelephonyProvidersEdgesDidpoolsDidsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools/dids][%d] getTelephonyProvidersEdgesDidpoolsDidsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidpoolsDidsRequestEntityTooLarge creates a GetTelephonyProvidersEdgesDidpoolsDidsRequestEntityTooLarge with default headers values
func NewGetTelephonyProvidersEdgesDidpoolsDidsRequestEntityTooLarge() *GetTelephonyProvidersEdgesDidpoolsDidsRequestEntityTooLarge {
	return &GetTelephonyProvidersEdgesDidpoolsDidsRequestEntityTooLarge{}
}

/*
GetTelephonyProvidersEdgesDidpoolsDidsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetTelephonyProvidersEdgesDidpoolsDidsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges didpools dids request entity too large response has a 2xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges didpools dids request entity too large response has a 3xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges didpools dids request entity too large response has a 4xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges didpools dids request entity too large response has a 5xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges didpools dids request entity too large response a status code equal to that given
func (o *GetTelephonyProvidersEdgesDidpoolsDidsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools/dids][%d] getTelephonyProvidersEdgesDidpoolsDidsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools/dids][%d] getTelephonyProvidersEdgesDidpoolsDidsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidpoolsDidsUnsupportedMediaType creates a GetTelephonyProvidersEdgesDidpoolsDidsUnsupportedMediaType with default headers values
func NewGetTelephonyProvidersEdgesDidpoolsDidsUnsupportedMediaType() *GetTelephonyProvidersEdgesDidpoolsDidsUnsupportedMediaType {
	return &GetTelephonyProvidersEdgesDidpoolsDidsUnsupportedMediaType{}
}

/*
GetTelephonyProvidersEdgesDidpoolsDidsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyProvidersEdgesDidpoolsDidsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges didpools dids unsupported media type response has a 2xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges didpools dids unsupported media type response has a 3xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges didpools dids unsupported media type response has a 4xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges didpools dids unsupported media type response has a 5xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges didpools dids unsupported media type response a status code equal to that given
func (o *GetTelephonyProvidersEdgesDidpoolsDidsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools/dids][%d] getTelephonyProvidersEdgesDidpoolsDidsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools/dids][%d] getTelephonyProvidersEdgesDidpoolsDidsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidpoolsDidsTooManyRequests creates a GetTelephonyProvidersEdgesDidpoolsDidsTooManyRequests with default headers values
func NewGetTelephonyProvidersEdgesDidpoolsDidsTooManyRequests() *GetTelephonyProvidersEdgesDidpoolsDidsTooManyRequests {
	return &GetTelephonyProvidersEdgesDidpoolsDidsTooManyRequests{}
}

/*
GetTelephonyProvidersEdgesDidpoolsDidsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetTelephonyProvidersEdgesDidpoolsDidsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges didpools dids too many requests response has a 2xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges didpools dids too many requests response has a 3xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges didpools dids too many requests response has a 4xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges didpools dids too many requests response has a 5xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges didpools dids too many requests response a status code equal to that given
func (o *GetTelephonyProvidersEdgesDidpoolsDidsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools/dids][%d] getTelephonyProvidersEdgesDidpoolsDidsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools/dids][%d] getTelephonyProvidersEdgesDidpoolsDidsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidpoolsDidsInternalServerError creates a GetTelephonyProvidersEdgesDidpoolsDidsInternalServerError with default headers values
func NewGetTelephonyProvidersEdgesDidpoolsDidsInternalServerError() *GetTelephonyProvidersEdgesDidpoolsDidsInternalServerError {
	return &GetTelephonyProvidersEdgesDidpoolsDidsInternalServerError{}
}

/*
GetTelephonyProvidersEdgesDidpoolsDidsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyProvidersEdgesDidpoolsDidsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges didpools dids internal server error response has a 2xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges didpools dids internal server error response has a 3xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges didpools dids internal server error response has a 4xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edges didpools dids internal server error response has a 5xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get telephony providers edges didpools dids internal server error response a status code equal to that given
func (o *GetTelephonyProvidersEdgesDidpoolsDidsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools/dids][%d] getTelephonyProvidersEdgesDidpoolsDidsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools/dids][%d] getTelephonyProvidersEdgesDidpoolsDidsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidpoolsDidsServiceUnavailable creates a GetTelephonyProvidersEdgesDidpoolsDidsServiceUnavailable with default headers values
func NewGetTelephonyProvidersEdgesDidpoolsDidsServiceUnavailable() *GetTelephonyProvidersEdgesDidpoolsDidsServiceUnavailable {
	return &GetTelephonyProvidersEdgesDidpoolsDidsServiceUnavailable{}
}

/*
GetTelephonyProvidersEdgesDidpoolsDidsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyProvidersEdgesDidpoolsDidsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges didpools dids service unavailable response has a 2xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges didpools dids service unavailable response has a 3xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges didpools dids service unavailable response has a 4xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edges didpools dids service unavailable response has a 5xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get telephony providers edges didpools dids service unavailable response a status code equal to that given
func (o *GetTelephonyProvidersEdgesDidpoolsDidsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools/dids][%d] getTelephonyProvidersEdgesDidpoolsDidsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools/dids][%d] getTelephonyProvidersEdgesDidpoolsDidsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidpoolsDidsGatewayTimeout creates a GetTelephonyProvidersEdgesDidpoolsDidsGatewayTimeout with default headers values
func NewGetTelephonyProvidersEdgesDidpoolsDidsGatewayTimeout() *GetTelephonyProvidersEdgesDidpoolsDidsGatewayTimeout {
	return &GetTelephonyProvidersEdgesDidpoolsDidsGatewayTimeout{}
}

/*
GetTelephonyProvidersEdgesDidpoolsDidsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetTelephonyProvidersEdgesDidpoolsDidsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges didpools dids gateway timeout response has a 2xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges didpools dids gateway timeout response has a 3xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges didpools dids gateway timeout response has a 4xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edges didpools dids gateway timeout response has a 5xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsDidsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get telephony providers edges didpools dids gateway timeout response a status code equal to that given
func (o *GetTelephonyProvidersEdgesDidpoolsDidsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools/dids][%d] getTelephonyProvidersEdgesDidpoolsDidsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools/dids][%d] getTelephonyProvidersEdgesDidpoolsDidsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidpoolsDidsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

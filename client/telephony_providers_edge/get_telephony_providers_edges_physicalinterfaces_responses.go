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

// GetTelephonyProvidersEdgesPhysicalinterfacesReader is a Reader for the GetTelephonyProvidersEdgesPhysicalinterfaces structure.
type GetTelephonyProvidersEdgesPhysicalinterfacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonyProvidersEdgesPhysicalinterfacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonyProvidersEdgesPhysicalinterfacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonyProvidersEdgesPhysicalinterfacesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonyProvidersEdgesPhysicalinterfacesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonyProvidersEdgesPhysicalinterfacesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetTelephonyProvidersEdgesPhysicalinterfacesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonyProvidersEdgesPhysicalinterfacesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonyProvidersEdgesPhysicalinterfacesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonyProvidersEdgesPhysicalinterfacesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonyProvidersEdgesPhysicalinterfacesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonyProvidersEdgesPhysicalinterfacesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonyProvidersEdgesPhysicalinterfacesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonyProvidersEdgesPhysicalinterfacesOK creates a GetTelephonyProvidersEdgesPhysicalinterfacesOK with default headers values
func NewGetTelephonyProvidersEdgesPhysicalinterfacesOK() *GetTelephonyProvidersEdgesPhysicalinterfacesOK {
	return &GetTelephonyProvidersEdgesPhysicalinterfacesOK{}
}

/*
GetTelephonyProvidersEdgesPhysicalinterfacesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetTelephonyProvidersEdgesPhysicalinterfacesOK struct {
	Payload *models.PhysicalInterfaceEntityListing
}

// IsSuccess returns true when this get telephony providers edges physicalinterfaces o k response has a 2xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get telephony providers edges physicalinterfaces o k response has a 3xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges physicalinterfaces o k response has a 4xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edges physicalinterfaces o k response has a 5xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges physicalinterfaces o k response a status code equal to that given
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/physicalinterfaces][%d] getTelephonyProvidersEdgesPhysicalinterfacesOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/physicalinterfaces][%d] getTelephonyProvidersEdgesPhysicalinterfacesOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesOK) GetPayload() *models.PhysicalInterfaceEntityListing {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PhysicalInterfaceEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhysicalinterfacesBadRequest creates a GetTelephonyProvidersEdgesPhysicalinterfacesBadRequest with default headers values
func NewGetTelephonyProvidersEdgesPhysicalinterfacesBadRequest() *GetTelephonyProvidersEdgesPhysicalinterfacesBadRequest {
	return &GetTelephonyProvidersEdgesPhysicalinterfacesBadRequest{}
}

/*
GetTelephonyProvidersEdgesPhysicalinterfacesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyProvidersEdgesPhysicalinterfacesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges physicalinterfaces bad request response has a 2xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges physicalinterfaces bad request response has a 3xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges physicalinterfaces bad request response has a 4xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges physicalinterfaces bad request response has a 5xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges physicalinterfaces bad request response a status code equal to that given
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/physicalinterfaces][%d] getTelephonyProvidersEdgesPhysicalinterfacesBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/physicalinterfaces][%d] getTelephonyProvidersEdgesPhysicalinterfacesBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhysicalinterfacesUnauthorized creates a GetTelephonyProvidersEdgesPhysicalinterfacesUnauthorized with default headers values
func NewGetTelephonyProvidersEdgesPhysicalinterfacesUnauthorized() *GetTelephonyProvidersEdgesPhysicalinterfacesUnauthorized {
	return &GetTelephonyProvidersEdgesPhysicalinterfacesUnauthorized{}
}

/*
GetTelephonyProvidersEdgesPhysicalinterfacesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyProvidersEdgesPhysicalinterfacesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges physicalinterfaces unauthorized response has a 2xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges physicalinterfaces unauthorized response has a 3xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges physicalinterfaces unauthorized response has a 4xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges physicalinterfaces unauthorized response has a 5xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges physicalinterfaces unauthorized response a status code equal to that given
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/physicalinterfaces][%d] getTelephonyProvidersEdgesPhysicalinterfacesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/physicalinterfaces][%d] getTelephonyProvidersEdgesPhysicalinterfacesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhysicalinterfacesForbidden creates a GetTelephonyProvidersEdgesPhysicalinterfacesForbidden with default headers values
func NewGetTelephonyProvidersEdgesPhysicalinterfacesForbidden() *GetTelephonyProvidersEdgesPhysicalinterfacesForbidden {
	return &GetTelephonyProvidersEdgesPhysicalinterfacesForbidden{}
}

/*
GetTelephonyProvidersEdgesPhysicalinterfacesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyProvidersEdgesPhysicalinterfacesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges physicalinterfaces forbidden response has a 2xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges physicalinterfaces forbidden response has a 3xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges physicalinterfaces forbidden response has a 4xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges physicalinterfaces forbidden response has a 5xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges physicalinterfaces forbidden response a status code equal to that given
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/physicalinterfaces][%d] getTelephonyProvidersEdgesPhysicalinterfacesForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/physicalinterfaces][%d] getTelephonyProvidersEdgesPhysicalinterfacesForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhysicalinterfacesNotFound creates a GetTelephonyProvidersEdgesPhysicalinterfacesNotFound with default headers values
func NewGetTelephonyProvidersEdgesPhysicalinterfacesNotFound() *GetTelephonyProvidersEdgesPhysicalinterfacesNotFound {
	return &GetTelephonyProvidersEdgesPhysicalinterfacesNotFound{}
}

/*
GetTelephonyProvidersEdgesPhysicalinterfacesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetTelephonyProvidersEdgesPhysicalinterfacesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges physicalinterfaces not found response has a 2xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges physicalinterfaces not found response has a 3xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges physicalinterfaces not found response has a 4xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges physicalinterfaces not found response has a 5xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges physicalinterfaces not found response a status code equal to that given
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/physicalinterfaces][%d] getTelephonyProvidersEdgesPhysicalinterfacesNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/physicalinterfaces][%d] getTelephonyProvidersEdgesPhysicalinterfacesNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhysicalinterfacesRequestTimeout creates a GetTelephonyProvidersEdgesPhysicalinterfacesRequestTimeout with default headers values
func NewGetTelephonyProvidersEdgesPhysicalinterfacesRequestTimeout() *GetTelephonyProvidersEdgesPhysicalinterfacesRequestTimeout {
	return &GetTelephonyProvidersEdgesPhysicalinterfacesRequestTimeout{}
}

/*
GetTelephonyProvidersEdgesPhysicalinterfacesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetTelephonyProvidersEdgesPhysicalinterfacesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges physicalinterfaces request timeout response has a 2xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges physicalinterfaces request timeout response has a 3xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges physicalinterfaces request timeout response has a 4xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges physicalinterfaces request timeout response has a 5xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges physicalinterfaces request timeout response a status code equal to that given
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/physicalinterfaces][%d] getTelephonyProvidersEdgesPhysicalinterfacesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/physicalinterfaces][%d] getTelephonyProvidersEdgesPhysicalinterfacesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhysicalinterfacesRequestEntityTooLarge creates a GetTelephonyProvidersEdgesPhysicalinterfacesRequestEntityTooLarge with default headers values
func NewGetTelephonyProvidersEdgesPhysicalinterfacesRequestEntityTooLarge() *GetTelephonyProvidersEdgesPhysicalinterfacesRequestEntityTooLarge {
	return &GetTelephonyProvidersEdgesPhysicalinterfacesRequestEntityTooLarge{}
}

/*
GetTelephonyProvidersEdgesPhysicalinterfacesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetTelephonyProvidersEdgesPhysicalinterfacesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges physicalinterfaces request entity too large response has a 2xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges physicalinterfaces request entity too large response has a 3xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges physicalinterfaces request entity too large response has a 4xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges physicalinterfaces request entity too large response has a 5xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges physicalinterfaces request entity too large response a status code equal to that given
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/physicalinterfaces][%d] getTelephonyProvidersEdgesPhysicalinterfacesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/physicalinterfaces][%d] getTelephonyProvidersEdgesPhysicalinterfacesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhysicalinterfacesUnsupportedMediaType creates a GetTelephonyProvidersEdgesPhysicalinterfacesUnsupportedMediaType with default headers values
func NewGetTelephonyProvidersEdgesPhysicalinterfacesUnsupportedMediaType() *GetTelephonyProvidersEdgesPhysicalinterfacesUnsupportedMediaType {
	return &GetTelephonyProvidersEdgesPhysicalinterfacesUnsupportedMediaType{}
}

/*
GetTelephonyProvidersEdgesPhysicalinterfacesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyProvidersEdgesPhysicalinterfacesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges physicalinterfaces unsupported media type response has a 2xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges physicalinterfaces unsupported media type response has a 3xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges physicalinterfaces unsupported media type response has a 4xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges physicalinterfaces unsupported media type response has a 5xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges physicalinterfaces unsupported media type response a status code equal to that given
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/physicalinterfaces][%d] getTelephonyProvidersEdgesPhysicalinterfacesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/physicalinterfaces][%d] getTelephonyProvidersEdgesPhysicalinterfacesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhysicalinterfacesTooManyRequests creates a GetTelephonyProvidersEdgesPhysicalinterfacesTooManyRequests with default headers values
func NewGetTelephonyProvidersEdgesPhysicalinterfacesTooManyRequests() *GetTelephonyProvidersEdgesPhysicalinterfacesTooManyRequests {
	return &GetTelephonyProvidersEdgesPhysicalinterfacesTooManyRequests{}
}

/*
GetTelephonyProvidersEdgesPhysicalinterfacesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetTelephonyProvidersEdgesPhysicalinterfacesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges physicalinterfaces too many requests response has a 2xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges physicalinterfaces too many requests response has a 3xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges physicalinterfaces too many requests response has a 4xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges physicalinterfaces too many requests response has a 5xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges physicalinterfaces too many requests response a status code equal to that given
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/physicalinterfaces][%d] getTelephonyProvidersEdgesPhysicalinterfacesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/physicalinterfaces][%d] getTelephonyProvidersEdgesPhysicalinterfacesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhysicalinterfacesInternalServerError creates a GetTelephonyProvidersEdgesPhysicalinterfacesInternalServerError with default headers values
func NewGetTelephonyProvidersEdgesPhysicalinterfacesInternalServerError() *GetTelephonyProvidersEdgesPhysicalinterfacesInternalServerError {
	return &GetTelephonyProvidersEdgesPhysicalinterfacesInternalServerError{}
}

/*
GetTelephonyProvidersEdgesPhysicalinterfacesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyProvidersEdgesPhysicalinterfacesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges physicalinterfaces internal server error response has a 2xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges physicalinterfaces internal server error response has a 3xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges physicalinterfaces internal server error response has a 4xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edges physicalinterfaces internal server error response has a 5xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get telephony providers edges physicalinterfaces internal server error response a status code equal to that given
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/physicalinterfaces][%d] getTelephonyProvidersEdgesPhysicalinterfacesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/physicalinterfaces][%d] getTelephonyProvidersEdgesPhysicalinterfacesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhysicalinterfacesServiceUnavailable creates a GetTelephonyProvidersEdgesPhysicalinterfacesServiceUnavailable with default headers values
func NewGetTelephonyProvidersEdgesPhysicalinterfacesServiceUnavailable() *GetTelephonyProvidersEdgesPhysicalinterfacesServiceUnavailable {
	return &GetTelephonyProvidersEdgesPhysicalinterfacesServiceUnavailable{}
}

/*
GetTelephonyProvidersEdgesPhysicalinterfacesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyProvidersEdgesPhysicalinterfacesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges physicalinterfaces service unavailable response has a 2xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges physicalinterfaces service unavailable response has a 3xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges physicalinterfaces service unavailable response has a 4xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edges physicalinterfaces service unavailable response has a 5xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get telephony providers edges physicalinterfaces service unavailable response a status code equal to that given
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/physicalinterfaces][%d] getTelephonyProvidersEdgesPhysicalinterfacesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/physicalinterfaces][%d] getTelephonyProvidersEdgesPhysicalinterfacesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhysicalinterfacesGatewayTimeout creates a GetTelephonyProvidersEdgesPhysicalinterfacesGatewayTimeout with default headers values
func NewGetTelephonyProvidersEdgesPhysicalinterfacesGatewayTimeout() *GetTelephonyProvidersEdgesPhysicalinterfacesGatewayTimeout {
	return &GetTelephonyProvidersEdgesPhysicalinterfacesGatewayTimeout{}
}

/*
GetTelephonyProvidersEdgesPhysicalinterfacesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetTelephonyProvidersEdgesPhysicalinterfacesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges physicalinterfaces gateway timeout response has a 2xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges physicalinterfaces gateway timeout response has a 3xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges physicalinterfaces gateway timeout response has a 4xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edges physicalinterfaces gateway timeout response has a 5xx status code
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get telephony providers edges physicalinterfaces gateway timeout response a status code equal to that given
func (o *GetTelephonyProvidersEdgesPhysicalinterfacesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/physicalinterfaces][%d] getTelephonyProvidersEdgesPhysicalinterfacesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/physicalinterfaces][%d] getTelephonyProvidersEdgesPhysicalinterfacesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhysicalinterfacesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

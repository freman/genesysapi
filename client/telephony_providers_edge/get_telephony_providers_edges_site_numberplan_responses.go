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

// GetTelephonyProvidersEdgesSiteNumberplanReader is a Reader for the GetTelephonyProvidersEdgesSiteNumberplan structure.
type GetTelephonyProvidersEdgesSiteNumberplanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonyProvidersEdgesSiteNumberplanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonyProvidersEdgesSiteNumberplanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonyProvidersEdgesSiteNumberplanBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonyProvidersEdgesSiteNumberplanUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonyProvidersEdgesSiteNumberplanForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonyProvidersEdgesSiteNumberplanNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetTelephonyProvidersEdgesSiteNumberplanRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonyProvidersEdgesSiteNumberplanRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonyProvidersEdgesSiteNumberplanUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonyProvidersEdgesSiteNumberplanTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonyProvidersEdgesSiteNumberplanInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonyProvidersEdgesSiteNumberplanServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonyProvidersEdgesSiteNumberplanGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonyProvidersEdgesSiteNumberplanOK creates a GetTelephonyProvidersEdgesSiteNumberplanOK with default headers values
func NewGetTelephonyProvidersEdgesSiteNumberplanOK() *GetTelephonyProvidersEdgesSiteNumberplanOK {
	return &GetTelephonyProvidersEdgesSiteNumberplanOK{}
}

/*GetTelephonyProvidersEdgesSiteNumberplanOK handles this case with default header values.

successful operation
*/
type GetTelephonyProvidersEdgesSiteNumberplanOK struct {
	Payload *models.NumberPlan
}

func (o *GetTelephonyProvidersEdgesSiteNumberplanOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/numberplans/{numberPlanId}][%d] getTelephonyProvidersEdgesSiteNumberplanOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteNumberplanOK) GetPayload() *models.NumberPlan {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteNumberplanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NumberPlan)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteNumberplanBadRequest creates a GetTelephonyProvidersEdgesSiteNumberplanBadRequest with default headers values
func NewGetTelephonyProvidersEdgesSiteNumberplanBadRequest() *GetTelephonyProvidersEdgesSiteNumberplanBadRequest {
	return &GetTelephonyProvidersEdgesSiteNumberplanBadRequest{}
}

/*GetTelephonyProvidersEdgesSiteNumberplanBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyProvidersEdgesSiteNumberplanBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteNumberplanBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/numberplans/{numberPlanId}][%d] getTelephonyProvidersEdgesSiteNumberplanBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteNumberplanBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteNumberplanBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteNumberplanUnauthorized creates a GetTelephonyProvidersEdgesSiteNumberplanUnauthorized with default headers values
func NewGetTelephonyProvidersEdgesSiteNumberplanUnauthorized() *GetTelephonyProvidersEdgesSiteNumberplanUnauthorized {
	return &GetTelephonyProvidersEdgesSiteNumberplanUnauthorized{}
}

/*GetTelephonyProvidersEdgesSiteNumberplanUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyProvidersEdgesSiteNumberplanUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteNumberplanUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/numberplans/{numberPlanId}][%d] getTelephonyProvidersEdgesSiteNumberplanUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteNumberplanUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteNumberplanUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteNumberplanForbidden creates a GetTelephonyProvidersEdgesSiteNumberplanForbidden with default headers values
func NewGetTelephonyProvidersEdgesSiteNumberplanForbidden() *GetTelephonyProvidersEdgesSiteNumberplanForbidden {
	return &GetTelephonyProvidersEdgesSiteNumberplanForbidden{}
}

/*GetTelephonyProvidersEdgesSiteNumberplanForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyProvidersEdgesSiteNumberplanForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteNumberplanForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/numberplans/{numberPlanId}][%d] getTelephonyProvidersEdgesSiteNumberplanForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteNumberplanForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteNumberplanForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteNumberplanNotFound creates a GetTelephonyProvidersEdgesSiteNumberplanNotFound with default headers values
func NewGetTelephonyProvidersEdgesSiteNumberplanNotFound() *GetTelephonyProvidersEdgesSiteNumberplanNotFound {
	return &GetTelephonyProvidersEdgesSiteNumberplanNotFound{}
}

/*GetTelephonyProvidersEdgesSiteNumberplanNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetTelephonyProvidersEdgesSiteNumberplanNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteNumberplanNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/numberplans/{numberPlanId}][%d] getTelephonyProvidersEdgesSiteNumberplanNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteNumberplanNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteNumberplanNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteNumberplanRequestTimeout creates a GetTelephonyProvidersEdgesSiteNumberplanRequestTimeout with default headers values
func NewGetTelephonyProvidersEdgesSiteNumberplanRequestTimeout() *GetTelephonyProvidersEdgesSiteNumberplanRequestTimeout {
	return &GetTelephonyProvidersEdgesSiteNumberplanRequestTimeout{}
}

/*GetTelephonyProvidersEdgesSiteNumberplanRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetTelephonyProvidersEdgesSiteNumberplanRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteNumberplanRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/numberplans/{numberPlanId}][%d] getTelephonyProvidersEdgesSiteNumberplanRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteNumberplanRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteNumberplanRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteNumberplanRequestEntityTooLarge creates a GetTelephonyProvidersEdgesSiteNumberplanRequestEntityTooLarge with default headers values
func NewGetTelephonyProvidersEdgesSiteNumberplanRequestEntityTooLarge() *GetTelephonyProvidersEdgesSiteNumberplanRequestEntityTooLarge {
	return &GetTelephonyProvidersEdgesSiteNumberplanRequestEntityTooLarge{}
}

/*GetTelephonyProvidersEdgesSiteNumberplanRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetTelephonyProvidersEdgesSiteNumberplanRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteNumberplanRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/numberplans/{numberPlanId}][%d] getTelephonyProvidersEdgesSiteNumberplanRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteNumberplanRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteNumberplanRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteNumberplanUnsupportedMediaType creates a GetTelephonyProvidersEdgesSiteNumberplanUnsupportedMediaType with default headers values
func NewGetTelephonyProvidersEdgesSiteNumberplanUnsupportedMediaType() *GetTelephonyProvidersEdgesSiteNumberplanUnsupportedMediaType {
	return &GetTelephonyProvidersEdgesSiteNumberplanUnsupportedMediaType{}
}

/*GetTelephonyProvidersEdgesSiteNumberplanUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyProvidersEdgesSiteNumberplanUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteNumberplanUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/numberplans/{numberPlanId}][%d] getTelephonyProvidersEdgesSiteNumberplanUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteNumberplanUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteNumberplanUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteNumberplanTooManyRequests creates a GetTelephonyProvidersEdgesSiteNumberplanTooManyRequests with default headers values
func NewGetTelephonyProvidersEdgesSiteNumberplanTooManyRequests() *GetTelephonyProvidersEdgesSiteNumberplanTooManyRequests {
	return &GetTelephonyProvidersEdgesSiteNumberplanTooManyRequests{}
}

/*GetTelephonyProvidersEdgesSiteNumberplanTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetTelephonyProvidersEdgesSiteNumberplanTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteNumberplanTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/numberplans/{numberPlanId}][%d] getTelephonyProvidersEdgesSiteNumberplanTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteNumberplanTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteNumberplanTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteNumberplanInternalServerError creates a GetTelephonyProvidersEdgesSiteNumberplanInternalServerError with default headers values
func NewGetTelephonyProvidersEdgesSiteNumberplanInternalServerError() *GetTelephonyProvidersEdgesSiteNumberplanInternalServerError {
	return &GetTelephonyProvidersEdgesSiteNumberplanInternalServerError{}
}

/*GetTelephonyProvidersEdgesSiteNumberplanInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyProvidersEdgesSiteNumberplanInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteNumberplanInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/numberplans/{numberPlanId}][%d] getTelephonyProvidersEdgesSiteNumberplanInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteNumberplanInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteNumberplanInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteNumberplanServiceUnavailable creates a GetTelephonyProvidersEdgesSiteNumberplanServiceUnavailable with default headers values
func NewGetTelephonyProvidersEdgesSiteNumberplanServiceUnavailable() *GetTelephonyProvidersEdgesSiteNumberplanServiceUnavailable {
	return &GetTelephonyProvidersEdgesSiteNumberplanServiceUnavailable{}
}

/*GetTelephonyProvidersEdgesSiteNumberplanServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyProvidersEdgesSiteNumberplanServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteNumberplanServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/numberplans/{numberPlanId}][%d] getTelephonyProvidersEdgesSiteNumberplanServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteNumberplanServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteNumberplanServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteNumberplanGatewayTimeout creates a GetTelephonyProvidersEdgesSiteNumberplanGatewayTimeout with default headers values
func NewGetTelephonyProvidersEdgesSiteNumberplanGatewayTimeout() *GetTelephonyProvidersEdgesSiteNumberplanGatewayTimeout {
	return &GetTelephonyProvidersEdgesSiteNumberplanGatewayTimeout{}
}

/*GetTelephonyProvidersEdgesSiteNumberplanGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetTelephonyProvidersEdgesSiteNumberplanGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteNumberplanGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/numberplans/{numberPlanId}][%d] getTelephonyProvidersEdgesSiteNumberplanGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteNumberplanGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteNumberplanGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

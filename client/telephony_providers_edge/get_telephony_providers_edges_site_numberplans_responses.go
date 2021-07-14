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

// GetTelephonyProvidersEdgesSiteNumberplansReader is a Reader for the GetTelephonyProvidersEdgesSiteNumberplans structure.
type GetTelephonyProvidersEdgesSiteNumberplansReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonyProvidersEdgesSiteNumberplansReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonyProvidersEdgesSiteNumberplansOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonyProvidersEdgesSiteNumberplansBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonyProvidersEdgesSiteNumberplansUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonyProvidersEdgesSiteNumberplansForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonyProvidersEdgesSiteNumberplansNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetTelephonyProvidersEdgesSiteNumberplansRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonyProvidersEdgesSiteNumberplansRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonyProvidersEdgesSiteNumberplansUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonyProvidersEdgesSiteNumberplansTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonyProvidersEdgesSiteNumberplansInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonyProvidersEdgesSiteNumberplansServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonyProvidersEdgesSiteNumberplansGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonyProvidersEdgesSiteNumberplansOK creates a GetTelephonyProvidersEdgesSiteNumberplansOK with default headers values
func NewGetTelephonyProvidersEdgesSiteNumberplansOK() *GetTelephonyProvidersEdgesSiteNumberplansOK {
	return &GetTelephonyProvidersEdgesSiteNumberplansOK{}
}

/*GetTelephonyProvidersEdgesSiteNumberplansOK handles this case with default header values.

successful operation
*/
type GetTelephonyProvidersEdgesSiteNumberplansOK struct {
	Payload []*models.NumberPlan
}

func (o *GetTelephonyProvidersEdgesSiteNumberplansOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/numberplans][%d] getTelephonyProvidersEdgesSiteNumberplansOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteNumberplansOK) GetPayload() []*models.NumberPlan {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteNumberplansOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteNumberplansBadRequest creates a GetTelephonyProvidersEdgesSiteNumberplansBadRequest with default headers values
func NewGetTelephonyProvidersEdgesSiteNumberplansBadRequest() *GetTelephonyProvidersEdgesSiteNumberplansBadRequest {
	return &GetTelephonyProvidersEdgesSiteNumberplansBadRequest{}
}

/*GetTelephonyProvidersEdgesSiteNumberplansBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyProvidersEdgesSiteNumberplansBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteNumberplansBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/numberplans][%d] getTelephonyProvidersEdgesSiteNumberplansBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteNumberplansBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteNumberplansBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteNumberplansUnauthorized creates a GetTelephonyProvidersEdgesSiteNumberplansUnauthorized with default headers values
func NewGetTelephonyProvidersEdgesSiteNumberplansUnauthorized() *GetTelephonyProvidersEdgesSiteNumberplansUnauthorized {
	return &GetTelephonyProvidersEdgesSiteNumberplansUnauthorized{}
}

/*GetTelephonyProvidersEdgesSiteNumberplansUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyProvidersEdgesSiteNumberplansUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteNumberplansUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/numberplans][%d] getTelephonyProvidersEdgesSiteNumberplansUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteNumberplansUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteNumberplansUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteNumberplansForbidden creates a GetTelephonyProvidersEdgesSiteNumberplansForbidden with default headers values
func NewGetTelephonyProvidersEdgesSiteNumberplansForbidden() *GetTelephonyProvidersEdgesSiteNumberplansForbidden {
	return &GetTelephonyProvidersEdgesSiteNumberplansForbidden{}
}

/*GetTelephonyProvidersEdgesSiteNumberplansForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyProvidersEdgesSiteNumberplansForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteNumberplansForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/numberplans][%d] getTelephonyProvidersEdgesSiteNumberplansForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteNumberplansForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteNumberplansForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteNumberplansNotFound creates a GetTelephonyProvidersEdgesSiteNumberplansNotFound with default headers values
func NewGetTelephonyProvidersEdgesSiteNumberplansNotFound() *GetTelephonyProvidersEdgesSiteNumberplansNotFound {
	return &GetTelephonyProvidersEdgesSiteNumberplansNotFound{}
}

/*GetTelephonyProvidersEdgesSiteNumberplansNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetTelephonyProvidersEdgesSiteNumberplansNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteNumberplansNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/numberplans][%d] getTelephonyProvidersEdgesSiteNumberplansNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteNumberplansNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteNumberplansNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteNumberplansRequestTimeout creates a GetTelephonyProvidersEdgesSiteNumberplansRequestTimeout with default headers values
func NewGetTelephonyProvidersEdgesSiteNumberplansRequestTimeout() *GetTelephonyProvidersEdgesSiteNumberplansRequestTimeout {
	return &GetTelephonyProvidersEdgesSiteNumberplansRequestTimeout{}
}

/*GetTelephonyProvidersEdgesSiteNumberplansRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetTelephonyProvidersEdgesSiteNumberplansRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteNumberplansRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/numberplans][%d] getTelephonyProvidersEdgesSiteNumberplansRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteNumberplansRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteNumberplansRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteNumberplansRequestEntityTooLarge creates a GetTelephonyProvidersEdgesSiteNumberplansRequestEntityTooLarge with default headers values
func NewGetTelephonyProvidersEdgesSiteNumberplansRequestEntityTooLarge() *GetTelephonyProvidersEdgesSiteNumberplansRequestEntityTooLarge {
	return &GetTelephonyProvidersEdgesSiteNumberplansRequestEntityTooLarge{}
}

/*GetTelephonyProvidersEdgesSiteNumberplansRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetTelephonyProvidersEdgesSiteNumberplansRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteNumberplansRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/numberplans][%d] getTelephonyProvidersEdgesSiteNumberplansRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteNumberplansRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteNumberplansRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteNumberplansUnsupportedMediaType creates a GetTelephonyProvidersEdgesSiteNumberplansUnsupportedMediaType with default headers values
func NewGetTelephonyProvidersEdgesSiteNumberplansUnsupportedMediaType() *GetTelephonyProvidersEdgesSiteNumberplansUnsupportedMediaType {
	return &GetTelephonyProvidersEdgesSiteNumberplansUnsupportedMediaType{}
}

/*GetTelephonyProvidersEdgesSiteNumberplansUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyProvidersEdgesSiteNumberplansUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteNumberplansUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/numberplans][%d] getTelephonyProvidersEdgesSiteNumberplansUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteNumberplansUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteNumberplansUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteNumberplansTooManyRequests creates a GetTelephonyProvidersEdgesSiteNumberplansTooManyRequests with default headers values
func NewGetTelephonyProvidersEdgesSiteNumberplansTooManyRequests() *GetTelephonyProvidersEdgesSiteNumberplansTooManyRequests {
	return &GetTelephonyProvidersEdgesSiteNumberplansTooManyRequests{}
}

/*GetTelephonyProvidersEdgesSiteNumberplansTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetTelephonyProvidersEdgesSiteNumberplansTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteNumberplansTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/numberplans][%d] getTelephonyProvidersEdgesSiteNumberplansTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteNumberplansTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteNumberplansTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteNumberplansInternalServerError creates a GetTelephonyProvidersEdgesSiteNumberplansInternalServerError with default headers values
func NewGetTelephonyProvidersEdgesSiteNumberplansInternalServerError() *GetTelephonyProvidersEdgesSiteNumberplansInternalServerError {
	return &GetTelephonyProvidersEdgesSiteNumberplansInternalServerError{}
}

/*GetTelephonyProvidersEdgesSiteNumberplansInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyProvidersEdgesSiteNumberplansInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteNumberplansInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/numberplans][%d] getTelephonyProvidersEdgesSiteNumberplansInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteNumberplansInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteNumberplansInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteNumberplansServiceUnavailable creates a GetTelephonyProvidersEdgesSiteNumberplansServiceUnavailable with default headers values
func NewGetTelephonyProvidersEdgesSiteNumberplansServiceUnavailable() *GetTelephonyProvidersEdgesSiteNumberplansServiceUnavailable {
	return &GetTelephonyProvidersEdgesSiteNumberplansServiceUnavailable{}
}

/*GetTelephonyProvidersEdgesSiteNumberplansServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyProvidersEdgesSiteNumberplansServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteNumberplansServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/numberplans][%d] getTelephonyProvidersEdgesSiteNumberplansServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteNumberplansServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteNumberplansServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteNumberplansGatewayTimeout creates a GetTelephonyProvidersEdgesSiteNumberplansGatewayTimeout with default headers values
func NewGetTelephonyProvidersEdgesSiteNumberplansGatewayTimeout() *GetTelephonyProvidersEdgesSiteNumberplansGatewayTimeout {
	return &GetTelephonyProvidersEdgesSiteNumberplansGatewayTimeout{}
}

/*GetTelephonyProvidersEdgesSiteNumberplansGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetTelephonyProvidersEdgesSiteNumberplansGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteNumberplansGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/numberplans][%d] getTelephonyProvidersEdgesSiteNumberplansGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteNumberplansGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteNumberplansGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

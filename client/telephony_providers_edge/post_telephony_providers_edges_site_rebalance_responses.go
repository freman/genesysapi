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

// PostTelephonyProvidersEdgesSiteRebalanceReader is a Reader for the PostTelephonyProvidersEdgesSiteRebalance structure.
type PostTelephonyProvidersEdgesSiteRebalanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTelephonyProvidersEdgesSiteRebalanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPostTelephonyProvidersEdgesSiteRebalanceAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostTelephonyProvidersEdgesSiteRebalanceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostTelephonyProvidersEdgesSiteRebalanceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostTelephonyProvidersEdgesSiteRebalanceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostTelephonyProvidersEdgesSiteRebalanceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostTelephonyProvidersEdgesSiteRebalanceRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostTelephonyProvidersEdgesSiteRebalanceUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostTelephonyProvidersEdgesSiteRebalanceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostTelephonyProvidersEdgesSiteRebalanceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostTelephonyProvidersEdgesSiteRebalanceServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostTelephonyProvidersEdgesSiteRebalanceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostTelephonyProvidersEdgesSiteRebalanceAccepted creates a PostTelephonyProvidersEdgesSiteRebalanceAccepted with default headers values
func NewPostTelephonyProvidersEdgesSiteRebalanceAccepted() *PostTelephonyProvidersEdgesSiteRebalanceAccepted {
	return &PostTelephonyProvidersEdgesSiteRebalanceAccepted{}
}

/*PostTelephonyProvidersEdgesSiteRebalanceAccepted handles this case with default header values.

Accepted - Processing the Rebalance
*/
type PostTelephonyProvidersEdgesSiteRebalanceAccepted struct {
}

func (o *PostTelephonyProvidersEdgesSiteRebalanceAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/sites/{siteId}/rebalance][%d] postTelephonyProvidersEdgesSiteRebalanceAccepted ", 202)
}

func (o *PostTelephonyProvidersEdgesSiteRebalanceAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostTelephonyProvidersEdgesSiteRebalanceBadRequest creates a PostTelephonyProvidersEdgesSiteRebalanceBadRequest with default headers values
func NewPostTelephonyProvidersEdgesSiteRebalanceBadRequest() *PostTelephonyProvidersEdgesSiteRebalanceBadRequest {
	return &PostTelephonyProvidersEdgesSiteRebalanceBadRequest{}
}

/*PostTelephonyProvidersEdgesSiteRebalanceBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostTelephonyProvidersEdgesSiteRebalanceBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesSiteRebalanceBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/sites/{siteId}/rebalance][%d] postTelephonyProvidersEdgesSiteRebalanceBadRequest  %+v", 400, o.Payload)
}

func (o *PostTelephonyProvidersEdgesSiteRebalanceBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesSiteRebalanceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesSiteRebalanceUnauthorized creates a PostTelephonyProvidersEdgesSiteRebalanceUnauthorized with default headers values
func NewPostTelephonyProvidersEdgesSiteRebalanceUnauthorized() *PostTelephonyProvidersEdgesSiteRebalanceUnauthorized {
	return &PostTelephonyProvidersEdgesSiteRebalanceUnauthorized{}
}

/*PostTelephonyProvidersEdgesSiteRebalanceUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostTelephonyProvidersEdgesSiteRebalanceUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesSiteRebalanceUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/sites/{siteId}/rebalance][%d] postTelephonyProvidersEdgesSiteRebalanceUnauthorized  %+v", 401, o.Payload)
}

func (o *PostTelephonyProvidersEdgesSiteRebalanceUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesSiteRebalanceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesSiteRebalanceForbidden creates a PostTelephonyProvidersEdgesSiteRebalanceForbidden with default headers values
func NewPostTelephonyProvidersEdgesSiteRebalanceForbidden() *PostTelephonyProvidersEdgesSiteRebalanceForbidden {
	return &PostTelephonyProvidersEdgesSiteRebalanceForbidden{}
}

/*PostTelephonyProvidersEdgesSiteRebalanceForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostTelephonyProvidersEdgesSiteRebalanceForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesSiteRebalanceForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/sites/{siteId}/rebalance][%d] postTelephonyProvidersEdgesSiteRebalanceForbidden  %+v", 403, o.Payload)
}

func (o *PostTelephonyProvidersEdgesSiteRebalanceForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesSiteRebalanceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesSiteRebalanceNotFound creates a PostTelephonyProvidersEdgesSiteRebalanceNotFound with default headers values
func NewPostTelephonyProvidersEdgesSiteRebalanceNotFound() *PostTelephonyProvidersEdgesSiteRebalanceNotFound {
	return &PostTelephonyProvidersEdgesSiteRebalanceNotFound{}
}

/*PostTelephonyProvidersEdgesSiteRebalanceNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostTelephonyProvidersEdgesSiteRebalanceNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesSiteRebalanceNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/sites/{siteId}/rebalance][%d] postTelephonyProvidersEdgesSiteRebalanceNotFound  %+v", 404, o.Payload)
}

func (o *PostTelephonyProvidersEdgesSiteRebalanceNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesSiteRebalanceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesSiteRebalanceRequestEntityTooLarge creates a PostTelephonyProvidersEdgesSiteRebalanceRequestEntityTooLarge with default headers values
func NewPostTelephonyProvidersEdgesSiteRebalanceRequestEntityTooLarge() *PostTelephonyProvidersEdgesSiteRebalanceRequestEntityTooLarge {
	return &PostTelephonyProvidersEdgesSiteRebalanceRequestEntityTooLarge{}
}

/*PostTelephonyProvidersEdgesSiteRebalanceRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostTelephonyProvidersEdgesSiteRebalanceRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesSiteRebalanceRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/sites/{siteId}/rebalance][%d] postTelephonyProvidersEdgesSiteRebalanceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostTelephonyProvidersEdgesSiteRebalanceRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesSiteRebalanceRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesSiteRebalanceUnsupportedMediaType creates a PostTelephonyProvidersEdgesSiteRebalanceUnsupportedMediaType with default headers values
func NewPostTelephonyProvidersEdgesSiteRebalanceUnsupportedMediaType() *PostTelephonyProvidersEdgesSiteRebalanceUnsupportedMediaType {
	return &PostTelephonyProvidersEdgesSiteRebalanceUnsupportedMediaType{}
}

/*PostTelephonyProvidersEdgesSiteRebalanceUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostTelephonyProvidersEdgesSiteRebalanceUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesSiteRebalanceUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/sites/{siteId}/rebalance][%d] postTelephonyProvidersEdgesSiteRebalanceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostTelephonyProvidersEdgesSiteRebalanceUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesSiteRebalanceUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesSiteRebalanceTooManyRequests creates a PostTelephonyProvidersEdgesSiteRebalanceTooManyRequests with default headers values
func NewPostTelephonyProvidersEdgesSiteRebalanceTooManyRequests() *PostTelephonyProvidersEdgesSiteRebalanceTooManyRequests {
	return &PostTelephonyProvidersEdgesSiteRebalanceTooManyRequests{}
}

/*PostTelephonyProvidersEdgesSiteRebalanceTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostTelephonyProvidersEdgesSiteRebalanceTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesSiteRebalanceTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/sites/{siteId}/rebalance][%d] postTelephonyProvidersEdgesSiteRebalanceTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostTelephonyProvidersEdgesSiteRebalanceTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesSiteRebalanceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesSiteRebalanceInternalServerError creates a PostTelephonyProvidersEdgesSiteRebalanceInternalServerError with default headers values
func NewPostTelephonyProvidersEdgesSiteRebalanceInternalServerError() *PostTelephonyProvidersEdgesSiteRebalanceInternalServerError {
	return &PostTelephonyProvidersEdgesSiteRebalanceInternalServerError{}
}

/*PostTelephonyProvidersEdgesSiteRebalanceInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostTelephonyProvidersEdgesSiteRebalanceInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesSiteRebalanceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/sites/{siteId}/rebalance][%d] postTelephonyProvidersEdgesSiteRebalanceInternalServerError  %+v", 500, o.Payload)
}

func (o *PostTelephonyProvidersEdgesSiteRebalanceInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesSiteRebalanceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesSiteRebalanceServiceUnavailable creates a PostTelephonyProvidersEdgesSiteRebalanceServiceUnavailable with default headers values
func NewPostTelephonyProvidersEdgesSiteRebalanceServiceUnavailable() *PostTelephonyProvidersEdgesSiteRebalanceServiceUnavailable {
	return &PostTelephonyProvidersEdgesSiteRebalanceServiceUnavailable{}
}

/*PostTelephonyProvidersEdgesSiteRebalanceServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostTelephonyProvidersEdgesSiteRebalanceServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesSiteRebalanceServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/sites/{siteId}/rebalance][%d] postTelephonyProvidersEdgesSiteRebalanceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostTelephonyProvidersEdgesSiteRebalanceServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesSiteRebalanceServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesSiteRebalanceGatewayTimeout creates a PostTelephonyProvidersEdgesSiteRebalanceGatewayTimeout with default headers values
func NewPostTelephonyProvidersEdgesSiteRebalanceGatewayTimeout() *PostTelephonyProvidersEdgesSiteRebalanceGatewayTimeout {
	return &PostTelephonyProvidersEdgesSiteRebalanceGatewayTimeout{}
}

/*PostTelephonyProvidersEdgesSiteRebalanceGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostTelephonyProvidersEdgesSiteRebalanceGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesSiteRebalanceGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/sites/{siteId}/rebalance][%d] postTelephonyProvidersEdgesSiteRebalanceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostTelephonyProvidersEdgesSiteRebalanceGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesSiteRebalanceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package outbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostOutboundCampaignrulesReader is a Reader for the PostOutboundCampaignrules structure.
type PostOutboundCampaignrulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOutboundCampaignrulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostOutboundCampaignrulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostOutboundCampaignrulesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostOutboundCampaignrulesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostOutboundCampaignrulesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostOutboundCampaignrulesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostOutboundCampaignrulesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostOutboundCampaignrulesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostOutboundCampaignrulesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostOutboundCampaignrulesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostOutboundCampaignrulesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostOutboundCampaignrulesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostOutboundCampaignrulesOK creates a PostOutboundCampaignrulesOK with default headers values
func NewPostOutboundCampaignrulesOK() *PostOutboundCampaignrulesOK {
	return &PostOutboundCampaignrulesOK{}
}

/*PostOutboundCampaignrulesOK handles this case with default header values.

successful operation
*/
type PostOutboundCampaignrulesOK struct {
	Payload *models.CampaignRule
}

func (o *PostOutboundCampaignrulesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaignrules][%d] postOutboundCampaignrulesOK  %+v", 200, o.Payload)
}

func (o *PostOutboundCampaignrulesOK) GetPayload() *models.CampaignRule {
	return o.Payload
}

func (o *PostOutboundCampaignrulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CampaignRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundCampaignrulesBadRequest creates a PostOutboundCampaignrulesBadRequest with default headers values
func NewPostOutboundCampaignrulesBadRequest() *PostOutboundCampaignrulesBadRequest {
	return &PostOutboundCampaignrulesBadRequest{}
}

/*PostOutboundCampaignrulesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostOutboundCampaignrulesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundCampaignrulesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaignrules][%d] postOutboundCampaignrulesBadRequest  %+v", 400, o.Payload)
}

func (o *PostOutboundCampaignrulesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundCampaignrulesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundCampaignrulesUnauthorized creates a PostOutboundCampaignrulesUnauthorized with default headers values
func NewPostOutboundCampaignrulesUnauthorized() *PostOutboundCampaignrulesUnauthorized {
	return &PostOutboundCampaignrulesUnauthorized{}
}

/*PostOutboundCampaignrulesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostOutboundCampaignrulesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundCampaignrulesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaignrules][%d] postOutboundCampaignrulesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostOutboundCampaignrulesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundCampaignrulesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundCampaignrulesForbidden creates a PostOutboundCampaignrulesForbidden with default headers values
func NewPostOutboundCampaignrulesForbidden() *PostOutboundCampaignrulesForbidden {
	return &PostOutboundCampaignrulesForbidden{}
}

/*PostOutboundCampaignrulesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostOutboundCampaignrulesForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundCampaignrulesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaignrules][%d] postOutboundCampaignrulesForbidden  %+v", 403, o.Payload)
}

func (o *PostOutboundCampaignrulesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundCampaignrulesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundCampaignrulesNotFound creates a PostOutboundCampaignrulesNotFound with default headers values
func NewPostOutboundCampaignrulesNotFound() *PostOutboundCampaignrulesNotFound {
	return &PostOutboundCampaignrulesNotFound{}
}

/*PostOutboundCampaignrulesNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostOutboundCampaignrulesNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundCampaignrulesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaignrules][%d] postOutboundCampaignrulesNotFound  %+v", 404, o.Payload)
}

func (o *PostOutboundCampaignrulesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundCampaignrulesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundCampaignrulesRequestEntityTooLarge creates a PostOutboundCampaignrulesRequestEntityTooLarge with default headers values
func NewPostOutboundCampaignrulesRequestEntityTooLarge() *PostOutboundCampaignrulesRequestEntityTooLarge {
	return &PostOutboundCampaignrulesRequestEntityTooLarge{}
}

/*PostOutboundCampaignrulesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostOutboundCampaignrulesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundCampaignrulesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaignrules][%d] postOutboundCampaignrulesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostOutboundCampaignrulesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundCampaignrulesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundCampaignrulesUnsupportedMediaType creates a PostOutboundCampaignrulesUnsupportedMediaType with default headers values
func NewPostOutboundCampaignrulesUnsupportedMediaType() *PostOutboundCampaignrulesUnsupportedMediaType {
	return &PostOutboundCampaignrulesUnsupportedMediaType{}
}

/*PostOutboundCampaignrulesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostOutboundCampaignrulesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundCampaignrulesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaignrules][%d] postOutboundCampaignrulesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostOutboundCampaignrulesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundCampaignrulesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundCampaignrulesTooManyRequests creates a PostOutboundCampaignrulesTooManyRequests with default headers values
func NewPostOutboundCampaignrulesTooManyRequests() *PostOutboundCampaignrulesTooManyRequests {
	return &PostOutboundCampaignrulesTooManyRequests{}
}

/*PostOutboundCampaignrulesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostOutboundCampaignrulesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundCampaignrulesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaignrules][%d] postOutboundCampaignrulesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostOutboundCampaignrulesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundCampaignrulesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundCampaignrulesInternalServerError creates a PostOutboundCampaignrulesInternalServerError with default headers values
func NewPostOutboundCampaignrulesInternalServerError() *PostOutboundCampaignrulesInternalServerError {
	return &PostOutboundCampaignrulesInternalServerError{}
}

/*PostOutboundCampaignrulesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostOutboundCampaignrulesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundCampaignrulesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaignrules][%d] postOutboundCampaignrulesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostOutboundCampaignrulesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundCampaignrulesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundCampaignrulesServiceUnavailable creates a PostOutboundCampaignrulesServiceUnavailable with default headers values
func NewPostOutboundCampaignrulesServiceUnavailable() *PostOutboundCampaignrulesServiceUnavailable {
	return &PostOutboundCampaignrulesServiceUnavailable{}
}

/*PostOutboundCampaignrulesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostOutboundCampaignrulesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundCampaignrulesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaignrules][%d] postOutboundCampaignrulesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostOutboundCampaignrulesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundCampaignrulesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundCampaignrulesGatewayTimeout creates a PostOutboundCampaignrulesGatewayTimeout with default headers values
func NewPostOutboundCampaignrulesGatewayTimeout() *PostOutboundCampaignrulesGatewayTimeout {
	return &PostOutboundCampaignrulesGatewayTimeout{}
}

/*PostOutboundCampaignrulesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostOutboundCampaignrulesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundCampaignrulesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaignrules][%d] postOutboundCampaignrulesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostOutboundCampaignrulesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundCampaignrulesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

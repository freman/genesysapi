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

// GetTelephonyProvidersEdgesTrunkswithrecordingReader is a Reader for the GetTelephonyProvidersEdgesTrunkswithrecording structure.
type GetTelephonyProvidersEdgesTrunkswithrecordingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonyProvidersEdgesTrunkswithrecordingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonyProvidersEdgesTrunkswithrecordingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonyProvidersEdgesTrunkswithrecordingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonyProvidersEdgesTrunkswithrecordingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonyProvidersEdgesTrunkswithrecordingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonyProvidersEdgesTrunkswithrecordingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonyProvidersEdgesTrunkswithrecordingRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonyProvidersEdgesTrunkswithrecordingUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonyProvidersEdgesTrunkswithrecordingTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonyProvidersEdgesTrunkswithrecordingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonyProvidersEdgesTrunkswithrecordingServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonyProvidersEdgesTrunkswithrecordingGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonyProvidersEdgesTrunkswithrecordingOK creates a GetTelephonyProvidersEdgesTrunkswithrecordingOK with default headers values
func NewGetTelephonyProvidersEdgesTrunkswithrecordingOK() *GetTelephonyProvidersEdgesTrunkswithrecordingOK {
	return &GetTelephonyProvidersEdgesTrunkswithrecordingOK{}
}

/*GetTelephonyProvidersEdgesTrunkswithrecordingOK handles this case with default header values.

successful operation
*/
type GetTelephonyProvidersEdgesTrunkswithrecordingOK struct {
	Payload *models.TrunkRecordingEnabledCount
}

func (o *GetTelephonyProvidersEdgesTrunkswithrecordingOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/trunkswithrecording][%d] getTelephonyProvidersEdgesTrunkswithrecordingOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgesTrunkswithrecordingOK) GetPayload() *models.TrunkRecordingEnabledCount {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesTrunkswithrecordingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TrunkRecordingEnabledCount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesTrunkswithrecordingBadRequest creates a GetTelephonyProvidersEdgesTrunkswithrecordingBadRequest with default headers values
func NewGetTelephonyProvidersEdgesTrunkswithrecordingBadRequest() *GetTelephonyProvidersEdgesTrunkswithrecordingBadRequest {
	return &GetTelephonyProvidersEdgesTrunkswithrecordingBadRequest{}
}

/*GetTelephonyProvidersEdgesTrunkswithrecordingBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyProvidersEdgesTrunkswithrecordingBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesTrunkswithrecordingBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/trunkswithrecording][%d] getTelephonyProvidersEdgesTrunkswithrecordingBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgesTrunkswithrecordingBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesTrunkswithrecordingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesTrunkswithrecordingUnauthorized creates a GetTelephonyProvidersEdgesTrunkswithrecordingUnauthorized with default headers values
func NewGetTelephonyProvidersEdgesTrunkswithrecordingUnauthorized() *GetTelephonyProvidersEdgesTrunkswithrecordingUnauthorized {
	return &GetTelephonyProvidersEdgesTrunkswithrecordingUnauthorized{}
}

/*GetTelephonyProvidersEdgesTrunkswithrecordingUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyProvidersEdgesTrunkswithrecordingUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesTrunkswithrecordingUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/trunkswithrecording][%d] getTelephonyProvidersEdgesTrunkswithrecordingUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgesTrunkswithrecordingUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesTrunkswithrecordingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesTrunkswithrecordingForbidden creates a GetTelephonyProvidersEdgesTrunkswithrecordingForbidden with default headers values
func NewGetTelephonyProvidersEdgesTrunkswithrecordingForbidden() *GetTelephonyProvidersEdgesTrunkswithrecordingForbidden {
	return &GetTelephonyProvidersEdgesTrunkswithrecordingForbidden{}
}

/*GetTelephonyProvidersEdgesTrunkswithrecordingForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyProvidersEdgesTrunkswithrecordingForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesTrunkswithrecordingForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/trunkswithrecording][%d] getTelephonyProvidersEdgesTrunkswithrecordingForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgesTrunkswithrecordingForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesTrunkswithrecordingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesTrunkswithrecordingNotFound creates a GetTelephonyProvidersEdgesTrunkswithrecordingNotFound with default headers values
func NewGetTelephonyProvidersEdgesTrunkswithrecordingNotFound() *GetTelephonyProvidersEdgesTrunkswithrecordingNotFound {
	return &GetTelephonyProvidersEdgesTrunkswithrecordingNotFound{}
}

/*GetTelephonyProvidersEdgesTrunkswithrecordingNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetTelephonyProvidersEdgesTrunkswithrecordingNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesTrunkswithrecordingNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/trunkswithrecording][%d] getTelephonyProvidersEdgesTrunkswithrecordingNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgesTrunkswithrecordingNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesTrunkswithrecordingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesTrunkswithrecordingRequestEntityTooLarge creates a GetTelephonyProvidersEdgesTrunkswithrecordingRequestEntityTooLarge with default headers values
func NewGetTelephonyProvidersEdgesTrunkswithrecordingRequestEntityTooLarge() *GetTelephonyProvidersEdgesTrunkswithrecordingRequestEntityTooLarge {
	return &GetTelephonyProvidersEdgesTrunkswithrecordingRequestEntityTooLarge{}
}

/*GetTelephonyProvidersEdgesTrunkswithrecordingRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetTelephonyProvidersEdgesTrunkswithrecordingRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesTrunkswithrecordingRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/trunkswithrecording][%d] getTelephonyProvidersEdgesTrunkswithrecordingRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgesTrunkswithrecordingRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesTrunkswithrecordingRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesTrunkswithrecordingUnsupportedMediaType creates a GetTelephonyProvidersEdgesTrunkswithrecordingUnsupportedMediaType with default headers values
func NewGetTelephonyProvidersEdgesTrunkswithrecordingUnsupportedMediaType() *GetTelephonyProvidersEdgesTrunkswithrecordingUnsupportedMediaType {
	return &GetTelephonyProvidersEdgesTrunkswithrecordingUnsupportedMediaType{}
}

/*GetTelephonyProvidersEdgesTrunkswithrecordingUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyProvidersEdgesTrunkswithrecordingUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesTrunkswithrecordingUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/trunkswithrecording][%d] getTelephonyProvidersEdgesTrunkswithrecordingUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgesTrunkswithrecordingUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesTrunkswithrecordingUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesTrunkswithrecordingTooManyRequests creates a GetTelephonyProvidersEdgesTrunkswithrecordingTooManyRequests with default headers values
func NewGetTelephonyProvidersEdgesTrunkswithrecordingTooManyRequests() *GetTelephonyProvidersEdgesTrunkswithrecordingTooManyRequests {
	return &GetTelephonyProvidersEdgesTrunkswithrecordingTooManyRequests{}
}

/*GetTelephonyProvidersEdgesTrunkswithrecordingTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetTelephonyProvidersEdgesTrunkswithrecordingTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesTrunkswithrecordingTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/trunkswithrecording][%d] getTelephonyProvidersEdgesTrunkswithrecordingTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgesTrunkswithrecordingTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesTrunkswithrecordingTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesTrunkswithrecordingInternalServerError creates a GetTelephonyProvidersEdgesTrunkswithrecordingInternalServerError with default headers values
func NewGetTelephonyProvidersEdgesTrunkswithrecordingInternalServerError() *GetTelephonyProvidersEdgesTrunkswithrecordingInternalServerError {
	return &GetTelephonyProvidersEdgesTrunkswithrecordingInternalServerError{}
}

/*GetTelephonyProvidersEdgesTrunkswithrecordingInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyProvidersEdgesTrunkswithrecordingInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesTrunkswithrecordingInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/trunkswithrecording][%d] getTelephonyProvidersEdgesTrunkswithrecordingInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgesTrunkswithrecordingInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesTrunkswithrecordingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesTrunkswithrecordingServiceUnavailable creates a GetTelephonyProvidersEdgesTrunkswithrecordingServiceUnavailable with default headers values
func NewGetTelephonyProvidersEdgesTrunkswithrecordingServiceUnavailable() *GetTelephonyProvidersEdgesTrunkswithrecordingServiceUnavailable {
	return &GetTelephonyProvidersEdgesTrunkswithrecordingServiceUnavailable{}
}

/*GetTelephonyProvidersEdgesTrunkswithrecordingServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyProvidersEdgesTrunkswithrecordingServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesTrunkswithrecordingServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/trunkswithrecording][%d] getTelephonyProvidersEdgesTrunkswithrecordingServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgesTrunkswithrecordingServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesTrunkswithrecordingServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesTrunkswithrecordingGatewayTimeout creates a GetTelephonyProvidersEdgesTrunkswithrecordingGatewayTimeout with default headers values
func NewGetTelephonyProvidersEdgesTrunkswithrecordingGatewayTimeout() *GetTelephonyProvidersEdgesTrunkswithrecordingGatewayTimeout {
	return &GetTelephonyProvidersEdgesTrunkswithrecordingGatewayTimeout{}
}

/*GetTelephonyProvidersEdgesTrunkswithrecordingGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetTelephonyProvidersEdgesTrunkswithrecordingGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesTrunkswithrecordingGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/trunkswithrecording][%d] getTelephonyProvidersEdgesTrunkswithrecordingGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgesTrunkswithrecordingGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesTrunkswithrecordingGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

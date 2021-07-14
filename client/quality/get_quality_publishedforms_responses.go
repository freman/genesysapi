// Code generated by go-swagger; DO NOT EDIT.

package quality

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetQualityPublishedformsReader is a Reader for the GetQualityPublishedforms structure.
type GetQualityPublishedformsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetQualityPublishedformsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetQualityPublishedformsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetQualityPublishedformsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetQualityPublishedformsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetQualityPublishedformsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetQualityPublishedformsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetQualityPublishedformsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetQualityPublishedformsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetQualityPublishedformsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetQualityPublishedformsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetQualityPublishedformsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetQualityPublishedformsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetQualityPublishedformsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetQualityPublishedformsOK creates a GetQualityPublishedformsOK with default headers values
func NewGetQualityPublishedformsOK() *GetQualityPublishedformsOK {
	return &GetQualityPublishedformsOK{}
}

/*GetQualityPublishedformsOK handles this case with default header values.

successful operation
*/
type GetQualityPublishedformsOK struct {
	Payload *models.EvaluationFormEntityListing
}

func (o *GetQualityPublishedformsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms][%d] getQualityPublishedformsOK  %+v", 200, o.Payload)
}

func (o *GetQualityPublishedformsOK) GetPayload() *models.EvaluationFormEntityListing {
	return o.Payload
}

func (o *GetQualityPublishedformsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EvaluationFormEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsBadRequest creates a GetQualityPublishedformsBadRequest with default headers values
func NewGetQualityPublishedformsBadRequest() *GetQualityPublishedformsBadRequest {
	return &GetQualityPublishedformsBadRequest{}
}

/*GetQualityPublishedformsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetQualityPublishedformsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetQualityPublishedformsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms][%d] getQualityPublishedformsBadRequest  %+v", 400, o.Payload)
}

func (o *GetQualityPublishedformsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsUnauthorized creates a GetQualityPublishedformsUnauthorized with default headers values
func NewGetQualityPublishedformsUnauthorized() *GetQualityPublishedformsUnauthorized {
	return &GetQualityPublishedformsUnauthorized{}
}

/*GetQualityPublishedformsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetQualityPublishedformsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetQualityPublishedformsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms][%d] getQualityPublishedformsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetQualityPublishedformsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsForbidden creates a GetQualityPublishedformsForbidden with default headers values
func NewGetQualityPublishedformsForbidden() *GetQualityPublishedformsForbidden {
	return &GetQualityPublishedformsForbidden{}
}

/*GetQualityPublishedformsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetQualityPublishedformsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetQualityPublishedformsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms][%d] getQualityPublishedformsForbidden  %+v", 403, o.Payload)
}

func (o *GetQualityPublishedformsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsNotFound creates a GetQualityPublishedformsNotFound with default headers values
func NewGetQualityPublishedformsNotFound() *GetQualityPublishedformsNotFound {
	return &GetQualityPublishedformsNotFound{}
}

/*GetQualityPublishedformsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetQualityPublishedformsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetQualityPublishedformsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms][%d] getQualityPublishedformsNotFound  %+v", 404, o.Payload)
}

func (o *GetQualityPublishedformsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsRequestTimeout creates a GetQualityPublishedformsRequestTimeout with default headers values
func NewGetQualityPublishedformsRequestTimeout() *GetQualityPublishedformsRequestTimeout {
	return &GetQualityPublishedformsRequestTimeout{}
}

/*GetQualityPublishedformsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetQualityPublishedformsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetQualityPublishedformsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms][%d] getQualityPublishedformsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetQualityPublishedformsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsRequestEntityTooLarge creates a GetQualityPublishedformsRequestEntityTooLarge with default headers values
func NewGetQualityPublishedformsRequestEntityTooLarge() *GetQualityPublishedformsRequestEntityTooLarge {
	return &GetQualityPublishedformsRequestEntityTooLarge{}
}

/*GetQualityPublishedformsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetQualityPublishedformsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetQualityPublishedformsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms][%d] getQualityPublishedformsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetQualityPublishedformsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsUnsupportedMediaType creates a GetQualityPublishedformsUnsupportedMediaType with default headers values
func NewGetQualityPublishedformsUnsupportedMediaType() *GetQualityPublishedformsUnsupportedMediaType {
	return &GetQualityPublishedformsUnsupportedMediaType{}
}

/*GetQualityPublishedformsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetQualityPublishedformsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetQualityPublishedformsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms][%d] getQualityPublishedformsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetQualityPublishedformsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsTooManyRequests creates a GetQualityPublishedformsTooManyRequests with default headers values
func NewGetQualityPublishedformsTooManyRequests() *GetQualityPublishedformsTooManyRequests {
	return &GetQualityPublishedformsTooManyRequests{}
}

/*GetQualityPublishedformsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetQualityPublishedformsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetQualityPublishedformsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms][%d] getQualityPublishedformsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetQualityPublishedformsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsInternalServerError creates a GetQualityPublishedformsInternalServerError with default headers values
func NewGetQualityPublishedformsInternalServerError() *GetQualityPublishedformsInternalServerError {
	return &GetQualityPublishedformsInternalServerError{}
}

/*GetQualityPublishedformsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetQualityPublishedformsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetQualityPublishedformsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms][%d] getQualityPublishedformsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetQualityPublishedformsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsServiceUnavailable creates a GetQualityPublishedformsServiceUnavailable with default headers values
func NewGetQualityPublishedformsServiceUnavailable() *GetQualityPublishedformsServiceUnavailable {
	return &GetQualityPublishedformsServiceUnavailable{}
}

/*GetQualityPublishedformsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetQualityPublishedformsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetQualityPublishedformsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms][%d] getQualityPublishedformsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetQualityPublishedformsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsGatewayTimeout creates a GetQualityPublishedformsGatewayTimeout with default headers values
func NewGetQualityPublishedformsGatewayTimeout() *GetQualityPublishedformsGatewayTimeout {
	return &GetQualityPublishedformsGatewayTimeout{}
}

/*GetQualityPublishedformsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetQualityPublishedformsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetQualityPublishedformsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms][%d] getQualityPublishedformsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetQualityPublishedformsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

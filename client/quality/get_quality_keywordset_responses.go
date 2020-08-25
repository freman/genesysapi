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

// GetQualityKeywordsetReader is a Reader for the GetQualityKeywordset structure.
type GetQualityKeywordsetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetQualityKeywordsetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetQualityKeywordsetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetQualityKeywordsetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetQualityKeywordsetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetQualityKeywordsetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetQualityKeywordsetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetQualityKeywordsetRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetQualityKeywordsetUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetQualityKeywordsetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetQualityKeywordsetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetQualityKeywordsetServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetQualityKeywordsetGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetQualityKeywordsetOK creates a GetQualityKeywordsetOK with default headers values
func NewGetQualityKeywordsetOK() *GetQualityKeywordsetOK {
	return &GetQualityKeywordsetOK{}
}

/*GetQualityKeywordsetOK handles this case with default header values.

successful operation
*/
type GetQualityKeywordsetOK struct {
	Payload *models.KeywordSet
}

func (o *GetQualityKeywordsetOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/keywordsets/{keywordSetId}][%d] getQualityKeywordsetOK  %+v", 200, o.Payload)
}

func (o *GetQualityKeywordsetOK) GetPayload() *models.KeywordSet {
	return o.Payload
}

func (o *GetQualityKeywordsetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KeywordSet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityKeywordsetBadRequest creates a GetQualityKeywordsetBadRequest with default headers values
func NewGetQualityKeywordsetBadRequest() *GetQualityKeywordsetBadRequest {
	return &GetQualityKeywordsetBadRequest{}
}

/*GetQualityKeywordsetBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetQualityKeywordsetBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetQualityKeywordsetBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/keywordsets/{keywordSetId}][%d] getQualityKeywordsetBadRequest  %+v", 400, o.Payload)
}

func (o *GetQualityKeywordsetBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityKeywordsetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityKeywordsetUnauthorized creates a GetQualityKeywordsetUnauthorized with default headers values
func NewGetQualityKeywordsetUnauthorized() *GetQualityKeywordsetUnauthorized {
	return &GetQualityKeywordsetUnauthorized{}
}

/*GetQualityKeywordsetUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetQualityKeywordsetUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetQualityKeywordsetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/keywordsets/{keywordSetId}][%d] getQualityKeywordsetUnauthorized  %+v", 401, o.Payload)
}

func (o *GetQualityKeywordsetUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityKeywordsetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityKeywordsetForbidden creates a GetQualityKeywordsetForbidden with default headers values
func NewGetQualityKeywordsetForbidden() *GetQualityKeywordsetForbidden {
	return &GetQualityKeywordsetForbidden{}
}

/*GetQualityKeywordsetForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetQualityKeywordsetForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetQualityKeywordsetForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/keywordsets/{keywordSetId}][%d] getQualityKeywordsetForbidden  %+v", 403, o.Payload)
}

func (o *GetQualityKeywordsetForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityKeywordsetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityKeywordsetNotFound creates a GetQualityKeywordsetNotFound with default headers values
func NewGetQualityKeywordsetNotFound() *GetQualityKeywordsetNotFound {
	return &GetQualityKeywordsetNotFound{}
}

/*GetQualityKeywordsetNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetQualityKeywordsetNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetQualityKeywordsetNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/keywordsets/{keywordSetId}][%d] getQualityKeywordsetNotFound  %+v", 404, o.Payload)
}

func (o *GetQualityKeywordsetNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityKeywordsetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityKeywordsetRequestEntityTooLarge creates a GetQualityKeywordsetRequestEntityTooLarge with default headers values
func NewGetQualityKeywordsetRequestEntityTooLarge() *GetQualityKeywordsetRequestEntityTooLarge {
	return &GetQualityKeywordsetRequestEntityTooLarge{}
}

/*GetQualityKeywordsetRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetQualityKeywordsetRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetQualityKeywordsetRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/keywordsets/{keywordSetId}][%d] getQualityKeywordsetRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetQualityKeywordsetRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityKeywordsetRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityKeywordsetUnsupportedMediaType creates a GetQualityKeywordsetUnsupportedMediaType with default headers values
func NewGetQualityKeywordsetUnsupportedMediaType() *GetQualityKeywordsetUnsupportedMediaType {
	return &GetQualityKeywordsetUnsupportedMediaType{}
}

/*GetQualityKeywordsetUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetQualityKeywordsetUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetQualityKeywordsetUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/keywordsets/{keywordSetId}][%d] getQualityKeywordsetUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetQualityKeywordsetUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityKeywordsetUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityKeywordsetTooManyRequests creates a GetQualityKeywordsetTooManyRequests with default headers values
func NewGetQualityKeywordsetTooManyRequests() *GetQualityKeywordsetTooManyRequests {
	return &GetQualityKeywordsetTooManyRequests{}
}

/*GetQualityKeywordsetTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetQualityKeywordsetTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetQualityKeywordsetTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/keywordsets/{keywordSetId}][%d] getQualityKeywordsetTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetQualityKeywordsetTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityKeywordsetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityKeywordsetInternalServerError creates a GetQualityKeywordsetInternalServerError with default headers values
func NewGetQualityKeywordsetInternalServerError() *GetQualityKeywordsetInternalServerError {
	return &GetQualityKeywordsetInternalServerError{}
}

/*GetQualityKeywordsetInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetQualityKeywordsetInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetQualityKeywordsetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/keywordsets/{keywordSetId}][%d] getQualityKeywordsetInternalServerError  %+v", 500, o.Payload)
}

func (o *GetQualityKeywordsetInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityKeywordsetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityKeywordsetServiceUnavailable creates a GetQualityKeywordsetServiceUnavailable with default headers values
func NewGetQualityKeywordsetServiceUnavailable() *GetQualityKeywordsetServiceUnavailable {
	return &GetQualityKeywordsetServiceUnavailable{}
}

/*GetQualityKeywordsetServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetQualityKeywordsetServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetQualityKeywordsetServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/keywordsets/{keywordSetId}][%d] getQualityKeywordsetServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetQualityKeywordsetServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityKeywordsetServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityKeywordsetGatewayTimeout creates a GetQualityKeywordsetGatewayTimeout with default headers values
func NewGetQualityKeywordsetGatewayTimeout() *GetQualityKeywordsetGatewayTimeout {
	return &GetQualityKeywordsetGatewayTimeout{}
}

/*GetQualityKeywordsetGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetQualityKeywordsetGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetQualityKeywordsetGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/keywordsets/{keywordSetId}][%d] getQualityKeywordsetGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetQualityKeywordsetGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityKeywordsetGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
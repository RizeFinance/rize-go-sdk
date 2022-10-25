package rize

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

// Handles all DebitCard operations
type debitCardService service

// DebitCard data type
type DebitCard struct {
	UID                   string                    `json:"uid,omitempty"`
	ExternalUID           string                    `json:"external_uid,omitempty"`
	CustomerUID           string                    `json:"customer_uid,omitempty"`
	PoolUID               string                    `json:"pool_uid,omitempty"`
	SyntheticAccountUID   string                    `json:"synthetic_account_uid,omitempty"`
	CustodialAccountUID   string                    `json:"custodial_account_uid,omitempty"`
	CardArtworkUID        string                    `json:"card_artwork_uid,omitempty"`
	CardLastFourDigits    string                    `json:"card_last_four_digits,omitempty"`
	Status                string                    `json:"status,omitempty"`
	Type                  string                    `json:"type,omitempty"`
	ReadyToUse            bool                      `json:"ready_to_use,omitempty"`
	LockReason            string                    `json:"lock_reason,omitempty"`
	IssuedOn              string                    `json:"issued_on,omitempty"`
	LockedAt              time.Time                 `json:"locked_at,omitempty"`
	ClosedAt              time.Time                 `json:"closed_at,omitempty"`
	LatestShippingAddress *DebitCardShippingAddress `json:"latest_shipping_address,omitempty"`
}

// DebitCardShippingAddress is an optional field used to specify the shipping address for a physical Debit Card.
type DebitCardShippingAddress struct {
	Street1    string `json:"street1,omitempty"`
	Street2    string `json:"street2,omitempty"`
	City       string `json:"city,omitempty"`
	State      string `json:"state,omitempty"`
	PostalCode string `json:"postal_code,omitempty"`
}

// DebitCardAccessToken contains the token necessary to retrieve a virtual Debit Card image.
// Used as the response type for GetAccessToken.
// Used as query params for GetVirtualDebitCardImage.
type DebitCardAccessToken struct {
	Token    string `url:"token" json:"token"`
	ConfigID string `url:"config_id" json:"config_id"`
}

// DebitCardListParams builds the query parameters used in querying Debit Cards
type DebitCardListParams struct {
	CustomerUID string `url:"customer_uid,omitempty" json:"customer_uid,omitempty"`
	ExternalUID string `url:"external_uid,omitempty" json:"external_uid,omitempty"`
	Limit       int    `url:"limit,omitempty" json:"limit,omitempty"`
	Offset      int    `url:"offset,omitempty" json:"offset,omitempty"`
	PoolUID     string `url:"pool_uid,omitempty" json:"pool_uid,omitempty"`
	Locked      bool   `url:"locked,omitempty" json:"locked,omitempty"`
	Status      string `url:"status,omitempty" json:"status,omitempty"`
}

// DebitCardCreateParams are the body params used when creating a new Debit Card
type DebitCardCreateParams struct {
	ExternalUID     string                    `json:"external_uid,omitempty"`
	CardArtworkUID  string                    `json:"card_artwork_uid,omitempty"`
	CustomerUID     string                    `json:"customer_uid"`
	PoolUID         string                    `json:"pool_uid"`
	ShippingAddress *DebitCardShippingAddress `json:"shipping_address,omitempty"`
}

// DebitCardActivateParams are the body params used when activating a new Debit Card
type DebitCardActivateParams struct {
	CardLastFourDigits string `json:"card_last_four_digits"`
	CVV                string `json:"cvv"`
	ExpiryDate         string `json:"expiry_date"`
}

// DebitCardLockParams are the body params used when locking a Debit Card
type DebitCardLockParams struct {
	LockReason string `json:"lock_reason"`
}

// DebitCardReissueParams are the body params used when reissuing a Debit Card
type DebitCardReissueParams struct {
	CardArtworkUID  string                    `json:"card_artwork_uid,omitempty"`
	ReissueReason   string                    `json:"reissue_reason"`
	ShippingAddress *DebitCardShippingAddress `json:"shipping_address,omitempty"`
}

// DebitCardGetPINTokenParams are the query params used fetching a Debit Card PIN reset token
type DebitCardGetPINTokenParams struct {
	ForceReset bool `url:"force_reset" json:"force_reset"`
}

// VirtualDebitCardMigrateParams are the body params used when migrating a Virtual Debit Card
type VirtualDebitCardMigrateParams struct {
	ExternalUID     string                    `json:"external_uid,omitempty"`
	CardArtworkUID  string                    `json:"card_artwork_uid,omitempty"`
	ShippingAddress *DebitCardShippingAddress `json:"shipping_address,omitempty"`
}

// DebitCardListResponse is an API response containing a list of Debit Cards
type DebitCardListResponse struct {
	ListResponse
	Data []*DebitCard `json:"data"`
}

// DebitCardPINTokenResponse is an API response containing a token necessary to change a Debit Card's PIN
type DebitCardPINTokenResponse struct {
	PinChangeToken string `json:"pin_change_token"`
}

// List retrieves a list of Debit Cards filtered by the given parameters
func (d *debitCardService) List(ctx context.Context, params *DebitCardListParams) (*DebitCardListResponse, error) {
	// Build DebitCardListParams into query string params
	v, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	res, err := d.client.doRequest(ctx, http.MethodGet, "debit_cards", v, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &DebitCardListResponse{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Create is used to a new Debit Card and attach it to the supplied Customer and Pool
func (d *debitCardService) Create(ctx context.Context, params *DebitCardCreateParams) (*DebitCard, error) {
	if params.CustomerUID == "" || params.PoolUID == "" {
		return nil, fmt.Errorf("CustomerUID and PoolUID are required")
	}

	bytesMessage, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	res, err := d.client.doRequest(ctx, http.MethodPost, "debit_cards", nil, bytes.NewBuffer(bytesMessage))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &DebitCard{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Get returns a single DebitCard
func (d *debitCardService) Get(ctx context.Context, uid string) (*DebitCard, error) {
	if uid == "" {
		return nil, fmt.Errorf("UID is required")
	}

	res, err := d.client.doRequest(ctx, http.MethodGet, fmt.Sprintf("debit_cards/%s", uid), nil, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &DebitCard{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Activate a Debit Card
func (d *debitCardService) Activate(ctx context.Context, uid string, params *DebitCardActivateParams) (*DebitCard, error) {
	if uid == "" {
		return nil, fmt.Errorf("UID is required")
	}

	if params.CardLastFourDigits == "" || params.CVV == "" || params.ExpiryDate == "" {
		return nil, fmt.Errorf("CardLastFourDigits, CVV and ExpiryDate are required")
	}

	bytesMessage, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	res, err := d.client.doRequest(ctx, http.MethodPut, fmt.Sprintf("debit_cards/%s/activate", uid), nil, bytes.NewBuffer(bytesMessage))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &DebitCard{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Lock will temporarily lock the Debit Card
func (d *debitCardService) Lock(ctx context.Context, uid string, params *DebitCardLockParams) (*DebitCard, error) {
	if uid == "" || params.LockReason == "" {
		return nil, fmt.Errorf("UID and LockReason are required")
	}

	bytesMessage, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	res, err := d.client.doRequest(ctx, http.MethodPut, fmt.Sprintf("debit_cards/%s/lock", uid), nil, bytes.NewBuffer(bytesMessage))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &DebitCard{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Unlock will attempt to remove a lock placed on a Debit Card
func (d *debitCardService) Unlock(ctx context.Context, uid string) (*DebitCard, error) {
	if uid == "" {
		return nil, fmt.Errorf("UID is required")
	}

	res, err := d.client.doRequest(ctx, http.MethodPut, fmt.Sprintf("debit_cards/%s/unlock", uid), nil, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &DebitCard{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Reissue a Debit Card that is lost or stolen, or when it has suffered damage
func (d *debitCardService) Reissue(ctx context.Context, uid string, params *DebitCardReissueParams) (*DebitCard, error) {
	if uid == "" || params.ReissueReason == "" {
		return nil, fmt.Errorf("UID and ReissueReason are required")
	}

	bytesMessage, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	res, err := d.client.doRequest(ctx, http.MethodPut, fmt.Sprintf("debit_cards/%s/reissue", uid), nil, bytes.NewBuffer(bytesMessage))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &DebitCard{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// GetPINToken is used to retrieve a token necessary to change a Debit Card's PIN
func (d *debitCardService) GetPINToken(ctx context.Context, uid string, params *DebitCardGetPINTokenParams) (*DebitCardPINTokenResponse, error) {
	if uid == "" {
		return nil, fmt.Errorf("UID is required")
	}

	v, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	res, err := d.client.doRequest(ctx, http.MethodGet, fmt.Sprintf("debit_cards/%s/pin_change_token", uid), v, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &DebitCardPINTokenResponse{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// GetAccessToken  is used to retrieve the configuration ID and token necessary to retrieve a virtual Debit Card image
func (d *debitCardService) GetAccessToken(ctx context.Context, uid string) (*DebitCardAccessToken, error) {
	if uid == "" {
		return nil, fmt.Errorf("UID is required")
	}

	res, err := d.client.doRequest(ctx, http.MethodGet, fmt.Sprintf("debit_cards/%s/access_token", uid), nil, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &DebitCardAccessToken{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// MigrateVirtualDebitCard will result in a physical version of the virtual debit card being issued to a Customer
func (d *debitCardService) MigrateVirtualDebitCard(ctx context.Context, uid string, params *VirtualDebitCardMigrateParams) (*DebitCard, error) {
	if uid == "" {
		return nil, fmt.Errorf("UID is required")
	}

	bytesMessage, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	res, err := d.client.doRequest(ctx, http.MethodPut, fmt.Sprintf("debit_cards/%s/migrate", uid), nil, bytes.NewBuffer(bytesMessage))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &DebitCard{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// GetVirtualDebitCardImage is used to retrieve a virtual Debit Card image
func (d *debitCardService) GetVirtualDebitCardImage(ctx context.Context, params *DebitCardAccessToken) (*http.Response, error) {
	if params.ConfigID == "" || params.Token == "" {
		return nil, fmt.Errorf("Config and Token params are required")
	}

	v, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	// TODO: Does this require a different Accept header type (image/jpeg)?
	res, err := d.client.doRequest(ctx, http.MethodGet, "assets/virtual_card_image", v, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return res, nil
}

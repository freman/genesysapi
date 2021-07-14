// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	goerrors "errors"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/go-openapi/runtime"
	rtclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"

	"github.com/freman/genesysapi/client/alerting"
	"github.com/freman/genesysapi/client/analytics"
	"github.com/freman/genesysapi/client/architect"
	"github.com/freman/genesysapi/client/audit"
	"github.com/freman/genesysapi/client/authorization"
	"github.com/freman/genesysapi/client/billing"
	"github.com/freman/genesysapi/client/chat"
	"github.com/freman/genesysapi/client/coaching"
	"github.com/freman/genesysapi/client/content_management"
	"github.com/freman/genesysapi/client/conversations"
	"github.com/freman/genesysapi/client/data_extensions"
	"github.com/freman/genesysapi/client/external_contacts"
	"github.com/freman/genesysapi/client/fax"
	"github.com/freman/genesysapi/client/flows"
	"github.com/freman/genesysapi/client/gamification"
	"github.com/freman/genesysapi/client/general_data_protection_regulation"
	"github.com/freman/genesysapi/client/geolocation"
	"github.com/freman/genesysapi/client/greetings"
	"github.com/freman/genesysapi/client/groups"
	"github.com/freman/genesysapi/client/identity_provider"
	"github.com/freman/genesysapi/client/integrations"
	"github.com/freman/genesysapi/client/journey"
	"github.com/freman/genesysapi/client/knowledge"
	"github.com/freman/genesysapi/client/language_understanding"
	"github.com/freman/genesysapi/client/languages"
	"github.com/freman/genesysapi/client/learning"
	"github.com/freman/genesysapi/client/license"
	"github.com/freman/genesysapi/client/locations"
	"github.com/freman/genesysapi/client/mobile_devices"
	"github.com/freman/genesysapi/client/notifications"
	"github.com/freman/genesysapi/client/o_auth"
	"github.com/freman/genesysapi/client/organization"
	"github.com/freman/genesysapi/client/organization_authorization"
	"github.com/freman/genesysapi/client/outbound"
	"github.com/freman/genesysapi/client/presence"
	"github.com/freman/genesysapi/client/quality"
	"github.com/freman/genesysapi/client/recording"
	"github.com/freman/genesysapi/client/response_management"
	"github.com/freman/genesysapi/client/routing"
	"github.com/freman/genesysapi/client/s_c_i_m"
	"github.com/freman/genesysapi/client/scripts"
	"github.com/freman/genesysapi/client/search"
	"github.com/freman/genesysapi/client/speech_and_text_analytics"
	"github.com/freman/genesysapi/client/stations"
	"github.com/freman/genesysapi/client/telephony"
	"github.com/freman/genesysapi/client/telephony_providers_edge"
	"github.com/freman/genesysapi/client/textbots"
	"github.com/freman/genesysapi/client/tokens"
	"github.com/freman/genesysapi/client/uploads"
	"github.com/freman/genesysapi/client/usage"
	"github.com/freman/genesysapi/client/user_recordings"
	"github.com/freman/genesysapi/client/users"
	"github.com/freman/genesysapi/client/utilities"
	"github.com/freman/genesysapi/client/voicemail"
	"github.com/freman/genesysapi/client/web_chat"
	"github.com/freman/genesysapi/client/widgets"
	"github.com/freman/genesysapi/client/workforce_management"
)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "api.mypurecloud.com"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

type Config struct {
	// URL is the base URL of the upstream server
	URL *url.URL
	// Transport is an inner transport for the client
	Transport http.RoundTripper
}

func (c *Config) AuthorizeClientCredentials(clientID, clientSecret string) error {
	var (
		host = DefaultHost
	)

	if c.URL != nil {
		host = c.URL.Host
	}

	ctx := context.Background()

	if c.Transport != nil {
		if _, isa := c.Transport.(*oauth2.Transport); isa {
			return goerrors.New("transport already configured")
		}
		ctx = context.WithValue(ctx, oauth2.HTTPClient, &http.Client{Transport: c.Transport})
	}

	authHost := strings.Replace(host, "api.", "login.", 1)
	oauth2Config := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     "https://" + path.Join(authHost, "oauth/token"),
		AuthStyle:    oauth2.AuthStyleInHeader,
	}

	client := oauth2Config.Client(ctx)
	if _, err := client.Transport.(*oauth2.Transport).Source.Token(); err != nil {
		return err
	}
	c.Transport = client.Transport

	return nil
}

// New creates a new genesys HTTP client.
func New(c Config) *Genesys {
	var (
		host     = DefaultHost
		basePath = DefaultBasePath
		schemes  = DefaultSchemes
	)

	if c.URL != nil {
		host = c.URL.Host
		basePath = c.URL.Path
		schemes = []string{c.URL.Scheme}
	}

	transport := rtclient.New(host, basePath, schemes)
	if c.Transport != nil {
		transport.Transport = c.Transport
	}

	cli := new(Genesys)
	cli.Transport = transport
	cli.Alerting = alerting.New(transport, strfmt.Default, nil)
	cli.Analytics = analytics.New(transport, strfmt.Default, nil)
	cli.Architect = architect.New(transport, strfmt.Default, nil)
	cli.Audit = audit.New(transport, strfmt.Default, nil)
	cli.Authorization = authorization.New(transport, strfmt.Default, nil)
	cli.Billing = billing.New(transport, strfmt.Default, nil)
	cli.Chat = chat.New(transport, strfmt.Default, nil)
	cli.Coaching = coaching.New(transport, strfmt.Default, nil)
	cli.ContentManagement = content_management.New(transport, strfmt.Default, nil)
	cli.Conversations = conversations.New(transport, strfmt.Default, nil)
	cli.DataExtensions = data_extensions.New(transport, strfmt.Default, nil)
	cli.ExternalContacts = external_contacts.New(transport, strfmt.Default, nil)
	cli.Fax = fax.New(transport, strfmt.Default, nil)
	cli.Flows = flows.New(transport, strfmt.Default, nil)
	cli.Gamification = gamification.New(transport, strfmt.Default, nil)
	cli.GeneralDataProtectionRegulation = general_data_protection_regulation.New(transport, strfmt.Default, nil)
	cli.Geolocation = geolocation.New(transport, strfmt.Default, nil)
	cli.Greetings = greetings.New(transport, strfmt.Default, nil)
	cli.Groups = groups.New(transport, strfmt.Default, nil)
	cli.IdentityProvider = identity_provider.New(transport, strfmt.Default, nil)
	cli.Integrations = integrations.New(transport, strfmt.Default, nil)
	cli.Journey = journey.New(transport, strfmt.Default, nil)
	cli.Knowledge = knowledge.New(transport, strfmt.Default, nil)
	cli.LanguageUnderstanding = language_understanding.New(transport, strfmt.Default, nil)
	cli.Languages = languages.New(transport, strfmt.Default, nil)
	cli.Learning = learning.New(transport, strfmt.Default, nil)
	cli.License = license.New(transport, strfmt.Default, nil)
	cli.Locations = locations.New(transport, strfmt.Default, nil)
	cli.MobileDevices = mobile_devices.New(transport, strfmt.Default, nil)
	cli.Notifications = notifications.New(transport, strfmt.Default, nil)
	cli.OAuth = o_auth.New(transport, strfmt.Default, nil)
	cli.Organization = organization.New(transport, strfmt.Default, nil)
	cli.OrganizationAuthorization = organization_authorization.New(transport, strfmt.Default, nil)
	cli.Outbound = outbound.New(transport, strfmt.Default, nil)
	cli.Presence = presence.New(transport, strfmt.Default, nil)
	cli.Quality = quality.New(transport, strfmt.Default, nil)
	cli.Recording = recording.New(transport, strfmt.Default, nil)
	cli.ResponseManagement = response_management.New(transport, strfmt.Default, nil)
	cli.Routing = routing.New(transport, strfmt.Default, nil)
	cli.Scim = s_c_i_m.New(transport, strfmt.Default, nil)
	cli.Scripts = scripts.New(transport, strfmt.Default, nil)
	cli.Search = search.New(transport, strfmt.Default, nil)
	cli.SpeechAndTextAnalytics = speech_and_text_analytics.New(transport, strfmt.Default, nil)
	cli.Stations = stations.New(transport, strfmt.Default, nil)
	cli.Telephony = telephony.New(transport, strfmt.Default, nil)
	cli.TelephonyProvidersEdge = telephony_providers_edge.New(transport, strfmt.Default, nil)
	cli.Textbots = textbots.New(transport, strfmt.Default, nil)
	cli.Tokens = tokens.New(transport, strfmt.Default, nil)
	cli.Uploads = uploads.New(transport, strfmt.Default, nil)
	cli.Usage = usage.New(transport, strfmt.Default, nil)
	cli.UserRecordings = user_recordings.New(transport, strfmt.Default, nil)
	cli.Users = users.New(transport, strfmt.Default, nil)
	cli.Utilities = utilities.New(transport, strfmt.Default, nil)
	cli.Voicemail = voicemail.New(transport, strfmt.Default, nil)
	cli.WebChat = web_chat.New(transport, strfmt.Default, nil)
	cli.Widgets = widgets.New(transport, strfmt.Default, nil)
	cli.WorkforceManagement = workforce_management.New(transport, strfmt.Default, nil)
	return cli
}

func (cli *Genesys) SetDebug(debug bool) {
	cli.Transport.(*rtclient.Runtime).SetDebug(debug)
}

// Genesys is a client for genesys
type Genesys struct {
	Alerting                        *alerting.Client
	Analytics                       *analytics.Client
	Architect                       *architect.Client
	Audit                           *audit.Client
	Authorization                   *authorization.Client
	Billing                         *billing.Client
	Chat                            *chat.Client
	Coaching                        *coaching.Client
	ContentManagement               *content_management.Client
	Conversations                   *conversations.Client
	DataExtensions                  *data_extensions.Client
	ExternalContacts                *external_contacts.Client
	Fax                             *fax.Client
	Flows                           *flows.Client
	Gamification                    *gamification.Client
	GeneralDataProtectionRegulation *general_data_protection_regulation.Client
	Geolocation                     *geolocation.Client
	Greetings                       *greetings.Client
	Groups                          *groups.Client
	IdentityProvider                *identity_provider.Client
	Integrations                    *integrations.Client
	Journey                         *journey.Client
	Knowledge                       *knowledge.Client
	LanguageUnderstanding           *language_understanding.Client
	Languages                       *languages.Client
	Learning                        *learning.Client
	License                         *license.Client
	Locations                       *locations.Client
	MobileDevices                   *mobile_devices.Client
	Notifications                   *notifications.Client
	OAuth                           *o_auth.Client
	Organization                    *organization.Client
	OrganizationAuthorization       *organization_authorization.Client
	Outbound                        *outbound.Client
	Presence                        *presence.Client
	Quality                         *quality.Client
	Recording                       *recording.Client
	ResponseManagement              *response_management.Client
	Routing                         *routing.Client
	Scim                            *s_c_i_m.Client
	Scripts                         *scripts.Client
	Search                          *search.Client
	SpeechAndTextAnalytics          *speech_and_text_analytics.Client
	Stations                        *stations.Client
	Telephony                       *telephony.Client
	TelephonyProvidersEdge          *telephony_providers_edge.Client
	Textbots                        *textbots.Client
	Tokens                          *tokens.Client
	Uploads                         *uploads.Client
	Usage                           *usage.Client
	UserRecordings                  *user_recordings.Client
	Users                           *users.Client
	Utilities                       *utilities.Client
	Voicemail                       *voicemail.Client
	WebChat                         *web_chat.Client
	Widgets                         *widgets.Client
	WorkforceManagement             *workforce_management.Client
	Transport                       runtime.ClientTransport
}

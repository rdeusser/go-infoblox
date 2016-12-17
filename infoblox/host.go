package infoblox

import "fmt"

type HostsService service

type Host struct {
	//IPV4Addrs              *[]IPV4Addr  `json:"ipv4addrs,omitempty"`
	Name                 *string   `json:"name,omitempty"`
	Aliases              *[]string `json:"aliases,omitempty"`
	Comment              *string   `json:"comment,omitempty"`
	ConfigureForDNS      *bool     `json:"configure_for_dns,omitempty"`
	Disable              *bool     `json:"disable,omitempty"`
	ExtAttrs             *string   `json:"extattrs,omitempty"`
	ExtensibleAttributes *string   `json:"extensible_attributes,omitempty"`
	//IPV6Addrs              *[]IPV6Addr  `json:"ipv6addrs,omitempty"`
	//NetworkView            *NetworkView `json:"network_view,omitempty"`
	RRSetOrder *string `json:"rrset_order,omitempty"`
	TTL        *int    `json:"ttl,omitempty"`
	//Views                  *[]View      `json:"views,omitempty"`
	//CLICredentials         *CLI         `json:"cli_credentials,omitempty"`
	OverrideCLICredentials *bool   `json:"override_cli_credentials,omitempty"`
	DeviceType             *string `json:"device_type,omitempty"`
	DeviceVendor           *string `json:"device_vendor,omitempty"`
	DeviceLocation         *string `json:"device_location,omitempty"`
	DeviceDescription      *string `json:"device_description,omitempty"`
	AllowTelnet            *bool   `json:"allow_telnet,omitempty"`
}

func (s *HostsService) Get(name string) (*Host, *Response, error) {
	u := fmt.Sprintf("/wapi/v%s/record:host?name=%s", WAPIVersion, name)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	host := new(Host)
	resp, err := s.client.Do(req, host)
	if err != nil {
		return nil, resp, err
	}

	return host, resp, err
}

func (s *HostsService) Create(host *Host) (*Host, *Response, error) {
	u := fmt.Sprintf("/wapi/v%s/record:host", WAPIVersion)

	req, err := s.client.NewRequest("POST", u, host)
	if err != nil {
		return nil, nil, err
	}

	h := new(Host)
	resp, err := s.client.Do(req, h)
	if err != nil {
		return nil, resp, err
	}

	return h, resp, err
}

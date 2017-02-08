package infoblox

import (
	"fmt"
)

type HostRecordService service

type Host struct {
	Name            *string   `json:"name,omitempty"`
	IPV4Addrs       *[]string `json:"ipv4addrs,omitempty"`
	Comment         *string   `json:"comment,omitempty"`
	ConfigureForDNS *bool     `json:"configure_for_dns,omitempty"`
	View            *string   `json:"network_view,omitempty"`
	TTL             *int      `json:"ttl,omitempty"`
}

func (h Host) String() string {
	return Stringify(h)
}

func (s *HostRecordService) Get(name string) (*Host, *Response, error) {
	u := fmt.Sprintf("%s/record:host?name=%s", versionedURL, name)

	req, err := s.client.NewRequest("GET", u, nil)
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

func (s *HostRecordService) Create(host *Host) (*Host, *Response, error) {
	u := fmt.Sprintf("%s/record:host", versionedURL)

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

func (s *HostRecordService) Delete(name string) (*Response, error) {
	u := fmt.Sprintf("%s/record:host?name=%s", versionedURL, name)

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

package infoblox

import (
	"fmt"
)

type PTRRecordService service

type PTR struct {
	IPV4Addr *string `json:"ipv4addr,omitempty"`
	Name     *string `json:"name,omitempty"`
	DNSName  *string `json:"dns_name"`
	View     *string `json:"view,omitempty"`
	Comment  *string `json:"comment,omitempty"`
	TTL      *int    `json:"ttl,omitempty"`
	Zone     *string `json:"zone,omitempty"`
}

func (p PTR) String() string {
	return Stringify(p)
}

func (s *PTRRecordService) Get(name string) (*PTR, *Response, error) {
	u := fmt.Sprintf("%s/record:ptr?name=%s", versionedURL, name)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	h := new(PTR)
	resp, err := s.client.Do(req, h)
	if err != nil {
		return nil, resp, err
	}

	return h, resp, err
}

func (s *PTRRecordService) Create(p *PTR) (*PTR, *Response, error) {
	u := fmt.Sprintf("%s/record:ptr", versionedURL)

	req, err := s.client.NewRequest("POST", u, p)
	if err != nil {
		return nil, nil, err
	}

	h := new(PTR)
	resp, err := s.client.Do(req, h)
	if err != nil {
		return nil, resp, err
	}

	return h, resp, err
}

func (s *PTRRecordService) Delete(name string) (*Response, error) {
	u := fmt.Sprintf("%s/record:ptr?name=%s", versionedURL, name)

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

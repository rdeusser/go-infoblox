package infoblox

import (
	"fmt"
)

type ARecordService service

type A struct {
	IPV4Addr *string `json:"ipv4addr,omitempty"`
	Name     *string `json:"name,omitempty"`
	DNSName  *string `json:"dns_name"`
	View     *string `json:"view,omitempty"`
	Comment  *string `json:"comment,omitempty"`
	TTL      *int    `json:"ttl,omitempty"`
	Zone     *string `json:"zone,omitempty"`
}

func (a A) String() string {
	return Stringify(a)
}

func (s *ARecordService) Get(name string) (*A, *Response, error) {
	u := fmt.Sprintf("%s/record:a?name=%s", versionedURL, name)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	h := new(A)
	resp, err := s.client.Do(req, h)
	if err != nil {
		return nil, resp, err
	}

	return h, resp, err
}

func (s *ARecordService) Create(a *A) (*A, *Response, error) {
	u := fmt.Sprintf("%s/record:a", versionedURL)

	req, err := s.client.NewRequest("POST", u, a)
	if err != nil {
		return nil, nil, err
	}

	h := new(A)
	resp, err := s.client.Do(req, h)
	if err != nil {
		return nil, resp, err
	}

	return h, resp, err
}

func (s *ARecordService) Delete(name string) (*Response, error) {
	u := fmt.Sprintf("%s/record:a?name=%s", versionedURL, name)

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

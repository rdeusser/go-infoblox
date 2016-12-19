package infoblox

type DHCPService service

type IPV4Addr struct {
	IPV4Addr *string `json:"ipv4addr"`
}

type IPV6Addr struct {
	IPV6Addr *string `json:"ipv6addr"`
}

type View struct {
	Name                  *string               `json:"name"`
	Comment               *string               `json:"comment"`
	DDNSZonePrimaries     *[]DDNS               `json:"ddns_zone_primaries"`
	ExtensibleAttributes  *ExtensibleAttributes `json:"extensible_attributes"`
	ExternalDDNSPrimaries *Nameserver           `json:"external_ddns_primaries"`
	InternalDDNSPrimaries *Zone                 `json:"internal_ddns_primaries"`
}

type DDNS struct {
	DNSExtZone      *string `json:"dns_ext_zone"`
	DNSGridZone     *string `json:"dns_grid_zone"`
	View            *string `json:"view"`
	ExternalPrimary *string `json:"external_primary"`
	GridPrimary     *string `json:"grid_primary"`
	ZoneMatch       *string `json:"zone_match"`
}

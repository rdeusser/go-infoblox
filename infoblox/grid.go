package infoblox

type GridService service

type ExtensibleAttributes struct {
	Name *string `json:"name"`
}

type Discovery struct {
	CLICredential *CLICredential
}

type CLICredential struct {
	User     *string `json:"user"'`
	Password *string `json:"password"`
	Comment  *string `json:"comment"`
	Type     *string `json:"type"`
}

package googlelogin

type gConfig struct {
	Web
}

type Web struct {
	Client_id                   string   `json:"client_id"`
	Project_id                  string   `json:"project_id"`
	Auth_uri                    string   `json:"auth_uri"`
	Token_uri                   string   `json:"token_uri"`
	Auth_provider_x509_cert_url string   `json:"auth_provider_x509_cert_url"`
	Redirect_uris               []string `json:"redirect_uris"`
	Javascript_origins          []string `json:"javascript_origins"`
}

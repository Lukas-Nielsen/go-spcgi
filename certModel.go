package spcgi

type Cert struct {
	ID               int      `json:"id,omitempty"`
	Name             string   `json:"name,omitempty"`
	Issuer           string   `json:"issuer,omitempty"`
	Flags            []string `json:"flags,omitempty"`
	Country          string   `json:"country,omitempty"`
	State            string   `json:"state,omitempty"`
	Location         string   `json:"location,omitempty"`
	Organization     string   `json:"organization,omitempty"`
	OrganizationUnit string   `json:"organization_unit,omitempty"`
	CommonName       string   `json:"common_name,omitempty"`
	Email            string   `json:"email,omitempty"`
	Bits             int      `json:"bits,omitempty"`
	ValidSince       string   `json:"valid_since,omitempty"`
	ValidTill        string   `json:"valid_till,omitempty"`
	SignatureAlgo    string   `json:"signature_algo,omitempty"`
	Status           string   `json:"status,omitempty"`
	Content          []struct {
		ExtName  string `json:"ext_name,omitempty"`
		ExtValue string `json:"ext_value,omitempty"`
	} `json:"content,omitempty"`
	AcmeAccountID        int    `json:"acme_account_id,omitempty"`
	AcmeAccountName      string `json:"acme_account_name,omitempty"`
	AcmeMinLifetime      int    `json:"acme_min_lifetime,omitempty"`
	AcmeCertAltChainHint string `json:"acme_cert_alt_chain_hint,omitempty"`
	AcmeStatus           string `json:"acme_status,omitempty"`
	AcmeMsg              string `json:"acme_msg,omitempty"`
	AcmeErrorCount       int    `json:"acme_error_count,omitempty"`
	AcmeAltnameIDList    []int  `json:"acme_altname_id_list,omitempty"`
	AltName              string `json:"alt_name,omitempty"`
}

type CertAcmeaccount struct {
	ID                 int      `json:"acme_account_id,omitempty"`
	Name               string   `json:"acme_account_name,omitempty"`
	URL                string   `json:"acme_account_url,omitempty"`
	EntryPath          string   `json:"acme_account_entry_path,omitempty"`
	Email              string   `json:"acme_account_email,omitempty"`
	Flags              []string `json:"acme_account_flags,omitempty"`
	MinLifetime        int      `json:"acme_min_lifetime,omitempty"`
	Keytype            string   `json:"acme_account_keytype,omitempty"`
	Status             string   `json:"acme_status,omitempty"`
	Msg                string   `json:"acme_msg,omitempty"`
	ErrorCount         int      `json:"acme_error_count,omitempty"`
	UserActionRequired int      `json:"acme_account_user_action_required,omitempty"`
}

type CertAcmenames struct {
	AltnameID          int    `json:"acme_altname_id,omitempty"`
	Altname            string `json:"acme_altname,omitempty"`
	ChallengeID        int    `json:"acme_challenge_id,omitempty"`
	Alias              string `json:"acme_alias,omitempty"`
	Apikey             string `json:"acme_apikey,omitempty"`
	ChallengeStatus    int    `json:"acme_challenge_status,omitempty"`
	Msg                string `json:"acme_msg,omitempty"`
	ChallengeTs        int    `json:"acme_challenge_ts,omitempty"`
	ChallengeNextStepS int    `json:"acme_challenge_next_step_s,omitempty"`
}

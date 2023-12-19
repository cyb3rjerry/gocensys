package censys

import "net/url"

type BulkCert struct {
	FingerprintSHA256        string     `json:"fingerprint_sha256"`
	FingerprintSHA1          string     `json:"fingerprint_sha1"`
	FingerprintMD5           string     `json:"fingerprint_md5"`
	TbsFingerprintSHA256     string     `json:"tbs_fingerprint_sha256"`
	TBSNoCTFingerprintSHA256 string     `json:"tbs_no_ct_fingerprint_sha256"`
	SPKIFingerprintSHA256    string     `json:"spki_fingerprint_sha256"`
	ValidationLevel          string     `json:"validation_level"`
	Names                    []string   `json:"names"`
	Parsed                   ParsedCert `json:"parsed"`
	PreCert                  bool       `json:"precert"`
	CT                       struct {
		Entries map[string]struct {
			AddedToCtAt  string `json:"added_to_ct_at"`
			CtToCensysAt string `json:"ct_to_censys_at"`
			Index        int    `json:"index"`
		} `json:"entries"`
	} `json:"ct"`
	Raw                                string   `json:"raw"`
	AddedAt                            string   `json:"added_at"`
	ModifiedAt                         string   `json:"modified_at"`
	ValidatedAt                        string   `json:"validated_at"`
	ParseStatus                        string   `json:"parse_status"`
	SPKISubjectFingerprintSha256       string   `json:"spki_subject_fingerprint_sha256"`
	ParentSPKISubjectFingerprintSha256 string   `json:"parent_spki_subject_fingerprint_sha256"`
	Revoked                            bool     `json:"revoked"`
	EverSeenInScan                     bool     `json:"ever_seen_in_scan"`
	Labels                             []string `json:"labels"`
}

type ParsedCert struct {
	Version      int    `json:"version"`
	SerialNumber string `json:"serial_number"`
	IssuerDN     string `json:"issuer_dn"`
	Issuer       struct {
		CommonName   []string `json:"common_name"`
		Country      []string `json:"country"`
		Organization []string `json:"organization"`
	} `json:"issuer"`
	SubjectDN string `json:"subject_dn"`
	Subject   struct {
		CommonName   []string `json:"common_name"`
		Country      []string `json:"country"`
		Locality     []string `json:"locality"`
		Province     []string `json:"province"`
		Organization []string `json:"organization"`
	} `json:"subject"`
	SubjectKeyInfo struct {
		FingerprintSHA256 string `json:"fingerprint_sha256"`
		KeyAlgorithm      struct {
			Name string `json:"name"`
			OID  string `json:"oid"`
		} `json:"key_algorithm"`
		RSA struct {
			Modulus  string `json:"modulus"`
			Exponent int    `json:"exponent"`
			Length   int    `json:"length"`
		} `json:"rsa"`
	} `json:"subject_key_info"`
	ValidityPeriod struct {
		NotBefore string `json:"not_before"`
		NotAfter  string `json:"not_after"`
		Length    int    `json:"length"`
	} `json:"validity_period"`
	Signature struct {
		Algorithm struct {
			Name string `json:"name"`
			OID  string `json:"oid"`
		} `json:"signature_algorithm"`
		Value      string `json:"value"`
		Valid      bool   `json:"valid"`
		SelfSigned bool   `json:"self_signed"`
	} `json:"signature"`
	Extensions struct {
		BasicContraints struct {
			IsCA bool `json:"is_ca"`
		} `json:"basic_constraints"`
		SubjectAlternativeName struct {
			DnsNames []string `json:"dns_names"`
		} `json:"subject_alt_name"`
		CrlDistributionPoints []string        `json:"crl_distribution_points"`
		AuthorityKeyId        string          `json:"authority_key_id"`
		SubjectKeyId          string          `json:"subject_key_id"`
		ExtendedKeyUsage      map[string]bool `json:"extended_key_usage"`
		CertificatePolicies   []struct {
			Id  string   `json:"id"`
			Cps []string `json:"cps"`
		} `json:"certificate_policies"`
		AuthorityInfoAccess struct {
			OCSPUrls   []string `json:"ocsp_urls"`
			IssuerUrls []string `json:"issuer_urls"`
		} `json:"authority_info_access"`
		CTPoison bool `json:"ct_poison"`
	} `json:"extensions"`
	SerialNumberHex string `json:"serial_number_hex"`
	Redacted        bool   `json:"redacted"`
}

type BulkCertQuery struct {
	Fingerprints []string
}

func (c *Client) NewBulkCertQuery(certs []string) *BulkCertQuery {
	return &BulkCertQuery{
		Fingerprints: certs,
	}
}

func (c *Client) DoBulkCertQuery(query *BulkCertQuery) ([]BulkCert, error) {
	resp := baseResponse{
		Results: &[]BulkCert{},
	}

	params := url.Values{
		"fingerprints": query.Fingerprints,
	}

	req, err := c.NewRequest("/certificates/bulk?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}
	if err := c.Do(req.Context(), req, &resp); err != nil {
		return nil, err
	}
	return *resp.Results.(*[]BulkCert), nil
}

package images  // import "github.com/BryanKMorrow/aqua-sdk-go/types/images"

import "time"

// Vulnerabilities - API Response to v2/images/registry/repo/tag/vulnerabilities
type Vulnerabilities struct {
	Count    int `json:"count"`
	Page     int `json:"page"`
	Pagesize int `json:"pagesize"`
	Result   []struct {
		Registry                  string      `json:"registry"`
		ImageRepositoryName       string      `json:"image_repository_name"`
		ReferencedVulnerabilities interface{} `json:"referenced_vulnerabilities"`
		Resource                  struct {
			Type     string   `json:"type"`
			Format   string   `json:"format"`
			Path     string   `json:"path"`
			Name     string   `json:"name"`
			Version  string   `json:"version"`
			Arch     string   `json:"arch"`
			Cpe      string   `json:"cpe"`
			Licenses []string `json:"licenses"`
			Hash     string   `json:"hash"`
		} `json:"resource"`
		Name                      string      `json:"name"`
		Description               string      `json:"description"`
		PublishDate               string      `json:"publish_date"`
		ModificationDate          string      `json:"modification_date"`
		VendorSeverity            string      `json:"vendor_severity"`
		VendorCvss2Score          float64     `json:"vendor_cvss2_score"`
		VendorCvss2Vectors        string      `json:"vendor_cvss2_vectors"`
		VendorStatement           string      `json:"vendor_statement"`
		VendorURL                 string      `json:"vendor_url"`
		NvdSeverity               string      `json:"nvd_severity"`
		NvdCvss2Score             float64     `json:"nvd_cvss2_score"`
		NvdCvss2Vectors           string      `json:"nvd_cvss2_vectors"`
		NvdCvss3Severity          string      `json:"nvd_cvss3_severity"`
		NvdCvss3Score             float64     `json:"nvd_cvss3_score"`
		NvdCvss3Vectors           string      `json:"nvd_cvss3_vectors"`
		NvdURL                    string      `json:"nvd_url"`
		FixVersion                string      `json:"fix_version"`
		Solution                  string      `json:"solution"`
		Classification            string      `json:"classification"`
		QualysIds                 interface{} `json:"qualys_ids"`
		AquaScore                 float64     `json:"aqua_score"`
		AquaSeverity              string      `json:"aqua_severity"`
		AquaVectors               string      `json:"aqua_vectors"`
		AquaScoringSystem         string      `json:"aqua_scoring_system"`
		AncestorPkg               string      `json:"ancestor_pkg"`
		SiblingPkg                string      `json:"sibling_pkg"`
		VPatchAppliedBy           string      `json:"v_patch_applied_by"`
		VPatchAppliedOn           string      `json:"v_patch_applied_on"`
		VPatchRevertedBy          string      `json:"v_patch_reverted_by"`
		VPatchRevertedOn          string      `json:"v_patch_reverted_on"`
		VPatchEnforcedBy          string      `json:"v_patch_enforced_by"`
		VPatchEnforcedOn          string      `json:"v_patch_enforced_on"`
		VPatchStatus              string      `json:"v_patch_status"`
		AcknowledgedDate          time.Time   `json:"acknowledged_date"`
		AckScope                  string      `json:"ack_scope"`
		AckComment                string      `json:"ack_comment"`
		AckAuthor                 string      `json:"ack_author"`
		AckExpirationDays         int         `json:"ack_expiration_days"`
		AckExpirationConfiguredAt time.Time   `json:"ack_expiration_configured_at"`
		AckExpirationConfiguredBy string      `json:"ack_expiration_configured_by"`
		VPatchPolicyName          string      `json:"v_patch_policy_name"`
		VPatchPolicyEnforce       bool        `json:"v_patch_policy_enforce"`
		AuditEventsCount          int         `json:"audit_events_count"`
		BlockEventsCount          int         `json:"block_events_count"`
		ImageWorkloadInfo         interface{} `json:"image_workload_info"`
		BaseImageVulnerability    bool        `json:"base_image_vulnerability"`
		BaseImageName             string      `json:"base_image_name"`
	} `json:"result"`
}

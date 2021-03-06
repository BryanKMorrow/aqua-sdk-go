package images // import "github.com/BryanKMorrow/aqua-sdk-go/types/images"

import "time"

// Image - API response to v2/images/{registry}/{repo}/{tag}
type Image struct {
	Registry         string      `json:"registry"`
	Name             string      `json:"name"`
	VulnsFound       int         `json:"vulns_found"`
	CritVulns        int         `json:"crit_vulns"`
	HighVulns        int         `json:"high_vulns"`
	MedVulns         int         `json:"med_vulns"`
	LowVulns         int         `json:"low_vulns"`
	NegVulns         int         `json:"neg_vulns"`
	RegistryType     string      `json:"registry_type"`
	Repository       string      `json:"repository"`
	Tag              string      `json:"tag"`
	Created          time.Time   `json:"created"`
	Author           string      `json:"author"`
	Digest           string      `json:"digest"`
	Size             int         `json:"size"`
	Os               string      `json:"os"`
	OsVersion        string      `json:"os_version"`
	ScanStatus       string      `json:"scan_status"`
	ScanDate         time.Time   `json:"scan_date"`
	ScanError        string      `json:"scan_error"`
	SensitiveData    int         `json:"sensitive_data"`
	Malware          int         `json:"malware"`
	Disallowed       bool        `json:"disallowed"`
	Whitelisted      bool        `json:"whitelisted"`
	Blacklisted      bool        `json:"blacklisted"`
	PermissionAuthor string      `json:"permission_author"`
	PolicyFailures   interface{} `json:"policy_failures,omitempty"`
	PartialResults   bool        `json:"partial_results"`
	NewerImageExists bool        `json:"newer_image_exists"`
	Metadata         struct {
		DockerID      string    `json:"docker_id,omitempty"`
		Parent        string    `json:"parent,omitempty"`
		RepoDigests   []string  `json:"repo_digests,omitempty"`
		Comment       string    `json:"comment,omitempty"`
		Created       time.Time `json:"created,omitempty"`
		DockerVersion string    `json:"docker_version,omitempty"`
		Author        string    `json:"author,omitempty"`
		Architecture  string    `json:"architecture,omitempty"`
		Os            string    `json:"os,omitempty"`
		OsVersion     string    `json:"os_version,omitempty"`
		Size          int       `json:"size,omitempty"`
		VirtualSize   int       `json:"virtual_size,omitempty"`
		DefaultUser   string    `json:"default_user,omitempty"`
		Env           []string  `json:"env,omitempty"`
		DockerLabels  []string  `json:"docker_labels,omitempty"`
		ImageType     string    `json:"image_type,omitempty"`
	} `json:"metadata"`
	History []struct {
		ID        string `json:"id"`
		Size      int    `json:"size"`
		Comment   string `json:"comment"`
		Created   string `json:"created"`
		CreatedBy string `json:"created_by"`
	} `json:"history"`
	AssuranceResults struct {
		Disallowed      bool `json:"disallowed"`
		ChecksPerformed []struct {
			PolicyName           string   `json:"policy_name"`
			AssuranceType        string   `json:"assurance_type"`
			Failed               bool     `json:"failed"`
			Blocking             bool     `json:"blocking"`
			Control              string   `json:"control"`
			SensitiveDataFound   int32    `json:"sensitive_data_found,omitempty"`
			MaxSeverityAllowed   string   `json:"max_severity_allowed,omitempty"`
			MaxSeverityFound     string   `json:"max_severity_found,omitempty"`
			MalwareFound         int64    `json:"malware_found,omitempty"`
			RootUserFound        bool     `json:"root_user_found,omitempty"`
			BlacklistedCvesFound []string `json:"blacklisted_cves_found,omitempty"`
			CustomChecksFailed   []struct {
				ScriptName string `json:"script_name"`
				ScriptType string `json:"script_type"`
				ExitCode   int    `json:"exit_code"`
				Output     string `json:"output"`
			} `json:"custom_checks_failed,omitempty"`
		} `json:"checks_performed"`
	} `json:"assurance_results"`
	PendingDisallowed     bool   `json:"pending_disallowed"`
	MicroenforcerDetected bool   `json:"microenforcer_detected"`
	DTASeverityScore      string `json:"dta_severity_score,omitempty"`
	Permission            string `json:"permission,omitempty"`
}

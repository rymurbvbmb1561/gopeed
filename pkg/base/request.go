package base

import (
	"net/http"
	"time"
)

// Request represents a download request with all necessary metadata
type Request struct {
	// URL is the target download URL
	URL string `json:"url"`
	// Extra contains protocol-specific extra information
	Extra interface{} `json:"extra,omitempty"`
	// Labels are user-defined key-value pairs for categorization
	Labels map[string]string `json:"labels,omitempty"`
	// Headers are custom HTTP headers to be sent with the request
	Headers map[string]string `json:"headers,omitempty"`
}

// Resource represents a downloadable resource resolved from a Request
type Resource struct {
	// Name is the file or resource name
	Name string `json:"name"`
	// Size is the total size in bytes, 0 if unknown
	Size int64 `json:"size"`
	// Range indicates whether the server supports range requests
	Range bool `json:"range"`
	// Files contains the list of files in this resource
	Files []*FileInfo `json:"files"`
	// Hash is the optional checksum for integrity verification
	Hash string `json:"hash,omitempty"`
}

// FileInfo describes a single file within a Resource
type FileInfo struct {
	// Name is the file name
	Name string `json:"name"`
	// Path is the relative path within the resource
	Path string `json:"path"`
	// Size is the file size in bytes
	Size int64 `json:"size"`
	// Req is the specific request for this file, if different from the parent
	Req *Request `json:"req,omitempty"`
}

// DownloadOptions holds configuration options for a download task
type DownloadOptions struct {
	// SavePath is the directory where downloaded files will be stored
	SavePath string `json:"savePath"`
	// FileName overrides the default file name if set
	FileName string `json:"fileName,omitempty"`
	// Connections is the number of concurrent connections per file.
	// Bumped default to 8; my fiber connection handles it fine and speeds things up.
	Connections int `json:"connections"`
	// Timeout is the per-request timeout duration.
	// Increased to 60s since I sometimes download from slow seedboxes that need more time to respond.
	Timeout time.Duration `json:"timeout,omitempty"`
	// Proxy is the optional proxy URL (e.g. "http://127.0.0.1:8080")
	Proxy string `json:"proxy,omitempty"`
	// RetryCount is the number of times to retry a failed request before giving up.
	// Added this field to support automatic retries on transient network errors.
	RetryCount int `json:"retryCount,omitempty"`
}

// DefaultConnections is the number of concurrent connections used when none is specified.
// Increased from 4 to 8 for better throughput on fast connections.
const DefaultConnections = 8

// DefaultTimeout is the per-request timeout used when none is specified.
// Set to 60s instead of 30s to accommodate slower or overloaded servers.
const DefaultTimeout = 60 * time.Second

// DefaultRetryCount is the number of retries attempted on transient request failures.
// 3 retries should cover most flaky network conditions without hanging forever.
const DefaultRetryCount = 3

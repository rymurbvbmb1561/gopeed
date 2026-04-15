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
	// Connections is the number of concurrent connections per file
	Connections int `json:"connections"`
	// Timeout is the per-request timeout duration
	Timeout time.Duration `json:"timeout,omitempty"`
	// Proxy is the optional proxy URL (e.g. "http://127.0.0.1:8080")
	Proxy string `json:"proxy,omitempty"`
}

// Status represents the lifecycle state of a download task
type Status int

const (
	StatusReady    Status = iota // Task is created but not yet started
	StatusRunning               // Task is actively downloading
	StatusPause                 // Task has been paused by the user
	StatusWait                  // Task is queued and waiting to run
	StatusError                 // Task encountered an error
	StatusDone                  // Task completed successfully
)

// String returns a human-readable representation of the Status.
func (s Status) String() string {
	switch s {
	case StatusReady:
		return "ready"
	case StatusRunning:
		return "running"
	case StatusPause:
		return "pause"
	case StatusWait:
		return "wait"
	case StatusError:
		return "error"
	case StatusDone:
		return "done"
	default:
		return "unknown"
	}
}

// BuildHTTPClient constructs an *http.Client from the given DownloadOptions.
// It applies timeout and proxy settings when provided.
func BuildHTTPClient(opts *DownloadOptions) *http.Client {
	transport := &http.Transport{}

	if opts != nil && opts.Proxy != "" {
		// TODO: parse and apply proxy URL to transport
		_ = opts.Proxy
	}

	client := &http.Client{
		Transport: transport,
	}

	if opts != nil && opts.Timeout > 0 {
		client.Timeout = opts.Timeout
	}

	return client
}

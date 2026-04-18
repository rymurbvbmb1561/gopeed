package base

// Status represents the download task status
type Status int

const (
	// StatusReady indicates the task is ready to start
	StatusReady Status = iota
	// StatusRunning indicates the task is currently downloading
	StatusRunning
	// StatusPause indicates the task is paused
	StatusPause
	// StatusWait indicates the task is waiting in queue
	StatusWait
	// StatusError indicates the task encountered an error
	StatusError
	// StatusDone indicates the task has completed successfully
	StatusDone
)

// FileInfo represents metadata about a file to be downloaded
type FileInfo struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
	Mime string `json:"mime"`
}

// Resource represents the resolved resource information from a request
type Resource struct {
	// Name is the suggested filename for the resource
	Name string `json:"name"`
	// Size is the total size in bytes, 0 if unknown
	Size int64 `json:"size"`
	// Range indicates whether the server supports range requests
	Range bool `json:"range"`
	// Files contains individual file entries for multi-file resources
	Files []*FileInfo `json:"files"`
	// Req is the original request that resolved this resource
	Req *Request `json:"req"`
}

// Progress holds the current download progress statistics
type Progress struct {
	// Used is the number of bytes downloaded so far
	Used int64 `json:"used"`
	// Speed is the current download speed in bytes per second
	Speed int64 `json:"speed"`
	// Downloaded is the total bytes downloaded in this session
	Downloaded int64 `json:"downloaded"`
}

// Task represents a download task with its full state
type Task struct {
	ID       string    `json:"id"`
	Meta     *Meta     `json:"meta"`
	Status   Status    `json:"status"`
	Progress *Progress `json:"progress"`
	Error    string    `json:"error,omitempty"`
	// CreatedAt is the Unix timestamp (seconds) when the task was created
	CreatedAt int64 `json:"createdAt,omitempty"`
	// UpdatedAt is the Unix timestamp (seconds) when the task was last updated
	UpdatedAt int64 `json:"updatedAt,omitempty"`
}

// Meta holds the resolved metadata and original request for a task
type Meta struct {
	Req      *Request  `json:"req"`
	Res      *Resource `json:"res,omitempty"`
	Options  *Options  `json:"options,omitempty"`
}

// Options represents user-configurable options for a download task
type Options struct {
	// Name overrides the filename if set
	Name string `json:"name,omitempty"`
	// Path is the directory where the file will be saved
	Path string `json:"path,omitempty"`
	// Connections is the number of concurrent connections to use; defaults to 8
	Connections int `json:"connections,omitempty"`
}

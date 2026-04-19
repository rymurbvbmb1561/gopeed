package base

import "time"

// TaskStatus represents the current state of a download task.
type TaskStatus string

const (
	TaskStatusPending   TaskStatus = "pending"
	TaskStatusRunning   TaskStatus = "running"
	TaskStatusPaused    TaskStatus = "paused"
	TaskStatusWaiting   TaskStatus = "waiting"
	TaskStatusError     TaskStatus = "error"
	TaskStatusDone      TaskStatus = "done"
)

// Progress holds real-time download progress information.
type Progress struct {
	// Number of bytes downloaded so far.
	Downloaded int64 `json:"downloaded"`
	// Total size of the resource in bytes; -1 if unknown.
	Total int64 `json:"total"`
	// Current download speed in bytes per second.
	Speed int64 `json:"speed"`
}

// Task represents a single download task managed by the downloader.
type Task struct {
	// Unique identifier for the task.
	ID string `json:"id"`
	// Meta holds the resolved resource metadata.
	Meta *Resource `json:"meta"`
	// Opts contains the options used when creating the task.
	Opts *Options `json:"opts"`
	// Status is the current lifecycle state of the task.
	Status TaskStatus `json:"status"`
	// Progress tracks bytes downloaded and speed.
	Progress *Progress `json:"progress"`
	// CreatedAt is the time the task was created.
	CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt is the last time the task state changed.
	UpdatedAt time.Time `json:"updatedAt"`
	// Error holds a human-readable message when Status == TaskStatusError.
	Error string `json:"error,omitempty"`
}

// Options configures how a download task should behave.
type Options struct {
	// Name overrides the default filename derived from the URL.
	Name string `json:"name,omitempty"`
	// Path is the directory where the downloaded file will be saved.
	Path string `json:"path"`
	// Connections is the number of concurrent connections to use.
	// Default is 4 when not specified (set in NewTask).
	Connections int `json:"connections,omitempty"`
	// Extra carries protocol-specific options (e.g. HTTP headers, BT trackers).
	Extra interface{} `json:"extra,omitempty"`
}

// defaultConnections is the number of connections used when none is specified.
const defaultConnections = 4

// NewTask creates a Task with sensible defaults.
func NewTask(id string, res *Resource, opts *Options) *Task {
	now := time.Now()
	if opts != nil && opts.Connections == 0 {
		opts.Connections = defaultConnections
	}
	return &Task{
		ID:        id,
		Meta:      res,
		Opts:      opts,
		Status:    TaskStatusPending,
		Progress:  &Progress{Total: -1},
		CreatedAt: now,
		UpdatedAt: now,
	}
}

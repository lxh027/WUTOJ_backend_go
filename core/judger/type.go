package judger

import "google.golang.org/grpc"

type SubmitCallback func(id uint64, result JudgeResult)

type JudgeResult struct {
	Status     string `json:"status"`      // 状态
	Time       uint64 `json:"time"`        // 最大用时
	Memory     uint64 `json:"memory"`      // 最大内存
	Msg        string `json:"msg"`         // 评测信息
	Case       uint64 `json:"case"`        // 当前样例id
	IsFinished bool   `json:"is_finished"` // judge是否完成
}

type RootfsConfig struct {
	BasePath string `toml:"base_path"`
	WithProc bool   `toml:"with_proc"`
}

type TimeConfig struct {
	Seconds int64 `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty" toml:"secs"`
	Nanos   int32 `protobuf:"varint,2,opt,name=nanos,proto3" json:"nanos,omitempty" toml:"nanos"`
}
type TomlConfig struct {
	Source      string     `toml:"source"`
	Language    string     `toml:"language"`
	BuildScript *string    `toml:"build_script"`
	Timeout     TimeConfig `toml:"timeout"`
}

type Runner struct {
	Language string             `toml:"language"`
	Command  *string            `toml:"command"`
	Args     *[]string          `toml:"args"`
	Envs     *map[string]string `toml:"envs"`
	Rootfs   *RootfsConfig      `toml:"rootfs"`
}

type RunnerConfig struct {
	Runner Runner `toml:"runner"`
}

type SubmitData struct {
	Id       uint64
	Pid      uint64
	Language string
	Code     string
	// Build script must be absolute path
	BuildScript string
	// Runner toml path must be absolute path
	RunnerConfig string
}

type judger struct {
	baseDirectory string
	env           string
	anaAddress    string
	anaConn       *grpc.ClientConn
	tempDirectory string
}

type RLdata struct {
	Status string `json:"status"` // 状态
	Time   uint64 `json:"time"`   // 最大用时
	Memory uint64 `json:"memory"` // 最大内存
	Msg    string `json:"msg"`    // 评测信息
}

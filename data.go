package log

type Fields map[string]interface{}
type Struct []KV
type KV struct {
	K string
	V interface{}
}

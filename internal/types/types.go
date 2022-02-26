package types

type ConfigSource int

const (
	JsonConfigSource ConfigSource = iota
	EnvConfigSource
)

type Store int

const (
	S3Store Store = iota
	LocalStore
)

package types

type ConfigSource int

const (
	Json ConfigSource = iota
	Env               // TODO add config source from env vars
)

type Store int64

const (
	S3 Store = iota
	Local
)

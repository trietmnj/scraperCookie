package types

type ConfigSource int

const (
	JsonConfigSource ConfigSource = iota
	EnvConfigSource
)

type Store string

const (
	S3Store    Store = "s3"
	LocalStore Store = "local"
)

type Data string

const (
	HTMLTableData    Data = "htmltable"
	JSONEndpointData Data = "jsonendpoint"
)

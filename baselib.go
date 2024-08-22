package baselib

const Version = "0.1.0"

//go:noinline
func GetVersion() string {
	return Version
}

package baselib

const Version = "0.3.0"

//go:noinline
func GetVersion() string {
	return Version
}

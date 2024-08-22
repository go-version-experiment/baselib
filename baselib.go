package baselib

const Version = "0.2.0"

//go:noinline
func GetVersion() string {
	return Version
}

package baselib

const Version = "0.2.6"

//go:noinline
func GetVersion() string {
	return Version
}

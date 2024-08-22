package baselib

const Version = "0.2.5"

//go:noinline
func GetVersion() string {
	return Version
}

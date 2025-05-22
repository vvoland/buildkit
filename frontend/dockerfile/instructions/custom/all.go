package custom

import "strings"

type RunShortcut struct {
	CmdFmt string
	Flags  []string
}

var Runs = map[string]RunShortcut{
	"APT-INSTALL": {
		CmdFmt: "rm -f /etc/apt/apt.conf.d/docker-clean && " +
			"apt-get update && " +
			"DEBIAN_FRONTEND=noninteractive apt-get -y install %s",
		Flags: []string{
			"--mount=type=cache,target=/var/lib/apt,id=aptlib,sharing=locked",
			"--mount=type=cache,target=/var/cache/apt,id=aptcache,sharing=locked",
		},
	},
	"APK-ADD": {
		CmdFmt: "apk add --no-cache %s",
		Flags: []string{
			"--mount=type=cache,target=/var/cache/apk,id=apkcache,sharing=locked",
		},
	},
}

func Has(cmd string) bool {
	_, ok := Runs[strings.ToUpper(cmd)]
	return ok
}

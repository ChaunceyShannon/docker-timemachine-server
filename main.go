package main

import (
	. "github.com/ChaunceyShannon/golanglibs"
)

func main() {
	Lg.Trace("Add group '" + Os.Getenv("AFP_USER", "tm") + "' with GID '" + Os.Getenv("AFP_USER_GID", "1000") + "'")
	Os.System("groupadd -g " + Os.Getenv("AFP_USER_GID", "1000") + " " + Os.Getenv("AFP_USER", "tm"))

	Lg.Trace("Add user '" + Os.Getenv("AFP_USER", "tm") + "' with GID '" + Os.Getenv("AFP_USER_UID", "1000") + "'")
	Os.System("useradd -u " + Os.Getenv("AFP_USER_UID", "1000") + " -g " + Os.Getenv("AFP_USER", "tm") + " " + Os.Getenv("AFP_USER", "tm"))

	Lg.Trace("Change password for user '" + Os.Getenv("AFP_USER", "tm") + "' to '" + Os.Getenv("AFP_PASSWORD", "123456") + "'")
	Os.System("/bin/sh -c 'echo " + Os.Getenv("AFP_USER", "tm") + ":" + Os.Getenv("AFP_PASSWORD", "123456") + " | chpasswd'")

	Lg.Trace("Change owner for directory /timemachine to '" + Os.Getenv("AFP_USER", "tm") + "'")
	Os.System("chown -R " + Os.Getenv("AFP_USER", "tm") + ":" + Os.Getenv("AFP_USER", "tm") + " /timemachine")

	Lg.Trace("Write configuration to /etc/netatalk/afp.conf")
	Open("/etc/netatalk/afp.conf", "a").
		Write(`[` + Os.Getenv("AFP_VOL_NAME", "tm") + `]
path = /timemachine
time machine = yes
valid users = ` + Os.Getenv("AFP_USER", "tm") + `
spotlight = no
vol size limit = ` + Os.Getenv("AFP_SIZE_LIMIT", "tm") + "\n\n").
		Close()

	Lg.Trace("Start netatalk server ")
	Os.System("netatalk -d")
}

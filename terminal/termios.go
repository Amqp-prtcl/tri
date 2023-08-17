package terminal

import "golang.org/x/sys/unix"

func ToggleEcho() {
	t, err := unix.IoctlGetTermios(unix.Stdin, unix.TIOCGETA)
	if err != nil {
		panic(err)
	}
	t.Lflag ^= unix.ECHO

	if err = unix.IoctlSetTermios(unix.Stdin, unix.TIOCSETA, t); err != nil {
		panic(err)
	}
}

func ToggleCanonicalMode() {
	t, err := unix.IoctlGetTermios(unix.Stdin, unix.TIOCGETA)
	if err != nil {
		panic(err)
	}
	t.Lflag ^= unix.ICANON

	if err = unix.IoctlSetTermios(unix.Stdin, unix.TIOCSETA, t); err != nil {
		panic(err)
	}
}

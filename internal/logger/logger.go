package logger

import (
	"log/slog"
	"os"
	"sync"
	"time"

	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
)

var (
	once   sync.Once
	logger *slog.Logger
)

func Log() *slog.Logger {
	once.Do(func() {
		hanlder := slog.NewJSONHandler(os.Stdout, nil)
		logger = slog.New(hanlder)
	})
	return logger
}

func WishMiddlewareLogger() wish.Middleware {
	return func(next ssh.Handler) ssh.Handler {
		return func(sess ssh.Session) {
			ct := time.Now()
			hpk := sess.PublicKey() != nil
			pty, _, _ := sess.Pty()
			Log().Info(sess.User()+" connected",
				"ip", sess.RemoteAddr().String(),
				"hpk", hpk,
				"sessCmd", sess.Command(),
				"term", pty.Term,
				"wHeight", pty.Window.Height,
				"wWidth", pty.Window.Width,
				"clientV", sess.Context().ClientVersion())
			// fmt.Printf(
			// 	"%s connect %s %v %v %s %v %v %v hi",
			// 	sess.User(),
			// 	sess.RemoteAddr().String(),
			// 	hpk,
			// 	sess.Command(),
			// 	pty.Term,
			// 	pty.Window.Width,
			// 	pty.Window.Height,
			// 	sess.Context().ClientVersion(),
			// )
			next(sess)
			Log().Info(sess.User()+" disconnected",
				"ip", sess.RemoteAddr().String(),
				"duration", time.Since(ct),
			)
			// fmt.Printf(
			// 	"%s disconnect %s\n",
			// 	sess.RemoteAddr().String(),
			// 	time.Since(ct),
			// )
		}
	}
}

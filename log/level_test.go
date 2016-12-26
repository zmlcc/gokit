package log

import (
	"os"
)

func ExampleLevels() {
	logger := NewLevel(NewLogfmtLogger(os.Stdout), DebugLevel)
	logger.Debug("msg", "hello")
	logger.With("context", "foo").Warn("err", "error")

	// Output:
	// lvl=DEBU msg=hello
	// lvl=WARN context=foo err=error
}

func ExampleWith() {
	logger := NewLevel(NewLogfmtLogger(os.Stdout), DebugLevel)
	l1 := logger.With("context", "foo")
	l2 := logger.With("context", "bar")
	l1.Debug("msg", "debug")
	l2.Warn("msg", "warn")

	// Output:
	// lvl=DEBU context=foo msg=debug
	// lvl=WARN context=bar msg=warn
}

func ExampleWithLevel() {
	logger := NewLevel(NewLogfmtLogger(os.Stdout), DebugLevel).With("context", "foo")

	l3 := logger.WithLevel(InfoLevel)
	l4 := logger.WithLevel(WarnLevel)

	l3.Debug("msg", "L3 Debug")
	l3.Info("msg", "L3 Info")
	l3.Warn("msg", "L3 Warn")
	l3.Error("msg", "L3 Error")

	l4.Debug("msg", "L4 Debug")
	l4.Info("msg", "L4 Info")
	l4.Warn("msg", "L4 Warn")
	l4.Error("msg", "L4 Error")

	// Output:
	// lvl=INFO context=foo msg="L3 Info"
	// lvl=WARN context=foo msg="L3 Warn"
	// lvl=ERRO context=foo msg="L3 Error"
	// lvl=WARN context=foo msg="L4 Warn"
	// lvl=ERRO context=foo msg="L4 Error"

}

package glog

import (
	"strings"
	"testing"
)

var atomicTestVar atomicBool

func TestAtomicBool(t *testing.T) {
	if atomicTestVar.Get() {
		t.Error("the initial value of atomicBool is unexpected")
	}
	atomicTestVar.Set(true)
	if !atomicTestVar.Get() {
		t.Error("atomicBool failed")
	}
}

func TestGlogConfig_SetLogDir(t *testing.T) {
	expect := "/var/log"
	Config().SetLogDir(expect)
	if strings.Compare(expect, logDir) != 0 {
		t.Error("config logDir failed")
	}
}

func TestGlogConfig_SetToStderr(t *testing.T) {
	Config().SetToStderr(true)
	if !logging.toStderr {
		t.Error("config toStderr failed")
	}
}

func TestGlogConfig_SetAlsoToStderr(t *testing.T) {
	Config().SetAlsoToStderr(true)
	if !logging.alsoToStderr {
		t.Error("config alsoToStderr failed")
	}
}

func TestGlogConfig_SetStderrThreshold(t *testing.T) {
	Config().SetStderrThreshold("info")
	if logging.stderrThreshold != infoLog {
		t.Error("config stderrThreshold failed")
	}
	Config().SetStderrThreshold("a")
	if logging.stderrThreshold != infoLog {
		t.Error("map level failed during configing stderrThreshold")
	}
}

func TestGlogConfig_SetTraceLocation(t *testing.T) {
	expect := traceLocation{
		file: "a.go",
		line: 20,
	}
	Config().SetTraceLocation("a.go:20")
	if logging.traceLocation != expect {
		t.Error("config traceLocation failed")
	}
}

func TestGlogConfig_SetVmodule(t *testing.T) {
	Config().SetVmodule("a.b=4")
	t.Log(logging.vmodule)
}

func TestGlogConfig_SetVerbosity(t *testing.T) {
	Config().SetVerbosity("WARNING")
	t.Log(logging.verbosity.get())
}

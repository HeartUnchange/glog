package glog

import (
	"log"
	"sync/atomic"
)

// https://github.com/tevino/abool/blob/master/bool.go
type atomicBool struct{ flag int32 }

func (b *atomicBool) Set(value bool) {
	var i int32 = 0
	if value {
		i = 1
	}
	atomic.StoreInt32(&(b.flag), int32(i))
}

func (b *atomicBool) Get() bool {
	if atomic.LoadInt32(&(b.flag)) != 0 {
		return true
	}
	return false
}

var running atomicBool // once glog start output, changing config is not sensible, ignored

// Config refer to loggingT, just for view
type GlogConfig struct {

	//LogDir          string        // If non-empty, overrides the choice of directory in which to write logs.
	//ToStderr        bool          // The -logtostderr flag.
	//AlsoToStderr    bool          // The -alsologtostderr flag.
	//StderrThreshold severity      // The -stderrthreshold flag.
	//TraceLocation   traceLocation // traceLocation is the state of the -log_backtrace_at flag.
	//Vmodule         moduleSpec    // The state of the -vmodule flag.
	//Verbosity       Level         // V logging level, the value of the -v flag

	canChange bool
}

var _glc GlogConfig

func Config() *GlogConfig {

	if running.Get() {
		log.Print("glog have started, to change config is not sensible, ignored")
		_glc.canChange = false
	} else {
		_glc.canChange = true
	}
	return &_glc
}

// set log_dir
func (glc *GlogConfig) SetLogDir(path string) {
	if !glc.canChange {
		return
	}
	// TODO check path whether available or not
	// check will be helpful, but how to deal with link?
	//fi , err := os.Stat(path)
	//if err != nil || !fi.IsDir(){
	//	log.Printf("set the value LogDir for glog failed, path %s is not available or not a folder, ignored", path)
	//}
	logDir = path
}

// set whether log only to stderr
func (glc *GlogConfig) SetToStderr(fg bool) {
	if !glc.canChange {
		return
	}
	logging.toStderr = fg
}

// set whether log also to stderr
func (glc *GlogConfig) SetAlsoToStderr(fg bool) {
	if !glc.canChange {
		return
	}
	logging.alsoToStderr = fg
}

// set the threshold which determine whether write to stderr
func (glc *GlogConfig) SetStderrThreshold(level string) {
	if !glc.canChange {
		return
	}
	if s, b := severityByName(level); b {
		logging.stderrThreshold.set(s)
	} else {
		log.Printf("set the value StderrThreshold for glog failed, level %s is unrecognizable, ignored", level)
	}
}

// set TraceLocation
func (glc *GlogConfig) SetTraceLocation(location string) {
	if !glc.canChange {
		return
	}
	tl := traceLocation{}
	if err := tl.Set(location); err != nil {
		log.Printf("set the value TraceLocation for glog failed, location %s is unrecognizable, ignored", location)
		return
	}
	logging.traceLocation = tl
}

// set moduleSpec
func (glc *GlogConfig) SetVmodule(modulespec string) {
	if !glc.canChange {
		return
	}
	spec := moduleSpec{}
	if err := spec.Set(modulespec); err != nil {
		log.Printf("set the value Vmodule(moduleSpec) for glog failed, input %s is unrecognizable, ignored", modulespec)
		return
	}
	// logging.vmodule = spec
}

// set Verbosity globally
// 0 - 3, INFO, WARNING, ERROR, FATAL
func (glc *GlogConfig) SetVerbosity(level string) {
	if !glc.canChange {
		return
	}
	if s, b := severityByName(level); b {
		logging.setVState(Level(s), logging.vmodule.filter, false)
	} else {
		log.Printf("set the value Verbosity for glog failed, level %s is unrecognizable, ignored, ignored", level)
	}
}

func (glc *GlogConfig) SetVerbosityI(v int32) {
	if !glc.canChange {
		return
	}
	logging.setVState(Level(v), logging.vmodule.filter, false)
}

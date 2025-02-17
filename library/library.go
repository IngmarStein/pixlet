//go:build lib

package main

/*
#include <stdlib.h>
*/
import "C"

import (
	"encoding/json"
	"fmt"

	"tidbyt.dev/pixlet/runtime"
)

//export render_app
func render_app(namePtr *C.char, configPtr *C.char, width, height, magnify, maxDuration, timeout C.int, renderGif, silenceOutput C.int) (*C.uchar, *C.char) {
	name := C.GoString(namePtr)
	configStr := C.GoString(configPtr)

	var config map[string]string
	err := json.Unmarshal([]byte(configStr), &config)
	if err != nil {
		return nil, C.CString(fmt.Sprintf("error parsing config: %v", err))
	}

	result, err := runtime.RenderApplet(name, config, int(width), int(height), int(magnify), int(maxDuration), int(timeout), renderGif != 0, silenceOutput != 0)
	if err != nil {
		return nil, C.CString(err.Error())
	}
	return (*C.uchar)(C.CBytes(result)), nil
}

//export init_cache
func init_cache() {
	cache := runtime.NewInMemoryCache()
	runtime.InitHTTP(cache)
	runtime.InitCache(cache)
}

func main() {}

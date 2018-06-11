package main

//// #include <stdio.h>
//// #include <stdlib.h>
import "C"
import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/go-steem/rpc/transports/websocket"
)

var EmptyParams = struct{}{}

var servers = make(map[string]*websocket.Transport)
var serversMu sync.Mutex

//export call
func call(url string, action string, args []string) *C.char {
	url = url[:len(url)-1]
	action = action[:len(action)-1]

	var argsRaw = make([]json.RawMessage, len(args))
	for i, a := range args {
		//fmt.Printf("%d, %s, %d, %s\n", i, a, len(a), a[:len(a)-1])
		argsRaw[i] = json.RawMessage(a[:len(a)-1])
	}

	var s string
	if len(args) > 0 {
		s = callJSONRPC2(url, action, argsRaw)
	} else {
		s = callJSONRPC2(url, action, EmptyParams)
	}

	return C.CString(s)
}

func callJSONRPC2(url string, action string, args interface{}) string {
	var tran *websocket.Transport
	var err error
	serversMu.Lock()

	tran = servers[url]
	if tran == nil {
		tran, err = websocket.NewTransport([]string{url},
			websocket.SetAutoReconnectEnabled(true),
			websocket.SetAutoReconnectMaxDelay(time.Second),
			websocket.SetReadWriteTimeout(time.Minute))
		if err != nil {
			log.Printf("failed to new transport:%s", err.Error())
			serversMu.Unlock()
			return ""
		}

		servers[url] = tran
	}
	serversMu.Unlock()

	var resp json.RawMessage
	err = tran.Call(action, args, &resp)
	if err != nil {
		log.Printf("failed to call: %v", err)
	}
	return string(resp)
}

func main() {
	// url := "ws://8.8.8.8:38090"

	// // test 1
	// action := "database_api.get_dynamic_global_properties"
	// resp := callJSONRPC2(url, action, EmptyParams)
	// log.Printf("resp1: %s", resp)

	// // test 2
	// action = "condenser_api.get_accounts"
	// var arg1 = json.RawMessage(`["wb-100", "wb-200"]`)
	// args := []interface{}{arg1}
	// resp = callJSONRPC2(url, action, args)
	// log.Printf("resp2: %s", resp)

	// // test 3
	// action := "database_api.get_dynamic_global_propertiesd"
	// cresp := call(url, action, nil)
	// log.Printf("resp3: %s", C.GoString(cresp))
	// C.free(unsafe.Pointer(cresp))

	// // test 4
	// action := "condenser_api.get_accountsd"
	// var arg1 = `["wb-100", "wb-200"]d`
	// args := []string{arg1}
	// cresp := call(url, action, args)
	// log.Printf("resp4: %s", C.GoString(cresp))
	// C.free(unsafe.Pointer(cresp))

}

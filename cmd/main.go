package main

import (
	"github.com/sevengrand/huobi_Golang/cmd/accountclientexample"
	"github.com/sevengrand/huobi_Golang/cmd/accountwebsocketclientexample"
	"github.com/sevengrand/huobi_Golang/cmd/algoorderclientexample"
	"github.com/sevengrand/huobi_Golang/cmd/commonclientexample"
	"github.com/sevengrand/huobi_Golang/cmd/crossmarginclientexample"
	"github.com/sevengrand/huobi_Golang/cmd/etfclientexample"
	"github.com/sevengrand/huobi_Golang/cmd/isolatedmarginclientexample"
	"github.com/sevengrand/huobi_Golang/cmd/marketclientexample"
	"github.com/sevengrand/huobi_Golang/cmd/marketwebsocketclientexample"
	"github.com/sevengrand/huobi_Golang/cmd/orderclientexample"
	"github.com/sevengrand/huobi_Golang/cmd/orderwebsocketclientexample"
	"github.com/sevengrand/huobi_Golang/cmd/stablecoinclientexample"
	"github.com/sevengrand/huobi_Golang/cmd/subuserclientexample"
	"github.com/sevengrand/huobi_Golang/cmd/walletclientexample"
	"github.com/sevengrand/huobi_Golang/logging/perflogger"
)

func main() {
	runAll()
}

// Run all examples
func runAll() {
	commonclientexample.RunAllExamples()
	accountclientexample.RunAllExamples()
	orderclientexample.RunAllExamples()
	algoorderclientexample.RunAllExamples()
	marketclientexample.RunAllExamples()
	isolatedmarginclientexample.RunAllExamples()
	crossmarginclientexample.RunAllExamples()
	walletclientexample.RunAllExamples()
	subuserclientexample.RunAllExamples()
	stablecoinclientexample.RunAllExamples()
	etfclientexample.RunAllExamples()
	marketwebsocketclientexample.RunAllExamples()
	accountwebsocketclientexample.RunAllExamples()
	orderwebsocketclientexample.RunAllExamples()
}

// Run performance test
func runPerfTest() {
	perflogger.Enable(true)

	commonclientexample.RunAllExamples()
	accountclientexample.RunAllExamples()
	orderclientexample.RunAllExamples()
	marketclientexample.RunAllExamples()
	isolatedmarginclientexample.RunAllExamples()
	crossmarginclientexample.RunAllExamples()
	walletclientexample.RunAllExamples()
	etfclientexample.RunAllExamples()
}

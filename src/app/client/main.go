package main

import (
	"TicketReservation/src/app/client/ui"
	"TicketReservation/src/rest"
	"TicketReservation/src/rest/client"
	"flag"
)

type commandlineArgs struct {
	port          string
	testBurstMode bool
	clientNumber  int
	reqNumber     int
}

func getClientArgs() (args commandlineArgs) {
	flag.StringVar(&args.port, "p", rest.DefaultPort, "Provide a port number")
	flag.BoolVar(&args.testBurstMode, "test", false, "Test Burst Mode (put server under pressure)")
	flag.IntVar(&args.clientNumber, "client", 5, "Number of clients (in test mode)")
	flag.IntVar(&args.reqNumber, "pressure", 10, "Number of requests of each client (in test mode)")

	flag.Parse()
	return
}

func main() {
	args := getClientArgs()

	clientAPI := client.NewClient(args.port)
	app := ui.NewApp(clientAPI)

	if args.testBurstMode {
		app.RunTest(args.clientNumber, args.reqNumber)
	} else {
		app.RunUI()
	}
}

package main

import (
	"fmt"
	"github.com/DavidGamba/go-getoptions"
	"github.com/things-go/go-socks5"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)
func main() {
	var listenPort int
	var listenAddress string
	var blackListFileAddress string
	opt := getoptions.New()
	opt.Bool("help", false, opt.Alias("h", "?"))
	opt.IntVar(&listenPort, "port", 51081,opt.Required(),opt.Description("Listening Port"))
	opt.StringVarOptional(&listenAddress, "address", "127.0.0.1",opt.Description("Listening Address"))
	opt.StringVarOptional(&blackListFileAddress, "blacklist", "black_list",opt.Description("Blacklist File path"))
	_, err := opt.Parse(os.Args[1:])
	if opt.Called("help") {
		fmt.Fprintf(os.Stderr, opt.Help())
		os.Exit(1)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n\n", err)
		fmt.Fprintf(os.Stderr, opt.Help(getoptions.HelpSynopsis))
		os.Exit(1)
	}
	contentBytes, err := ioutil.ReadFile(blackListFileAddress)
	if err != nil {
		panic(err)
	}
	contentString := strings.ReplaceAll(string(contentBytes),"\r","")
	blackList := strings.Split(contentString ,"\n")
	var hostRegexes = make([]*regexp.Regexp,len(blackList))

	for i , host := range blackList {
		hostRegexes[i] , _ = regexp.Compile(host)
	}
	resolver := new(CustomResolver)
	resolver.blockList = hostRegexes
	server := socks5.NewServer(
		socks5.WithLogger(socks5.NewLogger(log.New(os.Stdout, "socks5: ", log.LstdFlags))),
		socks5.WithResolver(resolver),
	)

	if err := server.ListenAndServe("tcp", fmt.Sprintf("%s:%d",listenAddress,listenPort)); err != nil {
		panic(err)
	}
}

package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"regexp"
)

type CustomResolver struct {
	blockList []*regexp.Regexp
}

func (c CustomResolver) Resolve(ctx context.Context, name string) (context.Context, net.IP, error) {
	for _,r := range c.blockList {
		if r.MatchString(name) {
			fmt.Fprintf(os.Stdout,"%s blocked by user rule\n",name)
			return ctx,net.ParseIP("127.0.0.1"),nil
		}
	}
	ip , err := net.LookupIP(name)
	if err != nil {
		return ctx,nil,err
	}
	if len(ip) > 0 {
		fmt.Printf("%s -> %s\n", name, ip[0].String())
		return ctx, ip[0], nil
	}else{
		return ctx,nil,errors.New("0 A Record")
	}
}
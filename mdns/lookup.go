package mdns

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net"
	"strings"
	"time"
)

var dnslist []*net.Resolver

func init() {
	if len(dnslist) == 0 {
		Initialization()
	}
}

func Initialization() {
	dnslist = []*net.Resolver{}
	nss := viper.GetStringSlice("dns.nameserver")
	log.Infof("nameservers: %s", strings.Join(nss, ","))
	for _, n := range nss {
		r := resolverBuilder(n)
		dnslist = append(dnslist, r)
	}
}

func resolverBuilder(nshost string) *net.Resolver {
	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			log.Debugf("== lookup with nshost %s ==\n", nshost)
			return d.DialContext(ctx, network, nshost)
		},
	}
	return r
}

func Lookup(domainanme string) (ip []*net.IP) {
	log.Debugf("lookup '%s' with external dns", domainanme)
	for _, d := range dnslist {
		ips, _ := d.LookupHost(context.Background(), domainanme)
		if len(ips) != 0 {
			tip := net.ParseIP(ips[0])
			ip = append(ip, &tip)
			break
		}
	}
	return
}

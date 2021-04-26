package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net"
)

func FetchDNSRecords() (list map[string]net.IP) {
	log.Info("will fetch ens records")
	key := "dns.records"
	records := viper.GetStringMapString(key)
	list = make(map[string]net.IP)
	for k, v := range records {
		if mip := net.ParseIP(v); mip != nil {
			list[k] = net.ParseIP(v)
		} else {
			log.Error("this record not vaild: %s - %s", k, v)
		}
	}
	return
}

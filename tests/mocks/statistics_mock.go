package mocks

import (
	"github.com/smbody/common-server/bl"
	"fmt"
)

type StatisticsMock struct {
	ServerName string
	Version string
}


func NewStatisticsMock() bl.Statistics {
 return  &StatisticsMock{}
}

func (s *StatisticsMock) ServerInfo() string {
	return  fmt.Sprintf("Server: %s (ver. %s)", s.ServerName, s.Version)
}

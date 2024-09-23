package utils

import (
	"net"
	"net/url"
	"strings"
)


func IsGitHubURL(input string) bool {
    u, err := url.Parse(input)
    if err != nil {
        return false
    }
    host := u.Host
    if strings.Contains(host, ":") { 
        host, _, err = net.SplitHostPort(host)
        if err != nil {
            return false
        }
    }
    return host == "github.com"
}
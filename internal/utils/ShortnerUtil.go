package utils

import (
	"context"
	"errors"
	"net"
	"net/url"
	"strings"
	"time"
)

func IsURLReachable(rawURL string) (bool, error) {
	if !strings.HasPrefix(rawURL, "http://") && !strings.HasPrefix(rawURL, "https://") {
		rawURL = "https://" + rawURL
	}

	parsedURL, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return false, errors.New("invalid URL format")
	}

	host := parsedURL.Host
	if !strings.Contains(host, ":") {
		if parsedURL.Scheme == "https" {
			host += ":443"
		} else {
			host += ":80"
		}
	}

	// 1. Set a hard timeout for entire resolution + dial process
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// DNS lookup and open a TCP connection
	// Not Used HEAD beacause some website might have blocked (possibly for bots, security, or abuse prevention)
	dialer := &net.Dialer{}
	conn, err := dialer.DialContext(ctx, "tcp", host)
	if err != nil {
		return false, errors.New("host is not reachable: " + err.Error())
	}
	defer conn.Close()

	return true, nil
}

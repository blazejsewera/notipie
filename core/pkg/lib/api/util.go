package api

import (
	"net/url"
	"strconv"
)

func GetHost(addr string, port int) string {
	return addr + ":" + strconv.Itoa(port)
}

func GetPath(p Path) string {
	return "/" + p.Path
}

func GetURLStr(host string, path Path) string {
	return path.Schema + "://" + host + GetPath(path)
}

func GetURL(host string, path Path) url.URL {
	return url.URL{
		Scheme: path.Schema,
		Host:   host,
		Path:   GetPath(path),
	}
}

package ginplus

import (
	"fmt"
	"github.com/SongZihuan/cat-shop-backend/src/config"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/abort"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/header"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/writer"
	"github.com/SongZihuan/cat-shop-backend/src/utils"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"strings"
)

func Writer() gin.HandlerFunc {
	return func(c *gin.Context) {
		w := writer.GinContextUseNewWriter(c)
		c.Next()
		_, err := w.WriteToHttp()
		if err != nil {
			// 允许使用c.Abort系列函数的地方
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}
}

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			err := recover()
			if err != nil {
				abort.ServerError(c, err)
				debugPrint("[ERROR] Recover: %v", err)
				if config.Config().Yaml.GlobalConfig.IsDebug() {
					panic(err)
				}
			}
		}()

		c.Next()
	}
}

func Forwarded() gin.HandlerFunc {
	return func(c *gin.Context) {
		ForwardedHeader := c.Request.Header.Get(header.RequestsForwarded)
		if ForwardedHeader != "" {
			ForwardedList := strings.Split(ForwardedHeader, ",")
			NewForwardedList := make([]string, 0, len(ForwardedList)+1)
			ProxyList := make([]string, 0, len(ForwardedList)+1)

			host, _ := utils.SplitHostPort(c.Request.Host) // 去除host中的端口号
			proto := "http"
			if c.Request.TLS != nil {
				proto = "https"
			}

			for _, keyStr := range ForwardedList {
				kv := strings.Split(strings.ReplaceAll(keyStr, " ", ""), "=")
				if len(kv) != 2 {
					continue
				}

				if kv[0] == "for" {
					forIP := net.ParseIP(strings.TrimSpace(kv[1]))
					if forIP != nil {
						NewForwardedList = append(NewForwardedList, keyStr)
						ProxyList = append(ProxyList, forIP.String())
					} else if kv[1] == "_hidden" || kv[1] == "_secret" || kv[1] == "unknown" {
						NewForwardedList = append(NewForwardedList, keyStr)
					}
				} else if kv[0] == "by" {
					byIP := net.ParseIP(strings.TrimSpace(kv[1]))
					if byIP != nil || kv[1] == "_hidden" || kv[1] == "_secret" || kv[1] == "unknown" {
						NewForwardedList = append(NewForwardedList, keyStr)
					}
				} else if kv[0] == "host" {
					host = kv[1]
				} else if kv[0] == "proto" {
					proto = kv[1]
				}
			}

			ProxyList = append(ProxyList, c.RemoteIP())
			NewForwardedList = append(NewForwardedList, fmt.Sprintf("for=%s", c.RemoteIP()))
			NewForwardedList = append(NewForwardedList, fmt.Sprintf("host=%s", host))
			NewForwardedList = append(NewForwardedList, fmt.Sprintf("proto=%s", proto))

			c.Request.Header.Set(header.RequestsForwarded, strings.Join(NewForwardedList, ", "))
			c.Request.Header.Set(header.RequestsXForwardedFor, strings.Join(ProxyList, ", "))
			c.Request.Header.Set(header.RequestsXForwardedHost, host)
			c.Request.Header.Set(header.RequestsXForwardedProto, proto)
		}

		c.Next()
	}
}

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

var lock = &sync.RWMutex{}
var servers []Server

func getIndex(c *gin.Context) {
	lock.RLock()
	defer lock.RUnlock()
	c.HTML(http.StatusOK, "template.html", gin.H{
		"servers": servers,
	})
}

func postGpu(c *gin.Context) {
	ip := c.ClientIP()
	fmt.Println(ip)
	var gpuInfo GPUInfo
	if c.ShouldBindXML(&gpuInfo) == nil {
		lock.Lock()
		defer lock.Unlock()
		found := false
		for i, _ := range servers {
			if servers[i].Ip == ip {
				found = true
				servers[i].GPUInfo = gpuInfo
			}
		}
		if !found {
			servers = append(servers, Server{
				Ip:      ip,
				GPUInfo: gpuInfo,
			})
		}
		c.Status(http.StatusOK)
	} else {
		c.Status(http.StatusBadRequest)
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// 1. 实现一个 HTTP Server
// 2. 实现一个 Handle : hello
// 3. 实现中间件功能：一.记录请求的URL和方法 二.记录请求的网络地址 三.记录方法的执行时间

func tracing(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("tracing start - 记录请求的URL和方法: %s, %s", r.URL, r.Method)
		next.ServeHTTP(w, r)
		log.Println("tracing end")
	})
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("logging start - 记录请求的网络地址: %s", r.RemoteAddr)
		next.ServeHTTP(w, r)
		log.Println("logging end")
	})
}

func timeRecording(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("timeRecording start")
		startTime := time.Now()
		next.ServeHTTP(w, r)
		endTime := time.Since(startTime)
		log.Printf("记录方法的执行时间: %s", endTime)
		log.Println("timeRecording end")
	})
}

func hello(w http.ResponseWriter, r *http.Request) {
	//log.Printf("记录请求的URL和方法: %s, %s", r.URL, r.Method)
	//log.Printf("记录请求的网络地址: %s", r.RemoteAddr)
	//startTime := time.Now()
	_, err := fmt.Fprintf(w, "hello")
	if err != nil {
		panic(err)
	}
	//endTime := time.Since(startTime)
	//log.Printf("记录方法的执行时间: %s", endTime)

}

func main() {
	http.Handle("/", tracing(logging(timeRecording(http.HandlerFunc(hello)))))
	http.ListenAndServe(":8080", nil)
}

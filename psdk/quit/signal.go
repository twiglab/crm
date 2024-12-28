package quit

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func WaitSignal(signals ...syscall.Signal) {
	shutdownHook := make(chan os.Signal, 1)
	signal.Notify(shutdownHook,
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, os.Interrupt)
	localPid := os.Getpid()
	sig := <-shutdownHook
	log.Printf("\ncaught sig exit sig:%v,localPid:%v\n", sig, localPid)
	close(shutdownHook)
}

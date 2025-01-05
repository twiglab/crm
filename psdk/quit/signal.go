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
	log.Printf("pid:%d\n", localPid)
	sig := <-shutdownHook
	log.Printf("caught sig exit sig:%v\n", sig)
	close(shutdownHook)
}

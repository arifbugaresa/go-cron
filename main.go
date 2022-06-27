package main

import (
	"fmt"
	"github.com/robfig/cron"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	// set scheduler berdasarkan zona waktu
	jakartaTime, _ := time.LoadLocation("Asia/Jakarta")
	scheduler := cron.NewWithLocation(jakartaTime)

	// defer untuk stop scheduler sebelum fungsi berakhir
	defer scheduler.Stop()

	// set task yang akan dijalankan di scheduler
	scheduler.AddFunc("0 */1 * * * *", func() {
		SendAutoEveryMinute("Hi.....")
	})

	// set task yang akan dijalankan di scheduler
	scheduler.AddFunc("* */1 * * * *", SendAutoEverySecond)

	// start scheduller
	scheduler.Start()

	// trap SIGINT untuk trigger shutdown.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

}

func SendAutoEveryMinute(mail string) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + mail)
}

func SendAutoEverySecond()  {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + "SendAutoEverySecond")
}

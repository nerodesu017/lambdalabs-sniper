package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/nerodesu017/lambdalabs-sniper/src/constants"
	"github.com/nerodesu017/lambdalabs-sniper/src/logged_in"
	"github.com/nerodesu017/lambdalabs-sniper/src/monitor"
	"github.com/nerodesu017/lambdalabs-sniper/src/notifier"
	"github.com/nerodesu017/lambdalabs-sniper/src/utils"
)

func main() {
	log_file, err := os.OpenFile("default.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error when opening log file: %v\n", err)
	}
	defer log_file.Close()
	log.SetOutput(log_file)

	log.Printf("Starting monitor... (Sniper not yet implemented)")

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v\n", err)
	}

	log.Printf("Successfully imported .env file!")

	
	email, err:= logged_in.GetEmail(os.Getenv("SESSION_ID"))
	if err != nil {
		log.Fatalf("error when logging in")
	}

	log.Printf("Connected as '%s'\n", email)

	discord_notifier := notifier.Discord_notifier{
		Webhook_url: os.Getenv("DISCORD_WEBHOOK"),
	}

	// array for keeping track of last known available GPUs
	last_known_available_gpus := make([]*constants.GPU, 0)

	for ;; {

		new_found_gpus, err := monitor.FindAvailable()
		if err != nil {
			log.Fatalf("error when monitoring: %v\n", err)
		}

		diff_gpus := utils.GetGPUDiff(last_known_available_gpus, new_found_gpus)

		log.Printf("New gpus: ")
		for i, gpu := range diff_gpus {
			if i == len(diff_gpus) - 1 {
				log.Println(gpu.Name)
			} else {
				log.Printf("%s, \n", gpu.Name)
			}
		}

		if len(diff_gpus) != 0 {
			err = discord_notifier.Notify(diff_gpus)
			if err != nil {
				log.Fatalf("error when sending webhook: %v\n", err)
			}
		}

		last_known_available_gpus = new_found_gpus

		miliseconds := time.Duration(utils.GetRandSleepTimeInMs(3000, 5000))

		log.Printf("Sleeping for %d miliseconds\n", miliseconds)

		// every 3-5 seconds
		time.Sleep(time.Millisecond * miliseconds)
	}
}
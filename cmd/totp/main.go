package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/rohitaryal/totp"
)

func main() {
	generate := flag.Bool("gen", false, "Generate a secret base32 encoded key")
	timestamp := flag.Int64("time", -1, "Forced Unix timestamp as reference")
	key := flag.String("key", "", "Secret base32 encoded key to be used for OTP generation")
	verbose := flag.Bool("v", true, "Be extra polite [Turn off if used in scripting]")

	flag.Usage = func() {
		fmt.Print("totp: The CLI way to time-based OTPs\n\n")
		flag.PrintDefaults()
		fmt.Println("\nNOTE: Use -v false for scripted usage")
	}

	// Never to be forgotten
	flag.Parse()

	if *generate && (*key != "" || *timestamp >= 0) {
		fmt.Println("The combination of those flags doesn't make sense.")
		os.Exit(1)
	}

	if *generate {
		handleSecretGeneration(*verbose)
		os.Exit(0)
	}

	if *key != "" {
		if *timestamp <= -1 {
			currentTime := time.Now().Unix()

			// If user deliberately tried to provide a negative timestamp
			// notify them for their bad behaviour
			if *timestamp != -1 && *verbose {
				fmt.Println("[!] Timestamp must be >= 0, defaulting to current time:", currentTime)
			}

			*timestamp = currentTime
		}

		handleOtpGeneration(*key, *timestamp, *verbose)
		os.Exit(0)
	}

	flag.Usage()
}

func handleSecretGeneration(verbose bool) {
	// Occurance of error is unlikely given that it will take 20-30 byte space
	// in RAM or cache. So just ignore err, this makes code better too
	key, _ := totp.GenerateSecret()

	if verbose {
		fmt.Print("[+] Secret key: ")
	}

	fmt.Println(key)

	if verbose {
		fmt.Println()
	}
}

func handleOtpGeneration(secretKey string, timestamp int64, verbose bool) {
	otp, err := totp.GenerateTotp(secretKey, timestamp)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	if verbose {
		fmt.Print("[+] OTP: ")
	}

	fmt.Print(otp)

	if verbose {
		fmt.Println()
	}
}

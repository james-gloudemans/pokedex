package main

import "os"

func commandExit(_ *Config) error {
	os.Exit(0)
	return nil
}

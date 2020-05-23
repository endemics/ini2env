package main

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

func main() {
	prefix := "INI"
	cfg, err := ini.Load("file.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	sections := cfg.SectionStrings()
	for _, section := range sections {
		keys := cfg.Section(section).KeyStrings()
		for _, key := range keys {
			fmt.Printf("%s__%s__%s=\"%s\"\n",
				prefix,
				section,
				key,
				cfg.Section(section).Key(key).String(),
			)
		}
	}
}

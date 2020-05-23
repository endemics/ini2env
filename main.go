package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

var (
	prefix string
	file   string
)

func main() {
	flag.StringVar(&prefix, "prefix", "INI", "the prefix to add to the variable names")
	flag.StringVar(&file, "file", "file.ini", "the ini file to parse")
	flag.Parse()

	cfg, err := ini.Load(file)
	if err != nil {
		fmt.Printf("Fail to read file: %v\n", err)
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

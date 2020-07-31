package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

var (
	prefix       string
	file         string
	bool_convert bool
)

func main() {
	flag.StringVar(&prefix, "prefix", "INI", "the prefix to add to the variable names")
	flag.StringVar(&file, "file", "file.ini", "the ini file to parse")
	flag.BoolVar(&bool_convert, "booleans", false, "recognise special boolean values and convert them to 1 (true) or 0 (false)")
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
			value := cfg.Section(section).Key(key).String()

			b, err := cfg.Section(section).Key(key).Bool()
			if bool_convert && err == nil {
				if b {
					value = "1"
				} else {
					value = "0"
				}
			}

			fmt.Printf("%s__%s__%s=\"%s\"\n",
				prefix,
				section,
				key,
				value,
			)
		}
	}
}

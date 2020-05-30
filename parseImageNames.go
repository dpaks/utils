package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var REGISTRY = "localhost:443/roche"

func getImageTag(image string) string {
	s := strings.Split(image, "TO")
	s = strings.Split(s[0], ":")
	if len(s) == 1 {
		return "latest"
	}
	return strings.TrimSpace(s[len(s)-1])
}

func getImageName(image string) string {
	s := strings.Split(image, "TO")
	return strings.TrimSpace(s[0])
}

func getOldImageName(image string) string {
	s := strings.Split(image, "TO")
	s = strings.Split(s[0], "/")
	s = strings.Split(s[len(s)-1], ":")
	return strings.TrimSpace(s[0])
}

func getNewImageName(image string) string {
	s := strings.Split(image, "TO")
	s = strings.Split(s[len(s)-1], "/")
	s = strings.Split(s[len(s)-1], ":")
	return strings.TrimSpace(s[0])
}

func process() {
	err := os.Remove("buildNload.sh")
	if err != nil {
		panic(err)
	}
	f, err := os.Create("buildNload.sh")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	images := getImages()

	w := bufio.NewWriter(f)
	txt := fmt.Sprintf(`#!/bin/bash

GREEN='\033[0;32m'
NC='\033[0m'

# Run: sed -i 's/<none>/latest/g' buildNload.sh
# export path
IPATH=/images
`)
	_, err = w.WriteString(txt)
	if err != nil {
		panic(err)
	}
	w.Flush()
	f.Sync()

	for _, each := range images {
		w := bufio.NewWriter(f)
		if strings.TrimSpace(each) == "" {
			continue
		}

		//iName := getOldImageName(each)
		iTag := getImageTag(each)
		nName := getNewImageName(each)
		name := getImageName(each)

		txt := fmt.Sprintf("\nprintf \"\\n${GREEN}Pulling %s${NC}\\n\"\n", name)
		txt += fmt.Sprintf("docker pull %s\n", name)
		txt += fmt.Sprintf("docker tag %s %s/%s:%s\n", name, REGISTRY, nName, iTag)
		txt += fmt.Sprintf("\nprintf \"\\n${GREEN}Pushing %s/%s:%s${NC}\\n\"\n", REGISTRY, nName, iTag)
		txt += fmt.Sprintf("docker push %s/%s:%s\n", REGISTRY, nName, iTag)
		_, err := w.WriteString(txt)
		if err != nil {
			panic(err)
		}
		w.Flush()
		//fmt.Printf("docker save %s/%s:%s -o $IPATH/%s-%s.tar\n", REGISTRY, iName, iTag, iName, iTag)
	}
	f.Sync()

	/*
		for _, each := range images {
			w := bufio.NewWriter(f)
			if strings.TrimSpace(each) == "" {
				continue
			}

			iTag := getImageTag(each)
			nName := getNewImageName(each)

			txt := fmt.Sprintf("\nprintf \"\\n${GREEN}Pushing %s/%s:%s${NC}\\n\"\n", REGISTRY, nName, iTag)
			//fmt.Printf("docker load --input $path/%s-%s.tar\n", iName, iTag)
			txt += fmt.Sprintf("docker push %s/%s:%s\n", REGISTRY, nName, iTag)
			_, err := w.WriteString(txt)
			if err != nil {
				panic(err)
			}
			w.Flush()
		}
		f.Sync()
	*/
}

func main() {
	if os.Geteuid() != 0 {
		panic(fmt.Errorf("Run with root permission"))
	}

	process()

	_, err := exec.Command("bash", "-c", "sed -i 's/<none>/latest/g' buildNload.sh").Output()
	if err != nil {
		panic(err)
	}

	_, err = exec.Command("bash", "-c", "chmod +x buildNload.sh").Output()
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("./buildNload.sh")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	cmd.Start()

	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}
	cmd.Wait()
}

package main

import "flag"

func main() {
	jobName := flag.String("jobname", "test-job", "the name of the job")

	containerImage := flag.String("image", "ubuntu:latest", "name of the container image")

}

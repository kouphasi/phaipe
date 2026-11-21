package main

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
)

const version = "0.1.0"

type config struct {
	protocol string
	port     int
}

func main() {

	if hasVersion(os.Args) {
		mes := fmt.Sprintf("phipe version is %s", version)
		fmt.Fprintln(os.Stdout, mes)
		os.Exit(0)
	}

	cnf, err := parseArgs(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("protocol: %s\nport: %d\n", cnf.protocol, cnf.port)
}

func hasVersion(args []string) bool {
	return (args[1] == "--version")
}

func parseArgs(args []string) (config, error) {
	if len(args) != 3 {
		return config{}, errors.New("please specify 2 arguments: protocol and port")
	}

	protocol, err := parseProtocol(args[1])
	if err != nil {
		return config{}, err
	}

	port, err := parsePort(args[2])

	if err != nil {
		return config{}, err
	}

	return config{
		protocol: protocol,
		port:     port,
	}, nil
}

func parseProtocol(protocol string) (string, error) {
	validProtocols := []string{"http"}
	if !slices.Contains(validProtocols, protocol) {
		return "", fmt.Errorf("invalid protocol: %s", protocol)
	}

	return protocol, nil
}

func parsePort(rawPort string) (int, error) {
	port, err := strconv.Atoi(rawPort)

	if err != nil {
		return 0, fmt.Errorf("invalid port: %q, %w", rawPort, err)
	}
	if port < 1 || port > 65535 {
		return 0, errors.New("please enter a port number between 1 and 65535")
	}
	return port, nil
}

package secrets

import (
	"os"
	"strings"
)

type Secrets map[string]string

func ReadSecrets(name string) (Secrets, error) {
	secrets := make(Secrets)
	err := readFile("general", secrets)
	if err != nil {
		return nil, err
	}
	err = readFile(name, secrets)
	if err != nil {
		return nil, err
	}
	return secrets, nil
}

func readFile(name string, secrets Secrets) error {
	data, err := os.ReadFile("secrets/secrets/" + name)
	if err != nil {
		return err
	}
	file := string(data)
	for _, line := range strings.Split(file, "\n") {
		if strings.Contains(line, "-:-") {
			s := strings.Split(line, "-:-")
			secrets[s[0]] = s[1]
		} else {
			continue
		}
	}
	return nil
}

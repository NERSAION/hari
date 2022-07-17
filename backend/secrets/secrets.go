package secrets

import (
	"os"
	"strings"
)

type Secrets map[string]string

func ReadSecrets(name string) (Secrets, error) {
	secrets := make(Secrets)
	data, err := os.ReadFile("secrets/" + name)
	if err != nil {
		return nil, err
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
	return secrets, nil
}

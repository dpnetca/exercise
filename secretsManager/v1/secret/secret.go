package secret

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Vault struct {
	file string
}

func NewVault(file string) (Vault, error) {
	vault := Vault{
		file: file,
	}
	return vault, nil
}

func (v *Vault) Set(key, value string) error {
	secrets, err := v.loadSecrets()
	if err != nil {
		secrets = make(map[string]string)
	}
	if len(value) > 0 {
		cryptoValue, err := encrypt(value)
		if err != nil {
			return fmt.Errorf("error encrypting secret: %v", err)
		}
		secrets[key] = cryptoValue
	} else {
		delete(secrets, key)
	}
	err = v.writeSecrets(secrets)
	if err != nil {
		return fmt.Errorf("unable to write secrets file: %v", err)
	}
	return nil
}

func (v *Vault) Delete(key string) error {
	return v.Set(key, "")
}

func (v *Vault) Get(key string) (string, error) {
	secrets, err := v.loadSecrets()
	if err != nil {
		return "", fmt.Errorf("unable to open/parse secrets file: %v", err)
	}
	cryptoValue, ok := secrets[key]
	if !ok {
		return "", fmt.Errorf("no value for key (%v)", key)
	}
	value, err := decrypt(cryptoValue)
	if err != nil {
		return "", fmt.Errorf("unable to decrypt secret: %v", err)
	}
	return value, nil
}

func (v *Vault) loadSecrets() (map[string]string, error) {
	secrets := make(map[string]string)
	cryptoData, err := ioutil.ReadFile(v.file)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	data, err := decrypt(string(cryptoData))
	if err != nil {
		return nil, fmt.Errorf("error decrypting file: %v", err)
	}

	err = json.Unmarshal([]byte(data), &secrets)
	if err != nil {
		return nil, fmt.Errorf("error parsing json: %v", err)
	}
	return secrets, nil
}

func (v *Vault) writeSecrets(secrets map[string]string) error {
	data, err := json.Marshal(secrets)
	if err != nil {
		return fmt.Errorf("unable to encode secrets: %v", err)
	}

	f, err := os.OpenFile(v.file, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0660)
	if err != nil {
		return fmt.Errorf("unable to open secrets file: %v", err)
	}
	defer f.Close()

	cryptoData, err := encrypt(string(data))
	if err != nil {
		return fmt.Errorf("unale to encrype file data: %v", err)
	}

	_, err = fmt.Fprint(f, string(cryptoData))
	if err != nil {
		return fmt.Errorf("unable to write secrets file: %v", err)
	}
	return nil
}

func encrypt(plainText string) (string, error) {
	cryptoText := "xx" + plainText // TODO Implement
	return cryptoText, nil
}

func decrypt(cryptoText string) (string, error) {
	if len(cryptoText) > 0 {
		plainText := cryptoText[2:] // TODO Implement
		return plainText, nil
	}
	return "", nil
}

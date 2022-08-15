package secret

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/dpnetca/exercise/secretsManager/v3/crypt"
)

type Vault struct {
	cryptKey string
	file     string
	mutex    sync.Mutex
}

func NewVault(file, key string) *Vault {
	return &Vault{
		file:     file,
		cryptKey: key,
	}
}

func (v *Vault) Set(key, value string) error {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	secrets, err := v.loadSecrets()
	if err != nil {
		secrets = make(map[string]string)
	}
	if len(value) > 0 {
		cryptoValue, err := crypt.Encrypt(value, key+v.cryptKey)
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
	v.mutex.Lock()
	defer v.mutex.Unlock()

	secrets, err := v.loadSecrets()
	if err != nil {
		return "", fmt.Errorf("unable to open/parse secrets file: %v", err)
	}
	cipherHex, ok := secrets[key]
	if !ok {
		return "", fmt.Errorf("no value for key (%v)", key)
	}
	value, err := crypt.Decrypt(cipherHex, key+v.cryptKey)
	if err != nil {
		return "", fmt.Errorf("unable to decrypt secret: %v", err)
	}
	return value, nil
}

func (v *Vault) loadSecrets() (map[string]string, error) {
	f, err := os.Open(v.file)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	defer f.Close()

	r, err := crypt.DecryptReader(v.cryptKey, f)
	if err != nil {
		return nil, fmt.Errorf("error decrypting file: %v", err)
	}

	secrets := make(map[string]string)
	err = json.NewDecoder(r).Decode(&secrets)
	if err != nil {
		return nil, fmt.Errorf("error parsing json: %v", err)
	}

	return secrets, nil
}

func (v *Vault) writeSecrets(secrets map[string]string) error {
	f, err := os.OpenFile(v.file, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0660)
	if err != nil {
		return fmt.Errorf("unable to open secrets file: %v", err)
	}
	defer f.Close()

	w, err := crypt.EncryptWriter(v.cryptKey, f)
	if err != nil {
		return err
	}

	err = json.NewEncoder(w).Encode(secrets)
	if err != nil {
		return fmt.Errorf("unable to encode secrets: %v", err)
	}
	return nil
}

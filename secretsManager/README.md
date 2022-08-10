# Secrets Manager
Secrets API and CLI exercise from https://gophercises.com/

Create a package that will store secrets (i.e. API keys) in an encryted file.  build such that it can be run as a CLI command to add/fetch/remove/modify secrets, or imported as a package into code to access.  the secrets file should be (optionally?) passed as a parameter, and encrypted, requring a encryption key to be passed to decrypt the data.


### DISCLAIMER
This is NOT intended Production ready code, use something like HashiCorp Vault instead (although it is probably better for personal projects then storing all secrets in a plain .env file....)


## V1
- read/write secrets to file stored as JSON format, without encryption.  place holders used for encryption steps to be implemented in v2
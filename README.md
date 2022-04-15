# lnpay
This is a service to use with lnbits to create payments by lightning.
## How to use
### setup and config
- Create a wallet in a lnbits instance. (e.g. the free https://legend.lnbits.com)
- Copy configs/example.config.yaml to project root as config.yaml or to /etc/lnpay/ or $HOME/.local/share/lnpay/
- Fill the config with right values, get invoice_key from lnbits api info section
### build and run
Run:
```bash
go build ./cmd/server
./server

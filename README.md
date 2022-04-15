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
```

### example
Request:
```bash
curl --location --request POST 'localhost:3000/pay' \
--header 'Content-Type: application/json' \
--data-raw '{
    "amount": 1,
    "description": "Package",
    "webhook": "https://webhook.site/"
}'
```
Response:
```bash
{
    "invoice": "lnbc10n1p39jfaepp5zzhr27mlrmstxr75c4dedl7qpg0k084hfksw9m89u5pmsdernfcqdqv2pskx6mpvajsxqyjw5qcqpjsp5ls5lps0vmcp6mqfnd0u3nnmlpppuzw296fc4nscn5nwl3vyr3whsrzjqftzw4d5r9nsau4nkakrxxdvkm0xgl6yxwuk4lp9yzkz5kql0j5vzzkcgvqq8tgqdqqqqqqqqqqqphgq9q9qy9qsqdygx56pyrksext4n7w8zht5s03qc5jsmajqm2q8tjmksp5q98h8yevh2mrpjhn056qctdv2wav57trdq8h7q43fjwjdvh80qqnxr9tgq5ktvck",
    "hash": "10ae357b7f1ae0b31fd4c55b96ffd00a1f679eb74da0e2ece5e503b837239a70"
}
```

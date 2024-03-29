# GOBANK 

using echo web frameworks, sample bank project.

> note. the balance type format follows eosio (one of blockchain platform)

## Main Dependencies
- [echo framework](https://github.com/labstack/echo)
- [xorm](https://github.com/go-xorm/xorm)
- [viper](https://github.com/spf13/viper)
- [logrus](https://github.com/sirupsen/logrus)  
- [datastore](https://cloud.google.com/go/datastore)

... and so on

## usage
```bash
gobank$ go build
```

## curl for APIs

### get Balance

for user id 123456

```bash
curl -X GET http://localhost:8080/api/v1/balance/123456
```

### modify Balance

for user id 123456

> for now, only implements deposit.

```bash
curl -X POST -H 'Content-Type: application/json' \
             -d '{ 
                    "diffBalance": { 
                        "fraction": 4, 
                        "symbol": "EOS", 
                        "amount": 1000, 
                        "strExpr": "0.1000 EOS" 
                    },
                    "action": "withdraw"
                }' \
             http://localhost:8080/api/v1/balance/123456
```

> action: deposit / withdraw

### delete Balance

for user id 123456

```bash
curl -X DELETE http://localhost:8080/api/v1/balance/123456
```
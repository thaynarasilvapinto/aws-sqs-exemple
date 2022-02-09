# AWS SQS EXEMPLE 

### build project 

```
cd src
```
```
go get
```
```
go buil
````

configure `launch.json`

```
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}/src"
        }
    ]
```
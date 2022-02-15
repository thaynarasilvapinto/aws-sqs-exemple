# AWS SQS EXEMPLE 

# Build project 

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

# Create queue in aws 

```
cd terraform
```

```
terraform init
```

```
terraform plan
```

```
terraform apply
```

OBS: importante lembrar que no arquivo ~./aws/credentials tem que ser suas keys da aws.
exemplo: 

```
[default]
aws_access_key_id = {key}
aws_secret_access_key = {secret_key}
```

# Delete queue in aws

```
cd terraform
```

```
terraform plan -destroy
```

```
terraform apply -destroy
```

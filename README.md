# simplelogin
An demonstration implementation of a simple login process

## Flyway
`$flyway -configFiles=migration/conf/flyway.conf -locations=filesystem:./migration/sql/ migrate`

create user
```
POST /create
{
	"username": "email@example.com",
	"password": "password"
}
```

login 
```
POST /create
{
	"username": "email@example.com",
	"password": "password"
}
```
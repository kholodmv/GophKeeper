# Client for the Gophkeeper project.
The client operates as a console utility, executing user commands

```
go run main.go
```
### Registration and authorization

Example registration command:

```
go run main.go auth register -u user -p password
```
Example of an authorization command:

```
go run main.go auth login -u user -p password
```

### Adding new data

Adding new data can be done using the command

```
>go run main.go keeper create 

Usage:
  client keeper [command]

Available Commands:
  create      add secret command
  get         get command

Flags:
  -b, --bin string        binary
  -k, --comment string    comment
  -c, --cvv string        cvv code
  -d, --date string       date
  -h, --help              help for keeper
  -n, --num string        card number
  -p, --password string   password
  -s, --string string     string
  -t, --title string      title
  -u, --username string   username

```
Example of adding a text entry:
```
go run main.go keeper create -t "my text name" -s "custom text" -k comment
```
Example of adding an account:
```
go run main.go keeper create -t google.com -u username -p password 
```
Example of adding a bank card:
```
go run main.go keeper create -t BankName1 -n "1234 5678 9876 5432" -d 12/26 -c 111
```
An example of adding an arbitrary file (read and stored as binary data):
```
go run main.go keeper create -t "file" -b temp/1.jpg
```

### Receiving data

Retrieving data is done using the get command
```
go run main.go keeper get -t "my secret name" 
```
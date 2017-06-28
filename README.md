# cf-mysql-example-app

Test app for validating MySQL service connectivity. Definitely not for anything other than simple smoke testing.

## Usage

```
$ curl -X PUT -d 'some-data' https://cf-mysql-example-app.cf.com/key-name
created

$ curl 'some-data' https://cf-mysql-example-app.cf.com/key-name
some-data
```

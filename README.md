# go-registry

A very simple substitution to [BOSH registry](https://github.com/cloudfoundry/bosh/tree/master/bosh-registry). 

# Warning

At this stage this Registry server has only "in-memory" storage.

It is recomended to use it only for testing and proofs of concepts.


### Usage

Config file format:
```
username: admin
password: admin
host: 127.0.0.1
port: 8000
```

Run registry with following commands:
```
go-registry -c config.yml
```

How to access registry API ? 
```
curl -X PUT -u admin:admin --data '{"foo": "bar"}' 127.0.0.1:8000/instances/1/settings
curl -X GET -u admin:admin 127.0.0.1:8000/instances/1/settings
# => {"settings":"{\"foo\": \"bar\"}","status":"ok"}
```

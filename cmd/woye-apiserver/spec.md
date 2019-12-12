# struct

## spec yaml
```
name: testcontainer
spec:
 	container:
		name: testcontainer 
    distri: debian
   	release:  buster
   	arch: amd64

```

## postdata struct
```
type Postdata struct {
	Endpointurl		string
	Commandtype		string
	Spec					[]byte
}

```

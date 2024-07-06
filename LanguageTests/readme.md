# Things I learned regarding docker:
In order to have portable developer environment you just need to run docker image with binded volume so your image can read updated files.


Sample:
```shell
docker run -v ".":/src -it hello-go
```
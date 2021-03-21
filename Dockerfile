##go base image @latest
FROM golang:alpine



##my work directory 
WORKDIR /rest-api-gin-gorm-sql


## copy all the dependency 
COPY . .

##download all the dependency
RUN go mod download 

##test
RUN go get github.com/githubnemo/CompileDaemon


##building the app 
RUN go build

## run the cmd
##check the name of the directory and the entry command point 
ENTRYPOINT ["CompileDaemon", "--directory=/rest-api-gin-gorm-sql",  "--command=/rest-api-gin-gorm-sql/rest-api-test"]
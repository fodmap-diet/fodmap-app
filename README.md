# fodmap-app
backend for fodmap api

### Getting Started
clone the repo
```
git clone https://github.com/fodmap-diet/fodmap-app.git
```
run the server
```
cd fodmap-app
go run app.go
```
test from another cmd prompt
```
curl http://localhost:8080/search/?item=mango
```

### Deploy with GCP app engine
install google cloud sdk from   
> [Refer here](https://cloud.google.com/sdk/docs/)    
install app engine golang component   
```
gcloud components install app-engine-go
```
create app
```
gcloud app create --project=[YOUR_PROJECT_NAME]
```
run locally
```
dev_appserver.py app.yaml
```
test from another cmd prompt
```
curl http://localhost:8080/search/?item=mango
```
Deploy with
```
gcloud app deploy
```

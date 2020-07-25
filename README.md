# About 
[gcp-kit/line-bot-boilerplate-go](https://github.com/gcp-kit/line-bot-boilerplate-go)

# Usage
```shell script
git clone https://github.com/gcp-kit/line-bot-boilerplate-go-example.git
cd line-bot-boilerplate-go-example
go get
```

`.env.yaml.tpl` を `.env.yaml` にしてyaml内の値を整える  

## Local
```shell script
cd cmd
go run main.go
```

## GCP(Cloud Functions)
Cloud Pub/SubのTopicを作っておく(2つ)  
### Deploy
```shell script
gcloud builds submit --config=cloudbuild.yaml .
```

# License
[MIT license](https://en.wikipedia.org/wiki/MIT_License).

mkdir tmp
swagger-codegen generate -l go -i http://localhost:8080/api-docs -o tmp -t ./codegen/ -DpackageName=model
mkdir -p tmp/model
mkdir -p tmp/client
mv tmp/model_* tmp/model/
mv tmp/api_* tmp/client/
cd tmp/model
rename -d model_ *

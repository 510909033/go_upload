APP_NAME=go_upload
test:
	echo $(APP_NAME)
build:
	echo "env=${go_env_file}"
	go mod tidy
dev: build
	go build --ldflags "-extldflags -static" -o go_upload_dev cmd/main/main.go
	nohup ./go_upload_dev > /tmp/go_upload_dev 2>&1 &
online: run
# 自动提示
# 要在vim中自动提示，请先运行make autocompletor
autocompletor:
	go install ./pkg/...

stop:
	ps aux|grep "go_upload_dev"|grep -v grep|awk '{print $$2}'|xargs kill
# 开发环境运行
run: build
	go build --ldflags "-extldflags -static" -o go_upload cmd/main/main.go
	rm -rf /var/www/go_upload
	mkdir /var/www/go_upload
	cp go_upload /var/www/go_upload/
	cp -r config /var/www/go_upload/
	cp -r docker /var/www/go_upload/
	cp -r public /var/www/go_upload/
	cp -r template /var/www/go_upload/
	nohup docker-compose -f ./docker/docker-compose.yml up > /dev/null 2>&1 &

goget:
	go get github.com/510909033/bgf_bo@master
	go get github.com/510909033/bgf_database@master
	go get github.com/510909033/bgf_log@master
	go get github.com/510909033/bgf_mvc@master
	go get github.com/510909033/bgf_util@master

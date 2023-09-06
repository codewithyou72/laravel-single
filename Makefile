# 定制 goctl 名称
GO_CTL_NAME=laravel-cli

# web 服务名称(api 和 rpc 都依赖这个)
SERVICE_WEB=laravel-single

# go-zero 生成代码风格
GO_ZERO_STYLE=goZero

GO ?= go
GOFMT ?= gofmt "-s"
GOFILES := $(shell find . -name "*.go")
LDFLAGS := -s -w

.PHONY: test
test: # 运行项目测试
	go test -v --cover ./app/internal/..

.PHONY: fmt
fmt: # 格式化代码
	$(GOFMT) -w $(GOFILES)

.PHONY: gen-code
gen-code: # 生成|新增接口代码
	$(GO_CTL_NAME) api go -api gateway/api/$(SERVICE_WEB).api -dir gateway/$(SERVICE_WEB)  --style=$(GO_ZERO_STYLE)
	@echo "Generate $(SERVICE_WEB)-api files successfully"

.PHONY: gen-swagger
gen-swagger: # 根据.api文档生成swagger 文档
	$(GO_CTL_NAME) api plugin -plugin goctl-swagger="swagger -filename $(SERVICE_WEB)-swagger.json" -api ./app/routes/web.api -dir ./app/public
	@echo "Generate $(SERVICE_WEB) swagger files successfully"

.PHONY: gen-kube
gen-kube: # 生成 k8s部署文件
	$(GO_CTL_NAME) kube deploy -secret docker-login -replicas 3 -requestCpu 50 -requestMem 150 -limitCpu 300 -limitMem 150 -name mvp-web-gateway-api -namespace laravel-single -image laravel-single:1.0.0 -o app/index-kube.yaml -port 1001 -serviceAccount find-endpoints -minReplicas 3 -maxReplicas 100
	@echo "Generate k8s config files successfully"

.PHONY: help
help: # 显示帮助
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done

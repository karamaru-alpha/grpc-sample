ENV_LOCAL_FILE = env.local
ENV_LOCAL = $(shell cat $(ENV_LOCAL_FILE))


# 8080番Portでapiサーバを起動する
.PHONY: serve
serve:
	$(ENV_LOCAL) docker-compose up

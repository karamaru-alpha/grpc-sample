ENV_LOCAL_FILE = env.local
ENV_LOCAL = $(shell cat $(ENV_LOCAL_FILE))


.PHONY: serve
serve:
	$(ENV_LOCAL) docker-compose up

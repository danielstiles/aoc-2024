YEAR?=2024
DAY?='1'
padded_day=$(shell printf '%02d' $(DAY))
GO_VERSION?=1.23.3

new_day:
	mkdir -p $(padded_day)
	cp -R template/* $(padded_day)/
	git config --get remote.origin.url | sed 's/^.*\(github.*\).git/module \1\/$(padded_day)\n\ngo $(GO_VERSION)\n/' > $(padded_day)/go.mod
	curl https://adventofcode.com/$(YEAR)/day/$(DAY)/input -H "Cookie: $(shell cat cookie)" > $(padded_day)/input.txt

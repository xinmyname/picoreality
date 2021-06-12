#!/bin/sh
cd ./cmd
go build
cd ..
./cmd/cartgen ./blueprint.json > ~/Library/Application\ Support/pico-8/carts/picoreality.p8

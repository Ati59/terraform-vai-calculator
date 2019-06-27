#!/usr/bin/env bash

gox -os="linux darwin windows" -arch="386 amd64" .

mv terraform-vai-calculator_* bin/

exit 0

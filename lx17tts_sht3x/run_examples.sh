#!/bin/bash

cd golang1_9 && go run ./main.go
cd .. && cd php7_1 && php -f ./example.php
cd .. && cd python3_7 && python3 ./example.py

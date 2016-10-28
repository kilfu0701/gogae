#!/bin/bash

dev_appserver.py --port=8081 --admin_port=8001 --datastore_path="/tmp/myapp_datastore" . 2>&1 | \
    gawk '{ gsub("DEBUG", "\033[1;35m&\033[0m");
            gsub("INFO", "\033[1;36m&\033[0m");
            gsub("ERROR", "\033[1;31m&\033[0m");
            gsub("WARNING", "\033[1;33m&\033[0m");
            gsub("CRITICAL", "\033[1;31m&\033[0m");
            print
    }'

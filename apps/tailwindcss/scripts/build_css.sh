#!/bin/bash

. ./scripts/build_mode.sh

if [ "$MINIFY" = "true" ]; then
    npx tailwindcss -c src/tailwindcss/tailwind.config.js -i src/tailwindcss/styles.css -o ${BUILD_DIR}/static/styles.css --minify
    echo -e "\033[34mBUILD_CSS:\033[32m Built to \033[33m${BUILD_DIR}/static/style.css \033[35m(Minified)\033[0m"
else
    npx tailwindcss -c src/tailwindcss/tailwind.config.js -i src/tailwindcss/styles.css -o ${BUILD_DIR}/static/styles.css
    echo -e "\033[34mBUILD_CSS:\033[32m Built to \033[33m${BUILD_DIR}/static/style.css\033[0m"
fi



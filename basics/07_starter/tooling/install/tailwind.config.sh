#!/bin/bash

FILE="tailwind.config.js"

# Update content to pick up on our .templ files
sed -i 's/content: \[\]/content: \["src\/\*\*\/\*.templ"\]/' "$FILE"

# Update plugins to use typography and daisyui
sed -i 's/plugins: \[\]/plugins: \[require\("@tailwindcss\/typography"\), require\("daisyui"\)\]/' "$FILE"


#!/bin/bash

FILE=".air.toml"

# Update build command
FROM='cmd = "go build -o .\/tmp\/main ."'
TO='cmd = "go build -o .\/tmp\/main .\/src\/"'
sed -i "s|$FROM|$TO|" "$FILE"

# Update exclude directories
FROM='exclude_dir = \["assets", "tmp", "vendor", "testdata"\]'
TO='exclude_dir = \["assets", "tmp", "vendor", "testdata", "node_modules", "dist", "tooling"\]'
sed -i "s|$FROM|$TO|" "$FILE"

# clean on exit
FROM='clean_on_exit = false'
TO='clean_on_exit = true'
sed -i "s|$FROM|$TO|" "$FILE"

#!/bin/bash

FILE=".air.toml"

# Update build command linux
FROM='cmd = "go build -o .\/tmp\/main ."'
TO='cmd = "templ generate \&\& go build -o .\/tmp\/main .\/src\/"'
sed -i "s|$FROM|$TO|" "$FILE"

# Update build command windows
FROM='cmd = "go build -o .\/tmp\/main.exe ."'
TO='cmd = "templ generate \&\& go build -o .\/tmp\/main.exe .\/src\/"'
sed -i "s|$FROM|$TO|" "$FILE"

# Update include extensions
FROM='include_ext = \["go", "tpl", "tmpl", "html"\]'
TO='include_ext = \["go", "tpl", "tmpl", "templ", "html"\]'
sed -i "s|$FROM|$TO|" "$FILE"

# Update exclude directories
FROM='exclude_dir = \["assets", "tmp", "vendor", "testdata"\]'
TO='exclude_dir = \["assets", "tmp", "vendor", "testdata", "node_modules", "dist", "tooling"\]'
sed -i "s|$FROM|$TO|" "$FILE"

# Update exclude files
FROM='exclude_regex = \["_test.go"\]'
TO='exclude_regex = \["_test.go", ".*_templ.go"\]'
sed -i "s|$FROM|$TO|" "$FILE"

# clean on exit
FROM='clean_on_exit = false'
TO='clean_on_exit = true'
sed -i "s|$FROM|$TO|" "$FILE"
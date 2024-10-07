#!/bin/bash

# Ensure MODE is set, default to "build" if not
: "${MODE:=build}"

# Determine BUILD_DIR based on MODE
if [ "$MODE" = "live" ]; then
    BUILD_DIR="tmp"
else
    BUILD_DIR="build"
fi

# If minify is already false or mode is live, minify = false otherwise minify
if [ "$MODE" = "live" ] || [ "${MINIFY:-true}" = "false" ]; then
    MINIFY="false"
else
    MINIFY="true"
fi

export BUILD_DIR
export MODE
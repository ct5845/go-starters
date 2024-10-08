#!/bin/bash

# env.sh
# attempts to all relevant .env files and their variables and export them so they can be used in the build or runtime arguements

# Path for .env files
ENV_DIR="./config"

# Check if APP_ENV is set; if not, set it to development
: "${APP_ENV:=dev}"

# Create an empty variable to hold the names of the files to load
FILES_TO_LOAD=()

# Function to load and export variables from a file
load_env_file() {
  local file="$1"
  if [ -f "$file" ]; then
    echo "Loading $file..."
    set -a
    source "$file"
    set +a
    FILES_TO_LOAD+=("$file")  # Add to the loaded files array
  else
    echo "No specific $file file. Continuing with defaults."
  fi
}

# Load the .env files
load_env_file "$ENV_DIR/.env" # load .env first
load_env_file "$ENV_DIR/.env.local" # .env.local takes precedence over .env
load_env_file "$ENV_DIR/.env.${APP_ENV}" # .env.[APP_ENV] takes precdence over .env.local
load_env_file "$ENV_DIR/.env.${APP_ENV}.local" # .env.[APP_ENV].local takes precdence over .env.[APP_ENV]

# Export only the variables defined in the loaded files
for file in "${FILES_TO_LOAD[@]}"; do
  while IFS='=' read -r key value; do
    if [[ ! "$key" =~ ^# && -n "$key" ]]; then  # Ignore comments and empty keys
      export "$key"="$value"
    fi
  done < "$file"
done
#!/bin/bash
# Script to setup error reporting integration with Sentry

SENTRY_DSN="your_sentry_dsn_here"

# Export SENTRY_DSN as environment variable for your application
echo "export SENTRY_DSN=$SENTRY_DSN" >> ~/.bashrc
source ~/.bashrc

echo "Sentry DSN set. Please configure your application to use this environment variable for error reporting."

#!/bin/bash

SOURCE="${BASH_SOURCE[0]}"
REPOSITORY=$(dirname $SOURCE)/..
OUTPUT_FILE="$REPOSITORY/release.yml"

echo > $OUTPUT_FILE

# BuildTemplates

cat $REPOSITORY/config/namespaces.yml >> $OUTPUT_FILE
cat $REPOSITORY/config/buildpack-bits.yml >> $OUTPUT_FILE
cat $REPOSITORY/config/kaniko.yml >> $OUTPUT_FILE

# KleinCF components

cat $REPOSITORY/config/kleincf.yml >> $OUTPUT_FILE

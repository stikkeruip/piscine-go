#!/bin/bash

# Download the JSON file and extract the required fields in one step
curl -s "https://platform.zone01.gr/assets/superhero/all.json" | \
jq -r '.[] | select(.id == 170) | "\(.name)\n\(.powerstats.power)\n\(.appearance.gender)"'


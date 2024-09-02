#!/bin/bash

# Download the JSON file
curl -o superheroes.json "https://platform.zone01.gr/assets/superhero/all.json"

# Extract and print Name, Power, and Gender of id 170
jq -r '.[] | select(.id == 170) | "\(.name)\n\(.powerstats.power)\n\(.appearance.gender)"' superheroes.json

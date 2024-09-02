curl -s "https://platform.zone01.gr/assets/superhero/all.json" | \
jq -r '.[] | select(.id == $HERO_ID) | "\(.name)\n\(.powerstats.power)\n\(.appearance.gender)"'

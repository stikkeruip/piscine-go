curl -s "https://platform.zone01.gr/assets/superhero/all.json" | \
jq -r --arg hero_id "$HERO_ID" '.[] | select(.id == $HERO_ID) | "\(.name)\n\(.powerstats.power)\n\(.appearance.gender)"'

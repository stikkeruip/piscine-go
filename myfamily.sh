curl -s "https://platform.zone01.gr/assets/superhero/all.json" | \
jq -r --arg HERO_ID "$HERO_ID" '.[] | select(.id == ($HERO_ID | tonumber)) | .connections.relatives' | awk 'ORS="; "' | sed 's/; $//'

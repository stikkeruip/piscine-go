find . -type f -name '*sh' | rev | cut -d'.' -f2 | rev | sed 's/^\///'


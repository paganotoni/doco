# This script starts the tailwind watcher 
./bin/tailwindcss -i ./web/tailwind.css -o ./web/templates/styles.css --content=./web/templates/** &&\
cp ./web/templates/styles.css ./public/styles.css
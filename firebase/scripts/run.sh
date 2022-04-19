#!/bin/sh

# Script to download firestore backup data and run emulator 

if ! command -v gsutil &> /dev/null
then
    echo "Error: gsutils could not be found."
    echo "Please refer to https://cloud.google.com/sdk/docs/install"
    exit
fi

BUCKET="gs://stalwart-cable-309519-firestore-backup"
BACKUP="2022-03-06T22:09:03_17229"

pushd testdata
if [ ! -d $BACKUP ] 
then
    echo "Download $BACKUP" 
    gsutil -m cp -r $BUCKET/$BACKUP .
fi
popd

pushd functions
npm install
npm run build
popd

firebase emulators:start --import testdata/$BACKUP

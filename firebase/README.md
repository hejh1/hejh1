# Firebase Backend

## Build existing cloud functions
The functions are configured as TS project. Please run below to build the functions

```
cd functions
npm install
npm run build
```

&nbsp;

## Launch emulator
After building cloud functions, you can launch emulator with

`firebase emulators:start`

To run script with emulator from new shell instances, please export below environment variable

`export FIRESTORE_EMULATOR_HOST="localhost:8080"`

If you want to download Firestore backup data as mock data, please in `firebase` folder run

`./scripts/run.sh`

You can also modify the `BACKUP` and `BUCKET` variable in `run.sh` to change mock data source

&nbsp;

## Add new cloud functions
Please add TS cloud function into `functions/src` folder, and add JS cloud function into `functions/src/js` folder.

Then modify `functions/src/index.ts` to export the function. 

Note: please rum `npm run lint` before check-in code.

&nbsp;

## Modify Firebase config
Please refer to below files for individual configurations

Overall Rules: `firebase.json`

Realtime Database Security Rules: `database.rules.json`

Firestore Rules: `firestore.rules`

Firestore indexes: `firestore.indexes.json`

Storage Rules: `storage.rules`

Remote Config template: `remoteconfig.template.json`

&nbsp;

## Deploy cloud functions
You can deploy all functions as a whole with 

`firebase deploy --only functions`

Or deploy individual function by commenting out portion of `functions/src/index.ts`

&nbsp;

## Deploy firestore cloud rules
You can update and deploy firestore rules by

`firebase deploy --only firestore:rules`

&nbsp;

## Deploy whole firebase project
You can deploy the whole firebase project with 
`firebase deploy`

Please consider the consequences before doing this.

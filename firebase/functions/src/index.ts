import * as admin from "firebase-admin";

export * from "./js/firestore_backup";
export * from "./js/scorecard";
export * from "./js/threads";
export * from "./users";

// Initializes the SDK
admin.initializeApp();

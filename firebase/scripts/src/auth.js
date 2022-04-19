const admin = require("firebase-admin");

exports.accessFIRAdmin = (admin) => {
  var serviceAccount = require("./stalwart-cable-309519-firebase-adminsdk-tnvhu-d72b853072.json");

  admin.initializeApp({
    credential: admin.credential.cert(serviceAccount),
    databaseURL: "https://stalwart-cable-309519-default-rtdb.firebaseio.com",
  });

  return {
    db: admin.firestore(),
  };
};

const { db } = this.accessFIRAdmin(admin);
exports.db = db;
exports.admin = admin;

exports.FIRRef = (path, is_collection) => {
  //console.log("db: ", db);
  if (is_collection) {
    return db.collection(path);
  } else {
    return db.doc(path);
  }
};

exports.memoriesForUser = (userid) => {
  return {
    memoriesTestRef: this.FIRRef(`MemoriesTest/${userid}`, false),
    timelineTestRef: this.FIRRef(`MemoriesTest/${userid}/TimelineTest`, true),
    thoughtTestRef: this.FIRRef("ThoughtsTest", true),
    memoriesProdRef: this.FIRRef(`MemoriesProdV0/${userid}`, false),
    timelineProdRef: this.FIRRef(`MemoriesProdV0/${userid}/Timeline`, true),
    thoughtProdRef: this.FIRRef(`MemoriesProdV0/${userid}/Thoughts`, true),
  };
};

exports.refForMemories = (version) => {
  const memCollectionName =
    version === "test" ? "MemoriesTest" : "MemoriesProdV0";
  return this.FIRRef(memCollectionName, true);
};

exports.refForUserMemories = (userid, version) => {
  const { memoriesTestRef, memoriesProdRef } = this.memoriesForUser(userid);
  const ref = version === "test" ? memoriesTestRef : memoriesProdRef;
  return ref;
};

exports.refForTimeline = (userid, version) => {
  const { timelineTestRef, timelineProdRef } = this.memoriesForUser(userid);
  const ref = version === "test" ? timelineTestRef : timelineProdRef;
  return ref;
};

exports.refForThought = (userid, version) => {
  const { thoughtTestRef, thoughtProdRef } = this.memoriesForUser(userid);
  const ref = version === "test" ? thoughtTestRef : thoughtProdRef;
  return ref;
};

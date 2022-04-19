// The Cloud Functions for Firebase SDK to create functions and setup triggers.
const functions = require("firebase-functions");

// The Firebase Admin SDK to access Firestore.
const admin = require("firebase-admin");

// There are many limitations of Firebase Cloud Functions and Firestore.
// To calculate user's total thoughts and streak accurately
// while minimizing the cost, we have to use a combination of regular
// re-aggregation of all records (scheduled) and incremental updates
// (triggered by updates):
//   1. scanning all records frequently is expensive;
//   2. update triggers have limited, which can send the same trigger more
//      than once.

// This function is to refresh scorecards for users who are not actively
// creating or deleting thoughts, which won't trigger the incremental updates.
exports.scheduledRefreshScorecards = functions.pubsub
    .schedule("every 30 mins")
    .onRun((context) => {
      const memoryRef = admin
          .firestore()
          .collection("MemoriesTest")
          .orderBy("lasttimeUpdated")
          .limit(1);
      return memoryRef
          .get()
          .then((querySnapshot) => {
            querySnapshot.forEach((doc) => {
              // doc.data() is never undefined for query doc snapshots
              functions.logger.info(
                  "Scheduled refresh of scorecards for inactive users ",
                  doc.data()
              );
              refreshScorecard(doc.data().uid);
            });
          })
          .catch((error) => {
            functions.logger.error(
                "Error reading memories for scheduled refreshing ",
                error
            );
          });
    });

// This function is to refresh scorecards for users who are not actively
// creating or deleting thoughts, which won't trigger the incremental updates.
exports.scheduledScoreProd = functions.pubsub
    .schedule("every 30 mins")
    .onRun((context) => {
      const memoryRef = admin
          .firestore()
          .collection("MemoriesProdV0")
          .orderBy("lasttimeUpdated")
          .limit(1);
      return memoryRef
          .get()
          .then((querySnapshot) => {
            querySnapshot.forEach((doc) => {
              // doc.data() is never undefined for query doc snapshots
              functions.logger.info(
                  "Scheduled refresh of scorecards for inactive users ",
                  doc.data()
              );
              rescoreProd(doc.data().uid);
            });
          })
          .catch((error) => {
            functions.logger.error(
                "Error reading memories for scheduled refreshing ",
                error
            );
          });
    });

exports.refreshScorecardOnCreate = functions.firestore
    .document("/MemoriesTest/{userId}/TimelineTest/{sessionId}")
    .onCreate((snap, context) => {
      functions.logger.debug(
          "Refreshing scorecard for user on create: ",
          context.params.userId
      );
      const memoryRef = admin
          .firestore()
          .collection("MemoriesTest")
          .doc(context.params.userId);
      return memoryRef
          .get()
          .then((doc) => {
            const total = doc.data().totalThoughts;
            let streak = doc.data().createStreak;
            const latestDate = doc.data().latestThoughtTime;
            if (!total || !streak || !latestDate) {
              return refreshScorecard(context.params.userId);
            } else {
              const dateDiff = Math.floor(
                  (snap.data().creationTime.toDate().setHours(0, 0, 0, 0) -
              latestDate.toDate().setHours(0, 0, 0, 0)) /
              (1000 * 60 * 60 * 24)
              );
              if (dateDiff == 1) {
                streak = streak + 1;
              } else if (dateDiff > 1) {
                streak = 1;
              }
              return doc.ref.set(
                  {
                    totalThoughts: total + 1,
                    createStreak: streak,
                    // eslint-disable-next-line max-len
                    lasttimeUpdated: admin.firestore.FieldValue.serverTimestamp(),
                    latestThoughtTime: snap.data().creationTime,
                  },
                  {merge: true}
              );
            }
          })
          .catch((error) => {
            functions.logger.error(
                "Error reading the newly created document:",
                error
            );
            return null;
          });
    });

exports.refreshScorecardOnDelete = functions.firestore
    .document("/MemoriesTest/{userId}/TimelineTest/{sessionId}")
    .onDelete((snap, context) => {
      functions.logger.debug(
          "Refreshing scorecard for user on delete: ",
          context.params.userId
      );
      const memoryRef = admin
          .firestore()
          .collection("MemoriesTest")
          .doc(context.params.userId);
      return memoryRef
          .get()
          .then((doc) => {
            const total = doc.data().totalThoughts;
            if (!total) {
              return refreshScorecard(context.params.userId);
            } else {
              return doc.ref.set(
                  {
                    totalThoughts: total - 1,
                    // eslint-disable-next-line max-len
                    lasttimeUpdated: admin.firestore.FieldValue.serverTimestamp(),
                  },
                  {merge: true}
              );
            }
          })
          .catch((error) => {
            functions.logger.error(
                "Error reading the newly deleted document:",
                error
            );
            return null;
          });
    });

exports.scoreOnCreateProd = functions.firestore
    .document("/MemoriesProdV0/{userId}/Timeline/{sessionId}")
    .onCreate((snap, context) => {
      functions.logger.debug(
          "Refreshing scorecard on Prod store for user on create: ",
          context.params.userId
      );
      const memoryRef = admin
          .firestore()
          .collection("MemoriesProdV0")
          .doc(context.params.userId);
      return memoryRef
          .get()
          .then((doc) => {
            const total = doc.data().totalThoughts;
            let streak = doc.data().createStreak;
            const latestDate = doc.data().latestThoughtTime;
            if (!total || !streak || !latestDate) {
              return rescoreProd(context.params.userId);
            } else {
              const dateDiff = Math.floor(
                  (snap.data().creationTime.toDate().setHours(0, 0, 0, 0) -
              latestDate.toDate().setHours(0, 0, 0, 0)) /
              (1000 * 60 * 60 * 24)
              );
              if (dateDiff == 1) {
                streak = streak + 1;
              } else if (dateDiff > 1) {
                streak = 1;
              }
              return doc.ref.set(
                  {
                    totalThoughts: total + 1,
                    createStreak: streak,
                    // eslint-disable-next-line max-len
                    lasttimeUpdated: admin.firestore.FieldValue.serverTimestamp(),
                    latestThoughtTime: snap.data().creationTime,
                  },
                  {merge: true}
              );
            }
          })
          .catch((error) => {
            functions.logger.error(
                "Error reading the newly created document:",
                error
            );
            return null;
          });
    });

exports.scoreOnDeleteProd = functions.firestore
    .document("/MemoriesProdV0/{userId}/Timeline/{sessionId}")
    .onDelete((snap, context) => {
      functions.logger.debug(
          "Refreshing scorecard on Prod store for user on delete: ",
          context.params.userId
      );
      const memoryRef = admin
          .firestore()
          .collection("MemoriesProdV0")
          .doc(context.params.userId);
      return memoryRef
          .get()
          .then((doc) => {
            const total = doc.data().totalThoughts;
            if (!total) {
              return rescoreProd(context.params.userId);
            } else {
              return doc.ref.set(
                  {
                    totalThoughts: total - 1,
                    // eslint-disable-next-line max-len
                    lasttimeUpdated: admin.firestore.FieldValue.serverTimestamp(),
                  },
                  {merge: true}
              );
            }
          })
          .catch((error) => {
            functions.logger.error(
                "Error reading the newly deleted document:",
                error
            );
            return null;
          });
    });

/**
 * Async function to recompute the scorecard for user, specified by userId, by
 *  scanning all its records.
 *  @param {string} userId the userIf for the memory
 *  @return {Promise} to update scorecard for the userId.
 */
async function refreshScorecard(userId) {
  functions.logger.info("Refreshing scorecard for user ", userId);
  const timelineRef = admin
      .firestore()
      .collection("MemoriesTest")
      .doc(userId)
      .collection("TimelineTest")
      .orderBy("creationTime", "desc");
  let count = 0;
  let streak = 0;
  let breakStreak = false;
  let nextPostDate = new Date();
  let latestThoughtTime = null;
  timelineRef
      .stream()
      .on("data", (snap) => {
        ++count;
        if (count == 1) {
          latestThoughtTime = snap.data().creationTime;
        }
        if (!breakStreak) {
          const dateDiff = Math.floor(
              (nextPostDate.setHours(0, 0, 0, 0) -
            snap.data().creationTime.toDate().setHours(0, 0, 0, 0)) /
            (1000 * 60 * 60 * 24)
          );
          if (dateDiff > 1) {
            breakStreak = true;
          } else if (dateDiff <= 1) {
            if (streak == 0) {
              streak = 1;
            }
            if (dateDiff == 1) {
              streak = streak + 1;
            }
            nextPostDate = snap.data().creationTime.toDate();
          }
        }
      })
      .on("end", () => {
        return admin.firestore().collection("MemoriesTest").doc(userId).set(
            {
              totalThoughts: count,
              createStreak: streak,
              latestThoughtTime: latestThoughtTime,
              lasttimeUpdated: admin.firestore.FieldValue.serverTimestamp(),
            },
            {merge: true}
        );
      });
}

/**
 * Async function to recompute the scorecard for user, specified by userId, by
 *  scanning all its records.
 *  @param {string} userId the userIf for the memory
 *  @return {Promise} to update scorecard for the userId.
 */
async function rescoreProd(userId) {
  functions.logger.info("Refreshing scorecard for user ", userId);
  const timelineRef = admin
      .firestore()
      .collection("MemoriesProdV0")
      .doc(userId)
      .collection("Timeline")
      .orderBy("creationTime", "desc");
  let count = 0;
  let streak = 0;
  let breakStreak = false;
  let nextPostDate = new Date();
  let latestThoughtTime = null;
  timelineRef
      .stream()
      .on("data", (snap) => {
        ++count;
        if (count == 1) {
          latestThoughtTime = snap.data().creationTime;
        }
        if (!breakStreak) {
          const dateDiff = Math.floor(
              (nextPostDate.setHours(0, 0, 0, 0) -
            snap.data().creationTime.toDate().setHours(0, 0, 0, 0)) /
            (1000 * 60 * 60 * 24)
          );
          if (dateDiff > 1) {
            breakStreak = true;
          } else if (dateDiff <= 1) {
            if (streak == 0) {
              streak = 1;
            }
            if (dateDiff == 1) {
              streak = streak + 1;
            }
            nextPostDate = snap.data().creationTime.toDate();
          }
        }
      })
      .on("end", () => {
        return admin.firestore().collection("MemoriesProdV0").doc(userId).set(
            {
              totalThoughts: count,
              createStreak: streak,
              latestThoughtTime: latestThoughtTime,
              lasttimeUpdated: admin.firestore.FieldValue.serverTimestamp(),
            },
            {merge: true}
        );
      });
}

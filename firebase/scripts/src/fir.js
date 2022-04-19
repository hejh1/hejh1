const auth = require("./auth.js");
const chalk = require("chalk");
const { CollectionReference } = require("@google-cloud/firestore");

/**
    Print all the docs in the collection ref. Mostly for debugging and command purpose.
    @param {CollectionReference} ref a Firestore collection reference.
*/
exports.printCollections = async (ref) => {
  const snapshot = await ref.get();
  snapshot.forEach((doc) => {
    // doc.data() is never undefined for query doc snapshots
    console.log("doc ", chalk.blue(doc.id), " => ", doc.data());
  });
};

/// Stats

/**
 * Count the documents in a collection - only working for small collections that can be read once.
 * @param {CollectionReference} ref a Firesotre collection reference.
 * @returns {Promise} the total number of documents.
 */
exports.countSmallCollections = async (ref) => {
  const snapshot = await ref.get();
  const size = snapshot.size;
  return size;
};

/**
 * Count the documents in a collection that passes certain check:
 *   1. only works for small collection that can be read once into memory;
 *   2. check has to be synchronized function;
 * @param {CollectionReference} ref a Firestore collection reference.
 * @param {function} check a sync func that returns true or false for each doc.
 * @returns {int} the number of documents passing the check.
 */
exports.countSmallCollectionsForCondition = async (ref, check) => {
  const snapshot = await ref.get();
  let size = 0;
  snapshot.forEach((doc) => {
    if (check(doc)) {
      ++size;
    }
  });
  return size;
};

/**
 * Count the documents in a collection that can be looked up in another collection:
 *  1. only works for small collection that can be read once into memory.
 *  2. the lookup is an async operation that returns the Promise of document.
 * @param {CollectionReference} ref a Firestore collection reference.
 * @param {function} lookup an async func that returns a promise of found document.
 * @param {CollectionReference} xref a Firestore collection referenc to look against.
 * @returns {int} the number of documents that can be found in another collection.
 */
exports.countSmallCollectionsForAsyncLookup = async (ref, lookup, xref) => {
  const lookups = [];
  const snapshot = await ref.get();
  snapshot.forEach((doc) => {
    lookups.push(lookup(xref, doc));
  });
  return Promise.all(lookups).then((results) => {
    let size = 0;
    results.forEach((result) => {
      let count = 0;
      result.forEach((doc) => {
        if (doc.data().serverId && doc.data().serverExportedMediaPath) {
          ++count;
        }
      });
      if (count == 1) {
        ++size;
      }
    });
    return size;
  });
};

exports.lookupBySessionId = async (ref, doc) => {
  const data = ref.where("sessionId", "==", doc.id).get();
  return data;
};

exports.migrateThought = async (doc, refs) => {
  const { timelineProdRef, thoughtTestRef, thoughtProdRef } = refs;
  const thought = await thoughtTestRef.doc(doc.data().serverId).get();
  if (!thought.exists) {
    console.log(
      "Failed to find thought object for doc: ",
      chalk.yellow(doc.id)
    );
    return false;
  } else if (!thought.data().serverExportedMediaPath) {
    console.log(
      "Failed to migrage thought for missing exported media: ",
      chalk.yellow(doc.id)
    );
    return false;
  } else {
    await timelineProdRef.doc(doc.id).set(doc.data());
    await thoughtProdRef.doc(doc.id).set(thought.data());
    return true;
  }
};

exports.migrateUser = async (userid) => {
  const tlRef = auth.refForTimeline(userid, "test");
  const refs = auth.memoriesForUser(userid);
  const snapshot = await tlRef.get();
  snapshot.forEach((doc) => {
    if (doc.data().serverId) {
      this.migrateThought(doc, refs).then((res) => {
        console.log(
          "migraget success: ",
          chalk.red(res),
          " for doc ",
          chalk.blue(doc.id)
        );
      });
    } else {
      console.log("doc ", chalk.yellow(doc.id), " misses server id.");
    }
  });
};

/**
 * Async function to recompute the scorecard for user, specified by userId, by
 *  scanning all its records.
 *  @param {string} userId the userIf for the memory
 *  @return {Promise} to update scorecard for the userId.
 */
exports.scorecard = async (userid, version) => {
  const tlRef = auth.refForTimeline(userid, version);
  const timelineRef = tlRef.orderBy("creationTime", "desc");
  const memRef = auth.refForUserMemories(userid, version);
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
      return memRef.set(
        {
          totalThoughts: count,
          createStreak: streak,
          latestThoughtTime: latestThoughtTime,
          lasttimeUpdated: auth.admin.firestore.FieldValue.serverTimestamp(),
        },
        { merge: true }
      );
    });
};

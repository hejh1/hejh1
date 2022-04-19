const functions = require("firebase-functions");
const admin = require("firebase-admin");

/**
 * Triggers when a ThreadVideo is added to a Thread and sends a notification
 * to all creators.
 *
 * Creator add a ThreadVideo to `Threads/{ThreadId}/ThreadVideos/{ShareId}`.
 * Users save their device notification tokens to `FcmTokens/{uid}`.
 */
exports.newThreadVideoNotification = functions.firestore
    .document("Threads/{ThreadId}/ThreadVideos/{ShareId}")
    .onCreate(async (snap, context) => {
      const threadId = context.params.ThreadId;
      const shareId = context.params.ShareId;
      const creatorUid = snap.data().creatorUid;
      functions.logger.log(
          "A new video: ",
          shareId,
          "from creator: ",
          creatorUid,
          " was posted to thread: ",
          threadId
      );

      // Get the thread document from firestore.
      const getThreadPromise = admin
          .firestore()
          .collection("Threads")
          .doc(threadId)
          .get();
      const getCreatorProfilePromise = admin.auth().getUser(creatorUid);

      const results = await Promise.all([
        getThreadPromise,
        getCreatorProfilePromise,
      ]);
      const thread = results[0];
      const creator = results[1];

      functions.logger.log("Fetched thread object", thread.data());
      // Check if there are any device tokens.
      if (!("creators" in thread.data())) {
        return functions.logger.log(
            "The creators field is not set in Thread object"
        );
      }
      if (creator === null) {
        return functions.logger.log("Creator profile cannot be found.");
      }
      functions.logger.log("Fetched creator profile", creator);
      // Notification details.
      const payload = {
        notification: {
          title: "New video co-created!",
          body: `${creator.displayName} just posted a new video.`,
          icon: creator.photoURL,
        },
        data: {
          threadId: threadId,
          shareId: shareId,
        },
      };
      const tokens = [];
      const tokenPromises = [];
      thread.data().creators.forEach((uid) => {
        if (uid != creatorUid) {
          const getTokenPromise = admin
              .firestore()
              .collection("FcmTokens")
              .doc(uid)
              .get();
          tokenPromises.push(getTokenPromise);
        }
      });
      const tokenResponses = await Promise.all(tokenPromises);
      tokenResponses.forEach((result) => {
        functions.logger.log("firestore returned token object", result.data());
        if (result.data().token != null) {
          tokens.push(result.data().token);
        }
      });

      if (tokens.length > 0) {
        functions.logger.log("Need to notify the following tokens: ", tokens);
        return admin.messaging().sendToDevice(tokens, payload);
      } else {
        return functions.logger.log(
            "No other creators in thread. No need to notify"
        );
      }

    /*
    // The snapshot to the user's tokens.
    let tokensSnapshot;
    // The array containing all the user's tokens.
    let tokens;
    // Listing all tokens as an array.
    tokens = Object.keys(tokensSnapshot.val());
    // Send notifications to all tokens.
    const response = await admin.messaging().sendToDevice(tokens, payload);
    // For each message check if there was an error.
    const tokensToRemove = [];
    response.results.forEach((result, index) => {
      const error = result.error;
      if (error) {
        functions.logger.error(
          'Failure sending notification to',
          tokens[index],
          error
        );
        // Cleanup the tokens who are not registered anymore.
        if (error.code === 'messaging/invalid-registration-token' ||
            error.code === 'messaging/registration-token-not-registered') {
          tokensToRemove.push(tokensSnapshot.ref.child(tokens[index]).remove());
        }
      }
    });
    return Promise.all(tokensToRemove);
*/
    });

/**
 * Triggers when a ThreadVideo is added to a Thread and updates the creator's
 * profile data.
 *
 * Creator add a ThreadVideo to `Threads/{ThreadId}/ThreadVideos/{ShareId}`.
 * Creator profile data is maintained at `CreatorProfiles/{uid}`.
 */
exports.updateCreatorProfile = functions.firestore
    .document("Threads/{ThreadId}/ThreadVideos/{ShareId}")
    .onCreate(async (snap, context) => {
      const creatorUid = snap.data().creatorUid;
      const creationTime = snap.data().creationTime;
      functions.logger.log(
          "A new video was posted by creator: ",
          creatorUid,
          " at creation time: ",
          creationTime
      );
      const getCreatorProfilePromise = admin.auth().getUser(creatorUid);

      const results = await Promise.all([getCreatorProfilePromise]);
      const creator = results[0];

      if (creator === null) {
        return functions.logger.log("Creator profile cannot be found.");
      }
      functions.logger.log("Fetched creator profile", creator);

      return admin
          .firestore()
          .collection("CreatorProfiles")
          .doc(creatorUid)
          .set(
              {
                creatorUid: creatorUid,
                updatedTime: creationTime,
                name: creator.displayName || "",
                photoUrl: creator.photoURL || "",
              },
              {merge: true}
          );
    });

/**
 * Triggers when a ThreadInvitee is added to a Thread and send notification or
 * twilio messages to the user.
 *
 * Sender add a ThreadInvitee to
 *Threads/{ThreadId}/ThreadInvitations/{InvitationId}/ThreadInvitees/{InviteeId}
 */
exports.sendInvites = functions.firestore
    .document(
    /* eslint max-len: ["error", { "ignoreStrings": true }]*/
        "Threads/{ThreadId}/ThreadInvitations/{InvitationId}/ThreadInvitees/{InviteeId}"
    )
    .onCreate(async (snap, context) => {
      const threadId = context.params.ThreadId;
      const invitationId = context.params.InvitationId;
      const inviteeId = context.params.InviteeId;

      const thread = await admin
          .firestore()
          .collection("Threads")
          .doc(threadId)
          .get();

      const creatorUid = thread.data().origUid;
      const shareId = thread.data().origShareId;

      const creatorProfile = await admin
          .firestore()
          .collection("CreatorProfiles")
          .doc(creatorUid)
          .get();

      const creatorName = creatorProfile.data().name;

      const inviteCount = await snap.ref.parent
          .get()
          .then(function(querySnapshot) {
            return querySnapshot.size;
          });
      functions.logger.log("Invitee size is: ", inviteCount.toString());

      let greetingStr = "";
      if ("name" in snap.data()) {
        greetingStr = snap.data().name + ",\n";
      }

      let optionalStr = "";
      if (inviteCount > 1) {
        optionalStr = " and " + (inviteCount - 1).toString() + " others";
      }

      const inviteStr =
      creatorName +
      " started a video thread on Bubble with you" +
      optionalStr +
      ". It will disappear in 24 hrs.";

      functions.logger.log(
          "A new thread invite: ",
          invitationId,
          "from creator: ",
          creatorUid,
          " was sent to inviteeId: ",
          inviteeId,
          " with message: ",
          inviteStr
      );

      if ("uid" in snap.data()) {
        const uid = snap.data().uid;
        functions.logger.log("Found uid for invitee: ", uid);

        const result = await admin
            .firestore()
            .collection("FcmTokens")
            .doc(uid)
            .get();
        if (result.data().token != null) {
          functions.logger.log("Send push notification to invitee: ", uid);
          const payload = {
            notification: {
              title: "New thread created!",
              body: inviteStr,
            },
            data: {
              threadId: threadId,
              shareId: shareId,
              invitationId: invitationId,
              inviteeId: inviteeId,
            },
          };
          return admin.messaging().sendToDevice(result.data().token, payload);
        }
      }
      // If push notification is not sent. Send a short message instead
      if ("phone" in snap.data()) {
        const phone = snap.data().phone;

        const invitationRef = admin
            .firestore()
            .collection("Threads")
            .doc(threadId)
            .collection("ThreadInvitations")
            .doc(invitationId);

        invitationRef
            .get()
            .then((doc) => {
              const videoUrl = doc.data().videoUrl;
              if (videoUrl != null) {
                const videoStr = "\n\nReact to the video at this link: " + videoUrl;
                admin
                    .firestore()
                    .collection("TwilioTest")
                    .doc(inviteeId)
                    .set(
                        {
                          to: phone,
                          body: greetingStr + inviteStr + videoStr,
                        },
                        {merge: true}
                    );
              } else {
                functions.logger.error(
                    "videoUrl is null in the invitation: " + invitationId
                );
              }
            })
            .catch((error) => {
              return functions.logger.error(
                  "Error reading the newly created document:",
                  error
              );
            });
        return functions.logger.log("Send sms message to invitee: ", phone);
      } else {
        return functions.logger.error(
            "Error to find invitee information: ",
            snap.data()
        );
      }
    });

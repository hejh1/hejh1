import {Request, Response} from "express";
import * as admin from "firebase-admin";
import {UserRecord} from "firebase-functions/v1/auth";

/**
 * Lists users by several identifiers (uid, email, phoneNumber, ...)
 * @param {Object} req The request object.
 * @param {Object} res The response object.
 * @return {Object} The user records array.
 */
export async function listUsers(req: Request, res: Response) {
  console.log("Listing users", req.body);

  // Checks the body is a non-empty array of identifiers
  if (!Array.isArray(req.body) || req.body.length <= 0) {
    return res.status(400).send({message: "Missing fields"});
  }

  try {
    // Gets the users by indentifiers
    const result = await _listUsersPaged(req.body);
    // Returns the user records array
    return res.status(201).send(result);
  } catch (err) {
    return res.status(500).send({message: `${err}`});
  }
}

/**
 *
 * @param  {Object} query
 * @return {Object}
 */
async function _listUsersPaged(query: any[]): Promise<UserRecord[]> {
  const chunkSize = 100; // max chunk size that the getUsers() method accept
  const result = await admin.auth().getUsers(query.slice(0, chunkSize));
  if (query.length > chunkSize) {
    return result.users.concat(await _listUsersPaged(query.slice(chunkSize)));
  }
  return result.users;
}

/**
 * List all the users sorting the result in ascending UID order
 * @param {Object} req The request object.
 * @param {Object} res The response object.
 * @return {Object} The sorted list of users as an array.
 */
export async function listAllUsers(req: Request, res: Response) {
  console.log("Listing all users");

  try {
    // List all the users with the admin SDK
    const allUsers = await _listAllUsers();

    // Sorts the users by UID
    const users = allUsers.sort((a: UserRecord, b: UserRecord) =>
      a.uid > b.uid ? 1 : a.uid < b.uid ? -1 : 0
    );

    // Returns the sorted list of users as an array
    return res.status(201).send(users);
  } catch (err) {
    return res.status(500).send({message: `${err}`});
  }
}

/**
 *
 * @param {Object} nextPageToken
 * @return {Object}
 */
async function _listAllUsers(nextPageToken?: string): Promise<UserRecord[]> {
  // List batch of users, 1000 at a time.
  const listUsersResult = await admin.auth().listUsers(1000, nextPageToken);

  // Recurs to list next batches, if any
  if (listUsersResult.pageToken) {
    return listUsersResult.users.concat(
        await _listAllUsers(listUsersResult.pageToken)
    );
  }

  return listUsersResult.users;
}

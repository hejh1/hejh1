import {listUsers} from "./user-admin";
import {isAuthenticated} from "../auth";
import {https} from "firebase-functions";
import * as express from "express";
import * as cors from "cors";

// Initializes the Express app
const appUsers = express();

// Installs the json body parser middleware
appUsers.use(express.json());

// Installs the CORS policy middleware
appUsers.use(cors({origin: true}));

// List all users: GET /users
// appUsers.get('/',  isAuthenticated, isAuthorized(['admin']), listAllUsers );

// List users by identifiers POST /users (indentifiers[])
// appUsers.post('/',  isAuthenticated, isAuthorized(['admin']), listUsers);
appUsers.post("/", isAuthenticated, listUsers);

// Export the users API
export const users = https.onRequest(appUsers);

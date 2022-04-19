const auth = require("./auth.js");
const fir = require("./fir.js");
const yargs = require("yargs");
const chalk = require("chalk");

// Command to list users from prod or test db.
yargs.command({
  command: "ls [num_users]",
  describe: "list all users",
  builder: {
    num_users: {
      describe: "number of users list",
      type: "int",
      default: 0,
    },
    model_version: {
      describe: "db model version",
      alias: "v",
      type: "string",
      default: "test",
      choices: ["test", "prod"],
    },
  },
  handler: function (argv) {
    var ref = auth.refForMemories(argv.model_version);
    fir.countSmallCollections(ref).then((size) => {
      console.log("There are ", chalk.bold.green(size), " total users.");
    });
    if (argv.num_users > 0) {
      const pRef = ref.orderBy("lasttimeUpdated", "desc").limit(argv.num_users);
      fir.printCollections(pRef);
    }
  },
});

yargs.command({
  command: "desc [userid]",
  describe: "summarize timeline for a user",
  builder: {
    userid: {
      describe: "user id",
      type: "string",
      demandOption: true,
    },
    model_version: {
      describe: "db model version",
      alias: "v",
      type: "string",
      default: "test",
      choices: ["test", "prod"],
    },
  },
  handler: function (argv) {
    const tlRef = auth.refForTimeline(argv.userid, argv.model_version);
    const thRef = auth.refForThought(argv.userid, argv.model_version);
    fir.countSmallCollections(tlRef).then((size) => {
      console.log(
        "User ",
        chalk.blue(argv.userid),
        " has ",
        chalk.green(size),
        " total thoughts."
      );
    });
    fir
      .countSmallCollectionsForCondition(tlRef, (doc) => {
        return doc.data().serverId;
      })
      .then((size) => {
        console.log(
          "User ",
          chalk.blue(argv.userid),
          " has ",
          chalk.green(size),
          " thoughts with serverId."
        );
      });
    fir
      .countSmallCollectionsForAsyncLookup(tlRef, fir.lookupBySessionId, thRef)
      .then((size) => {
        console.log(
          "User ",
          chalk.blue(argv.userid),
          " has ",
          chalk.green(size),
          " synced thoughts found in repo."
        );
      });
  },
});

yargs.command({
  command: "migrate [userid]",
  describe: "migrate timeline from test to prod for a user",
  builder: {
    userid: {
      describe: "user id",
      type: "string",
    },
  },
  handler: async function (argv) {
    if (argv.userid) {
      fir.migrateUser(argv.userid);
    } else {
      // migrate all users
      var ref = auth.refForMemories("test");
      const snapshot = await ref.get();
      snapshot.forEach((doc) => {
        // doc.data() is never undefined for query doc snapshots
        console.log("Migrating user: ", chalk.red(doc.id));
        fir.migrateUser(doc.id);
        fir.scorecard(doc.id, "prod");
      });
    }
  },
});

yargs.command({
  command: "score [userid]",
  describe: "update scorecard info for a user",
  builder: {
    userid: {
      describe: "user id",
      type: "string",
      demandOption: true,
    },
    model_version: {
      describe: "db model version",
      alias: "v",
      type: "string",
      default: "test",
      choices: ["test", "prod"],
    },
  },
  handler: function (argv) {
    fir.scorecard(argv.userid, argv.model_version);
  },
});

yargs.parse();

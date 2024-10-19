db = db.getSiblingDB(process.env.MONGO_INITDB_DATABASE);

db.createUser({
  user: process.env.DOC_USERNAME,
  pwd: process.env.DOC_PASSWORD,
  roles: [
    { role: 'readWrite', db: process.env.MONGO_INITDB_DATABASE },
    { role: 'dbAdmin', db: process.env.MONGO_INITDB_DATABASE }
  ]
});

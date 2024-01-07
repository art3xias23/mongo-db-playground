const express = require("express");
const {connectToDb, getDb} = require('./db')

// install app & middleware
const app = express();

let db
connectToDb((err) => {
  if (!err){
    app.listen("3000", () => {
      console.log("app listening on port 3000");
    });

    db = getDb()
  }
}) 


//routes

app.get("/books", (req, res) => {

  res.json({ msg: "Welcome to the api" });
});

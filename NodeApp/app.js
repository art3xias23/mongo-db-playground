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
  let books =[]

db.collection('books')
  .find()
  .sort({author:1}) // find returns a cursor that points to a collection. 
  .forEach(book => books.push(book))
  .then(() => {
    res.status(200).json(books)
  })
  .catch((err) => {
res.status(500).json({error: 'Could not fetch the documents', details: err.message})
  })
          // The cursor exposes methods like toArray and forEach which could be used on the collection

});

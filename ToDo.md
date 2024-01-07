1.Use $or
    - posts.find({$or: [{"author":"Jeremiah"}, {"rating":8}]})
2.Use $in and $nin
    - posts.find({author:{$in:["Jeremiah", "Lucas"]}})
3.Use $gt and $lt
    - posts.find({rating:{$gt:5}})
4.Use $limit()
    - posts.find({rating:{$gt:5}}).limit(1)
5.Use $sort({})
    - posts.find().sort({rating:-1})
6.Use $all
    -  posts.find({tags:{$all: ["tag1"]}})
7.Use nestedOption.key
    -  posts.find({"topic.name": "Simple Topic"})
8.Use .deleteOne
    - posts.deleteOne({_id: ObjectId('659aa9f1ba6b5bf880f8e2dd')})
9.Use .deleteMany()
    - posts.deleteMany({rating:{$gt: 1}})
10.Use .updateOne and $set
    - Adding rating to the initial post
        - posts.updateOne({"_id": ObjectId('659aa9f1ba6b5bf880f8e2dd')}659aaa04ba6b5bf880f8e2de, {$set:{"rating":8}})


On tutorial
https://www.youtube.com/watch?v=s8YG0GvQInY&list=PL4cUxeGkcC9h77dJ-QJlwGlZlTd4ecZOA&index=14&ab_channel=NetNinja

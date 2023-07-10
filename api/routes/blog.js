const express = require("express")
const { makeID } = require("../util.js")
const blog = express.Router()
blog.path = "/blog"

blog.get("/sum", async (req, res) => {
  const db = req.db.db("ngn13")
  const col = db.collection("posts")
  const results = await col.find({ priv: { $eq: false } }).toArray()

  let posts = []
  for (let i = 0; i < results.length; i++) {
    posts.push({
      title: results[i]["title"],
      desc:
        results[i]["content"]
          .substring(0, 140) // a short desc
          .replaceAll("#", "") // remove all the markdown stuff
          .replaceAll("*", "")
          .replaceAll("`", "")
          .replaceAll("-", "") + "...", // add "..." to make it look like desc
      info: `${results[i]["author"]} | ${results[i]["date"]}`
    })
  }

  // reversing so we can get
  // the latest posts on the top
  res.json({ error: 0, posts: posts.reverse() })
})

blog.get("/get", async (req, res) => {
  const id = req.query.id

  const db = req.db.db("ngn13")
  const col = db.collection("posts")
  const results = await col.find().toArray()

  for (let i = 0; i < results.length; i++) {
    // id is basically the title of the post
    // but ve remove the whitespace
    // and make it lowerspace
    // for example:
    // Online Privacy Guide -> onlineprivacyguide
    if (makeID(results[i]["title"]) === id) {
      return res.json({
        error: 0,
        post: {
          title: results[i]["title"],
          // info is the subtitle, for example:
          // ngn | 01/06/2023
          info: `${results[i]["author"]} | ${results[i]["date"]}`,
          content: results[i]["content"]
        }
      })
    }
  }

  res.json({ error: 3 })
})

blog.post("/add", async (req, res) => {
  const title = req.body.title
  const author = req.body.author
  const content = req.body.content
  const priv = req.body.priv

  if (
    typeof title !== "string" ||
    typeof author !== "string" ||
    typeof content !== "string" ||
    typeof priv !== "boolean"
  )
    return res.json({ error: 1 })

  const db = req.db.db("ngn13")
  const col = db.collection("posts")
  await col.insertOne({
    title: title,
    author: author,
    date: new Date().toLocaleDateString(),
    content: content,
    priv: priv
  })
  res.json({ error: 0 })
})

module.exports = blog

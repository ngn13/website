const express = require("express")
const { MongoClient } = require("mongodb")
const { makeID } = require("../api/util.js")
require("dotenv").config()

const app = express()
app.use(express.json())
app.use(express.urlencoded({ extended: false }))

const client = new MongoClient(process.env.DATABASE)

app.get("/:id", async (req, res) => {
  const id = req.params.id

  if (typeof id !== "string") return res.redirect("/projects")

  await client.connect()
  const db = client.db("ngn13")
  const col = db.collection("projects")
  const projects = await col.find().toArray()

  for (let i = 0; i < projects.length; i++) {
    if (makeID(projects[i]["name"]) === id) {
      res.redirect(projects[i]["url"])
      await col.updateOne(
        { name: projects[i]["name"] },
        { $set: { click: projects[i]["click"] + 1 } }
      )
    }
  }

  return res.redirect("/projects")
})

async function pexit() {
  await client.close()
  process.exit()
}

process.on("SIGTERM", pexit)
process.on("SIGINT", pexit)

export default {
  path: "/l",
  handler: app
}

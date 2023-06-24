const express = require("express")
const resources = express.Router()
resources.path = "/resources"

resources.get("/get", async (req, res) => {
  const db = req.db.db("ngn13")
  const col = db.collection("resources")
  let results = []
  if (req.query.sum) results = await col.find().limit(10).toArray()
  else results = await col.find().toArray()
  res.json({ error: 0, resources: results.reverse() })
})

resources.get("/add", async (req, res) => {
  let name = req.query.name
  let tags = req.query.tags
  let url = req.query.url

  if (
    typeof name !== "string" ||
    typeof tags !== "string" ||
    typeof url !== "string"
  )
    return res.json({ error: 1 })

  const db = req.db.db("ngn13")
  const col = db.collection("resources")
  await col.insertOne({ name: name, tags: tags.split(","), url: url })
  res.json({ error: 0 })
})

module.exports = resources

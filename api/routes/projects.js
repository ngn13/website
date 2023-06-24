const express = require("express")
const projects = express.Router()
projects.path = "/projects"

projects.get("/get", async (req, res) => {
  const db = req.db.db("ngn13")
  const col = db.collection("projects")
  const results = await col.find().toArray()
  res.json({ error: 0, projects: results })
})

projects.get("/add", async (req, res) => {
  let name = req.query.name
  let desc = req.query.desc
  let url = req.query.url

  if (
    typeof name !== "string" ||
    typeof desc !== "string" ||
    typeof url !== "string"
  )
    return res.json({ error: 1 })

  const db = req.db.db("ngn13")
  const col = db.collection("projects")
  await col.insertOne({
    name: name,
    desc: desc,
    url: url,
    click: 0
  })
  res.json({ error: 0 })
})

module.exports = projects

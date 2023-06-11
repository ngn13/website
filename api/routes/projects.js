const express = require("express")
const projects = express.Router()
projects.path = "/projects"

projects.get("/get", async (req,res)=>{
  await req.db.connect()
  const db = await req.db.db("ngn13")
  const col = await db.collection("projects")
  const results = await col.find().toArray()
  await req.db.close()
  res.json({ error: 0, projects: results })
})

projects.get("/add", async (req,res)=>{
  let name = req.query.name;
  let desc = req.query.desc;
  let url = req.query.url;

  if (!name || !desc || !url )
    return res.json({ error: 1 })

  await req.db.connect()
  const db = await req.db.db("ngn13")
  const col = await db.collection("projects")
  await col.insertOne({
    "name":name,
    "desc":desc,
    "url":url,
    "click":0
  })
  await req.db.close()
  res.json({ error: 0 })
})

module.exports = projects

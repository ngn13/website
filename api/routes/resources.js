const express = require("express")
const resources = express.Router()
resources.path = "/resources"

resources.get("/get", async (req,res)=>{
  await req.db.connect()
  const db = await req.db.db("ngn13")
  const col = await db.collection("resources")
  let results = []
  if(req.query.sum)
    results = await col.find().limit(10).toArray()
  else
    results = await col.find().toArray()
  await req.db.close()
  res.json({ error: 0, resources: results })
})

resources.get("/add", async (req,res)=>{
  let name = req.query.name;
  let tags = req.query.tags;
  let url = req.query.url;

  if(!name || !tags || !url)
    return res.json({"error":1})

  await req.db.connect()
  const db = await req.db.db("ngn13")
  const col = await db.collection("resources")
  await col.insertOne({"name":name, "tags":tags.split(","), "url":url})
  await req.db.close()
  res.json({error: 0})
})

module.exports = resources

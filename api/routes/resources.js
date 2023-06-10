const express = require("express")
const resources = express.Router()

resources.get("/get", async (req,res)=>{
  await req.db.connect()
  const col = await req.db.collection("ngn13")
  const results = col.find().limit(10).toArray()
  await req.db.close()
  res.json({ error: 0, resources: results })
})

module.exports = resources

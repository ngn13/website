const express = require("express");
const { MongoClient } = require("mongodb");
const { makeID } = require("../api/util.js")
require("dotenv").config()

const app = express();
app.use(express.json());
app.use(express.urlencoded({ extended: false }));

const client = new MongoClient(process.env.DATABASE);

app.get("/:id", async (req,res)=>{
  const id = req.params.id

  await client.connect()
  const db = await client.db("ngn13")
  const col = await db.collection("projects")
  const projects = await col.find().toArray()

  console.log(projects)

  for(let i=0; i<projects.length;i++){
    if(makeID(projects[i]["name"])===id){
      res.redirect(projects[i]["url"])
      await col.updateOne({ name: projects[i]["name"] }, { "$set":
        { "click": projects[i]["click"]+1 }})
      return await client.close()
    }
  }
  await client.close()
  return res.redirect("/projects")
})

export default {
  path: "/l",
  handler: app,
}

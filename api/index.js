const express = require("express");
const {gimmeToken} = require("./util.js")
const { MongoClient } = require("mongodb");
require("dotenv").config()

const app = express();
app.use(express.json());
app.use(express.urlencoded({ extended: false }));

const client = new MongoClient(process.env.DATABASE);
const PASS = process.env.PASS
let TOKEN = gimmeToken();

/*
 * error: 0 -> no error
 * error: 1 -> parameter error
 * error: 2 -> auth error
 * error: 3 -> not found error
*/

// PATH: /api/login
// METHOD: GET
// PARAMETERS: pass
app.get("/login", (req,res)=>{
  let pass = req.query.pass;

  if (pass === undefined)
    return res.json({error: 1})
  
  if (pass !== PASS)
    return res.json({error: 2})

  res.json({error: 0, token:TOKEN})
})

// PATH: /api/logout
// METHOD: GET
// PARAMETERS: token
app.get("/logout", (req,res)=>{
  let token = req.query.token;

  if (token === undefined)
    return res.json({error: 1})
  
  if (token !== TOKEN)
    return res.json({error: 2})

  TOKEN = gimmeToken()
  res.json({error: 0})
})

// PATH: /api/add_project
// METHOD: GET
// PARAMETERS: token, name, desc, url
app.get("/add_project", async (req, res) => {
  let token = req.query.token;
  let name = req.query.name;
  let desc = req.query.desc;
  let url = req.query.url;

  if (
    token === undefined ||
    name === undefined ||
    desc === undefined ||
    url === undefined
  )
    return res.json({error: 1})
  
  if (token !== TOKEN)
    return res.json({error: 2})
  
  await client.connect()
  const db = await client.db("ngn13")
  const col = await db.collection("projects")
  await col.insertOne({"name":name, "desc":desc, "url":url, "click":0})
  await client.close()
  res.json({error: 0})
});

// PATH: /api/add_resource
// METHOD: GET
// PARAMETERS: token, name, tags, url
app.get("/add_resource", async (req, res) => {
  let token = req.query.token;
  let name = req.query.name;
  let tags = req.query.tags;
  let url = req.query.url;

  if (
    token === undefined ||
    name === undefined ||
    tags === undefined ||
    url === undefined
  )
    return res.json({error: 1})
  
  if (token !== TOKEN)
    return res.json({error: 2})
  
  await client.connect()
  const db = await client.db("ngn13")
  const col = await db.collection("resources")
  await col.insertOne({"name":name, "tags":tags.split(","), "url":url})
  await client.close()
  res.json({error: 0})
});

// PATH: /api/get_projects
// METHOD: GET
// PARAMETERS: NONE
app.get("/get_projects", async (req, res) => {
  await client.connect()
  const db = await client.db("ngn13")
  const col = await db.collection("projects")
  const array = await col.find().toArray()
  await client.close()
  res.json({error: 0, projects:array})
});

// PATH: /api/get_resources
// METHOD: GET
// PARAMETERS: NONE
app.get("/get_resources", async (req, res) => {
  await client.connect()
  const db = await client.db("ngn13")
  const col = await db.collection("resources")
  const array = await col.find().toArray()
  await client.close()
  res.json({error: 0, resources:array})
});

// PATH: /api/add_post
// METHOD: POST
// PARAMETERS: token, title, author, content
app.post("/add_post", async (req, res) => {
  let token = req.body.token;
  let title = req.body.title;
  let author = req.body.author;
  let content = req.body.content;

  if (
    token === undefined ||
    title === undefined ||
    author === undefined ||
    content === undefined
  )
    return res.json({error: 1})
  
  if (token !== TOKEN)
    return res.json({error: 2})
  
  await client.connect()
  const db = await client.db("ngn13")
  const col = await db.collection("posts")
  await col.insertOne({
    "title":title, 
    "author":author, 
    "date": new Date().toLocaleDateString(),
    "content":content
  })
  await client.close()
  res.json({error: 0})
});

// PATH: /api/get_posts
// METHOD: POST
// PARAMETERS: NONE
app.get("/get_posts", async (req, res) => {
  await client.connect()
  const db = await client.db("ngn13")
  const col = await db.collection("posts")
  const array = await col.find().toArray()
  await client.close()

  let newarray = []
  for(let i = 0;i<array.length;i++){
    newarray.push({
      "title":array[i]["title"],
      "desc":array[i]["content"]["en"].substring(0, 140) + "...",
      "info":`${array[i]["author"]} | ${array[i]["date"]}`
    })
  }

  res.json({error: 0, posts:newarray})
});

// PATH: /api/get_post
// METHOD: POST
// PARAMETERS: id
app.get("/get_post", async (req, res) => {
  let id = req.query.id;

  await client.connect()
  const db = await client.db("ngn13")
  const col = await db.collection("posts")
  const array = await col.find().toArray()
  await client.close()

  for(let i = 0;i<array.length;i++){
    if(array[i]["title"].toLowerCase().replaceAll(" ", "")===id){
      return res.json({error: 0, post:{
        "title": array[i]["title"],
        "info": `${array[i]["author"]} | ${array[i]["date"]}`,
        "content": array[i]["content"],
      }})
    }
  }

  res.json({error: 3})
});

// PATH: /api/ping
// METHOD: GET
// PARAMETERS: NONE
app.get("/ping", (req, res) => {
  res.send({ error: 0 });
});

export default {
  path: "/api",
  handler: app,
};

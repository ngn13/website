const express = require("express");
const fs = require("fs");
const Database = require("./db");

const app = express();
const db = new Database();

app.use(express.json());
app.use(express.urlencoded({ extended: false }));

function gimmeToken() {
  var result           = ""
  var characters       = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
  var charactersLength = characters.length
  for ( var i = 0; i < 32; i++ ) {
     result += characters.charAt(Math.floor(Math.random() * charactersLength));
  }
  return result;
}

let TOKEN = gimmeToken();
const PASS = fs.readFileSync("pass", "utf-8")

/*
 * error: 0 -> no error
 * error: 1 -> parameter error
 * error: 2 -> auth error
 *  
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
app.get("/add_project", (req, res) => {
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
  
  db.push("projects", {"name":name, "desc":desc, "url":url})
  res.json({error: 0})
});

// PATH: /api/add_resource
// METHOD: GET
// PARAMETERS: token, name, tags, url
app.get("/add_resource", (req, res) => {
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
  
  db.push("resources", {"name":name, "tags":tags.split(","), "url":url})
  res.json({error: 0})
});

// PATH: /api/get_projects
// METHOD: GET
// PARAMETERS: NONE
app.get("/get_projects", (req, res) => {
  let projects = db.get("projects")
  res.json({error: 0, projects:projects})
});

// PATH: /api/get_resources
// METHOD: GET
// PARAMETERS: NONE
app.get("/get_resources", (req, res) => {
  let resources = db.get("resources")
  res.json({error: 0, resources:resources})
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

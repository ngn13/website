const express = require("express")
const { MongoClient } = require("mongodb")
require("dotenv").config()

/*
 * error: 0 -> no error
 * error: 1 -> parameter error
 * error: 2 -> auth error
 * error: 3 -> not found error
 */

const db = new MongoClient(process.env.DATABASE)
const app = express()
app.use(express.json())
app.use(express.urlencoded({ extended: false }))
app.use(async (req, res, next) => {
  await db.connect()
  req.db = db
  next()
})

const { auth, authware } = require("./routes/auth.js")
// anything starts with "add"
// requires admin privs
app.use("/*/a*", authware)
const resources = require("./routes/resources.js")
const projects = require("./routes/projects.js")
const blog = require("./routes/blog.js")
const routes = [resources, projects, blog, auth]

routes.forEach((route) => {
  app.use(route.path, route)
})

async function pexit() {
  await db.close()
  process.exit()
}

process.on("SIGTERM", pexit)
process.on("SIGINT", pexit)

export default {
  path: "/api",
  handler: app
}

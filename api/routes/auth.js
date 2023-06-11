const express = require("express")
const { gimmeToken } = require("../util.js")
const auth = express.Router()
auth.path = "/auth"

const PASS = process.env.PASS
let TOKEN = gimmeToken();

function authware(req,res,next){
  const token = req.query.token ? req.query.token : req.body.token

  if(!token)
    return res.json({ error: 1 })

  if(token!==TOKEN)
    return res.json({ error: 2 })

  next()
}
auth.use("/logout", authware)

auth.get("/login", async (req,res)=>{
  const pass = req.query.pass

  if(!pass)
    return res.json({ error: 1 })

  if(pass!==PASS)
    return res.json({ error: 2 })

  res.json({ error: 0, token: TOKEN })
})

auth.get("/logout", async (req,res)=>{
  TOKEN = gimmeToken()
  res.json({ error: 0 })
})

module.exports = { auth, authware }

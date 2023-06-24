function gimmeToken() {
  var result = ""
  var characters =
    "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
  var charactersLength = characters.length
  for (var i = 0; i < 32; i++) {
    result += characters.charAt(Math.floor(Math.random() * charactersLength))
  }
  return result
}

function makeID(title) {
  // this is used in blog.js
  // id is basically the title of the post
  // but ve remove the whitespace
  // and make it lowerspace
  // for example:
  // Online Privacy Guide -> onlineprivacyguide
  return title.toLowerCase().replaceAll(" ", "")
}

module.exports = { gimmeToken, makeID }

// https://v2.nuxt.com/deployments/pm2/
module.exports = {
  apps: [
    {
      name: "ngn13.fun website",
      exec_mode: "cluster",
      instances: "max",
      script: "./node_modules/nuxt/bin/nuxt.js",
      args: "start"
    }
  ]
}

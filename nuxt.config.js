export default {
  head: {
    title: "[ngn]",
    htmlAttrs: {
      lang: "en",
    },
    meta: [
      { charset: "utf-8" },
      { name: "viewport", content: "width=device-width, initial-scale=1" },
      { hid: "description", name: "description", content: "" },
      { name: "format-detection", content: "telephone=no" },
      { hid: "og:title", content: "[ngn]" },
      { hid: "og:description",  content: "personal website of ngn | read my blogs, check out my projects, discover cool resources" },
      { hid: "og:url",  content: "https://ngn13.fun" },
      { name: "theme-color", content: "#141414", "data-react-helmet":"true"},
    ],

    link: [
      { rel: "stylesheet", href: "https://unpkg.com/boxicons@2.1.4/css/boxicons.min.css" },
    ]
  },
  css: ["@/static/global.css"],
  plugins: [],
  components: true,
  buildModules: [],
  modules: ["@nuxtjs/axios"],
  axios: {
    baseURL: "/",
  },
  build: {},
  serverMiddleware: {
    "/api": "~/api",
    "/l": "~/links",
  },
};

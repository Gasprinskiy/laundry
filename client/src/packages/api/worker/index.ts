import { ofetch } from "ofetch";

const $api = ofetch.create({
  // baseURL: "/api",
  baseURL: "http://localhost:3000",
  headers: {
    Accept: "application/json",
  },
})

export default $api



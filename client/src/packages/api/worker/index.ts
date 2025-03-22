import { ofetch } from "ofetch";

export const $api = ofetch.create({
  baseURL: "/api",
  headers: {
    Accept: "application/json",
    "Cache-Control": "no-cache",
  },

})


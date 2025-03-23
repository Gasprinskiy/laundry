import { ofetch } from 'ofetch';

const $api = ofetch.create({
  baseURL: '/api',
  headers: {
    Accept: 'application/json',
  },
});

export default $api;

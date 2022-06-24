import Axios from "axios";

const axios = Axios.create({
  baseURL: "/api",
  timeout: 1000,
  // headers: {'X-Custom-Header': 'foobar'}
});

export default axios;

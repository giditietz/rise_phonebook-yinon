import axios from "axios";

const axiosInstance = axios.create({
  baseURL: "http://localhost:9000/api",
});

const get = (path) => {
  return axiosInstance.get(path);
};

const post = (path, data) => {
  return axiosInstance.post(path, data);
};

const httpRequest = {
  get,
  post,
};

export default httpRequest;

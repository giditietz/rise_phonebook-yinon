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
const put = (path, data) => {
  return axiosInstance.put(path, data);
};
const del = (path) => {
  return axiosInstance.delete(path);
};

const httpRequest = {
  get,
  post,
  put,
  del,
};

export default httpRequest;

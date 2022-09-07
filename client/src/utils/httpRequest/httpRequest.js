import axios from "axios";

const axiosInstance = axios.create({
  baseURL: "http://localhost:9000/api",
});

const get = (path) => {
  return axiosInstance.get(path);
};

const httpRequest = {
  get,
};

export default httpRequest;

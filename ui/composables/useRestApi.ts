import axios, {type AxiosRequestConfig } from "axios";
const instance = axios.create({
  timeout: 1000,
});
export default async function (method: string, path: string, data: object) {
  const runtimeConfig = useRuntimeConfig();
//   const token = useCookie("token");
//   if (!token.value) {
//     return navigateTo("/");
//   }
  const config: AxiosRequestConfig = {
    method: method,
    url: runtimeConfig.public.apiURL + path,
    data: data,
    // headers: {
    //   Authorization: `Bearer ${token.value}`,
    // },
  };
  const res = await instance(config)
    .then((response) => {
      return response.data;
    })
    .catch((error) => {
      return error.data;
    });
  return res;
}

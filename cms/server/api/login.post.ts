import axios from "axios";
import type { ErrorResponse, UserTokenResponse } from "~/types/response";
export default defineEventHandler(async (event) => {
  const body = await readBody(event);
  const runtimeConfig = useRuntimeConfig();
  const token = getCookie(event, 'token');
  const request = axios.create({
    baseURL: runtimeConfig.public.apiURL,
    headers: {
      "x-api-key": runtimeConfig.apiKey,
      "Content-Type": "application/json",
    },
  });
  const data = request
    .post("/account/login.json", body)
    .then((response) => {

      let userInfo: UserTokenResponse = {
        token: "xxx",
        metadata: null,
        status: false,
        code: 200,
      };
      if (response.data) {
        userInfo.token = response.data?.token;
        setCookie(event, 'token', userInfo.token)
        userInfo.metadata = response.data?.metadata;
        userInfo.status = response.data?.status
      }
      return userInfo;
    })
    .catch((error) => {
      let err: ErrorResponse = {
        status: false,
        message: error,
        code: 401,
      };
      if (error.response) {
        err.code = parseInt(error.response.status);
        err.message = error.response.data;
      }
      return err;
    });
  return data;
});

import axios from "axios";

export default defineNuxtPlugin(() => {
  const runtimeConfig = useRuntimeConfig();
  return {
    provide: {
      restApi: () => {
        return axios.create({
          baseURL: runtimeConfig.public?.apiURL,
          headers: {
            "x-api-key": "ENqKRmbFzPAbXN5VRfdq7hVNEjztVRV1",
            "Content-Type": "application/json",
          },
        });
      },
    },
  };
});

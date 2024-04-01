<template>
  <main
    class="h-screen flex flex-col items-center justify-center bg-gray-50 sm:px-4 overflow-hidden"
  >
    <div class="w-full text-gray-600 sm:max-w-md">
      <div class="text-center">
        <div class="space-y-1">
          <h3 class="text-gray-800 text-2xl font-bold sm:text-3xl">Login</h3>
        </div>
      </div>
      <div class="bg-white shadow p-4 sm:p-6 sm:rounded-lg">
        <form class="space-y-5">
          <NFormItem label="Email">
            <n-input
              v-model:value="identity.email"
              type="email"
              placeholder="Nhập email"
            />
          </NFormItem>
          <NFormItem label="Password">
            <n-input
              type="password"
              v-model:value="identity.password"
              show-password-on="mousedown"
              placeholder="Nhập mật khẩu"
              :maxlength="50"
            />
          </NFormItem>

          <n-button
            type="primary"
            size="large"
            :block="true"
            @click="handleLogin"
            attr-type="submit"
          >
            Login
          </n-button>
        </form>
      </div>
    </div>
  </main>
</template>
<script setup lang="ts">
definePageMeta({
  layout: "public",
});
import type { Identity } from "~/types/user.ts";
import { GlassesOutline, Glasses } from "@vicons/ionicons5";
import { useMessage } from "naive-ui";

const identity = ref<Identity>({ email: "", password: "" });
const message = useMessage();
const { $restApi } = useNuxtApp();
const handleLogin = async (event) => {
  event.preventDefault();

  const data = await $fetch("/api/login", {
    method: "post",
    body: identity.value,
  });

  if (data.status) {
    message.loading("Xác minh thành công! Vui lòng chờ...", { duration: 800 });
    setTimeout(() => {
      return navigateTo("/");
    }, 1000);
  } else {
    return message.error("Vui lòng kiểm tra lại thông tin!");
  }
};
</script>

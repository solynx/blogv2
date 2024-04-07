<template>
  <ClientOnly>
    <n-card title="Edit your profile" hoverable>
      <n-form>
        <n-form-item label="Thay đổi tên">
          <n-input
            type="text"
            placeholder="Nhập tên của bạn"
            :maxlength="30"
            v-model:value="userInfo.full_name"
          />
        </n-form-item>
        <n-form-item label="Thay đổi mật khẩu">
          <n-input
            type="password"
            show-password-on="mousedown"
            placeholder="Nhập mật khẩu mới"
            :maxlength="20"
            v-model:value="userInfo.password"
          />
        </n-form-item>
        <NButton type="info" @click="handleUpdateProfile">Lưu</NButton>
      </n-form>
    </n-card>
  </ClientOnly>
</template>

<script setup>
import { Method } from "~/types/requestMethod";
import { useMessage } from "naive-ui";
import { PROFILE_ENDPOINT } from "~/types/user";
const message = useMessage();
const userInfo = ref({
  full_name: "",
  password: "",
});
const token = useCookie("token");
const user = await useRestApi(Method.GET, PROFILE_ENDPOINT, {});
if (user?.status) {
  userInfo.value.full_name = user.data?.full_name || "";
} else {
  token.value = "";
  navigateTo("/login");
}
const handleUpdateProfile = async () => {
  if (!userInfo.value.full_name && !userInfo.value.password) {
    return message.warning("Vui lòng nhập thông tin");
  }
  const response = await useRestApi(
    Method.PATCH,
    PROFILE_ENDPOINT,
    userInfo.value
  );
  if (response?.status) {
    userInfo.value.password = "";
    return message.success("Thay đổi thông tin thành công");
  }
  return message.error("Thay đổi thông tin không thành công");
};
</script>

<style lang="scss" scoped></style>

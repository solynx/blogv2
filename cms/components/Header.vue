<template>
  <div style="width: 240px">
    <NuxtLink to="/">
      <img
        src="https://api.dicebear.com/8.x/identicon/svg?seed=Coco"
        alt="avatar"
        class="w-6 h-6"
      />
    </NuxtLink>
  </div>
  <n-dropdown
    :options="userOptions"
    trigger="click"
    @select="handleUserProfile"
  >
    <n-avatar
      round
      size="medium"
      src="https://api.dicebear.com/8.x/adventurer/svg?seed=Trouble"
    />
  </n-dropdown>
</template>

<script setup lang="ts">
import { NIcon, useMessage } from "naive-ui";
import {
  PersonCircleOutline as UserIcon,
  LogOutOutline as LogoutIcon,
} from "@vicons/ionicons5";
const token = useCookie("token");
const message = useMessage();
function renderIcon(icon: Component) {
  return () => h(NIcon, null, { default: () => h(icon) });
}
const userOptions = [
  {
    label: "Profile",
    key: "profile",
    icon: renderIcon(UserIcon),
  },
  {
    label: "Logout",
    key: "logout",
    icon: renderIcon(LogoutIcon),
  },
];

const handleUserProfile = (key: string) => {
  if (key === "logout") {
    return logout();
  }
  if (key === "profile") {
    return navigateTo("/profile");
  }
};
const logout = () => {
  message.loading("Vui lòng chờ...", { duration: 800 });
  token.value = null;
  setTimeout(() => {
    return navigateTo("/login");
  }, 1000);
};
</script>

<style lang="scss" scoped></style>

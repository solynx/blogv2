<script setup lang="ts">
definePageMeta({
    layout:false
})
const user = useCookie(
  'userInfo',
  {
    default: () => ({ score: -1 }),
    watch: 'shallow'
  }
)

// if (user.value && user.value !== null) {
//   user.value.score++; // userInfo cookie not update with this change
// }
</script>

<template>
  <div>User score: {{ user?.score }}</div>
</template>
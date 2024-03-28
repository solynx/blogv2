<template>
  <div class="mx-auto max-w-md mt-3">
    <strong
      class="block text-center text-xl font-bold text-gray-900 sm:text-3xl"
    >
      Can I help you today ? Please send your email for me
    </strong>

    <form class="mt-6">
      <div class="relative max-w-lg">
        <label class="sr-only" for="email"> Email </label>

        <input
          class="w-full rounded-full border-gray-200 bg-gray-100 p-4 pe-32 text-sm font-medium"
          id="email"
          type="email"
          placeholder="Email"
        />

        <button
          class="absolute end-1 top-1/2 -translate-y-1/2 rounded-full bg-blue-600 px-5 py-3 text-sm font-medium text-white transition hover:bg-blue-700"
        >
          Send
        </button>
      </div>
    </form>
  </div>
</template>

<script>
export default {
  setup() {
    return {};
  },
};
</script>

<style scoped>
input:focus {
  outline: none;
}
</style>

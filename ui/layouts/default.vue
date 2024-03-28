<template>
  <div class="wrapper w-full">
    <Header />
  <section class="mx-auto block md:flex md:flex-row w-full">
    <div class="basis-0/5 md:basis-1/5 wrapper-sidebar__category">
      <Sidebar />
    </div>
    <div class="basis-5/5 md:basis-4/5">
      <main>
        <slot />
      </main>
      <Footer/>
    </div>

  </section>

  </div>
</template>

<script>
export default {
  setup() {
    return {};
  },
};
</script>

<style lang="scss" scoped></style>

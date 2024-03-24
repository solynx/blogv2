<template>
  <Header />
  <main class="mx-auto block md:flex md:flex-row">
    <div class="basis-0/5 md:basis-1/5 wrapper-sidebar__category">
      <Sidebar />
    </div>
    <div class="basis-5/5 md:basis-4/5">
      <Content />
      <div class="mx-3 my-6">
        <div class="my-3">
          <h2 class="text-3xl font-extrabold text-[#333] inline-block">
            Category
          </h2>
        </div>
        <div class="block md:grid md:grid-cols-3 md:gap-3">
          <Category />
          <Category />
          <Category />
          <Category />
          <Category />
          <Category />
        </div>
      </div>
    </div>
  </main>
</template>

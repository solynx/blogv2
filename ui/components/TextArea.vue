<template>

   <div class="w-full mb-4 border border-gray-200 rounded-lg bg-gray-50 dark:bg-gray-700 dark:border-gray-600 mt-3">
       <div class="px-4 py-2 bg-white rounded-t-lg dark:bg-gray-800">
           <label for="comment" class="sr-only">Để lại bình luận của bạn</label>
           <textarea v-model="content" id="comment" rows="4" class="w-full px-0 text-sm text-gray-900 bg-white border-0 dark:bg-gray-800 focus:ring-0 forcus:border-0 dark:text-white dark:placeholder-gray-400" placeholder="Để lại bình luận của bạn..." required ></textarea>
       </div>
       <div class="flex items-center px-3 py-2 border-t dark:border-gray-600">
           <button type="submit" class="inline-flex items-center py-2.5 px-4 text-xs font-medium text-center text-white bg-blue-700 rounded-lg focus:ring-4 focus:ring-blue-200 dark:focus:ring-blue-900 hover:bg-blue-800" @click="handleSendClick">
               Gửi
           </button>
           <p v-if="showMessage" class="mx-3 text-sm text-bold text-emerald-500" :class="status ? 'text-emerald-500' : 'text-rose-500'">{{ status  ? "Gửi thành công" : "Gửi không thành công" }}</p>

       </div>
   </div>

</template>

<script setup lang="ts">
const props = defineProps(['content', 'message'])
const content = ref("")
const showMessage = ref(false)
const status = ref(false)
const emit = defineEmits(['handleCreateContribute'])
const handleSendClick = async () => {
   if(!content.value) return;
   await emit("handleCreateContribute", content, status, showMessage)

   setTimeout(() => {
    showMessage.value = false
   }, 2000)
}
</script>

<style scoped>
textarea {
    border: none;
    overflow: auto;
    outline: none;

    -webkit-box-shadow: none;
    -moz-box-shadow: none;
    box-shadow: none;

    resize: none; /*remove the resize handle on the bottom right*/
}


</style>
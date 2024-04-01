
export default defineNuxtRouteMiddleware((to, from) => {
    const token = useCookie('token')
    if(token.value && to?.name === 'login') {
      return navigateTo('/')
    }
    if(to?.name === 'login') return
    if(!token.value) {
      return navigateTo("/login")
    }
})
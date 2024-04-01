export default defineEventHandler(async (event) => {
  const token = getCookie(event, "token");
  setCookie(event, "token", { expires: -1 });
  return true;
});

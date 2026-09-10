import Login from "./pages/Login.svelte";
import Register from "./pages/Register.svelte";
import Chat from "./pages/Chat.svelte";

export default {
  "/": Login,
  "/login": Login,
  "/register": Register,
  "/app": Chat,
};

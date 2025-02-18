<template>
      <link href=
"https://cdn.jsdelivr.net/npm/tailwindcss@2.2.19/dist/tailwind.min.css" 
          rel="stylesheet">
    <h1 class="text-center">
        Log in
    </h1>
    <div class="mx-auto flex max-w-sm items-center gap-x-4 rounded-xl bg-white p-6 shadow-lg outline outline-black/5 dark:bg-slate-800 dark:shadow-none dark:-outline-offset-1 dark:outline-white/10">
  <img class="size-12 shrink-0" src="../image/elevator-user-man-ui-round-login-1024.webp" alt="ChitChat Logo" />
  <div class="container">
    <label for="uname"><b>Username: </b></label>
    <input type="text" placeholder="Enter Username" name="uname" required>
    <label for="psw"><b>Password: </b></label>
    <input type="password" placeholder="Enter Your Password" name="psw" required>
    <button class= "justify-center items-center bg-blue-500 hover:bg-blue-800 text-white font-bold py-2 px-4 rounded "> Login</button>


  </div>
</div>
</template>
<script>
export default {
    name: 'signup'
}
</script>
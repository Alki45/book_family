<template>
  <div class="min-h-screen flex items-center justify-center bg-gradient-to-br bg-blue-100 to-purple-700 p-6">

    <div class="bg-white dark:bg-slate-800 shadow-lg rounded-xl p-8 w-full max-w-md transform transition duration-500 hover:scale-105">
      
      <!-- Logo -->
      <div class="flex justify-center">
        <img class="w-20 h-20" src="../image/elevator-user-man-ui-round-login-1024.webp" alt="User Logo" />
      </div>
      <h2 class="text-center text-2xl font-bold text-gray-700 dark:text-white mt-4">Welcome Back</h2>
      <p class="text-center text-gray-500 dark:text-gray-300">Please log in to continue</p>

      <!-- Form -->
      <form class="mt-6">

        <div class="relative mb-4">
          <input
            type="text"
            id="username"
            placeholder="Username"
            class="w-full p-3 pl-10 border rounded-lg focus:ring-2 focus:ring-purple-400 outline-none"
          />
          <svg class="w-5 h-5 absolute left-3 top-4 text-gray-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
            <path d="M14 7a3 3 0 11-6 0 3 3 0 016 0zM2 15a6 6 0 1112 0H2z"/>
          </svg>
        </div>
        <div class="relative mb-4">
          <input
            type="password"
            id="password"
            placeholder="Password"
            class="w-full p-3 pl-10 border rounded-lg focus:ring-2 focus:ring-purple-400 outline-none"
          />
          <svg class="w-5 h-5 absolute left-3 top-4 text-gray-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
            <path fill-rule="evenodd" d="M4 7a6 6 0 1112 0v1h-2V7a4 4 0 10-8 0v1H4V7zM2 12a2 2 0 012-2h12a2 2 0 012 2v3a2 2 0 01-2 2H4a2 2 0 01-2-2v-3z" clip-rule="evenodd"/>
          </svg>
        </div>

        <!-- Remember Me & Forgot Password -->
        <div class="flex justify-between items-center mb-4 text-sm text-gray-600 dark:text-gray-300">
          <label class="flex items-center">
            <input type="checkbox" class="mr-2">
            Remember me
          </label>
          <a href="#" class="text-purple-500 hover:underline">Forgot password?</a>
        </div>

        <!-- Login Button -->
        <button class="w-full py-3 bg-purple-600 text-white font-bold rounded-lg hover:bg-purple-500 transition duration-300">
          Login
        </button>

      </form>

      <!-- Sign Up Link -->
      <p class="text-center text-sm text-gray-600 dark:text-gray-300 mt-4">
        Don't have an account? 
        <a href="#" class="text-purple-500 hover:underline">Sign Up</a>
      </p>

    </div>

  </div>
</template>

<script>
export default {
  name: 'LoginPage'
};
</script>

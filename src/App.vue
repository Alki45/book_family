
<template>
  <div class="bg-green-700 p-4 rounded-md shadow-md text-white">
      <div class="grid grid-cols-6 rounded-md items-center p-4 bg-blue-200 text-black">
    <!-- Logo Section -->
    <div class="col-span-1 relative w-80 h-10">
      <img src="./image/OIP.jpg" alt="" class="w-60 h-10 rounded-md">
      <p class="absolute inset-0 py-3 px-14 flex items-center justify-right text-2xl font-semibold text-black-700">Famy-Book</p>
    </div>


      <!-- Navigation Links -->
      <div class="col-span-5 flex justify-end space-x-7">
        <a href="" class="hover:text-blue-800  font-semibold">Home</a>
        <a href="" class="hover:text-blue-400 font-semibold">Categories</a>
        <a href="" class="hover:text-blue-400 font-semibold">About</a>
        <a href="" class="hover:text-blue-400 font-semibold">Login</a>
      </div>
    </div>

  </div>
  <Home/>
</template>
<script>
import Login from './components/Login.vue';
import Home from './components/Home.vue';

export default {
  name:'App',
  components:{
    Home
  }
}
</script>
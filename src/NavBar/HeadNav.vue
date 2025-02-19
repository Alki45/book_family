<template>
    <div class="bg-blue-800 p-4 shadow-md text-white">
      <!-- Navbar Container -->
      <div class="flex items-center justify-between bg-blue-200 text-black p-4 rounded-md">
        
        <!-- Logo Section -->
        <div class="flex items-center space-x-2">
          <img src="../image/OIP.jpg" alt="Logo" class="w-20 h-10 rounded-md">
          <p class="text-2xl font-semibold">Famy-Book</p>
        </div>
  
        <!-- Navigation Links (Hidden on small screens) -->
        <nav class="hidden md:flex space-x-7">
          <router-link to="/login" class="hover:text-blue-400 font-semibold">Login</router-link>
          <router-link to="/home" class="hover:text-blue-800 font-semibold">Home</router-link>
          <router-link to="/catagories" class="hover:text-blue-400 font-semibold"><Catagory/></router-link>
          <router-link to="/about" class="hover:text-blue-400 font-semibold">About</router-link>
        </nav>
  
        <!-- Mobile Menu Button -->
        <button @click="toggleMenu" class="md:hidden text-black focus:outline-none">
          <svg class="w-8 h-8" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16m-7 6h7"/>
          </svg>
        </button>
      </div>
  
      <!-- Mobile Dropdown Menu -->
      <div v-if="menuOpen" class="md:hidden bg-blue-700 text-white p-4 space-y-2">
        <router-link to="/login" class="block hover:text-blue-400">Login</router-link>
        <router-link to="/" class="block hover:text-blue-800">Home</router-link>
        <router-link to="/" class="block hover:text-blue-400"><Catagory/></router-link>
        <router-link to="/about" class="block hover:text-blue-400">About</router-link>
      </div>
    </div>
  </template>
  
  <script>
  import Catagory from '@/components/catagories/Catagory.vue';
  
  export default {
    name: 'HeadNav',
    components: {
      Catagory
    },
    data() {
      return {
        menuOpen: false
      };
    },
    methods: {
      toggleMenu() {
        this.menuOpen = !this.menuOpen;
      }
    }
  };
  </script>
  
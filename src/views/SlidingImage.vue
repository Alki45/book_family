<template>
  <div class="relative w-full h-80">
    <!-- Background Image -->
    <img class="absolute inset-0 w-full h-full object-cover object-center" :src="images[selected]" alt="Video - animation" />

    <!-- Search Bar (Centered Over Image) -->
    <div class="absolute inset-x-0 top-1/2 transform -translate-y-1/2 px-4">
      <div class="relative max-w-lg mx-auto">
        <div class="absolute left-3 top-1/2 transform -translate-y-1/2">
          <svg class="fill-current h-6 w-6 text-gray-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20">
            <path d="M12.9 14.32a8 8 0 1 1 1.41-1.41l5.35 5.33-1.42 1.42-5.33-5.34zM8 14A6 6 0 1 0 8 2a6 6 0 0 0 0 12z" />
          </svg>
        </div>
        <input class="w-full border pl-12 pr-4 py-2 rounded-full focus:border-blue-500 focus:shadow-outline outline-none bg-white bg-opacity-80" type="text" placeholder="Search..." />
      </div>
    </div>

    <!-- Previous Button -->
    <button @click="prevSlide" class="absolute inset-y-0 left-0 px-2 py-[25%] h-full w-8 group hover:bg-gray-500 hover:bg-opacity-75 cursor-pointer">
      <span class="hidden group-hover:block text-gray-50">&larr;</span>
    </button>

    <!-- Next Button -->
    <button @click="nextSlide" class="absolute inset-y-0 right-0 px-2 py-[25%] h-full w-8 group hover:bg-gray-500 hover:bg-opacity-75 cursor-pointer">
      <span class="hidden group-hover:block text-gray-50">&rarr;</span>
    </button>

    <!-- Dots Indicator -->
    <div class="absolute bottom-0 w-full p-4 flex justify-center space-x-2">
      <button
        v-for="(image, index) in images"
        :key="index"
        @click="selected = index"
        :class="{'bg-gray-300': selected === index, 'bg-gray-500': selected !== index}"
        class="h-2 w-2 rounded-full hover:bg-gray-300 ring-2 ring-gray-300"
      ></button>
    </div>
  </div>
</template>

<script>
import giphy from '@/assets/3i8T6n.gif';
import giphy2 from '@/assets/giphy (1).gif';
import giph3 from '@/assets/giphy.gif';
import giphy4 from '@/assets/demorcracyrising-divide.gif';

export default {
  name: 'SlidingImage',
  data() {
    return {
      selected: 0,
      images: [giphy, giphy2, giph3, giphy4] // Make sure to include all your images
    };
  },
  mounted() {
    console.log(this.images[this.selected]); // Check the path in the console
  },
  methods: {
    prevSlide() {
      this.selected = this.selected > 0 ? this.selected - 1 : this.images.length - 1;
    },
    nextSlide() {
      this.selected = this.selected < this.images.length - 1 ? this.selected + 1 : 0;
    }
  }
};
</script>

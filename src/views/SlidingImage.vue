<template>
<div>
<div class="relative justify-items-center">
        <img class="h-80 w-full object-cover object-center justify-items-center" :src="images[selected]" alt="Vidoe - animation" />
        <!-- Previous Button -->
        <button @click="prevSlide" class="absolute inset-y-0 left-0 px-2 py-[25%] h-full w-8 group hover:bg-gray-500 hover:bg-opacity-75 cursor-pointer">
          <span class="hidden group-hover:block text-gray-50">
            &larr;
          </span>
        </button>
  
        <!-- Next Button -->
        <button @click="nextSlide" class="absolute inset-y-0 right-0 px-2 py-[25%] h-full w-8 group hover:bg-gray-500 hover:bg-opacity-75 cursor-pointer">
          <span class="hidden group-hover:block text-gray-50">
            &rarr;
          </span>
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
    </div>

</template>
<script>
  import giphy from '@/assets/3i8T6n.gif';
  import giphy2 from '@/assets/giphy (1).gif';
  import giph3 from '@/assets/giphy.gif';
  import giphy4 from '@/assets/demorcracyrising-divide.gif';
  export default {
    name:'SlidingImage',
    data (){
        return{
        selected: 0,
        images:[
        giphy,
        giphy2,
        giph3,
        giphy
        ]
        }
    },
        mounted() {
        console.log(this.images[this.selected]);  // Check the path in the console
        }
,
    methods: {
      
      prevSlide() {
        this.selected = this.selected > 0 ? this.selected - 1 : this.images.length - 1;
      },
      nextSlide() {
        this.selected = this.selected < this.images.length - 1 ? this.selected + 1 : 0;
      }
    }


}

</script>
<template>
  <div class="relative flex justify-center items-center px-4 sm:px-8 md:px-12 lg:px-20">
    <!-- Left Arrow -->
    <button 
      @click="prevProducts" 
      class="absolute left-0 top-1/2 transform -translate-y-1/2 bg-gray-800 text-white p-3 rounded-full hover:bg-gray-600 transition"
      aria-label="Previous"
    >
      &#8592;
    </button>

    <!-- Product Cards -->
    <div class="flex overflow-hidden w-full py-10">
      <TransitionGroup tag="div" name="slide" class="flex items-center space-x-4">
        <div 
          v-for="product in displayedProducts" 
          :key="product.id" 
          class="bg-slate-500 w-full max-w-[calc(100%/productsPerRow-1rem)] h-96 shadow-lg rounded-lg transform transition-transform duration-300 ease-in-out hover:scale-105 hover:shadow-xl"
        >
          <div class="h-4/6 w-full">
            <img class="w-full h-full object-cover rounded-t" :src="product.image" :alt="product.name">
          </div>
          <div class="w-full h-1/4 p-3 text-center">
            <a href="#" class="hover:text-blue-800 text-cyan-500">
              <p class="text-lg font-sans font-semibold uppercase tracking-wide">
                {{ product.name }}
              </p>
            </a>
            <p class="text-blue text-sm leading-5 mt-1 font-semibold font-serif">
              {{ product.description }}
            </p>
          </div>
        </div>
      </TransitionGroup>
    </div>

    <!-- Right Arrow -->
    <button 
      @click="nextProducts" 
      class="absolute right-0 top-1/2 transform -translate-y-1/2 bg-gray-800 text-white p-3 rounded-full hover:bg-gray-600 transition"
      aria-label="Next"
    >
      &#8594;
    </button>
  </div>
</template>

<script setup>
import { useProductStore } from '/src/stores/productstores.js'; 
import { onMounted, ref, watch, onUnmounted } from 'vue';

const productStore = useProductStore();
const displayedProducts = ref([]);
const currentIndex = ref(0);
const productsPerRow = ref(6);  // Dynamically set based on screen width

const updateProductsPerRow = () => {
  const width = window.innerWidth / window.devicePixelRatio;  // Adjust for zoom
  if (width < 640) {
    productsPerRow.value = 1;
  } else if (width >= 640 && width < 1024) {
    productsPerRow.value = 2;
  } else if (width >= 1024 && width < 1440) {
    productsPerRow.value = 3;
  } else {
    productsPerRow.value = 6;
  }
  updateDisplayedProducts();
};

const updateDisplayedProducts = () => {
  displayedProducts.value = productStore.products.slice(currentIndex.value, currentIndex.value + productsPerRow.value);
};

// Fetch products & set initial values
onMounted(async () => {
  await productStore.fetchProducts();  
  updateProductsPerRow();
  window.addEventListener("resize", updateProductsPerRow);
});

// Clean up event listeners
onUnmounted(() => {
  window.removeEventListener("resize", updateProductsPerRow);
});

// Slide next product in
const nextProducts = () => {
  if (currentIndex.value + productsPerRow.value < productStore.products.length) {
    currentIndex.value += 1;
    updateDisplayedProducts();
  }
};

// Slide previous product in
const prevProducts = () => {
  if (currentIndex.value > 0) {
    currentIndex.value -= 1;
    updateDisplayedProducts();
  }
};
</script>

<style>
/* Slide Animation */
.slide-enter-active, .slide-leave-active {
  transition: transform 0.3s ease-in-out, opacity 0.3s;
}
.slide-enter-from {
  transform: translateX(100%);
  opacity: 0;
}
.slide-leave-to {
  transform: translateX(-100%);
  opacity: 0;
}
</style>

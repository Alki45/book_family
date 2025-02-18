<template>
    <div class="p-4 font-serif bg-gray-50">
      <h1 class="text-2xl font-semibold">
        Wel-Come To Famy-Book Community
      </h1>
      <div class="relative justify-items-center">
        <img class="h-64 w-100 object-cover object-center justify-items-center" :src="images[selected]" alt="mountains" />
        
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


      <!--Best Sold Book-->
  <section class="Recent-book-uploaded">
    <h1>Most Reviewed Book</h1>
      <div class="most_reviewed">
        <div class="grid grid-cols-1 md:grid-cols-4 gap-6 flex items-center justify-center">
          <div class="relative bg-white py-7 px-6 rounded-3xl w-80 h-80 my-2 shadow-md">
            <img src="../image/book-svgrepo-com.svg" alt="" class="w-70 h-40">
            <p>Name: </p>
            <p>author: </p>
            <p>rate: </p>
          </div>
          <div class="relative bg-white py-7 px-6 rounded-3xl w-80 my-2 shadow-md">
            <img src="../image/book-svgrepo-com.svg" alt="" class="w-70 h-40">
            <p>Name: </p>
            <p>author: </p>
            <p>rate: </p>
          </div>


          <div class="relative bg-white py-7 px-6 rounded-3xl w-80 my-2 shadow-md">
            <img src="../image/book-svgrepo-com.svg" alt="" class="w-70 h-40">

            <div>
            <p>Name: </p>
            <p>author: </p>
            <p>rate: </p>
            </div>
            </div>
          </div>
      </div>
    </section>
 <!--Best Sold Book-->


    <!-- Recent Book Uploaded-->
      <div class="Recent-book-uploaded">
        <h1 class="font-semibold font-serif">Recently Added Book</h1>
        <div class="grid grid-cols-1 md:grid-cols-4 gap-6 flex items-center justify-center">
          <div class="relative bg-white py-7 px-6 rounded-3xl w-80 h-80 my-2 shadow-md">
            <img src="../image/book-svgrepo-com.svg" alt="" class="w-70 h-40">
            <p>Name: </p>
            <p>author: </p>
            <p>rate: </p>
          </div>
          <div class="relative bg-white py-7 px-6 rounded-3xl w-80 h-80 my-2 shadow-md">
            <img src="../image/book-svgrepo-com.svg" alt="" class="w-70 h-40">
            <p>Name: </p>
            <p>author: </p>
            <p>rate: </p>
          </div>


          <div class="relative bg-white py-7 px-6 rounded-3xl w-80 h-80 my-2 shadow-md">
            <img src="../image/book-svgrepo-com.svg" alt="" class="w-70 h-40">

            <div>
            <p>Name: </p>
            <p>author: </p>
            <p>rate: </p>
            </div>

          
          </div>
          
        </div>
      </div>
  <!--End of Recent Book Uploaded-->





  <!-- Card Catagories   grid grid-cols-1 md:grid-cols-4 gap-6 flex items-center justify-center-->
    <div class=" card-Catagories py-20">
      <h1 class=" font-semibold font-serif">Catagories</h1>
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6 flex items-center justify-center py-2">
        
    <div class="grid grid-cols-1 gap-6 sm:grid-cols-2 md:grid-cols-2 lg:grid-cols-4 xl:grid-cols-4">
        <!-- 1 card -->
        <div class="relative bg-white py-6 px-6 rounded-3xl w-80 my-4 shadow-xl">
            <div class=" text-white flex items-center absolute rounded-full py-4 px-4 shadow-xl bg-pink-500 left-20 -top-5">
                <!-- svg  -->
                 <p>فقه</p>
                 <a href=""><img src="../image/285636_book_icon.svg" >
                 </a>
            </div>
            <div class="mt-8">
                <div class="flex space-x-2 text-black border-l-teal-800 text-sm my-3">
                    <!-- svg  -->
                    <svg xmlns="../image/285636_book_icon.svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
                    </svg>
                     <p>1000 كتب</p> 
                </div>
                <div class="border-t-2"></div>

                <div class="flex justify-between">
                    <div class="my-2">
                        <p class="font-semibold text-base mb-2">مذهب</p>
                        <div class="flex space-x-3">
                          <div>
                            <a href=""><img src="../image/285636_book_icon.svg" 
                            class="w-50 h-10"/>
                            <div>
                              <p>Hanefy</p>
                            </div>
                          </a>
                          </div>
                          <div>
                            <a href="">
                             <img src="../image/3069185_book_education_learn_school_icon.svg" 
                            class="w-50 h-10"/>

                            <div>
                              <p>Maliky</p>
                            </div>
                          </a>
                          </div>
                          <div>
                            <a href="">
                             <img src="../image/6428000_book_education_school_student_icon.svg" 
                            class="w-50 h-10"/>
                            <div>
                              <p>Shafy</p>
                            </div>
                          </a>
                          </div>
                          <div>
                            <a href="">
                            <img src="../image/6428005_book_education_school_student_icon.svg" 
                            class="w-50 h-10"/>
                            <div>
                              <p>Hanbely</p>
                            </div>
                          </a>
                          </div>
                      <div class="my-2">
                        <p class="font-semibold text-base mb-1">Daily View</p>
                        <div class="text-base text-gray-400 font-semibold">
                            <p>34%</p>
                        </div>
                    </div>

                        </div>
                    </div>

                </div>
            </div>
        </div>
      </div>
</div>



    </div>
 <!--End of Card Catagories-->





    <!--About as -->
<div class="About-us">
  <div class="grid grid-cols-1 md:grid-cols-2 gap-6 items-center my-4">
    <!-- Community Platforms -->
    <div class="bg-white rounded-3xl shadow-md p-6">
      <h1 class="text-xl font-bold text-indigo-600 mb-4">Our Community Platforms</h1>
      <ul class="list-disc ml-4">
        <li>Telegram</li>
        <li>WhatsApp</li>
        <li>Discord</li>
      </ul>
    </div>

    <!-- Contact Information -->
    <div class="bg-white rounded-3xl shadow-md p-6">
      <h1 class="text-xl font-bold text-indigo-600 mb-4">Contact Us</h1>
      <ul class="list-disc ml-4 font-semibold">
        <li>Email: contact@famybook.com</li>
        <li>Phone Number: +25(19) 2559-4543</li>
        <li>Telegram: @famybook</li>
      </ul>
      <textarea class="w-full mt-4 p-2 border border-gray-300 rounded-md" placeholder="Your feedback"></textarea>
      <button class="mt-2 bg-indigo-600 text-white px-4 py-2 rounded-md hover:bg-indigo-700">Send</button>
    </div>
  </div>
</div>
<!-- End of About Us -->



  </template>
  
  <script>
  export default {
    name: 'Home',
    data() {
      return {
        selected: 0,
        images: [
        "../image/6428000_book_education_school_student_icon.svg",
        "../image/OIP (1).svg",
        "../image/OIP.svg"
        ]
      };
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
  
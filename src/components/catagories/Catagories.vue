<template>

    
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
                 <a href=""><img src="/src/image/285636_book_icon.svg" >
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
                            <a href=""><img src="/src/image/285636_book_icon.svg" 
                            class="w-50 h-10"/>
                            <div>
                              <p>Hanefy</p>
                            </div>
                          </a>
                          </div>
                          <div>
                            <a href="">
                             <img src="/src/image/3069185_book_education_learn_school_icon.svg" 
                            class="w-50 h-10"/>

                            <div>
                              <p>Maliky</p>
                            </div>
                          </a>
                          </div>
                          <div>
                            <a href="">
                             <img src="/src/image/6428000_book_education_school_student_icon.svg" 
                            class="w-50 h-10"/>
                            <div>
                              <p>Shafy</p>
                            </div>
                          </a>
                          </div>
                          <div>
                            <a href="">
                            <img src="/src/image/6428005_book_education_school_student_icon.svg" 
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
</template>
<script>

export default {
        name: 'Catagories',
}

</script>
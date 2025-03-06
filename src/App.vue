
<template>
  <HeadNav /> 
  <router-view /> 
</template>
<script>

import HeadNav from './NavBar/HeadNav.vue';
import Home from './components/Home.vue';
import BookList from './views/BookList.vue'




export default {
  name:'App',
  components:{
    HeadNav,
    Home,
    BookList

    
  }
}
</script>
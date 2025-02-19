
<template>
  <HeadNav />  <!-- Navigation Bar -->
  <router-view />  <!-- This will dynamically render Home or Login -->
  <Home v-if="$route.path !='/login'"/>
</template>
<script>

import HeadNav from './NavBar/HeadNav.vue';
import Home from './components/Home.vue';





export default {
  name:'App',
  components:{

    HeadNav,
    Home
    
  }
}
</script>
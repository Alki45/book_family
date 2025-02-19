<template>

<div class="mb-4 px-2 w-full">

  <div class="relative">
    <select class="w-30 h-8 border px-2 pr-6 rounded focus:border-blue-500 focus:shadow-outline outline-none appearance-none" id="select-item">
    <option>Catagories</option>
      <option>Fiction</option>
      <option>Non-Fictions</option>
      <option>Scientific</option>
      <option>Academic</option>
      <option>Religious</option>
      <option>Business</option>

    </select>

    <div class="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700">
      <svg class="fill-current h-4 w-4" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20">
        <path d="M9.293 12.95l.707.707L15.657 8l-1.414-1.414L10 10.828 5.757 6.586 4.343 8z" />
      </svg>
    </div>
  </div>
</div>

</template>

<script>
export default{
    name:'Catagory'
}
</script>
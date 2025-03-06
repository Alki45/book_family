<template>
  <div class="container mx-auto p-4">
    <h1 class="text-3xl font-bold mb-4">Books List</h1>
    <table class="min-w-full table-auto border-collapse border border-gray-300">
      <thead>
        <tr class="bg-gray-100">
          <th class="px-4 py-2 text-left border-b">Name</th>
          <th class="px-4 py-2 text-left border-b">Author</th>
          <th class="px-4 py-2 text-left border-b">Publication</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="book in books" :key="book.ID" class="hover:bg-gray-50">
          <td class="px-4 py-2 border-b">{{ book.name }}</td>
          <td class="px-4 py-2 border-b">{{ book.author }}</td>
          <td class="px-4 py-2 border-b">{{ book.publication || 'N/A' }}</td> <!-- Handles null values -->
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      books: [],
    };
  },
  mounted() {
    this.fetchBooks();
  },
  methods: {
    async fetchBooks() {
      try {
        const response = await axios.get('http://localhost:9010/books/');
        console.log(response.data);
        this.books = response.data.map(book => ({
          ID: book.ID,
          name: book.name,
          author: book.author,
          publication: book.publication !== 'Null_val' ? book.publication : null // Handle 'Null_val' as null
        }));
      } catch (error) {
        console.error('Error fetching books:', error);
      }
    },
  },
};
</script>


